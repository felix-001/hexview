package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

const (
	AddrAreaLen         = 8
	GapAddrAreaHexArea  = 16 // 地址区域和hex区域之间的长度
	GapHexAreaAsciiArea = 16 // hex区域和ascii区域之间的长度
	BytesPerLine        = 16 // 一行显示16个字节
	AddrStartPos        = 8  // 地址区域偏移
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
		//self.painter.SetPen2(gui.NewQColor2(core.Qt__black))
		hexPos := AddrStartPos + AddrAreaLen*self.charW + GapAddrAreaHexArea
		firstLinePos := hexPos - (GapAddrAreaHexArea / 2)
		self.painter.DrawLine3(firstLinePos, event.Rect().Top(), firstLinePos, self.Height())
		// -1 ： 最后一个空格
		asciiPos := hexPos + (BytesPerLine*3-1)*self.charW + GapHexAreaAsciiArea
		linePos := asciiPos - (GapHexAreaAsciiArea / 2)
		self.painter.DrawLine3(linePos, event.Rect().Top(), linePos, self.Height())
		for line := 0; line < 40; line++ {
			//yPos := line * self.charH
			addr := fmt.Sprintf("%08X", line*BytesPerLine)
			self.drawText(AddrStartPos, line*self.charH, addr)
			//xPos := hexPos
			for i := 0; i < BytesPerLine; i++ {
				hex := fmt.Sprintf("%02X", self.data[line*BytesPerLine+i])
				self.drawText(hexPos+i*3*self.charW, line*self.charH, hex)
				//xPos += 3 * self.charW
			}
		}
	})
}

func main() {
	testStr := `
	QPainter performs low-level painting on widgets and other paint devices. The class can draw everything from simple lines to complex shapes like pies and chords. It can also draw aligned text and pixmaps. Normally, it draws in a "natural" coordinate system, but it can in addition do view and world transformation.
	QPainter performs low-level painting on widgets and other paint devices. The class can draw everything from simple lines to complex shapes like pies and chords. It can also draw aligned text and pixmaps. Normally, it draws in a "natural" coordinate system, but it can in addition do view and world transformation.
	QPainter performs low-level painting on widgets and other paint devices. The class can draw everything from simple lines to complex shapes like pies and chords. It can also draw aligned text and pixmaps. Normally, it draws in a "natural" coordinate system, but it can in addition do view and world transformation.
	QPainter performs low-level painting on widgets and other paint devices. The class can draw everything from simple lines to complex shapes like pies and chords. It can also draw aligned text and pixmaps. Normally, it draws in a "natural" coordinate system, but it can in addition do view and world transformation.
	QPainter performs low-level painting on widgets and other paint devices. The class can draw everything from simple lines to complex shapes like pies and chords. It can also draw aligned text and pixmaps. Normally, it draws in a "natural" coordinate system, but it can in addition do view and world transformation.
	QPainter performs low-level painting on widgets and other paint devices. The class can draw everything from simple lines to complex shapes like pies and chords. It can also draw aligned text and pixmaps. Normally, it draws in a "natural" coordinate system, but it can in addition do view and world transformation.
	QPainter performs low-level painting on widgets and other paint devices. The class can draw everything from simple lines to complex shapes like pies and chords. It can also draw aligned text and pixmaps. Normally, it draws in a "natural" coordinate system, but it can in addition do view and world transformation.
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
