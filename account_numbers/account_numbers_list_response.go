package account_numbers

type AccountNumbersListResponse struct {
	// The contents of the list.
	Data *[]AccountNumber `json:"data,omitempty"`

	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor,omitempty"`
}

// The contents of the list.
func (r *AccountNumbersListResponse) GetData() (Data []AccountNumber) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *AccountNumbersListResponse) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}
