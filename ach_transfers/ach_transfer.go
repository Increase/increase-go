package ach_transfers

type ACHTransfer struct {
	// The Account to which the transfer belongs.
	AccountId *string `json:"account_id,omitempty"`

	// The destination account number.
	AccountNumber *string `json:"account_number,omitempty"`

	// Additional information that will be sent to the recipient.
	Addendum *string `json:"addendum,omitempty"`

	// The transfer amount in USD cents. A positive amount indicates a credit transfer
	// pushing funds to the receiving account. A negative amount indicates a debit
	// transfer pulling funds from the receiving account.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For ACH transfers this is always equal to `usd`.
	Currency *string `json:"currency,omitempty"`

	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval *ACHTransferApproval `json:"approval,omitempty"`

	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation *ACHTransferCancellation `json:"cancellation,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The identifier of the External Account the transfer was made to, if any.
	ExternalAccountId *string `json:"external_account_id,omitempty"`

	// The ACH transfer's identifier.
	Id *string `json:"id,omitempty"`

	// The transfer's network.
	Network *string `json:"network,omitempty"`

	// If the receiving bank accepts the transfer but notifies that future transfers
	// should use different details, this will contain those details.
	NotificationOfChange *ACHTransferNotificationOfChange `json:"notification_of_change,omitempty"`

	// If your transfer is returned, this will contain details of the return.
	Return *ACHTransferReturn `json:"return,omitempty"`

	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber *string `json:"routing_number,omitempty"`

	// The descriptor that will show on the recipient's bank statement.
	StatementDescriptor *string `json:"statement_descriptor,omitempty"`

	// The lifecycle status of the transfer.
	Status *string `json:"status,omitempty"`

	// After the transfer is submitted to FedACH, this will contain supplemental
	// details.
	Submission *ACHTransferSubmission `json:"submission,omitempty"`

	// If the transfer was created from a template, this will be the template's ID.
	TemplateId *string `json:"template_id,omitempty"`

	// The ID for the transaction funding the transfer.
	TransactionId *string `json:"transaction_id,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `ach_transfer`.
	Type *string `json:"type,omitempty"`
}

// The Account to which the transfer belongs.
func (r *ACHTransfer) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The destination account number.
func (r *ACHTransfer) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// Additional information that will be sent to the recipient.
func (r *ACHTransfer) GetAddendum() (Addendum string) {
	if r != nil && r.Addendum != nil {
		Addendum = *r.Addendum
	}
	return
}

// The transfer amount in USD cents. A positive amount indicates a credit transfer
// pushing funds to the receiving account. A negative amount indicates a debit
// transfer pulling funds from the receiving account.
func (r *ACHTransfer) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
// currency. For ACH transfers this is always equal to `usd`.
func (r *ACHTransfer) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
func (r *ACHTransfer) GetApproval() (Approval ACHTransferApproval) {
	if r != nil && r.Approval != nil {
		Approval = *r.Approval
	}
	return
}

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
func (r *ACHTransfer) GetCancellation() (Cancellation ACHTransferCancellation) {
	if r != nil && r.Cancellation != nil {
		Cancellation = *r.Cancellation
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r *ACHTransfer) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The identifier of the External Account the transfer was made to, if any.
func (r *ACHTransfer) GetExternalAccountId() (ExternalAccountId string) {
	if r != nil && r.ExternalAccountId != nil {
		ExternalAccountId = *r.ExternalAccountId
	}
	return
}

// The ACH transfer's identifier.
func (r *ACHTransfer) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The transfer's network.
func (r *ACHTransfer) GetNetwork() (Network string) {
	if r != nil && r.Network != nil {
		Network = *r.Network
	}
	return
}

// If the receiving bank accepts the transfer but notifies that future transfers
// should use different details, this will contain those details.
func (r *ACHTransfer) GetNotificationOfChange() (NotificationOfChange ACHTransferNotificationOfChange) {
	if r != nil && r.NotificationOfChange != nil {
		NotificationOfChange = *r.NotificationOfChange
	}
	return
}

// If your transfer is returned, this will contain details of the return.
func (r *ACHTransfer) GetReturn() (Return ACHTransferReturn) {
	if r != nil && r.Return != nil {
		Return = *r.Return
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r *ACHTransfer) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The descriptor that will show on the recipient's bank statement.
func (r *ACHTransfer) GetStatementDescriptor() (StatementDescriptor string) {
	if r != nil && r.StatementDescriptor != nil {
		StatementDescriptor = *r.StatementDescriptor
	}
	return
}

// The lifecycle status of the transfer.
func (r *ACHTransfer) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// After the transfer is submitted to FedACH, this will contain supplemental
// details.
func (r *ACHTransfer) GetSubmission() (Submission ACHTransferSubmission) {
	if r != nil && r.Submission != nil {
		Submission = *r.Submission
	}
	return
}

// If the transfer was created from a template, this will be the template's ID.
func (r *ACHTransfer) GetTemplateId() (TemplateId string) {
	if r != nil && r.TemplateId != nil {
		TemplateId = *r.TemplateId
	}
	return
}

// The ID for the transaction funding the transfer.
func (r *ACHTransfer) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `ach_transfer`.
func (r *ACHTransfer) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
