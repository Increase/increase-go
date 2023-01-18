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

type CreateASupplementalDocumentForAnEntityParameters struct {
	// The identifier of the File containing the document.
	FileID string `json:"file_id"`
}

//
type Entity struct {
	// The entity's identifier.
	ID string `json:"id"`
	// The entity's legal structure.
	Structure EntityStructure `json:"structure"`
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
	Type EntityType `json:"type"`
	// The entity's description for display purposes.
	Description *string `json:"description"`
	// The relationship between your group and the entity.
	Relationship EntityRelationship `json:"relationship"`
	// Additional documentation associated with the entity.
	SupplementalDocuments []EntitySupplementalDocuments `json:"supplemental_documents"`
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

// The entity's description for display purposes.
func (r *Entity) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
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
	Name string `json:"name"`
	// The website of the corporation.
	Website *string `json:"website"`
	// The Employer Identification Number (EIN) for the corporation.
	TaxIdentifier *string `json:"tax_identifier"`
	// The two-letter United States Postal Service (USPS) abbreviation for the
	// corporation's state of incorporation.
	IncorporationState *string `json:"incorporation_state"`
	// The corporation's address.
	Address EntityCorporationAddress `json:"address"`
	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwners []EntityCorporationBeneficialOwners `json:"beneficial_owners"`
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

//
type EntityCorporationAddress struct {
	// The first line of the address.
	Line1 string `json:"line1"`
	// The second line of the address.
	Line2 *string `json:"line2"`
	// The city of the address.
	City string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state"`
	// The ZIP code of the address.
	Zip string `json:"zip"`
}

// The second line of the address.
func (r *EntityCorporationAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

type EntityCorporationBeneficialOwners struct {
	// Personal details for the beneficial owner.
	Individual EntityCorporationBeneficialOwnersIndividual `json:"individual"`
	// This person's role or title within the entity.
	CompanyTitle *string `json:"company_title"`
	// Why this person is considered a beneficial owner of the entity.
	Prong EntityCorporationBeneficialOwnersProng `json:"prong"`
}

// This person's role or title within the entity.
func (r *EntityCorporationBeneficialOwners) GetCompanyTitle() (CompanyTitle string) {
	if r != nil && r.CompanyTitle != nil {
		CompanyTitle = *r.CompanyTitle
	}
	return
}

//
type EntityCorporationBeneficialOwnersIndividual struct {
	// The person's legal name.
	Name string `json:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth string `json:"date_of_birth"`
	// The person's address.
	Address EntityCorporationBeneficialOwnersIndividualAddress `json:"address"`
	// A means of verifying the person's identity.
	Identification EntityCorporationBeneficialOwnersIndividualIdentification `json:"identification"`
}

//
type EntityCorporationBeneficialOwnersIndividualAddress struct {
	// The first line of the address.
	Line1 string `json:"line1"`
	// The second line of the address.
	Line2 *string `json:"line2"`
	// The city of the address.
	City string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state"`
	// The ZIP code of the address.
	Zip string `json:"zip"`
}

// The second line of the address.
func (r *EntityCorporationBeneficialOwnersIndividualAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

//
type EntityCorporationBeneficialOwnersIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityCorporationBeneficialOwnersIndividualIdentificationMethod `json:"method"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string `json:"number_last4"`
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
	Name string `json:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth string `json:"date_of_birth"`
	// The person's address.
	Address EntityNaturalPersonAddress `json:"address"`
	// A means of verifying the person's identity.
	Identification EntityNaturalPersonIdentification `json:"identification"`
}

//
type EntityNaturalPersonAddress struct {
	// The first line of the address.
	Line1 string `json:"line1"`
	// The second line of the address.
	Line2 *string `json:"line2"`
	// The city of the address.
	City string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state"`
	// The ZIP code of the address.
	Zip string `json:"zip"`
}

// The second line of the address.
func (r *EntityNaturalPersonAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

//
type EntityNaturalPersonIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityNaturalPersonIdentificationMethod `json:"method"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string `json:"number_last4"`
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
	Name string `json:"name"`
	// The two individuals that share control of the entity.
	Individuals []EntityJointIndividuals `json:"individuals"`
}

type EntityJointIndividuals struct {
	// The person's legal name.
	Name string `json:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth string `json:"date_of_birth"`
	// The person's address.
	Address EntityJointIndividualsAddress `json:"address"`
	// A means of verifying the person's identity.
	Identification EntityJointIndividualsIdentification `json:"identification"`
}

//
type EntityJointIndividualsAddress struct {
	// The first line of the address.
	Line1 string `json:"line1"`
	// The second line of the address.
	Line2 *string `json:"line2"`
	// The city of the address.
	City string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state"`
	// The ZIP code of the address.
	Zip string `json:"zip"`
}

// The second line of the address.
func (r *EntityJointIndividualsAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

//
type EntityJointIndividualsIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityJointIndividualsIdentificationMethod `json:"method"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string `json:"number_last4"`
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
	Name string `json:"name"`
	// Whether the trust is `revocable` or `irrevocable`.
	Category EntityTrustCategory `json:"category"`
	// The trust's address.
	Address EntityTrustAddress `json:"address"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state in
	// which the trust was formed.
	FormationState *string `json:"formation_state"`
	// The Employer Identification Number (EIN) of the trust itself.
	TaxIdentifier *string `json:"tax_identifier"`
	// The trustees of the trust.
	Trustees []EntityTrustTrustees `json:"trustees"`
	// The grantor of the trust. Will be present if the `category` is `revocable`.
	Grantor *EntityTrustGrantor `json:"grantor"`
	// The ID for the File containing the formation document of the trust.
	FormationDocumentFileID *string `json:"formation_document_file_id"`
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
	Line1 string `json:"line1"`
	// The second line of the address.
	Line2 *string `json:"line2"`
	// The city of the address.
	City string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state"`
	// The ZIP code of the address.
	Zip string `json:"zip"`
}

// The second line of the address.
func (r *EntityTrustAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

type EntityTrustTrustees struct {
	// The structure of the trustee. Will always be equal to `individual`.
	Structure EntityTrustTrusteesStructure `json:"structure"`
	// The individual trustee of the trust. Will be present if the trustee's
	// `structure` is equal to `individual`.
	Individual *EntityTrustTrusteesIndividual `json:"individual"`
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
	Name string `json:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth string `json:"date_of_birth"`
	// The person's address.
	Address EntityTrustTrusteesIndividualAddress `json:"address"`
	// A means of verifying the person's identity.
	Identification EntityTrustTrusteesIndividualIdentification `json:"identification"`
}

//
type EntityTrustTrusteesIndividualAddress struct {
	// The first line of the address.
	Line1 string `json:"line1"`
	// The second line of the address.
	Line2 *string `json:"line2"`
	// The city of the address.
	City string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state"`
	// The ZIP code of the address.
	Zip string `json:"zip"`
}

// The second line of the address.
func (r *EntityTrustTrusteesIndividualAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

//
type EntityTrustTrusteesIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityTrustTrusteesIndividualIdentificationMethod `json:"method"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string `json:"number_last4"`
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
	Name string `json:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth string `json:"date_of_birth"`
	// The person's address.
	Address EntityTrustGrantorAddress `json:"address"`
	// A means of verifying the person's identity.
	Identification EntityTrustGrantorIdentification `json:"identification"`
}

//
type EntityTrustGrantorAddress struct {
	// The first line of the address.
	Line1 string `json:"line1"`
	// The second line of the address.
	Line2 *string `json:"line2"`
	// The city of the address.
	City string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state"`
	// The ZIP code of the address.
	Zip string `json:"zip"`
}

// The second line of the address.
func (r *EntityTrustGrantorAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

//
type EntityTrustGrantorIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityTrustGrantorIdentificationMethod `json:"method"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string `json:"number_last4"`
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
	FileID string `json:"file_id"`
}

type CreateAnEntityParameters struct {
	// The type of Entity to create.
	Structure CreateAnEntityParametersStructure `json:"structure"`
	// Details of the corporation entity to create. Required if `structure` is equal to
	// `corporation`.
	Corporation CreateAnEntityParametersCorporation `json:"corporation,omitempty"`
	// Details of the natural person entity to create. Required if `structure` is equal
	// to `natural_person`.
	NaturalPerson CreateAnEntityParametersNaturalPerson `json:"natural_person,omitempty"`
	// Details of the joint entity to create. Required if `structure` is equal to
	// `joint`.
	Joint CreateAnEntityParametersJoint `json:"joint,omitempty"`
	// Details of the trust entity to create. Required if `structure` is equal to
	// `trust`.
	Trust CreateAnEntityParametersTrust `json:"trust,omitempty"`
	// The description you choose to give the entity.
	Description string `json:"description,omitempty"`
	// The relationship between your group and the entity.
	Relationship CreateAnEntityParametersRelationship `json:"relationship"`
	// Additional documentation associated with the entity.
	SupplementalDocuments []CreateAnEntityParametersSupplementalDocuments `json:"supplemental_documents,omitempty"`
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
	Name string `json:"name"`
	// The website of the corporation.
	Website string `json:"website,omitempty"`
	// The Employer Identification Number (EIN) for the corporation.
	TaxIdentifier string `json:"tax_identifier"`
	// The two-letter United States Postal Service (USPS) abbreviation for the
	// corporation's state of incorporation.
	IncorporationState string `json:"incorporation_state,omitempty"`
	// The corporation's address.
	Address CreateAnEntityParametersCorporationAddress `json:"address"`
	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwners []CreateAnEntityParametersCorporationBeneficialOwners `json:"beneficial_owners"`
}

//
type CreateAnEntityParametersCorporationAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 string `json:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 string `json:"line2,omitempty"`
	// The city of the address.
	City string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state"`
	// The ZIP code of the address.
	Zip string `json:"zip"`
}

type CreateAnEntityParametersCorporationBeneficialOwners struct {
	// Personal details for the beneficial owner.
	Individual CreateAnEntityParametersCorporationBeneficialOwnersIndividual `json:"individual"`
	// This person's role or title within the entity.
	CompanyTitle string `json:"company_title,omitempty"`
	// Why this person is considered a beneficial owner of the entity.
	Prong CreateAnEntityParametersCorporationBeneficialOwnersProng `json:"prong"`
}

//
type CreateAnEntityParametersCorporationBeneficialOwnersIndividual struct {
	// The person's legal name.
	Name string `json:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth string `json:"date_of_birth"`
	// The individual's address.
	Address CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress `json:"address"`
	// A means of verifying the person's identity.
	Identification CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification `json:"identification"`
}

//
type CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 string `json:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 string `json:"line2,omitempty"`
	// The city of the address.
	City string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state"`
	// The ZIP code of the address.
	Zip string `json:"zip"`
}

//
type CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethod `json:"method"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number string `json:"number"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport `json:"passport,omitempty"`
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
	FileID string `json:"file_id"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate string `json:"expiration_date"`
	// The country that issued the passport.
	Country string `json:"country"`
}

type CreateAnEntityParametersCorporationBeneficialOwnersProng string

const (
	CreateAnEntityParametersCorporationBeneficialOwnersProngOwnership CreateAnEntityParametersCorporationBeneficialOwnersProng = "ownership"
	CreateAnEntityParametersCorporationBeneficialOwnersProngControl   CreateAnEntityParametersCorporationBeneficialOwnersProng = "control"
)

//
type CreateAnEntityParametersNaturalPerson struct {
	// The person's legal name.
	Name string `json:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth string `json:"date_of_birth"`
	// The individual's address.
	Address CreateAnEntityParametersNaturalPersonAddress `json:"address"`
	// A means of verifying the person's identity.
	Identification CreateAnEntityParametersNaturalPersonIdentification `json:"identification"`
}

//
type CreateAnEntityParametersNaturalPersonAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 string `json:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 string `json:"line2,omitempty"`
	// The city of the address.
	City string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state"`
	// The ZIP code of the address.
	Zip string `json:"zip"`
}

//
type CreateAnEntityParametersNaturalPersonIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method CreateAnEntityParametersNaturalPersonIdentificationMethod `json:"method"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number string `json:"number"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport CreateAnEntityParametersNaturalPersonIdentificationPassport `json:"passport,omitempty"`
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
	FileID string `json:"file_id"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate string `json:"expiration_date"`
	// The country that issued the passport.
	Country string `json:"country"`
}

//
type CreateAnEntityParametersJoint struct {
	// The name of the joint entity.
	Name string `json:"name,omitempty"`
	// The two individuals that share control of the entity.
	Individuals []CreateAnEntityParametersJointIndividuals `json:"individuals"`
}

type CreateAnEntityParametersJointIndividuals struct {
	// The person's legal name.
	Name string `json:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth string `json:"date_of_birth"`
	// The individual's address.
	Address CreateAnEntityParametersJointIndividualsAddress `json:"address"`
	// A means of verifying the person's identity.
	Identification CreateAnEntityParametersJointIndividualsIdentification `json:"identification"`
}

