package arraypool

import "sync"

var defaultBufferCapacity int = 128

type ArrayBuffer[T any] struct {
	Buffer []T
}

type ArrayPool[T any] struct {
	pool sync.Pool
}

// create a new instance of pool
func NewPool[T any]() *ArrayPool[T] {
	ap := &ArrayPool[T]{}
	ap.pool = sync.Pool{
		New: func() any {
			return &ArrayBuffer[T]{
				Buffer: make([]T, 0, defaultBufferCapacity),
			}
		},
	}
	return ap
}

// get an existing instance from the pool. If it does not exist sync.pool creates one
func (ap *ArrayPool[T]) Get() *ArrayBuffer[T] {
	return ap.pool.Get().(*ArrayBuffer[T])
}

// put the buffer back in the pool after use, an additional flag to resize the buffer is required
func (ap *ArrayPool[T]) Put(ab *ArrayBuffer[T], resize bool) {
	if resize {
		ab.Resize()
	}
	ap.pool.Put(ab)
}

// write to the Buffer
func (ab *ArrayBuffer[T]) Write(val T) {
	ab.Buffer = append(ab.Buffer, val)
}

// clear the values of the buffer
func (ab *ArrayBuffer[T]) ClearAll() {
	clear(ab.Buffer)
}

// resize the buffer back to zero, maintains the capacity
func (ab *ArrayBuffer[T]) Resize() {
	ab.Buffer = ab.Buffer[:0]
}
