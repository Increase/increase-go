package transactions

type TransactionsSourceCardRefundRetrieveResponse struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *string `json:"currency,omitempty"`

	// The identifier for the Transaction this refunds, if any.
	CardSettlementTransactionId *string `json:"card_settlement_transaction_id,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type *string `json:"type,omitempty"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *TransactionsSourceCardRefundRetrieveResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *TransactionsSourceCardRefundRetrieveResponse) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier for the Transaction this refunds, if any.
func (r *TransactionsSourceCardRefundRetrieveResponse) GetCardSettlementTransactionId() (CardSettlementTransactionId string) {
	if r != nil && r.CardSettlementTransactionId != nil {
		CardSettlementTransactionId = *r.CardSettlementTransactionId
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_refund`.
func (r *TransactionsSourceCardRefundRetrieveResponse) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
