package infrastructure

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func GetLetterFromUser() (rune, error) {
	isAlpha := regexp.MustCompile(`^[A-Za-z]$`).MatchString
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter one letter: ")

		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		slog.Info("User input", slog.String("input", input))

		if err != nil {
			slog.Error("reading user input", slog.String("error", err.Error()))
			return 0, err
		}

		if len(input) == 1 && isAlpha(input) {
			return unicode.ToLower(rune(input[0])), nil
		}

		fmt.Println("Wrong input. Please enter one letter.")
	}
}
