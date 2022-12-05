package wire_transfers

type WireTransfer struct {
	// The wire transfer's identifier.
	Id *string `json:"id,omitempty"`

	// The message that will show on the recipient's bank statement.
	MessageToRecipient *string `json:"message_to_recipient,omitempty"`

	// The transfer amount in USD cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For wire transfers this is always equal to `usd`.
	Currency *string `json:"currency,omitempty"`

	// The destination account number.
	AccountNumber *string `json:"account_number,omitempty"`

	// The Account to which the transfer belongs.
	AccountId *string `json:"account_id,omitempty"`

	// The identifier of the External Account the transfer was made to, if any.
	ExternalAccountId *string `json:"external_account_id,omitempty"`

	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber *string `json:"routing_number,omitempty"`

	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval *WireTransferApproval `json:"approval,omitempty"`

	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation *WireTransferCancellation `json:"cancellation,omitempty"`

	// If your transfer is reversed, this will contain details of the reversal.
	Reversal *WireTransferReversal `json:"reversal,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The transfer's network.
	Network *string `json:"network,omitempty"`

	// The lifecycle status of the transfer.
	Status *string `json:"status,omitempty"`

	// After the transfer is submitted to Fedwire, this will contain supplemental
	// details.
	Submission *WireTransferSubmission `json:"submission,omitempty"`

	// If the transfer was created from a template, this will be the template's ID.
	TemplateId *string `json:"template_id,omitempty"`

	// The ID for the transaction funding the transfer.
	TransactionId *string `json:"transaction_id,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `wire_transfer`.
	Type *string `json:"type,omitempty"`
}

// The wire transfer's identifier.
func (r *WireTransfer) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The message that will show on the recipient's bank statement.
func (r *WireTransfer) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

// The transfer amount in USD cents.
func (r *WireTransfer) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
// currency. For wire transfers this is always equal to `usd`.
func (r *WireTransfer) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The destination account number.
func (r *WireTransfer) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The Account to which the transfer belongs.
func (r *WireTransfer) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The identifier of the External Account the transfer was made to, if any.
func (r *WireTransfer) GetExternalAccountId() (ExternalAccountId string) {
	if r != nil && r.ExternalAccountId != nil {
		ExternalAccountId = *r.ExternalAccountId
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r *WireTransfer) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
func (r *WireTransfer) GetApproval() (Approval WireTransferApproval) {
	if r != nil && r.Approval != nil {
		Approval = *r.Approval
	}
	return
}

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
func (r *WireTransfer) GetCancellation() (Cancellation WireTransferCancellation) {
	if r != nil && r.Cancellation != nil {
		Cancellation = *r.Cancellation
	}
	return
}

// If your transfer is reversed, this will contain details of the reversal.
func (r *WireTransfer) GetReversal() (Reversal WireTransferReversal) {
	if r != nil && r.Reversal != nil {
		Reversal = *r.Reversal
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r *WireTransfer) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The transfer's network.
func (r *WireTransfer) GetNetwork() (Network string) {
	if r != nil && r.Network != nil {
		Network = *r.Network
	}
	return
}

// The lifecycle status of the transfer.
func (r *WireTransfer) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// After the transfer is submitted to Fedwire, this will contain supplemental
// details.
func (r *WireTransfer) GetSubmission() (Submission WireTransferSubmission) {
	if r != nil && r.Submission != nil {
		Submission = *r.Submission
	}
	return
}

// If the transfer was created from a template, this will be the template's ID.
func (r *WireTransfer) GetTemplateId() (TemplateId string) {
	if r != nil && r.TemplateId != nil {
		TemplateId = *r.TemplateId
	}
	return
}

// The ID for the transaction funding the transfer.
func (r *WireTransfer) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `wire_transfer`.
func (r *WireTransfer) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
