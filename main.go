package main

import (
	"time"

	"github.com/adominguez793/pokedexcli/internal/pokeapi"
	"github.com/adominguez793/pokedexcli/internal/pokecache"
)

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

	fiveMin := time.Minute * 5
	cachePtr := pokecache.NewCache(fiveMin)

	Repl(&cfg, cachePtr)
}
