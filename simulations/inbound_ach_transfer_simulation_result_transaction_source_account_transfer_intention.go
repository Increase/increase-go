package simulations

type InboundACHTransferSimulationResultTransactionSourceAccountTransferIntention struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *string `json:"currency,omitempty"`

	// The description you chose to give the transfer.
	Description *string `json:"description,omitempty"`

	// The identifier of the Account to where the Account Transfer was sent.
	DestinationAccountId *string `json:"destination_account_id,omitempty"`

	// The identifier of the Account from where the Account Transfer was sent.
	SourceAccountId *string `json:"source_account_id,omitempty"`

	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferId *string `json:"transfer_id,omitempty"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *InboundACHTransferSimulationResultTransactionSourceAccountTransferIntention) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *InboundACHTransferSimulationResultTransactionSourceAccountTransferIntention) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The description you chose to give the transfer.
func (r *InboundACHTransferSimulationResultTransactionSourceAccountTransferIntention) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The identifier of the Account to where the Account Transfer was sent.
func (r *InboundACHTransferSimulationResultTransactionSourceAccountTransferIntention) GetDestinationAccountId() (DestinationAccountId string) {
	if r != nil && r.DestinationAccountId != nil {
		DestinationAccountId = *r.DestinationAccountId
	}
	return
}

// The identifier of the Account from where the Account Transfer was sent.
func (r *InboundACHTransferSimulationResultTransactionSourceAccountTransferIntention) GetSourceAccountId() (SourceAccountId string) {
	if r != nil && r.SourceAccountId != nil {
		SourceAccountId = *r.SourceAccountId
	}
	return
}

// The identifier of the Account Transfer that led to this Pending Transaction.
func (r *InboundACHTransferSimulationResultTransactionSourceAccountTransferIntention) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}
