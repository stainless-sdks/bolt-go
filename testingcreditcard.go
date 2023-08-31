// File generated from our OpenAPI spec by Stainless.

package bolt

import (
	"context"
	"net/http"

	"github.com/bolt/bolt-go/internal/requestconfig"
	"github.com/bolt/bolt-go/option"
)

// TestingCreditCardService contains methods and other services that help with
// interacting with the bolt API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewTestingCreditCardService] method
// instead.
type TestingCreditCardService struct {
	Options []option.RequestOption
}

// NewTestingCreditCardService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTestingCreditCardService(opts ...option.RequestOption) (r *TestingCreditCardService) {
	r = &TestingCreditCardService{}
	r.Options = opts
	return
}

// Retrieve test credit card information. This includes its token, which is
// generated against the `4111 1111 1111 1004` test card.
func (r *TestingCreditCardService) Get(ctx context.Context, opts ...option.RequestOption) (res *CreditCard, err error) {
	opts = append(r.Options[:], opts...)
	path := "testing/credit-cards"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
