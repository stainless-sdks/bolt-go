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

func TestAccountPaymentMethodNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := bolt.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Account.PaymentMethods.New(context.TODO(), bolt.AccountPaymentMethodNewParams{
		Token: bolt.F("a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0"),
		Tag:   bolt.F(bolt.AccountPaymentMethodNewParamsTagCreditCard),
		BillingAddress: bolt.F[bolt.AccountPaymentMethodNewParamsBillingAddress](bolt.AccountPaymentMethodNewParamsBillingAddressAddressReferenceID(bolt.AccountPaymentMethodNewParamsBillingAddressAddressReferenceID{
			Tag: bolt.F(bolt.AccountPaymentMethodNewParamsBillingAddressAddressReferenceIDTagID),
			ID:  bolt.F("D4g3h5tBuVYK9"),
		})),
		Bin:             bolt.F("411111"),
		Expiration:      bolt.F("2025-03"),
		Last4:           bolt.F("1004"),
		Network:         bolt.F(bolt.AccountPaymentMethodNewParamsNetworkVisa),
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

func TestAccountPaymentMethodDelete(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := bolt.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	err := client.Account.PaymentMethods.Delete(
		context.TODO(),
		"D4g3h5tBuVYK9",
		bolt.AccountPaymentMethodDeleteParams{
			XPublishableKey: bolt.F("string"),
		},
	)
	if err != nil {
		var apierr *bolt.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
