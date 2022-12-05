package simulations

type ACHTransfersDeclinedTransactionSourceCheckDeclineCreateInboundResponse struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	AuxiliaryOnUs *string `json:"auxiliary_on_us,omitempty"`

	// Why the check was declined.
	Reason *string `json:"reason,omitempty"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *ACHTransfersDeclinedTransactionSourceCheckDeclineCreateInboundResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransfersDeclinedTransactionSourceCheckDeclineCreateInboundResponse) GetAuxiliaryOnUs() (AuxiliaryOnUs string) {
	if r != nil && r.AuxiliaryOnUs != nil {
		AuxiliaryOnUs = *r.AuxiliaryOnUs
	}
	return
}

// Why the check was declined.
func (r *ACHTransfersDeclinedTransactionSourceCheckDeclineCreateInboundResponse) GetReason() (Reason string) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}
