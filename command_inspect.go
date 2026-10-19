package main

import "fmt"

func commandInspect(params []string, cfg *config) error {
	if len(params) == 0 {
		return fmt.Errorf("Please provide a pokemon name")
	}
	pokemonName := params[0]

	pokemon, ok := cfg.pokemon[pokemonName]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.Base_Stat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokeType.Type.Name)
	}
	return nil
}
