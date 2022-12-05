package external_accounts

type ExternalAccount struct {
	// The External Account's identifier.
	Id *string `json:"id,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the External Account was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The External Account's description for display purposes.
	Description *string `json:"description,omitempty"`

	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber *string `json:"routing_number,omitempty"`

	// The destination account number.
	AccountNumber *string `json:"account_number,omitempty"`

	// The type of the account to which the transfer will be sent.
	Funding *string `json:"funding,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `external_account`.
	Type *string `json:"type,omitempty"`
}

// The External Account's identifier.
func (r *ExternalAccount) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the External Account was created.
func (r *ExternalAccount) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The External Account's description for display purposes.
func (r *ExternalAccount) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r *ExternalAccount) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The destination account number.
func (r *ExternalAccount) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The type of the account to which the transfer will be sent.
func (r *ExternalAccount) GetFunding() (Funding string) {
	if r != nil && r.Funding != nil {
		Funding = *r.Funding
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `external_account`.
func (r *ExternalAccount) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
