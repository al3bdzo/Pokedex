package pokeapi


import (
	"net/http"
	"encoding/json"
	"io"
	"fmt"
)


func (c *Client) CatchPokemon(name string) (Pokemon, error){
	url := baseURL + "/pokemon/" + name 

	var pokemon Pokemon 

	data, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return Pokemon{}, fmt.Errorf("HTTP error: %v %s", res.StatusCode, http.StatusText(res.StatusCode))
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	return pokemon, nil
}