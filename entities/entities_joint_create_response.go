package entities

type EntitiesJointCreateResponse struct {
	// The entity's name.
	Name *string `json:"name,omitempty"`

	// The two individuals that share control of the entity.
	Individuals *[]EntitiesJointIndividualsCreateResponse `json:"individuals,omitempty"`
}

// The entity's name.
func (r *EntitiesJointCreateResponse) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The two individuals that share control of the entity.
func (r *EntitiesJointCreateResponse) GetIndividuals() (Individuals []EntitiesJointIndividualsCreateResponse) {
	if r != nil && r.Individuals != nil {
		Individuals = *r.Individuals
	}
	return
}
