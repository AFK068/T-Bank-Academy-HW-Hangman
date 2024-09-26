package application

import (
	"fmt"
	"log/slog"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
	"github.com/es-debug/backend-academy-2024-go-template/pkg/apperrors"
)

func ManageGame() {
	game, err := InitializeGame()
	if err != nil {
		slog.Error("initializing game", slog.String("error", err.Error()))
		fmt.Println("Error while initializing game. \nError: ", apperrors.UnwrapError(err))

		return
	}

	slog.Info("Game initialized", slog.String("word", game.GetWordAndHint().Word))
	RunGameLoop(game)
}

func RunGameLoop(game *domain.Game) {
	for game.GetAttempts() <= game.GetMaxAttempts() {
		infrastructure.PrintGameMenu(game)

		input, err := infrastructure.GetLetterFromUser()
		if err != nil {
			slog.Error("getting letter from user", slog.String("error", err.Error()))
			fmt.Println(err)

			return
		}

		message := game.LetterGuessed(input)
		fmt.Println(message)

		if gameIsOver, message := game.GameIsOver(); gameIsOver {
			infrastructure.PrintGameMenu(game) // Print the final state of the game
			fmt.Println(message)

			slog.Info("Game is over", slog.String("message to user", message))

			return
		}
	}
}
