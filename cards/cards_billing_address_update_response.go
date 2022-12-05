package cards

type CardsBillingAddressUpdateResponse struct {
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
func (r *CardsBillingAddressUpdateResponse) GetLine1() (Line1 string) {
	if r != nil && r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the billing address.
func (r *CardsBillingAddressUpdateResponse) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the billing address.
func (r *CardsBillingAddressUpdateResponse) GetCity() (City string) {
	if r != nil && r.City != nil {
		City = *r.City
	}
	return
}

// The US state of the billing address.
func (r *CardsBillingAddressUpdateResponse) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

// The postal code of the billing address.
func (r *CardsBillingAddressUpdateResponse) GetPostalCode() (PostalCode string) {
	if r != nil && r.PostalCode != nil {
		PostalCode = *r.PostalCode
	}
	return
}
