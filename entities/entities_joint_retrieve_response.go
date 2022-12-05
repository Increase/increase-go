package entities

type EntitiesJointRetrieveResponse struct {
	// The entity's name.
	Name *string `json:"name,omitempty"`

	// The two individuals that share control of the entity.
	Individuals *[]EntitiesJointIndividualsRetrieveResponse `json:"individuals,omitempty"`
}

// The entity's name.
func (r *EntitiesJointRetrieveResponse) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The two individuals that share control of the entity.
func (r *EntitiesJointRetrieveResponse) GetIndividuals() (Individuals []EntitiesJointIndividualsRetrieveResponse) {
	if r != nil && r.Individuals != nil {
		Individuals = *r.Individuals
	}
	return
}
