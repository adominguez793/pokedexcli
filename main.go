package main

import (
	"fmt"
	"log"

	"github.com/adominguez793/pokedexcli/internal/pokeapi"
)

func main() {
	// Repl()

	cli := pokeapi.NewClient()
	location, err := cli.PokeapiReq()
	if err != nil {
		log.Fatal("***** at the very end ...")
	}
	for _, place := range location.Results {
		fmt.Printf("%s\n", place.Name)
	}
}
