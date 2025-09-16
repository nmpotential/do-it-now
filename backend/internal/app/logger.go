package app

import (
	"log/slog"
	"os"
)

// NewLogger returns a structured slog.Logger.
// In production, logs as JSON; in dev, logs are human-readable.
func NewLogger(env string) *slog.Logger {
	var handler slog.Handler
	if env == "production" {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		handler = slog.NewTextHandler(os.Stdout, nil)
	}
	return slog.New(handler)
}
