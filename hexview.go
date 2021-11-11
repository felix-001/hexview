package main

import (
	"os"

	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type HexView struct {
	widgets.QAbstractScrollArea
}

func New() *HexView {
	return &HexView{QAbstractScrollArea: *widgets.NewQAbstractScrollArea(nil)}
}

func (self *HexView) Show() {
	self.ConnectPaintEvent(func(event *gui.QPaintEvent) {
		painter := gui.NewQPainter2(self.Viewport())
		painter.DrawText3(10, 20, "hello world")
		painter.DrawLine3(20, 30, 80, 90)
	})
}

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(800, 700)
	window.SetWindowTitle("hexview")
	hexview := New()
	window.SetCentralWidget(hexview)
	hexview.Show()
	window.Show()
	app.Exec()
}
