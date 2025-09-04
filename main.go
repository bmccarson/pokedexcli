package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	cleanedInput := []string{}

	words := strings.Fields(text)

	for _, word := range words {
		clean := strings.ToLower(word)

		cleanedInput = append(cleanedInput, clean)
	}

	return cleanedInput
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex> ")

		scanner.Scan()
		input := cleanInput(scanner.Text())

		word := input[0]

		fmt.Printf("Your command was: %s\n", string(word))
	}
}
