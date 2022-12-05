package ach_transfers

type ACHTransfersDataApprovalListResponse struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt *string `json:"approved_at,omitempty"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was approved.
func (r *ACHTransfersDataApprovalListResponse) GetApprovedAt() (ApprovedAt string) {
	if r != nil && r.ApprovedAt != nil {
		ApprovedAt = *r.ApprovedAt
	}
	return
}
