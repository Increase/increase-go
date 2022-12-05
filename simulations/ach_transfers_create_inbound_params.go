package simulations

type ACHTransfersCreateInboundParams struct {
	// The identifier of the Account Number the inbound ACH Transfer is for.
	AccountNumberId *string `json:"account_number_id,omitempty"`

	// The transfer amount in cents. A positive amount originates a credit transfer
	// pushing funds to the receiving account. A negative amount originates a debit
	// transfer pulling funds from the receiving account.
	Amount *int64 `json:"amount,omitempty"`
}

// The identifier of the Account Number the inbound ACH Transfer is for.
func (r *ACHTransfersCreateInboundParams) GetAccountNumberId() (AccountNumberId string) {
	if r != nil && r.AccountNumberId != nil {
		AccountNumberId = *r.AccountNumberId
	}
	return
}

// The transfer amount in cents. A positive amount originates a credit transfer
// pushing funds to the receiving account. A negative amount originates a debit
// transfer pulling funds from the receiving account.
func (r *ACHTransfersCreateInboundParams) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}
