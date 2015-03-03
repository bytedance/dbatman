package pool

func floorlog2(size uint) int {
	var idx int = 0
	for size > 1 {
		size >>= 1
		idx++
	}
	return idx
}
