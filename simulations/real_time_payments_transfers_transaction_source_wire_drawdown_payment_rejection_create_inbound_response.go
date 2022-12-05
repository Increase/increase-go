package simulations

type RealTimePaymentsTransfersTransactionSourceWireDrawdownPaymentRejectionCreateInboundResponse struct {
	TransferId *string `json:"transfer_id,omitempty"`
}

func (r *RealTimePaymentsTransfersTransactionSourceWireDrawdownPaymentRejectionCreateInboundResponse) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
