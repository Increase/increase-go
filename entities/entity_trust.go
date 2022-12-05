package entities

type EntityTrust struct {
	// The trust's name
	Name *string `json:"name,omitempty"`

	// Whether the trust is `revocable` or `irrevocable`.
	Category *string `json:"category,omitempty"`

	// The trust's address.
	Address *EntityTrustAddress `json:"address,omitempty"`

	// The two-letter United States Postal Service (USPS) abbreviation for the state in
	// which the trust was formed.
	FormationState *string `json:"formation_state,omitempty"`

	// The Employer Identification Number (EIN) of the trust itself.
	TaxIdentifier *string `json:"tax_identifier,omitempty"`

	// The trustees of the trust.
	Trustees *[]EntityTrustTrustee `json:"trustees,omitempty"`

	// The grantor of the trust. Will be present if the `category` is `revocable`.
	Grantor *EntityTrustGrantor `json:"grantor,omitempty"`

	// The ID for the File containing the formation document of the trust.
	FormationDocumentFileId *string `json:"formation_document_file_id,omitempty"`
}

// The trust's name
func (r *EntityTrust) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// Whether the trust is `revocable` or `irrevocable`.
func (r *EntityTrust) GetCategory() (Category string) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// The trust's address.
func (r *EntityTrust) GetAddress() (Address EntityTrustAddress) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the state in
// which the trust was formed.
func (r *EntityTrust) GetFormationState() (FormationState string) {
	if r != nil && r.FormationState != nil {
		FormationState = *r.FormationState
	}
	return
}

// The Employer Identification Number (EIN) of the trust itself.
func (r *EntityTrust) GetTaxIdentifier() (TaxIdentifier string) {
	if r != nil && r.TaxIdentifier != nil {
		TaxIdentifier = *r.TaxIdentifier
	}
	return
}

// The trustees of the trust.
func (r *EntityTrust) GetTrustees() (Trustees []EntityTrustTrustee) {
	if r != nil && r.Trustees != nil {
		Trustees = *r.Trustees
	}
	return
}

// The grantor of the trust. Will be present if the `category` is `revocable`.
func (r *EntityTrust) GetGrantor() (Grantor EntityTrustGrantor) {
	if r != nil && r.Grantor != nil {
		Grantor = *r.Grantor
	}
	return
}

// The ID for the File containing the formation document of the trust.
func (r *EntityTrust) GetFormationDocumentFileId() (FormationDocumentFileId string) {
	if r != nil && r.FormationDocumentFileId != nil {
		FormationDocumentFileId = *r.FormationDocumentFileId
	}
	return
}
