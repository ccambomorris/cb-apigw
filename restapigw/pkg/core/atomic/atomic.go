// atomic - 각 데이터 형식에 대한 Atomic 제공 패키지
package atomic

import "sync/atomic"

// ===== [ Constants and Variables ] =====
const ()

var ()

// ===== [ Types ] =====

type (
	Bool struct {
		core Int32
	}
	Int32 struct {
		core int32
	}
	Uint32 struct {
		core uint32
	}
)

// ===== [ Implementations ] =====

func (obj *Int32) Load() int32 {
	return atomic.LoadInt32(&obj.core)
}

func (obj *Int32) Store(val int32) {
	atomic.StoreInt32(&obj.core, val)
}

func (obj *Int32) Add(delta int32) (new int32) {
	return atomic.AddInt32(&obj.core, delta)
}

func (obj *Int32) Sub(delta int32) (new int32) {
	return atomic.AddInt32(&obj.core, -delta)
}

func (obj *Int32) Increment() (new int32) {
	return atomic.AddInt32(&obj.core, 1)
}

func (obj *Int32) Decrement() (new int32) {
	return atomic.AddInt32(&obj.core, -1)
}

func (obj *Int32) CompareAndSwap(old int32, new int32) (swapped bool) {
	return atomic.CompareAndSwapInt32(&obj.core, old, new)
}

func (obj *Int32) Swap(new int32) (old int32) {
	return atomic.SwapInt32(&obj.core, new)
}

func (obj *Uint32) Load() uint32 {
	return atomic.LoadUint32(&obj.core)
}

func (obj *Uint32) Store(val uint32) {
	atomic.StoreUint32(&obj.core, val)
}

func (obj *Uint32) Add(delta uint32) (new uint32) {
	return atomic.AddUint32(&obj.core, delta)
}

func (obj *Uint32) Sub(delta uint32) (new uint32) {
	return atomic.AddUint32(&obj.core, ^uint32(delta-1))
}

func (obj *Uint32) Increment() (new uint32) {
	return atomic.AddUint32(&obj.core, 1)
}

func (obj *Uint32) Decrement() (new uint32) {
	return atomic.AddUint32(&obj.core, ^uint32(0))
}

func (obj *Uint32) CompareAndSwap(old uint32, new uint32) (swapped bool) {
	return atomic.CompareAndSwapUint32(&obj.core, old, new)
}

func (obj *Uint32) Swap(new uint32) (old uint32) {
	return atomic.SwapUint32(&obj.core, new)
}

func (obj *Bool) Load() bool {
	return int32tobool(obj.core.Load())
}

func (obj *Bool) Store(val bool) {
	obj.core.Store(booltoint32(val))
}

func (obj *Bool) CompareAndSwap(old bool, new bool) (swapped bool) {
	return obj.core.CompareAndSwap(booltoint32(old), booltoint32(new))
}

func (obj *Bool) Swap(new bool) (old bool) {
	return int32tobool(obj.core.Swap(booltoint32(new)))
}

// ===== [ Private Functions ] =====

func booltoint32(val bool) int32 {
	if val {
		return 1
	}
	return 0
}

func int32tobool(val int32) bool {
	return val == 1
}

// ===== [ Public Functions ] =====
