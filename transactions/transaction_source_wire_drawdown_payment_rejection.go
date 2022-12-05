package transactions

type TransactionSourceWireDrawdownPaymentRejection struct {
	TransferId *string `json:"transfer_id,omitempty"`
}

func (r *TransactionSourceWireDrawdownPaymentRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
