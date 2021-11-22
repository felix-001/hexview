package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
	std_gui "github.com/therecipe/qt/gui"
	std_widgets "github.com/therecipe/qt/widgets"
)

func cGoFreePacked(ptr unsafe.Pointer) { std_core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_Moc_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return []byte(gs)
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type HexView_ITF interface {
	std_widgets.QAbstractScrollArea_ITF
	HexView_PTR() *HexView
}

func (ptr *HexView) HexView_PTR() *HexView {
	return ptr
}

func (ptr *HexView) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractScrollArea_PTR().Pointer()
	}
	return nil
}

func (ptr *HexView) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractScrollArea_PTR().SetPointer(p)
	}
}

func PointerFromHexView(ptr HexView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.HexView_PTR().Pointer()
	}
	return nil
}

func NewHexViewFromPointer(ptr unsafe.Pointer) (n *HexView) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(HexView)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *HexView:
			n = deduced

		case *std_widgets.QAbstractScrollArea:
			n = &HexView{QAbstractScrollArea: *deduced}

		default:
			n = new(HexView)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackHexViewb5bee9_Constructor
func callbackHexViewb5bee9_Constructor(ptr unsafe.Pointer) {
	this := NewHexViewFromPointer(ptr)
	qt.Register(ptr, this)
}

func HexView_QRegisterMetaType() int {
	return int(int32(C.HexViewb5bee9_HexViewb5bee9_QRegisterMetaType()))
}

func (ptr *HexView) QRegisterMetaType() int {
	return int(int32(C.HexViewb5bee9_HexViewb5bee9_QRegisterMetaType()))
}

func HexView_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.HexViewb5bee9_HexViewb5bee9_QRegisterMetaType2(typeNameC)))
}

func (ptr *HexView) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.HexViewb5bee9_HexViewb5bee9_QRegisterMetaType2(typeNameC)))
}

func HexView_QmlRegisterType() int {
	return int(int32(C.HexViewb5bee9_HexViewb5bee9_QmlRegisterType()))
}

func (ptr *HexView) QmlRegisterType() int {
	return int(int32(C.HexViewb5bee9_HexViewb5bee9_QmlRegisterType()))
}

func HexView_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.HexViewb5bee9_HexViewb5bee9_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *HexView) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.HexViewb5bee9_HexViewb5bee9_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func HexView_QmlRegisterUncreatableType(uri string, versionMajor int, versionMinor int, qmlName string, message string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	var messageC *C.char
	if message != "" {
		messageC = C.CString(message)
		defer C.free(unsafe.Pointer(messageC))
	}
	return int(int32(C.HexViewb5bee9_HexViewb5bee9_QmlRegisterUncreatableType(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC, C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})))
}

func (ptr *HexView) QmlRegisterUncreatableType(uri string, versionMajor int, versionMinor int, qmlName string, message string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	var messageC *C.char
	if message != "" {
		messageC = C.CString(message)
		defer C.free(unsafe.Pointer(messageC))
	}
	return int(int32(C.HexViewb5bee9_HexViewb5bee9_QmlRegisterUncreatableType(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC, C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})))
}

func (ptr *HexView) __scrollBarWidgets_atList(i int) *std_widgets.QWidget {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQWidgetFromPointer(C.HexViewb5bee9___scrollBarWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *HexView) __scrollBarWidgets_setList(i std_widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9___scrollBarWidgets_setList(ptr.Pointer(), std_widgets.PointerFromQWidget(i))
	}
}

func (ptr *HexView) __scrollBarWidgets_newList() unsafe.Pointer {
	return C.HexViewb5bee9___scrollBarWidgets_newList(ptr.Pointer())
}

func (ptr *HexView) __actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.HexViewb5bee9___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *HexView) __actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9___actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *HexView) __actions_newList() unsafe.Pointer {
	return C.HexViewb5bee9___actions_newList(ptr.Pointer())
}

func (ptr *HexView) __addActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.HexViewb5bee9___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *HexView) __addActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9___addActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *HexView) __addActions_actions_newList() unsafe.Pointer {
	return C.HexViewb5bee9___addActions_actions_newList(ptr.Pointer())
}

