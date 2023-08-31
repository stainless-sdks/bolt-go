// File generated from our OpenAPI spec by Stainless.

package bolt

import (
	"time"

	"github.com/bolt/bolt-go/internal/apijson"
	"github.com/bolt/bolt-go/option"
)

// TestingService contains methods and other services that help with interacting
// with the bolt API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewTestingService] method instead.
type TestingService struct {
	Options     []option.RequestOption
	Accounts    *TestingAccountService
	CreditCards *TestingCreditCardService
	Shipments   *TestingShipmentService
}

// NewTestingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTestingService(opts ...option.RequestOption) (r *TestingService) {
	r = &TestingService{}
	r.Options = opts
	r.Accounts = NewTestingAccountService(opts...)
	r.CreditCards = NewTestingCreditCardService(opts...)
	r.Shipments = NewTestingShipmentService(opts...)
	return
}

type AccountTest struct {
	DeactivateAt time.Time             `json:"deactivate_at,required" format:"date-time"`
	Email        string                `json:"email,required"`
	EmailState   AccountTestEmailState `json:"email_state,required"`
	OauthCode    string                `json:"oauth_code,required"`
	OtpCode      string                `json:"otp_code,required"`
	Phone        string                `json:"phone,required"`
	PhoneState   AccountTestPhoneState `json:"phone_state,required"`
	JSON         accountTestJSON
}

// accountTestJSON contains the JSON metadata for the struct [AccountTest]
type accountTestJSON struct {
	DeactivateAt apijson.Field
	Email        apijson.Field
	EmailState   apijson.Field
	OauthCode    apijson.Field
	OtpCode      apijson.Field
	Phone        apijson.Field
	PhoneState   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTestEmailState string

const (
	AccountTestEmailStateMissing    AccountTestEmailState = "missing"
	AccountTestEmailStateUnverified AccountTestEmailState = "unverified"
	AccountTestEmailStateVerified   AccountTestEmailState = "verified"
)

type AccountTestPhoneState string

const (
	AccountTestPhoneStateMissing    AccountTestPhoneState = "missing"
	AccountTestPhoneStateUnverified AccountTestPhoneState = "unverified"
	AccountTestPhoneStateVerified   AccountTestPhoneState = "verified"
)

type CreditCard struct {
	// The expiration date of the credit card. TODO TO MAKE EXPIRATION REUSABLE
	Expiration string `json:"expiration,required" format:"^\\d{4}-\\d{2}$"`
	// The last 4 digits of the credit card number.
	Last4 string `json:"last4,required" format:"^\\d{4}$"`
	// The credit card network.
	Network CreditCardNetwork `json:"network,required"`
	JSON    creditCardJSON
}

// creditCardJSON contains the JSON metadata for the struct [CreditCard]
type creditCardJSON struct {
	Expiration  apijson.Field
	Last4       apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditCard) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The credit card network.
type CreditCardNetwork string

const (
	CreditCardNetworkVisa         CreditCardNetwork = "visa"
	CreditCardNetworkMastercard   CreditCardNetwork = "mastercard"
	CreditCardNetworkAmex         CreditCardNetwork = "amex"
	CreditCardNetworkDiscover     CreditCardNetwork = "discover"
	CreditCardNetworkJcb          CreditCardNetwork = "jcb"
	CreditCardNetworkUnionpay     CreditCardNetwork = "unionpay"
	CreditCardNetworkAlliancedata CreditCardNetwork = "alliancedata"
	CreditCardNetworkCitiplcc     CreditCardNetwork = "citiplcc"
)
