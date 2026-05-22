package main

import (
	"errors"
	"fmt"

	"github.com/simddev/godex/internal/pokeapi"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("usage: explore <location-area>")
	}
	area, err := pokeapi.GetLocationAreaDetail(args[0], cfg.cache)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, e := range area.PokemonEncounters {
		fmt.Println(" -", e.Pokemon.Name)
	}
	return nil
}
