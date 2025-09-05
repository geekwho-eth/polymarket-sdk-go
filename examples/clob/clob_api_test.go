package data

import (
	"context"
	"encoding/json"
	"testing"

	clobpricing "github.com/geekwho-eth/polymarket-sdk-go/pkg/clob/pricing"
	sdkclient "github.com/geekwho-eth/polymarket-sdk-go/pkg/sdk"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDataAPI_Core_GetHolders(t *testing.T) {
	Convey("Scenario: Clob API", t, func() {
		Convey("Background: Call GetPrice", func() {
			const tokenID = "32813337802210923206070576333957167904670667570617189383542552986489903253171"

			c, err := sdkclient.NewClob()
			So(err, ShouldBeNil)

			svc := c.ClobService().ClobAPI().Pricing()
			res, err := svc.GetPrice(
				context.Background(),
				&clobpricing.GetPriceQuery{TokenID: tokenID, Side: clobpricing.MarketSideSell},
				&clobpricing.GetPriceOptions{},
			)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)

			// Basic assertions
			So(res.Price, ShouldNotBeEmpty)

			// Pretty print
			pretty, _ := json.MarshalIndent(res, "", "  ")
			t.Logf("GetPrice response:\n%s", string(pretty))
		})
	})
}
