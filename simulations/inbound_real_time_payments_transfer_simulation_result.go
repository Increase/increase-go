package simulations

type InboundRealTimePaymentsTransferSimulationResult struct {
	// If the Real Time Payments Transfer attempt succeeds, this will contain the
	// resulting [Transaction](#transactions) object. The Transaction's `source` will
	// be of `category: inbound_real_time_payments_transfer_confirmation`.
	Transaction *InboundRealTimePaymentsTransferSimulationResultTransaction `json:"transaction,omitempty"`

	// If the Real Time Payments Transfer attempt fails, this will contain the
	// resulting [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of
	// `category: inbound_real_time_payments_transfer_decline`.
	DeclinedTransaction *InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction `json:"declined_transaction,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `inbound_real_time_payments_transfer_simulation_result`.
	Type *string `json:"type,omitempty"`
}

// If the Real Time Payments Transfer attempt succeeds, this will contain the
// resulting [Transaction](#transactions) object. The Transaction's `source` will
// be of `category: inbound_real_time_payments_transfer_confirmation`.
func (r *InboundRealTimePaymentsTransferSimulationResult) GetTransaction() (Transaction InboundRealTimePaymentsTransferSimulationResultTransaction) {
	if r != nil && r.Transaction != nil {
		Transaction = *r.Transaction
	}
	return
}

// If the Real Time Payments Transfer attempt fails, this will contain the
// resulting [Declined Transaction](#declined-transactions) object. The Declined
// Transaction's `source` will be of
// `category: inbound_real_time_payments_transfer_decline`.
func (r *InboundRealTimePaymentsTransferSimulationResult) GetDeclinedTransaction() (DeclinedTransaction InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction) {
	if r != nil && r.DeclinedTransaction != nil {
		DeclinedTransaction = *r.DeclinedTransaction
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `inbound_real_time_payments_transfer_simulation_result`.
func (r *InboundRealTimePaymentsTransferSimulationResult) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
