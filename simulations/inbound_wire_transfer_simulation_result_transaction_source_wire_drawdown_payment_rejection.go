package simulations

type InboundWireTransferSimulationResultTransactionSourceWireDrawdownPaymentRejection struct {
	TransferId *string `json:"transfer_id,omitempty"`
}

func (r *InboundWireTransferSimulationResultTransactionSourceWireDrawdownPaymentRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
