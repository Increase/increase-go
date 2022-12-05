package card_disputes

type CardDisputesCreateParams struct {
	// The Transaction you wish to dispute. This Transaction must have a `source_type`
	// of `card_settlement`.
	DisputedTransactionId *string `json:"disputed_transaction_id,omitempty"`

	// Why you are disputing this Transaction.
	Explanation *string `json:"explanation,omitempty"`
}

// The Transaction you wish to dispute. This Transaction must have a `source_type`
// of `card_settlement`.
func (r *CardDisputesCreateParams) GetDisputedTransactionId() (DisputedTransactionId string) {
	if r != nil && r.DisputedTransactionId != nil {
		DisputedTransactionId = *r.DisputedTransactionId
	}
	return
}

// Why you are disputing this Transaction.
func (r *CardDisputesCreateParams) GetExplanation() (Explanation string) {
	if r != nil && r.Explanation != nil {
		Explanation = *r.Explanation
	}
	return
}
