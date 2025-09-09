package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/bmccarson/pokedexcli/internal/pokecache"
)

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func GetPokemon(url string, pokemon string, cache *pokecache.PokeCache) (Pokemon, error) {
	p := Pokemon{}

	data, inCache := cache.Get(url + "pokemon/" + pokemon)
	if inCache == false {
		res, err := http.Get(url + "pokemon/" + pokemon)
		if err != nil {
			return p, err
		}
		data, err = io.ReadAll(res.Body)
		res.Body.Close()

		if res.StatusCode > 299 {
			return p, fmt.Errorf("Response failed: Status Code %d", res.StatusCode)
		}

		cache.Add(url, data)
	}

	err := json.Unmarshal(data, &p)
	if err != nil {
		return p, errors.New("Error converting JSON response")
	}

	return p, nil
}
