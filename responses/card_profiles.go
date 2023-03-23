package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/pagination"
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
	ID             pjson.Metadata
	CreatedAt      pjson.Metadata
	Status         pjson.Metadata
	Description    pjson.Metadata
	DigitalWallets pjson.Metadata
	Type           pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardProfile using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardProfile) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	TextColor             pjson.Metadata
	IssuerName            pjson.Metadata
	CardDescription       pjson.Metadata
	ContactWebsite        pjson.Metadata
	ContactEmail          pjson.Metadata
	ContactPhone          pjson.Metadata
	BackgroundImageFileID pjson.Metadata
	AppIconFileID         pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardProfileDigitalWallets
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardProfileDigitalWallets) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Red   pjson.Metadata
	Green pjson.Metadata
	Blue  pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardProfileDigitalWalletsTextColor using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *CardProfileDigitalWalletsTextColor) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CardProfileType string

const (
	CardProfileTypeCardProfile CardProfileType = "card_profile"
)

type CardProfileList struct {
	// The contents of the list.
	Data []CardProfile `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       CardProfileListJSON
}

type CardProfileListJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardProfileList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardProfileList) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
