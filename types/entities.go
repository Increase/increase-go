package types

import (
	"fmt"
	"increase/core"
	"increase/core/pjson"
	"increase/pagination"
)

type Entity struct {
	// The entity's identifier.
	ID *string `pjson:"id"`
	// The entity's legal structure.
	Structure *EntityStructure `pjson:"structure"`
	// Details of the corporation entity. Will be present if `structure` is equal to
	// `corporation`.
	Corporation *EntityCorporation `pjson:"corporation"`
	// Details of the natural person entity. Will be present if `structure` is equal to
	// `natural_person`.
	NaturalPerson *EntityNaturalPerson `pjson:"natural_person"`
	// Details of the joint entity. Will be present if `structure` is equal to `joint`.
	Joint *EntityJoint `pjson:"joint"`
	// Details of the trust entity. Will be present if `structure` is equal to `trust`.
	Trust *EntityTrust `pjson:"trust"`
	// A constant representing the object's type. For this resource it will always be
	// `entity`.
	Type *EntityType `pjson:"type"`
	// The entity's description for display purposes.
	Description *string `pjson:"description"`
	// The relationship between your group and the entity.
	Relationship *EntityRelationship `pjson:"relationship"`
	// Additional documentation associated with the entity.
	SupplementalDocuments *[]EntitySupplementalDocuments `pjson:"supplemental_documents"`
	jsonFields            map[string]interface{}         `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into Entity using the internal
// pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *Entity) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes Entity into an array of bytes using the gjson library.
// Members of the `Extras` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Entity) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r Entity) String() (result string) {
	return fmt.Sprintf("&Entity{ID:%s Structure:%s Corporation:%s NaturalPerson:%s Joint:%s Trust:%s Type:%s Description:%s Relationship:%s SupplementalDocuments:%s}", core.FmtP(r.ID), core.FmtP(r.Structure), r.Corporation, r.NaturalPerson, r.Joint, r.Trust, core.FmtP(r.Type), core.FmtP(r.Description), core.FmtP(r.Relationship), core.Fmt(r.SupplementalDocuments))
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
	Name *string `pjson:"name"`
	// The website of the corporation.
	Website *string `pjson:"website"`
	// The Employer Identification Number (EIN) for the corporation.
	TaxIdentifier *string `pjson:"tax_identifier"`
	// The two-letter United States Postal Service (USPS) abbreviation for the
	// corporation's state of incorporation.
	IncorporationState *string `pjson:"incorporation_state"`
	// The corporation's address.
	Address *EntityCorporationAddress `pjson:"address"`
	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwners *[]EntityCorporationBeneficialOwners `pjson:"beneficial_owners"`
	jsonFields       map[string]interface{}               `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EntityCorporation using the
// internal pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *EntityCorporation) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityCorporation into an array of bytes using the gjson
// library. Members of the `Extras` field are serialized into the top-level, and
// will overwrite known members of the same name.
func (r *EntityCorporation) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityCorporation) String() (result string) {
	return fmt.Sprintf("&EntityCorporation{Name:%s Website:%s TaxIdentifier:%s IncorporationState:%s Address:%s BeneficialOwners:%s}", core.FmtP(r.Name), core.FmtP(r.Website), core.FmtP(r.TaxIdentifier), core.FmtP(r.IncorporationState), r.Address, core.Fmt(r.BeneficialOwners))
}

type EntityCorporationAddress struct {
	// The first line of the address.
	Line1 *string `pjson:"line1"`
	// The second line of the address.
	Line2 *string `pjson:"line2"`
	// The city of the address.
	City *string `pjson:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `pjson:"state"`
	// The ZIP code of the address.
	Zip        *string                `pjson:"zip"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EntityCorporationAddress
// using the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *EntityCorporationAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityCorporationAddress into an array of bytes using the
// gjson library. Members of the `Extras` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *EntityCorporationAddress) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityCorporationAddress) String() (result string) {
	return fmt.Sprintf("&EntityCorporationAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", core.FmtP(r.Line1), core.FmtP(r.Line2), core.FmtP(r.City), core.FmtP(r.State), core.FmtP(r.Zip))
}

type EntityCorporationBeneficialOwners struct {
	// Personal details for the beneficial owner.
	Individual *EntityCorporationBeneficialOwnersIndividual `pjson:"individual"`
	// This person's role or title within the entity.
	CompanyTitle *string `pjson:"company_title"`
	// Why this person is considered a beneficial owner of the entity.
	Prong      *EntityCorporationBeneficialOwnersProng `pjson:"prong"`
	jsonFields map[string]interface{}                  `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// EntityCorporationBeneficialOwners using the internal pjson library. Unrecognized
// fields are stored in the `Extras` property.
func (r *EntityCorporationBeneficialOwners) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityCorporationBeneficialOwners into an array of bytes
// using the gjson library. Members of the `Extras` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EntityCorporationBeneficialOwners) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityCorporationBeneficialOwners) String() (result string) {
	return fmt.Sprintf("&EntityCorporationBeneficialOwners{Individual:%s CompanyTitle:%s Prong:%s}", r.Individual, core.FmtP(r.CompanyTitle), core.FmtP(r.Prong))
}

type EntityCorporationBeneficialOwnersIndividual struct {
	// The person's legal name.
	Name *string `pjson:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `pjson:"date_of_birth"`
	// The person's address.
	Address *EntityCorporationBeneficialOwnersIndividualAddress `pjson:"address"`
	// A means of verifying the person's identity.
	Identification *EntityCorporationBeneficialOwnersIndividualIdentification `pjson:"identification"`
	jsonFields     map[string]interface{}                                     `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// EntityCorporationBeneficialOwnersIndividual using the internal pjson library.
// Unrecognized fields are stored in the `Extras` property.
func (r *EntityCorporationBeneficialOwnersIndividual) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityCorporationBeneficialOwnersIndividual into an array
// of bytes using the gjson library. Members of the `Extras` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *EntityCorporationBeneficialOwnersIndividual) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityCorporationBeneficialOwnersIndividual) String() (result string) {
	return fmt.Sprintf("&EntityCorporationBeneficialOwnersIndividual{Name:%s DateOfBirth:%s Address:%s Identification:%s}", core.FmtP(r.Name), core.FmtP(r.DateOfBirth), r.Address, r.Identification)
}

type EntityCorporationBeneficialOwnersIndividualAddress struct {
	// The first line of the address.
	Line1 *string `pjson:"line1"`
	// The second line of the address.
	Line2 *string `pjson:"line2"`
	// The city of the address.
	City *string `pjson:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `pjson:"state"`
	// The ZIP code of the address.
	Zip        *string                `pjson:"zip"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// EntityCorporationBeneficialOwnersIndividualAddress using the internal pjson
// library. Unrecognized fields are stored in the `Extras` property.
func (r *EntityCorporationBeneficialOwnersIndividualAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityCorporationBeneficialOwnersIndividualAddress into
// an array of bytes using the gjson library. Members of the `Extras` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *EntityCorporationBeneficialOwnersIndividualAddress) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityCorporationBeneficialOwnersIndividualAddress) String() (result string) {
	return fmt.Sprintf("&EntityCorporationBeneficialOwnersIndividualAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", core.FmtP(r.Line1), core.FmtP(r.Line2), core.FmtP(r.City), core.FmtP(r.State), core.FmtP(r.Zip))
}

type EntityCorporationBeneficialOwnersIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method *EntityCorporationBeneficialOwnersIndividualIdentificationMethod `pjson:"method"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 *string                `pjson:"number_last4"`
	jsonFields  map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// EntityCorporationBeneficialOwnersIndividualIdentification using the internal
// pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *EntityCorporationBeneficialOwnersIndividualIdentification) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityCorporationBeneficialOwnersIndividualIdentification
// into an array of bytes using the gjson library. Members of the `Extras` field
// are serialized into the top-level, and will overwrite known members of the same
// name.
func (r *EntityCorporationBeneficialOwnersIndividualIdentification) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityCorporationBeneficialOwnersIndividualIdentification) String() (result string) {
	return fmt.Sprintf("&EntityCorporationBeneficialOwnersIndividualIdentification{Method:%s NumberLast4:%s}", core.FmtP(r.Method), core.FmtP(r.NumberLast4))
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
	Name *string `pjson:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `pjson:"date_of_birth"`
	// The person's address.
	Address *EntityNaturalPersonAddress `pjson:"address"`
	// A means of verifying the person's identity.
	Identification *EntityNaturalPersonIdentification `pjson:"identification"`
	jsonFields     map[string]interface{}             `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EntityNaturalPerson using the
// internal pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *EntityNaturalPerson) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityNaturalPerson into an array of bytes using the
// gjson library. Members of the `Extras` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *EntityNaturalPerson) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityNaturalPerson) String() (result string) {
	return fmt.Sprintf("&EntityNaturalPerson{Name:%s DateOfBirth:%s Address:%s Identification:%s}", core.FmtP(r.Name), core.FmtP(r.DateOfBirth), r.Address, r.Identification)
}

