package main

type cliCommand struct {
	name        string
	description string
	callback    func(*configCommand) error
}

type configCommand struct {
	next     string
	previous string
}

func initConf() configCommand {
	var conf configCommand
	conf.next = "https://pokeapi.co/api/v2/location-area/"
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
			name:         "map",
			description: "Returns the next 20 locations. Use mapb to reverse.",
			callback: CommandMap,
		},
		"mapb": {
			name:         "map",
			description: "Returns the previous 20 locations. Use map to reverse.",
			callback: CommandMapb,
		},
	}
}
