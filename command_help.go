package main

import "fmt"

func commandHelp(params []string, cfg *config) error {
	var cmdList []string
	for _, command := range cfg.commands {
		cmdList = append(cmdList, fmt.Sprintf("%s: %s", command.name, command.description))
	}
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range cmdList {
		fmt.Println(cmd)
	}
	return nil
}
