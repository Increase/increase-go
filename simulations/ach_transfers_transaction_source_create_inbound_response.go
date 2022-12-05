package simulations

type ACHTransfersTransactionSourceCreateInboundResponse struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category *string `json:"category,omitempty"`

	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention *ACHTransfersTransactionSourceAccountTransferIntentionCreateInboundResponse `json:"account_transfer_intention,omitempty"`

	// A ACH Check Conversion Return object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_check_conversion_return`.
	ACHCheckConversionReturn *ACHTransfersTransactionSourceACHCheckConversionReturnCreateInboundResponse `json:"ach_check_conversion_return,omitempty"`

	// A ACH Check Conversion object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_check_conversion`.
	ACHCheckConversion *ACHTransfersTransactionSourceACHCheckConversionCreateInboundResponse `json:"ach_check_conversion,omitempty"`

	// A ACH Transfer Intention object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention *ACHTransfersTransactionSourceACHTransferIntentionCreateInboundResponse `json:"ach_transfer_intention,omitempty"`

	// A ACH Transfer Rejection object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection *ACHTransfersTransactionSourceACHTransferRejectionCreateInboundResponse `json:"ach_transfer_rejection,omitempty"`

	// A ACH Transfer Return object. This field will be present in the JSON response if
	// and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn *ACHTransfersTransactionSourceACHTransferReturnCreateInboundResponse `json:"ach_transfer_return,omitempty"`

	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance *ACHTransfersTransactionSourceCardDisputeAcceptanceCreateInboundResponse `json:"card_dispute_acceptance,omitempty"`

	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund *ACHTransfersTransactionSourceCardRefundCreateInboundResponse `json:"card_refund,omitempty"`

	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement *ACHTransfersTransactionSourceCardSettlementCreateInboundResponse `json:"card_settlement,omitempty"`

	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance *ACHTransfersTransactionSourceCheckDepositAcceptanceCreateInboundResponse `json:"check_deposit_acceptance,omitempty"`

	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn *ACHTransfersTransactionSourceCheckDepositReturnCreateInboundResponse `json:"check_deposit_return,omitempty"`

	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention *ACHTransfersTransactionSourceCheckTransferIntentionCreateInboundResponse `json:"check_transfer_intention,omitempty"`

	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection *ACHTransfersTransactionSourceCheckTransferRejectionCreateInboundResponse `json:"check_transfer_rejection,omitempty"`

	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest *ACHTransfersTransactionSourceCheckTransferStopPaymentRequestCreateInboundResponse `json:"check_transfer_stop_payment_request,omitempty"`

	// A Dispute Resolution object. This field will be present in the JSON response if
	// and only if `category` is equal to `dispute_resolution`.
	DisputeResolution *ACHTransfersTransactionSourceDisputeResolutionCreateInboundResponse `json:"dispute_resolution,omitempty"`

	// A Empyreal Cash Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `empyreal_cash_deposit`.
	EmpyrealCashDeposit *ACHTransfersTransactionSourceEmpyrealCashDepositCreateInboundResponse `json:"empyreal_cash_deposit,omitempty"`

	// A Inbound ACH Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer *ACHTransfersTransactionSourceInboundACHTransferCreateInboundResponse `json:"inbound_ach_transfer,omitempty"`

	// A Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck *ACHTransfersTransactionSourceInboundCheckCreateInboundResponse `json:"inbound_check,omitempty"`

	// A Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer *ACHTransfersTransactionSourceInboundInternationalACHTransferCreateInboundResponse `json:"inbound_international_ach_transfer,omitempty"`

	// A Inbound Real Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation *ACHTransfersTransactionSourceInboundRealTimePaymentsTransferConfirmationCreateInboundResponse `json:"inbound_real_time_payments_transfer_confirmation,omitempty"`

	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal *ACHTransfersTransactionSourceInboundWireDrawdownPaymentReversalCreateInboundResponse `json:"inbound_wire_drawdown_payment_reversal,omitempty"`

	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment *ACHTransfersTransactionSourceInboundWireDrawdownPaymentCreateInboundResponse `json:"inbound_wire_drawdown_payment,omitempty"`

	// A Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal *ACHTransfersTransactionSourceInboundWireReversalCreateInboundResponse `json:"inbound_wire_reversal,omitempty"`

	// A Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer *ACHTransfersTransactionSourceInboundWireTransferCreateInboundResponse `json:"inbound_wire_transfer,omitempty"`

	// A Internal Source object. This field will be present in the JSON response if and
	// only if `category` is equal to `internal_source`.
	InternalSource *ACHTransfersTransactionSourceInternalSourceCreateInboundResponse `json:"internal_source,omitempty"`

	// A Deprecated Card Refund object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_route_refund`.
	CardRouteRefund *ACHTransfersTransactionSourceCardRouteRefundCreateInboundResponse `json:"card_route_refund,omitempty"`

	// A Deprecated Card Settlement object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_settlement`.
	CardRouteSettlement *ACHTransfersTransactionSourceCardRouteSettlementCreateInboundResponse `json:"card_route_settlement,omitempty"`

	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds *ACHTransfersTransactionSourceSampleFundsCreateInboundResponse `json:"sample_funds,omitempty"`

	// A Wire Drawdown Payment Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_intention`.
	WireDrawdownPaymentIntention *ACHTransfersTransactionSourceWireDrawdownPaymentIntentionCreateInboundResponse `json:"wire_drawdown_payment_intention,omitempty"`

	// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_rejection`.
	WireDrawdownPaymentRejection *ACHTransfersTransactionSourceWireDrawdownPaymentRejectionCreateInboundResponse `json:"wire_drawdown_payment_rejection,omitempty"`

	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention *ACHTransfersTransactionSourceWireTransferIntentionCreateInboundResponse `json:"wire_transfer_intention,omitempty"`

	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection *ACHTransfersTransactionSourceWireTransferRejectionCreateInboundResponse `json:"wire_transfer_rejection,omitempty"`
}

