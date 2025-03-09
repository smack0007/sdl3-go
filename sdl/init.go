package sdl

//#include "init.h"
import "C"

type AppResult C.SDL_AppResult

const (
	APP_CONTINUE AppResult = C.SDL_APP_CONTINUE
	APP_SUCCESS  AppResult = C.SDL_APP_SUCCESS
	APP_FAILURE  AppResult = C.SDL_APP_FAILURE
)

type AppInit_func[AppState any] func(argv []string) (AppResult, *AppState)
type AppIterate_func[AppState any] func(appState *AppState) AppResult
type AppEvent_func[AppState any] func(appState *AppState, event *Event) AppResult
type AppQuit_func[AppState any] func(appState *AppState, result AppResult)
