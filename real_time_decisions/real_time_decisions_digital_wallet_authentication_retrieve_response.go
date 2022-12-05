package real_time_decisions

type RealTimeDecisionsDigitalWalletAuthenticationRetrieveResponse struct {
	// Whether your application successfully delivered the one-time passcode.
	Result *string `json:"result,omitempty"`

	// The identifier of the Card that is being tokenized.
	CardId *string `json:"card_id,omitempty"`

	// The digital wallet app being used.
	DigitalWallet *string `json:"digital_wallet,omitempty"`

	// The channel to send the card user their one-time passcode.
	Channel *string `json:"channel,omitempty"`

	// The one-time passcode to send the card user.
	OneTimePasscode *string `json:"one_time_passcode,omitempty"`

	// The phone number to send the one-time passcode to if `channel` is equal to
	// `sms`.
	Phone *string `json:"phone,omitempty"`

	// The email to send the one-time passcode to if `channel` is equal to `email`.
	Email *string `json:"email,omitempty"`
}

// Whether your application successfully delivered the one-time passcode.
func (r *RealTimeDecisionsDigitalWalletAuthenticationRetrieveResponse) GetResult() (Result string) {
	if r != nil && r.Result != nil {
		Result = *r.Result
	}
	return
}

// The identifier of the Card that is being tokenized.
func (r *RealTimeDecisionsDigitalWalletAuthenticationRetrieveResponse) GetCardId() (CardId string) {
	if r != nil && r.CardId != nil {
		CardId = *r.CardId
	}
	return
}

// The digital wallet app being used.
func (r *RealTimeDecisionsDigitalWalletAuthenticationRetrieveResponse) GetDigitalWallet() (DigitalWallet string) {
	if r != nil && r.DigitalWallet != nil {
		DigitalWallet = *r.DigitalWallet
	}
	return
}

// The channel to send the card user their one-time passcode.
func (r *RealTimeDecisionsDigitalWalletAuthenticationRetrieveResponse) GetChannel() (Channel string) {
	if r != nil && r.Channel != nil {
		Channel = *r.Channel
	}
	return
}

// The one-time passcode to send the card user.
func (r *RealTimeDecisionsDigitalWalletAuthenticationRetrieveResponse) GetOneTimePasscode() (OneTimePasscode string) {
	if r != nil && r.OneTimePasscode != nil {
		OneTimePasscode = *r.OneTimePasscode
	}
	return
}

// The phone number to send the one-time passcode to if `channel` is equal to
// `sms`.
func (r *RealTimeDecisionsDigitalWalletAuthenticationRetrieveResponse) GetPhone() (Phone string) {
	if r != nil && r.Phone != nil {
		Phone = *r.Phone
	}
	return
}

// The email to send the one-time passcode to if `channel` is equal to `email`.
func (r *RealTimeDecisionsDigitalWalletAuthenticationRetrieveResponse) GetEmail() (Email string) {
	if r != nil && r.Email != nil {
		Email = *r.Email
	}
	return
}
