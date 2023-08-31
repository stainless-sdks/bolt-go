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

func TestPaymentNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := bolt.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Payments.New(context.TODO(), bolt.PaymentNewParams{
		Cart: bolt.F(bolt.PaymentNewParamsCart{
			Amounts: bolt.F(bolt.PaymentNewParamsCartAmounts{
				Total:    bolt.F(int64(900)),
				Currency: bolt.F("USD"),
				Tax:      bolt.F(int64(900)),
			}),
			OrderReference:   bolt.F("order_100"),
			OrderDescription: bolt.F("Order #1234567890"),
			DisplayID:        bolt.F("215614191"),
			Shipments: bolt.F([]bolt.PaymentNewParamsCartShipment{{
				Address: bolt.F[bolt.PaymentNewParamsCartShipmentsAddress](bolt.PaymentNewParamsCartShipmentsAddressAddressReferenceID(bolt.PaymentNewParamsCartShipmentsAddressAddressReferenceID{
					Tag: bolt.F(bolt.PaymentNewParamsCartShipmentsAddressAddressReferenceIDTagID),
					ID:  bolt.F("D4g3h5tBuVYK9"),
				})),
				Cost: bolt.F(bolt.PaymentNewParamsCartShipmentsCost{
					Total:    bolt.F(int64(900)),
					Currency: bolt.F("USD"),
					Tax:      bolt.F(int64(900)),
				}),
				Carrier: bolt.F("FedEx"),
			}, {
				Address: bolt.F[bolt.PaymentNewParamsCartShipmentsAddress](bolt.PaymentNewParamsCartShipmentsAddressAddressReferenceID(bolt.PaymentNewParamsCartShipmentsAddressAddressReferenceID{
					Tag: bolt.F(bolt.PaymentNewParamsCartShipmentsAddressAddressReferenceIDTagID),
					ID:  bolt.F("D4g3h5tBuVYK9"),
				})),
				Cost: bolt.F(bolt.PaymentNewParamsCartShipmentsCost{
					Total:    bolt.F(int64(900)),
					Currency: bolt.F("USD"),
					Tax:      bolt.F(int64(900)),
				}),
				Carrier: bolt.F("FedEx"),
			}, {
				Address: bolt.F[bolt.PaymentNewParamsCartShipmentsAddress](bolt.PaymentNewParamsCartShipmentsAddressAddressReferenceID(bolt.PaymentNewParamsCartShipmentsAddressAddressReferenceID{
					Tag: bolt.F(bolt.PaymentNewParamsCartShipmentsAddressAddressReferenceIDTagID),
					ID:  bolt.F("D4g3h5tBuVYK9"),
				})),
				Cost: bolt.F(bolt.PaymentNewParamsCartShipmentsCost{
					Total:    bolt.F(int64(900)),
					Currency: bolt.F("USD"),
					Tax:      bolt.F(int64(900)),
				}),
				Carrier: bolt.F("FedEx"),
			}}),
			Discounts: bolt.F([]bolt.PaymentNewParamsCartDiscount{{
				Amounts: bolt.F(bolt.PaymentNewParamsCartDiscountsAmounts{
					Total:    bolt.F(int64(900)),
					Currency: bolt.F("USD"),
					Tax:      bolt.F(int64(900)),
				}),
				Code:       bolt.F("SUMMER10DISCOUNT"),
				DetailsURL: bolt.F("https://www.example.com/SUMMER-SALE"),
			}, {
				Amounts: bolt.F(bolt.PaymentNewParamsCartDiscountsAmounts{
					Total:    bolt.F(int64(900)),
					Currency: bolt.F("USD"),
					Tax:      bolt.F(int64(900)),
				}),
				Code:       bolt.F("SUMMER10DISCOUNT"),
				DetailsURL: bolt.F("https://www.example.com/SUMMER-SALE"),
			}, {
				Amounts: bolt.F(bolt.PaymentNewParamsCartDiscountsAmounts{
					Total:    bolt.F(int64(900)),
					Currency: bolt.F("USD"),
					Tax:      bolt.F(int64(900)),
				}),
				Code:       bolt.F("SUMMER10DISCOUNT"),
				DetailsURL: bolt.F("https://www.example.com/SUMMER-SALE"),
			}}),
			Items: bolt.F([]bolt.PaymentNewParamsCartItem{{
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
		PaymentMethod: bolt.F(bolt.PaymentNewParamsPaymentMethod{
			Tag: bolt.F(bolt.PaymentNewParamsPaymentMethodTagSavedPaymentMethod),
			ID:  bolt.F("id"),
		}),
		XPublishableKey: bolt.F("string"),
	})
	if err != nil {
		var apierr *bolt.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
