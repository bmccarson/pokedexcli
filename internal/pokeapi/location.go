package pokeapi

import (
	"encoding/json"
	"fmt"
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
	endpoint := fmt.Sprintf("%s/location-area/", url)

	res, err := http.Get(endpoint)
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
