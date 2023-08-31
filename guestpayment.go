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

// GuestPaymentService contains methods and other services that help with
// interacting with the bolt API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewGuestPaymentService] method instead.
type GuestPaymentService struct {
	Options []option.RequestOption
}

// NewGuestPaymentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGuestPaymentService(opts ...option.RequestOption) (r *GuestPaymentService) {
	r = &GuestPaymentService{}
	r.Options = opts
	return
}

// Initialize a Bolt payment token that will be used to reference this payment to
// Bolt when it is updated or finalized for guest shoppers.
func (r *GuestPaymentService) New(ctx context.Context, params GuestPaymentNewParams, opts ...option.RequestOption) (res *GuestPaymentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "guest/payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type GuestPaymentNewResponse struct {
	ID     string                        `json:"id"`
	Action GuestPaymentNewResponseAction `json:"action"`
	Status GuestPaymentNewResponseStatus `json:"status"`
	JSON   guestPaymentNewResponseJSON
}

// guestPaymentNewResponseJSON contains the JSON metadata for the struct
// [GuestPaymentNewResponse]
type guestPaymentNewResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GuestPaymentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GuestPaymentNewResponseAction struct {
	Method GuestPaymentNewResponseActionMethod `json:"method"`
	Type   GuestPaymentNewResponseActionType   `json:"type"`
	URL    string                              `json:"url" format:"uri"`
	JSON   guestPaymentNewResponseActionJSON
}

