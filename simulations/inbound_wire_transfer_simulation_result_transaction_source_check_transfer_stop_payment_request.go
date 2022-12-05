package simulations

type InboundWireTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferId *string `json:"transfer_id,omitempty"`

	// The transaction ID of the corresponding credit transaction.
	TransactionId *string `json:"transaction_id,omitempty"`

	// The time the stop-payment was requested.
	RequestedAt *string `json:"requested_at,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type *string `json:"type,omitempty"`
}

// The ID of the check transfer that was stopped.
func (r *InboundWireTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

// The transaction ID of the corresponding credit transaction.
func (r *InboundWireTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

// The time the stop-payment was requested.
func (r *InboundWireTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest) GetRequestedAt() (RequestedAt string) {
	if r != nil && r.RequestedAt != nil {
		RequestedAt = *r.RequestedAt
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_stop_payment_request`.
func (r *InboundWireTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
