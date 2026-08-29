package main 

import "fmt"

func commandHelp(conf *config) error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for _, command := range conf.supportedCommands{
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}