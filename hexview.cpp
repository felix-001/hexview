#include <QApplication>
#include <QMainWindow>
#include <QScrollBar>
#include <QPainter>
#include <QPaintEvent>
#include <QAbstractScrollArea>
#include <stdio.h>

enum {
	kAddrAreaLen         = 8,
	kGapAddrAreaHexArea  = 16, // 地址区域和hex区域之间的长度
	kGapHexAreaAsciiArea = 16, // hex区域和ascii区域之间的长度
	kBytesPerLine        = 16, // 一行显示16个字节
	kAddrStartPos        = 8,  // 地址区域偏移
};

#define log(args...) printf("%s %d ", __FILE__, __LINE__);printf(args);printf("\n")

class HexView : public QAbstractScrollArea 
{
private:
	int font_width_;
	int font_height_;
	int hex_pos_;
	int ascii_pos_;
	int first_line_pos_;
	int second_line_pos_;
	int sel_start_;
	int sel_end_;
	QByteArray data_;
public:
	HexView();
	~HexView(){}
public:
	void setData(QByteArray data){ data_ = data; }
	void setSelection(int start, int end){ 
		sel_start_ = start; 
		sel_end_ = end;
		//log("start:%d end:%d", sel_start_, sel_end_);
	}
protected:
	void paintEvent(QPaintEvent *event) override;
	void mousePressEvent(QMouseEvent *event) override;
	void mouseMoveEvent(QMouseEvent * event) override;
private:
	void drawBasic(QPainter *painter, QPaintEvent *event);
	bool isSelected(int offset);
	bool isSelectedEnd(int offset);
	int cursorOffset(QPoint point);
};

HexView::HexView():
	sel_start_(-1),
	sel_end_(-1)
{
	setFont(QFont("Menlo", 13));
	//setFocusPolicy(Qt::StrongFocus);

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
	verticalScrollBar()->setPageStep(viewport()->size().height() / font_height_);
	verticalScrollBar()->setRange(0, 30);
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


void HexView::paintEvent(QPaintEvent *event)
{
	QPainter painter(viewport());
	drawBasic(&painter, event);
	int nbLine = data_.size()/kBytesPerLine;
	int lineIdx = verticalScrollBar()->value();
	//printf("lineIdx:%d\n", lineIdx);
	for (int line = lineIdx; line < nbLine; line++) {
		QString addr = QString("%1").arg(line * kBytesPerLine, 8, 16, QChar('0'));
		//int y = (line-lineIdx+1)*font_height_;
		int y = (line-lineIdx)*font_height_;
		painter.drawText(kAddrStartPos, y, addr.length()*font_width_, font_height_, Qt::AlignTop, addr);
		for (int i = 0; i<kBytesPerLine; i++) {
			int offset = (line - lineIdx)*kBytesPerLine + i;
			int x = hex_pos_ + i*3*font_width_;
			if (isSelected(offset)) {
				QColor color = palette().color(QPalette::Active, QPalette::Highlight);
				int w = font_width_*3;
				if (isSelectedEnd(offset)) {
					w = font_width_*2;
				}
				QRectF rect(x, y, w, font_height_);
				painter.fillRect(rect, color);
			}
			char c = data_.at(offset);
			QString hex = QString::number(c, 16).toUpper();
			painter.drawText(x, y, font_width_*2, font_height_, Qt::AlignTop, hex);
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
	}
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
				   "hello world, this is a test, good"));
	//hexView->setSelection(10, 28);
	window.setCentralWidget(hexView);
	window.show();
	return app.exec();
 }