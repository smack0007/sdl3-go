package sdl

//#include "SDL_go.h"
import "C"

const ()

func Delay(ms uint32) {
	C.SDL_Delay(C.Uint32(ms))
}
