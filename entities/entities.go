package entities

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type EntityService struct {
	Requester             core.Requester
	get                   func(context.Context, string, *core.CoreRequest, interface{}) error
	post                  func(context.Context, string, *core.CoreRequest, interface{}) error
	patch                 func(context.Context, string, *core.CoreRequest, interface{}) error
	put                   func(context.Context, string, *core.CoreRequest, interface{}) error
	delete                func(context.Context, string, *core.CoreRequest, interface{}) error
	SupplementalDocuments *SupplementalDocumentService
}

func NewEntityService(requester core.Requester) (r *EntityService) {
	r = &EntityService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	r.SupplementalDocuments = NewSupplementalDocumentService(requester)
	return
}

type PreloadedEntityService struct {
	Entities              *EntityService
	SupplementalDocuments *PreloadedSupplementalDocumentService
}

func (r *PreloadedEntityService) Init(service *EntityService) {
	r.Entities = service
	r.SupplementalDocuments = &PreloadedSupplementalDocumentService{}
	r.SupplementalDocuments.Init(r.Entities.SupplementalDocuments)
}

func NewPreloadedEntityService(service *EntityService) (r *PreloadedEntityService) {
	r = &PreloadedEntityService{}
	r.Init(service)
	return
}

type CreateASupplementalDocumentForAnEntityParameters struct {
	// The identifier of the File containing the document.
	FileID *string `json:"file_id"`
}

// The identifier of the File containing the document.
func (r *CreateASupplementalDocumentForAnEntityParameters) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

