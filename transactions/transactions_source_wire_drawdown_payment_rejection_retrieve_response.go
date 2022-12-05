package transactions

type TransactionsSourceWireDrawdownPaymentRejectionRetrieveResponse struct {
	TransferId *string `json:"transfer_id,omitempty"`
}

func (r *TransactionsSourceWireDrawdownPaymentRejectionRetrieveResponse) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
