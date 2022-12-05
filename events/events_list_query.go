package events

type EventsListQuery struct {
	// Return the page of entries after this one.
	Cursor *string `json:"cursor,omitempty"`

	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `json:"limit,omitempty"`

	// Filter Events to those belonging to the object with the provided identifier.
	AssociatedObjectId *string                   `json:"associated_object_id,omitempty"`
	CreatedAt          *EventsCreatedAtListQuery `json:"created_at,omitempty"`
	Category           *EventsCategoryListQuery  `json:"category,omitempty"`
}

// Return the page of entries after this one.
func (r *EventsListQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *EventsListQuery) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Events to those belonging to the object with the provided identifier.
func (r *EventsListQuery) GetAssociatedObjectId() (AssociatedObjectId string) {
	if r != nil && r.AssociatedObjectId != nil {
		AssociatedObjectId = *r.AssociatedObjectId
	}
	return
}

func (r *EventsListQuery) GetCreatedAt() (CreatedAt EventsCreatedAtListQuery) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r *EventsListQuery) GetCategory() (Category EventsCategoryListQuery) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}
