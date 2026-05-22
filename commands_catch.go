package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/simddev/godex/internal/pokeapi"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("usage: catch <pokemon>")
	}
	name := args[0]
	pokemon, err := pokeapi.GetPokemon(name, cfg.cache)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	// Higher base experience = harder to catch. Threshold scales so most
	// common Pokemon (~50 XP) are easy and legendaries (~300 XP) are hard.
	if rand.Intn(pokemon.BaseExperience) < 40 {
		fmt.Printf("%s was caught!\n", name)
		cfg.pokedex[name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}
