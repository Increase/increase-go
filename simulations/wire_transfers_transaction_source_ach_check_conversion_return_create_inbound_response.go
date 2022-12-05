package simulations

type WireTransfersTransactionSourceACHCheckConversionReturnCreateInboundResponse struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// Why the transfer was returned.
	ReturnReasonCode *string `json:"return_reason_code,omitempty"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *WireTransfersTransactionSourceACHCheckConversionReturnCreateInboundResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// Why the transfer was returned.
func (r *WireTransfersTransactionSourceACHCheckConversionReturnCreateInboundResponse) GetReturnReasonCode() (ReturnReasonCode string) {
	if r != nil && r.ReturnReasonCode != nil {
		ReturnReasonCode = *r.ReturnReasonCode
	}
	return
}
