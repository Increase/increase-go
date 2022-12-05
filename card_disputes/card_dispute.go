package card_disputes

type CardDispute struct {
	// The Card Dispute identifier.
	Id *string `json:"id,omitempty"`

	// Why you disputed the Transaction in question.
	Explanation *string `json:"explanation,omitempty"`

	// The results of the Dispute investigation.
	Status *string `json:"status,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The identifier of the Transaction that was disputed.
	DisputedTransactionId *string `json:"disputed_transaction_id,omitempty"`

	// If the Card Dispute's status is `accepted`, this will contain details of the
	// successful dispute.
	Acceptance *CardDisputeAcceptance `json:"acceptance,omitempty"`

	// If the Card Dispute's status is `rejected`, this will contain details of the
	// unsuccessful dispute.
	Rejection *CardDisputeRejection `json:"rejection,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `card_dispute`.
	Type *string `json:"type,omitempty"`
}

// The Card Dispute identifier.
func (r *CardDispute) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// Why you disputed the Transaction in question.
func (r *CardDispute) GetExplanation() (Explanation string) {
	if r != nil && r.Explanation != nil {
		Explanation = *r.Explanation
	}
	return
}

// The results of the Dispute investigation.
func (r *CardDispute) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card Dispute was created.
func (r *CardDispute) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The identifier of the Transaction that was disputed.
func (r *CardDispute) GetDisputedTransactionId() (DisputedTransactionId string) {
	if r != nil && r.DisputedTransactionId != nil {
		DisputedTransactionId = *r.DisputedTransactionId
	}
	return
}

// If the Card Dispute's status is `accepted`, this will contain details of the
// successful dispute.
func (r *CardDispute) GetAcceptance() (Acceptance CardDisputeAcceptance) {
	if r != nil && r.Acceptance != nil {
		Acceptance = *r.Acceptance
	}
	return
}

// If the Card Dispute's status is `rejected`, this will contain details of the
// unsuccessful dispute.
func (r *CardDispute) GetRejection() (Rejection CardDisputeRejection) {
	if r != nil && r.Rejection != nil {
		Rejection = *r.Rejection
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_dispute`.
func (r *CardDispute) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
