package sdl

//#include "SDL_go.h"
import "C"

const (
	INIT_TIMER          = uint32(C.SDL_INIT_TIMER)
	INIT_AUDIO          = uint32(C.SDL_INIT_AUDIO)
	INIT_VIDEO          = uint32(C.SDL_INIT_VIDEO)
	INIT_JOYSTICK       = uint32(C.SDL_INIT_JOYSTICK)
	INIT_HAPTIC         = uint32(C.SDL_INIT_HAPTIC)
	INIT_GAMECONTROLLER = uint32(C.SDL_INIT_GAMECONTROLLER)
	INIT_EVENTS         = uint32(C.SDL_INIT_EVENTS)
	INIT_NOPARACHUTE    = uint32(C.SDL_INIT_NOPARACHUTE)
	INIT_SENSOR         = uint32(C.SDL_INIT_SENSOR)
	INIT_EVERYTHING     = uint32(C.SDL_INIT_EVERYTHING)
)

func Init(flags uint32) error {
	return mapErrorCode(int(C.SDL_Init(C.Uint32(flags))))
}

func Quit() {
	C.SDL_Quit()
}
