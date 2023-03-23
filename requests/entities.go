package requests

import (
	"fmt"
	"net/url"
	"time"

	"github.com/increase/increase-go/core"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type Entity struct {
	// The entity's identifier.
	ID fields.Field[string] `json:"id,required"`
	// The entity's legal structure.
	Structure fields.Field[EntityStructure] `json:"structure,required"`
	// Details of the corporation entity. Will be present if `structure` is equal to
	// `corporation`.
	Corporation fields.Field[EntityCorporation] `json:"corporation,required,nullable"`
	// Details of the natural person entity. Will be present if `structure` is equal to
	// `natural_person`.
	NaturalPerson fields.Field[EntityNaturalPerson] `json:"natural_person,required,nullable"`
	// Details of the joint entity. Will be present if `structure` is equal to `joint`.
	Joint fields.Field[EntityJoint] `json:"joint,required,nullable"`
	// Details of the trust entity. Will be present if `structure` is equal to `trust`.
	Trust fields.Field[EntityTrust] `json:"trust,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `entity`.
	Type fields.Field[EntityType] `json:"type,required"`
	// The entity's description for display purposes.
	Description fields.Field[string] `json:"description,required,nullable"`
	// The relationship between your group and the entity.
	Relationship fields.Field[EntityRelationship] `json:"relationship,required"`
	// Additional documentation associated with the entity.
	SupplementalDocuments fields.Field[[]EntitySupplementalDocuments] `json:"supplemental_documents,required"`
}

// MarshalJSON serializes Entity into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Entity) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Entity) String() (result string) {
	return fmt.Sprintf("&Entity{ID:%s Structure:%s Corporation:%s NaturalPerson:%s Joint:%s Trust:%s Type:%s Description:%s Relationship:%s SupplementalDocuments:%s}", r.ID, r.Structure, r.Corporation, r.NaturalPerson, r.Joint, r.Trust, r.Type, r.Description, r.Relationship, core.Fmt(r.SupplementalDocuments))
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
	Name fields.Field[string] `json:"name,required"`
	// The website of the corporation.
	Website fields.Field[string] `json:"website,required,nullable"`
	// The Employer Identification Number (EIN) for the corporation.
	TaxIdentifier fields.Field[string] `json:"tax_identifier,required,nullable"`
	// The two-letter United States Postal Service (USPS) abbreviation for the
	// corporation's state of incorporation.
	IncorporationState fields.Field[string] `json:"incorporation_state,required,nullable"`
	// The corporation's address.
	Address fields.Field[EntityCorporationAddress] `json:"address,required"`
	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwners fields.Field[[]EntityCorporationBeneficialOwners] `json:"beneficial_owners,required"`
}

// MarshalJSON serializes EntityCorporation into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *EntityCorporation) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityCorporation) String() (result string) {
	return fmt.Sprintf("&EntityCorporation{Name:%s Website:%s TaxIdentifier:%s IncorporationState:%s Address:%s BeneficialOwners:%s}", r.Name, r.Website, r.TaxIdentifier, r.IncorporationState, r.Address, core.Fmt(r.BeneficialOwners))
}

type EntityCorporationAddress struct {
	// The first line of the address.
	Line1 fields.Field[string] `json:"line1,required"`
	// The second line of the address.
	Line2 fields.Field[string] `json:"line2,required,nullable"`
	// The city of the address.
	City fields.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State fields.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip fields.Field[string] `json:"zip,required"`
}

// MarshalJSON serializes EntityCorporationAddress into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EntityCorporationAddress) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityCorporationAddress) String() (result string) {
	return fmt.Sprintf("&EntityCorporationAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", r.Line1, r.Line2, r.City, r.State, r.Zip)
}

type EntityCorporationBeneficialOwners struct {
	// Personal details for the beneficial owner.
	Individual fields.Field[EntityCorporationBeneficialOwnersIndividual] `json:"individual,required"`
	// This person's role or title within the entity.
	CompanyTitle fields.Field[string] `json:"company_title,required,nullable"`
	// Why this person is considered a beneficial owner of the entity.
	Prong fields.Field[EntityCorporationBeneficialOwnersProng] `json:"prong,required"`
}

// MarshalJSON serializes EntityCorporationBeneficialOwners into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *EntityCorporationBeneficialOwners) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityCorporationBeneficialOwners) String() (result string) {
	return fmt.Sprintf("&EntityCorporationBeneficialOwners{Individual:%s CompanyTitle:%s Prong:%s}", r.Individual, r.CompanyTitle, r.Prong)
}

type EntityCorporationBeneficialOwnersIndividual struct {
	// The person's legal name.
	Name fields.Field[string] `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth fields.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// The person's address.
	Address fields.Field[EntityCorporationBeneficialOwnersIndividualAddress] `json:"address,required"`
	// A means of verifying the person's identity.
	Identification fields.Field[EntityCorporationBeneficialOwnersIndividualIdentification] `json:"identification,required"`
}

// MarshalJSON serializes EntityCorporationBeneficialOwnersIndividual into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *EntityCorporationBeneficialOwnersIndividual) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityCorporationBeneficialOwnersIndividual) String() (result string) {
	return fmt.Sprintf("&EntityCorporationBeneficialOwnersIndividual{Name:%s DateOfBirth:%s Address:%s Identification:%s}", r.Name, r.DateOfBirth, r.Address, r.Identification)
}

type EntityCorporationBeneficialOwnersIndividualAddress struct {
	// The first line of the address.
	Line1 fields.Field[string] `json:"line1,required"`
	// The second line of the address.
	Line2 fields.Field[string] `json:"line2,required,nullable"`
	// The city of the address.
	City fields.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State fields.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip fields.Field[string] `json:"zip,required"`
}

// MarshalJSON serializes EntityCorporationBeneficialOwnersIndividualAddress into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *EntityCorporationBeneficialOwnersIndividualAddress) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityCorporationBeneficialOwnersIndividualAddress) String() (result string) {
	return fmt.Sprintf("&EntityCorporationBeneficialOwnersIndividualAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", r.Line1, r.Line2, r.City, r.State, r.Zip)
}

type EntityCorporationBeneficialOwnersIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method fields.Field[EntityCorporationBeneficialOwnersIndividualIdentificationMethod] `json:"method,required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 fields.Field[string] `json:"number_last4,required"`
}

// MarshalJSON serializes EntityCorporationBeneficialOwnersIndividualIdentification
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *EntityCorporationBeneficialOwnersIndividualIdentification) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityCorporationBeneficialOwnersIndividualIdentification) String() (result string) {
	return fmt.Sprintf("&EntityCorporationBeneficialOwnersIndividualIdentification{Method:%s NumberLast4:%s}", r.Method, r.NumberLast4)
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
	Name fields.Field[string] `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth fields.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// The person's address.
	Address fields.Field[EntityNaturalPersonAddress] `json:"address,required"`
	// A means of verifying the person's identity.
	Identification fields.Field[EntityNaturalPersonIdentification] `json:"identification,required"`
}

// MarshalJSON serializes EntityNaturalPerson into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EntityNaturalPerson) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityNaturalPerson) String() (result string) {
	return fmt.Sprintf("&EntityNaturalPerson{Name:%s DateOfBirth:%s Address:%s Identification:%s}", r.Name, r.DateOfBirth, r.Address, r.Identification)
}

type EntityNaturalPersonAddress struct {
	// The first line of the address.
	Line1 fields.Field[string] `json:"line1,required"`
	// The second line of the address.
	Line2 fields.Field[string] `json:"line2,required,nullable"`
	// The city of the address.
	City fields.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State fields.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip fields.Field[string] `json:"zip,required"`
}

// MarshalJSON serializes EntityNaturalPersonAddress into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EntityNaturalPersonAddress) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityNaturalPersonAddress) String() (result string) {
	return fmt.Sprintf("&EntityNaturalPersonAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", r.Line1, r.Line2, r.City, r.State, r.Zip)
}

type EntityNaturalPersonIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method fields.Field[EntityNaturalPersonIdentificationMethod] `json:"method,required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 fields.Field[string] `json:"number_last4,required"`
}

// MarshalJSON serializes EntityNaturalPersonIdentification into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *EntityNaturalPersonIdentification) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityNaturalPersonIdentification) String() (result string) {
	return fmt.Sprintf("&EntityNaturalPersonIdentification{Method:%s NumberLast4:%s}", r.Method, r.NumberLast4)
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
	Name fields.Field[string] `json:"name,required"`
	// The two individuals that share control of the entity.
	Individuals fields.Field[[]EntityJointIndividuals] `json:"individuals,required"`
}

// MarshalJSON serializes EntityJoint into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *EntityJoint) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityJoint) String() (result string) {
	return fmt.Sprintf("&EntityJoint{Name:%s Individuals:%s}", r.Name, core.Fmt(r.Individuals))
}

type EntityJointIndividuals struct {
	// The person's legal name.
	Name fields.Field[string] `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth fields.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// The person's address.
	Address fields.Field[EntityJointIndividualsAddress] `json:"address,required"`
	// A means of verifying the person's identity.
	Identification fields.Field[EntityJointIndividualsIdentification] `json:"identification,required"`
}

// MarshalJSON serializes EntityJointIndividuals into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EntityJointIndividuals) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityJointIndividuals) String() (result string) {
	return fmt.Sprintf("&EntityJointIndividuals{Name:%s DateOfBirth:%s Address:%s Identification:%s}", r.Name, r.DateOfBirth, r.Address, r.Identification)
}

type EntityJointIndividualsAddress struct {
	// The first line of the address.
	Line1 fields.Field[string] `json:"line1,required"`
	// The second line of the address.
	Line2 fields.Field[string] `json:"line2,required,nullable"`
	// The city of the address.
	City fields.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State fields.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip fields.Field[string] `json:"zip,required"`
}

// MarshalJSON serializes EntityJointIndividualsAddress into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *EntityJointIndividualsAddress) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityJointIndividualsAddress) String() (result string) {
	return fmt.Sprintf("&EntityJointIndividualsAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", r.Line1, r.Line2, r.City, r.State, r.Zip)
}

type EntityJointIndividualsIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method fields.Field[EntityJointIndividualsIdentificationMethod] `json:"method,required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 fields.Field[string] `json:"number_last4,required"`
}

// MarshalJSON serializes EntityJointIndividualsIdentification into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *EntityJointIndividualsIdentification) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityJointIndividualsIdentification) String() (result string) {
	return fmt.Sprintf("&EntityJointIndividualsIdentification{Method:%s NumberLast4:%s}", r.Method, r.NumberLast4)
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
	Name fields.Field[string] `json:"name,required"`
	// Whether the trust is `revocable` or `irrevocable`.
	Category fields.Field[EntityTrustCategory] `json:"category,required"`
	// The trust's address.
	Address fields.Field[EntityTrustAddress] `json:"address,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state in
	// which the trust was formed.
	FormationState fields.Field[string] `json:"formation_state,required,nullable"`
	// The Employer Identification Number (EIN) of the trust itself.
	TaxIdentifier fields.Field[string] `json:"tax_identifier,required,nullable"`
	// The trustees of the trust.
	Trustees fields.Field[[]EntityTrustTrustees] `json:"trustees,required"`
	// The grantor of the trust. Will be present if the `category` is `revocable`.
	Grantor fields.Field[EntityTrustGrantor] `json:"grantor,required,nullable"`
	// The ID for the File containing the formation document of the trust.
	FormationDocumentFileID fields.Field[string] `json:"formation_document_file_id,required,nullable"`
}

