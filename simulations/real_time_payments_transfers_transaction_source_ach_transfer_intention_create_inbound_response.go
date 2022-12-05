package simulations

type RealTimePaymentsTransfersTransactionSourceACHTransferIntentionCreateInboundResponse struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	AccountNumber *string `json:"account_number,omitempty"`

	RoutingNumber *string `json:"routing_number,omitempty"`

	StatementDescriptor *string `json:"statement_descriptor,omitempty"`

	// The identifier of the ACH Transfer that led to this Transaction.
	TransferId *string `json:"transfer_id,omitempty"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *RealTimePaymentsTransfersTransactionSourceACHTransferIntentionCreateInboundResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *RealTimePaymentsTransfersTransactionSourceACHTransferIntentionCreateInboundResponse) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *RealTimePaymentsTransfersTransactionSourceACHTransferIntentionCreateInboundResponse) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *RealTimePaymentsTransfersTransactionSourceACHTransferIntentionCreateInboundResponse) GetStatementDescriptor() (StatementDescriptor string) {
	if r != nil && r.StatementDescriptor != nil {
		StatementDescriptor = *r.StatementDescriptor
	}
	return
}

// The identifier of the ACH Transfer that led to this Transaction.
func (r *RealTimePaymentsTransfersTransactionSourceACHTransferIntentionCreateInboundResponse) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
