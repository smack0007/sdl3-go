package sdl

/*
#define SDL_MAIN_HANDLED
#include <SDL3/SDL_main.h>

extern SDL_AppResult _SDL_AppInit();
extern SDL_AppResult _SDL_AppEvent(SDL_Event *event);
extern SDL_AppResult _SDL_AppIterate();
extern void _SDL_AppQuit(SDL_AppResult result);

static inline SDL_AppResult __SDL_AppInit(void **appstate, int argc, char *argv[])
{
	return _SDL_AppInit();
}

static inline SDL_AppResult __SDL_AppIterate(void *appstate)
{
	return _SDL_AppIterate();
}

static inline SDL_AppResult __SDL_AppEvent(void *appstate, SDL_Event *event)
{
	return _SDL_AppEvent(event);
}

static inline void __SDL_AppQuit(void *appstate, SDL_AppResult result)
{
	_SDL_AppQuit(result);
}

static inline int _SDL_EnterAppMainCallbacks()
{
	char* argv[] = {};

	return SDL_EnterAppMainCallbacks(
		0,
		argv,
		__SDL_AppInit,
		__SDL_AppIterate,
		__SDL_AppEvent,
		__SDL_AppQuit
	);
}
*/
import "C"

import (
	"reflect"
	"runtime"
)

const ()

type enterAppMainCallbacksState struct {
	argv       []string
	appInit    reflect.Value
	appIterate reflect.Value
	appEvent   reflect.Value
	appQuit    reflect.Value
	appState   reflect.Value
}

var _enterAppMainCallbacksState enterAppMainCallbacksState

func EnterAppMainCallbacks[AppState any](
	argv []string,
	appinit AppInit_func[AppState],
	appiterate AppIterate_func[AppState],
	appevent AppEvent_func[AppState],
	appquit AppQuit_func[AppState],
) int {
	runtime.LockOSThread()

	_enterAppMainCallbacksState = enterAppMainCallbacksState{
		argv:       argv,
		appInit:    reflect.ValueOf(appinit),
		appIterate: reflect.ValueOf(appiterate),
		appEvent:   reflect.ValueOf(appevent),
		appQuit:    reflect.ValueOf(appquit),
	}

	return int(
		C._SDL_EnterAppMainCallbacks(),
	)
}

//export _SDL_AppInit
func _SDL_AppInit() AppResult {
	values := _enterAppMainCallbacksState.appInit.Call([]reflect.Value{reflect.ValueOf(_enterAppMainCallbacksState.argv)})

	result := values[0].Interface().(AppResult)
	_enterAppMainCallbacksState.appState = values[1]

	return result
}

//export _SDL_AppIterate
func _SDL_AppIterate() AppResult {
	values := _enterAppMainCallbacksState.appIterate.Call([]reflect.Value{
		_enterAppMainCallbacksState.appState,
	})

	return values[0].Interface().(AppResult)
}

//export _SDL_AppEvent
func _SDL_AppEvent(event *C.SDL_Event) AppResult {
	goEvent := cEventPointerToEventPointer(event)

	values := _enterAppMainCallbacksState.appEvent.Call([]reflect.Value{_enterAppMainCallbacksState.appState, reflect.ValueOf(goEvent)})

	return values[0].Interface().(AppResult)
}

//export _SDL_AppQuit
func _SDL_AppQuit(result AppResult) {
	_enterAppMainCallbacksState.appQuit.Call([]reflect.Value{
		_enterAppMainCallbacksState.appState,
		reflect.ValueOf(result),
	})
}
