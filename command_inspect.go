package main

import (
	"fmt"

	"github.com/adominguez793/pokedexcli/internal/pokecache"
)

func commandInspect(cfg *Config, cache *pokecache.Cache, arg string) error {
	if arg != Pokedex[arg].Name {
		fmt.Printf("%s is not in your Pokedex\n", arg)
		return nil
	}

	fmt.Printf("Name: %s\n", Pokedex[arg].Name)
	fmt.Printf("Height: %d\n", Pokedex[arg].Height)
	fmt.Printf("Weight %d\n", Pokedex[arg].Weight)

	fmt.Println("Stats:")
	for _, pokeStats := range Pokedex[arg].Stats {
		fmt.Printf("  - %s: %d\n", pokeStats.Stat.Name, pokeStats.BaseStat)
	}

	if len(Pokedex[arg].Types) == 1 {
		fmt.Println("Type:")
	}
	if len(Pokedex[arg].Types) > 1 {
		fmt.Println("Types:")
	}
	for _, pokeType := range Pokedex[arg].Types {
		fmt.Printf("  - %s\n", pokeType.Type.Name)
	}

	return nil
}
