package img

/*
#include <stdlib.h>
#include <SDL3_image/SDL_image.h>
*/
import "C"
import (
	"unsafe"

	"github.com/smack0007/sdl-go/sdl"
)

func Load(file string) (*sdl.Surface, error) {
	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	result := C.IMG_Load(c_file)

	if result == nil {
		return nil, sdl.GetError()
	}

	return (*sdl.Surface)(unsafe.Pointer(result)), sdl.PointerToError(result)
}
