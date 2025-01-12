package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const locationAreaUrl = "https://pokeapi.co/api/v2/location-area"

func commandMap(config *CliConfig) error {
	var url string
	if config == nil || config.Next == nil || *config.Next == "" {
		url = locationAreaUrl
	} else {
		url = *config.Next
	}

	res, err := http.Get(url)

	if err != nil {
		return fmt.Errorf("error fetching location-area data: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("fetching location-area data failed with status code: %v", res.Status)
	}

	var lar LocationAreaResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&lar)
	if err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	for _, la := range lar.Results {
		fmt.Println(la.Name)
	}

	config.Previous = lar.Previous
	config.Next = lar.Next

	return nil
}

func commandMapb(config *CliConfig) error {

	if config == nil || config.Previous == nil || *config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	url := *config.Previous
	res, err := http.Get(url)

	if err != nil {
		return fmt.Errorf("error fetching location-area data: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("fetching location-area data failed with status code: %v", res.Status)
	}

	var lar LocationAreaResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&lar)
	if err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	for _, la := range lar.Results {
		fmt.Println(la.Name)
	}

	config.Previous = lar.Previous
	config.Next = lar.Next

	return nil
}
