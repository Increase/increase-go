package cards

type CardsDataBillingAddressListResponse struct {
	// The first line of the billing address.
	Line1 *string `json:"line1,omitempty"`

	// The second line of the billing address.
	Line2 *string `json:"line2,omitempty"`

	// The city of the billing address.
	City *string `json:"city,omitempty"`

	// The US state of the billing address.
	State *string `json:"state,omitempty"`

	// The postal code of the billing address.
	PostalCode *string `json:"postal_code,omitempty"`
}

// The first line of the billing address.
func (r *CardsDataBillingAddressListResponse) GetLine1() (Line1 string) {
	if r != nil && r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the billing address.
func (r *CardsDataBillingAddressListResponse) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the billing address.
func (r *CardsDataBillingAddressListResponse) GetCity() (City string) {
	if r != nil && r.City != nil {
		City = *r.City
	}
	return
}

// The US state of the billing address.
func (r *CardsDataBillingAddressListResponse) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

// The postal code of the billing address.
func (r *CardsDataBillingAddressListResponse) GetPostalCode() (PostalCode string) {
	if r != nil && r.PostalCode != nil {
		PostalCode = *r.PostalCode
	}
	return
}
