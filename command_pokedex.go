package main

import "fmt"

func commandPokedex(cfg *config, options ...string) error {
	if cfg == nil {
		return fmt.Errorf("config cannot be nil")
	}

	cfg.pokeapiClient.PrintPokedex()
	return nil
}
