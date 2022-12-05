package simulations

type DigitalWalletTokenRequestsCreateParams struct {
	// The identifier of the Card to be authorized.
	CardId *string `json:"card_id,omitempty"`
}

// The identifier of the Card to be authorized.
func (r *DigitalWalletTokenRequestsCreateParams) GetCardId() (CardId string) {
	if r != nil && r.CardId != nil {
		CardId = *r.CardId
	}
	return
}
