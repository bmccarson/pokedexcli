package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bmccarson/pokedexcli/internal/commands"
	"github.com/bmccarson/pokedexcli/internal/state"
)

const (
	APIEndpoint       string        = "https://pokeapi.co/api/v2/"
	APICachePurgeTime time.Duration = time.Second * 30
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
	inputCommands := commands.Init()
	data := state.Init(APIEndpoint, APICachePurgeTime)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex> ")

		scanner.Scan()
		input := cleanInput(scanner.Text())

		command := input[0]
		arg := input[1]

		if key, exists := inputCommands[command]; exists {
			err := key.Callback(&data, arg)

			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("command does not exisist")
		}
	}
}
