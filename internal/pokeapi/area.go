package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/bmccarson/pokedexcli/internal/pokecache"
)

type Area struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetArea(url string, cache *pokecache.PokeCache) (Area, error) {
	area := Area{}

	data, inCache := cache.Get(url)
	if inCache == false {
		res, err := http.Get(url)
		if err != nil {
			return area, err
		}
		data, err = io.ReadAll(res.Body)
		res.Body.Close()

		if res.StatusCode > 299 {
			return area, fmt.Errorf("Response failed: Status Code %d", res.StatusCode)
		}

		cache.Add(url, data)
	}

	err := json.Unmarshal(data, &area)
	if err != nil {
		return area, errors.New("Error converting JSON response")
	}

	return area, nil
}
