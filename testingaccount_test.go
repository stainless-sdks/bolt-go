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

func TestTestingAccountNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	t.Skip("skipped: tests are disabled for the time being")
	client := bolt.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithPublishableKey("ABC.123.345"),
	)
	_, err := client.Testing.Accounts.New(context.TODO(), bolt.TestingAccountNewParams{
		DeactivateAt: bolt.F(time.Now()),
		EmailState:   bolt.F(bolt.TestingAccountNewParamsEmailStateMissing),
		PhoneState:   bolt.F(bolt.TestingAccountNewParamsPhoneStateMissing),
		HasAddress:   bolt.F(true),
		IsMigrated:   bolt.F(true),
	})
	if err != nil {
		var apierr *bolt.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
