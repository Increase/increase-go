package cards

type CardsCreateParams struct {
	// The Account the card should belong to.
	AccountId *string `json:"account_id,omitempty"`

	// The description you choose to give the card.
	Description *string `json:"description,omitempty"`

	// The card's billing address.
	BillingAddress *CardsBillingAddressCreateParams `json:"billing_address,omitempty"`

	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet *CardsDigitalWalletCreateParams `json:"digital_wallet,omitempty"`
}

// The Account the card should belong to.
func (r *CardsCreateParams) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The description you choose to give the card.
func (r *CardsCreateParams) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The card's billing address.
func (r *CardsCreateParams) GetBillingAddress() (BillingAddress CardsBillingAddressCreateParams) {
	if r != nil && r.BillingAddress != nil {
		BillingAddress = *r.BillingAddress
	}
	return
}

// The contact information used in the two-factor steps for digital wallet card
// creation. At least one field must be present to complete the digital wallet
// steps.
func (r *CardsCreateParams) GetDigitalWallet() (DigitalWallet CardsDigitalWalletCreateParams) {
	if r != nil && r.DigitalWallet != nil {
		DigitalWallet = *r.DigitalWallet
	}
	return
}
