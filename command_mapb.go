package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/adominguez793/pokedexcli/internal/pokeapi"
	"github.com/adominguez793/pokedexcli/internal/pokecache"
)

func commandMapB(cfg *Config, cache *pokecache.Cache) error {
	if cfg.Previous == nil {
		fmt.Println("This is the first page of Pokemon Area Locations ...")
		return errors.New("First page!")
	}
	fullURL := cfg.Previous

	var Local pokeapi.LocationArea
	cacheVal, confirmation := cache.Get(*fullURL)
	if confirmation == true {
		err := json.Unmarshal(cacheVal, &Local)
		if err != nil {
			return errors.New("Had trouble retrieving cache for last 30 Pokemon Area Locations")
		}
		for _, locale := range Local.Results {
			fmt.Printf("%s\n", locale.Name)
		}
		cfg.Next = Local.Next
		cfg.Previous = Local.Previous
		return nil
	}

	cli := pokeapi.NewClient()
	location, err := cli.PokeapiReq(*fullURL, cache)
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
