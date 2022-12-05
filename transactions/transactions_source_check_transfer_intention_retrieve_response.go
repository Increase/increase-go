package transactions

type TransactionsSourceCheckTransferIntentionRetrieveResponse struct {
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

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency *string `json:"currency,omitempty"`

	// The name that will be printed on the check.
	RecipientName *string `json:"recipient_name,omitempty"`

	// The identifier of the Check Transfer with which this is associated.
	TransferId *string `json:"transfer_id,omitempty"`
}

// The street address of the check's destination.
func (r *TransactionsSourceCheckTransferIntentionRetrieveResponse) GetAddressLine1() (AddressLine1 string) {
	if r != nil && r.AddressLine1 != nil {
		AddressLine1 = *r.AddressLine1
	}
	return
}

// The second line of the address of the check's destination.
func (r *TransactionsSourceCheckTransferIntentionRetrieveResponse) GetAddressLine2() (AddressLine2 string) {
	if r != nil && r.AddressLine2 != nil {
		AddressLine2 = *r.AddressLine2
	}
	return
}

// The city of the check's destination.
func (r *TransactionsSourceCheckTransferIntentionRetrieveResponse) GetAddressCity() (AddressCity string) {
	if r != nil && r.AddressCity != nil {
		AddressCity = *r.AddressCity
	}
	return
}

// The state of the check's destination.
func (r *TransactionsSourceCheckTransferIntentionRetrieveResponse) GetAddressState() (AddressState string) {
	if r != nil && r.AddressState != nil {
		AddressState = *r.AddressState
	}
	return
}

// The postal code of the check's destination.
func (r *TransactionsSourceCheckTransferIntentionRetrieveResponse) GetAddressZip() (AddressZip string) {
	if r != nil && r.AddressZip != nil {
		AddressZip = *r.AddressZip
	}
	return
}

// The transfer amount in USD cents.
func (r *TransactionsSourceCheckTransferIntentionRetrieveResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
func (r *TransactionsSourceCheckTransferIntentionRetrieveResponse) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The name that will be printed on the check.
func (r *TransactionsSourceCheckTransferIntentionRetrieveResponse) GetRecipientName() (RecipientName string) {
	if r != nil && r.RecipientName != nil {
		RecipientName = *r.RecipientName
	}
	return
}

// The identifier of the Check Transfer with which this is associated.
func (r *TransactionsSourceCheckTransferIntentionRetrieveResponse) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
