package real_time_decisions

type RealTimeDecisionsDigitalWalletTokenApprovalActionParams struct {
	// The identifier of the Card Profile to assign to the Digital Wallet token.
	CardProfileId *string `json:"card_profile_id,omitempty"`

	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone *string `json:"phone,omitempty"`

	// An email address that can be used to verify the cardholder via one-time
	// passcode.
	Email *string `json:"email,omitempty"`
}

// The identifier of the Card Profile to assign to the Digital Wallet token.
func (r *RealTimeDecisionsDigitalWalletTokenApprovalActionParams) GetCardProfileId() (CardProfileId string) {
	if r != nil && r.CardProfileId != nil {
		CardProfileId = *r.CardProfileId
	}
	return
}

// A phone number that can be used to verify the cardholder via one-time passcode
// over SMS.
func (r *RealTimeDecisionsDigitalWalletTokenApprovalActionParams) GetPhone() (Phone string) {
	if r != nil && r.Phone != nil {
		Phone = *r.Phone
	}
	return
}

// An email address that can be used to verify the cardholder via one-time
// passcode.
func (r *RealTimeDecisionsDigitalWalletTokenApprovalActionParams) GetEmail() (Email string) {
	if r != nil && r.Email != nil {
		Email = *r.Email
	}
	return
}
