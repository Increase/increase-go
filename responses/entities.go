package responses

import (
	"time"

	apijson "github.com/increase/increase-go/core/json"
)

type Entity struct {
	// The entity's identifier.
	ID string `json:"id,required"`
	// The entity's legal structure.
	Structure EntityStructure `json:"structure,required"`
	// Details of the corporation entity. Will be present if `structure` is equal to
	// `corporation`.
	Corporation EntityCorporation `json:"corporation,required,nullable"`
	// Details of the natural person entity. Will be present if `structure` is equal to
	// `natural_person`.
	NaturalPerson EntityNaturalPerson `json:"natural_person,required,nullable"`
	// Details of the joint entity. Will be present if `structure` is equal to `joint`.
	Joint EntityJoint `json:"joint,required,nullable"`
	// Details of the trust entity. Will be present if `structure` is equal to `trust`.
	Trust EntityTrust `json:"trust,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `entity`.
	Type EntityType `json:"type,required"`
	// The entity's description for display purposes.
	Description string `json:"description,required,nullable"`
	// The relationship between your group and the entity.
	Relationship EntityRelationship `json:"relationship,required"`
	// Additional documentation associated with the entity.
	SupplementalDocuments []EntitySupplementalDocuments `json:"supplemental_documents,required"`
	JSON                  EntityJSON
}

type EntityJSON struct {
	ID                    apijson.Metadata
	Structure             apijson.Metadata
	Corporation           apijson.Metadata
	NaturalPerson         apijson.Metadata
	Joint                 apijson.Metadata
	Trust                 apijson.Metadata
	Type                  apijson.Metadata
	Description           apijson.Metadata
	Relationship          apijson.Metadata
	SupplementalDocuments apijson.Metadata
	Raw                   []byte
	Extras                map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Entity using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Entity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityStructure string

const (
	EntityStructureCorporation   EntityStructure = "corporation"
	EntityStructureNaturalPerson EntityStructure = "natural_person"
	EntityStructureJoint         EntityStructure = "joint"
	EntityStructureTrust         EntityStructure = "trust"
)

type EntityCorporation struct {
	// The legal name of the corporation.
	Name string `json:"name,required"`
	// The website of the corporation.
	Website string `json:"website,required,nullable"`
	// The Employer Identification Number (EIN) for the corporation.
	TaxIdentifier string `json:"tax_identifier,required,nullable"`
	// The two-letter United States Postal Service (USPS) abbreviation for the
	// corporation's state of incorporation.
	IncorporationState string `json:"incorporation_state,required,nullable"`
	// The corporation's address.
	Address EntityCorporationAddress `json:"address,required"`
	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwners []EntityCorporationBeneficialOwners `json:"beneficial_owners,required"`
	JSON             EntityCorporationJSON
}

type EntityCorporationJSON struct {
	Name               apijson.Metadata
	Website            apijson.Metadata
	TaxIdentifier      apijson.Metadata
	IncorporationState apijson.Metadata
	Address            apijson.Metadata
	BeneficialOwners   apijson.Metadata
	Raw                []byte
	Extras             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EntityCorporation using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EntityCorporation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityCorporationAddress struct {
	// The first line of the address.
	Line1 string `json:"line1,required"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The city of the address.
	City string `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state,required"`
	// The ZIP code of the address.
	Zip  string `json:"zip,required"`
	JSON EntityCorporationAddressJSON
}

