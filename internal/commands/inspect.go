package commands

import (
	"fmt"

	"github.com/bmccarson/pokedexcli/internal/state"
)

func Inspect(data *state.DataStore, pokemon string) error {
	p, exists := data.PokemonContainer[pokemon]

	if exists {
		fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", p.Name, p.Height, p.Weight)
		fmt.Println("Stats:")
		for _, s := range p.Stats {
			fmt.Printf(" - %s: %d\n", s.Stat.Name, s.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range p.Types {
			fmt.Printf(" - %s\n", t.Type.Name)
		}
	} else {
		return fmt.Errorf("You have not caught %s", pokemon)
	}
	return nil
}
