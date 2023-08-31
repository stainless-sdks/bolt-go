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

type CallbackURLs struct {
	AccountPage                   string `json:"account_page" format:"uri"`
	BaseDomain                    string `json:"base_domain" format:"uri"`
	ConfirmationRedirect          string `json:"confirmation_redirect" format:"uri"`
	CreateOrder                   string `json:"create_order" format:"uri"`
	Debug                         string `json:"debug" format:"uri"`
	GetAccount                    string `json:"get_account" format:"uri"`
	MobileAppDomain               string `json:"mobile_app_domain" format:"uri"`
	OauthLogout                   string `json:"oauth_logout" format:"uri"`
	OauthRedirect                 string `json:"oauth_redirect" format:"uri"`
	PrivacyPolicy                 string `json:"privacy_policy" format:"uri"`
	ProductInfo                   string `json:"product_info" format:"uri"`
	RemoteAPI                     string `json:"remote_api" format:"uri"`
	Shipping                      string `json:"shipping" format:"uri"`
	SupportPage                   string `json:"support_page" format:"uri"`
	Tax                           string `json:"tax" format:"uri"`
	TermsOfService                string `json:"terms_of_service" format:"uri"`
	UniversalMerchantAPI          string `json:"universal_merchant_api" format:"uri"`
	UpdateCart                    string `json:"update_cart" format:"uri"`
	ValidateAdditionalAccountData string `json:"validate_additional_account_data" format:"uri"`
	JSON                          callbackURLsJSON
}

// callbackURLsJSON contains the JSON metadata for the struct [CallbackURLs]
type callbackURLsJSON struct {
	AccountPage                   apijson.Field
	BaseDomain                    apijson.Field
	ConfirmationRedirect          apijson.Field
	CreateOrder                   apijson.Field
	Debug                         apijson.Field
	GetAccount                    apijson.Field
	MobileAppDomain               apijson.Field
	OauthLogout                   apijson.Field
	OauthRedirect                 apijson.Field
	PrivacyPolicy                 apijson.Field
	ProductInfo                   apijson.Field
	RemoteAPI                     apijson.Field
	Shipping                      apijson.Field
	SupportPage                   apijson.Field
	Tax                           apijson.Field
	TermsOfService                apijson.Field
	UniversalMerchantAPI          apijson.Field
	UpdateCart                    apijson.Field
	ValidateAdditionalAccountData apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *CallbackURLs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
