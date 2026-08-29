package main

import (
	"fmt"
	"errors"
)



func commandMap(conf *config) error {
	locs, err := conf.pokeapiClient.ListLocations(conf.nextURL)
	if err != nil {
		return err
	}

	conf.nextURL = locations.Next
	conf.previousURL = locations.Previous

	printMap(locations)
	return nil
}

func commandMapb(conf *config) error {
	if conf.previousURL == nil {
		return errors.New("no previous page to go back to")
	}

	locs, err := conf.pokeapiClient.ListLocations(conf.previousURL)
	if err != nil {
		return err
	}

	conf.nextURL = locations.Next
	conf.previousURL = locations.Previous
	printMap(locations)
	
	return nil
}

func printMap(locs locations) {
	for _, location := range locs.Results {
		fmt.Println(location.Name)
	}
}
