package transactions

type TransactionsDataSourceInboundWireDrawdownPaymentListResponse struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	BeneficiaryAddressLine1 *string `json:"beneficiary_address_line1,omitempty"`

	BeneficiaryAddressLine2 *string `json:"beneficiary_address_line2,omitempty"`

	BeneficiaryAddressLine3 *string `json:"beneficiary_address_line3,omitempty"`

	BeneficiaryName *string `json:"beneficiary_name,omitempty"`

	BeneficiaryReference *string `json:"beneficiary_reference,omitempty"`

	Description *string `json:"description,omitempty"`

	InputMessageAccountabilityData *string `json:"input_message_accountability_data,omitempty"`

	OriginatorAddressLine1 *string `json:"originator_address_line1,omitempty"`

	OriginatorAddressLine2 *string `json:"originator_address_line2,omitempty"`

	OriginatorAddressLine3 *string `json:"originator_address_line3,omitempty"`

	OriginatorName *string `json:"originator_name,omitempty"`

	OriginatorToBeneficiaryInformation *string `json:"originator_to_beneficiary_information,omitempty"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionsDataSourceInboundWireDrawdownPaymentListResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *TransactionsDataSourceInboundWireDrawdownPaymentListResponse) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r != nil && r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

func (r *TransactionsDataSourceInboundWireDrawdownPaymentListResponse) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r != nil && r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

func (r *TransactionsDataSourceInboundWireDrawdownPaymentListResponse) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r != nil && r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

func (r *TransactionsDataSourceInboundWireDrawdownPaymentListResponse) GetBeneficiaryName() (BeneficiaryName string) {
	if r != nil && r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

func (r *TransactionsDataSourceInboundWireDrawdownPaymentListResponse) GetBeneficiaryReference() (BeneficiaryReference string) {
	if r != nil && r.BeneficiaryReference != nil {
		BeneficiaryReference = *r.BeneficiaryReference
	}
	return
}

func (r *TransactionsDataSourceInboundWireDrawdownPaymentListResponse) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

func (r *TransactionsDataSourceInboundWireDrawdownPaymentListResponse) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

func (r *TransactionsDataSourceInboundWireDrawdownPaymentListResponse) GetOriginatorAddressLine1() (OriginatorAddressLine1 string) {
	if r != nil && r.OriginatorAddressLine1 != nil {
		OriginatorAddressLine1 = *r.OriginatorAddressLine1
	}
	return
}

func (r *TransactionsDataSourceInboundWireDrawdownPaymentListResponse) GetOriginatorAddressLine2() (OriginatorAddressLine2 string) {
	if r != nil && r.OriginatorAddressLine2 != nil {
		OriginatorAddressLine2 = *r.OriginatorAddressLine2
	}
	return
}

func (r *TransactionsDataSourceInboundWireDrawdownPaymentListResponse) GetOriginatorAddressLine3() (OriginatorAddressLine3 string) {
	if r != nil && r.OriginatorAddressLine3 != nil {
		OriginatorAddressLine3 = *r.OriginatorAddressLine3
	}
	return
}

func (r *TransactionsDataSourceInboundWireDrawdownPaymentListResponse) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *TransactionsDataSourceInboundWireDrawdownPaymentListResponse) GetOriginatorToBeneficiaryInformation() (OriginatorToBeneficiaryInformation string) {
	if r != nil && r.OriginatorToBeneficiaryInformation != nil {
		OriginatorToBeneficiaryInformation = *r.OriginatorToBeneficiaryInformation
	}
	return
}
