package pool

import (
	"reflect"
)

const cacheSliceCap = 10240

type (
	// SlicePool holds bufs.
	SlicePool struct {
		pools []chan interface{}

		New       func(l int, c int) interface{}
		checkType func(interface{}) bool
	}
)

func NewSlicePool(NewFunc func(l int, c int) interface{}, check func(i interface{}) bool) *SlicePool {
	p := new(SlicePool)

	p.New = NewFunc
	p.checkType = check

	p.pools = make([]chan interface{}, maxSliceType+1)
	min := floorlog2(minSliceSize)
	max := floorlog2(maxSliceSize)
	for i := min; i <= max; i++ {
		// return 2^i size slice
		p.pools[i] = make(chan interface{}, cacheSliceCap)
	}

	return p
}

// borrow a buf from the pool.
func (p *SlicePool) Borrow(size int) interface{} {

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

func (p *SlicePool) borrow(idx int) interface{} {
	var b interface{}
	select {
	case b = <-p.pools[idx]:
	default:
		b = p.New(0, 1<<uint(idx))
	}
	return b
}

// Return returns a buf to the pool.
func (p *SlicePool) Return(b interface{}) {
	if p.checkType(b) == false {
		panic("interface is not property type!")
	}

	v := reflect.ValueOf(b)
	if v.Cap() > maxSliceSize || v.Cap() < minSliceSize {
		return // too big or too small, let it go
	}

	idx := floorlog2(uint(v.Cap()))
	rs := 1 << uint(idx)
	select {
	case p.pools[idx] <- v.Slice3(0, rs, rs).Interface():
	default:
		// let it go, let it go
	}
}
