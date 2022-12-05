package cards

type CardDetails struct {
	// The identifier for the Card for which sensitive details have been returned.
	CardId *string `json:"card_id,omitempty"`

	// The card number.
	PrimaryAccountNumber *string `json:"primary_account_number,omitempty"`

	// The month the card expires in MM format (e.g., August is 08).
	ExpirationMonth *string `json:"expiration_month,omitempty"`

	// The year the card expires in YYYY format (e.g., 2025).
	ExpirationYear *string `json:"expiration_year,omitempty"`

	// The three-digit verification code for the card. It's also known as the Card
	// Verification Code (CVC), the Card Verification Value (CVV), or the Card
	// Identification (CID).
	VerificationCode *string `json:"verification_code,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `card_details`.
	Type *string `json:"type,omitempty"`
}

// The identifier for the Card for which sensitive details have been returned.
func (r *CardDetails) GetCardId() (CardId string) {
	if r != nil && r.CardId != nil {
		CardId = *r.CardId
	}
	return
}

// The card number.
func (r *CardDetails) GetPrimaryAccountNumber() (PrimaryAccountNumber string) {
	if r != nil && r.PrimaryAccountNumber != nil {
		PrimaryAccountNumber = *r.PrimaryAccountNumber
	}
	return
}

// The month the card expires in MM format (e.g., August is 08).
func (r *CardDetails) GetExpirationMonth() (ExpirationMonth string) {
	if r != nil && r.ExpirationMonth != nil {
		ExpirationMonth = *r.ExpirationMonth
	}
	return
}

// The year the card expires in YYYY format (e.g., 2025).
func (r *CardDetails) GetExpirationYear() (ExpirationYear string) {
	if r != nil && r.ExpirationYear != nil {
		ExpirationYear = *r.ExpirationYear
	}
	return
}

// The three-digit verification code for the card. It's also known as the Card
// Verification Code (CVC), the Card Verification Value (CVV), or the Card
// Identification (CID).
func (r *CardDetails) GetVerificationCode() (VerificationCode string) {
	if r != nil && r.VerificationCode != nil {
		VerificationCode = *r.VerificationCode
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_details`.
func (r *CardDetails) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
