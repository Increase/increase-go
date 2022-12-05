package routing_numbers

type RoutingNumbersListQuery struct {
	// Return the page of entries after this one.
	Cursor *string `json:"cursor,omitempty"`

	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `json:"limit,omitempty"`

	// Filter financial institutions by routing number.
	RoutingNumber *string `json:"routing_number,omitempty"`
}

// Return the page of entries after this one.
func (r *RoutingNumbersListQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *RoutingNumbersListQuery) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter financial institutions by routing number.
func (r *RoutingNumbersListQuery) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}
