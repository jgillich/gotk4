// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"

//export _gotk4_gdk3_EventFunc
func _gotk4_gdk3_EventFunc(arg1 *C.GdkEvent, arg2 C.gpointer) {
	var fn EventFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(EventFunc)
	}

	var _event *Event // out

	{
		v := (*Event)(gextras.NewStructNative(unsafe.Pointer(arg1)))
		v = CopyEventer(v)
		_event = v
	}

	fn(_event)
}

//export _gotk4_gdk3_SeatGrabPrepareFunc
func _gotk4_gdk3_SeatGrabPrepareFunc(arg1 *C.GdkSeat, arg2 *C.GdkWindow, arg3 C.gpointer) {
	var fn SeatGrabPrepareFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(SeatGrabPrepareFunc)
	}

	var _seat Seater     // out
	var _window Windower // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Seater is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Seater)
			return ok
		})
		rv, ok := casted.(Seater)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Seater")
		}
		_seat = rv
	}
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Windower)
			return ok
		})
		rv, ok := casted.(Windower)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
		}
		_window = rv
	}

	fn(_seat, _window)
}

//export _gotk4_gdk3_WindowChildFunc
func _gotk4_gdk3_WindowChildFunc(arg1 *C.GdkWindow, arg2 C.gpointer) (cret C.gboolean) {
	var fn WindowChildFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(WindowChildFunc)
	}

	var _window Windower // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Windower)
			return ok
		})
		rv, ok := casted.(Windower)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
		}
		_window = rv
	}

	ok := fn(_window)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gdk3_Device_ConnectChanged
func _gotk4_gdk3_Device_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gdk3_Device_ConnectToolChanged
func _gotk4_gdk3_Device_ConnectToolChanged(arg0 C.gpointer, arg1 *C.GdkDeviceTool, arg2 C.guintptr) {
	var f func(tool *DeviceTool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(tool *DeviceTool))
	}

	var _tool *DeviceTool // out

	_tool = wrapDeviceTool(coreglib.Take(unsafe.Pointer(arg1)))

	f(_tool)
}

//export _gotk4_gdk3_DeviceManager_ConnectDeviceAdded
func _gotk4_gdk3_DeviceManager_ConnectDeviceAdded(arg0 C.gpointer, arg1 *C.GdkDevice, arg2 C.guintptr) {
	var f func(device Devicer)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(device Devicer))
	}

	var _device Devicer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Devicer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Devicer)
			return ok
		})
		rv, ok := casted.(Devicer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
		}
		_device = rv
	}

	f(_device)
}

//export _gotk4_gdk3_DeviceManager_ConnectDeviceChanged
func _gotk4_gdk3_DeviceManager_ConnectDeviceChanged(arg0 C.gpointer, arg1 *C.GdkDevice, arg2 C.guintptr) {
	var f func(device Devicer)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(device Devicer))
	}

	var _device Devicer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Devicer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Devicer)
			return ok
		})
		rv, ok := casted.(Devicer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
		}
		_device = rv
	}

	f(_device)
}

//export _gotk4_gdk3_DeviceManager_ConnectDeviceRemoved
func _gotk4_gdk3_DeviceManager_ConnectDeviceRemoved(arg0 C.gpointer, arg1 *C.GdkDevice, arg2 C.guintptr) {
	var f func(device Devicer)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(device Devicer))
	}

	var _device Devicer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Devicer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Devicer)
			return ok
		})
		rv, ok := casted.(Devicer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
		}
		_device = rv
	}

	f(_device)
}

//export _gotk4_gdk3_Display_ConnectClosed
func _gotk4_gdk3_Display_ConnectClosed(arg0 C.gpointer, arg1 C.gboolean, arg2 C.guintptr) {
	var f func(isError bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(isError bool))
	}

	var _isError bool // out

	if arg1 != 0 {
		_isError = true
	}

	f(_isError)
}

