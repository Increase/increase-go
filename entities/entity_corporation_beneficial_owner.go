package entities

type EntityCorporationBeneficialOwner struct {
	// Personal details for the beneficial owner.
	Individual *EntityCorporationBeneficialOwnerIndividual `json:"individual,omitempty"`

	// This person's role or title within the entity.
	CompanyTitle *string `json:"company_title,omitempty"`

	// Why this person is considered a beneficial owner of the entity.
	Prong *string `json:"prong,omitempty"`
}

// Personal details for the beneficial owner.
func (r *EntityCorporationBeneficialOwner) GetIndividual() (Individual EntityCorporationBeneficialOwnerIndividual) {
	if r != nil && r.Individual != nil {
		Individual = *r.Individual
	}
	return
}

// This person's role or title within the entity.
func (r *EntityCorporationBeneficialOwner) GetCompanyTitle() (CompanyTitle string) {
	if r != nil && r.CompanyTitle != nil {
		CompanyTitle = *r.CompanyTitle
	}
	return
}

// Why this person is considered a beneficial owner of the entity.
func (r *EntityCorporationBeneficialOwner) GetProng() (Prong string) {
	if r != nil && r.Prong != nil {
		Prong = *r.Prong
	}
	return
}
