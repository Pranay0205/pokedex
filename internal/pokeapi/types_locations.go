package pokeapi

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type RespLocation struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Result   []location `json:"results"`
}