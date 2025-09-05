package pricing

// MarketSide* constants define the side of the market for pricing queries.
const (
	MarketSideBuy  = "BUY"
	MarketSideSell = "SELL"
)

// GetPriceOptions holds optional parameters for GetPrice.
type GetPriceOptions struct {
	// TimeoutSeconds overrides per-request timeout if provided.
	TimeoutSeconds *int32
}

// PriceItem corresponds to one token entry in the Price response.
type PriceItem struct {
	Price string `json:"price"`
}

/*
GetPriceQuery specifies pricing query parameters.

- TokenID: the market token identifier to quote
- Side: BUY or SELL (use MarketSideBuy/MarketSideSell)
*/
type GetPriceQuery struct {
	TokenID string
	Side    string
}
