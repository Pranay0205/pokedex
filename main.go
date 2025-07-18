package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}
// cleanInput splits the input text by whitespace, trims spaces, and lowercases each part.
// It also trims leading and trailing whitespace from the input.
func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	fields := strings.Fields(text) 
	for i, field := range fields {
		fields[i] = strings.ToLower(field)
	}
	return fields
}
