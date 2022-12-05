package card_profiles

type CardProfilesCreateParams struct {
	// A description you can use to identify the Card Profile.
	Description *string `json:"description,omitempty"`

	// How Cards should appear in digital wallets such as Apple Pay. Different wallets
	// will use these values to render card artwork appropriately for their app.
	DigitalWallets *CardProfilesDigitalWalletsCreateParams `json:"digital_wallets,omitempty"`
}

// A description you can use to identify the Card Profile.
func (r *CardProfilesCreateParams) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// How Cards should appear in digital wallets such as Apple Pay. Different wallets
// will use these values to render card artwork appropriately for their app.
func (r *CardProfilesCreateParams) GetDigitalWallets() (DigitalWallets CardProfilesDigitalWalletsCreateParams) {
	if r != nil && r.DigitalWallets != nil {
		DigitalWallets = *r.DigitalWallets
	}
	return
}
