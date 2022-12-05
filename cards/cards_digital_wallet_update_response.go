package cards

type CardsDigitalWalletUpdateResponse struct {
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email *string `json:"email,omitempty"`

	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone *string `json:"phone,omitempty"`

	// The card profile assigned to this digital card. Card profiles may also be
	// assigned at the program level.
	CardProfileId *string `json:"card_profile_id,omitempty"`
}

// An email address that can be used to verify the cardholder via one-time passcode
// over email.
func (r *CardsDigitalWalletUpdateResponse) GetEmail() (Email string) {
	if r != nil && r.Email != nil {
		Email = *r.Email
	}
	return
}

// A phone number that can be used to verify the cardholder via one-time passcode
// over SMS.
func (r *CardsDigitalWalletUpdateResponse) GetPhone() (Phone string) {
	if r != nil && r.Phone != nil {
		Phone = *r.Phone
	}
	return
}

// The card profile assigned to this digital card. Card profiles may also be
// assigned at the program level.
func (r *CardsDigitalWalletUpdateResponse) GetCardProfileId() (CardProfileId string) {
	if r != nil && r.CardProfileId != nil {
		CardProfileId = *r.CardProfileId
	}
	return
}
