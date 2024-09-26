package main

import (
	"fmt"
	"log/slog"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
)

func main() {
	logger, err := infrastructure.InitLogger()
	if err != nil {
		fmt.Println("failed to initialize logger:", err)
		return
	}

	slog.SetDefault(logger.Logger)

	defer func() {
		if err := infrastructure.CloseLogger(logger); err != nil {
			fmt.Println("failed to close logger:", err)
		}
	}()

	application.ManageGame()
}
