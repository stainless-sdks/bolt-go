// File generated from our OpenAPI spec by Stainless.

package bolt_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/bolt/bolt-go"
	"github.com/bolt/bolt-go/internal/testutil"
	"github.com/bolt/bolt-go/option"
)

func TestTestingShipmentNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	t.Skip("skipped: tests are disabled for the time being")
	client := bolt.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithPublishableKey("ABC.123.345"),
	)
	err := client.Testing.Shipments.New(context.TODO(), bolt.TestingShipmentNewParams{
		Status: bolt.F(bolt.TestingShipmentNewParamsStatusInTransit),
		TrackingDetails: bolt.F([]bolt.TestingShipmentNewParamsTrackingDetail{{
			EventDate:   bolt.F("2014-08-21:T14:24:00Z"),
			Status:      bolt.F(bolt.TestingShipmentNewParamsTrackingDetailsStatusUnknown),
			Message:     bolt.F("Billing information received"),
			Locality:    bolt.F("San Francisco"),
			Region:      bolt.F("CA"),
			PostalCode:  bolt.F("string"),
			CountryCode: bolt.F("US"),
		}, {
			EventDate:   bolt.F("2014-08-21:T14:24:00Z"),
			Status:      bolt.F(bolt.TestingShipmentNewParamsTrackingDetailsStatusUnknown),
			Message:     bolt.F("Billing information received"),
			Locality:    bolt.F("San Francisco"),
			Region:      bolt.F("CA"),
			PostalCode:  bolt.F("string"),
			CountryCode: bolt.F("US"),
		}, {
			EventDate:   bolt.F("2014-08-21:T14:24:00Z"),
			Status:      bolt.F(bolt.TestingShipmentNewParamsTrackingDetailsStatusUnknown),
			Message:     bolt.F("Billing information received"),
			Locality:    bolt.F("San Francisco"),
			Region:      bolt.F("CA"),
			PostalCode:  bolt.F("string"),
			CountryCode: bolt.F("US"),
		}}),
		TrackingNumber: bolt.F("MockBolt-143292"),
		DeliveryDate:   bolt.F(time.Now()),
	})
	if err != nil {
		var apierr *bolt.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
