package markets

import (
	"context"
	"time"

	"github.com/geekwho-eth/polymarket-sdk-go/pkg/sdk/rest"
)

/*
Service defines Gamma Markets operations.
*/
type Service interface {
	// ListMarkets lists markets with optional filters. Both market and opts may be nil.
	// If market is provided and market.Slug is non-empty, it takes precedence over opts.Slug.
	ListMarkets(ctx context.Context, market *MarketQuery, opts *ListMarketsOptions) ([]MarketList, error)
}

type service struct {
	c rest.Client
}

/*
New constructs the Markets service.
*/
func New(c rest.Client) Service {
	return &service{c: c}
}

/*
ListMarkets implements GET /markets with chain-style request.
*/
func (s *service) ListMarkets(ctx context.Context, event *MarketQuery, opts *ListMarketsOptions) (result []MarketList, err error) {
	// timeout
	var timeout time.Duration
	if opts != nil && opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}

	// init result
	result = []MarketList{}

	// chain-style call using event.Slug if provided
	var slug string
	if event != nil {
		slug = event.Slug
	}

	err = s.c.
		Get("/markets").
		Params("slug", slug).
		Timeout(timeout).
		SendRequest(ctx).
		DecodeInto(&result)

	return
}
