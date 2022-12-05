package ach_prenotifications

type ACHPrenotificationsListQuery struct {
	// Return the page of entries after this one.
	Cursor *string `json:"cursor,omitempty"`

	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit     *int64                                 `json:"limit,omitempty"`
	CreatedAt *ACHPrenotificationsCreatedAtListQuery `json:"created_at,omitempty"`
}

// Return the page of entries after this one.
func (r *ACHPrenotificationsListQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ACHPrenotificationsListQuery) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

func (r *ACHPrenotificationsListQuery) GetCreatedAt() (CreatedAt ACHPrenotificationsCreatedAtListQuery) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}
