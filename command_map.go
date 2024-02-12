package main

import (
	"fmt"
	"log"

	"github.com/adominguez793/pokedexcli/internal/pokeapi"
)

func commandMap() error {
	cli := pokeapi.NewClient()
	location, err := cli.PokeapiReq()
	if err != nil {
		log.Fatal("Error trying to generate next 30 Pokemon Area Locations ...")
		return err
	}
	for _, place := range location.Results {
		fmt.Printf("%s\n", place.Name)
	}
	return nil
}