type EntityNaturalPersonAddress struct {
	// The first line of the address.
	Line1 *string `pjson:"line1"`
	// The second line of the address.
	Line2 *string `pjson:"line2"`
	// The city of the address.
	City *string `pjson:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `pjson:"state"`
	// The ZIP code of the address.
	Zip        *string                `pjson:"zip"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EntityNaturalPersonAddress
// using the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *EntityNaturalPersonAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityNaturalPersonAddress into an array of bytes using
// the gjson library. Members of the `Extras` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EntityNaturalPersonAddress) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityNaturalPersonAddress) String() (result string) {
	return fmt.Sprintf("&EntityNaturalPersonAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", core.FmtP(r.Line1), core.FmtP(r.Line2), core.FmtP(r.City), core.FmtP(r.State), core.FmtP(r.Zip))
}

type EntityNaturalPersonIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method *EntityNaturalPersonIdentificationMethod `pjson:"method"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 *string                `pjson:"number_last4"`
	jsonFields  map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// EntityNaturalPersonIdentification using the internal pjson library. Unrecognized
// fields are stored in the `Extras` property.
func (r *EntityNaturalPersonIdentification) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityNaturalPersonIdentification into an array of bytes
// using the gjson library. Members of the `Extras` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EntityNaturalPersonIdentification) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityNaturalPersonIdentification) String() (result string) {
	return fmt.Sprintf("&EntityNaturalPersonIdentification{Method:%s NumberLast4:%s}", core.FmtP(r.Method), core.FmtP(r.NumberLast4))
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
	Name *string `pjson:"name"`
	// The two individuals that share control of the entity.
	Individuals *[]EntityJointIndividuals `pjson:"individuals"`
	jsonFields  map[string]interface{}    `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EntityJoint using the
// internal pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *EntityJoint) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityJoint into an array of bytes using the gjson
// library. Members of the `Extras` field are serialized into the top-level, and
// will overwrite known members of the same name.
func (r *EntityJoint) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityJoint) String() (result string) {
	return fmt.Sprintf("&EntityJoint{Name:%s Individuals:%s}", core.FmtP(r.Name), core.Fmt(r.Individuals))
}

type EntityJointIndividuals struct {
	// The person's legal name.
	Name *string `pjson:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `pjson:"date_of_birth"`
	// The person's address.
	Address *EntityJointIndividualsAddress `pjson:"address"`
	// A means of verifying the person's identity.
	Identification *EntityJointIndividualsIdentification `pjson:"identification"`
	jsonFields     map[string]interface{}                `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EntityJointIndividuals using
// the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *EntityJointIndividuals) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityJointIndividuals into an array of bytes using the
// gjson library. Members of the `Extras` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *EntityJointIndividuals) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityJointIndividuals) String() (result string) {
	return fmt.Sprintf("&EntityJointIndividuals{Name:%s DateOfBirth:%s Address:%s Identification:%s}", core.FmtP(r.Name), core.FmtP(r.DateOfBirth), r.Address, r.Identification)
}

type EntityJointIndividualsAddress struct {
	// The first line of the address.
	Line1 *string `pjson:"line1"`
	// The second line of the address.
	Line2 *string `pjson:"line2"`
	// The city of the address.
	City *string `pjson:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `pjson:"state"`
	// The ZIP code of the address.
	Zip        *string                `pjson:"zip"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EntityJointIndividualsAddress
// using the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *EntityJointIndividualsAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityJointIndividualsAddress into an array of bytes
// using the gjson library. Members of the `Extras` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EntityJointIndividualsAddress) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityJointIndividualsAddress) String() (result string) {
	return fmt.Sprintf("&EntityJointIndividualsAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", core.FmtP(r.Line1), core.FmtP(r.Line2), core.FmtP(r.City), core.FmtP(r.State), core.FmtP(r.Zip))
}

type EntityJointIndividualsIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method *EntityJointIndividualsIdentificationMethod `pjson:"method"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 *string                `pjson:"number_last4"`
	jsonFields  map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// EntityJointIndividualsIdentification using the internal pjson library.
// Unrecognized fields are stored in the `Extras` property.
func (r *EntityJointIndividualsIdentification) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityJointIndividualsIdentification into an array of
// bytes using the gjson library. Members of the `Extras` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *EntityJointIndividualsIdentification) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityJointIndividualsIdentification) String() (result string) {
	return fmt.Sprintf("&EntityJointIndividualsIdentification{Method:%s NumberLast4:%s}", core.FmtP(r.Method), core.FmtP(r.NumberLast4))
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
	Name *string `pjson:"name"`
	// Whether the trust is `revocable` or `irrevocable`.
	Category *EntityTrustCategory `pjson:"category"`
	// The trust's address.
	Address *EntityTrustAddress `pjson:"address"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state in
	// which the trust was formed.
	FormationState *string `pjson:"formation_state"`
	// The Employer Identification Number (EIN) of the trust itself.
	TaxIdentifier *string `pjson:"tax_identifier"`
	// The trustees of the trust.
	Trustees *[]EntityTrustTrustees `pjson:"trustees"`
	// The grantor of the trust. Will be present if the `category` is `revocable`.
	Grantor *EntityTrustGrantor `pjson:"grantor"`
	// The ID for the File containing the formation document of the trust.
	FormationDocumentFileID *string                `pjson:"formation_document_file_id"`
	jsonFields              map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EntityTrust using the
// internal pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *EntityTrust) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityTrust into an array of bytes using the gjson
// library. Members of the `Extras` field are serialized into the top-level, and
// will overwrite known members of the same name.
func (r *EntityTrust) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityTrust) String() (result string) {
	return fmt.Sprintf("&EntityTrust{Name:%s Category:%s Address:%s FormationState:%s TaxIdentifier:%s Trustees:%s Grantor:%s FormationDocumentFileID:%s}", core.FmtP(r.Name), core.FmtP(r.Category), r.Address, core.FmtP(r.FormationState), core.FmtP(r.TaxIdentifier), core.Fmt(r.Trustees), r.Grantor, core.FmtP(r.FormationDocumentFileID))
}

type EntityTrustCategory string

const (
	EntityTrustCategoryRevocable   EntityTrustCategory = "revocable"
	EntityTrustCategoryIrrevocable EntityTrustCategory = "irrevocable"
)

type EntityTrustAddress struct {
	// The first line of the address.
	Line1 *string `pjson:"line1"`
	// The second line of the address.
	Line2 *string `pjson:"line2"`
	// The city of the address.
	City *string `pjson:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `pjson:"state"`
	// The ZIP code of the address.
	Zip        *string                `pjson:"zip"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EntityTrustAddress using the
// internal pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *EntityTrustAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityTrustAddress into an array of bytes using the gjson
// library. Members of the `Extras` field are serialized into the top-level, and
// will overwrite known members of the same name.
func (r *EntityTrustAddress) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityTrustAddress) String() (result string) {
	return fmt.Sprintf("&EntityTrustAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", core.FmtP(r.Line1), core.FmtP(r.Line2), core.FmtP(r.City), core.FmtP(r.State), core.FmtP(r.Zip))
}

