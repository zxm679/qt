package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QOffscreenSurface struct {
	core.QObject
	QSurface
}

type QOffscreenSurface_ITF interface {
	core.QObject_ITF
	QSurface_ITF
	QOffscreenSurface_PTR() *QOffscreenSurface
}

func (p *QOffscreenSurface) Pointer() unsafe.Pointer {
	return p.QObject_PTR().Pointer()
}

func (p *QOffscreenSurface) SetPointer(ptr unsafe.Pointer) {
	p.QObject_PTR().SetPointer(ptr)
	p.QSurface_PTR().SetPointer(ptr)
}

func PointerFromQOffscreenSurface(ptr QOffscreenSurface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOffscreenSurface_PTR().Pointer()
	}
	return nil
}

func NewQOffscreenSurfaceFromPointer(ptr unsafe.Pointer) *QOffscreenSurface {
	var n = new(QOffscreenSurface)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QOffscreenSurface_") {
		n.SetObjectName("QOffscreenSurface_" + qt.Identifier())
	}
	return n
}

func (ptr *QOffscreenSurface) QOffscreenSurface_PTR() *QOffscreenSurface {
	return ptr
}

func NewQOffscreenSurface(targetScreen QScreen_ITF) *QOffscreenSurface {
	defer qt.Recovering("QOffscreenSurface::QOffscreenSurface")

	return NewQOffscreenSurfaceFromPointer(C.QOffscreenSurface_NewQOffscreenSurface(PointerFromQScreen(targetScreen)))
}

func (ptr *QOffscreenSurface) Create() {
	defer qt.Recovering("QOffscreenSurface::create")

	if ptr.Pointer() != nil {
		C.QOffscreenSurface_Create(ptr.Pointer())
	}
}

func (ptr *QOffscreenSurface) Destroy() {
	defer qt.Recovering("QOffscreenSurface::destroy")

	if ptr.Pointer() != nil {
		C.QOffscreenSurface_Destroy(ptr.Pointer())
	}
}

func (ptr *QOffscreenSurface) IsValid() bool {
	defer qt.Recovering("QOffscreenSurface::isValid")

	if ptr.Pointer() != nil {
		return C.QOffscreenSurface_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QOffscreenSurface) Screen() *QScreen {
	defer qt.Recovering("QOffscreenSurface::screen")

	if ptr.Pointer() != nil {
		return NewQScreenFromPointer(C.QOffscreenSurface_Screen(ptr.Pointer()))
	}
	return nil
}

func (ptr *QOffscreenSurface) ConnectScreenChanged(f func(screen *QScreen)) {
	defer qt.Recovering("connect QOffscreenSurface::screenChanged")

	if ptr.Pointer() != nil {
		C.QOffscreenSurface_ConnectScreenChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "screenChanged", f)
	}
}

func (ptr *QOffscreenSurface) DisconnectScreenChanged() {
	defer qt.Recovering("disconnect QOffscreenSurface::screenChanged")

	if ptr.Pointer() != nil {
		C.QOffscreenSurface_DisconnectScreenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "screenChanged")
	}
}

//export callbackQOffscreenSurfaceScreenChanged
func callbackQOffscreenSurfaceScreenChanged(ptrName *C.char, screen unsafe.Pointer) {
	defer qt.Recovering("callback QOffscreenSurface::screenChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "screenChanged"); signal != nil {
		signal.(func(*QScreen))(NewQScreenFromPointer(screen))
	}

}

func (ptr *QOffscreenSurface) SetFormat(format QSurfaceFormat_ITF) {
	defer qt.Recovering("QOffscreenSurface::setFormat")

	if ptr.Pointer() != nil {
		C.QOffscreenSurface_SetFormat(ptr.Pointer(), PointerFromQSurfaceFormat(format))
	}
}

func (ptr *QOffscreenSurface) SetScreen(newScreen QScreen_ITF) {
	defer qt.Recovering("QOffscreenSurface::setScreen")

	if ptr.Pointer() != nil {
		C.QOffscreenSurface_SetScreen(ptr.Pointer(), PointerFromQScreen(newScreen))
	}
}

func (ptr *QOffscreenSurface) Size() *core.QSize {
	defer qt.Recovering("QOffscreenSurface::size")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QOffscreenSurface_Size(ptr.Pointer()))
	}
	return nil
}

func (ptr *QOffscreenSurface) SurfaceType() QSurface__SurfaceType {
	defer qt.Recovering("QOffscreenSurface::surfaceType")

	if ptr.Pointer() != nil {
		return QSurface__SurfaceType(C.QOffscreenSurface_SurfaceType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QOffscreenSurface) DestroyQOffscreenSurface() {
	defer qt.Recovering("QOffscreenSurface::~QOffscreenSurface")

	if ptr.Pointer() != nil {
		C.QOffscreenSurface_DestroyQOffscreenSurface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QOffscreenSurface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QOffscreenSurface::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QOffscreenSurface) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QOffscreenSurface::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQOffscreenSurfaceTimerEvent
func callbackQOffscreenSurfaceTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QOffscreenSurface::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QOffscreenSurface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QOffscreenSurface::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QOffscreenSurface) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QOffscreenSurface::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQOffscreenSurfaceChildEvent
func callbackQOffscreenSurfaceChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QOffscreenSurface::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QOffscreenSurface) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QOffscreenSurface::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QOffscreenSurface) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QOffscreenSurface::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQOffscreenSurfaceCustomEvent
func callbackQOffscreenSurfaceCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QOffscreenSurface::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}