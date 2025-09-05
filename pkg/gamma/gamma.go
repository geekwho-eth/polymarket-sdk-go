package gamma

import (
	"github.com/geekwho-eth/polymarket-sdk-go/pkg/gamma/events"
	"github.com/geekwho-eth/polymarket-sdk-go/pkg/gamma/markets"

	"github.com/geekwho-eth/polymarket-sdk-go/pkg/sdk/rest"
)

// Interface defines the Gamma application-level API surface.
type Interface interface {
	// Markets service accessor
	Markets() markets.Service
	// Events service accessor
	Events() events.Service
}

type gammaAPI struct {
	c rest.Client
}

// New constructs the Gamma API using a rest Client.
func New(c rest.Client) Interface {
	return &gammaAPI{c: c}
}

func (g *gammaAPI) Markets() markets.Service {
	return markets.New(g.c)
}

func (g *gammaAPI) Events() events.Service {
	return events.New(g.c)
}
