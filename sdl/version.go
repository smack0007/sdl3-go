package sdl

//#include "SDL_go.h"
import "C"

const ()

type Version struct {
	Major, Minor, Patch uint8
}

func GetRevision() string {
	return C.GoString(C.SDL_GetRevision())
}

func GetVersion() int {
	return int(C.SDL_GetVersion())
}
