package entities

type EntitiesTrustTrusteesIndividualCreateResponse struct {
	// The person's legal name.
	Name *string `json:"name,omitempty"`

	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `json:"date_of_birth,omitempty"`

	// The person's address.
	Address *EntitiesTrustTrusteesIndividualAddressCreateResponse `json:"address,omitempty"`

	// A means of verifying the person's identity.
	Identification *EntitiesTrustTrusteesIndividualIdentificationCreateResponse `json:"identification,omitempty"`
}

// The person's legal name.
func (r *EntitiesTrustTrusteesIndividualCreateResponse) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The person's date of birth in YYYY-MM-DD format.
func (r *EntitiesTrustTrusteesIndividualCreateResponse) GetDateOfBirth() (DateOfBirth string) {
	if r != nil && r.DateOfBirth != nil {
		DateOfBirth = *r.DateOfBirth
	}
	return
}

// The person's address.
func (r *EntitiesTrustTrusteesIndividualCreateResponse) GetAddress() (Address EntitiesTrustTrusteesIndividualAddressCreateResponse) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// A means of verifying the person's identity.
func (r *EntitiesTrustTrusteesIndividualCreateResponse) GetIdentification() (Identification EntitiesTrustTrusteesIndividualIdentificationCreateResponse) {
	if r != nil && r.Identification != nil {
		Identification = *r.Identification
	}
	return
}
