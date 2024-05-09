// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// EntityBeneficialOwnerService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityBeneficialOwnerService] method instead.
type EntityBeneficialOwnerService struct {
	Options []option.RequestOption
}

// NewEntityBeneficialOwnerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEntityBeneficialOwnerService(opts ...option.RequestOption) (r *EntityBeneficialOwnerService) {
	r = &EntityBeneficialOwnerService{}
	r.Options = opts
	return
}

// Create a beneficial owner for a corporate Entity
func (r *EntityBeneficialOwnerService) New(ctx context.Context, body EntityBeneficialOwnerNewParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := "entity_beneficial_owners"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Archive a beneficial owner for a corporate Entity
func (r *EntityBeneficialOwnerService) Archive(ctx context.Context, body EntityBeneficialOwnerArchiveParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := "entity_beneficial_owners/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update the address for a beneficial owner belonging to a corporate Entity
func (r *EntityBeneficialOwnerService) UpdateAddress(ctx context.Context, body EntityBeneficialOwnerUpdateAddressParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := "entity_beneficial_owners/address"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EntityBeneficialOwnerNewParams struct {
	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwner param.Field[EntityBeneficialOwnerNewParamsBeneficialOwner] `json:"beneficial_owner,required"`
	// The identifier of the Entity to associate with the new Beneficial Owner.
	EntityID param.Field[string] `json:"entity_id,required"`
}

func (r EntityBeneficialOwnerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The identifying details of anyone controlling or owning 25% or more of the
// corporation.
type EntityBeneficialOwnerNewParamsBeneficialOwner struct {
	// Personal details for the beneficial owner.
	Individual param.Field[EntityBeneficialOwnerNewParamsBeneficialOwnerIndividual] `json:"individual,required"`
	// Why this person is considered a beneficial owner of the entity. At least one
	// option is required, if a person is both a control person and owner, submit an
	// array containing both.
	Prongs param.Field[[]EntityBeneficialOwnerNewParamsBeneficialOwnerProng] `json:"prongs,required"`
	// This person's role or title within the entity.
	CompanyTitle param.Field[string] `json:"company_title"`
}

func (r EntityBeneficialOwnerNewParamsBeneficialOwner) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Personal details for the beneficial owner.
type EntityBeneficialOwnerNewParamsBeneficialOwnerIndividual struct {
	// The individual's physical address. Mail receiving locations like PO Boxes and
	// PMB's are disallowed.
	Address param.Field[EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualAddress] `json:"address,required"`
	// The person's date of birth in YYYY-MM-DD format.
	DateOfBirth param.Field[time.Time] `json:"date_of_birth,required" format:"date"`
	// A means of verifying the person's identity.
	Identification param.Field[EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentification] `json:"identification,required"`
	// The person's legal name.
	Name param.Field[string] `json:"name,required"`
	// The identification method for an individual can only be a passport, driver's
	// license, or other document if you've confirmed the individual does not have a US
	// tax id (either a Social Security Number or Individual Taxpayer Identification
	// Number).
	ConfirmedNoUsTaxID param.Field[bool] `json:"confirmed_no_us_tax_id"`
}

func (r EntityBeneficialOwnerNewParamsBeneficialOwnerIndividual) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The individual's physical address. Mail receiving locations like PO Boxes and
// PMB's are disallowed.
type EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualAddress struct {
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

func (r EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A means of verifying the person's identity.
type EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentification struct {
	// A method that can be used to verify the individual's identity.
	Method param.Field[EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethod] `json:"method,required"`
	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number param.Field[string] `json:"number,required"`
	// Information about the United States driver's license used for identification.
	// Required if `method` is equal to `drivers_license`.
	DriversLicense param.Field[EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationDriversLicense] `json:"drivers_license"`
	// Information about the identification document provided. Required if `method` is
	// equal to `other`.
	Other param.Field[EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationOther] `json:"other"`
	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport param.Field[EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationPassport] `json:"passport"`
}

func (r EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A method that can be used to verify the individual's identity.
type EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethod string

const (
	// A social security number.
	EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethodSocialSecurityNumber EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethod = "social_security_number"
	// An individual taxpayer identification number (ITIN).
	EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethod = "individual_taxpayer_identification_number"
	// A passport number.
	EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethodPassport EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethod = "passport"
	// A driver's license number.
	EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethodDriversLicense EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethod = "drivers_license"
	// Another identifying document.
	EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethodOther EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethod = "other"
)

func (r EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethod) IsKnown() bool {
	switch r {
	case EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethodSocialSecurityNumber, EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethodIndividualTaxpayerIdentificationNumber, EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethodPassport, EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethodDriversLicense, EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethodOther:
		return true
	}
	return false
}

// Information about the United States driver's license used for identification.
// Required if `method` is equal to `drivers_license`.
type EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationDriversLicense struct {
	// The driver's license's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The identifier of the File containing the front of the driver's license.
	FileID param.Field[string] `json:"file_id,required"`
	// The state that issued the provided driver's license.
	State param.Field[string] `json:"state,required"`
	// The identifier of the File containing the back of the driver's license.
	BackFileID param.Field[string] `json:"back_file_id"`
}

func (r EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationDriversLicense) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information about the identification document provided. Required if `method` is
// equal to `other`.
type EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationOther struct {
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

func (r EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationOther) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Information about the passport used for identification. Required if `method` is
// equal to `passport`.
type EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationPassport struct {
	// The country that issued the passport.
	Country param.Field[string] `json:"country,required"`
	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate param.Field[time.Time] `json:"expiration_date,required" format:"date"`
	// The identifier of the File containing the passport.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationPassport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityBeneficialOwnerNewParamsBeneficialOwnerProng string

const (
	// A person with 25% or greater direct or indirect ownership of the entity.
	EntityBeneficialOwnerNewParamsBeneficialOwnerProngOwnership EntityBeneficialOwnerNewParamsBeneficialOwnerProng = "ownership"
	// A person who manages, directs, or has significant control of the entity.
	EntityBeneficialOwnerNewParamsBeneficialOwnerProngControl EntityBeneficialOwnerNewParamsBeneficialOwnerProng = "control"
)

func (r EntityBeneficialOwnerNewParamsBeneficialOwnerProng) IsKnown() bool {
	switch r {
	case EntityBeneficialOwnerNewParamsBeneficialOwnerProngOwnership, EntityBeneficialOwnerNewParamsBeneficialOwnerProngControl:
		return true
	}
	return false
}

type EntityBeneficialOwnerArchiveParams struct {
	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwnerID param.Field[string] `json:"beneficial_owner_id,required"`
	// The identifier of the Entity to retrieve.
	EntityID param.Field[string] `json:"entity_id,required"`
}

func (r EntityBeneficialOwnerArchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityBeneficialOwnerUpdateAddressParams struct {
	// The individual's physical address. Mail receiving locations like PO Boxes and
	// PMB's are disallowed.
	Address param.Field[EntityBeneficialOwnerUpdateAddressParamsAddress] `json:"address,required"`
	// The identifying details of anyone controlling or owning 25% or more of the
	// corporation.
	BeneficialOwnerID param.Field[string] `json:"beneficial_owner_id,required"`
	// The identifier of the Entity to retrieve.
	EntityID param.Field[string] `json:"entity_id,required"`
}

func (r EntityBeneficialOwnerUpdateAddressParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The individual's physical address. Mail receiving locations like PO Boxes and
// PMB's are disallowed.
type EntityBeneficialOwnerUpdateAddressParamsAddress struct {
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

func (r EntityBeneficialOwnerUpdateAddressParamsAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
