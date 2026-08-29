package main

import (
	"fmt"
	"errors"
)



func commandMap(conf *config) error {
	res, err := http.Get(conf.nextURL)
	if err != nil {
		return fmt.Errorf("error getting the response: %v", err)
	}
	defer res.Body.Close()

	var locations locations
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return fmt.Errorf("error parsing json: %v", err)
	}

	conf.nextURL = locations.Next
	conf.previousURL = locations.Previous
	printMap(locations)
	return nil
}

func commandMapb(conf *config) error {
	if conf.previousURL == "" {
		return errors.New("no previous page to go back to")
	}
	res, err := http.Get(conf.previousURL)
	if err != nil {
		return fmt.Errorf("error getting the response: %v", err)
	}
	defer res.Body.Close()

	var locations locations
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return fmt.Errorf("error parsing json: %v", err)
	}

	conf.nextURL = locations.Next
	conf.previousURL = locations.Previous
	printMap(locations)
	return nil
}

func printMap(locs locations) error {
	for _, location := range locs.Results {
		fmt.Println(location.Name)
	}
	return nil
}
