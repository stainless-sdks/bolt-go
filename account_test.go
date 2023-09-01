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

func TestAccountGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	t.Skip("skipped: tests are disabled for the time being")
	client := bolt.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithPublishableKey("ABC.123.345"),
	)
	_, err := client.Accounts.Get(context.TODO())
	if err != nil {
		var apierr *bolt.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountExists(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	t.Skip("skipped: tests are disabled for the time being")
	client := bolt.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithPublishableKey("ABC.123.345"),
	)
	err := client.Accounts.Exists(context.TODO(), bolt.AccountExistsParams{
		Identifier: bolt.F(bolt.AccountExistsParamsIdentifier{
			IdentifierType:  bolt.F(bolt.AccountExistsParamsIdentifierIdentifierTypeEmail),
			IdentifierValue: bolt.F("alice@example.com"),
		}),
	})
	if err != nil {
		var apierr *bolt.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
