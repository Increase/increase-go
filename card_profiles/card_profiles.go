package card_profiles

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type CardProfileService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewCardProfileService(requester core.Requester) (r *CardProfileService) {
	r = &CardProfileService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

type PreloadedCardProfileService struct {
	CardProfiles *CardProfileService
}

func (r *PreloadedCardProfileService) Init(service *CardProfileService) {
	r.CardProfiles = service
}

func NewPreloadedCardProfileService(service *CardProfileService) (r *PreloadedCardProfileService) {
	r = &PreloadedCardProfileService{}
	r.Init(service)
	return
}

//
type CardProfile struct {
	// The Card Profile identifier.
	ID *string `json:"id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was created.
	CreatedAt *string `json:"created_at"`
	// The status of the Card Profile.
	Status *CardProfileStatus `json:"status"`
	// A description you can use to identify the Card Profile.
	Description *string `json:"description"`
	// How Cards should appear in digital wallets such as Apple Pay. Different wallets
	// will use these values to render card artwork appropriately for their app.
	DigitalWallets *CardProfileDigitalWallets `json:"digital_wallets"`
	// A constant representing the object's type. For this resource it will always be
	// `card_profile`.
	Type *CardProfileType `json:"type"`
}

// The Card Profile identifier.
func (r *CardProfile) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card Dispute was created.
func (r *CardProfile) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The status of the Card Profile.
func (r *CardProfile) GetStatus() (Status CardProfileStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// A description you can use to identify the Card Profile.
func (r *CardProfile) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// How Cards should appear in digital wallets such as Apple Pay. Different wallets
// will use these values to render card artwork appropriately for their app.
func (r *CardProfile) GetDigitalWallets() (DigitalWallets CardProfileDigitalWallets) {
	if r != nil && r.DigitalWallets != nil {
		DigitalWallets = *r.DigitalWallets
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_profile`.
func (r *CardProfile) GetType() (Type CardProfileType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type CardProfileStatus string

const (
	CardProfileStatusPending  CardProfileStatus = "pending"
	CardProfileStatusRejected CardProfileStatus = "rejected"
	CardProfileStatusActive   CardProfileStatus = "active"
	CardProfileStatusArchived CardProfileStatus = "archived"
)

//
type CardProfileDigitalWallets struct {
	// The Card's text color, specified as an RGB triple.
	TextColor *CardProfileDigitalWalletsTextColor `json:"text_color"`
	// A user-facing description for whoever is issuing the card.
	IssuerName *string `json:"issuer_name"`
	// A user-facing description for the card itself.
	CardDescription *string `json:"card_description"`
	// A website the user can visit to view and receive support for their card.
	ContactWebsite *string `json:"contact_website"`
	// An email address the user can contact to receive support for their card.
	ContactEmail *string `json:"contact_email"`
	// A phone number the user can contact to receive support for their card.
	ContactPhone *string `json:"contact_phone"`
	// The identifier of the File containing the card's front image.
	BackgroundImageFileID *string `json:"background_image_file_id"`
	// The identifier of the File containing the card's icon image.
	AppIconFileID *string `json:"app_icon_file_id"`
}

// The Card's text color, specified as an RGB triple.
func (r *CardProfileDigitalWallets) GetTextColor() (TextColor CardProfileDigitalWalletsTextColor) {
	if r != nil && r.TextColor != nil {
		TextColor = *r.TextColor
	}
	return
}

// A user-facing description for whoever is issuing the card.
func (r *CardProfileDigitalWallets) GetIssuerName() (IssuerName string) {
	if r != nil && r.IssuerName != nil {
		IssuerName = *r.IssuerName
	}
	return
}

// A user-facing description for the card itself.
func (r *CardProfileDigitalWallets) GetCardDescription() (CardDescription string) {
	if r != nil && r.CardDescription != nil {
		CardDescription = *r.CardDescription
	}
	return
}

// A website the user can visit to view and receive support for their card.
func (r *CardProfileDigitalWallets) GetContactWebsite() (ContactWebsite string) {
	if r != nil && r.ContactWebsite != nil {
		ContactWebsite = *r.ContactWebsite
	}
	return
}

// An email address the user can contact to receive support for their card.
func (r *CardProfileDigitalWallets) GetContactEmail() (ContactEmail string) {
	if r != nil && r.ContactEmail != nil {
		ContactEmail = *r.ContactEmail
	}
	return
}

// A phone number the user can contact to receive support for their card.
func (r *CardProfileDigitalWallets) GetContactPhone() (ContactPhone string) {
	if r != nil && r.ContactPhone != nil {
		ContactPhone = *r.ContactPhone
	}
	return
}

// The identifier of the File containing the card's front image.
func (r *CardProfileDigitalWallets) GetBackgroundImageFileID() (BackgroundImageFileID string) {
	if r != nil && r.BackgroundImageFileID != nil {
		BackgroundImageFileID = *r.BackgroundImageFileID
	}
	return
}

// The identifier of the File containing the card's icon image.
func (r *CardProfileDigitalWallets) GetAppIconFileID() (AppIconFileID string) {
	if r != nil && r.AppIconFileID != nil {
		AppIconFileID = *r.AppIconFileID
	}
	return
}

//
type CardProfileDigitalWalletsTextColor struct {
	// The value of the red channel in the RGB color.
	Red *int `json:"red"`
	// The value of the green channel in the RGB color.
	Green *int `json:"green"`
	// The value of the blue channel in the RGB color.
	Blue *int `json:"blue"`
}

// The value of the red channel in the RGB color.
func (r *CardProfileDigitalWalletsTextColor) GetRed() (Red int) {
	if r != nil && r.Red != nil {
		Red = *r.Red
	}
	return
}

// The value of the green channel in the RGB color.
func (r *CardProfileDigitalWalletsTextColor) GetGreen() (Green int) {
	if r != nil && r.Green != nil {
		Green = *r.Green
	}
	return
}

// The value of the blue channel in the RGB color.
func (r *CardProfileDigitalWalletsTextColor) GetBlue() (Blue int) {
	if r != nil && r.Blue != nil {
		Blue = *r.Blue
	}
	return
}

type CardProfileType string

const (
	CardProfileTypeCardProfile CardProfileType = "card_profile"
)

type CreateACardProfileParameters struct {
	// A description you can use to identify the Card Profile.
	Description *string `json:"description"`
	// How Cards should appear in digital wallets such as Apple Pay. Different wallets
	// will use these values to render card artwork appropriately for their app.
	DigitalWallets *CreateACardProfileParametersDigitalWallets `json:"digital_wallets"`
}

// A description you can use to identify the Card Profile.
func (r *CreateACardProfileParameters) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// How Cards should appear in digital wallets such as Apple Pay. Different wallets
// will use these values to render card artwork appropriately for their app.
func (r *CreateACardProfileParameters) GetDigitalWallets() (DigitalWallets CreateACardProfileParametersDigitalWallets) {
	if r != nil && r.DigitalWallets != nil {
		DigitalWallets = *r.DigitalWallets
	}
	return
}

//
type CreateACardProfileParametersDigitalWallets struct {
	// The Card's text color, specified as an RGB triple. The default is white.
	TextColor *CreateACardProfileParametersDigitalWalletsTextColor `json:"text_color,omitempty"`
	// A user-facing description for whoever is issuing the card.
	IssuerName *string `json:"issuer_name"`
	// A user-facing description for the card itself.
	CardDescription *string `json:"card_description"`
	// A website the user can visit to view and receive support for their card.
	ContactWebsite *string `json:"contact_website,omitempty"`
	// An email address the user can contact to receive support for their card.
	ContactEmail *string `json:"contact_email,omitempty"`
	// A phone number the user can contact to receive support for their card.
	ContactPhone *string `json:"contact_phone,omitempty"`
	// The identifier of the File containing the card's front image.
	BackgroundImageFileID *string `json:"background_image_file_id"`
	// The identifier of the File containing the card's icon image.
	AppIconFileID *string `json:"app_icon_file_id"`
}

// The Card's text color, specified as an RGB triple. The default is white.
func (r *CreateACardProfileParametersDigitalWallets) GetTextColor() (TextColor CreateACardProfileParametersDigitalWalletsTextColor) {
	if r != nil && r.TextColor != nil {
		TextColor = *r.TextColor
	}
	return
}

// A user-facing description for whoever is issuing the card.
func (r *CreateACardProfileParametersDigitalWallets) GetIssuerName() (IssuerName string) {
	if r != nil && r.IssuerName != nil {
		IssuerName = *r.IssuerName
	}
	return
}

// A user-facing description for the card itself.
func (r *CreateACardProfileParametersDigitalWallets) GetCardDescription() (CardDescription string) {
	if r != nil && r.CardDescription != nil {
		CardDescription = *r.CardDescription
	}
	return
}

// A website the user can visit to view and receive support for their card.
func (r *CreateACardProfileParametersDigitalWallets) GetContactWebsite() (ContactWebsite string) {
	if r != nil && r.ContactWebsite != nil {
		ContactWebsite = *r.ContactWebsite
	}
	return
}

// An email address the user can contact to receive support for their card.
func (r *CreateACardProfileParametersDigitalWallets) GetContactEmail() (ContactEmail string) {
	if r != nil && r.ContactEmail != nil {
		ContactEmail = *r.ContactEmail
	}
	return
}

// A phone number the user can contact to receive support for their card.
func (r *CreateACardProfileParametersDigitalWallets) GetContactPhone() (ContactPhone string) {
	if r != nil && r.ContactPhone != nil {
		ContactPhone = *r.ContactPhone
	}
	return
}

// The identifier of the File containing the card's front image.
func (r *CreateACardProfileParametersDigitalWallets) GetBackgroundImageFileID() (BackgroundImageFileID string) {
	if r != nil && r.BackgroundImageFileID != nil {
		BackgroundImageFileID = *r.BackgroundImageFileID
	}
	return
}

// The identifier of the File containing the card's icon image.
func (r *CreateACardProfileParametersDigitalWallets) GetAppIconFileID() (AppIconFileID string) {
	if r != nil && r.AppIconFileID != nil {
		AppIconFileID = *r.AppIconFileID
	}
	return
}

//
type CreateACardProfileParametersDigitalWalletsTextColor struct {
	// The value of the red channel in the RGB color.
	Red *int `json:"red"`
	// The value of the green channel in the RGB color.
	Green *int `json:"green"`
	// The value of the blue channel in the RGB color.
	Blue *int `json:"blue"`
}

// The value of the red channel in the RGB color.
func (r *CreateACardProfileParametersDigitalWalletsTextColor) GetRed() (Red int) {
	if r != nil && r.Red != nil {
		Red = *r.Red
	}
	return
}

// The value of the green channel in the RGB color.
func (r *CreateACardProfileParametersDigitalWalletsTextColor) GetGreen() (Green int) {
	if r != nil && r.Green != nil {
		Green = *r.Green
	}
	return
}

// The value of the blue channel in the RGB color.
func (r *CreateACardProfileParametersDigitalWalletsTextColor) GetBlue() (Blue int) {
	if r != nil && r.Blue != nil {
		Blue = *r.Blue
	}
	return
}

type ListCardProfilesQuery struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  *int                         `query:"limit"`
	Status *ListCardProfilesQueryStatus `query:"status"`
}

// Return the page of entries after this one.
func (r *ListCardProfilesQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListCardProfilesQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

func (r *ListCardProfilesQuery) GetStatus() (Status ListCardProfilesQueryStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

type ListCardProfilesQueryStatus struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In *[]ListCardProfilesQueryStatusIn `json:"in,omitempty"`
}

// Return results whose value is in the provided list. For GET requests, this
// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
func (r *ListCardProfilesQueryStatus) GetIn() (In []ListCardProfilesQueryStatusIn) {
	if r != nil && r.In != nil {
		In = *r.In
	}
	return
}

type ListCardProfilesQueryStatusIn string

const (
	ListCardProfilesQueryStatusInPending  ListCardProfilesQueryStatusIn = "pending"
	ListCardProfilesQueryStatusInRejected ListCardProfilesQueryStatusIn = "rejected"
	ListCardProfilesQueryStatusInActive   ListCardProfilesQueryStatusIn = "active"
	ListCardProfilesQueryStatusInArchived ListCardProfilesQueryStatusIn = "archived"
)

//
type CardProfileList struct {
	// The contents of the list.
	Data *[]CardProfile `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// The contents of the list.
func (r *CardProfileList) GetData() (Data []CardProfile) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *CardProfileList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *CardProfileService) Create(ctx context.Context, body *CreateACardProfileParameters, opts ...*core.RequestOpts) (res *CardProfile, err error) {
	err = r.post(
		ctx,
		"/card_profiles",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *PreloadedCardProfileService) Create(ctx context.Context, body *CreateACardProfileParameters, opts ...*core.RequestOpts) (res *CardProfile, err error) {
	err = r.CardProfiles.post(
		ctx,
		"/card_profiles",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *CardProfileService) Retrieve(ctx context.Context, card_profile_id string, opts ...*core.RequestOpts) (res *CardProfile, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/card_profiles/%s", card_profile_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *PreloadedCardProfileService) Retrieve(ctx context.Context, card_profile_id string, opts ...*core.RequestOpts) (res *CardProfile, err error) {
	err = r.CardProfiles.get(
		ctx,
		fmt.Sprintf("/card_profiles/%s", card_profile_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

type CardProfilesPage struct {
	*pagination.Page[CardProfile]
}

func (r *CardProfilesPage) CardProfile() *CardProfile {
	return r.Current()
}

func (r *CardProfilesPage) GetNextPage() (*CardProfilesPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &CardProfilesPage{page}, nil
	}
}

func (r *CardProfileService) List(ctx context.Context, query *ListCardProfilesQuery, opts ...*core.RequestOpts) (res *CardProfilesPage, err error) {
	page := &CardProfilesPage{
		Page: &pagination.Page[CardProfile]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/card_profiles",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *PreloadedCardProfileService) List(ctx context.Context, query *ListCardProfilesQuery, opts ...*core.RequestOpts) (res *CardProfilesPage, err error) {
	page := &CardProfilesPage{
		Page: &pagination.Page[CardProfile]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/card_profiles",
			},
			Requester: r.CardProfiles.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}
