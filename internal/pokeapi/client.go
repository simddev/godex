package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/simddev/godex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type LocationAreaDetail struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func GetLocationAreaDetail(name string, cache *pokecache.Cache) (LocationAreaDetail, error) {
	url := baseURL + "/location-area/" + name
	if data, ok := cache.Get(url); ok {
		var result LocationAreaDetail
		if err := json.Unmarshal(data, &result); err != nil {
			return LocationAreaDetail{}, err
		}
		return result, nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return LocationAreaDetail{}, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaDetail{}, err
	}
	cache.Add(url, data)
	var result LocationAreaDetail
	if err := json.Unmarshal(data, &result); err != nil {
		return LocationAreaDetail{}, err
	}
	return result, nil
}

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string, cache *pokecache.Cache) (LocationAreaResponse, error) {
	if url == "" {
		url = baseURL + "/location-area/"
	}

	if data, ok := cache.Get(url); ok {
		var result LocationAreaResponse
		if err := json.Unmarshal(data, &result); err != nil {
			return LocationAreaResponse{}, err
		}
		return result, nil
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

	cache.Add(url, data)

	var result LocationAreaResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return LocationAreaResponse{}, err
	}
	return result, nil
}
