package infrastructure

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

type Logger struct {
	*slog.Logger
	file *os.File
}

func CloseLogger(logger *Logger) error {
	err := logger.file.Close()
	if err != nil {
		return fmt.Errorf("closing file: %w", err)
	}

	return nil
}

func InitLogger() (*Logger, error) {
	absPath, err := filepath.Abs(filepath.Join("..", "..", "var", "logs.log"))
	if err != nil {
		return nil, fmt.Errorf("getting absolute path: %w", err)
	}

	file, err := os.OpenFile(absPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}

	logger := &Logger{
		slog.New(slog.NewTextHandler(file, nil)),
		file,
	}

	return logger, nil
}
