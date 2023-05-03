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
func (r *CardProfileService) Get(ctx context.Context, card_profile_id string, opts ...option.RequestOption) (res *CardProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("card_profiles/%s", card_profile_id)
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

// This contains artwork and metadata relating to a Card's appearance in digital
// wallet apps like Apple Pay and Google Pay. For more information, see our guide
// on [digital card artwork](https://increase.com/documentation/card-art)
type CardProfile struct {
	// The Card Profile identifier.
	ID string `json:"id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The status of the Card Profile.
	Status CardProfileStatus `json:"status,required"`
	// A description you can use to identify the Card Profile.
	Description string `json:"description,required"`
	// How Cards should appear in digital wallets such as Apple Pay. Different wallets
	// will use these values to render card artwork appropriately for their app.
	DigitalWallets CardProfileDigitalWallets `json:"digital_wallets,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_profile`.
	Type CardProfileType `json:"type,required"`
	JSON cardProfileJSON
}

// cardProfileJSON contains the JSON metadata for the struct [CardProfile]
type cardProfileJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Status         apijson.Field
	Description    apijson.Field
	DigitalWallets apijson.Field
	Type           apijson.Field
	raw            string
	Extras         map[string]apijson.Field
}

func (r *CardProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardProfileStatus string

const (
	CardProfileStatusPending  CardProfileStatus = "pending"
	CardProfileStatusRejected CardProfileStatus = "rejected"
	CardProfileStatusActive   CardProfileStatus = "active"
	CardProfileStatusArchived CardProfileStatus = "archived"
)

// How Cards should appear in digital wallets such as Apple Pay. Different wallets
// will use these values to render card artwork appropriately for their app.
type CardProfileDigitalWallets struct {
	// The Card's text color, specified as an RGB triple.
	TextColor CardProfileDigitalWalletsTextColor `json:"text_color,required"`
	// A user-facing description for whoever is issuing the card.
	IssuerName string `json:"issuer_name,required"`
	// A user-facing description for the card itself.
	CardDescription string `json:"card_description,required"`
	// A website the user can visit to view and receive support for their card.
	ContactWebsite string `json:"contact_website,required,nullable"`
	// An email address the user can contact to receive support for their card.
	ContactEmail string `json:"contact_email,required,nullable"`
	// A phone number the user can contact to receive support for their card.
	ContactPhone string `json:"contact_phone,required,nullable"`
	// The identifier of the File containing the card's front image.
	BackgroundImageFileID string `json:"background_image_file_id,required"`
	// The identifier of the File containing the card's icon image.
	AppIconFileID string `json:"app_icon_file_id,required"`
	JSON          cardProfileDigitalWalletsJSON
}

// cardProfileDigitalWalletsJSON contains the JSON metadata for the struct
// [CardProfileDigitalWallets]
type cardProfileDigitalWalletsJSON struct {
	TextColor             apijson.Field
	IssuerName            apijson.Field
	CardDescription       apijson.Field
	ContactWebsite        apijson.Field
	ContactEmail          apijson.Field
	ContactPhone          apijson.Field
	BackgroundImageFileID apijson.Field
	AppIconFileID         apijson.Field
	raw                   string
	Extras                map[string]apijson.Field
}

func (r *CardProfileDigitalWallets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Card's text color, specified as an RGB triple.
type CardProfileDigitalWalletsTextColor struct {
	// The value of the red channel in the RGB color.
	Red int64 `json:"red,required"`
	// The value of the green channel in the RGB color.
	Green int64 `json:"green,required"`
	// The value of the blue channel in the RGB color.
	Blue int64 `json:"blue,required"`
	JSON cardProfileDigitalWalletsTextColorJSON
}

// cardProfileDigitalWalletsTextColorJSON contains the JSON metadata for the struct
// [CardProfileDigitalWalletsTextColor]
type cardProfileDigitalWalletsTextColorJSON struct {
	Red    apijson.Field
	Green  apijson.Field
	Blue   apijson.Field
	raw    string
	Extras map[string]apijson.Field
}

func (r *CardProfileDigitalWalletsTextColor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardProfileType string

const (
	CardProfileTypeCardProfile CardProfileType = "card_profile"
)

type CardProfileNewParams struct {
	// A description you can use to identify the Card Profile.
	Description field.Field[string] `json:"description,required"`
	// How Cards should appear in digital wallets such as Apple Pay. Different wallets
	// will use these values to render card artwork appropriately for their app.
	DigitalWallets field.Field[CardProfileNewParamsDigitalWallets] `json:"digital_wallets,required"`
}

func (r CardProfileNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How Cards should appear in digital wallets such as Apple Pay. Different wallets
// will use these values to render card artwork appropriately for their app.
type CardProfileNewParamsDigitalWallets struct {
	// The Card's text color, specified as an RGB triple. The default is white.
	TextColor field.Field[CardProfileNewParamsDigitalWalletsTextColor] `json:"text_color"`
	// A user-facing description for whoever is issuing the card.
	IssuerName field.Field[string] `json:"issuer_name,required"`
	// A user-facing description for the card itself.
	CardDescription field.Field[string] `json:"card_description,required"`
	// A website the user can visit to view and receive support for their card.
	ContactWebsite field.Field[string] `json:"contact_website"`
	// An email address the user can contact to receive support for their card.
	ContactEmail field.Field[string] `json:"contact_email"`
	// A phone number the user can contact to receive support for their card.
	ContactPhone field.Field[string] `json:"contact_phone"`
	// The identifier of the File containing the card's front image.
	BackgroundImageFileID field.Field[string] `json:"background_image_file_id,required"`
	// The identifier of the File containing the card's icon image.
	AppIconFileID field.Field[string] `json:"app_icon_file_id,required"`
}

// The Card's text color, specified as an RGB triple. The default is white.
type CardProfileNewParamsDigitalWalletsTextColor struct {
	// The value of the red channel in the RGB color.
	Red field.Field[int64] `json:"red,required"`
	// The value of the green channel in the RGB color.
	Green field.Field[int64] `json:"green,required"`
	// The value of the blue channel in the RGB color.
	Blue field.Field[int64] `json:"blue,required"`
}

type CardProfileListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  field.Field[int64]                       `query:"limit"`
	Status field.Field[CardProfileListParamsStatus] `query:"status"`
}

// URLQuery serializes [CardProfileListParams]'s query parameters as `url.Values`.
func (r CardProfileListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type CardProfileListParamsStatus struct {
	// Filter Card Profiles for those with the specified status or statuses. For GET
	// requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In field.Field[[]CardProfileListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [CardProfileListParamsStatus]'s query parameters as
// `url.Values`.
func (r CardProfileListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type CardProfileListParamsStatusIn string

const (
	CardProfileListParamsStatusInPending  CardProfileListParamsStatusIn = "pending"
	CardProfileListParamsStatusInRejected CardProfileListParamsStatusIn = "rejected"
	CardProfileListParamsStatusInActive   CardProfileListParamsStatusIn = "active"
	CardProfileListParamsStatusInArchived CardProfileListParamsStatusIn = "archived"
)

// A list of Card Profile objects
type CardProfileListResponse struct {
	// The contents of the list.
	Data []CardProfile `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       cardProfileListResponseJSON
}

// cardProfileListResponseJSON contains the JSON metadata for the struct
// [CardProfileListResponse]
type cardProfileListResponseJSON struct {
	Data       apijson.Field
	NextCursor apijson.Field
	raw        string
	Extras     map[string]apijson.Field
}

func (r *CardProfileListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
