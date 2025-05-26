package logger

import (
	"io"
	"log/slog"
	"os"
	"strings"
)

type LoggerOptions struct {
	Level string
	Sink  io.Writer
}

func DefaultLoggerOptions() *LoggerOptions {
	return &LoggerOptions{
		Level: "INFO",
		Sink:  os.Stdout,
	}
}

func parseLogLevel(level string) slog.Leveler {
	switch strings.ToLower(level) {
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	case "debug":
		return slog.LevelDebug
	default:
		return slog.LevelInfo
	}
}

func New(opts *LoggerOptions) *slog.Logger {
	if opts == nil {
		opts = DefaultLoggerOptions()
	}
	handlerOpts := &slog.HandlerOptions{
		Level: parseLogLevel(opts.Level),
	}
	return slog.New(
		slog.NewJSONHandler(opts.Sink, handlerOpts),
	)
}
