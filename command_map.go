package main

import (
	"encoding/json"
	"fmt"
)

func commandMap(params []string, cfg *config) error {
	urlToCall := "https://pokeapi.co/api/v2/location-area"
	if cfg.Next != "" {
		urlToCall = cfg.Next
	}

	return printResults(urlToCall, cfg)
}

func commandMapBack(params []string, cfg *config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	return printResults(cfg.Previous, cfg)
}

func printResults(url string, cfg *config) error {
	result, err := fetchResults(url, cfg)
	if err != nil {
		return err
	}

	for _, res := range result {
		fmt.Println(res)
	}
	return nil
}

type locationResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []location `json:"results"`
}

type location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func fetchResults(url string, cfg *config) ([]string, error) {
	var data []byte
	cacheData, ok := cfg.cache.Get(url)
	if !ok {
		apiRes, err := callPokeApi(url)
		if err != nil {
			return nil, err
		}
		data = apiRes
		cfg.cache.Add(url, apiRes)
	} else {
		data = cacheData
	}

	jsonData := locationResponse{}

	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, fmt.Errorf("Error parsing response: %v", err)
	}

	var locations []string
	for _, result := range jsonData.Results {
		locations = append(locations, result.Name)
	}

	cfg.Next = jsonData.Next
	cfg.Previous = jsonData.Previous

	return locations, nil
}
