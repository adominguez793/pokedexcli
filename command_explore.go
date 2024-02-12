package main

import (
	"encoding/json"
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

	fmt.Printf("Exploring %s...\n", arg)

	cacheVal, confirmation := cache.Get(arg)
	if confirmation == true {
		var NameLocation pokeapi.NameLocationArea
		err := json.Unmarshal(cacheVal, &NameLocation)
		if err != nil {
			fmt.Println("Error unmarshalling var dat in explore command")
			return errors.New("Error unmarshalling var dat in explore command")
		}
		fmt.Println("Found Pokemon:")
		for _, pokemon := range NameLocation.PokemonEncounters {
			fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
		}
		return nil
	}

	cli := pokeapi.NewClient()
	location, err := cli.NamePokeapiReq(arg, cache)
	if err != nil {
		fmt.Println(err)
		return errors.New("Error with pokeapi request in explore command")
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
