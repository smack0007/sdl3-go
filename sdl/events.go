package sdl

/*
#include "events.h"

SDL_Event _event;

static inline bool _SDL_PollEvent()
{
  return SDL_PollEvent(&_event);
}
*/
import "C"

import (
	"unsafe"
)

type EventType uint32

const (
	EVENT_FIRST                         EventType = C.SDL_EVENT_FIRST
	EVENT_QUIT                          EventType = C.SDL_EVENT_QUIT
	EVENT_TERMINATING                   EventType = C.SDL_EVENT_TERMINATING
	EVENT_LOW_MEMORY                    EventType = C.SDL_EVENT_LOW_MEMORY
	EVENT_WILL_ENTER_BACKGROUND         EventType = C.SDL_EVENT_WILL_ENTER_BACKGROUND
	EVENT_DID_ENTER_BACKGROUND          EventType = C.SDL_EVENT_DID_ENTER_BACKGROUND
	EVENT_WILL_ENTER_FOREGROUND         EventType = C.SDL_EVENT_WILL_ENTER_FOREGROUND
	EVENT_DID_ENTER_FOREGROUND          EventType = C.SDL_EVENT_DID_ENTER_FOREGROUND
	EVENT_LOCALE_CHANGED                EventType = C.SDL_EVENT_LOCALE_CHANGED
	EVENT_SYSTEM_THEME_CHANGED          EventType = C.SDL_EVENT_SYSTEM_THEME_CHANGED
	EVENT_DISPLAY_ORIENTATION           EventType = C.SDL_EVENT_DISPLAY_ORIENTATION
	EVENT_DISPLAY_ADDED                 EventType = C.SDL_EVENT_DISPLAY_ADDED
	EVENT_DISPLAY_REMOVED               EventType = C.SDL_EVENT_DISPLAY_REMOVED
	EVENT_DISPLAY_MOVED                 EventType = C.SDL_EVENT_DISPLAY_MOVED
	EVENT_DISPLAY_DESKTOP_MODE_CHANGED  EventType = C.SDL_EVENT_DISPLAY_DESKTOP_MODE_CHANGED
	EVENT_DISPLAY_CURRENT_MODE_CHANGED  EventType = C.SDL_EVENT_DISPLAY_CURRENT_MODE_CHANGED
	EVENT_DISPLAY_CONTENT_SCALE_CHANGED EventType = C.SDL_EVENT_DISPLAY_CONTENT_SCALE_CHANGED
	EVENT_DISPLAY_FIRST                 EventType = C.SDL_EVENT_DISPLAY_FIRST
	EVENT_DISPLAY_LAST                  EventType = C.SDL_EVENT_DISPLAY_LAST
	EVENT_WINDOW_SHOWN                  EventType = C.SDL_EVENT_WINDOW_SHOWN
	EVENT_WINDOW_HIDDEN                 EventType = C.SDL_EVENT_WINDOW_HIDDEN
	EVENT_WINDOW_EXPOSED                EventType = C.SDL_EVENT_WINDOW_EXPOSED
	EVENT_WINDOW_MOVED                  EventType = C.SDL_EVENT_WINDOW_MOVED
	EVENT_WINDOW_RESIZED                EventType = C.SDL_EVENT_WINDOW_RESIZED
	EVENT_WINDOW_PIXEL_SIZE_CHANGED     EventType = C.SDL_EVENT_WINDOW_PIXEL_SIZE_CHANGED
	EVENT_WINDOW_METAL_VIEW_RESIZED     EventType = C.SDL_EVENT_WINDOW_METAL_VIEW_RESIZED
	EVENT_WINDOW_MINIMIZED              EventType = C.SDL_EVENT_WINDOW_MINIMIZED
	EVENT_WINDOW_MAXIMIZED              EventType = C.SDL_EVENT_WINDOW_MAXIMIZED
	EVENT_WINDOW_RESTORED               EventType = C.SDL_EVENT_WINDOW_RESTORED
	EVENT_WINDOW_MOUSE_ENTER            EventType = C.SDL_EVENT_WINDOW_MOUSE_ENTER
	EVENT_WINDOW_MOUSE_LEAVE            EventType = C.SDL_EVENT_WINDOW_MOUSE_LEAVE
	EVENT_WINDOW_FOCUS_GAINED           EventType = C.SDL_EVENT_WINDOW_FOCUS_GAINED
	EVENT_WINDOW_FOCUS_LOST             EventType = C.SDL_EVENT_WINDOW_FOCUS_LOST
	EVENT_WINDOW_CLOSE_REQUESTED        EventType = C.SDL_EVENT_WINDOW_CLOSE_REQUESTED
	EVENT_WINDOW_HIT_TEST               EventType = C.SDL_EVENT_WINDOW_HIT_TEST
	EVENT_WINDOW_ICCPROF_CHANGED        EventType = C.SDL_EVENT_WINDOW_ICCPROF_CHANGED
	EVENT_WINDOW_DISPLAY_CHANGED        EventType = C.SDL_EVENT_WINDOW_DISPLAY_CHANGED
	EVENT_WINDOW_DISPLAY_SCALE_CHANGED  EventType = C.SDL_EVENT_WINDOW_DISPLAY_SCALE_CHANGED
	EVENT_WINDOW_SAFE_AREA_CHANGED      EventType = C.SDL_EVENT_WINDOW_SAFE_AREA_CHANGED
	EVENT_WINDOW_OCCLUDED               EventType = C.SDL_EVENT_WINDOW_OCCLUDED
	EVENT_WINDOW_ENTER_FULLSCREEN       EventType = C.SDL_EVENT_WINDOW_ENTER_FULLSCREEN
	EVENT_WINDOW_LEAVE_FULLSCREEN       EventType = C.SDL_EVENT_WINDOW_LEAVE_FULLSCREEN
	EVENT_WINDOW_DESTROYED              EventType = C.SDL_EVENT_WINDOW_DESTROYED
	EVENT_WINDOW_HDR_STATE_CHANGED      EventType = C.SDL_EVENT_WINDOW_HDR_STATE_CHANGED
	EVENT_WINDOW_FIRST                  EventType = C.SDL_EVENT_WINDOW_FIRST
	EVENT_WINDOW_LAST                   EventType = C.SDL_EVENT_WINDOW_LAST
	EVENT_KEY_DOWN                      EventType = C.SDL_EVENT_KEY_DOWN
	EVENT_KEY_UP                        EventType = C.SDL_EVENT_KEY_UP
	EVENT_TEXT_EDITING                  EventType = C.SDL_EVENT_TEXT_EDITING
	EVENT_TEXT_INPUT                    EventType = C.SDL_EVENT_TEXT_INPUT
	EVENT_KEYMAP_CHANGED                EventType = C.SDL_EVENT_KEYMAP_CHANGED
	EVENT_KEYBOARD_ADDED                EventType = C.SDL_EVENT_KEYBOARD_ADDED
	EVENT_KEYBOARD_REMOVED              EventType = C.SDL_EVENT_KEYBOARD_REMOVED
	EVENT_TEXT_EDITING_CANDIDATES       EventType = C.SDL_EVENT_TEXT_EDITING_CANDIDATES
	EVENT_MOUSE_MOTION                  EventType = C.SDL_EVENT_MOUSE_MOTION
	EVENT_MOUSE_BUTTON_DOWN             EventType = C.SDL_EVENT_MOUSE_BUTTON_DOWN
	EVENT_MOUSE_BUTTON_UP               EventType = C.SDL_EVENT_MOUSE_BUTTON_UP
	EVENT_MOUSE_WHEEL                   EventType = C.SDL_EVENT_MOUSE_WHEEL
	EVENT_MOUSE_ADDED                   EventType = C.SDL_EVENT_MOUSE_ADDED
	EVENT_MOUSE_REMOVED                 EventType = C.SDL_EVENT_MOUSE_REMOVED
	EVENT_JOYSTICK_AXIS_MOTION          EventType = C.SDL_EVENT_JOYSTICK_AXIS_MOTION
	EVENT_JOYSTICK_BALL_MOTION          EventType = C.SDL_EVENT_JOYSTICK_BALL_MOTION
	EVENT_JOYSTICK_HAT_MOTION           EventType = C.SDL_EVENT_JOYSTICK_HAT_MOTION
	EVENT_JOYSTICK_BUTTON_DOWN          EventType = C.SDL_EVENT_JOYSTICK_BUTTON_DOWN
	EVENT_JOYSTICK_BUTTON_UP            EventType = C.SDL_EVENT_JOYSTICK_BUTTON_UP
	EVENT_JOYSTICK_ADDED                EventType = C.SDL_EVENT_JOYSTICK_ADDED
	EVENT_JOYSTICK_REMOVED              EventType = C.SDL_EVENT_JOYSTICK_REMOVED
	EVENT_JOYSTICK_BATTERY_UPDATED      EventType = C.SDL_EVENT_JOYSTICK_BATTERY_UPDATED
	EVENT_JOYSTICK_UPDATE_COMPLETE      EventType = C.SDL_EVENT_JOYSTICK_UPDATE_COMPLETE
	EVENT_GAMEPAD_AXIS_MOTION           EventType = C.SDL_EVENT_GAMEPAD_AXIS_MOTION
	EVENT_GAMEPAD_BUTTON_DOWN           EventType = C.SDL_EVENT_GAMEPAD_BUTTON_DOWN
	EVENT_GAMEPAD_BUTTON_UP             EventType = C.SDL_EVENT_GAMEPAD_BUTTON_UP
	EVENT_GAMEPAD_ADDED                 EventType = C.SDL_EVENT_GAMEPAD_ADDED
	EVENT_GAMEPAD_REMOVED               EventType = C.SDL_EVENT_GAMEPAD_REMOVED
	EVENT_GAMEPAD_REMAPPED              EventType = C.SDL_EVENT_GAMEPAD_REMAPPED
	EVENT_GAMEPAD_TOUCHPAD_DOWN         EventType = C.SDL_EVENT_GAMEPAD_TOUCHPAD_DOWN
	EVENT_GAMEPAD_TOUCHPAD_MOTION       EventType = C.SDL_EVENT_GAMEPAD_TOUCHPAD_MOTION
	EVENT_GAMEPAD_TOUCHPAD_UP           EventType = C.SDL_EVENT_GAMEPAD_TOUCHPAD_UP
	EVENT_GAMEPAD_SENSOR_UPDATE         EventType = C.SDL_EVENT_GAMEPAD_SENSOR_UPDATE
	EVENT_GAMEPAD_UPDATE_COMPLETE       EventType = C.SDL_EVENT_GAMEPAD_UPDATE_COMPLETE
	EVENT_GAMEPAD_STEAM_HANDLE_UPDATED  EventType = C.SDL_EVENT_GAMEPAD_STEAM_HANDLE_UPDATED
	EVENT_FINGER_DOWN                   EventType = C.SDL_EVENT_FINGER_DOWN
	EVENT_FINGER_UP                     EventType = C.SDL_EVENT_FINGER_UP
	EVENT_FINGER_MOTION                 EventType = C.SDL_EVENT_FINGER_MOTION
	EVENT_FINGER_CANCELED               EventType = C.SDL_EVENT_FINGER_CANCELED
	EVENT_CLIPBOARD_UPDATE              EventType = C.SDL_EVENT_CLIPBOARD_UPDATE
	EVENT_DROP_FILE                     EventType = C.SDL_EVENT_DROP_FILE
	EVENT_DROP_TEXT                     EventType = C.SDL_EVENT_DROP_TEXT
	EVENT_DROP_BEGIN                    EventType = C.SDL_EVENT_DROP_BEGIN
	EVENT_DROP_COMPLETE                 EventType = C.SDL_EVENT_DROP_COMPLETE
	EVENT_DROP_POSITION                 EventType = C.SDL_EVENT_DROP_POSITION
	EVENT_AUDIO_DEVICE_ADDED            EventType = C.SDL_EVENT_AUDIO_DEVICE_ADDED
	EVENT_AUDIO_DEVICE_REMOVED          EventType = C.SDL_EVENT_AUDIO_DEVICE_REMOVED
	EVENT_AUDIO_DEVICE_FORMAT_CHANGED   EventType = C.SDL_EVENT_AUDIO_DEVICE_FORMAT_CHANGED
	EVENT_SENSOR_UPDATE                 EventType = C.SDL_EVENT_SENSOR_UPDATE
	EVENT_PEN_PROXIMITY_IN              EventType = C.SDL_EVENT_PEN_PROXIMITY_IN
	EVENT_PEN_PROXIMITY_OUT             EventType = C.SDL_EVENT_PEN_PROXIMITY_OUT
	EVENT_PEN_DOWN                      EventType = C.SDL_EVENT_PEN_DOWN
	EVENT_PEN_UP                        EventType = C.SDL_EVENT_PEN_UP
	EVENT_PEN_BUTTON_DOWN               EventType = C.SDL_EVENT_PEN_BUTTON_DOWN
	EVENT_PEN_BUTTON_UP                 EventType = C.SDL_EVENT_PEN_BUTTON_UP
	EVENT_PEN_MOTION                    EventType = C.SDL_EVENT_PEN_MOTION
	EVENT_PEN_AXIS                      EventType = C.SDL_EVENT_PEN_AXIS
	EVENT_CAMERA_DEVICE_ADDED           EventType = C.SDL_EVENT_CAMERA_DEVICE_ADDED
	EVENT_CAMERA_DEVICE_REMOVED         EventType = C.SDL_EVENT_CAMERA_DEVICE_REMOVED
	EVENT_CAMERA_DEVICE_APPROVED        EventType = C.SDL_EVENT_CAMERA_DEVICE_APPROVED
	EVENT_CAMERA_DEVICE_DENIED          EventType = C.SDL_EVENT_CAMERA_DEVICE_DENIED
	EVENT_RENDER_TARGETS_RESET          EventType = C.SDL_EVENT_RENDER_TARGETS_RESET
	EVENT_RENDER_DEVICE_RESET           EventType = C.SDL_EVENT_RENDER_DEVICE_RESET
	EVENT_RENDER_DEVICE_LOST            EventType = C.SDL_EVENT_RENDER_DEVICE_LOST
	EVENT_PRIVATE0                      EventType = C.SDL_EVENT_PRIVATE0
	EVENT_PRIVATE1                      EventType = C.SDL_EVENT_PRIVATE1
	EVENT_PRIVATE2                      EventType = C.SDL_EVENT_PRIVATE2
	EVENT_PRIVATE3                      EventType = C.SDL_EVENT_PRIVATE3
	EVENT_POLL_SENTINEL                 EventType = C.SDL_EVENT_POLL_SENTINEL
	EVENT_USER                          EventType = C.SDL_EVENT_USER
	EVENT_LAST                          EventType = C.SDL_EVENT_LAST
)

