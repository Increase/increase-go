package simulations

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversion struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The identifier of the File containing an image of the returned check.
	FileId *string `json:"file_id,omitempty"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversion) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The identifier of the File containing an image of the returned check.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversion) GetFileId() (FileId string) {
	if r != nil && r.FileId != nil {
		FileId = *r.FileId
	}
	return
}
