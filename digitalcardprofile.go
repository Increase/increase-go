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

// DigitalCardProfileService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDigitalCardProfileService] method
// instead.
type DigitalCardProfileService struct {
	Options []option.RequestOption
}

// NewDigitalCardProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDigitalCardProfileService(opts ...option.RequestOption) (r *DigitalCardProfileService) {
	r = &DigitalCardProfileService{}
	r.Options = opts
	return
}

// Create a Digital Card Profile
func (r *DigitalCardProfileService) New(ctx context.Context, body DigitalCardProfileNewParams, opts ...option.RequestOption) (res *DigitalCardProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := "digital_card_profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Digital Card Profile
func (r *DigitalCardProfileService) Get(ctx context.Context, digitalCardProfileID string, opts ...option.RequestOption) (res *DigitalCardProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("digital_card_profiles/%s", digitalCardProfileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Card Profiles
func (r *DigitalCardProfileService) List(ctx context.Context, query DigitalCardProfileListParams, opts ...option.RequestOption) (res *shared.Page[DigitalCardProfile], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "digital_card_profiles"
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

// List Card Profiles
func (r *DigitalCardProfileService) ListAutoPaging(ctx context.Context, query DigitalCardProfileListParams, opts ...option.RequestOption) *shared.PageAutoPager[DigitalCardProfile] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Archive a Digital Card Profile
func (r *DigitalCardProfileService) Archive(ctx context.Context, digitalCardProfileID string, opts ...option.RequestOption) (res *DigitalCardProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("digital_card_profiles/%s/archive", digitalCardProfileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Clones a Digital Card Profile
func (r *DigitalCardProfileService) Clone(ctx context.Context, digitalCardProfileID string, body DigitalCardProfileCloneParams, opts ...option.RequestOption) (res *DigitalCardProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("digital_card_profiles/%s/clone", digitalCardProfileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This contains artwork and metadata relating to a Card's appearance in digital
// wallet apps like Apple Pay and Google Pay. For more information, see our guide
// on [digital card artwork](https://increase.com/documentation/card-art).
type DigitalCardProfile struct {
	// The Card Profile identifier.
	ID string `json:"id,required"`
	// The identifier of the File containing the card's icon image.
	AppIconFileID string `json:"app_icon_file_id,required"`
	// The identifier of the File containing the card's front image.
	BackgroundImageFileID string `json:"background_image_file_id,required"`
	// A user-facing description for the card itself.
	CardDescription string `json:"card_description,required"`
	// An email address the user can contact to receive support for their card.
	ContactEmail string `json:"contact_email,required,nullable"`
	// A phone number the user can contact to receive support for their card.
	ContactPhone string `json:"contact_phone,required,nullable"`
	// A website the user can visit to view and receive support for their card.
	ContactWebsite string `json:"contact_website,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A description you can use to identify the Card Profile.
	Description string `json:"description,required"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// Whether this Digital Card Profile is the default for all cards in its Increase
	// group.
	IsDefault bool `json:"is_default,required"`
	// A user-facing description for whoever is issuing the card.
	IssuerName string `json:"issuer_name,required"`
	// The status of the Card Profile.
	Status DigitalCardProfileStatus `json:"status,required"`
	// The Card's text color, specified as an RGB triple.
	TextColor DigitalCardProfileTextColor `json:"text_color,required"`
	// A constant representing the object's type. For this resource it will always be
	// `digital_card_profile`.
	Type DigitalCardProfileType `json:"type,required"`
	JSON digitalCardProfileJSON `json:"-"`
}

// digitalCardProfileJSON contains the JSON metadata for the struct
// [DigitalCardProfile]
type digitalCardProfileJSON struct {
	ID                    apijson.Field
	AppIconFileID         apijson.Field
	BackgroundImageFileID apijson.Field
	CardDescription       apijson.Field
	ContactEmail          apijson.Field
	ContactPhone          apijson.Field
	ContactWebsite        apijson.Field
	CreatedAt             apijson.Field
	Description           apijson.Field
	IdempotencyKey        apijson.Field
	IsDefault             apijson.Field
	IssuerName            apijson.Field
	Status                apijson.Field
	TextColor             apijson.Field
	Type                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *DigitalCardProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the Card Profile.
type DigitalCardProfileStatus string

const (
	// The Card Profile is awaiting review from Increase and/or processing by card
	// networks.
	DigitalCardProfileStatusPending DigitalCardProfileStatus = "pending"
	// There is an issue with the Card Profile preventing it from use.
	DigitalCardProfileStatusRejected DigitalCardProfileStatus = "rejected"
	// The Card Profile can be assigned to Cards.
	DigitalCardProfileStatusActive DigitalCardProfileStatus = "active"
	// The Card Profile is no longer in use.
	DigitalCardProfileStatusArchived DigitalCardProfileStatus = "archived"
)

// The Card's text color, specified as an RGB triple.
type DigitalCardProfileTextColor struct {
	// The value of the blue channel in the RGB color.
	Blue int64 `json:"blue,required"`
	// The value of the green channel in the RGB color.
	Green int64 `json:"green,required"`
	// The value of the red channel in the RGB color.
	Red  int64                           `json:"red,required"`
	JSON digitalCardProfileTextColorJSON `json:"-"`
}

// digitalCardProfileTextColorJSON contains the JSON metadata for the struct
// [DigitalCardProfileTextColor]
type digitalCardProfileTextColorJSON struct {
	Blue        apijson.Field
	Green       apijson.Field
	Red         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalCardProfileTextColor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A constant representing the object's type. For this resource it will always be
// `digital_card_profile`.
type DigitalCardProfileType string

const (
	DigitalCardProfileTypeDigitalCardProfile DigitalCardProfileType = "digital_card_profile"
)

type DigitalCardProfileNewParams struct {
	// The identifier of the File containing the card's icon image.
	AppIconFileID param.Field[string] `json:"app_icon_file_id,required"`
	// The identifier of the File containing the card's front image.
	BackgroundImageFileID param.Field[string] `json:"background_image_file_id,required"`
	// A user-facing description for the card itself.
	CardDescription param.Field[string] `json:"card_description,required"`
	// A description you can use to identify the Card Profile.
	Description param.Field[string] `json:"description,required"`
	// A user-facing description for whoever is issuing the card.
	IssuerName param.Field[string] `json:"issuer_name,required"`
	// An email address the user can contact to receive support for their card.
	ContactEmail param.Field[string] `json:"contact_email"`
	// A phone number the user can contact to receive support for their card.
	ContactPhone param.Field[string] `json:"contact_phone"`
	// A website the user can visit to view and receive support for their card.
	ContactWebsite param.Field[string] `json:"contact_website"`
	// The Card's text color, specified as an RGB triple. The default is white.
	TextColor param.Field[DigitalCardProfileNewParamsTextColor] `json:"text_color"`
}

func (r DigitalCardProfileNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The Card's text color, specified as an RGB triple. The default is white.
type DigitalCardProfileNewParamsTextColor struct {
	// The value of the blue channel in the RGB color.
	Blue param.Field[int64] `json:"blue,required"`
	// The value of the green channel in the RGB color.
	Green param.Field[int64] `json:"green,required"`
	// The value of the red channel in the RGB color.
	Red param.Field[int64] `json:"red,required"`
}

func (r DigitalCardProfileNewParamsTextColor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DigitalCardProfileListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  param.Field[int64]                              `query:"limit"`
	Status param.Field[DigitalCardProfileListParamsStatus] `query:"status"`
}

// URLQuery serializes [DigitalCardProfileListParams]'s query parameters as
// `url.Values`.
func (r DigitalCardProfileListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DigitalCardProfileListParamsStatus struct {
	// Filter Digital Card Profiles for those with the specified digital wallet status
	// or statuses. For GET requests, this should be encoded as a comma-delimited
	// string, such as `?in=one,two,three`.
	In param.Field[[]DigitalCardProfileListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [DigitalCardProfileListParamsStatus]'s query parameters as
// `url.Values`.
func (r DigitalCardProfileListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DigitalCardProfileListParamsStatusIn string

const (
	// The Card Profile is awaiting review from Increase and/or processing by card
	// networks.
	DigitalCardProfileListParamsStatusInPending DigitalCardProfileListParamsStatusIn = "pending"
	// There is an issue with the Card Profile preventing it from use.
	DigitalCardProfileListParamsStatusInRejected DigitalCardProfileListParamsStatusIn = "rejected"
	// The Card Profile can be assigned to Cards.
	DigitalCardProfileListParamsStatusInActive DigitalCardProfileListParamsStatusIn = "active"
	// The Card Profile is no longer in use.
	DigitalCardProfileListParamsStatusInArchived DigitalCardProfileListParamsStatusIn = "archived"
)

type DigitalCardProfileCloneParams struct {
	// The identifier of the File containing the card's icon image.
	AppIconFileID param.Field[string] `json:"app_icon_file_id"`
	// The identifier of the File containing the card's front image.
	BackgroundImageFileID param.Field[string] `json:"background_image_file_id"`
	// A user-facing description for the card itself.
	CardDescription param.Field[string] `json:"card_description"`
	// An email address the user can contact to receive support for their card.
	ContactEmail param.Field[string] `json:"contact_email"`
	// A phone number the user can contact to receive support for their card.
	ContactPhone param.Field[string] `json:"contact_phone"`
	// A website the user can visit to view and receive support for their card.
	ContactWebsite param.Field[string] `json:"contact_website"`
	// A description you can use to identify the Card Profile.
	Description param.Field[string] `json:"description"`
	// A user-facing description for whoever is issuing the card.
	IssuerName param.Field[string] `json:"issuer_name"`
	// The Card's text color, specified as an RGB triple. The default is white.
	TextColor param.Field[DigitalCardProfileCloneParamsTextColor] `json:"text_color"`
}

func (r DigitalCardProfileCloneParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The Card's text color, specified as an RGB triple. The default is white.
type DigitalCardProfileCloneParamsTextColor struct {
	// The value of the blue channel in the RGB color.
	Blue param.Field[int64] `json:"blue,required"`
	// The value of the green channel in the RGB color.
	Green param.Field[int64] `json:"green,required"`
	// The value of the red channel in the RGB color.
	Red param.Field[int64] `json:"red,required"`
}

func (r DigitalCardProfileCloneParamsTextColor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
