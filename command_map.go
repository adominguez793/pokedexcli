package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/adominguez793/pokedexcli/internal/pokeapi"
	"github.com/adominguez793/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

func commandMap(cfg *Config, cache *pokecache.Cache, arg string) error {
	fullURL := baseURL + "/location-area"
	if cfg.Next != nil {
		fullURL = *cfg.Next
	}

	var Local pokeapi.LocationArea
	cacheVal, confirmation := cache.Get(fullURL)
	if confirmation == true {
		err := json.Unmarshal(cacheVal, &Local)
		if err != nil {
			return errors.New("Had trouble with Unmarshal after retrieving cache")
		}
		for _, locale := range Local.Results {
			fmt.Printf("%s\n", locale.Name)
		}
		cfg.Next = Local.Next
		cfg.Previous = Local.Previous
		return nil
	}

	cli := pokeapi.NewClient()
	location, err := cli.PokeapiReq(fullURL, cache)
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
