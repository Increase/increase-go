package wire_drawdown_requests

type WireDrawdownRequestsListResponse struct {
	// The contents of the list.
	Data *[]WireDrawdownRequest `json:"data,omitempty"`

	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor,omitempty"`
}

// The contents of the list.
func (r *WireDrawdownRequestsListResponse) GetData() (Data []WireDrawdownRequest) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *WireDrawdownRequestsListResponse) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}
