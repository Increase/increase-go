// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// EntityService contains methods and other services that help with interacting
// with the increase API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewEntityService] method instead.
type EntityService struct {
	Options               []option.RequestOption
	SupplementalDocuments *EntitySupplementalDocumentService
}

// NewEntityService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
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
func (r *EntityService) Get(ctx context.Context, entityID string, opts ...option.RequestOption) (res *Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("entities/%s", entityID)
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

// Entities are the legal entities that own accounts. They can be people,
// corporations, partnerships, or trusts.
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
	// Additional documentation associated with the entity. This is limited to the
	// first 10 documents for an entity. If an entity has more than 10 documents, use
	// the GET /entity_supplemental_documents list endpoint to retrieve them.
	SupplementalDocuments []EntitySupplementalDocuments `json:"supplemental_documents,required"`
	JSON                  entityJSON
}

// entityJSON contains the JSON metadata for the struct [Entity]
type entityJSON struct {
	ID                    apijson.Field
	Structure             apijson.Field
	Corporation           apijson.Field
	NaturalPerson         apijson.Field
	Joint                 apijson.Field
	Trust                 apijson.Field
	Type                  apijson.Field
	Description           apijson.Field
	Relationship          apijson.Field
	SupplementalDocuments apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Entity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The entity's legal structure.
type EntityStructure string

const (
	EntityStructureCorporation   EntityStructure = "corporation"
	EntityStructureNaturalPerson EntityStructure = "natural_person"
	EntityStructureJoint         EntityStructure = "joint"
	EntityStructureTrust         EntityStructure = "trust"
)

// Details of the corporation entity. Will be present if `structure` is equal to
// `corporation`.
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
	JSON             entityCorporationJSON
}

// entityCorporationJSON contains the JSON metadata for the struct
// [EntityCorporation]
type entityCorporationJSON struct {
	Name               apijson.Field
	Website            apijson.Field
	TaxIdentifier      apijson.Field
	IncorporationState apijson.Field
	Address            apijson.Field
	BeneficialOwners   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EntityCorporation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The corporation's address.
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
	JSON entityCorporationAddressJSON
}

// entityCorporationAddressJSON contains the JSON metadata for the struct
// [EntityCorporationAddress]
type entityCorporationAddressJSON struct {
	Line1       apijson.Field
	Line2       apijson.Field
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

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
	JSON  entityCorporationBeneficialOwnersJSON
}

// entityCorporationBeneficialOwnersJSON contains the JSON metadata for the struct
// [EntityCorporationBeneficialOwners]
type entityCorporationBeneficialOwnersJSON struct {
	Individual   apijson.Field
	CompanyTitle apijson.Field
	Prong        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EntityCorporationBeneficialOwners) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Personal details for the beneficial owner.
type EntityCorporationBeneficialOwnersIndividual struct {
	// The person's legal name.
	Name string `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth time.Time `json:"date_of_birth,required" format:"date"`
	// The person's address.
	Address EntityCorporationBeneficialOwnersIndividualAddress `json:"address,required"`
	// A means of verifying the person's identity.
	Identification EntityCorporationBeneficialOwnersIndividualIdentification `json:"identification,required"`
	JSON           entityCorporationBeneficialOwnersIndividualJSON
}

// entityCorporationBeneficialOwnersIndividualJSON contains the JSON metadata for
// the struct [EntityCorporationBeneficialOwnersIndividual]
type entityCorporationBeneficialOwnersIndividualJSON struct {
	Name           apijson.Field
	DateOfBirth    apijson.Field
	Address        apijson.Field
	Identification apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityCorporationBeneficialOwnersIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The person's address.
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
	JSON entityCorporationBeneficialOwnersIndividualAddressJSON
}

// entityCorporationBeneficialOwnersIndividualAddressJSON contains the JSON
// metadata for the struct [EntityCorporationBeneficialOwnersIndividualAddress]
type entityCorporationBeneficialOwnersIndividualAddressJSON struct {
	Line1       apijson.Field
	Line2       apijson.Field
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityCorporationBeneficialOwnersIndividualAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A means of verifying the person's identity.
type EntityCorporationBeneficialOwnersIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityCorporationBeneficialOwnersIndividualIdentificationMethod `json:"method,required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string `json:"number_last4,required"`
	JSON        entityCorporationBeneficialOwnersIndividualIdentificationJSON
}

