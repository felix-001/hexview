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
	int char_w_;
	int char_h_;
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
	setFont(QFont("Courier", 13));
	setFocusPolicy(Qt::StrongFocus);

	char_w_ = fontMetrics().horizontalAdvance(QLatin1Char('9'));
	//char_h_ = fontMetrics().height();
	QFontMetrics fm(QFont("Courier", 13));
	char_h_ = fm.height();
	hex_pos_ = kAddrStartPos + kAddrAreaLen*char_w_ + kGapAddrAreaHexArea;
	first_line_pos_ = hex_pos_ - (kGapAddrAreaHexArea / 2);
	// -1 : 最后一个空格
	ascii_pos_ = hex_pos_ + (kBytesPerLine*3-1)*char_w_ + kGapHexAreaAsciiArea;
	second_line_pos_ = ascii_pos_ - (kGapHexAreaAsciiArea / 2);

}

void HexView::drawBasic(QPainter *painter, QPaintEvent *event)
{
	verticalScrollBar()->setPageStep(viewport()->size().height() / char_h_);
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
		int y = (line-lineIdx+1)*char_h_;
		painter.drawText(kAddrStartPos, y, addr);
		for (int i = 0; i<kBytesPerLine; i++) {
			int offset = (line - lineIdx)*kBytesPerLine + i;
			int x = hex_pos_ + i*3*char_w_;
			if (isSelected(offset)) {
				QColor color = Qt::gray; //palette().color(QPalette::Active, QPalette::Highlight);
				painter.fillRect(x, y-char_h_, char_w_*3, char_h_, color);
			}
			char c = data_.at(offset);
			QString hex = QString::number(c, 16).toUpper();
			painter.drawText(x, y, hex);
			if ((c < 0x20) || (c > 0x7e)) {
				c = '.';
			}
			painter.drawText(ascii_pos_ + i*char_w_, y, QString(c));
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