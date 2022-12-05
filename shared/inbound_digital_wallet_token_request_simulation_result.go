package shared

type InboundDigitalWalletTokenRequestSimulationResult struct {
	// If the simulated tokenization attempt was declined, this field contains details
	// as to why.
	DeclineReason *string `json:"decline_reason,omitempty"`

	// If the simulated tokenization attempt was accepted, this field contains the id
	// of the Digital Wallet Token that was created.
	DigitalWalletTokenId *string `json:"digital_wallet_token_id,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `inbound_digital_wallet_token_request_simulation_result`.
	Type *string `json:"type,omitempty"`
}

// If the simulated tokenization attempt was declined, this field contains details
// as to why.
func (r *InboundDigitalWalletTokenRequestSimulationResult) GetDeclineReason() (DeclineReason string) {
	if r != nil && r.DeclineReason != nil {
		DeclineReason = *r.DeclineReason
	}
	return
}

// If the simulated tokenization attempt was accepted, this field contains the id
// of the Digital Wallet Token that was created.
func (r *InboundDigitalWalletTokenRequestSimulationResult) GetDigitalWalletTokenId() (DigitalWalletTokenId string) {
	if r != nil && r.DigitalWalletTokenId != nil {
		DigitalWalletTokenId = *r.DigitalWalletTokenId
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `inbound_digital_wallet_token_request_simulation_result`.
func (r *InboundDigitalWalletTokenRequestSimulationResult) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
