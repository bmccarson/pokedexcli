package state

type DataStore struct {
	APIEndpoint         string
	NextLocationURL     string
	PreviousLocationURL string
}

func Init(apiEndpoint string) DataStore {
	locationEndpoint := apiEndpoint + "location-area?offset=0&limit=20"

	return DataStore{
		APIEndpoint:     apiEndpoint,
		NextLocationURL: locationEndpoint,
	}
}
