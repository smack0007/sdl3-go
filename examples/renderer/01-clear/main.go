/*
 * This example code creates an SDL window and renderer, and then clears the
 * window to a different color every frame, so you'll effectively get a window
 * that's smoothly fading between colors.
 *
 * This code is public domain. Feel free to use it for any purpose!
 */

package main

import (
	"math"
	"os"

	"github.com/smack0007/sdl3-go/sdl"
)

type AppState struct {
	/* We will use this renderer to draw into this window every frame. */
	window   *sdl.Window
	renderer *sdl.Renderer
}

func main() {
	/* use the callbacks */
	os.Exit(sdl.EnterAppMainCallbacks(os.Args, AppInit, AppIterate, AppEvent, AppQuit))
}

/* This function runs once at startup. */
func AppInit(argv []string) (sdl.AppResult, *AppState) {
	sdl.SetAppMetadata("Example Renderer Clear", "1.0", "com.example.renderer-clear")

	err := sdl.Init(sdl.INIT_VIDEO)

	if err != nil {
		sdl.Log("Couldn't initialize SDL: %s", err)
		return sdl.APP_FAILURE, nil
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/renderer/clear", 640, 480, 0)

	if err != nil {
		sdl.Log("Couldn't create window/renderer: %s", err)
		return sdl.APP_FAILURE, nil
	}

	/* carry on with the program! */
	return sdl.APP_CONTINUE, &AppState{
		window:   window,
		renderer: renderer,
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
	now := (float64)(sdl.GetTicks()) / 1000.0 /* convert from milliseconds to seconds. */
	/* choose the color for the frame we will draw. The sine wave trick makes it fade between colors smoothly. */
	red := (float32)(0.5 + 0.5*math.Sin(now))
	green := (float32)(0.5 + 0.5*math.Sin(now+math.Pi*2/3))
	blue := (float32)(0.5 + 0.5*math.Sin(now+math.Pi*4/3))
	sdl.SetRenderDrawColorFloat(appState.renderer, red, green, blue, sdl.ALPHA_OPAQUE_FLOAT) /* new color, full alpha. */

	/* clear the window to the draw color. */
	sdl.RenderClear(appState.renderer)

	/* put the newly-cleared rendering on the screen. */
	sdl.RenderPresent(appState.renderer)

	return sdl.APP_CONTINUE /* carry on with the program! */
}

/* This function runs once at shutdown. */
func AppQuit(appState *AppState, result sdl.AppResult) {
	/* SDL will clean up the window/renderer for us. */
}
