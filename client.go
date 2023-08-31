// File generated from our OpenAPI spec by Stainless.

package bolt

import (
	"os"

	"github.com/bolt/bolt-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the bolt API. You should not instantiate this client directly,
// and instead use the [NewClient] method instead.
type Client struct {
	Options  []option.RequestOption
	Accounts *AccountService
	Payments *PaymentService
	Guest    *GuestService
	Merchant *MerchantService
	Webhooks *WebhookService
	Testing  *TestingService
}

// NewClient generates a new client with the default option read from the
// environment (BOLT_API_KEY). The option passed in as arguments are applied after
// these default arguments, and all option will be passed down to the services and
// requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("BOLT_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	opts = append(defaults, opts...)

	r = &Client{Options: opts}

	r.Accounts = NewAccountService(opts...)
	r.Payments = NewPaymentService(opts...)
	r.Guest = NewGuestService(opts...)
	r.Merchant = NewMerchantService(opts...)
	r.Webhooks = NewWebhookService(opts...)
	r.Testing = NewTestingService(opts...)

	return
}
