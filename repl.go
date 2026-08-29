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
	callback func(*config) error
}

type config struct {
	supportedCommands map[string]cliCommand
	nextURL *string
	previousURL *string
	pokeapiClient pokepai.Client
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

		command, ok := conf.supportedCommands[input[0]]
		if ok {
			err := command.callback(conf)
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
	}
}