package main

import "fmt"

func CommandExplore(conf *configCommand, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("arg missing. Please seek help on explore")
	}
	if len(args) > 1 {
		return fmt.Errorf("too many args. PLease seek help on explore")
	}
	locationName := args[0]
	fmt.Println("Exploring " + locationName + "...")
	locationAreaResult, err := conf.locationClient.GetLocationDetails(locationName)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")

	for _, pokemonEncounter := range locationAreaResult.PokemonEncounters {
		fmt.Println(" - " + pokemonEncounter.Pokemon.Name)
	}
	
	return nil
}