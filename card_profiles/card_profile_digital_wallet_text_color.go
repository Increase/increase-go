package card_profiles

type CardProfileDigitalWalletTextColor struct {
	// The value of the red channel in the RGB color.
	Red *int64 `json:"red,omitempty"`

	// The value of the green channel in the RGB color.
	Green *int64 `json:"green,omitempty"`

	// The value of the blue channel in the RGB color.
	Blue *int64 `json:"blue,omitempty"`
}

// The value of the red channel in the RGB color.
func (r *CardProfileDigitalWalletTextColor) GetRed() (Red int64) {
	if r != nil && r.Red != nil {
		Red = *r.Red
	}
	return
}

// The value of the green channel in the RGB color.
func (r *CardProfileDigitalWalletTextColor) GetGreen() (Green int64) {
	if r != nil && r.Green != nil {
		Green = *r.Green
	}
	return
}

// The value of the blue channel in the RGB color.
func (r *CardProfileDigitalWalletTextColor) GetBlue() (Blue int64) {
	if r != nil && r.Blue != nil {
		Blue = *r.Blue
	}
	return
}
