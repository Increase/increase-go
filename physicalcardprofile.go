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

// PhysicalCardProfileService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPhysicalCardProfileService]
// method instead.
type PhysicalCardProfileService struct {
	Options []option.RequestOption
}

// NewPhysicalCardProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPhysicalCardProfileService(opts ...option.RequestOption) (r *PhysicalCardProfileService) {
	r = &PhysicalCardProfileService{}
	r.Options = opts
	return
}

// Create a Physical Card Profile
func (r *PhysicalCardProfileService) New(ctx context.Context, body PhysicalCardProfileNewParams, opts ...option.RequestOption) (res *PhysicalCardProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := "physical_card_profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Card Profile
func (r *PhysicalCardProfileService) Get(ctx context.Context, physicalCardProfileID string, opts ...option.RequestOption) (res *PhysicalCardProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("physical_card_profiles/%s", physicalCardProfileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Physical Card Profiles
func (r *PhysicalCardProfileService) List(ctx context.Context, query PhysicalCardProfileListParams, opts ...option.RequestOption) (res *shared.Page[PhysicalCardProfile], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "physical_card_profiles"
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

// List Physical Card Profiles
func (r *PhysicalCardProfileService) ListAutoPaging(ctx context.Context, query PhysicalCardProfileListParams, opts ...option.RequestOption) *shared.PageAutoPager[PhysicalCardProfile] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Archive a Physical Card Profile
func (r *PhysicalCardProfileService) Archive(ctx context.Context, physicalCardProfileID string, opts ...option.RequestOption) (res *PhysicalCardProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("physical_card_profiles/%s/archive", physicalCardProfileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Clone a Physical Card Profile
func (r *PhysicalCardProfileService) Clone(ctx context.Context, physicalCardProfileID string, body PhysicalCardProfileCloneParams, opts ...option.RequestOption) (res *PhysicalCardProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("physical_card_profiles/%s/clone", physicalCardProfileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This contains artwork and metadata relating to a Physical Card's appearance. For
// more information, see our guide on
// [physical card artwork](https://increase.com/documentation/card-art-physical-cards).
type PhysicalCardProfile struct {
	// The Card Profile identifier.
	ID string `json:"id,required"`
	// The identifier of the File containing the physical card's back image.
	BackImageFileID string `json:"back_image_file_id,required,nullable"`
	// The identifier of the File containing the physical card's carrier image.
	CarrierImageFileID string `json:"carrier_image_file_id,required,nullable"`
	// A phone number the user can contact to receive support for their card.
	ContactPhone string `json:"contact_phone,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The creator of this Physical Card Profile.
	Creator PhysicalCardProfileCreator `json:"creator,required"`
	// A description you can use to identify the Physical Card Profile.
	Description string `json:"description,required"`
	// The identifier of the File containing the physical card's front image.
	FrontImageFileID string `json:"front_image_file_id,required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// Whether this Physical Card Profile is the default for all cards in its Increase
	// group.
	IsDefault bool `json:"is_default,required"`
	// The status of the Physical Card Profile.
	Status PhysicalCardProfileStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `physical_card_profile`.
	Type PhysicalCardProfileType `json:"type,required"`
	JSON physicalCardProfileJSON `json:"-"`
}

// physicalCardProfileJSON contains the JSON metadata for the struct
// [PhysicalCardProfile]
type physicalCardProfileJSON struct {
	ID                 apijson.Field
	BackImageFileID    apijson.Field
	CarrierImageFileID apijson.Field
	ContactPhone       apijson.Field
	CreatedAt          apijson.Field
	Creator            apijson.Field
	Description        apijson.Field
	FrontImageFileID   apijson.Field
	IdempotencyKey     apijson.Field
	IsDefault          apijson.Field
	Status             apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhysicalCardProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The creator of this Physical Card Profile.
type PhysicalCardProfileCreator string

const (
	// This Physical Card Profile was created by Increase.
	PhysicalCardProfileCreatorIncrease PhysicalCardProfileCreator = "increase"
	// This Physical Card Profile was created by you.
	PhysicalCardProfileCreatorUser PhysicalCardProfileCreator = "user"
)

// The status of the Physical Card Profile.
type PhysicalCardProfileStatus string

const (
	// The Card Profile has not yet been processed by Increase.
	PhysicalCardProfileStatusPendingCreating PhysicalCardProfileStatus = "pending_creating"
	// The card profile is awaiting review by Increase.
	PhysicalCardProfileStatusPendingReviewing PhysicalCardProfileStatus = "pending_reviewing"
	// There is an issue with the Physical Card Profile preventing it from use.
	PhysicalCardProfileStatusRejected PhysicalCardProfileStatus = "rejected"
	// The card profile is awaiting submission to the fulfillment provider.
	PhysicalCardProfileStatusPendingSubmitting PhysicalCardProfileStatus = "pending_submitting"
	// The Physical Card Profile has been submitted to the fulfillment provider and is
	// ready to use.
	PhysicalCardProfileStatusActive PhysicalCardProfileStatus = "active"
	// The Physical Card Profile has been archived.
	PhysicalCardProfileStatusArchived PhysicalCardProfileStatus = "archived"
)

// A constant representing the object's type. For this resource it will always be
// `physical_card_profile`.
type PhysicalCardProfileType string

const (
	PhysicalCardProfileTypePhysicalCardProfile PhysicalCardProfileType = "physical_card_profile"
)

type PhysicalCardProfileNewParams struct {
	// The identifier of the File containing the physical card's carrier image.
	CarrierImageFileID param.Field[string] `json:"carrier_image_file_id,required"`
	// A phone number the user can contact to receive support for their card.
	ContactPhone param.Field[string] `json:"contact_phone,required"`
	// A description you can use to identify the Card Profile.
	Description param.Field[string] `json:"description,required"`
	// The identifier of the File containing the physical card's front image.
	FrontImageFileID param.Field[string] `json:"front_image_file_id,required"`
}

func (r PhysicalCardProfileNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PhysicalCardProfileListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  param.Field[int64]                               `query:"limit"`
	Status param.Field[PhysicalCardProfileListParamsStatus] `query:"status"`
}

// URLQuery serializes [PhysicalCardProfileListParams]'s query parameters as
// `url.Values`.
func (r PhysicalCardProfileListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type PhysicalCardProfileListParamsStatus struct {
	// Filter Physical Card Profiles for those with the specified statuses. For GET
	// requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In param.Field[[]PhysicalCardProfileListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [PhysicalCardProfileListParamsStatus]'s query parameters as
// `url.Values`.
func (r PhysicalCardProfileListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type PhysicalCardProfileListParamsStatusIn string

const (
	// The Card Profile has not yet been processed by Increase.
	PhysicalCardProfileListParamsStatusInPendingCreating PhysicalCardProfileListParamsStatusIn = "pending_creating"
	// The card profile is awaiting review by Increase.
	PhysicalCardProfileListParamsStatusInPendingReviewing PhysicalCardProfileListParamsStatusIn = "pending_reviewing"
	// There is an issue with the Physical Card Profile preventing it from use.
	PhysicalCardProfileListParamsStatusInRejected PhysicalCardProfileListParamsStatusIn = "rejected"
	// The card profile is awaiting submission to the fulfillment provider.
	PhysicalCardProfileListParamsStatusInPendingSubmitting PhysicalCardProfileListParamsStatusIn = "pending_submitting"
	// The Physical Card Profile has been submitted to the fulfillment provider and is
	// ready to use.
	PhysicalCardProfileListParamsStatusInActive PhysicalCardProfileListParamsStatusIn = "active"
	// The Physical Card Profile has been archived.
	PhysicalCardProfileListParamsStatusInArchived PhysicalCardProfileListParamsStatusIn = "archived"
)

type PhysicalCardProfileCloneParams struct {
	// The identifier of the File containing the physical card's carrier image.
	CarrierImageFileID param.Field[string] `json:"carrier_image_file_id"`
	// A phone number the user can contact to receive support for their card.
	ContactPhone param.Field[string] `json:"contact_phone"`
	// A description you can use to identify the Card Profile.
	Description param.Field[string] `json:"description"`
	// The identifier of the File containing the physical card's front image.
	FrontImageFileID param.Field[string] `json:"front_image_file_id"`
}

func (r PhysicalCardProfileCloneParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
