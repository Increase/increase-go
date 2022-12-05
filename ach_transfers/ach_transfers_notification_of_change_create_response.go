package ach_transfers

type ACHTransfersNotificationOfChangeCreateResponse struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the notification occurred.
	CreatedAt *string `json:"created_at,omitempty"`

	// The type of change that occurred.
	ChangeCode *string `json:"change_code,omitempty"`

	// The corrected data.
	CorrectedData *string `json:"corrected_data,omitempty"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the notification occurred.
func (r *ACHTransfersNotificationOfChangeCreateResponse) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The type of change that occurred.
func (r *ACHTransfersNotificationOfChangeCreateResponse) GetChangeCode() (ChangeCode string) {
	if r != nil && r.ChangeCode != nil {
		ChangeCode = *r.ChangeCode
	}
	return
}

// The corrected data.
func (r *ACHTransfersNotificationOfChangeCreateResponse) GetCorrectedData() (CorrectedData string) {
	if r != nil && r.CorrectedData != nil {
		CorrectedData = *r.CorrectedData
	}
	return
}
