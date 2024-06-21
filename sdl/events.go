package sdl

//#include <SDL.h>
//SDL_Event _event;
import "C"

import (
	"unsafe"
)

const (
	QUIT uint32 = C.SDL_QUIT
)

type Event interface {
	getType() uint32
}

type CommonEvent struct {
	Type      uint32
	Timestamp int
}

func (e CommonEvent) getType() uint32 {
	return e.Type
}

type QuitEvent struct {
	Type      uint32
	Timestamp int
}

func (e QuitEvent) getType() uint32 {
	return e.Type
}

func PollEvent() Event {
	result := int(C.SDL_PollEvent(&C._event))

	if result == 0 {
		return nil
	}

	event := (*C.SDL_CommonEvent)(unsafe.Pointer(&C._event))

	switch event._type {
	case C.SDL_QUIT:
		return QuitEvent{
			Type:      uint32(event._type),
			Timestamp: int(event.timestamp),
		}
	default:
		return CommonEvent{
			Type:      uint32(event._type),
			Timestamp: int(event.timestamp),
		}
	}
}
