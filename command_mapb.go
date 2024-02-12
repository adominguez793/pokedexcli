package main

import (
	"errors"
	"fmt"

	"github.com/adominguez793/pokedexcli/internal/pokeapi"
)

func commandMapB(cfg *Config) error {
	if cfg.Previous == nil {
		fmt.Println("This is the first page of Pokemon Area Locations ...")
		return errors.New("First page!")
	}
	fullURL := cfg.Previous

	cli := pokeapi.NewClient()
	location, err := cli.PokeapiReq(*fullURL)
	if err != nil {
		return err
	}
	for _, place := range location.Results {
		fmt.Printf("%s\n", place.Name)
	}
	cfg.Previous = location.Previous
	cfg.Next = location.Next

	return nil
}