// MarshalJSON serializes EntityTrust into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *EntityTrust) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityTrust) String() (result string) {
	return fmt.Sprintf("&EntityTrust{Name:%s Category:%s Address:%s FormationState:%s TaxIdentifier:%s Trustees:%s Grantor:%s FormationDocumentFileID:%s}", r.Name, r.Category, r.Address, r.FormationState, r.TaxIdentifier, core.Fmt(r.Trustees), r.Grantor, r.FormationDocumentFileID)
}

type EntityTrustCategory string

const (
	EntityTrustCategoryRevocable   EntityTrustCategory = "revocable"
	EntityTrustCategoryIrrevocable EntityTrustCategory = "irrevocable"
)

type EntityTrustAddress struct {
	// The first line of the address.
	Line1 fields.Field[string] `json:"line1,required"`
	// The second line of the address.
	Line2 fields.Field[string] `json:"line2,required,nullable"`
	// The city of the address.
	City fields.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State fields.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip fields.Field[string] `json:"zip,required"`
}

// MarshalJSON serializes EntityTrustAddress into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *EntityTrustAddress) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityTrustAddress) String() (result string) {
	return fmt.Sprintf("&EntityTrustAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", r.Line1, r.Line2, r.City, r.State, r.Zip)
}

type EntityTrustTrustees struct {
	// The structure of the trustee. Will always be equal to `individual`.
	Structure fields.Field[EntityTrustTrusteesStructure] `json:"structure,required"`
	// The individual trustee of the trust. Will be present if the trustee's
	// `structure` is equal to `individual`.
	Individual fields.Field[EntityTrustTrusteesIndividual] `json:"individual,required,nullable"`
}

// MarshalJSON serializes EntityTrustTrustees into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EntityTrustTrustees) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityTrustTrustees) String() (result string) {
	return fmt.Sprintf("&EntityTrustTrustees{Structure:%s Individual:%s}", r.Structure, r.Individual)
}

type EntityTrustTrusteesStructure string

const (
	EntityTrustTrusteesStructureIndividual EntityTrustTrusteesStructure = "individual"
)

type EntityTrustTrusteesIndividual struct {
	// The person's legal name.
	Name fields.Field[string] `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth fields.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// The person's address.
	Address fields.Field[EntityTrustTrusteesIndividualAddress] `json:"address,required"`
	// A means of verifying the person's identity.
	Identification fields.Field[EntityTrustTrusteesIndividualIdentification] `json:"identification,required"`
}

// MarshalJSON serializes EntityTrustTrusteesIndividual into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *EntityTrustTrusteesIndividual) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityTrustTrusteesIndividual) String() (result string) {
	return fmt.Sprintf("&EntityTrustTrusteesIndividual{Name:%s DateOfBirth:%s Address:%s Identification:%s}", r.Name, r.DateOfBirth, r.Address, r.Identification)
}

type EntityTrustTrusteesIndividualAddress struct {
	// The first line of the address.
	Line1 fields.Field[string] `json:"line1,required"`
	// The second line of the address.
	Line2 fields.Field[string] `json:"line2,required,nullable"`
	// The city of the address.
	City fields.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State fields.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip fields.Field[string] `json:"zip,required"`
}

// MarshalJSON serializes EntityTrustTrusteesIndividualAddress into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *EntityTrustTrusteesIndividualAddress) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityTrustTrusteesIndividualAddress) String() (result string) {
	return fmt.Sprintf("&EntityTrustTrusteesIndividualAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", r.Line1, r.Line2, r.City, r.State, r.Zip)
}

type EntityTrustTrusteesIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method fields.Field[EntityTrustTrusteesIndividualIdentificationMethod] `json:"method,required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 fields.Field[string] `json:"number_last4,required"`
}

// MarshalJSON serializes EntityTrustTrusteesIndividualIdentification into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *EntityTrustTrusteesIndividualIdentification) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityTrustTrusteesIndividualIdentification) String() (result string) {
	return fmt.Sprintf("&EntityTrustTrusteesIndividualIdentification{Method:%s NumberLast4:%s}", r.Method, r.NumberLast4)
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
	Name fields.Field[string] `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth fields.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// The person's address.
	Address fields.Field[EntityTrustGrantorAddress] `json:"address,required"`
	// A means of verifying the person's identity.
	Identification fields.Field[EntityTrustGrantorIdentification] `json:"identification,required"`
}

// MarshalJSON serializes EntityTrustGrantor into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *EntityTrustGrantor) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityTrustGrantor) String() (result string) {
	return fmt.Sprintf("&EntityTrustGrantor{Name:%s DateOfBirth:%s Address:%s Identification:%s}", r.Name, r.DateOfBirth, r.Address, r.Identification)
}

type EntityTrustGrantorAddress struct {
	// The first line of the address.
	Line1 fields.Field[string] `json:"line1,required"`
	// The second line of the address.
	Line2 fields.Field[string] `json:"line2,required,nullable"`
	// The city of the address.
	City fields.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State fields.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip fields.Field[string] `json:"zip,required"`
}

// MarshalJSON serializes EntityTrustGrantorAddress into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EntityTrustGrantorAddress) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityTrustGrantorAddress) String() (result string) {
	return fmt.Sprintf("&EntityTrustGrantorAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", r.Line1, r.Line2, r.City, r.State, r.Zip)
}

type EntityTrustGrantorIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method fields.Field[EntityTrustGrantorIdentificationMethod] `json:"method,required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 fields.Field[string] `json:"number_last4,required"`
}

// MarshalJSON serializes EntityTrustGrantorIdentification into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *EntityTrustGrantorIdentification) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntityTrustGrantorIdentification) String() (result string) {
	return fmt.Sprintf("&EntityTrustGrantorIdentification{Method:%s NumberLast4:%s}", r.Method, r.NumberLast4)
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
	FileID fields.Field[string] `json:"file_id,required"`
}

// MarshalJSON serializes EntitySupplementalDocuments into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EntitySupplementalDocuments) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EntitySupplementalDocuments) String() (result string) {
	return fmt.Sprintf("&EntitySupplementalDocuments{FileID:%s}", r.FileID)
}

