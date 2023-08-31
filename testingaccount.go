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

// TestingAccountService contains methods and other services that help with
// interacting with the bolt API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewTestingAccountService] method instead.
type TestingAccountService struct {
	Options []option.RequestOption
}

// NewTestingAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTestingAccountService(opts ...option.RequestOption) (r *TestingAccountService) {
	r = &TestingAccountService{}
	r.Options = opts
	return
}

// Create a Bolt shopper account for testing purposes.
func (r *TestingAccountService) New(ctx context.Context, body TestingAccountNewParams, opts ...option.RequestOption) (res *AccountTest, err error) {
	opts = append(r.Options[:], opts...)
	path := "testing/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TestingAccountNewParams struct {
	DeactivateAt param.Field[time.Time]                         `json:"deactivate_at,required" format:"date-time"`
	EmailState   param.Field[TestingAccountNewParamsEmailState] `json:"email_state,required"`
	PhoneState   param.Field[TestingAccountNewParamsPhoneState] `json:"phone_state,required"`
	HasAddress   param.Field[bool]                              `json:"has_address"`
	IsMigrated   param.Field[bool]                              `json:"is_migrated"`
}

func (r TestingAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TestingAccountNewParamsEmailState string

const (
	TestingAccountNewParamsEmailStateMissing    TestingAccountNewParamsEmailState = "missing"
	TestingAccountNewParamsEmailStateUnverified TestingAccountNewParamsEmailState = "unverified"
	TestingAccountNewParamsEmailStateVerified   TestingAccountNewParamsEmailState = "verified"
)

type TestingAccountNewParamsPhoneState string

const (
	TestingAccountNewParamsPhoneStateMissing    TestingAccountNewParamsPhoneState = "missing"
	TestingAccountNewParamsPhoneStateUnverified TestingAccountNewParamsPhoneState = "unverified"
	TestingAccountNewParamsPhoneStateVerified   TestingAccountNewParamsPhoneState = "verified"
)
