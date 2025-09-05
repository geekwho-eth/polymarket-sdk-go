package clob

import (
	"github.com/geekwho-eth/polymarket-sdk-go/pkg/clob/pricing"
	"github.com/geekwho-eth/polymarket-sdk-go/pkg/sdk/rest"
)

/*
Interface exposes the CLOB (Central Limit Order Book) application services.

Use this entry to access domain-specific services like Pricing.
*/
type Interface interface {
	Pricing() pricing.Service
}

type api struct {
	c rest.Client
}

// New constructs the Data API entry.
func New(c rest.Client) Interface {
	return &api{c: c}
}

func (a *api) Pricing() pricing.Service {
	return pricing.New(a.c)
}
