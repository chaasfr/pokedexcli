package main

import "fmt"

func commandHelp(conf *configCommand, args []string) error {
	helpMessage := `Welcome to the Pokedex!
Usage:

`
	for _, command := range getCommands() {
		helpMessage += fmt.Sprintf("%s: %s\n", command.name, command.description)
	}
	fmt.Println(helpMessage)
	return nil
}
