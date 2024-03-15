package slogrus

import (
	"log/slog"
	"os"
)

type LogLevel int

const (
	Debug LogLevel = -4
	Info  LogLevel = 0
	Warn  LogLevel = 4
	Error LogLevel = 8
)

type Logger struct {
	logger *slog.Logger
}

type Entry struct {
	attrs  []any
	logger *slog.Logger
}

type Fields map[string]interface{}

func New() *Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	return &Logger{logger: logger}
}

func NewWithHandler(handler *slog.Handler) *Logger {
	logger := slog.New(*handler)
	slog.SetDefault(logger)

	return &Logger{logger: logger}
}

func (l *Logger) WithFields(fields ...Fields) *Entry {
	attrs := make([]slog.Attr, 0, len(fields))
	if len(fields) > 0 {
		for k, v := range fields[0] {
			attrs = append(attrs, slog.Any(k, v))
		}
	}

	return &Entry{
		attrs:  convertToAny(attrs),
		logger: l.logger,
	}
}

func (l *Logger) Debug(message string) {
	l.logger.Debug(message)
}
func (l *Logger) Info(message string) {
	l.logger.Info(message)
}
func (l *Logger) Warn(message string) {
	l.logger.Warn(message)
}
func (l *Logger) Error(message string) {
	l.logger.Error(message)
}

func (e *Entry) Debug(message string) {
	e.logger.Debug(message, e.attrs...)
}

func (e *Entry) Info(message string) {
	e.logger.Info(message, e.attrs...)
}

func (e *Entry) Error(message string) {
	e.logger.Error(message, e.attrs...)
}

func (e *Entry) Warn(message string) {
	e.logger.Warn(message, e.attrs...)
}

func convertToAny(attrs []slog.Attr) []any {
	converted := make([]any, len(attrs))
	for i, attr := range attrs {
		converted[i] = any(attr)
	}
	return converted
}
