package main

import (
	"fmt"
)

func commandMap(cfg *config, options ...string) error {
	if cfg == nil {
		return fmt.Errorf("config cannot be nil")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	return nil
}

func commandMapb(cfg *config, options ...string) error {
	if cfg == nil {
		return fmt.Errorf("config cannot be nil")
	}

	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("you're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	return nil
}
