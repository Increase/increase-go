package entities

type EntityTrustTrusteeIndividual struct {
	// The person's legal name.
	Name *string `json:"name,omitempty"`

	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `json:"date_of_birth,omitempty"`

	// The person's address.
	Address *EntityTrustTrusteeIndividualAddress `json:"address,omitempty"`

	// A means of verifying the person's identity.
	Identification *EntityTrustTrusteeIndividualIdentification `json:"identification,omitempty"`
}

// The person's legal name.
func (r *EntityTrustTrusteeIndividual) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The person's date of birth in YYYY-MM-DD format.
func (r *EntityTrustTrusteeIndividual) GetDateOfBirth() (DateOfBirth string) {
	if r != nil && r.DateOfBirth != nil {
		DateOfBirth = *r.DateOfBirth
	}
	return
}

// The person's address.
func (r *EntityTrustTrusteeIndividual) GetAddress() (Address EntityTrustTrusteeIndividualAddress) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// A means of verifying the person's identity.
func (r *EntityTrustTrusteeIndividual) GetIdentification() (Identification EntityTrustTrusteeIndividualIdentification) {
	if r != nil && r.Identification != nil {
		Identification = *r.Identification
	}
	return
}
