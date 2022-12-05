package wire_drawdown_requests

type WireDrawdownRequestsCreateParams struct {
	// The Account Number to which the recipient should send funds.
	AccountNumberId *string `json:"account_number_id,omitempty"`

	// The amount requested from the recipient, in cents.
	Amount *int64 `json:"amount,omitempty"`

	// A message the recipient will see as part of the request.
	MessageToRecipient *string `json:"message_to_recipient,omitempty"`

	// The drawdown request's recipient's account number.
	RecipientAccountNumber *string `json:"recipient_account_number,omitempty"`

	// The drawdown request's recipient's routing number.
	RecipientRoutingNumber *string `json:"recipient_routing_number,omitempty"`

	// The drawdown request's recipient's name.
	RecipientName *string `json:"recipient_name,omitempty"`

	// Line 1 of the drawdown request's recipient's address.
	RecipientAddressLine1 *string `json:"recipient_address_line1,omitempty"`

	// Line 2 of the drawdown request's recipient's address.
	RecipientAddressLine2 *string `json:"recipient_address_line2,omitempty"`

	// Line 3 of the drawdown request's recipient's address.
	RecipientAddressLine3 *string `json:"recipient_address_line3,omitempty"`
}

// The Account Number to which the recipient should send funds.
func (r *WireDrawdownRequestsCreateParams) GetAccountNumberId() (AccountNumberId string) {
	if r != nil && r.AccountNumberId != nil {
		AccountNumberId = *r.AccountNumberId
	}
	return
}

// The amount requested from the recipient, in cents.
func (r *WireDrawdownRequestsCreateParams) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// A message the recipient will see as part of the request.
func (r *WireDrawdownRequestsCreateParams) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

// The drawdown request's recipient's account number.
func (r *WireDrawdownRequestsCreateParams) GetRecipientAccountNumber() (RecipientAccountNumber string) {
	if r != nil && r.RecipientAccountNumber != nil {
		RecipientAccountNumber = *r.RecipientAccountNumber
	}
	return
}

// The drawdown request's recipient's routing number.
func (r *WireDrawdownRequestsCreateParams) GetRecipientRoutingNumber() (RecipientRoutingNumber string) {
	if r != nil && r.RecipientRoutingNumber != nil {
		RecipientRoutingNumber = *r.RecipientRoutingNumber
	}
	return
}

// The drawdown request's recipient's name.
func (r *WireDrawdownRequestsCreateParams) GetRecipientName() (RecipientName string) {
	if r != nil && r.RecipientName != nil {
		RecipientName = *r.RecipientName
	}
	return
}

// Line 1 of the drawdown request's recipient's address.
func (r *WireDrawdownRequestsCreateParams) GetRecipientAddressLine1() (RecipientAddressLine1 string) {
	if r != nil && r.RecipientAddressLine1 != nil {
		RecipientAddressLine1 = *r.RecipientAddressLine1
	}
	return
}

// Line 2 of the drawdown request's recipient's address.
func (r *WireDrawdownRequestsCreateParams) GetRecipientAddressLine2() (RecipientAddressLine2 string) {
	if r != nil && r.RecipientAddressLine2 != nil {
		RecipientAddressLine2 = *r.RecipientAddressLine2
	}
	return
}

// Line 3 of the drawdown request's recipient's address.
func (r *WireDrawdownRequestsCreateParams) GetRecipientAddressLine3() (RecipientAddressLine3 string) {
	if r != nil && r.RecipientAddressLine3 != nil {
		RecipientAddressLine3 = *r.RecipientAddressLine3
	}
	return
}
