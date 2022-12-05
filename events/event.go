package events

type Event struct {
	// The identifier of the object that generated this Event.
	AssociatedObjectId *string `json:"associated_object_id,omitempty"`

	// The type of the object that generated this Event.
	AssociatedObjectType *string `json:"associated_object_type,omitempty"`

	// The category of the Event. We may add additional possible values for this enum
	// over time; your application should be able to handle such additions gracefully.
	Category *string `json:"category,omitempty"`

	// The time the Event was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The Event identifier.
	Id *string `json:"id,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `event`.
	Type *string `json:"type,omitempty"`
}

// The identifier of the object that generated this Event.
func (r *Event) GetAssociatedObjectId() (AssociatedObjectId string) {
	if r != nil && r.AssociatedObjectId != nil {
		AssociatedObjectId = *r.AssociatedObjectId
	}
	return
}

// The type of the object that generated this Event.
func (r *Event) GetAssociatedObjectType() (AssociatedObjectType string) {
	if r != nil && r.AssociatedObjectType != nil {
		AssociatedObjectType = *r.AssociatedObjectType
	}
	return
}

// The category of the Event. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
func (r *Event) GetCategory() (Category string) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// The time the Event was created.
func (r *Event) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The Event identifier.
func (r *Event) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `event`.
func (r *Event) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
