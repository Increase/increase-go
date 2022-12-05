package wire_drawdown_requests

type WireDrawdownRequest struct {
	// A constant representing the object's type. For this resource it will always be
	// `wire_drawdown_request`.
	Type *string `json:"type,omitempty"`

	// The Wire drawdown request identifier.
	Id *string `json:"id,omitempty"`

	// The Account Number to which the recipient of this request is being requested to
	// send funds.
	AccountNumberId *string `json:"account_number_id,omitempty"`

	// The drawdown request's recipient's account number.
	RecipientAccountNumber *string `json:"recipient_account_number,omitempty"`

	// The drawdown request's recipient's routing number.
	RecipientRoutingNumber *string `json:"recipient_routing_number,omitempty"`

	// The amount being requested in cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being
	// requested. Will always be "USD".
	Currency *string `json:"currency,omitempty"`

	// The message the recipient will see as part of the drawdown request.
	MessageToRecipient *string `json:"message_to_recipient,omitempty"`

	// The drawdown request's recipient's name.
	RecipientName *string `json:"recipient_name,omitempty"`

	// Line 1 of the drawdown request's recipient's address.
	RecipientAddressLine1 *string `json:"recipient_address_line1,omitempty"`

	// Line 2 of the drawdown request's recipient's address.
	RecipientAddressLine2 *string `json:"recipient_address_line2,omitempty"`

	// Line 3 of the drawdown request's recipient's address.
	RecipientAddressLine3 *string `json:"recipient_address_line3,omitempty"`

	// After the drawdown request is submitted to Fedwire, this will contain
	// supplemental details.
	Submission *WireDrawdownRequestSubmission `json:"submission,omitempty"`

	// If the recipient fulfills the drawdown request by sending funds, then this will
	// be the identifier of the corresponding Transaction.
	FulfillmentTransactionId *string `json:"fulfillment_transaction_id,omitempty"`

	// The lifecycle status of the drawdown request.
	Status *string `json:"status,omitempty"`
}

// A constant representing the object's type. For this resource it will always be
// `wire_drawdown_request`.
func (r *WireDrawdownRequest) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

// The Wire drawdown request identifier.
func (r *WireDrawdownRequest) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The Account Number to which the recipient of this request is being requested to
// send funds.
func (r *WireDrawdownRequest) GetAccountNumberId() (AccountNumberId string) {
	if r != nil && r.AccountNumberId != nil {
		AccountNumberId = *r.AccountNumberId
	}
	return
}

// The drawdown request's recipient's account number.
func (r *WireDrawdownRequest) GetRecipientAccountNumber() (RecipientAccountNumber string) {
	if r != nil && r.RecipientAccountNumber != nil {
		RecipientAccountNumber = *r.RecipientAccountNumber
	}
	return
}

// The drawdown request's recipient's routing number.
func (r *WireDrawdownRequest) GetRecipientRoutingNumber() (RecipientRoutingNumber string) {
	if r != nil && r.RecipientRoutingNumber != nil {
		RecipientRoutingNumber = *r.RecipientRoutingNumber
	}
	return
}

// The amount being requested in cents.
func (r *WireDrawdownRequest) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being
// requested. Will always be "USD".
func (r *WireDrawdownRequest) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The message the recipient will see as part of the drawdown request.
func (r *WireDrawdownRequest) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

// The drawdown request's recipient's name.
func (r *WireDrawdownRequest) GetRecipientName() (RecipientName string) {
	if r != nil && r.RecipientName != nil {
		RecipientName = *r.RecipientName
	}
	return
}

// Line 1 of the drawdown request's recipient's address.
func (r *WireDrawdownRequest) GetRecipientAddressLine1() (RecipientAddressLine1 string) {
	if r != nil && r.RecipientAddressLine1 != nil {
		RecipientAddressLine1 = *r.RecipientAddressLine1
	}
	return
}

// Line 2 of the drawdown request's recipient's address.
func (r *WireDrawdownRequest) GetRecipientAddressLine2() (RecipientAddressLine2 string) {
	if r != nil && r.RecipientAddressLine2 != nil {
		RecipientAddressLine2 = *r.RecipientAddressLine2
	}
	return
}

// Line 3 of the drawdown request's recipient's address.
func (r *WireDrawdownRequest) GetRecipientAddressLine3() (RecipientAddressLine3 string) {
	if r != nil && r.RecipientAddressLine3 != nil {
		RecipientAddressLine3 = *r.RecipientAddressLine3
	}
	return
}

// After the drawdown request is submitted to Fedwire, this will contain
// supplemental details.
func (r *WireDrawdownRequest) GetSubmission() (Submission WireDrawdownRequestSubmission) {
	if r != nil && r.Submission != nil {
		Submission = *r.Submission
	}
	return
}

// If the recipient fulfills the drawdown request by sending funds, then this will
// be the identifier of the corresponding Transaction.
func (r *WireDrawdownRequest) GetFulfillmentTransactionId() (FulfillmentTransactionId string) {
	if r != nil && r.FulfillmentTransactionId != nil {
		FulfillmentTransactionId = *r.FulfillmentTransactionId
	}
	return
}

// The lifecycle status of the drawdown request.
func (r *WireDrawdownRequest) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}
