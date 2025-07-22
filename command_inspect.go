package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandInspect(c *config, args []string) error {
	
	
	if len(args) < 1 {
    return errors.New("no pokemon name provided")
	}
	

	pokemonName := strings.ToLower(args[0])
	Pokemon, exists := c.caughtPokemon[pokemonName]

	if !exists {
		fmt.Printf("%s pokemon wasn't caught", pokemonName)
		return nil
	}

	fmt.Printf("Name: %s\n", Pokemon.Name)
	fmt.Printf("Height: %d\n", Pokemon.Height)
	fmt.Printf("Weight: %d\n", Pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range Pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, pokemonType := range Pokemon.Types {
			fmt.Printf("  - %s\n", pokemonType.Type.Name)
}
	return nil
}