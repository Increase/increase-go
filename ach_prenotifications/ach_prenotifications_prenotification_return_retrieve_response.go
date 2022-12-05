package ach_prenotifications

type ACHPrenotificationsPrenotificationReturnRetrieveResponse struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Prenotification was returned.
	CreatedAt *string `json:"created_at,omitempty"`

	// Why the Prenotification was returned.
	ReturnReasonCode *string `json:"return_reason_code,omitempty"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Prenotification was returned.
func (r *ACHPrenotificationsPrenotificationReturnRetrieveResponse) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// Why the Prenotification was returned.
func (r *ACHPrenotificationsPrenotificationReturnRetrieveResponse) GetReturnReasonCode() (ReturnReasonCode string) {
	if r != nil && r.ReturnReasonCode != nil {
		ReturnReasonCode = *r.ReturnReasonCode
	}
	return
}