type CreateAnEntityParameters struct {
	// The type of Entity to create.
	Structure fields.Field[CreateAnEntityParametersStructure] `json:"structure,required"`
	// Details of the corporation entity to create. Required if `structure` is equal to
	// `corporation`.
	Corporation fields.Field[CreateAnEntityParametersCorporation] `json:"corporation"`
	// Details of the natural person entity to create. Required if `structure` is equal
	// to `natural_person`. Natural people entities should be submitted with
	// `social_security_number` or `individual_taxpayer_identification_number`
	// identification methods.
	NaturalPerson fields.Field[CreateAnEntityParametersNaturalPerson] `json:"natural_person"`
	// Details of the joint entity to create. Required if `structure` is equal to
	// `joint`.
	Joint fields.Field[CreateAnEntityParametersJoint] `json:"joint"`
	// Details of the trust entity to create. Required if `structure` is equal to
	// `trust`.
	Trust fields.Field[CreateAnEntityParametersTrust] `json:"trust"`
	// The description you choose to give the entity.
	Description fields.Field[string] `json:"description"`
	// The relationship between your group and the entity.
	Relationship fields.Field[CreateAnEntityParametersRelationship] `json:"relationship,required"`
	// Additional documentation associated with the entity.
	SupplementalDocuments fields.Field[[]CreateAnEntityParametersSupplementalDocuments] `json:"supplemental_documents"`
}

// MarshalJSON serializes CreateAnEntityParameters into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CreateAnEntityParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateAnEntityParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParameters{Structure:%s Corporation:%s NaturalPerson:%s Joint:%s Trust:%s Description:%s Relationship:%s SupplementalDocuments:%s}", r.Structure, r.Corporation, r.NaturalPerson, r.Joint, r.Trust, r.Description, r.Relationship, core.Fmt(r.SupplementalDocuments))
}

type CreateAnEntityParametersStructure string

const (
	CreateAnEntityParametersStructureCorporation   CreateAnEntityParametersStructure = "corporation"
	CreateAnEntityParametersStructureNaturalPerson CreateAnEntityParametersStructure = "natural_person"
	CreateAnEntityParametersStructureJoint         CreateAnEntityParametersStructure = "joint"
	CreateAnEntityParametersStructureTrust         CreateAnEntityParametersStructure = "trust"
)

type CreateAnEntityParametersCorporation struct {
	// The legal name of the corporation.
	Name fields.Field[string] `json:"name,required"`
	// The website of the corporation.
	Website fields.Field[string] `json:"website"`
	// The Employer Identification Number (EIN) for the corporation.
	TaxIdentifier fields.Field[string] `json:"tax_identifier,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the
	// corporation's state of incorporation.
	IncorporationState fields.Field[string] `json:"incorporation_state"`
	// The corporation's address.
	Address fields.Field[CreateAnEntityParametersCorporationAddress] `json:"address,required"`
	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwners fields.Field[[]CreateAnEntityParametersCorporationBeneficialOwners] `json:"beneficial_owners,required"`
}

func (r CreateAnEntityParametersCorporation) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersCorporation{Name:%s Website:%s TaxIdentifier:%s IncorporationState:%s Address:%s BeneficialOwners:%s}", r.Name, r.Website, r.TaxIdentifier, r.IncorporationState, r.Address, core.Fmt(r.BeneficialOwners))
}

type CreateAnEntityParametersCorporationAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 fields.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 fields.Field[string] `json:"line2"`
	// The city of the address.
	City fields.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State fields.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip fields.Field[string] `json:"zip,required"`
}

func (r CreateAnEntityParametersCorporationAddress) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersCorporationAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", r.Line1, r.Line2, r.City, r.State, r.Zip)
}

type CreateAnEntityParametersCorporationBeneficialOwners struct {
	// Personal details for the beneficial owner.
	Individual fields.Field[CreateAnEntityParametersCorporationBeneficialOwnersIndividual] `json:"individual,required"`
	// This person's role or title within the entity.
	CompanyTitle fields.Field[string] `json:"company_title"`
	// Why this person is considered a beneficial owner of the entity.
	Prong fields.Field[CreateAnEntityParametersCorporationBeneficialOwnersProng] `json:"prong,required"`
}

func (r CreateAnEntityParametersCorporationBeneficialOwners) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersCorporationBeneficialOwners{Individual:%s CompanyTitle:%s Prong:%s}", r.Individual, r.CompanyTitle, r.Prong)
}

type CreateAnEntityParametersCorporationBeneficialOwnersIndividual struct {
	// The person's legal name.
	Name fields.Field[string] `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth fields.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// The individual's address.
	Address fields.Field[CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress] `json:"address,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID fields.Field[bool] `json:"confirmed_no_us_tax_id"`
	// A means of verifying the person's identity.
	Identification fields.Field[CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification] `json:"identification,required"`
}

func (r CreateAnEntityParametersCorporationBeneficialOwnersIndividual) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersCorporationBeneficialOwnersIndividual{Name:%s DateOfBirth:%s Address:%s ConfirmedNoUsTaxID:%s Identification:%s}", r.Name, r.DateOfBirth, r.Address, r.ConfirmedNoUsTaxID, r.Identification)
}

type CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 fields.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 fields.Field[string] `json:"line2"`
	// The city of the address.
	City fields.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State fields.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip fields.Field[string] `json:"zip,required"`
}

func (r CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", r.Line1, r.Line2, r.City, r.State, r.Zip)
}

type CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method fields.Field[CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethod] `json:"method,required"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number fields.Field[string] `json:"number,required"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport fields.Field[CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport] `json:"passport"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense fields.Field[CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other fields.Field[CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationOther] `json:"other"`
}

func (r CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification{Method:%s Number:%s Passport:%s DriversLicense:%s Other:%s}", r.Method, r.Number, r.Passport, r.DriversLicense, r.Other)
}

type CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethod string

const (
	CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber                   CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethod = "social_security_number"
	CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethodPassport                               CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethod = "passport"
	CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethodDriversLicense                         CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethod = "drivers_license"
	CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethodOther                                  CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethod = "other"
)

type CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport struct {
	// The identifier of the File containing the passport.
	FileID fields.Field[string] `json:"file_id,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate fields.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The country that issued the passport.
	Country fields.Field[string] `json:"country,required"`
}

func (r CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport{FileID:%s ExpirationDate:%s Country:%s}", r.FileID, r.ExpirationDate, r.Country)
}

type CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationDriversLicense struct {
	// The identifier of the File containing the driver's license.
	FileID fields.Field[string] `json:"file_id,required"`
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate fields.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The state that issued the provided driver's license.
	State fields.Field[string] `json:"state,required"`
}

func (r CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationDriversLicense) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationDriversLicense{FileID:%s ExpirationDate:%s State:%s}", r.FileID, r.ExpirationDate, r.State)
}

type CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document.
	Country fields.Field[string] `json:"country,required"`
	// A description of the document submitted.
	Description fields.Field[string] `json:"description,required"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate fields.Field[time.Time] `json:"expiration_date" format:"date"`
	// The identifier of the File containing the document.
	FileID fields.Field[string] `json:"file_id,required"`
}

func (r CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationOther) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationOther{Country:%s Description:%s ExpirationDate:%s FileID:%s}", r.Country, r.Description, r.ExpirationDate, r.FileID)
}

