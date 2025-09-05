package pricing

import (
	"context"
	"time"

	"github.com/geekwho-eth/polymarket-sdk-go/pkg/sdk/rest"
)

/*
Service defines CLOB Pricing operations.
*/
type Service interface {
	// GetMarketPrice returns the current price for a token on a given side.
	// Endpoint: GET /price?token_id=...&side=BUY|SELL
	GetMarketPrice(ctx context.Context, query *GetMarketPriceQuery, opts *GetMarketPriceOptions) (PriceItem, error)
}

type service struct {
	c rest.Client
}

func New(c rest.Client) Service {
	return &service{c: c}
}

/*
GetMarketPrice implements GET /price with the chainable request style:

	c.ClobService().ClobAPI().Pricing().
	  GetMarketPrice(ctx, &GetMarketPriceQuery{TokenID: "...", Side: MarketSideBuy}, &GetMarketPriceOptions{...})
*/
func (s *service) GetMarketPrice(ctx context.Context, query *GetMarketPriceQuery, opts *GetMarketPriceOptions) (result PriceItem, err error) {
	// timeout
	var timeout time.Duration
	if opts != nil && opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}

	result = PriceItem{}

	err = s.c.
		Get("/price").
		Params("token_id", query.TokenID).
		Params("side", query.Side).
		Timeout(timeout).
		SendRequest(ctx).
		DecodeInto(&result)

	return
}
