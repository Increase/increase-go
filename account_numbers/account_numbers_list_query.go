package account_numbers

type AccountNumbersListQuery struct {
	// Return the page of entries after this one.
	Cursor *string `json:"cursor,omitempty"`

	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `json:"limit,omitempty"`

	// The status to retrieve Account Numbers for.
	Status *string `json:"status,omitempty"`

	// Filter Account Numbers to those belonging to the specified Account.
	AccountId *string `json:"account_id,omitempty"`
}

// Return the page of entries after this one.
func (r *AccountNumbersListQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *AccountNumbersListQuery) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// The status to retrieve Account Numbers for.
func (r *AccountNumbersListQuery) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// Filter Account Numbers to those belonging to the specified Account.
func (r *AccountNumbersListQuery) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}
