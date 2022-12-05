package simulations

type RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	ForeignExchangeIndicator *string `json:"foreign_exchange_indicator,omitempty"`

	ForeignExchangeReferenceIndicator *string `json:"foreign_exchange_reference_indicator,omitempty"`

	ForeignExchangeReference *string `json:"foreign_exchange_reference,omitempty"`

	DestinationCountryCode *string `json:"destination_country_code,omitempty"`

	DestinationCurrencyCode *string `json:"destination_currency_code,omitempty"`

	ForeignPaymentAmount *int64 `json:"foreign_payment_amount,omitempty"`

	ForeignTraceNumber *string `json:"foreign_trace_number,omitempty"`

	InternationalTransactionTypeCode *string `json:"international_transaction_type_code,omitempty"`

	OriginatingCurrencyCode *string `json:"originating_currency_code,omitempty"`

	OriginatingDepositoryFinancialInstitutionName *string `json:"originating_depository_financial_institution_name,omitempty"`

	OriginatingDepositoryFinancialInstitutionIdQualifier *string `json:"originating_depository_financial_institution_id_qualifier,omitempty"`

	OriginatingDepositoryFinancialInstitutionId *string `json:"originating_depository_financial_institution_id,omitempty"`

	OriginatingDepositoryFinancialInstitutionBranchCountry *string `json:"originating_depository_financial_institution_branch_country,omitempty"`

	OriginatorCity *string `json:"originator_city,omitempty"`

	OriginatorCompanyEntryDescription *string `json:"originator_company_entry_description,omitempty"`

	OriginatorCountry *string `json:"originator_country,omitempty"`

	OriginatorIdentification *string `json:"originator_identification,omitempty"`

	OriginatorName *string `json:"originator_name,omitempty"`

	OriginatorPostalCode *string `json:"originator_postal_code,omitempty"`

	OriginatorStreetAddress *string `json:"originator_street_address,omitempty"`

	OriginatorStateOrProvince *string `json:"originator_state_or_province,omitempty"`

	PaymentRelatedInformation *string `json:"payment_related_information,omitempty"`

	PaymentRelatedInformation2 *string `json:"payment_related_information2,omitempty"`

	ReceiverIdentificationNumber *string `json:"receiver_identification_number,omitempty"`

	ReceiverStreetAddress *string `json:"receiver_street_address,omitempty"`

	ReceiverCity *string `json:"receiver_city,omitempty"`

	ReceiverStateOrProvince *string `json:"receiver_state_or_province,omitempty"`

	ReceiverCountry *string `json:"receiver_country,omitempty"`

	ReceiverPostalCode *string `json:"receiver_postal_code,omitempty"`

	ReceivingCompanyOrIndividualName *string `json:"receiving_company_or_individual_name,omitempty"`

	ReceivingDepositoryFinancialInstitutionName *string `json:"receiving_depository_financial_institution_name,omitempty"`

	ReceivingDepositoryFinancialInstitutionIdQualifier *string `json:"receiving_depository_financial_institution_id_qualifier,omitempty"`

	ReceivingDepositoryFinancialInstitutionId *string `json:"receiving_depository_financial_institution_id,omitempty"`

	ReceivingDepositoryFinancialInstitutionCountry *string `json:"receiving_depository_financial_institution_country,omitempty"`

	TraceNumber *string `json:"trace_number,omitempty"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetForeignExchangeIndicator() (ForeignExchangeIndicator string) {
	if r != nil && r.ForeignExchangeIndicator != nil {
		ForeignExchangeIndicator = *r.ForeignExchangeIndicator
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetForeignExchangeReferenceIndicator() (ForeignExchangeReferenceIndicator string) {
	if r != nil && r.ForeignExchangeReferenceIndicator != nil {
		ForeignExchangeReferenceIndicator = *r.ForeignExchangeReferenceIndicator
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetForeignExchangeReference() (ForeignExchangeReference string) {
	if r != nil && r.ForeignExchangeReference != nil {
		ForeignExchangeReference = *r.ForeignExchangeReference
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetDestinationCountryCode() (DestinationCountryCode string) {
	if r != nil && r.DestinationCountryCode != nil {
		DestinationCountryCode = *r.DestinationCountryCode
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetDestinationCurrencyCode() (DestinationCurrencyCode string) {
	if r != nil && r.DestinationCurrencyCode != nil {
		DestinationCurrencyCode = *r.DestinationCurrencyCode
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetForeignPaymentAmount() (ForeignPaymentAmount int64) {
	if r != nil && r.ForeignPaymentAmount != nil {
		ForeignPaymentAmount = *r.ForeignPaymentAmount
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetForeignTraceNumber() (ForeignTraceNumber string) {
	if r != nil && r.ForeignTraceNumber != nil {
		ForeignTraceNumber = *r.ForeignTraceNumber
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetInternationalTransactionTypeCode() (InternationalTransactionTypeCode string) {
	if r != nil && r.InternationalTransactionTypeCode != nil {
		InternationalTransactionTypeCode = *r.InternationalTransactionTypeCode
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetOriginatingCurrencyCode() (OriginatingCurrencyCode string) {
	if r != nil && r.OriginatingCurrencyCode != nil {
		OriginatingCurrencyCode = *r.OriginatingCurrencyCode
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetOriginatingDepositoryFinancialInstitutionName() (OriginatingDepositoryFinancialInstitutionName string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionName != nil {
		OriginatingDepositoryFinancialInstitutionName = *r.OriginatingDepositoryFinancialInstitutionName
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetOriginatingDepositoryFinancialInstitutionIdQualifier() (OriginatingDepositoryFinancialInstitutionIdQualifier string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionIdQualifier != nil {
		OriginatingDepositoryFinancialInstitutionIdQualifier = *r.OriginatingDepositoryFinancialInstitutionIdQualifier
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetOriginatingDepositoryFinancialInstitutionId() (OriginatingDepositoryFinancialInstitutionId string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionId != nil {
		OriginatingDepositoryFinancialInstitutionId = *r.OriginatingDepositoryFinancialInstitutionId
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetOriginatingDepositoryFinancialInstitutionBranchCountry() (OriginatingDepositoryFinancialInstitutionBranchCountry string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionBranchCountry != nil {
		OriginatingDepositoryFinancialInstitutionBranchCountry = *r.OriginatingDepositoryFinancialInstitutionBranchCountry
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetOriginatorCity() (OriginatorCity string) {
	if r != nil && r.OriginatorCity != nil {
		OriginatorCity = *r.OriginatorCity
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r != nil && r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetOriginatorCountry() (OriginatorCountry string) {
	if r != nil && r.OriginatorCountry != nil {
		OriginatorCountry = *r.OriginatorCountry
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetOriginatorIdentification() (OriginatorIdentification string) {
	if r != nil && r.OriginatorIdentification != nil {
		OriginatorIdentification = *r.OriginatorIdentification
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetOriginatorPostalCode() (OriginatorPostalCode string) {
	if r != nil && r.OriginatorPostalCode != nil {
		OriginatorPostalCode = *r.OriginatorPostalCode
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetOriginatorStreetAddress() (OriginatorStreetAddress string) {
	if r != nil && r.OriginatorStreetAddress != nil {
		OriginatorStreetAddress = *r.OriginatorStreetAddress
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetOriginatorStateOrProvince() (OriginatorStateOrProvince string) {
	if r != nil && r.OriginatorStateOrProvince != nil {
		OriginatorStateOrProvince = *r.OriginatorStateOrProvince
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetPaymentRelatedInformation() (PaymentRelatedInformation string) {
	if r != nil && r.PaymentRelatedInformation != nil {
		PaymentRelatedInformation = *r.PaymentRelatedInformation
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetPaymentRelatedInformation2() (PaymentRelatedInformation2 string) {
	if r != nil && r.PaymentRelatedInformation2 != nil {
		PaymentRelatedInformation2 = *r.PaymentRelatedInformation2
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetReceiverIdentificationNumber() (ReceiverIdentificationNumber string) {
	if r != nil && r.ReceiverIdentificationNumber != nil {
		ReceiverIdentificationNumber = *r.ReceiverIdentificationNumber
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetReceiverStreetAddress() (ReceiverStreetAddress string) {
	if r != nil && r.ReceiverStreetAddress != nil {
		ReceiverStreetAddress = *r.ReceiverStreetAddress
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetReceiverCity() (ReceiverCity string) {
	if r != nil && r.ReceiverCity != nil {
		ReceiverCity = *r.ReceiverCity
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetReceiverStateOrProvince() (ReceiverStateOrProvince string) {
	if r != nil && r.ReceiverStateOrProvince != nil {
		ReceiverStateOrProvince = *r.ReceiverStateOrProvince
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetReceiverCountry() (ReceiverCountry string) {
	if r != nil && r.ReceiverCountry != nil {
		ReceiverCountry = *r.ReceiverCountry
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetReceiverPostalCode() (ReceiverPostalCode string) {
	if r != nil && r.ReceiverPostalCode != nil {
		ReceiverPostalCode = *r.ReceiverPostalCode
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetReceivingCompanyOrIndividualName() (ReceivingCompanyOrIndividualName string) {
	if r != nil && r.ReceivingCompanyOrIndividualName != nil {
		ReceivingCompanyOrIndividualName = *r.ReceivingCompanyOrIndividualName
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetReceivingDepositoryFinancialInstitutionName() (ReceivingDepositoryFinancialInstitutionName string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionName != nil {
		ReceivingDepositoryFinancialInstitutionName = *r.ReceivingDepositoryFinancialInstitutionName
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetReceivingDepositoryFinancialInstitutionIdQualifier() (ReceivingDepositoryFinancialInstitutionIdQualifier string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionIdQualifier != nil {
		ReceivingDepositoryFinancialInstitutionIdQualifier = *r.ReceivingDepositoryFinancialInstitutionIdQualifier
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetReceivingDepositoryFinancialInstitutionId() (ReceivingDepositoryFinancialInstitutionId string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionId != nil {
		ReceivingDepositoryFinancialInstitutionId = *r.ReceivingDepositoryFinancialInstitutionId
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetReceivingDepositoryFinancialInstitutionCountry() (ReceivingDepositoryFinancialInstitutionCountry string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionCountry != nil {
		ReceivingDepositoryFinancialInstitutionCountry = *r.ReceivingDepositoryFinancialInstitutionCountry
	}
	return
}

func (r *RealTimePaymentsTransfersDeclinedTransactionSourceInternationalACHDeclineCreateInboundResponse) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}
