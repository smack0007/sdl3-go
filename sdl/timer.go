package sdl

//#include "SDL_go.h"
import "C"

const ()

func Delay(ms uint32) {
	C.SDL_Delay(C.Uint32(ms))
}

func GetTicks() uint32 {
	return uint32(C.SDL_GetTicks())
}

func GetTicks64() uint64 {
	return uint64(C.SDL_GetTicks64())
}
