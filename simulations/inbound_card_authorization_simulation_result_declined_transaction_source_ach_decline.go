package simulations

type InboundCardAuthorizationSimulationResultDeclinedTransactionSourceACHDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	OriginatorCompanyName *string `json:"originator_company_name,omitempty"`

	OriginatorCompanyDescriptiveDate *string `json:"originator_company_descriptive_date,omitempty"`

	OriginatorCompanyDiscretionaryData *string `json:"originator_company_discretionary_data,omitempty"`

	OriginatorCompanyId *string `json:"originator_company_id,omitempty"`

	// Why the ACH transfer was declined.
	Reason *string `json:"reason,omitempty"`

	ReceiverIdNumber *string `json:"receiver_id_number,omitempty"`

	ReceiverName *string `json:"receiver_name,omitempty"`

	TraceNumber *string `json:"trace_number,omitempty"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *InboundCardAuthorizationSimulationResultDeclinedTransactionSourceACHDecline) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *InboundCardAuthorizationSimulationResultDeclinedTransactionSourceACHDecline) GetOriginatorCompanyName() (OriginatorCompanyName string) {
	if r != nil && r.OriginatorCompanyName != nil {
		OriginatorCompanyName = *r.OriginatorCompanyName
	}
	return
}

func (r *InboundCardAuthorizationSimulationResultDeclinedTransactionSourceACHDecline) GetOriginatorCompanyDescriptiveDate() (OriginatorCompanyDescriptiveDate string) {
	if r != nil && r.OriginatorCompanyDescriptiveDate != nil {
		OriginatorCompanyDescriptiveDate = *r.OriginatorCompanyDescriptiveDate
	}
	return
}

func (r *InboundCardAuthorizationSimulationResultDeclinedTransactionSourceACHDecline) GetOriginatorCompanyDiscretionaryData() (OriginatorCompanyDiscretionaryData string) {
	if r != nil && r.OriginatorCompanyDiscretionaryData != nil {
		OriginatorCompanyDiscretionaryData = *r.OriginatorCompanyDiscretionaryData
	}
	return
}

func (r *InboundCardAuthorizationSimulationResultDeclinedTransactionSourceACHDecline) GetOriginatorCompanyId() (OriginatorCompanyId string) {
	if r != nil && r.OriginatorCompanyId != nil {
		OriginatorCompanyId = *r.OriginatorCompanyId
	}
	return
}

// Why the ACH transfer was declined.
func (r *InboundCardAuthorizationSimulationResultDeclinedTransactionSourceACHDecline) GetReason() (Reason string) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

func (r *InboundCardAuthorizationSimulationResultDeclinedTransactionSourceACHDecline) GetReceiverIdNumber() (ReceiverIdNumber string) {
	if r != nil && r.ReceiverIdNumber != nil {
		ReceiverIdNumber = *r.ReceiverIdNumber
	}
	return
}

func (r *InboundCardAuthorizationSimulationResultDeclinedTransactionSourceACHDecline) GetReceiverName() (ReceiverName string) {
	if r != nil && r.ReceiverName != nil {
		ReceiverName = *r.ReceiverName
	}
	return
}

func (r *InboundCardAuthorizationSimulationResultDeclinedTransactionSourceACHDecline) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}
