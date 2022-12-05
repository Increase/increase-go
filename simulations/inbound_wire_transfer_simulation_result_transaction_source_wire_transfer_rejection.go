package simulations

type InboundWireTransferSimulationResultTransactionSourceWireTransferRejection struct {
	TransferId *string `json:"transfer_id,omitempty"`
}

func (r *InboundWireTransferSimulationResultTransactionSourceWireTransferRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
