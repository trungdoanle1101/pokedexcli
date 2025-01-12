package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, options ...string) error {
	if cfg == nil {
		return fmt.Errorf("config cannot be nil")
	}

	if len(options) != 1 {
		return fmt.Errorf("please provide a valid pokemon name")
	}

	pokemonName := options[0]

	respPokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	// threshold for catching pokemon
	// if larger than threshold then pokemon is caught
	// otherwise escape
	// TODO: The larger the base experience, the larger the threshold
	threshold := 0.5
	randomNum := rand.Float64()

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if randomNum <= threshold {
		fmt.Println(pokemonName, "escaped!")
		return nil
	}

	fmt.Println(pokemonName, "was caught!")
	cfg.pokeapiClient.RegisterPokemon(respPokemon)
	
	return nil

}
