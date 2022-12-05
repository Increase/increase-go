package simulations

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferId *string `json:"transfer_id,omitempty"`
}

// The identifier of the ACH Transfer that led to this Transaction.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
