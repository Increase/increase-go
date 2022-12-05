package card_profiles

type CardProfile struct {
	// The Card Profile identifier.
	Id *string `json:"id,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The status of the Card Profile.
	Status *string `json:"status,omitempty"`

	// A description you can use to identify the Card Profile.
	Description *string `json:"description,omitempty"`

	// How Cards should appear in digital wallets such as Apple Pay. Different wallets
	// will use these values to render card artwork appropriately for their app.
	DigitalWallets *CardProfileDigitalWallet `json:"digital_wallets,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `card_profile`.
	Type *string `json:"type,omitempty"`
}

// The Card Profile identifier.
func (r *CardProfile) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card Dispute was created.
func (r *CardProfile) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The status of the Card Profile.
func (r *CardProfile) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// A description you can use to identify the Card Profile.
func (r *CardProfile) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// How Cards should appear in digital wallets such as Apple Pay. Different wallets
// will use these values to render card artwork appropriately for their app.
func (r *CardProfile) GetDigitalWallets() (DigitalWallets CardProfileDigitalWallet) {
	if r != nil && r.DigitalWallets != nil {
		DigitalWallets = *r.DigitalWallets
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_profile`.
func (r *CardProfile) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
