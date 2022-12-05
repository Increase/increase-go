package simulations

type CardAuthorizationSimulation struct {
	// If the authorization attempt succeeds, this will contain the resulting Pending
	// Transaction object. The Pending Transaction's `source` will be of
	// `category: card_authorization`.
	PendingTransaction *CardAuthorizationSimulationPendingTransaction `json:"pending_transaction,omitempty"`

	// If the authorization attempt fails, this will contain the resulting
	// [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of `category: card_decline`.
	DeclinedTransaction *CardAuthorizationSimulationDeclinedTransaction `json:"declined_transaction,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `inbound_card_authorization_simulation_result`.
	Type *string `json:"type,omitempty"`
}

// If the authorization attempt succeeds, this will contain the resulting Pending
// Transaction object. The Pending Transaction's `source` will be of
// `category: card_authorization`.
func (r *CardAuthorizationSimulation) GetPendingTransaction() (PendingTransaction CardAuthorizationSimulationPendingTransaction) {
	if r != nil && r.PendingTransaction != nil {
		PendingTransaction = *r.PendingTransaction
	}
	return
}

// If the authorization attempt fails, this will contain the resulting
// [Declined Transaction](#declined-transactions) object. The Declined
// Transaction's `source` will be of `category: card_decline`.
func (r *CardAuthorizationSimulation) GetDeclinedTransaction() (DeclinedTransaction CardAuthorizationSimulationDeclinedTransaction) {
	if r != nil && r.DeclinedTransaction != nil {
		DeclinedTransaction = *r.DeclinedTransaction
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `inbound_card_authorization_simulation_result`.
func (r *CardAuthorizationSimulation) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
