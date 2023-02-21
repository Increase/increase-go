package types

import (
	"fmt"
	"increase/core"
	"increase/core/pjson"
)

type DigitalWalletTokenRequestCreateResponse struct {
	// If the simulated tokenization attempt was declined, this field contains details
	// as to why.
	DeclineReason *DigitalWalletTokenRequestCreateResponseDeclineReason `pjson:"decline_reason"`
	// If the simulated tokenization attempt was accepted, this field contains the id
	// of the Digital Wallet Token that was created.
	DigitalWalletTokenID *string `pjson:"digital_wallet_token_id"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_digital_wallet_token_request_simulation_result`.
	Type       *DigitalWalletTokenRequestCreateResponseType `pjson:"type"`
	jsonFields map[string]interface{}                       `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// DigitalWalletTokenRequestCreateResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *DigitalWalletTokenRequestCreateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DigitalWalletTokenRequestCreateResponse into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *DigitalWalletTokenRequestCreateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// If the simulated tokenization attempt was declined, this field contains details
// as to why.
func (r *DigitalWalletTokenRequestCreateResponse) GetDeclineReason() (DeclineReason DigitalWalletTokenRequestCreateResponseDeclineReason) {
	if r != nil && r.DeclineReason != nil {
		DeclineReason = *r.DeclineReason
	}
	return
}

// If the simulated tokenization attempt was accepted, this field contains the id
// of the Digital Wallet Token that was created.
func (r *DigitalWalletTokenRequestCreateResponse) GetDigitalWalletTokenID() (DigitalWalletTokenID string) {
	if r != nil && r.DigitalWalletTokenID != nil {
		DigitalWalletTokenID = *r.DigitalWalletTokenID
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `inbound_digital_wallet_token_request_simulation_result`.
func (r *DigitalWalletTokenRequestCreateResponse) GetType() (Type DigitalWalletTokenRequestCreateResponseType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r DigitalWalletTokenRequestCreateResponse) String() (result string) {
	return fmt.Sprintf("&DigitalWalletTokenRequestCreateResponse{DeclineReason:%s DigitalWalletTokenID:%s Type:%s}", core.FmtP(r.DeclineReason), core.FmtP(r.DigitalWalletTokenID), core.FmtP(r.Type))
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
	CardID     *string                `pjson:"card_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// SimulateDigitalWalletProvisioningForACardParameters using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *SimulateDigitalWalletProvisioningForACardParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes SimulateDigitalWalletProvisioningForACardParameters into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *SimulateDigitalWalletProvisioningForACardParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the Card to be authorized.
func (r *SimulateDigitalWalletProvisioningForACardParameters) GetCardID() (CardID string) {
	if r != nil && r.CardID != nil {
		CardID = *r.CardID
	}
	return
}

func (r SimulateDigitalWalletProvisioningForACardParameters) String() (result string) {
	return fmt.Sprintf("&SimulateDigitalWalletProvisioningForACardParameters{CardID:%s}", core.FmtP(r.CardID))
}
