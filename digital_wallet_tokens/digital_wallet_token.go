package digital_wallet_tokens

type DigitalWalletToken struct {
	// The Digital Wallet Token identifier.
	Id *string `json:"id,omitempty"`

	// The identifier for the Card this Digital Wallet Token belongs to.
	CardId *string `json:"card_id,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// This indicates if payments can be made with the Digital Wallet Token.
	Status *string `json:"status,omitempty"`

	// The digital wallet app being used.
	TokenRequestor *string `json:"token_requestor,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `digital_wallet_token`.
	Type *string `json:"type,omitempty"`
}

// The Digital Wallet Token identifier.
func (r *DigitalWalletToken) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The identifier for the Card this Digital Wallet Token belongs to.
func (r *DigitalWalletToken) GetCardId() (CardId string) {
	if r != nil && r.CardId != nil {
		CardId = *r.CardId
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card was created.
func (r *DigitalWalletToken) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// This indicates if payments can be made with the Digital Wallet Token.
func (r *DigitalWalletToken) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// The digital wallet app being used.
func (r *DigitalWalletToken) GetTokenRequestor() (TokenRequestor string) {
	if r != nil && r.TokenRequestor != nil {
		TokenRequestor = *r.TokenRequestor
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `digital_wallet_token`.
func (r *DigitalWalletToken) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
