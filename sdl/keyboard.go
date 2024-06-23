package sdl

//#include "SDL_go.h"
import "C"

const ()

type Keycode int32

type Keysym struct {
	Scancode Scancode
	Sym      Keycode
	Mod      Keymod
}

func createKeysym(data []byte) Keysym {
	return Keysym{
		Scancode: Scancode(readUint32(data, C.offsetof_SDL_Keysym_scancode)),
		Sym:      Keycode(readInt32(data, C.offsetof_SDL_Keysym_sym)),
		Mod:      Keymod(readUint16(data, C.offsetof_SDL_Keysym_mod)),
	}
}
