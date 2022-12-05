package event_subscriptions

type EventSubscriptionsListResponse struct {
	// The contents of the list.
	Data *[]EventSubscription `json:"data,omitempty"`

	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor,omitempty"`
}

// The contents of the list.
func (r *EventSubscriptionsListResponse) GetData() (Data []EventSubscription) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *EventSubscriptionsListResponse) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}
