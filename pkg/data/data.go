package data

import (
	"github.com/geekwho-eth/polymarket-sdk-go/pkg/data/core"
	"github.com/geekwho-eth/polymarket-sdk-go/pkg/data/misc"
	"github.com/geekwho-eth/polymarket-sdk-go/pkg/sdk/rest"
)

// Interface exposes Data API services.
type Interface interface {
	Core() core.Service
	Misc() misc.Service
}

type api struct {
	c rest.Client
}

// New constructs the Data API entry.
func New(c rest.Client) Interface {
	return &api{c: c}
}

func (a *api) Misc() misc.Service {
	return misc.New(a.c)
}

func (a *api) Core() core.Service {
	return core.New(a.c)
}
