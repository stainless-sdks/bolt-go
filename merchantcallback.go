// File generated from our OpenAPI spec by Stainless.

package bolt

import (
	"context"
	"net/http"

	"github.com/bolt/bolt-go/internal/apijson"
	"github.com/bolt/bolt-go/internal/param"
	"github.com/bolt/bolt-go/internal/requestconfig"
	"github.com/bolt/bolt-go/option"
)

// MerchantCallbackService contains methods and other services that help with
// interacting with the bolt API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewMerchantCallbackService] method
// instead.
type MerchantCallbackService struct {
	Options []option.RequestOption
}

// NewMerchantCallbackService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMerchantCallbackService(opts ...option.RequestOption) (r *MerchantCallbackService) {
	r = &MerchantCallbackService{}
	r.Options = opts
	return
}

// Update and configure callback URLs on the merchant such as OAuth URLs.
func (r *MerchantCallbackService) Update(ctx context.Context, params MerchantCallbackUpdateParams, opts ...option.RequestOption) (res *CallbackURLs, err error) {
	opts = append(r.Options[:], opts...)
	path := "merchant/callbacks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Return callback URLs configured on the merchant such as OAuth URLs.
func (r *MerchantCallbackService) List(ctx context.Context, query MerchantCallbackListParams, opts ...option.RequestOption) (res *CallbackURLs, err error) {
	opts = append(r.Options[:], opts...)
	path := "merchant/callbacks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type MerchantCallbackUpdateParams struct {
	XPublishableKey               param.Field[string] `header:"X-Publishable-Key,required"`
	AccountPage                   param.Field[string] `json:"account_page" format:"uri"`
	BaseDomain                    param.Field[string] `json:"base_domain" format:"uri"`
	ConfirmationRedirect          param.Field[string] `json:"confirmation_redirect" format:"uri"`
	CreateOrder                   param.Field[string] `json:"create_order" format:"uri"`
	Debug                         param.Field[string] `json:"debug" format:"uri"`
	GetAccount                    param.Field[string] `json:"get_account" format:"uri"`
	MobileAppDomain               param.Field[string] `json:"mobile_app_domain" format:"uri"`
	OauthLogout                   param.Field[string] `json:"oauth_logout" format:"uri"`
	OauthRedirect                 param.Field[string] `json:"oauth_redirect" format:"uri"`
	PrivacyPolicy                 param.Field[string] `json:"privacy_policy" format:"uri"`
	ProductInfo                   param.Field[string] `json:"product_info" format:"uri"`
	RemoteAPI                     param.Field[string] `json:"remote_api" format:"uri"`
	Shipping                      param.Field[string] `json:"shipping" format:"uri"`
	SupportPage                   param.Field[string] `json:"support_page" format:"uri"`
	Tax                           param.Field[string] `json:"tax" format:"uri"`
	TermsOfService                param.Field[string] `json:"terms_of_service" format:"uri"`
	UniversalMerchantAPI          param.Field[string] `json:"universal_merchant_api" format:"uri"`
	UpdateCart                    param.Field[string] `json:"update_cart" format:"uri"`
	ValidateAdditionalAccountData param.Field[string] `json:"validate_additional_account_data" format:"uri"`
}

func (r MerchantCallbackUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MerchantCallbackListParams struct {
	XPublishableKey param.Field[string] `header:"X-Publishable-Key,required"`
}
