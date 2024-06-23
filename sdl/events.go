package sdl

//#include "SDL_go.h"
import "C"

import (
	"unsafe"
)

type EventType uint32

const (
	FIRSTEVENT      EventType = C.SDL_FIRSTEVENT
	QUIT            EventType = C.SDL_QUIT
	DISPLAYEVENT    EventType = C.SDL_DISPLAYEVENT
	WINDOWEVENT     EventType = C.SDL_WINDOWEVENT
	WM_EVENT        EventType = C.SDL_SYSWMEVENT
	KEYDOWN         EventType = C.SDL_KEYDOWN
	KEYUP           EventType = C.SDL_KEYUP
	TEXTEDITING     EventType = C.SDL_TEXTEDITING
	TEXTINPUT       EventType = C.SDL_TEXTINPUT
	KEYMAPCHANGED   EventType = C.SDL_KEYMAPCHANGED
	TEXTEDITING_EXT EventType = C.SDL_TEXTEDITING_EXT
	MOUSEMOTION     EventType = C.SDL_MOUSEMOTION
	MOUSEBUTTONDOWN EventType = C.SDL_MOUSEBUTTONDOWN
	MOUSEBUTTONUP   EventType = C.SDL_MOUSEBUTTONUP
	MOUSEWHEEL      EventType = C.SDL_MOUSEWHEEL

	PRESSED  int = C.SDL_PRESSED
	RELEASED int = C.SDL_RELEASED
)

/*
 * Event
 */
type WindowEvent struct{ data []byte }
type KeyboardEvent struct{ data []byte }
type MouseMotionEvent struct{ data []byte }
type MouseButtonEvent struct{ data []byte }

type Event struct {
	event C.SDL_Event
	data  []byte

	Window WindowEvent
	Key    KeyboardEvent
	Motion MouseMotionEvent
	Button MouseButtonEvent
}

func updateEventData(event *Event, data []byte) {
	event.data = data
	event.Window.data = data
	event.Key.data = data
	event.Motion.data = data
	event.Button.data = data
}

func (event Event) Type() EventType {
	return EventType(readUint32(event.data, C.offsetof_SDL_CommonEvent_type))
}

/*
 * WindowEvent
 */

func (event WindowEvent) Type() EventType {
	return EventType(readUint32(event.data, C.offsetof_SDL_CommonEvent_type))
}

func (event WindowEvent) Timestamp() uint32 {
	return readUint32(event.data, C.offsetof_SDL_CommonEvent_timestamp)
}

func (event WindowEvent) WindowID() uint32 {
	return readUint32(event.data, C.offsetof_SDL_WindowEvent_windowID)
}

func (event WindowEvent) Event() WindowEventID {
	return WindowEventID(readUint8(event.data, C.offsetof_SDL_WindowEvent_event))
}

func (event WindowEvent) Data1() int32 {
	return readInt32(event.data, C.offsetof_SDL_WindowEvent_data1)
}

func (event WindowEvent) Data2() int32 {
	return readInt32(event.data, C.offsetof_SDL_WindowEvent_data2)
}

/*
 * KeyboardEvent
 */

func (event KeyboardEvent) Type() EventType {
	return EventType(readUint32(event.data, C.offsetof_SDL_CommonEvent_type))
}

func (event KeyboardEvent) Timestamp() uint32 {
	return readUint32(event.data, C.offsetof_SDL_CommonEvent_timestamp)
}

func (event KeyboardEvent) WindowID() uint32 {
	return readUint32(event.data, C.offsetof_SDL_KeyboardEvent_windowID)
}

func (event KeyboardEvent) State() uint8 {
	return readUint8(event.data, C.offsetof_SDL_KeyboardEvent_state)
}

func (event KeyboardEvent) Repeat() uint8 {
	return readUint8(event.data, C.offsetof_SDL_KeyboardEvent_repeat)
}

func (event KeyboardEvent) Keysym() Keysym {
	return createKeysym(readBytes(event.data, C.offsetof_SDL_KeyboardEvent_keysym, C.sizeof_SDL_Keysym))
}

/*
 * MouseMotionEvent
 */

func (event MouseMotionEvent) Type() EventType {
	return EventType(readUint32(event.data, C.offsetof_SDL_CommonEvent_type))
}

func (event MouseMotionEvent) Timestamp() uint32 {
	return readUint32(event.data, C.offsetof_SDL_CommonEvent_timestamp)
}

func (event MouseMotionEvent) WindowID() uint32 {
	return readUint32(event.data, C.offsetof_SDL_MouseMotionEvent_windowID)
}

func (event MouseMotionEvent) Which() uint32 {
	return readUint32(event.data, C.offsetof_SDL_MouseMotionEvent_which)
}

func (event MouseMotionEvent) State() uint32 {
	return readUint32(event.data, C.offsetof_SDL_MouseMotionEvent_state)
}

func (event MouseMotionEvent) X() int32 {
	return readInt32(event.data, C.offsetof_SDL_MouseMotionEvent_x)
}

func (event MouseMotionEvent) Y() int32 {
	return readInt32(event.data, C.offsetof_SDL_MouseMotionEvent_y)
}

func (event MouseMotionEvent) Xrel() int32 {
	return readInt32(event.data, C.offsetof_SDL_MouseMotionEvent_xrel)
}

func (event MouseMotionEvent) Yrel() int32 {
	return readInt32(event.data, C.offsetof_SDL_MouseMotionEvent_yrel)
}

/*
 * MouseButtonEvent
 */

func (event MouseButtonEvent) Type() EventType {
	return EventType(readUint32(event.data, C.offsetof_SDL_CommonEvent_type))
}

func (event MouseButtonEvent) Timestamp() uint32 {
	return readUint32(event.data, C.offsetof_SDL_CommonEvent_timestamp)
}

func (event MouseButtonEvent) WindowID() uint32 {
	return readUint32(event.data, C.offsetof_SDL_MouseButtonEvent_windowID)
}

func (event MouseButtonEvent) Which() uint32 {
	return readUint32(event.data, C.offsetof_SDL_MouseButtonEvent_which)
}

func (event MouseButtonEvent) Button() uint8 {
	return readUint8(event.data, C.offsetof_SDL_MouseButtonEvent_button)
}

func (event MouseButtonEvent) State() uint8 {
	return readUint8(event.data, C.offsetof_SDL_MouseButtonEvent_state)
}

func (event MouseButtonEvent) Clicks() uint8 {
	return readUint8(event.data, C.offsetof_SDL_MouseButtonEvent_clicks)
}

func (event MouseButtonEvent) X() int32 {
	return readInt32(event.data, C.offsetof_SDL_MouseButtonEvent_x)
}

func (event MouseButtonEvent) Y() int32 {
	return readInt32(event.data, C.offsetof_SDL_MouseButtonEvent_y)
}

/*
 * Functions
 */

func PollEvent(event *Event) int {
	result := int(C.SDL_PollEvent((*C.SDL_Event)(unsafe.Pointer(&event.event))))

	if result != 0 {
		updateEventData(event, C.GoBytes(unsafe.Pointer(&event.event), C.sizeof_SDL_Event))
	} else {
		updateEventData(event, nil)
	}

	return result
}