//
type CreateAnEntityParametersJointIndividualsAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 string `json:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 string `json:"line2,omitempty"`
	// The city of the address.
	City string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state"`
	// The ZIP code of the address.
	Zip string `json:"zip"`
}

//
type CreateAnEntityParametersJointIndividualsIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method CreateAnEntityParametersJointIndividualsIdentificationMethod `json:"method"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number string `json:"number"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport CreateAnEntityParametersJointIndividualsIdentificationPassport `json:"passport,omitempty"`
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
	FileID string `json:"file_id"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate string `json:"expiration_date"`
	// The country that issued the passport.
	Country string `json:"country"`
}

//
type CreateAnEntityParametersTrust struct {
	// The legal name of the trust.
	Name string `json:"name"`
	// Whether the trust is `revocable` or `irrevocable`. Irrevocable trusts require
	// their own Employer Identification Number. Revocable trusts require information
	// about the individual `grantor` who created the trust.
	Category CreateAnEntityParametersTrustCategory `json:"category"`
	// The Employer Identification Number (EIN) for the trust. Required if `category`
	// is equal to `irrevocable`.
	TaxIdentifier string `json:"tax_identifier,omitempty"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state in
	// which the trust was formed.
	FormationState string `json:"formation_state,omitempty"`
	// The trust's address.
	Address CreateAnEntityParametersTrustAddress `json:"address"`
	// The identifier of the File containing the formation document of the trust.
	FormationDocumentFileID string `json:"formation_document_file_id,omitempty"`
	// The trustees of the trust.
	Trustees []CreateAnEntityParametersTrustTrustees `json:"trustees"`
	// The grantor of the trust. Required if `category` is equal to `revocable`.
	Grantor CreateAnEntityParametersTrustGrantor `json:"grantor,omitempty"`
}

