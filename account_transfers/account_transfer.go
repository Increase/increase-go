package account_transfers

type AccountTransfer struct {
	// The account transfer's identifier.
	Id *string `json:"id,omitempty"`

	// The transfer amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The Account to which the transfer belongs.
	AccountId *string `json:"account_id,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *string `json:"currency,omitempty"`

	// The destination account's identifier.
	DestinationAccountId *string `json:"destination_account_id,omitempty"`

	// The ID for the transaction receiving the transfer.
	DestinationTransactionId *string `json:"destination_transaction_id,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The description that will show on the transactions.
	Description *string `json:"description,omitempty"`

	// The transfer's network.
	Network *string `json:"network,omitempty"`

	// The lifecycle status of the transfer.
	Status *string `json:"status,omitempty"`

	// If the transfer was created from a template, this will be the template's ID.
	TemplateId *string `json:"template_id,omitempty"`

	// The ID for the transaction funding the transfer.
	TransactionId *string `json:"transaction_id,omitempty"`

	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval *AccountTransferApproval `json:"approval,omitempty"`

	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation *AccountTransferCancellation `json:"cancellation,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `account_transfer`.
	Type *string `json:"type,omitempty"`
}

// The account transfer's identifier.
func (r *AccountTransfer) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The transfer amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *AccountTransfer) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The Account to which the transfer belongs.
func (r *AccountTransfer) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *AccountTransfer) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The destination account's identifier.
func (r *AccountTransfer) GetDestinationAccountId() (DestinationAccountId string) {
	if r != nil && r.DestinationAccountId != nil {
		DestinationAccountId = *r.DestinationAccountId
	}
	return
}

// The ID for the transaction receiving the transfer.
func (r *AccountTransfer) GetDestinationTransactionId() (DestinationTransactionId string) {
	if r != nil && r.DestinationTransactionId != nil {
		DestinationTransactionId = *r.DestinationTransactionId
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r *AccountTransfer) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The description that will show on the transactions.
func (r *AccountTransfer) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The transfer's network.
func (r *AccountTransfer) GetNetwork() (Network string) {
	if r != nil && r.Network != nil {
		Network = *r.Network
	}
	return
}

// The lifecycle status of the transfer.
func (r *AccountTransfer) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// If the transfer was created from a template, this will be the template's ID.
func (r *AccountTransfer) GetTemplateId() (TemplateId string) {
	if r != nil && r.TemplateId != nil {
		TemplateId = *r.TemplateId
	}
	return
}

// The ID for the transaction funding the transfer.
func (r *AccountTransfer) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
func (r *AccountTransfer) GetApproval() (Approval AccountTransferApproval) {
	if r != nil && r.Approval != nil {
		Approval = *r.Approval
	}
	return
}

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
func (r *AccountTransfer) GetCancellation() (Cancellation AccountTransferCancellation) {
	if r != nil && r.Cancellation != nil {
		Cancellation = *r.Cancellation
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `account_transfer`.
func (r *AccountTransfer) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
