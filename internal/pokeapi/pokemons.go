package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chaasfr/pokedexcli/internal/pokecache"
)

func NewPokemonClient(cache *pokecache.Cache) *PokemonClient {
	var client PokemonClient
	client.Http = *http.DefaultClient
	client.Cache = cache

	return &client
}

func (client *PokemonClient) GetPokemonBaseExperience(name string) (*Pokemon, error) {
	url := fmt.Sprintf("%s%s", PokemonAPIBaseUrl, name)

	resbyte, err := client.getBytesFromAPI(url)
	if err != nil {
		return nil, err
	}
	var pokemon Pokemon

	if err := json.Unmarshal(resbyte, &pokemon); err != nil {
		return nil, fmt.Errorf("Pokemon not found %s", name)
	}
	return &pokemon, nil
}
