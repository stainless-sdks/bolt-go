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

func TestMerchantCallbackUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := bolt.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithPublishableKey("ABC.123.345"),
	)
	_, err := client.Merchants.Callbacks.Update(context.TODO(), bolt.MerchantCallbackUpdateParams{
		AccountPage:                   bolt.F("https://www.example.com/account"),
		BaseDomain:                    bolt.F("https://www.example.com/"),
		ConfirmationRedirect:          bolt.F("https://www.example.com/bolt/redirect"),
		CreateOrder:                   bolt.F("https://www.example.com/bolt/order"),
		Debug:                         bolt.F("https://www.example.com/bolt/debug"),
		GetAccount:                    bolt.F("https://www.example.com/bolt/account"),
		MobileAppDomain:               bolt.F("https://m.example.com/"),
		OauthLogout:                   bolt.F("https://www.example.com/bolt/logout"),
		OauthRedirect:                 bolt.F("https://www.example.com/bolt/oauth"),
		PrivacyPolicy:                 bolt.F("https://www.example.com/privacy-policy"),
		ProductInfo:                   bolt.F("https://www.example.com/bolt/product"),
		RemoteAPI:                     bolt.F("https://www.example.com/bolt/remote-api"),
		Shipping:                      bolt.F("https://www.example.com/bolt/shipping"),
		SupportPage:                   bolt.F("https://www.example.com/help"),
		Tax:                           bolt.F("https://www.example.com/bolt/tax"),
		TermsOfService:                bolt.F("https://www.example.com/terms-of-service"),
		UniversalMerchantAPI:          bolt.F("https://www.example.com/bolt/merchant-api"),
		UpdateCart:                    bolt.F("https://www.example.com/bolt/cart"),
		ValidateAdditionalAccountData: bolt.F("https://www.example.com/bolt/validate-account"),
	})
	if err != nil {
		var apierr *bolt.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMerchantCallbackList(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := bolt.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithPublishableKey("ABC.123.345"),
	)
	_, err := client.Merchants.Callbacks.List(context.TODO())
	if err != nil {
		var apierr *bolt.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
