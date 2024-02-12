package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func Repl() {
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
		availableCommands[command].callback()
	}
}

func cleanInput(input string) []string {
	cleanInput := strings.ToLower(input)
	cleanInputSlice := strings.Fields(cleanInput)
	return cleanInputSlice
}