//export _gotk4_gdk3_Display_ConnectMonitorAdded
func _gotk4_gdk3_Display_ConnectMonitorAdded(arg0 C.gpointer, arg1 *C.GdkMonitor, arg2 C.guintptr) {
	var f func(monitor *Monitor)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(monitor *Monitor))
	}

	var _monitor *Monitor // out

	_monitor = wrapMonitor(coreglib.Take(unsafe.Pointer(arg1)))

	f(_monitor)
}

//export _gotk4_gdk3_Display_ConnectMonitorRemoved
func _gotk4_gdk3_Display_ConnectMonitorRemoved(arg0 C.gpointer, arg1 *C.GdkMonitor, arg2 C.guintptr) {
	var f func(monitor *Monitor)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(monitor *Monitor))
	}

	var _monitor *Monitor // out

	_monitor = wrapMonitor(coreglib.Take(unsafe.Pointer(arg1)))

	f(_monitor)
}

//export _gotk4_gdk3_Display_ConnectOpened
func _gotk4_gdk3_Display_ConnectOpened(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gdk3_Display_ConnectSeatAdded
func _gotk4_gdk3_Display_ConnectSeatAdded(arg0 C.gpointer, arg1 *C.GdkSeat, arg2 C.guintptr) {
	var f func(seat Seater)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(seat Seater))
	}

	var _seat Seater // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Seater is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Seater)
			return ok
		})
		rv, ok := casted.(Seater)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Seater")
		}
		_seat = rv
	}

	f(_seat)
}

//export _gotk4_gdk3_Display_ConnectSeatRemoved
func _gotk4_gdk3_Display_ConnectSeatRemoved(arg0 C.gpointer, arg1 *C.GdkSeat, arg2 C.guintptr) {
	var f func(seat Seater)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(seat Seater))
	}

	var _seat Seater // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Seater is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Seater)
			return ok
		})
		rv, ok := casted.(Seater)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Seater")
		}
		_seat = rv
	}

	f(_seat)
}

//export _gotk4_gdk3_DisplayManager_ConnectDisplayOpened
func _gotk4_gdk3_DisplayManager_ConnectDisplayOpened(arg0 C.gpointer, arg1 *C.GdkDisplay, arg2 C.guintptr) {
	var f func(display *Display)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(display *Display))
	}

	var _display *Display // out

	_display = wrapDisplay(coreglib.Take(unsafe.Pointer(arg1)))

	f(_display)
}

//export _gotk4_gdk3_DragContext_ConnectActionChanged
func _gotk4_gdk3_DragContext_ConnectActionChanged(arg0 C.gpointer, arg1 C.GdkDragAction, arg2 C.guintptr) {
	var f func(action DragAction)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(action DragAction))
	}

	var _action DragAction // out

	_action = DragAction(arg1)

	f(_action)
}

//export _gotk4_gdk3_DragContext_ConnectCancel
func _gotk4_gdk3_DragContext_ConnectCancel(arg0 C.gpointer, arg1 C.GdkDragCancelReason, arg2 C.guintptr) {
	var f func(reason DragCancelReason)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(reason DragCancelReason))
	}

	var _reason DragCancelReason // out

	_reason = DragCancelReason(arg1)

	f(_reason)
}

//export _gotk4_gdk3_DragContext_ConnectDNDFinished
func _gotk4_gdk3_DragContext_ConnectDNDFinished(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gdk3_DragContext_ConnectDropPerformed
func _gotk4_gdk3_DragContext_ConnectDropPerformed(arg0 C.gpointer, arg1 C.gint, arg2 C.guintptr) {
	var f func(time int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(time int))
	}

	var _time int // out

	_time = int(arg1)

	f(_time)
}

