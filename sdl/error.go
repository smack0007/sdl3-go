package sdl

/*
#include <SDL3/SDL_error.h>
*/
import "C"

import (
	"errors"
)

func mapErrorBool(result bool) error {
	if result {
		return nil
	}

	return GetError()
}

func mapErrorPointer(pointer any) error {
	if pointer != nil {
		return nil
	}

	return GetError()
}

func GetError() error {
	return errors.New(
		C.GoString(
			C.SDL_GetError(),
		),
	)
}
