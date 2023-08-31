// File generated from our OpenAPI spec by Stainless.

package bolt_test

import (
	"context"
	"testing"

	"github.com/bolt/bolt-go"
	"github.com/bolt/bolt-go/internal/testutil"
	"github.com/bolt/bolt-go/option"
)

func TestUsage(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := bolt.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithPublishableKey("ABC.123.345"),
	)
	account, err := client.Accounts.Get(context.TODO())
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", account.Profile.Email)
}
