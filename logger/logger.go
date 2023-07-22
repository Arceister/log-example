package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	logLevel    string
	logLocation *os.File
}

func NewLogger(
	logLevel string,
	logLocation *os.File,
) *Logger {
	return &Logger{
		logLevel:    logLevel,
		logLocation: logLocation,
	}
}

func (l *Logger) Setup() zerolog.Logger {
	switch l.logLevel {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}
	logger := zerolog.New(l.logLocation).With().Timestamp().Logger()
	return logger
}
