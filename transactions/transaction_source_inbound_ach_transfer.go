package transactions

type TransactionSourceInboundACHTransfer struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	OriginatorCompanyName *string `json:"originator_company_name,omitempty"`

	OriginatorCompanyDescriptiveDate *string `json:"originator_company_descriptive_date,omitempty"`

	OriginatorCompanyDiscretionaryData *string `json:"originator_company_discretionary_data,omitempty"`

	OriginatorCompanyEntryDescription *string `json:"originator_company_entry_description,omitempty"`

	OriginatorCompanyId *string `json:"originator_company_id,omitempty"`

	ReceiverIdNumber *string `json:"receiver_id_number,omitempty"`

	ReceiverName *string `json:"receiver_name,omitempty"`

	TraceNumber *string `json:"trace_number,omitempty"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *TransactionSourceInboundACHTransfer) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetOriginatorCompanyName() (OriginatorCompanyName string) {
	if r != nil && r.OriginatorCompanyName != nil {
		OriginatorCompanyName = *r.OriginatorCompanyName
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetOriginatorCompanyDescriptiveDate() (OriginatorCompanyDescriptiveDate string) {
	if r != nil && r.OriginatorCompanyDescriptiveDate != nil {
		OriginatorCompanyDescriptiveDate = *r.OriginatorCompanyDescriptiveDate
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetOriginatorCompanyDiscretionaryData() (OriginatorCompanyDiscretionaryData string) {
	if r != nil && r.OriginatorCompanyDiscretionaryData != nil {
		OriginatorCompanyDiscretionaryData = *r.OriginatorCompanyDiscretionaryData
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r != nil && r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetOriginatorCompanyId() (OriginatorCompanyId string) {
	if r != nil && r.OriginatorCompanyId != nil {
		OriginatorCompanyId = *r.OriginatorCompanyId
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetReceiverIdNumber() (ReceiverIdNumber string) {
	if r != nil && r.ReceiverIdNumber != nil {
		ReceiverIdNumber = *r.ReceiverIdNumber
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetReceiverName() (ReceiverName string) {
	if r != nil && r.ReceiverName != nil {
		ReceiverName = *r.ReceiverName
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}
