package card_profiles

type CardProfilesDigitalWalletsRetrieveResponse struct {
	// The Card's text color, specified as an RGB triple.
	TextColor *CardProfilesDigitalWalletsTextColorRetrieveResponse `json:"text_color,omitempty"`

	// A user-facing description for whoever is issuing the card.
	IssuerName *string `json:"issuer_name,omitempty"`

	// A user-facing description for the card itself.
	CardDescription *string `json:"card_description,omitempty"`

	// A website the user can visit to view and receive support for their card.
	ContactWebsite *string `json:"contact_website,omitempty"`

	// An email address the user can contact to receive support for their card.
	ContactEmail *string `json:"contact_email,omitempty"`

	// A phone number the user can contact to receive support for their card.
	ContactPhone *string `json:"contact_phone,omitempty"`

	// The identifier of the File containing the card's front image.
	BackgroundImageFileId *string `json:"background_image_file_id,omitempty"`

	// The identifier of the File containing the card's icon image.
	AppIconFileId *string `json:"app_icon_file_id,omitempty"`
}

// The Card's text color, specified as an RGB triple.
func (r *CardProfilesDigitalWalletsRetrieveResponse) GetTextColor() (TextColor CardProfilesDigitalWalletsTextColorRetrieveResponse) {
	if r != nil && r.TextColor != nil {
		TextColor = *r.TextColor
	}
	return
}

// A user-facing description for whoever is issuing the card.
func (r *CardProfilesDigitalWalletsRetrieveResponse) GetIssuerName() (IssuerName string) {
	if r != nil && r.IssuerName != nil {
		IssuerName = *r.IssuerName
	}
	return
}

// A user-facing description for the card itself.
func (r *CardProfilesDigitalWalletsRetrieveResponse) GetCardDescription() (CardDescription string) {
	if r != nil && r.CardDescription != nil {
		CardDescription = *r.CardDescription
	}
	return
}

// A website the user can visit to view and receive support for their card.
func (r *CardProfilesDigitalWalletsRetrieveResponse) GetContactWebsite() (ContactWebsite string) {
	if r != nil && r.ContactWebsite != nil {
		ContactWebsite = *r.ContactWebsite
	}
	return
}

// An email address the user can contact to receive support for their card.
func (r *CardProfilesDigitalWalletsRetrieveResponse) GetContactEmail() (ContactEmail string) {
	if r != nil && r.ContactEmail != nil {
		ContactEmail = *r.ContactEmail
	}
	return
}

// A phone number the user can contact to receive support for their card.
func (r *CardProfilesDigitalWalletsRetrieveResponse) GetContactPhone() (ContactPhone string) {
	if r != nil && r.ContactPhone != nil {
		ContactPhone = *r.ContactPhone
	}
	return
}

// The identifier of the File containing the card's front image.
func (r *CardProfilesDigitalWalletsRetrieveResponse) GetBackgroundImageFileId() (BackgroundImageFileId string) {
	if r != nil && r.BackgroundImageFileId != nil {
		BackgroundImageFileId = *r.BackgroundImageFileId
	}
	return
}

// The identifier of the File containing the card's icon image.
func (r *CardProfilesDigitalWalletsRetrieveResponse) GetAppIconFileId() (AppIconFileId string) {
	if r != nil && r.AppIconFileId != nil {
		AppIconFileId = *r.AppIconFileId
	}
	return
}
