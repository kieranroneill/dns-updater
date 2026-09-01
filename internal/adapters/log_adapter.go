package adapters

import (
	"log/slog"

	"gopkg.in/natefinch/lumberjack.v2"
)

type LogAdapter struct {
	logger *slog.Logger
}

/**
 * constructors
 */

func NewLogAdapter(logPath string) *LogAdapter {
	_lumberjack := &lumberjack.Logger{
		Filename: logPath,
		MaxSize:  10, // MB
		Compress: true,
	}
	handler := slog.NewTextHandler(_lumberjack, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: false,
	})

	return &LogAdapter{
		logger: slog.New(handler),
	}
}

/**
 * public methods
 */

func (l *LogAdapter) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

func (l *LogAdapter) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *LogAdapter) Warn(msg string, args ...any) {
	l.logger.Info(msg, args...)
}
