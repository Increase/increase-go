// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// SimulationCardAuthenticationService contains methods and other services that
// help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationCardAuthenticationService] method instead.
type SimulationCardAuthenticationService struct {
	Options []option.RequestOption
}

// NewSimulationCardAuthenticationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationCardAuthenticationService(opts ...option.RequestOption) (r *SimulationCardAuthenticationService) {
	r = &SimulationCardAuthenticationService{}
	r.Options = opts
	return
}

// Simulates a Card Authentication attempt on a [Card](#cards). The attempt always
// results in a [Card Payment](#card_payments) being created, either with a status
// that allows further action or a terminal failed status.
func (r *SimulationCardAuthenticationService) New(ctx context.Context, body SimulationCardAuthenticationNewParams, opts ...option.RequestOption) (res *CardPayment, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "simulations/card_authentications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Simulates an attempt at a Card Authentication Challenge. This updates the
// `card_authentications` object under the [Card Payment](#card_payments). You can
// also attempt the challenge by navigating to
// https://dashboard.increase.com/card_authentication_simulation/:card_payment_id.
func (r *SimulationCardAuthenticationService) ChallengeAttempts(ctx context.Context, cardPaymentID string, body SimulationCardAuthenticationChallengeAttemptsParams, opts ...option.RequestOption) (res *CardPayment, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardPaymentID == "" {
		err = errors.New("missing required card_payment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("simulations/card_authentications/%s/challenge_attempts", cardPaymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Simulates starting a Card Authentication Challenge for an existing Card
// Authentication. This updates the `card_authentications` object under the
// [Card Payment](#card_payments). To attempt the challenge, use the
// `/simulations/card_authentications/:card_payment_id/challenge_attempts` endpoint
// or navigate to
// https://dashboard.increase.com/card_authentication_simulation/:card_payment_id.
func (r *SimulationCardAuthenticationService) Challenges(ctx context.Context, cardPaymentID string, opts ...option.RequestOption) (res *CardPayment, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardPaymentID == "" {
		err = errors.New("missing required card_payment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("simulations/card_authentications/%s/challenges", cardPaymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type SimulationCardAuthenticationNewParams struct {
	// The identifier of the Card to be authorized.
	CardID param.Field[string] `json:"card_id" api:"required"`
	// The category of the card authentication attempt.
	Category param.Field[SimulationCardAuthenticationNewParamsCategory] `json:"category"`
	// The device channel of the card authentication attempt.
	DeviceChannel param.Field[SimulationCardAuthenticationNewParamsDeviceChannel] `json:"device_channel"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID param.Field[string] `json:"merchant_acceptor_id"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode param.Field[string] `json:"merchant_category_code"`
	// The country the merchant resides in.
	MerchantCountry param.Field[string] `json:"merchant_country"`
	// The name of the merchant
	MerchantName param.Field[string] `json:"merchant_name"`
	// The purchase amount in cents.
	PurchaseAmount param.Field[int64] `json:"purchase_amount"`
}

func (r SimulationCardAuthenticationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The category of the card authentication attempt.
type SimulationCardAuthenticationNewParamsCategory string

const (
	SimulationCardAuthenticationNewParamsCategoryPaymentAuthentication    SimulationCardAuthenticationNewParamsCategory = "payment_authentication"
	SimulationCardAuthenticationNewParamsCategoryNonPaymentAuthentication SimulationCardAuthenticationNewParamsCategory = "non_payment_authentication"
)

func (r SimulationCardAuthenticationNewParamsCategory) IsKnown() bool {
	switch r {
	case SimulationCardAuthenticationNewParamsCategoryPaymentAuthentication, SimulationCardAuthenticationNewParamsCategoryNonPaymentAuthentication:
		return true
	}
	return false
}

// The device channel of the card authentication attempt.
type SimulationCardAuthenticationNewParamsDeviceChannel string

const (
	SimulationCardAuthenticationNewParamsDeviceChannelApp                       SimulationCardAuthenticationNewParamsDeviceChannel = "app"
	SimulationCardAuthenticationNewParamsDeviceChannelBrowser                   SimulationCardAuthenticationNewParamsDeviceChannel = "browser"
	SimulationCardAuthenticationNewParamsDeviceChannelThreeDSRequestorInitiated SimulationCardAuthenticationNewParamsDeviceChannel = "three_ds_requestor_initiated"
)

func (r SimulationCardAuthenticationNewParamsDeviceChannel) IsKnown() bool {
	switch r {
	case SimulationCardAuthenticationNewParamsDeviceChannelApp, SimulationCardAuthenticationNewParamsDeviceChannelBrowser, SimulationCardAuthenticationNewParamsDeviceChannelThreeDSRequestorInitiated:
		return true
	}
	return false
}

type SimulationCardAuthenticationChallengeAttemptsParams struct {
	// The one-time code to be validated.
	OneTimeCode param.Field[string] `json:"one_time_code" api:"required"`
}

func (r SimulationCardAuthenticationChallengeAttemptsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
