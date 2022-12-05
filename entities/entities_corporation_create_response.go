package entities

type EntitiesCorporationCreateResponse struct {
	// The legal name of the corporation.
	Name *string `json:"name,omitempty"`

	// The website of the corporation.
	Website *string `json:"website,omitempty"`

	// The Employer Identification Number (EIN) for the corporation.
	TaxIdentifier *string `json:"tax_identifier,omitempty"`

	// The two-letter United States Postal Service (USPS) abbreviation for the
	// corporation's state of incorporation.
	IncorporationState *string `json:"incorporation_state,omitempty"`

	// The corporation's address.
	Address *EntitiesCorporationAddressCreateResponse `json:"address,omitempty"`

	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwners *[]EntitiesCorporationBeneficialOwnersCreateResponse `json:"beneficial_owners,omitempty"`
}

// The legal name of the corporation.
func (r *EntitiesCorporationCreateResponse) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The website of the corporation.
func (r *EntitiesCorporationCreateResponse) GetWebsite() (Website string) {
	if r != nil && r.Website != nil {
		Website = *r.Website
	}
	return
}

// The Employer Identification Number (EIN) for the corporation.
func (r *EntitiesCorporationCreateResponse) GetTaxIdentifier() (TaxIdentifier string) {
	if r != nil && r.TaxIdentifier != nil {
		TaxIdentifier = *r.TaxIdentifier
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the
// corporation's state of incorporation.
func (r *EntitiesCorporationCreateResponse) GetIncorporationState() (IncorporationState string) {
	if r != nil && r.IncorporationState != nil {
		IncorporationState = *r.IncorporationState
	}
	return
}

// The corporation's address.
func (r *EntitiesCorporationCreateResponse) GetAddress() (Address EntitiesCorporationAddressCreateResponse) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// The identifying details of anyone controlling or owning 25% or more of the
// corporation.
func (r *EntitiesCorporationCreateResponse) GetBeneficialOwners() (BeneficialOwners []EntitiesCorporationBeneficialOwnersCreateResponse) {
	if r != nil && r.BeneficialOwners != nil {
		BeneficialOwners = *r.BeneficialOwners
	}
	return
}
