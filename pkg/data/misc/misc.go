package misc

import (
	"context"
	"fmt"
	"time"

	"github.com/geekwho-eth/polymarket-sdk-go/pkg/sdk/rest"
)

// Service exposes miscellaneous Data API operations.
type Service interface {
	// GetLiveVolume fetches holders by condition IDs (comma-separated).
	// Example path: GET /holders?conditionIds=...
	GetLiveVolume(ctx context.Context, query *GetLiveVolumeQuery, opts *GetLiveVolumeOptions) ([]LiveVolumeItem, error)
}

type service struct {
	c rest.Client
}

func New(c rest.Client) Service {
	return &service{c: c}
}

// GetLiveVolume implements GET /holders with chain-style request.
func (s *service) GetLiveVolume(ctx context.Context, query *GetLiveVolumeQuery, opts *GetLiveVolumeOptions) (result []LiveVolumeItem, err error) {
	// timeout
	var timeout time.Duration
	if opts != nil && opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}

	result = []LiveVolumeItem{}

	err = s.c.
		Get("/live-volume").
		Params("id", fmt.Sprintf("%d", query.ID)).
		Timeout(timeout).
		SendRequest(ctx).
		DecodeInto(&result)

	return
}
