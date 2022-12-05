package entities

type EntitiesTrustGrantorCreateResponse struct {
	// The person's legal name.
	Name *string `json:"name,omitempty"`

	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `json:"date_of_birth,omitempty"`

	// The person's address.
	Address *EntitiesTrustGrantorAddressCreateResponse `json:"address,omitempty"`

	// A means of verifying the person's identity.
	Identification *EntitiesTrustGrantorIdentificationCreateResponse `json:"identification,omitempty"`
}

// The person's legal name.
func (r *EntitiesTrustGrantorCreateResponse) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The person's date of birth in YYYY-MM-DD format.
func (r *EntitiesTrustGrantorCreateResponse) GetDateOfBirth() (DateOfBirth string) {
	if r != nil && r.DateOfBirth != nil {
		DateOfBirth = *r.DateOfBirth
	}
	return
}

// The person's address.
func (r *EntitiesTrustGrantorCreateResponse) GetAddress() (Address EntitiesTrustGrantorAddressCreateResponse) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// A means of verifying the person's identity.
func (r *EntitiesTrustGrantorCreateResponse) GetIdentification() (Identification EntitiesTrustGrantorIdentificationCreateResponse) {
	if r != nil && r.Identification != nil {
		Identification = *r.Identification
	}
	return
}