//
type Entity struct {
	// The entity's identifier.
	ID *string `json:"id"`
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
func (r *Entity) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
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
	FormationDocumentFileID *string `json:"formation_document_file_id"`
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
func (r *EntityTrust) GetFormationDocumentFileID() (FormationDocumentFileID string) {
	if r != nil && r.FormationDocumentFileID != nil {
		FormationDocumentFileID = *r.FormationDocumentFileID
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
	FileID *string `json:"file_id"`
}

// The File containing the document.
func (r *EntitySupplementalDocuments) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

type CreateAnEntityParameters struct {
	// The type of Entity to create.
	Structure *CreateAnEntityParametersStructure `json:"structure"`
	// Details of the corporation entity to create. Required if `structure` is equal to
	// `corporation`.
	Corporation *CreateAnEntityParametersCorporation `json:"corporation,omitempty"`
	// Details of the natural person entity to create. Required if `structure` is equal
	// to `natural_person`.
	NaturalPerson *CreateAnEntityParametersNaturalPerson `json:"natural_person,omitempty"`
	// Details of the joint entity to create. Required if `structure` is equal to
	// `joint`.
	Joint *CreateAnEntityParametersJoint `json:"joint,omitempty"`
	// Details of the trust entity to create. Required if `structure` is equal to
	// `trust`.
	Trust *CreateAnEntityParametersTrust `json:"trust,omitempty"`
	// The description you choose to give the entity.
	Description *string `json:"description,omitempty"`
	// The relationship between your group and the entity.
	Relationship *CreateAnEntityParametersRelationship `json:"relationship"`
	// Additional documentation associated with the entity.
	SupplementalDocuments *[]CreateAnEntityParametersSupplementalDocuments `json:"supplemental_documents,omitempty"`
}

// The type of Entity to create.
func (r *CreateAnEntityParameters) GetStructure() (Structure CreateAnEntityParametersStructure) {
	if r != nil && r.Structure != nil {
		Structure = *r.Structure
	}
	return
}

// Details of the corporation entity to create. Required if `structure` is equal to
// `corporation`.
func (r *CreateAnEntityParameters) GetCorporation() (Corporation CreateAnEntityParametersCorporation) {
	if r != nil && r.Corporation != nil {
		Corporation = *r.Corporation
	}
	return
}

// Details of the natural person entity to create. Required if `structure` is equal
// to `natural_person`.
func (r *CreateAnEntityParameters) GetNaturalPerson() (NaturalPerson CreateAnEntityParametersNaturalPerson) {
	if r != nil && r.NaturalPerson != nil {
		NaturalPerson = *r.NaturalPerson
	}
	return
}

// Details of the joint entity to create. Required if `structure` is equal to
// `joint`.
func (r *CreateAnEntityParameters) GetJoint() (Joint CreateAnEntityParametersJoint) {
	if r != nil && r.Joint != nil {
		Joint = *r.Joint
	}
	return
}

// Details of the trust entity to create. Required if `structure` is equal to
// `trust`.
func (r *CreateAnEntityParameters) GetTrust() (Trust CreateAnEntityParametersTrust) {
	if r != nil && r.Trust != nil {
		Trust = *r.Trust
	}
	return
}

// The description you choose to give the entity.
func (r *CreateAnEntityParameters) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The relationship between your group and the entity.
func (r *CreateAnEntityParameters) GetRelationship() (Relationship CreateAnEntityParametersRelationship) {
	if r != nil && r.Relationship != nil {
		Relationship = *r.Relationship
	}
	return
}

// Additional documentation associated with the entity.
func (r *CreateAnEntityParameters) GetSupplementalDocuments() (SupplementalDocuments []CreateAnEntityParametersSupplementalDocuments) {
	if r != nil && r.SupplementalDocuments != nil {
		SupplementalDocuments = *r.SupplementalDocuments
	}
	return
}

type CreateAnEntityParametersStructure string

const (
	CreateAnEntityParametersStructureCorporation   CreateAnEntityParametersStructure = "corporation"
	CreateAnEntityParametersStructureNaturalPerson CreateAnEntityParametersStructure = "natural_person"
	CreateAnEntityParametersStructureJoint         CreateAnEntityParametersStructure = "joint"
	CreateAnEntityParametersStructureTrust         CreateAnEntityParametersStructure = "trust"
)

//
type CreateAnEntityParametersCorporation struct {
	// The legal name of the corporation.
	Name *string `json:"name"`
	// The website of the corporation.
	Website *string `json:"website,omitempty"`
	// The Employer Identification Number (EIN) for the corporation.
	TaxIdentifier *string `json:"tax_identifier"`
	// The two-letter United States Postal Service (USPS) abbreviation for the
	// corporation's state of incorporation.
	IncorporationState *string `json:"incorporation_state,omitempty"`
	// The corporation's address.
	Address *CreateAnEntityParametersCorporationAddress `json:"address"`
	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwners *[]CreateAnEntityParametersCorporationBeneficialOwners `json:"beneficial_owners"`
}

// The legal name of the corporation.
func (r *CreateAnEntityParametersCorporation) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The website of the corporation.
func (r *CreateAnEntityParametersCorporation) GetWebsite() (Website string) {
	if r != nil && r.Website != nil {
		Website = *r.Website
	}
	return
}

// The Employer Identification Number (EIN) for the corporation.
func (r *CreateAnEntityParametersCorporation) GetTaxIdentifier() (TaxIdentifier string) {
	if r != nil && r.TaxIdentifier != nil {
		TaxIdentifier = *r.TaxIdentifier
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the
// corporation's state of incorporation.
func (r *CreateAnEntityParametersCorporation) GetIncorporationState() (IncorporationState string) {
	if r != nil && r.IncorporationState != nil {
		IncorporationState = *r.IncorporationState
	}
	return
}

// The corporation's address.
func (r *CreateAnEntityParametersCorporation) GetAddress() (Address CreateAnEntityParametersCorporationAddress) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// The identifying details of anyone controlling or owning 25% or more of the
// corporation.
func (r *CreateAnEntityParametersCorporation) GetBeneficialOwners() (BeneficialOwners []CreateAnEntityParametersCorporationBeneficialOwners) {
	if r != nil && r.BeneficialOwners != nil {
		BeneficialOwners = *r.BeneficialOwners
	}
	return
}

//
type CreateAnEntityParametersCorporationAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 *string `json:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 *string `json:"line2,omitempty"`
	// The city of the address.
	City *string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `json:"state"`
	// The ZIP code of the address.
	Zip *string `json:"zip"`
}

// The first line of the address. This is usually the street number and street.
func (r *CreateAnEntityParametersCorporationAddress) GetLine1() (Line1 string) {
	if r != nil && r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the address. This might be the floor or room number.
func (r *CreateAnEntityParametersCorporationAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the address.
func (r *CreateAnEntityParametersCorporationAddress) GetCity() (City string) {
	if r != nil && r.City != nil {
		City = *r.City
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the state of
// the address.
func (r *CreateAnEntityParametersCorporationAddress) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

// The ZIP code of the address.
func (r *CreateAnEntityParametersCorporationAddress) GetZip() (Zip string) {
	if r != nil && r.Zip != nil {
		Zip = *r.Zip
	}
	return
}

type CreateAnEntityParametersCorporationBeneficialOwners struct {
	// Personal details for the beneficial owner.
	Individual *CreateAnEntityParametersCorporationBeneficialOwnersIndividual `json:"individual"`
	// This person's role or title within the entity.
	CompanyTitle *string `json:"company_title,omitempty"`
	// Why this person is considered a beneficial owner of the entity.
	Prong *CreateAnEntityParametersCorporationBeneficialOwnersProng `json:"prong"`
}

// Personal details for the beneficial owner.
func (r *CreateAnEntityParametersCorporationBeneficialOwners) GetIndividual() (Individual CreateAnEntityParametersCorporationBeneficialOwnersIndividual) {
	if r != nil && r.Individual != nil {
		Individual = *r.Individual
	}
	return
}

// This person's role or title within the entity.
func (r *CreateAnEntityParametersCorporationBeneficialOwners) GetCompanyTitle() (CompanyTitle string) {
	if r != nil && r.CompanyTitle != nil {
		CompanyTitle = *r.CompanyTitle
	}
	return
}

// Why this person is considered a beneficial owner of the entity.
func (r *CreateAnEntityParametersCorporationBeneficialOwners) GetProng() (Prong CreateAnEntityParametersCorporationBeneficialOwnersProng) {
	if r != nil && r.Prong != nil {
		Prong = *r.Prong
	}
	return
}

//
type CreateAnEntityParametersCorporationBeneficialOwnersIndividual struct {
	// The person's legal name.
	Name *string `json:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `json:"date_of_birth"`
	// The individual's address.
	Address *CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress `json:"address"`
	// A means of verifying the person's identity.
	Identification *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification `json:"identification"`
}

// The person's legal name.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividual) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The person's date of birth in YYYY-MM-DD format.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividual) GetDateOfBirth() (DateOfBirth string) {
	if r != nil && r.DateOfBirth != nil {
		DateOfBirth = *r.DateOfBirth
	}
	return
}

// The individual's address.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividual) GetAddress() (Address CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// A means of verifying the person's identity.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividual) GetIdentification() (Identification CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification) {
	if r != nil && r.Identification != nil {
		Identification = *r.Identification
	}
	return
}

//
type CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 *string `json:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 *string `json:"line2,omitempty"`
	// The city of the address.
	City *string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `json:"state"`
	// The ZIP code of the address.
	Zip *string `json:"zip"`
}

// The first line of the address. This is usually the street number and street.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress) GetLine1() (Line1 string) {
	if r != nil && r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the address. This might be the floor or room number.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the address.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress) GetCity() (City string) {
	if r != nil && r.City != nil {
		City = *r.City
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the state of
// the address.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

// The ZIP code of the address.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress) GetZip() (Zip string) {
	if r != nil && r.Zip != nil {
		Zip = *r.Zip
	}
	return
}

//
type CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethod `json:"method"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number *string `json:"number"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport `json:"passport,omitempty"`
}

// A method that can be used to verify the individual's identity.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification) GetMethod() (Method CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethod) {
	if r != nil && r.Method != nil {
		Method = *r.Method
	}
	return
}

// An identification number that can be used to verify the individual's identity,
// such as a social security number.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification) GetNumber() (Number string) {
	if r != nil && r.Number != nil {
		Number = *r.Number
	}
	return
}

// Information about the passport used for identification. Required if `method` is
// equal to `passport`.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification) GetPassport() (Passport CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport) {
	if r != nil && r.Passport != nil {
		Passport = *r.Passport
	}
	return
}

type CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethod string

const (
	CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber                   CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethod = "social_security_number"
	CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethodPassport                               CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethod = "passport"
)

//
type CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport struct {
	// The identifier of the File containing the passport.
	FileID *string `json:"file_id"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `json:"expiration_date"`
	// The country that issued the passport.
	Country *string `json:"country"`
}

