package sdl

/*
#include "version.h"
static inline int _SDL_VERSIONNUM_MAJOR(int version) { return SDL_VERSIONNUM_MAJOR(version); }
static inline int _SDL_VERSIONNUM_MINOR(int version) { return SDL_VERSIONNUM_MINOR(version); }
static inline int _SDL_VERSIONNUM_MICRO(int version) { return SDL_VERSIONNUM_MICRO(version); }
*/
import "C"

const (
	REVISION string = C.SDL_REVISION
	VERSION  int    = C.SDL_VERSION
)

type Version struct {
	Major, Minor, Patch uint8
}

func GetRevision() string {
	return C.GoString(C.SDL_GetRevision())
}

func GetVersion() int {
	return int(C.SDL_GetVersion())
}

func VERSIONNUM_MAJOR(version int) int {
	return int(C._SDL_VERSIONNUM_MAJOR(C.int(version)))
}

func VERSIONNUM_MINOR(version int) int {
	return int(C._SDL_VERSIONNUM_MINOR(C.int(version)))
}

func VERSIONNUM_MICRO(version int) int {
	return int(C._SDL_VERSIONNUM_MICRO(C.int(version)))
}
