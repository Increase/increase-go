package ach_transfers

type ACHTransfersCreateParams struct {
	// The identifier for the account that will send the transfer.
	AccountId *string `json:"account_id,omitempty"`

	// The account number for the destination account.
	AccountNumber *string `json:"account_number,omitempty"`

	// Additional information that will be sent to the recipient.
	Addendum *string `json:"addendum,omitempty"`

	// The transfer amount in cents. A positive amount originates a credit transfer
	// pushing funds to the receiving account. A negative amount originates a debit
	// transfer pulling funds from the receiving account.
	Amount *int64 `json:"amount,omitempty"`

	// The description of the date of the transfer.
	CompanyDescriptiveDate *string `json:"company_descriptive_date,omitempty"`

	// The data you choose to associate with the transfer.
	CompanyDiscretionaryData *string `json:"company_discretionary_data,omitempty"`

	// The description of the transfer you wish to be shown to the recipient.
	CompanyEntryDescription *string `json:"company_entry_description,omitempty"`

	// The name by which the recipient knows you.
	CompanyName *string `json:"company_name,omitempty"`

	// The transfer effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate *string `json:"effective_date,omitempty"`

	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `account_number`, `routing_number`, and `funding` must be absent.
	ExternalAccountId *string `json:"external_account_id,omitempty"`

	// The type of the account to which the transfer will be sent.
	Funding *string `json:"funding,omitempty"`

	// Your identifer for the transfer recipient.
	IndividualId *string `json:"individual_id,omitempty"`

	// The name of the transfer recipient. This value is information and not verified
	// by the recipient's bank.
	IndividualName *string `json:"individual_name,omitempty"`

	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber *string `json:"routing_number,omitempty"`

	// The Standard Entry Class (SEC) code to use for the transfer.
	StandardEntryClassCode *string `json:"standard_entry_class_code,omitempty"`

	// The description you choose to give the transfer. This will be shown to the
	// recipient.
	StatementDescriptor *string `json:"statement_descriptor,omitempty"`
}

// The identifier for the account that will send the transfer.
func (r *ACHTransfersCreateParams) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The account number for the destination account.
func (r *ACHTransfersCreateParams) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// Additional information that will be sent to the recipient.
func (r *ACHTransfersCreateParams) GetAddendum() (Addendum string) {
	if r != nil && r.Addendum != nil {
		Addendum = *r.Addendum
	}
	return
}

// The transfer amount in cents. A positive amount originates a credit transfer
// pushing funds to the receiving account. A negative amount originates a debit
// transfer pulling funds from the receiving account.
func (r *ACHTransfersCreateParams) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description of the date of the transfer.
func (r *ACHTransfersCreateParams) GetCompanyDescriptiveDate() (CompanyDescriptiveDate string) {
	if r != nil && r.CompanyDescriptiveDate != nil {
		CompanyDescriptiveDate = *r.CompanyDescriptiveDate
	}
	return
}

// The data you choose to associate with the transfer.
func (r *ACHTransfersCreateParams) GetCompanyDiscretionaryData() (CompanyDiscretionaryData string) {
	if r != nil && r.CompanyDiscretionaryData != nil {
		CompanyDiscretionaryData = *r.CompanyDiscretionaryData
	}
	return
}

// The description of the transfer you wish to be shown to the recipient.
func (r *ACHTransfersCreateParams) GetCompanyEntryDescription() (CompanyEntryDescription string) {
	if r != nil && r.CompanyEntryDescription != nil {
		CompanyEntryDescription = *r.CompanyEntryDescription
	}
	return
}

// The name by which the recipient knows you.
func (r *ACHTransfersCreateParams) GetCompanyName() (CompanyName string) {
	if r != nil && r.CompanyName != nil {
		CompanyName = *r.CompanyName
	}
	return
}

// The transfer effective date in
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
func (r *ACHTransfersCreateParams) GetEffectiveDate() (EffectiveDate string) {
	if r != nil && r.EffectiveDate != nil {
		EffectiveDate = *r.EffectiveDate
	}
	return
}

// The ID of an External Account to initiate a transfer to. If this parameter is
// provided, `account_number`, `routing_number`, and `funding` must be absent.
func (r *ACHTransfersCreateParams) GetExternalAccountId() (ExternalAccountId string) {
	if r != nil && r.ExternalAccountId != nil {
		ExternalAccountId = *r.ExternalAccountId
	}
	return
}

// The type of the account to which the transfer will be sent.
func (r *ACHTransfersCreateParams) GetFunding() (Funding string) {
	if r != nil && r.Funding != nil {
		Funding = *r.Funding
	}
	return
}

// Your identifer for the transfer recipient.
func (r *ACHTransfersCreateParams) GetIndividualId() (IndividualId string) {
	if r != nil && r.IndividualId != nil {
		IndividualId = *r.IndividualId
	}
	return
}

// The name of the transfer recipient. This value is information and not verified
// by the recipient's bank.
func (r *ACHTransfersCreateParams) GetIndividualName() (IndividualName string) {
	if r != nil && r.IndividualName != nil {
		IndividualName = *r.IndividualName
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
// destination account.
func (r *ACHTransfersCreateParams) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The Standard Entry Class (SEC) code to use for the transfer.
func (r *ACHTransfersCreateParams) GetStandardEntryClassCode() (StandardEntryClassCode string) {
	if r != nil && r.StandardEntryClassCode != nil {
		StandardEntryClassCode = *r.StandardEntryClassCode
	}
	return
}

// The description you choose to give the transfer. This will be shown to the
// recipient.
func (r *ACHTransfersCreateParams) GetStatementDescriptor() (StatementDescriptor string) {
	if r != nil && r.StatementDescriptor != nil {
		StatementDescriptor = *r.StatementDescriptor
	}
	return
}