func (ptr *HexView) __insertActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.HexViewb5bee9___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *HexView) __insertActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9___insertActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *HexView) __insertActions_actions_newList() unsafe.Pointer {
	return C.HexViewb5bee9___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *HexView) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.HexViewb5bee9___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *HexView) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *HexView) __children_newList() unsafe.Pointer {
	return C.HexViewb5bee9___children_newList(ptr.Pointer())
}

func (ptr *HexView) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.HexViewb5bee9___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *HexView) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *HexView) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.HexViewb5bee9___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *HexView) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.HexViewb5bee9___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *HexView) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *HexView) __findChildren_newList() unsafe.Pointer {
	return C.HexViewb5bee9___findChildren_newList(ptr.Pointer())
}

func (ptr *HexView) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.HexViewb5bee9___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *HexView) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *HexView) __findChildren_newList3() unsafe.Pointer {
	return C.HexViewb5bee9___findChildren_newList3(ptr.Pointer())
}

func NewHexView(parent std_widgets.QWidget_ITF) *HexView {
	HexView_QRegisterMetaType()
	tmpValue := NewHexViewFromPointer(C.HexViewb5bee9_NewHexView(std_widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackHexViewb5bee9_DestroyHexView
func callbackHexViewb5bee9_DestroyHexView(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~HexView"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewHexViewFromPointer(ptr).DestroyHexViewDefault()
	}
}

func (ptr *HexView) ConnectDestroyHexView(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~HexView"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~HexView", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~HexView", unsafe.Pointer(&f))
		}
	}
}

func (ptr *HexView) DisconnectDestroyHexView() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~HexView")
	}
}

func (ptr *HexView) DestroyHexView() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.HexViewb5bee9_DestroyHexView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *HexView) DestroyHexViewDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.HexViewb5bee9_DestroyHexViewDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackHexViewb5bee9_ContextMenuEvent
func callbackHexViewb5bee9_ContextMenuEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		(*(*func(*std_gui.QContextMenuEvent))(signal))(std_gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewHexViewFromPointer(ptr).ContextMenuEventDefault(std_gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *HexView) ContextMenuEventDefault(e std_gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_ContextMenuEventDefault(ptr.Pointer(), std_gui.PointerFromQContextMenuEvent(e))
	}
}

//export callbackHexViewb5bee9_DragEnterEvent
func callbackHexViewb5bee9_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*std_gui.QDragEnterEvent))(signal))(std_gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).DragEnterEventDefault(std_gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *HexView) DragEnterEventDefault(event std_gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_DragEnterEventDefault(ptr.Pointer(), std_gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackHexViewb5bee9_DragLeaveEvent
func callbackHexViewb5bee9_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		(*(*func(*std_gui.QDragLeaveEvent))(signal))(std_gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).DragLeaveEventDefault(std_gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *HexView) DragLeaveEventDefault(event std_gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_DragLeaveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackHexViewb5bee9_DragMoveEvent
func callbackHexViewb5bee9_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		(*(*func(*std_gui.QDragMoveEvent))(signal))(std_gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).DragMoveEventDefault(std_gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *HexView) DragMoveEventDefault(event std_gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_DragMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackHexViewb5bee9_DropEvent
func callbackHexViewb5bee9_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		(*(*func(*std_gui.QDropEvent))(signal))(std_gui.NewQDropEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).DropEventDefault(std_gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *HexView) DropEventDefault(event std_gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_DropEventDefault(ptr.Pointer(), std_gui.PointerFromQDropEvent(event))
	}
}

//export callbackHexViewb5bee9_Event
func callbackHexViewb5bee9_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHexViewFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(event)))))
}

func (ptr *HexView) EventDefault(event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.HexViewb5bee9_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackHexViewb5bee9_KeyPressEvent
func callbackHexViewb5bee9_KeyPressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(e))
	} else {
		NewHexViewFromPointer(ptr).KeyPressEventDefault(std_gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *HexView) KeyPressEventDefault(e std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_KeyPressEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(e))
	}
}

