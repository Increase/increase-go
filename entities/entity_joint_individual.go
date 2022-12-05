package entities

type EntityJointIndividual struct {
	// The person's legal name.
	Name *string `json:"name,omitempty"`

	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `json:"date_of_birth,omitempty"`

	// The person's address.
	Address *EntityJointIndividualAddress `json:"address,omitempty"`

	// A means of verifying the person's identity.
	Identification *EntityJointIndividualIdentification `json:"identification,omitempty"`
}

// The person's legal name.
func (r *EntityJointIndividual) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The person's date of birth in YYYY-MM-DD format.
func (r *EntityJointIndividual) GetDateOfBirth() (DateOfBirth string) {
	if r != nil && r.DateOfBirth != nil {
		DateOfBirth = *r.DateOfBirth
	}
	return
}

// The person's address.
func (r *EntityJointIndividual) GetAddress() (Address EntityJointIndividualAddress) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// A means of verifying the person's identity.
func (r *EntityJointIndividual) GetIdentification() (Identification EntityJointIndividualIdentification) {
	if r != nil && r.Identification != nil {
		Identification = *r.Identification
	}
	return
}
