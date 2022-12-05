package pending_transactions

type PendingTransactionsSourceCheckTransferInstructionRetrieveResponse struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency *string `json:"currency,omitempty"`

	// The identifier of the Check Transfer that led to this Pending Transaction.
	TransferId *string `json:"transfer_id,omitempty"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionsSourceCheckTransferInstructionRetrieveResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
func (r *PendingTransactionsSourceCheckTransferInstructionRetrieveResponse) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Check Transfer that led to this Pending Transaction.
func (r *PendingTransactionsSourceCheckTransferInstructionRetrieveResponse) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
