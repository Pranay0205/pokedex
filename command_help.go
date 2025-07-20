package main

import (
	"fmt"
)



func commandHelp(c *config, arg []string) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Commands:\n")

	if len(registred_commands) == 0 {
		return fmt.Errorf("no commands registry")
	}

	i := 1
	for _, value := range registred_commands {
			fmt.Printf("%d) %s: %s\n", i, value.name, value.description)
			i++
	}

	return nil
}