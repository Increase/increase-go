package simulations

type CardsSettlementParams struct {
	// The identifier of the Card to create a settlement on.
	CardId *string `json:"card_id,omitempty"`

	// The identifier of the Pending Transaction for the Card Authorization you wish to
	// settle.
	PendingTransactionId *string `json:"pending_transaction_id,omitempty"`

	// The amount to be settled. This defaults to the amount of the Pending Transaction
	// being settled.
	Amount *int64 `json:"amount,omitempty"`
}

// The identifier of the Card to create a settlement on.
func (r *CardsSettlementParams) GetCardId() (CardId string) {
	if r != nil && r.CardId != nil {
		CardId = *r.CardId
	}
	return
}

// The identifier of the Pending Transaction for the Card Authorization you wish to
// settle.
func (r *CardsSettlementParams) GetPendingTransactionId() (PendingTransactionId string) {
	if r != nil && r.PendingTransactionId != nil {
		PendingTransactionId = *r.PendingTransactionId
	}
	return
}

// The amount to be settled. This defaults to the amount of the Pending Transaction
// being settled.
func (r *CardsSettlementParams) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}
