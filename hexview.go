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
	painter       *gui.QPainter
}

func New() *HexView {
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
		painter:             nil,
		hexPos:              hexPos,
		firstLinePos:        firstLinePos,
		asciiPos:            asciiPos,
		secondLinePos:       secondLinePos,
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

func (self *HexView) addr2Text(addr int) string {
	return fmt.Sprintf("%08X", addr)
}

func (self *HexView) drawBasic(event *gui.QPaintEvent) {
	self.painter = gui.NewQPainter2(self.Viewport())
	self.VerticalScrollBar().SetPageStep(self.Viewport().Height() / self.charH)
	self.VerticalScrollBar().SetRangeDefault(0, 100)
	self.painter.DrawLine3(self.firstLinePos, event.Rect().Top(), self.firstLinePos, self.Height())
	self.painter.DrawLine3(self.secondLinePos, event.Rect().Top(), self.secondLinePos, self.Height())
}

// 只显示一个条目的数据，参考wireshark的hexview
func (self *HexView) Show(data []byte) {
	self.ConnectPaintEvent(func(event *gui.QPaintEvent) {
		self.drawBasic(event)
		self.painter.FillRect5(100, 60, 100, 50, gui.NewQColor2(core.Qt__lightGray))
		nbLine := len(data) / BytesPerLine
		for line := 0; line < nbLine; line++ {
			self.drawText(AddrStartPos, line*self.charH, self.addr2Text(line*BytesPerLine))
			for i := 0; i < BytesPerLine; i++ {
				self.drawText(self.hexPos+i*3*self.charW, line*self.charH, fmt.Sprintf("%02X", data[line*BytesPerLine+i]))
				b := data[line*BytesPerLine+i]
				ascii := string(b)
				if b < 0x20 || b > 0x7e {
					ascii = "."
				}
				self.drawChar(self.asciiPos+i*self.charW, line*self.charH, ascii)
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
	hexview := New()
	window.SetCentralWidget(hexview)
	hexview.Show([]byte(testStr))
	window.Show()
	app.Exec()
}