/*
 * Event
 */
type Event []byte

func (event Event) Type() EventType {
	return EventType(readUint32(event, C.offsetof_SDL_CommonEvent_type))
}

/*
 * WindowEvent
 */

type WindowEvent Event

func (event Event) Window() WindowEvent {
	return WindowEvent(event)
}

func (event WindowEvent) Type() EventType {
	return EventType(readUint32(event, C.offsetof_SDL_CommonEvent_type))
}

func (event WindowEvent) Timestamp() uint32 {
	return readUint32(event, C.offsetof_SDL_CommonEvent_timestamp)
}

func (event WindowEvent) WindowID() WindowID {
	return WindowID(readUint32(event, C.offsetof_SDL_WindowEvent_windowID))
}

func (event WindowEvent) Data1() int32 {
	return readInt32(event, C.offsetof_SDL_WindowEvent_data1)
}

func (event WindowEvent) Data2() int32 {
	return readInt32(event, C.offsetof_SDL_WindowEvent_data2)
}

/*
 * KeyboardEvent
 */

type KeyboardEvent Event

func (event Event) Key() KeyboardEvent {
	return KeyboardEvent(event)
}

func (event KeyboardEvent) Type() EventType {
	return EventType(readUint32(event, C.offsetof_SDL_CommonEvent_type))
}

