package account_transfers

type AccountTransfersDataCancellationListResponse struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt *string `json:"canceled_at,omitempty"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Transfer was canceled.
func (r *AccountTransfersDataCancellationListResponse) GetCanceledAt() (CanceledAt string) {
	if r != nil && r.CanceledAt != nil {
		CanceledAt = *r.CanceledAt
	}
	return
}
