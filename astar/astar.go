package astar

import (
	"container/heap"
	"fmt"
	"sort"
)

func nodeEqual(a *PfNode, b *PfNode) bool {
	return a.pos == b.pos
}

func getReturnPathFrom(n *PfNode) [][2]int {
	path := [][2]int{}

	cn := n

	for cn != nil {
		path = append(path, [2]int{cn.pos.x, cn.pos.y})
		cn = cn.parent
	}

	return path
}

func isNodeBlocked(x, y int, blockerPts *[]pt) bool {
	for bi := 0; bi < len(*blockerPts); bi++ {
		if x == (*blockerPts)[bi].x && y == (*blockerPts)[bi].y {
			return true
		}
	}

	return false
}

func getValidAdjNodes(cn *PfNode, steps *map[string]pt, blockerPts *[]pt, colSize, rowSize int) *map[string]pt {
	childrenMap := make(map[string]pt)
	blockMap := map[string]bool{"t": false, "r": false, "b": false, "l": false}

	// need a "sorted map", because adj steps need to be processed first
	i := 0
	keySlice := make([]string, 8, 8)
	for k, _ := range *steps {
		keySlice[i] = k
		i++
	}

	sort.Slice(keySlice, func(i, j int) bool {
		l1, l2 := len(keySlice[i]), len(keySlice[j])
		if l1 != l2 {
			return l1 < l2
		}

		return keySlice[i] < keySlice[j]
	})

	for i := 0; i < len(keySlice); i++ {
		k := keySlice[i]
		v := (*steps)[k]
		nx := cn.pos.x + v.x
		ny := cn.pos.y + v.y

		if nx < 0 || nx >= colSize || ny < 0 || ny >= rowSize {
			continue
		}

		isBlock := isNodeBlocked(nx, ny, blockerPts)

		if isBlock {
			blockMap[k] = true
			continue
		}

		if len(k) == 2 {
			if blockMap[string(k[0])] && blockMap[string(k[1])] {
				continue
			}
		}

		childrenMap[k] = NewPt(nx, ny)
	}

	return &childrenMap
}

func isInCloseSlice(pt pt, closeSlice *[]*PfNode) bool {
	for i := 0; i < len(*closeSlice); i++ {
		if pt == (*closeSlice)[i].pos {
			return true
		}
	}

	return false
}

func getIndexInOpenHeap(pt pt, openHeap *pfHeap) int {
	for i := 0; i < openHeap.Len(); i++ {
		if pt == (*openHeap)[i].pos {
			return i
		}
	}

	return -1
}

func buildHeapByInit(array []*PfNode) *pfHeap {
	pfheap := &pfHeap{}
	*pfheap = array
	heap.Init(pfheap)
	return pfheap
}

func StartPathFinding(colSize, rowSize int, start [2]int, end [2]int, blockers [][2]int) [][2]int {
	fmt.Println("StartPathFinding")

	sn := NewPfNode(NewPt(start[0], start[1]), nil, 0, 0, 0)
	en := NewPfNode(NewPt(end[0], end[1]), nil, 0, 0, 0)

	blockerPts := make([]pt, len(blockers))
	for i := 0; i < len(blockers); i++ {
		blockerPts[i] = NewPt(blockers[i][0], blockers[i][1])
	}

	openHeap := buildHeapByInit([]*PfNode{sn})
	closeSlice := []*PfNode{}

	steps := &map[string]pt{"t": NewPt(0, -1), "r": NewPt(1, 0), "b": NewPt(0, 1), "l": NewPt(-1, 0), "tr": NewPt(1, -1), "br": NewPt(1, 1), "bl": NewPt(-1, 1), "tl": NewPt(-1, -1)}

	for openHeap.Len() > 0 {
		cn := heap.Pop(openHeap).(*PfNode)
		closeSlice = append(closeSlice, cn)

		if nodeEqual(cn, en) {
			return getReturnPathFrom(cn)
		}

		validNodes := getValidAdjNodes(cn, steps, &blockerPts, colSize, rowSize)
		for k, v := range *validNodes {
			child := v

			if isInCloseSlice(child, &closeSlice) {
				continue
			}

			var g int
			if len(k) == 2 {
				g = cn.g + 15
			} else {
				g = cn.g + 10
			}
			h := square(child, en.pos)
			f := g + h

			existingIndex := getIndexInOpenHeap(child, openHeap)
			if existingIndex != -1 {
				node := (*openHeap)[existingIndex]
				if node.g > g {
					node.g = g
					node.h = h
					node.f = f
					node.parent = cn
					heap.Fix(openHeap, existingIndex)
				}
				continue
			}

			heap.Push(openHeap, NewPfNode(child, cn, f, g, h))
		}
	}

	return nil
}

func square(a, b pt) int {
	c := a.x - b.x
	d := a.y - b.y
	return c*c + d*d
}
