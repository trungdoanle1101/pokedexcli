package pokeapi

import "github.com/trungdoanle1101/pokedexcli/internal/pokedex"

func (c *Client) GetFromPokedex(pokemonName string) (pokedex.Pokemon, bool) {
	return c.pokedex.Get(pokemonName)
}
