package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type HexView struct {
	widgets.QAbstractScrollArea
	charWidth int
}

func New() *HexView {
	scrollArea := *widgets.NewQAbstractScrollArea(nil)
	ch := core.NewQChar8(core.NewQLatin1Char("9"))
	charWidth := scrollArea.FontMetrics().HorizontalAdvance2(ch)
	return &HexView{
		QAbstractScrollArea: scrollArea,
		charWidth:           charWidth,
	}
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
