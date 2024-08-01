package sdl

//#include "SDL_go.h"
import "C"

import (
	"fmt"
	"unsafe"
)

type LogPriority C.SDL_LogPriority

const (
	// SDL_LogCategory
	LOG_CATEGORY_APPLICATION int = C.SDL_LOG_CATEGORY_APPLICATION
	LOG_CATEGORY_ERROR       int = C.SDL_LOG_CATEGORY_ERROR
	LOG_CATEGORY_ASSERT      int = C.SDL_LOG_CATEGORY_ASSERT
	LOG_CATEGORY_SYSTEM      int = C.SDL_LOG_CATEGORY_SYSTEM
	LOG_CATEGORY_AUDIO       int = C.SDL_LOG_CATEGORY_AUDIO
	LOG_CATEGORY_VIDEO       int = C.SDL_LOG_CATEGORY_VIDEO
	LOG_CATEGORY_RENDER      int = C.SDL_LOG_CATEGORY_RENDER
	LOG_CATEGORY_INPUT       int = C.SDL_LOG_CATEGORY_INPUT
	LOG_CATEGORY_TEST        int = C.SDL_LOG_CATEGORY_TEST
	LOG_CATEGORY_CUSTOM      int = C.SDL_LOG_CATEGORY_CUSTOM

	// SDL_LogPriority
	LOG_PRIORITY_VERBOSE  LogPriority = C.SDL_LOG_PRIORITY_VERBOSE
	LOG_PRIORITY_DEBUG    LogPriority = C.SDL_LOG_PRIORITY_DEBUG
	LOG_PRIORITY_INFO     LogPriority = C.SDL_LOG_PRIORITY_INFO
	LOG_PRIORITY_WARN     LogPriority = C.SDL_LOG_PRIORITY_WARN
	LOG_PRIORITY_ERROR    LogPriority = C.SDL_LOG_PRIORITY_ERROR
	LOG_PRIORITY_CRITICAL LogPriority = C.SDL_LOG_PRIORITY_CRITICAL
	NUM_LOG_PRIORITIES    LogPriority = C.SDL_NUM_LOG_PRIORITIES
)

func Log(str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	C._SDL_Log(cStr)
}

func LogCritical(
	category int,
	str string,
	args ...interface{},
) {
	str = fmt.Sprintf(str, args...)

	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	C._SDL_LogCritical(C.int(category), cStr)
}

func LogDebug(
	category int,
	str string,
	args ...interface{},
) {
	str = fmt.Sprintf(str, args...)

	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	C._SDL_LogDebug(C.int(category), cStr)
}

func LogError(
	category int,
	str string,
	args ...interface{},
) {
	str = fmt.Sprintf(str, args...)

	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	C._SDL_LogError(C.int(category), cStr)
}

func LogGetPriority(category int) LogPriority {
	return LogPriority(C.SDL_LogGetPriority(C.int(category)))
}

func LogInfo(
	category int,
	str string,
	args ...interface{},
) {
	str = fmt.Sprintf(str, args...)

	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	C._SDL_LogInfo(C.int(category), cStr)
}

func LogMessage(
	category int,
	priority LogPriority,
	str string,
	args ...interface{},
) {
	str = fmt.Sprintf(str, args...)

	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	C._SDL_LogMessage(C.int(category), C.SDL_LogPriority(priority), cStr)
}

func LogResetPriorities() {
	C.SDL_LogResetPriorities()
}

func LogSetAllPriority(
	priority LogPriority,
) {
	C.SDL_LogSetAllPriority(C.SDL_LogPriority(priority))
}

// TODO: Implement LogSetOutputFunction.
// func LogSetOutputFunction() {
// }

func LogSetPriority(
	category int,
	priority LogPriority,
) {
	C.SDL_LogSetPriority(C.int(category), C.SDL_LogPriority(priority))
}

func LogVerbose(
	category int,
	str string,
	args ...interface{},
) {
	str = fmt.Sprintf(str, args...)

	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	C._SDL_LogVerbose(C.int(category), cStr)
}

func LogWarn(
	category int,
	str string,
	args ...interface{},
) {
	str = fmt.Sprintf(str, args...)

	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	C._SDL_LogWarn(C.int(category), cStr)
}
