package declined_transactions

type DeclinedTransactionsSourceRetrieveResponse struct {
	// The type of decline that took place. We may add additional possible values for
	// this enum over time; your application should be able to handle such additions
	// gracefully.
	Category *string `json:"category,omitempty"`

	// A ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline *DeclinedTransactionsSourceACHDeclineRetrieveResponse `json:"ach_decline,omitempty"`

	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline *DeclinedTransactionsSourceCardDeclineRetrieveResponse `json:"card_decline,omitempty"`

	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline *DeclinedTransactionsSourceCheckDeclineRetrieveResponse `json:"check_decline,omitempty"`

	// A Inbound Real Time Payments Transfer Decline object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline *DeclinedTransactionsSourceInboundRealTimePaymentsTransferDeclineRetrieveResponse `json:"inbound_real_time_payments_transfer_decline,omitempty"`

	// A International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline *DeclinedTransactionsSourceInternationalACHDeclineRetrieveResponse `json:"international_ach_decline,omitempty"`

	// A Deprecated Card Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_decline`.
	CardRouteDecline *DeclinedTransactionsSourceCardRouteDeclineRetrieveResponse `json:"card_route_decline,omitempty"`
}

// The type of decline that took place. We may add additional possible values for
// this enum over time; your application should be able to handle such additions
// gracefully.
func (r *DeclinedTransactionsSourceRetrieveResponse) GetCategory() (Category string) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// A ACH Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `ach_decline`.
func (r *DeclinedTransactionsSourceRetrieveResponse) GetACHDecline() (ACHDecline DeclinedTransactionsSourceACHDeclineRetrieveResponse) {
	if r != nil && r.ACHDecline != nil {
		ACHDecline = *r.ACHDecline
	}
	return
}

// A Card Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `card_decline`.
func (r *DeclinedTransactionsSourceRetrieveResponse) GetCardDecline() (CardDecline DeclinedTransactionsSourceCardDeclineRetrieveResponse) {
	if r != nil && r.CardDecline != nil {
		CardDecline = *r.CardDecline
	}
	return
}

// A Check Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `check_decline`.
func (r *DeclinedTransactionsSourceRetrieveResponse) GetCheckDecline() (CheckDecline DeclinedTransactionsSourceCheckDeclineRetrieveResponse) {
	if r != nil && r.CheckDecline != nil {
		CheckDecline = *r.CheckDecline
	}
	return
}

// A Inbound Real Time Payments Transfer Decline object. This field will be present
// in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_decline`.
func (r *DeclinedTransactionsSourceRetrieveResponse) GetInboundRealTimePaymentsTransferDecline() (InboundRealTimePaymentsTransferDecline DeclinedTransactionsSourceInboundRealTimePaymentsTransferDeclineRetrieveResponse) {
	if r != nil && r.InboundRealTimePaymentsTransferDecline != nil {
		InboundRealTimePaymentsTransferDecline = *r.InboundRealTimePaymentsTransferDecline
	}
	return
}

// A International ACH Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `international_ach_decline`.
func (r *DeclinedTransactionsSourceRetrieveResponse) GetInternationalACHDecline() (InternationalACHDecline DeclinedTransactionsSourceInternationalACHDeclineRetrieveResponse) {
	if r != nil && r.InternationalACHDecline != nil {
		InternationalACHDecline = *r.InternationalACHDecline
	}
	return
}

// A Deprecated Card Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_decline`.
func (r *DeclinedTransactionsSourceRetrieveResponse) GetCardRouteDecline() (CardRouteDecline DeclinedTransactionsSourceCardRouteDeclineRetrieveResponse) {
	if r != nil && r.CardRouteDecline != nil {
		CardRouteDecline = *r.CardRouteDecline
	}
	return
}
