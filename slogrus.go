package slogrus

import (
	"fmt"
	"log/slog"
	"os"
)

type LogLevel int

const ErrorKey = "error"

const (
	Debug LogLevel = -4
	Info  LogLevel = 0
	Warn  LogLevel = 4
	Error LogLevel = 8
)

type Logger struct {
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

func NewWithHandler(handler slog.Handler) *Logger {
	logger := slog.New(handler)
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

func (l *Logger) WithField(key string, value interface{}) *Entry {
	attrs := make([]slog.Attr, 0, 1)
	attrs = append(attrs, slog.Any(key, value))

	return &Entry{
		attrs:  convertToAny(attrs),
		logger: l.logger,
	}
}

func (l *Logger) WithError(err error) *Entry {
	return l.WithField(ErrorKey, err)
}

func (l *Logger) GetSlogLogger() *slog.Logger {
	return l.logger
}

func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(fmt.Sprint(args...))
}
func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(fmt.Sprint(args...))
}
func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(fmt.Sprint(args...))
}
func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(fmt.Sprint(args...))
}

func convertToAny(attrs []slog.Attr) []any {
	converted := make([]any, len(attrs))
	for i, attr := range attrs {
		converted[i] = any(attr)
	}
	return converted
}