type CreateAnEntityParametersCorporationBeneficialOwnersProng string

const (
	CreateAnEntityParametersCorporationBeneficialOwnersProngOwnership CreateAnEntityParametersCorporationBeneficialOwnersProng = "ownership"
	CreateAnEntityParametersCorporationBeneficialOwnersProngControl   CreateAnEntityParametersCorporationBeneficialOwnersProng = "control"
)

type CreateAnEntityParametersNaturalPerson struct {
	// The person's legal name.
	Name fields.Field[string] `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth fields.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// The individual's address.
	Address fields.Field[CreateAnEntityParametersNaturalPersonAddress] `json:"address,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID fields.Field[bool] `json:"confirmed_no_us_tax_id"`
	// A means of verifying the person's identity.
	Identification fields.Field[CreateAnEntityParametersNaturalPersonIdentification] `json:"identification,required"`
}

func (r CreateAnEntityParametersNaturalPerson) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersNaturalPerson{Name:%s DateOfBirth:%s Address:%s ConfirmedNoUsTaxID:%s Identification:%s}", r.Name, r.DateOfBirth, r.Address, r.ConfirmedNoUsTaxID, r.Identification)
}

type CreateAnEntityParametersNaturalPersonAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 fields.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 fields.Field[string] `json:"line2"`
	// The city of the address.
	City fields.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State fields.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip fields.Field[string] `json:"zip,required"`
}

func (r CreateAnEntityParametersNaturalPersonAddress) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersNaturalPersonAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", r.Line1, r.Line2, r.City, r.State, r.Zip)
}

type CreateAnEntityParametersNaturalPersonIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method fields.Field[CreateAnEntityParametersNaturalPersonIdentificationMethod] `json:"method,required"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number fields.Field[string] `json:"number,required"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport fields.Field[CreateAnEntityParametersNaturalPersonIdentificationPassport] `json:"passport"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense fields.Field[CreateAnEntityParametersNaturalPersonIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other fields.Field[CreateAnEntityParametersNaturalPersonIdentificationOther] `json:"other"`
}

func (r CreateAnEntityParametersNaturalPersonIdentification) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersNaturalPersonIdentification{Method:%s Number:%s Passport:%s DriversLicense:%s Other:%s}", r.Method, r.Number, r.Passport, r.DriversLicense, r.Other)
}

type CreateAnEntityParametersNaturalPersonIdentificationMethod string

const (
	CreateAnEntityParametersNaturalPersonIdentificationMethodSocialSecurityNumber                   CreateAnEntityParametersNaturalPersonIdentificationMethod = "social_security_number"
	CreateAnEntityParametersNaturalPersonIdentificationMethodIndividualTaxpayerIdentificationNumber CreateAnEntityParametersNaturalPersonIdentificationMethod = "individual_taxpayer_identification_number"
	CreateAnEntityParametersNaturalPersonIdentificationMethodPassport                               CreateAnEntityParametersNaturalPersonIdentificationMethod = "passport"
	CreateAnEntityParametersNaturalPersonIdentificationMethodDriversLicense                         CreateAnEntityParametersNaturalPersonIdentificationMethod = "drivers_license"
	CreateAnEntityParametersNaturalPersonIdentificationMethodOther                                  CreateAnEntityParametersNaturalPersonIdentificationMethod = "other"
)

type CreateAnEntityParametersNaturalPersonIdentificationPassport struct {
	// The identifier of the File containing the passport.
	FileID fields.Field[string] `json:"file_id,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate fields.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The country that issued the passport.
	Country fields.Field[string] `json:"country,required"`
}

func (r CreateAnEntityParametersNaturalPersonIdentificationPassport) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersNaturalPersonIdentificationPassport{FileID:%s ExpirationDate:%s Country:%s}", r.FileID, r.ExpirationDate, r.Country)
}

type CreateAnEntityParametersNaturalPersonIdentificationDriversLicense struct {
	// The identifier of the File containing the driver's license.
	FileID fields.Field[string] `json:"file_id,required"`
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate fields.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The state that issued the provided driver's license.
	State fields.Field[string] `json:"state,required"`
}

func (r CreateAnEntityParametersNaturalPersonIdentificationDriversLicense) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersNaturalPersonIdentificationDriversLicense{FileID:%s ExpirationDate:%s State:%s}", r.FileID, r.ExpirationDate, r.State)
}

type CreateAnEntityParametersNaturalPersonIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document.
	Country fields.Field[string] `json:"country,required"`
	// A description of the document submitted.
	Description fields.Field[string] `json:"description,required"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate fields.Field[time.Time] `json:"expiration_date" format:"date"`
	// The identifier of the File containing the document.
	FileID fields.Field[string] `json:"file_id,required"`
}

func (r CreateAnEntityParametersNaturalPersonIdentificationOther) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersNaturalPersonIdentificationOther{Country:%s Description:%s ExpirationDate:%s FileID:%s}", r.Country, r.Description, r.ExpirationDate, r.FileID)
}

type CreateAnEntityParametersJoint struct {
	// The name of the joint entity.
	Name fields.Field[string] `json:"name"`
	// The two individuals that share control of the entity.
	Individuals fields.Field[[]CreateAnEntityParametersJointIndividuals] `json:"individuals,required"`
}

func (r CreateAnEntityParametersJoint) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersJoint{Name:%s Individuals:%s}", r.Name, core.Fmt(r.Individuals))
}

type CreateAnEntityParametersJointIndividuals struct {
	// The person's legal name.
	Name fields.Field[string] `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth fields.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// The individual's address.
	Address fields.Field[CreateAnEntityParametersJointIndividualsAddress] `json:"address,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID fields.Field[bool] `json:"confirmed_no_us_tax_id"`
	// A means of verifying the person's identity.
	Identification fields.Field[CreateAnEntityParametersJointIndividualsIdentification] `json:"identification,required"`
}

func (r CreateAnEntityParametersJointIndividuals) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersJointIndividuals{Name:%s DateOfBirth:%s Address:%s ConfirmedNoUsTaxID:%s Identification:%s}", r.Name, r.DateOfBirth, r.Address, r.ConfirmedNoUsTaxID, r.Identification)
}

type CreateAnEntityParametersJointIndividualsAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 fields.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 fields.Field[string] `json:"line2"`
	// The city of the address.
	City fields.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State fields.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip fields.Field[string] `json:"zip,required"`
}

func (r CreateAnEntityParametersJointIndividualsAddress) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersJointIndividualsAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", r.Line1, r.Line2, r.City, r.State, r.Zip)
}

type CreateAnEntityParametersJointIndividualsIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method fields.Field[CreateAnEntityParametersJointIndividualsIdentificationMethod] `json:"method,required"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number fields.Field[string] `json:"number,required"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport fields.Field[CreateAnEntityParametersJointIndividualsIdentificationPassport] `json:"passport"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense fields.Field[CreateAnEntityParametersJointIndividualsIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other fields.Field[CreateAnEntityParametersJointIndividualsIdentificationOther] `json:"other"`
}

func (r CreateAnEntityParametersJointIndividualsIdentification) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersJointIndividualsIdentification{Method:%s Number:%s Passport:%s DriversLicense:%s Other:%s}", r.Method, r.Number, r.Passport, r.DriversLicense, r.Other)
}

type CreateAnEntityParametersJointIndividualsIdentificationMethod string

const (
	CreateAnEntityParametersJointIndividualsIdentificationMethodSocialSecurityNumber                   CreateAnEntityParametersJointIndividualsIdentificationMethod = "social_security_number"
	CreateAnEntityParametersJointIndividualsIdentificationMethodIndividualTaxpayerIdentificationNumber CreateAnEntityParametersJointIndividualsIdentificationMethod = "individual_taxpayer_identification_number"
	CreateAnEntityParametersJointIndividualsIdentificationMethodPassport                               CreateAnEntityParametersJointIndividualsIdentificationMethod = "passport"
	CreateAnEntityParametersJointIndividualsIdentificationMethodDriversLicense                         CreateAnEntityParametersJointIndividualsIdentificationMethod = "drivers_license"
	CreateAnEntityParametersJointIndividualsIdentificationMethodOther                                  CreateAnEntityParametersJointIndividualsIdentificationMethod = "other"
)

type CreateAnEntityParametersJointIndividualsIdentificationPassport struct {
	// The identifier of the File containing the passport.
	FileID fields.Field[string] `json:"file_id,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate fields.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The country that issued the passport.
	Country fields.Field[string] `json:"country,required"`
}

func (r CreateAnEntityParametersJointIndividualsIdentificationPassport) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersJointIndividualsIdentificationPassport{FileID:%s ExpirationDate:%s Country:%s}", r.FileID, r.ExpirationDate, r.Country)
}

type CreateAnEntityParametersJointIndividualsIdentificationDriversLicense struct {
	// The identifier of the File containing the driver's license.
	FileID fields.Field[string] `json:"file_id,required"`
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate fields.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The state that issued the provided driver's license.
	State fields.Field[string] `json:"state,required"`
}

func (r CreateAnEntityParametersJointIndividualsIdentificationDriversLicense) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersJointIndividualsIdentificationDriversLicense{FileID:%s ExpirationDate:%s State:%s}", r.FileID, r.ExpirationDate, r.State)
}

type CreateAnEntityParametersJointIndividualsIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document.
	Country fields.Field[string] `json:"country,required"`
	// A description of the document submitted.
	Description fields.Field[string] `json:"description,required"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate fields.Field[time.Time] `json:"expiration_date" format:"date"`
	// The identifier of the File containing the document.
	FileID fields.Field[string] `json:"file_id,required"`
}

func (r CreateAnEntityParametersJointIndividualsIdentificationOther) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersJointIndividualsIdentificationOther{Country:%s Description:%s ExpirationDate:%s FileID:%s}", r.Country, r.Description, r.ExpirationDate, r.FileID)
}

