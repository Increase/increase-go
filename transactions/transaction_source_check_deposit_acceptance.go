package transactions

type TransactionSourceCheckDepositAcceptance struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *string `json:"currency,omitempty"`

	// The ID of the Check Deposit that led to the Transaction.
	CheckDepositId *string `json:"check_deposit_id,omitempty"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceCheckDepositAcceptance) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *TransactionSourceCheckDepositAcceptance) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The ID of the Check Deposit that led to the Transaction.
func (r *TransactionSourceCheckDepositAcceptance) GetCheckDepositId() (CheckDepositId string) {
	if r != nil && r.CheckDepositId != nil {
		CheckDepositId = *r.CheckDepositId
	}
	return
}
