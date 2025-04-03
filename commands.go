package main

import (
	"fmt"
	"math/rand"
	"os"

	"Pokedex/internal/pokeapi"
)

// Initialize random seed for catching pokemon
func init() {
	// In newer Go versions (1.20+), we don't need to explicitly seed the random number generator
	// For older versions compatibility, you can uncomment:
	// rand.Seed(time.Now().UnixNano())
}

type Config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
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
		"explore": {
			name:        "explore",
			description: "Display a list of Pokémon in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokémon by name",
			callback:    commandCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View your caught Pokémon",
			callback:    commandPokedex,
		},
		"inspect": {
			name:        "inspect",
			description: "View detailed information about a caught Pokémon",
			callback:    commandInspect,
		},
	}
}

func commandHelp(cfg *Config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()

	return nil
}

func commandExit(cfg *Config, args []string) error {
	fmt.Println("Exiting the program.")
	os.Exit(0)
	return nil
}

func commandMap(cfg *Config, args []string) error {
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

func commandMapBack(cfg *Config, args []string) error {
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

func commandExplore(cfg *Config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a location area name")
		return nil
	}

	locationAreaName := args[0]
	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationArea.Name)
	fmt.Println("Found Pokémon:")

	if len(locationArea.PokemonEncounters) == 0 {
		fmt.Println("No Pokémon found in this area!")
		return nil
	}

	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *Config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a Pokémon name")
		return nil
	}

	// Initialize the caught Pokemon map if it doesn't exist
	if cfg.caughtPokemon == nil {
		cfg.caughtPokemon = make(map[string]pokeapi.Pokemon)
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	// Get Pokemon information
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	// Calculate catch probability (lower base experience = easier to catch)
	// Base formula: catchChance = 50 - baseExperience/10 with minimum of 10%
	catchChance := 50.0 - float64(pokemon.BaseExperience)/10.0
	if catchChance < 10.0 {
		catchChance = 10.0
	}
	if catchChance > 90.0 {
		catchChance = 90.0
	}

	// Generate random number and check if caught
	randNum := rand.Float64() * 100.0
	caught := randNum <= catchChance

	if caught {
		fmt.Printf("%s was caught!\n", pokemonName)
		// Store the caught Pokemon in the Pokedex
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

func commandPokedex(cfg *Config, args []string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("Your Pokedex is empty. Try catching some Pokemon first!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s (Base Experience: %d)\n", pokemon.Name, pokemon.BaseExperience)
	}

	return nil
}

func commandInspect(cfg *Config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Please provide the name of a Pokemon to inspect")
		return nil
	}

	pokemonName := args[0]
	pokemon, exists := cfg.caughtPokemon[pokemonName]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf("  - %s\n", typeInfo.Type.Name)
	}

	return nil
}
