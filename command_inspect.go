package main

import(
	"fmt"
)

func CommandInspect(conf *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("too few argument")
	}

	if len(args) > 1 {
		return fmt.Errorf("too many argument")
	}

	pokemonName := args[0]

	pokemon, ok := conf.pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	
	fmt.Println(pokemon.GetDetails())

	return nil
}