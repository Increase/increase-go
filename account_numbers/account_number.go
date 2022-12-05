package account_numbers

type AccountNumber struct {
	// The identifier for the account this Account Number belongs to.
	AccountId *string `json:"account_id,omitempty"`

	// The account number.
	AccountNumber *string `json:"account_number,omitempty"`

	// The Account Number identifier.
	Id *string `json:"id,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
	// Number was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The name you choose for the Account Number.
	Name *string `json:"name,omitempty"`

	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber *string `json:"routing_number,omitempty"`

	// This indicates if payments can be made to the Account Number.
	Status *string `json:"status,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `account_number`.
	Type *string `json:"type,omitempty"`
}

// The identifier for the account this Account Number belongs to.
func (r *AccountNumber) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The account number.
func (r *AccountNumber) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The Account Number identifier.
func (r *AccountNumber) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
// Number was created.
func (r *AccountNumber) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The name you choose for the Account Number.
func (r *AccountNumber) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r *AccountNumber) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// This indicates if payments can be made to the Account Number.
func (r *AccountNumber) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `account_number`.
func (r *AccountNumber) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