// entityCorporationBeneficialOwnersIndividualIdentificationJSON contains the JSON
// metadata for the struct
// [EntityCorporationBeneficialOwnersIndividualIdentification]
type entityCorporationBeneficialOwnersIndividualIdentificationJSON struct {
	Method      apijson.Field
	NumberLast4 apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityCorporationBeneficialOwnersIndividualIdentification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A method that can be used to verify the individual's identity.
type EntityCorporationBeneficialOwnersIndividualIdentificationMethod string

const (
	EntityCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber                   EntityCorporationBeneficialOwnersIndividualIdentificationMethod = "social_security_number"
	EntityCorporationBeneficialOwnersIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber EntityCorporationBeneficialOwnersIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	EntityCorporationBeneficialOwnersIndividualIdentificationMethodPassport                               EntityCorporationBeneficialOwnersIndividualIdentificationMethod = "passport"
	EntityCorporationBeneficialOwnersIndividualIdentificationMethodDriversLicense                         EntityCorporationBeneficialOwnersIndividualIdentificationMethod = "drivers_license"
	EntityCorporationBeneficialOwnersIndividualIdentificationMethodOther                                  EntityCorporationBeneficialOwnersIndividualIdentificationMethod = "other"
)

// Why this person is considered a beneficial owner of the entity.
type EntityCorporationBeneficialOwnersProng string

const (
	EntityCorporationBeneficialOwnersProngOwnership EntityCorporationBeneficialOwnersProng = "ownership"
	EntityCorporationBeneficialOwnersProngControl   EntityCorporationBeneficialOwnersProng = "control"
)

// Details of the natural person entity. Will be present if `structure` is equal to
// `natural_person`.
type EntityNaturalPerson struct {
	// The person's legal name.
	Name string `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth time.Time `json:"date_of_birth,required" format:"date"`
	// The person's address.
	Address EntityNaturalPersonAddress `json:"address,required"`
	// A means of verifying the person's identity.
	Identification EntityNaturalPersonIdentification `json:"identification,required"`
	JSON           entityNaturalPersonJSON
}

// entityNaturalPersonJSON contains the JSON metadata for the struct
// [EntityNaturalPerson]
type entityNaturalPersonJSON struct {
	Name           apijson.Field
	DateOfBirth    apijson.Field
	Address        apijson.Field
	Identification apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityNaturalPerson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The person's address.
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
	JSON entityNaturalPersonAddressJSON
}

// entityNaturalPersonAddressJSON contains the JSON metadata for the struct
// [EntityNaturalPersonAddress]
type entityNaturalPersonAddressJSON struct {
	Line1       apijson.Field
	Line2       apijson.Field
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityNaturalPersonAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A means of verifying the person's identity.
type EntityNaturalPersonIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityNaturalPersonIdentificationMethod `json:"method,required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string `json:"number_last4,required"`
	JSON        entityNaturalPersonIdentificationJSON
}

// entityNaturalPersonIdentificationJSON contains the JSON metadata for the struct
// [EntityNaturalPersonIdentification]
type entityNaturalPersonIdentificationJSON struct {
	Method      apijson.Field
	NumberLast4 apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityNaturalPersonIdentification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A method that can be used to verify the individual's identity.
type EntityNaturalPersonIdentificationMethod string

const (
	EntityNaturalPersonIdentificationMethodSocialSecurityNumber                   EntityNaturalPersonIdentificationMethod = "social_security_number"
	EntityNaturalPersonIdentificationMethodIndividualTaxpayerIdentificationNumber EntityNaturalPersonIdentificationMethod = "individual_taxpayer_identification_number"
	EntityNaturalPersonIdentificationMethodPassport                               EntityNaturalPersonIdentificationMethod = "passport"
	EntityNaturalPersonIdentificationMethodDriversLicense                         EntityNaturalPersonIdentificationMethod = "drivers_license"
	EntityNaturalPersonIdentificationMethodOther                                  EntityNaturalPersonIdentificationMethod = "other"
)

// Details of the joint entity. Will be present if `structure` is equal to `joint`.
type EntityJoint struct {
	// The entity's name.
	Name string `json:"name,required"`
	// The two individuals that share control of the entity.
	Individuals []EntityJointIndividuals `json:"individuals,required"`
	JSON        entityJointJSON
}

// entityJointJSON contains the JSON metadata for the struct [EntityJoint]
type entityJointJSON struct {
	Name        apijson.Field
	Individuals apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

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
	JSON           entityJointIndividualsJSON
}

// entityJointIndividualsJSON contains the JSON metadata for the struct
// [EntityJointIndividuals]
type entityJointIndividualsJSON struct {
	Name           apijson.Field
	DateOfBirth    apijson.Field
	Address        apijson.Field
	Identification apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityJointIndividuals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The person's address.
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
	JSON entityJointIndividualsAddressJSON
}

// entityJointIndividualsAddressJSON contains the JSON metadata for the struct
// [EntityJointIndividualsAddress]
type entityJointIndividualsAddressJSON struct {
	Line1       apijson.Field
	Line2       apijson.Field
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityJointIndividualsAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A means of verifying the person's identity.
type EntityJointIndividualsIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityJointIndividualsIdentificationMethod `json:"method,required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string `json:"number_last4,required"`
	JSON        entityJointIndividualsIdentificationJSON
}

// entityJointIndividualsIdentificationJSON contains the JSON metadata for the
// struct [EntityJointIndividualsIdentification]
type entityJointIndividualsIdentificationJSON struct {
	Method      apijson.Field
	NumberLast4 apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityJointIndividualsIdentification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A method that can be used to verify the individual's identity.
type EntityJointIndividualsIdentificationMethod string

const (
	EntityJointIndividualsIdentificationMethodSocialSecurityNumber                   EntityJointIndividualsIdentificationMethod = "social_security_number"
	EntityJointIndividualsIdentificationMethodIndividualTaxpayerIdentificationNumber EntityJointIndividualsIdentificationMethod = "individual_taxpayer_identification_number"
	EntityJointIndividualsIdentificationMethodPassport                               EntityJointIndividualsIdentificationMethod = "passport"
	EntityJointIndividualsIdentificationMethodDriversLicense                         EntityJointIndividualsIdentificationMethod = "drivers_license"
	EntityJointIndividualsIdentificationMethodOther                                  EntityJointIndividualsIdentificationMethod = "other"
)

// Details of the trust entity. Will be present if `structure` is equal to `trust`.
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
	JSON                    entityTrustJSON
}

// entityTrustJSON contains the JSON metadata for the struct [EntityTrust]
type entityTrustJSON struct {
	Name                    apijson.Field
	Category                apijson.Field
	Address                 apijson.Field
	FormationState          apijson.Field
	TaxIdentifier           apijson.Field
	Trustees                apijson.Field
	Grantor                 apijson.Field
	FormationDocumentFileID apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *EntityTrust) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the trust is `revocable` or `irrevocable`.
type EntityTrustCategory string

const (
	EntityTrustCategoryRevocable   EntityTrustCategory = "revocable"
	EntityTrustCategoryIrrevocable EntityTrustCategory = "irrevocable"
)

// The trust's address.
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
	JSON entityTrustAddressJSON
}

// entityTrustAddressJSON contains the JSON metadata for the struct
// [EntityTrustAddress]
type entityTrustAddressJSON struct {
	Line1       apijson.Field
	Line2       apijson.Field
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityTrustAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityTrustTrustees struct {
	// The structure of the trustee. Will always be equal to `individual`.
	Structure EntityTrustTrusteesStructure `json:"structure,required"`
	// The individual trustee of the trust. Will be present if the trustee's
	// `structure` is equal to `individual`.
	Individual EntityTrustTrusteesIndividual `json:"individual,required,nullable"`
	JSON       entityTrustTrusteesJSON
}

// entityTrustTrusteesJSON contains the JSON metadata for the struct
// [EntityTrustTrustees]
type entityTrustTrusteesJSON struct {
	Structure   apijson.Field
	Individual  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityTrustTrustees) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The structure of the trustee. Will always be equal to `individual`.
type EntityTrustTrusteesStructure string

const (
	EntityTrustTrusteesStructureIndividual EntityTrustTrusteesStructure = "individual"
)

// The individual trustee of the trust. Will be present if the trustee's
// `structure` is equal to `individual`.
type EntityTrustTrusteesIndividual struct {
	// The person's legal name.
	Name string `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth time.Time `json:"date_of_birth,required" format:"date"`
	// The person's address.
	Address EntityTrustTrusteesIndividualAddress `json:"address,required"`
	// A means of verifying the person's identity.
	Identification EntityTrustTrusteesIndividualIdentification `json:"identification,required"`
	JSON           entityTrustTrusteesIndividualJSON
}

// entityTrustTrusteesIndividualJSON contains the JSON metadata for the struct
// [EntityTrustTrusteesIndividual]
type entityTrustTrusteesIndividualJSON struct {
	Name           apijson.Field
	DateOfBirth    apijson.Field
	Address        apijson.Field
	Identification apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityTrustTrusteesIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The person's address.
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
	JSON entityTrustTrusteesIndividualAddressJSON
}

// entityTrustTrusteesIndividualAddressJSON contains the JSON metadata for the
// struct [EntityTrustTrusteesIndividualAddress]
type entityTrustTrusteesIndividualAddressJSON struct {
	Line1       apijson.Field
	Line2       apijson.Field
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityTrustTrusteesIndividualAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A means of verifying the person's identity.
type EntityTrustTrusteesIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityTrustTrusteesIndividualIdentificationMethod `json:"method,required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string `json:"number_last4,required"`
	JSON        entityTrustTrusteesIndividualIdentificationJSON
}

// entityTrustTrusteesIndividualIdentificationJSON contains the JSON metadata for
// the struct [EntityTrustTrusteesIndividualIdentification]
type entityTrustTrusteesIndividualIdentificationJSON struct {
	Method      apijson.Field
	NumberLast4 apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityTrustTrusteesIndividualIdentification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A method that can be used to verify the individual's identity.
type EntityTrustTrusteesIndividualIdentificationMethod string

const (
	EntityTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber                   EntityTrustTrusteesIndividualIdentificationMethod = "social_security_number"
	EntityTrustTrusteesIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber EntityTrustTrusteesIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	EntityTrustTrusteesIndividualIdentificationMethodPassport                               EntityTrustTrusteesIndividualIdentificationMethod = "passport"
	EntityTrustTrusteesIndividualIdentificationMethodDriversLicense                         EntityTrustTrusteesIndividualIdentificationMethod = "drivers_license"
	EntityTrustTrusteesIndividualIdentificationMethodOther                                  EntityTrustTrusteesIndividualIdentificationMethod = "other"
)

// The grantor of the trust. Will be present if the `category` is `revocable`.
type EntityTrustGrantor struct {
	// The person's legal name.
	Name string `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth time.Time `json:"date_of_birth,required" format:"date"`
	// The person's address.
	Address EntityTrustGrantorAddress `json:"address,required"`
	// A means of verifying the person's identity.
	Identification EntityTrustGrantorIdentification `json:"identification,required"`
	JSON           entityTrustGrantorJSON
}

// entityTrustGrantorJSON contains the JSON metadata for the struct
// [EntityTrustGrantor]
type entityTrustGrantorJSON struct {
	Name           apijson.Field
	DateOfBirth    apijson.Field
	Address        apijson.Field
	Identification apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityTrustGrantor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The person's address.
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
	JSON entityTrustGrantorAddressJSON
}

// entityTrustGrantorAddressJSON contains the JSON metadata for the struct
// [EntityTrustGrantorAddress]
type entityTrustGrantorAddressJSON struct {
	Line1       apijson.Field
	Line2       apijson.Field
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityTrustGrantorAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A means of verifying the person's identity.
type EntityTrustGrantorIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityTrustGrantorIdentificationMethod `json:"method,required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string `json:"number_last4,required"`
	JSON        entityTrustGrantorIdentificationJSON
}

// entityTrustGrantorIdentificationJSON contains the JSON metadata for the struct
// [EntityTrustGrantorIdentification]
type entityTrustGrantorIdentificationJSON struct {
	Method      apijson.Field
	NumberLast4 apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityTrustGrantorIdentification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A method that can be used to verify the individual's identity.
type EntityTrustGrantorIdentificationMethod string

const (
	EntityTrustGrantorIdentificationMethodSocialSecurityNumber                   EntityTrustGrantorIdentificationMethod = "social_security_number"
	EntityTrustGrantorIdentificationMethodIndividualTaxpayerIdentificationNumber EntityTrustGrantorIdentificationMethod = "individual_taxpayer_identification_number"
	EntityTrustGrantorIdentificationMethodPassport                               EntityTrustGrantorIdentificationMethod = "passport"
	EntityTrustGrantorIdentificationMethodDriversLicense                         EntityTrustGrantorIdentificationMethod = "drivers_license"
	EntityTrustGrantorIdentificationMethodOther                                  EntityTrustGrantorIdentificationMethod = "other"
)

// A constant representing the object's type. For this resource it will always be
// `entity`.
type EntityType string

const (
	EntityTypeEntity EntityType = "entity"
)

// The relationship between your group and the entity.
type EntityRelationship string

const (
	EntityRelationshipAffiliated    EntityRelationship = "affiliated"
	EntityRelationshipInformational EntityRelationship = "informational"
	EntityRelationshipUnaffiliated  EntityRelationship = "unaffiliated"
)

// Supplemental Documents are uploaded files connected to an Entity during
// onboarding.
type EntitySupplementalDocuments struct {
	// The File containing the document.
	FileID string `json:"file_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the
	// Supplemental Document was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A constant representing the object's type. For this resource it will always be
	// `entity_supplemental_document`.
	Type EntitySupplementalDocumentsType `json:"type,required"`
	JSON entitySupplementalDocumentsJSON
}

// entitySupplementalDocumentsJSON contains the JSON metadata for the struct
// [EntitySupplementalDocuments]
type entitySupplementalDocumentsJSON struct {
	FileID      apijson.Field
	CreatedAt   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitySupplementalDocuments) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A constant representing the object's type. For this resource it will always be
// `entity_supplemental_document`.
type EntitySupplementalDocumentsType string

const (
	EntitySupplementalDocumentsTypeEntitySupplementalDocument EntitySupplementalDocumentsType = "entity_supplemental_document"
)

type EntityNewParams struct {
	// The relationship between your group and the entity.
	Relationship param.Field[EntityNewParamsRelationship] `json:"relationship,required"`
	// The type of Entity to create.
	Structure param.Field[EntityNewParamsStructure] `json:"structure,required"`
	// Details of the corporation entity to create. Required if `structure` is equal to
	// `corporation`.
	Corporation param.Field[EntityNewParamsCorporation] `json:"corporation"`
	// The description you choose to give the entity.
	Description param.Field[string] `json:"description"`
	// Details of the joint entity to create. Required if `structure` is equal to
	// `joint`.
	Joint param.Field[EntityNewParamsJoint] `json:"joint"`
	// Details of the natural person entity to create. Required if `structure` is equal
	// to `natural_person`. Natural people entities should be submitted with
	// `social_security_number` or `individual_taxpayer_identification_number`
	// identification methods.
	NaturalPerson param.Field[EntityNewParamsNaturalPerson] `json:"natural_person"`
	// Additional documentation associated with the entity.
	SupplementalDocuments param.Field[[]EntityNewParamsSupplementalDocuments] `json:"supplemental_documents"`
	// Details of the trust entity to create. Required if `structure` is equal to
	// `trust`.
	Trust param.Field[EntityNewParamsTrust] `json:"trust"`
}

func (r EntityNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The relationship between your group and the entity.
type EntityNewParamsRelationship string

const (
	EntityNewParamsRelationshipAffiliated    EntityNewParamsRelationship = "affiliated"
	EntityNewParamsRelationshipInformational EntityNewParamsRelationship = "informational"
	EntityNewParamsRelationshipUnaffiliated  EntityNewParamsRelationship = "unaffiliated"
)

// The type of Entity to create.
type EntityNewParamsStructure string

const (
	EntityNewParamsStructureCorporation   EntityNewParamsStructure = "corporation"
	EntityNewParamsStructureNaturalPerson EntityNewParamsStructure = "natural_person"
	EntityNewParamsStructureJoint         EntityNewParamsStructure = "joint"
	EntityNewParamsStructureTrust         EntityNewParamsStructure = "trust"
)

// Details of the corporation entity to create. Required if `structure` is equal to
// `corporation`.
type EntityNewParamsCorporation struct {
	// The legal name of the corporation.
	Name param.Field[string] `json:"name,required"`
	// The website of the corporation.
	Website param.Field[string] `json:"website"`
	// The Employer Identification Number (EIN) for the corporation.
	TaxIdentifier param.Field[string] `json:"tax_identifier,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the
	// corporation's state of incorporation.
	IncorporationState param.Field[string] `json:"incorporation_state"`
	// The corporation's address.
	Address param.Field[EntityNewParamsCorporationAddress] `json:"address,required"`
	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwners param.Field[[]EntityNewParamsCorporationBeneficialOwners] `json:"beneficial_owners,required"`
}

func (r EntityNewParamsCorporation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The corporation's address.
type EntityNewParamsCorporationAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
	// The city of the address.
	City param.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State param.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip param.Field[string] `json:"zip,required"`
}

func (r EntityNewParamsCorporationAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityNewParamsCorporationBeneficialOwners struct {
	// Personal details for the beneficial owner.
	Individual param.Field[EntityNewParamsCorporationBeneficialOwnersIndividual] `json:"individual,required"`
	// This person's role or title within the entity.
	CompanyTitle param.Field[string] `json:"company_title"`
	// Why this person is considered a beneficial owner of the entity.
	Prong param.Field[EntityNewParamsCorporationBeneficialOwnersProng] `json:"prong,required"`
}

func (r EntityNewParamsCorporationBeneficialOwners) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Personal details for the beneficial owner.
type EntityNewParamsCorporationBeneficialOwnersIndividual struct {
	// The person's legal name.
	Name param.Field[string] `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth param.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// The individual's address.
	Address param.Field[EntityNewParamsCorporationBeneficialOwnersIndividualAddress] `json:"address,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID param.Field[bool] `json:"confirmed_no_us_tax_id"`
	// A means of verifying the person's identity.
	Identification param.Field[EntityNewParamsCorporationBeneficialOwnersIndividualIdentification] `json:"identification,required"`
}

func (r EntityNewParamsCorporationBeneficialOwnersIndividual) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The individual's address.
type EntityNewParamsCorporationBeneficialOwnersIndividualAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
	// The city of the address.
	City param.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State param.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip param.Field[string] `json:"zip,required"`
}

func (r EntityNewParamsCorporationBeneficialOwnersIndividualAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A means of verifying the person's identity.
type EntityNewParamsCorporationBeneficialOwnersIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method param.Field[EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod] `json:"method,required"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number param.Field[string] `json:"number,required"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport param.Field[EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationPassport] `json:"passport"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense param.Field[EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other param.Field[EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationOther] `json:"other"`
}

func (r EntityNewParamsCorporationBeneficialOwnersIndividualIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A method that can be used to verify the individual's identity.
type EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod string

const (
	EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber                   EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod = "social_security_number"
	EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodPassport                               EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod = "passport"
	EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodDriversLicense                         EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod = "drivers_license"
	EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodOther                                  EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod = "other"
)

// Information about the passport used for identification. Required if `method` is
// equal to `passport`.
type EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationPassport struct {
	// The identifier of the File containing the passport.
	FileID param.Field[string] `json:"file_id,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The country that issued the passport.
	Country param.Field[string] `json:"country,required"`
}

func (r EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationPassport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information about the United States driver's license used for identification.
// Required if `method` is equal to `drivers_license`.
type EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationDriversLicense struct {
	// The identifier of the File containing the driver's license.
	FileID param.Field[string] `json:"file_id,required"`
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The state that issued the provided driver's license.
	State param.Field[string] `json:"state,required"`
}

func (r EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationDriversLicense) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information about the identification document provided. Required if `method` is
// equal to `other`.
type EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document.
	Country param.Field[string] `json:"country,required"`
	// A description of the document submitted.
	Description param.Field[string] `json:"description,required"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date" format:"date"`
	// The identifier of the File containing the document.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationOther) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Why this person is considered a beneficial owner of the entity.
type EntityNewParamsCorporationBeneficialOwnersProng string

const (
	EntityNewParamsCorporationBeneficialOwnersProngOwnership EntityNewParamsCorporationBeneficialOwnersProng = "ownership"
	EntityNewParamsCorporationBeneficialOwnersProngControl   EntityNewParamsCorporationBeneficialOwnersProng = "control"
)

// Details of the joint entity to create. Required if `structure` is equal to
// `joint`.
type EntityNewParamsJoint struct {
	// The name of the joint entity.
	Name param.Field[string] `json:"name"`
	// The two individuals that share control of the entity.
	Individuals param.Field[[]EntityNewParamsJointIndividuals] `json:"individuals,required"`
}

func (r EntityNewParamsJoint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityNewParamsJointIndividuals struct {
	// The person's legal name.
	Name param.Field[string] `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth param.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// The individual's address.
	Address param.Field[EntityNewParamsJointIndividualsAddress] `json:"address,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID param.Field[bool] `json:"confirmed_no_us_tax_id"`
	// A means of verifying the person's identity.
	Identification param.Field[EntityNewParamsJointIndividualsIdentification] `json:"identification,required"`
}

func (r EntityNewParamsJointIndividuals) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The individual's address.
type EntityNewParamsJointIndividualsAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
	// The city of the address.
	City param.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State param.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip param.Field[string] `json:"zip,required"`
}

func (r EntityNewParamsJointIndividualsAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A means of verifying the person's identity.
type EntityNewParamsJointIndividualsIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method param.Field[EntityNewParamsJointIndividualsIdentificationMethod] `json:"method,required"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number param.Field[string] `json:"number,required"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport param.Field[EntityNewParamsJointIndividualsIdentificationPassport] `json:"passport"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense param.Field[EntityNewParamsJointIndividualsIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other param.Field[EntityNewParamsJointIndividualsIdentificationOther] `json:"other"`
}

func (r EntityNewParamsJointIndividualsIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A method that can be used to verify the individual's identity.
type EntityNewParamsJointIndividualsIdentificationMethod string

const (
	EntityNewParamsJointIndividualsIdentificationMethodSocialSecurityNumber                   EntityNewParamsJointIndividualsIdentificationMethod = "social_security_number"
	EntityNewParamsJointIndividualsIdentificationMethodIndividualTaxpayerIdentificationNumber EntityNewParamsJointIndividualsIdentificationMethod = "individual_taxpayer_identification_number"
	EntityNewParamsJointIndividualsIdentificationMethodPassport                               EntityNewParamsJointIndividualsIdentificationMethod = "passport"
	EntityNewParamsJointIndividualsIdentificationMethodDriversLicense                         EntityNewParamsJointIndividualsIdentificationMethod = "drivers_license"
	EntityNewParamsJointIndividualsIdentificationMethodOther                                  EntityNewParamsJointIndividualsIdentificationMethod = "other"
)

// Information about the passport used for identification. Required if `method` is
// equal to `passport`.
type EntityNewParamsJointIndividualsIdentificationPassport struct {
	// The identifier of the File containing the passport.
	FileID param.Field[string] `json:"file_id,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The country that issued the passport.
	Country param.Field[string] `json:"country,required"`
}

func (r EntityNewParamsJointIndividualsIdentificationPassport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information about the United States driver's license used for identification.
// Required if `method` is equal to `drivers_license`.
type EntityNewParamsJointIndividualsIdentificationDriversLicense struct {
	// The identifier of the File containing the driver's license.
	FileID param.Field[string] `json:"file_id,required"`
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The state that issued the provided driver's license.
	State param.Field[string] `json:"state,required"`
}

func (r EntityNewParamsJointIndividualsIdentificationDriversLicense) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information about the identification document provided. Required if `method` is
// equal to `other`.
type EntityNewParamsJointIndividualsIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document.
	Country param.Field[string] `json:"country,required"`
	// A description of the document submitted.
	Description param.Field[string] `json:"description,required"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date" format:"date"`
	// The identifier of the File containing the document.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r EntityNewParamsJointIndividualsIdentificationOther) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details of the natural person entity to create. Required if `structure` is equal
// to `natural_person`. Natural people entities should be submitted with
// `social_security_number` or `individual_taxpayer_identification_number`
// identification methods.
type EntityNewParamsNaturalPerson struct {
	// The person's legal name.
	Name param.Field[string] `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth param.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// The individual's address.
	Address param.Field[EntityNewParamsNaturalPersonAddress] `json:"address,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID param.Field[bool] `json:"confirmed_no_us_tax_id"`
	// A means of verifying the person's identity.
	Identification param.Field[EntityNewParamsNaturalPersonIdentification] `json:"identification,required"`
}

func (r EntityNewParamsNaturalPerson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The individual's address.
type EntityNewParamsNaturalPersonAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
	// The city of the address.
	City param.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State param.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip param.Field[string] `json:"zip,required"`
}

func (r EntityNewParamsNaturalPersonAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A means of verifying the person's identity.
type EntityNewParamsNaturalPersonIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method param.Field[EntityNewParamsNaturalPersonIdentificationMethod] `json:"method,required"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number param.Field[string] `json:"number,required"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport param.Field[EntityNewParamsNaturalPersonIdentificationPassport] `json:"passport"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense param.Field[EntityNewParamsNaturalPersonIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other param.Field[EntityNewParamsNaturalPersonIdentificationOther] `json:"other"`
}

func (r EntityNewParamsNaturalPersonIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A method that can be used to verify the individual's identity.
type EntityNewParamsNaturalPersonIdentificationMethod string

const (
	EntityNewParamsNaturalPersonIdentificationMethodSocialSecurityNumber                   EntityNewParamsNaturalPersonIdentificationMethod = "social_security_number"
	EntityNewParamsNaturalPersonIdentificationMethodIndividualTaxpayerIdentificationNumber EntityNewParamsNaturalPersonIdentificationMethod = "individual_taxpayer_identification_number"
	EntityNewParamsNaturalPersonIdentificationMethodPassport                               EntityNewParamsNaturalPersonIdentificationMethod = "passport"
	EntityNewParamsNaturalPersonIdentificationMethodDriversLicense                         EntityNewParamsNaturalPersonIdentificationMethod = "drivers_license"
	EntityNewParamsNaturalPersonIdentificationMethodOther                                  EntityNewParamsNaturalPersonIdentificationMethod = "other"
)

// Information about the passport used for identification. Required if `method` is
// equal to `passport`.
type EntityNewParamsNaturalPersonIdentificationPassport struct {
	// The identifier of the File containing the passport.
	FileID param.Field[string] `json:"file_id,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The country that issued the passport.
	Country param.Field[string] `json:"country,required"`
}

func (r EntityNewParamsNaturalPersonIdentificationPassport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information about the United States driver's license used for identification.
// Required if `method` is equal to `drivers_license`.
type EntityNewParamsNaturalPersonIdentificationDriversLicense struct {
	// The identifier of the File containing the driver's license.
	FileID param.Field[string] `json:"file_id,required"`
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The state that issued the provided driver's license.
	State param.Field[string] `json:"state,required"`
}

func (r EntityNewParamsNaturalPersonIdentificationDriversLicense) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information about the identification document provided. Required if `method` is
// equal to `other`.
type EntityNewParamsNaturalPersonIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document.
	Country param.Field[string] `json:"country,required"`
	// A description of the document submitted.
	Description param.Field[string] `json:"description,required"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date" format:"date"`
	// The identifier of the File containing the document.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r EntityNewParamsNaturalPersonIdentificationOther) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityNewParamsSupplementalDocuments struct {
	// The identifier of the File containing the document.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r EntityNewParamsSupplementalDocuments) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details of the trust entity to create. Required if `structure` is equal to
// `trust`.
type EntityNewParamsTrust struct {
	// The legal name of the trust.
	Name param.Field[string] `json:"name,required"`
	// Whether the trust is `revocable` or `irrevocable`. Irrevocable trusts require
	// their own Employer Identification Number. Revocable trusts require information
	// about the individual `grantor` who created the trust.
	Category param.Field[EntityNewParamsTrustCategory] `json:"category,required"`
	// The Employer Identification Number (EIN) for the trust. Required if `category`
	// is equal to `irrevocable`.
	TaxIdentifier param.Field[string] `json:"tax_identifier"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state in
	// which the trust was formed.
	FormationState param.Field[string] `json:"formation_state"`
	// The trust's address.
	Address param.Field[EntityNewParamsTrustAddress] `json:"address,required"`
	// The identifier of the File containing the formation document of the trust.
	FormationDocumentFileID param.Field[string] `json:"formation_document_file_id"`
	// The trustees of the trust.
	Trustees param.Field[[]EntityNewParamsTrustTrustees] `json:"trustees,required"`
	// The grantor of the trust. Required if `category` is equal to `revocable`.
	Grantor param.Field[EntityNewParamsTrustGrantor] `json:"grantor"`
}

func (r EntityNewParamsTrust) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the trust is `revocable` or `irrevocable`. Irrevocable trusts require
// their own Employer Identification Number. Revocable trusts require information
// about the individual `grantor` who created the trust.
type EntityNewParamsTrustCategory string

const (
	EntityNewParamsTrustCategoryRevocable   EntityNewParamsTrustCategory = "revocable"
	EntityNewParamsTrustCategoryIrrevocable EntityNewParamsTrustCategory = "irrevocable"
)

// The trust's address.
type EntityNewParamsTrustAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
	// The city of the address.
	City param.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State param.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip param.Field[string] `json:"zip,required"`
}

func (r EntityNewParamsTrustAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityNewParamsTrustTrustees struct {
	// The structure of the trustee.
	Structure param.Field[EntityNewParamsTrustTrusteesStructure] `json:"structure,required"`
	// Details of the individual trustee. Required when the trustee `structure` is
	// equal to `individual`.
	Individual param.Field[EntityNewParamsTrustTrusteesIndividual] `json:"individual"`
}

func (r EntityNewParamsTrustTrustees) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The structure of the trustee.
type EntityNewParamsTrustTrusteesStructure string

const (
	EntityNewParamsTrustTrusteesStructureIndividual EntityNewParamsTrustTrusteesStructure = "individual"
)

// Details of the individual trustee. Required when the trustee `structure` is
// equal to `individual`.
type EntityNewParamsTrustTrusteesIndividual struct {
	// The person's legal name.
	Name param.Field[string] `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth param.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// The individual's address.
	Address param.Field[EntityNewParamsTrustTrusteesIndividualAddress] `json:"address,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID param.Field[bool] `json:"confirmed_no_us_tax_id"`
	// A means of verifying the person's identity.
	Identification param.Field[EntityNewParamsTrustTrusteesIndividualIdentification] `json:"identification,required"`
}

func (r EntityNewParamsTrustTrusteesIndividual) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The individual's address.
type EntityNewParamsTrustTrusteesIndividualAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
	// The city of the address.
	City param.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State param.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip param.Field[string] `json:"zip,required"`
}

func (r EntityNewParamsTrustTrusteesIndividualAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A means of verifying the person's identity.
type EntityNewParamsTrustTrusteesIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method param.Field[EntityNewParamsTrustTrusteesIndividualIdentificationMethod] `json:"method,required"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number param.Field[string] `json:"number,required"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport param.Field[EntityNewParamsTrustTrusteesIndividualIdentificationPassport] `json:"passport"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense param.Field[EntityNewParamsTrustTrusteesIndividualIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other param.Field[EntityNewParamsTrustTrusteesIndividualIdentificationOther] `json:"other"`
}

func (r EntityNewParamsTrustTrusteesIndividualIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A method that can be used to verify the individual's identity.
type EntityNewParamsTrustTrusteesIndividualIdentificationMethod string

const (
	EntityNewParamsTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber                   EntityNewParamsTrustTrusteesIndividualIdentificationMethod = "social_security_number"
	EntityNewParamsTrustTrusteesIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber EntityNewParamsTrustTrusteesIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	EntityNewParamsTrustTrusteesIndividualIdentificationMethodPassport                               EntityNewParamsTrustTrusteesIndividualIdentificationMethod = "passport"
	EntityNewParamsTrustTrusteesIndividualIdentificationMethodDriversLicense                         EntityNewParamsTrustTrusteesIndividualIdentificationMethod = "drivers_license"
	EntityNewParamsTrustTrusteesIndividualIdentificationMethodOther                                  EntityNewParamsTrustTrusteesIndividualIdentificationMethod = "other"
)

// Information about the passport used for identification. Required if `method` is
// equal to `passport`.
type EntityNewParamsTrustTrusteesIndividualIdentificationPassport struct {
	// The identifier of the File containing the passport.
	FileID param.Field[string] `json:"file_id,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The country that issued the passport.
	Country param.Field[string] `json:"country,required"`
}

func (r EntityNewParamsTrustTrusteesIndividualIdentificationPassport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information about the United States driver's license used for identification.
// Required if `method` is equal to `drivers_license`.
type EntityNewParamsTrustTrusteesIndividualIdentificationDriversLicense struct {
	// The identifier of the File containing the driver's license.
	FileID param.Field[string] `json:"file_id,required"`
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The state that issued the provided driver's license.
	State param.Field[string] `json:"state,required"`
}

func (r EntityNewParamsTrustTrusteesIndividualIdentificationDriversLicense) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information about the identification document provided. Required if `method` is
// equal to `other`.
type EntityNewParamsTrustTrusteesIndividualIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document.
	Country param.Field[string] `json:"country,required"`
	// A description of the document submitted.
	Description param.Field[string] `json:"description,required"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date" format:"date"`
	// The identifier of the File containing the document.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r EntityNewParamsTrustTrusteesIndividualIdentificationOther) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The grantor of the trust. Required if `category` is equal to `revocable`.
type EntityNewParamsTrustGrantor struct {
	// The person's legal name.
	Name param.Field[string] `json:"name,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth param.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// The individual's address.
	Address param.Field[EntityNewParamsTrustGrantorAddress] `json:"address,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID param.Field[bool] `json:"confirmed_no_us_tax_id"`
	// A means of verifying the person's identity.
	Identification param.Field[EntityNewParamsTrustGrantorIdentification] `json:"identification,required"`
}

func (r EntityNewParamsTrustGrantor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The individual's address.
type EntityNewParamsTrustGrantorAddress struct {
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
	// The city of the address.
	City param.Field[string] `json:"city,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State param.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip param.Field[string] `json:"zip,required"`
}

func (r EntityNewParamsTrustGrantorAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A means of verifying the person's identity.
type EntityNewParamsTrustGrantorIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method param.Field[EntityNewParamsTrustGrantorIdentificationMethod] `json:"method,required"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number param.Field[string] `json:"number,required"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport param.Field[EntityNewParamsTrustGrantorIdentificationPassport] `json:"passport"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense param.Field[EntityNewParamsTrustGrantorIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other param.Field[EntityNewParamsTrustGrantorIdentificationOther] `json:"other"`
}

func (r EntityNewParamsTrustGrantorIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A method that can be used to verify the individual's identity.
type EntityNewParamsTrustGrantorIdentificationMethod string

const (
	EntityNewParamsTrustGrantorIdentificationMethodSocialSecurityNumber                   EntityNewParamsTrustGrantorIdentificationMethod = "social_security_number"
	EntityNewParamsTrustGrantorIdentificationMethodIndividualTaxpayerIdentificationNumber EntityNewParamsTrustGrantorIdentificationMethod = "individual_taxpayer_identification_number"
	EntityNewParamsTrustGrantorIdentificationMethodPassport                               EntityNewParamsTrustGrantorIdentificationMethod = "passport"
	EntityNewParamsTrustGrantorIdentificationMethodDriversLicense                         EntityNewParamsTrustGrantorIdentificationMethod = "drivers_license"
	EntityNewParamsTrustGrantorIdentificationMethodOther                                  EntityNewParamsTrustGrantorIdentificationMethod = "other"
)

// Information about the passport used for identification. Required if `method` is
// equal to `passport`.
type EntityNewParamsTrustGrantorIdentificationPassport struct {
	// The identifier of the File containing the passport.
	FileID param.Field[string] `json:"file_id,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The country that issued the passport.
	Country param.Field[string] `json:"country,required"`
}

func (r EntityNewParamsTrustGrantorIdentificationPassport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information about the United States driver's license used for identification.
// Required if `method` is equal to `drivers_license`.
type EntityNewParamsTrustGrantorIdentificationDriversLicense struct {
	// The identifier of the File containing the driver's license.
	FileID param.Field[string] `json:"file_id,required"`
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The state that issued the provided driver's license.
	State param.Field[string] `json:"state,required"`
}

func (r EntityNewParamsTrustGrantorIdentificationDriversLicense) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information about the identification document provided. Required if `method` is
// equal to `other`.
type EntityNewParamsTrustGrantorIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document.
	Country param.Field[string] `json:"country,required"`
	// A description of the document submitted.
	Description param.Field[string] `json:"description,required"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date" format:"date"`
	// The identifier of the File containing the document.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r EntityNewParamsTrustGrantorIdentificationOther) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityListParams struct {
	CreatedAt param.Field[EntityListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [EntityListParams]'s query parameters as `url.Values`.
func (r EntityListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EntityListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After param.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter param.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore param.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes [EntityListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r EntityListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// A list of Entity objects
type EntityListResponse struct {
	// The contents of the list.
	Data []Entity `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       entityListResponseJSON
}

// entityListResponseJSON contains the JSON metadata for the struct
// [EntityListResponse]
type entityListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
