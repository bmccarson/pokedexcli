package commands

import (
	"errors"
	"fmt"

	"github.com/bmccarson/pokedexcli/internal/pokeapi"
	"github.com/bmccarson/pokedexcli/internal/state"
)

func Map(data *state.DataStore, _ string) error {
	area, err := pokeapi.GetArea(data.NextAreaURL, data.APICache)

	if err != nil {
		return err
	}

	for _, loc := range area.Results {
		fmt.Println(loc.Name)
	}

	data.NextAreaURL = area.Next
	data.PreviousAreaURL = area.Previous

	return nil
}

func Mapb(data *state.DataStore, _ string) error {
	if data.PreviousAreaURL == "" {
		return errors.New("can not go to previous from first location")
	}

	area, err := pokeapi.GetArea(data.PreviousAreaURL, data.APICache)

	if err != nil {
		return err
	}

	for _, loc := range area.Results {
		fmt.Println(loc.Name)
	}

	data.NextAreaURL = area.Next
	data.PreviousAreaURL = area.Previous

	return nil
}
