// File generated from our OpenAPI spec by Stainless.

package bolt

import (
	"github.com/bolt/bolt-go/option"
)

// GuestService contains methods and other services that help with interacting with
// the bolt API. Note, unlike clients, this service does not read variables from
// the environment automatically. You should not instantiate this service directly,
// and instead use the [NewGuestService] method instead.
type GuestService struct {
	Options  []option.RequestOption
	Payments *GuestPaymentService
}

// NewGuestService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGuestService(opts ...option.RequestOption) (r *GuestService) {
	r = &GuestService{}
	r.Options = opts
	r.Payments = NewGuestPaymentService(opts...)
	return
}
