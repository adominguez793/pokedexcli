package main

import (
	"errors"
	"fmt"

	"github.com/adominguez793/pokedexcli/internal/pokeapi"
	"github.com/adominguez793/pokedexcli/internal/pokecache"
)

func commandExplore(cfg *Config, cache *pokecache.Cache, arg string) error {
	if arg == "" {
		fmt.Println("Missing input for explore command ...")
		return errors.New("Missing input for explore command ... ")
	}

	cli := pokeapi.NewClient()
	location, err := cli.NamePokeapiReq(arg, cache)
	if err != nil {
		fmt.Println(err)
		return errors.New("Error with pokeapi request in explore command")
	}
	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf("%s\n", pokemon.Pokemon.Name)
	}
	return nil
}
