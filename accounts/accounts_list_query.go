package accounts

type AccountsListQuery struct {
	// Return the page of entries after this one.
	Cursor *string `json:"cursor,omitempty"`

	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `json:"limit,omitempty"`

	// Filter Accounts for those belonging to the specified Entity.
	EntityId *string `json:"entity_id,omitempty"`

	// Filter Accounts for those with the specified status.
	Status *string `json:"status,omitempty"`
}

// Return the page of entries after this one.
func (r *AccountsListQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *AccountsListQuery) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Accounts for those belonging to the specified Entity.
func (r *AccountsListQuery) GetEntityId() (EntityId string) {
	if r != nil && r.EntityId != nil {
		EntityId = *r.EntityId
	}
	return
}

// Filter Accounts for those with the specified status.
func (r *AccountsListQuery) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}
