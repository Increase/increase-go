package entities

type EntitiesNaturalPersonIdentificationPassportCreateParams struct {
	// The identifier of the File containing the passport.
	FileId *string `json:"file_id,omitempty"`

	// The passport's expiration date in YYYY-MM-DD format.
	ExpirationDate *string `json:"expiration_date,omitempty"`

	// The country that issued the passport.
	Country *string `json:"country,omitempty"`
}

// The identifier of the File containing the passport.
func (r *EntitiesNaturalPersonIdentificationPassportCreateParams) GetFileId() (FileId string) {
	if r != nil && r.FileId != nil {
		FileId = *r.FileId
	}
	return
}

// The passport's expiration date in YYYY-MM-DD format.
func (r *EntitiesNaturalPersonIdentificationPassportCreateParams) GetExpirationDate() (ExpirationDate string) {
	if r != nil && r.ExpirationDate != nil {
		ExpirationDate = *r.ExpirationDate
	}
	return
}

// The country that issued the passport.
func (r *EntitiesNaturalPersonIdentificationPassportCreateParams) GetCountry() (Country string) {
	if r != nil && r.Country != nil {
		Country = *r.Country
	}
	return
}
