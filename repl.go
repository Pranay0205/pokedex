package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Pranay0205/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}


func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)	
	for {
		fmt.Print("Pokedex > ")
		var prompt string

		if scanner.Scan() {
			prompt = scanner.Text()
		}

		prompt = strings.ToLower(strings.TrimSpace(prompt))
		if len(prompt) == 0 {
			continue
		}

		command := strings.Fields(prompt)[0]
		current_command, exists := registred_commands[command]
		if !exists {
			fmt.Printf("Unknown Command")
			continue
		}
		
		err := current_command.callback(cfg)

		if err != nil {
			fmt.Println(err)
		}
		continue
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(strings.TrimSpace(text))
	fields := strings.Fields(text) 
	return fields
}
