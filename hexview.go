package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

const (
	AddrAreaLen         = 10
	GapAddrAreaHexArea  = 10 // 地址区域和hex区域之间的长度
	GapHexAreaAsciiArea = 16 // hex区域和ascii区域之间的长度
	BytesPerLine        = 16 // 一行显示16个字节
)

type HexView struct {
	widgets.QAbstractScrollArea
	charW   int
	charH   int
	data    []byte
	painter *gui.QPainter
}

func New(data []byte) *HexView {
	scrollArea := *widgets.NewQAbstractScrollArea(nil)
	ch := core.NewQChar8(core.NewQLatin1Char("9"))
	charW := scrollArea.FontMetrics().HorizontalAdvance2(ch)
	charH := scrollArea.FontMetrics().Height()
	return &HexView{
		QAbstractScrollArea: scrollArea,
		charW:               charW,
		charH:               charH,
		data:                data,
		painter:             nil, //gui.NewQPainter2(scrollArea.Viewport()),
	}
}

func (self *HexView) drawChar(x, y int, c string) {
	rect := core.NewQRect4(x, y, self.charW, self.charH)
	self.painter.DrawText4(rect, int(core.Qt__AlignHCenter|core.Qt__AlignVCenter), c, nil)
}

func (self *HexView) drawText(x, y int, s string) {
	for _, ch := range s {
		self.drawChar(x, y, string(ch))
		x += self.charW
	}
}

func (self *HexView) Show() {
	self.ConnectPaintEvent(func(event *gui.QPaintEvent) {
		self.painter = gui.NewQPainter2(self.Viewport())
		self.painter.SetPen2(gui.NewQColor2(core.Qt__black))
		hexPos := AddrAreaLen*self.charW + GapAddrAreaHexArea
		asciiPos := hexPos + BytesPerLine*3*self.charW + GapHexAreaAsciiArea
		linePos := asciiPos - (GapHexAreaAsciiArea / 2)
		self.painter.DrawLine3(linePos, event.Rect().Top(), linePos, self.Height())
		for line := 0; line < 5; line++ {
			yPos := (line + 1) * self.charH
			addr := fmt.Sprintf("%08x", line*BytesPerLine)
			self.drawText(0, yPos, addr)
			/*
				x := 0
				for i := 0; i < 8; i++ {
					rect := core.NewQRect4(x, yPos, self.charW, self.charH)
					self.painter.DrawText4(rect, int(core.Qt__AlignHCenter|core.Qt__AlignVCenter), string(addr[i]), nil)
					//painter.DrawText3(x, yPos, string(addr[i]))
					x += self.charW
				}
			*/
			//painter.DrawText3(0, yPos, fmt.Sprintf("%08x", line*BytesPerLine))
			xPos := hexPos
			for i := 0; i < BytesPerLine; i++ {
				self.drawText(xPos, yPos, fmt.Sprintf("%02x", self.data[line*BytesPerLine+i]&0xf0))
				/*
					rect := core.NewQRect4(xPos, yPos, self.charW, self.charH)
					painter.DrawText4(rect, int(core.Qt__AlignHCenter|core.Qt__AlignVCenter), fmt.Sprintf("%x", self.data[line*BytesPerLine+i]&0xf0>>4), nil)
					rect = core.NewQRect4(xPos+self.charW, yPos, self.charW, self.charH)
					painter.DrawText4(rect, int(core.Qt__AlignHCenter|core.Qt__AlignVCenter), fmt.Sprintf("%x", self.data[line*BytesPerLine+i]&0x0f), nil)
				*/
				//painter.DrawText3(xPos, yPos, fmt.Sprintf("%x", self.data[line*BytesPerLine+i]&0xf0>>4))
				//painter.DrawText3(xPos+self.charW, yPos, fmt.Sprintf("%x", self.data[line*BytesPerLine+i]&0x0f))
				xPos += 3 * self.charW
			}
		}
	})
}

func main() {
	testStr := `
	QPainter performs low-level painting on widgets and other paint devices. The class can draw everything from simple lines to complex shapes like pies and chords. It can also draw aligned text and pixmaps. Normally, it draws in a "natural" coordinate system, but it can in addition do view and world transformation.
`

	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(800, 700)
	window.SetWindowTitle("hexview")
	hexview := New([]byte(testStr))
	window.SetCentralWidget(hexview)
	hexview.Show()
	window.Show()
	app.Exec()
}