//export callbackHexViewb5bee9_MinimumSizeHint
func callbackHexViewb5bee9_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewHexViewFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *HexView) MinimumSizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.HexViewb5bee9_MinimumSizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackHexViewb5bee9_MouseDoubleClickEvent
func callbackHexViewb5bee9_MouseDoubleClickEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(e))
	} else {
		NewHexViewFromPointer(ptr).MouseDoubleClickEventDefault(std_gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *HexView) MouseDoubleClickEventDefault(e std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_MouseDoubleClickEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(e))
	}
}

//export callbackHexViewb5bee9_MouseMoveEvent
func callbackHexViewb5bee9_MouseMoveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(e))
	} else {
		NewHexViewFromPointer(ptr).MouseMoveEventDefault(std_gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *HexView) MouseMoveEventDefault(e std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_MouseMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(e))
	}
}

//export callbackHexViewb5bee9_MousePressEvent
func callbackHexViewb5bee9_MousePressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(e))
	} else {
		NewHexViewFromPointer(ptr).MousePressEventDefault(std_gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *HexView) MousePressEventDefault(e std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_MousePressEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(e))
	}
}

//export callbackHexViewb5bee9_MouseReleaseEvent
func callbackHexViewb5bee9_MouseReleaseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(e))
	} else {
		NewHexViewFromPointer(ptr).MouseReleaseEventDefault(std_gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *HexView) MouseReleaseEventDefault(e std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_MouseReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(e))
	}
}

//export callbackHexViewb5bee9_PaintEvent
func callbackHexViewb5bee9_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		(*(*func(*std_gui.QPaintEvent))(signal))(std_gui.NewQPaintEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).PaintEventDefault(std_gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *HexView) PaintEventDefault(event std_gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_PaintEventDefault(ptr.Pointer(), std_gui.PointerFromQPaintEvent(event))
	}
}

//export callbackHexViewb5bee9_ResizeEvent
func callbackHexViewb5bee9_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		(*(*func(*std_gui.QResizeEvent))(signal))(std_gui.NewQResizeEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).ResizeEventDefault(std_gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *HexView) ResizeEventDefault(event std_gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_ResizeEventDefault(ptr.Pointer(), std_gui.PointerFromQResizeEvent(event))
	}
}

//export callbackHexViewb5bee9_ScrollContentsBy
func callbackHexViewb5bee9_ScrollContentsBy(ptr unsafe.Pointer, dx C.int, dy C.int) {
	if signal := qt.GetSignal(ptr, "scrollContentsBy"); signal != nil {
		(*(*func(int, int))(signal))(int(int32(dx)), int(int32(dy)))
	} else {
		NewHexViewFromPointer(ptr).ScrollContentsByDefault(int(int32(dx)), int(int32(dy)))
	}
}

func (ptr *HexView) ScrollContentsByDefault(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_ScrollContentsByDefault(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)))
	}
}

//export callbackHexViewb5bee9_SetupViewport
func callbackHexViewb5bee9_SetupViewport(ptr unsafe.Pointer, viewport unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setupViewport"); signal != nil {
		(*(*func(*std_widgets.QWidget))(signal))(std_widgets.NewQWidgetFromPointer(viewport))
	} else {
		NewHexViewFromPointer(ptr).SetupViewportDefault(std_widgets.NewQWidgetFromPointer(viewport))
	}
}

func (ptr *HexView) SetupViewportDefault(viewport std_widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_SetupViewportDefault(ptr.Pointer(), std_widgets.PointerFromQWidget(viewport))
	}
}

//export callbackHexViewb5bee9_SizeHint
func callbackHexViewb5bee9_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewHexViewFromPointer(ptr).SizeHintDefault())
}

func (ptr *HexView) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.HexViewb5bee9_SizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackHexViewb5bee9_ViewportEvent
func callbackHexViewb5bee9_ViewportEvent(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "viewportEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHexViewFromPointer(ptr).ViewportEventDefault(std_core.NewQEventFromPointer(event)))))
}

func (ptr *HexView) ViewportEventDefault(event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.HexViewb5bee9_ViewportEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackHexViewb5bee9_ViewportSizeHint
func callbackHexViewb5bee9_ViewportSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "viewportSizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewHexViewFromPointer(ptr).ViewportSizeHintDefault())
}