// The identifier of the File containing the passport.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

// The passport's expiration date in YYYY-MM-DD format.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport) GetExpirationDate() (ExpirationDate string) {
	if r != nil && r.ExpirationDate != nil {
		ExpirationDate = *r.ExpirationDate
	}
	return
}

// The country that issued the passport.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport) GetCountry() (Country string) {
	if r != nil && r.Country != nil {
		Country = *r.Country
	}
	return
}

type CreateAnEntityParametersCorporationBeneficialOwnersProng string

const (
	CreateAnEntityParametersCorporationBeneficialOwnersProngOwnership CreateAnEntityParametersCorporationBeneficialOwnersProng = "ownership"
	CreateAnEntityParametersCorporationBeneficialOwnersProngControl   CreateAnEntityParametersCorporationBeneficialOwnersProng = "control"
)

//
type CreateAnEntityParametersNaturalPerson struct {
	// The person's legal name.
	Name *string `json:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `json:"date_of_birth"`
	// The individual's address.
	Address *CreateAnEntityParametersNaturalPersonAddress `json:"address"`
	// A means of verifying the person's identity.
	Identification *CreateAnEntityParametersNaturalPersonIdentification `json:"identification"`
}

// The person's legal name.
func (r *CreateAnEntityParametersNaturalPerson) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The person's date of birth in YYYY-MM-DD format.
func (r *CreateAnEntityParametersNaturalPerson) GetDateOfBirth() (DateOfBirth string) {
	if r != nil && r.DateOfBirth != nil {
		DateOfBirth = *r.DateOfBirth
	}
	return
}

// The individual's address.
func (r *CreateAnEntityParametersNaturalPerson) GetAddress() (Address CreateAnEntityParametersNaturalPersonAddress) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// A means of verifying the person's identity.
func (r *CreateAnEntityParametersNaturalPerson) GetIdentification() (Identification CreateAnEntityParametersNaturalPersonIdentification) {
	if r != nil && r.Identification != nil {
		Identification = *r.Identification
	}
	return
}

//
type CreateAnEntityParametersNaturalPersonAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 *string `json:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 *string `json:"line2,omitempty"`
	// The city of the address.
	City *string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `json:"state"`
	// The ZIP code of the address.
	Zip *string `json:"zip"`
}

