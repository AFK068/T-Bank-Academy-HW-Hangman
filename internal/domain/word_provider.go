package domain

import (
	"fmt"
	"strings"

	"crypto/rand"
	"math/big"
)

type WordHintPair struct {
	Word string
	Hint string
}

type DefaultWordProvider struct {
	Words           map[Difficulty]map[Category][]WordHintPair
	AllDifficulties []Difficulty
	AllCategories   []Category
}

func (dwp *DefaultWordProvider) UpdateUniqueCategoriesAndDifficulties() error {
	if len(dwp.Words) == 0 {
		return &NotFoundError{Message: "no difficulty found in data"}
	}

	uniqueDifficulties := make(map[Difficulty]bool)
	uniqueCategories := make(map[Category]bool)

	for diff, categories := range dwp.Words {
		if len(categories) == 0 {
			return &NotFoundError{Message: "no category found in data"}
		}

		lowerDiff := Difficulty(strings.ToLower(string(diff)))
		if !uniqueDifficulties[lowerDiff] {
			dwp.AllDifficulties = append(dwp.AllDifficulties, lowerDiff)
			uniqueDifficulties[lowerDiff] = true
		}

		for ctg := range categories {
			lowerCtg := Category(strings.ToLower(string(ctg)))
			if !uniqueCategories[lowerCtg] {
				dwp.AllCategories = append(dwp.AllCategories, lowerCtg)
				uniqueCategories[lowerCtg] = true
			}
		}
	}

	dwp.NormalizeCase() // Normalize case of all words

	return nil
}

func (dwp *DefaultWordProvider) NormalizeCase() {
	normalizedWords := make(map[Difficulty]map[Category][]WordHintPair)

	for diff, categories := range dwp.Words {
		lowerDiff := Difficulty(strings.ToLower(string(diff)))

		if _, exists := normalizedWords[lowerDiff]; !exists {
			normalizedWords[lowerDiff] = make(map[Category][]WordHintPair)
		}

		for ctg, wordAndHintPairs := range categories {
			lowerCtg := Category(strings.ToLower(string(ctg)))
			normalizedWords[lowerDiff][lowerCtg] = wordAndHintPairs
		}
	}

	dwp.Words = normalizedWords
}

func (dwp *DefaultWordProvider) GetRandomDifficulty() (Difficulty, error) {
	difficulties := make([]Difficulty, 0, len(dwp.Words))
	for diff := range dwp.Words {
		difficulties = append(difficulties, diff)
	}

	if len(difficulties) == 0 {
		return "", &NotFoundError{Message: "no difficulty found to get random value"}
	}

	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(difficulties))))
	if err != nil {
		return "", err
	}

	return difficulties[nBig.Int64()], nil
}

func (dwp *DefaultWordProvider) GetRandomCategoryFromDifficulty(diff Difficulty) (Category, error) {
	categories := make([]Category, 0, len(dwp.Words[diff]))
	for ctg := range dwp.Words[diff] {
		categories = append(categories, ctg)
	}

	if len(categories) == 0 {
		return "", &NotFoundError{Message: "no category found to get random value"}
	}

	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(categories))))
	if err != nil {
		return "", err
	}

	return categories[nBig.Int64()], nil
}

func (dwp *DefaultWordProvider) GetRandomWordAndHintFromCategory(ctg Category, diff Difficulty) (WordHintPair, error) {
	words, exists := dwp.Words[diff][ctg]
	if !exists || len(words) == 0 {
		return WordHintPair{}, &NotFoundError{
			Message: fmt.Sprintf("No words and hints in category '%s' with difficulty '%s'", ctg, diff),
		}
	}

	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(words))))
	if err != nil {
		return WordHintPair{}, err
	}

	return words[nBig.Int64()], nil
}