func (ptr *HexView) ViewportSizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.HexViewb5bee9_ViewportSizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackHexViewb5bee9_WheelEvent
func callbackHexViewb5bee9_WheelEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		(*(*func(*std_gui.QWheelEvent))(signal))(std_gui.NewQWheelEventFromPointer(e))
	} else {
		NewHexViewFromPointer(ptr).WheelEventDefault(std_gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *HexView) WheelEventDefault(e std_gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_WheelEventDefault(ptr.Pointer(), std_gui.PointerFromQWheelEvent(e))
	}
}

//export callbackHexViewb5bee9_ChangeEvent
func callbackHexViewb5bee9_ChangeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(ev))
	} else {
		NewHexViewFromPointer(ptr).ChangeEventDefault(std_core.NewQEventFromPointer(ev))
	}
}

func (ptr *HexView) ChangeEventDefault(ev std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_ChangeEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(ev))
	}
}

//export callbackHexViewb5bee9_ActionEvent
func callbackHexViewb5bee9_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		(*(*func(*std_gui.QActionEvent))(signal))(std_gui.NewQActionEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).ActionEventDefault(std_gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *HexView) ActionEventDefault(event std_gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_ActionEventDefault(ptr.Pointer(), std_gui.PointerFromQActionEvent(event))
	}
}

//export callbackHexViewb5bee9_Close
func callbackHexViewb5bee9_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewHexViewFromPointer(ptr).CloseDefault())))
}

func (ptr *HexView) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.HexViewb5bee9_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackHexViewb5bee9_CloseEvent
func callbackHexViewb5bee9_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		(*(*func(*std_gui.QCloseEvent))(signal))(std_gui.NewQCloseEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).CloseEventDefault(std_gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *HexView) CloseEventDefault(event std_gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_CloseEventDefault(ptr.Pointer(), std_gui.PointerFromQCloseEvent(event))
	}
}

//export callbackHexViewb5bee9_CustomContextMenuRequested
func callbackHexViewb5bee9_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		(*(*func(*std_core.QPoint))(signal))(std_core.NewQPointFromPointer(pos))
	}

}

//export callbackHexViewb5bee9_EnterEvent
func callbackHexViewb5bee9_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).EnterEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *HexView) EnterEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_EnterEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackHexViewb5bee9_FocusInEvent
func callbackHexViewb5bee9_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).FocusInEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *HexView) FocusInEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_FocusInEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackHexViewb5bee9_FocusNextPrevChild
func callbackHexViewb5bee9_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(bool) bool)(signal))(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHexViewFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *HexView) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.HexViewb5bee9_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackHexViewb5bee9_FocusOutEvent
func callbackHexViewb5bee9_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).FocusOutEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *HexView) FocusOutEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_FocusOutEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackHexViewb5bee9_HasHeightForWidth
func callbackHexViewb5bee9_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewHexViewFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *HexView) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.HexViewb5bee9_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackHexViewb5bee9_HeightForWidth
func callbackHexViewb5bee9_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32((*(*func(int) int)(signal))(int(int32(w)))))
	}

	return C.int(int32(NewHexViewFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *HexView) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.HexViewb5bee9_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackHexViewb5bee9_Hide
func callbackHexViewb5bee9_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewHexViewFromPointer(ptr).HideDefault()
	}
}

func (ptr *HexView) HideDefault() {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_HideDefault(ptr.Pointer())
	}
}

//export callbackHexViewb5bee9_HideEvent
func callbackHexViewb5bee9_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		(*(*func(*std_gui.QHideEvent))(signal))(std_gui.NewQHideEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).HideEventDefault(std_gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *HexView) HideEventDefault(event std_gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_HideEventDefault(ptr.Pointer(), std_gui.PointerFromQHideEvent(event))
	}
}

//export callbackHexViewb5bee9_InitPainter
func callbackHexViewb5bee9_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		(*(*func(*std_gui.QPainter))(signal))(std_gui.NewQPainterFromPointer(painter))
	} else {
		NewHexViewFromPointer(ptr).InitPainterDefault(std_gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *HexView) InitPainterDefault(painter std_gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_InitPainterDefault(ptr.Pointer(), std_gui.PointerFromQPainter(painter))
	}
}

