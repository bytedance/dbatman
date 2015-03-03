package pool

import (
	"math/rand"
	"reflect"
	"testing"
	"unsafe"
)

func getBytes(i interface{}, tb testing.TB) []byte {
	var b []byte
	var ok bool
	if b, ok = i.([]byte); !ok {
		tb.Fatal(i, "is not bytes slice type!")
	}
	return b
}

func getInts(i interface{}, tb testing.TB) []int {
	var b []int
	var ok bool
	if b, ok = i.([]int); !ok {
		tb.Fatal(i, "is not int slice type!")
	}
	return b
}

func checkBytes(i interface{}) bool {
	_, ok := i.([]byte)
	return ok
}

func checkInts(i interface{}) bool {
	_, ok := i.([]int)
	return ok
}

func TestSyncPoolConsist(t *testing.T) {

	isSyncPool := NewSliceSyncPool(
		func(l int, c int) interface{} { return make([]int, l, c) },
		checkInts,
	)

	testPoolConsist(isSyncPool, t)
}

func testPoolConsist(isPool PoolI, t testing.TB) {

	contents := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	b := getInts(isPool.Borrow(0), t)
	b = append(b, contents[:]...)
	isPool.Return(b)

	nb := getInts(isPool.Borrow(8), t)

	if (*reflect.SliceHeader)(unsafe.Pointer(&nb)).Data != (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data {
		t.Fatal("not the same underly buffer!")
	}
}

func TestSyncHugePool(t *testing.T) {
	isSyncPool := NewSliceSyncPool(
		func(l int, c int) interface{} { return make([]int, l, c) },
		checkInts,
	)

	testHugePool(isSyncPool, t)
}

func testHugePool(isPool PoolI, tb testing.TB) {
	b := getInts(isPool.Borrow(maxSliceSize+1), tb)
	isPool.Return(b) // should not pool this really big buffer

	nb := getInts(isPool.Borrow(maxSliceSize), tb)
	if (*reflect.SliceHeader)(unsafe.Pointer(&nb)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data {
		tb.Fatal("these two buffer should be different underly array!")
	}
}

func TestSyncPoolEdgeCondition(t *testing.T) {
	bsSyncPool := NewSliceSyncPool(
		func(l int, c int) interface{} { return make([]byte, l, c) },
		checkBytes,
	)

	testPoolEdgeCondition(bsSyncPool, t)
}

func testPoolEdgeCondition(bsSyncPool PoolI, t testing.TB) {

	for i := 1; i <= minSliceSize; i++ {
		s := bsSyncPool.Borrow(i)
		b := getBytes(s, t)

		if len(b) != i {
			t.Fatal("len:", len(b), "not match required size:", i)
		}

		if cap(b) != minSliceSize {
			t.Fatal("cap:", cap(b), "not match minSliceSize:", minSliceSize)
		}

		bsSyncPool.Return(b)
	}

	for i := minSliceSize + 1; i <= maxSliceSize; i++ {
		s := bsSyncPool.Borrow(i)
		b := getBytes(s, t)

		if len(b) != i {
			t.Fatal("len:", len(b), "not match required size:", i)
		}

		fl := floorlog2(uint(i))
		if 1<<uint(fl) < i {
			fl += 1
		}
		if cap(b) != 1<<uint(fl) {
			t.Fatal("cap:", cap(b), "not match real cap size:", 1<<uint(fl))
		}

		bsSyncPool.Return(b)
	}
}

func TestSyncDifferentTypePanic(t *testing.T) {

	bsSyncPool := NewSliceSyncPool(
		func(l int, c int) interface{} { return make([]byte, l, c) },
		checkBytes,
	)

	testDifferentTypePanic(bsSyncPool, t)
}

func testDifferentTypePanic(bsPool PoolI, t testing.TB) {
	defer func() {
		if r := recover(); r == nil {
			panic("no receive panic")
		}
	}()

	bsPool.Return(make([]int, 0, 8))
}

func BenchmarkSliceSyncBorrowReturn(t *testing.B) {
	bsSyncPool := NewSliceSyncPool(
		func(l int, c int) interface{} { return make([]byte, l, c) },
		checkBytes,
	)

	for i := 0; i < t.N; i++ {
		size := rand.Intn(maxSliceSize)
		if size == 0 {
			continue
		}

		v := bsSyncPool.Borrow(size)
		b := getBytes(v, t)

		if len(b) != size || cap(b) < len(b) {
			t.Fatal("length:", len(b), "is less than cap:", cap(b))
		}

		bsSyncPool.Return(b)
	}
}
