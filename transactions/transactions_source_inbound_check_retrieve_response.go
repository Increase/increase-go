package transactions

type TransactionsSourceInboundCheckRetrieveResponse struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *string `json:"currency,omitempty"`

	CheckNumber *string `json:"check_number,omitempty"`

	CheckFrontImageFileId *string `json:"check_front_image_file_id,omitempty"`

	CheckRearImageFileId *string `json:"check_rear_image_file_id,omitempty"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *TransactionsSourceInboundCheckRetrieveResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *TransactionsSourceInboundCheckRetrieveResponse) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *TransactionsSourceInboundCheckRetrieveResponse) GetCheckNumber() (CheckNumber string) {
	if r != nil && r.CheckNumber != nil {
		CheckNumber = *r.CheckNumber
	}
	return
}

func (r *TransactionsSourceInboundCheckRetrieveResponse) GetCheckFrontImageFileId() (CheckFrontImageFileId string) {
	if r != nil && r.CheckFrontImageFileId != nil {
		CheckFrontImageFileId = *r.CheckFrontImageFileId
	}
	return
}

func (r *TransactionsSourceInboundCheckRetrieveResponse) GetCheckRearImageFileId() (CheckRearImageFileId string) {
	if r != nil && r.CheckRearImageFileId != nil {
		CheckRearImageFileId = *r.CheckRearImageFileId
	}
	return
}
