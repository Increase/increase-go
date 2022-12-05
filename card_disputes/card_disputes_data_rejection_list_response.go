package card_disputes

type CardDisputesDataRejectionListResponse struct {
	// Why the Card Dispute was rejected.
	Explanation *string `json:"explanation,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was rejected.
	RejectedAt *string `json:"rejected_at,omitempty"`

	// The identifier of the Card Dispute that was rejected.
	CardDisputeId *string `json:"card_dispute_id,omitempty"`
}

// Why the Card Dispute was rejected.
func (r *CardDisputesDataRejectionListResponse) GetExplanation() (Explanation string) {
	if r != nil && r.Explanation != nil {
		Explanation = *r.Explanation
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card Dispute was rejected.
func (r *CardDisputesDataRejectionListResponse) GetRejectedAt() (RejectedAt string) {
	if r != nil && r.RejectedAt != nil {
		RejectedAt = *r.RejectedAt
	}
	return
}

// The identifier of the Card Dispute that was rejected.
func (r *CardDisputesDataRejectionListResponse) GetCardDisputeId() (CardDisputeId string) {
	if r != nil && r.CardDisputeId != nil {
		CardDisputeId = *r.CardDisputeId
	}
	return
}
