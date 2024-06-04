package astar

type pt struct {
	x int
	y int
}

func NewPt(x, y int) pt {
	return pt{x: x, y: y}
}

type PfNode struct {
	pos    pt
	parent *PfNode
	f      int
	g      int
	h      int
}

func NewPfNode(pos pt, parent *PfNode, f, g, h int) *PfNode {
	pfn := &PfNode{pos, parent, f, g, h}
	return pfn
}
