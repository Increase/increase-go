package types

import (
	"fmt"
	"net/url"
	"time"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/pagination"
)

type CardProfile struct {
	// The Card Profile identifier.
	ID *string `pjson:"id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was created.
	CreatedAt *time.Time `pjson:"created_at" format:"2006-01-02T15:04:05Z07:00"`
	// The status of the Card Profile.
	Status *CardProfileStatus `pjson:"status"`
	// A description you can use to identify the Card Profile.
	Description *string `pjson:"description"`
	// How Cards should appear in digital wallets such as Apple Pay. Different wallets
	// will use these values to render card artwork appropriately for their app.
	DigitalWallets *CardProfileDigitalWallets `pjson:"digital_wallets"`
	// A constant representing the object's type. For this resource it will always be
	// `card_profile`.
	Type       *CardProfileType       `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CardProfile using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardProfile) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardProfile into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CardProfile) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The Card Profile identifier.
func (r CardProfile) GetID() (ID string) {
	if r.ID != nil {
		ID = *r.ID
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card Dispute was created.
func (r CardProfile) GetCreatedAt() (CreatedAt time.Time) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The status of the Card Profile.
func (r CardProfile) GetStatus() (Status CardProfileStatus) {
	if r.Status != nil {
		Status = *r.Status
	}
	return
}

// A description you can use to identify the Card Profile.
func (r CardProfile) GetDescription() (Description string) {
	if r.Description != nil {
		Description = *r.Description
	}
	return
}

// How Cards should appear in digital wallets such as Apple Pay. Different wallets
// will use these values to render card artwork appropriately for their app.
func (r CardProfile) GetDigitalWallets() (DigitalWallets CardProfileDigitalWallets) {
	if r.DigitalWallets != nil {
		DigitalWallets = *r.DigitalWallets
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_profile`.
func (r CardProfile) GetType() (Type CardProfileType) {
	if r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r CardProfile) String() (result string) {
	return fmt.Sprintf("&CardProfile{ID:%s CreatedAt:%s Status:%s Description:%s DigitalWallets:%s Type:%s}", core.FmtP(r.ID), core.FmtP(r.CreatedAt), core.FmtP(r.Status), core.FmtP(r.Description), r.DigitalWallets, core.FmtP(r.Type))
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
	TextColor *CardProfileDigitalWalletsTextColor `pjson:"text_color"`
	// A user-facing description for whoever is issuing the card.
	IssuerName *string `pjson:"issuer_name"`
	// A user-facing description for the card itself.
	CardDescription *string `pjson:"card_description"`
	// A website the user can visit to view and receive support for their card.
	ContactWebsite *string `pjson:"contact_website"`
	// An email address the user can contact to receive support for their card.
	ContactEmail *string `pjson:"contact_email"`
	// A phone number the user can contact to receive support for their card.
	ContactPhone *string `pjson:"contact_phone"`
	// The identifier of the File containing the card's front image.
	BackgroundImageFileID *string `pjson:"background_image_file_id"`
	// The identifier of the File containing the card's icon image.
	AppIconFileID *string                `pjson:"app_icon_file_id"`
	jsonFields    map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CardProfileDigitalWallets
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardProfileDigitalWallets) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardProfileDigitalWallets into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CardProfileDigitalWallets) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The Card's text color, specified as an RGB triple.
func (r CardProfileDigitalWallets) GetTextColor() (TextColor CardProfileDigitalWalletsTextColor) {
	if r.TextColor != nil {
		TextColor = *r.TextColor
	}
	return
}

// A user-facing description for whoever is issuing the card.
func (r CardProfileDigitalWallets) GetIssuerName() (IssuerName string) {
	if r.IssuerName != nil {
		IssuerName = *r.IssuerName
	}
	return
}

// A user-facing description for the card itself.
func (r CardProfileDigitalWallets) GetCardDescription() (CardDescription string) {
	if r.CardDescription != nil {
		CardDescription = *r.CardDescription
	}
	return
}

// A website the user can visit to view and receive support for their card.
func (r CardProfileDigitalWallets) GetContactWebsite() (ContactWebsite string) {
	if r.ContactWebsite != nil {
		ContactWebsite = *r.ContactWebsite
	}
	return
}

// An email address the user can contact to receive support for their card.
func (r CardProfileDigitalWallets) GetContactEmail() (ContactEmail string) {
	if r.ContactEmail != nil {
		ContactEmail = *r.ContactEmail
	}
	return
}

// A phone number the user can contact to receive support for their card.
func (r CardProfileDigitalWallets) GetContactPhone() (ContactPhone string) {
	if r.ContactPhone != nil {
		ContactPhone = *r.ContactPhone
	}
	return
}

// The identifier of the File containing the card's front image.
func (r CardProfileDigitalWallets) GetBackgroundImageFileID() (BackgroundImageFileID string) {
	if r.BackgroundImageFileID != nil {
		BackgroundImageFileID = *r.BackgroundImageFileID
	}
	return
}

// The identifier of the File containing the card's icon image.
func (r CardProfileDigitalWallets) GetAppIconFileID() (AppIconFileID string) {
	if r.AppIconFileID != nil {
		AppIconFileID = *r.AppIconFileID
	}
	return
}

func (r CardProfileDigitalWallets) String() (result string) {
	return fmt.Sprintf("&CardProfileDigitalWallets{TextColor:%s IssuerName:%s CardDescription:%s ContactWebsite:%s ContactEmail:%s ContactPhone:%s BackgroundImageFileID:%s AppIconFileID:%s}", r.TextColor, core.FmtP(r.IssuerName), core.FmtP(r.CardDescription), core.FmtP(r.ContactWebsite), core.FmtP(r.ContactEmail), core.FmtP(r.ContactPhone), core.FmtP(r.BackgroundImageFileID), core.FmtP(r.AppIconFileID))
}

type CardProfileDigitalWalletsTextColor struct {
	// The value of the red channel in the RGB color.
	Red *int64 `pjson:"red"`
	// The value of the green channel in the RGB color.
	Green *int64 `pjson:"green"`
	// The value of the blue channel in the RGB color.
	Blue       *int64                 `pjson:"blue"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CardProfileDigitalWalletsTextColor using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *CardProfileDigitalWalletsTextColor) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardProfileDigitalWalletsTextColor into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CardProfileDigitalWalletsTextColor) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The value of the red channel in the RGB color.
func (r CardProfileDigitalWalletsTextColor) GetRed() (Red int64) {
	if r.Red != nil {
		Red = *r.Red
	}
	return
}

// The value of the green channel in the RGB color.
func (r CardProfileDigitalWalletsTextColor) GetGreen() (Green int64) {
	if r.Green != nil {
		Green = *r.Green
	}
	return
}

// The value of the blue channel in the RGB color.
func (r CardProfileDigitalWalletsTextColor) GetBlue() (Blue int64) {
	if r.Blue != nil {
		Blue = *r.Blue
	}
	return
}

func (r CardProfileDigitalWalletsTextColor) String() (result string) {
	return fmt.Sprintf("&CardProfileDigitalWalletsTextColor{Red:%s Green:%s Blue:%s}", core.FmtP(r.Red), core.FmtP(r.Green), core.FmtP(r.Blue))
}

type CardProfileType string

const (
	CardProfileTypeCardProfile CardProfileType = "card_profile"
)

type CreateACardProfileParameters struct {
	// A description you can use to identify the Card Profile.
	Description *string `pjson:"description"`
	// How Cards should appear in digital wallets such as Apple Pay. Different wallets
	// will use these values to render card artwork appropriately for their app.
	DigitalWallets *CreateACardProfileParametersDigitalWallets `pjson:"digital_wallets"`
	jsonFields     map[string]interface{}                      `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CreateACardProfileParameters
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CreateACardProfileParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateACardProfileParameters into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CreateACardProfileParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// A description you can use to identify the Card Profile.
func (r CreateACardProfileParameters) GetDescription() (Description string) {
	if r.Description != nil {
		Description = *r.Description
	}
	return
}

// How Cards should appear in digital wallets such as Apple Pay. Different wallets
// will use these values to render card artwork appropriately for their app.
func (r CreateACardProfileParameters) GetDigitalWallets() (DigitalWallets CreateACardProfileParametersDigitalWallets) {
	if r.DigitalWallets != nil {
		DigitalWallets = *r.DigitalWallets
	}
	return
}

func (r CreateACardProfileParameters) String() (result string) {
	return fmt.Sprintf("&CreateACardProfileParameters{Description:%s DigitalWallets:%s}", core.FmtP(r.Description), r.DigitalWallets)
}

type CreateACardProfileParametersDigitalWallets struct {
	// The Card's text color, specified as an RGB triple. The default is white.
	TextColor *CreateACardProfileParametersDigitalWalletsTextColor `pjson:"text_color"`
	// A user-facing description for whoever is issuing the card.
	IssuerName *string `pjson:"issuer_name"`
	// A user-facing description for the card itself.
	CardDescription *string `pjson:"card_description"`
	// A website the user can visit to view and receive support for their card.
	ContactWebsite *string `pjson:"contact_website"`
	// An email address the user can contact to receive support for their card.
	ContactEmail *string `pjson:"contact_email"`
	// A phone number the user can contact to receive support for their card.
	ContactPhone *string `pjson:"contact_phone"`
	// The identifier of the File containing the card's front image.
	BackgroundImageFileID *string `pjson:"background_image_file_id"`
	// The identifier of the File containing the card's icon image.
	AppIconFileID *string                `pjson:"app_icon_file_id"`
	jsonFields    map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateACardProfileParametersDigitalWallets using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *CreateACardProfileParametersDigitalWallets) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateACardProfileParametersDigitalWallets into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateACardProfileParametersDigitalWallets) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The Card's text color, specified as an RGB triple. The default is white.
func (r CreateACardProfileParametersDigitalWallets) GetTextColor() (TextColor CreateACardProfileParametersDigitalWalletsTextColor) {
	if r.TextColor != nil {
		TextColor = *r.TextColor
	}
	return
}

// A user-facing description for whoever is issuing the card.
func (r CreateACardProfileParametersDigitalWallets) GetIssuerName() (IssuerName string) {
	if r.IssuerName != nil {
		IssuerName = *r.IssuerName
	}
	return
}

// A user-facing description for the card itself.
func (r CreateACardProfileParametersDigitalWallets) GetCardDescription() (CardDescription string) {
	if r.CardDescription != nil {
		CardDescription = *r.CardDescription
	}
	return
}

// A website the user can visit to view and receive support for their card.
func (r CreateACardProfileParametersDigitalWallets) GetContactWebsite() (ContactWebsite string) {
	if r.ContactWebsite != nil {
		ContactWebsite = *r.ContactWebsite
	}
	return
}

// An email address the user can contact to receive support for their card.
func (r CreateACardProfileParametersDigitalWallets) GetContactEmail() (ContactEmail string) {
	if r.ContactEmail != nil {
		ContactEmail = *r.ContactEmail
	}
	return
}

// A phone number the user can contact to receive support for their card.
func (r CreateACardProfileParametersDigitalWallets) GetContactPhone() (ContactPhone string) {
	if r.ContactPhone != nil {
		ContactPhone = *r.ContactPhone
	}
	return
}

// The identifier of the File containing the card's front image.
func (r CreateACardProfileParametersDigitalWallets) GetBackgroundImageFileID() (BackgroundImageFileID string) {
	if r.BackgroundImageFileID != nil {
		BackgroundImageFileID = *r.BackgroundImageFileID
	}
	return
}

// The identifier of the File containing the card's icon image.
func (r CreateACardProfileParametersDigitalWallets) GetAppIconFileID() (AppIconFileID string) {
	if r.AppIconFileID != nil {
		AppIconFileID = *r.AppIconFileID
	}
	return
}

func (r CreateACardProfileParametersDigitalWallets) String() (result string) {
	return fmt.Sprintf("&CreateACardProfileParametersDigitalWallets{TextColor:%s IssuerName:%s CardDescription:%s ContactWebsite:%s ContactEmail:%s ContactPhone:%s BackgroundImageFileID:%s AppIconFileID:%s}", r.TextColor, core.FmtP(r.IssuerName), core.FmtP(r.CardDescription), core.FmtP(r.ContactWebsite), core.FmtP(r.ContactEmail), core.FmtP(r.ContactPhone), core.FmtP(r.BackgroundImageFileID), core.FmtP(r.AppIconFileID))
}

type CreateACardProfileParametersDigitalWalletsTextColor struct {
	// The value of the red channel in the RGB color.
	Red *int64 `pjson:"red"`
	// The value of the green channel in the RGB color.
	Green *int64 `pjson:"green"`
	// The value of the blue channel in the RGB color.
	Blue       *int64                 `pjson:"blue"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateACardProfileParametersDigitalWalletsTextColor using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *CreateACardProfileParametersDigitalWalletsTextColor) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateACardProfileParametersDigitalWalletsTextColor into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateACardProfileParametersDigitalWalletsTextColor) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The value of the red channel in the RGB color.
func (r CreateACardProfileParametersDigitalWalletsTextColor) GetRed() (Red int64) {
	if r.Red != nil {
		Red = *r.Red
	}
	return
}

// The value of the green channel in the RGB color.
func (r CreateACardProfileParametersDigitalWalletsTextColor) GetGreen() (Green int64) {
	if r.Green != nil {
		Green = *r.Green
	}
	return
}

// The value of the blue channel in the RGB color.
func (r CreateACardProfileParametersDigitalWalletsTextColor) GetBlue() (Blue int64) {
	if r.Blue != nil {
		Blue = *r.Blue
	}
	return
}

func (r CreateACardProfileParametersDigitalWalletsTextColor) String() (result string) {
	return fmt.Sprintf("&CreateACardProfileParametersDigitalWalletsTextColor{Red:%s Green:%s Blue:%s}", core.FmtP(r.Red), core.FmtP(r.Green), core.FmtP(r.Blue))
}

type CardProfileListParams struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit      *int64                       `query:"limit"`
	Status     *CardProfileListParamsStatus `query:"status"`
	jsonFields map[string]interface{}       `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CardProfileListParams using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardProfileListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardProfileListParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CardProfileListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes CardProfileListParams into a url.Values of the query
// parameters associated with this value
func (r *CardProfileListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r CardProfileListParams) GetCursor() (Cursor string) {
	if r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r CardProfileListParams) GetLimit() (Limit int64) {
	if r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

func (r CardProfileListParams) GetStatus() (Status CardProfileListParamsStatus) {
	if r.Status != nil {
		Status = *r.Status
	}
	return
}

func (r CardProfileListParams) String() (result string) {
	return fmt.Sprintf("&CardProfileListParams{Cursor:%s Limit:%s Status:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), r.Status)
}

type CardProfileListParamsStatus struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In         *[]CardProfileListParamsStatusIn `pjson:"in"`
	jsonFields map[string]interface{}           `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CardProfileListParamsStatus
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardProfileListParamsStatus) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardProfileListParamsStatus into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CardProfileListParamsStatus) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes CardProfileListParamsStatus into a url.Values of the query
// parameters associated with this value
func (r *CardProfileListParamsStatus) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return results whose value is in the provided list. For GET requests, this
// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
func (r CardProfileListParamsStatus) GetIn() (In []CardProfileListParamsStatusIn) {
	if r.In != nil {
		In = *r.In
	}
	return
}

