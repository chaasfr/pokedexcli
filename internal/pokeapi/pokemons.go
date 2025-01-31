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


func (p Pokemon) GetDetails() string {
	result := ""
	
	result += fmt.Sprintf("Name: %s\n",p.Name)
	result += fmt.Sprintf("Height: %v\n",p.Height)
	result += fmt.Sprintf("Weight: %v\n",p.Weight)
	result += "Stats:\n"
	
	for _, stat := range p.Stats {
		result += fmt.Sprintf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	result += "Types:\n"
	for _, t := range p.Types {
		result += fmt.Sprintf("  - %s", t.Type.Name)
	}

	return result
}