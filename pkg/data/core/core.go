package core

import (
	"context"
	"strings"
	"time"

	"github.com/geekwho-eth/polymarket-sdk-go/pkg/sdk/rest"
)

// Service exposes miscellaneous Data API operations.
type Service interface {
	// GetHolders fetches holders by condition IDs (comma-separated).
	// Example path: GET /holders?conditionIds=...
	GetHolders(ctx context.Context, query *GetHoldersQuery, opts *GetHoldersOptions) ([]HolderItem, error)
}

type service struct {
	c rest.Client
}

func New(c rest.Client) Service {
	return &service{c: c}
}

// GetHolders implements GET /holders with chain-style request.
func (s *service) GetHolders(ctx context.Context, query *GetHoldersQuery, opts *GetHoldersOptions) (result []HolderItem, err error) {
	// timeout
	var timeout time.Duration
	if opts != nil && opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}

	result = []HolderItem{}

	err = s.c.
		Get("/holders").
		Params("limit", "100").
		Params("market", strings.Join(query.Market, ",")).
		Timeout(timeout).
		SendRequest(ctx).
		DecodeInto(&result)

	return
}
