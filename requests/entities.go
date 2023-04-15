package requests

import (
	"net/url"
	"time"

	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
)

type EntityNewParams struct {
	// The type of Entity to create.
	Structure field.Field[EntityNewParamsStructure] `json:"structure,required"`
	// Details of the corporation entity to create. Required if `structure` is equal to
	// `corporation`.
	Corporation field.Field[EntityNewParamsCorporation] `json:"corporation"`
	// Details of the natural person entity to create. Required if `structure` is equal
	// to `natural_person`. Natural people entities should be submitted with
	// `social_security_number` or `individual_taxpayer_identification_number`
	// identification methods.
	NaturalPerson field.Field[EntityNewParamsNaturalPerson] `json:"natural_person"`
	// Details of the joint entity to create. Required if `structure` is equal to
	// `joint`.
	Joint field.Field[EntityNewParamsJoint] `json:"joint"`
	// Details of the trust entity to create. Required if `structure` is equal to
	// `trust`.
	Trust field.Field[EntityNewParamsTrust] `json:"trust"`
	// The description you choose to give the entity.
	Description field.Field[string] `json:"description"`
	// The relationship between your group and the entity.
	Relationship field.Field[EntityNewParamsRelationship] `json:"relationship,required"`
	// Additional documentation associated with the entity.
	SupplementalDocuments field.Field[[]EntityNewParamsSupplementalDocuments] `json:"supplemental_documents"`
}

// MarshalJSON serializes EntityNewParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r EntityNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type EntityNewParamsStructure string

const (
	EntityNewParamsStructureCorporation   EntityNewParamsStructure = "corporation"
	EntityNewParamsStructureNaturalPerson EntityNewParamsStructure = "natural_person"
	EntityNewParamsStructureJoint         EntityNewParamsStructure = "joint"
	EntityNewParamsStructureTrust         EntityNewParamsStructure = "trust"
)

