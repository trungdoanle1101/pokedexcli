package main

import "fmt"

func commandHelp(cfg *config, options ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	supportedCommands := getCommands()

	for _, command := range supportedCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
