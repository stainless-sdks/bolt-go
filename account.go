// File generated from our OpenAPI spec by Stainless.

package bolt

import (
	"context"
	"net/http"
	"net/url"
	"reflect"

	"github.com/bolt/bolt-go/internal/apijson"
	"github.com/bolt/bolt-go/internal/apiquery"
	"github.com/bolt/bolt-go/internal/param"
	"github.com/bolt/bolt-go/internal/requestconfig"
	"github.com/bolt/bolt-go/option"
	"github.com/tidwall/gjson"
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
func (r *AccountService) Get(ctx context.Context, query AccountGetParams, opts ...option.RequestOption) (res *Account, err error) {
	opts = append(r.Options[:], opts...)
	path := "account"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Determine whether or not an identifier is associated with an existing Bolt
// account.
func (r *AccountService) Exists(ctx context.Context, params AccountExistsParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "account/exists"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, nil, opts...)
	return
}

type Account struct {
	Addresses      []Address       `json:"addresses,required"`
	PaymentMethods []PaymentMethod `json:"payment_methods,required"`
	Profile        Profile         `json:"profile"`
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

type Address struct {
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
	JSON           addressJSON
}

// addressJSON contains the JSON metadata for the struct [Address]
type addressJSON struct {
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

func (r *Address) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentMethod struct {
	Tag PaymentMethodTag `json:".tag,required"`
	ID  string           `json:"id"`
	// The ID of credit card's billing address
	BillingAddressID string `json:"billing_address_id"`
	JSON             paymentMethodJSON
	CreditCard
}

// paymentMethodJSON contains the JSON metadata for the struct [PaymentMethod]
type paymentMethodJSON struct {
	Tag              apijson.Field
	ID               apijson.Field
	BillingAddressID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PaymentMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentMethodTag string

const (
	PaymentMethodTagCreditCard PaymentMethodTag = "credit_card"
)

// The credit card's billing address
//
// Union satisfied by [PaymentMethodBillingAddressAddressReferenceID] or
// [PaymentMethodBillingAddressAddressReferenceExplicit].
type PaymentMethodBillingAddress interface {
	implementsPaymentMethodBillingAddress()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PaymentMethodBillingAddress)(nil)).Elem(),
		".tag",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"id\"",
			Type:               reflect.TypeOf(PaymentMethodBillingAddressAddressReferenceID{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"explicit\"",
			Type:               reflect.TypeOf(PaymentMethodBillingAddressAddressReferenceExplicit{}),
		},
	)
}

type PaymentMethodBillingAddressAddressReferenceID struct {
	// The address's ID
	ID string `json:"id,required"`
	// The type of address reference
	Tag  PaymentMethodBillingAddressAddressReferenceIDTag `json:".tag,required"`
	JSON paymentMethodBillingAddressAddressReferenceIDJSON
}

// paymentMethodBillingAddressAddressReferenceIDJSON contains the JSON metadata for
// the struct [PaymentMethodBillingAddressAddressReferenceID]
type paymentMethodBillingAddressAddressReferenceIDJSON struct {
	ID          apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentMethodBillingAddressAddressReferenceID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of address reference
type PaymentMethodBillingAddressAddressReferenceIDTag string

const (
	PaymentMethodBillingAddressAddressReferenceIDTagID PaymentMethodBillingAddressAddressReferenceIDTag = "id"
)

type PaymentMethodBillingAddressAddressReferenceExplicit struct {
	ID string `json:"id,required"`
	// The type of address reference
	Tag            PaymentMethodBillingAddressAddressReferenceExplicitTag `json:".tag,required"`
	CountryCode    string                                                 `json:"country_code,required"`
	FirstName      string                                                 `json:"first_name,required"`
	LastName       string                                                 `json:"last_name,required"`
	Locality       string                                                 `json:"locality,required"`
	PostalCode     string                                                 `json:"postal_code,required"`
	StreetAddress1 string                                                 `json:"street_address1,required"`
	Company        string                                                 `json:"company"`
	Email          string                                                 `json:"email" format:"email"`
	Phone          string                                                 `json:"phone" format:"phone"`
	Region         string                                                 `json:"region"`
	StreetAddress2 string                                                 `json:"street_address2"`
	JSON           paymentMethodBillingAddressAddressReferenceExplicitJSON
}

// paymentMethodBillingAddressAddressReferenceExplicitJSON contains the JSON
// metadata for the struct [PaymentMethodBillingAddressAddressReferenceExplicit]
type paymentMethodBillingAddressAddressReferenceExplicitJSON struct {
	ID             apijson.Field
	Tag            apijson.Field
	CountryCode    apijson.Field
	FirstName      apijson.Field
	LastName       apijson.Field
	Locality       apijson.Field
	PostalCode     apijson.Field
	StreetAddress1 apijson.Field
	Company        apijson.Field
	Email          apijson.Field
	Phone          apijson.Field
	Region         apijson.Field
	StreetAddress2 apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PaymentMethodBillingAddressAddressReferenceExplicit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of address reference
type PaymentMethodBillingAddressAddressReferenceExplicitTag string

const (
	PaymentMethodBillingAddressAddressReferenceExplicitTagExplicit PaymentMethodBillingAddressAddressReferenceExplicitTag = "explicit"
)

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

type AccountGetParams struct {
	XPublishableKey param.Field[string] `header:"X-Publishable-Key,required"`
}

type AccountExistsParams struct {
	// A type and value combination that defines the identifier used to detect the
	// existence of an account.
	Identifier      param.Field[AccountExistsParamsIdentifier] `query:"identifier,required"`
	XPublishableKey param.Field[string]                        `header:"X-Publishable-Key,required"`
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
