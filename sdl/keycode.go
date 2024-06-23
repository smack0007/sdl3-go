package sdl

//#include "SDL_go.h"
import "C"

type Keymod uint16

const (
	KMOD_NONE     Keymod = C.KMOD_NONE
	KMOD_LSHIFT   Keymod = C.KMOD_LSHIFT
	KMOD_RSHIFT   Keymod = C.KMOD_RSHIFT
	KMOD_LCTRL    Keymod = C.KMOD_LCTRL
	KMOD_RCTRL    Keymod = C.KMOD_RCTRL
	KMOD_LALT     Keymod = C.KMOD_LALT
	KMOD_RALT     Keymod = C.KMOD_RALT
	KMOD_LGUI     Keymod = C.KMOD_LGUI
	KMOD_RGUI     Keymod = C.KMOD_RGUI
	KMOD_NUM      Keymod = C.KMOD_NUM
	KMOD_CAPS     Keymod = C.KMOD_CAPS
	KMOD_MODE     Keymod = C.KMOD_MODE
	KMOD_SCROLL   Keymod = C.KMOD_SCROLL
	KMOD_CTRL     Keymod = C.KMOD_CTRL
	KMOD_SHIFT    Keymod = C.KMOD_SHIFT
	KMOD_ALT      Keymod = C.KMOD_ALT
	KMOD_GUI      Keymod = C.KMOD_GUI
	KMOD_RESERVED Keymod = C.KMOD_RESERVED
)
