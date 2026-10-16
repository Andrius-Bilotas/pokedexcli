package main

import (
	"encoding/json"
	"fmt"
)

type locationArea struct {
	Encounters []pokemonEncounter `json:"pokemon_encounters"`
}

type pokemonEncounter struct {
	Pokemon pokemon `json:"pokemon"`
}

type pokemon struct {
	Name string `json:"name"`
}

func commandExplore(params []string, cfg *config) error {
	if len(params) == 0 {
		return fmt.Errorf("Please provide an area name")
	}
	area := params[0]
	fmt.Printf("Exploring %s...\n", area)
	url := "https://pokeapi.co/api/v2/location-area/" + area

	var data []byte
	cacheData, ok := cfg.cache.Get(url)
	if !ok {
		apiRes, err := callPokeApi(url)
		if err != nil {
			return err
		}
		data = apiRes
		cfg.cache.Add(url, apiRes)
	} else {
		data = cacheData
	}

	jsonData := locationArea{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return fmt.Errorf("Error parsing response: %v", err)
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range jsonData.Encounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
