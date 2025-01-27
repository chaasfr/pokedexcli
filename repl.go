package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	trimLowerText := strings.Fields(lowerText)
	return trimLowerText
}

func replStart(){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		textClean := cleanInput(text)

		if len(textClean) > 0 {
			toPrint := fmt.Sprintf("Your command was: %s", textClean[0])
			fmt.Println(toPrint)
		} else {
			fmt.Println("Please provide a command")
		}
	}
}