type CreateAnEntityParametersTrust struct {
	// The legal name of the trust.
	Name fields.Field[string] `json:"name,required"`
	// Whether the trust is `revocable` or `irrevocable`. Irrevocable trusts require
	// their own Employer Identification Number. Revocable trusts require information
	// about the individual `grantor` who created the trust.
	Category fields.Field[CreateAnEntityParametersTrustCategory] `json:"category,required"`
	// The Employer Identification Number (EIN) for the trust. Required if `category`
	// is equal to `irrevocable`.
	TaxIdentifier fields.Field[string] `json:"tax_identifier"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state in
	// which the trust was formed.
	FormationState fields.Field[string] `json:"formation_state"`
	// The trust's address.
	Address fields.Field[CreateAnEntityParametersTrustAddress] `json:"address,required"`
	// The identifier of the File containing the formation document of the trust.
	FormationDocumentFileID fields.Field[string] `json:"formation_document_file_id"`
	// The trustees of the trust.
	Trustees fields.Field[[]CreateAnEntityParametersTrustTrustees] `json:"trustees,required"`
	// The grantor of the trust. Required if `category` is equal to `revocable`.
	Grantor fields.Field[CreateAnEntityParametersTrustGrantor] `json:"grantor"`
}

func (r CreateAnEntityParametersTrust) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrust{Name:%s Category:%s TaxIdentifier:%s FormationState:%s Address:%s FormationDocumentFileID:%s Trustees:%s Grantor:%s}", r.Name, r.Category, r.TaxIdentifier, r.FormationState, r.Address, r.FormationDocumentFileID, core.Fmt(r.Trustees), r.Grantor)
}

type CreateAnEntityParametersTrustCategory string

const (
	CreateAnEntityParametersTrustCategoryRevocable   CreateAnEntityParametersTrustCategory = "revocable"
	CreateAnEntityParametersTrustCategoryIrrevocable CreateAnEntityParametersTrustCategory = "irrevocable"
)

type CreateAnEntityParametersTrustAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 fields.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 fields.Field[string] `json:"line2"`
	// The city of the address.
	City fields.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State fields.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip fields.Field[string] `json:"zip,required"`
}

func (r CreateAnEntityParametersTrustAddress) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", r.Line1, r.Line2, r.City, r.State, r.Zip)
}

type CreateAnEntityParametersTrustTrustees struct {
	// The structure of the trustee.
	Structure fields.Field[CreateAnEntityParametersTrustTrusteesStructure] `json:"structure,required"`
	// Details of the individual trustee. Required when the trustee `structure` is
	// equal to `individual`.
	Individual fields.Field[CreateAnEntityParametersTrustTrusteesIndividual] `json:"individual"`
}

func (r CreateAnEntityParametersTrustTrustees) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustTrustees{Structure:%s Individual:%s}", r.Structure, r.Individual)
}

type CreateAnEntityParametersTrustTrusteesStructure string

const (
	CreateAnEntityParametersTrustTrusteesStructureIndividual CreateAnEntityParametersTrustTrusteesStructure = "individual"
)

type CreateAnEntityParametersTrustTrusteesIndividual struct {
	// The person's legal name.
	Name fields.Field[string] `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth fields.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// The individual's address.
	Address fields.Field[CreateAnEntityParametersTrustTrusteesIndividualAddress] `json:"address,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID fields.Field[bool] `json:"confirmed_no_us_tax_id"`
	// A means of verifying the person's identity.
	Identification fields.Field[CreateAnEntityParametersTrustTrusteesIndividualIdentification] `json:"identification,required"`
}

func (r CreateAnEntityParametersTrustTrusteesIndividual) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustTrusteesIndividual{Name:%s DateOfBirth:%s Address:%s ConfirmedNoUsTaxID:%s Identification:%s}", r.Name, r.DateOfBirth, r.Address, r.ConfirmedNoUsTaxID, r.Identification)
}

type CreateAnEntityParametersTrustTrusteesIndividualAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 fields.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 fields.Field[string] `json:"line2"`
	// The city of the address.
	City fields.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State fields.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip fields.Field[string] `json:"zip,required"`
}

func (r CreateAnEntityParametersTrustTrusteesIndividualAddress) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustTrusteesIndividualAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", r.Line1, r.Line2, r.City, r.State, r.Zip)
}

type CreateAnEntityParametersTrustTrusteesIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method fields.Field[CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethod] `json:"method,required"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number fields.Field[string] `json:"number,required"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport fields.Field[CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport] `json:"passport"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense fields.Field[CreateAnEntityParametersTrustTrusteesIndividualIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other fields.Field[CreateAnEntityParametersTrustTrusteesIndividualIdentificationOther] `json:"other"`
}

func (r CreateAnEntityParametersTrustTrusteesIndividualIdentification) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustTrusteesIndividualIdentification{Method:%s Number:%s Passport:%s DriversLicense:%s Other:%s}", r.Method, r.Number, r.Passport, r.DriversLicense, r.Other)
}

type CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethod string

const (
	CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber                   CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethod = "social_security_number"
	CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethodPassport                               CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethod = "passport"
	CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethodDriversLicense                         CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethod = "drivers_license"
	CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethodOther                                  CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethod = "other"
)

type CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport struct {
	// The identifier of the File containing the passport.
	FileID fields.Field[string] `json:"file_id,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate fields.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The country that issued the passport.
	Country fields.Field[string] `json:"country,required"`
}

func (r CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport{FileID:%s ExpirationDate:%s Country:%s}", r.FileID, r.ExpirationDate, r.Country)
}

type CreateAnEntityParametersTrustTrusteesIndividualIdentificationDriversLicense struct {
	// The identifier of the File containing the driver's license.
	FileID fields.Field[string] `json:"file_id,required"`
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate fields.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The state that issued the provided driver's license.
	State fields.Field[string] `json:"state,required"`
}

func (r CreateAnEntityParametersTrustTrusteesIndividualIdentificationDriversLicense) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustTrusteesIndividualIdentificationDriversLicense{FileID:%s ExpirationDate:%s State:%s}", r.FileID, r.ExpirationDate, r.State)
}

type CreateAnEntityParametersTrustTrusteesIndividualIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document.
	Country fields.Field[string] `json:"country,required"`
	// A description of the document submitted.
	Description fields.Field[string] `json:"description,required"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate fields.Field[time.Time] `json:"expiration_date" format:"date"`
	// The identifier of the File containing the document.
	FileID fields.Field[string] `json:"file_id,required"`
}

