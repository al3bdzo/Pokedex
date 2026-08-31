package main 

import (
	"fmt"
)

func commandPokedex(conf *config, args ...string) error{
	if len(conf.user.cuaghtPokemons) == 0 {
		return fmt.Errorf("You didn't catch any pokemons!")
	}

	fmt.Println("Your Pokedex:")
	for key, _ := range conf.user.cuaghtPokemons {
		fmt.Printf("  - %s\n", key)
	}
	return nil
}