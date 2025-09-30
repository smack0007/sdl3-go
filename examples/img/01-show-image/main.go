package main

import (
	"os"
	"runtime"

	"github.com/smack0007/sdl3-go/img"
	"github.com/smack0007/sdl3-go/sdl"
)

const (
	WINDOW_WIDTH  = 640
	WINDOW_HEIGHT = 480
	WINDOW_TITLE  = "SDL_image 01"
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

	surface, err := img.Load("assets/sample.png")

	// io, err := sdl.IOFromFile("assets/sample.png", "r")

	// if err != nil {
	// 	sdl.Log("Couldn't load sample.png: %s", err)
	// 	return 1
	// }

	// surface, err := img.Load_IO(io, true)

	if err != nil {
		sdl.Log("Couldn't load sample.png: %s", err)
		return 1
	}

	texture, err := sdl.CreateTextureFromSurface(renderer, surface)

	if err != nil {
		sdl.Log("Couldn't create texture: %s", err)
		return 1
	}

	sdl.DestroySurface(surface)
	defer sdl.DestroyTexture(texture)

	sdl.RenderTexture(renderer, texture, nil, nil)
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
