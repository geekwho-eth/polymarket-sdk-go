package events

import "time"

// EventList is a simplified Polymarket Event model for listing.
// Fields included are based on the official docs and your provided JSON sample.
type EventList struct {
	ID               string     `json:"id"`
	Ticker           string     `json:"ticker"`
	Slug             string     `json:"slug"`
	Title            string     `json:"title"`
	Description      string     `json:"description"`
	ResolutionSource string     `json:"resolutionSource"`
	StartDate        *time.Time `json:"startDate"`
	CreationDate     *time.Time `json:"creationDate"`
	EndDate          *time.Time `json:"endDate"`
	Image            string     `json:"image"`
	Icon             string     `json:"icon"`
	Active           bool       `json:"active"`
	Closed           bool       `json:"closed"`
	Archived         bool       `json:"archived"`
	New              bool       `json:"new"`
	Featured         bool       `json:"featured"`
	Restricted       bool       `json:"restricted"`

	Liquidity    float64 `json:"liquidity"`
	Volume       float64 `json:"volume"`
	OpenInterest float64 `json:"openInterest"`

	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`

	// Nested markets (simplified)
	Markets []EventMarket `json:"markets"`

	// Tags (simplified)
	Tags []EventTag `json:"tags"`
}

type EventMarket struct {
	ID            string     `json:"id"`
	Question      string     `json:"question"`
	ConditionID   string     `json:"conditionId"`
	Slug          string     `json:"slug"`
	EndDate       *time.Time `json:"endDate"`
	StartDate     *time.Time `json:"startDate"`
	Image         string     `json:"image"`
	Icon          string     `json:"icon"`
	Description   string     `json:"description"`
	Outcomes      string     `json:"outcomes"`
	OutcomePrices string     `json:"outcomePrices"`
	Volume        string     `json:"volume"`
	Active        bool       `json:"active"`
	Closed        bool       `json:"closed"`
	//ClobTokenIds  []string   `json:"clobTokenIds"`
}

type EventTag struct {
	ID        string     `json:"id"`
	Label     string     `json:"label"`
	Slug      string     `json:"slug"`
	ForceShow bool       `json:"forceShow"`
	ForceHide bool       `json:"forceHide"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// EventQuery identifies specific event filters that take precedence over options.
type EventQuery struct {
	// Slug identifies a specific event by slug. If provided, it takes precedence
	// over opts.Slug for the request query.
	Slug string
}

// ListEventsOptions holds query filters for GET /events.
type ListEventsOptions struct {
	// Slug kept for backward compatibility; ignored if event.Slug is provided.
	Slug     string // e.g. "clarity-act-signed-into-law-in-2025"
	Search   string
	Active   *bool
	Category string
	Limit    int
	Cursor   string

	// TimeoutSeconds overrides per-request timeout if provided.
	TimeoutSeconds *int32
}
