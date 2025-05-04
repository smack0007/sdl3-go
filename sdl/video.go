package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL_video.h>
*/
import "C"

import (
	"unsafe"
)

type FlashOperation C.SDL_FlashOperation
type Window C.SDL_Window
type WindowFlags uint32
type WindowEventID uint8
type WindowID uint32

const (
	FLASH_CANCEL        FlashOperation = C.SDL_FLASH_CANCEL
	FLASH_BRIEFLY       FlashOperation = C.SDL_FLASH_BRIEFLY
	FLASH_UNTIL_FOCUSED FlashOperation = C.SDL_FLASH_UNTIL_FOCUSED

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

	WINDOWPOS_UNDEFINED int = C.SDL_WINDOWPOS_UNDEFINED
	WINDOWPOS_CENTERED  int = C.SDL_WINDOWPOS_CENTERED
)

func CreateWindow(
	title string,
	w int,
	h int,
	flags WindowFlags,
) (*Window, error) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	result := (*Window)(unsafe.Pointer(C.SDL_CreateWindow(
		c_title,
		C.int(w),
		C.int(h),
		C.SDL_WindowFlags(flags),
	)))

	return result, PointerToError(result)
}

func DestroyWindow(window *Window) {
	C.SDL_DestroyWindow((*C.SDL_Window)(unsafe.Pointer(window)))
}

func FlashWindow(window *Window, operation FlashOperation) error {
	return BoolToError(
		bool(
			C.SDL_FlashWindow(
				(*C.SDL_Window)(unsafe.Pointer(window)),
				C.SDL_FlashOperation(operation),
			),
		),
	)
}

func GetWindowSizeInPixels(window *Window) (w int, h int, err error) {
	err = BoolToError(
		bool(
			C.SDL_GetWindowSizeInPixels(
				(*C.SDL_Window)(unsafe.Pointer(window)),
				(*C.int)(unsafe.Pointer(&w)),
				(*C.int)(unsafe.Pointer(&h)),
			),
		),
	)

	return
}

func GetWindowSurface(window *Window) (*Surface, error) {
	result := (*Surface)(unsafe.Pointer(
		C.SDL_GetWindowSurface(
			(*C.SDL_Window)(unsafe.Pointer(window)),
		),
	))

	return result, PointerToError(result)
}

func HideWindow(window *Window) error {
	return BoolToError(
		bool(
			C.SDL_HideWindow(
				(*C.SDL_Window)(unsafe.Pointer(window)),
			),
		),
	)
}

func SetWindowTitle(window *Window, title string) error {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	return BoolToError(
		bool(
			C.SDL_SetWindowTitle(
				(*C.SDL_Window)(unsafe.Pointer(window)),
				c_title,
			),
		),
	)
}

func ShowWindow(window *Window) error {
	return BoolToError(
		bool(
			C.SDL_ShowWindow(
				(*C.SDL_Window)(unsafe.Pointer(window)),
			),
		),
	)
}

func UpdateWindowSurface(window *Window) error {
	return BoolToError(bool(C.SDL_UpdateWindowSurface((*C.SDL_Window)(unsafe.Pointer(window)))))
}
