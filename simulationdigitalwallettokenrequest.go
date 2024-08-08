// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// SimulationDigitalWalletTokenRequestService contains methods and other services
// that help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationDigitalWalletTokenRequestService] method instead.
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
func (r *SimulationDigitalWalletTokenRequestService) New(ctx context.Context, body SimulationDigitalWalletTokenRequestNewParams, opts ...option.RequestOption) (res *SimulationDigitalWalletTokenRequestNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/digital_wallet_token_requests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The results of a Digital Wallet Token simulation.
type SimulationDigitalWalletTokenRequestNewResponse struct {
	// If the simulated tokenization attempt was declined, this field contains details
	// as to why.
	DeclineReason SimulationDigitalWalletTokenRequestNewResponseDeclineReason `json:"decline_reason,required,nullable"`
	// If the simulated tokenization attempt was accepted, this field contains the id
	// of the Digital Wallet Token that was created.
	DigitalWalletTokenID string `json:"digital_wallet_token_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_digital_wallet_token_request_simulation_result`.
	Type SimulationDigitalWalletTokenRequestNewResponseType `json:"type,required"`
	JSON simulationDigitalWalletTokenRequestNewResponseJSON `json:"-"`
}

// simulationDigitalWalletTokenRequestNewResponseJSON contains the JSON metadata
// for the struct [SimulationDigitalWalletTokenRequestNewResponse]
type simulationDigitalWalletTokenRequestNewResponseJSON struct {
	DeclineReason        apijson.Field
	DigitalWalletTokenID apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SimulationDigitalWalletTokenRequestNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r simulationDigitalWalletTokenRequestNewResponseJSON) RawJSON() string {
	return r.raw
}

// If the simulated tokenization attempt was declined, this field contains details
// as to why.
type SimulationDigitalWalletTokenRequestNewResponseDeclineReason string

const (
	// The card is not active.
	SimulationDigitalWalletTokenRequestNewResponseDeclineReasonCardNotActive SimulationDigitalWalletTokenRequestNewResponseDeclineReason = "card_not_active"
	// The card does not have a two-factor authentication method.
	SimulationDigitalWalletTokenRequestNewResponseDeclineReasonNoVerificationMethod SimulationDigitalWalletTokenRequestNewResponseDeclineReason = "no_verification_method"
	// Your webhook timed out when evaluating the token provisioning attempt.
	SimulationDigitalWalletTokenRequestNewResponseDeclineReasonWebhookTimedOut SimulationDigitalWalletTokenRequestNewResponseDeclineReason = "webhook_timed_out"
	// Your webhook declined the token provisioning attempt.
	SimulationDigitalWalletTokenRequestNewResponseDeclineReasonWebhookDeclined SimulationDigitalWalletTokenRequestNewResponseDeclineReason = "webhook_declined"
)

func (r SimulationDigitalWalletTokenRequestNewResponseDeclineReason) IsKnown() bool {
	switch r {
	case SimulationDigitalWalletTokenRequestNewResponseDeclineReasonCardNotActive, SimulationDigitalWalletTokenRequestNewResponseDeclineReasonNoVerificationMethod, SimulationDigitalWalletTokenRequestNewResponseDeclineReasonWebhookTimedOut, SimulationDigitalWalletTokenRequestNewResponseDeclineReasonWebhookDeclined:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `inbound_digital_wallet_token_request_simulation_result`.
type SimulationDigitalWalletTokenRequestNewResponseType string

const (
	SimulationDigitalWalletTokenRequestNewResponseTypeInboundDigitalWalletTokenRequestSimulationResult SimulationDigitalWalletTokenRequestNewResponseType = "inbound_digital_wallet_token_request_simulation_result"
)

func (r SimulationDigitalWalletTokenRequestNewResponseType) IsKnown() bool {
	switch r {
	case SimulationDigitalWalletTokenRequestNewResponseTypeInboundDigitalWalletTokenRequestSimulationResult:
		return true
	}
	return false
}

type SimulationDigitalWalletTokenRequestNewParams struct {
	// The identifier of the Card to be authorized.
	CardID param.Field[string] `json:"card_id,required"`
}

func (r SimulationDigitalWalletTokenRequestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
