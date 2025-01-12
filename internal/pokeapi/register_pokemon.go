package pokeapi

import "github.com/trungdoanle1101/pokedexcli/internal/pokedex"

func (c *Client) RegisterPokemon(resp RespPokemon) pokedex.Pokemon {
	stats := make([]struct {
		Stat    string
		BaseVal int
	}, 0)

	types := make([]string, 0)
	for _, stat := range resp.Stats {
		stats = append(stats, struct {
			Stat    string
			BaseVal int
		}{Stat: stat.Stat.Name, BaseVal: stat.BaseStat})
	}

	for _, t := range resp.Types {
		types = append(types, t.Type.Name)
	}

	pokemon := pokedex.NewPokemon(resp.Name, resp.BaseExperience, resp.Height, resp.Weight, stats, types)

	c.pokedex.Add(pokemon)

	return pokemon
}
