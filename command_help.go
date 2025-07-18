package main

import (
	"fmt"
)



func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")

	if len(registred_commands) == 0 {
		return fmt.Errorf("no commands registry")
	}

	for _, value := range registred_commands {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}

	return nil
}