type CreateAnEntityParametersTrustCategory string

const (
	CreateAnEntityParametersTrustCategoryRevocable   CreateAnEntityParametersTrustCategory = "revocable"
	CreateAnEntityParametersTrustCategoryIrrevocable CreateAnEntityParametersTrustCategory = "irrevocable"
)

//
type CreateAnEntityParametersTrustAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 string `json:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 string `json:"line2,omitempty"`
	// The city of the address.
	City string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state"`
	// The ZIP code of the address.
	Zip string `json:"zip"`
}

type CreateAnEntityParametersTrustTrustees struct {
	// The structure of the trustee.
	Structure CreateAnEntityParametersTrustTrusteesStructure `json:"structure"`
	// Details of the individual trustee. Required when the trustee `structure` is
	// equal to `individual`.
	Individual CreateAnEntityParametersTrustTrusteesIndividual `json:"individual,omitempty"`
}

type CreateAnEntityParametersTrustTrusteesStructure string

const (
	CreateAnEntityParametersTrustTrusteesStructureIndividual CreateAnEntityParametersTrustTrusteesStructure = "individual"
)

//
type CreateAnEntityParametersTrustTrusteesIndividual struct {
	// The person's legal name.
	Name string `json:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth string `json:"date_of_birth"`
	// The individual's address.
	Address CreateAnEntityParametersTrustTrusteesIndividualAddress `json:"address"`
	// A means of verifying the person's identity.
	Identification CreateAnEntityParametersTrustTrusteesIndividualIdentification `json:"identification"`
}

//
type CreateAnEntityParametersTrustTrusteesIndividualAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 string `json:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 string `json:"line2,omitempty"`
	// The city of the address.
	City string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state"`
	// The ZIP code of the address.
	Zip string `json:"zip"`
}

