// File generated from our OpenAPI spec by Stainless.

package bolt

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bolt/bolt-go/internal/apijson"
	"github.com/bolt/bolt-go/internal/param"
	"github.com/bolt/bolt-go/internal/requestconfig"
	"github.com/bolt/bolt-go/option"
)

// AccountAddressService contains methods and other services that help with
// interacting with the bolt API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccountAddressService] method instead.
type AccountAddressService struct {
	Options []option.RequestOption
}

// NewAccountAddressService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountAddressService(opts ...option.RequestOption) (r *AccountAddressService) {
	r = &AccountAddressService{}
	r.Options = opts
	return
}

// Add an address to the shopper's account
func (r *AccountAddressService) New(ctx context.Context, body AccountAddressNewParams, opts ...option.RequestOption) (res *AccountAddress, err error) {
	opts = append(r.Options[:], opts...)
	path := "account/addresses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Edit an existing address on the shopper's account. This does not edit addresses
// that are already associated with other resources, such as transactions or
// shipments.
func (r *AccountAddressService) Update(ctx context.Context, id string, body AccountAddressUpdateParams, opts ...option.RequestOption) (res *AccountAddress, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account/addresses/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete an existing address. Deleting an address does not invalidate transactions
// or shipments that are associated with it.
func (r *AccountAddressService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("account/addresses/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type AccountAddress struct {
	ID             string `json:"id,required"`
	CountryCode    string `json:"country_code,required"`
	FirstName      string `json:"first_name,required"`
	LastName       string `json:"last_name,required"`
	Locality       string `json:"locality,required"`
	PostalCode     string `json:"postal_code,required"`
	StreetAddress1 string `json:"street_address1,required"`
	Company        string `json:"company"`
	Email          string `json:"email" format:"email"`
	IsDefault      bool   `json:"is_default"`
	Phone          string `json:"phone" format:"phone"`
	Region         string `json:"region"`
	StreetAddress2 string `json:"street_address2"`
	JSON           accountAddressJSON
}

// accountAddressJSON contains the JSON metadata for the struct [AccountAddress]
type accountAddressJSON struct {
	ID             apijson.Field
	CountryCode    apijson.Field
	FirstName      apijson.Field
	LastName       apijson.Field
	Locality       apijson.Field
	PostalCode     apijson.Field
	StreetAddress1 apijson.Field
	Company        apijson.Field
	Email          apijson.Field
	IsDefault      apijson.Field
	Phone          apijson.Field
	Region         apijson.Field
	StreetAddress2 apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressNewParams struct {
	CountryCode    param.Field[string] `json:"country_code,required"`
	FirstName      param.Field[string] `json:"first_name,required"`
	LastName       param.Field[string] `json:"last_name,required"`
	Locality       param.Field[string] `json:"locality,required"`
	PostalCode     param.Field[string] `json:"postal_code,required"`
	StreetAddress1 param.Field[string] `json:"street_address1,required"`
	Company        param.Field[string] `json:"company"`
	Email          param.Field[string] `json:"email" format:"email"`
	IsDefault      param.Field[bool]   `json:"is_default"`
	Phone          param.Field[string] `json:"phone" format:"phone"`
	Region         param.Field[string] `json:"region"`
	StreetAddress2 param.Field[string] `json:"street_address2"`
}

func (r AccountAddressNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAddressUpdateParams struct {
	CountryCode    param.Field[string] `json:"country_code,required"`
	FirstName      param.Field[string] `json:"first_name,required"`
	LastName       param.Field[string] `json:"last_name,required"`
	Locality       param.Field[string] `json:"locality,required"`
	PostalCode     param.Field[string] `json:"postal_code,required"`
	StreetAddress1 param.Field[string] `json:"street_address1,required"`
	Company        param.Field[string] `json:"company"`
	Email          param.Field[string] `json:"email" format:"email"`
	IsDefault      param.Field[bool]   `json:"is_default"`
	Phone          param.Field[string] `json:"phone" format:"phone"`
	Region         param.Field[string] `json:"region"`
	StreetAddress2 param.Field[string] `json:"street_address2"`
}

func (r AccountAddressUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
