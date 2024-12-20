package logger

import (
	"log/slog"
	"os"
)

// NewLogger returns new instance of logrus logger with defined log level and formatting
func NewLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stderr, opts))
	slog.SetDefault(logger)

	return logger
}

func Error(err error) slog.Attr {
	return slog.Any("error", err)
}
