package ach_prenotifications

type ACHPrenotification struct {
	// The ACH Prenotification's identifier.
	Id *string `json:"id,omitempty"`

	// The destination account number.
	AccountNumber *string `json:"account_number,omitempty"`

	// Additional information for the recipient.
	Addendum *string `json:"addendum,omitempty"`

	// The description of the date of the notification.
	CompanyDescriptiveDate *string `json:"company_descriptive_date,omitempty"`

	// Optional data associated with the notification.
	CompanyDiscretionaryData *string `json:"company_discretionary_data,omitempty"`

	// The description of the notification.
	CompanyEntryDescription *string `json:"company_entry_description,omitempty"`

	// The name by which you know the company.
	CompanyName *string `json:"company_name,omitempty"`

	// If the notification is for a future credit or debit.
	CreditDebitIndicator *string `json:"credit_debit_indicator,omitempty"`

	// The effective date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate *string `json:"effective_date,omitempty"`

	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber *string `json:"routing_number,omitempty"`

	// If your prenotification is returned, this will contain details of the return.
	PrenotificationReturn *ACHPrenotificationPrenotificationReturn `json:"prenotification_return,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the prenotification was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The lifecycle status of the ACH Prenotification.
	Status *string `json:"status,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `ach_prenotification`.
	Type *string `json:"type,omitempty"`
}

// The ACH Prenotification's identifier.
func (r *ACHPrenotification) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The destination account number.
func (r *ACHPrenotification) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// Additional information for the recipient.
func (r *ACHPrenotification) GetAddendum() (Addendum string) {
	if r != nil && r.Addendum != nil {
		Addendum = *r.Addendum
	}
	return
}

// The description of the date of the notification.
func (r *ACHPrenotification) GetCompanyDescriptiveDate() (CompanyDescriptiveDate string) {
	if r != nil && r.CompanyDescriptiveDate != nil {
		CompanyDescriptiveDate = *r.CompanyDescriptiveDate
	}
	return
}

// Optional data associated with the notification.
func (r *ACHPrenotification) GetCompanyDiscretionaryData() (CompanyDiscretionaryData string) {
	if r != nil && r.CompanyDiscretionaryData != nil {
		CompanyDiscretionaryData = *r.CompanyDiscretionaryData
	}
	return
}

// The description of the notification.
func (r *ACHPrenotification) GetCompanyEntryDescription() (CompanyEntryDescription string) {
	if r != nil && r.CompanyEntryDescription != nil {
		CompanyEntryDescription = *r.CompanyEntryDescription
	}
	return
}

// The name by which you know the company.
func (r *ACHPrenotification) GetCompanyName() (CompanyName string) {
	if r != nil && r.CompanyName != nil {
		CompanyName = *r.CompanyName
	}
	return
}

// If the notification is for a future credit or debit.
func (r *ACHPrenotification) GetCreditDebitIndicator() (CreditDebitIndicator string) {
	if r != nil && r.CreditDebitIndicator != nil {
		CreditDebitIndicator = *r.CreditDebitIndicator
	}
	return
}

// The effective date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
func (r *ACHPrenotification) GetEffectiveDate() (EffectiveDate string) {
	if r != nil && r.EffectiveDate != nil {
		EffectiveDate = *r.EffectiveDate
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r *ACHPrenotification) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// If your prenotification is returned, this will contain details of the return.
func (r *ACHPrenotification) GetPrenotificationReturn() (PrenotificationReturn ACHPrenotificationPrenotificationReturn) {
	if r != nil && r.PrenotificationReturn != nil {
		PrenotificationReturn = *r.PrenotificationReturn
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the prenotification was created.
func (r *ACHPrenotification) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The lifecycle status of the ACH Prenotification.
func (r *ACHPrenotification) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `ach_prenotification`.
func (r *ACHPrenotification) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
