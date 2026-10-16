package main

import (
	"time"

	"github.com/andrius-bilotas/pokedexcli/internal"
)

type config struct {
	commands map[string]cliCommand
	Next     string
	Previous string
	cache    *internal.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(params []string, cfg *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Provides a list of locations. Subsequent calls move to next page",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Returns previous page of locations if available",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Takes in location area name and provides pokemon available in that area",
			callback:    commandExplore,
		},
	}
}

func main() {
	config := config{
		commands: getCommands(),
		cache:    internal.NewCache(20 * time.Second),
	}
	repl(&config)
}
