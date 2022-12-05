package simulations

type ACHTransfersTransactionSourceSampleFundsCreateInboundResponse struct {
	// Where the sample funds came from.
	Originator *string `json:"originator,omitempty"`
}

// Where the sample funds came from.
func (r *ACHTransfersTransactionSourceSampleFundsCreateInboundResponse) GetOriginator() (Originator string) {
	if r != nil && r.Originator != nil {
		Originator = *r.Originator
	}
	return
}
