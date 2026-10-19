package main

import "fmt"

func commandPokedex(params []string, cfg *config) error {
	if len(cfg.pokemon) == 0 {
		return fmt.Errorf("Your pokedex is empty... Populate it by catching pokemon!")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
