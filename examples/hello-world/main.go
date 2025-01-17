package main

import (
	"os"

	SDL "github.com/smack0007/sdl-go/sdl"
)

func main() {
	os.Exit(run())
}

func run() int {
	err := SDL.Init(SDL.INIT_VIDEO)

	if err != nil {
		SDL.LogError(SDL.LOG_CATEGORY_APPLICATION, "Failed initialize SDL: %s", err)
		return 1
	}
	defer SDL.Quit()

	SDL.SetLogPriority(SDL.LOG_CATEGORY_APPLICATION, SDL.LOG_PRIORITY_DEBUG)

	window, renderer, err := SDL.CreateWindowAndRenderer("Hello World!", 800, 600, SDL.WINDOW_OCCLUDED)

	if err != nil {
		SDL.LogError(SDL.LOG_CATEGORY_APPLICATION, "Failed to create window and renderer: %s", err)
		return 1
	}

	defer SDL.DestroyWindow(window)
	defer SDL.DestroyRenderer(renderer)

	SDL.SetWindowTitle(window, "Hello World!")

	SDL.LogInfo(SDL.LOG_CATEGORY_APPLICATION, "Initialized.")

	done := false
	for !done {
		for event := SDL.PollEvent(); event != nil; event = SDL.PollEvent() {
			switch event.Type() {

			case SDL.EVENT_QUIT:
				done = true

			case SDL.EVENT_WINDOW_MOUSE_ENTER:
				SDL.LogDebug(SDL.LOG_CATEGORY_APPLICATION, "Enter")

			case SDL.EVENT_WINDOW_MOUSE_LEAVE:
				SDL.LogDebug(SDL.LOG_CATEGORY_APPLICATION, "Leave")

			case SDL.EVENT_KEY_DOWN:
				SDL.LogDebug(SDL.LOG_CATEGORY_APPLICATION, "Key Down %d %t", event.Key().Key(), event.Key().Down())

			case SDL.EVENT_KEY_UP:
				SDL.LogDebug(SDL.LOG_CATEGORY_APPLICATION, "Key Up %d %t", event.Key().Key(), event.Key().Down())

			case SDL.EVENT_MOUSE_MOTION:
				SDL.LogDebug(SDL.LOG_CATEGORY_APPLICATION, "Mouse Motion (%f, %f)", event.Motion().X(), event.Motion().Y())

			case SDL.EVENT_MOUSE_BUTTON_DOWN:
				SDL.LogDebug(SDL.LOG_CATEGORY_APPLICATION, "Mouse Button Down %d", event.Button().Button())

			case SDL.EVENT_MOUSE_BUTTON_UP:
				SDL.LogDebug(SDL.LOG_CATEGORY_APPLICATION, "Mouse Button Up %d", event.Button().Button())
			}
		}

		draw(renderer)
		SDL.Delay(1)
	}

	SDL.LogInfo(SDL.LOG_CATEGORY_APPLICATION, "Shuting down...")
	return 0
}

func draw(renderer *SDL.Renderer) {
	SDL.SetRenderDrawColor(renderer, 100, 149, 237, 255)
	SDL.RenderClear(renderer)

	SDL.SetRenderDrawColor(renderer, 255, 0, 0, 255)

	points := []SDL.FPoint{
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
	SDL.RenderPoints(renderer, points, 36)

	SDL.RenderPresent(renderer)
}
