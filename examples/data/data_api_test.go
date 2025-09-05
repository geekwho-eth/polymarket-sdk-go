package data

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	datacore "github.com/geekwho-eth/polymarket-sdk-go/pkg/data/core"
	datamisc "github.com/geekwho-eth/polymarket-sdk-go/pkg/data/misc"
	sdkclient "github.com/geekwho-eth/polymarket-sdk-go/pkg/sdk"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDataAPI_Misc_GetLiveVolume_Mock(t *testing.T) {
	Convey("Scenario: Data API", t, func() {
		Convey("Background: Call GetLiveVolume", func() {
			const marketID = 27541

			// Mock server returning the provided JSON for GET /live-volume?conditionIds=...
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/live-volume" {
					http.NotFound(w, r)
					return
				}
				idStr := fmt.Sprintf("%d", marketID)
				if r.URL.Query().Get("ID") != idStr {
					http.Error(w, "missing or invalid marketID", http.StatusBadRequest)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(mockHoldersJSON))
			}))
			defer ts.Close()

			c, err := sdkclient.NewData()
			So(err, ShouldBeNil)

			svc := c.DataService().DataAPI().Misc()
			list, err := svc.GetLiveVolume(
				context.Background(),
				&datamisc.GetLiveVolumeQuery{ID: marketID},
				&datamisc.GetLiveVolumeOptions{},
			)
			So(err, ShouldBeNil)
			So(list, ShouldNotBeNil)
			So(len(list), ShouldBeGreaterThan, 0)

			// Basic assertions
			So(len(list), ShouldBeGreaterThan, 0)
			So(list[0].Total, ShouldBeGreaterThan, 0)
			So(len(list[0].Markets), ShouldBeGreaterThan, 0)
			So(list[0].Markets[0].Market, ShouldNotBeEmpty)
			So(list[0].Markets[0].Value, ShouldBeGreaterThan, 0)

			// Pretty print
			pretty, _ := json.MarshalIndent(list, "", "  ")
			t.Logf("GetLiveVolume response:\n%s", string(pretty))
		})
	})
}

func TestDataAPI_Core_GetHolders(t *testing.T) {
	Convey("Scenario: Data API", t, func() {
		Convey("Background: Call GetHolders", func() {
			const conditionID = "0xd744703d2ddcd360ca311fb1a330bf569eda57366292510d75418dc11f76cd00"

			c, err := sdkclient.NewData()
			So(err, ShouldBeNil)

			svc := c.DataService().DataAPI().Core()
			list, err := svc.GetHolders(
				context.Background(),
				&datacore.GetHoldersQuery{Market: []string{conditionID}},
				&datacore.GetHoldersOptions{},
			)
			So(err, ShouldBeNil)
			So(list, ShouldNotBeNil)
			So(len(list), ShouldBeGreaterThan, 0)

			// Basic assertions
			So(len(list), ShouldBeGreaterThan, 0)
			So(list[0].Token, ShouldNotBeEmpty)
			So(len(list[0].Holders), ShouldBeGreaterThan, 0)
			So(list[0].Holders[0].Asset, ShouldNotBeEmpty)
			So(list[0].Holders[0].Amount, ShouldBeGreaterThan, 0)

			// Pretty print
			pretty, _ := json.MarshalIndent(list, "", "  ")
			t.Logf("GetHolders response:\n%s", string(pretty))
		})
	})
}
