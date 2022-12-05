package oauth_connections

type OauthConnectionsListResponse struct {
	// The contents of the list.
	Data *[]OauthConnection `json:"data,omitempty"`

	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor,omitempty"`
}

// The contents of the list.
func (r *OauthConnectionsListResponse) GetData() (Data []OauthConnection) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *OauthConnectionsListResponse) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}
