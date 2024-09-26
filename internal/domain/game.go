package domain

import (
	"strings"
)

const MaxAttempts int = 7

type Game struct {
	wordAndHint WordHintPair
	category    Category
	difficulty  Difficulty
	guesses     map[rune]bool
	attempts    int
	maxAttempts int
}

func NewGame(wordAndHint WordHintPair, ctg Category, diff Difficulty) (*Game, error) {
	wordAndHint.Word = strings.ToLower(wordAndHint.Word)
	wordAndHint.Hint = strings.ToLower(wordAndHint.Hint)

	if wordAndHint.Word == "" || wordAndHint.Hint == "" || ctg == "" || diff == "" {
		return nil, &InvalidLengthError{Message: "Invalid game parameters"}
	}

	return &Game{
		wordAndHint: wordAndHint,
		category:    ctg,
		difficulty:  diff,
		guesses:     make(map[rune]bool),
		attempts:    0,
		maxAttempts: max(1, MaxAttempts),
	}, nil
}

func (game *Game) GetDifficulty() Difficulty {
	return game.difficulty
}

func (game *Game) GetCategory() Category {
	return game.category
}

func (game *Game) GetWordAndHint() WordHintPair {
	return game.wordAndHint
}

func (game *Game) GetAttempts() int {
	return game.attempts
}

func (game *Game) GetMaxAttempts() int {
	return game.maxAttempts
}

func (game *Game) GetGuesses() map[rune]bool {
	return game.guesses
}

func (game *Game) SetGuesses(guesses map[rune]bool) {
	game.guesses = guesses
}

func (game *Game) SetAttempts(attempts int) {
	game.attempts = attempts
}

func (game *Game) SetMaxAttempts(maxAttempts int) {
	game.maxAttempts = maxAttempts
}

func (game *Game) GameIsOver() (isOver bool, message string) {
	if game.WordGuessed() {
		return true, "Word guessed. You win!"
	}

	if game.attempts >= game.maxAttempts {
		return true, "Max attempts reached. You lose! \n" + "Word: " + game.wordAndHint.Word
	}

	return false, ""
}

func (game *Game) GetWordWithGuesses() string {
	var wordWithGuesses strings.Builder

	for _, letter := range game.wordAndHint.Word {
		if game.guesses[letter] {
			wordWithGuesses.WriteRune(letter)
		} else {
			if letter == ' ' {
				wordWithGuesses.WriteRune(' ')
			} else {
				wordWithGuesses.WriteRune('_')
			}
		}
	}

	return wordWithGuesses.String()
}

func (game *Game) WordGuessed() bool {
	for _, letter := range game.wordAndHint.Word {
		if !game.guesses[letter] && letter != ' ' {
			return false
		}
	}

	return true
}

func (game *Game) LetterGuessed(letter rune) string {
	if game.guesses[letter] {
		return "Letter already guessed"
	}

	game.guesses[letter] = true
	letterGuessed := strings.Contains(game.wordAndHint.Word, string(letter))

	if !letterGuessed {
		game.attempts++
		return "Letter not in word"
	}

	return "Letter guessed"
}
