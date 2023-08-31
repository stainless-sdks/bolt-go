// File generated from our OpenAPI spec by Stainless.

package bolt

import (
	"context"
	"net/http"
	"net/url"

	"github.com/bolt/bolt-go/internal/apijson"
	"github.com/bolt/bolt-go/internal/apiquery"
	"github.com/bolt/bolt-go/internal/param"
	"github.com/bolt/bolt-go/internal/requestconfig"
	"github.com/bolt/bolt-go/option"
)

// AccountService contains methods and other services that help with interacting
// with the bolt API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewAccountService] method instead.
type AccountService struct {
	Options        []option.RequestOption
	Addresses      *AccountAddressService
	PaymentMethods *AccountPaymentMethodService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	r.Addresses = NewAccountAddressService(opts...)
	r.PaymentMethods = NewAccountPaymentMethodService(opts...)
	return
}

// Retrieve a shopper's account details, such as addresses and payment information
func (r *AccountService) Get(ctx context.Context, opts ...option.RequestOption) (res *Account, err error) {
	opts = append(r.Options[:], opts...)
	path := "account"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Determine whether or not an identifier is associated with an existing Bolt
// account.
func (r *AccountService) Exists(ctx context.Context, query AccountExistsParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "account/exists"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type Account struct {
	Addresses      []AccountAddress `json:"addresses,required"`
	PaymentMethods []PaymentMethod  `json:"payment_methods,required"`
	Profile        Profile          `json:"profile"`
	JSON           accountJSON
}

// accountJSON contains the JSON metadata for the struct [Account]
type accountJSON struct {
	Addresses      apijson.Field
	PaymentMethods apijson.Field
	Profile        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Account) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Profile struct {
	// An email address.
	Email string `json:"email"`
	// The given name of the person associated with this record.
	FirstName string `json:"first_name"`
	// The last name of the person associated with this record.
	LastName string `json:"last_name"`
	// A phone number following E164 standards, in its globalized format, i.e.
	// prepended with a plus sign.
	Phone string `json:"phone"`
	JSON  profileJSON
}

// profileJSON contains the JSON metadata for the struct [Profile]
type profileJSON struct {
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	Phone       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Profile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountExistsParams struct {
	// A type and value combination that defines the identifier used to detect the
	// existence of an account.
	Identifier param.Field[AccountExistsParamsIdentifier] `query:"identifier,required"`
}

// URLQuery serializes [AccountExistsParams]'s query parameters as `url.Values`.
func (r AccountExistsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A type and value combination that defines the identifier used to detect the
// existence of an account.
type AccountExistsParamsIdentifier struct {
	// The type of identifier
	IdentifierType param.Field[AccountExistsParamsIdentifierIdentifierType] `query:"identifier_type,required"`
	// The value of the identifier. The value must be valid for the specified
	// `identifier_type`
	IdentifierValue param.Field[string] `query:"identifier_value,required"`
}

// URLQuery serializes [AccountExistsParamsIdentifier]'s query parameters as
// `url.Values`.
func (r AccountExistsParamsIdentifier) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of identifier
type AccountExistsParamsIdentifierIdentifierType string

const (
	AccountExistsParamsIdentifierIdentifierTypeEmail       AccountExistsParamsIdentifierIdentifierType = "email"
	AccountExistsParamsIdentifierIdentifierTypeEmailSha256 AccountExistsParamsIdentifierIdentifierType = "email_sha256"
)
