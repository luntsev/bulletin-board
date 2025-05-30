package logger

import (
	"bulletin-board/config"
	"os"

	"github.com/rs/zerolog"
)

func NewLogger(conf *config.LogConfig) *zerolog.Logger {
	zerolog.SetGlobalLevel(conf.LogLevel)
	var logger zerolog.Logger
	if conf.LogFormat == "json" {
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	} else {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
	}

	return &logger
}
