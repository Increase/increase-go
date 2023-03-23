package responses

import (
	pjson "github.com/increase/increase-go/core/json"
)

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
	JSON DigitalWalletTokenRequestCreateResponseJSON
}

type DigitalWalletTokenRequestCreateResponseJSON struct {
	DeclineReason        pjson.Metadata
	DigitalWalletTokenID pjson.Metadata
	Type                 pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// DigitalWalletTokenRequestCreateResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *DigitalWalletTokenRequestCreateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
