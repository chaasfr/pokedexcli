package main

import (
	"fmt"
	"github.com/chaasfr/pokedexcli/internal/pokeapi"
)

func CommandMap(conf *configCommand, args []string) error {
	if conf.locationClient.NextLocationsUrl == "" {
		fmt.Println("you're on the last page")
		return nil
	}
	return getLocations(conf.locationClient.GetNextLocations)
}

func CommandMapb(conf *configCommand, args []string) error {
	if conf.locationClient.PreviousLocationsUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	return getLocations(conf.locationClient.GetPreviousLocations)
}

func getLocations( getLocations func()(*pokeapi.LocationBulkResult, error)) error {
	locationBulkResults, err := getLocations()
	if err != nil {
		return err
	}
	for _, location := range locationBulkResults.Results {
		fmt.Println(location.Name)
	}
	return nil
}
