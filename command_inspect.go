package main 

import (
	"fmt"
)


func commandInspect(conf *config, args ...string) error { 
	if len(args) != 1 {
		return fmt.Errorf("you need to supply a pokemon name")
	}

	pokemon, ok := conf.user.cuaghtPokemons[args[0]]

	if !ok {
		return fmt.Errorf("you haven't caught this pokemon!")
	}

	fmt.Printf("Name: %s\nHeight: %v\nWeight: %v\n", pokemon.Name, pokemon.Height, pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %v\n", stat.Stat.Name, stat.Base_stat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}