package pokeapi 

import (
	"net/http"
	"encoding/json"
	"io"
	"fmt"
)


func (c *Client) ExploreLocation(location string) (LocationPokemones, error) {
	url := baseURL + "/location-area/" + location

	var pokes LocationPokemones 

	data, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(data, &pokes)
		if err != nil {
			return LocationPokemones{}, err
		}
		return pokes, nil
	}
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationPokemones{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationPokemones{}, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return LocationPokemones{}, fmt.Errorf("HTTP error: %v %s", res.StatusCode, http.StatusText(res.StatusCode))
	}


	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationPokemones{}, err
	}

	err = json.Unmarshal(data, &pokes)
	if err != nil {
		return LocationPokemones{}, err
	}

	c.cache.Add(url, data)

	return pokes, nil
}