// The type of transaction that took place. We may add additional possible values
// for this enum over time; your application should be able to handle such
// additions gracefully.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetCategory() (Category string) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// A Account Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_intention`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetAccountTransferIntention() (AccountTransferIntention ACHTransfersTransactionSourceAccountTransferIntentionCreateInboundResponse) {
	if r != nil && r.AccountTransferIntention != nil {
		AccountTransferIntention = *r.AccountTransferIntention
	}
	return
}

// A ACH Check Conversion Return object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_check_conversion_return`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetACHCheckConversionReturn() (ACHCheckConversionReturn ACHTransfersTransactionSourceACHCheckConversionReturnCreateInboundResponse) {
	if r != nil && r.ACHCheckConversionReturn != nil {
		ACHCheckConversionReturn = *r.ACHCheckConversionReturn
	}
	return
}

// A ACH Check Conversion object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_check_conversion`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetACHCheckConversion() (ACHCheckConversion ACHTransfersTransactionSourceACHCheckConversionCreateInboundResponse) {
	if r != nil && r.ACHCheckConversion != nil {
		ACHCheckConversion = *r.ACHCheckConversion
	}
	return
}

// A ACH Transfer Intention object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_intention`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetACHTransferIntention() (ACHTransferIntention ACHTransfersTransactionSourceACHTransferIntentionCreateInboundResponse) {
	if r != nil && r.ACHTransferIntention != nil {
		ACHTransferIntention = *r.ACHTransferIntention
	}
	return
}

// A ACH Transfer Rejection object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_rejection`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetACHTransferRejection() (ACHTransferRejection ACHTransfersTransactionSourceACHTransferRejectionCreateInboundResponse) {
	if r != nil && r.ACHTransferRejection != nil {
		ACHTransferRejection = *r.ACHTransferRejection
	}
	return
}

// A ACH Transfer Return object. This field will be present in the JSON response if
// and only if `category` is equal to `ach_transfer_return`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetACHTransferReturn() (ACHTransferReturn ACHTransfersTransactionSourceACHTransferReturnCreateInboundResponse) {
	if r != nil && r.ACHTransferReturn != nil {
		ACHTransferReturn = *r.ACHTransferReturn
	}
	return
}

// A Card Dispute Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `card_dispute_acceptance`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetCardDisputeAcceptance() (CardDisputeAcceptance ACHTransfersTransactionSourceCardDisputeAcceptanceCreateInboundResponse) {
	if r != nil && r.CardDisputeAcceptance != nil {
		CardDisputeAcceptance = *r.CardDisputeAcceptance
	}
	return
}

// A Card Refund object. This field will be present in the JSON response if and
// only if `category` is equal to `card_refund`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetCardRefund() (CardRefund ACHTransfersTransactionSourceCardRefundCreateInboundResponse) {
	if r != nil && r.CardRefund != nil {
		CardRefund = *r.CardRefund
	}
	return
}

