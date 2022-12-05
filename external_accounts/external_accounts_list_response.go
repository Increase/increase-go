package external_accounts

type ExternalAccountsListResponse struct {
	// The contents of the list.
	Data *[]ExternalAccount `json:"data,omitempty"`

	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor,omitempty"`
}

// The contents of the list.
func (r *ExternalAccountsListResponse) GetData() (Data []ExternalAccount) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *ExternalAccountsListResponse) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}
