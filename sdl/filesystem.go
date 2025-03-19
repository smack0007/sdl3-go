package sdl

/*
#include <SDL3/SDL_filesystem.h>
*/
import "C"

func GetBasePath() string {
	return C.GoString(
		C.SDL_GetBasePath(),
	)
}
