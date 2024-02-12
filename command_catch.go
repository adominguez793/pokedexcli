package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/adominguez793/pokedexcli/internal/pokeapi"
	"github.com/adominguez793/pokedexcli/internal/pokecache"
)

var Pokedex = make(map[string]pokeapi.PokeInfo)

func commandCatch(cfg *Config, cache *pokecache.Cache, arg string) error {
	if arg == "" {
		fmt.Println("Missing input for explore command ...")
		return errors.New("Missing input for explore command ... ")
	}
	client := pokeapi.NewClient()
	pokeInfo, err := client.CatchPokeapiReq(arg)
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}

	catchOdds := probabilityByLevel(pokeInfo.BaseExperience)
	fmt.Printf("\nThrowing a Pokeball at %s...\n", pokeInfo.Name)

	if rand.Float64() <= catchOdds {
		fmt.Printf("%s was caught!\n", pokeInfo.Name)
		_, ok := Pokedex[pokeInfo.Name]
		if !ok {
			Pokedex[pokeInfo.Name] = pokeInfo
		}
		// Once the Pokemon is caught, add it to the user's Pokedex.
		// I used a map[string]Pokemon to keep track of caught Pokemon.
		// You'll want to store the Pokemon's data so that in the next step we can use it.
	} else {
		fmt.Printf("%s escaped!\n", pokeInfo.Name)
	}

	return nil
}

func probabilityByLevel(level int) float64 {
	return 1 / (float64(level) / 70.0)
}

// type Pokedex struct {
// 	caughtPokemon map[string]pokemonEntry
// }
