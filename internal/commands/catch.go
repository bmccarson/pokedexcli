package commands

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bmccarson/pokedexcli/internal/pokeapi"
	"github.com/bmccarson/pokedexcli/internal/state"
)

func Catch(data *state.DataStore, pokemon string) error {
	p, err := pokeapi.GetPokemon(data.APIEndpoint, pokemon, data.APICache)

	if err != nil {
		return fmt.Errorf("Pokemon does not exsist: %s", pokemon)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", p.Name)

	result := attemptCatch(p)

	if result {
		fmt.Printf("%s was caught!\n", p.Name)
		data.PokemonContainer[p.Name] = p
	} else {
		fmt.Printf("%s got away!\n", p.Name)
	}

	return nil
}

func attemptCatch(pokemon pokeapi.Pokemon) bool {
	attempValue := rand.New(rand.NewSource(time.Now().UnixNano())).Float64()
	catchValue := 1 / (1 + float64(pokemon.BaseExperience)/100.0)

	if attempValue < catchValue {
		return true
	}
	return false
}
