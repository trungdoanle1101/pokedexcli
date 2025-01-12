package main

type LocationArea struct {
	Name string `json:"name"`
}

type LocationAreaResponse struct {
	Previous *string `json:"previous"`
	Next *string `json:"next"`
	Results []LocationArea `json:"results"`
}
