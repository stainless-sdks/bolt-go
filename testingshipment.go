// File generated from our OpenAPI spec by Stainless.

package bolt

import (
	"context"
	"net/http"
	"time"

	"github.com/bolt/bolt-go/internal/apijson"
	"github.com/bolt/bolt-go/internal/param"
	"github.com/bolt/bolt-go/internal/requestconfig"
	"github.com/bolt/bolt-go/option"
)

// TestingShipmentService contains methods and other services that help with
// interacting with the bolt API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewTestingShipmentService] method
// instead.
type TestingShipmentService struct {
	Options []option.RequestOption
}

// NewTestingShipmentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTestingShipmentService(opts ...option.RequestOption) (r *TestingShipmentService) {
	r = &TestingShipmentService{}
	r.Options = opts
	return
}

// Simulate a shipment tracking update, such as those that Bolt would ingest from
// third-party shipping providers that would allow updating tracking and delivery
// information to shipments associated with orders.
func (r *TestingShipmentService) New(ctx context.Context, body TestingShipmentNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "testing/shipments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type TestingShipmentNewParams struct {
	// The shipment's status.
	Status param.Field[TestingShipmentNewParamsStatus] `json:"status,required"`
	// A list of tracking updates that contain the shipment's status, location, and any
	// unique messages.
	TrackingDetails param.Field[[]TestingShipmentNewParamsTrackingDetail] `json:"tracking_details,required"`
	// The carrier's tracking number for the shipment. Must be prefixed with
	// `MockBolt`.
	TrackingNumber param.Field[string] `json:"tracking_number,required"`
	// The shipment's actual or estimated delivery date.
	DeliveryDate param.Field[time.Time] `json:"delivery_date" format:"date-time"`
}

func (r TestingShipmentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The shipment's status.
type TestingShipmentNewParamsStatus string

const (
	TestingShipmentNewParamsStatusInTransit TestingShipmentNewParamsStatus = "in_transit"
	TestingShipmentNewParamsStatusCancelled TestingShipmentNewParamsStatus = "cancelled"
	TestingShipmentNewParamsStatusFailure   TestingShipmentNewParamsStatus = "failure"
	TestingShipmentNewParamsStatusDelivered TestingShipmentNewParamsStatus = "delivered"
)

type TestingShipmentNewParamsTrackingDetail struct {
	// The country associated this this set of tracking details, if any.
	CountryCode param.Field[string] `json:"country_code"`
	// The tracking detail's timestamp.
	EventDate param.Field[string] `json:"event_date"`
	// The locality associated this this set of tracking details, if any.
	Locality param.Field[string] `json:"locality"`
	// An arbitrary, human-readable message associated with this set of tracking
	// details.
	Message param.Field[string] `json:"message"`
	// The postal associated this this set of tracking details, if any.
	PostalCode param.Field[string] `json:"postal_code"`
	// The region associated this this set of tracking details, if any.
	Region param.Field[string]                                        `json:"region"`
	Status param.Field[TestingShipmentNewParamsTrackingDetailsStatus] `json:"status"`
}

func (r TestingShipmentNewParamsTrackingDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TestingShipmentNewParamsTrackingDetailsStatus string

const (
	TestingShipmentNewParamsTrackingDetailsStatusUnknown            TestingShipmentNewParamsTrackingDetailsStatus = "unknown"
	TestingShipmentNewParamsTrackingDetailsStatusPreTransit         TestingShipmentNewParamsTrackingDetailsStatus = "pre_transit"
	TestingShipmentNewParamsTrackingDetailsStatusInTransit          TestingShipmentNewParamsTrackingDetailsStatus = "in_transit"
	TestingShipmentNewParamsTrackingDetailsStatusOutForDelivery     TestingShipmentNewParamsTrackingDetailsStatus = "out_for_delivery"
	TestingShipmentNewParamsTrackingDetailsStatusDelivered          TestingShipmentNewParamsTrackingDetailsStatus = "delivered"
	TestingShipmentNewParamsTrackingDetailsStatusAvailableForPickup TestingShipmentNewParamsTrackingDetailsStatus = "available_for_pickup"
	TestingShipmentNewParamsTrackingDetailsStatusReturnToSender     TestingShipmentNewParamsTrackingDetailsStatus = "return_to_sender"
	TestingShipmentNewParamsTrackingDetailsStatusFailure            TestingShipmentNewParamsTrackingDetailsStatus = "failure"
	TestingShipmentNewParamsTrackingDetailsStatusCancelled          TestingShipmentNewParamsTrackingDetailsStatus = "cancelled"
	TestingShipmentNewParamsTrackingDetailsStatusError              TestingShipmentNewParamsTrackingDetailsStatus = "error"
)
