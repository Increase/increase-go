package pending_transactions

type PendingTransactionsSourceWireDrawdownPaymentInstructionRetrieveResponse struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	AccountNumber *string `json:"account_number,omitempty"`

	RoutingNumber *string `json:"routing_number,omitempty"`

	MessageToRecipient *string `json:"message_to_recipient,omitempty"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionsSourceWireDrawdownPaymentInstructionRetrieveResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *PendingTransactionsSourceWireDrawdownPaymentInstructionRetrieveResponse) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *PendingTransactionsSourceWireDrawdownPaymentInstructionRetrieveResponse) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *PendingTransactionsSourceWireDrawdownPaymentInstructionRetrieveResponse) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}
