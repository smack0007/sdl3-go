package main

import (
	"os"

	SDL "github.com/smack0007/sdl-go/sdl"
)

func main() {
	err := SDL.Init(SDL.INIT_VIDEO)

	if err != nil {
		SDL.LogError(SDL.LOG_CATEGORY_APPLICATION, "Failed initialize SDL: %s", err)
		os.Exit(1)
	}
	defer SDL.Quit()

	SDL.LogSetPriority(SDL.LOG_CATEGORY_APPLICATION, SDL.LOG_PRIORITY_DEBUG)

	window, renderer, err := SDL.CreateWindowAndRenderer(800, 600, SDL.WINDOW_SHOWN)

	if err != nil {
		SDL.LogError(SDL.LOG_CATEGORY_APPLICATION, "Failed to create window and renderer: %s", err)
		os.Exit(1)
	}

	defer SDL.DestroyWindow(window)
	defer SDL.DestroyRenderer(renderer)

	SDL.SetWindowTitle(window, "Hello World!")

	SDL.LogInfo(SDL.LOG_CATEGORY_APPLICATION, "Initialized.")

	done := false
	for !done {
		for event := SDL.PollEvent(); event != nil; event = SDL.PollEvent() {
			switch event.Type() {

			case SDL.QUIT:
				done = true

			case SDL.WINDOWEVENT:
				switch event.Window().Event() {
				case SDL.WINDOWEVENT_ENTER:
					SDL.LogDebug(SDL.LOG_CATEGORY_APPLICATION, "Enter")

				case SDL.WINDOWEVENT_LEAVE:
					SDL.LogDebug(SDL.LOG_CATEGORY_APPLICATION, "Leave")
				}

			case SDL.KEYDOWN:
				SDL.LogDebug(SDL.LOG_CATEGORY_APPLICATION, "Key Down %d", event.Key().Keysym())

			case SDL.KEYUP:
				SDL.LogDebug(SDL.LOG_CATEGORY_APPLICATION, "Key Up %d", event.Key().Keysym())

			case SDL.MOUSEMOTION:
				SDL.LogDebug(SDL.LOG_CATEGORY_APPLICATION, "Mouse Motion (%d, %d)", event.Motion().X(), event.Motion().Y())

			case SDL.MOUSEBUTTONDOWN:
				SDL.LogDebug(SDL.LOG_CATEGORY_APPLICATION, "Mouse Button Down %d", event.Button().Button())

			case SDL.MOUSEBUTTONUP:
				SDL.LogDebug(SDL.LOG_CATEGORY_APPLICATION, "Mouse Button Up %d", event.Button().Button())
			}
		}

		draw(renderer)
		SDL.Delay(1)
	}

	SDL.LogInfo(SDL.LOG_CATEGORY_APPLICATION, "Shuting down...")
}

func draw(renderer *SDL.Renderer) {
	SDL.SetRenderDrawColor(renderer, 100, 149, 237, 255)
	SDL.RenderClear(renderer)

	SDL.SetRenderDrawColor(renderer, 255, 0, 0, 255)

	points := []SDL.Point{
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 2, Y: 0},
		{X: 3, Y: 0},
		{X: 4, Y: 0},
		{X: 5, Y: 0},

		{X: 0, Y: 1},
		{X: 1, Y: 1},
		{X: 2, Y: 1},
		{X: 3, Y: 1},
		{X: 4, Y: 1},
		{X: 5, Y: 1},

		{X: 0, Y: 2},
		{X: 1, Y: 2},
		{X: 2, Y: 2},
		{X: 3, Y: 2},
		{X: 4, Y: 2},
		{X: 5, Y: 2},

		{X: 0, Y: 3},
		{X: 1, Y: 3},
		{X: 2, Y: 3},
		{X: 3, Y: 3},
		{X: 4, Y: 3},
		{X: 5, Y: 3},

		{X: 0, Y: 4},
		{X: 1, Y: 4},
		{X: 2, Y: 4},
		{X: 3, Y: 4},
		{X: 4, Y: 4},
		{X: 5, Y: 4},

		{X: 0, Y: 5},
		{X: 1, Y: 5},
		{X: 2, Y: 5},
		{X: 3, Y: 5},
		{X: 4, Y: 5},
		{X: 5, Y: 5},
	}
	SDL.RenderDrawPoints(renderer, points, 36)

	SDL.RenderPresent(renderer)
}
