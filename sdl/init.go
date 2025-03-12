package sdl

//#include "init.h"
import "C"
import (
	"unsafe"
)

type AppResult C.SDL_AppResult
type InitFlags uint32

const (
	APP_CONTINUE AppResult = C.SDL_APP_CONTINUE
	APP_SUCCESS  AppResult = C.SDL_APP_SUCCESS
	APP_FAILURE  AppResult = C.SDL_APP_FAILURE

	INIT_AUDIO    InitFlags = C.SDL_INIT_AUDIO
	INIT_VIDEO    InitFlags = C.SDL_INIT_VIDEO
	INIT_JOYSTICK InitFlags = C.SDL_INIT_JOYSTICK
	INIT_HAPTIC   InitFlags = C.SDL_INIT_HAPTIC
	INIT_GAMEPAD  InitFlags = C.SDL_INIT_GAMEPAD
	INIT_EVENTS   InitFlags = C.SDL_INIT_EVENTS
	INIT_SENSOR   InitFlags = C.SDL_INIT_SENSOR
	INIT_CAMERA   InitFlags = C.SDL_INIT_CAMERA

	PROP_APP_METADATA_NAME_STRING       string = C.SDL_PROP_APP_METADATA_NAME_STRING
	PROP_APP_METADATA_VERSION_STRING    string = C.SDL_PROP_APP_METADATA_VERSION_STRING
	PROP_APP_METADATA_IDENTIFIER_STRING string = C.SDL_PROP_APP_METADATA_IDENTIFIER_STRING
	PROP_APP_METADATA_CREATOR_STRING    string = C.SDL_PROP_APP_METADATA_CREATOR_STRING
	PROP_APP_METADATA_COPYRIGHT_STRING  string = C.SDL_PROP_APP_METADATA_COPYRIGHT_STRING
	PROP_APP_METADATA_URL_STRING        string = C.SDL_PROP_APP_METADATA_URL_STRING
	PROP_APP_METADATA_TYPE_STRING       string = C.SDL_PROP_APP_METADATA_TYPE_STRING
)

type AppInit_func[AppState any] func(argv []string) (AppResult, *AppState)
type AppIterate_func[AppState any] func(appState *AppState) AppResult
type AppEvent_func[AppState any] func(appState *AppState, event *Event) AppResult
type AppQuit_func[AppState any] func(appState *AppState, result AppResult)

func Init(flags InitFlags) error {
	return mapErrorBool(bool(C.SDL_Init(C.Uint32(flags))))
}

func Quit() {
	C.SDL_Quit()
}

func SetAppMetadata(appname string, appversion string, appidentifier string) error {
	c_appname := C.CString(appname)
	defer C.free(unsafe.Pointer(c_appname))

	c_appversion := C.CString(appversion)
	defer C.free(unsafe.Pointer(c_appversion))

	c_appidentifier := C.CString(appidentifier)
	defer C.free(unsafe.Pointer(c_appidentifier))

	return mapErrorBool(
		bool(
			C.SDL_SetAppMetadata(
				c_appname,
				c_appversion,
				c_appidentifier,
			),
		),
	)
}
