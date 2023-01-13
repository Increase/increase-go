package shared

//
type Entity struct {
	// The entity's identifier.
	Id *string `json:"id"`
	// The entity's legal structure.
	Structure *EntityStructure `json:"structure"`
	// Details of the corporation entity. Will be present if `structure` is equal to
	// `corporation`.
	Corporation *EntityCorporation `json:"corporation"`
	// Details of the natural person entity. Will be present if `structure` is equal to
	// `natural_person`.
	NaturalPerson *EntityNaturalPerson `json:"natural_person"`
	// Details of the joint entity. Will be present if `structure` is equal to `joint`.
	Joint *EntityJoint `json:"joint"`
	// Details of the trust entity. Will be present if `structure` is equal to `trust`.
	Trust *EntityTrust `json:"trust"`
	// A constant representing the object's type. For this resource it will always be
	// `entity`.
	Type *EntityType `json:"type"`
	// The entity's description for display purposes.
	Description *string `json:"description"`
	// The relationship between your group and the entity.
	Relationship *EntityRelationship `json:"relationship"`
	// Additional documentation associated with the entity.
	SupplementalDocuments *[]EntitySupplementalDocuments `json:"supplemental_documents"`
}

// The entity's identifier.
func (r *Entity) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The entity's legal structure.
func (r *Entity) GetStructure() (Structure EntityStructure) {
	if r != nil && r.Structure != nil {
		Structure = *r.Structure
	}
	return
}

// Details of the corporation entity. Will be present if `structure` is equal to
// `corporation`.
func (r *Entity) GetCorporation() (Corporation EntityCorporation) {
	if r != nil && r.Corporation != nil {
		Corporation = *r.Corporation
	}
	return
}

// Details of the natural person entity. Will be present if `structure` is equal to
// `natural_person`.
func (r *Entity) GetNaturalPerson() (NaturalPerson EntityNaturalPerson) {
	if r != nil && r.NaturalPerson != nil {
		NaturalPerson = *r.NaturalPerson
	}
	return
}

// Details of the joint entity. Will be present if `structure` is equal to `joint`.
func (r *Entity) GetJoint() (Joint EntityJoint) {
	if r != nil && r.Joint != nil {
		Joint = *r.Joint
	}
	return
}

