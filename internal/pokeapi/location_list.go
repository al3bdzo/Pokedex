package pokeapi


import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (Locations, error){
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close

	var locs locations{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locs); err != nil {
		return locations{}, err
	}

	return locs, nil
}