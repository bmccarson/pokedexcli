package commands

import (
	"fmt"

	"github.com/bmccarson/pokedexcli/internal/state"
)

func Help(_ *state.DataStore, _ string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	commands := Init()

	for _, v := range commands {
		fmt.Printf("%s: %s\n", v.Name, v.Description)
	}
	return nil
}