type EntityTrustTrustees struct {
	// The structure of the trustee. Will always be equal to `individual`.
	Structure *EntityTrustTrusteesStructure `pjson:"structure"`
	// The individual trustee of the trust. Will be present if the trustee's
	// `structure` is equal to `individual`.
	Individual *EntityTrustTrusteesIndividual `pjson:"individual"`
	jsonFields map[string]interface{}         `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EntityTrustTrustees using the
// internal pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *EntityTrustTrustees) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityTrustTrustees into an array of bytes using the
// gjson library. Members of the `Extras` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *EntityTrustTrustees) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityTrustTrustees) String() (result string) {
	return fmt.Sprintf("&EntityTrustTrustees{Structure:%s Individual:%s}", core.FmtP(r.Structure), r.Individual)
}

type EntityTrustTrusteesStructure string

const (
	EntityTrustTrusteesStructureIndividual EntityTrustTrusteesStructure = "individual"
)

type EntityTrustTrusteesIndividual struct {
	// The person's legal name.
	Name *string `pjson:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `pjson:"date_of_birth"`
	// The person's address.
	Address *EntityTrustTrusteesIndividualAddress `pjson:"address"`
	// A means of verifying the person's identity.
	Identification *EntityTrustTrusteesIndividualIdentification `pjson:"identification"`
	jsonFields     map[string]interface{}                       `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EntityTrustTrusteesIndividual
// using the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *EntityTrustTrusteesIndividual) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityTrustTrusteesIndividual into an array of bytes
// using the gjson library. Members of the `Extras` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EntityTrustTrusteesIndividual) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityTrustTrusteesIndividual) String() (result string) {
	return fmt.Sprintf("&EntityTrustTrusteesIndividual{Name:%s DateOfBirth:%s Address:%s Identification:%s}", core.FmtP(r.Name), core.FmtP(r.DateOfBirth), r.Address, r.Identification)
}

type EntityTrustTrusteesIndividualAddress struct {
	// The first line of the address.
	Line1 *string `pjson:"line1"`
	// The second line of the address.
	Line2 *string `pjson:"line2"`
	// The city of the address.
	City *string `pjson:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `pjson:"state"`
	// The ZIP code of the address.
	Zip        *string                `pjson:"zip"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// EntityTrustTrusteesIndividualAddress using the internal pjson library.
// Unrecognized fields are stored in the `Extras` property.
func (r *EntityTrustTrusteesIndividualAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityTrustTrusteesIndividualAddress into an array of
// bytes using the gjson library. Members of the `Extras` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *EntityTrustTrusteesIndividualAddress) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityTrustTrusteesIndividualAddress) String() (result string) {
	return fmt.Sprintf("&EntityTrustTrusteesIndividualAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", core.FmtP(r.Line1), core.FmtP(r.Line2), core.FmtP(r.City), core.FmtP(r.State), core.FmtP(r.Zip))
}

type EntityTrustTrusteesIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method *EntityTrustTrusteesIndividualIdentificationMethod `pjson:"method"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 *string                `pjson:"number_last4"`
	jsonFields  map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// EntityTrustTrusteesIndividualIdentification using the internal pjson library.
// Unrecognized fields are stored in the `Extras` property.
func (r *EntityTrustTrusteesIndividualIdentification) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityTrustTrusteesIndividualIdentification into an array
// of bytes using the gjson library. Members of the `Extras` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *EntityTrustTrusteesIndividualIdentification) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityTrustTrusteesIndividualIdentification) String() (result string) {
	return fmt.Sprintf("&EntityTrustTrusteesIndividualIdentification{Method:%s NumberLast4:%s}", core.FmtP(r.Method), core.FmtP(r.NumberLast4))
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
	Name *string `pjson:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `pjson:"date_of_birth"`
	// The person's address.
	Address *EntityTrustGrantorAddress `pjson:"address"`
	// A means of verifying the person's identity.
	Identification *EntityTrustGrantorIdentification `pjson:"identification"`
	jsonFields     map[string]interface{}            `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EntityTrustGrantor using the
// internal pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *EntityTrustGrantor) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityTrustGrantor into an array of bytes using the gjson
// library. Members of the `Extras` field are serialized into the top-level, and
// will overwrite known members of the same name.
func (r *EntityTrustGrantor) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityTrustGrantor) String() (result string) {
	return fmt.Sprintf("&EntityTrustGrantor{Name:%s DateOfBirth:%s Address:%s Identification:%s}", core.FmtP(r.Name), core.FmtP(r.DateOfBirth), r.Address, r.Identification)
}

type EntityTrustGrantorAddress struct {
	// The first line of the address.
	Line1 *string `pjson:"line1"`
	// The second line of the address.
	Line2 *string `pjson:"line2"`
	// The city of the address.
	City *string `pjson:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `pjson:"state"`
	// The ZIP code of the address.
	Zip        *string                `pjson:"zip"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EntityTrustGrantorAddress
// using the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *EntityTrustGrantorAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityTrustGrantorAddress into an array of bytes using
// the gjson library. Members of the `Extras` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EntityTrustGrantorAddress) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityTrustGrantorAddress) String() (result string) {
	return fmt.Sprintf("&EntityTrustGrantorAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", core.FmtP(r.Line1), core.FmtP(r.Line2), core.FmtP(r.City), core.FmtP(r.State), core.FmtP(r.Zip))
}

type EntityTrustGrantorIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method *EntityTrustGrantorIdentificationMethod `pjson:"method"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 *string                `pjson:"number_last4"`
	jsonFields  map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// EntityTrustGrantorIdentification using the internal pjson library. Unrecognized
// fields are stored in the `Extras` property.
func (r *EntityTrustGrantorIdentification) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityTrustGrantorIdentification into an array of bytes
// using the gjson library. Members of the `Extras` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EntityTrustGrantorIdentification) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityTrustGrantorIdentification) String() (result string) {
	return fmt.Sprintf("&EntityTrustGrantorIdentification{Method:%s NumberLast4:%s}", core.FmtP(r.Method), core.FmtP(r.NumberLast4))
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
	FileID     *string                `pjson:"file_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EntitySupplementalDocuments
// using the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *EntitySupplementalDocuments) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntitySupplementalDocuments into an array of bytes using
// the gjson library. Members of the `Extras` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EntitySupplementalDocuments) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The File containing the document.
func (r *EntitySupplementalDocuments) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

func (r EntitySupplementalDocuments) String() (result string) {
	return fmt.Sprintf("&EntitySupplementalDocuments{FileID:%s}", core.FmtP(r.FileID))
}

type CreateAnEntityParameters struct {
	// The type of Entity to create.
	Structure *CreateAnEntityParametersStructure `pjson:"structure"`
	// Details of the corporation entity to create. Required if `structure` is equal to
	// `corporation`.
	Corporation *CreateAnEntityParametersCorporation `pjson:"corporation"`
	// Details of the natural person entity to create. Required if `structure` is equal
	// to `natural_person`. Natural people entities should be submitted with
	// `social_security_number` or `individual_taxpayer_identification_number`
	// identification methods.
	NaturalPerson *CreateAnEntityParametersNaturalPerson `pjson:"natural_person"`
	// Details of the joint entity to create. Required if `structure` is equal to
	// `joint`.
	Joint *CreateAnEntityParametersJoint `pjson:"joint"`
	// Details of the trust entity to create. Required if `structure` is equal to
	// `trust`.
	Trust *CreateAnEntityParametersTrust `pjson:"trust"`
	// The description you choose to give the entity.
	Description *string `pjson:"description"`
	// The relationship between your group and the entity.
	Relationship *CreateAnEntityParametersRelationship `pjson:"relationship"`
	// Additional documentation associated with the entity.
	SupplementalDocuments *[]CreateAnEntityParametersSupplementalDocuments `pjson:"supplemental_documents"`
	jsonFields            map[string]interface{}                           `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CreateAnEntityParameters
// using the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *CreateAnEntityParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParameters into an array of bytes using the
// gjson library. Members of the `Extras` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CreateAnEntityParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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
// to `natural_person`. Natural people entities should be submitted with
// `social_security_number` or `individual_taxpayer_identification_number`
// identification methods.
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

func (r CreateAnEntityParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParameters{Structure:%s Corporation:%s NaturalPerson:%s Joint:%s Trust:%s Description:%s Relationship:%s SupplementalDocuments:%s}", core.FmtP(r.Structure), r.Corporation, r.NaturalPerson, r.Joint, r.Trust, core.FmtP(r.Description), core.FmtP(r.Relationship), core.Fmt(r.SupplementalDocuments))
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
	Name *string `pjson:"name"`
	// The website of the corporation.
	Website *string `pjson:"website"`
	// The Employer Identification Number (EIN) for the corporation.
	TaxIdentifier *string `pjson:"tax_identifier"`
	// The two-letter United States Postal Service (USPS) abbreviation for the
	// corporation's state of incorporation.
	IncorporationState *string `pjson:"incorporation_state"`
	// The corporation's address.
	Address *CreateAnEntityParametersCorporationAddress `pjson:"address"`
	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwners *[]CreateAnEntityParametersCorporationBeneficialOwners `pjson:"beneficial_owners"`
	jsonFields       map[string]interface{}                                 `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersCorporation using the internal pjson library.
// Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersCorporation) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersCorporation into an array of
// bytes using the gjson library. Members of the `Extras` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnEntityParametersCorporation) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r CreateAnEntityParametersCorporation) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersCorporation{Name:%s Website:%s TaxIdentifier:%s IncorporationState:%s Address:%s BeneficialOwners:%s}", core.FmtP(r.Name), core.FmtP(r.Website), core.FmtP(r.TaxIdentifier), core.FmtP(r.IncorporationState), r.Address, core.Fmt(r.BeneficialOwners))
}

type CreateAnEntityParametersCorporationAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 *string `pjson:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 *string `pjson:"line2"`
	// The city of the address.
	City *string `pjson:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `pjson:"state"`
	// The ZIP code of the address.
	Zip        *string                `pjson:"zip"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersCorporationAddress using the internal pjson library.
// Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersCorporationAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersCorporationAddress into an array
// of bytes using the gjson library. Members of the `Extras` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *CreateAnEntityParametersCorporationAddress) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r CreateAnEntityParametersCorporationAddress) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersCorporationAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", core.FmtP(r.Line1), core.FmtP(r.Line2), core.FmtP(r.City), core.FmtP(r.State), core.FmtP(r.Zip))
}

type CreateAnEntityParametersCorporationBeneficialOwners struct {
	// Personal details for the beneficial owner.
	Individual *CreateAnEntityParametersCorporationBeneficialOwnersIndividual `pjson:"individual"`
	// This person's role or title within the entity.
	CompanyTitle *string `pjson:"company_title"`
	// Why this person is considered a beneficial owner of the entity.
	Prong      *CreateAnEntityParametersCorporationBeneficialOwnersProng `pjson:"prong"`
	jsonFields map[string]interface{}                                    `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersCorporationBeneficialOwners using the internal pjson
