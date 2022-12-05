package simulations

type InboundCardAuthorizationSimulationResultPendingTransactionSourceWireDrawdownPaymentInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	AccountNumber *string `json:"account_number,omitempty"`

	RoutingNumber *string `json:"routing_number,omitempty"`

	MessageToRecipient *string `json:"message_to_recipient,omitempty"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *InboundCardAuthorizationSimulationResultPendingTransactionSourceWireDrawdownPaymentInstruction) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *InboundCardAuthorizationSimulationResultPendingTransactionSourceWireDrawdownPaymentInstruction) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *InboundCardAuthorizationSimulationResultPendingTransactionSourceWireDrawdownPaymentInstruction) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *InboundCardAuthorizationSimulationResultPendingTransactionSourceWireDrawdownPaymentInstruction) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}
