package transactions

type TransactionsSourceSampleFundsRetrieveResponse struct {
	// Where the sample funds came from.
	Originator *string `json:"originator,omitempty"`
}

// Where the sample funds came from.
func (r *TransactionsSourceSampleFundsRetrieveResponse) GetOriginator() (Originator string) {
	if r != nil && r.Originator != nil {
		Originator = *r.Originator
	}
	return
}
