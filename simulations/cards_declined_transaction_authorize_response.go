package simulations

type CardsDeclinedTransactionAuthorizeResponse struct {
	// The identifier for the Account the Declined Transaction belongs to.
	AccountId *string `json:"account_id,omitempty"`

	// The Declined Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transcation's Account.
	Currency *string `json:"currency,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt *string `json:"created_at,omitempty"`

	// This is the description the vendor provides.
	Description *string `json:"description,omitempty"`

	// The Declined Transaction identifier.
	Id *string `json:"id,omitempty"`

	// The identifier for the route this Declined Transaction came through. Routes are
	// things like cards and ACH details.
	RouteId *string `json:"route_id,omitempty"`

	// The type of the route this Declined Transaction came through.
	RouteType *string `json:"route_type,omitempty"`

	// This is an object giving more details on the network-level event that caused the
	// Declined Transaction. For example, for a card transaction this lists the
	// merchant's industry and location. Note that for backwards compatibility reasons,
	// additional undocumented keys may appear in this object. These should be treated
	// as deprecated and will be removed in the future.
	Source *CardsDeclinedTransactionSourceAuthorizeResponse `json:"source,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type *string `json:"type,omitempty"`
}

// The identifier for the Account the Declined Transaction belongs to.
func (r *CardsDeclinedTransactionAuthorizeResponse) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The Declined Transaction amount in the minor unit of its currency. For dollars,
// for example, this is cents.
func (r *CardsDeclinedTransactionAuthorizeResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
// Transaction's currency. This will match the currency on the Declined
// Transcation's Account.
func (r *CardsDeclinedTransactionAuthorizeResponse) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
// Transaction occured.
func (r *CardsDeclinedTransactionAuthorizeResponse) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// This is the description the vendor provides.
func (r *CardsDeclinedTransactionAuthorizeResponse) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Declined Transaction identifier.
func (r *CardsDeclinedTransactionAuthorizeResponse) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The identifier for the route this Declined Transaction came through. Routes are
// things like cards and ACH details.
func (r *CardsDeclinedTransactionAuthorizeResponse) GetRouteId() (RouteId string) {
	if r != nil && r.RouteId != nil {
		RouteId = *r.RouteId
	}
	return
}

// The type of the route this Declined Transaction came through.
func (r *CardsDeclinedTransactionAuthorizeResponse) GetRouteType() (RouteType string) {
	if r != nil && r.RouteType != nil {
		RouteType = *r.RouteType
	}
	return
}

// This is an object giving more details on the network-level event that caused the
// Declined Transaction. For example, for a card transaction this lists the
// merchant's industry and location. Note that for backwards compatibility reasons,
// additional undocumented keys may appear in this object. These should be treated
// as deprecated and will be removed in the future.
func (r *CardsDeclinedTransactionAuthorizeResponse) GetSource() (Source CardsDeclinedTransactionSourceAuthorizeResponse) {
	if r != nil && r.Source != nil {
		Source = *r.Source
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `declined_transaction`.
func (r *CardsDeclinedTransactionAuthorizeResponse) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
