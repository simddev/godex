package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/simddev/godex/internal/pokecache"
)

type config struct {
	nextURL *string
	prevURL *string
	cache   *pokecache.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 location areas",
			callback:    commandMapb,
		},
	}
}

func commandHelp(cfg *config) error {
	cmds := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range cmds {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func main() {
	cmds := getCommands()
	cfg := &config{cache: pokecache.NewCache(5 * time.Minute)}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		cmd, ok := cmds[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(cfg); err != nil {
			fmt.Println(err)
		}
	}
}
