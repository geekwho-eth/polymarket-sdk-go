package markets

import "time"

/*
MarketList is a simplified representation of a market returned by /markets.
*/
type MarketList struct {
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

// MarketQuery identifies specific market filters that take precedence over options.
type MarketQuery struct {
	Limit     int
	Offset    int
	Order     string
	Ascending bool
	ID        []int

	// Slug identifies a specific market by slug. If provided, it takes precedence
	// over opts.Slug for the request query.
	Slug string // e.g. "clarity-act-signed-into-law-in-2025"

	ClobTokenIDs []string
	ConditionIDs []string
}

// ListMarketsOptions holds query filters for GET /markets.
type ListMarketsOptions struct {
	// TimeoutSeconds overrides per-request timeout if provided.
	TimeoutSeconds *int32
}
