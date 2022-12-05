package simulations

type CardsPendingTransactionSourceWireDrawdownPaymentInstructionAuthorizeResponse struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	AccountNumber *string `json:"account_number,omitempty"`

	RoutingNumber *string `json:"routing_number,omitempty"`

	MessageToRecipient *string `json:"message_to_recipient,omitempty"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *CardsPendingTransactionSourceWireDrawdownPaymentInstructionAuthorizeResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *CardsPendingTransactionSourceWireDrawdownPaymentInstructionAuthorizeResponse) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *CardsPendingTransactionSourceWireDrawdownPaymentInstructionAuthorizeResponse) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *CardsPendingTransactionSourceWireDrawdownPaymentInstructionAuthorizeResponse) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}
