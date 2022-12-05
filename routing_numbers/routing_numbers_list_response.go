package routing_numbers

type RoutingNumbersListResponse struct {
	// The contents of the list.
	Data *[]RoutingNumber `json:"data,omitempty"`

	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor,omitempty"`
}

// The contents of the list.
func (r *RoutingNumbersListResponse) GetData() (Data []RoutingNumber) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *RoutingNumbersListResponse) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}
