package sdl

//#include "SDL_go.h"
import "C"
import "unsafe"

type Renderer C.SDL_Renderer

const ()

func CreateWindowAndRenderer(
	width int,
	height int,
	window_flags WindowFlags,
	window **Window,
	renderer **Renderer,
) int {
	return (int)(C.SDL_CreateWindowAndRenderer(
		C.int(width),
		C.int(height),
		C.Uint32(window_flags),
		(**C.SDL_Window)(unsafe.Pointer(window)),
		(**C.SDL_Renderer)(unsafe.Pointer(renderer)),
	))
}

func DestroyRenderer(renderer *Renderer) {
	C.SDL_DestroyRenderer((*C.SDL_Renderer)(unsafe.Pointer(renderer)))
}

func RenderClear(renderer *Renderer) int {
	return (int)(C.SDL_RenderClear(
		(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
	))
}

func RenderPresent(renderer *Renderer) {
	C.SDL_RenderPresent(
		(*C.SDL_Renderer)(unsafe.Pointer(renderer)),
	)
}
