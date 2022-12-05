package pending_transactions

type PendingTransactionsSourceWireTransferInstructionRetrieveResponse struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	AccountNumber *string `json:"account_number,omitempty"`

	RoutingNumber *string `json:"routing_number,omitempty"`

	MessageToRecipient *string `json:"message_to_recipient,omitempty"`

	TransferId *string `json:"transfer_id,omitempty"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionsSourceWireTransferInstructionRetrieveResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *PendingTransactionsSourceWireTransferInstructionRetrieveResponse) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *PendingTransactionsSourceWireTransferInstructionRetrieveResponse) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *PendingTransactionsSourceWireTransferInstructionRetrieveResponse) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r *PendingTransactionsSourceWireTransferInstructionRetrieveResponse) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
