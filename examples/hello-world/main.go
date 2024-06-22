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

	event := sdl.Event{}
	done := false
	for !done {
		for sdl.PollEvent(&event) > 0 {
			switch event.Type() {

			case sdl.QUIT:
				done = true

			case sdl.WINDOWEVENT:
				switch event.Window.Event() {
				case sdl.WINDOWEVENT_ENTER:
					fmt.Println("Enter")

				case sdl.WINDOWEVENT_LEAVE:
					fmt.Println("Leave")
				}

			case sdl.KEYDOWN:
				fmt.Println(event.Key.Keysym())
			}
		}
	}
}
