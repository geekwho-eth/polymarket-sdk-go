package events

import (
	"context"
	"time"

	"github.com/geekwho-eth/polymarket-sdk-go/pkg/sdk/rest"
)

/*
Service defines Gamma Events operations.
*/
type Service interface {
	// ListEvents lists events with optional filters. Both event and opts may be nil.
	// If event is provided and event.Slug is non-empty, it takes precedence over opts.Slug.
	ListEvents(ctx context.Context, event *EventQuery, opts *ListEventsOptions) ([]EventList, error)
}

type service struct {
	c rest.Client
}

/*
New constructs the Events service.
*/
func New(c rest.Client) Service {
	return &service{c: c}
}

// ListEvents implements GET /events with chain-style request.
func (s *service) ListEvents(ctx context.Context, event *EventQuery, opts *ListEventsOptions) (result []EventList, err error) {
	// timeout
	var timeout time.Duration
	if opts != nil && opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}

	// init result
	result = []EventList{}

	// chain-style call using event.Slug if provided
	var slug string
	if event != nil {
		slug = event.Slug
	}

	err = s.c.
		Get("/events").
		Params("slug", slug).
		Timeout(timeout).
		SendRequest(ctx).
		DecodeInto(&result)

	return
}
