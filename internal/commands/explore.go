package commands

import (
	"fmt"

	"github.com/bmccarson/pokedexcli/internal/pokeapi"
	"github.com/bmccarson/pokedexcli/internal/state"
)

func Explore(data *state.DataStore, location string) error {
	loc, err := pokeapi.GetLocation(data.APIEndpoint, location, data.APICache)

	if err != nil {
		return err
	}

	for _, pokemon := range loc.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
