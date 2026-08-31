package main

import (
	"fmt"
	"errors"
)

func commandExplore(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("provide a location name")
	}

	pokes, err := conf.pokeapiClient.ExploreLocation(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring: %s.... \n", args[0])
	fmt.Println("Found Pokemon: ")
	for _, poke := range pokes.PokemonEncounters {
		fmt.Printf(" - %s\n",poke.Pokemon.Name)
	}
	return nil
}