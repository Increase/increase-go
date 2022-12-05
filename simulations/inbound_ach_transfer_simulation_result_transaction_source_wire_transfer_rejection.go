package simulations

type InboundACHTransferSimulationResultTransactionSourceWireTransferRejection struct {
	TransferId *string `json:"transfer_id,omitempty"`
}

func (r *InboundACHTransferSimulationResultTransactionSourceWireTransferRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
