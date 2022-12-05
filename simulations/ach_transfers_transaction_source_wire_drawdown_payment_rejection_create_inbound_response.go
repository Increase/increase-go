package simulations

type ACHTransfersTransactionSourceWireDrawdownPaymentRejectionCreateInboundResponse struct {
	TransferId *string `json:"transfer_id,omitempty"`
}

func (r *ACHTransfersTransactionSourceWireDrawdownPaymentRejectionCreateInboundResponse) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
