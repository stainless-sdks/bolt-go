// File generated from our OpenAPI spec by Stainless.

package bolt

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/bolt/bolt-go/internal/apijson"
	"github.com/bolt/bolt-go/internal/param"
	"github.com/bolt/bolt-go/internal/requestconfig"
	"github.com/bolt/bolt-go/option"
	"github.com/tidwall/gjson"
)

// WebhookService contains methods and other services that help with interacting
// with the bolt API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts
	return
}

// Retrieve information for an existing webhook.
func (r *WebhookService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Create a new webhook to receive notifications from Bolt about various events,
// such as transaction status.
func (r *WebhookService) Update(ctx context.Context, body WebhookUpdateParams, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = append(r.Options[:], opts...)
	path := "webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve information about all existing webhooks.
func (r *WebhookService) List(ctx context.Context, opts ...option.RequestOption) (res *WebhookListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an existing webhook. You will no longer receive notifications from Bolt
// about its events.
func (r *WebhookService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Webhook struct {
	Event WebhookEvent `json:"event,required"`
	// The webhook's URL
	URL string `json:"url,required" format:"uri"`
	// The webhook's ID
	ID string `json:"id"`
	// The time at which the webhook was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	JSON      webhookJSON
}

// webhookJSON contains the JSON metadata for the struct [Webhook]
type webhookJSON struct {
	Event       apijson.Field
	URL         apijson.Field
	ID          apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Webhook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [WebhookEventEventGroup] or [WebhookEventEventList].
type WebhookEvent interface {
	implementsWebhookEvent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WebhookEvent)(nil)).Elem(),
		".tag",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"group\"",
			Type:               reflect.TypeOf(WebhookEventEventGroup{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"list\"",
			Type:               reflect.TypeOf(WebhookEventEventList{}),
		},
	)
}

type WebhookEventEventGroup struct {
	Tag        WebhookEventEventGroupTag        `json:".tag,required"`
	EventGroup WebhookEventEventGroupEventGroup `json:"event_group,required"`
	JSON       webhookEventEventGroupJSON
}

// webhookEventEventGroupJSON contains the JSON metadata for the struct
// [WebhookEventEventGroup]
type webhookEventEventGroupJSON struct {
	Tag         apijson.Field
	EventGroup  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookEventEventGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WebhookEventEventGroup) implementsWebhookEvent() {}

type WebhookEventEventGroupTag string

const (
	WebhookEventEventGroupTagGroup WebhookEventEventGroupTag = "group"
)

type WebhookEventEventGroupEventGroup string

const (
	WebhookEventEventGroupEventGroupAll WebhookEventEventGroupEventGroup = "all"
)

type WebhookEventEventList struct {
	Tag       WebhookEventEventListTag         `json:".tag,required"`
	EventList []WebhookEventEventListEventList `json:"event_list,required"`
	JSON      webhookEventEventListJSON
}

// webhookEventEventListJSON contains the JSON metadata for the struct
// [WebhookEventEventList]
type webhookEventEventListJSON struct {
	Tag         apijson.Field
	EventList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookEventEventList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WebhookEventEventList) implementsWebhookEvent() {}

type WebhookEventEventListTag string

const (
	WebhookEventEventListTagList WebhookEventEventListTag = "list"
)

type WebhookEventEventListEventList string

const (
	WebhookEventEventListEventListPayment                WebhookEventEventListEventList = "payment"
	WebhookEventEventListEventListCredit                 WebhookEventEventListEventList = "credit"
	WebhookEventEventListEventListCapture                WebhookEventEventListEventList = "capture"
	WebhookEventEventListEventListVoid                   WebhookEventEventListEventList = "void"
	WebhookEventEventListEventListAuth                   WebhookEventEventListEventList = "auth"
	WebhookEventEventListEventListPending                WebhookEventEventListEventList = "pending"
	WebhookEventEventListEventListRejectedIrreversible   WebhookEventEventListEventList = "rejected_irreversible"
	WebhookEventEventListEventListRejectedReversible     WebhookEventEventListEventList = "rejected_reversible"
	WebhookEventEventListEventListFailedPayment          WebhookEventEventListEventList = "failed_payment"
	WebhookEventEventListEventListNewsletterSubscription WebhookEventEventListEventList = "newsletter_subscription"
	WebhookEventEventListEventListRiskInsights           WebhookEventEventListEventList = "risk_insights"
	WebhookEventEventListEventListCreditCardDeleted      WebhookEventEventListEventList = "credit_card_deleted"
)

type WebhookListResponse struct {
	Webhooks []Webhook `json:"webhooks"`
	JSON     webhookListResponseJSON
}

// webhookListResponseJSON contains the JSON metadata for the struct
// [WebhookListResponse]
type webhookListResponseJSON struct {
	Webhooks    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateParams struct {
	Event param.Field[WebhookUpdateParamsEvent] `json:"event,required"`
	// The webhook's URL
	URL param.Field[string] `json:"url,required" format:"uri"`
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [WebhookUpdateParamsEventEventGroup],
// [WebhookUpdateParamsEventEventList].
type WebhookUpdateParamsEvent interface {
	implementsWebhookUpdateParamsEvent()
}

type WebhookUpdateParamsEventEventGroup struct {
	Tag        param.Field[WebhookUpdateParamsEventEventGroupTag]        `json:".tag,required"`
	EventGroup param.Field[WebhookUpdateParamsEventEventGroupEventGroup] `json:"event_group,required"`
}

func (r WebhookUpdateParamsEventEventGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WebhookUpdateParamsEventEventGroup) implementsWebhookUpdateParamsEvent() {}

type WebhookUpdateParamsEventEventGroupTag string

const (
	WebhookUpdateParamsEventEventGroupTagGroup WebhookUpdateParamsEventEventGroupTag = "group"
)

type WebhookUpdateParamsEventEventGroupEventGroup string

const (
	WebhookUpdateParamsEventEventGroupEventGroupAll WebhookUpdateParamsEventEventGroupEventGroup = "all"
)

type WebhookUpdateParamsEventEventList struct {
	Tag       param.Field[WebhookUpdateParamsEventEventListTag]         `json:".tag,required"`
	EventList param.Field[[]WebhookUpdateParamsEventEventListEventList] `json:"event_list,required"`
}

func (r WebhookUpdateParamsEventEventList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WebhookUpdateParamsEventEventList) implementsWebhookUpdateParamsEvent() {}

type WebhookUpdateParamsEventEventListTag string

const (
	WebhookUpdateParamsEventEventListTagList WebhookUpdateParamsEventEventListTag = "list"
)

type WebhookUpdateParamsEventEventListEventList string

const (
	WebhookUpdateParamsEventEventListEventListPayment                WebhookUpdateParamsEventEventListEventList = "payment"
	WebhookUpdateParamsEventEventListEventListCredit                 WebhookUpdateParamsEventEventListEventList = "credit"
	WebhookUpdateParamsEventEventListEventListCapture                WebhookUpdateParamsEventEventListEventList = "capture"
	WebhookUpdateParamsEventEventListEventListVoid                   WebhookUpdateParamsEventEventListEventList = "void"
	WebhookUpdateParamsEventEventListEventListAuth                   WebhookUpdateParamsEventEventListEventList = "auth"
	WebhookUpdateParamsEventEventListEventListPending                WebhookUpdateParamsEventEventListEventList = "pending"
	WebhookUpdateParamsEventEventListEventListRejectedIrreversible   WebhookUpdateParamsEventEventListEventList = "rejected_irreversible"
	WebhookUpdateParamsEventEventListEventListRejectedReversible     WebhookUpdateParamsEventEventListEventList = "rejected_reversible"
	WebhookUpdateParamsEventEventListEventListFailedPayment          WebhookUpdateParamsEventEventListEventList = "failed_payment"
	WebhookUpdateParamsEventEventListEventListNewsletterSubscription WebhookUpdateParamsEventEventListEventList = "newsletter_subscription"
	WebhookUpdateParamsEventEventListEventListRiskInsights           WebhookUpdateParamsEventEventListEventList = "risk_insights"
	WebhookUpdateParamsEventEventListEventListCreditCardDeleted      WebhookUpdateParamsEventEventListEventList = "credit_card_deleted"
)
