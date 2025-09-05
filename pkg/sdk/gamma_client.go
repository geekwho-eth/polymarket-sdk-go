package sdk

import (
	"strings"

	"github.com/geekwho-eth/polymarket-sdk-go/pkg/gamma"
	"github.com/geekwho-eth/polymarket-sdk-go/pkg/sdk/rest"
)

var gammaCfg *Config

/*
NewGamma constructs a Client configured for the Gamma API.

Use the returned client to access GammaService().GammaAPI() and chain into Events()/Markets().
*/
func NewGamma() (*Client, error) {
	if gammaCfg == nil {
		gammaCfg = &Config{
			BaseURL: GammaEndpoint,
			Timeout: defaultTimeout,
		}
	}
	gammaCfg.setDefaults()

	baseURL := strings.TrimRight(gammaCfg.BaseURL, "/")

	c := &Client{}

	r := rest.New(baseURL)
	// Wire application-level APIs
	c.gammaAPI = gamma.New(r)

	return c, nil
}

/*
GammaServiceInterface is the accessor surface for Gamma APIs.

  - GammaAPI(): returns the application-level Gamma entry for chaining,
    e.g. GammaService().GammaAPI().Events().ListEvents(...)
*/
type GammaServiceInterface interface {
	GammaAPI() gamma.Interface
}

/*
GammaService returns an accessor for Gamma APIs.

Use the returned value to obtain the application-level Gamma interface via GammaAPI()
and then chain into concrete services such as Events().
*/
func (c *Client) GammaService() GammaServiceInterface {
	return &gammaService{api: c.gammaAPI}
}

type gammaService struct {
	api gamma.Interface
}

func (g *gammaService) GammaAPI() gamma.Interface { return g.api }
