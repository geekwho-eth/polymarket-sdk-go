package gamma

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	gammaevents "github.com/geekwho-eth/polymarket-sdk-go/pkg/gamma/events"
	gammamarkets "github.com/geekwho-eth/polymarket-sdk-go/pkg/gamma/markets"
	sdkclient "github.com/geekwho-eth/polymarket-sdk-go/pkg/sdk"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGammaAPI_Events_ListEvents(t *testing.T) {
	Convey("Scenario: Gamma API", t, func() {
		Convey("Background: Call ListEvents", func() {
			const wantSlug = "clarity-act-signed-into-law-in-2025"

			c, err := sdkclient.NewGamma()
			So(err, ShouldBeNil)

			svc := c.GammaService().GammaAPI().Events()
			list, err := svc.ListEvents(
				context.Background(),
				&gammaevents.EventQuery{Slug: wantSlug},
				&gammaevents.ListEventsOptions{},
			)
			So(err, ShouldBeNil)
			So(list, ShouldNotBeNil)
			So(list, ShouldHaveLength, 1)

			ev := list[0]
			So(ev.Slug, ShouldEqual, wantSlug)
			So(ev.ID, ShouldEqual, "27541")
			So(ev.Title, ShouldEqual, "Clarity Act signed into law in 2025?")
			So(ev.Markets, ShouldHaveLength, 1)
			So(ev.Markets[0].ConditionID, ShouldEqual, "0xd744703d2ddcd360ca311fb1a330bf569eda57366292510d75418dc11f76cd00")
			//So(ev.Markets[0].ClobTokenIds, ShouldHaveLength, 2)
			//So(ev.Markets[0].ClobTokenIds[0], ShouldEqual, "32813337802210923206070576333957167904670667570617189383542552986489903253171")

			// Print the parsed response as JSON for visual inspection.
			pretty, _ := json.MarshalIndent(list, "", "  ")
			t.Logf("ListEvents response:\n%s", string(pretty))
		})
	})
}

func TestGammaAPI_Events_ListEvents_Mock(t *testing.T) {
	Convey("Scenario: Gamma API", t, func() {
		Convey("Background: Call ListEvents with httptest", func() {
			const wantSlug = "clarity-act-signed-into-law-in-2025"

			// Mock server returning the provided JSON for GET /events?slug=...
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/events" {
					http.NotFound(w, r)
					return
				}
				if r.URL.Query().Get("slug") != wantSlug {
					http.Error(w, "missing or invalid slug", http.StatusBadRequest)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(mockListEventsJSON))
			}))
			defer ts.Close()

			c, err := sdkclient.NewGamma()
			So(err, ShouldBeNil)

			svc := c.GammaService().GammaAPI().Events()
			list, err := svc.ListEvents(
				context.Background(),
				&gammaevents.EventQuery{Slug: wantSlug},
				&gammaevents.ListEventsOptions{},
			)
			So(err, ShouldBeNil)
			So(list, ShouldNotBeNil)
			So(list, ShouldHaveLength, 1)

			ev := list[0]
			So(ev.Slug, ShouldEqual, wantSlug)
			So(ev.ID, ShouldEqual, "27541")
			So(ev.Title, ShouldEqual, "Clarity Act signed into law in 2025?")
			So(ev.Markets, ShouldHaveLength, 1)
			So(ev.Markets[0].ConditionID, ShouldEqual, "0xd744703d2ddcd360ca311fb1a330bf569eda57366292510d75418dc11f76cd00")
			//So(ev.Markets[0].ClobTokenIds, ShouldHaveLength, 2)
			//So(ev.Markets[0].ClobTokenIds[0], ShouldEqual, "32813337802210923206070576333957167904670667570617189383542552986489903253171")

			// Print the parsed response as JSON for visual inspection.
			pretty, _ := json.MarshalIndent(list, "", "  ")
			t.Logf("ListEvents response:\n%s", string(pretty))
		})
	})
}

func TestGammaAPI_Markets_ListMarkets(t *testing.T) {
	Convey("Scenario: Gamma API", t, func() {
		Convey("Background: Call ListMarkets", func() {
			const wantSlug = "clarity-act-signed-into-law-in-2025"

			c, err := sdkclient.NewGamma()
			So(err, ShouldBeNil)

			svc := c.GammaService().GammaAPI().Markets()
			list, err := svc.ListMarkets(
				context.Background(),
				&gammamarkets.MarketQuery{Slug: wantSlug},
				&gammamarkets.ListMarketsOptions{},
			)
			So(err, ShouldBeNil)
			So(list, ShouldNotBeNil)
			So(list, ShouldHaveLength, 1)

			ev := list[0]
			So(ev.Slug, ShouldEqual, wantSlug)
			So(ev.ID, ShouldEqual, "553363")
			So(ev.ConditionID, ShouldEqual, "0xd744703d2ddcd360ca311fb1a330bf569eda57366292510d75418dc11f76cd00")

			// Print the parsed response as JSON for visual inspection.
			pretty, _ := json.MarshalIndent(list, "", "  ")
			t.Logf("ListEvents response:\n%s", string(pretty))
		})
	})
}
