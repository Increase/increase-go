package simulations

type ACHTransferSimulationTransaction struct {
	// The identifier for the Account the Transaction belongs to.
	AccountId *string `json:"account_id,omitempty"`

	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transcation's
	// Account.
	Currency *string `json:"currency,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt *string `json:"created_at,omitempty"`

	// For a Transaction related to a transfer, this is the description you provide.
	// For a Transaction related to a payment, this is the description the vendor
	// provides.
	Description *string `json:"description,omitempty"`

	// The Transaction identifier.
	Id *string `json:"id,omitempty"`

	// The identifier for the route this Transaction came through. Routes are things
	// like cards and ACH details.
	RouteId *string `json:"route_id,omitempty"`

	// The type of the route this Transaction came through.
	RouteType *string `json:"route_type,omitempty"`

	// This is an object giving more details on the network-level event that caused the
	// Transaction. Note that for backwards compatibility reasons, additional
	// undocumented keys may appear in this object. These should be treated as
	// deprecated and will be removed in the future.
	Source *InboundACHTransferSimulationResultTransactionSource `json:"source,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type *string `json:"type,omitempty"`
}

// The identifier for the Account the Transaction belongs to.
func (r *ACHTransferSimulationTransaction) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The Transaction amount in the minor unit of its currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransaction) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// Transaction's currency. This will match the currency on the Transcation's
// Account.
func (r *ACHTransferSimulationTransaction) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
// Transaction occured.
func (r *ACHTransferSimulationTransaction) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// For a Transaction related to a transfer, this is the description you provide.
// For a Transaction related to a payment, this is the description the vendor
// provides.
func (r *ACHTransferSimulationTransaction) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Transaction identifier.
func (r *ACHTransferSimulationTransaction) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The identifier for the route this Transaction came through. Routes are things
// like cards and ACH details.
func (r *ACHTransferSimulationTransaction) GetRouteId() (RouteId string) {
	if r != nil && r.RouteId != nil {
		RouteId = *r.RouteId
	}
	return
}

// The type of the route this Transaction came through.
func (r *ACHTransferSimulationTransaction) GetRouteType() (RouteType string) {
	if r != nil && r.RouteType != nil {
		RouteType = *r.RouteType
	}
	return
}

// This is an object giving more details on the network-level event that caused the
// Transaction. Note that for backwards compatibility reasons, additional
// undocumented keys may appear in this object. These should be treated as
// deprecated and will be removed in the future.
func (r *ACHTransferSimulationTransaction) GetSource() (Source InboundACHTransferSimulationResultTransactionSource) {
	if r != nil && r.Source != nil {
		Source = *r.Source
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `transaction`.
func (r *ACHTransferSimulationTransaction) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
