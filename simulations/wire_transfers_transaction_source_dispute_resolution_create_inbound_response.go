package simulations

type WireTransfersTransactionSourceDisputeResolutionCreateInboundResponse struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *string `json:"currency,omitempty"`

	// The identifier of the Transaction that was disputed.
	DisputedTransactionId *string `json:"disputed_transaction_id,omitempty"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *WireTransfersTransactionSourceDisputeResolutionCreateInboundResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *WireTransfersTransactionSourceDisputeResolutionCreateInboundResponse) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Transaction that was disputed.
func (r *WireTransfersTransactionSourceDisputeResolutionCreateInboundResponse) GetDisputedTransactionId() (DisputedTransactionId string) {
	if r != nil && r.DisputedTransactionId != nil {
		DisputedTransactionId = *r.DisputedTransactionId
	}
	return
}
