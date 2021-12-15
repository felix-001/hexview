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
	void setSelection(int start, int end){ sel_start_ = start; sel_end_ = end; }
protected:
	void paintEvent(QPaintEvent *event);
private:
	void drawBasic(QPainter *painter, QPaintEvent *event);
	bool isSelected(int offset);
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
				QRectF rect(x, y, font_width_*3, font_height_);
				painter.fillRect(rect, color);
			}
			char c = data_.at(offset);
			QString hex = QString::number(c, 16).toUpper();
			painter.drawText(x, y, font_width_*2, font_height_, Qt::AlignTop, hex);
			if ((c < 0x20) || (c > 0x7e)) {
				c = '.';
			}
			painter.drawText(ascii_pos_ + i*font_width_, y, font_width_, font_height_, Qt::AlignTop, QString(c));
		}
	}
}

int main(int argc, char *argv[])
{
	QApplication app(argc, argv);
	QMainWindow window;
	window.setMinimumSize(700, 700);
	window.setWindowTitle("hexview");
	HexView *hexView = new HexView;
	hexView->setData(QByteArray("hello world, this is a test, good luck, fast food place"));
	hexView->setSelection(10, 28);
	window.setCentralWidget(hexView);
	window.show();
	return app.exec();
 }