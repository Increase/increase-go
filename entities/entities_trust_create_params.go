package entities

type EntitiesTrustCreateParams struct {
	// The legal name of the trust.
	Name *string `json:"name,omitempty"`

	// Whether the trust is `revocable` or `irrevocable`. Irrevocable trusts require
	// their own Employer Identification Number. Revocable trusts require information
	// about the individual `grantor` who created the trust.
	Category *string `json:"category,omitempty"`

	// The Employer Identification Number (EIN) for the trust. Required if `category`
	// is equal to `irrevocable`.
	TaxIdentifier *string `json:"tax_identifier,omitempty"`

	// The two-letter United States Postal Service (USPS) abbreviation for the state in
	// which the trust was formed.
	FormationState *string `json:"formation_state,omitempty"`

	// The trust's address.
	Address *EntitiesTrustAddressCreateParams `json:"address,omitempty"`

	// The identifier of the File containing the formation document of the trust.
	FormationDocumentFileId *string `json:"formation_document_file_id,omitempty"`

	// The trustees of the trust.
	Trustees *[]EntitiesTrustTrusteesCreateParams `json:"trustees,omitempty"`

	// The grantor of the trust. Required if `category` is equal to `revocable`.
	Grantor *EntitiesTrustGrantorCreateParams `json:"grantor,omitempty"`
}

// The legal name of the trust.
func (r *EntitiesTrustCreateParams) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// Whether the trust is `revocable` or `irrevocable`. Irrevocable trusts require
// their own Employer Identification Number. Revocable trusts require information
// about the individual `grantor` who created the trust.
func (r *EntitiesTrustCreateParams) GetCategory() (Category string) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// The Employer Identification Number (EIN) for the trust. Required if `category`
// is equal to `irrevocable`.
func (r *EntitiesTrustCreateParams) GetTaxIdentifier() (TaxIdentifier string) {
	if r != nil && r.TaxIdentifier != nil {
		TaxIdentifier = *r.TaxIdentifier
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the state in
// which the trust was formed.
func (r *EntitiesTrustCreateParams) GetFormationState() (FormationState string) {
	if r != nil && r.FormationState != nil {
		FormationState = *r.FormationState
	}
	return
}

// The trust's address.
func (r *EntitiesTrustCreateParams) GetAddress() (Address EntitiesTrustAddressCreateParams) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// The identifier of the File containing the formation document of the trust.
func (r *EntitiesTrustCreateParams) GetFormationDocumentFileId() (FormationDocumentFileId string) {
	if r != nil && r.FormationDocumentFileId != nil {
		FormationDocumentFileId = *r.FormationDocumentFileId
	}
	return
}

// The trustees of the trust.
func (r *EntitiesTrustCreateParams) GetTrustees() (Trustees []EntitiesTrustTrusteesCreateParams) {
	if r != nil && r.Trustees != nil {
		Trustees = *r.Trustees
	}
	return
}

// The grantor of the trust. Required if `category` is equal to `revocable`.
func (r *EntitiesTrustCreateParams) GetGrantor() (Grantor EntitiesTrustGrantorCreateParams) {
	if r != nil && r.Grantor != nil {
		Grantor = *r.Grantor
	}
	return
}
