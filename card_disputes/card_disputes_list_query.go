package card_disputes

type CardDisputesListQuery struct {
	// Return the page of entries after this one.
	Cursor *string `json:"cursor,omitempty"`

	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit     *int64                          `json:"limit,omitempty"`
	CreatedAt *CardDisputesCreatedAtListQuery `json:"created_at,omitempty"`
	Status    *CardDisputesStatusListQuery    `json:"status,omitempty"`
}

// Return the page of entries after this one.
func (r *CardDisputesListQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *CardDisputesListQuery) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

func (r *CardDisputesListQuery) GetCreatedAt() (CreatedAt CardDisputesCreatedAtListQuery) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r *CardDisputesListQuery) GetStatus() (Status CardDisputesStatusListQuery) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}
