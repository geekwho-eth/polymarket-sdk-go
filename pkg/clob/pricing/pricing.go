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
	// GetPrice returns the current price for a token on a given side.
	// Endpoint: GET /price?token_id=...&side=BUY|SELL
	GetPrice(ctx context.Context, query *GetPriceQuery, opts *GetPriceOptions) (PriceItem, error)
}

type service struct {
	c rest.Client
}

func New(c rest.Client) Service {
	return &service{c: c}
}

/*
GetPrice implements GET /price with the chainable request style:

	c.ClobService().ClobAPI().Pricing().
	  GetPrice(ctx, &GetPriceQuery{TokenID: "...", Side: MarketSideBuy}, &GetPriceOptions{...})
*/
func (s *service) GetPrice(ctx context.Context, query *GetPriceQuery, opts *GetPriceOptions) (result PriceItem, err error) {
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
