package main

import (
	"os"

	"github.com/adominguez793/pokedexcli/internal/pokecache"
)

func commandExit(cfg *Config, cache *pokecache.Cache, arg string) error {
	os.Exit(0)
	return nil
}
