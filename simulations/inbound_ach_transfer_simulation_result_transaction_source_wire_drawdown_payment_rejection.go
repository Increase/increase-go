package simulations

type InboundACHTransferSimulationResultTransactionSourceWireDrawdownPaymentRejection struct {
	TransferId *string `json:"transfer_id,omitempty"`
}

func (r *InboundACHTransferSimulationResultTransactionSourceWireDrawdownPaymentRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
