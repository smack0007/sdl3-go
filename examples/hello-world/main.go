package main

import (
	"os"

	"github.com/smack0007/sdl-go/sdl"
)

type AppState struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

func main() {
	os.Exit(sdl.EnterAppMainCallbacks(os.Args, AppInit, AppIterate, AppEvent, AppQuit))
}

func AppInit(argv []string) (sdl.AppResult, *AppState) {
	err := sdl.Init(sdl.INIT_VIDEO)

	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, "Failed initialize SDL: %s", err)
		return sdl.APP_FAILURE, nil
	}

	sdl.SetLogPriority(sdl.LOG_CATEGORY_APPLICATION, sdl.LOG_PRIORITY_DEBUG)

	window, renderer, err := sdl.CreateWindowAndRenderer("Hello World!", 800, 600, sdl.WINDOW_HIDDEN)

	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, "Failed to create window and renderer: %s", err)
		return sdl.APP_FAILURE, nil
	}

	sdl.SetWindowTitle(window, "Hello World!")

	sdl.LogInfo(sdl.LOG_CATEGORY_APPLICATION, "Initialized.")

	draw(renderer)
	sdl.ShowWindow(window)

	return sdl.APP_CONTINUE, &AppState{window: window, renderer: renderer}
}

func AppIterate(appState *AppState) sdl.AppResult {
	draw(appState.renderer)

	return sdl.APP_CONTINUE
}

func AppEvent(appState *AppState, event *sdl.Event) sdl.AppResult {
	switch event.Type() {
	case sdl.EVENT_QUIT:
		return sdl.APP_SUCCESS

	case sdl.EVENT_WINDOW_MINIMIZED:
		sdl.LogDebug(sdl.LOG_CATEGORY_APPLICATION, "Minimized")
		sdl.FlashWindow(appState.window, sdl.FLASH_UNTIL_FOCUSED)

	case sdl.EVENT_WINDOW_MOUSE_ENTER:
		sdl.LogDebug(sdl.LOG_CATEGORY_APPLICATION, "Enter")

	case sdl.EVENT_WINDOW_MOUSE_LEAVE:
		sdl.LogDebug(sdl.LOG_CATEGORY_APPLICATION, "Leave")

	case sdl.EVENT_KEY_DOWN:
		sdl.LogDebug(sdl.LOG_CATEGORY_APPLICATION, "Key Down %d %t", event.Key().Key(), event.Key().Down())

	case sdl.EVENT_KEY_UP:
		sdl.LogDebug(sdl.LOG_CATEGORY_APPLICATION, "Key Up %d %t", event.Key().Key(), event.Key().Down())

	case sdl.EVENT_MOUSE_MOTION:
		sdl.LogDebug(sdl.LOG_CATEGORY_APPLICATION, "Mouse Motion (%f, %f)", event.Motion().X(), event.Motion().Y())

	case sdl.EVENT_MOUSE_BUTTON_DOWN:
		sdl.LogDebug(sdl.LOG_CATEGORY_APPLICATION, "Mouse Button Down %d", event.Button().Button())

	case sdl.EVENT_MOUSE_BUTTON_UP:
		sdl.LogDebug(sdl.LOG_CATEGORY_APPLICATION, "Mouse Button Up %d", event.Button().Button())
	}

	return sdl.APP_CONTINUE
}

func AppQuit(appState *AppState, result sdl.AppResult) {
	sdl.LogInfo(sdl.LOG_CATEGORY_APPLICATION, "Shuting down...")
	sdl.DestroyRenderer(appState.renderer)
	sdl.DestroyWindow(appState.window)
	sdl.Quit()
}

func draw(renderer *sdl.Renderer) {
	sdl.SetRenderDrawColor(renderer, 100, 149, 237, 255)
	sdl.RenderClear(renderer)

	sdl.SetRenderDrawColor(renderer, 255, 0, 0, 255)

	points := []sdl.FPoint{
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
	sdl.RenderPoints(renderer, points, 36)

	sdl.RenderPresent(renderer)
}
