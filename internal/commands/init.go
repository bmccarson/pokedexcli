package commands

import "github.com/bmccarson/pokedexcli/internal/state"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(data *state.DataStore) error
}

const BaseURL = "https://pokeapi.co/api/v2"

func Init() map[string]CliCommand {
	commands := make(map[string]CliCommand)

	commands["map"] = CliCommand{
		Name:        "map",
		Description: "display map locations",
		Callback:    Map,
	}
	commands["mapb"] = CliCommand{
		Name:        "mapb",
		Description: "display previous map locations",
		Callback:    Mapb,
	}
	commands["exit"] = CliCommand{
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    Exit,
	}
	commands["help"] = CliCommand{
		Name:        "help",
		Description: "Displays a help message",
		Callback:    Help,
	}

	return commands
}
