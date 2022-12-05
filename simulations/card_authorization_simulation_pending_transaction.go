package simulations

type CardAuthorizationSimulationPendingTransaction struct {
	// The identifier for the account this Pending Transaction belongs to.
	AccountId *string `json:"account_id,omitempty"`

	// The Pending Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Pending
	// Transaction's currency. This will match the currency on the Pending
	// Transcation's Account.
	Currency *string `json:"currency,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the Pending
	// Transaction occured.
	CreatedAt *string `json:"created_at,omitempty"`

	// For a Pending Transaction related to a transfer, this is the description you
	// provide. For a Pending Transaction related to a payment, this is the description
	// the vendor provides.
	Description *string `json:"description,omitempty"`

	// The Pending Transaction identifier.
	Id *string `json:"id,omitempty"`

	// The identifier for the route this Pending Transaction came through. Routes are
	// things like cards and ACH details.
	RouteId *string `json:"route_id,omitempty"`

	// The type of the route this Pending Transaction came through.
	RouteType *string `json:"route_type,omitempty"`

	// This is an object giving more details on the network-level event that caused the
	// Pending Transaction. For example, for a card transaction this lists the
	// merchant's industry and location.
	Source *InboundCardAuthorizationSimulationResultPendingTransactionSource `json:"source,omitempty"`

	// Whether the Pending Transaction has been confirmed and has an associated
	// Transaction.
	Status *string `json:"status,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `pending_transaction`.
	Type *string `json:"type,omitempty"`
}

// The identifier for the account this Pending Transaction belongs to.
func (r *CardAuthorizationSimulationPendingTransaction) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The Pending Transaction amount in the minor unit of its currency. For dollars,
// for example, this is cents.
func (r *CardAuthorizationSimulationPendingTransaction) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Pending
// Transaction's currency. This will match the currency on the Pending
// Transcation's Account.
func (r *CardAuthorizationSimulationPendingTransaction) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the Pending
// Transaction occured.
func (r *CardAuthorizationSimulationPendingTransaction) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// For a Pending Transaction related to a transfer, this is the description you
// provide. For a Pending Transaction related to a payment, this is the description
// the vendor provides.
func (r *CardAuthorizationSimulationPendingTransaction) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Pending Transaction identifier.
func (r *CardAuthorizationSimulationPendingTransaction) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The identifier for the route this Pending Transaction came through. Routes are
// things like cards and ACH details.
func (r *CardAuthorizationSimulationPendingTransaction) GetRouteId() (RouteId string) {
	if r != nil && r.RouteId != nil {
		RouteId = *r.RouteId
	}
	return
}

// The type of the route this Pending Transaction came through.
func (r *CardAuthorizationSimulationPendingTransaction) GetRouteType() (RouteType string) {
	if r != nil && r.RouteType != nil {
		RouteType = *r.RouteType
	}
	return
}

// This is an object giving more details on the network-level event that caused the
// Pending Transaction. For example, for a card transaction this lists the
// merchant's industry and location.
func (r *CardAuthorizationSimulationPendingTransaction) GetSource() (Source InboundCardAuthorizationSimulationResultPendingTransactionSource) {
	if r != nil && r.Source != nil {
		Source = *r.Source
	}
	return
}

// Whether the Pending Transaction has been confirmed and has an associated
// Transaction.
func (r *CardAuthorizationSimulationPendingTransaction) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `pending_transaction`.
func (r *CardAuthorizationSimulationPendingTransaction) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
