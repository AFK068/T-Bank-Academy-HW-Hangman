package infrastructure

import (
	"fmt"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func PrintGameMenu(game *domain.Game) {
	titleCaser := cases.Title(language.Und, cases.NoLower)
	attempts := game.GetAttempts()
	maxAttempts := game.GetMaxAttempts()

	category := titleCaser.String(string(game.GetCategory()))
	difficulty := titleCaser.String(string(game.GetDifficulty()))

	fmt.Println()
	fmt.Println("╔════════════════════════════════════════════════╗")
	fmt.Println("║                  Hangman Game                  ║")
	fmt.Println("╠════════════════════════════════════════════════╣")

	fmt.Printf("║ Attempts: %d/%-*d ║\n", attempts, 35-len(fmt.Sprintf("%d", attempts)), maxAttempts)
	fmt.Printf("║ Difficulty: %-34s ║\n", truncateString(difficulty, 34))

	fmt.Println("╠════════════════════════════════════════════════╣")

	fmt.Printf("║ Category: %-36s ║\n", truncateString(category, 36))
	fmt.Printf("║ Word: %-40s ║\n", game.GetWordWithGuesses())

	if game.GetAttempts() >= game.GetMaxAttempts()/2 {
		fmt.Println("╠════════════════════════════════════════════════╣")
		fmt.Printf("║ Hint: %-40s ║\n", truncateString(game.GetWordAndHint().Hint, 40))
	}

	printHangmanStage(attempts, maxAttempts)
}

// PrintHangmanStage prints the hangman stage based on the number of attempts and the maximum number of attempts.
func printHangmanStage(attempts, maxAttempts int) {
	if maxAttempts <= 0 {
		maxAttempts = 1
	}

	totalStages := len(hangmanStages)
	if totalStages == 0 {
		return
	}

	stagesPerAttempt := float64(totalStages-1) / float64(maxAttempts)
	stageIndex := int(stagesPerAttempt * float64(attempts))

	if attempts >= maxAttempts {
		stageIndex = totalStages - 1
	}

	fmt.Println(hangmanStages[stageIndex])
}

func truncateString(str string, maxLen int) string {
	if len(str) > maxLen {
		return str[:maxLen-3] + "..."
	}

	return str
}

var hangmanStages = []string{
	`╠════════════════════════════════════════════════╣
║ +---+                                          ║
║ |                                              ║
║ |                                              ║
║ |                                              ║
║ |                                              ║
║/|\                                             ║
╚════════════════════════════════════════════════╝`,
	`╠════════════════════════════════════════════════╣
║ +---+                                          ║
║ |   |                                          ║
║ |                                              ║
║ |                                              ║
║ |                                              ║
║/|\                                             ║
╚════════════════════════════════════════════════╝`,
	`╠════════════════════════════════════════════════╣
║ +---+                                          ║
║ |   |                                          ║
║ |   0                                          ║
║ |                                              ║
║ |                                              ║
║/|\                                             ║
╚════════════════════════════════════════════════╝`,
	`╠════════════════════════════════════════════════╣
║ +---+                                          ║
║ |   |                                          ║
║ |   0                                          ║
║ |   |                                          ║
║ |                                              ║
║/|\                                             ║
╚════════════════════════════════════════════════╝`,
	`╠════════════════════════════════════════════════╣
║ +---+                                          ║
║ |   |                                          ║
║ |   0                                          ║
║ |  /|                                          ║
║ |                                              ║
║/|\                                             ║
╚════════════════════════════════════════════════╝`,
	`╠════════════════════════════════════════════════╣
║ +---+                                          ║
║ |   |                                          ║
║ |   0                                          ║
║ |  /|\                                         ║
║ |                                              ║
║/|\                                             ║
╚════════════════════════════════════════════════╝`,
	`╠════════════════════════════════════════════════╣
║ +---+                                          ║
║ |   |                                          ║
║ |   0                                          ║
║ |  /|\                                         ║
║ |  /                                           ║
║/|\                                             ║
╚════════════════════════════════════════════════╝`,
	`╠════════════════════════════════════════════════╣
║ +---+                                          ║
║ |   |                                          ║
║ |   0                                          ║
║ |  /|\                                         ║
║ |  / \                                         ║
║/|\                                             ║
╚════════════════════════════════════════════════╝`,
}
