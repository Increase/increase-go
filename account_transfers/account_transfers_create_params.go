package account_transfers

type AccountTransfersCreateParams struct {
	// The identifier for the account that will send the transfer.
	AccountId *string `json:"account_id,omitempty"`

	// The transfer amount in the minor unit of the account currency. For dollars, for
	// example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The description you choose to give the transfer.
	Description *string `json:"description,omitempty"`

	// The identifier for the account that will receive the transfer.
	DestinationAccountId *string `json:"destination_account_id,omitempty"`
}

// The identifier for the account that will send the transfer.
func (r *AccountTransfersCreateParams) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The transfer amount in the minor unit of the account currency. For dollars, for
// example, this is cents.
func (r *AccountTransfersCreateParams) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description you choose to give the transfer.
func (r *AccountTransfersCreateParams) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The identifier for the account that will receive the transfer.
func (r *AccountTransfersCreateParams) GetDestinationAccountId() (DestinationAccountId string) {
	if r != nil && r.DestinationAccountId != nil {
		DestinationAccountId = *r.DestinationAccountId
	}
	return
}
