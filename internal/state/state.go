package state

import (
	"time"

	"github.com/bmccarson/pokedexcli/internal/pokecache"
)

type DataStore struct {
	APIEndpoint     string
	APICache        *pokecache.PokeCache
	NextAreaURL     string
	PreviousAreaURL string
}

func Init(apiEndpoint string, cachePurgeInterval time.Duration) DataStore {
	locationEndpoint := apiEndpoint + "location-area?offset=0&limit=20"

	return DataStore{
		APIEndpoint: apiEndpoint,
		APICache:    pokecache.New(cachePurgeInterval),
		NextAreaURL: locationEndpoint,
	}
}
