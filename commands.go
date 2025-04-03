package main

import (
	"fmt"
	"os"

	"Pokedex/internal/pokeapi"
)

type Config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 location areas",
			callback:    commandMapBack,
		},
	}
}

func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()

	return nil
}

func commandExit(cfg *Config) error {
	fmt.Println("Exiting the program.")
	os.Exit(0)
	return nil
}

func commandMap(cfg *Config) error {
	resp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBack(cfg *Config) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	resp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}

	return nil
}