func (event KeyboardEvent) Timestamp() uint32 {
	return readUint32(event, C.offsetof_SDL_CommonEvent_timestamp)
}

func (event KeyboardEvent) WindowID() WindowID {
	return WindowID(readUint32(event, C.offsetof_SDL_KeyboardEvent_windowID))
}

func (event KeyboardEvent) Which() KeyboardID {
	return KeyboardID(readUint32(event, C.offsetof_SDL_KeyboardEvent_which))
}

func (event KeyboardEvent) Scancode() Scancode {
	return Scancode(readUint32(event, C.offsetof_SDL_KeyboardEvent_scancode))
}

func (event KeyboardEvent) Key() Keycode {
	return Keycode(readInt32(event, C.offsetof_SDL_KeyboardEvent_key))
}

func (event KeyboardEvent) Mod() Keycode {
	return Keycode(readInt32(event, C.offsetof_SDL_KeyboardEvent_mod))
}

func (event KeyboardEvent) Raw() uint16 {
	return readUint16(event, C.offsetof_SDL_KeyboardEvent_raw)
}

func (event KeyboardEvent) Down() bool {
	return readBool(event, C.offsetof_SDL_KeyboardEvent_down)
}

func (event KeyboardEvent) Repeat() bool {
	return readBool(event, C.offsetof_SDL_KeyboardEvent_repeat)
}

