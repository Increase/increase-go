package account_transfers

type AccountTransfersListQuery struct {
	// Return the page of entries after this one.
	Cursor *string `json:"cursor,omitempty"`

	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `json:"limit,omitempty"`

	// Filter Account Transfers to those that originated from the specified Account.
	AccountId *string                             `json:"account_id,omitempty"`
	CreatedAt *AccountTransfersCreatedAtListQuery `json:"created_at,omitempty"`
}

// Return the page of entries after this one.
func (r *AccountTransfersListQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *AccountTransfersListQuery) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Account Transfers to those that originated from the specified Account.
func (r *AccountTransfersListQuery) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

func (r *AccountTransfersListQuery) GetCreatedAt() (CreatedAt AccountTransfersCreatedAtListQuery) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}