type EntityCorporationAddressJSON struct {
	Line1  apijson.Metadata
	Line2  apijson.Metadata
	City   apijson.Metadata
	State  apijson.Metadata
	Zip    apijson.Metadata
	Raw    []byte
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EntityCorporationAddress
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *EntityCorporationAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityCorporationBeneficialOwners struct {
	// Personal details for the beneficial owner.
	Individual EntityCorporationBeneficialOwnersIndividual `json:"individual,required"`
	// This person's role or title within the entity.
	CompanyTitle string `json:"company_title,required,nullable"`
	// Why this person is considered a beneficial owner of the entity.
	Prong EntityCorporationBeneficialOwnersProng `json:"prong,required"`
	JSON  EntityCorporationBeneficialOwnersJSON
}

type EntityCorporationBeneficialOwnersJSON struct {
	Individual   apijson.Metadata
	CompanyTitle apijson.Metadata
	Prong        apijson.Metadata
	Raw          []byte
	Extras       map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// EntityCorporationBeneficialOwners using the internal json library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *EntityCorporationBeneficialOwners) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityCorporationBeneficialOwnersIndividual struct {
	// The person's legal name.
	Name string `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth time.Time `json:"date_of_birth,required" format:"date"`
	// The person's address.
	Address EntityCorporationBeneficialOwnersIndividualAddress `json:"address,required"`
	// A means of verifying the person's identity.
	Identification EntityCorporationBeneficialOwnersIndividualIdentification `json:"identification,required"`
	JSON           EntityCorporationBeneficialOwnersIndividualJSON
}

type EntityCorporationBeneficialOwnersIndividualJSON struct {
	Name           apijson.Metadata
	DateOfBirth    apijson.Metadata
	Address        apijson.Metadata
	Identification apijson.Metadata
	Raw            []byte
	Extras         map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// EntityCorporationBeneficialOwnersIndividual using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *EntityCorporationBeneficialOwnersIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityCorporationBeneficialOwnersIndividualAddress struct {
	// The first line of the address.
	Line1 string `json:"line1,required"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The city of the address.
	City string `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state,required"`
	// The ZIP code of the address.
	Zip  string `json:"zip,required"`
	JSON EntityCorporationBeneficialOwnersIndividualAddressJSON
}

type EntityCorporationBeneficialOwnersIndividualAddressJSON struct {
	Line1  apijson.Metadata
	Line2  apijson.Metadata
	City   apijson.Metadata
	State  apijson.Metadata
	Zip    apijson.Metadata
	Raw    []byte
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// EntityCorporationBeneficialOwnersIndividualAddress using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *EntityCorporationBeneficialOwnersIndividualAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityCorporationBeneficialOwnersIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityCorporationBeneficialOwnersIndividualIdentificationMethod `json:"method,required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string `json:"number_last4,required"`
	JSON        EntityCorporationBeneficialOwnersIndividualIdentificationJSON
}

type EntityCorporationBeneficialOwnersIndividualIdentificationJSON struct {
	Method      apijson.Metadata
	NumberLast4 apijson.Metadata
	Raw         []byte
	Extras      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// EntityCorporationBeneficialOwnersIndividualIdentification using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *EntityCorporationBeneficialOwnersIndividualIdentification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityCorporationBeneficialOwnersIndividualIdentificationMethod string

const (
	EntityCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber                   EntityCorporationBeneficialOwnersIndividualIdentificationMethod = "social_security_number"
	EntityCorporationBeneficialOwnersIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber EntityCorporationBeneficialOwnersIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	EntityCorporationBeneficialOwnersIndividualIdentificationMethodPassport                               EntityCorporationBeneficialOwnersIndividualIdentificationMethod = "passport"
	EntityCorporationBeneficialOwnersIndividualIdentificationMethodDriversLicense                         EntityCorporationBeneficialOwnersIndividualIdentificationMethod = "drivers_license"
	EntityCorporationBeneficialOwnersIndividualIdentificationMethodOther                                  EntityCorporationBeneficialOwnersIndividualIdentificationMethod = "other"
)

type EntityCorporationBeneficialOwnersProng string

const (
	EntityCorporationBeneficialOwnersProngOwnership EntityCorporationBeneficialOwnersProng = "ownership"
	EntityCorporationBeneficialOwnersProngControl   EntityCorporationBeneficialOwnersProng = "control"
)

type EntityNaturalPerson struct {
	// The person's legal name.
	Name string `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth time.Time `json:"date_of_birth,required" format:"date"`
	// The person's address.
	Address EntityNaturalPersonAddress `json:"address,required"`
	// A means of verifying the person's identity.
	Identification EntityNaturalPersonIdentification `json:"identification,required"`
	JSON           EntityNaturalPersonJSON
}

type EntityNaturalPersonJSON struct {
	Name           apijson.Metadata
	DateOfBirth    apijson.Metadata
	Address        apijson.Metadata
	Identification apijson.Metadata
	Raw            []byte
	Extras         map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EntityNaturalPerson using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EntityNaturalPerson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityNaturalPersonAddress struct {
	// The first line of the address.
	Line1 string `json:"line1,required"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The city of the address.
	City string `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state,required"`
	// The ZIP code of the address.
	Zip  string `json:"zip,required"`
	JSON EntityNaturalPersonAddressJSON
}

type EntityNaturalPersonAddressJSON struct {
	Line1  apijson.Metadata
	Line2  apijson.Metadata
	City   apijson.Metadata
	State  apijson.Metadata
	Zip    apijson.Metadata
	Raw    []byte
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EntityNaturalPersonAddress
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *EntityNaturalPersonAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityNaturalPersonIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityNaturalPersonIdentificationMethod `json:"method,required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string `json:"number_last4,required"`
	JSON        EntityNaturalPersonIdentificationJSON
}

type EntityNaturalPersonIdentificationJSON struct {
	Method      apijson.Metadata
	NumberLast4 apijson.Metadata
	Raw         []byte
	Extras      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// EntityNaturalPersonIdentification using the internal json library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *EntityNaturalPersonIdentification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityNaturalPersonIdentificationMethod string

const (
	EntityNaturalPersonIdentificationMethodSocialSecurityNumber                   EntityNaturalPersonIdentificationMethod = "social_security_number"
	EntityNaturalPersonIdentificationMethodIndividualTaxpayerIdentificationNumber EntityNaturalPersonIdentificationMethod = "individual_taxpayer_identification_number"
	EntityNaturalPersonIdentificationMethodPassport                               EntityNaturalPersonIdentificationMethod = "passport"
	EntityNaturalPersonIdentificationMethodDriversLicense                         EntityNaturalPersonIdentificationMethod = "drivers_license"
	EntityNaturalPersonIdentificationMethodOther                                  EntityNaturalPersonIdentificationMethod = "other"
)

type EntityJoint struct {
	// The entity's name.
	Name string `json:"name,required"`
	// The two individuals that share control of the entity.
	Individuals []EntityJointIndividuals `json:"individuals,required"`
	JSON        EntityJointJSON
}

type EntityJointJSON struct {
	Name        apijson.Metadata
	Individuals apijson.Metadata
	Raw         []byte
	Extras      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EntityJoint using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EntityJoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityJointIndividuals struct {
	// The person's legal name.
	Name string `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth time.Time `json:"date_of_birth,required" format:"date"`
	// The person's address.
	Address EntityJointIndividualsAddress `json:"address,required"`
	// A means of verifying the person's identity.
	Identification EntityJointIndividualsIdentification `json:"identification,required"`
	JSON           EntityJointIndividualsJSON
}

type EntityJointIndividualsJSON struct {
	Name           apijson.Metadata
	DateOfBirth    apijson.Metadata
	Address        apijson.Metadata
	Identification apijson.Metadata
	Raw            []byte
	Extras         map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EntityJointIndividuals using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EntityJointIndividuals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityJointIndividualsAddress struct {
	// The first line of the address.
	Line1 string `json:"line1,required"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The city of the address.
	City string `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state,required"`
	// The ZIP code of the address.
	Zip  string `json:"zip,required"`
	JSON EntityJointIndividualsAddressJSON
}

type EntityJointIndividualsAddressJSON struct {
	Line1  apijson.Metadata
	Line2  apijson.Metadata
	City   apijson.Metadata
	State  apijson.Metadata
	Zip    apijson.Metadata
	Raw    []byte
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EntityJointIndividualsAddress
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *EntityJointIndividualsAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityJointIndividualsIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityJointIndividualsIdentificationMethod `json:"method,required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string `json:"number_last4,required"`
	JSON        EntityJointIndividualsIdentificationJSON
}

type EntityJointIndividualsIdentificationJSON struct {
	Method      apijson.Metadata
	NumberLast4 apijson.Metadata
	Raw         []byte
	Extras      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// EntityJointIndividualsIdentification using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *EntityJointIndividualsIdentification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityJointIndividualsIdentificationMethod string

const (
	EntityJointIndividualsIdentificationMethodSocialSecurityNumber                   EntityJointIndividualsIdentificationMethod = "social_security_number"
	EntityJointIndividualsIdentificationMethodIndividualTaxpayerIdentificationNumber EntityJointIndividualsIdentificationMethod = "individual_taxpayer_identification_number"
	EntityJointIndividualsIdentificationMethodPassport                               EntityJointIndividualsIdentificationMethod = "passport"
	EntityJointIndividualsIdentificationMethodDriversLicense                         EntityJointIndividualsIdentificationMethod = "drivers_license"
	EntityJointIndividualsIdentificationMethodOther                                  EntityJointIndividualsIdentificationMethod = "other"
)

type EntityTrust struct {
	// The trust's name
	Name string `json:"name,required"`
	// Whether the trust is `revocable` or `irrevocable`.
	Category EntityTrustCategory `json:"category,required"`
	// The trust's address.
	Address EntityTrustAddress `json:"address,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state in
	// which the trust was formed.
	FormationState string `json:"formation_state,required,nullable"`
	// The Employer Identification Number (EIN) of the trust itself.
	TaxIdentifier string `json:"tax_identifier,required,nullable"`
	// The trustees of the trust.
	Trustees []EntityTrustTrustees `json:"trustees,required"`
	// The grantor of the trust. Will be present if the `category` is `revocable`.
	Grantor EntityTrustGrantor `json:"grantor,required,nullable"`
	// The ID for the File containing the formation document of the trust.
	FormationDocumentFileID string `json:"formation_document_file_id,required,nullable"`
	JSON                    EntityTrustJSON
}

type EntityTrustJSON struct {
	Name                    apijson.Metadata
	Category                apijson.Metadata
	Address                 apijson.Metadata
	FormationState          apijson.Metadata
	TaxIdentifier           apijson.Metadata
	Trustees                apijson.Metadata
	Grantor                 apijson.Metadata
	FormationDocumentFileID apijson.Metadata
	Raw                     []byte
	Extras                  map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EntityTrust using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EntityTrust) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityTrustCategory string

const (
	EntityTrustCategoryRevocable   EntityTrustCategory = "revocable"
	EntityTrustCategoryIrrevocable EntityTrustCategory = "irrevocable"
)

type EntityTrustAddress struct {
	// The first line of the address.
	Line1 string `json:"line1,required"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The city of the address.
	City string `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state,required"`
	// The ZIP code of the address.
	Zip  string `json:"zip,required"`
	JSON EntityTrustAddressJSON
}

type EntityTrustAddressJSON struct {
	Line1  apijson.Metadata
	Line2  apijson.Metadata
	City   apijson.Metadata
	State  apijson.Metadata
	Zip    apijson.Metadata
	Raw    []byte
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EntityTrustAddress using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EntityTrustAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityTrustTrustees struct {
	// The structure of the trustee. Will always be equal to `individual`.
	Structure EntityTrustTrusteesStructure `json:"structure,required"`
	// The individual trustee of the trust. Will be present if the trustee's
	// `structure` is equal to `individual`.
	Individual EntityTrustTrusteesIndividual `json:"individual,required,nullable"`
	JSON       EntityTrustTrusteesJSON
}

type EntityTrustTrusteesJSON struct {
	Structure  apijson.Metadata
	Individual apijson.Metadata
	Raw        []byte
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EntityTrustTrustees using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EntityTrustTrustees) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityTrustTrusteesStructure string

const (
	EntityTrustTrusteesStructureIndividual EntityTrustTrusteesStructure = "individual"
)

type EntityTrustTrusteesIndividual struct {
	// The person's legal name.
	Name string `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth time.Time `json:"date_of_birth,required" format:"date"`
	// The person's address.
	Address EntityTrustTrusteesIndividualAddress `json:"address,required"`
	// A means of verifying the person's identity.
	Identification EntityTrustTrusteesIndividualIdentification `json:"identification,required"`
	JSON           EntityTrustTrusteesIndividualJSON
}

type EntityTrustTrusteesIndividualJSON struct {
	Name           apijson.Metadata
	DateOfBirth    apijson.Metadata
	Address        apijson.Metadata
	Identification apijson.Metadata
	Raw            []byte
	Extras         map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EntityTrustTrusteesIndividual
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *EntityTrustTrusteesIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityTrustTrusteesIndividualAddress struct {
	// The first line of the address.
	Line1 string `json:"line1,required"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The city of the address.
	City string `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state,required"`
	// The ZIP code of the address.
	Zip  string `json:"zip,required"`
	JSON EntityTrustTrusteesIndividualAddressJSON
}

type EntityTrustTrusteesIndividualAddressJSON struct {
	Line1  apijson.Metadata
	Line2  apijson.Metadata
	City   apijson.Metadata
	State  apijson.Metadata
	Zip    apijson.Metadata
	Raw    []byte
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// EntityTrustTrusteesIndividualAddress using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *EntityTrustTrusteesIndividualAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityTrustTrusteesIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityTrustTrusteesIndividualIdentificationMethod `json:"method,required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string `json:"number_last4,required"`
	JSON        EntityTrustTrusteesIndividualIdentificationJSON
}

type EntityTrustTrusteesIndividualIdentificationJSON struct {
	Method      apijson.Metadata
	NumberLast4 apijson.Metadata
	Raw         []byte
	Extras      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// EntityTrustTrusteesIndividualIdentification using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *EntityTrustTrusteesIndividualIdentification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityTrustTrusteesIndividualIdentificationMethod string

const (
	EntityTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber                   EntityTrustTrusteesIndividualIdentificationMethod = "social_security_number"
	EntityTrustTrusteesIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber EntityTrustTrusteesIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	EntityTrustTrusteesIndividualIdentificationMethodPassport                               EntityTrustTrusteesIndividualIdentificationMethod = "passport"
	EntityTrustTrusteesIndividualIdentificationMethodDriversLicense                         EntityTrustTrusteesIndividualIdentificationMethod = "drivers_license"
	EntityTrustTrusteesIndividualIdentificationMethodOther                                  EntityTrustTrusteesIndividualIdentificationMethod = "other"
)

type EntityTrustGrantor struct {
	// The person's legal name.
	Name string `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth time.Time `json:"date_of_birth,required" format:"date"`
	// The person's address.
	Address EntityTrustGrantorAddress `json:"address,required"`
	// A means of verifying the person's identity.
	Identification EntityTrustGrantorIdentification `json:"identification,required"`
	JSON           EntityTrustGrantorJSON
}

type EntityTrustGrantorJSON struct {
	Name           apijson.Metadata
	DateOfBirth    apijson.Metadata
	Address        apijson.Metadata
	Identification apijson.Metadata
	Raw            []byte
	Extras         map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EntityTrustGrantor using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EntityTrustGrantor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityTrustGrantorAddress struct {
	// The first line of the address.
	Line1 string `json:"line1,required"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The city of the address.
	City string `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state,required"`
	// The ZIP code of the address.
	Zip  string `json:"zip,required"`
	JSON EntityTrustGrantorAddressJSON
}

type EntityTrustGrantorAddressJSON struct {
	Line1  apijson.Metadata
	Line2  apijson.Metadata
	City   apijson.Metadata
	State  apijson.Metadata
	Zip    apijson.Metadata
	Raw    []byte
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EntityTrustGrantorAddress
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *EntityTrustGrantorAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityTrustGrantorIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityTrustGrantorIdentificationMethod `json:"method,required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string `json:"number_last4,required"`
	JSON        EntityTrustGrantorIdentificationJSON
}

type EntityTrustGrantorIdentificationJSON struct {
	Method      apijson.Metadata
	NumberLast4 apijson.Metadata
	Raw         []byte
	Extras      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// EntityTrustGrantorIdentification using the internal json library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *EntityTrustGrantorIdentification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityTrustGrantorIdentificationMethod string

const (
	EntityTrustGrantorIdentificationMethodSocialSecurityNumber                   EntityTrustGrantorIdentificationMethod = "social_security_number"
	EntityTrustGrantorIdentificationMethodIndividualTaxpayerIdentificationNumber EntityTrustGrantorIdentificationMethod = "individual_taxpayer_identification_number"
	EntityTrustGrantorIdentificationMethodPassport                               EntityTrustGrantorIdentificationMethod = "passport"
	EntityTrustGrantorIdentificationMethodDriversLicense                         EntityTrustGrantorIdentificationMethod = "drivers_license"
	EntityTrustGrantorIdentificationMethodOther                                  EntityTrustGrantorIdentificationMethod = "other"
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
	FileID string `json:"file_id,required"`
	JSON   EntitySupplementalDocumentsJSON
}

type EntitySupplementalDocumentsJSON struct {
	FileID apijson.Metadata
	Raw    []byte
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EntitySupplementalDocuments
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *EntitySupplementalDocuments) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityListResponse struct {
	// The contents of the list.
	Data []Entity `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       EntityListResponseJSON
}

type EntityListResponseJSON struct {
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	Raw        []byte
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EntityListResponse using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EntityListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
