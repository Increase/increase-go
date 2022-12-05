package check_deposits

type CheckDepositsCreateParams struct {
	// The identifier for the Account to deposit the check in.
	AccountId *string `json:"account_id,omitempty"`

	// The deposit amount in the minor unit of the account currency. For dollars, for
	// example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The currency to use for the deposit.
	Currency *string `json:"currency,omitempty"`

	// The File containing the check's front image.
	FrontImageFileId *string `json:"front_image_file_id,omitempty"`

	// The File containing the check's back image.
	BackImageFileId *string `json:"back_image_file_id,omitempty"`
}

// The identifier for the Account to deposit the check in.
func (r *CheckDepositsCreateParams) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The deposit amount in the minor unit of the account currency. For dollars, for
// example, this is cents.
func (r *CheckDepositsCreateParams) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The currency to use for the deposit.
func (r *CheckDepositsCreateParams) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The File containing the check's front image.
func (r *CheckDepositsCreateParams) GetFrontImageFileId() (FrontImageFileId string) {
	if r != nil && r.FrontImageFileId != nil {
		FrontImageFileId = *r.FrontImageFileId
	}
	return
}

// The File containing the check's back image.
func (r *CheckDepositsCreateParams) GetBackImageFileId() (BackImageFileId string) {
	if r != nil && r.BackImageFileId != nil {
		BackImageFileId = *r.BackImageFileId
	}
	return
}
