package simulations

type WireTransfersTransactionSourceCreateInboundResponse struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category *string `json:"category,omitempty"`

	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention *WireTransfersTransactionSourceAccountTransferIntentionCreateInboundResponse `json:"account_transfer_intention,omitempty"`

	// A ACH Check Conversion Return object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_check_conversion_return`.
	ACHCheckConversionReturn *WireTransfersTransactionSourceACHCheckConversionReturnCreateInboundResponse `json:"ach_check_conversion_return,omitempty"`

	// A ACH Check Conversion object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_check_conversion`.
	ACHCheckConversion *WireTransfersTransactionSourceACHCheckConversionCreateInboundResponse `json:"ach_check_conversion,omitempty"`

	// A ACH Transfer Intention object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention *WireTransfersTransactionSourceACHTransferIntentionCreateInboundResponse `json:"ach_transfer_intention,omitempty"`

	// A ACH Transfer Rejection object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection *WireTransfersTransactionSourceACHTransferRejectionCreateInboundResponse `json:"ach_transfer_rejection,omitempty"`

	// A ACH Transfer Return object. This field will be present in the JSON response if
	// and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn *WireTransfersTransactionSourceACHTransferReturnCreateInboundResponse `json:"ach_transfer_return,omitempty"`

	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance *WireTransfersTransactionSourceCardDisputeAcceptanceCreateInboundResponse `json:"card_dispute_acceptance,omitempty"`

	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund *WireTransfersTransactionSourceCardRefundCreateInboundResponse `json:"card_refund,omitempty"`

	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement *WireTransfersTransactionSourceCardSettlementCreateInboundResponse `json:"card_settlement,omitempty"`

	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance *WireTransfersTransactionSourceCheckDepositAcceptanceCreateInboundResponse `json:"check_deposit_acceptance,omitempty"`

	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn *WireTransfersTransactionSourceCheckDepositReturnCreateInboundResponse `json:"check_deposit_return,omitempty"`

	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention *WireTransfersTransactionSourceCheckTransferIntentionCreateInboundResponse `json:"check_transfer_intention,omitempty"`

	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection *WireTransfersTransactionSourceCheckTransferRejectionCreateInboundResponse `json:"check_transfer_rejection,omitempty"`

	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest *WireTransfersTransactionSourceCheckTransferStopPaymentRequestCreateInboundResponse `json:"check_transfer_stop_payment_request,omitempty"`

	// A Dispute Resolution object. This field will be present in the JSON response if
	// and only if `category` is equal to `dispute_resolution`.
	DisputeResolution *WireTransfersTransactionSourceDisputeResolutionCreateInboundResponse `json:"dispute_resolution,omitempty"`

	// A Empyreal Cash Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `empyreal_cash_deposit`.
	EmpyrealCashDeposit *WireTransfersTransactionSourceEmpyrealCashDepositCreateInboundResponse `json:"empyreal_cash_deposit,omitempty"`

	// A Inbound ACH Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer *WireTransfersTransactionSourceInboundACHTransferCreateInboundResponse `json:"inbound_ach_transfer,omitempty"`

	// A Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck *WireTransfersTransactionSourceInboundCheckCreateInboundResponse `json:"inbound_check,omitempty"`

	// A Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer *WireTransfersTransactionSourceInboundInternationalACHTransferCreateInboundResponse `json:"inbound_international_ach_transfer,omitempty"`

	// A Inbound Real Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation *WireTransfersTransactionSourceInboundRealTimePaymentsTransferConfirmationCreateInboundResponse `json:"inbound_real_time_payments_transfer_confirmation,omitempty"`

	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal *WireTransfersTransactionSourceInboundWireDrawdownPaymentReversalCreateInboundResponse `json:"inbound_wire_drawdown_payment_reversal,omitempty"`

	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment *WireTransfersTransactionSourceInboundWireDrawdownPaymentCreateInboundResponse `json:"inbound_wire_drawdown_payment,omitempty"`

	// A Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal *WireTransfersTransactionSourceInboundWireReversalCreateInboundResponse `json:"inbound_wire_reversal,omitempty"`

	// A Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer *WireTransfersTransactionSourceInboundWireTransferCreateInboundResponse `json:"inbound_wire_transfer,omitempty"`

	// A Internal Source object. This field will be present in the JSON response if and
	// only if `category` is equal to `internal_source`.
	InternalSource *WireTransfersTransactionSourceInternalSourceCreateInboundResponse `json:"internal_source,omitempty"`

	// A Deprecated Card Refund object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_route_refund`.
	CardRouteRefund *WireTransfersTransactionSourceCardRouteRefundCreateInboundResponse `json:"card_route_refund,omitempty"`

	// A Deprecated Card Settlement object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_settlement`.
	CardRouteSettlement *WireTransfersTransactionSourceCardRouteSettlementCreateInboundResponse `json:"card_route_settlement,omitempty"`

	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds *WireTransfersTransactionSourceSampleFundsCreateInboundResponse `json:"sample_funds,omitempty"`

	// A Wire Drawdown Payment Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_intention`.
	WireDrawdownPaymentIntention *WireTransfersTransactionSourceWireDrawdownPaymentIntentionCreateInboundResponse `json:"wire_drawdown_payment_intention,omitempty"`

	// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_rejection`.
	WireDrawdownPaymentRejection *WireTransfersTransactionSourceWireDrawdownPaymentRejectionCreateInboundResponse `json:"wire_drawdown_payment_rejection,omitempty"`

	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention *WireTransfersTransactionSourceWireTransferIntentionCreateInboundResponse `json:"wire_transfer_intention,omitempty"`

	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection *WireTransfersTransactionSourceWireTransferRejectionCreateInboundResponse `json:"wire_transfer_rejection,omitempty"`
}

