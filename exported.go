package slogrus

import (
	"log/slog"
	"os"
)

var (
	std = NewWithHandler(slog.NewTextHandler(os.Stdout, nil))
)

func WithFields(fields Fields) *Entry {
	return std.WithFields(fields)
}
