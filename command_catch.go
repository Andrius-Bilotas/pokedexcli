package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type detailedPokemon struct {
	Height          int           `json:"height"`
	Weight          int           `json:"weight"`
	Base_experience int           `json:"base_experience"`
	Stats           []pokemonStat `json:"stats"`
	Types           []pokemonType `json:"types"`
	pokemon
}

type pokemonStat struct {
	Base_Stat int  `json:"base_stat"`
	Effort    int  `json:"effort"`
	Stat      stat `json:"stat"`
}

type stat struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type pokemonType struct {
	Slot int  `json:"slot"`
	Type stat `json:"type"`
}

const coeficient = 100

func commandCatch(params []string, cfg *config) error {
	if len(params) == 0 {
		return fmt.Errorf("Please provide a pokemon name")
	}
	name := params[0]
	pokemon, err := getPokemonDetails(name, cfg)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	chance := coeficient / (coeficient + float64(pokemon.Base_experience))
	caught := rand.Float64() < chance

	if caught {
		fmt.Printf("%s was caught!\n", name)
		cfg.pokemon[name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}

func getPokemonDetails(pokemon string, cfg *config) (detailedPokemon, error) {
	var data []byte
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	cacheData, ok := cfg.cache.Get(url)
	if !ok {
		apiRes, err := callPokeApi(url)
		if err != nil {
			return detailedPokemon{}, err
		}
		data = apiRes
		cfg.cache.Add(url, apiRes)
	} else {
		data = cacheData
	}

	jsonData := detailedPokemon{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return detailedPokemon{}, fmt.Errorf("Error parsing response: %v", err)
	}
	return jsonData, nil
}
