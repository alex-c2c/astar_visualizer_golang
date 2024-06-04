package astar

type pfHeap []*PfNode

func (h pfHeap) Len() int {
	return len(h)
}

func (h pfHeap) Less(i, j int) bool {
	return h[i].f < h[j].f
}

func (h pfHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *pfHeap) Push(x interface{}) {
	*h = append(*h, x.(*PfNode))
}

func (h *pfHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
