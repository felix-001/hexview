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
	charW         int
	charH         int
	hexPos        int
	firstLinePos  int
	secondLinePos int
	asciiPos      int
	data          []byte
	painter       *gui.QPainter
}

func New(data []byte) *HexView {
	scrollArea := *widgets.NewQAbstractScrollArea(nil)
	scrollArea.SetFocusPolicy(core.Qt__StrongFocus)
	ch := core.NewQChar8(core.NewQLatin1Char("9"))
	charW := scrollArea.FontMetrics().HorizontalAdvance2(ch)
	charH := scrollArea.FontMetrics().Height()
	hexPos := AddrStartPos + AddrAreaLen*charW + GapAddrAreaHexArea
	firstLinePos := hexPos - (GapAddrAreaHexArea / 2)
	// -1 : 最后一个空格
	asciiPos := hexPos + (BytesPerLine*3-1)*charW + GapHexAreaAsciiArea
	secondLinePos := asciiPos - (GapHexAreaAsciiArea / 2)
	return &HexView{
		QAbstractScrollArea: scrollArea,
		charW:               charW,
		charH:               charH,
		data:                data,
		painter:             nil,
		hexPos:              hexPos,
		firstLinePos:        firstLinePos,
		asciiPos:            asciiPos,
		secondLinePos:       secondLinePos,
	}
}

func (self *HexView) drawChar(painter *gui.QPainter, x, y int, c string) {
	rect := core.NewQRect4(x, y, self.charW, self.charH)
	painter.DrawText4(rect, int(core.Qt__AlignHCenter|core.Qt__AlignVCenter), c, nil)
}

func (self *HexView) drawText(painter *gui.QPainter, x, y int, s string) {
	for _, ch := range s {
		self.drawChar(painter, x, y, string(ch))
		x += self.charW
	}
}

func (self *HexView) addr2Text(addr int) string {
	return fmt.Sprintf("%08X", addr)
}

func (self *HexView) Show() {
	self.ConnectPaintEvent(func(event *gui.QPaintEvent) {
		//log.Println("ConnectPaintEvent")
		//if self.painter == nil {
		//self.painter = gui.NewQPainter2(self.Viewport())
		self.painter = gui.NewQPainter2(self.Viewport())
		//}
		self.VerticalScrollBar().SetPageStep(self.Viewport().Height() / self.charH)
		self.VerticalScrollBar().SetRangeDefault(0, 1000)
		//self.painter.DrawLine3(self.firstLinePos, event.Rect().Top(), self.firstLinePos, self.Height())
		//self.painter.DrawLine3(self.secondLinePos, event.Rect().Top(), self.secondLinePos, self.Height())
		self.painter.DrawLine3(self.firstLinePos, event.Rect().Top(), self.firstLinePos, self.Height())
		self.painter.DrawLine3(self.secondLinePos, event.Rect().Top(), self.secondLinePos, self.Height())
		for line := 0; line < 60; line++ {
			self.drawText(self.painter, AddrStartPos, line*self.charH, self.addr2Text(line*BytesPerLine))
			for i := 0; i < BytesPerLine; i++ {
				self.drawText(self.painter, self.hexPos+i*3*self.charW, line*self.charH, fmt.Sprintf("%02X", self.data[line*BytesPerLine+i]))
				b := self.data[line*BytesPerLine+i]
				ascii := string(b)
				if b < 0x20 || b > 0x7e {
					ascii = "."
				}
				self.drawChar(self.painter, self.asciiPos+i*self.charW, line*self.charH, ascii)
			}
		}
		self.painter.End()
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
	window.SetMinimumSize2(620, 700)
	window.SetWindowTitle("hexview")
	hexview := New([]byte(testStr))
	window.SetCentralWidget(hexview)
	hexview.Show()
	window.Show()
	app.Exec()
}
