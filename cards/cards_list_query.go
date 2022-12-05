package cards

type CardsListQuery struct {
	// Return the page of entries after this one.
	Cursor *string `json:"cursor,omitempty"`

	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `json:"limit,omitempty"`

	// Filter Cards to ones belonging to the specified Account.
	AccountId *string                  `json:"account_id,omitempty"`
	CreatedAt *CardsCreatedAtListQuery `json:"created_at,omitempty"`
}

// Return the page of entries after this one.
func (r *CardsListQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *CardsListQuery) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Cards to ones belonging to the specified Account.
func (r *CardsListQuery) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

func (r *CardsListQuery) GetCreatedAt() (CreatedAt CardsCreatedAtListQuery) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}
