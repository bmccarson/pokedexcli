package commands

import (
	"errors"
	"fmt"

	"github.com/bmccarson/pokedexcli/internal/pokeapi"
	"github.com/bmccarson/pokedexcli/internal/state"
)

func Map(data *state.DataStore) error {
	locations, err := pokeapi.GetLoactions(data.NextLocationURL)

	if err != nil {
		return err
	}

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	data.NextLocationURL = locations.Next
	data.PreviousLocationURL = locations.Previous

	return nil
}

func Mapb(data *state.DataStore) error {
	if data.PreviousLocationURL == "" {
		return errors.New("can not go to previous from first location")
	}

	locations, err := pokeapi.GetLoactions(data.PreviousLocationURL)

	if err != nil {
		return err
	}

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	data.NextLocationURL = locations.Next
	data.PreviousLocationURL = locations.Previous

	return nil
}
