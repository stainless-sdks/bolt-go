// File generated from our OpenAPI spec by Stainless.

package bolt

import (
	"github.com/bolt/bolt-go/internal/apijson"
	"github.com/bolt/bolt-go/option"
)

// MerchantService contains methods and other services that help with interacting
// with the bolt API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewMerchantService] method instead.
type MerchantService struct {
	Options     []option.RequestOption
	Callbacks   *MerchantCallbackService
	Identifiers *MerchantIdentifierService
}

// NewMerchantService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMerchantService(opts ...option.RequestOption) (r *MerchantService) {
	r = &MerchantService{}
	r.Options = opts
	r.Callbacks = NewMerchantCallbackService(opts...)
	r.Identifiers = NewMerchantIdentifierService(opts...)
	return
}

type Identifiers struct {
	MerchantDivisions []IdentifiersMerchantDivision `json:"merchant_divisions,required"`
	MerchantID        string                        `json:"merchant_id,required"`
	SigningSecret     string                        `json:"signing_secret,required"`
	JSON              identifiersJSON
}

// identifiersJSON contains the JSON metadata for the struct [Identifiers]
type identifiersJSON struct {
	MerchantDivisions apijson.Field
	MerchantID        apijson.Field
	SigningSecret     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Identifiers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentifiersMerchantDivision struct {
	PublishableKey string `json:"publishable_key,required"`
	JSON           identifiersMerchantDivisionJSON
}

// identifiersMerchantDivisionJSON contains the JSON metadata for the struct
// [IdentifiersMerchantDivision]
type identifiersMerchantDivisionJSON struct {
	PublishableKey apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IdentifiersMerchantDivision) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
