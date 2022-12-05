package transactions

type TransactionsListResponse struct {
	// The contents of the list.
	Data *[]Transaction `json:"data,omitempty"`

	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor,omitempty"`
}

// The contents of the list.
func (r *TransactionsListResponse) GetData() (Data []Transaction) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *TransactionsListResponse) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}
