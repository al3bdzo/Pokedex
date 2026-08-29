package main

import (
	"time"
	"github.com/al3bdzo/Pokedex/internal/pokecache"
)

func main() {
	conf := &config{
		supportedCommands: getCommands(),
		pokeCache: pokecache.NewCache(5 * time.Second),
	}
	repl(conf)
}


