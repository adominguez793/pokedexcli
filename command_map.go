package main

import (
	"fmt"
	"log"

	"github.com/adominguez793/pokedexcli/internal/pokeapi"
)

const baseURL = "https://pokeapi.co/api/v2"

func commandMap(cfg *Config) error {
	fullURL := baseURL + "/location-area"
	if cfg.Next != nil {
		fullURL = *cfg.Next
	}

	cli := pokeapi.NewClient()
	location, err := cli.PokeapiReq(fullURL)
	if err != nil {
		log.Fatal("Error trying to generate next 30 Pokemon Area Locations ...")
		return err
	}
	for _, place := range location.Results {
		fmt.Printf("%s\n", place.Name)
	}
	cfg.Next = location.Next
	cfg.Previous = location.Previous

	return nil
}