// Details of the trust entity. Will be present if `structure` is equal to `trust`.
func (r *Entity) GetTrust() (Trust EntityTrust) {
	if r != nil && r.Trust != nil {
		Trust = *r.Trust
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `entity`.
func (r *Entity) GetType() (Type EntityType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

// The entity's description for display purposes.
func (r *Entity) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The relationship between your group and the entity.
func (r *Entity) GetRelationship() (Relationship EntityRelationship) {
	if r != nil && r.Relationship != nil {
		Relationship = *r.Relationship
	}
	return
}

// Additional documentation associated with the entity.
func (r *Entity) GetSupplementalDocuments() (SupplementalDocuments []EntitySupplementalDocuments) {
	if r != nil && r.SupplementalDocuments != nil {
		SupplementalDocuments = *r.SupplementalDocuments
	}
	return
}

type EntityStructure string

const (
	EntityStructureCorporation   EntityStructure = "corporation"
	EntityStructureNaturalPerson EntityStructure = "natural_person"
	EntityStructureJoint         EntityStructure = "joint"
	EntityStructureTrust         EntityStructure = "trust"
)

//
type EntityCorporation struct {
	// The legal name of the corporation.
	Name *string `json:"name"`
	// The website of the corporation.
	Website *string `json:"website"`
	// The Employer Identification Number (EIN) for the corporation.
	TaxIdentifier *string `json:"tax_identifier"`
	// The two-letter United States Postal Service (USPS) abbreviation for the
	// corporation's state of incorporation.
	IncorporationState *string `json:"incorporation_state"`
	// The corporation's address.
	Address *EntityCorporationAddress `json:"address"`
	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwners *[]EntityCorporationBeneficialOwners `json:"beneficial_owners"`
}

// The legal name of the corporation.
func (r *EntityCorporation) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The website of the corporation.
func (r *EntityCorporation) GetWebsite() (Website string) {
	if r != nil && r.Website != nil {
		Website = *r.Website
	}
	return
}

// The Employer Identification Number (EIN) for the corporation.
func (r *EntityCorporation) GetTaxIdentifier() (TaxIdentifier string) {
	if r != nil && r.TaxIdentifier != nil {
		TaxIdentifier = *r.TaxIdentifier
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the
// corporation's state of incorporation.
func (r *EntityCorporation) GetIncorporationState() (IncorporationState string) {
	if r != nil && r.IncorporationState != nil {
		IncorporationState = *r.IncorporationState
	}
	return
}

// The corporation's address.
func (r *EntityCorporation) GetAddress() (Address EntityCorporationAddress) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// The identifying details of anyone controlling or owning 25% or more of the
// corporation.
func (r *EntityCorporation) GetBeneficialOwners() (BeneficialOwners []EntityCorporationBeneficialOwners) {
	if r != nil && r.BeneficialOwners != nil {
		BeneficialOwners = *r.BeneficialOwners
	}
	return
}

//
type EntityCorporationAddress struct {
	// The first line of the address.
	Line1 *string `json:"line1"`
	// The second line of the address.
	Line2 *string `json:"line2"`
	// The city of the address.
	City *string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `json:"state"`
	// The ZIP code of the address.
	Zip *string `json:"zip"`
}

// The first line of the address.
func (r *EntityCorporationAddress) GetLine1() (Line1 string) {
	if r != nil && r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the address.
func (r *EntityCorporationAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the address.
func (r *EntityCorporationAddress) GetCity() (City string) {
	if r != nil && r.City != nil {
		City = *r.City
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the state of
// the address.
func (r *EntityCorporationAddress) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

// The ZIP code of the address.
func (r *EntityCorporationAddress) GetZip() (Zip string) {
	if r != nil && r.Zip != nil {
		Zip = *r.Zip
	}
	return
}

type EntityCorporationBeneficialOwners struct {
	// Personal details for the beneficial owner.
	Individual *EntityCorporationBeneficialOwnersIndividual `json:"individual"`
	// This person's role or title within the entity.
	CompanyTitle *string `json:"company_title"`
	// Why this person is considered a beneficial owner of the entity.
	Prong *EntityCorporationBeneficialOwnersProng `json:"prong"`
}

// Personal details for the beneficial owner.
func (r *EntityCorporationBeneficialOwners) GetIndividual() (Individual EntityCorporationBeneficialOwnersIndividual) {
	if r != nil && r.Individual != nil {
		Individual = *r.Individual
	}
	return
}

// This person's role or title within the entity.
func (r *EntityCorporationBeneficialOwners) GetCompanyTitle() (CompanyTitle string) {
	if r != nil && r.CompanyTitle != nil {
		CompanyTitle = *r.CompanyTitle
	}
	return
}

// Why this person is considered a beneficial owner of the entity.
func (r *EntityCorporationBeneficialOwners) GetProng() (Prong EntityCorporationBeneficialOwnersProng) {
	if r != nil && r.Prong != nil {
		Prong = *r.Prong
	}
	return
}

//
type EntityCorporationBeneficialOwnersIndividual struct {
	// The person's legal name.
	Name *string `json:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `json:"date_of_birth"`
	// The person's address.
	Address *EntityCorporationBeneficialOwnersIndividualAddress `json:"address"`
	// A means of verifying the person's identity.
	Identification *EntityCorporationBeneficialOwnersIndividualIdentification `json:"identification"`
}

// The person's legal name.
func (r *EntityCorporationBeneficialOwnersIndividual) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The person's date of birth in YYYY-MM-DD format.
func (r *EntityCorporationBeneficialOwnersIndividual) GetDateOfBirth() (DateOfBirth string) {
	if r != nil && r.DateOfBirth != nil {
		DateOfBirth = *r.DateOfBirth
	}
	return
}

// The person's address.
func (r *EntityCorporationBeneficialOwnersIndividual) GetAddress() (Address EntityCorporationBeneficialOwnersIndividualAddress) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// A means of verifying the person's identity.
func (r *EntityCorporationBeneficialOwnersIndividual) GetIdentification() (Identification EntityCorporationBeneficialOwnersIndividualIdentification) {
	if r != nil && r.Identification != nil {
		Identification = *r.Identification
	}
	return
}

//
type EntityCorporationBeneficialOwnersIndividualAddress struct {
	// The first line of the address.
	Line1 *string `json:"line1"`
	// The second line of the address.
	Line2 *string `json:"line2"`
	// The city of the address.
	City *string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `json:"state"`
	// The ZIP code of the address.
	Zip *string `json:"zip"`
}

// The first line of the address.
func (r *EntityCorporationBeneficialOwnersIndividualAddress) GetLine1() (Line1 string) {
	if r != nil && r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the address.
func (r *EntityCorporationBeneficialOwnersIndividualAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the address.
func (r *EntityCorporationBeneficialOwnersIndividualAddress) GetCity() (City string) {
	if r != nil && r.City != nil {
		City = *r.City
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the state of
// the address.
func (r *EntityCorporationBeneficialOwnersIndividualAddress) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

// The ZIP code of the address.
func (r *EntityCorporationBeneficialOwnersIndividualAddress) GetZip() (Zip string) {
	if r != nil && r.Zip != nil {
		Zip = *r.Zip
	}
	return
}

//
type EntityCorporationBeneficialOwnersIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method *EntityCorporationBeneficialOwnersIndividualIdentificationMethod `json:"method"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 *string `json:"number_last4"`
}

// A method that can be used to verify the individual's identity.
func (r *EntityCorporationBeneficialOwnersIndividualIdentification) GetMethod() (Method EntityCorporationBeneficialOwnersIndividualIdentificationMethod) {
	if r != nil && r.Method != nil {
		Method = *r.Method
	}
	return
}

// The last 4 digits of the identification number that can be used to verify the
// individual's identity.
func (r *EntityCorporationBeneficialOwnersIndividualIdentification) GetNumberLast4() (NumberLast4 string) {
	if r != nil && r.NumberLast4 != nil {
		NumberLast4 = *r.NumberLast4
	}
	return
}

type EntityCorporationBeneficialOwnersIndividualIdentificationMethod string

const (
	EntityCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber                   EntityCorporationBeneficialOwnersIndividualIdentificationMethod = "social_security_number"
	EntityCorporationBeneficialOwnersIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber EntityCorporationBeneficialOwnersIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	EntityCorporationBeneficialOwnersIndividualIdentificationMethodPassport                               EntityCorporationBeneficialOwnersIndividualIdentificationMethod = "passport"
)

type EntityCorporationBeneficialOwnersProng string

const (
	EntityCorporationBeneficialOwnersProngOwnership EntityCorporationBeneficialOwnersProng = "ownership"
	EntityCorporationBeneficialOwnersProngControl   EntityCorporationBeneficialOwnersProng = "control"
)

//
type EntityNaturalPerson struct {
	// The person's legal name.
	Name *string `json:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `json:"date_of_birth"`
	// The person's address.
	Address *EntityNaturalPersonAddress `json:"address"`
	// A means of verifying the person's identity.
	Identification *EntityNaturalPersonIdentification `json:"identification"`
}

// The person's legal name.
func (r *EntityNaturalPerson) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The person's date of birth in YYYY-MM-DD format.
func (r *EntityNaturalPerson) GetDateOfBirth() (DateOfBirth string) {
	if r != nil && r.DateOfBirth != nil {
		DateOfBirth = *r.DateOfBirth
	}
	return
}

// The person's address.
func (r *EntityNaturalPerson) GetAddress() (Address EntityNaturalPersonAddress) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// A means of verifying the person's identity.
func (r *EntityNaturalPerson) GetIdentification() (Identification EntityNaturalPersonIdentification) {
	if r != nil && r.Identification != nil {
		Identification = *r.Identification
	}
	return
}

//
type EntityNaturalPersonAddress struct {
	// The first line of the address.
	Line1 *string `json:"line1"`
	// The second line of the address.
	Line2 *string `json:"line2"`
	// The city of the address.
	City *string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `json:"state"`
	// The ZIP code of the address.
	Zip *string `json:"zip"`
}

// The first line of the address.
func (r *EntityNaturalPersonAddress) GetLine1() (Line1 string) {
	if r != nil && r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the address.
func (r *EntityNaturalPersonAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the address.
func (r *EntityNaturalPersonAddress) GetCity() (City string) {
	if r != nil && r.City != nil {
		City = *r.City
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the state of
// the address.
func (r *EntityNaturalPersonAddress) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

// The ZIP code of the address.
func (r *EntityNaturalPersonAddress) GetZip() (Zip string) {
	if r != nil && r.Zip != nil {
		Zip = *r.Zip
	}
	return
}

//
type EntityNaturalPersonIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method *EntityNaturalPersonIdentificationMethod `json:"method"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 *string `json:"number_last4"`
}

// A method that can be used to verify the individual's identity.
func (r *EntityNaturalPersonIdentification) GetMethod() (Method EntityNaturalPersonIdentificationMethod) {
	if r != nil && r.Method != nil {
		Method = *r.Method
	}
	return
}

// The last 4 digits of the identification number that can be used to verify the
// individual's identity.
func (r *EntityNaturalPersonIdentification) GetNumberLast4() (NumberLast4 string) {
	if r != nil && r.NumberLast4 != nil {
		NumberLast4 = *r.NumberLast4
	}
	return
}

type EntityNaturalPersonIdentificationMethod string

const (
	EntityNaturalPersonIdentificationMethodSocialSecurityNumber                   EntityNaturalPersonIdentificationMethod = "social_security_number"
	EntityNaturalPersonIdentificationMethodIndividualTaxpayerIdentificationNumber EntityNaturalPersonIdentificationMethod = "individual_taxpayer_identification_number"
	EntityNaturalPersonIdentificationMethodPassport                               EntityNaturalPersonIdentificationMethod = "passport"
)

//
type EntityJoint struct {
	// The entity's name.
	Name *string `json:"name"`
	// The two individuals that share control of the entity.
	Individuals *[]EntityJointIndividuals `json:"individuals"`
}

// The entity's name.
func (r *EntityJoint) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The two individuals that share control of the entity.
func (r *EntityJoint) GetIndividuals() (Individuals []EntityJointIndividuals) {
	if r != nil && r.Individuals != nil {
		Individuals = *r.Individuals
	}
	return
}

type EntityJointIndividuals struct {
	// The person's legal name.
	Name *string `json:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `json:"date_of_birth"`
	// The person's address.
	Address *EntityJointIndividualsAddress `json:"address"`
	// A means of verifying the person's identity.
	Identification *EntityJointIndividualsIdentification `json:"identification"`
}

// The person's legal name.
func (r *EntityJointIndividuals) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The person's date of birth in YYYY-MM-DD format.
func (r *EntityJointIndividuals) GetDateOfBirth() (DateOfBirth string) {
	if r != nil && r.DateOfBirth != nil {
		DateOfBirth = *r.DateOfBirth
	}
	return
}

// The person's address.
func (r *EntityJointIndividuals) GetAddress() (Address EntityJointIndividualsAddress) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// A means of verifying the person's identity.
func (r *EntityJointIndividuals) GetIdentification() (Identification EntityJointIndividualsIdentification) {
	if r != nil && r.Identification != nil {
		Identification = *r.Identification
	}
	return
}

//
type EntityJointIndividualsAddress struct {
	// The first line of the address.
	Line1 *string `json:"line1"`
	// The second line of the address.
	Line2 *string `json:"line2"`
	// The city of the address.
	City *string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `json:"state"`
	// The ZIP code of the address.
	Zip *string `json:"zip"`
}

// The first line of the address.
func (r *EntityJointIndividualsAddress) GetLine1() (Line1 string) {
	if r != nil && r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the address.
func (r *EntityJointIndividualsAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the address.
func (r *EntityJointIndividualsAddress) GetCity() (City string) {
	if r != nil && r.City != nil {
		City = *r.City
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the state of
// the address.
func (r *EntityJointIndividualsAddress) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

// The ZIP code of the address.
func (r *EntityJointIndividualsAddress) GetZip() (Zip string) {
	if r != nil && r.Zip != nil {
		Zip = *r.Zip
	}
	return
}

//
type EntityJointIndividualsIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method *EntityJointIndividualsIdentificationMethod `json:"method"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 *string `json:"number_last4"`
}

// A method that can be used to verify the individual's identity.
func (r *EntityJointIndividualsIdentification) GetMethod() (Method EntityJointIndividualsIdentificationMethod) {
	if r != nil && r.Method != nil {
		Method = *r.Method
	}
	return
}

// The last 4 digits of the identification number that can be used to verify the
// individual's identity.
func (r *EntityJointIndividualsIdentification) GetNumberLast4() (NumberLast4 string) {
	if r != nil && r.NumberLast4 != nil {
		NumberLast4 = *r.NumberLast4
	}
	return
}

type EntityJointIndividualsIdentificationMethod string

const (
	EntityJointIndividualsIdentificationMethodSocialSecurityNumber                   EntityJointIndividualsIdentificationMethod = "social_security_number"
	EntityJointIndividualsIdentificationMethodIndividualTaxpayerIdentificationNumber EntityJointIndividualsIdentificationMethod = "individual_taxpayer_identification_number"
	EntityJointIndividualsIdentificationMethodPassport                               EntityJointIndividualsIdentificationMethod = "passport"
)

//
type EntityTrust struct {
	// The trust's name
	Name *string `json:"name"`
	// Whether the trust is `revocable` or `irrevocable`.
	Category *EntityTrustCategory `json:"category"`
	// The trust's address.
	Address *EntityTrustAddress `json:"address"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state in
	// which the trust was formed.
	FormationState *string `json:"formation_state"`
	// The Employer Identification Number (EIN) of the trust itself.
	TaxIdentifier *string `json:"tax_identifier"`
	// The trustees of the trust.
	Trustees *[]EntityTrustTrustees `json:"trustees"`
	// The grantor of the trust. Will be present if the `category` is `revocable`.
	Grantor *EntityTrustGrantor `json:"grantor"`
	// The ID for the File containing the formation document of the trust.
	FormationDocumentFileId *string `json:"formation_document_file_id"`
}

// The trust's name
func (r *EntityTrust) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// Whether the trust is `revocable` or `irrevocable`.
func (r *EntityTrust) GetCategory() (Category EntityTrustCategory) {
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
func (r *EntityTrust) GetTrustees() (Trustees []EntityTrustTrustees) {
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

type EntityTrustCategory string

const (
	EntityTrustCategoryRevocable   EntityTrustCategory = "revocable"
	EntityTrustCategoryIrrevocable EntityTrustCategory = "irrevocable"
)

//
type EntityTrustAddress struct {
	// The first line of the address.
	Line1 *string `json:"line1"`
	// The second line of the address.
	Line2 *string `json:"line2"`
	// The city of the address.
	City *string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `json:"state"`
	// The ZIP code of the address.
	Zip *string `json:"zip"`
}

// The first line of the address.
func (r *EntityTrustAddress) GetLine1() (Line1 string) {
	if r != nil && r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the address.
func (r *EntityTrustAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the address.
func (r *EntityTrustAddress) GetCity() (City string) {
	if r != nil && r.City != nil {
		City = *r.City
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the state of
// the address.
func (r *EntityTrustAddress) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

// The ZIP code of the address.
func (r *EntityTrustAddress) GetZip() (Zip string) {
	if r != nil && r.Zip != nil {
		Zip = *r.Zip
	}
	return
}

type EntityTrustTrustees struct {
	// The structure of the trustee. Will always be equal to `individual`.
	Structure *EntityTrustTrusteesStructure `json:"structure"`
	// The individual trustee of the trust. Will be present if the trustee's
	// `structure` is equal to `individual`.
	Individual *EntityTrustTrusteesIndividual `json:"individual"`
}

// The structure of the trustee. Will always be equal to `individual`.
func (r *EntityTrustTrustees) GetStructure() (Structure EntityTrustTrusteesStructure) {
	if r != nil && r.Structure != nil {
		Structure = *r.Structure
	}
	return
}

// The individual trustee of the trust. Will be present if the trustee's
// `structure` is equal to `individual`.
func (r *EntityTrustTrustees) GetIndividual() (Individual EntityTrustTrusteesIndividual) {
	if r != nil && r.Individual != nil {
		Individual = *r.Individual
	}
	return
}

type EntityTrustTrusteesStructure string

const (
	EntityTrustTrusteesStructureIndividual EntityTrustTrusteesStructure = "individual"
)

//
type EntityTrustTrusteesIndividual struct {
	// The person's legal name.
	Name *string `json:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `json:"date_of_birth"`
	// The person's address.
	Address *EntityTrustTrusteesIndividualAddress `json:"address"`
	// A means of verifying the person's identity.
	Identification *EntityTrustTrusteesIndividualIdentification `json:"identification"`
}

// The person's legal name.
func (r *EntityTrustTrusteesIndividual) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The person's date of birth in YYYY-MM-DD format.
func (r *EntityTrustTrusteesIndividual) GetDateOfBirth() (DateOfBirth string) {
	if r != nil && r.DateOfBirth != nil {
		DateOfBirth = *r.DateOfBirth
	}
	return
}

// The person's address.
func (r *EntityTrustTrusteesIndividual) GetAddress() (Address EntityTrustTrusteesIndividualAddress) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// A means of verifying the person's identity.
func (r *EntityTrustTrusteesIndividual) GetIdentification() (Identification EntityTrustTrusteesIndividualIdentification) {
	if r != nil && r.Identification != nil {
		Identification = *r.Identification
	}
	return
}

//
type EntityTrustTrusteesIndividualAddress struct {
	// The first line of the address.
	Line1 *string `json:"line1"`
	// The second line of the address.
	Line2 *string `json:"line2"`
	// The city of the address.
	City *string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `json:"state"`
	// The ZIP code of the address.
	Zip *string `json:"zip"`
}

// The first line of the address.
func (r *EntityTrustTrusteesIndividualAddress) GetLine1() (Line1 string) {
	if r != nil && r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the address.
func (r *EntityTrustTrusteesIndividualAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the address.
func (r *EntityTrustTrusteesIndividualAddress) GetCity() (City string) {
	if r != nil && r.City != nil {
		City = *r.City
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the state of
// the address.
func (r *EntityTrustTrusteesIndividualAddress) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

// The ZIP code of the address.
func (r *EntityTrustTrusteesIndividualAddress) GetZip() (Zip string) {
	if r != nil && r.Zip != nil {
		Zip = *r.Zip
	}
	return
}

//
type EntityTrustTrusteesIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method *EntityTrustTrusteesIndividualIdentificationMethod `json:"method"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 *string `json:"number_last4"`
}

// A method that can be used to verify the individual's identity.
func (r *EntityTrustTrusteesIndividualIdentification) GetMethod() (Method EntityTrustTrusteesIndividualIdentificationMethod) {
	if r != nil && r.Method != nil {
		Method = *r.Method
	}
	return
}

// The last 4 digits of the identification number that can be used to verify the
// individual's identity.
func (r *EntityTrustTrusteesIndividualIdentification) GetNumberLast4() (NumberLast4 string) {
	if r != nil && r.NumberLast4 != nil {
		NumberLast4 = *r.NumberLast4
	}
	return
}

type EntityTrustTrusteesIndividualIdentificationMethod string

const (
	EntityTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber                   EntityTrustTrusteesIndividualIdentificationMethod = "social_security_number"
	EntityTrustTrusteesIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber EntityTrustTrusteesIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	EntityTrustTrusteesIndividualIdentificationMethodPassport                               EntityTrustTrusteesIndividualIdentificationMethod = "passport"
)

//
type EntityTrustGrantor struct {
	// The person's legal name.
	Name *string `json:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `json:"date_of_birth"`
	// The person's address.
	Address *EntityTrustGrantorAddress `json:"address"`
	// A means of verifying the person's identity.
	Identification *EntityTrustGrantorIdentification `json:"identification"`
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

//
type EntityTrustGrantorAddress struct {
	// The first line of the address.
	Line1 *string `json:"line1"`
	// The second line of the address.
	Line2 *string `json:"line2"`
	// The city of the address.
	City *string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `json:"state"`
	// The ZIP code of the address.
	Zip *string `json:"zip"`
}

// The first line of the address.
func (r *EntityTrustGrantorAddress) GetLine1() (Line1 string) {
	if r != nil && r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the address.
func (r *EntityTrustGrantorAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the address.
func (r *EntityTrustGrantorAddress) GetCity() (City string) {
	if r != nil && r.City != nil {
		City = *r.City
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the state of
// the address.
func (r *EntityTrustGrantorAddress) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

// The ZIP code of the address.
func (r *EntityTrustGrantorAddress) GetZip() (Zip string) {
	if r != nil && r.Zip != nil {
		Zip = *r.Zip
	}
	return
}

//
type EntityTrustGrantorIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method *EntityTrustGrantorIdentificationMethod `json:"method"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 *string `json:"number_last4"`
}

// A method that can be used to verify the individual's identity.
func (r *EntityTrustGrantorIdentification) GetMethod() (Method EntityTrustGrantorIdentificationMethod) {
	if r != nil && r.Method != nil {
		Method = *r.Method
	}
	return
}

// The last 4 digits of the identification number that can be used to verify the
// individual's identity.
func (r *EntityTrustGrantorIdentification) GetNumberLast4() (NumberLast4 string) {
	if r != nil && r.NumberLast4 != nil {
		NumberLast4 = *r.NumberLast4
	}
	return
}

type EntityTrustGrantorIdentificationMethod string

const (
	EntityTrustGrantorIdentificationMethodSocialSecurityNumber                   EntityTrustGrantorIdentificationMethod = "social_security_number"
	EntityTrustGrantorIdentificationMethodIndividualTaxpayerIdentificationNumber EntityTrustGrantorIdentificationMethod = "individual_taxpayer_identification_number"
	EntityTrustGrantorIdentificationMethodPassport                               EntityTrustGrantorIdentificationMethod = "passport"
)

type EntityType string

const (
	EntityTypeEntity EntityType = "entity"
)

type EntityRelationship string

const (
	EntityRelationshipAffiliated    EntityRelationship = "affiliated"
	EntityRelationshipInformational EntityRelationship = "informational"
	EntityRelationshipUnaffiliated  EntityRelationship = "unaffiliated"
)

type EntitySupplementalDocuments struct {
	// The File containing the document.
	FileId *string `json:"file_id"`
}

// The File containing the document.
func (r *EntitySupplementalDocuments) GetFileId() (FileId string) {
	if r != nil && r.FileId != nil {
		FileId = *r.FileId
	}
	return
}

//
type InboundDigitalWalletTokenRequestSimulationResult struct {
	// If the simulated tokenization attempt was declined, this field contains details
	// as to why.
	DeclineReason *InboundDigitalWalletTokenRequestSimulationResultDeclineReason `json:"decline_reason"`
	// If the simulated tokenization attempt was accepted, this field contains the id
	// of the Digital Wallet Token that was created.
	DigitalWalletTokenId *string `json:"digital_wallet_token_id"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_digital_wallet_token_request_simulation_result`.
	Type *InboundDigitalWalletTokenRequestSimulationResultType `json:"type"`
}

// If the simulated tokenization attempt was declined, this field contains details
// as to why.
func (r *InboundDigitalWalletTokenRequestSimulationResult) GetDeclineReason() (DeclineReason InboundDigitalWalletTokenRequestSimulationResultDeclineReason) {
	if r != nil && r.DeclineReason != nil {
		DeclineReason = *r.DeclineReason
	}
	return
}

// If the simulated tokenization attempt was accepted, this field contains the id
// of the Digital Wallet Token that was created.
func (r *InboundDigitalWalletTokenRequestSimulationResult) GetDigitalWalletTokenId() (DigitalWalletTokenId string) {
	if r != nil && r.DigitalWalletTokenId != nil {
		DigitalWalletTokenId = *r.DigitalWalletTokenId
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `inbound_digital_wallet_token_request_simulation_result`.
func (r *InboundDigitalWalletTokenRequestSimulationResult) GetType() (Type InboundDigitalWalletTokenRequestSimulationResultType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type InboundDigitalWalletTokenRequestSimulationResultDeclineReason string

const (
	InboundDigitalWalletTokenRequestSimulationResultDeclineReasonCardNotActive        InboundDigitalWalletTokenRequestSimulationResultDeclineReason = "card_not_active"
	InboundDigitalWalletTokenRequestSimulationResultDeclineReasonNoVerificationMethod InboundDigitalWalletTokenRequestSimulationResultDeclineReason = "no_verification_method"
	InboundDigitalWalletTokenRequestSimulationResultDeclineReasonWebhookTimedOut      InboundDigitalWalletTokenRequestSimulationResultDeclineReason = "webhook_timed_out"
	InboundDigitalWalletTokenRequestSimulationResultDeclineReasonWebhookDeclined      InboundDigitalWalletTokenRequestSimulationResultDeclineReason = "webhook_declined"
)

type InboundDigitalWalletTokenRequestSimulationResultType string

const (
	InboundDigitalWalletTokenRequestSimulationResultTypeInboundDigitalWalletTokenRequestSimulationResult InboundDigitalWalletTokenRequestSimulationResultType = "inbound_digital_wallet_token_request_simulation_result"
)
