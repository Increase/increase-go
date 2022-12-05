package simulations

type InboundACHTransferSimulationResultTransactionSourceCheckTransferRejection struct {
	// The identifier of the Check Transfer that led to this Transaction.
	TransferId *string `json:"transfer_id,omitempty"`
}

// The identifier of the Check Transfer that led to this Transaction.
func (r *InboundACHTransferSimulationResultTransactionSourceCheckTransferRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
