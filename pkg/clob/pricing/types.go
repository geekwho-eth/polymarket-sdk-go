package pricing

// MarketSide* constants define the side of the market for pricing queries.
const (
	MarketSideBuy  = "BUY"
	MarketSideSell = "SELL"
)

// GetMarketPriceOptions holds optional parameters for GetMarketPrice.
type GetMarketPriceOptions struct {
	// TimeoutSeconds overrides per-request timeout if provided.
	TimeoutSeconds *int32
}

// PriceItem corresponds to one token entry in the Price response.
type PriceItem struct {
	Price float64 `json:"price,string"`
}

/*
GetMarketPriceQuery specifies pricing query parameters.

- TokenID: the market token identifier to quote
- Side: BUY or SELL (use MarketSideBuy/MarketSideSell)
*/
type GetMarketPriceQuery struct {
	TokenID string
	Side    string
}

/*
GetMidPointPriceQuery specifies pricing query parameters.

- TokenID: the market token identifier to quote
*/
type GetMidPointPriceQuery struct {
	TokenID string
}

// GetMidPointPriceOptions holds optional parameters for GetMarketPrice.
type GetMidPointPriceOptions struct {
	// TimeoutSeconds overrides per-request timeout if provided.
	TimeoutSeconds *int32
}

// MidPointPriceItem corresponds to one token entry in the Price response.
type MidPointPriceItem struct {
	Mid float64 `json:"mid,string"`
}
