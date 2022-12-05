package cards

type Card struct {
	// The card identifier.
	Id *string `json:"id,omitempty"`

	// The identifier for the account this card belongs to.
	AccountId *string `json:"account_id,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The card's description for display purposes.
	Description *string `json:"description,omitempty"`

	// The last 4 digits of the Card's Primary Account Number.
	Last4 *string `json:"last4,omitempty"`

	// The month the card expires in MM format (e.g., August is 08).
	ExpirationMonth *string `json:"expiration_month,omitempty"`

	// The year the card expires in YYYY format (e.g., 2025).
	ExpirationYear *string `json:"expiration_year,omitempty"`

	// This indicates if payments can be made with the card.
	Status *string `json:"status,omitempty"`

	// The Card's billing address.
	BillingAddress *CardBillingAddress `json:"billing_address,omitempty"`

	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet *CardDigitalWallet `json:"digital_wallet,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `card`.
	Type *string `json:"type,omitempty"`
}

// The card identifier.
func (r *Card) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The identifier for the account this card belongs to.
func (r *Card) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card was created.
func (r *Card) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The card's description for display purposes.
func (r *Card) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The last 4 digits of the Card's Primary Account Number.
func (r *Card) GetLast4() (Last4 string) {
	if r != nil && r.Last4 != nil {
		Last4 = *r.Last4
	}
	return
}

// The month the card expires in MM format (e.g., August is 08).
func (r *Card) GetExpirationMonth() (ExpirationMonth string) {
	if r != nil && r.ExpirationMonth != nil {
		ExpirationMonth = *r.ExpirationMonth
	}
	return
}

// The year the card expires in YYYY format (e.g., 2025).
func (r *Card) GetExpirationYear() (ExpirationYear string) {
	if r != nil && r.ExpirationYear != nil {
		ExpirationYear = *r.ExpirationYear
	}
	return
}

// This indicates if payments can be made with the card.
func (r *Card) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// The Card's billing address.
func (r *Card) GetBillingAddress() (BillingAddress CardBillingAddress) {
	if r != nil && r.BillingAddress != nil {
		BillingAddress = *r.BillingAddress
	}
	return
}

// The contact information used in the two-factor steps for digital wallet card
// creation. At least one field must be present to complete the digital wallet
// steps.
func (r *Card) GetDigitalWallet() (DigitalWallet CardDigitalWallet) {
	if r != nil && r.DigitalWallet != nil {
		DigitalWallet = *r.DigitalWallet
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card`.
func (r *Card) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
