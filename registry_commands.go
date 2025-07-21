package main

type cliCommand struct {
	name string
	description string
	callback func(c *config, arg []string) error
}

var registred_commands map[string]cliCommand

func init() {
		registred_commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays all available commands (Usage: help)",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex application (Usage: exit)",
			callback:    commandExit,
		},
		"map":{
			name: "map",
			description: "Displays next 20 location areas (Usage: map)",
			callback: commandMap,
		},
		"mapb":{
			name: "mapb",
			description: "Displays previous 20 location areas (Usage: mapb)",
			callback: commandMapb,
		},
		"explore":{
			name: "explore",
			description: "Explores a location area for Pokemon (Usage: explore <location-name>)",
			callback: commandExplore,
		},
		"catch": {
				name: "catch",
				description: "Catches a Pokemon (Usage: catch <pokemon-name>)", 
				callback: commandCatch,
		},
	}
}