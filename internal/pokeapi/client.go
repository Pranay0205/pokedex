package pokeapi

import (
	"net/http"
	"time"

	"github.com/Pranay0205/pokedexcli/internal/pokecache"
)


type Client struct {
	httpClient http.Client
	pokeCache *pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		pokeCache : pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}