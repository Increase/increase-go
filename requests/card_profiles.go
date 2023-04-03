package requests

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/fields"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
)

type CreateACardProfileParameters struct {
	// A description you can use to identify the Card Profile.
	Description fields.Field[string] `json:"description,required"`
	// How Cards should appear in digital wallets such as Apple Pay. Different wallets
	// will use these values to render card artwork appropriately for their app.
	DigitalWallets fields.Field[CreateACardProfileParametersDigitalWallets] `json:"digital_wallets,required"`
}

// MarshalJSON serializes CreateACardProfileParameters into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CreateACardProfileParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateACardProfileParameters) String() (result string) {
	return fmt.Sprintf("&CreateACardProfileParameters{Description:%s DigitalWallets:%s}", r.Description, r.DigitalWallets)
}

type CreateACardProfileParametersDigitalWallets struct {
	// The Card's text color, specified as an RGB triple. The default is white.
	TextColor fields.Field[CreateACardProfileParametersDigitalWalletsTextColor] `json:"text_color"`
	// A user-facing description for whoever is issuing the card.
	IssuerName fields.Field[string] `json:"issuer_name,required"`
	// A user-facing description for the card itself.
	CardDescription fields.Field[string] `json:"card_description,required"`
	// A website the user can visit to view and receive support for their card.
	ContactWebsite fields.Field[string] `json:"contact_website"`
	// An email address the user can contact to receive support for their card.
	ContactEmail fields.Field[string] `json:"contact_email"`
	// A phone number the user can contact to receive support for their card.
	ContactPhone fields.Field[string] `json:"contact_phone"`
	// The identifier of the File containing the card's front image.
	BackgroundImageFileID fields.Field[string] `json:"background_image_file_id,required"`
	// The identifier of the File containing the card's icon image.
	AppIconFileID fields.Field[string] `json:"app_icon_file_id,required"`
}

func (r CreateACardProfileParametersDigitalWallets) String() (result string) {
	return fmt.Sprintf("&CreateACardProfileParametersDigitalWallets{TextColor:%s IssuerName:%s CardDescription:%s ContactWebsite:%s ContactEmail:%s ContactPhone:%s BackgroundImageFileID:%s AppIconFileID:%s}", r.TextColor, r.IssuerName, r.CardDescription, r.ContactWebsite, r.ContactEmail, r.ContactPhone, r.BackgroundImageFileID, r.AppIconFileID)
}

type CreateACardProfileParametersDigitalWalletsTextColor struct {
	// The value of the red channel in the RGB color.
	Red fields.Field[int64] `json:"red,required"`
	// The value of the green channel in the RGB color.
	Green fields.Field[int64] `json:"green,required"`
	// The value of the blue channel in the RGB color.
	Blue fields.Field[int64] `json:"blue,required"`
}

func (r CreateACardProfileParametersDigitalWalletsTextColor) String() (result string) {
	return fmt.Sprintf("&CreateACardProfileParametersDigitalWalletsTextColor{Red:%s Green:%s Blue:%s}", r.Red, r.Green, r.Blue)
}

type CardProfileListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  fields.Field[int64]                       `query:"limit"`
	Status fields.Field[CardProfileListParamsStatus] `query:"status"`
}

// URLQuery serializes CardProfileListParams into a url.Values of the query
// parameters associated with this value
func (r *CardProfileListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r CardProfileListParams) String() (result string) {
	return fmt.Sprintf("&CardProfileListParams{Cursor:%s Limit:%s Status:%s}", r.Cursor, r.Limit, r.Status)
}

type CardProfileListParamsStatus struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In fields.Field[[]CardProfileListParamsStatusIn] `query:"in"`
}

// URLQuery serializes CardProfileListParamsStatus into a url.Values of the query
// parameters associated with this value
func (r *CardProfileListParamsStatus) URLQuery() (v url.Values) {
	return query.Marshal(r)
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
