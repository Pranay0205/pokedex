package main

import (
	"time"

	"github.com/Pranay0205/pokedexcli/internal/pokeapi"
)

func main() {

	pokeClient := pokeapi.NewClient(5 * time.Second, 5* time.Minute)


	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.PokemonResp{},
		maxBaseExperience: 1000,
	}
	startRepl(cfg)
}