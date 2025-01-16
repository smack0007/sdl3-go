package sdl

//#include "SDL_go.h"
import "C"

type InitFlags uint32

const (
	INIT_AUDIO    InitFlags = C.SDL_INIT_AUDIO
	INIT_VIDEO    InitFlags = C.SDL_INIT_VIDEO
	INIT_JOYSTICK InitFlags = C.SDL_INIT_JOYSTICK
	INIT_HAPTIC   InitFlags = C.SDL_INIT_HAPTIC
	INIT_GAMEPAD  InitFlags = C.SDL_INIT_GAMEPAD
	INIT_EVENTS   InitFlags = C.SDL_INIT_EVENTS
	INIT_SENSOR   InitFlags = C.SDL_INIT_SENSOR
	INIT_CAMERA   InitFlags = C.SDL_INIT_CAMERA
)

func Init(flags InitFlags) error {
	return mapErrorBool(bool(C.SDL_Init(C.Uint32(flags))))
}

func Quit() {
	C.SDL_Quit()
}
