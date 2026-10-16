package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if scanner.Err() != nil {
			panic(scanner.Err())
		}
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}
		cmd := words[0]
		var params []string
		if len(words) > 1 {
			params = words[1:]
		}

		if command, exists := cfg.commands[cmd]; exists {
			err := command.callback(params, cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
