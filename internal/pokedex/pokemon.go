package pokedex

import "fmt"

type Pokedex struct {
	dex map[string]Pokemon
}

type Pokemon struct {
	Name           string
	BaseExperience int
	Height         int
	Weight         int
	Stats          []struct {
		Stat    string
		BaseVal int
	}
	Types []string
}

func NewPokedex() Pokedex {
	return Pokedex{dex: make(map[string]Pokemon)}
}

func NewPokemon(name string, baseExperience, height, weight int, stats []struct {
	Stat    string
	BaseVal int
}, types []string) Pokemon {
	pokemon := Pokemon{
		Name:           name,
		BaseExperience: baseExperience,
		Height:         height,
		Weight:         weight,
		Stats:          stats,
		Types:          types,
	}

	return pokemon
}

func (pd *Pokedex) Add(pokemon Pokemon) {
	pd.dex[pokemon.Name] = pokemon
}

func (pd *Pokedex) Exists(name string) bool {
	_, ok := pd.dex[name]
	return ok
}

func (pd *Pokedex) Get(name string) (Pokemon, bool) {
	pokemon, ok := pd.dex[name]
	if !ok {
		return Pokemon{}, false
	}
	return pokemon, true
}

func (pd *Pokedex) Print() {
	for name := range pd.dex {
		fmt.Println(" - ", name)
	}
}
