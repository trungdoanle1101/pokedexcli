package main

import "fmt"

func commandExplore(cfg *config, options ...string) error {
	if cfg == nil {
		return fmt.Errorf("config cannot be nil")
	}

	if len(options) != 1 {
		return fmt.Errorf("please provide one area name")
	}

	area := options[0]
	ladResp, err := cfg.pokeapiClient.GetLocationArea(area)
	if err != nil {
		return err
	}

	fmt.Println("Exploring", area)
	if len(ladResp.PokemonEncounters) > 0 {
		fmt.Println("Found Pokemon:")
	}

	for _, encounter := range ladResp.PokemonEncounters {
		fmt.Println(" - ", encounter.Pokemon.Name)
	}

	return nil

}
