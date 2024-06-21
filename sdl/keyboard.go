package sdl

//#include <SDL.h>
import "C"

const ()

type Keycode int32

type Keysym struct {
	Scancode Scancode
	Sym      Keycode
	Mod      Keymod
}
