package sdl

//#include "SDL_go.h"
import "C"

import (
	"errors"
)

func mapErrorCode(code int) error {
	if code >= 0 {
		return nil
	}

	return GetError()
}

func mapErrorPointer(pointer interface{}) error {
	if pointer != nil {
		return nil
	}

	return GetError()
}

func GetError() error {
	return errors.New(C.GoString(C.SDL_GetError()))
}
