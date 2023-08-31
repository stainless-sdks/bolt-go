// File generated from our OpenAPI spec by Stainless.

package bolt_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bolt/bolt-go"
	"github.com/bolt/bolt-go/internal/testutil"
	"github.com/bolt/bolt-go/option"
)

func TestAccountAddressNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := bolt.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Account.Addresses.New(context.TODO(), bolt.AccountAddressNewParams{
		CountryCode:     bolt.F("US"),
		FirstName:       bolt.F("Alice"),
		LastName:        bolt.F("Baker"),
		Locality:        bolt.F("San Francisco"),
		PostalCode:      bolt.F("94105"),
		StreetAddress1:  bolt.F("535 Mission St, Ste 1401"),
		XPublishableKey: bolt.F("string"),
		Company:         bolt.F("ACME Corporation"),
		Email:           bolt.F("alice@example.com"),
		IsDefault:       bolt.F(true),
		Phone:           bolt.F("+14155550199"),
		Region:          bolt.F("CA"),
		StreetAddress2:  bolt.F("c/o Shipping Department"),
	})
	if err != nil {
		var apierr *bolt.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAddressUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := bolt.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Account.Addresses.Update(
		context.TODO(),
		"D4g3h5tBuVYK9",
		bolt.AccountAddressUpdateParams{
			CountryCode:     bolt.F("US"),
			FirstName:       bolt.F("Alice"),
			LastName:        bolt.F("Baker"),
			Locality:        bolt.F("San Francisco"),
			PostalCode:      bolt.F("94105"),
			StreetAddress1:  bolt.F("535 Mission St, Ste 1401"),
			XPublishableKey: bolt.F("string"),
			Company:         bolt.F("ACME Corporation"),
			Email:           bolt.F("alice@example.com"),
			IsDefault:       bolt.F(true),
			Phone:           bolt.F("+14155550199"),
			Region:          bolt.F("CA"),
			StreetAddress2:  bolt.F("c/o Shipping Department"),
		},
	)
	if err != nil {
		var apierr *bolt.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAddressDelete(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := bolt.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	err := client.Account.Addresses.Delete(
		context.TODO(),
		"D4g3h5tBuVYK9",
		bolt.AccountAddressDeleteParams{
			XPublishableKey: bolt.F("string"),
		},
	)
	if err != nil {
		var apierr *bolt.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}