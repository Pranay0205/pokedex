package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)



func (c *Client) ListLocations(pageURL *string) (RespLocation, error){
		
		url := baseURL + "/location-area"

		if pageURL != nil {
			url = *pageURL
		}

		locationResp := RespLocation{}

		if val, exists := c.pokeCache.Get(url); exists {
			err := json.Unmarshal(val, &locationResp)

			if err != nil {
				return RespLocation{}, err
			}

			return locationResp, nil
		}

		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return RespLocation{}, err
		}
		
		resp, err := c.httpClient.Do(req)

		if err != nil {
			return RespLocation{}, err
		}
		
		defer resp.Body.Close()
		
		data, err := io.ReadAll(resp.Body)
	
		if err != nil {
			return RespLocation{}, err
		}

		c.pokeCache.Add(url, data)
		
		err = json.Unmarshal(data, &locationResp)
		if err != nil {
			return RespLocation{}, err
		}

		return locationResp, nil

}

func (c *Client) ExploreLocation(locationName *string) (RespExplore, error) {

		url := baseURL + "/location-area/" + *locationName
		
		exploreResp := RespExplore{}
		if val, exists := c.pokeCache.Get(url); exists {
			err := json.Unmarshal(val, &exploreResp)

			if err != nil {
				return RespExplore{}, err
			}

			return exploreResp, nil
		}

		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return RespExplore{}, err
		}

		resp, err := c.httpClient.Do(req)
		
		if err != nil {
			return RespExplore{}, err
		}

		data, err := io.ReadAll(resp.Body)

		if err != nil {
			return RespExplore{}, err
		}

		err = json.Unmarshal(data, &exploreResp)

		if err != nil {
			return RespExplore{}, err
		}

		return exploreResp, nil

}