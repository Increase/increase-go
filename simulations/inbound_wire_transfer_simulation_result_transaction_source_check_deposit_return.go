package simulations

type InboundWireTransferSimulationResultTransactionSourceCheckDepositReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt *string `json:"returned_at,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *string `json:"currency,omitempty"`

	// The identifier of the Check Deposit that was returned.
	CheckDepositId *string `json:"check_deposit_id,omitempty"`

	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionId *string `json:"transaction_id,omitempty"`

	ReturnReason *string `json:"return_reason,omitempty"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *InboundWireTransferSimulationResultTransactionSourceCheckDepositReturn) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the check deposit was returned.
func (r *InboundWireTransferSimulationResultTransactionSourceCheckDepositReturn) GetReturnedAt() (ReturnedAt string) {
	if r != nil && r.ReturnedAt != nil {
		ReturnedAt = *r.ReturnedAt
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *InboundWireTransferSimulationResultTransactionSourceCheckDepositReturn) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Check Deposit that was returned.
func (r *InboundWireTransferSimulationResultTransactionSourceCheckDepositReturn) GetCheckDepositId() (CheckDepositId string) {
	if r != nil && r.CheckDepositId != nil {
		CheckDepositId = *r.CheckDepositId
	}
	return
}

// The identifier of the transaction that reversed the original check deposit
// transaction.
func (r *InboundWireTransferSimulationResultTransactionSourceCheckDepositReturn) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

func (r *InboundWireTransferSimulationResultTransactionSourceCheckDepositReturn) GetReturnReason() (ReturnReason string) {
	if r != nil && r.ReturnReason != nil {
		ReturnReason = *r.ReturnReason
	}
	return
}
