package astar

// MinHeap -- Priority Queue

// ci --> currentIndex
// ei --> endIndex
// pi --> parentIndex
// si --> swapIndex

type CustomHeap []*PfNode

func (h *CustomHeap) Push(value *PfNode) {
	*h = append(*h, value)
	h.siftUp()
}

func (h *CustomHeap) Pop() *PfNode {
	n := len(*h)

	h.swap(0, n-1)

	removeValue := (*h)[n-1]

	*h = (*h)[:n-1]

	h.siftDown(0, n-2)

	return removeValue
}

func (h *CustomHeap) Fix() {
	lastNonLeafIndex := (len(*h) - 2) / 2

	ei := h.Len() - 1
	for ci := lastNonLeafIndex; ci >= 0; ci-- {
		h.siftDown(ci, ei)
	}
}

func (h CustomHeap) Len() int {
	return len(h)
}

func (h *CustomHeap) siftDown(ci, ei int) {
	lci := ci*2 + 1
	for lci <= ei {
		rci := ci*2 + 2
		if rci > ei {
			rci = -1
		}

		si := lci
		if rci != -1 && (*h)[rci].f < (*h)[lci].f {
			si = rci
		}

		if (*h)[si].f < (*h)[ci].f {
			h.swap(si, ci)
			ci = si
			lci = ci*2 + 1
		} else {
			return
		}
	}
}

func (h *CustomHeap) siftUp() {
	ci := len(*h) - 1
	pi := (ci - 1) / 2
	for ci > 0 && (*h)[ci].f < (*h)[pi].f {
		h.swap(ci, pi)
		ci = pi
		pi = (ci - 1) / 2
	}
}

func (h CustomHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
