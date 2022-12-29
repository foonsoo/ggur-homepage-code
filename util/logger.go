package util

import (
	"io"

	"github.com/rs/zerolog"
)

type (
	Logger = zerolog.Logger
)

func NewLogger(w io.Writer) Logger {
	return zerolog.New(w)
}

func ParseLevel(levelStr string) (zerolog.Level, error) {
	return zerolog.ParseLevel(levelStr)
}

func SetGlobalLevel(level zerolog.Level) {
	zerolog.SetGlobalLevel(level)
}
