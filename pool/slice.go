package pool

import (
	"reflect"
	"sync"
)

const maxSliceType = 16
const minSliceSize = 1 << 3            // 8 len slice
const maxSliceSize = 1 << maxSliceType // 64k len slice

type (
	PoolI interface {
		Borrow(size int) interface{}
		Return(b interface{})
	}

	// SliceSyncPool holds bufs.
	syncPool struct {
		capV int
		lenV int
		*sync.Pool
	}

	SliceSyncPool struct {
		pools []*syncPool

		New       func(l int, c int) interface{}
		checkType func(interface{}) bool
	}
)

func newSyncPool(NewFunc func(l int, c int) interface{}, lv int, cv int) *syncPool {
	p := new(syncPool)
	p.capV = cv
	p.lenV = lv
	p.Pool = &sync.Pool{New: func() interface{} { return NewFunc(p.lenV, p.capV) }}
	return p
}

func NewSliceSyncPool(NewFunc func(l int, c int) interface{}, check func(interface{}) bool) *SliceSyncPool {
	p := new(SliceSyncPool)

	p.New = NewFunc
	p.checkType = check

	p.pools = make([]*syncPool, maxSliceType+1)
	min := floorlog2(minSliceSize)
	max := floorlog2(maxSliceSize)
	for i := min; i <= max; i++ {
		// return 2^i size slice
		p.pools[i] = newSyncPool(NewFunc, 0, 1<<uint(i))
	}

	return p
}

// borrow a buf from the pool.
func (p *SliceSyncPool) Borrow(size int) interface{} {

	var ret interface{}

	if size > maxSliceSize {
		return p.New(size, 2*size)
	} else if size < minSliceSize {
		// small than 8 len's slice all return 8 cap interface{}
		ret = p.borrow(floorlog2(minSliceSize))
	} else {
		idx := floorlog2(uint(size))
		if 1<<uint(idx) < size {
			idx++
		}

		ret = p.borrow(idx)
	}

	// Return must check if it is slice, so here we assume the interface
	// has right type
	return reflect.ValueOf(ret).Slice(0, size).Interface()
}

func (p *SliceSyncPool) borrow(idx int) interface{} {
	return p.pools[idx].Get()
}

// Return returns a buf to the pool.
func (p *SliceSyncPool) Return(b interface{}) {

	if p.checkType(b) == false {
		panic("interface is not property type!")
	}

	v := reflect.ValueOf(b)
	if v.Cap() > maxSliceSize || v.Cap() < minSliceSize {
		return // too big or too small, let it go
	}

	idx := floorlog2(uint(v.Cap()))
	rs := 1 << uint(idx)
	p.pools[idx].Put(v.Slice3(0, rs, rs).Interface())
}
