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

// BeneficialOwnerService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBeneficialOwnerService] method instead.
type BeneficialOwnerService struct {
	Options []option.RequestOption
}

// NewBeneficialOwnerService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBeneficialOwnerService(opts ...option.RequestOption) (r *BeneficialOwnerService) {
	r = &BeneficialOwnerService{}
	r.Options = opts
	return
}

// Retrieve a Beneficial Owner
func (r *BeneficialOwnerService) Get(ctx context.Context, entityBeneficialOwnerID string, opts ...option.RequestOption) (res *EntityBeneficialOwner, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityBeneficialOwnerID == "" {
		err = errors.New("missing required entity_beneficial_owner_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("entity_beneficial_owners/%s", entityBeneficialOwnerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a Beneficial Owner
func (r *BeneficialOwnerService) Update(ctx context.Context, entityBeneficialOwnerID string, body BeneficialOwnerUpdateParams, opts ...option.RequestOption) (res *EntityBeneficialOwner, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityBeneficialOwnerID == "" {
		err = errors.New("missing required entity_beneficial_owner_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("entity_beneficial_owners/%s", entityBeneficialOwnerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List Beneficial Owners
func (r *BeneficialOwnerService) List(ctx context.Context, query BeneficialOwnerListParams, opts ...option.RequestOption) (res *pagination.Page[EntityBeneficialOwner], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "entity_beneficial_owners"
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

// List Beneficial Owners
func (r *BeneficialOwnerService) ListAutoPaging(ctx context.Context, query BeneficialOwnerListParams, opts ...option.RequestOption) *pagination.PageAutoPager[EntityBeneficialOwner] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type EntityBeneficialOwner struct {
	// The identifier of this beneficial owner.
	ID string `json:"id" api:"required"`
	// This person's role or title within the entity.
	CompanyTitle string `json:"company_title" api:"required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the
	// Beneficial Owner was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Personal details for the beneficial owner.
	Individual EntityBeneficialOwnerIndividual `json:"individual" api:"required"`
	// Why this person is considered a beneficial owner of the entity.
	Prongs []EntityBeneficialOwnerProng `json:"prongs" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `entity_beneficial_owner`.
	Type        EntityBeneficialOwnerType `json:"type" api:"required"`
	ExtraFields map[string]interface{}    `json:"-" api:"extrafields"`
	JSON        entityBeneficialOwnerJSON `json:"-"`
}

// entityBeneficialOwnerJSON contains the JSON metadata for the struct
// [EntityBeneficialOwner]
type entityBeneficialOwnerJSON struct {
	ID           apijson.Field
	CompanyTitle apijson.Field
	CreatedAt    apijson.Field
	Individual   apijson.Field
	Prongs       apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EntityBeneficialOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityBeneficialOwnerJSON) RawJSON() string {
	return r.raw
}

// Personal details for the beneficial owner.
type EntityBeneficialOwnerIndividual struct {
	// The person's address.
	Address EntityBeneficialOwnerIndividualAddress `json:"address" api:"required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth time.Time `json:"date_of_birth" api:"required" format:"date"`
	// A means of verifying the person's identity.
	Identification EntityBeneficialOwnerIndividualIdentification `json:"identification" api:"required"`
	// The person's legal name.
	Name string                              `json:"name" api:"required"`
	JSON entityBeneficialOwnerIndividualJSON `json:"-"`
}

// entityBeneficialOwnerIndividualJSON contains the JSON metadata for the struct
// [EntityBeneficialOwnerIndividual]
type entityBeneficialOwnerIndividualJSON struct {
	Address        apijson.Field
	DateOfBirth    apijson.Field
	Identification apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityBeneficialOwnerIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityBeneficialOwnerIndividualJSON) RawJSON() string {
	return r.raw
}

// The person's address.
type EntityBeneficialOwnerIndividualAddress struct {
	// The city, district, town, or village of the address.
	City string `json:"city" api:"required,nullable"`
	// The two-letter ISO 3166-1 alpha-2 code for the country of the address.
	Country string `json:"country" api:"required"`
	// The first line of the address.
	Line1 string `json:"line1" api:"required"`
	// The second line of the address.
	Line2 string `json:"line2" api:"required,nullable"`
	// The two-letter United States Postal Service (USPS) abbreviation for the US
	// state, province, or region of the address.
	State string `json:"state" api:"required,nullable"`
	// The ZIP or postal code of the address.
	Zip  string                                     `json:"zip" api:"required,nullable"`
	JSON entityBeneficialOwnerIndividualAddressJSON `json:"-"`
}

// entityBeneficialOwnerIndividualAddressJSON contains the JSON metadata for the
// struct [EntityBeneficialOwnerIndividualAddress]
type entityBeneficialOwnerIndividualAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityBeneficialOwnerIndividualAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityBeneficialOwnerIndividualAddressJSON) RawJSON() string {
	return r.raw
}

// A means of verifying the person's identity.
type EntityBeneficialOwnerIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method EntityBeneficialOwnerIndividualIdentificationMethod `json:"method" api:"required"`
	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 string                                            `json:"number_last4" api:"required"`
	ExtraFields map[string]interface{}                            `json:"-" api:"extrafields"`
	JSON        entityBeneficialOwnerIndividualIdentificationJSON `json:"-"`
}

// entityBeneficialOwnerIndividualIdentificationJSON contains the JSON metadata for
// the struct [EntityBeneficialOwnerIndividualIdentification]
type entityBeneficialOwnerIndividualIdentificationJSON struct {
	Method      apijson.Field
	NumberLast4 apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityBeneficialOwnerIndividualIdentification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityBeneficialOwnerIndividualIdentificationJSON) RawJSON() string {
	return r.raw
}

// A method that can be used to verify the individual's identity.
type EntityBeneficialOwnerIndividualIdentificationMethod string

const (
	EntityBeneficialOwnerIndividualIdentificationMethodSocialSecurityNumber                   EntityBeneficialOwnerIndividualIdentificationMethod = "social_security_number"
	EntityBeneficialOwnerIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber EntityBeneficialOwnerIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	EntityBeneficialOwnerIndividualIdentificationMethodPassport                               EntityBeneficialOwnerIndividualIdentificationMethod = "passport"
	EntityBeneficialOwnerIndividualIdentificationMethodDriversLicense                         EntityBeneficialOwnerIndividualIdentificationMethod = "drivers_license"
	EntityBeneficialOwnerIndividualIdentificationMethodOther                                  EntityBeneficialOwnerIndividualIdentificationMethod = "other"
)

func (r EntityBeneficialOwnerIndividualIdentificationMethod) IsKnown() bool {
	switch r {
	case EntityBeneficialOwnerIndividualIdentificationMethodSocialSecurityNumber, EntityBeneficialOwnerIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber, EntityBeneficialOwnerIndividualIdentificationMethodPassport, EntityBeneficialOwnerIndividualIdentificationMethodDriversLicense, EntityBeneficialOwnerIndividualIdentificationMethodOther:
		return true
	}
	return false
}

type EntityBeneficialOwnerProng string

const (
	EntityBeneficialOwnerProngOwnership EntityBeneficialOwnerProng = "ownership"
	EntityBeneficialOwnerProngControl   EntityBeneficialOwnerProng = "control"
)

func (r EntityBeneficialOwnerProng) IsKnown() bool {
	switch r {
	case EntityBeneficialOwnerProngOwnership, EntityBeneficialOwnerProngControl:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `entity_beneficial_owner`.
type EntityBeneficialOwnerType string

const (
	EntityBeneficialOwnerTypeEntityBeneficialOwner EntityBeneficialOwnerType = "entity_beneficial_owner"
)

func (r EntityBeneficialOwnerType) IsKnown() bool {
	switch r {
	case EntityBeneficialOwnerTypeEntityBeneficialOwner:
		return true
	}
	return false
}

type BeneficialOwnerUpdateParams struct {
	// The individual's physical address. Mail receiving locations like PO Boxes and
	// PMB's are disallowed.
	Address param.Field[BeneficialOwnerUpdateParamsAddress] `json:"address"`
}

func (r BeneficialOwnerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The individual's physical address. Mail receiving locations like PO Boxes and
// PMB's are disallowed.
type BeneficialOwnerUpdateParamsAddress struct {
	// The city, district, town, or village of the address.
	City param.Field[string] `json:"city" api:"required"`
	// The two-letter ISO 3166-1 alpha-2 code for the country of the address.
	Country param.Field[string] `json:"country" api:"required"`
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1" api:"required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
	// The two-letter United States Postal Service (USPS) abbreviation for the US
	// state, province, or region of the address. Required in certain countries.
	State param.Field[string] `json:"state"`
	// The ZIP or postal code of the address. Required in certain countries.
	Zip param.Field[string] `json:"zip"`
}

func (r BeneficialOwnerUpdateParamsAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BeneficialOwnerListParams struct {
	// The identifier of the Entity to list beneficial owners for. Only `corporation`
	// entities have beneficial owners.
	EntityID param.Field[string] `query:"entity_id" api:"required"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [BeneficialOwnerListParams]'s query parameters as
// `url.Values`.
func (r BeneficialOwnerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
