package main

import (
	"fmt"

	"github.com/adominguez793/pokedexcli/internal/pokecache"
)

func commandPokedex(cfg *Config, cache *pokecache.Cache, arg string) error {
	fmt.Println("Your Pokedex:")
	for _, poke := range Pokedex {
		fmt.Printf("  - %s\n", poke.Name)
	}

	return nil
}
