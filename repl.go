package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"Pokedex/internal/pokeapi"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	config := &Config{
		pokeapiClient: pokeapi.NewClient(),
	}

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			words := cleanInput(input)
			if len(words) == 0 {
				continue
			}

			commandName := words[0]

			command, exists := getCommands()[commandName]
			if exists {
				err := command.callback(config)
				if err != nil {
					fmt.Println("Error executing command:", err)
				}
				continue
			} else {
				fmt.Println("Unknown command:", commandName)
				continue
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			break
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	trimmedText := strings.TrimSpace(lowerText)
	words := strings.Fields(trimmedText)
	return words
}