//export callbackHexViewb5bee9_InputMethodEvent
func callbackHexViewb5bee9_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		(*(*func(*std_gui.QInputMethodEvent))(signal))(std_gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).InputMethodEventDefault(std_gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *HexView) InputMethodEventDefault(event std_gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_InputMethodEventDefault(ptr.Pointer(), std_gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackHexViewb5bee9_InputMethodQuery
func callbackHexViewb5bee9_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(std_core.Qt__InputMethodQuery) *std_core.QVariant)(signal))(std_core.Qt__InputMethodQuery(query)))
	}

	return std_core.PointerFromQVariant(NewHexViewFromPointer(ptr).InputMethodQueryDefault(std_core.Qt__InputMethodQuery(query)))
}

func (ptr *HexView) InputMethodQueryDefault(query std_core.Qt__InputMethodQuery) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.HexViewb5bee9_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		qt.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackHexViewb5bee9_KeyReleaseEvent
func callbackHexViewb5bee9_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).KeyReleaseEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *HexView) KeyReleaseEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_KeyReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackHexViewb5bee9_LeaveEvent
func callbackHexViewb5bee9_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).LeaveEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *HexView) LeaveEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_LeaveEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackHexViewb5bee9_Lower
func callbackHexViewb5bee9_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewHexViewFromPointer(ptr).LowerDefault()
	}
}

func (ptr *HexView) LowerDefault() {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_LowerDefault(ptr.Pointer())
	}
}

//export callbackHexViewb5bee9_Metric
func callbackHexViewb5bee9_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32((*(*func(std_gui.QPaintDevice__PaintDeviceMetric) int)(signal))(std_gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewHexViewFromPointer(ptr).MetricDefault(std_gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *HexView) MetricDefault(m std_gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.HexViewb5bee9_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackHexViewb5bee9_MoveEvent
func callbackHexViewb5bee9_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		(*(*func(*std_gui.QMoveEvent))(signal))(std_gui.NewQMoveEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).MoveEventDefault(std_gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *HexView) MoveEventDefault(event std_gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_MoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMoveEvent(event))
	}
}

//export callbackHexViewb5bee9_NativeEvent
func callbackHexViewb5bee9_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result *C.long) C.char {
	var resultR int
	if result != nil {
		resultR = int(int32(*result))
		defer func() { *result = C.long(int32(resultR)) }()
	}
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QByteArray, unsafe.Pointer, *int) bool)(signal))(std_core.NewQByteArrayFromPointer(eventType), message, &resultR))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHexViewFromPointer(ptr).NativeEventDefault(std_core.NewQByteArrayFromPointer(eventType), message, &resultR))))
}

func (ptr *HexView) NativeEventDefault(eventType std_core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {
	if ptr.Pointer() != nil {
		var resultC C.long
		if result != nil {
			resultC = C.long(int32(*result))
			defer func() { *result = int(int32(resultC)) }()
		}
		return int8(C.HexViewb5bee9_NativeEventDefault(ptr.Pointer(), std_core.PointerFromQByteArray(eventType), message, &resultC)) != 0
	}
	return false
}

//export callbackHexViewb5bee9_PaintEngine
func callbackHexViewb5bee9_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return std_gui.PointerFromQPaintEngine((*(*func() *std_gui.QPaintEngine)(signal))())
	}

	return std_gui.PointerFromQPaintEngine(NewHexViewFromPointer(ptr).PaintEngineDefault())
}

func (ptr *HexView) PaintEngineDefault() *std_gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return std_gui.NewQPaintEngineFromPointer(C.HexViewb5bee9_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackHexViewb5bee9_Raise
func callbackHexViewb5bee9_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewHexViewFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *HexView) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_RaiseDefault(ptr.Pointer())
	}
}

//export callbackHexViewb5bee9_Repaint
func callbackHexViewb5bee9_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewHexViewFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *HexView) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_RepaintDefault(ptr.Pointer())
	}
}

//export callbackHexViewb5bee9_SetDisabled
func callbackHexViewb5bee9_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		(*(*func(bool))(signal))(int8(disable) != 0)
	} else {
		NewHexViewFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *HexView) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackHexViewb5bee9_SetEnabled
