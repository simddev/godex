package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {
	fmt.Println("Your Pokedex:")
	for name := range cfg.pokedex {
		fmt.Println(" -", name)
	}
	return nil
}