// guestPaymentNewResponseActionJSON contains the JSON metadata for the struct
// [GuestPaymentNewResponseAction]
type guestPaymentNewResponseActionJSON struct {
	Method      apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GuestPaymentNewResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GuestPaymentNewResponseActionMethod string

const (
	GuestPaymentNewResponseActionMethodGet  GuestPaymentNewResponseActionMethod = "GET"
	GuestPaymentNewResponseActionMethodPost GuestPaymentNewResponseActionMethod = "POST"
)

type GuestPaymentNewResponseActionType string

const (
	GuestPaymentNewResponseActionTypeRedirect GuestPaymentNewResponseActionType = "redirect"
	GuestPaymentNewResponseActionTypeFinalize GuestPaymentNewResponseActionType = "finalize"
)

type GuestPaymentNewResponseStatus string

const (
	GuestPaymentNewResponseStatusAwaitingUserConfirmation GuestPaymentNewResponseStatus = "awaiting_user_confirmation"
	GuestPaymentNewResponseStatusPaymentReady             GuestPaymentNewResponseStatus = "payment_ready"
	GuestPaymentNewResponseStatusUpdatePaymentMethod      GuestPaymentNewResponseStatus = "update_payment_method"
	GuestPaymentNewResponseStatusSuccess                  GuestPaymentNewResponseStatus = "success"
)

type GuestPaymentNewParams struct {
	Cart            param.Field[GuestPaymentNewParamsCart]          `json:"cart,required"`
	PaymentMethod   param.Field[GuestPaymentNewParamsPaymentMethod] `json:"payment_method,required"`
	XPublishableKey param.Field[string]                             `header:"X-Publishable-Key,required"`
}

func (r GuestPaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GuestPaymentNewParamsCart struct {
	Amounts param.Field[GuestPaymentNewParamsCartAmounts] `json:"amounts,required"`
	// This value is used by Bolt as an external reference to a given order. This
	// reference must be unique per successful transaction.
	OrderReference param.Field[string]                              `json:"order_reference,required"`
	Discounts      param.Field[[]GuestPaymentNewParamsCartDiscount] `json:"discounts"`
	// This field corresponds to the merchant's order reference associated with this
	// Bolt transaction.
	DisplayID param.Field[string]                          `json:"display_id"`
	Items     param.Field[[]GuestPaymentNewParamsCartItem] `json:"items"`
	// Used optionally to pass additional information like order numbers or other IDs
	// as needed.
	OrderDescription param.Field[string]                              `json:"order_description"`
	Shipments        param.Field[[]GuestPaymentNewParamsCartShipment] `json:"shipments"`
}

func (r GuestPaymentNewParamsCart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GuestPaymentNewParamsCartAmounts struct {
	Currency param.Field[string] `json:"currency,required"`
	// The total amount, in cents, including its items and taxes (if applicable).
	Total param.Field[int64] `json:"total,required"`
	// The total tax amount, in cents for all of the items associated with the cart.
	Tax param.Field[int64] `json:"tax"`
}

func (r GuestPaymentNewParamsCartAmounts) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GuestPaymentNewParamsCartDiscount struct {
	Amounts param.Field[GuestPaymentNewParamsCartDiscountsAmounts] `json:"amounts,required"`
	// Discount code.
	Code param.Field[string] `json:"code"`
	// Used to provide a link to additional details, such as a landing page, associated
	// with the discount offering.
	DetailsURL param.Field[string] `json:"details_url" format:"uri"`
}

func (r GuestPaymentNewParamsCartDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GuestPaymentNewParamsCartDiscountsAmounts struct {
	Currency param.Field[string] `json:"currency,required"`
	// The total amount, in cents, including its items and taxes (if applicable).
	Total param.Field[int64] `json:"total,required"`
	// The total tax amount, in cents for all of the items associated with the cart.
	Tax param.Field[int64] `json:"tax"`
}

func (r GuestPaymentNewParamsCartDiscountsAmounts) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GuestPaymentNewParamsCartItem struct {
	// The name of a given item.
	Name param.Field[string] `json:"name,required"`
	// The number of units that comprise this cart item.
	Quantity param.Field[int64] `json:"quantity,required"`
	// This value is used by Bolt as an external reference to a given item.
	Reference param.Field[string] `json:"reference,required"`
	// The total amount, in cents, of the item including its taxes if applicable.
	TotalAmount param.Field[int64] `json:"total_amount,required"`
	// The price of one unit of the item; for example, the price of one pack of socks.
	UnitPrice param.Field[int64] `json:"unit_price,required"`
	// A human-readable description of this cart item.
	Description param.Field[string] `json:"description"`
	// Used to provide a link to the image associated with the item.
	ImageURL param.Field[string] `json:"image_url" format:"uri"`
}

func (r GuestPaymentNewParamsCartItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GuestPaymentNewParamsCartShipment struct {
	// The Address object is used for shipping, and physical store address use cases.
	Address param.Field[GuestPaymentNewParamsCartShipmentsAddress] `json:"address"`
	// The name of the carrier selected.
	Carrier param.Field[string]                                 `json:"carrier"`
	Cost    param.Field[GuestPaymentNewParamsCartShipmentsCost] `json:"cost"`
}

func (r GuestPaymentNewParamsCartShipment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The Address object is used for shipping, and physical store address use cases.
//
// Satisfied by [GuestPaymentNewParamsCartShipmentsAddressAddressReferenceID],
// [GuestPaymentNewParamsCartShipmentsAddressAddressReferenceExplicit].
type GuestPaymentNewParamsCartShipmentsAddress interface {
	implementsGuestPaymentNewParamsCartShipmentsAddress()
}

type GuestPaymentNewParamsCartShipmentsAddressAddressReferenceID struct {
	// The address's ID
	ID param.Field[string] `json:"id,required"`
	// The type of address reference
	Tag param.Field[GuestPaymentNewParamsCartShipmentsAddressAddressReferenceIDTag] `json:".tag,required"`
}

func (r GuestPaymentNewParamsCartShipmentsAddressAddressReferenceID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GuestPaymentNewParamsCartShipmentsAddressAddressReferenceID) implementsGuestPaymentNewParamsCartShipmentsAddress() {
}

// The type of address reference
type GuestPaymentNewParamsCartShipmentsAddressAddressReferenceIDTag string

const (
	GuestPaymentNewParamsCartShipmentsAddressAddressReferenceIDTagID GuestPaymentNewParamsCartShipmentsAddressAddressReferenceIDTag = "id"
)

type GuestPaymentNewParamsCartShipmentsAddressAddressReferenceExplicit struct {
	// The type of address reference
	Tag            param.Field[GuestPaymentNewParamsCartShipmentsAddressAddressReferenceExplicitTag] `json:".tag,required"`
	CountryCode    param.Field[string]                                                               `json:"country_code,required"`
	FirstName      param.Field[string]                                                               `json:"first_name,required"`
	LastName       param.Field[string]                                                               `json:"last_name,required"`
	Locality       param.Field[string]                                                               `json:"locality,required"`
	PostalCode     param.Field[string]                                                               `json:"postal_code,required"`
	StreetAddress1 param.Field[string]                                                               `json:"street_address1,required"`
	Company        param.Field[string]                                                               `json:"company"`
	Email          param.Field[string]                                                               `json:"email" format:"email"`
	Phone          param.Field[string]                                                               `json:"phone" format:"phone"`
	Region         param.Field[string]                                                               `json:"region"`
	StreetAddress2 param.Field[string]                                                               `json:"street_address2"`
}

func (r GuestPaymentNewParamsCartShipmentsAddressAddressReferenceExplicit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GuestPaymentNewParamsCartShipmentsAddressAddressReferenceExplicit) implementsGuestPaymentNewParamsCartShipmentsAddress() {
}

// The type of address reference
type GuestPaymentNewParamsCartShipmentsAddressAddressReferenceExplicitTag string

const (
	GuestPaymentNewParamsCartShipmentsAddressAddressReferenceExplicitTagExplicit GuestPaymentNewParamsCartShipmentsAddressAddressReferenceExplicitTag = "explicit"
)

type GuestPaymentNewParamsCartShipmentsCost struct {
	Currency param.Field[string] `json:"currency,required"`
	// The total amount, in cents, including its items and taxes (if applicable).
	Total param.Field[int64] `json:"total,required"`
	// The total tax amount, in cents for all of the items associated with the cart.
	Tax param.Field[int64] `json:"tax"`
}

func (r GuestPaymentNewParamsCartShipmentsCost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GuestPaymentNewParamsPaymentMethod struct {
	Tag param.Field[GuestPaymentNewParamsPaymentMethodTag] `json:".tag,required"`
	// Redirect URL for canceled PayPal transaction.
	Cancel param.Field[string] `json:"cancel,required" format:"uri"`
	// Redirect URL for successful PayPal transaction.
	Success param.Field[string] `json:"success,required" format:"uri"`
}

func (r GuestPaymentNewParamsPaymentMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GuestPaymentNewParamsPaymentMethodTag string

const (
	GuestPaymentNewParamsPaymentMethodTagPaypal GuestPaymentNewParamsPaymentMethodTag = "paypal"
)
