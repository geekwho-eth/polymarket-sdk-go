package sdk

import (
	"strings"

	"github.com/geekwho-eth/polymarket-sdk-go/pkg/clob"
	"github.com/geekwho-eth/polymarket-sdk-go/pkg/sdk/rest"
)

var clobCfg *Config

/*
NewClob constructs a Client configured for the CLOB API.

Use the returned client to access ClobService().ClobAPI() and chain into Pricing()/other services.
*/
func NewClob() (*Client, error) {
	if clobCfg == nil {
		clobCfg = &Config{
			BaseURL: ClobEndpoint,
			Timeout: defaultTimeout,
		}
	}
	clobCfg.setDefaults()

	baseURL := strings.TrimRight(clobCfg.BaseURL, "/")

	c := &Client{}

	r := rest.New(baseURL)
	// Wire application-level APIs
	c.clobAPI = clob.New(r)

	return c, nil
}

// ClobServiceInterface accessor to expose ClobAPI().
type ClobServiceInterface interface {
	ClobAPI() clob.Interface
}

// ClobService returns accessor for Clob API with BaseURL wiring.
func (c *Client) ClobService() ClobServiceInterface {
	return &clobService{api: c.clobAPI}
}

// clobService implements DataServiceInterface.
type clobService struct {
	api clob.Interface
}

func (d *clobService) ClobAPI() clob.Interface { return d.api }
