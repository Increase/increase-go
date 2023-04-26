package requests

import (
	"net/url"

	"github.com/increase/increase-go/core/field"
	apijson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
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
	return query.Marshal(r)
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
	return query.Marshal(r)
}

type CardProfileListParamsStatusIn string

const (
	CardProfileListParamsStatusInPending  CardProfileListParamsStatusIn = "pending"
	CardProfileListParamsStatusInRejected CardProfileListParamsStatusIn = "rejected"
	CardProfileListParamsStatusInActive   CardProfileListParamsStatusIn = "active"
	CardProfileListParamsStatusInArchived CardProfileListParamsStatusIn = "archived"
)