//export _gotk4_gdk3_FrameClock_ConnectAfterPaint
func _gotk4_gdk3_FrameClock_ConnectAfterPaint(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gdk3_FrameClock_ConnectBeforePaint
func _gotk4_gdk3_FrameClock_ConnectBeforePaint(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gdk3_FrameClock_ConnectFlushEvents
func _gotk4_gdk3_FrameClock_ConnectFlushEvents(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gdk3_FrameClock_ConnectLayout
func _gotk4_gdk3_FrameClock_ConnectLayout(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gdk3_FrameClock_ConnectPaint
func _gotk4_gdk3_FrameClock_ConnectPaint(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gdk3_FrameClock_ConnectResumeEvents
func _gotk4_gdk3_FrameClock_ConnectResumeEvents(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gdk3_FrameClock_ConnectUpdate
func _gotk4_gdk3_FrameClock_ConnectUpdate(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gdk3_Keymap_ConnectDirectionChanged
func _gotk4_gdk3_Keymap_ConnectDirectionChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gdk3_Keymap_ConnectKeysChanged
func _gotk4_gdk3_Keymap_ConnectKeysChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gdk3_Keymap_ConnectStateChanged
func _gotk4_gdk3_Keymap_ConnectStateChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gdk3_Monitor_ConnectInvalidate
func _gotk4_gdk3_Monitor_ConnectInvalidate(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gdk3_Screen_ConnectCompositedChanged
func _gotk4_gdk3_Screen_ConnectCompositedChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gdk3_Screen_ConnectMonitorsChanged
func _gotk4_gdk3_Screen_ConnectMonitorsChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gdk3_Screen_ConnectSizeChanged
func _gotk4_gdk3_Screen_ConnectSizeChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gdk3_Seat_ConnectDeviceAdded
func _gotk4_gdk3_Seat_ConnectDeviceAdded(arg0 C.gpointer, arg1 *C.GdkDevice, arg2 C.guintptr) {
	var f func(device Devicer)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(device Devicer))
	}

	var _device Devicer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Devicer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Devicer)
			return ok
		})
		rv, ok := casted.(Devicer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
		}
		_device = rv
	}

	f(_device)
}

//export _gotk4_gdk3_Seat_ConnectDeviceRemoved
func _gotk4_gdk3_Seat_ConnectDeviceRemoved(arg0 C.gpointer, arg1 *C.GdkDevice, arg2 C.guintptr) {
	var f func(device Devicer)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(device Devicer))
	}

	var _device Devicer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Devicer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Devicer)
			return ok
		})
		rv, ok := casted.(Devicer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
		}
		_device = rv
	}

	f(_device)
}

//export _gotk4_gdk3_Seat_ConnectToolAdded
func _gotk4_gdk3_Seat_ConnectToolAdded(arg0 C.gpointer, arg1 *C.GdkDeviceTool, arg2 C.guintptr) {
	var f func(tool *DeviceTool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(tool *DeviceTool))
	}

	var _tool *DeviceTool // out

	_tool = wrapDeviceTool(coreglib.Take(unsafe.Pointer(arg1)))

	f(_tool)
}

//export _gotk4_gdk3_Seat_ConnectToolRemoved
func _gotk4_gdk3_Seat_ConnectToolRemoved(arg0 C.gpointer, arg1 *C.GdkDeviceTool, arg2 C.guintptr) {
	var f func(tool *DeviceTool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(tool *DeviceTool))
	}

	var _tool *DeviceTool // out

	_tool = wrapDeviceTool(coreglib.Take(unsafe.Pointer(arg1)))

	f(_tool)
}

//export _gotk4_gdk3_WindowClass_create_surface
func _gotk4_gdk3_WindowClass_create_surface(arg0 *C.GdkWindow, arg1 C.gint, arg2 C.gint) (cret *C.cairo_surface_t) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[WindowOverrides](instance0)
	if overrides.CreateSurface == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected WindowOverrides.CreateSurface, got none")
	}

	var _width int  // out
	var _height int // out

	_width = int(arg1)
	_height = int(arg2)

	surface := overrides.CreateSurface(_width, _height)

	var _ *cairo.Surface

	cret = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))

	return cret
}

