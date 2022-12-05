package simulations

type WireTransfersTransactionSourceWireDrawdownPaymentRejectionCreateInboundResponse struct {
	TransferId *string `json:"transfer_id,omitempty"`
}

func (r *WireTransfersTransactionSourceWireDrawdownPaymentRejectionCreateInboundResponse) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
