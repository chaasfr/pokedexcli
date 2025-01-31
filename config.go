package main

import (
	"time"
	"github.com/chaasfr/pokedexcli/internal/pokeapi"
	"github.com/chaasfr/pokedexcli/internal/pokecache"
)


type config struct {
	cache          *pokecache.Cache
	pokedex        map[string]*pokeapi.Pokemon
	locationClient *pokeapi.LocationCLient
	pokemonClient  *pokeapi.PokemonClient
}

func initConf() config {
	var conf config
	conf.cache = pokecache.NewCache(5 * time.Second)
	conf.locationClient = pokeapi.NewLocationClient(conf.cache)
	conf.pokemonClient = pokeapi.NewPokemonClient(conf.cache)
	conf.pokedex = make(map[string]*pokeapi.Pokemon)
	return conf
}