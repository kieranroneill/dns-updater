package utilities

import (
	"log/slog"
	"os"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	logger    *slog.Logger
	Directory string
	Filename  string
}

/**
 * constructors
 */

func NewLogger() (*Logger, error) {
	filename := "dns-updater.log"
	directory, err := logDirectory()
	if err != nil {
		return nil, err
	}

	_lumberjack := &lumberjack.Logger{
		Filename: filepath.Join(directory, filename),
		MaxSize:  10, // MB
		Compress: true,
	}

	handler := slog.NewTextHandler(_lumberjack, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: false,
	})

	return &Logger{
		logger:    slog.New(handler),
		Directory: directory,
		Filename:  filename,
	}, nil
}

/**
 * private static methods
 */

// logDirectory returns the directory where logs are stored.
//
// If the directory does not exist, it is created.
//
// Returns:
//   - The directory where logs are stored.
//   - An error if the directory could not be created.
func logDirectory() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	directory := filepath.Join(home, ".dns-updater", "logs")
	if err = os.MkdirAll(directory, 0o700); err != nil {
		return "", err
	}

	return directory, nil
}

/**
 * public methods
 */

func (l *Logger) FilePath() string {
	return filepath.Join(l.Directory, l.Filename)
}

func (l *Logger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.logger.Info(msg, args...)
}
