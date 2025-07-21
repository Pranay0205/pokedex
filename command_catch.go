package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strings"
)

const (maxBaseExperience = 1000)

func commandCatch(cfg *config, args []string) error {

	pokemonName := strings.ToLower(args[0])

	if len(args) == 0 {
    return errors.New("no pokemon name provided")
	}
	
	pokemonResp, err := cfg.pokeapiClient.PokemonList(pokemonName)
	
	if err != nil {
		return err
	}
	
	fmt.Printf("Throwing a Pokeball at %s%s%s...\n", ColorBlue, pokemonName, ColorReset)
	
	randomNum := rand.IntN(maxBaseExperience)

	if randomNum <= int(pokemonResp.BaseExperience) {
		
		fmt.Printf("%s%s escaped!%s\n", ColorRed, pokemonName, ColorReset)
		return nil
	}

	cfg.caughtPokemon[pokemonName] = pokemonResp

	fmt.Printf("%s%s was caught!%s\n", ColorGreen, pokemonName, ColorReset)
	
	return nil
}

