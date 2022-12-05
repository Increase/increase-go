package simulations

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejection struct {
	TransferId *string `json:"transfer_id,omitempty"`
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
