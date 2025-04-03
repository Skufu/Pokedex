package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Define the command registry
var commandRegistry = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin) // Create a new scanner to read input

	for {
		if scanner.Scan() { // Wait for user input
			input := scanner.Text()           // Get the user's input
			cleanedInput := cleanInput(input) // Clean the input

			if len(cleanedInput) > 0 {
				firstWord := cleanedInput[0] // Capture the first word

				// Check if the command exists in the registry
				if command, exists := commandRegistry[firstWord]; exists {
					if err := command.callback(); err != nil {
						fmt.Println("Error executing command:", err)
					}
				} else {
					fmt.Println("Unknown command:", firstWord)
				}
			}
		}

		if err := scanner.Err(); err != nil { // Check for errors
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

// Function to check if the word is unwanted
func isUnwantedWord(word string) bool {
	unwantedWords := []string{"better", "kinda", "is"}
	for _, unwanted := range unwantedWords {
		if word == unwanted {
			return true
		}
	}
	return false
}

// New function to handle exit command
func commandExit() error {
	fmt.Println("Exiting the program.")
	os.Exit(0) // Exit the program
	return nil // This line will never be reached
}
