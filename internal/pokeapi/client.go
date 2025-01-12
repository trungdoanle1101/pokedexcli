package pokeapi

import (
	"net/http"
	"time"

	"github.com/trungdoanle1101/pokedexcli/internal/pokecache"
	"github.com/trungdoanle1101/pokedexcli/internal/pokedex"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
	pokedex    pokedex.Pokedex
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache:   pokecache.NewCache(5 * time.Second),
		pokedex: pokedex.NewPokedex(),
	}
}
