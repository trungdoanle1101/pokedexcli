package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}

	defer resp.Body.Close()

	locationsResp := RespShallowLocations{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}
	return locationsResp, nil
}
