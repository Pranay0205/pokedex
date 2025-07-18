package main

type cliCommand struct {
	name string
	description string
	callback func() error
}

var registred_commands map[string]cliCommand

func init() {
		registred_commands = map[string]cliCommand{
        "help": {
            name:        "help",
            description: "Get help with Pokedex",
            callback:    commandHelp,
        },
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
    }
}