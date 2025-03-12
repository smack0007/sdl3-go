package sdl

//#include "render.h"
import "C"

import (
	"unsafe"
)

type Renderer C.SDL_Renderer

const ()

func CreateWindowAndRenderer(
	title string,
	width, height int,
	window_flags WindowFlags,
) (window *Window, renderer *Renderer, err error) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	err = mapErrorBool(
		bool(
			C.SDL_CreateWindowAndRenderer(
				c_title,
				C.int(width),
				C.int(height),
				C.SDL_WindowFlags(window_flags),
				(**C.SDL_Window)(unsafe.Pointer(&window)),
				(**C.SDL_Renderer)(unsafe.Pointer(&renderer)),
			),
		),
	)

	return
}

func DestroyRenderer(renderer *Renderer) {
	C.SDL_DestroyRenderer((*C.SDL_Renderer)(unsafe.Pointer(renderer)))
}

func GetRenderDrawColor(renderer *Renderer) (r, g, b, a uint8, err error) {
	err = mapErrorBool(
		bool(
			C.SDL_GetRenderDrawColor(
				(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
				(*C.Uint8)(unsafe.Pointer(&r)),
				(*C.Uint8)(unsafe.Pointer(&g)),
				(*C.Uint8)(unsafe.Pointer(&b)),
				(*C.Uint8)(unsafe.Pointer(&a)),
			),
		),
	)

	return
}

func RenderClear(renderer *Renderer) error {
	return mapErrorBool(
		bool(
			C.SDL_RenderClear(
				(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
			),
		),
	)
}

func RenderFillRect(renderer *Renderer, rect *FRect) error {
	return mapErrorBool((bool)(C.SDL_RenderFillRect(
		(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
		(*C.SDL_FRect)(unsafe.Pointer(rect)),
	)))
}

func RenderLine(renderer *Renderer, x1 float32, y1 float32, x2 float32, y2 float32) error {
	return mapErrorBool((bool)(C.SDL_RenderLine(
		(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
		C.float(x1),
		C.float(y1),
		C.float(x2),
		C.float(y2),
	)))
}

func RenderPoint(renderer *Renderer, x float32, y float32) error {
	return mapErrorBool((bool)(C.SDL_RenderPoint(
		(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
		C.float(x),
		C.float(y),
	)))
}

func RenderPoints(renderer *Renderer, points []FPoint, count int) error {
	return mapErrorBool((bool)(C.SDL_RenderPoints(
		(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
		(*C.SDL_FPoint)(unsafe.Pointer(&points[0])),
		C.int(count),
	)))
}

func RenderRect(renderer *Renderer, rect *FRect) error {
	return mapErrorBool(
		bool(
			C.SDL_RenderRect(
				(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
				(*C.SDL_FRect)(unsafe.Pointer(rect)),
			),
		),
	)
}

func GetRenderScale(renderer *Renderer) (scaleX float32, scaleY float32, err error) {
	err = mapErrorBool(
		bool(
			C.SDL_GetRenderScale(
				(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
				(*C.float)(unsafe.Pointer(&scaleX)),
				(*C.float)(unsafe.Pointer(&scaleY)),
			),
		),
	)

	return
}

func RenderPresent(renderer *Renderer) error {
	return mapErrorBool(
		bool(
			C.SDL_RenderPresent(
				(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
			),
		),
	)
}

func SetRenderDrawColor(renderer *Renderer, r uint8, g uint8, b uint8, a uint8) error {
	return mapErrorBool(
		bool(
			C.SDL_SetRenderDrawColor(
				(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
				C.Uint8(r),
				C.Uint8(g),
				C.Uint8(b),
				C.Uint8(a),
			),
		),
	)
}

func SetRenderDrawColorFloat(renderer *Renderer, r float32, g float32, b float32, a float32) error {
	return mapErrorBool(
		bool(
			C.SDL_SetRenderDrawColorFloat(
				(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
				C.float(r),
				C.float(g),
				C.float(b),
				C.float(a),
			),
		),
	)
}

func SetRenderScale(renderer *Renderer, scaleX float32, scaleY float32) error {
	return mapErrorBool(
		bool(
			C.SDL_SetRenderScale(
				(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
				C.float(scaleX),
				C.float(scaleY),
			),
		),
	)
}
