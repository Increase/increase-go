package simulations

type WireTransfersTransactionSourceWireTransferRejectionCreateInboundResponse struct {
	TransferId *string `json:"transfer_id,omitempty"`
}

func (r *WireTransfersTransactionSourceWireTransferRejectionCreateInboundResponse) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
