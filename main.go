package main

import "github.com/adominguez793/pokedexcli/internal/pokeapi"

type Config struct {
	httpClient pokeapi.Client
	Next       *string
	Previous   *string
}

func main() {
	cfg := Config{
		httpClient: pokeapi.NewClient(),
		Next:       nil,
		Previous:   nil,
	}

	Repl(&cfg)
}
