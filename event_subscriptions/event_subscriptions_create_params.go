package event_subscriptions

type EventSubscriptionsCreateParams struct {
	// The URL you'd like us to send webhooks to.
	URL *string `json:"url,omitempty"`

	// The key that will be used to sign webhooks. If no value is passed, a random
	// string will be used as default.
	SharedSecret *string `json:"shared_secret,omitempty"`

	// If specified, this subscription will only receive webhooks for Events with the
	// specified `category`.
	SelectedEventCategory *string `json:"selected_event_category,omitempty"`
}

// The URL you'd like us to send webhooks to.
func (r *EventSubscriptionsCreateParams) GetURL() (URL string) {
	if r != nil && r.URL != nil {
		URL = *r.URL
	}
	return
}

// The key that will be used to sign webhooks. If no value is passed, a random
// string will be used as default.
func (r *EventSubscriptionsCreateParams) GetSharedSecret() (SharedSecret string) {
	if r != nil && r.SharedSecret != nil {
		SharedSecret = *r.SharedSecret
	}
	return
}

// If specified, this subscription will only receive webhooks for Events with the
// specified `category`.
func (r *EventSubscriptionsCreateParams) GetSelectedEventCategory() (SelectedEventCategory string) {
	if r != nil && r.SelectedEventCategory != nil {
		SelectedEventCategory = *r.SelectedEventCategory
	}
	return
}
