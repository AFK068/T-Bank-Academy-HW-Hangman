package integration_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGame_LetterGuessed_success(t *testing.T) {
	type args struct {
		letter rune
	}

	tests := []struct {
		name  string
		setup func() *domain.Game
		args  args
		want  string
	}{
		{
			name: "correct guess",
			setup: func() *domain.Game {
				game, _ := domain.NewGame(domain.WordHintPair{Word: "apple", Hint: "A fruit"}, "fruits", "easy")
				game.SetGuesses(map[rune]bool{'a': true, 'p': true})
				game.SetAttempts(2)
				game.SetMaxAttempts(5)
				return game
			},
			args: args{letter: 'l'},
			want: "Letter guessed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := tt.setup()
			if got := game.LetterGuessed(tt.args.letter); got != tt.want {
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func TestGame_LetterGuessed_failure(t *testing.T) {
	type args struct {
		letter rune
	}

	tests := []struct {
		name  string
		setup func() *domain.Game
		args  args
		want  string
	}{
		{
			name: "incorrect guess",
			setup: func() *domain.Game {
				game, _ := domain.NewGame(domain.WordHintPair{Word: "apple", Hint: "A fruit"}, "fruits", "easy")
				game.SetGuesses(map[rune]bool{'a': true, 'p': true})
				game.SetAttempts(2)
				game.SetMaxAttempts(5)
				return game
			},
			args: args{letter: 'z'},
			want: "Letter not in word",
		},
		{
			name: "already guessed letter",
			setup: func() *domain.Game {
				game, _ := domain.NewGame(domain.WordHintPair{Word: "apple", Hint: "A fruit"}, "fruits", "easy")
				game.SetGuesses(map[rune]bool{'a': true, 'p': true, 'l': true})
				game.SetAttempts(3)
				game.SetMaxAttempts(5)
				return game
			},
			args: args{letter: 'l'},
			want: "Letter already guessed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := tt.setup()
			if got := game.LetterGuessed(tt.args.letter); got != tt.want {
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func TestGame_GetWordWithGuesses(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *domain.Game
		want  string
	}{
		{
			name: "all letters guessed",
			setup: func() *domain.Game {
				game, _ := domain.NewGame(domain.WordHintPair{Word: "apple", Hint: "A fruit"}, "fruits", "easy")
				game.SetGuesses(map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true})
				game.SetAttempts(0)
				game.SetMaxAttempts(5)
				return game
			},
			want: "apple",
		},
		{
			name: "no letters guessed",
			setup: func() *domain.Game {
				game, _ := domain.NewGame(domain.WordHintPair{Word: "apple", Hint: "A fruit"}, "fruits", "easy")
				game.SetGuesses(map[rune]bool{})
				game.SetAttempts(0)
				game.SetMaxAttempts(5)
				return game
			},
			want: "_____",
		},
		{
			name: "some letters guessed",
			setup: func() *domain.Game {
				game, _ := domain.NewGame(domain.WordHintPair{Word: "apple", Hint: "A fruit"}, "fruits", "easy")
				game.SetGuesses(map[rune]bool{'a': true, 'p': true})
				game.SetAttempts(0)
				game.SetMaxAttempts(5)
				return game
			},
			want: "app__",
		},
		{
			name: "word with spaces",
			setup: func() *domain.Game {
				game, _ := domain.NewGame(domain.WordHintPair{Word: "a b", Hint: "A phrase"}, "phrases", "easy")
				game.SetGuesses(map[rune]bool{'a': true})
				game.SetAttempts(0)
				game.SetMaxAttempts(5)
				return game
			},
			want: "a _",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := tt.setup()
			got := game.GetWordWithGuesses()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestGame_WordGuessed_success(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *domain.Game
		want  bool
	}{
		{
			name: "all letters guessed",
			setup: func() *domain.Game {
				game, _ := domain.NewGame(domain.WordHintPair{Word: "apple", Hint: "A fruit"}, "fruits", "easy")
				game.SetGuesses(map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true})
				return game
			},
			want: true,
		},
		{
			name: "word with spaces, all letters guessed",
			setup: func() *domain.Game {
				game, _ := domain.NewGame(domain.WordHintPair{Word: "a b", Hint: "A phrase"}, "phrases", "easy")
				game.SetGuesses(map[rune]bool{'a': true, 'b': true})
				return game
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := tt.setup()
			got := game.WordGuessed()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestGame_WordGuessed_failure(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *domain.Game
		want  bool
	}{
		{
			name: "no letters guessed",
			setup: func() *domain.Game {
				game, _ := domain.NewGame(domain.WordHintPair{Word: "apple", Hint: "A fruit"}, "fruits", "easy")
				game.SetGuesses(map[rune]bool{})
				return game
			},
			want: false,
		},
		{
			name: "some letters guessed",
			setup: func() *domain.Game {
				game, _ := domain.NewGame(domain.WordHintPair{Word: "apple", Hint: "A fruit"}, "fruits", "easy")
				game.SetGuesses(map[rune]bool{'a': true, 'p': true})
				return game
			},
			want: false,
		},
		{
			name: "word with spaces, some letters guessed",
			setup: func() *domain.Game {
				game, _ := domain.NewGame(domain.WordHintPair{Word: "a b", Hint: "A phrase"}, "phrases", "easy")
				game.SetGuesses(map[rune]bool{'a': true})
				return game
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := tt.setup()
			got := game.WordGuessed()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestGame_GameIsOver_success(t *testing.T) {
	tests := []struct {
		name        string
		setup       func() *domain.Game
		wantIsOver  bool
		wantMessage string
	}{
		{
			name: "word guessed",
			setup: func() *domain.Game {
				game, _ := domain.NewGame(domain.WordHintPair{Word: "apple", Hint: "A fruit"}, "fruits", "easy")
				game.SetGuesses(map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true})
				return game
			},
			wantIsOver:  true,
			wantMessage: "Word guessed. You win!",
		},
		{
			name: "max attempts reached",
			setup: func() *domain.Game {
				game, _ := domain.NewGame(domain.WordHintPair{Word: "apple", Hint: "A fruit"}, "fruits", "easy")
				game.SetGuesses(map[rune]bool{})
				game.SetAttempts(5)
				game.SetMaxAttempts(5)
				return game
			},
			wantIsOver:  true,
			wantMessage: "Max attempts reached. You lose! \nWord: apple",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := tt.setup()
			gotIsOver, gotMessage := game.GameIsOver()
			require.Equal(t, tt.wantIsOver, gotIsOver)
			require.Equal(t, tt.wantMessage, gotMessage)
		})
	}
}

func TestGame_GameIsOver_failure(t *testing.T) {
	tests := []struct {
		name        string
		setup       func() *domain.Game
		wantIsOver  bool
		wantMessage string
	}{
		{
			name: "game continues",
			setup: func() *domain.Game {
				game, _ := domain.NewGame(domain.WordHintPair{Word: "apple", Hint: "A fruit"}, "fruits", "easy")
				game.SetGuesses(map[rune]bool{'a': true, 'p': true})
				game.SetAttempts(2)
				game.SetMaxAttempts(5)
				return game
			},
			wantIsOver:  false,
			wantMessage: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := tt.setup()
			gotIsOver, gotMessage := game.GameIsOver()
			require.Equal(t, tt.wantIsOver, gotIsOver)
			require.Equal(t, tt.wantMessage, gotMessage)
		})
	}
}

func TestGame_NewGame_success(t *testing.T) {
	tests := []struct {
		name        string
		wordAndHint domain.WordHintPair
		category    domain.Category
		difficulty  domain.Difficulty
		wantErr     bool
	}{
		{
			name:        "valid word length",
			wordAndHint: domain.WordHintPair{Word: "apple", Hint: "A fruit"},
			category:    "fruits",
			difficulty:  "easy",
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := domain.NewGame(tt.wordAndHint, tt.category, tt.difficulty)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestGame_NewGame_failure(t *testing.T) {
	tests := []struct {
		name        string
		wordAndHint domain.WordHintPair
		category    domain.Category
		difficulty  domain.Difficulty
		wantErr     bool
	}{
		{
			name:        "word empty",
			wordAndHint: domain.WordHintPair{Word: "", Hint: "Too short"},
			category:    "fruits",
			difficulty:  "easy",
			wantErr:     true,
		},
		{
			name:        "hint empty",
			wordAndHint: domain.WordHintPair{Word: "thisisaverylongword", Hint: ""},
			category:    "fruits",
			difficulty:  "easy",
			wantErr:     true,
		},
		{
			name:        "empty category",
			wordAndHint: domain.WordHintPair{Word: "apple", Hint: "A fruit"},
			category:    "",
			difficulty:  "easy",
			wantErr:     true,
		},
		{
			name:        "empty difficulty",
			wordAndHint: domain.WordHintPair{Word: "apple", Hint: "A fruit"},
			category:    "fruits",
			difficulty:  "",
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := domain.NewGame(tt.wordAndHint, tt.category, tt.difficulty)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDefaultWordProvider_GetRandomDifficulty_success(t *testing.T) {
	tests := []struct {
		name    string
		words   map[domain.Difficulty]map[domain.Category][]domain.WordHintPair
		wantErr bool
	}{
		{
			name: "successful case with difficulties",
			words: map[domain.Difficulty]map[domain.Category][]domain.WordHintPair{
				domain.Difficulty("easy"):   {},
				domain.Difficulty("medium"): {},
				domain.Difficulty("hard"):   {},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dwp := &domain.DefaultWordProvider{
				Words: tt.words,
			}

			got, err := dwp.GetRandomDifficulty()

			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "no difficulty found")
			} else {
				require.NoError(t, err)

				_, exists := tt.words[got]
				assert.True(t, exists, "returned difficulty should exist in the map")
			}
		})
	}
}

func TestDefaultWordProvider_GetRandomDifficulty_failure(t *testing.T) {
	tests := []struct {
		name    string
		words   map[domain.Difficulty]map[domain.Category][]domain.WordHintPair
		wantErr bool
	}{
		{
			name:    "no difficulties found",
			words:   map[domain.Difficulty]map[domain.Category][]domain.WordHintPair{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dwp := &domain.DefaultWordProvider{
				Words: tt.words,
			}

			got, err := dwp.GetRandomDifficulty()

			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "no difficulty found")
			} else {
				require.NoError(t, err)

				_, exists := tt.words[got]
				assert.True(t, exists, "returned difficulty should exist in the map")
			}
		})
	}
}

func TestDefaultWordProvider_GetRandomCategoryFromDifficulty_success(t *testing.T) {
	tests := []struct {
		name    string
		words   map[domain.Difficulty]map[domain.Category][]domain.WordHintPair
		diff    domain.Difficulty
		wantErr bool
	}{
		{
			name: "successful case with categories",
			words: map[domain.Difficulty]map[domain.Category][]domain.WordHintPair{
				domain.Difficulty("easy"): {
					domain.Category("animals"): {},
					domain.Category("fruits"):  {},
				},
			},
			diff:    domain.Difficulty("easy"),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dwp := &domain.DefaultWordProvider{
				Words: tt.words,
			}

			got, err := dwp.GetRandomCategoryFromDifficulty(tt.diff)

			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "no category found")
			} else {
				require.NoError(t, err)

				_, exists := tt.words[tt.diff][got]
				assert.True(t, exists, "returned category should exist in the map for given difficulty")
			}
		})
	}
}

func TestDefaultWordProvider_GetRandomCategoryFromDifficulty_failure(t *testing.T) {
	tests := []struct {
		name    string
		words   map[domain.Difficulty]map[domain.Category][]domain.WordHintPair
		diff    domain.Difficulty
		wantErr bool
	}{
		{
			name: "difficulty not found",
			words: map[domain.Difficulty]map[domain.Category][]domain.WordHintPair{
				domain.Difficulty("easy"): {
					domain.Category("animals"): {},
				},
			},
			diff:    domain.Difficulty("hard"),
			wantErr: true,
		},
		{
			name: "no categories in difficulty",
			words: map[domain.Difficulty]map[domain.Category][]domain.WordHintPair{
				domain.Difficulty("easy"): {},
			},
			diff:    domain.Difficulty("easy"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dwp := &domain.DefaultWordProvider{
				Words: tt.words,
			}

			got, err := dwp.GetRandomCategoryFromDifficulty(tt.diff)

			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "no category found")
			} else {
				require.NoError(t, err)

				_, exists := tt.words[tt.diff][got]
				assert.True(t, exists, "returned category should exist in the map for given difficulty")
			}
		})
	}
}

func TestDefaultWordProvider_GetRandomWordAndHintFromCategory_success(t *testing.T) {
	type fields struct {
		Words map[domain.Difficulty]map[domain.Category][]domain.WordHintPair
	}

	type args struct {
		ctg  domain.Category
		diff domain.Difficulty
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful case",
			fields: fields{
				Words: map[domain.Difficulty]map[domain.Category][]domain.WordHintPair{
					domain.Difficulty("easy"): {
						domain.Category("animals"): {
							{Word: "cat", Hint: "a small domesticated carnivorous mammal"},
							{Word: "dog", Hint: "a domesticated carnivorous mammal"},
						},
					},
				},
			},
			args: args{
				ctg:  domain.Category("animals"),
				diff: domain.Difficulty("easy"),
			},
			wantErr: false,
		},
		{
			name: "single word in category",
			fields: fields{
				Words: map[domain.Difficulty]map[domain.Category][]domain.WordHintPair{
					domain.Difficulty("easy"): {
						domain.Category("animals"): {
							{Word: "cat", Hint: "a small domesticated carnivorous mammal"},
						},
					},
				},
			},
			args: args{
				ctg:  domain.Category("animals"),
				diff: domain.Difficulty("easy"),
			},
			wantErr: false,
		},
		{
			name: "multiple words in category",
			fields: fields{
				Words: map[domain.Difficulty]map[domain.Category][]domain.WordHintPair{
					domain.Difficulty("easy"): {
						domain.Category("animals"): {
							{Word: "cat", Hint: "a small domesticated carnivorous mammal"},
							{Word: "dog", Hint: "a domesticated carnivorous mammal"},
						},
					},
				},
			},
			args: args{
				ctg:  domain.Category("animals"),
				diff: domain.Difficulty("easy"),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dwp := &domain.DefaultWordProvider{
				Words: tt.fields.Words,
			}

			got, err := dwp.GetRandomWordAndHintFromCategory(tt.args.ctg, tt.args.diff)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Contains(t, tt.fields.Words[tt.args.diff][tt.args.ctg], got)
			}
		})
	}
}

func TestDefaultWordProvider_GetRandomWordAndHintFromCategory_failure(t *testing.T) {
	type fields struct {
		Words map[domain.Difficulty]map[domain.Category][]domain.WordHintPair
	}

	type args struct {
		ctg  domain.Category
		diff domain.Difficulty
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "category not found",
			fields: fields{
				Words: map[domain.Difficulty]map[domain.Category][]domain.WordHintPair{
					domain.Difficulty("easy"): {
						domain.Category("animals"): {
							{Word: "cat", Hint: "a small domesticated carnivorous mammal"},
						},
					},
				},
			},
			args: args{
				ctg:  domain.Category("fruits"),
				diff: domain.Difficulty("easy"),
			},
			wantErr: true,
		},
		{
			name: "difficulty not found",
			fields: fields{
				Words: map[domain.Difficulty]map[domain.Category][]domain.WordHintPair{
					domain.Difficulty("easy"): {
						domain.Category("animals"): {
							{Word: "cat", Hint: "a small domesticated carnivorous mammal"},
						},
					},
				},
			},
			args: args{
				ctg:  domain.Category("animals"),
				diff: domain.Difficulty("hard"),
			},
			wantErr: true,
		},
		{
			name: "empty category",
			fields: fields{
				Words: map[domain.Difficulty]map[domain.Category][]domain.WordHintPair{
					domain.Difficulty("easy"): {
						domain.Category("animals"): {
							{Word: "cat", Hint: "a small domesticated carnivorous mammal"},
						},
					},
				},
			},
			args: args{
				ctg:  domain.Category(""),
				diff: domain.Difficulty("easy"),
			},
			wantErr: true,
		},
		{
			name: "empty difficulty",
			fields: fields{
				Words: map[domain.Difficulty]map[domain.Category][]domain.WordHintPair{
					domain.Difficulty("easy"): {
						domain.Category("animals"): {
							{Word: "cat", Hint: "a small domesticated carnivorous mammal"},
						},
					},
				},
			},
			args: args{
				ctg:  domain.Category("animals"),
				diff: domain.Difficulty(""),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dwp := &domain.DefaultWordProvider{
				Words: tt.fields.Words,
			}

			got, err := dwp.GetRandomWordAndHintFromCategory(tt.args.ctg, tt.args.diff)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Contains(t, tt.fields.Words[tt.args.diff][tt.args.ctg], got)
			}
		})
	}
}

func TestDefaultWordProvider_NormalizeCase(t *testing.T) {
	tests := []struct {
		name     string
		input    map[domain.Difficulty]map[domain.Category][]domain.WordHintPair
		expected map[domain.Difficulty]map[domain.Category][]domain.WordHintPair
	}{
		{
			name: "mixed case",
			input: map[domain.Difficulty]map[domain.Category][]domain.WordHintPair{
				"Easy": {
					"Fruits": {
						{Word: "Apple", Hint: "A fruit"},
					},
				},
				"HARD": {
					"VEGETABLES": {
						{Word: "Carrot", Hint: "A vegetable"},
					},
				},
			},
			expected: map[domain.Difficulty]map[domain.Category][]domain.WordHintPair{
				"easy": {
					"fruits": {
						{Word: "Apple", Hint: "A fruit"},
					},
				},
				"hard": {
					"vegetables": {
						{Word: "Carrot", Hint: "A vegetable"},
					},
				},
			},
		},
		{
			name: "already normalized",
			input: map[domain.Difficulty]map[domain.Category][]domain.WordHintPair{
				"easy": {
					"fruits": {
						{Word: "Apple", Hint: "A fruit"},
					},
				},
				"hard": {
					"vegetables": {
						{Word: "Carrot", Hint: "A vegetable"},
					},
				},
			},
			expected: map[domain.Difficulty]map[domain.Category][]domain.WordHintPair{
				"easy": {
					"fruits": {
						{Word: "Apple", Hint: "A fruit"},
					},
				},
				"hard": {
					"vegetables": {
						{Word: "Carrot", Hint: "A vegetable"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dwp := &domain.DefaultWordProvider{
				Words: tt.input,
			}
			dwp.NormalizeCase()
			assert.Equal(t, tt.expected, dwp.Words)
		})
	}
}

func TestDefaultWordProvider_UpdateUniqueCategoriesAndDifficulties_success(t *testing.T) {
	tests := []struct {
		name          string
		input         map[domain.Difficulty]map[domain.Category][]domain.WordHintPair
		expectedDiffs []domain.Difficulty
		expectedCats  []domain.Category
		expectedError error
	}{
		{
			name: "mixed case",
			input: map[domain.Difficulty]map[domain.Category][]domain.WordHintPair{
				"Easy": {
					"Fruits": {
						{Word: "Apple", Hint: "A fruit"},
					},
				},
				"HARD": {
					"VEGETABLES": {
						{Word: "Carrot", Hint: "A vegetable"},
					},
				},
			},
			expectedDiffs: []domain.Difficulty{"easy", "hard"},
			expectedCats:  []domain.Category{"fruits", "vegetables"},
			expectedError: nil,
		},
		{
			name: "duplicate entries",
			input: map[domain.Difficulty]map[domain.Category][]domain.WordHintPair{
				"Easy": {
					"Fruits": {
						{Word: "Apple", Hint: "A fruit"},
					},
				},
				"easy": {
					"fruits": {
						{Word: "Banana", Hint: "Another fruit"},
					},
				},
			},
			expectedDiffs: []domain.Difficulty{"easy"},
			expectedCats:  []domain.Category{"fruits"},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dwp := &domain.DefaultWordProvider{
				Words: tt.input,
			}
			err := dwp.UpdateUniqueCategoriesAndDifficulties()

			if tt.expectedError != nil {
				require.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				require.NoError(t, err)
				assert.ElementsMatch(t, tt.expectedDiffs, dwp.AllDifficulties)
				assert.ElementsMatch(t, tt.expectedCats, dwp.AllCategories)
			}
		})
	}
}

func TestDefaultWordProvider_UpdateUniqueCategoriesAndDifficulties_failure(t *testing.T) {
	tests := []struct {
		name          string
		input         map[domain.Difficulty]map[domain.Category][]domain.WordHintPair
		expectedDiffs []domain.Difficulty
		expectedCats  []domain.Category
		expectedError error
	}{
		{
			name:          "empty difficulties",
			input:         map[domain.Difficulty]map[domain.Category][]domain.WordHintPair{},
			expectedDiffs: []domain.Difficulty{},
			expectedCats:  []domain.Category{},
			expectedError: &domain.NotFoundError{Message: "no difficulty found in data"},
		},
		{
			name: "empty categories",
			input: map[domain.Difficulty]map[domain.Category][]domain.WordHintPair{
				"Easy": {},
			},
			expectedDiffs: []domain.Difficulty{},
			expectedCats:  []domain.Category{},
			expectedError: &domain.NotFoundError{Message: "no category found in data"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dwp := &domain.DefaultWordProvider{
				Words: tt.input,
			}
			err := dwp.UpdateUniqueCategoriesAndDifficulties()

			if tt.expectedError != nil {
				require.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				require.NoError(t, err)
				assert.ElementsMatch(t, tt.expectedDiffs, dwp.AllDifficulties)
				assert.ElementsMatch(t, tt.expectedCats, dwp.AllCategories)
			}
		})
	}
}
