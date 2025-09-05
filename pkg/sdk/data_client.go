package sdk

import (
	"strings"

	"github.com/geekwho-eth/polymarket-sdk-go/pkg/data"
	"github.com/geekwho-eth/polymarket-sdk-go/pkg/sdk/rest"
)

var dataCfg *Config

/*
NewData constructs a Client configured for the Data API.

Use the returned client to access DataService().DataAPI() and chain into sub-services (e.g., Misc()).
*/
func NewData() (*Client, error) {
	if dataCfg == nil {
		dataCfg = &Config{
			BaseURL: DataEndpoint,
			Timeout: defaultTimeout,
		}
	}
	dataCfg.setDefaults()

	baseURL := strings.TrimRight(dataCfg.BaseURL, "/")

	c := &Client{}

	r := rest.New(baseURL)
	// Wire application-level APIs
	c.dataAPI = data.New(r)

	return c, nil
}

// DataServiceInterface accessor to expose DataAPI().
type DataServiceInterface interface {
	DataAPI() data.Interface
}

/*
DataService returns an accessor for the Data API.

It exposes a fluent entry to DataAPI(), from which you can drill down to sub-services
(e.g., Misc()) and call APIs like GetLiveVolume().
*/
func (c *Client) DataService() DataServiceInterface {
	return &dataService{api: c.dataAPI}
}

// dataService implements DataServiceInterface.
type dataService struct {
	api data.Interface
}

func (d *dataService) DataAPI() data.Interface { return d.api }