type EntityNewParamsCorporation struct {
	// The legal name of the corporation.
	Name field.Field[string] `json:"name,required"`
	// The website of the corporation.
	Website field.Field[string] `json:"website"`
	// The Employer Identification Number (EIN) for the corporation.
	TaxIdentifier field.Field[string] `json:"tax_identifier,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the
	// corporation's state of incorporation.
	IncorporationState field.Field[string] `json:"incorporation_state"`
	// The corporation's address.
	Address field.Field[EntityNewParamsCorporationAddress] `json:"address,required"`
	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwners field.Field[[]EntityNewParamsCorporationBeneficialOwners] `json:"beneficial_owners,required"`
}

type EntityNewParamsCorporationAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 field.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 field.Field[string] `json:"line2"`
	// The city of the address.
	City field.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State field.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip field.Field[string] `json:"zip,required"`
}

type EntityNewParamsCorporationBeneficialOwners struct {
	// Personal details for the beneficial owner.
	Individual field.Field[EntityNewParamsCorporationBeneficialOwnersIndividual] `json:"individual,required"`
	// This person's role or title within the entity.
	CompanyTitle field.Field[string] `json:"company_title"`
	// Why this person is considered a beneficial owner of the entity.
	Prong field.Field[EntityNewParamsCorporationBeneficialOwnersProng] `json:"prong,required"`
}

type EntityNewParamsCorporationBeneficialOwnersIndividual struct {
	// The person's legal name.
	Name field.Field[string] `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth field.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// The individual's address.
	Address field.Field[EntityNewParamsCorporationBeneficialOwnersIndividualAddress] `json:"address,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID field.Field[bool] `json:"confirmed_no_us_tax_id"`
	// A means of verifying the person's identity.
	Identification field.Field[EntityNewParamsCorporationBeneficialOwnersIndividualIdentification] `json:"identification,required"`
}

type EntityNewParamsCorporationBeneficialOwnersIndividualAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 field.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 field.Field[string] `json:"line2"`
	// The city of the address.
	City field.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State field.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip field.Field[string] `json:"zip,required"`
}

type EntityNewParamsCorporationBeneficialOwnersIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method field.Field[EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod] `json:"method,required"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number field.Field[string] `json:"number,required"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport field.Field[EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationPassport] `json:"passport"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense field.Field[EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other field.Field[EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationOther] `json:"other"`
}

type EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod string

const (
	EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber                   EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod = "social_security_number"
	EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodPassport                               EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod = "passport"
	EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodDriversLicense                         EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod = "drivers_license"
	EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodOther                                  EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod = "other"
)

type EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationPassport struct {
	// The identifier of the File containing the passport.
	FileID field.Field[string] `json:"file_id,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate field.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The country that issued the passport.
	Country field.Field[string] `json:"country,required"`
}

type EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationDriversLicense struct {
	// The identifier of the File containing the driver's license.
	FileID field.Field[string] `json:"file_id,required"`
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate field.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The state that issued the provided driver's license.
	State field.Field[string] `json:"state,required"`
}

type EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document.
	Country field.Field[string] `json:"country,required"`
	// A description of the document submitted.
	Description field.Field[string] `json:"description,required"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate field.Field[time.Time] `json:"expiration_date" format:"date"`
	// The identifier of the File containing the document.
	FileID field.Field[string] `json:"file_id,required"`
}

type EntityNewParamsCorporationBeneficialOwnersProng string

const (
	EntityNewParamsCorporationBeneficialOwnersProngOwnership EntityNewParamsCorporationBeneficialOwnersProng = "ownership"
	EntityNewParamsCorporationBeneficialOwnersProngControl   EntityNewParamsCorporationBeneficialOwnersProng = "control"
)

type EntityNewParamsNaturalPerson struct {
	// The person's legal name.
	Name field.Field[string] `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth field.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// The individual's address.
	Address field.Field[EntityNewParamsNaturalPersonAddress] `json:"address,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID field.Field[bool] `json:"confirmed_no_us_tax_id"`
	// A means of verifying the person's identity.
	Identification field.Field[EntityNewParamsNaturalPersonIdentification] `json:"identification,required"`
}

type EntityNewParamsNaturalPersonAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 field.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 field.Field[string] `json:"line2"`
	// The city of the address.
	City field.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State field.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip field.Field[string] `json:"zip,required"`
}

type EntityNewParamsNaturalPersonIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method field.Field[EntityNewParamsNaturalPersonIdentificationMethod] `json:"method,required"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number field.Field[string] `json:"number,required"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport field.Field[EntityNewParamsNaturalPersonIdentificationPassport] `json:"passport"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense field.Field[EntityNewParamsNaturalPersonIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other field.Field[EntityNewParamsNaturalPersonIdentificationOther] `json:"other"`
}

type EntityNewParamsNaturalPersonIdentificationMethod string

const (
	EntityNewParamsNaturalPersonIdentificationMethodSocialSecurityNumber                   EntityNewParamsNaturalPersonIdentificationMethod = "social_security_number"
	EntityNewParamsNaturalPersonIdentificationMethodIndividualTaxpayerIdentificationNumber EntityNewParamsNaturalPersonIdentificationMethod = "individual_taxpayer_identification_number"
	EntityNewParamsNaturalPersonIdentificationMethodPassport                               EntityNewParamsNaturalPersonIdentificationMethod = "passport"
	EntityNewParamsNaturalPersonIdentificationMethodDriversLicense                         EntityNewParamsNaturalPersonIdentificationMethod = "drivers_license"
	EntityNewParamsNaturalPersonIdentificationMethodOther                                  EntityNewParamsNaturalPersonIdentificationMethod = "other"
)

type EntityNewParamsNaturalPersonIdentificationPassport struct {
	// The identifier of the File containing the passport.
	FileID field.Field[string] `json:"file_id,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate field.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The country that issued the passport.
	Country field.Field[string] `json:"country,required"`
}

type EntityNewParamsNaturalPersonIdentificationDriversLicense struct {
	// The identifier of the File containing the driver's license.
	FileID field.Field[string] `json:"file_id,required"`
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate field.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The state that issued the provided driver's license.
	State field.Field[string] `json:"state,required"`
}

type EntityNewParamsNaturalPersonIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document.
	Country field.Field[string] `json:"country,required"`
	// A description of the document submitted.
	Description field.Field[string] `json:"description,required"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate field.Field[time.Time] `json:"expiration_date" format:"date"`
	// The identifier of the File containing the document.
	FileID field.Field[string] `json:"file_id,required"`
}

type EntityNewParamsJoint struct {
	// The name of the joint entity.
	Name field.Field[string] `json:"name"`
	// The two individuals that share control of the entity.
	Individuals field.Field[[]EntityNewParamsJointIndividuals] `json:"individuals,required"`
}

type EntityNewParamsJointIndividuals struct {
	// The person's legal name.
	Name field.Field[string] `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth field.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// The individual's address.
	Address field.Field[EntityNewParamsJointIndividualsAddress] `json:"address,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID field.Field[bool] `json:"confirmed_no_us_tax_id"`
	// A means of verifying the person's identity.
	Identification field.Field[EntityNewParamsJointIndividualsIdentification] `json:"identification,required"`
}

type EntityNewParamsJointIndividualsAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 field.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 field.Field[string] `json:"line2"`
	// The city of the address.
	City field.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State field.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip field.Field[string] `json:"zip,required"`
}

type EntityNewParamsJointIndividualsIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method field.Field[EntityNewParamsJointIndividualsIdentificationMethod] `json:"method,required"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number field.Field[string] `json:"number,required"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport field.Field[EntityNewParamsJointIndividualsIdentificationPassport] `json:"passport"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense field.Field[EntityNewParamsJointIndividualsIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other field.Field[EntityNewParamsJointIndividualsIdentificationOther] `json:"other"`
}

type EntityNewParamsJointIndividualsIdentificationMethod string

const (
	EntityNewParamsJointIndividualsIdentificationMethodSocialSecurityNumber                   EntityNewParamsJointIndividualsIdentificationMethod = "social_security_number"
	EntityNewParamsJointIndividualsIdentificationMethodIndividualTaxpayerIdentificationNumber EntityNewParamsJointIndividualsIdentificationMethod = "individual_taxpayer_identification_number"
	EntityNewParamsJointIndividualsIdentificationMethodPassport                               EntityNewParamsJointIndividualsIdentificationMethod = "passport"
	EntityNewParamsJointIndividualsIdentificationMethodDriversLicense                         EntityNewParamsJointIndividualsIdentificationMethod = "drivers_license"
	EntityNewParamsJointIndividualsIdentificationMethodOther                                  EntityNewParamsJointIndividualsIdentificationMethod = "other"
)

type EntityNewParamsJointIndividualsIdentificationPassport struct {
	// The identifier of the File containing the passport.
	FileID field.Field[string] `json:"file_id,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate field.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The country that issued the passport.
	Country field.Field[string] `json:"country,required"`
}

type EntityNewParamsJointIndividualsIdentificationDriversLicense struct {
	// The identifier of the File containing the driver's license.
	FileID field.Field[string] `json:"file_id,required"`
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate field.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The state that issued the provided driver's license.
	State field.Field[string] `json:"state,required"`
}

type EntityNewParamsJointIndividualsIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document.
	Country field.Field[string] `json:"country,required"`
	// A description of the document submitted.
	Description field.Field[string] `json:"description,required"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate field.Field[time.Time] `json:"expiration_date" format:"date"`
	// The identifier of the File containing the document.
	FileID field.Field[string] `json:"file_id,required"`
}

