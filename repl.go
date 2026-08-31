package main

import (
	"strings"
	"fmt"
	"os"
	"bufio"

	"github.com/al3bdzo/Pokedex/internal/pokeapi"
)

type cliCommand struct {
	name string
	description string
	callback func(*config, ...string) error
}

type config struct {
	supportedCommands map[string]cliCommand
	nextURL *string
	previousURL *string
	pokeapiClient pokeapi.Client
	user user
}

type user struct {
	cuaghtPokemons map[string]pokeapi.Pokemon
}


func cleanInput(text string) []string{
	return strings.Fields(strings.ToLower(text))
}

func repl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		arg1 := ""
		if len(input) > 1 {
			arg1 = input[1]
		}

		command, ok := conf.supportedCommands[input[0]]
		if ok {
			err := command.callback(conf, arg1)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help" : {
			name : "help", 
			description: "Displays a help message", 
			callback: commandHelp,
		},
		"exit" : {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"map" : {
			name: "map",
			description: "display the location of 20 location areas in pokemon world",
			callback: commandMap,
		},
		"mapb" : {
			name: "mapb",
			description: "displays the prevuios 20 location areas in pokemon world",
			callback: commandMapb,
		},
		"explore" : {
			name: "explore <area_name>",
			description: "lists all the Pokemon located in an area; take an argument: <area_name>",
			callback: commandExplore,
		},
		"catch" : {
			name: "catch <pokemon_name>",
			description: "catches a pokemon and adds to users's Pokedex; takes an argument: <pokemon_name>",
			callback: commandCatch,
		},
	}
}