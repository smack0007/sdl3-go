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
	WINDOW_FULLSCREEN          WindowFlags = C.SDL_WINDOW_FULLSCREEN
	WINDOW_OPENGL              WindowFlags = C.SDL_WINDOW_OPENGL
	WINDOW_OCCLUDED            WindowFlags = C.SDL_WINDOW_OCCLUDED
	WINDOW_HIDDEN              WindowFlags = C.SDL_WINDOW_HIDDEN
	WINDOW_BORDERLESS          WindowFlags = C.SDL_WINDOW_BORDERLESS
	WINDOW_RESIZABLE           WindowFlags = C.SDL_WINDOW_RESIZABLE
	WINDOW_MINIMIZED           WindowFlags = C.SDL_WINDOW_MINIMIZED
	WINDOW_MAXIMIZED           WindowFlags = C.SDL_WINDOW_MAXIMIZED
	WINDOW_MOUSE_GRABBED       WindowFlags = C.SDL_WINDOW_MOUSE_GRABBED
	WINDOW_INPUT_FOCUS         WindowFlags = C.SDL_WINDOW_INPUT_FOCUS
	WINDOW_MOUSE_FOCUS         WindowFlags = C.SDL_WINDOW_MOUSE_FOCUS
	WINDOW_EXTERNAL            WindowFlags = C.SDL_WINDOW_EXTERNAL
	WINDOW_MODAL               WindowFlags = C.SDL_WINDOW_MODAL
	WINDOW_HIGH_PIXEL_DENSITY  WindowFlags = C.SDL_WINDOW_HIGH_PIXEL_DENSITY
	WINDOW_MOUSE_CAPTURE       WindowFlags = C.SDL_WINDOW_MOUSE_CAPTURE
	WINDOW_MOUSE_RELATIVE_MODE WindowFlags = C.SDL_WINDOW_MOUSE_RELATIVE_MODE
	WINDOW_ALWAYS_ON_TOP       WindowFlags = C.SDL_WINDOW_ALWAYS_ON_TOP
	WINDOW_UTILITY             WindowFlags = C.SDL_WINDOW_UTILITY
	WINDOW_TOOLTIP             WindowFlags = C.SDL_WINDOW_TOOLTIP
	WINDOW_POPUP_MENU          WindowFlags = C.SDL_WINDOW_POPUP_MENU
	WINDOW_KEYBOARD_GRABBED    WindowFlags = C.SDL_WINDOW_KEYBOARD_GRABBED
	WINDOW_VULKAN              WindowFlags = C.SDL_WINDOW_VULKAN
	WINDOW_METAL               WindowFlags = C.SDL_WINDOW_METAL
	WINDOW_TRANSPARENT         WindowFlags = C.SDL_WINDOW_TRANSPARENT
	WINDOW_NOT_FOCUSABLE       WindowFlags = C.SDL_WINDOW_NOT_FOCUSABLE

	WINDOWPOS_CENTERED = int(C.SDL_WINDOWPOS_CENTERED)
)

func CreateWindow(
	title string,
	w int,
	h int,
	flags WindowFlags,
) (*Window, error) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	result := (*Window)(unsafe.Pointer(C.SDL_CreateWindow(
		cTitle,
		C.int(w),
		C.int(h),
		C.SDL_WindowFlags(flags),
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

func UpdateWindowSurface(window *Window) error {
	return mapErrorBool(bool(C.SDL_UpdateWindowSurface((*C.SDL_Window)(unsafe.Pointer(window)))))
}
