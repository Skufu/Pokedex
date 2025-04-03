package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetLocationArea retrieves details about a specific location area by name
func (c *Client) GetLocationArea(name string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area/" + name

	// Check if the data is already in the cache
	if cachedData, ok := c.cache.Get(url); ok {
		// Found in cache, deserialize and return
		fmt.Println("Cache hit!")
		locationArea := LocationAreaResponse{}
		err := json.Unmarshal(cachedData, &locationArea)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return locationArea, nil
	}

	fmt.Println("Cache miss, fetching from API...")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	// Handle 404 error
	if resp.StatusCode == http.StatusNotFound {
		return LocationAreaResponse{}, fmt.Errorf("location area %s not found", name)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	// Store the raw response in the cache
	c.cache.Add(url, dat)

	locationArea := LocationAreaResponse{}
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationArea, nil
}
