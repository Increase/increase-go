package entities

type EntityTrustGrantor struct {
	// The person's legal name.
	Name *string `json:"name,omitempty"`

	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `json:"date_of_birth,omitempty"`

	// The person's address.
	Address *EntityTrustGrantorAddress `json:"address,omitempty"`

	// A means of verifying the person's identity.
	Identification *EntityTrustGrantorIdentification `json:"identification,omitempty"`
}

// The person's legal name.
func (r *EntityTrustGrantor) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The person's date of birth in YYYY-MM-DD format.
func (r *EntityTrustGrantor) GetDateOfBirth() (DateOfBirth string) {
	if r != nil && r.DateOfBirth != nil {
		DateOfBirth = *r.DateOfBirth
	}
	return
}

// The person's address.
func (r *EntityTrustGrantor) GetAddress() (Address EntityTrustGrantorAddress) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// A means of verifying the person's identity.
func (r *EntityTrustGrantor) GetIdentification() (Identification EntityTrustGrantorIdentification) {
	if r != nil && r.Identification != nil {
		Identification = *r.Identification
	}
	return
}
