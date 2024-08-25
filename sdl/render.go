package sdl

//#include "SDL_go.h"
import "C"

import (
	"unsafe"
)

type Renderer C.SDL_Renderer

const ()

func CreateWindowAndRenderer(
	width, height int,
	window_flags WindowFlags,
) (window *Window, renderer *Renderer, err error) {
	errorCode := (int)(C.SDL_CreateWindowAndRenderer(
		C.int(width),
		C.int(height),
		C.Uint32(window_flags),
		(**C.SDL_Window)(unsafe.Pointer(&window)),
		(**C.SDL_Renderer)(unsafe.Pointer(&renderer)),
	))

	err = mapErrorCode(errorCode)

	return
}

func DestroyRenderer(renderer *Renderer) {
	C.SDL_DestroyRenderer((*C.SDL_Renderer)(unsafe.Pointer(renderer)))
}

func GetRenderDrawColor(renderer *Renderer) (r, g, b, a uint8, err error) {
	errorCode := int(C.SDL_GetRenderDrawColor(
		(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
		(*C.Uint8)(unsafe.Pointer(&r)),
		(*C.Uint8)(unsafe.Pointer(&g)),
		(*C.Uint8)(unsafe.Pointer(&b)),
		(*C.Uint8)(unsafe.Pointer(&a)),
	))

	err = mapErrorCode(errorCode)

	return
}

func RenderClear(renderer *Renderer) error {
	return mapErrorCode((int)(C.SDL_RenderClear(
		(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
	)))
}

func RenderDrawLine(renderer *Renderer, x1 int, y1 int, x2 int, y2 int) error {
	return mapErrorCode((int)(C.SDL_RenderDrawLine(
		(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
		C.int(x1),
		C.int(y1),
		C.int(x2),
		C.int(y2),
	)))
}

func RenderDrawPoint(renderer *Renderer, x int, y int) error {
	return mapErrorCode((int)(C.SDL_RenderDrawPoint(
		(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
		C.int(x),
		C.int(y),
	)))
}

func RenderDrawPoints(renderer *Renderer, points []Point, count int) error {
	return mapErrorCode((int)(C.SDL_RenderDrawPoints(
		(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
		(*C.SDL_Point)(unsafe.Pointer(&points[0])),
		C.int(count),
	)))
}

func RenderFillRect(renderer *Renderer, rect *Rect) error {
	return mapErrorCode((int)(C.SDL_RenderFillRect(
		(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
		(*C.SDL_Rect)(unsafe.Pointer(rect)),
	)))
}

func RenderGetScale(renderer *Renderer, scaleX *float32, scaleY *float32) {
	C.SDL_RenderGetScale(
		(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
		(*C.float)(unsafe.Pointer(scaleX)),
		(*C.float)(unsafe.Pointer(scaleY)),
	)
}

func RenderPresent(renderer *Renderer) {
	C.SDL_RenderPresent(
		(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
	)
}

func RenderSetScale(renderer *Renderer, scaleX float32, scaleY float32) error {
	return mapErrorCode(int(C.SDL_RenderSetScale(
		(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
		C.float(scaleX),
		C.float(scaleY),
	)))
}

func SetRenderDrawColor(renderer *Renderer, r uint8, g uint8, b uint8, a uint8) error {
	return mapErrorCode(int(C.SDL_SetRenderDrawColor(
		(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
		C.Uint8(r),
		C.Uint8(g),
		C.Uint8(b),
		C.Uint8(a),
	)))
}
