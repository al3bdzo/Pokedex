package main

import (
	"time"
	"github.com/al3bdzo/Pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	conf := &config{
		supportedCommands: getCommands(),
		pokeapiClient: pokeClient,
		user: user{
			cuaghtPokemons: make(map[string]pokeapi.Pokemon),
		},
	}
	repl(conf)
}


