package simulations

type CardsAuthorizeParams struct {
	// The authorization amount in cents.
	Amount *int64 `json:"amount,omitempty"`

	// The identifier of the Card to be authorized.
	CardId *string `json:"card_id,omitempty"`

	// The identifier of the Digital Wallet Token to be authorized.
	DigitalWalletTokenId *string `json:"digital_wallet_token_id,omitempty"`
}

// The authorization amount in cents.
func (r *CardsAuthorizeParams) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The identifier of the Card to be authorized.
func (r *CardsAuthorizeParams) GetCardId() (CardId string) {
	if r != nil && r.CardId != nil {
		CardId = *r.CardId
	}
	return
}

// The identifier of the Digital Wallet Token to be authorized.
func (r *CardsAuthorizeParams) GetDigitalWalletTokenId() (DigitalWalletTokenId string) {
	if r != nil && r.DigitalWalletTokenId != nil {
		DigitalWalletTokenId = *r.DigitalWalletTokenId
	}
	return
}
