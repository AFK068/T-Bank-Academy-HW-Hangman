package cmd

import (
	"flag"
	"fmt"
	"log/slog"
	"slices"
	"strings"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

var (
	difficulty = flag.String("difficulty", "random", "Game difficulty level (random, easy, hard)")
	category   = flag.String("category", "random", "Category of words to use (random, Animals, Fruits, Cars, Cities, Countries, Hobbies)")
)

func ParseFlag(dwp *domain.DefaultWordProvider) (domain.Category, domain.Difficulty, error) {
	flag.Parse()

	inputCategory := domain.Category(strings.ToLower(*category))
	inputDifficulty := domain.Difficulty(strings.ToLower(*difficulty))

	if !slices.Contains(dwp.AllDifficulties, inputDifficulty) {
		slog.Info(
			"Difficulty not found in list of available values or flags is missing, so the default value is set - random",
			slog.String("difficulty", string(inputDifficulty)),
		)

		randomDifficulty, err := dwp.GetRandomDifficulty()
		if err != nil {
			slog.Error(
				"getting random difficulty",
				slog.String("error", err.Error()),
			)

			return "", "", fmt.Errorf("getting random difficulty: %w", err)
		}

		fmt.Printf("Difficulty not found in list of available values or flags is missing, so the default value is set - random \n")

		inputDifficulty = randomDifficulty
	}

	if !slices.Contains(dwp.AllCategories, inputCategory) {
		slog.Info(
			"Category not found in list of available values or flags is missing, so the default value is set - random",
			slog.String("category", string(inputCategory)),
		)

		randomCategory, err := dwp.GetRandomCategoryFromDifficulty(inputDifficulty)
		if err != nil {
			slog.Error(
				"getting random category",
				slog.String("error", err.Error()),
			)

			return "", "", fmt.Errorf("getting random category: %w", err)
		}

		fmt.Printf("Category not found in list of available values or flags is missing, so the default value is set - random \n")

		inputCategory = randomCategory
	}

	return inputCategory, inputDifficulty, nil
}
