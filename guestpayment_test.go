// File generated from our OpenAPI spec by Stainless.

package bolt_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bolt/bolt-go"
	"github.com/bolt/bolt-go/internal/testutil"
	"github.com/bolt/bolt-go/option"
)

func TestGuestPaymentNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	t.Skip("skipped: tests are disabled for the time being")
	client := bolt.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithPublishableKey("ABC.123.345"),
	)
	_, err := client.Guests.Payments.New(context.TODO(), bolt.GuestPaymentNewParams{
		Cart: bolt.F(bolt.GuestPaymentNewParamsCart{
			Amounts: bolt.F(bolt.GuestPaymentNewParamsCartAmounts{
				Total:    bolt.F(int64(900)),
				Currency: bolt.F("USD"),
				Tax:      bolt.F(int64(900)),
			}),
			OrderReference:   bolt.F("order_100"),
			OrderDescription: bolt.F("Order #1234567890"),
			DisplayID:        bolt.F("215614191"),
			Shipments: bolt.F([]bolt.GuestPaymentNewParamsCartShipment{{
				Address: bolt.F[bolt.GuestPaymentNewParamsCartShipmentsAddress](bolt.GuestPaymentNewParamsCartShipmentsAddressAddressReferenceID(bolt.GuestPaymentNewParamsCartShipmentsAddressAddressReferenceID{
					Tag: bolt.F(bolt.GuestPaymentNewParamsCartShipmentsAddressAddressReferenceIDTagID),
					ID:  bolt.F("D4g3h5tBuVYK9"),
				})),
				Cost: bolt.F(bolt.GuestPaymentNewParamsCartShipmentsCost{
					Total:    bolt.F(int64(900)),
					Currency: bolt.F("USD"),
					Tax:      bolt.F(int64(900)),
				}),
				Carrier: bolt.F("FedEx"),
			}, {
				Address: bolt.F[bolt.GuestPaymentNewParamsCartShipmentsAddress](bolt.GuestPaymentNewParamsCartShipmentsAddressAddressReferenceID(bolt.GuestPaymentNewParamsCartShipmentsAddressAddressReferenceID{
					Tag: bolt.F(bolt.GuestPaymentNewParamsCartShipmentsAddressAddressReferenceIDTagID),
					ID:  bolt.F("D4g3h5tBuVYK9"),
				})),
				Cost: bolt.F(bolt.GuestPaymentNewParamsCartShipmentsCost{
					Total:    bolt.F(int64(900)),
					Currency: bolt.F("USD"),
					Tax:      bolt.F(int64(900)),
				}),
				Carrier: bolt.F("FedEx"),
			}, {
				Address: bolt.F[bolt.GuestPaymentNewParamsCartShipmentsAddress](bolt.GuestPaymentNewParamsCartShipmentsAddressAddressReferenceID(bolt.GuestPaymentNewParamsCartShipmentsAddressAddressReferenceID{
					Tag: bolt.F(bolt.GuestPaymentNewParamsCartShipmentsAddressAddressReferenceIDTagID),
					ID:  bolt.F("D4g3h5tBuVYK9"),
				})),
				Cost: bolt.F(bolt.GuestPaymentNewParamsCartShipmentsCost{
					Total:    bolt.F(int64(900)),
					Currency: bolt.F("USD"),
					Tax:      bolt.F(int64(900)),
				}),
				Carrier: bolt.F("FedEx"),
			}}),
			Discounts: bolt.F([]bolt.GuestPaymentNewParamsCartDiscount{{
				Amounts: bolt.F(bolt.GuestPaymentNewParamsCartDiscountsAmounts{
					Total:    bolt.F(int64(900)),
					Currency: bolt.F("USD"),
					Tax:      bolt.F(int64(900)),
				}),
				Code:       bolt.F("SUMMER10DISCOUNT"),
				DetailsURL: bolt.F("https://www.example.com/SUMMER-SALE"),
			}, {
				Amounts: bolt.F(bolt.GuestPaymentNewParamsCartDiscountsAmounts{
					Total:    bolt.F(int64(900)),
					Currency: bolt.F("USD"),
					Tax:      bolt.F(int64(900)),
				}),
				Code:       bolt.F("SUMMER10DISCOUNT"),
				DetailsURL: bolt.F("https://www.example.com/SUMMER-SALE"),
			}, {
				Amounts: bolt.F(bolt.GuestPaymentNewParamsCartDiscountsAmounts{
					Total:    bolt.F(int64(900)),
					Currency: bolt.F("USD"),
					Tax:      bolt.F(int64(900)),
				}),
				Code:       bolt.F("SUMMER10DISCOUNT"),
				DetailsURL: bolt.F("https://www.example.com/SUMMER-SALE"),
			}}),
			Items: bolt.F([]bolt.GuestPaymentNewParamsCartItem{{
				Name:        bolt.F("Bolt Swag Bag"),
				Reference:   bolt.F("item_100"),
				Description: bolt.F("Large tote with Bolt logo."),
				TotalAmount: bolt.F(int64(1000)),
				UnitPrice:   bolt.F(int64(1000)),
				Quantity:    bolt.F(int64(1)),
				ImageURL:    bolt.F("https://www.example.com/products/123456/images/1.png"),
			}, {
				Name:        bolt.F("Bolt Swag Bag"),
				Reference:   bolt.F("item_100"),
				Description: bolt.F("Large tote with Bolt logo."),
				TotalAmount: bolt.F(int64(1000)),
				UnitPrice:   bolt.F(int64(1000)),
				Quantity:    bolt.F(int64(1)),
				ImageURL:    bolt.F("https://www.example.com/products/123456/images/1.png"),
			}, {
				Name:        bolt.F("Bolt Swag Bag"),
				Reference:   bolt.F("item_100"),
				Description: bolt.F("Large tote with Bolt logo."),
				TotalAmount: bolt.F(int64(1000)),
				UnitPrice:   bolt.F(int64(1000)),
				Quantity:    bolt.F(int64(1)),
				ImageURL:    bolt.F("https://www.example.com/products/123456/images/1.png"),
			}}),
		}),
		PaymentMethod: bolt.F(bolt.GuestPaymentNewParamsPaymentMethod{
			Tag:     bolt.F(bolt.GuestPaymentNewParamsPaymentMethodTagPaypal),
			Success: bolt.F("www.example.com/handle_paypal_success"),
			Cancel:  bolt.F("www.example.com/handle_paypal_cancel"),
		}),
	})
	if err != nil {
		var apierr *bolt.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
