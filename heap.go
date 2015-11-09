package sensorserver

// Min Heap

func (h MinFloat32Heap) Len() int           { return len(h) }
func (h MinFloat32Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinFloat32Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinFloat32Heap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(float32))
}

func (h *MinFloat32Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
