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

	result := (*sdl.Surface)(unsafe.Pointer(
		C.IMG_Load(c_file),
	))

	return result, sdl.PointerToError(unsafe.Pointer(result))
}

func Load_IO(src *sdl.IOStream, closeio bool) (*sdl.Surface, error) {
	result := (*sdl.Surface)(unsafe.Pointer(
		C.IMG_Load_IO(
			(*C.SDL_IOStream)(unsafe.Pointer(src)),
			(C.bool)(closeio),
		),
	))

	return result, sdl.PointerToError(unsafe.Pointer(result))
}
