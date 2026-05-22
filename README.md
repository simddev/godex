# Godex

A Pokedex CLI built in Go. Explore the Pokemon world, catch Pokemon, and build your collection.

## Usage

```
go build
./godex
```

## Commands

| Command | Description |
|---------|-------------|
| `help` | Display available commands |
| `map` | Show the next 20 location areas |
| `mapb` | Show the previous 20 location areas |
| `explore <area>` | List Pokemon found in a location area |
| `catch <pokemon>` | Attempt to catch a Pokemon |
| `inspect <pokemon>` | View stats for a caught Pokemon |
| `pokedex` | List all caught Pokemon |
| `exit` | Exit the program |

## Example Session

```
Pokedex > map
canalave-city-area
eterna-city-area
...
Pokedex > explore canalave-city-area
Found Pokemon:
 - tentacool
 - magikarp
 ...
Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
magikarp was caught!
Pokedex > inspect magikarp
Name: magikarp
Height: 9
Weight: 100
Stats:
  -hp: 20
  -attack: 10
  ...
Pokedex > pokedex
Your Pokedex:
 - magikarp
```

## Data

Pokemon data is sourced from [PokeAPI](https://pokeapi.co/). Responses are cached in memory to avoid redundant network requests.
