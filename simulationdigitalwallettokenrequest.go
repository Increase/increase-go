// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// SimulationDigitalWalletTokenRequestService contains methods and other services
// that help with interacting with the increase API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewSimulationDigitalWalletTokenRequestService] method instead.
type SimulationDigitalWalletTokenRequestService struct {
	Options []option.RequestOption
}

// NewSimulationDigitalWalletTokenRequestService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewSimulationDigitalWalletTokenRequestService(opts ...option.RequestOption) (r *SimulationDigitalWalletTokenRequestService) {
	r = &SimulationDigitalWalletTokenRequestService{}
	r.Options = opts
	return
}

// Simulates a user attempting add a [Card](#cards) to a digital wallet such as
// Apple Pay.
func (r *SimulationDigitalWalletTokenRequestService) New(ctx context.Context, body SimulationDigitalWalletTokenRequestNewParams, opts ...option.RequestOption) (res *DigitalWalletTokenRequestCreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/digital_wallet_token_requests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The results of a Digital Wallet Token simulation.
type DigitalWalletTokenRequestCreateResponse struct {
	// If the simulated tokenization attempt was declined, this field contains details
	// as to why.
	DeclineReason DigitalWalletTokenRequestCreateResponseDeclineReason `json:"decline_reason,required,nullable"`
	// If the simulated tokenization attempt was accepted, this field contains the id
	// of the Digital Wallet Token that was created.
	DigitalWalletTokenID string `json:"digital_wallet_token_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_digital_wallet_token_request_simulation_result`.
	Type DigitalWalletTokenRequestCreateResponseType `json:"type,required"`
	JSON digitalWalletTokenRequestCreateResponseJSON
}

// digitalWalletTokenRequestCreateResponseJSON contains the JSON metadata for the
// struct [DigitalWalletTokenRequestCreateResponse]
type digitalWalletTokenRequestCreateResponseJSON struct {
	DeclineReason        apijson.Field
	DigitalWalletTokenID apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DigitalWalletTokenRequestCreateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DigitalWalletTokenRequestCreateResponseDeclineReason string

const (
	DigitalWalletTokenRequestCreateResponseDeclineReasonCardNotActive        DigitalWalletTokenRequestCreateResponseDeclineReason = "card_not_active"
	DigitalWalletTokenRequestCreateResponseDeclineReasonNoVerificationMethod DigitalWalletTokenRequestCreateResponseDeclineReason = "no_verification_method"
	DigitalWalletTokenRequestCreateResponseDeclineReasonWebhookTimedOut      DigitalWalletTokenRequestCreateResponseDeclineReason = "webhook_timed_out"
	DigitalWalletTokenRequestCreateResponseDeclineReasonWebhookDeclined      DigitalWalletTokenRequestCreateResponseDeclineReason = "webhook_declined"
)

type DigitalWalletTokenRequestCreateResponseType string

const (
	DigitalWalletTokenRequestCreateResponseTypeInboundDigitalWalletTokenRequestSimulationResult DigitalWalletTokenRequestCreateResponseType = "inbound_digital_wallet_token_request_simulation_result"
)

type SimulationDigitalWalletTokenRequestNewParams struct {
	// The identifier of the Card to be authorized.
	CardID param.Field[string] `json:"card_id,required"`
}

func (r SimulationDigitalWalletTokenRequestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