type EntityNewParamsTrust struct {
	// The legal name of the trust.
	Name field.Field[string] `json:"name,required"`
	// Whether the trust is `revocable` or `irrevocable`. Irrevocable trusts require
	// their own Employer Identification Number. Revocable trusts require information
	// about the individual `grantor` who created the trust.
	Category field.Field[EntityNewParamsTrustCategory] `json:"category,required"`
	// The Employer Identification Number (EIN) for the trust. Required if `category`
	// is equal to `irrevocable`.
	TaxIdentifier field.Field[string] `json:"tax_identifier"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state in
	// which the trust was formed.
	FormationState field.Field[string] `json:"formation_state"`
	// The trust's address.
	Address field.Field[EntityNewParamsTrustAddress] `json:"address,required"`
	// The identifier of the File containing the formation document of the trust.
	FormationDocumentFileID field.Field[string] `json:"formation_document_file_id"`
	// The trustees of the trust.
	Trustees field.Field[[]EntityNewParamsTrustTrustees] `json:"trustees,required"`
	// The grantor of the trust. Required if `category` is equal to `revocable`.
	Grantor field.Field[EntityNewParamsTrustGrantor] `json:"grantor"`
}

type EntityNewParamsTrustCategory string

const (
	EntityNewParamsTrustCategoryRevocable   EntityNewParamsTrustCategory = "revocable"
	EntityNewParamsTrustCategoryIrrevocable EntityNewParamsTrustCategory = "irrevocable"
)

type EntityNewParamsTrustAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 field.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 field.Field[string] `json:"line2"`
	// The city of the address.
	City field.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State field.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip field.Field[string] `json:"zip,required"`
}

type EntityNewParamsTrustTrustees struct {
	// The structure of the trustee.
	Structure field.Field[EntityNewParamsTrustTrusteesStructure] `json:"structure,required"`
	// Details of the individual trustee. Required when the trustee `structure` is
	// equal to `individual`.
	Individual field.Field[EntityNewParamsTrustTrusteesIndividual] `json:"individual"`
}

type EntityNewParamsTrustTrusteesStructure string

const (
	EntityNewParamsTrustTrusteesStructureIndividual EntityNewParamsTrustTrusteesStructure = "individual"
)

type EntityNewParamsTrustTrusteesIndividual struct {
	// The person's legal name.
	Name field.Field[string] `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth field.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// The individual's address.
	Address field.Field[EntityNewParamsTrustTrusteesIndividualAddress] `json:"address,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID field.Field[bool] `json:"confirmed_no_us_tax_id"`
	// A means of verifying the person's identity.
	Identification field.Field[EntityNewParamsTrustTrusteesIndividualIdentification] `json:"identification,required"`
}

type EntityNewParamsTrustTrusteesIndividualAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 field.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 field.Field[string] `json:"line2"`
	// The city of the address.
	City field.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State field.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip field.Field[string] `json:"zip,required"`
}

type EntityNewParamsTrustTrusteesIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method field.Field[EntityNewParamsTrustTrusteesIndividualIdentificationMethod] `json:"method,required"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number field.Field[string] `json:"number,required"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport field.Field[EntityNewParamsTrustTrusteesIndividualIdentificationPassport] `json:"passport"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense field.Field[EntityNewParamsTrustTrusteesIndividualIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other field.Field[EntityNewParamsTrustTrusteesIndividualIdentificationOther] `json:"other"`
}

type EntityNewParamsTrustTrusteesIndividualIdentificationMethod string

const (
	EntityNewParamsTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber                   EntityNewParamsTrustTrusteesIndividualIdentificationMethod = "social_security_number"
	EntityNewParamsTrustTrusteesIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber EntityNewParamsTrustTrusteesIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	EntityNewParamsTrustTrusteesIndividualIdentificationMethodPassport                               EntityNewParamsTrustTrusteesIndividualIdentificationMethod = "passport"
	EntityNewParamsTrustTrusteesIndividualIdentificationMethodDriversLicense                         EntityNewParamsTrustTrusteesIndividualIdentificationMethod = "drivers_license"
	EntityNewParamsTrustTrusteesIndividualIdentificationMethodOther                                  EntityNewParamsTrustTrusteesIndividualIdentificationMethod = "other"
)

type EntityNewParamsTrustTrusteesIndividualIdentificationPassport struct {
	// The identifier of the File containing the passport.
	FileID field.Field[string] `json:"file_id,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate field.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The country that issued the passport.
	Country field.Field[string] `json:"country,required"`
}

type EntityNewParamsTrustTrusteesIndividualIdentificationDriversLicense struct {
	// The identifier of the File containing the driver's license.
	FileID field.Field[string] `json:"file_id,required"`
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate field.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The state that issued the provided driver's license.
	State field.Field[string] `json:"state,required"`
}

type EntityNewParamsTrustTrusteesIndividualIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document.
	Country field.Field[string] `json:"country,required"`
	// A description of the document submitted.
	Description field.Field[string] `json:"description,required"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate field.Field[time.Time] `json:"expiration_date" format:"date"`
	// The identifier of the File containing the document.
	FileID field.Field[string] `json:"file_id,required"`
}

type EntityNewParamsTrustGrantor struct {
	// The person's legal name.
	Name field.Field[string] `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth field.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// The individual's address.
	Address field.Field[EntityNewParamsTrustGrantorAddress] `json:"address,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID field.Field[bool] `json:"confirmed_no_us_tax_id"`
	// A means of verifying the person's identity.
	Identification field.Field[EntityNewParamsTrustGrantorIdentification] `json:"identification,required"`
}

type EntityNewParamsTrustGrantorAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 field.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 field.Field[string] `json:"line2"`
	// The city of the address.
	City field.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State field.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip field.Field[string] `json:"zip,required"`
}

type EntityNewParamsTrustGrantorIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method field.Field[EntityNewParamsTrustGrantorIdentificationMethod] `json:"method,required"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number field.Field[string] `json:"number,required"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport field.Field[EntityNewParamsTrustGrantorIdentificationPassport] `json:"passport"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense field.Field[EntityNewParamsTrustGrantorIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other field.Field[EntityNewParamsTrustGrantorIdentificationOther] `json:"other"`
}

type EntityNewParamsTrustGrantorIdentificationMethod string

const (
	EntityNewParamsTrustGrantorIdentificationMethodSocialSecurityNumber                   EntityNewParamsTrustGrantorIdentificationMethod = "social_security_number"
	EntityNewParamsTrustGrantorIdentificationMethodIndividualTaxpayerIdentificationNumber EntityNewParamsTrustGrantorIdentificationMethod = "individual_taxpayer_identification_number"
	EntityNewParamsTrustGrantorIdentificationMethodPassport                               EntityNewParamsTrustGrantorIdentificationMethod = "passport"
	EntityNewParamsTrustGrantorIdentificationMethodDriversLicense                         EntityNewParamsTrustGrantorIdentificationMethod = "drivers_license"
	EntityNewParamsTrustGrantorIdentificationMethodOther                                  EntityNewParamsTrustGrantorIdentificationMethod = "other"
)

type EntityNewParamsTrustGrantorIdentificationPassport struct {
	// The identifier of the File containing the passport.
	FileID field.Field[string] `json:"file_id,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate field.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The country that issued the passport.
	Country field.Field[string] `json:"country,required"`
}

type EntityNewParamsTrustGrantorIdentificationDriversLicense struct {
	// The identifier of the File containing the driver's license.
	FileID field.Field[string] `json:"file_id,required"`
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate field.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The state that issued the provided driver's license.
	State field.Field[string] `json:"state,required"`
}

type EntityNewParamsTrustGrantorIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document.
	Country field.Field[string] `json:"country,required"`
	// A description of the document submitted.
	Description field.Field[string] `json:"description,required"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate field.Field[time.Time] `json:"expiration_date" format:"date"`
	// The identifier of the File containing the document.
	FileID field.Field[string] `json:"file_id,required"`
}

type EntityNewParamsRelationship string

const (
	EntityNewParamsRelationshipAffiliated    EntityNewParamsRelationship = "affiliated"
	EntityNewParamsRelationshipInformational EntityNewParamsRelationship = "informational"
	EntityNewParamsRelationshipUnaffiliated  EntityNewParamsRelationship = "unaffiliated"
)

type EntityNewParamsSupplementalDocuments struct {
	// The identifier of the File containing the document.
	FileID field.Field[string] `json:"file_id,required"`
}

type EntityListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
}

// URLQuery serializes EntityListParams into a url.Values of the query parameters
// associated with this value
func (r EntityListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