// library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersCorporationBeneficialOwners) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersCorporationBeneficialOwners into
// an array of bytes using the gjson library. Members of the `Extras` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateAnEntityParametersCorporationBeneficialOwners) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r CreateAnEntityParametersCorporationBeneficialOwners) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersCorporationBeneficialOwners{Individual:%s CompanyTitle:%s Prong:%s}", r.Individual, core.FmtP(r.CompanyTitle), core.FmtP(r.Prong))
}

type CreateAnEntityParametersCorporationBeneficialOwnersIndividual struct {
	// The person's legal name.
	Name *string `pjson:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `pjson:"date_of_birth"`
	// The individual's address.
	Address *CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress `pjson:"address"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID *bool `pjson:"confirmed_no_us_tax_id"`
	// A means of verifying the person's identity.
	Identification *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification `pjson:"identification"`
	jsonFields     map[string]interface{}                                                       `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersCorporationBeneficialOwnersIndividual using the internal
// pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividual) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CreateAnEntityParametersCorporationBeneficialOwnersIndividual into an array of
// bytes using the gjson library. Members of the `Extras` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividual) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

// The identification method for an individual can only be a passport, driver's
// license, or other document if you've confirmed the individual does not have a US
// tax id (either a Social Security Number or Individual Taxpayer Identification
// Number).
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividual) GetConfirmedNoUsTaxID() (ConfirmedNoUsTaxID bool) {
	if r != nil && r.ConfirmedNoUsTaxID != nil {
		ConfirmedNoUsTaxID = *r.ConfirmedNoUsTaxID
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

func (r CreateAnEntityParametersCorporationBeneficialOwnersIndividual) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersCorporationBeneficialOwnersIndividual{Name:%s DateOfBirth:%s Address:%s ConfirmedNoUsTaxID:%s Identification:%s}", core.FmtP(r.Name), core.FmtP(r.DateOfBirth), r.Address, core.FmtP(r.ConfirmedNoUsTaxID), r.Identification)
}

type CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 *string `pjson:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 *string `pjson:"line2"`
	// The city of the address.
	City *string `pjson:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `pjson:"state"`
	// The ZIP code of the address.
	Zip        *string                `pjson:"zip"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress using the
// internal pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress into an
// array of bytes using the gjson library. Members of the `Extras` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", core.FmtP(r.Line1), core.FmtP(r.Line2), core.FmtP(r.City), core.FmtP(r.State), core.FmtP(r.Zip))
}

type CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethod `pjson:"method"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number *string `pjson:"number"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport `pjson:"passport"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationDriversLicense `pjson:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other      *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationOther `pjson:"other"`
	jsonFields map[string]interface{}                                                            `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification
// using the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification into
// an array of bytes using the gjson library. Members of the `Extras` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

// Information about the United States driver's license used for identification.
// Required if `method` is equal to `drivers_license`.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification) GetDriversLicense() (DriversLicense CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationDriversLicense) {
	if r != nil && r.DriversLicense != nil {
		DriversLicense = *r.DriversLicense
	}
	return
}

// Information about the identification document provided. Required if `method` is
// equal to `other`.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification) GetOther() (Other CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationOther) {
	if r != nil && r.Other != nil {
		Other = *r.Other
	}
	return
}

func (r CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification{Method:%s Number:%s Passport:%s DriversLicense:%s Other:%s}", core.FmtP(r.Method), core.FmtP(r.Number), r.Passport, r.DriversLicense, r.Other)
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
	FileID *string `pjson:"file_id"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `pjson:"expiration_date"`
	// The country that issued the passport.
	Country    *string                `pjson:"country"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport
// using the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport
// into an array of bytes using the gjson library. Members of the `Extras` field
// are serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport{FileID:%s ExpirationDate:%s Country:%s}", core.FmtP(r.FileID), core.FmtP(r.ExpirationDate), core.FmtP(r.Country))
}

type CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationDriversLicense struct {
	// The identifier of the File containing the driver's license.
	FileID *string `pjson:"file_id"`
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `pjson:"expiration_date"`
	// The state that issued the provided driver's license.
	State      *string                `pjson:"state"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationDriversLicense
// using the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationDriversLicense) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationDriversLicense
// into an array of bytes using the gjson library. Members of the `Extras` field
// are serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationDriversLicense) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the File containing the driver's license.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationDriversLicense) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

// The driver's license's expiration date in YYYY-MM-DD format.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationDriversLicense) GetExpirationDate() (ExpirationDate string) {
	if r != nil && r.ExpirationDate != nil {
		ExpirationDate = *r.ExpirationDate
	}
	return
}

// The state that issued the provided driver's license.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationDriversLicense) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

func (r CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationDriversLicense) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationDriversLicense{FileID:%s ExpirationDate:%s State:%s}", core.FmtP(r.FileID), core.FmtP(r.ExpirationDate), core.FmtP(r.State))
}

type CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document.
	Country *string `pjson:"country"`
	// A description of the document submitted.
	Description *string `pjson:"description"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `pjson:"expiration_date"`
	// The identifier of the File containing the document.
	FileID     *string                `pjson:"file_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationOther
// using the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationOther) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationOther
// into an array of bytes using the gjson library. Members of the `Extras` field
// are serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationOther) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The two-character ISO 3166-1 code representing the country that issued the
// document.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationOther) GetCountry() (Country string) {
	if r != nil && r.Country != nil {
		Country = *r.Country
	}
	return
}

// A description of the document submitted.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationOther) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The document's expiration date in YYYY-MM-DD format.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationOther) GetExpirationDate() (ExpirationDate string) {
	if r != nil && r.ExpirationDate != nil {
		ExpirationDate = *r.ExpirationDate
	}
	return
}

// The identifier of the File containing the document.
func (r *CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationOther) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

func (r CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationOther) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationOther{Country:%s Description:%s ExpirationDate:%s FileID:%s}", core.FmtP(r.Country), core.FmtP(r.Description), core.FmtP(r.ExpirationDate), core.FmtP(r.FileID))
}

type CreateAnEntityParametersCorporationBeneficialOwnersProng string

const (
	CreateAnEntityParametersCorporationBeneficialOwnersProngOwnership CreateAnEntityParametersCorporationBeneficialOwnersProng = "ownership"
	CreateAnEntityParametersCorporationBeneficialOwnersProngControl   CreateAnEntityParametersCorporationBeneficialOwnersProng = "control"
)

type CreateAnEntityParametersNaturalPerson struct {
	// The person's legal name.
	Name *string `pjson:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `pjson:"date_of_birth"`
	// The individual's address.
	Address *CreateAnEntityParametersNaturalPersonAddress `pjson:"address"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID *bool `pjson:"confirmed_no_us_tax_id"`
	// A means of verifying the person's identity.
	Identification *CreateAnEntityParametersNaturalPersonIdentification `pjson:"identification"`
	jsonFields     map[string]interface{}                               `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersNaturalPerson using the internal pjson library.
// Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersNaturalPerson) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersNaturalPerson into an array of
// bytes using the gjson library. Members of the `Extras` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnEntityParametersNaturalPerson) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

// The identification method for an individual can only be a passport, driver's
// license, or other document if you've confirmed the individual does not have a US
// tax id (either a Social Security Number or Individual Taxpayer Identification
// Number).
func (r *CreateAnEntityParametersNaturalPerson) GetConfirmedNoUsTaxID() (ConfirmedNoUsTaxID bool) {
	if r != nil && r.ConfirmedNoUsTaxID != nil {
		ConfirmedNoUsTaxID = *r.ConfirmedNoUsTaxID
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

func (r CreateAnEntityParametersNaturalPerson) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersNaturalPerson{Name:%s DateOfBirth:%s Address:%s ConfirmedNoUsTaxID:%s Identification:%s}", core.FmtP(r.Name), core.FmtP(r.DateOfBirth), r.Address, core.FmtP(r.ConfirmedNoUsTaxID), r.Identification)
}

type CreateAnEntityParametersNaturalPersonAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 *string `pjson:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 *string `pjson:"line2"`
	// The city of the address.
	City *string `pjson:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `pjson:"state"`
	// The ZIP code of the address.
	Zip        *string                `pjson:"zip"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersNaturalPersonAddress using the internal pjson library.
// Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersNaturalPersonAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersNaturalPersonAddress into an
// array of bytes using the gjson library. Members of the `Extras` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateAnEntityParametersNaturalPersonAddress) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r CreateAnEntityParametersNaturalPersonAddress) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersNaturalPersonAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", core.FmtP(r.Line1), core.FmtP(r.Line2), core.FmtP(r.City), core.FmtP(r.State), core.FmtP(r.Zip))
}

type CreateAnEntityParametersNaturalPersonIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method *CreateAnEntityParametersNaturalPersonIdentificationMethod `pjson:"method"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number *string `pjson:"number"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport *CreateAnEntityParametersNaturalPersonIdentificationPassport `pjson:"passport"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense *CreateAnEntityParametersNaturalPersonIdentificationDriversLicense `pjson:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other      *CreateAnEntityParametersNaturalPersonIdentificationOther `pjson:"other"`
	jsonFields map[string]interface{}                                    `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersNaturalPersonIdentification using the internal pjson
// library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersNaturalPersonIdentification) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersNaturalPersonIdentification into
// an array of bytes using the gjson library. Members of the `Extras` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateAnEntityParametersNaturalPersonIdentification) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

// Information about the United States driver's license used for identification.
// Required if `method` is equal to `drivers_license`.
func (r *CreateAnEntityParametersNaturalPersonIdentification) GetDriversLicense() (DriversLicense CreateAnEntityParametersNaturalPersonIdentificationDriversLicense) {
	if r != nil && r.DriversLicense != nil {
		DriversLicense = *r.DriversLicense
	}
	return
}

// Information about the identification document provided. Required if `method` is
// equal to `other`.
func (r *CreateAnEntityParametersNaturalPersonIdentification) GetOther() (Other CreateAnEntityParametersNaturalPersonIdentificationOther) {
	if r != nil && r.Other != nil {
		Other = *r.Other
	}
	return
}

func (r CreateAnEntityParametersNaturalPersonIdentification) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersNaturalPersonIdentification{Method:%s Number:%s Passport:%s DriversLicense:%s Other:%s}", core.FmtP(r.Method), core.FmtP(r.Number), r.Passport, r.DriversLicense, r.Other)
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
	FileID *string `pjson:"file_id"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `pjson:"expiration_date"`
	// The country that issued the passport.
	Country    *string                `pjson:"country"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersNaturalPersonIdentificationPassport using the internal
// pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersNaturalPersonIdentificationPassport) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CreateAnEntityParametersNaturalPersonIdentificationPassport into an array of
// bytes using the gjson library. Members of the `Extras` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnEntityParametersNaturalPersonIdentificationPassport) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r CreateAnEntityParametersNaturalPersonIdentificationPassport) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersNaturalPersonIdentificationPassport{FileID:%s ExpirationDate:%s Country:%s}", core.FmtP(r.FileID), core.FmtP(r.ExpirationDate), core.FmtP(r.Country))
}

type CreateAnEntityParametersNaturalPersonIdentificationDriversLicense struct {
	// The identifier of the File containing the driver's license.
	FileID *string `pjson:"file_id"`
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `pjson:"expiration_date"`
	// The state that issued the provided driver's license.
	State      *string                `pjson:"state"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersNaturalPersonIdentificationDriversLicense using the
// internal pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersNaturalPersonIdentificationDriversLicense) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CreateAnEntityParametersNaturalPersonIdentificationDriversLicense into an array
// of bytes using the gjson library. Members of the `Extras` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *CreateAnEntityParametersNaturalPersonIdentificationDriversLicense) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the File containing the driver's license.
func (r *CreateAnEntityParametersNaturalPersonIdentificationDriversLicense) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

// The driver's license's expiration date in YYYY-MM-DD format.
func (r *CreateAnEntityParametersNaturalPersonIdentificationDriversLicense) GetExpirationDate() (ExpirationDate string) {
	if r != nil && r.ExpirationDate != nil {
		ExpirationDate = *r.ExpirationDate
	}
	return
}

// The state that issued the provided driver's license.
func (r *CreateAnEntityParametersNaturalPersonIdentificationDriversLicense) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

func (r CreateAnEntityParametersNaturalPersonIdentificationDriversLicense) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersNaturalPersonIdentificationDriversLicense{FileID:%s ExpirationDate:%s State:%s}", core.FmtP(r.FileID), core.FmtP(r.ExpirationDate), core.FmtP(r.State))
}

type CreateAnEntityParametersNaturalPersonIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document.
	Country *string `pjson:"country"`
	// A description of the document submitted.
	Description *string `pjson:"description"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `pjson:"expiration_date"`
	// The identifier of the File containing the document.
	FileID     *string                `pjson:"file_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersNaturalPersonIdentificationOther using the internal
// pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersNaturalPersonIdentificationOther) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersNaturalPersonIdentificationOther
// into an array of bytes using the gjson library. Members of the `Extras` field
// are serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateAnEntityParametersNaturalPersonIdentificationOther) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The two-character ISO 3166-1 code representing the country that issued the
// document.
func (r *CreateAnEntityParametersNaturalPersonIdentificationOther) GetCountry() (Country string) {
	if r != nil && r.Country != nil {
		Country = *r.Country
	}
	return
}

// A description of the document submitted.
func (r *CreateAnEntityParametersNaturalPersonIdentificationOther) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The document's expiration date in YYYY-MM-DD format.
func (r *CreateAnEntityParametersNaturalPersonIdentificationOther) GetExpirationDate() (ExpirationDate string) {
	if r != nil && r.ExpirationDate != nil {
		ExpirationDate = *r.ExpirationDate
	}
	return
}

// The identifier of the File containing the document.
func (r *CreateAnEntityParametersNaturalPersonIdentificationOther) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

func (r CreateAnEntityParametersNaturalPersonIdentificationOther) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersNaturalPersonIdentificationOther{Country:%s Description:%s ExpirationDate:%s FileID:%s}", core.FmtP(r.Country), core.FmtP(r.Description), core.FmtP(r.ExpirationDate), core.FmtP(r.FileID))
}

