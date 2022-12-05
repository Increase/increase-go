package simulations

type InboundACHTransferSimulationResultDeclinedTransactionSource struct {
	// The type of decline that took place. We may add additional possible values for
	// this enum over time; your application should be able to handle such additions
	// gracefully.
	Category *string `json:"category,omitempty"`

	// A ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline *InboundACHTransferSimulationResultDeclinedTransactionSourceACHDecline `json:"ach_decline,omitempty"`

	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline *InboundACHTransferSimulationResultDeclinedTransactionSourceCardDecline `json:"card_decline,omitempty"`

	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline *InboundACHTransferSimulationResultDeclinedTransactionSourceCheckDecline `json:"check_decline,omitempty"`

	// A Inbound Real Time Payments Transfer Decline object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline *InboundACHTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline `json:"inbound_real_time_payments_transfer_decline,omitempty"`

	// A International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline *InboundACHTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline `json:"international_ach_decline,omitempty"`

	// A Deprecated Card Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_decline`.
	CardRouteDecline *InboundACHTransferSimulationResultDeclinedTransactionSourceCardRouteDecline `json:"card_route_decline,omitempty"`
}

// The type of decline that took place. We may add additional possible values for
// this enum over time; your application should be able to handle such additions
// gracefully.
func (r *InboundACHTransferSimulationResultDeclinedTransactionSource) GetCategory() (Category string) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// A ACH Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `ach_decline`.
func (r *InboundACHTransferSimulationResultDeclinedTransactionSource) GetACHDecline() (ACHDecline InboundACHTransferSimulationResultDeclinedTransactionSourceACHDecline) {
	if r != nil && r.ACHDecline != nil {
		ACHDecline = *r.ACHDecline
	}
	return
}

// A Card Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `card_decline`.
func (r *InboundACHTransferSimulationResultDeclinedTransactionSource) GetCardDecline() (CardDecline InboundACHTransferSimulationResultDeclinedTransactionSourceCardDecline) {
	if r != nil && r.CardDecline != nil {
		CardDecline = *r.CardDecline
	}
	return
}

// A Check Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `check_decline`.
func (r *InboundACHTransferSimulationResultDeclinedTransactionSource) GetCheckDecline() (CheckDecline InboundACHTransferSimulationResultDeclinedTransactionSourceCheckDecline) {
	if r != nil && r.CheckDecline != nil {
		CheckDecline = *r.CheckDecline
	}
	return
}

// A Inbound Real Time Payments Transfer Decline object. This field will be present
// in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_decline`.
func (r *InboundACHTransferSimulationResultDeclinedTransactionSource) GetInboundRealTimePaymentsTransferDecline() (InboundRealTimePaymentsTransferDecline InboundACHTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) {
	if r != nil && r.InboundRealTimePaymentsTransferDecline != nil {
		InboundRealTimePaymentsTransferDecline = *r.InboundRealTimePaymentsTransferDecline
	}
	return
}

// A International ACH Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `international_ach_decline`.
func (r *InboundACHTransferSimulationResultDeclinedTransactionSource) GetInternationalACHDecline() (InternationalACHDecline InboundACHTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) {
	if r != nil && r.InternationalACHDecline != nil {
		InternationalACHDecline = *r.InternationalACHDecline
	}
	return
}

// A Deprecated Card Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_decline`.
func (r *InboundACHTransferSimulationResultDeclinedTransactionSource) GetCardRouteDecline() (CardRouteDecline InboundACHTransferSimulationResultDeclinedTransactionSourceCardRouteDecline) {
	if r != nil && r.CardRouteDecline != nil {
		CardRouteDecline = *r.CardRouteDecline
	}
	return
}
