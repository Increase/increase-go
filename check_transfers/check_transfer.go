package check_transfers

type CheckTransfer struct {
	// The identifier of the Account from which funds will be transferred.
	AccountId *string `json:"account_id,omitempty"`

	// The street address of the check's destination.
	AddressLine1 *string `json:"address_line1,omitempty"`

	// The second line of the address of the check's destination.
	AddressLine2 *string `json:"address_line2,omitempty"`

	// The city of the check's destination.
	AddressCity *string `json:"address_city,omitempty"`

	// The state of the check's destination.
	AddressState *string `json:"address_state,omitempty"`

	// The postal code of the check's destination.
	AddressZip *string `json:"address_zip,omitempty"`

	// The transfer amount in USD cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency *string `json:"currency,omitempty"`

	// The Check transfer's identifier.
	Id *string `json:"id,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check was mailed.
	MailedAt *string `json:"mailed_at,omitempty"`

	// The descriptor that is printed on the check.
	Message *string `json:"message,omitempty"`

	// The name that will be printed on the check.
	RecipientName *string `json:"recipient_name,omitempty"`

	// The lifecycle status of the transfer.
	Status *string `json:"status,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check was submitted.
	SubmittedAt *string `json:"submitted_at,omitempty"`

	// After the transfer is submitted, this will contain supplemental details.
	Submission *CheckTransferSubmission `json:"submission,omitempty"`

	// If the transfer was created from a template, this will be the template's ID.
	TemplateId *string `json:"template_id,omitempty"`

	// The ID for the transaction caused by the transfer.
	TransactionId *string `json:"transaction_id,omitempty"`

	// After a stop-payment is requested on the check, this will contain supplemental
	// details.
	StopPaymentRequest *CheckTransferStopPaymentRequest `json:"stop_payment_request,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `check_transfer`.
	Type *string `json:"type,omitempty"`
}

// The identifier of the Account from which funds will be transferred.
func (r *CheckTransfer) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The street address of the check's destination.
func (r *CheckTransfer) GetAddressLine1() (AddressLine1 string) {
	if r != nil && r.AddressLine1 != nil {
		AddressLine1 = *r.AddressLine1
	}
	return
}

// The second line of the address of the check's destination.
func (r *CheckTransfer) GetAddressLine2() (AddressLine2 string) {
	if r != nil && r.AddressLine2 != nil {
		AddressLine2 = *r.AddressLine2
	}
	return
}

// The city of the check's destination.
func (r *CheckTransfer) GetAddressCity() (AddressCity string) {
	if r != nil && r.AddressCity != nil {
		AddressCity = *r.AddressCity
	}
	return
}

// The state of the check's destination.
func (r *CheckTransfer) GetAddressState() (AddressState string) {
	if r != nil && r.AddressState != nil {
		AddressState = *r.AddressState
	}
	return
}

// The postal code of the check's destination.
func (r *CheckTransfer) GetAddressZip() (AddressZip string) {
	if r != nil && r.AddressZip != nil {
		AddressZip = *r.AddressZip
	}
	return
}

// The transfer amount in USD cents.
func (r *CheckTransfer) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r *CheckTransfer) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
func (r *CheckTransfer) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The Check transfer's identifier.
func (r *CheckTransfer) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the check was mailed.
func (r *CheckTransfer) GetMailedAt() (MailedAt string) {
	if r != nil && r.MailedAt != nil {
		MailedAt = *r.MailedAt
	}
	return
}

// The descriptor that is printed on the check.
func (r *CheckTransfer) GetMessage() (Message string) {
	if r != nil && r.Message != nil {
		Message = *r.Message
	}
	return
}

// The name that will be printed on the check.
func (r *CheckTransfer) GetRecipientName() (RecipientName string) {
	if r != nil && r.RecipientName != nil {
		RecipientName = *r.RecipientName
	}
	return
}

// The lifecycle status of the transfer.
func (r *CheckTransfer) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the check was submitted.
func (r *CheckTransfer) GetSubmittedAt() (SubmittedAt string) {
	if r != nil && r.SubmittedAt != nil {
		SubmittedAt = *r.SubmittedAt
	}
	return
}

// After the transfer is submitted, this will contain supplemental details.
func (r *CheckTransfer) GetSubmission() (Submission CheckTransferSubmission) {
	if r != nil && r.Submission != nil {
		Submission = *r.Submission
	}
	return
}

// If the transfer was created from a template, this will be the template's ID.
func (r *CheckTransfer) GetTemplateId() (TemplateId string) {
	if r != nil && r.TemplateId != nil {
		TemplateId = *r.TemplateId
	}
	return
}

// The ID for the transaction caused by the transfer.
func (r *CheckTransfer) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

// After a stop-payment is requested on the check, this will contain supplemental
// details.
func (r *CheckTransfer) GetStopPaymentRequest() (StopPaymentRequest CheckTransferStopPaymentRequest) {
	if r != nil && r.StopPaymentRequest != nil {
		StopPaymentRequest = *r.StopPaymentRequest
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer`.
func (r *CheckTransfer) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
