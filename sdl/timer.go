package sdl

/*
#include <SDL3/SDL_timer.h>
*/
import "C"

const ()

func Delay(ms uint32) {
	C.SDL_Delay(C.Uint32(ms))
}

func GetTicks() uint64 {
	return uint64(C.SDL_GetTicks())
}
