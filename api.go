package main

import (
	"fmt"
	"io"
	"net/http"
)

func callPokeApi(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Received non-OK status response: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading respoinse data: %v", err)
	}

	return data, nil
}
