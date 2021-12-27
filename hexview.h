#include <QApplication>
#include <QMainWindow>
#include <QScrollBar>
#include <QPainter>
#include <QPaintEvent>
#include <QMenu>
#include <QAbstractScrollArea>
#include <QtGlobal>

class HexView : public QAbstractScrollArea 
{
	Q_OBJECT
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
public Q_SLOTS:
	void copyHex();
	void copyAscii();
protected:
	void paintEvent(QPaintEvent *event) override;
	void mousePressEvent(QMouseEvent *event) override;
	void mouseMoveEvent(QMouseEvent * event) override;
	void contextMenuEvent(QContextMenuEvent *event) override;
private:
	void drawBasic(QPainter *painter, QPaintEvent *event);
	bool isSelected(int offset);
	bool isSelectedEnd(int offset);
	int cursorOffset(QPoint point);
	bool hasSelectedText();
	void fillWithSpace(int offset, QString &hex);
	void drawHex(QPainter &painter, int offse, int i, int y, char c);
	void drawAscii(QPainter &painter, int offset, int i, int y, char c);
	void drawLine(QPainter &painter, int lineOffset, int y);
	void drawAddr(QPainter &painter, int line, int y);
};