// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
	"github.com/Increase/increase-go/packages/pagination"
)

// EntityService contains methods and other services that help with interacting
// with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityService] method instead.
type EntityService struct {
	Options []option.RequestOption
}

// NewEntityService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEntityService(opts ...option.RequestOption) (r *EntityService) {
	r = &EntityService{}
	r.Options = opts
	return
}

// Create an Entity
func (r *EntityService) New(ctx context.Context, body EntityNewParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "entities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an Entity
func (r *EntityService) Get(ctx context.Context, entityID string, opts ...option.RequestOption) (res *Entity, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an Entity
func (r *EntityService) Update(ctx context.Context, entityID string, body EntityUpdateParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List Entities
func (r *EntityService) List(ctx context.Context, query EntityListParams, opts ...option.RequestOption) (res *pagination.Page[Entity], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
func (r *EntityService) ListAutoPaging(ctx context.Context, query EntityListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Entity] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Archive an Entity
func (r *EntityService) Archive(ctx context.Context, entityID string, opts ...option.RequestOption) (res *Entity, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/archive", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Archive a beneficial owner for a corporate Entity
func (r *EntityService) ArchiveBeneficialOwner(ctx context.Context, entityID string, body EntityArchiveBeneficialOwnerParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/archive_beneficial_owner", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Depending on your program, you may be required to re-confirm an Entity's details
// on a recurring basis. After making any required updates, call this endpoint to
// record that your user confirmed their details.
func (r *EntityService) Confirm(ctx context.Context, entityID string, body EntityConfirmParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/confirm", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a beneficial owner for a corporate Entity
func (r *EntityService) NewBeneficialOwner(ctx context.Context, entityID string, body EntityNewBeneficialOwnerParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/create_beneficial_owner", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a Natural Person or Corporation's address
func (r *EntityService) UpdateAddress(ctx context.Context, entityID string, body EntityUpdateAddressParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/update_address", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update the address for a beneficial owner belonging to a corporate Entity
func (r *EntityService) UpdateBeneficialOwnerAddress(ctx context.Context, entityID string, body EntityUpdateBeneficialOwnerAddressParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/update_beneficial_owner_address", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update the industry code for a corporate Entity
func (r *EntityService) UpdateIndustryCode(ctx context.Context, entityID string, body EntityUpdateIndustryCodeParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/update_industry_code", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Entities are the legal entities that own accounts. They can be people,
// corporations, partnerships, government authorities, or trusts.
type Entity struct {
	// The entity's identifier.
	ID string `json:"id,required"`
	// Details of the corporation entity. Will be present if `structure` is equal to
	// `corporation`.
	Corporation EntityCorporation `json:"corporation,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Entity
	// was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The entity's description for display purposes.
	Description string `json:"description,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the
	// Entity's details were most recently confirmed.
	DetailsConfirmedAt time.Time `json:"details_confirmed_at,required,nullable" format:"date-time"`
	// Details of the government authority entity. Will be present if `structure` is
	// equal to `government_authority`.
	GovernmentAuthority EntityGovernmentAuthority `json:"government_authority,required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// Details of the joint entity. Will be present if `structure` is equal to `joint`.
	Joint EntityJoint `json:"joint,required,nullable"`
	// Details of the natural person entity. Will be present if `structure` is equal to
	// `natural_person`.
	NaturalPerson EntityNaturalPerson `json:"natural_person,required,nullable"`
	// An assessment of the entity’s potential risk of involvement in financial crimes,
	// such as money laundering.
	RiskRating EntityRiskRating `json:"risk_rating,required,nullable"`
	// The status of the entity.
	Status EntityStatus `json:"status,required"`
	// The entity's legal structure.
	Structure EntityStructure `json:"structure,required"`
	// Additional documentation associated with the entity. This is limited to the
	// first 10 documents for an entity. If an entity has more than 10 documents, use
	// the GET /entity_supplemental_documents list endpoint to retrieve them.
	SupplementalDocuments []EntitySupplementalDocument `json:"supplemental_documents,required"`
	// If you are using a third-party service for identity verification, you can use
	// this field to associate this Entity with the identifier that represents them in
	// that service.
	ThirdPartyVerification EntityThirdPartyVerification `json:"third_party_verification,required,nullable"`
	// Details of the trust entity. Will be present if `structure` is equal to `trust`.
	Trust EntityTrust `json:"trust,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `entity`.
	Type EntityType `json:"type,required"`
	JSON entityJSON `json:"-"`
}

// entityJSON contains the JSON metadata for the struct [Entity]
type entityJSON struct {
	ID                     apijson.Field
	Corporation            apijson.Field
	CreatedAt              apijson.Field
	Description            apijson.Field
	DetailsConfirmedAt     apijson.Field
	GovernmentAuthority    apijson.Field
	IdempotencyKey         apijson.Field
	Joint                  apijson.Field
	NaturalPerson          apijson.Field
	RiskRating             apijson.Field
	Status                 apijson.Field
	Structure              apijson.Field
	SupplementalDocuments  apijson.Field
	ThirdPartyVerification apijson.Field
	Trust                  apijson.Field
	Type                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *Entity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityJSON) RawJSON() string {
	return r.raw
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
	// The numeric North American Industry Classification System (NAICS) code submitted
	// for the corporation.
	IndustryCode string `json:"industry_code,required,nullable"`
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
	IndustryCode       apijson.Field
	Name               apijson.Field
	TaxIdentifier      apijson.Field
	Website            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EntityCorporation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityCorporationJSON) RawJSON() string {
	return r.raw
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

func (r entityCorporationAddressJSON) RawJSON() string {
	return r.raw
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

func (r entityCorporationBeneficialOwnerJSON) RawJSON() string {
	return r.raw
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

func (r entityCorporationBeneficialOwnersIndividualJSON) RawJSON() string {
	return r.raw
}

// The person's address.
type EntityCorporationBeneficialOwnersIndividualAddress struct {
	// The city, district, town, or village of the address.
	City string `json:"city,required,nullable"`
	// The two-letter ISO 3166-1 alpha-2 code for the country of the address.
	Country string `json:"country,required"`
	// The first line of the address.
	Line1 string `json:"line1,required"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The two-letter United States Postal Service (USPS) abbreviation for the US
	// state, province, or region of the address.
	State string `json:"state,required,nullable"`
	// The ZIP or postal code of the address.
	Zip  string                                                 `json:"zip,required,nullable"`
	JSON entityCorporationBeneficialOwnersIndividualAddressJSON `json:"-"`
}

// entityCorporationBeneficialOwnersIndividualAddressJSON contains the JSON
// metadata for the struct [EntityCorporationBeneficialOwnersIndividualAddress]
type entityCorporationBeneficialOwnersIndividualAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
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

func (r entityCorporationBeneficialOwnersIndividualAddressJSON) RawJSON() string {
	return r.raw
}

// A means of verifying the person's identity.
type EntityCorporationBeneficialOwnersIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityCorporationBeneficialOwnersIndividualIdentificationMethod `json:"method,required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string                                                        `json:"number_last4,required"`
	ExtraFields map[string]interface{}                                        `json:"-,extras"`
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

func (r entityCorporationBeneficialOwnersIndividualIdentificationJSON) RawJSON() string {
	return r.raw
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

func (r EntityCorporationBeneficialOwnersIndividualIdentificationMethod) IsKnown() bool {
	switch r {
	case EntityCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber, EntityCorporationBeneficialOwnersIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber, EntityCorporationBeneficialOwnersIndividualIdentificationMethodPassport, EntityCorporationBeneficialOwnersIndividualIdentificationMethodDriversLicense, EntityCorporationBeneficialOwnersIndividualIdentificationMethodOther:
		return true
	}
	return false
}

// Why this person is considered a beneficial owner of the entity.
type EntityCorporationBeneficialOwnersProng string

const (
	EntityCorporationBeneficialOwnersProngOwnership EntityCorporationBeneficialOwnersProng = "ownership"
	EntityCorporationBeneficialOwnersProngControl   EntityCorporationBeneficialOwnersProng = "control"
)

func (r EntityCorporationBeneficialOwnersProng) IsKnown() bool {
	switch r {
	case EntityCorporationBeneficialOwnersProngOwnership, EntityCorporationBeneficialOwnersProngControl:
		return true
	}
	return false
}

// Details of the government authority entity. Will be present if `structure` is
// equal to `government_authority`.
type EntityGovernmentAuthority struct {
	// The government authority's address.
	Address EntityGovernmentAuthorityAddress `json:"address,required"`
	// The identifying details of authorized persons of the government authority.
	AuthorizedPersons []EntityGovernmentAuthorityAuthorizedPerson `json:"authorized_persons,required"`
	// The category of the government authority.
	Category EntityGovernmentAuthorityCategory `json:"category,required"`
	// The government authority's name.
	Name string `json:"name,required"`
	// The Employer Identification Number (EIN) of the government authority.
	TaxIdentifier string `json:"tax_identifier,required,nullable"`
	// The government authority's website.
	Website string                        `json:"website,required,nullable"`
	JSON    entityGovernmentAuthorityJSON `json:"-"`
}

// entityGovernmentAuthorityJSON contains the JSON metadata for the struct
// [EntityGovernmentAuthority]
type entityGovernmentAuthorityJSON struct {
	Address           apijson.Field
	AuthorizedPersons apijson.Field
	Category          apijson.Field
	Name              apijson.Field
	TaxIdentifier     apijson.Field
	Website           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EntityGovernmentAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityGovernmentAuthorityJSON) RawJSON() string {
	return r.raw
}

// The government authority's address.
type EntityGovernmentAuthorityAddress struct {
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
	Zip  string                               `json:"zip,required"`
	JSON entityGovernmentAuthorityAddressJSON `json:"-"`
}

// entityGovernmentAuthorityAddressJSON contains the JSON metadata for the struct
// [EntityGovernmentAuthorityAddress]
type entityGovernmentAuthorityAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityGovernmentAuthorityAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityGovernmentAuthorityAddressJSON) RawJSON() string {
	return r.raw
}

type EntityGovernmentAuthorityAuthorizedPerson struct {
	// The identifier of this authorized person.
	AuthorizedPersonID string `json:"authorized_person_id,required"`
	// The person's legal name.
	Name string                                        `json:"name,required"`
	JSON entityGovernmentAuthorityAuthorizedPersonJSON `json:"-"`
}

// entityGovernmentAuthorityAuthorizedPersonJSON contains the JSON metadata for the
// struct [EntityGovernmentAuthorityAuthorizedPerson]
type entityGovernmentAuthorityAuthorizedPersonJSON struct {
	AuthorizedPersonID apijson.Field
	Name               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EntityGovernmentAuthorityAuthorizedPerson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityGovernmentAuthorityAuthorizedPersonJSON) RawJSON() string {
	return r.raw
}

// The category of the government authority.
type EntityGovernmentAuthorityCategory string

const (
	EntityGovernmentAuthorityCategoryMunicipality    EntityGovernmentAuthorityCategory = "municipality"
	EntityGovernmentAuthorityCategoryStateAgency     EntityGovernmentAuthorityCategory = "state_agency"
	EntityGovernmentAuthorityCategoryStateGovernment EntityGovernmentAuthorityCategory = "state_government"
	EntityGovernmentAuthorityCategoryFederalAgency   EntityGovernmentAuthorityCategory = "federal_agency"
)

func (r EntityGovernmentAuthorityCategory) IsKnown() bool {
	switch r {
	case EntityGovernmentAuthorityCategoryMunicipality, EntityGovernmentAuthorityCategoryStateAgency, EntityGovernmentAuthorityCategoryStateGovernment, EntityGovernmentAuthorityCategoryFederalAgency:
		return true
	}
	return false
}

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

func (r entityJointJSON) RawJSON() string {
	return r.raw
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

func (r entityJointIndividualJSON) RawJSON() string {
	return r.raw
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

func (r entityJointIndividualsAddressJSON) RawJSON() string {
	return r.raw
}

// A means of verifying the person's identity.
type EntityJointIndividualsIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityJointIndividualsIdentificationMethod `json:"method,required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string                                   `json:"number_last4,required"`
	ExtraFields map[string]interface{}                   `json:"-,extras"`
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

func (r entityJointIndividualsIdentificationJSON) RawJSON() string {
	return r.raw
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

func (r EntityJointIndividualsIdentificationMethod) IsKnown() bool {
	switch r {
	case EntityJointIndividualsIdentificationMethodSocialSecurityNumber, EntityJointIndividualsIdentificationMethodIndividualTaxpayerIdentificationNumber, EntityJointIndividualsIdentificationMethodPassport, EntityJointIndividualsIdentificationMethodDriversLicense, EntityJointIndividualsIdentificationMethodOther:
		return true
	}
	return false
}

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

func (r entityNaturalPersonJSON) RawJSON() string {
	return r.raw
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

func (r entityNaturalPersonAddressJSON) RawJSON() string {
	return r.raw
}

// A means of verifying the person's identity.
type EntityNaturalPersonIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityNaturalPersonIdentificationMethod `json:"method,required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string                                `json:"number_last4,required"`
	ExtraFields map[string]interface{}                `json:"-,extras"`
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

func (r entityNaturalPersonIdentificationJSON) RawJSON() string {
	return r.raw
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

func (r EntityNaturalPersonIdentificationMethod) IsKnown() bool {
	switch r {
	case EntityNaturalPersonIdentificationMethodSocialSecurityNumber, EntityNaturalPersonIdentificationMethodIndividualTaxpayerIdentificationNumber, EntityNaturalPersonIdentificationMethodPassport, EntityNaturalPersonIdentificationMethodDriversLicense, EntityNaturalPersonIdentificationMethodOther:
		return true
	}
	return false
}

// An assessment of the entity’s potential risk of involvement in financial crimes,
// such as money laundering.
type EntityRiskRating struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the risk
	// rating was performed.
	RatedAt time.Time `json:"rated_at,required" format:"date-time"`
	// The rating given to this entity.
	Rating EntityRiskRatingRating `json:"rating,required"`
	JSON   entityRiskRatingJSON   `json:"-"`
}

// entityRiskRatingJSON contains the JSON metadata for the struct
// [EntityRiskRating]
type entityRiskRatingJSON struct {
	RatedAt     apijson.Field
	Rating      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityRiskRating) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityRiskRatingJSON) RawJSON() string {
	return r.raw
}

// The rating given to this entity.
type EntityRiskRatingRating string

const (
	EntityRiskRatingRatingLow    EntityRiskRatingRating = "low"
	EntityRiskRatingRatingMedium EntityRiskRatingRating = "medium"
	EntityRiskRatingRatingHigh   EntityRiskRatingRating = "high"
)

func (r EntityRiskRatingRating) IsKnown() bool {
	switch r {
	case EntityRiskRatingRatingLow, EntityRiskRatingRatingMedium, EntityRiskRatingRatingHigh:
		return true
	}
	return false
}

// The status of the entity.
type EntityStatus string

const (
	EntityStatusActive   EntityStatus = "active"
	EntityStatusArchived EntityStatus = "archived"
	EntityStatusDisabled EntityStatus = "disabled"
)

func (r EntityStatus) IsKnown() bool {
	switch r {
	case EntityStatusActive, EntityStatusArchived, EntityStatusDisabled:
		return true
	}
	return false
}

// The entity's legal structure.
type EntityStructure string

const (
	EntityStructureCorporation         EntityStructure = "corporation"
	EntityStructureNaturalPerson       EntityStructure = "natural_person"
	EntityStructureJoint               EntityStructure = "joint"
	EntityStructureTrust               EntityStructure = "trust"
	EntityStructureGovernmentAuthority EntityStructure = "government_authority"
)

func (r EntityStructure) IsKnown() bool {
	switch r {
	case EntityStructureCorporation, EntityStructureNaturalPerson, EntityStructureJoint, EntityStructureTrust, EntityStructureGovernmentAuthority:
		return true
	}
	return false
}

// If you are using a third-party service for identity verification, you can use
// this field to associate this Entity with the identifier that represents them in
// that service.
type EntityThirdPartyVerification struct {
	// The reference identifier for the third party verification.
	Reference string `json:"reference,required"`
	// The vendor that was used to perform the verification.
	Vendor EntityThirdPartyVerificationVendor `json:"vendor,required"`
	JSON   entityThirdPartyVerificationJSON   `json:"-"`
}

// entityThirdPartyVerificationJSON contains the JSON metadata for the struct
// [EntityThirdPartyVerification]
type entityThirdPartyVerificationJSON struct {
	Reference   apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityThirdPartyVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityThirdPartyVerificationJSON) RawJSON() string {
	return r.raw
}

// The vendor that was used to perform the verification.
type EntityThirdPartyVerificationVendor string

const (
	EntityThirdPartyVerificationVendorAlloy   EntityThirdPartyVerificationVendor = "alloy"
	EntityThirdPartyVerificationVendorMiddesk EntityThirdPartyVerificationVendor = "middesk"
	EntityThirdPartyVerificationVendorOscilar EntityThirdPartyVerificationVendor = "oscilar"
)

func (r EntityThirdPartyVerificationVendor) IsKnown() bool {
	switch r {
	case EntityThirdPartyVerificationVendorAlloy, EntityThirdPartyVerificationVendorMiddesk, EntityThirdPartyVerificationVendorOscilar:
		return true
	}
	return false
}

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

func (r entityTrustJSON) RawJSON() string {
	return r.raw
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

func (r entityTrustAddressJSON) RawJSON() string {
	return r.raw
}

// Whether the trust is `revocable` or `irrevocable`.
type EntityTrustCategory string

const (
	EntityTrustCategoryRevocable   EntityTrustCategory = "revocable"
	EntityTrustCategoryIrrevocable EntityTrustCategory = "irrevocable"
)

func (r EntityTrustCategory) IsKnown() bool {
	switch r {
	case EntityTrustCategoryRevocable, EntityTrustCategoryIrrevocable:
		return true
	}
	return false
}

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

func (r entityTrustGrantorJSON) RawJSON() string {
	return r.raw
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

func (r entityTrustGrantorAddressJSON) RawJSON() string {
	return r.raw
}

// A means of verifying the person's identity.
type EntityTrustGrantorIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityTrustGrantorIdentificationMethod `json:"method,required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string                               `json:"number_last4,required"`
	ExtraFields map[string]interface{}               `json:"-,extras"`
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

func (r entityTrustGrantorIdentificationJSON) RawJSON() string {
	return r.raw
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

func (r EntityTrustGrantorIdentificationMethod) IsKnown() bool {
	switch r {
	case EntityTrustGrantorIdentificationMethodSocialSecurityNumber, EntityTrustGrantorIdentificationMethodIndividualTaxpayerIdentificationNumber, EntityTrustGrantorIdentificationMethodPassport, EntityTrustGrantorIdentificationMethodDriversLicense, EntityTrustGrantorIdentificationMethodOther:
		return true
	}
	return false
}

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

func (r entityTrustTrusteeJSON) RawJSON() string {
	return r.raw
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

func (r entityTrustTrusteesIndividualJSON) RawJSON() string {
	return r.raw
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

func (r entityTrustTrusteesIndividualAddressJSON) RawJSON() string {
	return r.raw
}

// A means of verifying the person's identity.
type EntityTrustTrusteesIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityTrustTrusteesIndividualIdentificationMethod `json:"method,required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string                                          `json:"number_last4,required"`
	ExtraFields map[string]interface{}                          `json:"-,extras"`
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

func (r entityTrustTrusteesIndividualIdentificationJSON) RawJSON() string {
	return r.raw
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

func (r EntityTrustTrusteesIndividualIdentificationMethod) IsKnown() bool {
	switch r {
	case EntityTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber, EntityTrustTrusteesIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber, EntityTrustTrusteesIndividualIdentificationMethodPassport, EntityTrustTrusteesIndividualIdentificationMethodDriversLicense, EntityTrustTrusteesIndividualIdentificationMethodOther:
		return true
	}
	return false
}

// The structure of the trustee. Will always be equal to `individual`.
type EntityTrustTrusteesStructure string

const (
	EntityTrustTrusteesStructureIndividual EntityTrustTrusteesStructure = "individual"
)

func (r EntityTrustTrusteesStructure) IsKnown() bool {
	switch r {
	case EntityTrustTrusteesStructureIndividual:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `entity`.
type EntityType string

const (
	EntityTypeEntity EntityType = "entity"
)

func (r EntityType) IsKnown() bool {
	switch r {
	case EntityTypeEntity:
		return true
	}
	return false
}

type EntityNewParams struct {
	// The type of Entity to create.
	Structure param.Field[EntityNewParamsStructure] `json:"structure,required"`
	// Details of the corporation entity to create. Required if `structure` is equal to
	// `corporation`.
	Corporation param.Field[EntityNewParamsCorporation] `json:"corporation"`
	// The description you choose to give the entity.
	Description param.Field[string] `json:"description"`
	// Details of the Government Authority entity to create. Required if `structure` is
	// equal to `government_authority`.
	GovernmentAuthority param.Field[EntityNewParamsGovernmentAuthority] `json:"government_authority"`
	// Details of the joint entity to create. Required if `structure` is equal to
	// `joint`.
	Joint param.Field[EntityNewParamsJoint] `json:"joint"`
	// Details of the natural person entity to create. Required if `structure` is equal
	// to `natural_person`. Natural people entities should be submitted with
	// `social_security_number` or `individual_taxpayer_identification_number`
	// identification methods.
	NaturalPerson param.Field[EntityNewParamsNaturalPerson] `json:"natural_person"`
	// An assessment of the entity’s potential risk of involvement in financial crimes,
	// such as money laundering.
	RiskRating param.Field[EntityNewParamsRiskRating] `json:"risk_rating"`
	// Additional documentation associated with the entity.
	SupplementalDocuments param.Field[[]EntityNewParamsSupplementalDocument] `json:"supplemental_documents"`
	// If you are using a third-party service for identity verification, you can use
	// this field to associate this Entity with the identifier that represents them in
	// that service.
	ThirdPartyVerification param.Field[EntityNewParamsThirdPartyVerification] `json:"third_party_verification"`
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
	EntityNewParamsStructureCorporation         EntityNewParamsStructure = "corporation"
	EntityNewParamsStructureNaturalPerson       EntityNewParamsStructure = "natural_person"
	EntityNewParamsStructureJoint               EntityNewParamsStructure = "joint"
	EntityNewParamsStructureTrust               EntityNewParamsStructure = "trust"
	EntityNewParamsStructureGovernmentAuthority EntityNewParamsStructure = "government_authority"
)

func (r EntityNewParamsStructure) IsKnown() bool {
	switch r {
	case EntityNewParamsStructureCorporation, EntityNewParamsStructureNaturalPerson, EntityNewParamsStructureJoint, EntityNewParamsStructureTrust, EntityNewParamsStructureGovernmentAuthority:
		return true
	}
	return false
}

// Details of the corporation entity to create. Required if `structure` is equal to
// `corporation`.
type EntityNewParamsCorporation struct {
	// The entity's physical address. Mail receiving locations like PO Boxes and PMB's
	// are disallowed.
	Address param.Field[EntityNewParamsCorporationAddress] `json:"address,required"`
	// The identifying details of each person who owns 25% or more of the business and
	// one control person, like the CEO, CFO, or other executive. You can submit
	// between 1 and 5 people to this list.
	BeneficialOwners param.Field[[]EntityNewParamsCorporationBeneficialOwner] `json:"beneficial_owners,required"`
	// The legal name of the corporation.
	Name param.Field[string] `json:"name,required"`
	// The Employer Identification Number (EIN) for the corporation.
	TaxIdentifier param.Field[string] `json:"tax_identifier,required"`
	// If the entity is exempt from the requirement to submit beneficial owners,
	// provide the justification. If a reason is provided, you do not need to submit a
	// list of beneficial owners.
	BeneficialOwnershipExemptionReason param.Field[EntityNewParamsCorporationBeneficialOwnershipExemptionReason] `json:"beneficial_ownership_exemption_reason"`
	// The two-letter United States Postal Service (USPS) abbreviation for the
	// corporation's state of incorporation.
	IncorporationState param.Field[string] `json:"incorporation_state"`
	// The North American Industry Classification System (NAICS) code for the
	// corporation's primary line of business. This is a number, like `5132` for
	// `Software Publishers`. A full list of classification codes is available
	// [here](https://increase.com/documentation/data-dictionary#north-american-industry-classification-system-codes).
	IndustryCode param.Field[string] `json:"industry_code"`
	// The website of the corporation.
	Website param.Field[string] `json:"website"`
}

func (r EntityNewParamsCorporation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The entity's physical address. Mail receiving locations like PO Boxes and PMB's
// are disallowed.
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
	CompanyTitle param.Field[string]    `json:"company_title"`
	ExtraFields  map[string]interface{} `json:"-,extras"`
}

func (r EntityNewParamsCorporationBeneficialOwner) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Personal details for the beneficial owner.
type EntityNewParamsCorporationBeneficialOwnersIndividual struct {
	// The individual's physical address. Mail receiving locations like PO Boxes and
	// PMB's are disallowed.
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

// The individual's physical address. Mail receiving locations like PO Boxes and
// PMB's are disallowed.
type EntityNewParamsCorporationBeneficialOwnersIndividualAddress struct {
	// The city, district, town, or village of the address.
	City param.Field[string] `json:"city,required"`
	// The two-letter ISO 3166-1 alpha-2 code for the country of the address.
	Country param.Field[string] `json:"country,required"`
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
	// The two-letter United States Postal Service (USPS) abbreviation for the US
	// state, province, or region of the address. Required in certain countries.
	State param.Field[string] `json:"state"`
	// The ZIP or postal code of the address. Required in certain countries.
	Zip param.Field[string] `json:"zip"`
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
	Passport    param.Field[EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationPassport] `json:"passport"`
	ExtraFields map[string]interface{}                                                                  `json:"-,extras"`
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

func (r EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethod) IsKnown() bool {
	switch r {
	case EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber, EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber, EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodPassport, EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodDriversLicense, EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodOther:
		return true
	}
	return false
}

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
	// document (e.g., `US`).
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
	// The two-character ISO 3166-1 code representing the country that issued the
	// document (e.g., `US`).
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
	EntityNewParamsCorporationBeneficialOwnersProngOwnership EntityNewParamsCorporationBeneficialOwnersProng = "ownership"
	EntityNewParamsCorporationBeneficialOwnersProngControl   EntityNewParamsCorporationBeneficialOwnersProng = "control"
)

func (r EntityNewParamsCorporationBeneficialOwnersProng) IsKnown() bool {
	switch r {
	case EntityNewParamsCorporationBeneficialOwnersProngOwnership, EntityNewParamsCorporationBeneficialOwnersProngControl:
		return true
	}
	return false
}

// If the entity is exempt from the requirement to submit beneficial owners,
// provide the justification. If a reason is provided, you do not need to submit a
// list of beneficial owners.
type EntityNewParamsCorporationBeneficialOwnershipExemptionReason string

const (
	EntityNewParamsCorporationBeneficialOwnershipExemptionReasonRegulatedFinancialInstitution EntityNewParamsCorporationBeneficialOwnershipExemptionReason = "regulated_financial_institution"
	EntityNewParamsCorporationBeneficialOwnershipExemptionReasonPubliclyTradedCompany         EntityNewParamsCorporationBeneficialOwnershipExemptionReason = "publicly_traded_company"
	EntityNewParamsCorporationBeneficialOwnershipExemptionReasonPublicEntity                  EntityNewParamsCorporationBeneficialOwnershipExemptionReason = "public_entity"
	EntityNewParamsCorporationBeneficialOwnershipExemptionReasonOther                         EntityNewParamsCorporationBeneficialOwnershipExemptionReason = "other"
)

func (r EntityNewParamsCorporationBeneficialOwnershipExemptionReason) IsKnown() bool {
	switch r {
	case EntityNewParamsCorporationBeneficialOwnershipExemptionReasonRegulatedFinancialInstitution, EntityNewParamsCorporationBeneficialOwnershipExemptionReasonPubliclyTradedCompany, EntityNewParamsCorporationBeneficialOwnershipExemptionReasonPublicEntity, EntityNewParamsCorporationBeneficialOwnershipExemptionReasonOther:
		return true
	}
	return false
}

// Details of the Government Authority entity to create. Required if `structure` is
// equal to `government_authority`.
type EntityNewParamsGovernmentAuthority struct {
	// The entity's physical address. Mail receiving locations like PO Boxes and PMB's
	// are disallowed.
	Address param.Field[EntityNewParamsGovernmentAuthorityAddress] `json:"address,required"`
	// The identifying details of authorized officials acting on the entity's behalf.
	AuthorizedPersons param.Field[[]EntityNewParamsGovernmentAuthorityAuthorizedPerson] `json:"authorized_persons,required"`
	// The category of the government authority.
	Category param.Field[EntityNewParamsGovernmentAuthorityCategory] `json:"category,required"`
	// The legal name of the government authority.
	Name param.Field[string] `json:"name,required"`
	// The Employer Identification Number (EIN) for the government authority.
	TaxIdentifier param.Field[string] `json:"tax_identifier,required"`
	// The website of the government authority.
	Website param.Field[string] `json:"website"`
}

func (r EntityNewParamsGovernmentAuthority) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The entity's physical address. Mail receiving locations like PO Boxes and PMB's
// are disallowed.
type EntityNewParamsGovernmentAuthorityAddress struct {
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

func (r EntityNewParamsGovernmentAuthorityAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityNewParamsGovernmentAuthorityAuthorizedPerson struct {
	// The person's legal name.
	Name param.Field[string] `json:"name,required"`
}

func (r EntityNewParamsGovernmentAuthorityAuthorizedPerson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The category of the government authority.
type EntityNewParamsGovernmentAuthorityCategory string

const (
	EntityNewParamsGovernmentAuthorityCategoryMunicipality    EntityNewParamsGovernmentAuthorityCategory = "municipality"
	EntityNewParamsGovernmentAuthorityCategoryStateAgency     EntityNewParamsGovernmentAuthorityCategory = "state_agency"
	EntityNewParamsGovernmentAuthorityCategoryStateGovernment EntityNewParamsGovernmentAuthorityCategory = "state_government"
	EntityNewParamsGovernmentAuthorityCategoryFederalAgency   EntityNewParamsGovernmentAuthorityCategory = "federal_agency"
)

func (r EntityNewParamsGovernmentAuthorityCategory) IsKnown() bool {
	switch r {
	case EntityNewParamsGovernmentAuthorityCategoryMunicipality, EntityNewParamsGovernmentAuthorityCategoryStateAgency, EntityNewParamsGovernmentAuthorityCategoryStateGovernment, EntityNewParamsGovernmentAuthorityCategoryFederalAgency:
		return true
	}
	return false
}

// Details of the joint entity to create. Required if `structure` is equal to
// `joint`.
type EntityNewParamsJoint struct {
	// The two individuals that share control of the entity.
	Individuals param.Field[[]EntityNewParamsJointIndividual] `json:"individuals,required"`
}

func (r EntityNewParamsJoint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityNewParamsJointIndividual struct {
	// The individual's physical address. Mail receiving locations like PO Boxes and
	// PMB's are disallowed.
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

// The individual's physical address. Mail receiving locations like PO Boxes and
// PMB's are disallowed.
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
	Passport    param.Field[EntityNewParamsJointIndividualsIdentificationPassport] `json:"passport"`
	ExtraFields map[string]interface{}                                             `json:"-,extras"`
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

func (r EntityNewParamsJointIndividualsIdentificationMethod) IsKnown() bool {
	switch r {
	case EntityNewParamsJointIndividualsIdentificationMethodSocialSecurityNumber, EntityNewParamsJointIndividualsIdentificationMethodIndividualTaxpayerIdentificationNumber, EntityNewParamsJointIndividualsIdentificationMethodPassport, EntityNewParamsJointIndividualsIdentificationMethodDriversLicense, EntityNewParamsJointIndividualsIdentificationMethodOther:
		return true
	}
	return false
}

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
	// document (e.g., `US`).
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
	// The two-character ISO 3166-1 code representing the country that issued the
	// passport (e.g., `US`).
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
	// The individual's physical address. Mail receiving locations like PO Boxes and
	// PMB's are disallowed.
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

// The individual's physical address. Mail receiving locations like PO Boxes and
// PMB's are disallowed.
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
	Passport    param.Field[EntityNewParamsNaturalPersonIdentificationPassport] `json:"passport"`
	ExtraFields map[string]interface{}                                          `json:"-,extras"`
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

func (r EntityNewParamsNaturalPersonIdentificationMethod) IsKnown() bool {
	switch r {
	case EntityNewParamsNaturalPersonIdentificationMethodSocialSecurityNumber, EntityNewParamsNaturalPersonIdentificationMethodIndividualTaxpayerIdentificationNumber, EntityNewParamsNaturalPersonIdentificationMethodPassport, EntityNewParamsNaturalPersonIdentificationMethodDriversLicense, EntityNewParamsNaturalPersonIdentificationMethodOther:
		return true
	}
	return false
}

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
	// document (e.g., `US`).
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
	// The two-character ISO 3166-1 code representing the country that issued the
	// passport (e.g., `US`).
	Country param.Field[string] `json:"country,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The identifier of the File containing the passport.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r EntityNewParamsNaturalPersonIdentificationPassport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An assessment of the entity’s potential risk of involvement in financial crimes,
// such as money laundering.
type EntityNewParamsRiskRating struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the risk
	// rating was performed.
	RatedAt param.Field[time.Time] `json:"rated_at,required" format:"date-time"`
	// The rating given to this entity.
	Rating param.Field[EntityNewParamsRiskRatingRating] `json:"rating,required"`
}

func (r EntityNewParamsRiskRating) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rating given to this entity.
type EntityNewParamsRiskRatingRating string

const (
	EntityNewParamsRiskRatingRatingLow    EntityNewParamsRiskRatingRating = "low"
	EntityNewParamsRiskRatingRatingMedium EntityNewParamsRiskRatingRating = "medium"
	EntityNewParamsRiskRatingRatingHigh   EntityNewParamsRiskRatingRating = "high"
)

func (r EntityNewParamsRiskRatingRating) IsKnown() bool {
	switch r {
	case EntityNewParamsRiskRatingRatingLow, EntityNewParamsRiskRatingRatingMedium, EntityNewParamsRiskRatingRatingHigh:
		return true
	}
	return false
}

type EntityNewParamsSupplementalDocument struct {
	// The identifier of the File containing the document.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r EntityNewParamsSupplementalDocument) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If you are using a third-party service for identity verification, you can use
// this field to associate this Entity with the identifier that represents them in
// that service.
type EntityNewParamsThirdPartyVerification struct {
	// The reference identifier for the third party verification.
	Reference param.Field[string] `json:"reference,required"`
	// The vendor that was used to perform the verification.
	Vendor param.Field[EntityNewParamsThirdPartyVerificationVendor] `json:"vendor,required"`
}

func (r EntityNewParamsThirdPartyVerification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The vendor that was used to perform the verification.
type EntityNewParamsThirdPartyVerificationVendor string

const (
	EntityNewParamsThirdPartyVerificationVendorAlloy   EntityNewParamsThirdPartyVerificationVendor = "alloy"
	EntityNewParamsThirdPartyVerificationVendorMiddesk EntityNewParamsThirdPartyVerificationVendor = "middesk"
	EntityNewParamsThirdPartyVerificationVendorOscilar EntityNewParamsThirdPartyVerificationVendor = "oscilar"
)

func (r EntityNewParamsThirdPartyVerificationVendor) IsKnown() bool {
	switch r {
	case EntityNewParamsThirdPartyVerificationVendorAlloy, EntityNewParamsThirdPartyVerificationVendorMiddesk, EntityNewParamsThirdPartyVerificationVendorOscilar:
		return true
	}
	return false
}

// Details of the trust entity to create. Required if `structure` is equal to
// `trust`.
type EntityNewParamsTrust struct {
	// The trust's physical address. Mail receiving locations like PO Boxes and PMB's
	// are disallowed.
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

// The trust's physical address. Mail receiving locations like PO Boxes and PMB's
// are disallowed.
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
	EntityNewParamsTrustCategoryRevocable   EntityNewParamsTrustCategory = "revocable"
	EntityNewParamsTrustCategoryIrrevocable EntityNewParamsTrustCategory = "irrevocable"
)

func (r EntityNewParamsTrustCategory) IsKnown() bool {
	switch r {
	case EntityNewParamsTrustCategoryRevocable, EntityNewParamsTrustCategoryIrrevocable:
		return true
	}
	return false
}

type EntityNewParamsTrustTrustee struct {
	// The structure of the trustee.
	Structure param.Field[EntityNewParamsTrustTrusteesStructure] `json:"structure,required"`
	// Details of the individual trustee. Within the trustee object, this is required
	// if `structure` is equal to `individual`.
	Individual param.Field[EntityNewParamsTrustTrusteesIndividual] `json:"individual"`
}

func (r EntityNewParamsTrustTrustee) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The structure of the trustee.
type EntityNewParamsTrustTrusteesStructure string

const (
	EntityNewParamsTrustTrusteesStructureIndividual EntityNewParamsTrustTrusteesStructure = "individual"
)

func (r EntityNewParamsTrustTrusteesStructure) IsKnown() bool {
	switch r {
	case EntityNewParamsTrustTrusteesStructureIndividual:
		return true
	}
	return false
}

// Details of the individual trustee. Within the trustee object, this is required
// if `structure` is equal to `individual`.
type EntityNewParamsTrustTrusteesIndividual struct {
	// The individual's physical address. Mail receiving locations like PO Boxes and
	// PMB's are disallowed.
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

// The individual's physical address. Mail receiving locations like PO Boxes and
// PMB's are disallowed.
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
	Passport    param.Field[EntityNewParamsTrustTrusteesIndividualIdentificationPassport] `json:"passport"`
	ExtraFields map[string]interface{}                                                    `json:"-,extras"`
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

func (r EntityNewParamsTrustTrusteesIndividualIdentificationMethod) IsKnown() bool {
	switch r {
	case EntityNewParamsTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber, EntityNewParamsTrustTrusteesIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber, EntityNewParamsTrustTrusteesIndividualIdentificationMethodPassport, EntityNewParamsTrustTrusteesIndividualIdentificationMethodDriversLicense, EntityNewParamsTrustTrusteesIndividualIdentificationMethodOther:
		return true
	}
	return false
}

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
	// document (e.g., `US`).
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
	// The two-character ISO 3166-1 code representing the country that issued the
	// passport (e.g., `US`).
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
	// The individual's physical address. Mail receiving locations like PO Boxes and
	// PMB's are disallowed.
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

// The individual's physical address. Mail receiving locations like PO Boxes and
// PMB's are disallowed.
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
	Passport    param.Field[EntityNewParamsTrustGrantorIdentificationPassport] `json:"passport"`
	ExtraFields map[string]interface{}                                         `json:"-,extras"`
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

func (r EntityNewParamsTrustGrantorIdentificationMethod) IsKnown() bool {
	switch r {
	case EntityNewParamsTrustGrantorIdentificationMethodSocialSecurityNumber, EntityNewParamsTrustGrantorIdentificationMethodIndividualTaxpayerIdentificationNumber, EntityNewParamsTrustGrantorIdentificationMethodPassport, EntityNewParamsTrustGrantorIdentificationMethodDriversLicense, EntityNewParamsTrustGrantorIdentificationMethodOther:
		return true
	}
	return false
}

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
	// document (e.g., `US`).
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
	// The two-character ISO 3166-1 code representing the country that issued the
	// passport (e.g., `US`).
	Country param.Field[string] `json:"country,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The identifier of the File containing the passport.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r EntityNewParamsTrustGrantorIdentificationPassport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityUpdateParams struct {
	// Details of the corporation entity to update. If you specify this parameter and
	// the entity is not a corporation, the request will fail.
	Corporation param.Field[EntityUpdateParamsCorporation] `json:"corporation"`
	// When your user last confirmed the Entity's details. Depending on your program,
	// you may be required to affirmatively confirm details with your users on an
	// annual basis.
	DetailsConfirmedAt param.Field[time.Time] `json:"details_confirmed_at" format:"date-time"`
	// Details of the government authority entity to update. If you specify this
	// parameter and the entity is not a government authority, the request will fail.
	GovernmentAuthority param.Field[EntityUpdateParamsGovernmentAuthority] `json:"government_authority"`
	// Details of the natural person entity to update. If you specify this parameter
	// and the entity is not a natural person, the request will fail.
	NaturalPerson param.Field[EntityUpdateParamsNaturalPerson] `json:"natural_person"`
	// An assessment of the entity’s potential risk of involvement in financial crimes,
	// such as money laundering.
	RiskRating param.Field[EntityUpdateParamsRiskRating] `json:"risk_rating"`
	// If you are using a third-party service for identity verification, you can use
	// this field to associate this Entity with the identifier that represents them in
	// that service.
	ThirdPartyVerification param.Field[EntityUpdateParamsThirdPartyVerification] `json:"third_party_verification"`
	// Details of the trust entity to update. If you specify this parameter and the
	// entity is not a trust, the request will fail.
	Trust param.Field[EntityUpdateParamsTrust] `json:"trust"`
}

func (r EntityUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details of the corporation entity to update. If you specify this parameter and
// the entity is not a corporation, the request will fail.
type EntityUpdateParamsCorporation struct {
	// The entity's physical address. Mail receiving locations like PO Boxes and PMB's
	// are disallowed.
	Address param.Field[EntityUpdateParamsCorporationAddress] `json:"address"`
	// The North American Industry Classification System (NAICS) code for the
	// corporation's primary line of business. This is a number, like `5132` for
	// `Software Publishers`. A full list of classification codes is available
	// [here](https://increase.com/documentation/data-dictionary#north-american-industry-classification-system-codes).
	IndustryCode param.Field[string] `json:"industry_code"`
	// The legal name of the corporation.
	Name param.Field[string] `json:"name"`
}

func (r EntityUpdateParamsCorporation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The entity's physical address. Mail receiving locations like PO Boxes and PMB's
// are disallowed.
type EntityUpdateParamsCorporationAddress struct {
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

func (r EntityUpdateParamsCorporationAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details of the government authority entity to update. If you specify this
// parameter and the entity is not a government authority, the request will fail.
type EntityUpdateParamsGovernmentAuthority struct {
	// The entity's physical address. Mail receiving locations like PO Boxes and PMB's
	// are disallowed.
	Address param.Field[EntityUpdateParamsGovernmentAuthorityAddress] `json:"address"`
	// The legal name of the government authority.
	Name param.Field[string] `json:"name"`
}

func (r EntityUpdateParamsGovernmentAuthority) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The entity's physical address. Mail receiving locations like PO Boxes and PMB's
// are disallowed.
type EntityUpdateParamsGovernmentAuthorityAddress struct {
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

func (r EntityUpdateParamsGovernmentAuthorityAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details of the natural person entity to update. If you specify this parameter
// and the entity is not a natural person, the request will fail.
type EntityUpdateParamsNaturalPerson struct {
	// The entity's physical address. Mail receiving locations like PO Boxes and PMB's
	// are disallowed.
	Address param.Field[EntityUpdateParamsNaturalPersonAddress] `json:"address"`
	// The legal name of the natural person.
	Name param.Field[string] `json:"name"`
}

func (r EntityUpdateParamsNaturalPerson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The entity's physical address. Mail receiving locations like PO Boxes and PMB's
// are disallowed.
type EntityUpdateParamsNaturalPersonAddress struct {
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

func (r EntityUpdateParamsNaturalPersonAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An assessment of the entity’s potential risk of involvement in financial crimes,
// such as money laundering.
type EntityUpdateParamsRiskRating struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the risk
	// rating was performed.
	RatedAt param.Field[time.Time] `json:"rated_at,required" format:"date-time"`
	// The rating given to this entity.
	Rating param.Field[EntityUpdateParamsRiskRatingRating] `json:"rating,required"`
}

func (r EntityUpdateParamsRiskRating) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rating given to this entity.
type EntityUpdateParamsRiskRatingRating string

const (
	EntityUpdateParamsRiskRatingRatingLow    EntityUpdateParamsRiskRatingRating = "low"
	EntityUpdateParamsRiskRatingRatingMedium EntityUpdateParamsRiskRatingRating = "medium"
	EntityUpdateParamsRiskRatingRatingHigh   EntityUpdateParamsRiskRatingRating = "high"
)

func (r EntityUpdateParamsRiskRatingRating) IsKnown() bool {
	switch r {
	case EntityUpdateParamsRiskRatingRatingLow, EntityUpdateParamsRiskRatingRatingMedium, EntityUpdateParamsRiskRatingRatingHigh:
		return true
	}
	return false
}

// If you are using a third-party service for identity verification, you can use
// this field to associate this Entity with the identifier that represents them in
// that service.
type EntityUpdateParamsThirdPartyVerification struct {
	// The reference identifier for the third party verification.
	Reference param.Field[string] `json:"reference,required"`
	// The vendor that was used to perform the verification.
	Vendor param.Field[EntityUpdateParamsThirdPartyVerificationVendor] `json:"vendor,required"`
}

func (r EntityUpdateParamsThirdPartyVerification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The vendor that was used to perform the verification.
type EntityUpdateParamsThirdPartyVerificationVendor string

const (
	EntityUpdateParamsThirdPartyVerificationVendorAlloy   EntityUpdateParamsThirdPartyVerificationVendor = "alloy"
	EntityUpdateParamsThirdPartyVerificationVendorMiddesk EntityUpdateParamsThirdPartyVerificationVendor = "middesk"
	EntityUpdateParamsThirdPartyVerificationVendorOscilar EntityUpdateParamsThirdPartyVerificationVendor = "oscilar"
)

func (r EntityUpdateParamsThirdPartyVerificationVendor) IsKnown() bool {
	switch r {
	case EntityUpdateParamsThirdPartyVerificationVendorAlloy, EntityUpdateParamsThirdPartyVerificationVendorMiddesk, EntityUpdateParamsThirdPartyVerificationVendorOscilar:
		return true
	}
	return false
}

// Details of the trust entity to update. If you specify this parameter and the
// entity is not a trust, the request will fail.
type EntityUpdateParamsTrust struct {
	// The entity's physical address. Mail receiving locations like PO Boxes and PMB's
	// are disallowed.
	Address param.Field[EntityUpdateParamsTrustAddress] `json:"address"`
	// The legal name of the trust.
	Name param.Field[string] `json:"name"`
}

func (r EntityUpdateParamsTrust) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The entity's physical address. Mail receiving locations like PO Boxes and PMB's
// are disallowed.
type EntityUpdateParamsTrustAddress struct {
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

func (r EntityUpdateParamsTrustAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityListParams struct {
	CreatedAt param.Field[EntityListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
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
	EntityListParamsStatusInActive   EntityListParamsStatusIn = "active"
	EntityListParamsStatusInArchived EntityListParamsStatusIn = "archived"
	EntityListParamsStatusInDisabled EntityListParamsStatusIn = "disabled"
)

func (r EntityListParamsStatusIn) IsKnown() bool {
	switch r {
	case EntityListParamsStatusInActive, EntityListParamsStatusInArchived, EntityListParamsStatusInDisabled:
		return true
	}
	return false
}

type EntityArchiveBeneficialOwnerParams struct {
	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwnerID param.Field[string] `json:"beneficial_owner_id,required"`
}

func (r EntityArchiveBeneficialOwnerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityConfirmParams struct {
	// When your user confirmed the Entity's details. If not provided, the current time
	// will be used.
	ConfirmedAt param.Field[time.Time] `json:"confirmed_at" format:"date-time"`
}

func (r EntityConfirmParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityNewBeneficialOwnerParams struct {
	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwner param.Field[EntityNewBeneficialOwnerParamsBeneficialOwner] `json:"beneficial_owner,required"`
}

func (r EntityNewBeneficialOwnerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The identifying details of anyone controlling or owning 25% or more of the
// corporation.
type EntityNewBeneficialOwnerParamsBeneficialOwner struct {
	// Personal details for the beneficial owner.
	Individual param.Field[EntityNewBeneficialOwnerParamsBeneficialOwnerIndividual] `json:"individual,required"`
	// Why this person is considered a beneficial owner of the entity. At least one
	// option is required, if a person is both a control person and owner, submit an
	// array containing both.
	Prongs param.Field[[]EntityNewBeneficialOwnerParamsBeneficialOwnerProng] `json:"prongs,required"`
	// This person's role or title within the entity.
	CompanyTitle param.Field[string]    `json:"company_title"`
	ExtraFields  map[string]interface{} `json:"-,extras"`
}

func (r EntityNewBeneficialOwnerParamsBeneficialOwner) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Personal details for the beneficial owner.
type EntityNewBeneficialOwnerParamsBeneficialOwnerIndividual struct {
	// The individual's physical address. Mail receiving locations like PO Boxes and
	// PMB's are disallowed.
	Address param.Field[EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualAddress] `json:"address,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth param.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// A means of verifying the person's identity.
	Identification param.Field[EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentification] `json:"identification,required"`
	// The person's legal name.
	Name param.Field[string] `json:"name,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID param.Field[bool] `json:"confirmed_no_us_tax_id"`
}

func (r EntityNewBeneficialOwnerParamsBeneficialOwnerIndividual) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The individual's physical address. Mail receiving locations like PO Boxes and
// PMB's are disallowed.
type EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualAddress struct {
	// The city, district, town, or village of the address.
	City param.Field[string] `json:"city,required"`
	// The two-letter ISO 3166-1 alpha-2 code for the country of the address.
	Country param.Field[string] `json:"country,required"`
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
	// The two-letter United States Postal Service (USPS) abbreviation for the US
	// state, province, or region of the address. Required in certain countries.
	State param.Field[string] `json:"state"`
	// The ZIP or postal code of the address. Required in certain countries.
	Zip param.Field[string] `json:"zip"`
}

func (r EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A means of verifying the person's identity.
type EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method param.Field[EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationMethod] `json:"method,required"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number param.Field[string] `json:"number,required"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense param.Field[EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other param.Field[EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationOther] `json:"other"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport    param.Field[EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationPassport] `json:"passport"`
	ExtraFields map[string]interface{}                                                                     `json:"-,extras"`
}

func (r EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A method that can be used to verify the individual's identity.
type EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationMethod string

const (
	EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationMethodSocialSecurityNumber                   EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationMethod = "social_security_number"
	EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationMethodPassport                               EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationMethod = "passport"
	EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationMethodDriversLicense                         EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationMethod = "drivers_license"
	EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationMethodOther                                  EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationMethod = "other"
)

func (r EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationMethod) IsKnown() bool {
	switch r {
	case EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationMethodSocialSecurityNumber, EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber, EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationMethodPassport, EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationMethodDriversLicense, EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationMethodOther:
		return true
	}
	return false
}

// Information about the United States driver's license used for identification.
// Required if `method` is equal to `drivers_license`.
type EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationDriversLicense struct {
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The identifier of the File containing the front of the driver's license.
	FileID param.Field[string] `json:"file_id,required"`
	// The state that issued the provided driver's license.
	State param.Field[string] `json:"state,required"`
	// The identifier of the File containing the back of the driver's license.
	BackFileID param.Field[string] `json:"back_file_id"`
}

func (r EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationDriversLicense) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information about the identification document provided. Required if `method` is
// equal to `other`.
type EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationOther struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document (e.g., `US`).
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

func (r EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationOther) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information about the passport used for identification. Required if `method` is
// equal to `passport`.
type EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationPassport struct {
	// The two-character ISO 3166-1 code representing the country that issued the
	// document (e.g., `US`).
	Country param.Field[string] `json:"country,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The identifier of the File containing the passport.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationPassport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityNewBeneficialOwnerParamsBeneficialOwnerProng string

const (
	EntityNewBeneficialOwnerParamsBeneficialOwnerProngOwnership EntityNewBeneficialOwnerParamsBeneficialOwnerProng = "ownership"
	EntityNewBeneficialOwnerParamsBeneficialOwnerProngControl   EntityNewBeneficialOwnerParamsBeneficialOwnerProng = "control"
)

func (r EntityNewBeneficialOwnerParamsBeneficialOwnerProng) IsKnown() bool {
	switch r {
	case EntityNewBeneficialOwnerParamsBeneficialOwnerProngOwnership, EntityNewBeneficialOwnerParamsBeneficialOwnerProngControl:
		return true
	}
	return false
}

type EntityUpdateAddressParams struct {
	// The entity's physical address. Mail receiving locations like PO Boxes and PMB's
	// are disallowed.
	Address param.Field[EntityUpdateAddressParamsAddress] `json:"address,required"`
}

func (r EntityUpdateAddressParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The entity's physical address. Mail receiving locations like PO Boxes and PMB's
// are disallowed.
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

type EntityUpdateBeneficialOwnerAddressParams struct {
	// The individual's physical address. Mail receiving locations like PO Boxes and
	// PMB's are disallowed.
	Address param.Field[EntityUpdateBeneficialOwnerAddressParamsAddress] `json:"address,required"`
	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwnerID param.Field[string] `json:"beneficial_owner_id,required"`
}

func (r EntityUpdateBeneficialOwnerAddressParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The individual's physical address. Mail receiving locations like PO Boxes and
// PMB's are disallowed.
type EntityUpdateBeneficialOwnerAddressParamsAddress struct {
	// The city, district, town, or village of the address.
	City param.Field[string] `json:"city,required"`
	// The two-letter ISO 3166-1 alpha-2 code for the country of the address.
	Country param.Field[string] `json:"country,required"`
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
	// The two-letter United States Postal Service (USPS) abbreviation for the US
	// state, province, or region of the address. Required in certain countries.
	State param.Field[string] `json:"state"`
	// The ZIP or postal code of the address. Required in certain countries.
	Zip param.Field[string] `json:"zip"`
}

func (r EntityUpdateBeneficialOwnerAddressParamsAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityUpdateIndustryCodeParams struct {
	// The North American Industry Classification System (NAICS) code for the
	// corporation's primary line of business. This is a number, like `5132` for
	// `Software Publishers`. A full list of classification codes is available
	// [here](https://increase.com/documentation/data-dictionary#north-american-industry-classification-system-codes).
	IndustryCode param.Field[string] `json:"industry_code,required"`
}

func (r EntityUpdateIndustryCodeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
