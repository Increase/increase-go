package ach_transfers

type ACHTransfersCancellationCreateResponse struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt *string `json:"canceled_at,omitempty"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Transfer was canceled.
func (r *ACHTransfersCancellationCreateResponse) GetCanceledAt() (CanceledAt string) {
	if r != nil && r.CanceledAt != nil {
		CanceledAt = *r.CanceledAt
	}
	return
}
