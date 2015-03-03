package pool

import (
	"math/rand"
	"testing"
)

func TestChanPoolConsist(t *testing.T) {

	isPool := NewSlicePool(
		func(l int, c int) interface{} { return make([]int, l, c) },
		checkInts,
	)

	testPoolConsist(isPool, t)
}

func TestHugePool(t *testing.T) {
	isPool := NewSlicePool(
		func(l int, c int) interface{} { return make([]int, l, c) },
		checkInts,
	)

	testHugePool(isPool, t)
}

func TestPoolEdgeCondition(t *testing.T) {
	bsPool := NewSlicePool(
		func(l int, c int) interface{} { return make([]byte, l, c) },
		checkBytes,
	)

	testPoolEdgeCondition(bsPool, t)
}

func TestDifferentTypePanic(t *testing.T) {

	bsPool := NewSlicePool(
		func(l int, c int) interface{} { return make([]byte, l, c) },
		checkBytes,
	)

	testDifferentTypePanic(bsPool, t)
}
func TestPoolFull(t *testing.T) {

	bsPool := NewSlicePool(
		func(l int, c int) interface{} { return make([]byte, l, c) },
		checkBytes,
	)

	for i := 0; i <= cacheSliceCap+1; i++ {
		bsPool.Return(make([]byte, 0, 8))
	}
}

func BenchmarkSliceBorrowReturn(t *testing.B) {

	bytesPool := NewSlicePool(
		func(l int, c int) interface{} { return make([]byte, l, c) },
		checkBytes,
	)

	for i := 0; i < t.N; i++ {
		size := rand.Intn(maxSliceSize)
		if size == 0 {
			continue
		}

		v := bytesPool.Borrow(size)
		b, ok := v.([]byte)
		if !ok {
			t.Fatal(v, "is not slice type!")
		}

		if len(b) != size || cap(b) < len(b) {
			t.Fatal("length:", len(b), "is less than cap:", cap(b))
		} else {
			bytesPool.Return(b)
		}
	}
}
