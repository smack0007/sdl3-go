package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL_surface.h>
*/
import "C"
import (
	"unsafe"
)

type FlipMode C.SDL_FlipMode
type ScaleMode C.SDL_ScaleMode
type SurfaceFlags uint32

const (
	FLIP_NONE       FlipMode = C.SDL_FLIP_NONE
	FLIP_HORIZONTAL FlipMode = C.SDL_FLIP_HORIZONTAL
	FILP_VERTICAL   FlipMode = C.SDL_FLIP_VERTICAL

	PROP_SURFACE_SDR_WHITE_POINT_FLOAT   string = C.SDL_PROP_SURFACE_SDR_WHITE_POINT_FLOAT
	PROP_SURFACE_HDR_HEADROOM_FLOAT      string = C.SDL_PROP_SURFACE_HDR_HEADROOM_FLOAT
	PROP_SURFACE_TONEMAP_OPERATOR_STRING string = C.SDL_PROP_SURFACE_TONEMAP_OPERATOR_STRING

	SCALEMODE_NEAREST ScaleMode = C.SDL_SCALEMODE_NEAREST
	SCALEMODE_LINEAR  ScaleMode = C.SDL_SCALEMODE_LINEAR

	SURFACE_PREALLOCATED SurfaceFlags = C.SDL_SURFACE_PREALLOCATED
	SURFACE_LOCK_NEEDED  SurfaceFlags = C.SDL_SURFACE_LOCK_NEEDED
	SURFACE_LOCKED       SurfaceFlags = C.SDL_SURFACE_LOCKED
	SURFACE_SIMD_ALIGNED SurfaceFlags = C.SDL_SURFACE_SIMD_ALIGNED
)

//
// Structs
//

type Surface C.SDL_Surface

func (surface *Surface) Flags() SurfaceFlags { return SurfaceFlags(surface.flags) }

func (surface *Surface) Format() PixelFormat { return PixelFormat(surface.format) }

func (surface *Surface) W() int { return int(surface.w) }

func (surface *Surface) H() int { return int(surface.h) }

func (surface *Surface) Pitch() int { return int(surface.pitch) }

func (surface *Surface) Pixels() unsafe.Pointer { return unsafe.Pointer(surface.pixels) }

func (surface *Surface) Refcount() int { return int(surface.refcount) }

//
// Functions
//

func CreateSurface(width int, height int, format PixelFormat) (*Surface, error) {
	result := (*Surface)(
		C.SDL_CreateSurface(
			(C.int)(width),
			(C.int)(height),
			(C.SDL_PixelFormat)(format),
		),
	)

	return result, PointerToError(unsafe.Pointer(result))
}

func DestroySurface(surface *Surface) {
	C.SDL_DestroySurface(
		(*C.SDL_Surface)(surface),
	)
}

func FillSurfaceRect(dst *Surface, rect *Rect, color uint32) error {
	return BoolToError(
		bool(
			C.SDL_FillSurfaceRect(
				(*C.SDL_Surface)(dst),
				(*C.SDL_Rect)(unsafe.Pointer(rect)),
				(C.Uint32)(color),
			),
		),
	)
}

func LoadBMP(file string) (*Surface, error) {
	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	result := C.SDL_LoadBMP(c_file)

	if result == nil {
		return nil, GetError()
	}

	return (*Surface)(result), PointerToError(unsafe.Pointer(result))
}

func LoadBMP_IO(src *IOStream, closeio bool) (*Surface, error) {
	result := (*Surface)(unsafe.Pointer(
		C.SDL_LoadBMP_IO(
			(*C.SDL_IOStream)(unsafe.Pointer(src)),
			(C.bool)(closeio),
		),
	))

	return result, PointerToError(unsafe.Pointer(result))
}
