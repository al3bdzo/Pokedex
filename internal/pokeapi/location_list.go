package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
	"github.com/al3bdzo/Pokedex/internal/pokecache"
)

func (c *Client) ListLocations(pageURL *string, pokeCache *pokecache.Cache) (Locations, error){
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var locs Locations

	data, ok := pokeCache.Get(url)
	if ok {
		err := json.Unmarshal(data, &locs)
		if err != nil {
			return Locations{}, err
		}
		return locs, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, err
	}

	err = json.Unmarshal(data, &locs)
	if err != nil {
		return Locations{}, err
	}

	pokeCache.Add(url, data)

	return locs, nil
}