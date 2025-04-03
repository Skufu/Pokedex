package pokeapi

import (
	"net/http"
	"time"
)

// Client is the PokeAPI client
type Client struct {
	httpClient http.Client
}

// NewClient creates a new PokeAPI client with the given timeout
func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
