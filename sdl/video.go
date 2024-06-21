package sdl

//#include <SDL.h>
import "C"

import (
	"unsafe"
)

const (
	WINDOW_FULLSCREEN         = uint32(C.SDL_WINDOW_FULLSCREEN)
	WINDOW_OPENGL             = uint32(C.SDL_WINDOW_OPENGL)
	WINDOW_SHOWN              = uint32(C.SDL_WINDOW_SHOWN)
	WINDOW_HIDDEN             = uint32(C.SDL_WINDOW_HIDDEN)
	WINDOW_BORDERLESS         = uint32(C.SDL_WINDOW_BORDERLESS)
	WINDOW_RESIZABLE          = uint32(C.SDL_WINDOW_RESIZABLE)
	WINDOW_MINIMIZED          = uint32(C.SDL_WINDOW_MINIMIZED)
	WINDOW_MAXIMIZED          = uint32(C.SDL_WINDOW_MAXIMIZED)
	WINDOW_MOUSE_GRABBED      = uint32(C.SDL_WINDOW_MOUSE_GRABBED)
	WINDOW_INPUT_FOCUS        = uint32(C.SDL_WINDOW_INPUT_FOCUS)
	WINDOW_MOUSE_FOCUS        = uint32(C.SDL_WINDOW_MOUSE_FOCUS)
	WINDOW_FULLSCREEN_DESKTOP = uint32(C.SDL_WINDOW_FULLSCREEN_DESKTOP)
	WINDOW_FOREIGN            = uint32(C.SDL_WINDOW_FOREIGN)
	WINDOW_ALLOW_HIGHDPI      = uint32(C.SDL_WINDOW_ALLOW_HIGHDPI)
	WINDOW_MOUSE_CAPTURE      = uint32(C.SDL_WINDOW_MOUSE_CAPTURE)
	WINDOW_ALWAYS_ON_TOP      = uint32(C.SDL_WINDOW_ALWAYS_ON_TOP)
	WINDOW_SKIP_TASKBAR       = uint32(C.SDL_WINDOW_SKIP_TASKBAR)
	WINDOW_UTILITY            = uint32(C.SDL_WINDOW_UTILITY)
	WINDOW_TOOLTIP            = uint32(C.SDL_WINDOW_TOOLTIP)
	WINDOW_POPUP_MENU         = uint32(C.SDL_WINDOW_POPUP_MENU)
	WINDOW_KEYBOARD_GRABBED   = uint32(C.SDL_WINDOW_KEYBOARD_GRABBED)
	WINDOW_VULKAN             = uint32(C.SDL_WINDOW_VULKAN)
	WINDOW_METAL              = uint32(C.SDL_WINDOW_METAL)
	WINDOW_INPUT_GRABBED      = uint32(C.SDL_WINDOW_INPUT_GRABBED)

	WINDOWPOS_CENTERED = int(C.SDL_WINDOWPOS_CENTERED)
)

type Surface C.SDL_Surface

type Window C.SDL_Window

func CreateWindow(
	title string,
	x int,
	y int,
	w int,
	h int,
	flags uint32,
) *Window {
	return (*Window)(unsafe.Pointer(C.SDL_CreateWindow(
		C.CString(title),
		C.int(x),
		C.int(y),
		C.int(w),
		C.int(h),
		C.Uint32(flags),
	)))
}

func DestroyWindow(window *Window) {
	C.SDL_DestroyWindow((*C.SDL_Window)(unsafe.Pointer(window)))
}

func GetWindowSurface(window *Window) *Surface {
	return (*Surface)(unsafe.Pointer(C.SDL_GetWindowSurface((*C.SDL_Window)(unsafe.Pointer(window)))))
}

func UpdateWindowSurface(window *Window) int {
	return int(C.SDL_UpdateWindowSurface((*C.SDL_Window)(unsafe.Pointer(window))))
}
