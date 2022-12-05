package check_transfers

type CheckTransfersCreateParams struct {
	// The identifier for the account that will send the transfer.
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

	// The transfer amount in cents.
	Amount *int64 `json:"amount,omitempty"`

	// The descriptor that will be printed on the check.
	Message *string `json:"message,omitempty"`

	// The name that will be printed on the check.
	RecipientName *string `json:"recipient_name,omitempty"`
}

// The identifier for the account that will send the transfer.
func (r *CheckTransfersCreateParams) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The street address of the check's destination.
func (r *CheckTransfersCreateParams) GetAddressLine1() (AddressLine1 string) {
	if r != nil && r.AddressLine1 != nil {
		AddressLine1 = *r.AddressLine1
	}
	return
}

// The second line of the address of the check's destination.
func (r *CheckTransfersCreateParams) GetAddressLine2() (AddressLine2 string) {
	if r != nil && r.AddressLine2 != nil {
		AddressLine2 = *r.AddressLine2
	}
	return
}

// The city of the check's destination.
func (r *CheckTransfersCreateParams) GetAddressCity() (AddressCity string) {
	if r != nil && r.AddressCity != nil {
		AddressCity = *r.AddressCity
	}
	return
}

// The state of the check's destination.
func (r *CheckTransfersCreateParams) GetAddressState() (AddressState string) {
	if r != nil && r.AddressState != nil {
		AddressState = *r.AddressState
	}
	return
}

// The postal code of the check's destination.
func (r *CheckTransfersCreateParams) GetAddressZip() (AddressZip string) {
	if r != nil && r.AddressZip != nil {
		AddressZip = *r.AddressZip
	}
	return
}

// The transfer amount in cents.
func (r *CheckTransfersCreateParams) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The descriptor that will be printed on the check.
func (r *CheckTransfersCreateParams) GetMessage() (Message string) {
	if r != nil && r.Message != nil {
		Message = *r.Message
	}
	return
}

// The name that will be printed on the check.
func (r *CheckTransfersCreateParams) GetRecipientName() (RecipientName string) {
	if r != nil && r.RecipientName != nil {
		RecipientName = *r.RecipientName
	}
	return
}
