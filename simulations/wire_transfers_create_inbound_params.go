package simulations

type WireTransfersCreateInboundParams struct {
	// The identifier of the Account Number the inbound Wire Transfer is for.
	AccountNumberId *string `json:"account_number_id,omitempty"`

	// The transfer amount in cents. Must be positive.
	Amount *int64 `json:"amount,omitempty"`
}

// The identifier of the Account Number the inbound Wire Transfer is for.
func (r *WireTransfersCreateInboundParams) GetAccountNumberId() (AccountNumberId string) {
	if r != nil && r.AccountNumberId != nil {
		AccountNumberId = *r.AccountNumberId
	}
	return
}

// The transfer amount in cents. Must be positive.
func (r *WireTransfersCreateInboundParams) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}
