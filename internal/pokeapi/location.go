package pokeapi

import (
	"encoding/json"
	"net/http"
)

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLoactions(url string) (Locations, error) {
	locations := Locations{}

	res, err := http.Get(url)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return Locations{}, err
	}

	return locations, nil
}
