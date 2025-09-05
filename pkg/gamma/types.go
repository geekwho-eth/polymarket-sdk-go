package gamma

import "time"

// Market represents a Polymarket market (simplified).
type Market struct {
	ID        string     `json:"id"`
	Question  string     `json:"question"`
	Ticker    string     `json:"ticker"`
	Status    string     `json:"status"` // e.g., active, closed
	Currency  string     `json:"currency,omitempty"`
	Category  string     `json:"category,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	CloseTime *time.Time `json:"close_time,omitempty"`

	// Additional commonly seen fields
	Slug     string   `json:"slug,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	Source   string   `json:"source,omitempty"`
	Universe string   `json:"universe,omitempty"`
}

// ListMarketsOptions provides query filters for listing markets.
type ListMarketsOptions struct {
	Active   *bool  // filter active markets
	Search   string // substring or query keywords
	Category string // category filter
	Limit    int    // page size
	Cursor   string // pagination cursor
}
