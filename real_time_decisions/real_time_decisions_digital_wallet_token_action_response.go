package real_time_decisions

type RealTimeDecisionsDigitalWalletTokenActionResponse struct {
	// Whether or not the provisioning request was approved.
	Decision *string `json:"decision,omitempty"`

	// The identifier of the Card that is being tokenized.
	CardId *string `json:"card_id,omitempty"`

	// The digital wallet app being used.
	DigitalWallet *string `json:"digital_wallet,omitempty"`
}

// Whether or not the provisioning request was approved.
func (r *RealTimeDecisionsDigitalWalletTokenActionResponse) GetDecision() (Decision string) {
	if r != nil && r.Decision != nil {
		Decision = *r.Decision
	}
	return
}

// The identifier of the Card that is being tokenized.
func (r *RealTimeDecisionsDigitalWalletTokenActionResponse) GetCardId() (CardId string) {
	if r != nil && r.CardId != nil {
		CardId = *r.CardId
	}
	return
}

// The digital wallet app being used.
func (r *RealTimeDecisionsDigitalWalletTokenActionResponse) GetDigitalWallet() (DigitalWallet string) {
	if r != nil && r.DigitalWallet != nil {
		DigitalWallet = *r.DigitalWallet
	}
	return
}
