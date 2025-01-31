package main

import (
	"fmt"
	"math"
	"math/rand"
)

func CommandCatch(conf *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("too few argument")
	}

	if len(args) > 1 {
		return fmt.Errorf("too many argument")
	}

	name := args[0]

	pokemon, err := conf.pokemonClient.GetPokemonBaseExperience(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	fmt.Printf("pokemon: %s - baseExp: %v\n", name, pokemon.BaseExperience)

	if IsCaught(pokemon.BaseExperience) {
		fmt.Printf("%s was caught!\n", name)
		conf.pokedex[name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}

func IsCaught(baseExp int64) bool {
	captureThreshold := 40.0 + math.Sqrt(float64(baseExp))
	roll := rand.Float64() * 100
	return roll >= captureThreshold
}
