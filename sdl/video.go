package sdl

//#include "SDL_go.h"
import "C"

import (
	"unsafe"
)

type Surface C.SDL_Surface

type Window C.SDL_Window
type WindowFlags uint32
type WindowEventID uint8

const (
	WINDOW_FULLSCREEN         WindowFlags = C.SDL_WINDOW_FULLSCREEN
	WINDOW_OPENGL             WindowFlags = C.SDL_WINDOW_OPENGL
	WINDOW_SHOWN              WindowFlags = C.SDL_WINDOW_SHOWN
	WINDOW_HIDDEN             WindowFlags = C.SDL_WINDOW_HIDDEN
	WINDOW_BORDERLESS         WindowFlags = C.SDL_WINDOW_BORDERLESS
	WINDOW_RESIZABLE          WindowFlags = C.SDL_WINDOW_RESIZABLE
	WINDOW_MINIMIZED          WindowFlags = C.SDL_WINDOW_MINIMIZED
	WINDOW_MAXIMIZED          WindowFlags = C.SDL_WINDOW_MAXIMIZED
	WINDOW_MOUSE_GRABBED      WindowFlags = C.SDL_WINDOW_MOUSE_GRABBED
	WINDOW_INPUT_FOCUS        WindowFlags = C.SDL_WINDOW_INPUT_FOCUS
	WINDOW_MOUSE_FOCUS        WindowFlags = C.SDL_WINDOW_MOUSE_FOCUS
	WINDOW_FULLSCREEN_DESKTOP WindowFlags = C.SDL_WINDOW_FULLSCREEN_DESKTOP
	WINDOW_FOREIGN            WindowFlags = C.SDL_WINDOW_FOREIGN
	WINDOW_ALLOW_HIGHDPI      WindowFlags = C.SDL_WINDOW_ALLOW_HIGHDPI
	WINDOW_MOUSE_CAPTURE      WindowFlags = C.SDL_WINDOW_MOUSE_CAPTURE
	WINDOW_ALWAYS_ON_TOP      WindowFlags = C.SDL_WINDOW_ALWAYS_ON_TOP
	WINDOW_SKIP_TASKBAR       WindowFlags = C.SDL_WINDOW_SKIP_TASKBAR
	WINDOW_UTILITY            WindowFlags = C.SDL_WINDOW_UTILITY
	WINDOW_TOOLTIP            WindowFlags = C.SDL_WINDOW_TOOLTIP
	WINDOW_POPUP_MENU         WindowFlags = C.SDL_WINDOW_POPUP_MENU
	WINDOW_KEYBOARD_GRABBED   WindowFlags = C.SDL_WINDOW_KEYBOARD_GRABBED
	WINDOW_VULKAN             WindowFlags = C.SDL_WINDOW_VULKAN
	WINDOW_METAL              WindowFlags = C.SDL_WINDOW_METAL
	WINDOW_INPUT_GRABBED      WindowFlags = C.SDL_WINDOW_INPUT_GRABBED

	WINDOWEVENT_NONE            WindowEventID = C.SDL_WINDOWEVENT_NONE
	WINDOWEVENT_SHOWN           WindowEventID = C.SDL_WINDOWEVENT_SHOWN
	WINDOWEVENT_HIDDEN          WindowEventID = C.SDL_WINDOWEVENT_HIDDEN
	WINDOWEVENT_EXPOSED         WindowEventID = C.SDL_WINDOWEVENT_EXPOSED
	WINDOWEVENT_MOVED           WindowEventID = C.SDL_WINDOWEVENT_MOVED
	WINDOWEVENT_RESIZED         WindowEventID = C.SDL_WINDOWEVENT_RESIZED
	WINDOWEVENT_SIZE_CHANGED    WindowEventID = C.SDL_WINDOWEVENT_SIZE_CHANGED
	WINDOWEVENT_MINIMIZED       WindowEventID = C.SDL_WINDOWEVENT_MINIMIZED
	WINDOWEVENT_MAXIMIZED       WindowEventID = C.SDL_WINDOWEVENT_MAXIMIZED
	WINDOWEVENT_RESTORED        WindowEventID = C.SDL_WINDOWEVENT_RESTORED
	WINDOWEVENT_ENTER           WindowEventID = C.SDL_WINDOWEVENT_ENTER
	WINDOWEVENT_LEAVE           WindowEventID = C.SDL_WINDOWEVENT_LEAVE
	WINDOWEVENT_FOCUS_GAINED    WindowEventID = C.SDL_WINDOWEVENT_FOCUS_GAINED
	WINDOWEVENT_FOCUS_LOST      WindowEventID = C.SDL_WINDOWEVENT_FOCUS_LOST
	WINDOWEVENT_CLOSE           WindowEventID = C.SDL_WINDOWEVENT_CLOSE
	WINDOWEVENT_TAKE_FOCUS      WindowEventID = C.SDL_WINDOWEVENT_TAKE_FOCUS
	WINDOWEVENT_HIT_TEST        WindowEventID = C.SDL_WINDOWEVENT_HIT_TEST
	WINDOWEVENT_ICCPROF_CHANGED WindowEventID = C.SDL_WINDOWEVENT_ICCPROF_CHANGED
	WINDOWEVENT_DISPLAY_CHANGED WindowEventID = C.SDL_WINDOWEVENT_DISPLAY_CHANGED

	WINDOWPOS_CENTERED = int(C.SDL_WINDOWPOS_CENTERED)
)

func CreateWindow(
	title string,
	x int,
	y int,
	w int,
	h int,
	flags WindowFlags,
) (*Window, error) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	result := (*Window)(unsafe.Pointer(C.SDL_CreateWindow(
		cTitle,
		C.int(x),
		C.int(y),
		C.int(w),
		C.int(h),
		C.Uint32(flags),
	)))

	return result, mapErrorPointer(result)
}

func DestroyWindow(window *Window) {
	C.SDL_DestroyWindow((*C.SDL_Window)(unsafe.Pointer(window)))
}

func GetWindowSurface(window *Window) (*Surface, error) {
	result := (*Surface)(unsafe.Pointer(C.SDL_GetWindowSurface((*C.SDL_Window)(unsafe.Pointer(window)))))
	return result, mapErrorPointer(result)
}

func SetWindowTitle(window *Window, title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	C.SDL_SetWindowTitle(
		(*C.SDL_Window)(unsafe.Pointer(window)),
		cTitle,
	)
}

func UpdateWindowSurface(window *Window) int {
	return int(C.SDL_UpdateWindowSurface((*C.SDL_Window)(unsafe.Pointer(window))))
}
