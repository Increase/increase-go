package accounts

type Account struct {
	// The current balance of the Account in the minor unit of the currency. For
	// dollars, for example, this is cents.
	Balance *int64 `json:"balance,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
	// was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Account
	// currency.
	Currency *string `json:"currency,omitempty"`

	// The identifier for the Entity the Account belongs to.
	EntityId *string `json:"entity_id,omitempty"`

	// The identifier of an Entity that, while not owning the Account, is associated
	// with its activity.
	InformationalEntityId *string `json:"informational_entity_id,omitempty"`

	// The Account identifier.
	Id *string `json:"id,omitempty"`

	// The interest accrued but not yet paid, expressed as a string containing a
	// floating-point value.
	InterestAccrued *string `json:"interest_accrued,omitempty"`

	// The latest [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which
	// interest was accrued.
	InterestAccruedAt *string `json:"interest_accrued_at,omitempty"`

	// The name you choose for the Account.
	Name *string `json:"name,omitempty"`

	// The status of the Account.
	Status *string `json:"status,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `account`.
	Type *string `json:"type,omitempty"`
}

// The current balance of the Account in the minor unit of the currency. For
// dollars, for example, this is cents.
func (r *Account) GetBalance() (Balance int64) {
	if r != nil && r.Balance != nil {
		Balance = *r.Balance
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
// was created.
func (r *Account) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Account
// currency.
func (r *Account) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier for the Entity the Account belongs to.
func (r *Account) GetEntityId() (EntityId string) {
	if r != nil && r.EntityId != nil {
		EntityId = *r.EntityId
	}
	return
}

// The identifier of an Entity that, while not owning the Account, is associated
// with its activity.
func (r *Account) GetInformationalEntityId() (InformationalEntityId string) {
	if r != nil && r.InformationalEntityId != nil {
		InformationalEntityId = *r.InformationalEntityId
	}
	return
}

// The Account identifier.
func (r *Account) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The interest accrued but not yet paid, expressed as a string containing a
// floating-point value.
func (r *Account) GetInterestAccrued() (InterestAccrued string) {
	if r != nil && r.InterestAccrued != nil {
		InterestAccrued = *r.InterestAccrued
	}
	return
}

// The latest [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which
// interest was accrued.
func (r *Account) GetInterestAccruedAt() (InterestAccruedAt string) {
	if r != nil && r.InterestAccruedAt != nil {
		InterestAccruedAt = *r.InterestAccruedAt
	}
	return
}

// The name you choose for the Account.
func (r *Account) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The status of the Account.
func (r *Account) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `account`.
func (r *Account) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
