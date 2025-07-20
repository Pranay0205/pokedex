package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, arg []string) error{
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL) 
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = &locationResp.Next
	cfg.prevLocationsURL = &locationResp.Previous

	for _, loc := range locationResp.Result{
		fmt.Println(loc.Name)
	}
	
	return nil
}

func commandMapb(cfg *config, arg []string) error {
	if cfg.prevLocationsURL == nil || *cfg.prevLocationsURL == "" {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = &locationResp.Next
	cfg.prevLocationsURL = &locationResp.Previous

	for _, loc := range locationResp.Result {
		fmt.Println(loc.Name)
	}

	return nil
}