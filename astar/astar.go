package astar

import "fmt"

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

func getValidAdjNodes(cn *PfNode, adjSteps *[4]pt, diagSteps *[4]pt, blockerPts *[]pt, colSize, rowSize int) *[]pt {
	children := make([]pt, 0, 8)

	blockTop := false
	blockRight := false
	blockDown := false
	blockLeft := false

	for i := 0; i < 4; i++ {
		nx := cn.pos.x + adjSteps[i].x
		ny := cn.pos.y + adjSteps[i].y

		if nx < 0 || nx >= colSize {
			continue
		}

		if ny < 0 || ny >= rowSize {
			continue
		}

		isBlock := false
		for bi := 0; bi < len(*blockerPts); bi++ {
			if nx == (*blockerPts)[bi].x && ny == (*blockerPts)[bi].y {
				isBlock = true
				break
			}
		}

		if isBlock {
			if adjSteps[i].x == 0 && adjSteps[i].y == -1 {
				blockTop = true
			} else if adjSteps[i].x == 1 && adjSteps[i].y == 0 {
				blockRight = true
			} else if adjSteps[i].x == 0 && adjSteps[i].y == 1 {
				blockDown = true
			} else if adjSteps[i].x == -1 && adjSteps[i].y == 0 {
				blockLeft = true
			}
			continue
		}

		children = append(children, NewPt(nx, ny))
	}

	for i := 0; i < 4; i++ {
		nx := cn.pos.x + diagSteps[i].x
		ny := cn.pos.y + diagSteps[i].y

		if nx < 0 || nx >= colSize {
			continue
		}

		if ny < 0 || ny >= rowSize {
			continue
		}

		isBlock := false
		for bi := 0; bi < len(*blockerPts); bi++ {
			if nx == (*blockerPts)[bi].x && ny == (*blockerPts)[bi].y {
				isBlock = true
				break
			}
		}

		if isBlock {
			continue
		}

		if diagSteps[i].x == 1 && diagSteps[i].y == -1 && blockTop && blockRight {
			continue
		} else if diagSteps[i].x == 1 && diagSteps[i].y == 1 && blockRight && blockDown {
			continue
		} else if diagSteps[i].x == -1 && diagSteps[i].y == 1 && blockDown && blockLeft {
			continue
		} else if diagSteps[i].x == -1 && diagSteps[i].y == -1 && blockLeft && blockTop {
			continue
		}

		children = append(children, NewPt(nx, ny))
	}

	return &children
}

func isInCloseList(pt pt, closeList *[]*PfNode) bool {
	for i := 0; i < len(*closeList); i++ {
		if pt == (*closeList)[i].pos {
			return true
		}
	}

	return false
}

func isInOpenList(pt pt, g int, openList *[]*PfNode) bool {
	for i := 0; i < len(*openList); i++ {
		if pt == (*openList)[i].pos && g > (*openList)[i].g {
			return true
		}
	}

	return false
}

func StartPathFinding(colSize, rowSize int, start [2]int, end [2]int, blockers [][2]int) [][2]int {
	fmt.Println("StartPathFinding")

	sn := NewPfNode(NewPt(start[0], start[1]), nil, 0, 0, 0)
	en := NewPfNode(NewPt(end[0], end[1]), nil, 0, 0, 0)

	blockerPts := make([]pt, len(blockers))
	for i := 0; i < len(blockers); i++ {
		blockerPts[i] = NewPt(blockers[i][0], blockers[i][1])
	}

	openList := []*PfNode{sn}
	closeList := []*PfNode{}

	adjSteps := &[4]pt{NewPt(0, -1), NewPt(1, 0), NewPt(0, 1), NewPt(-1, 0)}
	diagSteps := &[4]pt{NewPt(1, -1), NewPt(1, 1), NewPt(-1, 1), NewPt(-1, -1)}

	for len(openList) > 0 {
		cn := openList[0]

		openList = openList[1:]
		closeList = append(closeList, cn)

		if nodeEqual(cn, en) {
			return getReturnPathFrom(cn)
		}

		children := getValidAdjNodes(cn, adjSteps, diagSteps, &blockerPts, colSize, rowSize)
		for i := 0; i < len(*children); i++ {
			child := (*children)[i]

			if isInCloseList(child, &closeList) {
				continue
			}

			g := cn.g + 1
			h := func(a, b pt) int {
				c := a.x - b.x
				d := a.y - b.y
				return c*c + d*d
			}(child, en.pos)
			f := g + h

			if isInOpenList(child, g, &openList) {
				continue
			}

			openList = append(openList, NewPfNode(child, cn, f, g, h))
		}
	}

	return nil
}
