package transactions

type TransactionsSourceInboundWireReversalRetrieveResponse struct {
	// The amount that was reversed.
	Amount *int64 `json:"amount,omitempty"`

	// The description on the reversal message from Fedwire.
	Description *string `json:"description,omitempty"`

	// The Fedwire cycle date for the wire reversal.
	InputCycleDate *string `json:"input_cycle_date,omitempty"`

	// The Fedwire sequence number.
	InputSequenceNumber *string `json:"input_sequence_number,omitempty"`

	// The Fedwire input source identifier.
	InputSource *string `json:"input_source,omitempty"`

	// The Fedwire transaction identifier.
	InputMessageAccountabilityData *string `json:"input_message_accountability_data,omitempty"`

	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData *string `json:"previous_message_input_message_accountability_data,omitempty"`

	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate *string `json:"previous_message_input_cycle_date,omitempty"`

	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber *string `json:"previous_message_input_sequence_number,omitempty"`

	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource *string `json:"previous_message_input_source,omitempty"`

	// Information included in the wire reversal for the receiving financial
	// institution.
	ReceiverFinancialInstitutionInformation *string `json:"receiver_financial_institution_information,omitempty"`

	// Additional financial institution information included in the wire reversal.
	FinancialInstitutionToFinancialInstitutionInformation *string `json:"financial_institution_to_financial_institution_information,omitempty"`
}

// The amount that was reversed.
func (r *TransactionsSourceInboundWireReversalRetrieveResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description on the reversal message from Fedwire.
func (r *TransactionsSourceInboundWireReversalRetrieveResponse) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Fedwire cycle date for the wire reversal.
func (r *TransactionsSourceInboundWireReversalRetrieveResponse) GetInputCycleDate() (InputCycleDate string) {
	if r != nil && r.InputCycleDate != nil {
		InputCycleDate = *r.InputCycleDate
	}
	return
}

// The Fedwire sequence number.
func (r *TransactionsSourceInboundWireReversalRetrieveResponse) GetInputSequenceNumber() (InputSequenceNumber string) {
	if r != nil && r.InputSequenceNumber != nil {
		InputSequenceNumber = *r.InputSequenceNumber
	}
	return
}

// The Fedwire input source identifier.
func (r *TransactionsSourceInboundWireReversalRetrieveResponse) GetInputSource() (InputSource string) {
	if r != nil && r.InputSource != nil {
		InputSource = *r.InputSource
	}
	return
}

// The Fedwire transaction identifier.
func (r *TransactionsSourceInboundWireReversalRetrieveResponse) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

// The Fedwire transaction identifier for the wire transfer that was reversed.
func (r *TransactionsSourceInboundWireReversalRetrieveResponse) GetPreviousMessageInputMessageAccountabilityData() (PreviousMessageInputMessageAccountabilityData string) {
	if r != nil && r.PreviousMessageInputMessageAccountabilityData != nil {
		PreviousMessageInputMessageAccountabilityData = *r.PreviousMessageInputMessageAccountabilityData
	}
	return
}

// The Fedwire cycle date for the wire transfer that was reversed.
func (r *TransactionsSourceInboundWireReversalRetrieveResponse) GetPreviousMessageInputCycleDate() (PreviousMessageInputCycleDate string) {
	if r != nil && r.PreviousMessageInputCycleDate != nil {
		PreviousMessageInputCycleDate = *r.PreviousMessageInputCycleDate
	}
	return
}

// The Fedwire sequence number for the wire transfer that was reversed.
func (r *TransactionsSourceInboundWireReversalRetrieveResponse) GetPreviousMessageInputSequenceNumber() (PreviousMessageInputSequenceNumber string) {
	if r != nil && r.PreviousMessageInputSequenceNumber != nil {
		PreviousMessageInputSequenceNumber = *r.PreviousMessageInputSequenceNumber
	}
	return
}

// The Fedwire input source identifier for the wire transfer that was reversed.
func (r *TransactionsSourceInboundWireReversalRetrieveResponse) GetPreviousMessageInputSource() (PreviousMessageInputSource string) {
	if r != nil && r.PreviousMessageInputSource != nil {
		PreviousMessageInputSource = *r.PreviousMessageInputSource
	}
	return
}

// Information included in the wire reversal for the receiving financial
// institution.
func (r *TransactionsSourceInboundWireReversalRetrieveResponse) GetReceiverFinancialInstitutionInformation() (ReceiverFinancialInstitutionInformation string) {
	if r != nil && r.ReceiverFinancialInstitutionInformation != nil {
		ReceiverFinancialInstitutionInformation = *r.ReceiverFinancialInstitutionInformation
	}
	return
}

// Additional financial institution information included in the wire reversal.
func (r *TransactionsSourceInboundWireReversalRetrieveResponse) GetFinancialInstitutionToFinancialInstitutionInformation() (FinancialInstitutionToFinancialInstitutionInformation string) {
	if r != nil && r.FinancialInstitutionToFinancialInstitutionInformation != nil {
		FinancialInstitutionToFinancialInstitutionInformation = *r.FinancialInstitutionToFinancialInstitutionInformation
	}
	return
}
