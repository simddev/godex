package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2"

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (LocationAreaResponse, error) {
	if url == "" {
		url = baseURL + "/location-area/"
	}
	resp, err := http.Get(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	var result LocationAreaResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return LocationAreaResponse{}, err
	}
	return result, nil
}
