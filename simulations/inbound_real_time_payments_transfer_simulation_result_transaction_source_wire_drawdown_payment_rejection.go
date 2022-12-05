package simulations

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentRejection struct {
	TransferId *string `json:"transfer_id,omitempty"`
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
