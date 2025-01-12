package main

import "fmt"

func commandInspect(cfg *config, options ...string) error {
	if cfg == nil {
		return fmt.Errorf("config cannot be nil")
	}

	if len(options) != 1 {
		return fmt.Errorf("please provide a valid pokemon name")
	}

	pokemonName := options[0]
	pokemon, ok := cfg.pokeapiClient.GetFromPokedex(pokemonName)
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, st := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", st.Stat, st.BaseVal)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Println(" -", t)
	}
	return nil
}