// The first line of the address. This is usually the street number and street.
func (r *CreateAnEntityParametersNaturalPersonAddress) GetLine1() (Line1 string) {
	if r != nil && r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the address. This might be the floor or room number.
func (r *CreateAnEntityParametersNaturalPersonAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the address.
func (r *CreateAnEntityParametersNaturalPersonAddress) GetCity() (City string) {
	if r != nil && r.City != nil {
		City = *r.City
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the state of
// the address.
func (r *CreateAnEntityParametersNaturalPersonAddress) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

// The ZIP code of the address.
func (r *CreateAnEntityParametersNaturalPersonAddress) GetZip() (Zip string) {
	if r != nil && r.Zip != nil {
		Zip = *r.Zip
	}
	return
}

//
type CreateAnEntityParametersNaturalPersonIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method *CreateAnEntityParametersNaturalPersonIdentificationMethod `json:"method"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number *string `json:"number"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport *CreateAnEntityParametersNaturalPersonIdentificationPassport `json:"passport,omitempty"`
}

// A method that can be used to verify the individual's identity.
func (r *CreateAnEntityParametersNaturalPersonIdentification) GetMethod() (Method CreateAnEntityParametersNaturalPersonIdentificationMethod) {
	if r != nil && r.Method != nil {
		Method = *r.Method
	}
	return
}

// An identification number that can be used to verify the individual's identity,
// such as a social security number.
func (r *CreateAnEntityParametersNaturalPersonIdentification) GetNumber() (Number string) {
	if r != nil && r.Number != nil {
		Number = *r.Number
	}
	return
}

// Information about the passport used for identification. Required if `method` is
// equal to `passport`.
func (r *CreateAnEntityParametersNaturalPersonIdentification) GetPassport() (Passport CreateAnEntityParametersNaturalPersonIdentificationPassport) {
	if r != nil && r.Passport != nil {
		Passport = *r.Passport
	}
	return
}

type CreateAnEntityParametersNaturalPersonIdentificationMethod string

const (
	CreateAnEntityParametersNaturalPersonIdentificationMethodSocialSecurityNumber                   CreateAnEntityParametersNaturalPersonIdentificationMethod = "social_security_number"
	CreateAnEntityParametersNaturalPersonIdentificationMethodIndividualTaxpayerIdentificationNumber CreateAnEntityParametersNaturalPersonIdentificationMethod = "individual_taxpayer_identification_number"
	CreateAnEntityParametersNaturalPersonIdentificationMethodPassport                               CreateAnEntityParametersNaturalPersonIdentificationMethod = "passport"
)

//
type CreateAnEntityParametersNaturalPersonIdentificationPassport struct {
	// The identifier of the File containing the passport.
	FileID *string `json:"file_id"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `json:"expiration_date"`
	// The country that issued the passport.
	Country *string `json:"country"`
}

// The identifier of the File containing the passport.
func (r *CreateAnEntityParametersNaturalPersonIdentificationPassport) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

// The passport's expiration date in YYYY-MM-DD format.
func (r *CreateAnEntityParametersNaturalPersonIdentificationPassport) GetExpirationDate() (ExpirationDate string) {
	if r != nil && r.ExpirationDate != nil {
		ExpirationDate = *r.ExpirationDate
	}
	return
}

// The country that issued the passport.
func (r *CreateAnEntityParametersNaturalPersonIdentificationPassport) GetCountry() (Country string) {
	if r != nil && r.Country != nil {
		Country = *r.Country
	}
	return
}

//
type CreateAnEntityParametersJoint struct {
	// The name of the joint entity.
	Name *string `json:"name,omitempty"`
	// The two individuals that share control of the entity.
	Individuals *[]CreateAnEntityParametersJointIndividuals `json:"individuals"`
}

// The name of the joint entity.
func (r *CreateAnEntityParametersJoint) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The two individuals that share control of the entity.
func (r *CreateAnEntityParametersJoint) GetIndividuals() (Individuals []CreateAnEntityParametersJointIndividuals) {
	if r != nil && r.Individuals != nil {
		Individuals = *r.Individuals
	}
	return
}

type CreateAnEntityParametersJointIndividuals struct {
	// The person's legal name.
	Name *string `json:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `json:"date_of_birth"`
	// The individual's address.
	Address *CreateAnEntityParametersJointIndividualsAddress `json:"address"`
	// A means of verifying the person's identity.
	Identification *CreateAnEntityParametersJointIndividualsIdentification `json:"identification"`
}

// The person's legal name.
func (r *CreateAnEntityParametersJointIndividuals) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The person's date of birth in YYYY-MM-DD format.
func (r *CreateAnEntityParametersJointIndividuals) GetDateOfBirth() (DateOfBirth string) {
	if r != nil && r.DateOfBirth != nil {
		DateOfBirth = *r.DateOfBirth
	}
	return
}

// The individual's address.
func (r *CreateAnEntityParametersJointIndividuals) GetAddress() (Address CreateAnEntityParametersJointIndividualsAddress) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// A means of verifying the person's identity.
func (r *CreateAnEntityParametersJointIndividuals) GetIdentification() (Identification CreateAnEntityParametersJointIndividualsIdentification) {
	if r != nil && r.Identification != nil {
		Identification = *r.Identification
	}
	return
}

//
type CreateAnEntityParametersJointIndividualsAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 *string `json:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 *string `json:"line2,omitempty"`
	// The city of the address.
	City *string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `json:"state"`
	// The ZIP code of the address.
	Zip *string `json:"zip"`
}

// The first line of the address. This is usually the street number and street.
func (r *CreateAnEntityParametersJointIndividualsAddress) GetLine1() (Line1 string) {
	if r != nil && r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the address. This might be the floor or room number.
func (r *CreateAnEntityParametersJointIndividualsAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the address.
func (r *CreateAnEntityParametersJointIndividualsAddress) GetCity() (City string) {
	if r != nil && r.City != nil {
		City = *r.City
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the state of
// the address.
func (r *CreateAnEntityParametersJointIndividualsAddress) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

// The ZIP code of the address.
func (r *CreateAnEntityParametersJointIndividualsAddress) GetZip() (Zip string) {
	if r != nil && r.Zip != nil {
		Zip = *r.Zip
	}
	return
}

//
type CreateAnEntityParametersJointIndividualsIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method *CreateAnEntityParametersJointIndividualsIdentificationMethod `json:"method"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number *string `json:"number"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport *CreateAnEntityParametersJointIndividualsIdentificationPassport `json:"passport,omitempty"`
}

// A method that can be used to verify the individual's identity.
func (r *CreateAnEntityParametersJointIndividualsIdentification) GetMethod() (Method CreateAnEntityParametersJointIndividualsIdentificationMethod) {
	if r != nil && r.Method != nil {
		Method = *r.Method
	}
	return
}

// An identification number that can be used to verify the individual's identity,
// such as a social security number.
func (r *CreateAnEntityParametersJointIndividualsIdentification) GetNumber() (Number string) {
	if r != nil && r.Number != nil {
		Number = *r.Number
	}
	return
}

// Information about the passport used for identification. Required if `method` is
// equal to `passport`.
func (r *CreateAnEntityParametersJointIndividualsIdentification) GetPassport() (Passport CreateAnEntityParametersJointIndividualsIdentificationPassport) {
	if r != nil && r.Passport != nil {
		Passport = *r.Passport
	}
	return
}

type CreateAnEntityParametersJointIndividualsIdentificationMethod string

const (
	CreateAnEntityParametersJointIndividualsIdentificationMethodSocialSecurityNumber                   CreateAnEntityParametersJointIndividualsIdentificationMethod = "social_security_number"
	CreateAnEntityParametersJointIndividualsIdentificationMethodIndividualTaxpayerIdentificationNumber CreateAnEntityParametersJointIndividualsIdentificationMethod = "individual_taxpayer_identification_number"
	CreateAnEntityParametersJointIndividualsIdentificationMethodPassport                               CreateAnEntityParametersJointIndividualsIdentificationMethod = "passport"
)

//
type CreateAnEntityParametersJointIndividualsIdentificationPassport struct {
	// The identifier of the File containing the passport.
	FileID *string `json:"file_id"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `json:"expiration_date"`
	// The country that issued the passport.
	Country *string `json:"country"`
}

// The identifier of the File containing the passport.
func (r *CreateAnEntityParametersJointIndividualsIdentificationPassport) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

// The passport's expiration date in YYYY-MM-DD format.
func (r *CreateAnEntityParametersJointIndividualsIdentificationPassport) GetExpirationDate() (ExpirationDate string) {
	if r != nil && r.ExpirationDate != nil {
		ExpirationDate = *r.ExpirationDate
	}
	return
}

// The country that issued the passport.
func (r *CreateAnEntityParametersJointIndividualsIdentificationPassport) GetCountry() (Country string) {
	if r != nil && r.Country != nil {
		Country = *r.Country
	}
	return
}

//
type CreateAnEntityParametersTrust struct {
	// The legal name of the trust.
	Name *string `json:"name"`
	// Whether the trust is `revocable` or `irrevocable`. Irrevocable trusts require
	// their own Employer Identification Number. Revocable trusts require information
	// about the individual `grantor` who created the trust.
	Category *CreateAnEntityParametersTrustCategory `json:"category"`
	// The Employer Identification Number (EIN) for the trust. Required if `category`
	// is equal to `irrevocable`.
	TaxIdentifier *string `json:"tax_identifier,omitempty"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state in
	// which the trust was formed.
	FormationState *string `json:"formation_state,omitempty"`
	// The trust's address.
	Address *CreateAnEntityParametersTrustAddress `json:"address"`
	// The identifier of the File containing the formation document of the trust.
	FormationDocumentFileID *string `json:"formation_document_file_id,omitempty"`
	// The trustees of the trust.
	Trustees *[]CreateAnEntityParametersTrustTrustees `json:"trustees"`
	// The grantor of the trust. Required if `category` is equal to `revocable`.
	Grantor *CreateAnEntityParametersTrustGrantor `json:"grantor,omitempty"`
}

// The legal name of the trust.
func (r *CreateAnEntityParametersTrust) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// Whether the trust is `revocable` or `irrevocable`. Irrevocable trusts require
// their own Employer Identification Number. Revocable trusts require information
// about the individual `grantor` who created the trust.
func (r *CreateAnEntityParametersTrust) GetCategory() (Category CreateAnEntityParametersTrustCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// The Employer Identification Number (EIN) for the trust. Required if `category`
// is equal to `irrevocable`.
func (r *CreateAnEntityParametersTrust) GetTaxIdentifier() (TaxIdentifier string) {
	if r != nil && r.TaxIdentifier != nil {
		TaxIdentifier = *r.TaxIdentifier
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the state in
// which the trust was formed.
func (r *CreateAnEntityParametersTrust) GetFormationState() (FormationState string) {
	if r != nil && r.FormationState != nil {
		FormationState = *r.FormationState
	}
	return
}

// The trust's address.
func (r *CreateAnEntityParametersTrust) GetAddress() (Address CreateAnEntityParametersTrustAddress) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// The identifier of the File containing the formation document of the trust.
func (r *CreateAnEntityParametersTrust) GetFormationDocumentFileID() (FormationDocumentFileID string) {
	if r != nil && r.FormationDocumentFileID != nil {
		FormationDocumentFileID = *r.FormationDocumentFileID
	}
	return
}

// The trustees of the trust.
func (r *CreateAnEntityParametersTrust) GetTrustees() (Trustees []CreateAnEntityParametersTrustTrustees) {
	if r != nil && r.Trustees != nil {
		Trustees = *r.Trustees
	}
	return
}

// The grantor of the trust. Required if `category` is equal to `revocable`.
func (r *CreateAnEntityParametersTrust) GetGrantor() (Grantor CreateAnEntityParametersTrustGrantor) {
	if r != nil && r.Grantor != nil {
		Grantor = *r.Grantor
	}
	return
}

type CreateAnEntityParametersTrustCategory string

const (
	CreateAnEntityParametersTrustCategoryRevocable   CreateAnEntityParametersTrustCategory = "revocable"
	CreateAnEntityParametersTrustCategoryIrrevocable CreateAnEntityParametersTrustCategory = "irrevocable"
)

//
type CreateAnEntityParametersTrustAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 *string `json:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 *string `json:"line2,omitempty"`
	// The city of the address.
	City *string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `json:"state"`
	// The ZIP code of the address.
	Zip *string `json:"zip"`
}

// The first line of the address. This is usually the street number and street.
func (r *CreateAnEntityParametersTrustAddress) GetLine1() (Line1 string) {
	if r != nil && r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the address. This might be the floor or room number.
func (r *CreateAnEntityParametersTrustAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the address.
func (r *CreateAnEntityParametersTrustAddress) GetCity() (City string) {
	if r != nil && r.City != nil {
		City = *r.City
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the state of
// the address.
func (r *CreateAnEntityParametersTrustAddress) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

// The ZIP code of the address.
func (r *CreateAnEntityParametersTrustAddress) GetZip() (Zip string) {
	if r != nil && r.Zip != nil {
		Zip = *r.Zip
	}
	return
}

type CreateAnEntityParametersTrustTrustees struct {
	// The structure of the trustee.
	Structure *CreateAnEntityParametersTrustTrusteesStructure `json:"structure"`
	// Details of the individual trustee. Required when the trustee `structure` is
	// equal to `individual`.
	Individual *CreateAnEntityParametersTrustTrusteesIndividual `json:"individual,omitempty"`
}

// The structure of the trustee.
func (r *CreateAnEntityParametersTrustTrustees) GetStructure() (Structure CreateAnEntityParametersTrustTrusteesStructure) {
	if r != nil && r.Structure != nil {
		Structure = *r.Structure
	}
	return
}

// Details of the individual trustee. Required when the trustee `structure` is
// equal to `individual`.
func (r *CreateAnEntityParametersTrustTrustees) GetIndividual() (Individual CreateAnEntityParametersTrustTrusteesIndividual) {
	if r != nil && r.Individual != nil {
		Individual = *r.Individual
	}
	return
}

type CreateAnEntityParametersTrustTrusteesStructure string

const (
	CreateAnEntityParametersTrustTrusteesStructureIndividual CreateAnEntityParametersTrustTrusteesStructure = "individual"
)

//
type CreateAnEntityParametersTrustTrusteesIndividual struct {
	// The person's legal name.
	Name *string `json:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `json:"date_of_birth"`
	// The individual's address.
	Address *CreateAnEntityParametersTrustTrusteesIndividualAddress `json:"address"`
	// A means of verifying the person's identity.
	Identification *CreateAnEntityParametersTrustTrusteesIndividualIdentification `json:"identification"`
}

// The person's legal name.
func (r *CreateAnEntityParametersTrustTrusteesIndividual) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The person's date of birth in YYYY-MM-DD format.
func (r *CreateAnEntityParametersTrustTrusteesIndividual) GetDateOfBirth() (DateOfBirth string) {
	if r != nil && r.DateOfBirth != nil {
		DateOfBirth = *r.DateOfBirth
	}
	return
}

// The individual's address.
func (r *CreateAnEntityParametersTrustTrusteesIndividual) GetAddress() (Address CreateAnEntityParametersTrustTrusteesIndividualAddress) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// A means of verifying the person's identity.
func (r *CreateAnEntityParametersTrustTrusteesIndividual) GetIdentification() (Identification CreateAnEntityParametersTrustTrusteesIndividualIdentification) {
	if r != nil && r.Identification != nil {
		Identification = *r.Identification
	}
	return
}

//
type CreateAnEntityParametersTrustTrusteesIndividualAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 *string `json:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 *string `json:"line2,omitempty"`
	// The city of the address.
	City *string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `json:"state"`
	// The ZIP code of the address.
	Zip *string `json:"zip"`
}

// The first line of the address. This is usually the street number and street.
func (r *CreateAnEntityParametersTrustTrusteesIndividualAddress) GetLine1() (Line1 string) {
	if r != nil && r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the address. This might be the floor or room number.
func (r *CreateAnEntityParametersTrustTrusteesIndividualAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the address.
func (r *CreateAnEntityParametersTrustTrusteesIndividualAddress) GetCity() (City string) {
	if r != nil && r.City != nil {
		City = *r.City
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the state of
// the address.
func (r *CreateAnEntityParametersTrustTrusteesIndividualAddress) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

// The ZIP code of the address.
func (r *CreateAnEntityParametersTrustTrusteesIndividualAddress) GetZip() (Zip string) {
	if r != nil && r.Zip != nil {
		Zip = *r.Zip
	}
	return
}

//
type CreateAnEntityParametersTrustTrusteesIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method *CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethod `json:"method"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number *string `json:"number"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport *CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport `json:"passport,omitempty"`
}

// A method that can be used to verify the individual's identity.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentification) GetMethod() (Method CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethod) {
	if r != nil && r.Method != nil {
		Method = *r.Method
	}
	return
}

// An identification number that can be used to verify the individual's identity,
// such as a social security number.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentification) GetNumber() (Number string) {
	if r != nil && r.Number != nil {
		Number = *r.Number
	}
	return
}

// Information about the passport used for identification. Required if `method` is
// equal to `passport`.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentification) GetPassport() (Passport CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport) {
	if r != nil && r.Passport != nil {
		Passport = *r.Passport
	}
	return
}

type CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethod string

const (
	CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber                   CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethod = "social_security_number"
	CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethodPassport                               CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethod = "passport"
)

//
type CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport struct {
	// The identifier of the File containing the passport.
	FileID *string `json:"file_id"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `json:"expiration_date"`
	// The country that issued the passport.
	Country *string `json:"country"`
}

// The identifier of the File containing the passport.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

// The passport's expiration date in YYYY-MM-DD format.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport) GetExpirationDate() (ExpirationDate string) {
	if r != nil && r.ExpirationDate != nil {
		ExpirationDate = *r.ExpirationDate
	}
	return
}

// The country that issued the passport.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport) GetCountry() (Country string) {
	if r != nil && r.Country != nil {
		Country = *r.Country
	}
	return
}

//
type CreateAnEntityParametersTrustGrantor struct {
	// The person's legal name.
	Name *string `json:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `json:"date_of_birth"`
	// The individual's address.
	Address *CreateAnEntityParametersTrustGrantorAddress `json:"address"`
	// A means of verifying the person's identity.
	Identification *CreateAnEntityParametersTrustGrantorIdentification `json:"identification"`
}

// The person's legal name.
func (r *CreateAnEntityParametersTrustGrantor) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The person's date of birth in YYYY-MM-DD format.
func (r *CreateAnEntityParametersTrustGrantor) GetDateOfBirth() (DateOfBirth string) {
	if r != nil && r.DateOfBirth != nil {
		DateOfBirth = *r.DateOfBirth
	}
	return
}

// The individual's address.
func (r *CreateAnEntityParametersTrustGrantor) GetAddress() (Address CreateAnEntityParametersTrustGrantorAddress) {
	if r != nil && r.Address != nil {
		Address = *r.Address
	}
	return
}

// A means of verifying the person's identity.
func (r *CreateAnEntityParametersTrustGrantor) GetIdentification() (Identification CreateAnEntityParametersTrustGrantorIdentification) {
	if r != nil && r.Identification != nil {
		Identification = *r.Identification
	}
	return
}

//
type CreateAnEntityParametersTrustGrantorAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 *string `json:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 *string `json:"line2,omitempty"`
	// The city of the address.
	City *string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `json:"state"`
	// The ZIP code of the address.
	Zip *string `json:"zip"`
}

// The first line of the address. This is usually the street number and street.
func (r *CreateAnEntityParametersTrustGrantorAddress) GetLine1() (Line1 string) {
	if r != nil && r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the address. This might be the floor or room number.
func (r *CreateAnEntityParametersTrustGrantorAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the address.
func (r *CreateAnEntityParametersTrustGrantorAddress) GetCity() (City string) {
	if r != nil && r.City != nil {
		City = *r.City
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the state of
// the address.
func (r *CreateAnEntityParametersTrustGrantorAddress) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

// The ZIP code of the address.
func (r *CreateAnEntityParametersTrustGrantorAddress) GetZip() (Zip string) {
	if r != nil && r.Zip != nil {
		Zip = *r.Zip
	}
	return
}

//
type CreateAnEntityParametersTrustGrantorIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method *CreateAnEntityParametersTrustGrantorIdentificationMethod `json:"method"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number *string `json:"number"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport *CreateAnEntityParametersTrustGrantorIdentificationPassport `json:"passport,omitempty"`
}

// A method that can be used to verify the individual's identity.
func (r *CreateAnEntityParametersTrustGrantorIdentification) GetMethod() (Method CreateAnEntityParametersTrustGrantorIdentificationMethod) {
	if r != nil && r.Method != nil {
		Method = *r.Method
	}
	return
}

// An identification number that can be used to verify the individual's identity,
// such as a social security number.
func (r *CreateAnEntityParametersTrustGrantorIdentification) GetNumber() (Number string) {
	if r != nil && r.Number != nil {
		Number = *r.Number
	}
	return
}

// Information about the passport used for identification. Required if `method` is
// equal to `passport`.
func (r *CreateAnEntityParametersTrustGrantorIdentification) GetPassport() (Passport CreateAnEntityParametersTrustGrantorIdentificationPassport) {
	if r != nil && r.Passport != nil {
		Passport = *r.Passport
	}
	return
}

type CreateAnEntityParametersTrustGrantorIdentificationMethod string

const (
	CreateAnEntityParametersTrustGrantorIdentificationMethodSocialSecurityNumber                   CreateAnEntityParametersTrustGrantorIdentificationMethod = "social_security_number"
	CreateAnEntityParametersTrustGrantorIdentificationMethodIndividualTaxpayerIdentificationNumber CreateAnEntityParametersTrustGrantorIdentificationMethod = "individual_taxpayer_identification_number"
	CreateAnEntityParametersTrustGrantorIdentificationMethodPassport                               CreateAnEntityParametersTrustGrantorIdentificationMethod = "passport"
)

//
type CreateAnEntityParametersTrustGrantorIdentificationPassport struct {
	// The identifier of the File containing the passport.
	FileID *string `json:"file_id"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `json:"expiration_date"`
	// The country that issued the passport.
	Country *string `json:"country"`
}

// The identifier of the File containing the passport.
func (r *CreateAnEntityParametersTrustGrantorIdentificationPassport) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

// The passport's expiration date in YYYY-MM-DD format.
func (r *CreateAnEntityParametersTrustGrantorIdentificationPassport) GetExpirationDate() (ExpirationDate string) {
	if r != nil && r.ExpirationDate != nil {
		ExpirationDate = *r.ExpirationDate
	}
	return
}

// The country that issued the passport.
func (r *CreateAnEntityParametersTrustGrantorIdentificationPassport) GetCountry() (Country string) {
	if r != nil && r.Country != nil {
		Country = *r.Country
	}
	return
}

type CreateAnEntityParametersRelationship string

const (
	CreateAnEntityParametersRelationshipAffiliated    CreateAnEntityParametersRelationship = "affiliated"
	CreateAnEntityParametersRelationshipInformational CreateAnEntityParametersRelationship = "informational"
	CreateAnEntityParametersRelationshipUnaffiliated  CreateAnEntityParametersRelationship = "unaffiliated"
)

type CreateAnEntityParametersSupplementalDocuments struct {
	// The identifier of the File containing the document.
	FileID *string `json:"file_id"`
}

// The identifier of the File containing the document.
func (r *CreateAnEntityParametersSupplementalDocuments) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

type ListEntitiesQuery struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int `query:"limit"`
}

// Return the page of entries after this one.
func (r *ListEntitiesQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListEntitiesQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

//
type EntityList struct {
	// The contents of the list.
	Data *[]Entity `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// The contents of the list.
func (r *EntityList) GetData() (Data []Entity) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *EntityList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *EntityService) Create(ctx context.Context, body *CreateAnEntityParameters, opts ...*core.RequestOpts) (res *Entity, err error) {
	err = r.post(
		ctx,
		"/entities",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *PreloadedEntityService) Create(ctx context.Context, body *CreateAnEntityParameters, opts ...*core.RequestOpts) (res *Entity, err error) {
	err = r.Entities.post(
		ctx,
		"/entities",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *EntityService) Retrieve(ctx context.Context, entity_id string, opts ...*core.RequestOpts) (res *Entity, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/entities/%s", entity_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *PreloadedEntityService) Retrieve(ctx context.Context, entity_id string, opts ...*core.RequestOpts) (res *Entity, err error) {
	err = r.Entities.get(
		ctx,
		fmt.Sprintf("/entities/%s", entity_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

type EntitiesPage struct {
	*pagination.Page[Entity]
}

func (r *EntitiesPage) Entity() *Entity {
	return r.Current()
}

func (r *EntitiesPage) GetNextPage() (*EntitiesPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &EntitiesPage{page}, nil
	}
}

func (r *EntityService) List(ctx context.Context, query *ListEntitiesQuery, opts ...*core.RequestOpts) (res *EntitiesPage, err error) {
	page := &EntitiesPage{
		Page: &pagination.Page[Entity]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/entities",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *PreloadedEntityService) List(ctx context.Context, query *ListEntitiesQuery, opts ...*core.RequestOpts) (res *EntitiesPage, err error) {
	page := &EntitiesPage{
		Page: &pagination.Page[Entity]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/entities",
			},
			Requester: r.Entities.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}
