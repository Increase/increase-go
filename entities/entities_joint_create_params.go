package entities

type EntitiesJointCreateParams struct {
	// The name of the joint entity.
	Name *string `json:"name,omitempty"`

	// The two individuals that share control of the entity.
	Individuals *[]EntitiesJointIndividualsCreateParams `json:"individuals,omitempty"`
}

// The name of the joint entity.
func (r *EntitiesJointCreateParams) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The two individuals that share control of the entity.
func (r *EntitiesJointCreateParams) GetIndividuals() (Individuals []EntitiesJointIndividualsCreateParams) {
	if r != nil && r.Individuals != nil {
		Individuals = *r.Individuals
	}
	return
}
