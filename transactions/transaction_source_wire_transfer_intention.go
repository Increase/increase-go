package transactions

type TransactionSourceWireTransferIntention struct {
	// The transfer amount in USD cents.
	Amount *int64 `json:"amount,omitempty"`

	// The destination account number.
	AccountNumber *string `json:"account_number,omitempty"`

	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber *string `json:"routing_number,omitempty"`

	// The message that will show on the recipient's bank statement.
	MessageToRecipient *string `json:"message_to_recipient,omitempty"`

	TransferId *string `json:"transfer_id,omitempty"`
}

// The transfer amount in USD cents.
func (r *TransactionSourceWireTransferIntention) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The destination account number.
func (r *TransactionSourceWireTransferIntention) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r *TransactionSourceWireTransferIntention) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The message that will show on the recipient's bank statement.
func (r *TransactionSourceWireTransferIntention) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r *TransactionSourceWireTransferIntention) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
