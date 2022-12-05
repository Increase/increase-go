package entities

type EntitiesTrustTrusteesIndividualCreateParams struct {
	// The person's legal name.
	Name *string `json:"name,omitempty"`

	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `json:"date_of_birth,omitempty"`

	// The individual's address.
	Address *EntitiesTrustTrusteesIndividualAddressCreateParams `json:"address,omitempty"`

	// A means of verifying the person's identity.
	Identification *EntitiesTrustTrusteesIndividualIdentificationCreateParams `json:"identification,omitempty"`
}

// The person's legal name.
func (r *EntitiesTrustTrusteesIndividualCreateParams) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The person's date of birth in YYYY-MM-DD format.
func (r *EntitiesTrustTrusteesIndividualCreateParams) GetDateOfBirth() (DateOfBirth string) {
	if r != nil && r.DateOfBirth != nil {
		DateOfBirth = *r.DateOfBirth
	}
	return
}

// The individual's address.
func (r *EntitiesTrustTrusteesIndividualCreateParams) GetAddress() (Address EntitiesTrustTrusteesIndividualAddressCreateParams) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// A means of verifying the person's identity.
func (r *EntitiesTrustTrusteesIndividualCreateParams) GetIdentification() (Identification EntitiesTrustTrusteesIndividualIdentificationCreateParams) {
	if r != nil && r.Identification != nil {
		Identification = *r.Identification
	}
	return
}
