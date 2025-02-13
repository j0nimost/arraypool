package arraypool

import (
	"testing"
)

func TestArrayPoolBasic(t *testing.T) {
	ap := NewPool[int]()
	apbuffer := ap.Get()
	if apbuffer == nil && len(apbuffer.Buffer) == defaultBufferCapacity {
		t.Errorf("Result; ArrayPoolBuffer is nil or length does not match defaultCapacity")
	}

	for i := 0; i < 10; i++ {
		apbuffer.Write(i)
	}
	if len(apbuffer.Buffer) != 10 {
		t.Errorf("Result; Expected Buffer Length %d actually got %d", 10, len(apbuffer.Buffer))
	}

	apbuffer.ClearAll()

	if len(apbuffer.Buffer) != 10 {
		t.Errorf("Result; Expected Buffer Length %d actually got %d", 10, len(apbuffer.Buffer))
	}

	ap.Put(apbuffer, true)

	if len(apbuffer.Buffer) != 0 {
		t.Errorf("Result; Expected Buffer Resize to 0 got %d", len(apbuffer.Buffer))
	}
	t.Logf("Len: %d, Cap: %d", len(apbuffer.Buffer), cap(apbuffer.Buffer))
}

func TestArrayPoolGet(t *testing.T) {
	ap := NewPool[int]()
	apbuffer := ap.Get()
	if apbuffer == nil {
		t.Error("Result; ArrayPoolBuffer is nil")
	}

	apbuffer.Write(4)
	ap.Put(apbuffer, false)
	apbuffer2 := ap.Get()

	if len(apbuffer.Buffer) != len(apbuffer2.Buffer) {
		t.Errorf("Result; Mismatch Lengths Expected Buffers of Equal Size")
	}

	if apbuffer.Buffer[0] != apbuffer2.Buffer[0] {
		t.Errorf("Result; Mismatch Values %d vs %d", apbuffer.Buffer[0], apbuffer2.Buffer[0])
	}
}

// benchmark against regular buffers
var ap *ArrayPool[int] = NewPool[int]()

func ArrayPoolBenchTest() {
	apBuffer := ap.Get()
	for i := 0; i < defaultBufferCapacity+2; i++ {
		apBuffer.Write(i)
	}
	ap.Put(apBuffer, true)
}

func SliceBenchTest() {
	sl := make([]int, 0, defaultBufferCapacity)

	for i := 0; i < defaultBufferCapacity+2; i++ {
		sl = append(sl, i)
	}
}

func BenchmarkArrayPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ArrayPoolBenchTest()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceBenchTest()
	}
}
