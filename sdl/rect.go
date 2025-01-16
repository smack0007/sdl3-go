package sdl

//#include "rect.h"
import "C"

const ()

type FPoint struct {
	X, Y float32
}

type FRect struct {
	X, Y, W, H float32
}

type Point struct {
	X, Y int32
}

type Rect struct {
	X, Y, W, H int32
}
