package infrastructure_test

import (
	"path/filepath"
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
	"github.com/es-debug/backend-academy-2024-go-template/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetLetterFromUser_success(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    rune
		wantErr bool
	}{
		{
			name:    "lowercase letter",
			input:   "a\n",
			want:    'a',
			wantErr: false,
		},
		{
			name:    "uppercase letter",
			input:   "A\n",
			want:    'a',
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restoreStdin, err := testutils.SimulateStdinInput(tt.input)
			require.NoError(t, err)
			defer restoreStdin()

			letter, err := infrastructure.GetLetterFromUser()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, letter)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, letter)
			}
		})
	}
}

func TestGetLetterFromUser_failure(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    rune
		wantErr bool
	}{
		{
			name:    "non-letter input",
			input:   "1\n",
			want:    0,
			wantErr: true,
		},
		{
			name:    "multiple characters",
			input:   "ab\n",
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restoreStdin, err := testutils.SimulateStdinInput(tt.input)
			require.NoError(t, err)
			defer restoreStdin()

			letter, err := infrastructure.GetLetterFromUser()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, letter)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, letter)
			}
		})
	}
}

func TestGameStateAfterEachInput(t *testing.T) {
	tests := []struct {
		name             string
		wordAndHint      domain.WordHintPair
		category         domain.Category
		difficulty       domain.Difficulty
		inputs           []rune
		expectedStates   []string
		expectedAttempts []int
	}{
		{
			name:             "correct guesses",
			wordAndHint:      domain.WordHintPair{Word: "apple", Hint: "A fruit"},
			category:         "fruits",
			difficulty:       "easy",
			inputs:           []rune{'a', 'p', 'l', 'e'},
			expectedStates:   []string{"a____", "app__", "appl_", "apple"},
			expectedAttempts: []int{0, 0, 0, 0},
		},
		{
			name:             "incorrect guesses",
			wordAndHint:      domain.WordHintPair{Word: "apple", Hint: "A fruit"},
			category:         "fruits",
			difficulty:       "easy",
			inputs:           []rune{'z', 'x', 'y'},
			expectedStates:   []string{"_____", "_____", "_____"},
			expectedAttempts: []int{1, 2, 3},
		},
		{
			name:             "mixed guesses",
			wordAndHint:      domain.WordHintPair{Word: "apple", Hint: "A fruit"},
			category:         "fruits",
			difficulty:       "easy",
			inputs:           []rune{'a', 'z', 'p', 'x', 'l', 'e'},
			expectedStates:   []string{"a____", "a____", "app__", "app__", "appl_", "apple"},
			expectedAttempts: []int{0, 1, 1, 2, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game, err := domain.NewGame(tt.wordAndHint, tt.category, tt.difficulty)
			require.NoError(t, err)

			var restoreStdinFuncs []func()

			for i, input := range tt.inputs {
				restoreStdin, err := testutils.SimulateStdinInput(string(input) + "\n")
				require.NoError(t, err)

				restoreStdinFuncs = append(restoreStdinFuncs, restoreStdin)

				game.LetterGuessed(input)
				assert.Equal(t, tt.expectedStates[i], game.GetWordWithGuesses())
				assert.Equal(t, tt.expectedAttempts[i], game.GetAttempts())
			}

			for _, restore := range restoreStdinFuncs {
				restore()
			}
		})
	}
}

func TestCreateProviderFromJSONFile_success(t *testing.T) {
	tests := []struct {
		name          string
		filePath      string
		expectedError string
	}{
		{
			name:          "valid JSON file",
			filePath:      filepath.Join("..", "..", "files", "words.json"),
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			absFilePath, err := filepath.Abs(tt.filePath)
			require.NoError(t, err)

			provider, err := infrastructure.CreateProviderFromJSONFile(absFilePath)
			if tt.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				assert.Nil(t, provider)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, provider)
			}
		})
	}
}

func TestCreateProviderFromJSONFile_failure(t *testing.T) {
	tests := []struct {
		name          string
		filePath      string
		expectedError string
	}{
		{
			name:          "non-existent file",
			filePath:      filepath.Join("..", "..", "files", "test", "test.json"),
			expectedError: "reading file",
		},
		{
			name:          "invalid JSON structure",
			filePath:      filepath.Join("..", "..", "files", "test", "words_invalid_test.json"),
			expectedError: "validating JSON",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			absFilePath, err := filepath.Abs(tt.filePath)
			require.NoError(t, err)

			provider, err := infrastructure.CreateProviderFromJSONFile(absFilePath)
			if tt.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				assert.Nil(t, provider)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, provider)
			}
		})
	}
}
