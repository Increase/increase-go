package responses

import (
	"time"

	apijson "github.com/increase/increase-go/core/json"
)

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
	Raw            []byte
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
	Raw                   []byte
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
	Raw    []byte
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
	Raw        []byte
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardProfileListResponse using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardProfileListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
