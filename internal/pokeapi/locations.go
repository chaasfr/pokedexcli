package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chaasfr/pokedexcli/internal/pokecache"
)

func NewLocationClient(cache *pokecache.Cache) *LocationCLient {
	var client LocationCLient
	client.Http = *http.DefaultClient
	client.Cache = cache
	client.NextLocationsUrl = LocationAPIBaseUrl

	return &client
}

func (client *LocationCLient) GetNextLocations() (*LocationBulkResult, error) {
	return client.getLocations(client.NextLocationsUrl)
}

func (client *LocationCLient) GetPreviousLocations() (*LocationBulkResult, error) {
	return client.getLocations(client.PreviousLocationsUrl)
}

func (client *LocationCLient) getLocations(url string) (*LocationBulkResult, error) {
	resbyte, err := client.getBytesFromAPI(url)
	if err != nil {
		return nil, err
	}

	var locationResult LocationBulkResult
	err = json.Unmarshal(resbyte, &locationResult)
	if err != nil {
		return nil, err
	}
	client.NextLocationsUrl = locationResult.Next
	client.PreviousLocationsUrl = locationResult.Previous

	return &locationResult, nil
}

func (client *LocationCLient) GetLocationDetails(locationName string) (*LocationAreaResult, error) {
	url := fmt.Sprintf("%s%s", LocationAPIBaseUrl, locationName)
	resbyte, err := client.getBytesFromAPI(url)
	if err != nil {
		return nil, err
	}
	var locationAreaResult LocationAreaResult
	if err = json.Unmarshal(resbyte, &locationAreaResult); err != nil {
		return nil, fmt.Errorf("location not found %s", locationName)
	}
	return &locationAreaResult, nil
}