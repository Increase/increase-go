package digital_wallet_tokens

type DigitalWalletTokensListResponse struct {
	// The contents of the list.
	Data *[]DigitalWalletToken `json:"data,omitempty"`

	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor,omitempty"`
}

// The contents of the list.
func (r *DigitalWalletTokensListResponse) GetData() (Data []DigitalWalletToken) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *DigitalWalletTokensListResponse) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}
