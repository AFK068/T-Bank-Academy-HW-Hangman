package infrastructure

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/xeipuuv/gojsonschema"
)

func CreateProviderFromJSONFile(filePath string) (*domain.DefaultWordProvider, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		slog.Error("reading JSON file", slog.String("filePath", filePath), slog.String("error", err.Error()))
		return nil, fmt.Errorf("reading file: %w", err)
	}

	if err := validateJSON(&data); err != nil {
		slog.Error("validating JSON file", slog.String("filePath", filePath), slog.String("error", err.Error()))
		return nil, fmt.Errorf("validating JSON: %w", err)
	}

	var wordsList map[domain.Difficulty]map[domain.Category][]domain.WordHintPair

	if err := json.Unmarshal(data, &wordsList); err != nil {
		slog.Error("unmarshalling JSON file", slog.String("filePath", filePath), slog.String("error", err.Error()))
		return nil, fmt.Errorf("unmarshalling JSON: %w", err)
	}

	provider := &domain.DefaultWordProvider{Words: wordsList}
	if err := provider.UpdateUniqueCategoriesAndDifficulties(); err != nil {
		slog.Error(
			"updating unique categories and difficulties",
			slog.String("filePath", filePath),
			slog.String("error", err.Error()))

		return nil, fmt.Errorf("updating unique categories and difficulties: %w", err)
	}

	return provider, nil
}

func validateJSON(jsonData *[]byte) error {
	absPathSchemaJSONFile, err := filepath.Abs(filepath.Join("..", "..", "files", "schema.json"))
	if err != nil {
		slog.Error("getting schema path", slog.String("error", err.Error()))
		return fmt.Errorf("getting schema path: %w", err)
	}

	schemaData, err := os.ReadFile(absPathSchemaJSONFile)
	if err != nil {
		slog.Error("reading schema file", slog.String("error", err.Error()))
		return fmt.Errorf("reading schema file: %w", err)
	}

	documentLoader := gojsonschema.NewStringLoader(string(*jsonData))
	schemaLoader := gojsonschema.NewStringLoader(string(schemaData))

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		slog.Error("validating file with schema", slog.String("error", err.Error()))
		return fmt.Errorf("validating data: %w", err)
	}

	if !result.Valid() {
		var validationErrors string
		for _, desc := range result.Errors() {
			validationErrors += fmt.Sprintf("- %s\n", desc)
		}

		slog.Error("validating JSON", slog.String("error", validationErrors))

		return fmt.Errorf("validating JSON: \n%s", validationErrors)
	}

	return nil
}
