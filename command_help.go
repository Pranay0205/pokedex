package main

import (
	"fmt"
)

func commandHelp(c *config, arg []string) error {
	fmt.Printf("%s%sWelcome to the Pokedex!%s\n", ColorBold, ColorCyan, ColorReset)
	fmt.Printf("%s%sCommands:%s\n", ColorBold, ColorYellow, ColorReset)

	if len(registred_commands) == 0 {
		return fmt.Errorf("no commands registry")
	}

	i := 1
	for _, value := range registred_commands {
		fmt.Printf("%s%d) %s%s%s: %s%s%s\n", 
			ColorGreen, i, 
			ColorBlue, value.name, ColorReset,
			ColorWhite, value.description, ColorReset)
		i++
	}

	return nil
}