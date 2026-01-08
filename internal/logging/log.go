package logging

import (
	"log/slog"
	"os"
)

var (
	defaultLogger *slog.Logger
	logLevel      = new(slog.LevelVar)
)

func SetupBaseLogger() {
	logLevel.Set(slog.LevelWarn)
	defaultLogger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: logLevel,
	}))
}

func ConfigureLogLevel(level slog.Level) {
	logLevel.Set(level)
}

func Debug(msg string, args ...any) {
	defaultLogger.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	defaultLogger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	defaultLogger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	defaultLogger.Error(msg, args...)
}
