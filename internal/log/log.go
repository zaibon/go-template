// internal/log/log.go
package log

import (
	"log/slog"
	"os"
	"strings"
)

type LoggerOption func(*loggerOptions)

type loggerOptions struct {
	format string
	output string
	level  slog.Level
}

func defaultLoggerOptions() *loggerOptions {
	return &loggerOptions{
		format: "json",
		output: "stdout",
		level:  slog.LevelInfo,
	}
}

func WithFormat(format string) LoggerOption {
	return func(o *loggerOptions) {
		o.format = format
	}
}

func WithOutput(output string) LoggerOption {
	return func(o *loggerOptions) {
		o.output = output
	}
}

func WithLevel(level string) LoggerOption {
	return func(o *loggerOptions) {
		switch strings.ToLower(level) {
		case "debug":
			o.level = slog.LevelDebug
		case "info":
			o.level = slog.LevelInfo
		case "warn":
			o.level = slog.LevelWarn
		case "error":
			o.level = slog.LevelError
		default:
			o.level = slog.LevelInfo
		}
	}
}

func NewLogger(opts ...LoggerOption) *slog.Logger {
	options := defaultLoggerOptions()
	for _, opt := range opts {
		opt(options)
	}

	var handler slog.Handler

	switch strings.ToLower(options.format) {
	case "json":
		handler = slog.NewJSONHandler(getOutputWriter(options.output), &slog.HandlerOptions{Level: options.level})
	case "text":
		handler = slog.NewTextHandler(getOutputWriter(options.output), &slog.HandlerOptions{Level: options.level})
	default:
		handler = slog.NewJSONHandler(getOutputWriter(options.output), &slog.HandlerOptions{Level: options.level}) // Default to JSON
	}

	return slog.New(handler)
}

func getOutputWriter(output string) *os.File {
	if strings.HasPrefix(output, "file:") {
		filePath := strings.TrimPrefix(output, "file:")
		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			slog.Warn("Failed to open log file, using stdout", "error", err)
			return os.Stdout
		}
		return file
	}
	return os.Stdout // Default to stdout
}
