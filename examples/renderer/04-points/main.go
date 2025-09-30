/*
 * This example creates an SDL window and renderer, and then draws some points
 * to it every frame.
 *
 * This code is public domain. Feel free to use it for any purpose!
 */

package main

import (
	"math/rand/v2"
	"os"

	"github.com/smack0007/sdl3-go/sdl"
)

const (
	WINDOW_WIDTH  = 640
	WINDOW_HEIGHT = 480

	NUM_POINTS            = 500
	MIN_PIXELS_PER_SECOND = 30 /* move at least this many pixels per second. */
	MAX_PIXELS_PER_SECOND = 60 /* move this many pixels per second at most. */
)

type AppState struct {
	/* We will use this renderer to draw into this window every frame. */
	window    *sdl.Window
	renderer  *sdl.Renderer
	last_time uint64

	points       []sdl.FPoint
	point_speeds []float32
}

func main() {
	/* use the callbacks */
	os.Exit(sdl.EnterAppMainCallbacks(os.Args, AppInit, AppIterate, AppEvent, AppQuit))
}

/* This function runs once at startup. */
func AppInit(argv []string) (sdl.AppResult, *AppState) {
	sdl.SetAppMetadata("Example Renderer Points", "1.0", "com.example.renderer-points")

	err := sdl.Init(sdl.INIT_VIDEO)

	if err != nil {
		sdl.Log("Couldn't initialize SDL: %s", err)
		return sdl.APP_FAILURE, nil
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/renderer/points", WINDOW_WIDTH, WINDOW_HEIGHT, 0)

	if err != nil {
		sdl.Log("Couldn't create window/renderer: %s", err)
		return sdl.APP_FAILURE, nil
	}

	points := make([]sdl.FPoint, NUM_POINTS)
	point_speeds := make([]float32, NUM_POINTS)

	/* set up the data for a bunch of points. */
	for i := 0; i < len(points); i += 1 {
		points[i].X = rand.Float32() * (float32(WINDOW_WIDTH))
		points[i].Y = rand.Float32() * (float32(WINDOW_HEIGHT))
		point_speeds[i] = MIN_PIXELS_PER_SECOND + (rand.Float32() * (MAX_PIXELS_PER_SECOND - MIN_PIXELS_PER_SECOND))
	}

	last_time := sdl.GetTicks()

	/* carry on with the program! */
	return sdl.APP_CONTINUE, &AppState{
		window:    window,
		renderer:  renderer,
		last_time: last_time,

		points:       points,
		point_speeds: point_speeds,
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
	now := sdl.GetTicks()
	elapsed := float32(now-appState.last_time) / float32(1000.0) /* seconds since last iteration */

	/* let's move all our points a little for a new frame. */
	for i := 0; i < len(appState.points); i += 1 {
		distance := elapsed * appState.point_speeds[i]
		appState.points[i].X += distance
		appState.points[i].Y += distance
		if (appState.points[i].X >= WINDOW_WIDTH) || (appState.points[i].Y >= WINDOW_HEIGHT) {
			/* off the screen; restart it elsewhere! */
			if rand.IntN(2) == 1 {
				appState.points[i].X = rand.Float32() * float32(WINDOW_WIDTH)
				appState.points[i].Y = 0.0
			} else {
				appState.points[i].X = 0.0
				appState.points[i].Y = rand.Float32() * float32(WINDOW_HEIGHT)
			}
			appState.point_speeds[i] = MIN_PIXELS_PER_SECOND + (rand.Float32() * (MAX_PIXELS_PER_SECOND - MIN_PIXELS_PER_SECOND))
		}
	}

	appState.last_time = now

	/* as you can see from this, rendering draws over whatever was drawn before it. */
	sdl.SetRenderDrawColor(appState.renderer, 0, 0, 0, sdl.ALPHA_OPAQUE)       /* black, full alpha */
	sdl.RenderClear(appState.renderer)                                         /* start with a blank canvas. */
	sdl.SetRenderDrawColor(appState.renderer, 255, 255, 255, sdl.ALPHA_OPAQUE) /* white, full alpha */
	sdl.RenderPoints(appState.renderer, appState.points, len(appState.points)) /* draw all the points! */

	/* You can also draw single points with SDL_RenderPoint(), but it's
	   cheaper (sometimes significantly so) to do them all at once. */

	sdl.RenderPresent(appState.renderer) /* put it all on the screen! */

	return sdl.APP_CONTINUE /* carry on with the program! */
}

/* This function runs once at shutdown. */
func AppQuit(appState *AppState, result sdl.AppResult) {
	/* SDL will clean up the window/renderer for us. */
}
