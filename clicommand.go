package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
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
		"catch": {
			name:        "catch",
			description: "Use catch [pokemon_name]. Tries to capture the pokemon",
			callback:    CommandCatch,
		},
	}
}
