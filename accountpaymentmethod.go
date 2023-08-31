// File generated from our OpenAPI spec by Stainless.

package bolt

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/bolt/bolt-go/internal/apijson"
	"github.com/bolt/bolt-go/internal/param"
	"github.com/bolt/bolt-go/internal/requestconfig"
	"github.com/bolt/bolt-go/option"
	"github.com/tidwall/gjson"
)

// AccountPaymentMethodService contains methods and other services that help with
// interacting with the bolt API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccountPaymentMethodService] method
// instead.
type AccountPaymentMethodService struct {
	Options []option.RequestOption
}

// NewAccountPaymentMethodService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountPaymentMethodService(opts ...option.RequestOption) (r *AccountPaymentMethodService) {
	r = &AccountPaymentMethodService{}
	r.Options = opts
	return
}

// Add a payment method to a shopper's Bolt account Wallet. For security purposes,
// this request must come from your backend because authentication requires the use
// of your private key.<br /> **Note**: Before using this API, the credit card
// details must be tokenized using Bolt's JavaScript library function, which is
// documented in
// [Install the Bolt Tokenizer](https://help.bolt.com/developers/references/bolt-tokenizer).
func (r *AccountPaymentMethodService) New(ctx context.Context, body AccountPaymentMethodNewParams, opts ...option.RequestOption) (res *PaymentMethod, err error) {
	opts = append(r.Options[:], opts...)
	path := "account/payment-methods"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete an existing payment method. Deleting a payment method does not invalidate
// transactions or orders that are associated with it.
func (r *AccountPaymentMethodService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("account/payment-methods/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
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

func (r PaymentMethodBillingAddressAddressReferenceID) implementsPaymentMethodBillingAddress() {}

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

func (r PaymentMethodBillingAddressAddressReferenceExplicit) implementsPaymentMethodBillingAddress() {
}

// The type of address reference
type PaymentMethodBillingAddressAddressReferenceExplicitTag string

const (
	PaymentMethodBillingAddressAddressReferenceExplicitTagExplicit PaymentMethodBillingAddressAddressReferenceExplicitTag = "explicit"
)

type AccountPaymentMethodNewParams struct {
	// The Bolt token associated to the credit card.
	Token param.Field[string]                           `json:"token,required"`
	Tag   param.Field[AccountPaymentMethodNewParamsTag] `json:".tag,required"`
	// The credit card's billing address
	BillingAddress param.Field[AccountPaymentMethodNewParamsBillingAddress] `json:"billing_address,required"`
	// The Bank Identification Number for the credit card. This is typically the first
	// 4-6 digits of the credit card number.
	Bin param.Field[string] `json:"bin,required" format:"^\\d+$"`
	// The expiration date of the credit card. TODO TO MAKE EXPIRATION REUSABLE
	Expiration param.Field[string] `json:"expiration,required" format:"^\\d{4}-\\d{2}$"`
	// The last 4 digits of the credit card number.
	Last4 param.Field[string] `json:"last4,required" format:"^\\d{4}$"`
	// The credit card network.
	Network param.Field[AccountPaymentMethodNewParamsNetwork] `json:"network,required"`
}

func (r AccountPaymentMethodNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountPaymentMethodNewParamsTag string

const (
	AccountPaymentMethodNewParamsTagCreditCard AccountPaymentMethodNewParamsTag = "credit_card"
)

// The credit card's billing address
//
// Satisfied by [AccountPaymentMethodNewParamsBillingAddressAddressReferenceID],
// [AccountPaymentMethodNewParamsBillingAddressAddressReferenceExplicit].
type AccountPaymentMethodNewParamsBillingAddress interface {
	implementsAccountPaymentMethodNewParamsBillingAddress()
}

type AccountPaymentMethodNewParamsBillingAddressAddressReferenceID struct {
	// The address's ID
	ID param.Field[string] `json:"id,required"`
	// The type of address reference
	Tag param.Field[AccountPaymentMethodNewParamsBillingAddressAddressReferenceIDTag] `json:".tag,required"`
}

func (r AccountPaymentMethodNewParamsBillingAddressAddressReferenceID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountPaymentMethodNewParamsBillingAddressAddressReferenceID) implementsAccountPaymentMethodNewParamsBillingAddress() {
}

// The type of address reference
type AccountPaymentMethodNewParamsBillingAddressAddressReferenceIDTag string

const (
	AccountPaymentMethodNewParamsBillingAddressAddressReferenceIDTagID AccountPaymentMethodNewParamsBillingAddressAddressReferenceIDTag = "id"
)

type AccountPaymentMethodNewParamsBillingAddressAddressReferenceExplicit struct {
	// The type of address reference
	Tag            param.Field[AccountPaymentMethodNewParamsBillingAddressAddressReferenceExplicitTag] `json:".tag,required"`
	CountryCode    param.Field[string]                                                                 `json:"country_code,required"`
	FirstName      param.Field[string]                                                                 `json:"first_name,required"`
	LastName       param.Field[string]                                                                 `json:"last_name,required"`
	Locality       param.Field[string]                                                                 `json:"locality,required"`
	PostalCode     param.Field[string]                                                                 `json:"postal_code,required"`
	StreetAddress1 param.Field[string]                                                                 `json:"street_address1,required"`
	Company        param.Field[string]                                                                 `json:"company"`
	Email          param.Field[string]                                                                 `json:"email" format:"email"`
	Phone          param.Field[string]                                                                 `json:"phone" format:"phone"`
	Region         param.Field[string]                                                                 `json:"region"`
	StreetAddress2 param.Field[string]                                                                 `json:"street_address2"`
}

func (r AccountPaymentMethodNewParamsBillingAddressAddressReferenceExplicit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountPaymentMethodNewParamsBillingAddressAddressReferenceExplicit) implementsAccountPaymentMethodNewParamsBillingAddress() {
}

// The type of address reference
type AccountPaymentMethodNewParamsBillingAddressAddressReferenceExplicitTag string

const (
	AccountPaymentMethodNewParamsBillingAddressAddressReferenceExplicitTagExplicit AccountPaymentMethodNewParamsBillingAddressAddressReferenceExplicitTag = "explicit"
)

// The credit card network.
type AccountPaymentMethodNewParamsNetwork string

const (
	AccountPaymentMethodNewParamsNetworkVisa         AccountPaymentMethodNewParamsNetwork = "visa"
	AccountPaymentMethodNewParamsNetworkMastercard   AccountPaymentMethodNewParamsNetwork = "mastercard"
	AccountPaymentMethodNewParamsNetworkAmex         AccountPaymentMethodNewParamsNetwork = "amex"
	AccountPaymentMethodNewParamsNetworkDiscover     AccountPaymentMethodNewParamsNetwork = "discover"
	AccountPaymentMethodNewParamsNetworkJcb          AccountPaymentMethodNewParamsNetwork = "jcb"
	AccountPaymentMethodNewParamsNetworkUnionpay     AccountPaymentMethodNewParamsNetwork = "unionpay"
	AccountPaymentMethodNewParamsNetworkAlliancedata AccountPaymentMethodNewParamsNetwork = "alliancedata"
	AccountPaymentMethodNewParamsNetworkCitiplcc     AccountPaymentMethodNewParamsNetwork = "citiplcc"
)