func (r CreateAnEntityParametersTrustTrusteesIndividualIdentificationOther) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustTrusteesIndividualIdentificationOther{Country:%s Description:%s ExpirationDate:%s FileID:%s}", r.Country, r.Description, r.ExpirationDate, r.FileID)
}

type CreateAnEntityParametersTrustGrantor struct {
	// The person's legal name.
	Name fields.Field[string] `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth fields.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// The individual's address.
	Address fields.Field[CreateAnEntityParametersTrustGrantorAddress] `json:"address,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID fields.Field[bool] `json:"confirmed_no_us_tax_id"`
	// A means of verifying the person's identity.
	Identification fields.Field[CreateAnEntityParametersTrustGrantorIdentification] `json:"identification,required"`
}

func (r CreateAnEntityParametersTrustGrantor) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustGrantor{Name:%s DateOfBirth:%s Address:%s ConfirmedNoUsTaxID:%s Identification:%s}", r.Name, r.DateOfBirth, r.Address, r.ConfirmedNoUsTaxID, r.Identification)
}

type CreateAnEntityParametersTrustGrantorAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 fields.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 fields.Field[string] `json:"line2"`
	// The city of the address.
	City fields.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State fields.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip fields.Field[string] `json:"zip,required"`
}

func (r CreateAnEntityParametersTrustGrantorAddress) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustGrantorAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", r.Line1, r.Line2, r.City, r.State, r.Zip)
}

type CreateAnEntityParametersTrustGrantorIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method fields.Field[CreateAnEntityParametersTrustGrantorIdentificationMethod] `json:"method,required"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number fields.Field[string] `json:"number,required"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport fields.Field[CreateAnEntityParametersTrustGrantorIdentificationPassport] `json:"passport"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense fields.Field[CreateAnEntityParametersTrustGrantorIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other fields.Field[CreateAnEntityParametersTrustGrantorIdentificationOther] `json:"other"`
}

func (r CreateAnEntityParametersTrustGrantorIdentification) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustGrantorIdentification{Method:%s Number:%s Passport:%s DriversLicense:%s Other:%s}", r.Method, r.Number, r.Passport, r.DriversLicense, r.Other)
}

type CreateAnEntityParametersTrustGrantorIdentificationMethod string

const (
	CreateAnEntityParametersTrustGrantorIdentificationMethodSocialSecurityNumber                   CreateAnEntityParametersTrustGrantorIdentificationMethod = "social_security_number"
	CreateAnEntityParametersTrustGrantorIdentificationMethodIndividualTaxpayerIdentificationNumber CreateAnEntityParametersTrustGrantorIdentificationMethod = "individual_taxpayer_identification_number"
	CreateAnEntityParametersTrustGrantorIdentificationMethodPassport                               CreateAnEntityParametersTrustGrantorIdentificationMethod = "passport"
	CreateAnEntityParametersTrustGrantorIdentificationMethodDriversLicense                         CreateAnEntityParametersTrustGrantorIdentificationMethod = "drivers_license"
	CreateAnEntityParametersTrustGrantorIdentificationMethodOther                                  CreateAnEntityParametersTrustGrantorIdentificationMethod = "other"
)

type CreateAnEntityParametersTrustGrantorIdentificationPassport struct {
	// The identifier of the File containing the passport.
	FileID fields.Field[string] `json:"file_id,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate fields.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The country that issued the passport.
	Country fields.Field[string] `json:"country,required"`
}

func (r CreateAnEntityParametersTrustGrantorIdentificationPassport) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustGrantorIdentificationPassport{FileID:%s ExpirationDate:%s Country:%s}", r.FileID, r.ExpirationDate, r.Country)
}

type CreateAnEntityParametersTrustGrantorIdentificationDriversLicense struct {
	// The identifier of the File containing the driver's license.
	FileID fields.Field[string] `json:"file_id,required"`
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate fields.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The state that issued the provided driver's license.
	State fields.Field[string] `json:"state,required"`
}

func (r CreateAnEntityParametersTrustGrantorIdentificationDriversLicense) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustGrantorIdentificationDriversLicense{FileID:%s ExpirationDate:%s State:%s}", r.FileID, r.ExpirationDate, r.State)
}

type CreateAnEntityParametersTrustGrantorIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document.
	Country fields.Field[string] `json:"country,required"`
	// A description of the document submitted.
	Description fields.Field[string] `json:"description,required"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate fields.Field[time.Time] `json:"expiration_date" format:"date"`
	// The identifier of the File containing the document.
	FileID fields.Field[string] `json:"file_id,required"`
}

func (r CreateAnEntityParametersTrustGrantorIdentificationOther) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustGrantorIdentificationOther{Country:%s Description:%s ExpirationDate:%s FileID:%s}", r.Country, r.Description, r.ExpirationDate, r.FileID)
}

type CreateAnEntityParametersRelationship string

const (
	CreateAnEntityParametersRelationshipAffiliated    CreateAnEntityParametersRelationship = "affiliated"
	CreateAnEntityParametersRelationshipInformational CreateAnEntityParametersRelationship = "informational"
	CreateAnEntityParametersRelationshipUnaffiliated  CreateAnEntityParametersRelationship = "unaffiliated"
)

type CreateAnEntityParametersSupplementalDocuments struct {
	// The identifier of the File containing the document.
	FileID fields.Field[string] `json:"file_id,required"`
}

func (r CreateAnEntityParametersSupplementalDocuments) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersSupplementalDocuments{FileID:%s}", r.FileID)
}

type EntityListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
}

// URLQuery serializes EntityListParams into a url.Values of the query parameters
// associated with this value
func (r *EntityListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r EntityListParams) String() (result string) {
	return fmt.Sprintf("&EntityListParams{Cursor:%s Limit:%s}", r.Cursor, r.Limit)
}
