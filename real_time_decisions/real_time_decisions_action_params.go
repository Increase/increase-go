package real_time_decisions

type RealTimeDecisionsActionParams struct {
	// If the Real-Time Decision relates to a card authorization attempt, this object
	// contains your response to the authorization.
	CardAuthorization *RealTimeDecisionsCardAuthorizationActionParams `json:"card_authorization,omitempty"`

	// If the Real-Time Decision relates to a digital wallet token provisioning
	// attempt, this object contains your response to the attempt.
	DigitalWalletToken *RealTimeDecisionsDigitalWalletTokenActionParams `json:"digital_wallet_token,omitempty"`

	// If the Real-Time Decision relates to a digital wallet authentication attempt,
	// this object contains your response to the authentication.
	DigitalWalletAuthentication *RealTimeDecisionsDigitalWalletAuthenticationActionParams `json:"digital_wallet_authentication,omitempty"`
}

// If the Real-Time Decision relates to a card authorization attempt, this object
// contains your response to the authorization.
func (r *RealTimeDecisionsActionParams) GetCardAuthorization() (CardAuthorization RealTimeDecisionsCardAuthorizationActionParams) {
	if r != nil && r.CardAuthorization != nil {
		CardAuthorization = *r.CardAuthorization
	}
	return
}

// If the Real-Time Decision relates to a digital wallet token provisioning
// attempt, this object contains your response to the attempt.
func (r *RealTimeDecisionsActionParams) GetDigitalWalletToken() (DigitalWalletToken RealTimeDecisionsDigitalWalletTokenActionParams) {
	if r != nil && r.DigitalWalletToken != nil {
		DigitalWalletToken = *r.DigitalWalletToken
	}
	return
}

// If the Real-Time Decision relates to a digital wallet authentication attempt,
// this object contains your response to the authentication.
func (r *RealTimeDecisionsActionParams) GetDigitalWalletAuthentication() (DigitalWalletAuthentication RealTimeDecisionsDigitalWalletAuthenticationActionParams) {
	if r != nil && r.DigitalWalletAuthentication != nil {
		DigitalWalletAuthentication = *r.DigitalWalletAuthentication
	}
	return
}
