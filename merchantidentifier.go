// File generated from our OpenAPI spec by Stainless.

package bolt

import (
	"context"
	"net/http"

	"github.com/bolt/bolt-go/internal/requestconfig"
	"github.com/bolt/bolt-go/option"
)

// MerchantIdentifierService contains methods and other services that help with
// interacting with the bolt API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewMerchantIdentifierService] method
// instead.
type MerchantIdentifierService struct {
	Options []option.RequestOption
}

// NewMerchantIdentifierService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMerchantIdentifierService(opts ...option.RequestOption) (r *MerchantIdentifierService) {
	r = &MerchantIdentifierService{}
	r.Options = opts
	return
}

// Return several identifiers for the merchant, such as public IDs, publishable
// keys, signing secrets, etc...
func (r *MerchantIdentifierService) List(ctx context.Context, opts ...option.RequestOption) (res *Identifiers, err error) {
	opts = append(r.Options[:], opts...)
	path := "merchant/identifiers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
