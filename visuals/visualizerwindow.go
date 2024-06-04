package visuals

import (
	"astar"
	"common"
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var startPointStatus int = StartPointEmpty
var endPointStatus int = EndPointEmpty
var blockersStatus int = BlockersIdle

var labelSP *canvas.Text // label Start Point
var labelEP *canvas.Text // label End Point
var labelB *canvas.Text  // Label Blockers

var buttonSP *widget.Button    // button Start Point
var buttonEP *widget.Button    // button End Point
var buttonSB *widget.Button    //button Set Blockers
var buttonCB *widget.Button    // button Clear Blockers
var buttonGRB *widget.Button   // button Generate Random Blockers
var buttonStart *widget.Button // button Start

var startNode *Node
var endNode *Node
var blockerNodes []*Node
var pathNodes []*Node

func clearStartNode() {
	if startNode != nil {
		startNode.SetNodeType(common.NodeTypeEmpty)
	}

	startNode = nil
	labelSP.Text = fmt.Sprintf("Start: [ , ]")
	labelSP.Refresh()
}

func setStartNode(node *Node) {
	if node == nil {
		return
	}

	if startNode != nil {
		startNode.SetNodeType(common.NodeTypeEmpty)
	}

	node.SetNodeType(common.NodeTypeStart)
	startNode = node

	labelSP.Text = fmt.Sprintf("Start: [ %d , %d ]", startNode.x, startNode.y)
	labelSP.Refresh()
}

func clearEndNode() {
	if endNode != nil {
		endNode.SetNodeType(common.NodeTypeEmpty)
	}

	endNode = nil
	labelEP.Text = fmt.Sprintf("End: [ , ]")
	labelEP.Refresh()
}

func setEndNode(node *Node) {
	if node == nil {
		return
	}

	if endNode != nil {
		endNode.SetNodeType(common.NodeTypeEmpty)
	}

	node.SetNodeType(common.NodeTypeEnd)
	endNode = node

	labelEP.Text = fmt.Sprintf("End: [ %d , %d ]", endNode.x, endNode.y)
	labelEP.Refresh()
}

func addBlockerNode(node *Node) {
	node.SetNodeType(common.NodeTypeBlocker)
	blockerNodes = append(blockerNodes, node)

	labelB.Text = fmt.Sprintf("Blockers: %d", len(blockerNodes))
	labelB.Refresh()
}

func removeBlockerNode(node *Node) {
	for i := 0; i < len(blockerNodes); i++ {
		if blockerNodes[i] == node {
			blockerNodes[i].SetNodeType(common.NodeTypeEmpty)
			blockerNodes[i] = blockerNodes[len(blockerNodes)-1]
			blockerNodes = blockerNodes[:len(blockerNodes)-1]

			return
		}
	}

	labelB.Text = fmt.Sprintf("Blockers: %d", len(blockerNodes))
	labelB.Refresh()
}

func clearBlockerNodes() {
	for i := 0; i < len(blockerNodes); i++ {
		if blockerNodes[i] != nil {
			blockerNodes[i].SetNodeType(common.NodeTypeEmpty)
		}
	}

	blockerNodes = []*Node{}

	labelB.Text = "Blockers: 0"
	labelB.Refresh()
}

func clearPathNodes() {
	for i := 0; i < len(pathNodes); i++ {
		if pathNodes[i] != nil {
			pathNodes[i].SetNodeType(common.NodeTypeEmpty)
		}
	}

	pathNodes = []*Node{}
}

func setStartPointStatus(status int) {
	switch status {
	case StartPointEmpty:
		buttonSP.Text = "Set Start Point"
	case StartPointSetting:
		buttonSP.Text = "Stop Setting Start Point..."
	case StartPointSet:
		buttonSP.Text = "Clear Start Point"
	default:
		fmt.Printf("Invalid status: %d", status)
		return
	}

	startPointStatus = status
	buttonSP.Refresh()
}

func setEndPointStatus(status int) {
	switch status {
	case EndPointEmpty:
		buttonEP.Text = "Set End Point"
	case EndPointSetting:
		buttonEP.Text = "Stop Setting End Point..."
	case EndPointSet:
		buttonEP.Text = "Clear End Point"
	default:
		fmt.Printf("Invalid status: %d", status)
		return
	}

	endPointStatus = status
	buttonEP.Refresh()
}

func setBlockersStatus(status int) {
	switch status {
	case BlockersIdle:
		buttonSB.Text = "Set Blockers"
	case BlockersSetting:
		buttonSB.Text = "Stop Setting Blockers..."
	default:
		fmt.Printf("Invalid status: %d", status)
		return
	}

	blockersStatus = status
	buttonSB.Refresh()
}

func refreshEndPointStatus() {
	if endNode == nil {
		setEndPointStatus(EndPointEmpty)
	} else {
		setEndPointStatus(EndPointSet)
	}
}

func refreshStartPointStatus() {
	if startNode == nil {
		setStartPointStatus(StartPointEmpty)
	} else {
		setStartPointStatus(StartPointSet)
	}
}

func displayPath(grid *[][]*Node, path *[][2]int) {
	for i := len(*path) - 1; i >= 0; i-- {
		n := (*grid)[(*path)[i][1]][(*path)[i][0]]
		if n.nodeType == common.NodeTypeEmpty {
			n.SetNodeType(common.NodeTypePath)
			pathNodes = append(pathNodes, n)
		}
	}
}

func nodePress(n *Node) {
	if startPointStatus == StartPointSetting {
		if n.nodeType == common.NodeTypeEmpty {
			setStartNode(n)
		}
	} else if endPointStatus == EndPointSetting {
		if n.nodeType == common.NodeTypeEmpty {
			setEndNode(n)
		}
	} else if blockersStatus == BlockersSetting {
		if n.nodeType == common.NodeTypeEmpty {
			addBlockerNode(n)
		} else if n.nodeType == common.NodeTypeBlocker {
			removeBlockerNode(n)
		}
	}
}

func buttonPressStartPoint() {
	switch startPointStatus {
	case StartPointEmpty:
		setStartPointStatus(StartPointSetting)
	case StartPointSetting:
		refreshStartPointStatus()
	case StartPointSet:
		if startNode != nil {
			startNode.SetNodeType(common.NodeTypeEmpty)
		}

		clearStartNode()
		setStartPointStatus(StartPointEmpty)
	}

	refreshEndPointStatus()
	setBlockersStatus(BlockersIdle)

	buttonSP.Refresh()
}

func buttonPressEndPoint() {
	switch endPointStatus {
	case EndPointEmpty:
		setEndPointStatus(EndPointSetting)
	case EndPointSetting:
		refreshEndPointStatus()
	case EndPointSet:
		if endNode != nil {
			endNode.SetNodeType(common.NodeTypeEmpty)
		}

		clearEndNode()
		setEndPointStatus(EndPointEmpty)
	}

	refreshStartPointStatus()
	setBlockersStatus(BlockersIdle)

	buttonEP.Refresh()
}

func buttonPressBlockersSet() {
	switch blockersStatus {
	case BlockersIdle:
		buttonSB.Text = "Setting Blockers..."
		blockersStatus = BlockersSetting
	case BlockersSetting:
		buttonSB.Text = "Set Blockers"
		blockersStatus = BlockersIdle
	}

	refreshStartPointStatus()
	refreshEndPointStatus()

	buttonSB.Refresh()
}

func buttonPressBlockersClear() {
	setBlockersStatus(BlockersIdle)

	clearBlockerNodes()
}

func buttonPressGenerateRandomBlockers() {

}

func buttonPressClearAll() {
	setStartPointStatus(StartPointEmpty)
	setEndPointStatus(EndPointEmpty)
	setBlockersStatus(BlockersIdle)

	clearStartNode()
	clearEndNode()
	clearBlockerNodes()
	clearPathNodes()
}

func getReturnPath(colSize, rowSize int, startNode *Node, endNode *Node, blockerNodes *[]*Node) *[][2]int {
	b := make([][2]int, len(*blockerNodes))
	for i := 0; i < len(*blockerNodes); i++ {
		b[i] = [2]int{(*blockerNodes)[i].x, (*blockerNodes)[i].y}
	}

	startTime := time.Now()
	path := astar.StartPathFinding(colSize, rowSize, [2]int{startNode.x, startNode.y}, [2]int{endNode.x, endNode.y}, b)
	elapseTime := time.Now().Sub(startTime)

	fmt.Printf("Elapse Time: %dms\n", elapseTime.Milliseconds())

	if path == nil {
		fmt.Print("Unable to find a path!")
	} else {
		fmt.Println("Path found:", path)
	}

	return &path
}

func CreateVisualizerWindow(app fyne.App, colSize int, rowSize int) {
	w := app.NewWindow(fmt.Sprintf("A-star Visualizer | Map size: %d x %d", colSize, rowSize))

	// Setting up Nodes in grid
	grid := make([][]*Node, rowSize)
	gridContainer := container.New(layout.NewGridLayout(colSize))
	for r := 0; r < rowSize; r++ {
		grid[r] = make([]*Node, colSize)
		for c := 0; c < colSize; c++ {
			n := NewNode(c, r, func(n *Node) {
				nodePress(n)
			})
			grid[r][c] = n
			gridContainer.Add(n)
		}
	}

	// UI
	labelSP = canvas.NewText(fmt.Sprintf("Start: [ , ]"), color.White)
	buttonSP = widget.NewButton("Set Start Point", func() {
		buttonPressStartPoint()
	})
	c1 := container.New(layout.NewHBoxLayout(), labelSP, buttonSP)

	labelEP = canvas.NewText(fmt.Sprintf("End: [ , ]"), color.White)
	buttonEP = widget.NewButton("Set End Point", func() {
		buttonPressEndPoint()
	})
	c2 := container.New(layout.NewHBoxLayout(), labelEP, buttonEP)

	labelB = canvas.NewText(fmt.Sprintf("Blockers: 0"), color.White)
	buttonSB = widget.NewButton("Set Blockers", func() {
		buttonPressBlockersSet()
	})
	buttonCB = widget.NewButton("Clear Blockers", func() {
		buttonPressBlockersClear()
	})
	buttonGRB = widget.NewButton("Generate Random Blockers", func() {
		buttonPressGenerateRandomBlockers()
	})
	c3 := container.New(layout.NewHBoxLayout(), labelB, buttonSB, buttonCB, buttonGRB)

	buttonClearAll := widget.NewButton("Clear All", buttonPressClearAll)
	buttonStart = widget.NewButton("Start Visualizer", func() {
		refreshStartPointStatus()
		refreshEndPointStatus()
		setBlockersStatus(BlockersIdle)

		if startNode == nil {
			dialog.NewInformation("Error", "Start node required", w).Show()
			return
		}

		if endNode == nil {
			dialog.NewInformation("Error", "End node required", w).Show()
			return
		}

		path := getReturnPath(colSize, rowSize, startNode, endNode, &blockerNodes)

		displayPath(&grid, path)
	})

	mainContent := container.New(layout.NewVBoxLayout(), c1, c2, c3, buttonClearAll, buttonStart, gridContainer)

	w.SetContent(mainContent)

	w.Resize(fyne.NewSize(NodeSize*float32(colSize), NodeSize*float32(rowSize)))
	w.Show()
}
