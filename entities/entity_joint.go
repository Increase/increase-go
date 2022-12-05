package entities

type EntityJoint struct {
	// The entity's name.
	Name *string `json:"name,omitempty"`

	// The two individuals that share control of the entity.
	Individuals *[]EntityJointIndividual `json:"individuals,omitempty"`
}

// The entity's name.
func (r *EntityJoint) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The two individuals that share control of the entity.
func (r *EntityJoint) GetIndividuals() (Individuals []EntityJointIndividual) {
	if r != nil && r.Individuals != nil {
		Individuals = *r.Individuals
	}
	return
}
