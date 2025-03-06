package main

import (
	"fmt"
	"strings"
)

func main() {
    fmt.Println("Hello, World!")
}

func cleanInput(text string) []string{
	lowerText := strings.ToLower(text)
	trimmedText := strings.TrimSpace(lowerText)

	words := strings.Fields(trimmedText)
	return words
}




