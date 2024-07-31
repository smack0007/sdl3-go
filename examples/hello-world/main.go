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

	var window *sdl.Window
	var renderer *sdl.Renderer
	result := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN, &window, &renderer)

	if result != 0 {
		fmt.Println("Failed to create window and renderer.")
		os.Exit(1)
	}

	defer sdl.DestroyWindow(window)
	defer sdl.DestroyRenderer(renderer)

	sdl.SetWindowTitle(window, "Hello World!")

	fmt.Println("Initialized.")

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
				fmt.Println("Key Down", event.Key.Keysym())

			case sdl.KEYUP:
				fmt.Println("Key Up", event.Key.Keysym())

			case sdl.MOUSEMOTION:
				fmt.Println("Mouse Motion", event.Motion.X(), event.Motion.Y())

			case sdl.MOUSEBUTTONDOWN:
				fmt.Println("Mouse Button Down", event.Button.Button())

			case sdl.MOUSEBUTTONUP:
				fmt.Println("Mouse Button Up", event.Button.Button())
			}
		}

		draw(renderer)
		sdl.Delay(1)
	}

	fmt.Println("Shuting down...")
}

func draw(renderer *sdl.Renderer) {
	sdl.SetRenderDrawColor(renderer, 100, 149, 237, 255)
	sdl.RenderClear(renderer)

	sdl.SetRenderDrawColor(renderer, 255, 0, 0, 255)

	points := []sdl.Point{
		sdl.Point{X: 0, Y: 0},
		sdl.Point{X: 1, Y: 0},
		sdl.Point{X: 2, Y: 0},
		sdl.Point{X: 3, Y: 0},
		sdl.Point{X: 4, Y: 0},
		sdl.Point{X: 5, Y: 0},

		sdl.Point{X: 0, Y: 1},
		sdl.Point{X: 1, Y: 1},
		sdl.Point{X: 2, Y: 1},
		sdl.Point{X: 3, Y: 1},
		sdl.Point{X: 4, Y: 1},
		sdl.Point{X: 5, Y: 1},

		sdl.Point{X: 0, Y: 2},
		sdl.Point{X: 1, Y: 2},
		sdl.Point{X: 2, Y: 2},
		sdl.Point{X: 3, Y: 2},
		sdl.Point{X: 4, Y: 2},
		sdl.Point{X: 5, Y: 2},

		sdl.Point{X: 0, Y: 3},
		sdl.Point{X: 1, Y: 3},
		sdl.Point{X: 2, Y: 3},
		sdl.Point{X: 3, Y: 3},
		sdl.Point{X: 4, Y: 3},
		sdl.Point{X: 5, Y: 3},

		sdl.Point{X: 0, Y: 4},
		sdl.Point{X: 1, Y: 4},
		sdl.Point{X: 2, Y: 4},
		sdl.Point{X: 3, Y: 4},
		sdl.Point{X: 4, Y: 4},
		sdl.Point{X: 5, Y: 4},

		sdl.Point{X: 0, Y: 5},
		sdl.Point{X: 1, Y: 5},
		sdl.Point{X: 2, Y: 5},
		sdl.Point{X: 3, Y: 5},
		sdl.Point{X: 4, Y: 5},
		sdl.Point{X: 5, Y: 5},
	}
	sdl.RenderDrawPoints(renderer, points, 36)

	sdl.RenderPresent(renderer)
}
