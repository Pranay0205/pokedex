package main

import (
	"errors"
	"fmt"
)



func commandExplore(cfg *config, arg []string) error{
		if len(arg) == 0 {
			return errors.New("no location provided")
		}

		name := &arg[0]

		exploreResp, err := cfg.pokeapiClient.ExploreLocation(name)

		fmt.Printf("Exploring %s...\n", exploreResp.Name)

		if err != nil {
			return err
		}
		
		fmt.Printf("Found Pokemon:\n")
		for i, item := range exploreResp.PokemonEncounters {
			fmt.Printf("%d) %s\n", i + 1, item.Pokemon.Name)
		}

		return nil
}