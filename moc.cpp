

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAbstractScrollArea>
#include <QAction>
#include <QActionEvent>
#include <QByteArray>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QHideEvent>
#include <QIcon>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QMetaMethod>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPoint>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QTabletEvent>
#include <QTimerEvent>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


class HexViewb5bee9: public QAbstractScrollArea
{
Q_OBJECT
public:
	HexViewb5bee9(QWidget *parent = Q_NULLPTR) : QAbstractScrollArea(parent) {qRegisterMetaType<quintptr>("quintptr");HexViewb5bee9_HexViewb5bee9_QRegisterMetaType();HexViewb5bee9_HexViewb5bee9_QRegisterMetaTypes();callbackHexViewb5bee9_Constructor(this);};
	 ~HexViewb5bee9() { callbackHexViewb5bee9_DestroyHexView(this); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackHexViewb5bee9_ContextMenuEvent(this, e); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackHexViewb5bee9_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackHexViewb5bee9_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackHexViewb5bee9_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackHexViewb5bee9_DropEvent(this, event); };
	bool event(QEvent * event) { return callbackHexViewb5bee9_Event(this, event) != 0; };
	void keyPressEvent(QKeyEvent * e) { callbackHexViewb5bee9_KeyPressEvent(this, e); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackHexViewb5bee9_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void mouseDoubleClickEvent(QMouseEvent * e) { callbackHexViewb5bee9_MouseDoubleClickEvent(this, e); };
	void mouseMoveEvent(QMouseEvent * e) { callbackHexViewb5bee9_MouseMoveEvent(this, e); };
	void mousePressEvent(QMouseEvent * e) { callbackHexViewb5bee9_MousePressEvent(this, e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackHexViewb5bee9_MouseReleaseEvent(this, e); };
	void paintEvent(QPaintEvent * event) { callbackHexViewb5bee9_PaintEvent(this, event); };
	void resizeEvent(QResizeEvent * event) { callbackHexViewb5bee9_ResizeEvent(this, event); };
	void scrollContentsBy(int dx, int dy) { callbackHexViewb5bee9_ScrollContentsBy(this, dx, dy); };
	void setupViewport(QWidget * viewport) { callbackHexViewb5bee9_SetupViewport(this, viewport); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackHexViewb5bee9_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	bool viewportEvent(QEvent * event) { return callbackHexViewb5bee9_ViewportEvent(this, event) != 0; };
	QSize viewportSizeHint() const { return *static_cast<QSize*>(callbackHexViewb5bee9_ViewportSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void wheelEvent(QWheelEvent * e) { callbackHexViewb5bee9_WheelEvent(this, e); };
	void changeEvent(QEvent * ev) { callbackHexViewb5bee9_ChangeEvent(this, ev); };
	void actionEvent(QActionEvent * event) { callbackHexViewb5bee9_ActionEvent(this, event); };
	bool close() { return callbackHexViewb5bee9_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackHexViewb5bee9_CloseEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackHexViewb5bee9_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void enterEvent(QEvent * event) { callbackHexViewb5bee9_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackHexViewb5bee9_FocusInEvent(this, event); };
	bool focusNextPrevChild(bool next) { return callbackHexViewb5bee9_FocusNextPrevChild(this, next) != 0; };
	void focusOutEvent(QFocusEvent * event) { callbackHexViewb5bee9_FocusOutEvent(this, event); };
	bool hasHeightForWidth() const { return callbackHexViewb5bee9_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackHexViewb5bee9_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackHexViewb5bee9_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackHexViewb5bee9_HideEvent(this, event); };
	void initPainter(QPainter * painter) const { callbackHexViewb5bee9_InitPainter(const_cast<void*>(static_cast<const void*>(this)), painter); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackHexViewb5bee9_InputMethodEvent(this, event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackHexViewb5bee9_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	void keyReleaseEvent(QKeyEvent * event) { callbackHexViewb5bee9_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackHexViewb5bee9_LeaveEvent(this, event); };
	void lower() { callbackHexViewb5bee9_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackHexViewb5bee9_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void moveEvent(QMoveEvent * event) { callbackHexViewb5bee9_MoveEvent(this, event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackHexViewb5bee9_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, result) != 0; };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackHexViewb5bee9_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	void raise() { callbackHexViewb5bee9_Raise(this); };
	void repaint() { callbackHexViewb5bee9_Repaint(this); };
	void setDisabled(bool disable) { callbackHexViewb5bee9_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackHexViewb5bee9_SetEnabled(this, vbo); };
	void setFocus() { callbackHexViewb5bee9_SetFocus2(this); };
	void setHidden(bool hidden) { callbackHexViewb5bee9_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); Moc_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackHexViewb5bee9_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackHexViewb5bee9_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackHexViewb5bee9_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackHexViewb5bee9_SetWindowTitle(this, vqsPacked); };
	void show() { callbackHexViewb5bee9_Show(this); };
	void showEvent(QShowEvent * event) { callbackHexViewb5bee9_ShowEvent(this, event); };
	void showFullScreen() { callbackHexViewb5bee9_ShowFullScreen(this); };
	void showMaximized() { callbackHexViewb5bee9_ShowMaximized(this); };
	void showMinimized() { callbackHexViewb5bee9_ShowMinimized(this); };
	void showNormal() { callbackHexViewb5bee9_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackHexViewb5bee9_TabletEvent(this, event); };
	void update() { callbackHexViewb5bee9_Update(this); };
	void updateMicroFocus() { callbackHexViewb5bee9_UpdateMicroFocus(this); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackHexViewb5bee9_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray* t3c6de1 = new QByteArray(title.toUtf8()); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1->prepend("WHITESPACE").constData()+10), t3c6de1->size()-10, t3c6de1 };callbackHexViewb5bee9_WindowTitleChanged(this, titlePacked); };
	void childEvent(QChildEvent * event) { callbackHexViewb5bee9_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackHexViewb5bee9_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackHexViewb5bee9_CustomEvent(this, event); };
	void deleteLater() { callbackHexViewb5bee9_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackHexViewb5bee9_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackHexViewb5bee9_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackHexViewb5bee9_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackHexViewb5bee9_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackHexViewb5bee9_TimerEvent(this, event); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(HexViewb5bee9*)


void HexViewb5bee9_HexViewb5bee9_QRegisterMetaTypes() {
}

int HexViewb5bee9_HexViewb5bee9_QRegisterMetaType()
{
	return qRegisterMetaType<HexViewb5bee9*>();
}

int HexViewb5bee9_HexViewb5bee9_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<HexViewb5bee9*>(const_cast<const char*>(typeName));
}

int HexViewb5bee9_HexViewb5bee9_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<HexViewb5bee9>();
#else
	return 0;
#endif
}

int HexViewb5bee9_HexViewb5bee9_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<HexViewb5bee9>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int HexViewb5bee9_HexViewb5bee9_QmlRegisterUncreatableType(char* uri, int versionMajor, int versionMinor, char* qmlName, struct Moc_PackedString message)
{
#ifdef QT_QML_LIB
	return qmlRegisterUncreatableType<HexViewb5bee9>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName), QString::fromUtf8(message.data, message.len));
#else
	return 0;
#endif
}

void* HexViewb5bee9___scrollBarWidgets_atList(void* ptr, int i)
{
	return ({QWidget * tmp = static_cast<QList<QWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QWidget *>*>(ptr)->size()-1) { static_cast<QList<QWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HexViewb5bee9___scrollBarWidgets_setList(void* ptr, void* i)
{
	static_cast<QList<QWidget *>*>(ptr)->append(static_cast<QWidget*>(i));
}

void* HexViewb5bee9___scrollBarWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWidget *>();
}

void* HexViewb5bee9___actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HexViewb5bee9___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* HexViewb5bee9___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* HexViewb5bee9___addActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HexViewb5bee9___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* HexViewb5bee9___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* HexViewb5bee9___insertActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HexViewb5bee9___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* HexViewb5bee9___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* HexViewb5bee9___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HexViewb5bee9___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* HexViewb5bee9___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* HexViewb5bee9___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void HexViewb5bee9___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* HexViewb5bee9___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* HexViewb5bee9___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HexViewb5bee9___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* HexViewb5bee9___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* HexViewb5bee9___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HexViewb5bee9___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* HexViewb5bee9___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* HexViewb5bee9_NewHexView(void* parent)
{
		return new HexViewb5bee9(static_cast<QWidget*>(parent));
}

void HexViewb5bee9_DestroyHexView(void* ptr)
{
	static_cast<HexViewb5bee9*>(ptr)->~HexViewb5bee9();
}

void HexViewb5bee9_DestroyHexViewDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void HexViewb5bee9_ContextMenuEventDefault(void* ptr, void* e)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void HexViewb5bee9_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void HexViewb5bee9_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void HexViewb5bee9_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void HexViewb5bee9_DropEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::dropEvent(static_cast<QDropEvent*>(event));
}

char HexViewb5bee9_EventDefault(void* ptr, void* event)
{
	return static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::event(static_cast<QEvent*>(event));
}

void HexViewb5bee9_KeyPressEventDefault(void* ptr, void* e)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::keyPressEvent(static_cast<QKeyEvent*>(e));
}

void* HexViewb5bee9_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void HexViewb5bee9_MouseDoubleClickEventDefault(void* ptr, void* e)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::mouseDoubleClickEvent(static_cast<QMouseEvent*>(e));
}

void HexViewb5bee9_MouseMoveEventDefault(void* ptr, void* e)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::mouseMoveEvent(static_cast<QMouseEvent*>(e));
}

void HexViewb5bee9_MousePressEventDefault(void* ptr, void* e)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::mousePressEvent(static_cast<QMouseEvent*>(e));
}

void HexViewb5bee9_MouseReleaseEventDefault(void* ptr, void* e)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::mouseReleaseEvent(static_cast<QMouseEvent*>(e));
}

