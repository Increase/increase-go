package transactions

type TransactionSourceWireTransferRejection struct {
	TransferId *string `json:"transfer_id,omitempty"`
}

func (r *TransactionSourceWireTransferRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
