package logger

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

type Logger struct {
	*zerolog.Logger
}

func NewLogger() *Logger {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return &Logger{&logger}
}
