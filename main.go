package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bmccarson/pokedexcli/internal/commands"
	"github.com/bmccarson/pokedexcli/internal/state"
)

const APIEndpoint = "https://pokeapi.co/api/v2/"

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
	inputCommands := commands.Init()
	data := state.Init(APIEndpoint, 5)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex> ")

		scanner.Scan()
		input := cleanInput(scanner.Text())

		command := input[0]

		if key, exists := inputCommands[command]; exists {
			err := key.Callback(&data)

			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("command does not exisist")
		}
	}
}
