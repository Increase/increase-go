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

// CardProfileService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCardProfileService] method
// instead.
type CardProfileService struct {
	Options []option.RequestOption
}

// NewCardProfileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCardProfileService(opts ...option.RequestOption) (r *CardProfileService) {
	r = &CardProfileService{}
	r.Options = opts
	return
}

// Create a Card Profile
func (r *CardProfileService) New(ctx context.Context, body CardProfileNewParams, opts ...option.RequestOption) (res *CardProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := "card_profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Card Profile
func (r *CardProfileService) Get(ctx context.Context, cardProfileID string, opts ...option.RequestOption) (res *CardProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("card_profiles/%s", cardProfileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Card Profiles
func (r *CardProfileService) List(ctx context.Context, query CardProfileListParams, opts ...option.RequestOption) (res *shared.Page[CardProfile], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "card_profiles"
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
func (r *CardProfileService) ListAutoPaging(ctx context.Context, query CardProfileListParams, opts ...option.RequestOption) *shared.PageAutoPager[CardProfile] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Archive an Card Profile
func (r *CardProfileService) Archive(ctx context.Context, cardProfileID string, opts ...option.RequestOption) (res *CardProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("card_profiles/%s/archive", cardProfileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// This contains artwork and metadata relating to a Card's appearance in digital
// wallet apps like Apple Pay and Google Pay. For more information, see our guide
// on [digital card artwork](https://increase.com/documentation/card-art).
type CardProfile struct {
	// The Card Profile identifier.
	ID string `json:"id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A description you can use to identify the Card Profile.
	Description string `json:"description,required"`
	// How Cards should appear in digital wallets such as Apple Pay. Different wallets
	// will use these values to render card artwork appropriately for their app.
	DigitalWallets CardProfileDigitalWallets `json:"digital_wallets,required"`
	// Whether this Card Profile is the default for all cards in its Increase group.
	IsDefault bool `json:"is_default,required"`
	// How physical cards should be designed and shipped.
	PhysicalCards CardProfilePhysicalCards `json:"physical_cards,required,nullable"`
	// The status of the Card Profile.
	Status CardProfileStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_profile`.
	Type CardProfileType `json:"type,required"`
	JSON cardProfileJSON `json:"-"`
}

// cardProfileJSON contains the JSON metadata for the struct [CardProfile]
type cardProfileJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Description    apijson.Field
	DigitalWallets apijson.Field
	IsDefault      apijson.Field
	PhysicalCards  apijson.Field
	Status         apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CardProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How Cards should appear in digital wallets such as Apple Pay. Different wallets
// will use these values to render card artwork appropriately for their app.
type CardProfileDigitalWallets struct {
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
	// A user-facing description for whoever is issuing the card.
	IssuerName string `json:"issuer_name,required"`
	// The Card's text color, specified as an RGB triple.
	TextColor CardProfileDigitalWalletsTextColor `json:"text_color,required"`
	JSON      cardProfileDigitalWalletsJSON      `json:"-"`
}

// cardProfileDigitalWalletsJSON contains the JSON metadata for the struct
// [CardProfileDigitalWallets]
type cardProfileDigitalWalletsJSON struct {
	AppIconFileID         apijson.Field
	BackgroundImageFileID apijson.Field
	CardDescription       apijson.Field
	ContactEmail          apijson.Field
	ContactPhone          apijson.Field
	ContactWebsite        apijson.Field
	IssuerName            apijson.Field
	TextColor             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CardProfileDigitalWallets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Card's text color, specified as an RGB triple.
type CardProfileDigitalWalletsTextColor struct {
	// The value of the blue channel in the RGB color.
	Blue int64 `json:"blue,required"`
	// The value of the green channel in the RGB color.
	Green int64 `json:"green,required"`
	// The value of the red channel in the RGB color.
	Red  int64                                  `json:"red,required"`
	JSON cardProfileDigitalWalletsTextColorJSON `json:"-"`
}

// cardProfileDigitalWalletsTextColorJSON contains the JSON metadata for the struct
// [CardProfileDigitalWalletsTextColor]
type cardProfileDigitalWalletsTextColorJSON struct {
	Blue        apijson.Field
	Green       apijson.Field
	Red         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardProfileDigitalWalletsTextColor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How physical cards should be designed and shipped.
type CardProfilePhysicalCards struct {
	// The identifier of the File containing the physical card's back image.
	BackImageFileID string `json:"back_image_file_id,required,nullable"`
	// The identifier of the File containing the physical card's carrier image.
	CarrierImageFileID string `json:"carrier_image_file_id,required,nullable"`
	// A phone number the user can contact to receive support for their card.
	ContactPhone string `json:"contact_phone,required,nullable"`
	// The identifier of the File containing the physical card's front image.
	FrontImageFileID string `json:"front_image_file_id,required,nullable"`
	// The status of the Physical Card Profile.
	Status CardProfilePhysicalCardsStatus `json:"status,required"`
	JSON   cardProfilePhysicalCardsJSON   `json:"-"`
}

// cardProfilePhysicalCardsJSON contains the JSON metadata for the struct
// [CardProfilePhysicalCards]
type cardProfilePhysicalCardsJSON struct {
	BackImageFileID    apijson.Field
	CarrierImageFileID apijson.Field
	ContactPhone       apijson.Field
	FrontImageFileID   apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CardProfilePhysicalCards) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the Physical Card Profile.
type CardProfilePhysicalCardsStatus string

const (
	// The Card Profile is not eligible for physical cards.
	CardProfilePhysicalCardsStatusNotEligible CardProfilePhysicalCardsStatus = "not_eligible"
	// There is an issue with the Physical Card Profile preventing it from use.
	CardProfilePhysicalCardsStatusRejected CardProfilePhysicalCardsStatus = "rejected"
	// The Card Profile has not yet been processed by Increase.
	CardProfilePhysicalCardsStatusPendingCreating CardProfilePhysicalCardsStatus = "pending_creating"
	// The card profile is awaiting review by Increase.
	CardProfilePhysicalCardsStatusPendingReviewing CardProfilePhysicalCardsStatus = "pending_reviewing"
	// The card profile is awaiting submission to the fulfillment provider.
	CardProfilePhysicalCardsStatusPendingSubmitting CardProfilePhysicalCardsStatus = "pending_submitting"
	// The Physical Card Profile has been submitted to the fulfillment provider and is
	// ready to use.
	CardProfilePhysicalCardsStatusActive CardProfilePhysicalCardsStatus = "active"
)

// The status of the Card Profile.
type CardProfileStatus string

const (
	// The Card Profile is awaiting review from Increase and/or processing by card
	// networks.
	CardProfileStatusPending CardProfileStatus = "pending"
	// There is an issue with the Card Profile preventing it from use.
	CardProfileStatusRejected CardProfileStatus = "rejected"
	// The Card Profile can be assigned to Cards.
	CardProfileStatusActive CardProfileStatus = "active"
	// The Card Profile is no longer in use.
	CardProfileStatusArchived CardProfileStatus = "archived"
)

// A constant representing the object's type. For this resource it will always be
// `card_profile`.
type CardProfileType string

const (
	CardProfileTypeCardProfile CardProfileType = "card_profile"
)

type CardProfileNewParams struct {
	// A description you can use to identify the Card Profile.
	Description param.Field[string] `json:"description,required"`
	// How Cards should appear in digital wallets such as Apple Pay. Different wallets
	// will use these values to render card artwork appropriately for their app.
	DigitalWallets param.Field[CardProfileNewParamsDigitalWallets] `json:"digital_wallets,required"`
	// How physical cards should be designed and shipped.
	PhysicalCards param.Field[CardProfileNewParamsPhysicalCards] `json:"physical_cards"`
}

func (r CardProfileNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How Cards should appear in digital wallets such as Apple Pay. Different wallets
// will use these values to render card artwork appropriately for their app.
type CardProfileNewParamsDigitalWallets struct {
	// The identifier of the File containing the card's icon image.
	AppIconFileID param.Field[string] `json:"app_icon_file_id,required"`
	// The identifier of the File containing the card's front image.
	BackgroundImageFileID param.Field[string] `json:"background_image_file_id,required"`
	// A user-facing description for the card itself.
	CardDescription param.Field[string] `json:"card_description,required"`
	// A user-facing description for whoever is issuing the card.
	IssuerName param.Field[string] `json:"issuer_name,required"`
	// An email address the user can contact to receive support for their card.
	ContactEmail param.Field[string] `json:"contact_email"`
	// A phone number the user can contact to receive support for their card.
	ContactPhone param.Field[string] `json:"contact_phone"`
	// A website the user can visit to view and receive support for their card.
	ContactWebsite param.Field[string] `json:"contact_website"`
	// The Card's text color, specified as an RGB triple. The default is white.
	TextColor param.Field[CardProfileNewParamsDigitalWalletsTextColor] `json:"text_color"`
}

func (r CardProfileNewParamsDigitalWallets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The Card's text color, specified as an RGB triple. The default is white.
type CardProfileNewParamsDigitalWalletsTextColor struct {
	// The value of the blue channel in the RGB color.
	Blue param.Field[int64] `json:"blue,required"`
	// The value of the green channel in the RGB color.
	Green param.Field[int64] `json:"green,required"`
	// The value of the red channel in the RGB color.
	Red param.Field[int64] `json:"red,required"`
}

func (r CardProfileNewParamsDigitalWalletsTextColor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How physical cards should be designed and shipped.
type CardProfileNewParamsPhysicalCards struct {
	// The identifier of the File containing the physical card's carrier image.
	CarrierImageFileID param.Field[string] `json:"carrier_image_file_id,required"`
	// A phone number the user can contact to receive support for their card.
	ContactPhone param.Field[string] `json:"contact_phone,required"`
	// The identifier of the File containing the physical card's front image.
	FrontImageFileID param.Field[string] `json:"front_image_file_id,required"`
}

func (r CardProfileNewParamsPhysicalCards) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardProfileListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit               param.Field[int64]                                    `query:"limit"`
	PhysicalCardsStatus param.Field[CardProfileListParamsPhysicalCardsStatus] `query:"physical_cards_status"`
	Status              param.Field[CardProfileListParamsStatus]              `query:"status"`
}

// URLQuery serializes [CardProfileListParams]'s query parameters as `url.Values`.
func (r CardProfileListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CardProfileListParamsPhysicalCardsStatus struct {
	// Filter Card Profiles for those with the specified physical card status or
	// statuses. For GET requests, this should be encoded as a comma-delimited string,
	// such as `?in=one,two,three`.
	In param.Field[[]CardProfileListParamsPhysicalCardsStatusIn] `query:"in"`
}

// URLQuery serializes [CardProfileListParamsPhysicalCardsStatus]'s query
// parameters as `url.Values`.
func (r CardProfileListParamsPhysicalCardsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CardProfileListParamsPhysicalCardsStatusIn string

const (
	// The Card Profile is not eligible for physical cards.
	CardProfileListParamsPhysicalCardsStatusInNotEligible CardProfileListParamsPhysicalCardsStatusIn = "not_eligible"
	// There is an issue with the Physical Card Profile preventing it from use.
	CardProfileListParamsPhysicalCardsStatusInRejected CardProfileListParamsPhysicalCardsStatusIn = "rejected"
	// The Card Profile has not yet been processed by Increase.
	CardProfileListParamsPhysicalCardsStatusInPendingCreating CardProfileListParamsPhysicalCardsStatusIn = "pending_creating"
	// The card profile is awaiting review by Increase.
	CardProfileListParamsPhysicalCardsStatusInPendingReviewing CardProfileListParamsPhysicalCardsStatusIn = "pending_reviewing"
	// The card profile is awaiting submission to the fulfillment provider.
	CardProfileListParamsPhysicalCardsStatusInPendingSubmitting CardProfileListParamsPhysicalCardsStatusIn = "pending_submitting"
	// The Physical Card Profile has been submitted to the fulfillment provider and is
	// ready to use.
	CardProfileListParamsPhysicalCardsStatusInActive CardProfileListParamsPhysicalCardsStatusIn = "active"
)

type CardProfileListParamsStatus struct {
	// Filter Card Profiles for those with the specified digital wallet status or
	// statuses. For GET requests, this should be encoded as a comma-delimited string,
	// such as `?in=one,two,three`.
	In param.Field[[]CardProfileListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [CardProfileListParamsStatus]'s query parameters as
// `url.Values`.
func (r CardProfileListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CardProfileListParamsStatusIn string

const (
	// The Card Profile is awaiting review from Increase and/or processing by card
	// networks.
	CardProfileListParamsStatusInPending CardProfileListParamsStatusIn = "pending"
	// There is an issue with the Card Profile preventing it from use.
	CardProfileListParamsStatusInRejected CardProfileListParamsStatusIn = "rejected"
	// The Card Profile can be assigned to Cards.
	CardProfileListParamsStatusInActive CardProfileListParamsStatusIn = "active"
	// The Card Profile is no longer in use.
	CardProfileListParamsStatusInArchived CardProfileListParamsStatusIn = "archived"
)
