package transactions

type TransactionsListQuery struct {
	// Return the page of entries after this one.
	Cursor *string `json:"cursor,omitempty"`

	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `json:"limit,omitempty"`

	// Filter Transactions for those belonging to the specified Account.
	AccountId *string `json:"account_id,omitempty"`

	// Filter Transactions for those belonging to the specified route.
	RouteId   *string                         `json:"route_id,omitempty"`
	CreatedAt *TransactionsCreatedAtListQuery `json:"created_at,omitempty"`
}

// Return the page of entries after this one.
func (r *TransactionsListQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *TransactionsListQuery) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Transactions for those belonging to the specified Account.
func (r *TransactionsListQuery) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// Filter Transactions for those belonging to the specified route.
func (r *TransactionsListQuery) GetRouteId() (RouteId string) {
	if r != nil && r.RouteId != nil {
		RouteId = *r.RouteId
	}
	return
}

func (r *TransactionsListQuery) GetCreatedAt() (CreatedAt TransactionsCreatedAtListQuery) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}
