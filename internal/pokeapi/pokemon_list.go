package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)


func (c *Client) PokemonList(pokemonName string) (PokemonResp, error) {
		url := baseURL + "/pokemon/" + pokemonName
	
		if data, exists := c.pokeCache.Get(url); exists{
			pokemon := PokemonResp{}
			err := json.Unmarshal(data, &pokemon)
			if err != nil {
				return PokemonResp{}, err
			}
			return pokemon, nil
		}

		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return PokemonResp{}, err
		}

		
		resp, err := c.httpClient.Do(req)
		
		if err != nil {
			return PokemonResp{}, nil
		}
		
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)

		if err != nil {
			return PokemonResp{}, err
		}

		c.pokeCache.Add(url, data)

		pokemon := PokemonResp{}
		
		err = json.Unmarshal(data, &pokemon)
		
		if err != nil {
			return PokemonResp{}, err
		}

		return pokemon, nil
}