void HexViewb5bee9_PaintEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::paintEvent(static_cast<QPaintEvent*>(event));
}

void HexViewb5bee9_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::resizeEvent(static_cast<QResizeEvent*>(event));
}

void HexViewb5bee9_ScrollContentsByDefault(void* ptr, int dx, int dy)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::scrollContentsBy(dx, dy);
}

void HexViewb5bee9_SetupViewportDefault(void* ptr, void* viewport)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::setupViewport(static_cast<QWidget*>(viewport));
}

void* HexViewb5bee9_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

char HexViewb5bee9_ViewportEventDefault(void* ptr, void* event)
{
	return static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::viewportEvent(static_cast<QEvent*>(event));
}

void* HexViewb5bee9_ViewportSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::viewportSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void HexViewb5bee9_WheelEventDefault(void* ptr, void* e)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::wheelEvent(static_cast<QWheelEvent*>(e));
}

void HexViewb5bee9_ChangeEventDefault(void* ptr, void* ev)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::changeEvent(static_cast<QEvent*>(ev));
}

void HexViewb5bee9_ActionEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::actionEvent(static_cast<QActionEvent*>(event));
}

char HexViewb5bee9_CloseDefault(void* ptr)
{
	return static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::close();
}

void HexViewb5bee9_CloseEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::closeEvent(static_cast<QCloseEvent*>(event));
}

