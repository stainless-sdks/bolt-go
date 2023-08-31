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

// PaymentService contains methods and other services that help with interacting
// with the bolt API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewPaymentService] method instead.
type PaymentService struct {
	Options []option.RequestOption
}

// NewPaymentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPaymentService(opts ...option.RequestOption) (r *PaymentService) {
	r = &PaymentService{}
	r.Options = opts
	return
}

// Initialize a Bolt payment token that will be used to reference this payment to
// Bolt when it is updated or finalized for logged in shoppers.
func (r *PaymentService) New(ctx context.Context, body PaymentNewParams, opts ...option.RequestOption) (res *PaymentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PaymentNewResponse struct {
	ID     string                   `json:"id"`
	Action PaymentNewResponseAction `json:"action"`
	Status PaymentNewResponseStatus `json:"status"`
	JSON   paymentNewResponseJSON
}

// paymentNewResponseJSON contains the JSON metadata for the struct
// [PaymentNewResponse]
type paymentNewResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentNewResponseAction struct {
	Method PaymentNewResponseActionMethod `json:"method"`
	Type   PaymentNewResponseActionType   `json:"type"`
	URL    string                         `json:"url" format:"uri"`
	JSON   paymentNewResponseActionJSON
}

// paymentNewResponseActionJSON contains the JSON metadata for the struct
// [PaymentNewResponseAction]
type paymentNewResponseActionJSON struct {
	Method      apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentNewResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentNewResponseActionMethod string

const (
	PaymentNewResponseActionMethodGet  PaymentNewResponseActionMethod = "GET"
	PaymentNewResponseActionMethodPost PaymentNewResponseActionMethod = "POST"
)

type PaymentNewResponseActionType string

const (
	PaymentNewResponseActionTypeRedirect PaymentNewResponseActionType = "redirect"
	PaymentNewResponseActionTypeFinalize PaymentNewResponseActionType = "finalize"
)

type PaymentNewResponseStatus string

const (
	PaymentNewResponseStatusAwaitingUserConfirmation PaymentNewResponseStatus = "awaiting_user_confirmation"
	PaymentNewResponseStatusPaymentReady             PaymentNewResponseStatus = "payment_ready"
	PaymentNewResponseStatusUpdatePaymentMethod      PaymentNewResponseStatus = "update_payment_method"
	PaymentNewResponseStatusSuccess                  PaymentNewResponseStatus = "success"
)

type PaymentNewParams struct {
	Cart          param.Field[PaymentNewParamsCart]          `json:"cart,required"`
	PaymentMethod param.Field[PaymentNewParamsPaymentMethod] `json:"payment_method,required"`
}

func (r PaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsCart struct {
	Amounts param.Field[PaymentNewParamsCartAmounts] `json:"amounts,required"`
	// This value is used by Bolt as an external reference to a given order. This
	// reference must be unique per successful transaction.
	OrderReference param.Field[string]                         `json:"order_reference,required"`
	Discounts      param.Field[[]PaymentNewParamsCartDiscount] `json:"discounts"`
	// This field corresponds to the merchant's order reference associated with this
	// Bolt transaction.
	DisplayID param.Field[string]                     `json:"display_id"`
	Items     param.Field[[]PaymentNewParamsCartItem] `json:"items"`
	// Used optionally to pass additional information like order numbers or other IDs
	// as needed.
	OrderDescription param.Field[string]                         `json:"order_description"`
	Shipments        param.Field[[]PaymentNewParamsCartShipment] `json:"shipments"`
}

func (r PaymentNewParamsCart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsCartAmounts struct {
	Currency param.Field[string] `json:"currency,required"`
	// The total amount, in cents, including its items and taxes (if applicable).
	Total param.Field[int64] `json:"total,required"`
	// The total tax amount, in cents for all of the items associated with the cart.
	Tax param.Field[int64] `json:"tax"`
}

func (r PaymentNewParamsCartAmounts) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsCartDiscount struct {
	Amounts param.Field[PaymentNewParamsCartDiscountsAmounts] `json:"amounts,required"`
	// Discount code.
	Code param.Field[string] `json:"code"`
	// Used to provide a link to additional details, such as a landing page, associated
	// with the discount offering.
	DetailsURL param.Field[string] `json:"details_url" format:"uri"`
}

func (r PaymentNewParamsCartDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsCartDiscountsAmounts struct {
	Currency param.Field[string] `json:"currency,required"`
	// The total amount, in cents, including its items and taxes (if applicable).
	Total param.Field[int64] `json:"total,required"`
	// The total tax amount, in cents for all of the items associated with the cart.
	Tax param.Field[int64] `json:"tax"`
}

func (r PaymentNewParamsCartDiscountsAmounts) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsCartItem struct {
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

func (r PaymentNewParamsCartItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsCartShipment struct {
	// The Address object is used for shipping, and physical store address use cases.
	Address param.Field[PaymentNewParamsCartShipmentsAddress] `json:"address"`
	// The name of the carrier selected.
	Carrier param.Field[string]                            `json:"carrier"`
	Cost    param.Field[PaymentNewParamsCartShipmentsCost] `json:"cost"`
}

func (r PaymentNewParamsCartShipment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The Address object is used for shipping, and physical store address use cases.
//
// Satisfied by [PaymentNewParamsCartShipmentsAddressAddressReferenceID],
// [PaymentNewParamsCartShipmentsAddressAddressReferenceExplicit].
type PaymentNewParamsCartShipmentsAddress interface {
	implementsPaymentNewParamsCartShipmentsAddress()
}

type PaymentNewParamsCartShipmentsAddressAddressReferenceID struct {
	// The address's ID
	ID param.Field[string] `json:"id,required"`
	// The type of address reference
	Tag param.Field[PaymentNewParamsCartShipmentsAddressAddressReferenceIDTag] `json:".tag,required"`
}

func (r PaymentNewParamsCartShipmentsAddressAddressReferenceID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PaymentNewParamsCartShipmentsAddressAddressReferenceID) implementsPaymentNewParamsCartShipmentsAddress() {
}

// The type of address reference
type PaymentNewParamsCartShipmentsAddressAddressReferenceIDTag string

const (
	PaymentNewParamsCartShipmentsAddressAddressReferenceIDTagID PaymentNewParamsCartShipmentsAddressAddressReferenceIDTag = "id"
)

type PaymentNewParamsCartShipmentsAddressAddressReferenceExplicit struct {
	// The type of address reference
	Tag            param.Field[PaymentNewParamsCartShipmentsAddressAddressReferenceExplicitTag] `json:".tag,required"`
	CountryCode    param.Field[string]                                                          `json:"country_code,required"`
	FirstName      param.Field[string]                                                          `json:"first_name,required"`
	LastName       param.Field[string]                                                          `json:"last_name,required"`
	Locality       param.Field[string]                                                          `json:"locality,required"`
	PostalCode     param.Field[string]                                                          `json:"postal_code,required"`
	StreetAddress1 param.Field[string]                                                          `json:"street_address1,required"`
	Company        param.Field[string]                                                          `json:"company"`
	Email          param.Field[string]                                                          `json:"email" format:"email"`
	Phone          param.Field[string]                                                          `json:"phone" format:"phone"`
	Region         param.Field[string]                                                          `json:"region"`
	StreetAddress2 param.Field[string]                                                          `json:"street_address2"`
}

func (r PaymentNewParamsCartShipmentsAddressAddressReferenceExplicit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PaymentNewParamsCartShipmentsAddressAddressReferenceExplicit) implementsPaymentNewParamsCartShipmentsAddress() {
}

// The type of address reference
type PaymentNewParamsCartShipmentsAddressAddressReferenceExplicitTag string

const (
	PaymentNewParamsCartShipmentsAddressAddressReferenceExplicitTagExplicit PaymentNewParamsCartShipmentsAddressAddressReferenceExplicitTag = "explicit"
)

type PaymentNewParamsCartShipmentsCost struct {
	Currency param.Field[string] `json:"currency,required"`
	// The total amount, in cents, including its items and taxes (if applicable).
	Total param.Field[int64] `json:"total,required"`
	// The total tax amount, in cents for all of the items associated with the cart.
	Tax param.Field[int64] `json:"tax"`
}

func (r PaymentNewParamsCartShipmentsCost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsPaymentMethod struct {
	// Payment ID of the saved Bolt Payment method.
	ID  param.Field[string]                           `json:"id,required"`
	Tag param.Field[PaymentNewParamsPaymentMethodTag] `json:".tag,required"`
}

func (r PaymentNewParamsPaymentMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsPaymentMethodTag string

const (
	PaymentNewParamsPaymentMethodTagSavedPaymentMethod PaymentNewParamsPaymentMethodTag = "saved_payment_method"
)