// The type of transaction that took place. We may add additional possible values
// for this enum over time; your application should be able to handle such
// additions gracefully.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetCategory() (Category string) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// A Account Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_intention`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetAccountTransferIntention() (AccountTransferIntention WireTransfersTransactionSourceAccountTransferIntentionCreateInboundResponse) {
	if r != nil && r.AccountTransferIntention != nil {
		AccountTransferIntention = *r.AccountTransferIntention
	}
	return
}

// A ACH Check Conversion Return object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_check_conversion_return`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetACHCheckConversionReturn() (ACHCheckConversionReturn WireTransfersTransactionSourceACHCheckConversionReturnCreateInboundResponse) {
	if r != nil && r.ACHCheckConversionReturn != nil {
		ACHCheckConversionReturn = *r.ACHCheckConversionReturn
	}
	return
}

// A ACH Check Conversion object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_check_conversion`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetACHCheckConversion() (ACHCheckConversion WireTransfersTransactionSourceACHCheckConversionCreateInboundResponse) {
	if r != nil && r.ACHCheckConversion != nil {
		ACHCheckConversion = *r.ACHCheckConversion
	}
	return
}

// A ACH Transfer Intention object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_intention`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetACHTransferIntention() (ACHTransferIntention WireTransfersTransactionSourceACHTransferIntentionCreateInboundResponse) {
	if r != nil && r.ACHTransferIntention != nil {
		ACHTransferIntention = *r.ACHTransferIntention
	}
	return
}

// A ACH Transfer Rejection object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_rejection`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetACHTransferRejection() (ACHTransferRejection WireTransfersTransactionSourceACHTransferRejectionCreateInboundResponse) {
	if r != nil && r.ACHTransferRejection != nil {
		ACHTransferRejection = *r.ACHTransferRejection
	}
	return
}

// A ACH Transfer Return object. This field will be present in the JSON response if
// and only if `category` is equal to `ach_transfer_return`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetACHTransferReturn() (ACHTransferReturn WireTransfersTransactionSourceACHTransferReturnCreateInboundResponse) {
	if r != nil && r.ACHTransferReturn != nil {
		ACHTransferReturn = *r.ACHTransferReturn
	}
	return
}

// A Card Dispute Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `card_dispute_acceptance`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetCardDisputeAcceptance() (CardDisputeAcceptance WireTransfersTransactionSourceCardDisputeAcceptanceCreateInboundResponse) {
	if r != nil && r.CardDisputeAcceptance != nil {
		CardDisputeAcceptance = *r.CardDisputeAcceptance
	}
	return
}

// A Card Refund object. This field will be present in the JSON response if and
// only if `category` is equal to `card_refund`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetCardRefund() (CardRefund WireTransfersTransactionSourceCardRefundCreateInboundResponse) {
	if r != nil && r.CardRefund != nil {
		CardRefund = *r.CardRefund
	}
	return
}

// A Card Settlement object. This field will be present in the JSON response if and
// only if `category` is equal to `card_settlement`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetCardSettlement() (CardSettlement WireTransfersTransactionSourceCardSettlementCreateInboundResponse) {
	if r != nil && r.CardSettlement != nil {
		CardSettlement = *r.CardSettlement
	}
	return
}

