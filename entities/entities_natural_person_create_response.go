package entities

type EntitiesNaturalPersonCreateResponse struct {
	// The person's legal name.
	Name *string `json:"name,omitempty"`

	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `json:"date_of_birth,omitempty"`

	// The person's address.
	Address *EntitiesNaturalPersonAddressCreateResponse `json:"address,omitempty"`

	// A means of verifying the person's identity.
	Identification *EntitiesNaturalPersonIdentificationCreateResponse `json:"identification,omitempty"`
}

// The person's legal name.
func (r *EntitiesNaturalPersonCreateResponse) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The person's date of birth in YYYY-MM-DD format.
func (r *EntitiesNaturalPersonCreateResponse) GetDateOfBirth() (DateOfBirth string) {
	if r != nil && r.DateOfBirth != nil {
		DateOfBirth = *r.DateOfBirth
	}
	return
}

// The person's address.
func (r *EntitiesNaturalPersonCreateResponse) GetAddress() (Address EntitiesNaturalPersonAddressCreateResponse) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// A means of verifying the person's identity.
func (r *EntitiesNaturalPersonCreateResponse) GetIdentification() (Identification EntitiesNaturalPersonIdentificationCreateResponse) {
	if r != nil && r.Identification != nil {
		Identification = *r.Identification
	}
	return
}