type CreateAnEntityParametersJoint struct {
	// The name of the joint entity.
	Name *string `pjson:"name"`
	// The two individuals that share control of the entity.
	Individuals *[]CreateAnEntityParametersJointIndividuals `pjson:"individuals"`
	jsonFields  map[string]interface{}                      `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CreateAnEntityParametersJoint
// using the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *CreateAnEntityParametersJoint) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersJoint into an array of bytes
// using the gjson library. Members of the `Extras` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CreateAnEntityParametersJoint) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r CreateAnEntityParametersJoint) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersJoint{Name:%s Individuals:%s}", core.FmtP(r.Name), core.Fmt(r.Individuals))
}

type CreateAnEntityParametersJointIndividuals struct {
	// The person's legal name.
	Name *string `pjson:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `pjson:"date_of_birth"`
	// The individual's address.
	Address *CreateAnEntityParametersJointIndividualsAddress `pjson:"address"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID *bool `pjson:"confirmed_no_us_tax_id"`
	// A means of verifying the person's identity.
	Identification *CreateAnEntityParametersJointIndividualsIdentification `pjson:"identification"`
	jsonFields     map[string]interface{}                                  `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersJointIndividuals using the internal pjson library.
// Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersJointIndividuals) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersJointIndividuals into an array of
// bytes using the gjson library. Members of the `Extras` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnEntityParametersJointIndividuals) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

// The identification method for an individual can only be a passport, driver's
// license, or other document if you've confirmed the individual does not have a US
// tax id (either a Social Security Number or Individual Taxpayer Identification
// Number).
func (r *CreateAnEntityParametersJointIndividuals) GetConfirmedNoUsTaxID() (ConfirmedNoUsTaxID bool) {
	if r != nil && r.ConfirmedNoUsTaxID != nil {
		ConfirmedNoUsTaxID = *r.ConfirmedNoUsTaxID
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

func (r CreateAnEntityParametersJointIndividuals) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersJointIndividuals{Name:%s DateOfBirth:%s Address:%s ConfirmedNoUsTaxID:%s Identification:%s}", core.FmtP(r.Name), core.FmtP(r.DateOfBirth), r.Address, core.FmtP(r.ConfirmedNoUsTaxID), r.Identification)
}

type CreateAnEntityParametersJointIndividualsAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 *string `pjson:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 *string `pjson:"line2"`
	// The city of the address.
	City *string `pjson:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `pjson:"state"`
	// The ZIP code of the address.
	Zip        *string                `pjson:"zip"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersJointIndividualsAddress using the internal pjson
// library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersJointIndividualsAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersJointIndividualsAddress into an
// array of bytes using the gjson library. Members of the `Extras` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateAnEntityParametersJointIndividualsAddress) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r CreateAnEntityParametersJointIndividualsAddress) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersJointIndividualsAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", core.FmtP(r.Line1), core.FmtP(r.Line2), core.FmtP(r.City), core.FmtP(r.State), core.FmtP(r.Zip))
}

type CreateAnEntityParametersJointIndividualsIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method *CreateAnEntityParametersJointIndividualsIdentificationMethod `pjson:"method"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number *string `pjson:"number"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport *CreateAnEntityParametersJointIndividualsIdentificationPassport `pjson:"passport"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense *CreateAnEntityParametersJointIndividualsIdentificationDriversLicense `pjson:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other      *CreateAnEntityParametersJointIndividualsIdentificationOther `pjson:"other"`
	jsonFields map[string]interface{}                                       `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersJointIndividualsIdentification using the internal pjson
// library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersJointIndividualsIdentification) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersJointIndividualsIdentification
// into an array of bytes using the gjson library. Members of the `Extras` field
// are serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateAnEntityParametersJointIndividualsIdentification) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

// Information about the United States driver's license used for identification.
// Required if `method` is equal to `drivers_license`.
func (r *CreateAnEntityParametersJointIndividualsIdentification) GetDriversLicense() (DriversLicense CreateAnEntityParametersJointIndividualsIdentificationDriversLicense) {
	if r != nil && r.DriversLicense != nil {
		DriversLicense = *r.DriversLicense
	}
	return
}

// Information about the identification document provided. Required if `method` is
// equal to `other`.
func (r *CreateAnEntityParametersJointIndividualsIdentification) GetOther() (Other CreateAnEntityParametersJointIndividualsIdentificationOther) {
	if r != nil && r.Other != nil {
		Other = *r.Other
	}
	return
}

func (r CreateAnEntityParametersJointIndividualsIdentification) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersJointIndividualsIdentification{Method:%s Number:%s Passport:%s DriversLicense:%s Other:%s}", core.FmtP(r.Method), core.FmtP(r.Number), r.Passport, r.DriversLicense, r.Other)
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
	FileID *string `pjson:"file_id"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `pjson:"expiration_date"`
	// The country that issued the passport.
	Country    *string                `pjson:"country"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersJointIndividualsIdentificationPassport using the
// internal pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersJointIndividualsIdentificationPassport) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CreateAnEntityParametersJointIndividualsIdentificationPassport into an array of
// bytes using the gjson library. Members of the `Extras` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnEntityParametersJointIndividualsIdentificationPassport) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r CreateAnEntityParametersJointIndividualsIdentificationPassport) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersJointIndividualsIdentificationPassport{FileID:%s ExpirationDate:%s Country:%s}", core.FmtP(r.FileID), core.FmtP(r.ExpirationDate), core.FmtP(r.Country))
}

type CreateAnEntityParametersJointIndividualsIdentificationDriversLicense struct {
	// The identifier of the File containing the driver's license.
	FileID *string `pjson:"file_id"`
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `pjson:"expiration_date"`
	// The state that issued the provided driver's license.
	State      *string                `pjson:"state"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersJointIndividualsIdentificationDriversLicense using the
// internal pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersJointIndividualsIdentificationDriversLicense) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CreateAnEntityParametersJointIndividualsIdentificationDriversLicense into an
// array of bytes using the gjson library. Members of the `Extras` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateAnEntityParametersJointIndividualsIdentificationDriversLicense) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the File containing the driver's license.
func (r *CreateAnEntityParametersJointIndividualsIdentificationDriversLicense) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

// The driver's license's expiration date in YYYY-MM-DD format.
func (r *CreateAnEntityParametersJointIndividualsIdentificationDriversLicense) GetExpirationDate() (ExpirationDate string) {
	if r != nil && r.ExpirationDate != nil {
		ExpirationDate = *r.ExpirationDate
	}
	return
}

// The state that issued the provided driver's license.
func (r *CreateAnEntityParametersJointIndividualsIdentificationDriversLicense) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

func (r CreateAnEntityParametersJointIndividualsIdentificationDriversLicense) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersJointIndividualsIdentificationDriversLicense{FileID:%s ExpirationDate:%s State:%s}", core.FmtP(r.FileID), core.FmtP(r.ExpirationDate), core.FmtP(r.State))
}

type CreateAnEntityParametersJointIndividualsIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document.
	Country *string `pjson:"country"`
	// A description of the document submitted.
	Description *string `pjson:"description"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `pjson:"expiration_date"`
	// The identifier of the File containing the document.
	FileID     *string                `pjson:"file_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersJointIndividualsIdentificationOther using the internal
// pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersJointIndividualsIdentificationOther) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CreateAnEntityParametersJointIndividualsIdentificationOther into an array of
// bytes using the gjson library. Members of the `Extras` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnEntityParametersJointIndividualsIdentificationOther) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The two-character ISO 3166-1 code representing the country that issued the
// document.
func (r *CreateAnEntityParametersJointIndividualsIdentificationOther) GetCountry() (Country string) {
	if r != nil && r.Country != nil {
		Country = *r.Country
	}
	return
}

// A description of the document submitted.
func (r *CreateAnEntityParametersJointIndividualsIdentificationOther) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The document's expiration date in YYYY-MM-DD format.
func (r *CreateAnEntityParametersJointIndividualsIdentificationOther) GetExpirationDate() (ExpirationDate string) {
	if r != nil && r.ExpirationDate != nil {
		ExpirationDate = *r.ExpirationDate
	}
	return
}

// The identifier of the File containing the document.
func (r *CreateAnEntityParametersJointIndividualsIdentificationOther) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

func (r CreateAnEntityParametersJointIndividualsIdentificationOther) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersJointIndividualsIdentificationOther{Country:%s Description:%s ExpirationDate:%s FileID:%s}", core.FmtP(r.Country), core.FmtP(r.Description), core.FmtP(r.ExpirationDate), core.FmtP(r.FileID))
}

type CreateAnEntityParametersTrust struct {
	// The legal name of the trust.
	Name *string `pjson:"name"`
	// Whether the trust is `revocable` or `irrevocable`. Irrevocable trusts require
	// their own Employer Identification Number. Revocable trusts require information
	// about the individual `grantor` who created the trust.
	Category *CreateAnEntityParametersTrustCategory `pjson:"category"`
	// The Employer Identification Number (EIN) for the trust. Required if `category`
	// is equal to `irrevocable`.
	TaxIdentifier *string `pjson:"tax_identifier"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state in
	// which the trust was formed.
	FormationState *string `pjson:"formation_state"`
	// The trust's address.
	Address *CreateAnEntityParametersTrustAddress `pjson:"address"`
	// The identifier of the File containing the formation document of the trust.
	FormationDocumentFileID *string `pjson:"formation_document_file_id"`
	// The trustees of the trust.
	Trustees *[]CreateAnEntityParametersTrustTrustees `pjson:"trustees"`
	// The grantor of the trust. Required if `category` is equal to `revocable`.
	Grantor    *CreateAnEntityParametersTrustGrantor `pjson:"grantor"`
	jsonFields map[string]interface{}                `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CreateAnEntityParametersTrust
// using the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *CreateAnEntityParametersTrust) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersTrust into an array of bytes
// using the gjson library. Members of the `Extras` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CreateAnEntityParametersTrust) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r CreateAnEntityParametersTrust) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrust{Name:%s Category:%s TaxIdentifier:%s FormationState:%s Address:%s FormationDocumentFileID:%s Trustees:%s Grantor:%s}", core.FmtP(r.Name), core.FmtP(r.Category), core.FmtP(r.TaxIdentifier), core.FmtP(r.FormationState), r.Address, core.FmtP(r.FormationDocumentFileID), core.Fmt(r.Trustees), r.Grantor)
}

type CreateAnEntityParametersTrustCategory string

const (
	CreateAnEntityParametersTrustCategoryRevocable   CreateAnEntityParametersTrustCategory = "revocable"
	CreateAnEntityParametersTrustCategoryIrrevocable CreateAnEntityParametersTrustCategory = "irrevocable"
)

type CreateAnEntityParametersTrustAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 *string `pjson:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 *string `pjson:"line2"`
	// The city of the address.
	City *string `pjson:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `pjson:"state"`
	// The ZIP code of the address.
	Zip        *string                `pjson:"zip"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersTrustAddress using the internal pjson library.
// Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersTrustAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersTrustAddress into an array of
// bytes using the gjson library. Members of the `Extras` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnEntityParametersTrustAddress) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r CreateAnEntityParametersTrustAddress) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", core.FmtP(r.Line1), core.FmtP(r.Line2), core.FmtP(r.City), core.FmtP(r.State), core.FmtP(r.Zip))
}

type CreateAnEntityParametersTrustTrustees struct {
	// The structure of the trustee.
	Structure *CreateAnEntityParametersTrustTrusteesStructure `pjson:"structure"`
	// Details of the individual trustee. Required when the trustee `structure` is
	// equal to `individual`.
	Individual *CreateAnEntityParametersTrustTrusteesIndividual `pjson:"individual"`
	jsonFields map[string]interface{}                           `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersTrustTrustees using the internal pjson library.
// Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersTrustTrustees) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersTrustTrustees into an array of
// bytes using the gjson library. Members of the `Extras` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnEntityParametersTrustTrustees) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r CreateAnEntityParametersTrustTrustees) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustTrustees{Structure:%s Individual:%s}", core.FmtP(r.Structure), r.Individual)
}

type CreateAnEntityParametersTrustTrusteesStructure string

const (
	CreateAnEntityParametersTrustTrusteesStructureIndividual CreateAnEntityParametersTrustTrusteesStructure = "individual"
)

type CreateAnEntityParametersTrustTrusteesIndividual struct {
	// The person's legal name.
	Name *string `pjson:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `pjson:"date_of_birth"`
	// The individual's address.
	Address *CreateAnEntityParametersTrustTrusteesIndividualAddress `pjson:"address"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID *bool `pjson:"confirmed_no_us_tax_id"`
	// A means of verifying the person's identity.
	Identification *CreateAnEntityParametersTrustTrusteesIndividualIdentification `pjson:"identification"`
	jsonFields     map[string]interface{}                                         `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersTrustTrusteesIndividual using the internal pjson
// library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersTrustTrusteesIndividual) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersTrustTrusteesIndividual into an
// array of bytes using the gjson library. Members of the `Extras` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateAnEntityParametersTrustTrusteesIndividual) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

// The identification method for an individual can only be a passport, driver's
// license, or other document if you've confirmed the individual does not have a US
// tax id (either a Social Security Number or Individual Taxpayer Identification
// Number).
func (r *CreateAnEntityParametersTrustTrusteesIndividual) GetConfirmedNoUsTaxID() (ConfirmedNoUsTaxID bool) {
	if r != nil && r.ConfirmedNoUsTaxID != nil {
		ConfirmedNoUsTaxID = *r.ConfirmedNoUsTaxID
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

func (r CreateAnEntityParametersTrustTrusteesIndividual) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustTrusteesIndividual{Name:%s DateOfBirth:%s Address:%s ConfirmedNoUsTaxID:%s Identification:%s}", core.FmtP(r.Name), core.FmtP(r.DateOfBirth), r.Address, core.FmtP(r.ConfirmedNoUsTaxID), r.Identification)
}

type CreateAnEntityParametersTrustTrusteesIndividualAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 *string `pjson:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 *string `pjson:"line2"`
	// The city of the address.
	City *string `pjson:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `pjson:"state"`
	// The ZIP code of the address.
	Zip        *string                `pjson:"zip"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersTrustTrusteesIndividualAddress using the internal pjson
// library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersTrustTrusteesIndividualAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersTrustTrusteesIndividualAddress
// into an array of bytes using the gjson library. Members of the `Extras` field
// are serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateAnEntityParametersTrustTrusteesIndividualAddress) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r CreateAnEntityParametersTrustTrusteesIndividualAddress) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustTrusteesIndividualAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", core.FmtP(r.Line1), core.FmtP(r.Line2), core.FmtP(r.City), core.FmtP(r.State), core.FmtP(r.Zip))
}

type CreateAnEntityParametersTrustTrusteesIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method *CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethod `pjson:"method"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number *string `pjson:"number"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport *CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport `pjson:"passport"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense *CreateAnEntityParametersTrustTrusteesIndividualIdentificationDriversLicense `pjson:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other      *CreateAnEntityParametersTrustTrusteesIndividualIdentificationOther `pjson:"other"`
	jsonFields map[string]interface{}                                              `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersTrustTrusteesIndividualIdentification using the internal
// pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentification) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CreateAnEntityParametersTrustTrusteesIndividualIdentification into an array of
// bytes using the gjson library. Members of the `Extras` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentification) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

// Information about the United States driver's license used for identification.
// Required if `method` is equal to `drivers_license`.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentification) GetDriversLicense() (DriversLicense CreateAnEntityParametersTrustTrusteesIndividualIdentificationDriversLicense) {
	if r != nil && r.DriversLicense != nil {
		DriversLicense = *r.DriversLicense
	}
	return
}

// Information about the identification document provided. Required if `method` is
// equal to `other`.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentification) GetOther() (Other CreateAnEntityParametersTrustTrusteesIndividualIdentificationOther) {
	if r != nil && r.Other != nil {
		Other = *r.Other
	}
	return
}

func (r CreateAnEntityParametersTrustTrusteesIndividualIdentification) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustTrusteesIndividualIdentification{Method:%s Number:%s Passport:%s DriversLicense:%s Other:%s}", core.FmtP(r.Method), core.FmtP(r.Number), r.Passport, r.DriversLicense, r.Other)
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
	FileID *string `pjson:"file_id"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `pjson:"expiration_date"`
	// The country that issued the passport.
	Country    *string                `pjson:"country"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport using the
// internal pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport into an
// array of bytes using the gjson library. Members of the `Extras` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport{FileID:%s ExpirationDate:%s Country:%s}", core.FmtP(r.FileID), core.FmtP(r.ExpirationDate), core.FmtP(r.Country))
}

