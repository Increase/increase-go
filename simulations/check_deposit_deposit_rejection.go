package simulations

type CheckDepositDepositRejection struct {
	// The rejected amount in the minor unit of check's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency *string `json:"currency,omitempty"`

	// Why the check deposit was rejected.
	Reason *string `json:"reason,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was rejected.
	RejectedAt *string `json:"rejected_at,omitempty"`
}

// The rejected amount in the minor unit of check's currency. For dollars, for
// example, this is cents.
func (r *CheckDepositDepositRejection) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
func (r *CheckDepositDepositRejection) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// Why the check deposit was rejected.
func (r *CheckDepositDepositRejection) GetReason() (Reason string) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the check deposit was rejected.
func (r *CheckDepositDepositRejection) GetRejectedAt() (RejectedAt string) {
	if r != nil && r.RejectedAt != nil {
		RejectedAt = *r.RejectedAt
	}
	return
}
