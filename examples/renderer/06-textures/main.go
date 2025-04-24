/*
 * This example creates an SDL window and renderer, and then draws some
 * rectangles to it every frame.
 *
 * This code is public domain. Feel free to use it for any purpose!
 */

package main

import (
	_ "embed"
	"os"

	"github.com/smack0007/sdl-go/sdl"
)

const (
	WINDOW_WIDTH  = 640
	WINDOW_HEIGHT = 480
)

//go:embed sample.bmp
var sampleBmp []byte

type AppState struct {
	/* We will use this renderer to draw into this window every frame. */
	window   *sdl.Window
	renderer *sdl.Renderer

	texture        *sdl.Texture
	texture_width  int
	texture_height int
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

	/* Textures are pixel data that we upload to the video hardware for fast drawing. Lots of 2D
	   engines refer to these as "sprites." We'll do a static texture (upload once, draw many
	   times) with data from a bitmap file. */

	/* SDL_Surface is pixel data the CPU can access. SDL_Texture is pixel data the GPU can access.
	   Load a .bmp into a surface, move it to a texture from there. */

	io, err := sdl.IOFromConstMem(&sampleBmp[0], (uint64)(len(sampleBmp)))
	if err != nil {
		sdl.Log("Couldn't create bitmap io: %s", err)
		return sdl.APP_FAILURE, nil
	}

	surface, err := sdl.LoadBMP_IO(io, true)
	if err != nil {
		sdl.Log("Couldn't load bitmap: %s", err)
		return sdl.APP_FAILURE, nil
	}

	texture_width := surface.W()
	texture_height := surface.H()

	texture, err := sdl.CreateTextureFromSurface(renderer, surface)
	if err != nil {
		sdl.Log("Couldn't create static texture: %s", err)
		return sdl.APP_FAILURE, nil
	}

	sdl.DestroySurface(surface) /* done with this, the texture has a copy of the pixels now. */

	/* carry on with the program! */
	return sdl.APP_CONTINUE, &AppState{
		window:   window,
		renderer: renderer,

		texture:        texture,
		texture_width:  texture_width,
		texture_height: texture_height,
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

	/* we'll have some textures move around over a few seconds. */
	direction := float32(-1.0)
	if (now % 2000) >= 1000 {
		direction = 1.0
	}

	scale := (float32(int(now%1000)-500) / float32(500.0)) * direction

	/* as you can see from this, rendering draws over whatever was drawn before it. */
	sdl.SetRenderDrawColor(appState.renderer, 0, 0, 0, sdl.ALPHA_OPAQUE) /* black, full alpha */
	sdl.RenderClear(appState.renderer)                                   /* start with a blank canvas. */

	/* Just draw the static texture a few times. You can think of it like a
	   stamp, there isn't a limit to the number of times you can draw with it. */

	/* top left */
	dst_rect.X = 100.0 * scale
	dst_rect.Y = 0.0
	dst_rect.W = float32(appState.texture_width)
	dst_rect.H = float32(appState.texture_height)
	sdl.RenderTexture(appState.renderer, appState.texture, nil, &dst_rect)

	/* center this one. */
	dst_rect.X = float32(WINDOW_WIDTH-appState.texture_width) / 2.0
	dst_rect.Y = float32(WINDOW_HEIGHT-appState.texture_height) / 2.0
	dst_rect.W = float32(appState.texture_width)
	dst_rect.H = float32(appState.texture_height)
	sdl.RenderTexture(appState.renderer, appState.texture, nil, &dst_rect)

	/* bottom right. */
	dst_rect.X = float32(WINDOW_WIDTH-appState.texture_width) - (100.0 * scale)
	dst_rect.Y = float32(WINDOW_HEIGHT - appState.texture_height)
	dst_rect.W = float32(appState.texture_width)
	dst_rect.H = float32(appState.texture_height)
	sdl.RenderTexture(appState.renderer, appState.texture, nil, &dst_rect)

	sdl.RenderPresent(appState.renderer) /* put it all on the screen! */

	return sdl.APP_CONTINUE /* carry on with the program! */
}

/* This function runs once at shutdown. */
func AppQuit(appState *AppState, result sdl.AppResult) {
	/* SDL will clean up the window/renderer for us. */
}
