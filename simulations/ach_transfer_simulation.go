package simulations

type ACHTransferSimulation struct {
	// If the ACH Transfer attempt succeeds, this will contain the resulting
	// [Transaction](#transactions) object. The Transaction's `source` will be of
	// `category: inbound_ach_transfer`.
	Transaction *ACHTransferSimulationTransaction `json:"transaction,omitempty"`

	// If the ACH Transfer attempt fails, this will contain the resulting
	// [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of `category: inbound_ach_transfer`.
	DeclinedTransaction *ACHTransferSimulationDeclinedTransaction `json:"declined_transaction,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `inbound_ach_transfer_simulation_result`.
	Type *string `json:"type,omitempty"`
}

// If the ACH Transfer attempt succeeds, this will contain the resulting
// [Transaction](#transactions) object. The Transaction's `source` will be of
// `category: inbound_ach_transfer`.
func (r *ACHTransferSimulation) GetTransaction() (Transaction ACHTransferSimulationTransaction) {
	if r != nil && r.Transaction != nil {
		Transaction = *r.Transaction
	}
	return
}

// If the ACH Transfer attempt fails, this will contain the resulting
// [Declined Transaction](#declined-transactions) object. The Declined
// Transaction's `source` will be of `category: inbound_ach_transfer`.
func (r *ACHTransferSimulation) GetDeclinedTransaction() (DeclinedTransaction ACHTransferSimulationDeclinedTransaction) {
	if r != nil && r.DeclinedTransaction != nil {
		DeclinedTransaction = *r.DeclinedTransaction
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `inbound_ach_transfer_simulation_result`.
func (r *ACHTransferSimulation) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
