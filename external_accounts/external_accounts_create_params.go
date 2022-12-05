package external_accounts

type ExternalAccountsCreateParams struct {
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber *string `json:"routing_number,omitempty"`

	// The account number for the destination account.
	AccountNumber *string `json:"account_number,omitempty"`

	// The type of the destination account. Defaults to `checking`.
	Funding *string `json:"funding,omitempty"`

	// The name you choose for the Account.
	Description *string `json:"description,omitempty"`
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
// destination account.
func (r *ExternalAccountsCreateParams) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The account number for the destination account.
func (r *ExternalAccountsCreateParams) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The type of the destination account. Defaults to `checking`.
func (r *ExternalAccountsCreateParams) GetFunding() (Funding string) {
	if r != nil && r.Funding != nil {
		Funding = *r.Funding
	}
	return
}

// The name you choose for the Account.
func (r *ExternalAccountsCreateParams) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}
