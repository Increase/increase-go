package routing_numbers

type RoutingNumber struct {
	// The name of the financial institution belonging to a routing number.
	Name *string `json:"name,omitempty"`

	// The nine digit routing number identifier.
	RoutingNumber *string `json:"routing_number,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `routing_number`.
	Type *string `json:"type,omitempty"`

	// This routing number's support for ACH Transfers.
	ACHTransfers *string `json:"ach_transfers,omitempty"`

	// This routing number's support for Real Time Payments Transfers.
	RealTimePaymentsTransfers *string `json:"real_time_payments_transfers,omitempty"`

	// This routing number's support for Wire Transfers.
	WireTransfers *string `json:"wire_transfers,omitempty"`
}

// The name of the financial institution belonging to a routing number.
func (r *RoutingNumber) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The nine digit routing number identifier.
func (r *RoutingNumber) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `routing_number`.
func (r *RoutingNumber) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

// This routing number's support for ACH Transfers.
func (r *RoutingNumber) GetACHTransfers() (ACHTransfers string) {
	if r != nil && r.ACHTransfers != nil {
		ACHTransfers = *r.ACHTransfers
	}
	return
}

// This routing number's support for Real Time Payments Transfers.
func (r *RoutingNumber) GetRealTimePaymentsTransfers() (RealTimePaymentsTransfers string) {
	if r != nil && r.RealTimePaymentsTransfers != nil {
		RealTimePaymentsTransfers = *r.RealTimePaymentsTransfers
	}
	return
}

// This routing number's support for Wire Transfers.
func (r *RoutingNumber) GetWireTransfers() (WireTransfers string) {
	if r != nil && r.WireTransfers != nil {
		WireTransfers = *r.WireTransfers
	}
	return
}
