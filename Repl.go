package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/adominguez793/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, *pokecache.Cache, string) error
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 30 Pokemon Area Locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays last 30 Pokemon Area Locations",
			callback:    commandMapB,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Explore the Pokedex",
			callback:    commandExplore,
		},
	}
}

func Repl(cfg *Config, cache *pokecache.Cache) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			continue
		}

		cleanInput := cleanInput(input)
		command := cleanInput[0]
		availableCommands := getCommand()
		if command != availableCommands[command].name {
			fmt.Println("That's not a real command...")
			continue
		}

		if len(cleanInput) > 1 {
			arg := cleanInput[1]
			availableCommands[command].callback(cfg, cache, arg)
		}
		if len(cleanInput) == 1 {
			arg := ""
			availableCommands[command].callback(cfg, cache, arg)
		}
	}
}

func cleanInput(input string) []string {
	cleanInput := strings.ToLower(input)
	cleanInputSlice := strings.Fields(cleanInput)
	return cleanInputSlice
}
