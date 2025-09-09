package commands

import (
	"errors"
	"fmt"

	"github.com/bmccarson/pokedexcli/internal/state"
)

func Pokedex(data *state.DataStore, _ string) error {
	if len(data.PokemonContainer) == 0 {
		return errors.New("You have not caught any Pokemon yet")
	}

	for _, v := range data.PokemonContainer {
		fmt.Printf("- %s\n", v.Name)
	}
	return nil
}
