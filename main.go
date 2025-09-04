package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
	},
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, v := range commands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}

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

		command := input[0]

		switch command {
		case "exit":
			commandExit()
		case "help":
			commandHelp()
		}
	}
}
