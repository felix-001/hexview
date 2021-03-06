#include <QApplication>
#include <QClipboard>
#include <QMainWindow>
#include <QScrollBar>
#include <QPainter>
#include <QPaintEvent>
#include <QMenu>
#include <QAbstractScrollArea>
#include <QtGlobal>
#include <stdio.h>
#include "hexview.h"

enum {
	kAddrAreaLen         = 8,
	kGapAddrAreaHexArea  = 16, // 地址区域和hex区域之间的长度
	kGapHexAreaAsciiArea = 16, // hex区域和ascii区域之间的长度
	kBytesPerLine        = 16, // 一行显示16个字节
	kAddrStartPos        = 8,  // 地址区域偏移
};

#define log(args...) printf("%s %d ", __FILE__, __LINE__);printf(args);printf("\n")

HexView::HexView():
	QAbstractScrollArea(),
	sel_start_(-1),
	sel_end_(-1)
{
	setFont(QFont("Menlo", 13));

	font_width_ = fontMetrics().horizontalAdvance("M");
	font_height_ = fontMetrics().height();
	hex_pos_ = kAddrStartPos + kAddrAreaLen*font_width_ + kGapAddrAreaHexArea;
	first_line_pos_ = hex_pos_ - (kGapAddrAreaHexArea / 2);
	// -1 : 最后一个空格
	ascii_pos_ = hex_pos_ + (kBytesPerLine*3-1)*font_width_ + kGapHexAreaAsciiArea;
	second_line_pos_ = ascii_pos_ - (kGapHexAreaAsciiArea / 2);

}

void HexView::drawBasic(QPainter *painter, QPaintEvent *event)
{
	int totalLines = data_.size() / kBytesPerLine;
	if(data_.size() % kBytesPerLine)
		totalLines++;
	int viewportLines = viewport()->size().height() / font_height_;
	verticalScrollBar()->setPageStep(viewportLines);
	verticalScrollBar()->setRange(0, totalLines - viewportLines);
	painter->drawLine(first_line_pos_, event->rect().top(), first_line_pos_, height());
	painter->drawLine(second_line_pos_, event->rect().top(), second_line_pos_, height());
}

bool HexView::isSelected(int offset)
{
	if (sel_start_ == -1) {
		return false;
	}
	if (sel_end_ == -1) {
		return false;
	}
	return (offset >= sel_start_ && offset <= sel_end_);
}

bool HexView::isSelectedEnd(int offset)
{
	if (offset == sel_end_) {
		return true;
	}
	return false;
}

// 3 = 两个hex字符+一个空格
#define CHARS_PER_HEX (3)
// 光标所在字节在整个buf的偏移
int HexView::cursorOffset(QPoint point)
{
	int hex_end_pos = ascii_pos_ - kGapHexAreaAsciiArea;
	if (point.x() < hex_pos_ || point.x() > hex_end_pos) {
		return -1;
	}
	int x = (point.x() - hex_pos_)/font_width_;
	// 一个hex在屏幕上显示的是两个字符
	x /= CHARS_PER_HEX;
	// line_idx是被滚动条隐藏的部分
	int line_idx = verticalScrollBar() -> value();
	int y = (point.y()/font_height_)*kBytesPerLine;
	int offset = x + y + line_idx*kBytesPerLine;
	return offset;
}

void HexView::mousePressEvent(QMouseEvent *event)
{
	if (event->button() == Qt::RightButton) {
		return;
	}
	//log("mousePressEvent");
	int offset = cursorOffset(event->pos());
	setSelection(offset, offset);
	viewport()->update();
}

void HexView::mouseMoveEvent(QMouseEvent * event)
{
	//log("mouseMoveEvent");
	int offset = cursorOffset(event->pos());
	if (offset > sel_start_) {
		sel_end_ = offset;
	} else {
		sel_start_ = offset;
	}
	viewport()->update();
}

bool HexView::isLineEnd(int offset)
{
	return (offset+1)%kBytesPerLine == 0;
}

void HexView::drawHexHighlight(QPainter &painter, int offset, int x, int y)
{
	QColor color = palette().color(QPalette::Active, QPalette::Highlight);
	int w = font_width_*3;
	if (isSelectedEnd(offset) || isLineEnd(offset)) {
		w = font_width_*2;
	}
	QRectF rect(x, y, w, font_height_);
	painter.fillRect(rect, color);
}

