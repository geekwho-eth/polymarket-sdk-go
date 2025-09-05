package sdk

import "time"

const (
	GammaEndpoint = "https://gamma-api.polymarket.com"
	DataEndpoint  = "https://data-api.polymarket.com"
	ClobEndpoint  = "https://clob.polymarket.com"

	defaultTimeout = 3 * time.Second
)

// Config holds SDK configuration.
type Config struct {
	// BaseURL is the API base endpoint for Gamma API.
	BaseURL string

	// APIKey is optional; if set, requests will include Authorization: Bearer <APIKey>.
	APIKey string

	// Timeout configures the underlying HTTP client timeout.
	Timeout time.Duration
}

func (c *Config) setDefaults() {
	if c.Timeout == 0 {
		c.Timeout = defaultTimeout
	}
	if c.BaseURL == "" {
		c.BaseURL = GammaEndpoint
	}
}
