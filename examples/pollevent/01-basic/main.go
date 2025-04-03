package main

import (
	"os"
	"runtime"

	"github.com/smack0007/sdl-go/sdl"
)

const (
	WINDOW_WIDTH  = 640
	WINDOW_HEIGHT = 480
	WINDOW_TITLE  = "SDL_PollEvent 01"
)

func main() {
	runtime.LockOSThread()
	os.Exit(run())
}

func run() int {
	err := sdl.Init(sdl.INIT_VIDEO)

	if err != nil {
		sdl.Log("Couldn't initialize SDL: %s", err)
		return 1
	}

	defer sdl.Quit()

	window, renderer, err := sdl.CreateWindowAndRenderer(WINDOW_TITLE, WINDOW_WIDTH, WINDOW_HEIGHT, sdl.WINDOW_OCCLUDED)

	if err != nil {
		sdl.Log("Couldn't create window/renderer: %s", err)
		return 1
	}

	defer sdl.DestroyWindow(window)
	defer sdl.DestroyRenderer(renderer)

	sdl.RenderPresent(renderer)

	shouldQuit := false

	var event sdl.Event
	for !shouldQuit {
		for sdl.PollEvent(&event) {
			switch event.Type() {
			case sdl.EVENT_QUIT:
				shouldQuit = true
			}
		}
	}

	return 0
}
