package loggerUtil

import (
	"os"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
)

func ZerologSetup(level zerolog.Level) zerolog.Logger {
	buildInfo, _ := debug.ReadBuildInfo()
	zerolog.SetGlobalLevel(level)
	return zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Str("go_version", buildInfo.GoVersion).
		Logger()
}
