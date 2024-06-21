package sdl

/*
#include <SDL.h>

#define offsetof_SDL_CommonEvent_type offsetof(SDL_CommonEvent, type)
#define offsetof_SDL_CommonEvent_timestamp offsetof(SDL_CommonEvent, timestamp)

#define offsetof_SDL_WindowEvent_windowID offsetof(SDL_WindowEvent, windowID)
#define offsetof_SDL_WindowEvent_event offsetof(SDL_WindowEvent, event)
#define offsetof_SDL_WindowEvent_data1 offsetof(SDL_WindowEvent, data1)
#define offsetof_SDL_WindowEvent_data2 offsetof(SDL_WindowEvent, data2)

#define offsetof_SDL_KeyboardEvent_windowID offsetof(SDL_KeyboardEvent, windowID)
#define offsetof_SDL_KeyboardEvent_state offsetof(SDL_KeyboardEvent, state)
#define offsetof_SDL_KeyboardEvent_repeat offsetof(SDL_KeyboardEvent, repeat)
#define offsetof_SDL_KeyboardEvent_keysym offsetof(SDL_KeyboardEvent, keysym)
*/
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

type Event struct {
	event C.SDL_Event
	data  []byte
}

func (event Event) Type() EventType {
	return EventType(readUint32(event.data, C.offsetof_SDL_CommonEvent_type))
}

/*
 * WindowEvent
 */

type WindowEvent struct {
	data []byte
}

func (event Event) Window() WindowEvent {
	return WindowEvent{event.data}
}

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

type KeyboardEvent struct {
	data []byte
}

func (event Event) Key() KeyboardEvent {
	return KeyboardEvent{event.data}
}

func (event KeyboardEvent) Type() EventType {
	return EventType(readUint32(event.data, C.offsetof_SDL_CommonEvent_type))
}

func (event KeyboardEvent) Timestamp() uint32 {
	return readUint32(event.data, C.offsetof_SDL_CommonEvent_timestamp)
}

func (event KeyboardEvent) State() uint8 {
	return readUint8(event.data, C.offsetof_SDL_KeyboardEvent_state)
}

func (event KeyboardEvent) Repeat() uint8 {
	return readUint8(event.data, C.offsetof_SDL_KeyboardEvent_repeat)
}

func (event KeyboardEvent) Keysym() Keysym {
	return Keysym{
		Scancode: Scancode(readUint32(event.data, C.offsetof_SDL_KeyboardEvent_keysym)),
		Sym:      Keycode(readInt32(event.data, C.offsetof_SDL_KeyboardEvent_keysym+4)),
		Mod:      Keymod(readUint16(event.data, C.offsetof_SDL_KeyboardEvent_keysym+8)),
	}
}

/*
 * Functions
 */

func PollEvent(event *Event) int {
	result := int(C.SDL_PollEvent((*C.SDL_Event)(unsafe.Pointer(&event.event))))

	if result != 0 {
		event.data = C.GoBytes(unsafe.Pointer(&event.event), C.sizeof_SDL_Event)
	} else {
		event.data = nil
	}

	return result
}
