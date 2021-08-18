// Package limiter - Rate Limit 처리용 TokenBucket 구현 패키지
package limiter

import (
	"sync"
	"time"

	"github.com/cloud-barista/cb-apigw/restapigw/pkg/core/atomic"
)

// ===== [ Constants and Variables ] =====
const (
	defaultInteraval  int = 100 // Fill interval 기본 값 (Millisencond)
	defaultFillTokens int = 2   // Fill token count 기본 값
)

var ()

// ===== [ Types ] =====
type (
	// TokenBucket - Ratelimit 처리를 위한 Token Bucket 형식
	TokenBucket struct {
		capacity       int
		fillInterval   time.Duration
		fillTokenCount int
		disposeFlag    atomic.Bool
		mu             sync.Mutex
		available      int
		cond           *sync.Cond
	}
)

// ===== [ Implementations ] =====

// asyncFillToken - FillInternval로 지정한 시간 간격마다 FillTokenCount로 지정된 수의 토큰을 추가하는 비동기 처리
func (tb *TokenBucket) asyncFillToken() {
	go func() {
		t := time.NewTicker(tb.fillInterval)
		defer t.Stop()
		for {
			if tb.disposeFlag.Load() {
				break
			}
			<-t.C
			tb.mu.Lock()
			if tb.available < tb.capacity {
				tb.available += tb.fillTokenCount
				if tb.available > tb.capacity {
					tb.available = tb.capacity
				}

				// 호출 중에 Lock 유지하는 것은 허용되지만 필수는 아님
				tb.cond.Broadcast()
			}
			tb.mu.Unlock()

			// select {
			// case <-t.C:
			// 	tb.mu.Lock()
			// 	if tb.available < tb.capacity {
			// 		tb.available += tb.fillTokenCount
			// 		if tb.available > tb.capacity {
			// 			tb.available = tb.capacity
			// 		}

			// 		// 호출 중에 Lock 유지하는 것은 허용되지만 필수는 아님
			// 		tb.cond.Broadcast()
			// 	}
			// 	tb.mu.Unlock()
			// }
		}
	}()
}

// checkAquireNum - 지정한 요청 Token 수가 최대치 이상인지 검증
// @param num: 사용할 Token 요청 수
func (tb *TokenBucket) checkAquireNum(num int) bool {
	return num <= tb.capacity
}

// TryAquire - 사용할 Token획득 시도
func (tb *TokenBucket) TryAquire() bool {
	return tb.TryAquireWithNum(1)
}

// WaitUntilAquire - 사용할 Token이 채워질 때까지 대기
func (tb *TokenBucket) WaitUntilAquire() bool {
	return tb.WaitUntilAquireWithNum(1)
}

// TryAquireWithNum - 지정한 수만큼의 Token 획득 시도
// @param num: 사용할 Token 요청 수
func (tb *TokenBucket) TryAquireWithNum(num int) bool {
	// 활용 가능한 Token 수 또는 최대 수를 넘어가면 불가처리
	if !tb.checkAquireNum(num) {
		return false
	}

	tb.mu.Lock()
	defer tb.mu.Unlock()
	if tb.available >= num {
		tb.available -= num
		return true
	}

	return false
}

// WaitUntilAquireWithNum - 지정한 수만큼의 Token이 채워질때까지 대기
// @param num: 사용할 Token 요청 수
func (tb *TokenBucket) WaitUntilAquireWithNum(num int) bool {
	// 활용 가능한 Token 수 또는 최대 수를 넘어가면 불가처리
	if !tb.checkAquireNum(num) {
		return false
	}

	for {
		tb.mu.Lock()
		if tb.available >= num {
			tb.available -= num
			tb.mu.Unlock()
			return true
		}

		tb.cond.Wait()
		tb.mu.Unlock()
	}
}

// Dispose - TokenBucket을 채우는 처리 중지
func (tb *TokenBucket) Dispose() {
	tb.disposeFlag.Store(true)
}

// ===== [ Private Functions ] =====
// ===== [ Public Functions ] =====

// NewTokenBucket - 지정한 허용 수 기준으로 TockenBucket 생성 (채워지는 기간과 수량은 기본값 사용)
// @param capacity: 초당 허용 수
func NewTokenBucket(capacity int) *TokenBucket {
	return NewTokenBucketWithFill(capacity, defaultInteraval, defaultFillTokens)
}

// NewTokenBucket - 지정한 정보에 따라서 TokenBucket 생성
// @param capacity: 초당 허용 수
// @param fillInterval: 토큰을 채울 기간 (millisecond)
// @param fillTokens: 채울 토큰 수
func NewTokenBucketWithFill(capacity int, fillInterval int, fillTokens int) *TokenBucket {
	if fillInterval == 0 {
		fillInterval = defaultInteraval
	}
	if fillTokens == 0 {
		fillTokens = defaultFillTokens
	}

	tb := &TokenBucket{
		capacity:       capacity,
		available:      capacity,
		fillInterval:   time.Duration(fillInterval) * time.Millisecond,
		fillTokenCount: fillTokens,
	}

	tb.cond = sync.NewCond(&tb.mu)
	tb.asyncFillToken()
	return tb
}
