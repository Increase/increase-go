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
	BeneficialOwners      *EntityBeneficialOwnerService
	SupplementalDocuments *EntitySupplementalDocumentService
}

// NewEntityService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEntityService(opts ...option.RequestOption) (r *EntityService) {
	r = &EntityService{}
	r.Options = opts
	r.BeneficialOwners = NewEntityBeneficialOwnerService(opts...)
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

// Archive an Entity
func (r *EntityService) Archive(ctx context.Context, entityID string, opts ...option.RequestOption) (res *Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("entities/%s/archive", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Update a Natural Person or Corporation's address
func (r *EntityService) UpdateAddress(ctx context.Context, entityID string, body EntityUpdateAddressParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("entities/%s/address", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Entities are the legal entities that own accounts. They can be people,
// corporations, partnerships, or trusts.
type Entity struct {
	// The entity's identifier.
	ID string `json:"id,required"`
	// Details of the corporation entity. Will be present if `structure` is equal to
	// `corporation`.
	Corporation EntityCorporation `json:"corporation,required,nullable"`
	// The entity's description for display purposes.
	Description string `json:"description,required,nullable"`
	// Details of the joint entity. Will be present if `structure` is equal to `joint`.
	Joint EntityJoint `json:"joint,required,nullable"`
	// Details of the natural person entity. Will be present if `structure` is equal to
	// `natural_person`.
	NaturalPerson EntityNaturalPerson `json:"natural_person,required,nullable"`
	// The status of the entity.
	Status EntityStatus `json:"status,required"`
	// The entity's legal structure.
	Structure EntityStructure `json:"structure,required"`
	// Additional documentation associated with the entity. This is limited to the
	// first 10 documents for an entity. If an entity has more than 10 documents, use
	// the GET /entity_supplemental_documents list endpoint to retrieve them.
	SupplementalDocuments []EntitySupplementalDocument `json:"supplemental_documents,required"`
	// Details of the trust entity. Will be present if `structure` is equal to `trust`.
	Trust EntityTrust `json:"trust,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `entity`.
	Type EntityType `json:"type,required"`
	JSON entityJSON `json:"-"`
}

// entityJSON contains the JSON metadata for the struct [Entity]
type entityJSON struct {
	ID                    apijson.Field
	Corporation           apijson.Field
	Description           apijson.Field
	Joint                 apijson.Field
	NaturalPerson         apijson.Field
	Status                apijson.Field
	Structure             apijson.Field
	SupplementalDocuments apijson.Field
	Trust                 apijson.Field
	Type                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Entity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details of the corporation entity. Will be present if `structure` is equal to
// `corporation`.
type EntityCorporation struct {
	// The corporation's address.
	Address EntityCorporationAddress `json:"address,required"`
	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwners []EntityCorporationBeneficialOwner `json:"beneficial_owners,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the
	// corporation's state of incorporation.
	IncorporationState string `json:"incorporation_state,required,nullable"`
	// The legal name of the corporation.
	Name string `json:"name,required"`
	// The Employer Identification Number (EIN) for the corporation.
	TaxIdentifier string `json:"tax_identifier,required,nullable"`
	// The website of the corporation.
	Website string                `json:"website,required,nullable"`
	JSON    entityCorporationJSON `json:"-"`
}

// entityCorporationJSON contains the JSON metadata for the struct
// [EntityCorporation]
type entityCorporationJSON struct {
	Address            apijson.Field
	BeneficialOwners   apijson.Field
	IncorporationState apijson.Field
	Name               apijson.Field
	TaxIdentifier      apijson.Field
	Website            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EntityCorporation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The corporation's address.
type EntityCorporationAddress struct {
	// The city of the address.
	City string `json:"city,required"`
	// The first line of the address.
	Line1 string `json:"line1,required"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state,required"`
	// The ZIP code of the address.
	Zip  string                       `json:"zip,required"`
	JSON entityCorporationAddressJSON `json:"-"`
}

// entityCorporationAddressJSON contains the JSON metadata for the struct
// [EntityCorporationAddress]
type entityCorporationAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityCorporationAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityCorporationBeneficialOwner struct {
	// The identifier of this beneficial owner.
	BeneficialOwnerID string `json:"beneficial_owner_id,required"`
	// This person's role or title within the entity.
	CompanyTitle string `json:"company_title,required,nullable"`
	// Personal details for the beneficial owner.
	Individual EntityCorporationBeneficialOwnersIndividual `json:"individual,required"`
	// Why this person is considered a beneficial owner of the entity.
	Prong EntityCorporationBeneficialOwnersProng `json:"prong,required"`
	JSON  entityCorporationBeneficialOwnerJSON   `json:"-"`
}

// entityCorporationBeneficialOwnerJSON contains the JSON metadata for the struct
// [EntityCorporationBeneficialOwner]
type entityCorporationBeneficialOwnerJSON struct {
	BeneficialOwnerID apijson.Field
	CompanyTitle      apijson.Field
	Individual        apijson.Field
	Prong             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EntityCorporationBeneficialOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Personal details for the beneficial owner.
type EntityCorporationBeneficialOwnersIndividual struct {
	// The person's address.
	Address EntityCorporationBeneficialOwnersIndividualAddress `json:"address,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth time.Time `json:"date_of_birth,required" format:"date"`
	// A means of verifying the person's identity.
	Identification EntityCorporationBeneficialOwnersIndividualIdentification `json:"identification,required"`
	// The person's legal name.
	Name string                                          `json:"name,required"`
	JSON entityCorporationBeneficialOwnersIndividualJSON `json:"-"`
}

// entityCorporationBeneficialOwnersIndividualJSON contains the JSON metadata for
// the struct [EntityCorporationBeneficialOwnersIndividual]
type entityCorporationBeneficialOwnersIndividualJSON struct {
	Address        apijson.Field
	DateOfBirth    apijson.Field
	Identification apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityCorporationBeneficialOwnersIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The person's address.
type EntityCorporationBeneficialOwnersIndividualAddress struct {
	// The city of the address.
	City string `json:"city,required"`
	// The first line of the address.
	Line1 string `json:"line1,required"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state,required"`
	// The ZIP code of the address.
	Zip  string                                                 `json:"zip,required"`
	JSON entityCorporationBeneficialOwnersIndividualAddressJSON `json:"-"`
}

// entityCorporationBeneficialOwnersIndividualAddressJSON contains the JSON
// metadata for the struct [EntityCorporationBeneficialOwnersIndividualAddress]
type entityCorporationBeneficialOwnersIndividualAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
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
	NumberLast4 string                                                        `json:"number_last4,required"`
	JSON        entityCorporationBeneficialOwnersIndividualIdentificationJSON `json:"-"`
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
	// A social security number.
	EntityCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber EntityCorporationBeneficialOwnersIndividualIdentificationMethod = "social_security_number"
	// An individual taxpayer identification number (ITIN).
	EntityCorporationBeneficialOwnersIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber EntityCorporationBeneficialOwnersIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	// A passport number.
	EntityCorporationBeneficialOwnersIndividualIdentificationMethodPassport EntityCorporationBeneficialOwnersIndividualIdentificationMethod = "passport"
	// A driver's license number.
	EntityCorporationBeneficialOwnersIndividualIdentificationMethodDriversLicense EntityCorporationBeneficialOwnersIndividualIdentificationMethod = "drivers_license"
	// Another identifying document.
	EntityCorporationBeneficialOwnersIndividualIdentificationMethodOther EntityCorporationBeneficialOwnersIndividualIdentificationMethod = "other"
)

// Why this person is considered a beneficial owner of the entity.
type EntityCorporationBeneficialOwnersProng string

const (
	// A person with 25% or greater direct or indirect ownership of the entity.
	EntityCorporationBeneficialOwnersProngOwnership EntityCorporationBeneficialOwnersProng = "ownership"
	// A person who manages, directs, or has significant control of the entity.
	EntityCorporationBeneficialOwnersProngControl EntityCorporationBeneficialOwnersProng = "control"
)

// Details of the joint entity. Will be present if `structure` is equal to `joint`.
type EntityJoint struct {
	// The two individuals that share control of the entity.
	Individuals []EntityJointIndividual `json:"individuals,required"`
	// The entity's name.
	Name string          `json:"name,required"`
	JSON entityJointJSON `json:"-"`
}

// entityJointJSON contains the JSON metadata for the struct [EntityJoint]
type entityJointJSON struct {
	Individuals apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityJoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EntityJointIndividual struct {
	// The person's address.
	Address EntityJointIndividualsAddress `json:"address,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth time.Time `json:"date_of_birth,required" format:"date"`
	// A means of verifying the person's identity.
	Identification EntityJointIndividualsIdentification `json:"identification,required"`
	// The person's legal name.
	Name string                    `json:"name,required"`
	JSON entityJointIndividualJSON `json:"-"`
}

// entityJointIndividualJSON contains the JSON metadata for the struct
// [EntityJointIndividual]
type entityJointIndividualJSON struct {
	Address        apijson.Field
	DateOfBirth    apijson.Field
	Identification apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityJointIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The person's address.
type EntityJointIndividualsAddress struct {
	// The city of the address.
	City string `json:"city,required"`
	// The first line of the address.
	Line1 string `json:"line1,required"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state,required"`
	// The ZIP code of the address.
	Zip  string                            `json:"zip,required"`
	JSON entityJointIndividualsAddressJSON `json:"-"`
}

// entityJointIndividualsAddressJSON contains the JSON metadata for the struct
// [EntityJointIndividualsAddress]
type entityJointIndividualsAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
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
	NumberLast4 string                                   `json:"number_last4,required"`
	JSON        entityJointIndividualsIdentificationJSON `json:"-"`
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
	// A social security number.
	EntityJointIndividualsIdentificationMethodSocialSecurityNumber EntityJointIndividualsIdentificationMethod = "social_security_number"
	// An individual taxpayer identification number (ITIN).
	EntityJointIndividualsIdentificationMethodIndividualTaxpayerIdentificationNumber EntityJointIndividualsIdentificationMethod = "individual_taxpayer_identification_number"
	// A passport number.
	EntityJointIndividualsIdentificationMethodPassport EntityJointIndividualsIdentificationMethod = "passport"
	// A driver's license number.
	EntityJointIndividualsIdentificationMethodDriversLicense EntityJointIndividualsIdentificationMethod = "drivers_license"
	// Another identifying document.
	EntityJointIndividualsIdentificationMethodOther EntityJointIndividualsIdentificationMethod = "other"
)

// Details of the natural person entity. Will be present if `structure` is equal to
// `natural_person`.
type EntityNaturalPerson struct {
	// The person's address.
	Address EntityNaturalPersonAddress `json:"address,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth time.Time `json:"date_of_birth,required" format:"date"`
	// A means of verifying the person's identity.
	Identification EntityNaturalPersonIdentification `json:"identification,required"`
	// The person's legal name.
	Name string                  `json:"name,required"`
	JSON entityNaturalPersonJSON `json:"-"`
}

// entityNaturalPersonJSON contains the JSON metadata for the struct
// [EntityNaturalPerson]
type entityNaturalPersonJSON struct {
	Address        apijson.Field
	DateOfBirth    apijson.Field
	Identification apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityNaturalPerson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The person's address.
type EntityNaturalPersonAddress struct {
	// The city of the address.
	City string `json:"city,required"`
	// The first line of the address.
	Line1 string `json:"line1,required"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state,required"`
	// The ZIP code of the address.
	Zip  string                         `json:"zip,required"`
	JSON entityNaturalPersonAddressJSON `json:"-"`
}

// entityNaturalPersonAddressJSON contains the JSON metadata for the struct
// [EntityNaturalPersonAddress]
type entityNaturalPersonAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
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
	NumberLast4 string                                `json:"number_last4,required"`
	JSON        entityNaturalPersonIdentificationJSON `json:"-"`
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
	// A social security number.
	EntityNaturalPersonIdentificationMethodSocialSecurityNumber EntityNaturalPersonIdentificationMethod = "social_security_number"
	// An individual taxpayer identification number (ITIN).
	EntityNaturalPersonIdentificationMethodIndividualTaxpayerIdentificationNumber EntityNaturalPersonIdentificationMethod = "individual_taxpayer_identification_number"
	// A passport number.
	EntityNaturalPersonIdentificationMethodPassport EntityNaturalPersonIdentificationMethod = "passport"
	// A driver's license number.
	EntityNaturalPersonIdentificationMethodDriversLicense EntityNaturalPersonIdentificationMethod = "drivers_license"
	// Another identifying document.
	EntityNaturalPersonIdentificationMethodOther EntityNaturalPersonIdentificationMethod = "other"
)

// The status of the entity.
type EntityStatus string

const (
	// The entity is active.
	EntityStatusActive EntityStatus = "active"
	// The entity is archived, and can no longer be used to create accounts.
	EntityStatusArchived EntityStatus = "archived"
	// The entity is temporarily disabled and cannot be used for financial activity.
	EntityStatusDisabled EntityStatus = "disabled"
)

// The entity's legal structure.
type EntityStructure string

const (
	// A corporation.
	EntityStructureCorporation EntityStructure = "corporation"
	// An individual person.
	EntityStructureNaturalPerson EntityStructure = "natural_person"
	// Multiple individual people.
	EntityStructureJoint EntityStructure = "joint"
	// A trust.
	EntityStructureTrust EntityStructure = "trust"
)

// Supplemental Documents are uploaded files connected to an Entity during
// onboarding.
type EntitySupplementalDocument struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the
	// Supplemental Document was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The File containing the document.
	FileID string `json:"file_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `entity_supplemental_document`.
	Type EntitySupplementalDocumentsType `json:"type,required"`
	JSON entitySupplementalDocumentJSON  `json:"-"`
}

// entitySupplementalDocumentJSON contains the JSON metadata for the struct
// [EntitySupplementalDocument]
type entitySupplementalDocumentJSON struct {
	CreatedAt   apijson.Field
	FileID      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitySupplementalDocument) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A constant representing the object's type. For this resource it will always be
// `entity_supplemental_document`.
type EntitySupplementalDocumentsType string

const (
	EntitySupplementalDocumentsTypeEntitySupplementalDocument EntitySupplementalDocumentsType = "entity_supplemental_document"
)

// Details of the trust entity. Will be present if `structure` is equal to `trust`.
type EntityTrust struct {
	// The trust's address.
	Address EntityTrustAddress `json:"address,required"`
	// Whether the trust is `revocable` or `irrevocable`.
	Category EntityTrustCategory `json:"category,required"`
	// The ID for the File containing the formation document of the trust.
	FormationDocumentFileID string `json:"formation_document_file_id,required,nullable"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state in
	// which the trust was formed.
	FormationState string `json:"formation_state,required,nullable"`
	// The grantor of the trust. Will be present if the `category` is `revocable`.
	Grantor EntityTrustGrantor `json:"grantor,required,nullable"`
	// The trust's name.
	Name string `json:"name,required"`
	// The Employer Identification Number (EIN) of the trust itself.
	TaxIdentifier string `json:"tax_identifier,required,nullable"`
	// The trustees of the trust.
	Trustees []EntityTrustTrustee `json:"trustees,required"`
	JSON     entityTrustJSON      `json:"-"`
}

// entityTrustJSON contains the JSON metadata for the struct [EntityTrust]
type entityTrustJSON struct {
	Address                 apijson.Field
	Category                apijson.Field
	FormationDocumentFileID apijson.Field
	FormationState          apijson.Field
	Grantor                 apijson.Field
	Name                    apijson.Field
	TaxIdentifier           apijson.Field
	Trustees                apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *EntityTrust) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The trust's address.
type EntityTrustAddress struct {
	// The city of the address.
	City string `json:"city,required"`
	// The first line of the address.
	Line1 string `json:"line1,required"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state,required"`
	// The ZIP code of the address.
	Zip  string                 `json:"zip,required"`
	JSON entityTrustAddressJSON `json:"-"`
}

// entityTrustAddressJSON contains the JSON metadata for the struct
// [EntityTrustAddress]
type entityTrustAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityTrustAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the trust is `revocable` or `irrevocable`.
type EntityTrustCategory string

const (
	// The trust is revocable by the grantor.
	EntityTrustCategoryRevocable EntityTrustCategory = "revocable"
	// The trust cannot be revoked.
	EntityTrustCategoryIrrevocable EntityTrustCategory = "irrevocable"
)

// The grantor of the trust. Will be present if the `category` is `revocable`.
type EntityTrustGrantor struct {
	// The person's address.
	Address EntityTrustGrantorAddress `json:"address,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth time.Time `json:"date_of_birth,required" format:"date"`
	// A means of verifying the person's identity.
	Identification EntityTrustGrantorIdentification `json:"identification,required"`
	// The person's legal name.
	Name string                 `json:"name,required"`
	JSON entityTrustGrantorJSON `json:"-"`
}

// entityTrustGrantorJSON contains the JSON metadata for the struct
// [EntityTrustGrantor]
type entityTrustGrantorJSON struct {
	Address        apijson.Field
	DateOfBirth    apijson.Field
	Identification apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityTrustGrantor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The person's address.
type EntityTrustGrantorAddress struct {
	// The city of the address.
	City string `json:"city,required"`
	// The first line of the address.
	Line1 string `json:"line1,required"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state,required"`
	// The ZIP code of the address.
	Zip  string                        `json:"zip,required"`
	JSON entityTrustGrantorAddressJSON `json:"-"`
}

// entityTrustGrantorAddressJSON contains the JSON metadata for the struct
// [EntityTrustGrantorAddress]
type entityTrustGrantorAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
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
	NumberLast4 string                               `json:"number_last4,required"`
	JSON        entityTrustGrantorIdentificationJSON `json:"-"`
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
	// A social security number.
	EntityTrustGrantorIdentificationMethodSocialSecurityNumber EntityTrustGrantorIdentificationMethod = "social_security_number"
	// An individual taxpayer identification number (ITIN).
	EntityTrustGrantorIdentificationMethodIndividualTaxpayerIdentificationNumber EntityTrustGrantorIdentificationMethod = "individual_taxpayer_identification_number"
	// A passport number.
	EntityTrustGrantorIdentificationMethodPassport EntityTrustGrantorIdentificationMethod = "passport"
	// A driver's license number.
	EntityTrustGrantorIdentificationMethodDriversLicense EntityTrustGrantorIdentificationMethod = "drivers_license"
	// Another identifying document.
	EntityTrustGrantorIdentificationMethodOther EntityTrustGrantorIdentificationMethod = "other"
)

type EntityTrustTrustee struct {
	// The individual trustee of the trust. Will be present if the trustee's
	// `structure` is equal to `individual`.
	Individual EntityTrustTrusteesIndividual `json:"individual,required,nullable"`
	// The structure of the trustee. Will always be equal to `individual`.
	Structure EntityTrustTrusteesStructure `json:"structure,required"`
	JSON      entityTrustTrusteeJSON       `json:"-"`
}

// entityTrustTrusteeJSON contains the JSON metadata for the struct
// [EntityTrustTrustee]
type entityTrustTrusteeJSON struct {
	Individual  apijson.Field
	Structure   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityTrustTrustee) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The individual trustee of the trust. Will be present if the trustee's
// `structure` is equal to `individual`.
type EntityTrustTrusteesIndividual struct {
	// The person's address.
	Address EntityTrustTrusteesIndividualAddress `json:"address,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth time.Time `json:"date_of_birth,required" format:"date"`
	// A means of verifying the person's identity.
	Identification EntityTrustTrusteesIndividualIdentification `json:"identification,required"`
	// The person's legal name.
	Name string                            `json:"name,required"`
	JSON entityTrustTrusteesIndividualJSON `json:"-"`
}

// entityTrustTrusteesIndividualJSON contains the JSON metadata for the struct
// [EntityTrustTrusteesIndividual]
type entityTrustTrusteesIndividualJSON struct {
	Address        apijson.Field
	DateOfBirth    apijson.Field
	Identification apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityTrustTrusteesIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The person's address.
type EntityTrustTrusteesIndividualAddress struct {
	// The city of the address.
	City string `json:"city,required"`
	// The first line of the address.
	Line1 string `json:"line1,required"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string `json:"state,required"`
	// The ZIP code of the address.
	Zip  string                                   `json:"zip,required"`
	JSON entityTrustTrusteesIndividualAddressJSON `json:"-"`
}

// entityTrustTrusteesIndividualAddressJSON contains the JSON metadata for the
// struct [EntityTrustTrusteesIndividualAddress]
type entityTrustTrusteesIndividualAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
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
	NumberLast4 string                                          `json:"number_last4,required"`
	JSON        entityTrustTrusteesIndividualIdentificationJSON `json:"-"`
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
	// A social security number.
	EntityTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber EntityTrustTrusteesIndividualIdentificationMethod = "social_security_number"
	// An individual taxpayer identification number (ITIN).
	EntityTrustTrusteesIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber EntityTrustTrusteesIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	// A passport number.
	EntityTrustTrusteesIndividualIdentificationMethodPassport EntityTrustTrusteesIndividualIdentificationMethod = "passport"
	// A driver's license number.
	EntityTrustTrusteesIndividualIdentificationMethodDriversLicense EntityTrustTrusteesIndividualIdentificationMethod = "drivers_license"
	// Another identifying document.
	EntityTrustTrusteesIndividualIdentificationMethodOther EntityTrustTrusteesIndividualIdentificationMethod = "other"
)

// The structure of the trustee. Will always be equal to `individual`.
type EntityTrustTrusteesStructure string

const (
	// The trustee is an individual.
	EntityTrustTrusteesStructureIndividual EntityTrustTrusteesStructure = "individual"
)

// A constant representing the object's type. For this resource it will always be
// `entity`.
type EntityType string

const (
	EntityTypeEntity EntityType = "entity"
)

type EntityNewParams struct {
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
	// The relationship between your group and the entity.
	Relationship param.Field[EntityNewParamsRelationship] `json:"relationship"`
	// Additional documentation associated with the entity.
	SupplementalDocuments param.Field[[]EntityNewParamsSupplementalDocument] `json:"supplemental_documents"`
	// Details of the trust entity to create. Required if `structure` is equal to
	// `trust`.
	Trust param.Field[EntityNewParamsTrust] `json:"trust"`
}

func (r EntityNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of Entity to create.
type EntityNewParamsStructure string

const (
	// A corporation.
	EntityNewParamsStructureCorporation EntityNewParamsStructure = "corporation"
	// An individual person.
	EntityNewParamsStructureNaturalPerson EntityNewParamsStructure = "natural_person"
	// Multiple individual people.
	EntityNewParamsStructureJoint EntityNewParamsStructure = "joint"
	// A trust.
	EntityNewParamsStructureTrust EntityNewParamsStructure = "trust"
)

// Details of the corporation entity to create. Required if `structure` is equal to
// `corporation`.
type EntityNewParamsCorporation struct {
	// The corporation's physical address. Post Office Boxes are disallowed.
	Address param.Field[EntityNewParamsCorporationAddress] `json:"address,required"`
	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwners param.Field[[]EntityNewParamsCorporationBeneficialOwner] `json:"beneficial_owners,required"`
	// The legal name of the corporation.
	Name param.Field[string] `json:"name,required"`
	// The Employer Identification Number (EIN) for the corporation.
	TaxIdentifier param.Field[string] `json:"tax_identifier,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the
	// corporation's state of incorporation.
	IncorporationState param.Field[string] `json:"incorporation_state"`
	// The website of the corporation.
	Website param.Field[string] `json:"website"`
}

func (r EntityNewParamsCorporation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The corporation's physical address. Post Office Boxes are disallowed.
type EntityNewParamsCorporationAddress struct {
	// The city of the address.
	City param.Field[string] `json:"city,required"`
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State param.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip param.Field[string] `json:"zip,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
}

func (r EntityNewParamsCorporationAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityNewParamsCorporationBeneficialOwner struct {
	// Personal details for the beneficial owner.
	Individual param.Field[EntityNewParamsCorporationBeneficialOwnersIndividual] `json:"individual,required"`
	// Why this person is considered a beneficial owner of the entity. At least one
	// option is required, if a person is both a control person and owner, submit an
	// array containing both.
	Prongs param.Field[[]EntityNewParamsCorporationBeneficialOwnersProng] `json:"prongs,required"`
	// This person's role or title within the entity.
	CompanyTitle param.Field[string] `json:"company_title"`
}

func (r EntityNewParamsCorporationBeneficialOwner) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Personal details for the beneficial owner.
type EntityNewParamsCorporationBeneficialOwnersIndividual struct {
	// The individual's physical address. Post Office Boxes are disallowed.
	Address param.Field[EntityNewParamsCorporationBeneficialOwnersIndividualAddress] `json:"address,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth param.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// A means of verifying the person's identity.
	Identification param.Field[EntityNewParamsCorporationBeneficialOwnersIndividualIdentification] `json:"identification,required"`
	// The person's legal name.
	Name param.Field[string] `json:"name,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID param.Field[bool] `json:"confirmed_no_us_tax_id"`
}

func (r EntityNewParamsCorporationBeneficialOwnersIndividual) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The individual's physical address. Post Office Boxes are disallowed.
type EntityNewParamsCorporationBeneficialOwnersIndividualAddress struct {
	// The city of the address.
	City param.Field[string] `json:"city,required"`
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State param.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip param.Field[string] `json:"zip,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
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
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense param.Field[EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other param.Field[EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationOther] `json:"other"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport param.Field[EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationPassport] `json:"passport"`
}

func (r EntityNewParamsCorporationBeneficialOwnersIndividualIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A method that can be used to verify the individual's identity.
type EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod string

const (
	// A social security number.
	EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod = "social_security_number"
	// An individual taxpayer identification number (ITIN).
	EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	// A passport number.
	EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodPassport EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod = "passport"
	// A driver's license number.
	EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodDriversLicense EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod = "drivers_license"
	// Another identifying document.
	EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodOther EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod = "other"
)

// Information about the United States driver's license used for identification.
// Required if `method` is equal to `drivers_license`.
type EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationDriversLicense struct {
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The identifier of the File containing the front of the driver's license.
	FileID param.Field[string] `json:"file_id,required"`
	// The state that issued the provided driver's license.
	State param.Field[string] `json:"state,required"`
	// The identifier of the File containing the back of the driver's license.
	BackFileID param.Field[string] `json:"back_file_id"`
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
	// The identifier of the File containing the front of the document.
	FileID param.Field[string] `json:"file_id,required"`
	// The identifier of the File containing the back of the document. Not every
	// document has a reverse side.
	BackFileID param.Field[string] `json:"back_file_id"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date" format:"date"`
}

func (r EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationOther) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information about the passport used for identification. Required if `method` is
// equal to `passport`.
type EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationPassport struct {
	// The country that issued the passport.
	Country param.Field[string] `json:"country,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The identifier of the File containing the passport.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationPassport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityNewParamsCorporationBeneficialOwnersProng string

const (
	// A person with 25% or greater direct or indirect ownership of the entity.
	EntityNewParamsCorporationBeneficialOwnersProngOwnership EntityNewParamsCorporationBeneficialOwnersProng = "ownership"
	// A person who manages, directs, or has significant control of the entity.
	EntityNewParamsCorporationBeneficialOwnersProngControl EntityNewParamsCorporationBeneficialOwnersProng = "control"
)

// Details of the joint entity to create. Required if `structure` is equal to
// `joint`.
type EntityNewParamsJoint struct {
	// The two individuals that share control of the entity.
	Individuals param.Field[[]EntityNewParamsJointIndividual] `json:"individuals,required"`
	// The name of the joint entity.
	Name param.Field[string] `json:"name"`
}

func (r EntityNewParamsJoint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityNewParamsJointIndividual struct {
	// The individual's physical address. Post Office Boxes are disallowed.
	Address param.Field[EntityNewParamsJointIndividualsAddress] `json:"address,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth param.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// A means of verifying the person's identity.
	Identification param.Field[EntityNewParamsJointIndividualsIdentification] `json:"identification,required"`
	// The person's legal name.
	Name param.Field[string] `json:"name,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID param.Field[bool] `json:"confirmed_no_us_tax_id"`
}

func (r EntityNewParamsJointIndividual) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The individual's physical address. Post Office Boxes are disallowed.
type EntityNewParamsJointIndividualsAddress struct {
	// The city of the address.
	City param.Field[string] `json:"city,required"`
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State param.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip param.Field[string] `json:"zip,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
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
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense param.Field[EntityNewParamsJointIndividualsIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other param.Field[EntityNewParamsJointIndividualsIdentificationOther] `json:"other"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport param.Field[EntityNewParamsJointIndividualsIdentificationPassport] `json:"passport"`
}

func (r EntityNewParamsJointIndividualsIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A method that can be used to verify the individual's identity.
type EntityNewParamsJointIndividualsIdentificationMethod string

const (
	// A social security number.
	EntityNewParamsJointIndividualsIdentificationMethodSocialSecurityNumber EntityNewParamsJointIndividualsIdentificationMethod = "social_security_number"
	// An individual taxpayer identification number (ITIN).
	EntityNewParamsJointIndividualsIdentificationMethodIndividualTaxpayerIdentificationNumber EntityNewParamsJointIndividualsIdentificationMethod = "individual_taxpayer_identification_number"
	// A passport number.
	EntityNewParamsJointIndividualsIdentificationMethodPassport EntityNewParamsJointIndividualsIdentificationMethod = "passport"
	// A driver's license number.
	EntityNewParamsJointIndividualsIdentificationMethodDriversLicense EntityNewParamsJointIndividualsIdentificationMethod = "drivers_license"
	// Another identifying document.
	EntityNewParamsJointIndividualsIdentificationMethodOther EntityNewParamsJointIndividualsIdentificationMethod = "other"
)

// Information about the United States driver's license used for identification.
// Required if `method` is equal to `drivers_license`.
type EntityNewParamsJointIndividualsIdentificationDriversLicense struct {
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The identifier of the File containing the front of the driver's license.
	FileID param.Field[string] `json:"file_id,required"`
	// The state that issued the provided driver's license.
	State param.Field[string] `json:"state,required"`
	// The identifier of the File containing the back of the driver's license.
	BackFileID param.Field[string] `json:"back_file_id"`
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
	// The identifier of the File containing the front of the document.
	FileID param.Field[string] `json:"file_id,required"`
	// The identifier of the File containing the back of the document. Not every
	// document has a reverse side.
	BackFileID param.Field[string] `json:"back_file_id"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date" format:"date"`
}

func (r EntityNewParamsJointIndividualsIdentificationOther) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information about the passport used for identification. Required if `method` is
// equal to `passport`.
type EntityNewParamsJointIndividualsIdentificationPassport struct {
	// The country that issued the passport.
	Country param.Field[string] `json:"country,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The identifier of the File containing the passport.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r EntityNewParamsJointIndividualsIdentificationPassport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details of the natural person entity to create. Required if `structure` is equal
// to `natural_person`. Natural people entities should be submitted with
// `social_security_number` or `individual_taxpayer_identification_number`
// identification methods.
type EntityNewParamsNaturalPerson struct {
	// The individual's physical address. Post Office Boxes are disallowed.
	Address param.Field[EntityNewParamsNaturalPersonAddress] `json:"address,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth param.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// A means of verifying the person's identity.
	Identification param.Field[EntityNewParamsNaturalPersonIdentification] `json:"identification,required"`
	// The person's legal name.
	Name param.Field[string] `json:"name,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID param.Field[bool] `json:"confirmed_no_us_tax_id"`
}

func (r EntityNewParamsNaturalPerson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The individual's physical address. Post Office Boxes are disallowed.
type EntityNewParamsNaturalPersonAddress struct {
	// The city of the address.
	City param.Field[string] `json:"city,required"`
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State param.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip param.Field[string] `json:"zip,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
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
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense param.Field[EntityNewParamsNaturalPersonIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other param.Field[EntityNewParamsNaturalPersonIdentificationOther] `json:"other"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport param.Field[EntityNewParamsNaturalPersonIdentificationPassport] `json:"passport"`
}

func (r EntityNewParamsNaturalPersonIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A method that can be used to verify the individual's identity.
type EntityNewParamsNaturalPersonIdentificationMethod string

const (
	// A social security number.
	EntityNewParamsNaturalPersonIdentificationMethodSocialSecurityNumber EntityNewParamsNaturalPersonIdentificationMethod = "social_security_number"
	// An individual taxpayer identification number (ITIN).
	EntityNewParamsNaturalPersonIdentificationMethodIndividualTaxpayerIdentificationNumber EntityNewParamsNaturalPersonIdentificationMethod = "individual_taxpayer_identification_number"
	// A passport number.
	EntityNewParamsNaturalPersonIdentificationMethodPassport EntityNewParamsNaturalPersonIdentificationMethod = "passport"
	// A driver's license number.
	EntityNewParamsNaturalPersonIdentificationMethodDriversLicense EntityNewParamsNaturalPersonIdentificationMethod = "drivers_license"
	// Another identifying document.
	EntityNewParamsNaturalPersonIdentificationMethodOther EntityNewParamsNaturalPersonIdentificationMethod = "other"
)

// Information about the United States driver's license used for identification.
// Required if `method` is equal to `drivers_license`.
type EntityNewParamsNaturalPersonIdentificationDriversLicense struct {
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The identifier of the File containing the front of the driver's license.
	FileID param.Field[string] `json:"file_id,required"`
	// The state that issued the provided driver's license.
	State param.Field[string] `json:"state,required"`
	// The identifier of the File containing the back of the driver's license.
	BackFileID param.Field[string] `json:"back_file_id"`
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
	// The identifier of the File containing the front of the document.
	FileID param.Field[string] `json:"file_id,required"`
	// The identifier of the File containing the back of the document. Not every
	// document has a reverse side.
	BackFileID param.Field[string] `json:"back_file_id"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date" format:"date"`
}

func (r EntityNewParamsNaturalPersonIdentificationOther) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information about the passport used for identification. Required if `method` is
// equal to `passport`.
type EntityNewParamsNaturalPersonIdentificationPassport struct {
	// The country that issued the passport.
	Country param.Field[string] `json:"country,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The identifier of the File containing the passport.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r EntityNewParamsNaturalPersonIdentificationPassport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The relationship between your group and the entity.
type EntityNewParamsRelationship string

const (
	// The entity is controlled by your group.
	EntityNewParamsRelationshipAffiliated EntityNewParamsRelationship = "affiliated"
	// The entity is for informational purposes only.
	EntityNewParamsRelationshipInformational EntityNewParamsRelationship = "informational"
	// The entity is not controlled by your group, but can still directly open
	// accounts.
	EntityNewParamsRelationshipUnaffiliated EntityNewParamsRelationship = "unaffiliated"
)

type EntityNewParamsSupplementalDocument struct {
	// The identifier of the File containing the document.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r EntityNewParamsSupplementalDocument) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details of the trust entity to create. Required if `structure` is equal to
// `trust`.
type EntityNewParamsTrust struct {
	// The trust's physical address. Post Office Boxes are disallowed.
	Address param.Field[EntityNewParamsTrustAddress] `json:"address,required"`
	// Whether the trust is `revocable` or `irrevocable`. Irrevocable trusts require
	// their own Employer Identification Number. Revocable trusts require information
	// about the individual `grantor` who created the trust.
	Category param.Field[EntityNewParamsTrustCategory] `json:"category,required"`
	// The legal name of the trust.
	Name param.Field[string] `json:"name,required"`
	// The trustees of the trust.
	Trustees param.Field[[]EntityNewParamsTrustTrustee] `json:"trustees,required"`
	// The identifier of the File containing the formation document of the trust.
	FormationDocumentFileID param.Field[string] `json:"formation_document_file_id"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state in
	// which the trust was formed.
	FormationState param.Field[string] `json:"formation_state"`
	// The grantor of the trust. Required if `category` is equal to `revocable`.
	Grantor param.Field[EntityNewParamsTrustGrantor] `json:"grantor"`
	// The Employer Identification Number (EIN) for the trust. Required if `category`
	// is equal to `irrevocable`.
	TaxIdentifier param.Field[string] `json:"tax_identifier"`
}

func (r EntityNewParamsTrust) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The trust's physical address. Post Office Boxes are disallowed.
type EntityNewParamsTrustAddress struct {
	// The city of the address.
	City param.Field[string] `json:"city,required"`
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State param.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip param.Field[string] `json:"zip,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
}

func (r EntityNewParamsTrustAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the trust is `revocable` or `irrevocable`. Irrevocable trusts require
// their own Employer Identification Number. Revocable trusts require information
// about the individual `grantor` who created the trust.
type EntityNewParamsTrustCategory string

const (
	// The trust is revocable by the grantor.
	EntityNewParamsTrustCategoryRevocable EntityNewParamsTrustCategory = "revocable"
	// The trust cannot be revoked.
	EntityNewParamsTrustCategoryIrrevocable EntityNewParamsTrustCategory = "irrevocable"
)

type EntityNewParamsTrustTrustee struct {
	// The structure of the trustee.
	Structure param.Field[EntityNewParamsTrustTrusteesStructure] `json:"structure,required"`
	// Details of the individual trustee. Required when the trustee `structure` is
	// equal to `individual`.
	Individual param.Field[EntityNewParamsTrustTrusteesIndividual] `json:"individual"`
}

func (r EntityNewParamsTrustTrustee) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The structure of the trustee.
type EntityNewParamsTrustTrusteesStructure string

const (
	// The trustee is an individual.
	EntityNewParamsTrustTrusteesStructureIndividual EntityNewParamsTrustTrusteesStructure = "individual"
)

// Details of the individual trustee. Required when the trustee `structure` is
// equal to `individual`.
type EntityNewParamsTrustTrusteesIndividual struct {
	// The individual's physical address. Post Office Boxes are disallowed.
	Address param.Field[EntityNewParamsTrustTrusteesIndividualAddress] `json:"address,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth param.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// A means of verifying the person's identity.
	Identification param.Field[EntityNewParamsTrustTrusteesIndividualIdentification] `json:"identification,required"`
	// The person's legal name.
	Name param.Field[string] `json:"name,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID param.Field[bool] `json:"confirmed_no_us_tax_id"`
}

func (r EntityNewParamsTrustTrusteesIndividual) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The individual's physical address. Post Office Boxes are disallowed.
type EntityNewParamsTrustTrusteesIndividualAddress struct {
	// The city of the address.
	City param.Field[string] `json:"city,required"`
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State param.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip param.Field[string] `json:"zip,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
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
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense param.Field[EntityNewParamsTrustTrusteesIndividualIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other param.Field[EntityNewParamsTrustTrusteesIndividualIdentificationOther] `json:"other"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport param.Field[EntityNewParamsTrustTrusteesIndividualIdentificationPassport] `json:"passport"`
}

func (r EntityNewParamsTrustTrusteesIndividualIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A method that can be used to verify the individual's identity.
type EntityNewParamsTrustTrusteesIndividualIdentificationMethod string

const (
	// A social security number.
	EntityNewParamsTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber EntityNewParamsTrustTrusteesIndividualIdentificationMethod = "social_security_number"
	// An individual taxpayer identification number (ITIN).
	EntityNewParamsTrustTrusteesIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber EntityNewParamsTrustTrusteesIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	// A passport number.
	EntityNewParamsTrustTrusteesIndividualIdentificationMethodPassport EntityNewParamsTrustTrusteesIndividualIdentificationMethod = "passport"
	// A driver's license number.
	EntityNewParamsTrustTrusteesIndividualIdentificationMethodDriversLicense EntityNewParamsTrustTrusteesIndividualIdentificationMethod = "drivers_license"
	// Another identifying document.
	EntityNewParamsTrustTrusteesIndividualIdentificationMethodOther EntityNewParamsTrustTrusteesIndividualIdentificationMethod = "other"
)

// Information about the United States driver's license used for identification.
// Required if `method` is equal to `drivers_license`.
type EntityNewParamsTrustTrusteesIndividualIdentificationDriversLicense struct {
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The identifier of the File containing the front of the driver's license.
	FileID param.Field[string] `json:"file_id,required"`
	// The state that issued the provided driver's license.
	State param.Field[string] `json:"state,required"`
	// The identifier of the File containing the back of the driver's license.
	BackFileID param.Field[string] `json:"back_file_id"`
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
	// The identifier of the File containing the front of the document.
	FileID param.Field[string] `json:"file_id,required"`
	// The identifier of the File containing the back of the document. Not every
	// document has a reverse side.
	BackFileID param.Field[string] `json:"back_file_id"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date" format:"date"`
}

func (r EntityNewParamsTrustTrusteesIndividualIdentificationOther) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information about the passport used for identification. Required if `method` is
// equal to `passport`.
type EntityNewParamsTrustTrusteesIndividualIdentificationPassport struct {
	// The country that issued the passport.
	Country param.Field[string] `json:"country,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The identifier of the File containing the passport.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r EntityNewParamsTrustTrusteesIndividualIdentificationPassport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The grantor of the trust. Required if `category` is equal to `revocable`.
type EntityNewParamsTrustGrantor struct {
	// The individual's physical address. Post Office Boxes are disallowed.
	Address param.Field[EntityNewParamsTrustGrantorAddress] `json:"address,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth param.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// A means of verifying the person's identity.
	Identification param.Field[EntityNewParamsTrustGrantorIdentification] `json:"identification,required"`
	// The person's legal name.
	Name param.Field[string] `json:"name,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID param.Field[bool] `json:"confirmed_no_us_tax_id"`
}

func (r EntityNewParamsTrustGrantor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The individual's physical address. Post Office Boxes are disallowed.
type EntityNewParamsTrustGrantorAddress struct {
	// The city of the address.
	City param.Field[string] `json:"city,required"`
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State param.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip param.Field[string] `json:"zip,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
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
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense param.Field[EntityNewParamsTrustGrantorIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other param.Field[EntityNewParamsTrustGrantorIdentificationOther] `json:"other"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport param.Field[EntityNewParamsTrustGrantorIdentificationPassport] `json:"passport"`
}

func (r EntityNewParamsTrustGrantorIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A method that can be used to verify the individual's identity.
type EntityNewParamsTrustGrantorIdentificationMethod string

const (
	// A social security number.
	EntityNewParamsTrustGrantorIdentificationMethodSocialSecurityNumber EntityNewParamsTrustGrantorIdentificationMethod = "social_security_number"
	// An individual taxpayer identification number (ITIN).
	EntityNewParamsTrustGrantorIdentificationMethodIndividualTaxpayerIdentificationNumber EntityNewParamsTrustGrantorIdentificationMethod = "individual_taxpayer_identification_number"
	// A passport number.
	EntityNewParamsTrustGrantorIdentificationMethodPassport EntityNewParamsTrustGrantorIdentificationMethod = "passport"
	// A driver's license number.
	EntityNewParamsTrustGrantorIdentificationMethodDriversLicense EntityNewParamsTrustGrantorIdentificationMethod = "drivers_license"
	// Another identifying document.
	EntityNewParamsTrustGrantorIdentificationMethodOther EntityNewParamsTrustGrantorIdentificationMethod = "other"
)

// Information about the United States driver's license used for identification.
// Required if `method` is equal to `drivers_license`.
type EntityNewParamsTrustGrantorIdentificationDriversLicense struct {
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The identifier of the File containing the front of the driver's license.
	FileID param.Field[string] `json:"file_id,required"`
	// The state that issued the provided driver's license.
	State param.Field[string] `json:"state,required"`
	// The identifier of the File containing the back of the driver's license.
	BackFileID param.Field[string] `json:"back_file_id"`
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
	// The identifier of the File containing the front of the document.
	FileID param.Field[string] `json:"file_id,required"`
	// The identifier of the File containing the back of the document. Not every
	// document has a reverse side.
	BackFileID param.Field[string] `json:"back_file_id"`
	// The document's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date" format:"date"`
}

func (r EntityNewParamsTrustGrantorIdentificationOther) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information about the passport used for identification. Required if `method` is
// equal to `passport`.
type EntityNewParamsTrustGrantorIdentificationPassport struct {
	// The country that issued the passport.
	Country param.Field[string] `json:"country,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The identifier of the File containing the passport.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r EntityNewParamsTrustGrantorIdentificationPassport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityListParams struct {
	CreatedAt param.Field[EntityListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  param.Field[int64]                  `query:"limit"`
	Status param.Field[EntityListParamsStatus] `query:"status"`
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

type EntityListParamsStatus struct {
	// Filter Entities for those with the specified status or statuses. For GET
	// requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In param.Field[[]EntityListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [EntityListParamsStatus]'s query parameters as `url.Values`.
func (r EntityListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EntityListParamsStatusIn string

const (
	// The entity is active.
	EntityListParamsStatusInActive EntityListParamsStatusIn = "active"
	// The entity is archived, and can no longer be used to create accounts.
	EntityListParamsStatusInArchived EntityListParamsStatusIn = "archived"
	// The entity is temporarily disabled and cannot be used for financial activity.
	EntityListParamsStatusInDisabled EntityListParamsStatusIn = "disabled"
)

type EntityUpdateAddressParams struct {
	// The entity's physical address. Post Office Boxes are disallowed.
	Address param.Field[EntityUpdateAddressParamsAddress] `json:"address,required"`
}

func (r EntityUpdateAddressParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The entity's physical address. Post Office Boxes are disallowed.
type EntityUpdateAddressParamsAddress struct {
	// The city of the address.
	City param.Field[string] `json:"city,required"`
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State param.Field[string] `json:"state,required"`
	// The ZIP code of the address.
	Zip param.Field[string] `json:"zip,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
}

func (r EntityUpdateAddressParamsAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
