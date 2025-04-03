package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// GetPokemon retrieves details about a specific Pokemon by name
func (c *Client) GetPokemon(name string) (Pokemon, error) {
	// Convert name to lowercase to match API format
	name = strings.ToLower(name)
	url := baseURL + "/pokemon/" + name

	// Check if the data is already in the cache
	if cachedData, ok := c.cache.Get(url); ok {
		// Found in cache, deserialize and return
		fmt.Println("Cache hit!")
		pokemon := Pokemon{}
		err := json.Unmarshal(cachedData, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	fmt.Println("Cache miss, fetching from API...")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	// Handle 404 error
	if resp.StatusCode == http.StatusNotFound {
		return Pokemon{}, fmt.Errorf("pokemon %s not found", name)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// Store the raw response in the cache
	c.cache.Add(url, dat)

	pokemon := Pokemon{}
	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
