package main

import (
	"os"

	"github.com/smack0007/sdl-go/sdl"
)

type AppState struct {
}

func main() {
	os.Exit(sdl.EnterAppMainCallbacks(os.Args, AppInit, AppIterate, AppEvent, AppQuit))
}

func AppInit(argv []string) (sdl.AppResult, *AppState) {
	return sdl.APP_CONTINUE, &AppState{}
}

func AppIterate(appState *AppState) sdl.AppResult {
	return sdl.APP_CONTINUE
}

func AppEvent(appState *AppState, event *sdl.Event) sdl.AppResult {
	return sdl.APP_CONTINUE
}

func AppQuit(appState *AppState, result sdl.AppResult) {
}
