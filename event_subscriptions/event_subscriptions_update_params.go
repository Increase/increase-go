package event_subscriptions

type EventSubscriptionsUpdateParams struct {
	// The status to update the Event Subscription with.
	Status *string `json:"status,omitempty"`
}

// The status to update the Event Subscription with.
func (r *EventSubscriptionsUpdateParams) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}
