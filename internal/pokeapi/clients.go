package pokeapi

import(
	"net/http"
	"io"
	"github.com/chaasfr/pokedexcli/internal/pokecache"
)

type Client struct {
	Http  http.Client
	Cache *pokecache.Cache
}

type LocationCLient struct {
	Client
	NextLocationsUrl     string
	PreviousLocationsUrl string
}

type PokemonClient struct {
	Client
}

func (client *Client) getBytesFromAPI(url string) ([]byte, error) {
	val, found := client.Cache.Get(url)
	if found {
		return val, nil
	}

	res, err := client.Http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resbyte, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	client.Cache.Add(url, resbyte)
	return resbyte, nil
}