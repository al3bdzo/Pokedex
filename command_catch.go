package main


import (
	"fmt"
	"errors"
	"math/rand"
)

func commandCatch(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("pass the pokemon name")
	}

	pokemon, err := conf.pokeapiClient.CatchPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	catching_chance := catchChance(pokemon.Base_experience)

	if rand.Float64() < catching_chance {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		conf.user.cuaghtPokemons[pokemon.Name] = pokemon
		fmt.Println("You may now inspect it with the inspect command")
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func catchChance(baseExp int) float64 {
	minExp := 36.0
	maxExp := 608.0

	normalized := (float64(baseExp) - minExp) / (maxExp - minExp)

	return 0.90 - normalized*-0.80
}