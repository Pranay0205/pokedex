package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)	
	for {
		fmt.Print("Pokedex > ")
		var prompt string

		if scanner.Scan() {
			prompt = scanner.Text()
		}

		prompt = strings.ToLower(strings.TrimSpace(prompt))
		if len(prompt) == 0 {
			fmt.Printf("No prompt entered")
		}

		command := strings.Fields(prompt)[0]
		_, exists := registred_commands[command]
		if !exists {
			fmt.Printf("Unknown Command")
		}

		registred_commands[command].callback(&config{})
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(strings.TrimSpace(text))
	fields := strings.Fields(text) 
	return fields
}
