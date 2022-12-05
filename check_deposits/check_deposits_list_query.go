package check_deposits

type CheckDepositsListQuery struct {
	// Return the page of entries after this one.
	Cursor *string `json:"cursor,omitempty"`

	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `json:"limit,omitempty"`

	// Filter Check Deposits to those belonging to the specified Account.
	AccountId *string                          `json:"account_id,omitempty"`
	CreatedAt *CheckDepositsCreatedAtListQuery `json:"created_at,omitempty"`
}

// Return the page of entries after this one.
func (r *CheckDepositsListQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *CheckDepositsListQuery) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Check Deposits to those belonging to the specified Account.
func (r *CheckDepositsListQuery) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

func (r *CheckDepositsListQuery) GetCreatedAt() (CreatedAt CheckDepositsCreatedAtListQuery) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}