//
type CreateAnEntityParametersTrustTrusteesIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethod `json:"method"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number string `json:"number"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport `json:"passport,omitempty"`
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
	FileID string `json:"file_id"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate string `json:"expiration_date"`
	// The country that issued the passport.
	Country string `json:"country"`
}

//
type CreateAnEntityParametersTrustGrantor struct {
	// The person's legal name.
	Name string `json:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth string `json:"date_of_birth"`
	// The individual's address.
	Address CreateAnEntityParametersTrustGrantorAddress `json:"address"`
	// A means of verifying the person's identity.
	Identification CreateAnEntityParametersTrustGrantorIdentification `json:"identification"`
}

//
type CreateAnEntityParametersTrustGrantorAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 string `json:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 string `json:"line2,omitempty"`
	// The city of the address.
	City string `json:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state"`
	// The ZIP code of the address.
	Zip string `json:"zip"`
}

//
type CreateAnEntityParametersTrustGrantorIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method CreateAnEntityParametersTrustGrantorIdentificationMethod `json:"method"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number string `json:"number"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport CreateAnEntityParametersTrustGrantorIdentificationPassport `json:"passport,omitempty"`
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
	FileID string `json:"file_id"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate string `json:"expiration_date"`
	// The country that issued the passport.
	Country string `json:"country"`
}

type CreateAnEntityParametersRelationship string

const (
	CreateAnEntityParametersRelationshipAffiliated    CreateAnEntityParametersRelationship = "affiliated"
	CreateAnEntityParametersRelationshipInformational CreateAnEntityParametersRelationship = "informational"
	CreateAnEntityParametersRelationshipUnaffiliated  CreateAnEntityParametersRelationship = "unaffiliated"
)

type CreateAnEntityParametersSupplementalDocuments struct {
	// The identifier of the File containing the document.
	FileID string `json:"file_id"`
}

type ListEntitiesQuery struct {
	// Return the page of entries after this one.
	Cursor string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit int `query:"limit"`
}

//
type EntityList struct {
	// The contents of the list.
	Data []Entity `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
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
