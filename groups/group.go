package groups

type Group struct {
	// If the Group is activated or not.
	ActivationStatus *string `json:"activation_status,omitempty"`

	// If the Group is allowed to create ACH debits.
	ACHDebitStatus *string `json:"ach_debit_status,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Group
	// was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The Group identifier.
	Id *string `json:"id,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `group`.
	Type *string `json:"type,omitempty"`
}

// If the Group is activated or not.
func (r *Group) GetActivationStatus() (ActivationStatus string) {
	if r != nil && r.ActivationStatus != nil {
		ActivationStatus = *r.ActivationStatus
	}
	return
}

// If the Group is allowed to create ACH debits.
func (r *Group) GetACHDebitStatus() (ACHDebitStatus string) {
	if r != nil && r.ACHDebitStatus != nil {
		ACHDebitStatus = *r.ACHDebitStatus
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Group
// was created.
func (r *Group) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The Group identifier.
func (r *Group) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `group`.
func (r *Group) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
