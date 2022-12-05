package declined_transactions

type DeclinedTransactionsListQuery struct {
	// Return the page of entries after this one.
	Cursor *string `json:"cursor,omitempty"`

	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `json:"limit,omitempty"`

	// Filter Declined Transactions to ones belonging to the specified Account.
	AccountId *string `json:"account_id,omitempty"`

	// Filter Declined Transactions to those belonging to the specified route.
	RouteId   *string                                 `json:"route_id,omitempty"`
	CreatedAt *DeclinedTransactionsCreatedAtListQuery `json:"created_at,omitempty"`
}

// Return the page of entries after this one.
func (r *DeclinedTransactionsListQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *DeclinedTransactionsListQuery) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Declined Transactions to ones belonging to the specified Account.
func (r *DeclinedTransactionsListQuery) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// Filter Declined Transactions to those belonging to the specified route.
func (r *DeclinedTransactionsListQuery) GetRouteId() (RouteId string) {
	if r != nil && r.RouteId != nil {
		RouteId = *r.RouteId
	}
	return
}

func (r *DeclinedTransactionsListQuery) GetCreatedAt() (CreatedAt DeclinedTransactionsCreatedAtListQuery) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}
