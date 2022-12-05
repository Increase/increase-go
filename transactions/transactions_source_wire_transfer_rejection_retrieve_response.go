package transactions

type TransactionsSourceWireTransferRejectionRetrieveResponse struct {
	TransferId *string `json:"transfer_id,omitempty"`
}

func (r *TransactionsSourceWireTransferRejectionRetrieveResponse) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
