# polymarket-sdk-go

A lightweight, chainable Go SDK for Polymarket, featuring a layered architecture (Project → Application → Service) inspired by [marmotedu-sdk-go](https://github.com/marmotedu/marmotedu-sdk-go). Its fluent, chainable API (e.g., GammaService().GammaAPI().Events().ListEvents(...)) ensures a clean and elegant codebase.

Layers
- Project: client (root SDK client)
- Application:
  - Gamma API via GammaService().GammaAPI()
  - Data API via DataService().DataAPI()
  - CLOB API via ClobService().ClobAPI()
- Service:
  - Gamma: Events().ListEvents(), Markets().ListMarkets()
  - Data: Misc().GetLiveVolume()
  - CLOB: Pricing().GetPrice()

Implemented APIs
- Gamma
  - Events
    - [x]  [/events](https://docs.polymarket.com/api-reference/events/list-events) -> ListEvents
  - Markets:
    - [x] [/markets](https://docs.polymarket.com/api-reference/markets/list-markets) -> ListMarkets
- Data
  - Misc:
    - [x] [/live-volume](https://docs.polymarket.com/api-reference/misc/get-live-volume-for-an-event) -> GetLiveVolume
- CLOB
  - Pricing:
    - [x] [/price](https://docs.polymarket.com/api-reference/pricing/get-market-price) -> GetPrice

Install
```bash
go get github.com/geekwho-eth/polymarket-sdk-go
```

Quick Start: Gamma Events
```go
package main

import (
  "context"
  "fmt"

  gammaevents "github.com/geekwho-eth/polymarket-sdk-go/pkg/gamma/events"
  sdk "github.com/geekwho-eth/polymarket-sdk-go/pkg/sdk"
)

func main() {
  // Use default endpoint https://gamma-api.polymarket.com
  c, err := sdk.NewGamma()
  if err != nil {
    panic(err)
  }

  svc := c.GammaService().GammaAPI().Events()
  list, err := svc.ListEvents(
    context.Background(),
    &gammaevents.EventQuery{Slug: "clarity-act-signed-into-law-in-2025"},
    &gammaevents.ListEventsOptions{},
  )
  if err != nil {
    panic(err)
  }
  fmt.Printf("events: %d\n", len(list))
  if len(list) > 0 {
    fmt.Printf("event[0]: %s - %s\n", list[0].ID, list[0].Title)
  }
}
```

Quick Start: Data API (Misc)
```go
package main

import (
  "context"
  "fmt"

  "github.com/geekwho-eth/polymarket-sdk-go/pkg/data/misc"
  sdk "github.com/geekwho-eth/polymarket-sdk-go/pkg/sdk"
)

func main() {
  // Uses default https://data-api.polymarket.com
  c, err := sdk.NewData()
  if err != nil {
    panic(err)
  }

  svc := c.DataService().DataAPI().Misc()
  resp, err := svc.GetLiveVolume(
    context.Background(),
    "0xd744703d2ddcd360ca311fb1a330bf569eda57366292510d75418dc11f76cd00",
    &misc.GetLiveVolumeOptions{},
  )
  if err != nil {
    panic(err)
  }
  fmt.Printf("tokens: %d\n", len(resp))
}
```

Quick Start: CLOB Pricing
```go
package main

import (
  "context"
  "fmt"

  clobpricing "github.com/geekwho-eth/polymarket-sdk-go/pkg/clob/pricing"
  sdk "github.com/geekwho-eth/polymarket-sdk-go/pkg/sdk"
)

func main() {
  // Uses default CLOB endpoint (configured in SDK)
  c, err := sdk.NewClob()
  if err != nil {
    panic(err)
  }

  svc := c.ClobService().ClobAPI().Pricing()
  res, err := svc.GetPrice(
    context.Background(),
    &clobpricing.GetPriceQuery{
      TokenID: "32813337802210923206070576333957167904670667570617189383542552986489903253171",
      Side:    clobpricing.MarketSideSell, // or clobpricing.MarketSideBuy
    },
    &clobpricing.GetPriceOptions{},
  )
  if err != nil {
    panic(err)
  }
  fmt.Printf("price: %s\n", res.Price)
}
```

Design
- pkg/sdk: SDK client bootstrapping and service accessors (NewGamma, NewData, NewClob, GammaService, DataService, ClobService)
- pkg/gamma: Gamma API application surface (Events, Markets)
- pkg/data: Data API application surface (Misc)
- pkg/clob: CLOB API application surface (Pricing; more like Orderbook, Spreads, Order to be added)

Notes
- Default endpoints:
  - Gamma: https://gamma-api.polymarket.com
  - Data: https://data-api.polymarket.com
  - CLOB: https://clob.polymarket.com
- All services follow a simple, chainable call style:
  - Get("/path").Params(...).Timeout(...).Do(ctx).DecodeInto(&out)
- For tests, prefer net/http/httptest to avoid real network calls.