/*
 * MouseMotionEvent
 */

type MouseMotionEvent Event

func (event Event) Motion() MouseMotionEvent {
	return MouseMotionEvent(event)
}

func (event MouseMotionEvent) Type() EventType {
	return EventType(readUint32(event, C.offsetof_SDL_CommonEvent_type))
}

func (event MouseMotionEvent) Timestamp() uint32 {
	return readUint32(event, C.offsetof_SDL_CommonEvent_timestamp)
}

func (event MouseMotionEvent) WindowID() WindowID {
	return WindowID(readUint32(event, C.offsetof_SDL_MouseMotionEvent_windowID))
}

func (event MouseMotionEvent) Which() uint32 {
	return readUint32(event, C.offsetof_SDL_MouseMotionEvent_which)
}

func (event MouseMotionEvent) State() uint32 {
	return readUint32(event, C.offsetof_SDL_MouseMotionEvent_state)
}

func (event MouseMotionEvent) X() float32 {
	return readFloat32(event, C.offsetof_SDL_MouseMotionEvent_x)
}

func (event MouseMotionEvent) Y() float32 {
	return readFloat32(event, C.offsetof_SDL_MouseMotionEvent_y)
}

func (event MouseMotionEvent) Xrel() float32 {
	return readFloat32(event, C.offsetof_SDL_MouseMotionEvent_xrel)
}

