package entities

type EntitiesCreateParams struct {
	// The type of Entity to create.
	Structure *string `json:"structure,omitempty"`

	// Details of the corporation entity to create. Required if `structure` is equal to
	// `corporation`.
	Corporation *EntitiesCorporationCreateParams `json:"corporation,omitempty"`

	// Details of the natural person entity to create. Required if `structure` is equal
	// to `natural_person`.
	NaturalPerson *EntitiesNaturalPersonCreateParams `json:"natural_person,omitempty"`

	// Details of the joint entity to create. Required if `structure` is equal to
	// `joint`.
	Joint *EntitiesJointCreateParams `json:"joint,omitempty"`

	// Details of the trust entity to create. Required if `structure` is equal to
	// `trust`.
	Trust *EntitiesTrustCreateParams `json:"trust,omitempty"`

	// The description you choose to give the entity.
	Description *string `json:"description,omitempty"`

	// The relationship between your group and the entity.
	Relationship *string `json:"relationship,omitempty"`
}

// The type of Entity to create.
func (r *EntitiesCreateParams) GetStructure() (Structure string) {
	if r != nil && r.Structure != nil {
		Structure = *r.Structure
	}
	return
}

// Details of the corporation entity to create. Required if `structure` is equal to
// `corporation`.
func (r *EntitiesCreateParams) GetCorporation() (Corporation EntitiesCorporationCreateParams) {
	if r != nil && r.Corporation != nil {
		Corporation = *r.Corporation
	}
	return
}

// Details of the natural person entity to create. Required if `structure` is equal
// to `natural_person`.
func (r *EntitiesCreateParams) GetNaturalPerson() (NaturalPerson EntitiesNaturalPersonCreateParams) {
	if r != nil && r.NaturalPerson != nil {
		NaturalPerson = *r.NaturalPerson
	}
	return
}

// Details of the joint entity to create. Required if `structure` is equal to
// `joint`.
func (r *EntitiesCreateParams) GetJoint() (Joint EntitiesJointCreateParams) {
	if r != nil && r.Joint != nil {
		Joint = *r.Joint
	}
	return
}

// Details of the trust entity to create. Required if `structure` is equal to
// `trust`.
func (r *EntitiesCreateParams) GetTrust() (Trust EntitiesTrustCreateParams) {
	if r != nil && r.Trust != nil {
		Trust = *r.Trust
	}
	return
}

// The description you choose to give the entity.
func (r *EntitiesCreateParams) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The relationship between your group and the entity.
func (r *EntitiesCreateParams) GetRelationship() (Relationship string) {
	if r != nil && r.Relationship != nil {
		Relationship = *r.Relationship
	}
	return
}
