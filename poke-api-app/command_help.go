package main

import "fmt"

func callbackHelp() error {
	availableCommands := getCommands()
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Here are your available commands:")
	for _, command := range availableCommands {
		fmt.Printf(" - %s: %s\n", command.name, command.description)
	}
	fmt.Println("")
	return nil
}
