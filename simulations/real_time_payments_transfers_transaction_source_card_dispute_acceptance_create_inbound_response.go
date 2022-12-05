package simulations

type RealTimePaymentsTransfersTransactionSourceCardDisputeAcceptanceCreateInboundResponse struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was accepted.
	AcceptedAt *string `json:"accepted_at,omitempty"`

	// The identifier of the Card Dispute that was accepted.
	CardDisputeId *string `json:"card_dispute_id,omitempty"`

	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionId *string `json:"transaction_id,omitempty"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card Dispute was accepted.
func (r *RealTimePaymentsTransfersTransactionSourceCardDisputeAcceptanceCreateInboundResponse) GetAcceptedAt() (AcceptedAt string) {
	if r != nil && r.AcceptedAt != nil {
		AcceptedAt = *r.AcceptedAt
	}
	return
}

// The identifier of the Card Dispute that was accepted.
func (r *RealTimePaymentsTransfersTransactionSourceCardDisputeAcceptanceCreateInboundResponse) GetCardDisputeId() (CardDisputeId string) {
	if r != nil && r.CardDisputeId != nil {
		CardDisputeId = *r.CardDisputeId
	}
	return
}

// The identifier of the Transaction that was created to return the disputed funds
// to your account.
func (r *RealTimePaymentsTransfersTransactionSourceCardDisputeAcceptanceCreateInboundResponse) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}
