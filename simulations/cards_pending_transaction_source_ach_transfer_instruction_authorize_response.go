package simulations

type CardsPendingTransactionSourceACHTransferInstructionAuthorizeResponse struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The identifier of the ACH Transfer that led to this Pending Transaction.
	TransferId *string `json:"transfer_id,omitempty"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *CardsPendingTransactionSourceACHTransferInstructionAuthorizeResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The identifier of the ACH Transfer that led to this Pending Transaction.
func (r *CardsPendingTransactionSourceACHTransferInstructionAuthorizeResponse) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
