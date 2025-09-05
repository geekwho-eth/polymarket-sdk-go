package data

import (
	"context"
	"encoding/json"
	"testing"

	clobpricing "github.com/geekwho-eth/polymarket-sdk-go/pkg/clob/pricing"
	sdkclient "github.com/geekwho-eth/polymarket-sdk-go/pkg/sdk"

	. "github.com/smartystreets/goconvey/convey"
)

func TestClobAPI_Pricing(t *testing.T) {
	Convey("Scenario: Clob API", t, func() {
		Convey("Background: Call GetMarketPrice", func() {
			const tokenID = "32813337802210923206070576333957167904670667570617189383542552986489903253171"

			c, err := sdkclient.NewClob()
			So(err, ShouldBeNil)

			svc := c.ClobService().ClobAPI().Pricing()
			res, err := svc.GetMarketPrice(
				context.Background(),
				&clobpricing.GetMarketPriceQuery{TokenID: tokenID, Side: clobpricing.MarketSideSell},
				&clobpricing.GetMarketPriceOptions{},
			)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)

			// Basic assertions
			So(res.Price, ShouldNotBeEmpty)

			// Pretty print
			pretty, _ := json.MarshalIndent(res, "", "  ")
			t.Logf("GetPrice response:\n%s", string(pretty))
		})
		Convey("Background: Call GetMidPointPrice", func() {
			const tokenID = "32813337802210923206070576333957167904670667570617189383542552986489903253171"

			c, err := sdkclient.NewClob()
			So(err, ShouldBeNil)

			svc := c.ClobService().ClobAPI().Pricing()
			res, err := svc.GetMidPointPrice(
				context.Background(),
				&clobpricing.GetMidPointPriceQuery{TokenID: tokenID},
				&clobpricing.GetMidPointPriceOptions{},
			)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)

			// Basic assertions
			So(res.Mid, ShouldNotBeEmpty)

			// Pretty print
			pretty, _ := json.MarshalIndent(res, "", "  ")
			t.Logf("GetPrice response:\n%s", string(pretty))
		})
	})
}
