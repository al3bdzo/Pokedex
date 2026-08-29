package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
)

func (c *Client) ListLocations(pageURL *string) (Locations, error){
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var locs Locations

	data, ok := c.cache.Get(url)
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

	c.cache.Add(url, data)

	return locs, nil
}