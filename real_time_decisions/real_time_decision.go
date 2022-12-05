package real_time_decisions

type RealTimeDecision struct {
	// The Real-Time Decision identifier.
	Id *string `json:"id,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Real-Time Decision was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// your application can no longer respond to the Real-Time Decision.
	TimeoutAt *string `json:"timeout_at,omitempty"`

	// The status of the Real-Time Decision.
	Status *string `json:"status,omitempty"`

	// The category of the Real-Time Decision.
	Category *string `json:"category,omitempty"`

	// Fields related to a card authorization.
	CardAuthorization *RealTimeDecisionCardAuthorization `json:"card_authorization,omitempty"`

	// Fields related to a digital wallet token provisioning attempt.
	DigitalWalletToken *RealTimeDecisionDigitalWalletToken `json:"digital_wallet_token,omitempty"`

	// Fields related to a digital wallet authentication attempt.
	DigitalWalletAuthentication *RealTimeDecisionDigitalWalletAuthentication `json:"digital_wallet_authentication,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `real_time_decision`.
	Type *string `json:"type,omitempty"`
}

// The Real-Time Decision identifier.
func (r *RealTimeDecision) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Real-Time Decision was created.
func (r *RealTimeDecision) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// your application can no longer respond to the Real-Time Decision.
func (r *RealTimeDecision) GetTimeoutAt() (TimeoutAt string) {
	if r != nil && r.TimeoutAt != nil {
		TimeoutAt = *r.TimeoutAt
	}
	return
}

// The status of the Real-Time Decision.
func (r *RealTimeDecision) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// The category of the Real-Time Decision.
func (r *RealTimeDecision) GetCategory() (Category string) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// Fields related to a card authorization.
func (r *RealTimeDecision) GetCardAuthorization() (CardAuthorization RealTimeDecisionCardAuthorization) {
	if r != nil && r.CardAuthorization != nil {
		CardAuthorization = *r.CardAuthorization
	}
	return
}

// Fields related to a digital wallet token provisioning attempt.
func (r *RealTimeDecision) GetDigitalWalletToken() (DigitalWalletToken RealTimeDecisionDigitalWalletToken) {
	if r != nil && r.DigitalWalletToken != nil {
		DigitalWalletToken = *r.DigitalWalletToken
	}
	return
}

// Fields related to a digital wallet authentication attempt.
func (r *RealTimeDecision) GetDigitalWalletAuthentication() (DigitalWalletAuthentication RealTimeDecisionDigitalWalletAuthentication) {
	if r != nil && r.DigitalWalletAuthentication != nil {
		DigitalWalletAuthentication = *r.DigitalWalletAuthentication
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `real_time_decision`.
func (r *RealTimeDecision) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
