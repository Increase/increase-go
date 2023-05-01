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

type CardProfileService struct {
	Options []option.RequestOption
}

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
	JSON CardProfileJSON
}

type CardProfileJSON struct {
	ID             apijson.Metadata
	CreatedAt      apijson.Metadata
	Status         apijson.Metadata
	Description    apijson.Metadata
	DigitalWallets apijson.Metadata
	Type           apijson.Metadata
	raw            string
	Extras         map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardProfile using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
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
	JSON          CardProfileDigitalWalletsJSON
}

type CardProfileDigitalWalletsJSON struct {
	TextColor             apijson.Metadata
	IssuerName            apijson.Metadata
	CardDescription       apijson.Metadata
	ContactWebsite        apijson.Metadata
	ContactEmail          apijson.Metadata
	ContactPhone          apijson.Metadata
	BackgroundImageFileID apijson.Metadata
	AppIconFileID         apijson.Metadata
	raw                   string
	Extras                map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardProfileDigitalWallets
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardProfileDigitalWallets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardProfileDigitalWalletsTextColor struct {
	// The value of the red channel in the RGB color.
	Red int64 `json:"red,required"`
	// The value of the green channel in the RGB color.
	Green int64 `json:"green,required"`
	// The value of the blue channel in the RGB color.
	Blue int64 `json:"blue,required"`
	JSON CardProfileDigitalWalletsTextColorJSON
}

type CardProfileDigitalWalletsTextColorJSON struct {
	Red    apijson.Metadata
	Green  apijson.Metadata
	Blue   apijson.Metadata
	raw    string
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardProfileDigitalWalletsTextColor using the internal json library. Unrecognized
// fields are stored in the `jsonFields` property.
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

// MarshalJSON serializes CardProfileNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r CardProfileNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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

// URLQuery serializes CardProfileListParams into a url.Values of the query
// parameters associated with this value
func (r CardProfileListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type CardProfileListParamsStatus struct {
	// Filter Card Profiles for those with the specified status or statuses. For GET
	// requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In field.Field[[]CardProfileListParamsStatusIn] `query:"in"`
}

// URLQuery serializes CardProfileListParamsStatus into a url.Values of the query
// parameters associated with this value
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

type CardProfileListResponse struct {
	// The contents of the list.
	Data []CardProfile `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       CardProfileListResponseJSON
}

type CardProfileListResponseJSON struct {
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardProfileListResponse using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardProfileListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
