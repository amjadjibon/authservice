package logger

import (
	"github.com/rs/zerolog"
)

func InitLogger(logLevel string) {
	zerolog.MessageFieldName = "message"
	zerolog.LevelFieldName = "level"
	zerolog.TimestampFieldName = "timestamp"
	switch logLevel {
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
