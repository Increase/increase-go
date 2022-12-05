package simulations

type CheckDeposit struct {
	// The deposit's identifier.
	Id *string `json:"id,omitempty"`

	// The deposited amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the deposit.
	Currency *string `json:"currency,omitempty"`

	// The status of the Check Deposit.
	Status *string `json:"status,omitempty"`

	// The Account the check was deposited into.
	AccountId *string `json:"account_id,omitempty"`

	// The ID for the File containing the image of the front of the check.
	FrontImageFileId *string `json:"front_image_file_id,omitempty"`

	// The ID for the File containing the image of the back of the check.
	BackImageFileId *string `json:"back_image_file_id,omitempty"`

	// The ID for the Transaction created by the deposit.
	TransactionId *string `json:"transaction_id,omitempty"`

	// If your deposit is successfully parsed and accepted by Increase, this will
	// contain details of the parsed check.
	DepositAcceptance *CheckDepositDepositAcceptance `json:"deposit_acceptance,omitempty"`

	// If your deposit is rejected by Increase, this will contain details as to why it
	// was rejected.
	DepositRejection *CheckDepositDepositRejection `json:"deposit_rejection,omitempty"`

	// If your deposit is returned, this will contain details as to why it was
	// returned.
	DepositReturn *CheckDepositDepositReturn `json:"deposit_return,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `check_deposit`.
	Type *string `json:"type,omitempty"`
}

// The deposit's identifier.
func (r *CheckDeposit) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The deposited amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *CheckDeposit) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r *CheckDeposit) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the deposit.
func (r *CheckDeposit) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The status of the Check Deposit.
func (r *CheckDeposit) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// The Account the check was deposited into.
func (r *CheckDeposit) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The ID for the File containing the image of the front of the check.
func (r *CheckDeposit) GetFrontImageFileId() (FrontImageFileId string) {
	if r != nil && r.FrontImageFileId != nil {
		FrontImageFileId = *r.FrontImageFileId
	}
	return
}

// The ID for the File containing the image of the back of the check.
func (r *CheckDeposit) GetBackImageFileId() (BackImageFileId string) {
	if r != nil && r.BackImageFileId != nil {
		BackImageFileId = *r.BackImageFileId
	}
	return
}

// The ID for the Transaction created by the deposit.
func (r *CheckDeposit) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

// If your deposit is successfully parsed and accepted by Increase, this will
// contain details of the parsed check.
func (r *CheckDeposit) GetDepositAcceptance() (DepositAcceptance CheckDepositDepositAcceptance) {
	if r != nil && r.DepositAcceptance != nil {
		DepositAcceptance = *r.DepositAcceptance
	}
	return
}

// If your deposit is rejected by Increase, this will contain details as to why it
// was rejected.
func (r *CheckDeposit) GetDepositRejection() (DepositRejection CheckDepositDepositRejection) {
	if r != nil && r.DepositRejection != nil {
		DepositRejection = *r.DepositRejection
	}
	return
}

// If your deposit is returned, this will contain details as to why it was
// returned.
func (r *CheckDeposit) GetDepositReturn() (DepositReturn CheckDepositDepositReturn) {
	if r != nil && r.DepositReturn != nil {
		DepositReturn = *r.DepositReturn
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `check_deposit`.
func (r *CheckDeposit) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
