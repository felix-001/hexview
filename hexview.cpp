#include <QApplication>
#include <QMainWindow>
#include <QAbstractScrollArea>

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
public:
	HexView();
	~HexView();
protected:
	void paintEvent(QPaintEvent *event);
	void keyPressEvent(QKeyEvent *event);
	void mouseMoveEvent(QMouseEvent *event);
	void mousePressEvent(QMouseEvent *event);
};

HexView::HexView()
{
	char_w_ = fontMetrics().horizontalAdvance(QLatin1Char('9'));
	char_h_ = fontMetrics().height();
	hex_pos_ = kAddrStartPos + kAddrAreaLen*char_w_ + kGapAddrAreaHexArea;
	first_line_pos_ = hex_pos_ - (kGapAddrAreaHexArea / 2);
	// -1 : 最后一个空格
	ascii_pos_ = hex_pos_ + (kBytesPerLine*3-1)*char_w_ + kGapHexAreaAsciiArea;
	second_line_pos_ = ascii_pos_ - (kGapHexAreaAsciiArea / 2);

	setFont(QFont("Courier", 10));
	setFocusPolicy(Qt::StrongFocus);
}

HexView::~HexView()
{

}

void HexView::paintEvent(QPaintEvent *event)
{

}
void HexView::keyPressEvent(QKeyEvent *event)
{

}

void HexView::mouseMoveEvent(QMouseEvent *event)
{

}

void HexView::mousePressEvent(QMouseEvent *event)
{

}

int main(int argc, char *argv[])
{
	QApplication app(argc, argv);
	QMainWindow window;
	window.setMinimumSize(620, 700);
	window.setWindowTitle("hexview");
	window.show();
	return app.exec();
 }