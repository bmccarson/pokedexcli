package commands

import (
	"fmt"
	"os"

	"github.com/bmccarson/pokedexcli/internal/state"
)

func Exit(_ *state.DataStore) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
