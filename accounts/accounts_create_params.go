package accounts

type AccountsCreateParams struct {
	// The identifier for the Entity that will own the Account.
	EntityId *string `json:"entity_id,omitempty"`

	// The identifier of an Entity that, while not owning the Account, is associated
	// with its activity. Its relationship to your group must be `informational`.
	InformationalEntityId *string `json:"informational_entity_id,omitempty"`

	// The name you choose for the Account.
	Name *string `json:"name,omitempty"`
}

// The identifier for the Entity that will own the Account.
func (r *AccountsCreateParams) GetEntityId() (EntityId string) {
	if r != nil && r.EntityId != nil {
		EntityId = *r.EntityId
	}
	return
}

// The identifier of an Entity that, while not owning the Account, is associated
// with its activity. Its relationship to your group must be `informational`.
func (r *AccountsCreateParams) GetInformationalEntityId() (InformationalEntityId string) {
	if r != nil && r.InformationalEntityId != nil {
		InformationalEntityId = *r.InformationalEntityId
	}
	return
}

// The name you choose for the Account.
func (r *AccountsCreateParams) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}