func callbackHexViewb5bee9_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewHexViewFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *HexView) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackHexViewb5bee9_SetFocus2
func callbackHexViewb5bee9_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewHexViewFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *HexView) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackHexViewb5bee9_SetHidden
func callbackHexViewb5bee9_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		(*(*func(bool))(signal))(int8(hidden) != 0)
	} else {
		NewHexViewFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *HexView) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackHexViewb5bee9_SetStyleSheet
func callbackHexViewb5bee9_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(styleSheet))
	} else {
		NewHexViewFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *HexView) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.HexViewb5bee9_SetStyleSheetDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackHexViewb5bee9_SetVisible
func callbackHexViewb5bee9_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	} else {
		NewHexViewFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *HexView) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackHexViewb5bee9_SetWindowModified
func callbackHexViewb5bee9_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewHexViewFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *HexView) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackHexViewb5bee9_SetWindowTitle
func callbackHexViewb5bee9_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	} else {
		NewHexViewFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *HexView) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.HexViewb5bee9_SetWindowTitleDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackHexViewb5bee9_Show
func callbackHexViewb5bee9_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewHexViewFromPointer(ptr).ShowDefault()
	}
}

func (ptr *HexView) ShowDefault() {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_ShowDefault(ptr.Pointer())
	}
}

//export callbackHexViewb5bee9_ShowEvent
func callbackHexViewb5bee9_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		(*(*func(*std_gui.QShowEvent))(signal))(std_gui.NewQShowEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).ShowEventDefault(std_gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *HexView) ShowEventDefault(event std_gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_ShowEventDefault(ptr.Pointer(), std_gui.PointerFromQShowEvent(event))
	}
}

//export callbackHexViewb5bee9_ShowFullScreen
func callbackHexViewb5bee9_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewHexViewFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *HexView) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackHexViewb5bee9_ShowMaximized
func callbackHexViewb5bee9_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewHexViewFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *HexView) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackHexViewb5bee9_ShowMinimized
func callbackHexViewb5bee9_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewHexViewFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *HexView) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackHexViewb5bee9_ShowNormal
func callbackHexViewb5bee9_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewHexViewFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *HexView) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackHexViewb5bee9_TabletEvent
func callbackHexViewb5bee9_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		(*(*func(*std_gui.QTabletEvent))(signal))(std_gui.NewQTabletEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).TabletEventDefault(std_gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *HexView) TabletEventDefault(event std_gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_TabletEventDefault(ptr.Pointer(), std_gui.PointerFromQTabletEvent(event))
	}
}

//export callbackHexViewb5bee9_Update
func callbackHexViewb5bee9_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewHexViewFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *HexView) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_UpdateDefault(ptr.Pointer())
	}
}

//export callbackHexViewb5bee9_UpdateMicroFocus
func callbackHexViewb5bee9_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewHexViewFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *HexView) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackHexViewb5bee9_WindowIconChanged
func callbackHexViewb5bee9_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		(*(*func(*std_gui.QIcon))(signal))(std_gui.NewQIconFromPointer(icon))
	}

}

//export callbackHexViewb5bee9_WindowTitleChanged
func callbackHexViewb5bee9_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	}

}

//export callbackHexViewb5bee9_ChildEvent
func callbackHexViewb5bee9_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *HexView) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackHexViewb5bee9_ConnectNotify
func callbackHexViewb5bee9_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewHexViewFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *HexView) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackHexViewb5bee9_CustomEvent
func callbackHexViewb5bee9_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *HexView) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackHexViewb5bee9_DeleteLater
func callbackHexViewb5bee9_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewHexViewFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *HexView) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.HexViewb5bee9_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackHexViewb5bee9_Destroyed
func callbackHexViewb5bee9_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackHexViewb5bee9_DisconnectNotify
func callbackHexViewb5bee9_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewHexViewFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *HexView) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackHexViewb5bee9_EventFilter
func callbackHexViewb5bee9_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHexViewFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *HexView) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.HexViewb5bee9_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackHexViewb5bee9_ObjectNameChanged
func callbackHexViewb5bee9_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackHexViewb5bee9_TimerEvent
func callbackHexViewb5bee9_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewHexViewFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *HexView) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HexViewb5bee9_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

func init() {
}
