package cards

type CardsUpdateParams struct {
	// The description you choose to give the card.
	Description *string `json:"description,omitempty"`

	// The status to update the Card with.
	Status *string `json:"status,omitempty"`

	// The card's updated billing address.
	BillingAddress *CardsBillingAddressUpdateParams `json:"billing_address,omitempty"`

	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet *CardsDigitalWalletUpdateParams `json:"digital_wallet,omitempty"`
}

// The description you choose to give the card.
func (r *CardsUpdateParams) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The status to update the Card with.
func (r *CardsUpdateParams) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// The card's updated billing address.
func (r *CardsUpdateParams) GetBillingAddress() (BillingAddress CardsBillingAddressUpdateParams) {
	if r != nil && r.BillingAddress != nil {
		BillingAddress = *r.BillingAddress
	}
	return
}

// The contact information used in the two-factor steps for digital wallet card
// creation. At least one field must be present to complete the digital wallet
// steps.
func (r *CardsUpdateParams) GetDigitalWallet() (DigitalWallet CardsDigitalWalletUpdateParams) {
	if r != nil && r.DigitalWallet != nil {
		DigitalWallet = *r.DigitalWallet
	}
	return
}