func (event MouseMotionEvent) Yrel() float32 {
	return readFloat32(event, C.offsetof_SDL_MouseMotionEvent_yrel)
}

/*
 * MouseButtonEvent
 */

type MouseButtonEvent Event

func (event Event) Button() MouseButtonEvent {
	return MouseButtonEvent(event)
}

func (event MouseButtonEvent) Type() EventType {
	return EventType(readUint32(event, C.offsetof_SDL_CommonEvent_type))
}

func (event MouseButtonEvent) Timestamp() uint32 {
	return readUint32(event, C.offsetof_SDL_CommonEvent_timestamp)
}

func (event MouseButtonEvent) WindowID() WindowID {
	return WindowID(readUint32(event, C.offsetof_SDL_MouseButtonEvent_windowID))
}

func (event MouseButtonEvent) Which() uint32 {
	return readUint32(event, C.offsetof_SDL_MouseButtonEvent_which)
}

func (event MouseButtonEvent) Button() uint8 {
	return readUint8(event, C.offsetof_SDL_MouseButtonEvent_button)
}

func (event MouseButtonEvent) Clicks() uint8 {
	return readUint8(event, C.offsetof_SDL_MouseButtonEvent_clicks)
}

func (event MouseButtonEvent) X() int32 {
	return readInt32(event, C.offsetof_SDL_MouseButtonEvent_x)
}

func (event MouseButtonEvent) Y() int32 {
	return readInt32(event, C.offsetof_SDL_MouseButtonEvent_y)
}

/*
 * Functions
 */

func cEventPointerToEvent(event *C.SDL_Event) Event {
	return C.GoBytes(unsafe.Pointer(event), C.sizeof_SDL_Event)
}

func PollEvent() Event {
	result := bool(C._SDL_PollEvent())

	if !result {
		return nil
	}

	return cEventPointerToEvent(&C._event)
}
