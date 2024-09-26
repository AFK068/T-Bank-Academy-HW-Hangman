package application

import (
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/es-debug/backend-academy-2024-go-template/cmd"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
)

func InitializeGame() (*domain.Game, error) {
	absPath, err := filepath.Abs(filepath.Join("..", "..", "files", "words.json"))
	if err != nil {
		slog.Error("getting absolute path to words.json file", slog.String("error", err.Error()))
		return nil, fmt.Errorf("getting absolute path: %w", err)
	}

	provider, err := infrastructure.CreateProviderFromJSONFile(absPath)
	if err != nil {
		slog.Error("creating provider from JSON file", slog.String("error", err.Error()))
		return nil, fmt.Errorf("creating provider from JSON file: %w", err)
	}

	ctg, diff, err := cmd.ParseFlag(provider)
	if err != nil {
		slog.Error("parsing flags", slog.String("error", err.Error()))
		return nil, fmt.Errorf("parsing flags: %w", err)
	}

	wordAndHint, err := provider.GetRandomWordAndHintFromCategory(ctg, diff)
	if err != nil {
		slog.Error("getting random word and hint", slog.String("error", err.Error()))
		return nil, fmt.Errorf("getting random word and hint: %w", err)
	}

	game, err := domain.NewGame(wordAndHint, ctg, diff)
	if err != nil {
		slog.Error("creating new game", slog.String("error", err.Error()))
		return nil, fmt.Errorf("creating new game: %w", err)
	}

	return game, nil
}