//export _gotk4_gdk3_WindowClass_from_embedder
func _gotk4_gdk3_WindowClass_from_embedder(arg0 *C.GdkWindow, arg1 C.gdouble, arg2 C.gdouble, arg3 *C.gdouble, arg4 *C.gdouble) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[WindowOverrides](instance0)
	if overrides.FromEmbedder == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected WindowOverrides.FromEmbedder, got none")
	}

	var _embedderX float64   // out
	var _embedderY float64   // out
	var _offscreenX *float64 // out
	var _offscreenY *float64 // out

	_embedderX = float64(arg1)
	_embedderY = float64(arg2)
	_offscreenX = (*float64)(unsafe.Pointer(arg3))
	_offscreenY = (*float64)(unsafe.Pointer(arg4))

	overrides.FromEmbedder(_embedderX, _embedderY, _offscreenX, _offscreenY)
}

//export _gotk4_gdk3_WindowClass_to_embedder
func _gotk4_gdk3_WindowClass_to_embedder(arg0 *C.GdkWindow, arg1 C.gdouble, arg2 C.gdouble, arg3 *C.gdouble, arg4 *C.gdouble) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[WindowOverrides](instance0)
	if overrides.ToEmbedder == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected WindowOverrides.ToEmbedder, got none")
	}

	var _offscreenX float64 // out
	var _offscreenY float64 // out
	var _embedderX *float64 // out
	var _embedderY *float64 // out

	_offscreenX = float64(arg1)
	_offscreenY = float64(arg2)
	_embedderX = (*float64)(unsafe.Pointer(arg3))
	_embedderY = (*float64)(unsafe.Pointer(arg4))

	overrides.ToEmbedder(_offscreenX, _offscreenY, _embedderX, _embedderY)
}

//export _gotk4_gdk3_Window_ConnectCreateSurface
func _gotk4_gdk3_Window_ConnectCreateSurface(arg0 C.gpointer, arg1 C.gint, arg2 C.gint, arg3 C.guintptr) (cret *C.cairo_surface_t) {
	var f func(width, height int) (surface *cairo.Surface)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(width, height int) (surface *cairo.Surface))
	}

	var _width int  // out
	var _height int // out

	_width = int(arg1)
	_height = int(arg2)

	surface := f(_width, _height)

	var _ *cairo.Surface

	cret = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))

	return cret
}

//export _gotk4_gdk3_Window_ConnectMovedToRect
func _gotk4_gdk3_Window_ConnectMovedToRect(arg0 C.gpointer, arg1 C.gpointer, arg2 C.gpointer, arg3 C.gboolean, arg4 C.gboolean, arg5 C.guintptr) {
	var f func(flippedRect, finalRect unsafe.Pointer, flippedX, flippedY bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg5))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(flippedRect, finalRect unsafe.Pointer, flippedX, flippedY bool))
	}

	var _flippedRect unsafe.Pointer // out
	var _finalRect unsafe.Pointer   // out
	var _flippedX bool              // out
	var _flippedY bool              // out

	_flippedRect = (unsafe.Pointer)(unsafe.Pointer(arg1))
	_finalRect = (unsafe.Pointer)(unsafe.Pointer(arg2))
	if arg3 != 0 {
		_flippedX = true
	}
	if arg4 != 0 {
		_flippedY = true
	}

	f(_flippedRect, _finalRect, _flippedX, _flippedY)
}

//export _gotk4_gdk3_Window_ConnectPickEmbeddedChild
func _gotk4_gdk3_Window_ConnectPickEmbeddedChild(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) (cret *C.GdkWindow) {
	var f func(x, y float64) (window Windower)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(x, y float64) (window Windower))
	}

	var _x float64 // out
	var _y float64 // out

	_x = float64(arg1)
	_y = float64(arg2)

	window := f(_x, _y)

	var _ Windower

	if window != nil {
		cret = (*C.GdkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	}

	return cret
}
