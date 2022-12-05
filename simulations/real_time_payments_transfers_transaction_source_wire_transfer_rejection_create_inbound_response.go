package simulations

type RealTimePaymentsTransfersTransactionSourceWireTransferRejectionCreateInboundResponse struct {
	TransferId *string `json:"transfer_id,omitempty"`
}

func (r *RealTimePaymentsTransfersTransactionSourceWireTransferRejectionCreateInboundResponse) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
