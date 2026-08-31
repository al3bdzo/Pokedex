package pokeapi

type Locations struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct{
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"results"`
}

type LocationPokemones struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Base_experience int `json:"base_expeience"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats []struct {
		Base_stat int `json:"base_stat"`
		Stat struct{
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct{
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}