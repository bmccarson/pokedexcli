package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/bmccarson/pokedexcli/internal/pokecache"
)

type Location struct {
	Id       int `json:"id"`
	Location struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"location"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func GetLocation(url string, loc string, cache *pokecache.PokeCache) (Location, error) {
	location := Location{}

	data, inCache := cache.Get(url + "location-area/" + loc)
	if inCache == false {
		res, err := http.Get(url)
		if err != nil {
			return location, err
		}
		data, err = io.ReadAll(res.Body)
		res.Body.Close()

		if res.StatusCode > 299 {
			return location, fmt.Errorf("Response failed: Status Code %d", res.StatusCode)
		}

		cache.Add(url, data)
		fmt.Println("Didnt use Cache")
	}

	err := json.Unmarshal(data, &location)
	if err != nil {
		return location, errors.New("Error converting JSON response")
	}

	return location, nil
}