// A Check Deposit Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_acceptance`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetCheckDepositAcceptance() (CheckDepositAcceptance WireTransfersTransactionSourceCheckDepositAcceptanceCreateInboundResponse) {
	if r != nil && r.CheckDepositAcceptance != nil {
		CheckDepositAcceptance = *r.CheckDepositAcceptance
	}
	return
}

// A Check Deposit Return object. This field will be present in the JSON response
// if and only if `category` is equal to `check_deposit_return`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetCheckDepositReturn() (CheckDepositReturn WireTransfersTransactionSourceCheckDepositReturnCreateInboundResponse) {
	if r != nil && r.CheckDepositReturn != nil {
		CheckDepositReturn = *r.CheckDepositReturn
	}
	return
}

// A Check Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_intention`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetCheckTransferIntention() (CheckTransferIntention WireTransfersTransactionSourceCheckTransferIntentionCreateInboundResponse) {
	if r != nil && r.CheckTransferIntention != nil {
		CheckTransferIntention = *r.CheckTransferIntention
	}
	return
}

// A Check Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_rejection`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetCheckTransferRejection() (CheckTransferRejection WireTransfersTransactionSourceCheckTransferRejectionCreateInboundResponse) {
	if r != nil && r.CheckTransferRejection != nil {
		CheckTransferRejection = *r.CheckTransferRejection
	}
	return
}

// A Check Transfer Stop Payment Request object. This field will be present in the
// JSON response if and only if `category` is equal to
// `check_transfer_stop_payment_request`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetCheckTransferStopPaymentRequest() (CheckTransferStopPaymentRequest WireTransfersTransactionSourceCheckTransferStopPaymentRequestCreateInboundResponse) {
	if r != nil && r.CheckTransferStopPaymentRequest != nil {
		CheckTransferStopPaymentRequest = *r.CheckTransferStopPaymentRequest
	}
	return
}

// A Dispute Resolution object. This field will be present in the JSON response if
// and only if `category` is equal to `dispute_resolution`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetDisputeResolution() (DisputeResolution WireTransfersTransactionSourceDisputeResolutionCreateInboundResponse) {
	if r != nil && r.DisputeResolution != nil {
		DisputeResolution = *r.DisputeResolution
	}
	return
}

// A Empyreal Cash Deposit object. This field will be present in the JSON response
// if and only if `category` is equal to `empyreal_cash_deposit`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetEmpyrealCashDeposit() (EmpyrealCashDeposit WireTransfersTransactionSourceEmpyrealCashDepositCreateInboundResponse) {
	if r != nil && r.EmpyrealCashDeposit != nil {
		EmpyrealCashDeposit = *r.EmpyrealCashDeposit
	}
	return
}

// A Inbound ACH Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_ach_transfer`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetInboundACHTransfer() (InboundACHTransfer WireTransfersTransactionSourceInboundACHTransferCreateInboundResponse) {
	if r != nil && r.InboundACHTransfer != nil {
		InboundACHTransfer = *r.InboundACHTransfer
	}
	return
}

// A Inbound Check object. This field will be present in the JSON response if and
// only if `category` is equal to `inbound_check`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetInboundCheck() (InboundCheck WireTransfersTransactionSourceInboundCheckCreateInboundResponse) {
	if r != nil && r.InboundCheck != nil {
		InboundCheck = *r.InboundCheck
	}
	return
}

// A Inbound International ACH Transfer object. This field will be present in the
// JSON response if and only if `category` is equal to
// `inbound_international_ach_transfer`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetInboundInternationalACHTransfer() (InboundInternationalACHTransfer WireTransfersTransactionSourceInboundInternationalACHTransferCreateInboundResponse) {
	if r != nil && r.InboundInternationalACHTransfer != nil {
		InboundInternationalACHTransfer = *r.InboundInternationalACHTransfer
	}
	return
}

// A Inbound Real Time Payments Transfer Confirmation object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_confirmation`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetInboundRealTimePaymentsTransferConfirmation() (InboundRealTimePaymentsTransferConfirmation WireTransfersTransactionSourceInboundRealTimePaymentsTransferConfirmationCreateInboundResponse) {
	if r != nil && r.InboundRealTimePaymentsTransferConfirmation != nil {
		InboundRealTimePaymentsTransferConfirmation = *r.InboundRealTimePaymentsTransferConfirmation
	}
	return
}

// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
// the JSON response if and only if `category` is equal to
// `inbound_wire_drawdown_payment_reversal`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetInboundWireDrawdownPaymentReversal() (InboundWireDrawdownPaymentReversal WireTransfersTransactionSourceInboundWireDrawdownPaymentReversalCreateInboundResponse) {
	if r != nil && r.InboundWireDrawdownPaymentReversal != nil {
		InboundWireDrawdownPaymentReversal = *r.InboundWireDrawdownPaymentReversal
	}
	return
}

// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetInboundWireDrawdownPayment() (InboundWireDrawdownPayment WireTransfersTransactionSourceInboundWireDrawdownPaymentCreateInboundResponse) {
	if r != nil && r.InboundWireDrawdownPayment != nil {
		InboundWireDrawdownPayment = *r.InboundWireDrawdownPayment
	}
	return
}

// A Inbound Wire Reversal object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_reversal`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetInboundWireReversal() (InboundWireReversal WireTransfersTransactionSourceInboundWireReversalCreateInboundResponse) {
	if r != nil && r.InboundWireReversal != nil {
		InboundWireReversal = *r.InboundWireReversal
	}
	return
}

// A Inbound Wire Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_transfer`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetInboundWireTransfer() (InboundWireTransfer WireTransfersTransactionSourceInboundWireTransferCreateInboundResponse) {
	if r != nil && r.InboundWireTransfer != nil {
		InboundWireTransfer = *r.InboundWireTransfer
	}
	return
}

// A Internal Source object. This field will be present in the JSON response if and
// only if `category` is equal to `internal_source`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetInternalSource() (InternalSource WireTransfersTransactionSourceInternalSourceCreateInboundResponse) {
	if r != nil && r.InternalSource != nil {
		InternalSource = *r.InternalSource
	}
	return
}

// A Deprecated Card Refund object. This field will be present in the JSON response
// if and only if `category` is equal to `card_route_refund`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetCardRouteRefund() (CardRouteRefund WireTransfersTransactionSourceCardRouteRefundCreateInboundResponse) {
	if r != nil && r.CardRouteRefund != nil {
		CardRouteRefund = *r.CardRouteRefund
	}
	return
}

// A Deprecated Card Settlement object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_settlement`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetCardRouteSettlement() (CardRouteSettlement WireTransfersTransactionSourceCardRouteSettlementCreateInboundResponse) {
	if r != nil && r.CardRouteSettlement != nil {
		CardRouteSettlement = *r.CardRouteSettlement
	}
	return
}

// A Sample Funds object. This field will be present in the JSON response if and
// only if `category` is equal to `sample_funds`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetSampleFunds() (SampleFunds WireTransfersTransactionSourceSampleFundsCreateInboundResponse) {
	if r != nil && r.SampleFunds != nil {
		SampleFunds = *r.SampleFunds
	}
	return
}

// A Wire Drawdown Payment Intention object. This field will be present in the JSON
// response if and only if `category` is equal to
// `wire_drawdown_payment_intention`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetWireDrawdownPaymentIntention() (WireDrawdownPaymentIntention WireTransfersTransactionSourceWireDrawdownPaymentIntentionCreateInboundResponse) {
	if r != nil && r.WireDrawdownPaymentIntention != nil {
		WireDrawdownPaymentIntention = *r.WireDrawdownPaymentIntention
	}
	return
}

// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to
// `wire_drawdown_payment_rejection`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetWireDrawdownPaymentRejection() (WireDrawdownPaymentRejection WireTransfersTransactionSourceWireDrawdownPaymentRejectionCreateInboundResponse) {
	if r != nil && r.WireDrawdownPaymentRejection != nil {
		WireDrawdownPaymentRejection = *r.WireDrawdownPaymentRejection
	}
	return
}

// A Wire Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_intention`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetWireTransferIntention() (WireTransferIntention WireTransfersTransactionSourceWireTransferIntentionCreateInboundResponse) {
	if r != nil && r.WireTransferIntention != nil {
		WireTransferIntention = *r.WireTransferIntention
	}
	return
}

// A Wire Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_rejection`.
func (r *WireTransfersTransactionSourceCreateInboundResponse) GetWireTransferRejection() (WireTransferRejection WireTransfersTransactionSourceWireTransferRejectionCreateInboundResponse) {
	if r != nil && r.WireTransferRejection != nil {
		WireTransferRejection = *r.WireTransferRejection
	}
	return
}
