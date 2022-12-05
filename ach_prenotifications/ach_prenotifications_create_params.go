package ach_prenotifications

type ACHPrenotificationsCreateParams struct {
	// The account number for the destination account.
	AccountNumber *string `json:"account_number,omitempty"`

	// Additional information that will be sent to the recipient.
	Addendum *string `json:"addendum,omitempty"`

	// The description of the date of the transfer.
	CompanyDescriptiveDate *string `json:"company_descriptive_date,omitempty"`

	// The data you choose to associate with the transfer.
	CompanyDiscretionaryData *string `json:"company_discretionary_data,omitempty"`

	// The description of the transfer you wish to be shown to the recipient.
	CompanyEntryDescription *string `json:"company_entry_description,omitempty"`

	// The name by which the recipient knows you.
	CompanyName *string `json:"company_name,omitempty"`

	// Whether the Prenotification is for a future debit or credit.
	CreditDebitIndicator *string `json:"credit_debit_indicator,omitempty"`

	// The transfer effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate *string `json:"effective_date,omitempty"`

	// Your identifer for the transfer recipient.
	IndividualId *string `json:"individual_id,omitempty"`

	// The name of the transfer recipient. This value is information and not verified
	// by the recipient's bank.
	IndividualName *string `json:"individual_name,omitempty"`

	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber *string `json:"routing_number,omitempty"`

	// The Standard Entry Class (SEC) code to use for the ACH Prenotification.
	StandardEntryClassCode *string `json:"standard_entry_class_code,omitempty"`
}

// The account number for the destination account.
func (r *ACHPrenotificationsCreateParams) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// Additional information that will be sent to the recipient.
func (r *ACHPrenotificationsCreateParams) GetAddendum() (Addendum string) {
	if r != nil && r.Addendum != nil {
		Addendum = *r.Addendum
	}
	return
}

// The description of the date of the transfer.
func (r *ACHPrenotificationsCreateParams) GetCompanyDescriptiveDate() (CompanyDescriptiveDate string) {
	if r != nil && r.CompanyDescriptiveDate != nil {
		CompanyDescriptiveDate = *r.CompanyDescriptiveDate
	}
	return
}

// The data you choose to associate with the transfer.
func (r *ACHPrenotificationsCreateParams) GetCompanyDiscretionaryData() (CompanyDiscretionaryData string) {
	if r != nil && r.CompanyDiscretionaryData != nil {
		CompanyDiscretionaryData = *r.CompanyDiscretionaryData
	}
	return
}

// The description of the transfer you wish to be shown to the recipient.
func (r *ACHPrenotificationsCreateParams) GetCompanyEntryDescription() (CompanyEntryDescription string) {
	if r != nil && r.CompanyEntryDescription != nil {
		CompanyEntryDescription = *r.CompanyEntryDescription
	}
	return
}

// The name by which the recipient knows you.
func (r *ACHPrenotificationsCreateParams) GetCompanyName() (CompanyName string) {
	if r != nil && r.CompanyName != nil {
		CompanyName = *r.CompanyName
	}
	return
}

// Whether the Prenotification is for a future debit or credit.
func (r *ACHPrenotificationsCreateParams) GetCreditDebitIndicator() (CreditDebitIndicator string) {
	if r != nil && r.CreditDebitIndicator != nil {
		CreditDebitIndicator = *r.CreditDebitIndicator
	}
	return
}

// The transfer effective date in
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
func (r *ACHPrenotificationsCreateParams) GetEffectiveDate() (EffectiveDate string) {
	if r != nil && r.EffectiveDate != nil {
		EffectiveDate = *r.EffectiveDate
	}
	return
}

// Your identifer for the transfer recipient.
func (r *ACHPrenotificationsCreateParams) GetIndividualId() (IndividualId string) {
	if r != nil && r.IndividualId != nil {
		IndividualId = *r.IndividualId
	}
	return
}

// The name of the transfer recipient. This value is information and not verified
// by the recipient's bank.
func (r *ACHPrenotificationsCreateParams) GetIndividualName() (IndividualName string) {
	if r != nil && r.IndividualName != nil {
		IndividualName = *r.IndividualName
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
// destination account.
func (r *ACHPrenotificationsCreateParams) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The Standard Entry Class (SEC) code to use for the ACH Prenotification.
func (r *ACHPrenotificationsCreateParams) GetStandardEntryClassCode() (StandardEntryClassCode string) {
	if r != nil && r.StandardEntryClassCode != nil {
		StandardEntryClassCode = *r.StandardEntryClassCode
	}
	return
}
