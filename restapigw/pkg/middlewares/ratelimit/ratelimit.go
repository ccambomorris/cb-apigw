// Package ratelimit - Token Bucket 기반의 Rate Limit 처리를 지원하는 패키지
package ratelimit

import (
	"context"

	"github.com/cloud-barista/cb-apigw/restapigw/pkg/core"
	"github.com/cloud-barista/cb-apigw/restapigw/pkg/errors"
	"github.com/cloud-barista/cb-apigw/restapigw/pkg/middlewares/ratelimit/backend"
	"github.com/cloud-barista/cb-apigw/restapigw/pkg/middlewares/ratelimit/limiter"
)

// ===== [ Constants and Variables ] =====

var (
	// ErrRouteLimited - Endpoint 기준 Rate Limit 제한인 경우의 오류
	ErrRouteLimited = errors.New("ERROR: Endpoint rate limit exceeded")
	// ErrClientLimited - Client 식별 기준 Rate Limit 제한인 경우의 오류
	ErrClientLimited = errors.New("ERROR: Endpoint(By Client) rate limit exceeded")
	// ErrProxyLimited - Rate Limit 제한인 경우의 오류
	ErrProxyLimited = core.NewWrappedError(503, "Proxy(Backend) rate limit exceeded", errors.New("ERROR: Proxy(Backend) rate limit exceeded"))
)

// ===== [ Types ] =====

// ILimiter - Rate limit 운영을 위한 인터페이스
type ILimiter interface {
	// Rate limit 초과 여부 검증
	Allow() bool
}

// LimiterStore - 지정한 키에 해당하는 Limiter 정보를 검증하는 함수 구조
type LimiterStore func(string) ILimiter

// RateLimiter - Rate limit 운영을 위한 Bucket Wrapper 구조
type RateLimiter struct {
	limiter *limiter.TokenBucket
}

// ===== [ Implementations ] =====

// Allow - Rate Limit 처리를 위해 Bucket에서 Token 사용이 가능한지를 검증하고, 1개의 Token을 사용한다.
func (rl RateLimiter) Allow() bool {
	return rl.limiter.TryAquire()
}

// ===== [ Private Functions ] =====
// ===== [ Public Functions ] =====

// NewLimitterStore - Rate limit 정보 저장을 위한 Limiter Store 생성
func NewLimitterStore(maxRate int, fillInterval int, fillCount int, backend backend.IBackend) LimiterStore {
	f := func() interface{} {
		return NewLimiterWithRate(maxRate, fillInterval, fillCount)
	}
	return func(t string) ILimiter {
		return backend.Load(t, f).(RateLimiter)
	}
}

// NewMemoryStore - Memory를 기반으로 하는 LimiterStore 생성
func NewMemoryStore(maxRate int, fillInterval int, fillCount int) LimiterStore {
	return NewLimitterStore(maxRate, fillInterval, fillCount, backend.DefaultShardedMemoryBackend(context.Background()))
}

// NewLimiterWithRate - 지정한 허용수를 기준으로 Rate Limiter 생성
func NewLimiterWithRate(maxRate int, fillInterval int, fillCount int) RateLimiter {
	return RateLimiter{
		limiter.NewTokenBucketWithFill(maxRate, fillInterval, fillCount),
	}
}
