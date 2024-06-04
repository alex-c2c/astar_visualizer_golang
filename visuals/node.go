package visuals

import (
	"common"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const NodeSize = 16.0

const (
	StartPointEmpty = iota
	StartPointSetting
	StartPointSet
)

const (
	EndPointEmpty = iota
	EndPointSetting
	EndPointSet
)

const (
	BlockersIdle = iota
	BlockersSetting
)

type nodeRenderer struct {
	bg   *canvas.Rectangle
	icon *canvas.Image
	cir  *canvas.Circle
	text *canvas.Text

	objects []fyne.CanvasObject

	node *Node
}

func (nr *nodeRenderer) MinSize() fyne.Size {
	//return fyne.NewSize(float32(NodeSize)+theme.Padding()*2, float32(NodeSize)+theme.Padding()*2)
	return fyne.NewSize(NodeSize, NodeSize)
}

func (nr *nodeRenderer) Layout(size fyne.Size) {
	switch nr.node.nodeType {
	case common.NodeTypeEmpty:
		nr.bg.Hidden = true
		nr.icon.Hidden = true
		nr.cir.Hidden = false
		nr.text.Hidden = true

		csize := float32(size.Height * 0.1)
		coffset := (NodeSize - csize) * 0.5

		nr.cir.FillColor = color.RGBA{R: 255, G: 255, B: 255, A: 64}
		nr.cir.Resize(fyne.NewSize(csize, csize))
		nr.cir.Move(fyne.NewPos(coffset, coffset))
	case common.NodeTypeStart:
		nr.bg.Hidden = true
		nr.icon.Hidden = true
		nr.cir.Hidden = true
		nr.text.Hidden = false

		nr.text.Text = "S"
		nr.text.Color = color.RGBA{R: 0, G: 128, B: 128, A: 255}
		nr.text.TextStyle = fyne.TextStyle{Bold: true}
		nr.text.TextSize = size.Height
		nr.text.Alignment = fyne.TextAlignCenter
		nr.text.Move(fyne.NewPos(size.Height*0.5, -size.Height*0.25))
	case common.NodeTypeEnd:
		nr.bg.Hidden = true
		nr.icon.Hidden = true
		nr.cir.Hidden = true
		nr.text.Hidden = false

		nr.text.Text = "E"
		nr.text.Color = color.RGBA{R: 64, G: 255, B: 128, A: 255}
		nr.text.TextStyle = fyne.TextStyle{Bold: true}
		nr.text.TextSize = size.Height
		nr.text.Alignment = fyne.TextAlignCenter
		nr.text.Move(fyne.NewPos(size.Height*0.5, -size.Height*0.25))
	case common.NodeTypeBlocker:
		nr.bg.Hidden = true
		nr.icon.Hidden = true
		nr.cir.Hidden = false
		nr.text.Hidden = true

		csize := float32(size.Height * 0.8)
		coffset := (NodeSize - csize) * 0.5

		nr.cir.FillColor = color.RGBA{R: 255, G: 64, B: 64, A: 255}
		nr.cir.Resize(fyne.NewSize(csize, csize))
		nr.cir.Move(fyne.NewPos(coffset, coffset))
	case common.NodeTypePath:
		nr.bg.Hidden = true
		nr.icon.Hidden = true
		nr.cir.Hidden = false
		nr.text.Hidden = true

		csize := float32(size.Height * 0.5)
		coffset := (NodeSize - csize) * 0.5

		nr.cir.FillColor = color.RGBA{R: 0, G: 255, B: 0, A: 128}
		nr.cir.Resize(fyne.NewSize(csize, csize))
		nr.cir.Move(fyne.NewPos(coffset, coffset))
	case common.NodeTypeOpenPath:
		nr.bg.Hidden = true
		nr.icon.Hidden = true
		nr.cir.Hidden = false
		nr.text.Hidden = true

		csize := float32(size.Height * 0.4)
		coffset := (NodeSize - csize) * 0.5

		nr.cir.FillColor = color.RGBA{R: 64, G: 64, B: 255, A: 128}
		nr.cir.Resize(fyne.NewSize(csize, csize))
		nr.cir.Move(fyne.NewPos(coffset, coffset))
	}
}

func (nr *nodeRenderer) ApplyTheme() {

}

func (nr *nodeRenderer) BackgroundColor() color.Color {
	return theme.ButtonColor()
}

func (nr *nodeRenderer) Refresh() {
	/*nr.icon.Hidden = nr.node.icon == nil
	if nr.node.icon != nil {
		nr.icon.Resource = nr.node.icon
	}*/

	nr.Layout(nr.node.Size())
	canvas.Refresh(nr.node)
}

func (nr *nodeRenderer) Objects() []fyne.CanvasObject {
	return nr.objects
}

func (nr *nodeRenderer) Destroy() {

}

type Node struct {
	widget.Button

	x        int
	y        int
	nodeType int
	icon     fyne.Resource

	//tap func(bool)
	tap func(*Node)
}

func (n *Node) Tapped(ev *fyne.PointEvent) {
	n.tap(n)
}

func (n *Node) TappedSecondary(ev *fyne.PointEvent) {
}

func (n *Node) SetIcon(icon fyne.Resource) {
	n.Icon = icon

	n.Refresh()
}

func (n *Node) CreateRenderer() fyne.WidgetRenderer {
	bg := canvas.NewRectangle(color.RGBA{255, 0, 255, 255})
	bg.Resize(fyne.NewSize(NodeSize, NodeSize))

	icon := canvas.NewImageFromResource(n.icon)
	icon.FillMode = canvas.ImageFillContain

	cir := canvas.NewCircle(color.White)

	text := canvas.NewText("", color.White)
	text.Alignment = fyne.TextAlignCenter

	objects := []fyne.CanvasObject{
		bg,
		icon,
		cir,
		text,
	}

	return &nodeRenderer{bg, icon, cir, text, objects, n}
}

func (n *Node) SetNodeType(nodeType int) bool {
	switch nodeType {
	case common.NodeTypeEmpty:
		fallthrough
	case common.NodeTypeStart:
		fallthrough
	case common.NodeTypeEnd:
		fallthrough
	case common.NodeTypeOpenPath:
		fallthrough
	case common.NodeTypePath:
		fallthrough
	case common.NodeTypeBlocker:
		n.nodeType = nodeType
		n.Refresh()
		return true
	}

	fmt.Println("Attempting to set an invalid Node Type: ", nodeType)
	return false
}

func NewNode(x int, y int, tap func(*Node)) *Node {
	node := &Node{x: x, y: y, nodeType: common.NodeTypeEmpty, icon: theme.ContentAddIcon(), tap: tap}
	node.ExtendBaseWidget(node)
	return node
}
