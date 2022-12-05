package declined_transactions

type DeclinedTransactionsListResponse struct {
	// The contents of the list.
	Data *[]DeclinedTransaction `json:"data,omitempty"`

	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor,omitempty"`
}

// The contents of the list.
func (r *DeclinedTransactionsListResponse) GetData() (Data []DeclinedTransaction) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *DeclinedTransactionsListResponse) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}
