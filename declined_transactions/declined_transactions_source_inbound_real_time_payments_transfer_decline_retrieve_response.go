package declined_transactions

type DeclinedTransactionsSourceInboundRealTimePaymentsTransferDeclineRetrieveResponse struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real Time Payments
	// transfer.
	Currency *string `json:"currency,omitempty"`

	// Why the transfer was declined.
	Reason *string `json:"reason,omitempty"`

	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName *string `json:"creditor_name,omitempty"`

	// The name provided by the sender of the transfer.
	DebtorName *string `json:"debtor_name,omitempty"`

	// The account number of the account that sent the transfer.
	DebtorAccountNumber *string `json:"debtor_account_number,omitempty"`

	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber *string `json:"debtor_routing_number,omitempty"`

	// The Real Time Payments network identification of the declined transfer.
	TransactionIdentification *string `json:"transaction_identification,omitempty"`

	// Additional information included with the transfer.
	RemittanceInformation *string `json:"remittance_information,omitempty"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *DeclinedTransactionsSourceInboundRealTimePaymentsTransferDeclineRetrieveResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
// transfer's currency. This will always be "USD" for a Real Time Payments
// transfer.
func (r *DeclinedTransactionsSourceInboundRealTimePaymentsTransferDeclineRetrieveResponse) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// Why the transfer was declined.
func (r *DeclinedTransactionsSourceInboundRealTimePaymentsTransferDeclineRetrieveResponse) GetReason() (Reason string) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

// The name the sender of the transfer specified as the recipient of the transfer.
func (r *DeclinedTransactionsSourceInboundRealTimePaymentsTransferDeclineRetrieveResponse) GetCreditorName() (CreditorName string) {
	if r != nil && r.CreditorName != nil {
		CreditorName = *r.CreditorName
	}
	return
}

// The name provided by the sender of the transfer.
func (r *DeclinedTransactionsSourceInboundRealTimePaymentsTransferDeclineRetrieveResponse) GetDebtorName() (DebtorName string) {
	if r != nil && r.DebtorName != nil {
		DebtorName = *r.DebtorName
	}
	return
}

// The account number of the account that sent the transfer.
func (r *DeclinedTransactionsSourceInboundRealTimePaymentsTransferDeclineRetrieveResponse) GetDebtorAccountNumber() (DebtorAccountNumber string) {
	if r != nil && r.DebtorAccountNumber != nil {
		DebtorAccountNumber = *r.DebtorAccountNumber
	}
	return
}

// The routing number of the account that sent the transfer.
func (r *DeclinedTransactionsSourceInboundRealTimePaymentsTransferDeclineRetrieveResponse) GetDebtorRoutingNumber() (DebtorRoutingNumber string) {
	if r != nil && r.DebtorRoutingNumber != nil {
		DebtorRoutingNumber = *r.DebtorRoutingNumber
	}
	return
}

// The Real Time Payments network identification of the declined transfer.
func (r *DeclinedTransactionsSourceInboundRealTimePaymentsTransferDeclineRetrieveResponse) GetTransactionIdentification() (TransactionIdentification string) {
	if r != nil && r.TransactionIdentification != nil {
		TransactionIdentification = *r.TransactionIdentification
	}
	return
}

// Additional information included with the transfer.
func (r *DeclinedTransactionsSourceInboundRealTimePaymentsTransferDeclineRetrieveResponse) GetRemittanceInformation() (RemittanceInformation string) {
	if r != nil && r.RemittanceInformation != nil {
		RemittanceInformation = *r.RemittanceInformation
	}
	return
}
