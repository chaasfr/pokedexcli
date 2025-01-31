package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	trimLowerText := strings.Fields(lowerText)
	return trimLowerText
}

func handleInput(input []string, conf *configCommand) error {
	if len(input) == 0 {
		return fmt.Errorf("no input provided")
	}
	action := input[0]

	if command, ok := getCommands()[action]; ok {
		return command.callback(conf, input[1:])
	}

	return fmt.Errorf("unknown action %s", action)
}

func replStart() {
	scanner := bufio.NewScanner(os.Stdin)
	conf := initConf()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		if err := handleInput(cleanInput((text)), &conf); err != nil {
			message := fmt.Sprintf("incorrect action: %s. Please use help to see the list of valid actions.", err)
			fmt.Println(message)
		}
	}
}
