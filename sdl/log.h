static inline void _SDL_Log(const char *str)
{
  SDL_Log("%s", str);
}

static inline void _SDL_LogCritical(int category,
                                    const char *str)
{
  SDL_LogCritical(category, str);
}

static inline void _SDL_LogDebug(int category,
                                 const char *str)
{
  SDL_LogDebug(category, str);
}

static inline void _SDL_LogError(int category,
                                 const char *str)
{
  SDL_LogError(category, str);
}

static inline void _SDL_LogInfo(int category,
                                const char *str)
{
  SDL_LogInfo(category, str);
}

static inline void _SDL_LogMessage(int category,
                                   SDL_LogPriority priority,
                                   const char *str)
{
  SDL_LogMessage(category, priority, str);
}

static inline void _SDL_LogVerbose(int category,
                                   const char *str)
{
  SDL_LogVerbose(category, str);
}

static inline void _SDL_LogWarn(int category,
                                const char *str)
{
  SDL_LogWarn(category, str);
}