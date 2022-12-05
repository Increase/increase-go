package transactions

type TransactionsDataSourceWireDrawdownPaymentRejectionListResponse struct {
	TransferId *string `json:"transfer_id,omitempty"`
}

func (r *TransactionsDataSourceWireDrawdownPaymentRejectionListResponse) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
