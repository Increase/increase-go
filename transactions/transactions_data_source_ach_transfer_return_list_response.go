package transactions

type TransactionsDataSourceACHTransferReturnListResponse struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// Why the ACH Transfer was returned.
	ReturnReasonCode *string `json:"return_reason_code,omitempty"`

	// The identifier of the ACH Transfer associated with this return.
	TransferId *string `json:"transfer_id,omitempty"`

	// The identifier of the Tranasaction associated with this return.
	TransactionId *string `json:"transaction_id,omitempty"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r *TransactionsDataSourceACHTransferReturnListResponse) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// Why the ACH Transfer was returned.
func (r *TransactionsDataSourceACHTransferReturnListResponse) GetReturnReasonCode() (ReturnReasonCode string) {
	if r != nil && r.ReturnReasonCode != nil {
		ReturnReasonCode = *r.ReturnReasonCode
	}
	return
}

// The identifier of the ACH Transfer associated with this return.
func (r *TransactionsDataSourceACHTransferReturnListResponse) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

// The identifier of the Tranasaction associated with this return.
func (r *TransactionsDataSourceACHTransferReturnListResponse) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}
