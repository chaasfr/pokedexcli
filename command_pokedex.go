package main

import "fmt"


func CommandPokedex(conf *config, args []string) error {
	if len(conf.pokedex) == 0 {
		fmt.Println("Your Pokedex is empty")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name, _ := range conf.pokedex {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}