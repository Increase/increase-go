package simulations

type ACHTransfersTransactionSourceWireDrawdownPaymentIntentionCreateInboundResponse struct {
	// The transfer amount in USD cents.
	Amount *int64 `json:"amount,omitempty"`

	AccountNumber *string `json:"account_number,omitempty"`

	RoutingNumber *string `json:"routing_number,omitempty"`

	MessageToRecipient *string `json:"message_to_recipient,omitempty"`

	TransferId *string `json:"transfer_id,omitempty"`
}

// The transfer amount in USD cents.
func (r *ACHTransfersTransactionSourceWireDrawdownPaymentIntentionCreateInboundResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransfersTransactionSourceWireDrawdownPaymentIntentionCreateInboundResponse) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *ACHTransfersTransactionSourceWireDrawdownPaymentIntentionCreateInboundResponse) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *ACHTransfersTransactionSourceWireDrawdownPaymentIntentionCreateInboundResponse) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r *ACHTransfersTransactionSourceWireDrawdownPaymentIntentionCreateInboundResponse) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
