package pending_transactions

type PendingTransactionsListQuery struct {
	// Return the page of entries after this one.
	Cursor *string `json:"cursor,omitempty"`

	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `json:"limit,omitempty"`

	// Filter pending transactions to those belonging to the specified Account.
	AccountId *string `json:"account_id,omitempty"`

	// Filter pending transactions to those belonging to the specified Route.
	RouteId *string `json:"route_id,omitempty"`

	// Filter pending transactions to those caused by the specified source.
	SourceId *string                             `json:"source_id,omitempty"`
	Status   *PendingTransactionsStatusListQuery `json:"status,omitempty"`
}

// Return the page of entries after this one.
func (r *PendingTransactionsListQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *PendingTransactionsListQuery) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter pending transactions to those belonging to the specified Account.
func (r *PendingTransactionsListQuery) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// Filter pending transactions to those belonging to the specified Route.
func (r *PendingTransactionsListQuery) GetRouteId() (RouteId string) {
	if r != nil && r.RouteId != nil {
		RouteId = *r.RouteId
	}
	return
}

// Filter pending transactions to those caused by the specified source.
func (r *PendingTransactionsListQuery) GetSourceId() (SourceId string) {
	if r != nil && r.SourceId != nil {
		SourceId = *r.SourceId
	}
	return
}

func (r *PendingTransactionsListQuery) GetStatus() (Status PendingTransactionsStatusListQuery) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}
