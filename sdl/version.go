package sdl

//#include "SDL_go.h"
import "C"
import "unsafe"

const ()

type Version struct {
	Major, Minor, Patch uint8
}

func GetRevision() string {
	return C.GoString(C.SDL_GetRevision())
}

func GetVersion(ver *Version) {
	C.SDL_GetVersion((*C.SDL_version)(unsafe.Pointer(ver)))
}
