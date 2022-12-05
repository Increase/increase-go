package digital_wallet_tokens

type DigitalWalletTokensListQuery struct {
	// Return the page of entries after this one.
	Cursor *string `json:"cursor,omitempty"`

	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `json:"limit,omitempty"`

	// Filter Digital Wallet Tokens to ones belonging to the specified Card.
	CardId    *string                                `json:"card_id,omitempty"`
	CreatedAt *DigitalWalletTokensCreatedAtListQuery `json:"created_at,omitempty"`
}

// Return the page of entries after this one.
func (r *DigitalWalletTokensListQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *DigitalWalletTokensListQuery) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Digital Wallet Tokens to ones belonging to the specified Card.
func (r *DigitalWalletTokensListQuery) GetCardId() (CardId string) {
	if r != nil && r.CardId != nil {
		CardId = *r.CardId
	}
	return
}

func (r *DigitalWalletTokensListQuery) GetCreatedAt() (CreatedAt DigitalWalletTokensCreatedAtListQuery) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}
