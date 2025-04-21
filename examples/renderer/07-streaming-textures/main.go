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

	TEXTURE_SIZE = 150
)

type AppState struct {
	/* We will use this renderer to draw into this window every frame. */
	window   *sdl.Window
	renderer *sdl.Renderer

	texture *sdl.Texture
}

func main() {
	/* use the callbacks */
	os.Exit(sdl.EnterAppMainCallbacks(os.Args, AppInit, AppIterate, AppEvent, AppQuit))
}

/* This function runs once at startup. */
func AppInit(argv []string) (sdl.AppResult, *AppState) {
	sdl.SetAppMetadata("Example Renderer Streaming Textures", "1.0", "com.example.renderer-streaming-textures")

	err := sdl.Init(sdl.INIT_VIDEO)

	if err != nil {
		sdl.Log("Couldn't initialize SDL: %s", err)
		return sdl.APP_FAILURE, nil
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/renderer/streaming-textures", WINDOW_WIDTH, WINDOW_HEIGHT, 0)

	if err != nil {
		sdl.Log("Couldn't create window/renderer: %s", err)
		return sdl.APP_FAILURE, nil
	}

	texture, err := sdl.CreateTexture(renderer, sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STREAMING, TEXTURE_SIZE, TEXTURE_SIZE)
	if err != nil {
		sdl.Log("Couldn't create streaming texture: %s", err)
		return sdl.APP_FAILURE, nil
	}

	/* carry on with the program! */
	return sdl.APP_CONTINUE, &AppState{
		window:   window,
		renderer: renderer,

		texture: texture,
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
	var dst_rect sdl.FRect
	now := sdl.GetTicks()
	var surface *sdl.Surface

	/* we'll have some color move around over a few seconds. */
	direction := float32(-1.0)
	if int32(now%2000) >= 1000 {
		direction = 1.0
	}
	scale := float32(int32(now%1000)-500) / float32(500.0) * direction

	/* To update a streaming texture, you need to lock it first. This gets you access to the pixels.
	   Note that this is considered a _write-only_ operation: the buffer you get from locking
	   might not acutally have the existing contents of the texture, and you have to write to every
	   locked pixel! */

	/* You can use SDL_LockTexture() to get an array of raw pixels, but we're going to use
	   SDL_LockTextureToSurface() here, because it wraps that array in a temporary SDL_Surface,
	   letting us use the surface drawing functions instead of lighting up individual pixels. */
	err := sdl.LockTextureToSurface(appState.texture, nil, &surface)
	if err == nil {
		var r sdl.Rect

		pixelFormatDetails, err := sdl.GetPixelFormatDetails(surface.Format())

		if err != nil {
			sdl.Log("Failed to get pixel format details: %s", err)
			return sdl.APP_FAILURE
		}

		sdl.FillSurfaceRect(surface, nil, sdl.MapRGB(pixelFormatDetails, nil, 0, 0, 0)) /* make the whole surface black */
		r.W = TEXTURE_SIZE
		r.H = TEXTURE_SIZE / 10
		r.X = 0
		r.Y = int32(float32(int32(TEXTURE_SIZE)-r.H) * ((scale + 1.0) / 2.0))
		sdl.FillSurfaceRect(surface, &r, sdl.MapRGB(pixelFormatDetails, nil, 0, 255, 0)) /* make a strip of the surface green */
		sdl.UnlockTexture(appState.texture)                                              /* upload the changes (and frees the temporary surface)! */
	}

	/* as you can see from this, rendering draws over whatever was drawn before it. */
	sdl.SetRenderDrawColor(appState.renderer, 66, 66, 66, sdl.ALPHA_OPAQUE) /* grey, full alpha */
	sdl.RenderClear(appState.renderer)                                      /* start with a blank canvas. */

	/* Just draw the static texture a few times. You can think of it like a
	   stamp, there isn't a limit to the number of times you can draw with it. */

	/* Center this one. It'll draw the latest version of the texture we drew while it was locked. */
	dst_rect.X = float32(WINDOW_WIDTH-TEXTURE_SIZE) / float32(2.0)
	dst_rect.Y = float32(WINDOW_HEIGHT-TEXTURE_SIZE) / float32(2.0)
	dst_rect.W = float32(TEXTURE_SIZE)
	dst_rect.H = float32(TEXTURE_SIZE)
	sdl.RenderTexture(appState.renderer, appState.texture, nil, &dst_rect)

	sdl.RenderPresent(appState.renderer) /* put it all on the screen! */

	return sdl.APP_CONTINUE /* carry on with the program! */
}

/* This function runs once at shutdown. */
func AppQuit(appState *AppState, result sdl.AppResult) {
	/* SDL will clean up the window/renderer for us. */
}
