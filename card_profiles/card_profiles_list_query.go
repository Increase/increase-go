package card_profiles

type CardProfilesListQuery struct {
	// Return the page of entries after this one.
	Cursor *string `json:"cursor,omitempty"`

	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  *int64                       `json:"limit,omitempty"`
	Status *CardProfilesStatusListQuery `json:"status,omitempty"`
}

// Return the page of entries after this one.
func (r *CardProfilesListQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *CardProfilesListQuery) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

func (r *CardProfilesListQuery) GetStatus() (Status CardProfilesStatusListQuery) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}
