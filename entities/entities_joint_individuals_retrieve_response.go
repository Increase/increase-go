package entities

type EntitiesJointIndividualsRetrieveResponse struct {
	// The person's legal name.
	Name *string `json:"name,omitempty"`

	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `json:"date_of_birth,omitempty"`

	// The person's address.
	Address *EntitiesJointIndividualsAddressRetrieveResponse `json:"address,omitempty"`

	// A means of verifying the person's identity.
	Identification *EntitiesJointIndividualsIdentificationRetrieveResponse `json:"identification,omitempty"`
}

// The person's legal name.
func (r *EntitiesJointIndividualsRetrieveResponse) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The person's date of birth in YYYY-MM-DD format.
func (r *EntitiesJointIndividualsRetrieveResponse) GetDateOfBirth() (DateOfBirth string) {
	if r != nil && r.DateOfBirth != nil {
		DateOfBirth = *r.DateOfBirth
	}
	return
}

// The person's address.
func (r *EntitiesJointIndividualsRetrieveResponse) GetAddress() (Address EntitiesJointIndividualsAddressRetrieveResponse) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// A means of verifying the person's identity.
func (r *EntitiesJointIndividualsRetrieveResponse) GetIdentification() (Identification EntitiesJointIndividualsIdentificationRetrieveResponse) {
	if r != nil && r.Identification != nil {
		Identification = *r.Identification
	}
	return
}
