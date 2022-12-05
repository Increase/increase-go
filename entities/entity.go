package entities

type Entity struct {
	// The entity's identifier.
	Id *string `json:"id,omitempty"`

	// The entity's legal structure.
	Structure *string `json:"structure,omitempty"`

	// Details of the corporation entity. Will be present if `structure` is equal to
	// `corporation`.
	Corporation *EntityCorporation `json:"corporation,omitempty"`

	// Details of the natural person entity. Will be present if `structure` is equal to
	// `natural_person`.
	NaturalPerson *EntityNaturalPerson `json:"natural_person,omitempty"`

	// Details of the joint entity. Will be present if `structure` is equal to `joint`.
	Joint *EntityJoint `json:"joint,omitempty"`

	// Details of the trust entity. Will be present if `structure` is equal to `trust`.
	Trust *EntityTrust `json:"trust,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `entity`.
	Type *string `json:"type,omitempty"`

	// The entity's description for display purposes.
	Description *string `json:"description,omitempty"`

	// The relationship between your group and the entity.
	Relationship *string `json:"relationship,omitempty"`
}

// The entity's identifier.
func (r *Entity) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The entity's legal structure.
func (r *Entity) GetStructure() (Structure string) {
	if r != nil && r.Structure != nil {
		Structure = *r.Structure
	}
	return
}

// Details of the corporation entity. Will be present if `structure` is equal to
// `corporation`.
func (r *Entity) GetCorporation() (Corporation EntityCorporation) {
	if r != nil && r.Corporation != nil {
		Corporation = *r.Corporation
	}
	return
}

// Details of the natural person entity. Will be present if `structure` is equal to
// `natural_person`.
func (r *Entity) GetNaturalPerson() (NaturalPerson EntityNaturalPerson) {
	if r != nil && r.NaturalPerson != nil {
		NaturalPerson = *r.NaturalPerson
	}
	return
}

// Details of the joint entity. Will be present if `structure` is equal to `joint`.
func (r *Entity) GetJoint() (Joint EntityJoint) {
	if r != nil && r.Joint != nil {
		Joint = *r.Joint
	}
	return
}

// Details of the trust entity. Will be present if `structure` is equal to `trust`.
func (r *Entity) GetTrust() (Trust EntityTrust) {
	if r != nil && r.Trust != nil {
		Trust = *r.Trust
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `entity`.
func (r *Entity) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

// The entity's description for display purposes.
func (r *Entity) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The relationship between your group and the entity.
func (r *Entity) GetRelationship() (Relationship string) {
	if r != nil && r.Relationship != nil {
		Relationship = *r.Relationship
	}
	return
}
