package simulations

type CheckDepositsDepositAcceptanceSubmitResponse struct {
	// The amount to be deposited in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *string `json:"currency,omitempty"`

	// The account number printed on the check.
	AccountNumber *string `json:"account_number,omitempty"`

	// The routing number printed on the check.
	RoutingNumber *string `json:"routing_number,omitempty"`

	// An additional line of metadata printed on the check. This typically includes the
	// check number.
	AuxiliaryOnUs *string `json:"auxiliary_on_us,omitempty"`
}

// The amount to be deposited in the minor unit of the transaction's currency. For
// dollars, for example, this is cents.
func (r *CheckDepositsDepositAcceptanceSubmitResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *CheckDepositsDepositAcceptanceSubmitResponse) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The account number printed on the check.
func (r *CheckDepositsDepositAcceptanceSubmitResponse) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The routing number printed on the check.
func (r *CheckDepositsDepositAcceptanceSubmitResponse) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// An additional line of metadata printed on the check. This typically includes the
// check number.
func (r *CheckDepositsDepositAcceptanceSubmitResponse) GetAuxiliaryOnUs() (AuxiliaryOnUs string) {
	if r != nil && r.AuxiliaryOnUs != nil {
		AuxiliaryOnUs = *r.AuxiliaryOnUs
	}
	return
}