void HexView::drawHex(QPainter &painter, int offset, int i, int y, char c)
{
	int x = hex_pos_ + i*3*font_width_;
	if (isSelected(offset)) {
		drawHexHighlight(painter, offset, x, y);	
	}
	QString hex = QString::number(c, 16).toUpper();
	painter.drawText(x, y, font_width_*2, font_height_, Qt::AlignTop, hex);
}

void HexView::drawAscii(QPainter &painter, int offset, int i, int y, char c)
{
	if ((c < 0x20) || (c > 0x7e)) {
		c = '.';
	}
	if (isSelected(offset)) {
		QColor color = palette().color(QPalette::Active, QPalette::Highlight);
		QRectF rect(ascii_pos_ + i*font_width_, y, font_width_, font_height_);
		painter.fillRect(rect, color);
	}
	painter.drawText(ascii_pos_ + i*font_width_, y, font_width_, font_height_, Qt::AlignTop, QString(c));
}

void HexView::drawLine(QPainter &painter, int line, int y)
{
	for (int i = 0; i<kBytesPerLine; i++) {
		int offset = line*kBytesPerLine + i;
		char c = data_.at(offset);
		drawHex(painter, offset, i, y, c);	
		drawAscii(painter, offset, i, y, c);
	}
}

void HexView::drawAddr(QPainter &painter, int line, int y)
{
	QString addr = QString("%1").arg(line * kBytesPerLine, 8, 16, QChar('0'));
	painter.drawText(kAddrStartPos, y, addr.length()*font_width_, font_height_, Qt::AlignTop, addr);
}

void HexView::paintEvent(QPaintEvent *event)
{
	QPainter painter(viewport());
	drawBasic(&painter, event);
	int nbLine = data_.size()/kBytesPerLine;
	int lineIdx = verticalScrollBar()->value();
	for (int line = lineIdx; line < nbLine; line++) {
		int y = (line-lineIdx)*font_height_;
		drawAddr(painter, line, y);
		drawLine(painter, line, y);	
	}
}

void HexView::contextMenuEvent(QContextMenuEvent *event)
{
	QMenu * menu = new QMenu(this);
	menu->addAction(tr("&Copy as Hex Dump"), this, SLOT(copyHex()));
	menu->addAction(tr("&Copy as Text"), this, SLOT(copyAscii()));
	menu->exec(event->globalPos());
	delete menu;
}

bool HexView::hasSelectedText()
{
	if (sel_start_ == -1) {
		return false;
	}
	if (sel_end_ == -1) {
		return false;
	}
	return true;
}

void HexView::copyAscii()
{
	int start = std::min(sel_start_, sel_end_);
	int end = std::max(sel_start_, sel_end_);
	QString ascii = "";
	while(start <= end) {
		char c = data_.at(start);
		if ((c < 0x20) || (c > 0x7e)) {
			c = '.';
		}	
		ascii += c;
		start++;
	}
	QApplication::clipboard()->setText(ascii);
}

void HexView::fillWithSpace(int offset, QString &hex)
{
	int left = offset%kBytesPerLine;
	if (!left) {
		return;
	}
	for (int i=0; i<left*3; i++) {
		hex += " ";
	}
}

void HexView::copyHex()
{
	int start = std::min(sel_start_, sel_end_);
	int end = std::max(sel_start_, sel_end_);
	int offset = start;
	QString hex = "";
	fillWithSpace(offset, hex);
	while(offset <= end) {
		char c = data_.at(offset);
		hex += QString::number(c, 16).toUpper();	
		if (isLineEnd(offset)) {
			hex += "\n";
		} else {
			hex += " ";
		}
		offset++;
	}
	if ((offset - start + 1)%kBytesPerLine != 0) {
		hex += "\n";
	}
	QApplication::clipboard()->setText(hex);
}

// TODO:
// 1. 滚动，根据数据量的大小
// 2. 支持hex搜索
// 	2.1 表格首先跳到对应的item
//	2.2 hex view显示相应的条目，搜索的hex高亮
// 3. 指定offset查看
//	3.1 表格先跳到对应的item
//	3.2 hex view显示相应的条目
// 4. isSelected是否需要使用坐标的偏移
// 5. Command + A 全选
// 6. shift + 鼠标左键选择
// 7. ascii区域右侧的竖线
int main(int argc, char *argv[])
{
	QApplication app(argc, argv);
	QMainWindow window;
	window.setMinimumSize(630, 700);
	window.setWindowTitle("hexview");
	HexView *hexView = new HexView;
	hexView->setData(QByteArray("hello world, this is a test, good luck, fast food place"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"
				   "hello world, this is a test, good"));
	//hexView->setSelection(10, 28);
	window.setCentralWidget(hexView);
	window.show();
	return app.exec();
 }