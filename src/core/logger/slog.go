package logger

import (
	"fmt"
	"log/slog"
	"my-go-template/src/core/config"
	"os"
)

var logger *slog.Logger

func init() {
	logLevel := new(slog.LevelVar)
	switch config.LogLevel() {
	case "debug":
		logLevel.Set(slog.LevelDebug)
	case "info":
	default:
		logLevel.Set(slog.LevelInfo)

	}
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
}

func Error(err error, arg ...any) {
	logger.Error(fmt.Sprintf("err: %s", err.Error()), arg...)
}

func Info(msg string, arg ...any) {
	logger.Info(msg, arg...)
}

func Debug(msg string, arg ...any) {
	logger.Debug(msg, arg...)
}
