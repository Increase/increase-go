package wire_transfers

type WireTransfersCreateParams struct {
	// The identifier for the account that will send the transfer.
	AccountId *string `json:"account_id,omitempty"`

	// The account number for the destination account.
	AccountNumber *string `json:"account_number,omitempty"`

	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber *string `json:"routing_number,omitempty"`

	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `account_number` and `routing_number` must be absent.
	ExternalAccountId *string `json:"external_account_id,omitempty"`

	// The transfer amount in cents.
	Amount *int64 `json:"amount,omitempty"`

	// The message that will show on the recipient's bank statement.
	MessageToRecipient *string `json:"message_to_recipient,omitempty"`

	// The beneficiary's name.
	BeneficiaryName *string `json:"beneficiary_name,omitempty"`

	// The beneficiary's address line 1.
	BeneficiaryAddressLine1 *string `json:"beneficiary_address_line1,omitempty"`

	// The beneficiary's address line 2.
	BeneficiaryAddressLine2 *string `json:"beneficiary_address_line2,omitempty"`

	// The beneficiary's address line 3.
	BeneficiaryAddressLine3 *string `json:"beneficiary_address_line3,omitempty"`
}

// The identifier for the account that will send the transfer.
func (r *WireTransfersCreateParams) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The account number for the destination account.
func (r *WireTransfersCreateParams) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
// destination account.
func (r *WireTransfersCreateParams) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The ID of an External Account to initiate a transfer to. If this parameter is
// provided, `account_number` and `routing_number` must be absent.
func (r *WireTransfersCreateParams) GetExternalAccountId() (ExternalAccountId string) {
	if r != nil && r.ExternalAccountId != nil {
		ExternalAccountId = *r.ExternalAccountId
	}
	return
}

// The transfer amount in cents.
func (r *WireTransfersCreateParams) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The message that will show on the recipient's bank statement.
func (r *WireTransfersCreateParams) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

// The beneficiary's name.
func (r *WireTransfersCreateParams) GetBeneficiaryName() (BeneficiaryName string) {
	if r != nil && r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

// The beneficiary's address line 1.
func (r *WireTransfersCreateParams) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r != nil && r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

// The beneficiary's address line 2.
func (r *WireTransfersCreateParams) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r != nil && r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

// The beneficiary's address line 3.
func (r *WireTransfersCreateParams) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r != nil && r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}
