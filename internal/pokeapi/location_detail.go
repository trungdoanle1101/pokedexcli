package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(area string) (RespLocationArea, error) {
	// baseURL = "https://pokeapi.co/api/v2"
	url := baseURL + "/location-area/" + area

	if data, ok := c.cache.Get(url); ok {
		ladResp := RespLocationArea{}
		err := json.Unmarshal(data, &ladResp)
		if err != nil {
			return RespLocationArea{}, err
		}
		return ladResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	ladResp := RespLocationArea{}
	err = json.Unmarshal(data, &ladResp)
	if err != nil {
		return RespLocationArea{}, err
	}

	c.cache.Add(url, data)

	return ladResp, nil

}
