package simulations

type ACHTransfersTransactionSourceWireTransferRejectionCreateInboundResponse struct {
	TransferId *string `json:"transfer_id,omitempty"`
}

func (r *ACHTransfersTransactionSourceWireTransferRejectionCreateInboundResponse) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
