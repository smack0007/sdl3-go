/*
 * This example creates an SDL window and renderer, and then draws some lines
 * to it every frame.
 *
 * This code is public domain. Feel free to use it for any purpose!
 */

package main

import (
	"math"
	"math/rand/v2"
	"os"

	"github.com/smack0007/sdl-go/sdl"
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
	sdl.SetAppMetadata("Example Renderer Lines", "1.0", "com.example.renderer-lines")

	err := sdl.Init(sdl.INIT_VIDEO)

	if err != nil {
		sdl.Log("Couldn't initialize SDL: %s", err)
		return sdl.APP_FAILURE, nil
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/renderer/lines", 640, 480, 0)

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

/*
	Lines (line segments, really) are drawn in terms of points: a set of

X and Y coordinates, one set for each end of the line.
(0, 0) is the top left of the window, and larger numbers go down
and to the right. This isn't how geometry works, but this is pretty
standard in 2D graphics.
*/
var line_points = []sdl.FPoint{
	{X: 100, Y: 354}, {X: 220, Y: 230}, {X: 140, Y: 230}, {X: 320, Y: 100}, {X: 500, Y: 230},
	{X: 420, Y: 230}, {X: 540, Y: 354}, {X: 400, Y: 354}, {X: 100, Y: 354},
}

/* This function runs once per frame, and is the heart of the program. */
func AppIterate(appState *AppState) sdl.AppResult {

	/* as you can see from this, rendering draws over whatever was drawn before it. */
	sdl.SetRenderDrawColor(appState.renderer, 100, 100, 100, sdl.ALPHA_OPAQUE) /* grey, full alpha */
	sdl.RenderClear(appState.renderer)                                         /* start with a blank canvas. */

	/* You can draw lines, one at a time, like these brown ones... */
	sdl.SetRenderDrawColor(appState.renderer, 127, 49, 32, sdl.ALPHA_OPAQUE)
	sdl.RenderLine(appState.renderer, 240, 450, 400, 450)
	sdl.RenderLine(appState.renderer, 240, 356, 400, 356)
	sdl.RenderLine(appState.renderer, 240, 356, 240, 450)
	sdl.RenderLine(appState.renderer, 400, 356, 400, 450)

	/* You can also draw a series of connected lines in a single batch... */
	sdl.SetRenderDrawColor(appState.renderer, 0, 255, 0, sdl.ALPHA_OPAQUE)
	sdl.RenderLines(appState.renderer, line_points, len(line_points))

	/* here's a bunch of lines drawn out from a center point in a circle. */
	/* we randomize the color of each line, so it functions as animation. */
	for i := 0; i < 360; i++ {
		size := float32(30.0)
		x := float32(320.0)
		y := float32(95.0 - size/2.0)
		sdl.SetRenderDrawColor(appState.renderer, uint8(rand.IntN(256)), uint8(rand.IntN(256)), uint8(rand.IntN(256)), sdl.ALPHA_OPAQUE)
		sdl.RenderLine(appState.renderer, x, y, x+float32(math.Sin(float64(i)))*size, y+float32(math.Cos(float64(i)))*size)
	}

	sdl.RenderPresent(appState.renderer) /* put it all on the screen! */

	return sdl.APP_CONTINUE /* carry on with the program! */
}

/* This function runs once at shutdown. */
func AppQuit(appState *AppState, result sdl.AppResult) {
	/* SDL will clean up the window/renderer for us. */
}
