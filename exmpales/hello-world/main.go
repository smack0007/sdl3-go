package main

import (
	"fmt"
	"os"

	"github.com/smack0007/sdl_go/sdl"
)

func main() {
	if sdl.Init(sdl.INIT_VIDEO) != 0 {
		fmt.Println("Failed initialize SDL.")
		os.Exit(1)
	}
	defer sdl.Quit()

	window := sdl.CreateWindow("Hello World!", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 800, 600, sdl.WINDOW_SHOWN)

	if window == nil {
		fmt.Println("Failed to create window.")
		os.Exit(1)
	}

	defer sdl.DestroyWindow(window)

	sdl.GetWindowSurface(window)
	sdl.UpdateWindowSurface(window)

	done := false
	for !done {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case sdl.QuitEvent:
				done = true
			}
		}
	}
}
