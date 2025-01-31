package main

import (
	"time"

	"github.com/chaasfr/pokedexcli/internal/pokeapi"
	"github.com/chaasfr/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*configCommand, []string) error
}

type configCommand struct {
	cache                *pokecache.Cache
	locationClient       *pokeapi.LocationCLient
}

func initConf() configCommand {
	var conf configCommand
	conf.cache = pokecache.NewCache(5 * time.Second)
	conf.locationClient = pokeapi.NewClient(conf.cache)
	return conf
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Returns the next 20 locations. Use mapb to reverse.",
			callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Returns the previous 20 locations. Use map to reverse.",
			callback:    CommandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Use explore [location_name]. Returns the pokemons found in this location",
			callback:    CommandExplore,
		},
	}
}
