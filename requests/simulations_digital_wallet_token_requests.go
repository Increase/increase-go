package requests

import (
	"fmt"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/fields"
)

type DigitalWalletTokenRequestCreateResponse struct {
	// If the simulated tokenization attempt was declined, this field contains details
	// as to why.
	DeclineReason fields.Field[DigitalWalletTokenRequestCreateResponseDeclineReason] `json:"decline_reason,required,nullable"`
	// If the simulated tokenization attempt was accepted, this field contains the id
	// of the Digital Wallet Token that was created.
	DigitalWalletTokenID fields.Field[string] `json:"digital_wallet_token_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_digital_wallet_token_request_simulation_result`.
	Type fields.Field[DigitalWalletTokenRequestCreateResponseType] `json:"type,required"`
}

// MarshalJSON serializes DigitalWalletTokenRequestCreateResponse into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *DigitalWalletTokenRequestCreateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DigitalWalletTokenRequestCreateResponse) String() (result string) {
	return fmt.Sprintf("&DigitalWalletTokenRequestCreateResponse{DeclineReason:%s DigitalWalletTokenID:%s Type:%s}", r.DeclineReason, r.DigitalWalletTokenID, r.Type)
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

type SimulateDigitalWalletProvisioningForACardParameters struct {
	// The identifier of the Card to be authorized.
	CardID fields.Field[string] `json:"card_id,required"`
}

// MarshalJSON serializes SimulateDigitalWalletProvisioningForACardParameters into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *SimulateDigitalWalletProvisioningForACardParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SimulateDigitalWalletProvisioningForACardParameters) String() (result string) {
	return fmt.Sprintf("&SimulateDigitalWalletProvisioningForACardParameters{CardID:%s}", r.CardID)
}
