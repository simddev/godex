package main

import (
	"fmt"

	"github.com/simddev/godex/internal/pokeapi"
)

func commandMap(cfg *config, args []string) error {
	url := ""
	if cfg.nextURL != nil {
		url = *cfg.nextURL
	}
	resp, err := pokeapi.GetLocationAreas(url, cfg.cache)
	if err != nil {
		return err
	}
	cfg.nextURL = resp.Next
	cfg.prevURL = resp.Previous
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb(cfg *config, args []string) error {
	if cfg.prevURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	resp, err := pokeapi.GetLocationAreas(*cfg.prevURL, cfg.cache)
	if err != nil {
		return err
	}
	cfg.nextURL = resp.Next
	cfg.prevURL = resp.Previous
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
	return nil
}
