package simulations

type WireTransferSimulation struct {
	// If the Wire Transfer attempt succeeds, this will contain the resulting
	// [Transaction](#transactions) object. The Transaction's `source` will be of
	// `category: inbound_wire_transfer`.
	Transaction *WireTransferSimulationTransaction `json:"transaction,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `inbound_wire_transfer_simulation_result`.
	Type *string `json:"type,omitempty"`
}

// If the Wire Transfer attempt succeeds, this will contain the resulting
// [Transaction](#transactions) object. The Transaction's `source` will be of
// `category: inbound_wire_transfer`.
func (r *WireTransferSimulation) GetTransaction() (Transaction WireTransferSimulationTransaction) {
	if r != nil && r.Transaction != nil {
		Transaction = *r.Transaction
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `inbound_wire_transfer_simulation_result`.
func (r *WireTransferSimulation) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