// A Card Settlement object. This field will be present in the JSON response if and
// only if `category` is equal to `card_settlement`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetCardSettlement() (CardSettlement ACHTransfersTransactionSourceCardSettlementCreateInboundResponse) {
	if r != nil && r.CardSettlement != nil {
		CardSettlement = *r.CardSettlement
	}
	return
}

// A Check Deposit Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_acceptance`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetCheckDepositAcceptance() (CheckDepositAcceptance ACHTransfersTransactionSourceCheckDepositAcceptanceCreateInboundResponse) {
	if r != nil && r.CheckDepositAcceptance != nil {
		CheckDepositAcceptance = *r.CheckDepositAcceptance
	}
	return
}

// A Check Deposit Return object. This field will be present in the JSON response
// if and only if `category` is equal to `check_deposit_return`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetCheckDepositReturn() (CheckDepositReturn ACHTransfersTransactionSourceCheckDepositReturnCreateInboundResponse) {
	if r != nil && r.CheckDepositReturn != nil {
		CheckDepositReturn = *r.CheckDepositReturn
	}
	return
}

// A Check Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_intention`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetCheckTransferIntention() (CheckTransferIntention ACHTransfersTransactionSourceCheckTransferIntentionCreateInboundResponse) {
	if r != nil && r.CheckTransferIntention != nil {
		CheckTransferIntention = *r.CheckTransferIntention
	}
	return
}

// A Check Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_rejection`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetCheckTransferRejection() (CheckTransferRejection ACHTransfersTransactionSourceCheckTransferRejectionCreateInboundResponse) {
	if r != nil && r.CheckTransferRejection != nil {
		CheckTransferRejection = *r.CheckTransferRejection
	}
	return
}

// A Check Transfer Stop Payment Request object. This field will be present in the
// JSON response if and only if `category` is equal to
// `check_transfer_stop_payment_request`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetCheckTransferStopPaymentRequest() (CheckTransferStopPaymentRequest ACHTransfersTransactionSourceCheckTransferStopPaymentRequestCreateInboundResponse) {
	if r != nil && r.CheckTransferStopPaymentRequest != nil {
		CheckTransferStopPaymentRequest = *r.CheckTransferStopPaymentRequest
	}
	return
}

// A Dispute Resolution object. This field will be present in the JSON response if
// and only if `category` is equal to `dispute_resolution`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetDisputeResolution() (DisputeResolution ACHTransfersTransactionSourceDisputeResolutionCreateInboundResponse) {
	if r != nil && r.DisputeResolution != nil {
		DisputeResolution = *r.DisputeResolution
	}
	return
}

// A Empyreal Cash Deposit object. This field will be present in the JSON response
// if and only if `category` is equal to `empyreal_cash_deposit`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetEmpyrealCashDeposit() (EmpyrealCashDeposit ACHTransfersTransactionSourceEmpyrealCashDepositCreateInboundResponse) {
	if r != nil && r.EmpyrealCashDeposit != nil {
		EmpyrealCashDeposit = *r.EmpyrealCashDeposit
	}
	return
}

// A Inbound ACH Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_ach_transfer`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetInboundACHTransfer() (InboundACHTransfer ACHTransfersTransactionSourceInboundACHTransferCreateInboundResponse) {
	if r != nil && r.InboundACHTransfer != nil {
		InboundACHTransfer = *r.InboundACHTransfer
	}
	return
}

// A Inbound Check object. This field will be present in the JSON response if and
// only if `category` is equal to `inbound_check`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetInboundCheck() (InboundCheck ACHTransfersTransactionSourceInboundCheckCreateInboundResponse) {
	if r != nil && r.InboundCheck != nil {
		InboundCheck = *r.InboundCheck
	}
	return
}

// A Inbound International ACH Transfer object. This field will be present in the
// JSON response if and only if `category` is equal to
// `inbound_international_ach_transfer`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetInboundInternationalACHTransfer() (InboundInternationalACHTransfer ACHTransfersTransactionSourceInboundInternationalACHTransferCreateInboundResponse) {
	if r != nil && r.InboundInternationalACHTransfer != nil {
		InboundInternationalACHTransfer = *r.InboundInternationalACHTransfer
	}
	return
}

// A Inbound Real Time Payments Transfer Confirmation object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_confirmation`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetInboundRealTimePaymentsTransferConfirmation() (InboundRealTimePaymentsTransferConfirmation ACHTransfersTransactionSourceInboundRealTimePaymentsTransferConfirmationCreateInboundResponse) {
	if r != nil && r.InboundRealTimePaymentsTransferConfirmation != nil {
		InboundRealTimePaymentsTransferConfirmation = *r.InboundRealTimePaymentsTransferConfirmation
	}
	return
}

// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
// the JSON response if and only if `category` is equal to
// `inbound_wire_drawdown_payment_reversal`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetInboundWireDrawdownPaymentReversal() (InboundWireDrawdownPaymentReversal ACHTransfersTransactionSourceInboundWireDrawdownPaymentReversalCreateInboundResponse) {
	if r != nil && r.InboundWireDrawdownPaymentReversal != nil {
		InboundWireDrawdownPaymentReversal = *r.InboundWireDrawdownPaymentReversal
	}
	return
}

// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetInboundWireDrawdownPayment() (InboundWireDrawdownPayment ACHTransfersTransactionSourceInboundWireDrawdownPaymentCreateInboundResponse) {
	if r != nil && r.InboundWireDrawdownPayment != nil {
		InboundWireDrawdownPayment = *r.InboundWireDrawdownPayment
	}
	return
}

// A Inbound Wire Reversal object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_reversal`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetInboundWireReversal() (InboundWireReversal ACHTransfersTransactionSourceInboundWireReversalCreateInboundResponse) {
	if r != nil && r.InboundWireReversal != nil {
		InboundWireReversal = *r.InboundWireReversal
	}
	return
}

// A Inbound Wire Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_transfer`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetInboundWireTransfer() (InboundWireTransfer ACHTransfersTransactionSourceInboundWireTransferCreateInboundResponse) {
	if r != nil && r.InboundWireTransfer != nil {
		InboundWireTransfer = *r.InboundWireTransfer
	}
	return
}

// A Internal Source object. This field will be present in the JSON response if and
// only if `category` is equal to `internal_source`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetInternalSource() (InternalSource ACHTransfersTransactionSourceInternalSourceCreateInboundResponse) {
	if r != nil && r.InternalSource != nil {
		InternalSource = *r.InternalSource
	}
	return
}

// A Deprecated Card Refund object. This field will be present in the JSON response
// if and only if `category` is equal to `card_route_refund`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetCardRouteRefund() (CardRouteRefund ACHTransfersTransactionSourceCardRouteRefundCreateInboundResponse) {
	if r != nil && r.CardRouteRefund != nil {
		CardRouteRefund = *r.CardRouteRefund
	}
	return
}

// A Deprecated Card Settlement object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_settlement`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetCardRouteSettlement() (CardRouteSettlement ACHTransfersTransactionSourceCardRouteSettlementCreateInboundResponse) {
	if r != nil && r.CardRouteSettlement != nil {
		CardRouteSettlement = *r.CardRouteSettlement
	}
	return
}

// A Sample Funds object. This field will be present in the JSON response if and
// only if `category` is equal to `sample_funds`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetSampleFunds() (SampleFunds ACHTransfersTransactionSourceSampleFundsCreateInboundResponse) {
	if r != nil && r.SampleFunds != nil {
		SampleFunds = *r.SampleFunds
	}
	return
}

// A Wire Drawdown Payment Intention object. This field will be present in the JSON
// response if and only if `category` is equal to
// `wire_drawdown_payment_intention`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetWireDrawdownPaymentIntention() (WireDrawdownPaymentIntention ACHTransfersTransactionSourceWireDrawdownPaymentIntentionCreateInboundResponse) {
	if r != nil && r.WireDrawdownPaymentIntention != nil {
		WireDrawdownPaymentIntention = *r.WireDrawdownPaymentIntention
	}
	return
}

// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to
// `wire_drawdown_payment_rejection`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetWireDrawdownPaymentRejection() (WireDrawdownPaymentRejection ACHTransfersTransactionSourceWireDrawdownPaymentRejectionCreateInboundResponse) {
	if r != nil && r.WireDrawdownPaymentRejection != nil {
		WireDrawdownPaymentRejection = *r.WireDrawdownPaymentRejection
	}
	return
}

// A Wire Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_intention`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetWireTransferIntention() (WireTransferIntention ACHTransfersTransactionSourceWireTransferIntentionCreateInboundResponse) {
	if r != nil && r.WireTransferIntention != nil {
		WireTransferIntention = *r.WireTransferIntention
	}
	return
}

// A Wire Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_rejection`.
func (r *ACHTransfersTransactionSourceCreateInboundResponse) GetWireTransferRejection() (WireTransferRejection ACHTransfersTransactionSourceWireTransferRejectionCreateInboundResponse) {
	if r != nil && r.WireTransferRejection != nil {
		WireTransferRejection = *r.WireTransferRejection
	}
	return
}
