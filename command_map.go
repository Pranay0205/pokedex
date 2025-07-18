package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type config struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Result   []location `json:"results"`
}



func commandMap(c *config) error{

	var res *http.Response
	var err error

	if c.Next == "" {
		res, err = http.Get("https://pokeapi.co/api/v2/location/")
	} else {
		res, err = http.Get(c.Next)
	}

	if err != nil {
		return fmt.Errorf("failed to get response from the api: %w", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(body, c)

	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil
}

func commandMapb(c *config) error {
	res, err := http.Get(c.Previous)

	if err != nil {
		return fmt.Errorf("failed to get response from the api: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(body, c)

	if err != nil {
		return fmt.Errorf("failed to unmarshal reponse: %w", err)
	}

	return nil
}