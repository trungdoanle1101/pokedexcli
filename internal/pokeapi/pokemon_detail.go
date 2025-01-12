package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (RespPokemon, error) {
	if name == "" {
		return RespPokemon{}, fmt.Errorf("pokemon name must not be empty")
	}

	url := baseURL + "/pokemon/" + name

	if data, ok := c.cache.Get(url); ok {
		respPokemon := RespPokemon{}
		err := json.Unmarshal(data, &respPokemon)
		if err != nil {
			return RespPokemon{}, err
		}
		return respPokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	respPokemon := RespPokemon{}
	err = json.Unmarshal(data, &respPokemon)
	if err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(url, data)
	return respPokemon, nil

}