void HexViewb5bee9_EnterEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::enterEvent(static_cast<QEvent*>(event));
}

void HexViewb5bee9_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::focusInEvent(static_cast<QFocusEvent*>(event));
}

char HexViewb5bee9_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::focusNextPrevChild(next != 0);
}

void HexViewb5bee9_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::focusOutEvent(static_cast<QFocusEvent*>(event));
}

char HexViewb5bee9_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::hasHeightForWidth();
}

int HexViewb5bee9_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::heightForWidth(w);
}

void HexViewb5bee9_HideDefault(void* ptr)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::hide();
}

void HexViewb5bee9_HideEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::hideEvent(static_cast<QHideEvent*>(event));
}

void HexViewb5bee9_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::initPainter(static_cast<QPainter*>(painter));
}

void HexViewb5bee9_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* HexViewb5bee9_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void HexViewb5bee9_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void HexViewb5bee9_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::leaveEvent(static_cast<QEvent*>(event));
}

void HexViewb5bee9_LowerDefault(void* ptr)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::lower();
}

int HexViewb5bee9_MetricDefault(void* ptr, long long m)
{
	return static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void HexViewb5bee9_MoveEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::moveEvent(static_cast<QMoveEvent*>(event));
}

char HexViewb5bee9_NativeEventDefault(void* ptr, void* eventType, void* message, long* result)
{
	return static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::nativeEvent(*static_cast<QByteArray*>(eventType), message, result);
}

void* HexViewb5bee9_PaintEngineDefault(void* ptr)
{
	return static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::paintEngine();
}

void HexViewb5bee9_RaiseDefault(void* ptr)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::raise();
}

void HexViewb5bee9_RepaintDefault(void* ptr)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::repaint();
}

void HexViewb5bee9_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::setDisabled(disable != 0);
}

void HexViewb5bee9_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::setEnabled(vbo != 0);
}

void HexViewb5bee9_SetFocus2Default(void* ptr)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::setFocus();
}

void HexViewb5bee9_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::setHidden(hidden != 0);
}

void HexViewb5bee9_SetStyleSheetDefault(void* ptr, struct Moc_PackedString styleSheet)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void HexViewb5bee9_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::setVisible(visible != 0);
}

void HexViewb5bee9_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::setWindowModified(vbo != 0);
}

void HexViewb5bee9_SetWindowTitleDefault(void* ptr, struct Moc_PackedString vqs)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void HexViewb5bee9_ShowDefault(void* ptr)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::show();
}

void HexViewb5bee9_ShowEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::showEvent(static_cast<QShowEvent*>(event));
}

void HexViewb5bee9_ShowFullScreenDefault(void* ptr)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::showFullScreen();
}

void HexViewb5bee9_ShowMaximizedDefault(void* ptr)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::showMaximized();
}

void HexViewb5bee9_ShowMinimizedDefault(void* ptr)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::showMinimized();
}

void HexViewb5bee9_ShowNormalDefault(void* ptr)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::showNormal();
}

void HexViewb5bee9_TabletEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::tabletEvent(static_cast<QTabletEvent*>(event));
}

void HexViewb5bee9_UpdateDefault(void* ptr)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::update();
}

void HexViewb5bee9_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::updateMicroFocus();
}

void HexViewb5bee9_ChildEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::childEvent(static_cast<QChildEvent*>(event));
}

void HexViewb5bee9_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void HexViewb5bee9_CustomEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::customEvent(static_cast<QEvent*>(event));
}

void HexViewb5bee9_DeleteLaterDefault(void* ptr)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::deleteLater();
}

void HexViewb5bee9_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char HexViewb5bee9_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}



void HexViewb5bee9_TimerEventDefault(void* ptr, void* event)
{
	static_cast<HexViewb5bee9*>(ptr)->QAbstractScrollArea::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
