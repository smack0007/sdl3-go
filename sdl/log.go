package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL_log.h>

static inline void _SDL_Log(const char *str)
{
  SDL_Log("%s", str);
}

static inline void _SDL_LogCritical(int category,
                                    const char *str)
{
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wformat"
	SDL_LogCritical(category, str);
#pragma clang diagnostic pop
}

static inline void _SDL_LogDebug(int category,
                                 const char *str)
{
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wformat"
	SDL_LogDebug(category, str);
#pragma clang diagnostic pop
}

static inline void _SDL_LogError(int category,
                                 const char *str)
{
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wformat"
	SDL_LogError(category, str);
#pragma clang diagnostic pop
}

static inline void _SDL_LogInfo(int category,
                                const char *str)
{
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wformat"
  SDL_LogInfo(category, str);
#pragma clang diagnostic pop
}

static inline void _SDL_LogMessage(int category,
                                   SDL_LogPriority priority,
                                   const char *str)
{
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wformat"
  SDL_LogMessage(category, priority, str);
#pragma clang diagnostic pop
}

static inline void _SDL_LogVerbose(int category,
                                   const char *str)
{
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wformat"
  SDL_LogVerbose(category, str);
#pragma clang diagnostic pop
}

static inline void _SDL_LogWarn(int category,
                                const char *str)
{
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wformat"
  SDL_LogWarn(category, str);
#pragma clang diagnostic pop
}
*/
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
	LOG_CATEGORY_GPU         int = C.SDL_LOG_CATEGORY_GPU
	LOG_CATEGORY_CUSTOM      int = C.SDL_LOG_CATEGORY_CUSTOM

	// SDL_LogPriority
	LOG_PRIORITY_INVALID  LogPriority = C.SDL_LOG_PRIORITY_INVALID
	LOG_PRIORITY_TRACE    LogPriority = C.SDL_LOG_PRIORITY_TRACE
	LOG_PRIORITY_VERBOSE  LogPriority = C.SDL_LOG_PRIORITY_VERBOSE
	LOG_PRIORITY_DEBUG    LogPriority = C.SDL_LOG_PRIORITY_DEBUG
	LOG_PRIORITY_INFO     LogPriority = C.SDL_LOG_PRIORITY_INFO
	LOG_PRIORITY_WARN     LogPriority = C.SDL_LOG_PRIORITY_WARN
	LOG_PRIORITY_ERROR    LogPriority = C.SDL_LOG_PRIORITY_ERROR
	LOG_PRIORITY_CRITICAL LogPriority = C.SDL_LOG_PRIORITY_CRITICAL
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

func GetLogPriority(category int) LogPriority {
	return LogPriority(C.SDL_GetLogPriority(C.int(category)))
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

func ResetLogPriorities() {
	C.SDL_ResetLogPriorities()
}

func SetLogPriorities(
	priority LogPriority,
) {
	C.SDL_SetLogPriorities(C.SDL_LogPriority(priority))
}

func SetLogPriority(
	category int,
	priority LogPriority,
) {
	C.SDL_SetLogPriority(C.int(category), C.SDL_LogPriority(priority))
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
