package entities

type EntitiesCorporationRetrieveResponse struct {
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
	Address *EntitiesCorporationAddressRetrieveResponse `json:"address,omitempty"`

	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwners *[]EntitiesCorporationBeneficialOwnersRetrieveResponse `json:"beneficial_owners,omitempty"`
}

// The legal name of the corporation.
func (r *EntitiesCorporationRetrieveResponse) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The website of the corporation.
func (r *EntitiesCorporationRetrieveResponse) GetWebsite() (Website string) {
	if r != nil && r.Website != nil {
		Website = *r.Website
	}
	return
}

// The Employer Identification Number (EIN) for the corporation.
func (r *EntitiesCorporationRetrieveResponse) GetTaxIdentifier() (TaxIdentifier string) {
	if r != nil && r.TaxIdentifier != nil {
		TaxIdentifier = *r.TaxIdentifier
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the
// corporation's state of incorporation.
func (r *EntitiesCorporationRetrieveResponse) GetIncorporationState() (IncorporationState string) {
	if r != nil && r.IncorporationState != nil {
		IncorporationState = *r.IncorporationState
	}
	return
}

// The corporation's address.
func (r *EntitiesCorporationRetrieveResponse) GetAddress() (Address EntitiesCorporationAddressRetrieveResponse) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// The identifying details of anyone controlling or owning 25% or more of the
// corporation.
func (r *EntitiesCorporationRetrieveResponse) GetBeneficialOwners() (BeneficialOwners []EntitiesCorporationBeneficialOwnersRetrieveResponse) {
	if r != nil && r.BeneficialOwners != nil {
		BeneficialOwners = *r.BeneficialOwners
	}
	return
}
