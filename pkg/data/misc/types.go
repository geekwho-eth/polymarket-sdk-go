package misc

// GetLiveVolumeOptions holds optional parameters for GetLiveVolume.
type GetLiveVolumeOptions struct {
	// TimeoutSeconds overrides per-request timeout if provided.
	TimeoutSeconds *int32
}

// LiveVolumeItem corresponds to one token entry in the holders response.
type LiveVolumeItem struct {
	Total   float64            `json:"total"`
	Markets []LiveVolumeMarket `json:"markets"`
}

type LiveVolumeMarket struct {
	Market string  `json:"market"`
	Value  float64 `json:"value"`
}

// CommonQuery identifies specific market filters that take precedence over options.
type CommonQuery struct {
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

type GetLiveVolumeQuery struct {
	ID int
}
