package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

type EntityService struct {
	Options               []option.RequestOption
	SupplementalDocuments *EntitySupplementalDocumentService
}

func NewEntityService(opts ...option.RequestOption) (r *EntityService) {
	r = &EntityService{}
	r.Options = opts
	r.SupplementalDocuments = NewEntitySupplementalDocumentService(opts...)
	return
}

// Create an Entity
func (r *EntityService) New(ctx context.Context, body EntityNewParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := "entities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an Entity
func (r *EntityService) Get(ctx context.Context, entity_id string, opts ...option.RequestOption) (res *Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("entities/%s", entity_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Entities
func (r *EntityService) List(ctx context.Context, query EntityListParams, opts ...option.RequestOption) (res *shared.Page[Entity], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "entities"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Entities
func (r *EntityService) ListAutoPaging(ctx context.Context, query EntityListParams, opts ...option.RequestOption) *shared.PageAutoPager[Entity] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

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
	raw                   string
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
	raw                string
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
	raw    string
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
	raw          string
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
	raw            string
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
	raw    string
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
	raw         string
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
	raw            string
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
	raw    string
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
	raw         string
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
	raw         string
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
	raw            string
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
	raw    string
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
	raw         string
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
	raw                     string
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
	raw    string
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
	raw        string
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
	raw            string
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
	raw    string
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
	raw         string
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
	raw            string
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
	raw    string
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
	raw         string
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
	raw    string
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EntitySupplementalDocuments
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *EntitySupplementalDocuments) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
	return apijson.MarshalRoot(r)
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
	Limit     field.Field[int64]                     `query:"limit"`
	CreatedAt field.Field[EntityListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes EntityListParams into a url.Values of the query parameters
// associated with this value
func (r EntityListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type EntityListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After field.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before field.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter field.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore field.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes EntityListParamsCreatedAt into a url.Values of the query
// parameters associated with this value
func (r EntityListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
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
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EntityListResponse using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EntityListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