type CreateAnEntityParametersTrustTrusteesIndividualIdentificationDriversLicense struct {
	// The identifier of the File containing the driver's license.
	FileID *string `pjson:"file_id"`
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `pjson:"expiration_date"`
	// The state that issued the provided driver's license.
	State      *string                `pjson:"state"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersTrustTrusteesIndividualIdentificationDriversLicense
// using the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentificationDriversLicense) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CreateAnEntityParametersTrustTrusteesIndividualIdentificationDriversLicense into
// an array of bytes using the gjson library. Members of the `Extras` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentificationDriversLicense) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the File containing the driver's license.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentificationDriversLicense) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

// The driver's license's expiration date in YYYY-MM-DD format.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentificationDriversLicense) GetExpirationDate() (ExpirationDate string) {
	if r != nil && r.ExpirationDate != nil {
		ExpirationDate = *r.ExpirationDate
	}
	return
}

// The state that issued the provided driver's license.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentificationDriversLicense) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

func (r CreateAnEntityParametersTrustTrusteesIndividualIdentificationDriversLicense) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustTrusteesIndividualIdentificationDriversLicense{FileID:%s ExpirationDate:%s State:%s}", core.FmtP(r.FileID), core.FmtP(r.ExpirationDate), core.FmtP(r.State))
}

type CreateAnEntityParametersTrustTrusteesIndividualIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document.
	Country *string `pjson:"country"`
	// A description of the document submitted.
	Description *string `pjson:"description"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `pjson:"expiration_date"`
	// The identifier of the File containing the document.
	FileID     *string                `pjson:"file_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersTrustTrusteesIndividualIdentificationOther using the
// internal pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentificationOther) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CreateAnEntityParametersTrustTrusteesIndividualIdentificationOther into an array
// of bytes using the gjson library. Members of the `Extras` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentificationOther) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The two-character ISO 3166-1 code representing the country that issued the
// document.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentificationOther) GetCountry() (Country string) {
	if r != nil && r.Country != nil {
		Country = *r.Country
	}
	return
}

// A description of the document submitted.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentificationOther) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The document's expiration date in YYYY-MM-DD format.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentificationOther) GetExpirationDate() (ExpirationDate string) {
	if r != nil && r.ExpirationDate != nil {
		ExpirationDate = *r.ExpirationDate
	}
	return
}

// The identifier of the File containing the document.
func (r *CreateAnEntityParametersTrustTrusteesIndividualIdentificationOther) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

func (r CreateAnEntityParametersTrustTrusteesIndividualIdentificationOther) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustTrusteesIndividualIdentificationOther{Country:%s Description:%s ExpirationDate:%s FileID:%s}", core.FmtP(r.Country), core.FmtP(r.Description), core.FmtP(r.ExpirationDate), core.FmtP(r.FileID))
}

type CreateAnEntityParametersTrustGrantor struct {
	// The person's legal name.
	Name *string `pjson:"name"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth *string `pjson:"date_of_birth"`
	// The individual's address.
	Address *CreateAnEntityParametersTrustGrantorAddress `pjson:"address"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID *bool `pjson:"confirmed_no_us_tax_id"`
	// A means of verifying the person's identity.
	Identification *CreateAnEntityParametersTrustGrantorIdentification `pjson:"identification"`
	jsonFields     map[string]interface{}                              `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersTrustGrantor using the internal pjson library.
// Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersTrustGrantor) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersTrustGrantor into an array of
// bytes using the gjson library. Members of the `Extras` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnEntityParametersTrustGrantor) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

// The identification method for an individual can only be a passport, driver's
// license, or other document if you've confirmed the individual does not have a US
// tax id (either a Social Security Number or Individual Taxpayer Identification
// Number).
func (r *CreateAnEntityParametersTrustGrantor) GetConfirmedNoUsTaxID() (ConfirmedNoUsTaxID bool) {
	if r != nil && r.ConfirmedNoUsTaxID != nil {
		ConfirmedNoUsTaxID = *r.ConfirmedNoUsTaxID
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

func (r CreateAnEntityParametersTrustGrantor) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustGrantor{Name:%s DateOfBirth:%s Address:%s ConfirmedNoUsTaxID:%s Identification:%s}", core.FmtP(r.Name), core.FmtP(r.DateOfBirth), r.Address, core.FmtP(r.ConfirmedNoUsTaxID), r.Identification)
}

type CreateAnEntityParametersTrustGrantorAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 *string `pjson:"line1"`
	// The second line of the address. This might be the floor or room number.
	Line2 *string `pjson:"line2"`
	// The city of the address.
	City *string `pjson:"city"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `pjson:"state"`
	// The ZIP code of the address.
	Zip        *string                `pjson:"zip"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersTrustGrantorAddress using the internal pjson library.
// Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersTrustGrantorAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersTrustGrantorAddress into an array
// of bytes using the gjson library. Members of the `Extras` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *CreateAnEntityParametersTrustGrantorAddress) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r CreateAnEntityParametersTrustGrantorAddress) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustGrantorAddress{Line1:%s Line2:%s City:%s State:%s Zip:%s}", core.FmtP(r.Line1), core.FmtP(r.Line2), core.FmtP(r.City), core.FmtP(r.State), core.FmtP(r.Zip))
}

type CreateAnEntityParametersTrustGrantorIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method *CreateAnEntityParametersTrustGrantorIdentificationMethod `pjson:"method"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number *string `pjson:"number"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport *CreateAnEntityParametersTrustGrantorIdentificationPassport `pjson:"passport"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense *CreateAnEntityParametersTrustGrantorIdentificationDriversLicense `pjson:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other      *CreateAnEntityParametersTrustGrantorIdentificationOther `pjson:"other"`
	jsonFields map[string]interface{}                                   `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersTrustGrantorIdentification using the internal pjson
// library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersTrustGrantorIdentification) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersTrustGrantorIdentification into
// an array of bytes using the gjson library. Members of the `Extras` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateAnEntityParametersTrustGrantorIdentification) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

// Information about the United States driver's license used for identification.
// Required if `method` is equal to `drivers_license`.
func (r *CreateAnEntityParametersTrustGrantorIdentification) GetDriversLicense() (DriversLicense CreateAnEntityParametersTrustGrantorIdentificationDriversLicense) {
	if r != nil && r.DriversLicense != nil {
		DriversLicense = *r.DriversLicense
	}
	return
}

// Information about the identification document provided. Required if `method` is
// equal to `other`.
func (r *CreateAnEntityParametersTrustGrantorIdentification) GetOther() (Other CreateAnEntityParametersTrustGrantorIdentificationOther) {
	if r != nil && r.Other != nil {
		Other = *r.Other
	}
	return
}

func (r CreateAnEntityParametersTrustGrantorIdentification) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustGrantorIdentification{Method:%s Number:%s Passport:%s DriversLicense:%s Other:%s}", core.FmtP(r.Method), core.FmtP(r.Number), r.Passport, r.DriversLicense, r.Other)
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
	FileID *string `pjson:"file_id"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `pjson:"expiration_date"`
	// The country that issued the passport.
	Country    *string                `pjson:"country"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersTrustGrantorIdentificationPassport using the internal
// pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersTrustGrantorIdentificationPassport) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CreateAnEntityParametersTrustGrantorIdentificationPassport into an array of
// bytes using the gjson library. Members of the `Extras` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnEntityParametersTrustGrantorIdentificationPassport) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r CreateAnEntityParametersTrustGrantorIdentificationPassport) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustGrantorIdentificationPassport{FileID:%s ExpirationDate:%s Country:%s}", core.FmtP(r.FileID), core.FmtP(r.ExpirationDate), core.FmtP(r.Country))
}

type CreateAnEntityParametersTrustGrantorIdentificationDriversLicense struct {
	// The identifier of the File containing the driver's license.
	FileID *string `pjson:"file_id"`
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `pjson:"expiration_date"`
	// The state that issued the provided driver's license.
	State      *string                `pjson:"state"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersTrustGrantorIdentificationDriversLicense using the
// internal pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersTrustGrantorIdentificationDriversLicense) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CreateAnEntityParametersTrustGrantorIdentificationDriversLicense into an array
// of bytes using the gjson library. Members of the `Extras` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *CreateAnEntityParametersTrustGrantorIdentificationDriversLicense) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the File containing the driver's license.
func (r *CreateAnEntityParametersTrustGrantorIdentificationDriversLicense) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

// The driver's license's expiration date in YYYY-MM-DD format.
func (r *CreateAnEntityParametersTrustGrantorIdentificationDriversLicense) GetExpirationDate() (ExpirationDate string) {
	if r != nil && r.ExpirationDate != nil {
		ExpirationDate = *r.ExpirationDate
	}
	return
}

// The state that issued the provided driver's license.
func (r *CreateAnEntityParametersTrustGrantorIdentificationDriversLicense) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

func (r CreateAnEntityParametersTrustGrantorIdentificationDriversLicense) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustGrantorIdentificationDriversLicense{FileID:%s ExpirationDate:%s State:%s}", core.FmtP(r.FileID), core.FmtP(r.ExpirationDate), core.FmtP(r.State))
}

type CreateAnEntityParametersTrustGrantorIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document.
	Country *string `pjson:"country"`
	// A description of the document submitted.
	Description *string `pjson:"description"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `pjson:"expiration_date"`
	// The identifier of the File containing the document.
	FileID     *string                `pjson:"file_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersTrustGrantorIdentificationOther using the internal pjson
// library. Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersTrustGrantorIdentificationOther) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersTrustGrantorIdentificationOther
// into an array of bytes using the gjson library. Members of the `Extras` field
// are serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateAnEntityParametersTrustGrantorIdentificationOther) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The two-character ISO 3166-1 code representing the country that issued the
// document.
func (r *CreateAnEntityParametersTrustGrantorIdentificationOther) GetCountry() (Country string) {
	if r != nil && r.Country != nil {
		Country = *r.Country
	}
	return
}

// A description of the document submitted.
func (r *CreateAnEntityParametersTrustGrantorIdentificationOther) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The document's expiration date in YYYY-MM-DD format.
func (r *CreateAnEntityParametersTrustGrantorIdentificationOther) GetExpirationDate() (ExpirationDate string) {
	if r != nil && r.ExpirationDate != nil {
		ExpirationDate = *r.ExpirationDate
	}
	return
}

// The identifier of the File containing the document.
func (r *CreateAnEntityParametersTrustGrantorIdentificationOther) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

func (r CreateAnEntityParametersTrustGrantorIdentificationOther) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersTrustGrantorIdentificationOther{Country:%s Description:%s ExpirationDate:%s FileID:%s}", core.FmtP(r.Country), core.FmtP(r.Description), core.FmtP(r.ExpirationDate), core.FmtP(r.FileID))
}

type CreateAnEntityParametersRelationship string

const (
	CreateAnEntityParametersRelationshipAffiliated    CreateAnEntityParametersRelationship = "affiliated"
	CreateAnEntityParametersRelationshipInformational CreateAnEntityParametersRelationship = "informational"
	CreateAnEntityParametersRelationshipUnaffiliated  CreateAnEntityParametersRelationship = "unaffiliated"
)

type CreateAnEntityParametersSupplementalDocuments struct {
	// The identifier of the File containing the document.
	FileID     *string                `pjson:"file_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEntityParametersSupplementalDocuments using the internal pjson library.
// Unrecognized fields are stored in the `Extras` property.
func (r *CreateAnEntityParametersSupplementalDocuments) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEntityParametersSupplementalDocuments into an
// array of bytes using the gjson library. Members of the `Extras` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateAnEntityParametersSupplementalDocuments) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the File containing the document.
func (r *CreateAnEntityParametersSupplementalDocuments) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

func (r CreateAnEntityParametersSupplementalDocuments) String() (result string) {
	return fmt.Sprintf("&CreateAnEntityParametersSupplementalDocuments{FileID:%s}", core.FmtP(r.FileID))
}

type EntityListParams struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit      *int64                 `query:"limit"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EntityListParams using the
// internal pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *EntityListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityListParams into an array of bytes using the gjson
// library. Members of the `Extras` field are serialized into the top-level, and
// will overwrite known members of the same name.
func (r *EntityListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Return the page of entries after this one.
func (r *EntityListParams) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *EntityListParams) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

func (r EntityListParams) String() (result string) {
	return fmt.Sprintf("&EntityListParams{Cursor:%s Limit:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit))
}

type EntityList struct {
	// The contents of the list.
	Data *[]Entity `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EntityList using the internal
// pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *EntityList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EntityList into an array of bytes using the gjson
// library. Members of the `Extras` field are serialized into the top-level, and
// will overwrite known members of the same name.
func (r *EntityList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r EntityList) String() (result string) {
	return fmt.Sprintf("&EntityList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type EntitiesPage struct {
	*pagination.Page[Entity]
}

func (r *EntitiesPage) Entity() *Entity {
	return r.Current()
}

func (r *EntitiesPage) NextPage() (*EntitiesPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &EntitiesPage{page}, nil
	}
}
