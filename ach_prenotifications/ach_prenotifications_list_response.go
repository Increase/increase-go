package ach_prenotifications

type ACHPrenotificationsListResponse struct {
	// The contents of the list.
	Data *[]ACHPrenotification `json:"data,omitempty"`

	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor,omitempty"`
}

// The contents of the list.
func (r *ACHPrenotificationsListResponse) GetData() (Data []ACHPrenotification) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *ACHPrenotificationsListResponse) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}
