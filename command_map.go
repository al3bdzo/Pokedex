package main

import (
	"fmt"
	"errors"
	"github.com/al3bdzo/Pokedex/internal/pokeapi"
)

func commandMap(conf *config, args ...string) error {
	locs, err := conf.pokeapiClient.ListLocations(conf.nextURL)
	if err != nil {
		return err
	}

	conf.nextURL = locs.Next
	conf.previousURL = locs.Previous

	printMap(locs)
	return nil
}

func commandMapb(conf *config, optional ...string) error {
	if conf.previousURL == nil {
		return errors.New("no previous page to go back to")
	}

	locs, err := conf.pokeapiClient.ListLocations(conf.previousURL)
	if err != nil {
		return err
	}

	conf.nextURL = locs.Next
	conf.previousURL = locs.Previous
	printMap(locs)

	return nil
}

func printMap(locs pokeapi.Locations) {
	for _, location := range locs.Results {
		fmt.Println(location.Name)
	}
}
