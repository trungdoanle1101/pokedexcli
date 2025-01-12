package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/trungdoanle1101/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}
		commandName := cleanedInput[0]
		supportedCommands := getCommands()

		command, ok := supportedCommands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		options := cleanedInput[1:]

		err := command.callback(config, options...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	cleanedWords := strings.Fields(loweredText)

	return cleanedWords
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays the Pokemon in the provided area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch the pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect the pokemon",
			callback:    commandInspect,
		},
	}
}