func (r CardProfileListParamsStatus) String() (result string) {
	return fmt.Sprintf("&CardProfileListParamsStatus{In:%s}", core.Fmt(r.In))
}

type CardProfileListParamsStatusIn string

const (
	CardProfileListParamsStatusInPending  CardProfileListParamsStatusIn = "pending"
	CardProfileListParamsStatusInRejected CardProfileListParamsStatusIn = "rejected"
	CardProfileListParamsStatusInActive   CardProfileListParamsStatusIn = "active"
	CardProfileListParamsStatusInArchived CardProfileListParamsStatusIn = "archived"
)

type CardProfileList struct {
	// The contents of the list.
	Data *[]CardProfile `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CardProfileList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardProfileList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardProfileList into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CardProfileList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes CardProfileList into a url.Values of the query parameters
// associated with this value
func (r *CardProfileList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r CardProfileList) GetData() (Data []CardProfile) {
	if r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r CardProfileList) GetNextCursor() (NextCursor string) {
	if r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r CardProfileList) String() (result string) {
	return fmt.Sprintf("&CardProfileList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type CardProfilesPage struct {
	*pagination.Page[CardProfile]
}

func (r *CardProfilesPage) CardProfile() *CardProfile {
	return r.Current()
}

func (r *CardProfilesPage) NextPage() (*CardProfilesPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &CardProfilesPage{page}, nil
	}
}
