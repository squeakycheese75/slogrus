package slogrus

import (
	"fmt"
	"log/slog"
)

type Entry struct {
	attrs  []any
	logger *slog.Logger
}

func (entry *Entry) Debug(args ...interface{}) {
	entry.logger.Debug(fmt.Sprint(args...))
}

func (entry *Entry) Info(args ...interface{}) {
	entry.logger.Info(fmt.Sprint(args...))
}

func (entry *Entry) Print(args ...interface{}) {
	entry.logger.Info(fmt.Sprint(args...))
}

func (entry *Entry) Error(args ...interface{}) {
	entry.logger.Info(fmt.Sprint(args...))
}

func (entry *Entry) Warn(args ...interface{}) {
	entry.logger.Warn(fmt.Sprint(args...))
}
