package event_subscriptions

type EventSubscription struct {
	// The event subscription identifier.
	Id *string `json:"id,omitempty"`

	// The time the event subscription was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// This indicates if we'll send notifications to this subscription.
	Status *string `json:"status,omitempty"`

	// If specified, this subscription will only receive webhooks for Events with the
	// specified `category`.
	SelectedEventCategory *string `json:"selected_event_category,omitempty"`

	// The webhook url where we'll send notifications.
	URL *string `json:"url,omitempty"`

	// The key that will be used to sign webhooks.
	SharedSecret *string `json:"shared_secret,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `event_subscription`.
	Type *string `json:"type,omitempty"`
}

// The event subscription identifier.
func (r *EventSubscription) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The time the event subscription was created.
func (r *EventSubscription) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// This indicates if we'll send notifications to this subscription.
func (r *EventSubscription) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// If specified, this subscription will only receive webhooks for Events with the
// specified `category`.
func (r *EventSubscription) GetSelectedEventCategory() (SelectedEventCategory string) {
	if r != nil && r.SelectedEventCategory != nil {
		SelectedEventCategory = *r.SelectedEventCategory
	}
	return
}

// The webhook url where we'll send notifications.
func (r *EventSubscription) GetURL() (URL string) {
	if r != nil && r.URL != nil {
		URL = *r.URL
	}
	return
}

// The key that will be used to sign webhooks.
func (r *EventSubscription) GetSharedSecret() (SharedSecret string) {
	if r != nil && r.SharedSecret != nil {
		SharedSecret = *r.SharedSecret
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `event_subscription`.
func (r *EventSubscription) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
