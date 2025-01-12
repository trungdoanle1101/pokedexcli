package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if data, ok := c.cache.Get(url); ok {
		locationResp := RespShallowLocations{}
		err := json.Unmarshal(data, &locationResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationResp, nil
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
	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return RespShallowLocations{}, err
	}

	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, data)

	return locationsResp, nil
}
