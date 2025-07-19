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


type RespExplore struct {
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	GameIndex            int64                 `json:"game_index"`            
	ID                   int64                 `json:"id"`                    
	Location             location              `json:"location"`              
	Name                 string                `json:"name"`                  
	Names                []Name                `json:"names"`                 
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`    
}

type EncounterMethodRate struct {
	EncounterMethod location                           `json:"encounter_method"`
	VersionDetails  []EncounterMethodRateVersionDetail `json:"version_details"` 
}

type EncounterMethodRateVersionDetail struct {
	Rate    int64    `json:"rate"`   
	Version location `json:"version"`
}

type Name struct {
	Language location `json:"language"`
	Name     string   `json:"name"`    
}

type PokemonEncounter struct {
	Pokemon        location                        `json:"pokemon"`        
	VersionDetails []PokemonEncounterVersionDetail `json:"version_details"`
}

type PokemonEncounterVersionDetail struct {
	EncounterDetails []EncounterDetail `json:"encounter_details"`
	MaxChance        int64             `json:"max_chance"`       
	Version          location          `json:"version"`          
}

type EncounterDetail struct {
	Chance          int64      `json:"chance"`          
	ConditionValues []location `json:"condition_values"`
	MaxLevel        int64      `json:"max_level"`       
	Method          location   `json:"method"`          
	MinLevel        int64      `json:"min_level"`       
}
