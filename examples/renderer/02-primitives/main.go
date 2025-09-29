/*
 * This example creates an SDL window and renderer, and then draws some lines,
 * rectangles and points to it every frame.
 *
 * This code is public domain. Feel free to use it for any purpose!
 */

package main

import (
	"math/rand/v2"
	"os"

	"github.com/smack0007/sdl3-go/sdl"
)

type AppState struct {
	/* We will use this renderer to draw into this window every frame. */
	window   *sdl.Window
	renderer *sdl.Renderer
	points   []sdl.FPoint
}

func main() {
	/* use the callbacks */
	os.Exit(sdl.EnterAppMainCallbacks(os.Args, AppInit, AppIterate, AppEvent, AppQuit))
}

/* This function runs once at startup. */
func AppInit(argv []string) (sdl.AppResult, *AppState) {
	sdl.SetAppMetadata("Example Renderer Primitives", "1.0", "com.example.renderer-primitives")

	err := sdl.Init(sdl.INIT_VIDEO)

	if err != nil {
		sdl.Log("Couldn't initialize SDL: %s", err)
		return sdl.APP_FAILURE, nil
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/renderer/primitives", 640, 480, 0)

	if err != nil {
		sdl.Log("Couldn't create window/renderer: %s", err)
		return sdl.APP_FAILURE, nil
	}

	points := make([]sdl.FPoint, 500)
	for i := 0; i < len(points); i += 1 {
		points[i].X = (rand.Float32() * 440.0) + 100.0
		points[i].Y = (rand.Float32() * 280.0) + 100.0
	}

	/* carry on with the program! */
	return sdl.APP_CONTINUE, &AppState{
		window:   window,
		renderer: renderer,
		points:   points,
	}
}

/* This function runs when a new event (mouse input, keypresses, etc) occurs. */
func AppEvent(appState *AppState, event *sdl.Event) sdl.AppResult {
	if event.Type() == sdl.EVENT_QUIT {
		return sdl.APP_SUCCESS /* end the program, reporting success to the OS. */
	}
	return sdl.APP_CONTINUE
}

/* This function runs once per frame, and is the heart of the program. */
func AppIterate(appState *AppState) sdl.AppResult {
	var rect sdl.FRect

	/* as you can see from this, rendering draws over whatever was drawn before it. */
	sdl.SetRenderDrawColor(appState.renderer, 33, 33, 33, sdl.ALPHA_OPAQUE) /* dark gray, full alpha */
	sdl.RenderClear(appState.renderer)                                      /* start with a blank canvas. */

	/* draw a filled rectangle in the middle of the canvas. */
	sdl.SetRenderDrawColor(appState.renderer, 0, 0, 255, sdl.ALPHA_OPAQUE) /* blue, full alpha */
	rect.X = 100
	rect.Y = 100
	rect.W = 440
	rect.H = 280
	sdl.RenderFillRect(appState.renderer, &rect)

	/* draw some points across the canvas. */
	sdl.SetRenderDrawColor(appState.renderer, 255, 0, 0, sdl.ALPHA_OPAQUE) /* red, full alpha */
	sdl.RenderPoints(appState.renderer, appState.points, len(appState.points))

	/* draw a unfilled rectangle in-set a little bit. */
	sdl.SetRenderDrawColor(appState.renderer, 0, 255, 0, sdl.ALPHA_OPAQUE) /* green, full alpha */
	rect.X += 30
	rect.Y += 30
	rect.W -= 60
	rect.H -= 60
	sdl.RenderRect(appState.renderer, &rect)

	/* draw two lines in an X across the whole canvas. */
	sdl.SetRenderDrawColor(appState.renderer, 255, 255, 0, sdl.ALPHA_OPAQUE) /* yellow, full alpha */
	sdl.RenderLine(appState.renderer, 0, 0, 640, 480)
	sdl.RenderLine(appState.renderer, 0, 480, 640, 0)

	sdl.RenderPresent(appState.renderer) /* put it all on the screen! */

	return sdl.APP_CONTINUE /* carry on with the program! */
}

/* This function runs once at shutdown. */
func AppQuit(appState *AppState, result sdl.AppResult) {
	/* SDL will clean up the window/renderer for us. */
}
