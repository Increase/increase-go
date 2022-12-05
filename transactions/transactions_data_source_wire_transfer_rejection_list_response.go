package transactions

type TransactionsDataSourceWireTransferRejectionListResponse struct {
	TransferId *string `json:"transfer_id,omitempty"`
}

func (r *TransactionsDataSourceWireTransferRejectionListResponse) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
