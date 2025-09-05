package sdk

import (
	"github.com/geekwho-eth/polymarket-sdk-go/pkg/clob"
	"github.com/geekwho-eth/polymarket-sdk-go/pkg/data"
	"github.com/geekwho-eth/polymarket-sdk-go/pkg/gamma"
)

// Client is the root SDK client for Polymarket APIs.
type Client struct {
	// gamma API wiring
	gammaAPI gamma.Interface

	dataAPI data.Interface

	clobAPI clob.Interface
}
