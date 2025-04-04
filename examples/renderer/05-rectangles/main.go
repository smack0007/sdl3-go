/*
 * This example creates an SDL window and renderer, and then draws some
 * rectangles to it every frame.
 *
 * This code is public domain. Feel free to use it for any purpose!
 */

package main

import (
	"os"

	"github.com/smack0007/sdl-go/sdl"
)

const (
	WINDOW_WIDTH  = 640
	WINDOW_HEIGHT = 480
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
	sdl.SetAppMetadata("Example Renderer Rectangles", "1.0", "com.example.renderer-rectangles")

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

var rects []sdl.FRect = make([]sdl.FRect, 16)

/* This function runs once per frame, and is the heart of the program. */
func AppIterate(appState *AppState) sdl.AppResult {
	now := sdl.GetTicks()

	/* we'll have the rectangles grow and shrink over a few seconds. */
	direction := float32(-1.0)
	if now%2000 >= 1000 {
		direction = 1.0
	}

	scale := float32(int(now%1000)-500) / float32(500.0) * direction

	/* as you can see from this, rendering draws over whatever was drawn before it. */
	sdl.SetRenderDrawColor(appState.renderer, 0, 0, 0, sdl.ALPHA_OPAQUE) /* black, full alpha */
	sdl.RenderClear(appState.renderer)                                   /* start with a blank canvas. */

	/* Rectangles are comprised of set of X and Y coordinates, plus width and
	   height. (0, 0) is the top left of the window, and larger numbers go
	   down and to the right. This isn't how geometry works, but this is
	   pretty standard in 2D graphics. */

	/* Let's draw a single rectangle (square, really). */
	rects[0].X = 100
	rects[0].Y = 100
	rects[0].W = 100 + (100 * scale)
	rects[0].H = 100 + (100 * scale)
	sdl.SetRenderDrawColor(appState.renderer, 255, 0, 0, sdl.ALPHA_OPAQUE) /* red, full alpha */
	sdl.RenderRect(appState.renderer, &rects[0])

	/* Now let's draw several rectangles with one function call. */
	for i := 0; i < 3; i += 1 {
		size := float32(i+1) * 50.0
		rects[i].W = size + (size * scale)
		rects[i].H = size + (size * scale)
		rects[i].X = (WINDOW_WIDTH - rects[i].W) / 2  /* center it. */
		rects[i].Y = (WINDOW_HEIGHT - rects[i].H) / 2 /* center it. */
	}
	sdl.SetRenderDrawColor(appState.renderer, 0, 255, 0, sdl.ALPHA_OPAQUE) /* green, full alpha */
	sdl.RenderRects(appState.renderer, rects, 3)                           /* draw three rectangles at once */

	/* those were rectangle _outlines_, really. You can also draw _filled_ rectangles! */
	rects[0].X = 400
	rects[0].Y = 50
	rects[0].W = 100 + (100 * scale)
	rects[0].H = 50 + (50 * scale)
	sdl.SetRenderDrawColor(appState.renderer, 0, 0, 255, sdl.ALPHA_OPAQUE) /* blue, full alpha */
	sdl.RenderFillRect(appState.renderer, &rects[0])

	/* ...and also fill a bunch of rectangles at once... */
	for i := 0; i < len(rects); i += 1 {
		w := float32(WINDOW_WIDTH / len(rects))
		h := float32(i) * 8.0
		rects[i].X = float32(i) * w
		rects[i].Y = WINDOW_HEIGHT - h
		rects[i].W = w
		rects[i].H = h
	}
	sdl.SetRenderDrawColor(appState.renderer, 255, 255, 255, sdl.ALPHA_OPAQUE) /* white, full alpha */
	sdl.RenderFillRects(appState.renderer, rects, len(rects))

	sdl.RenderPresent(appState.renderer) /* put it all on the screen! */

	return sdl.APP_CONTINUE /* carry on with the program! */
}

/* This function runs once at shutdown. */
func AppQuit(appState *AppState, result sdl.AppResult) {
	/* SDL will clean up the window/renderer for us. */
}
