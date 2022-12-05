package pending_transactions

type PendingTransactionsDataSourceCheckDepositInstructionListResponse struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *string `json:"currency,omitempty"`

	// The identifier of the File containing the image of the front of the check that
	// was deposited.
	FrontImageFileId *string `json:"front_image_file_id,omitempty"`

	// The identifier of the File containing the image of the back of the check that
	// was deposited.
	BackImageFileId *string `json:"back_image_file_id,omitempty"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionsDataSourceCheckDepositInstructionListResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *PendingTransactionsDataSourceCheckDepositInstructionListResponse) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the File containing the image of the front of the check that
// was deposited.
func (r *PendingTransactionsDataSourceCheckDepositInstructionListResponse) GetFrontImageFileId() (FrontImageFileId string) {
	if r != nil && r.FrontImageFileId != nil {
		FrontImageFileId = *r.FrontImageFileId
	}
	return
}

// The identifier of the File containing the image of the back of the check that
// was deposited.
func (r *PendingTransactionsDataSourceCheckDepositInstructionListResponse) GetBackImageFileId() (BackImageFileId string) {
	if r != nil && r.BackImageFileId != nil {
		BackImageFileId = *r.BackImageFileId
	}
	return
}
