package types

import (
	"fmt"
	"increase/core"
	"increase/core/pjson"
	"increase/pagination"
)

type Transaction struct {
	// The identifier for the Account the Transaction belongs to.
	AccountID *string `pjson:"account_id"`
	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transcation's
	// Account.
	Currency *TransactionCurrency `pjson:"currency"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt *string `pjson:"created_at"`
	// For a Transaction related to a transfer, this is the description you provide.
	// For a Transaction related to a payment, this is the description the vendor
	// provides.
	Description *string `pjson:"description"`
	// The Transaction identifier.
	ID *string `pjson:"id"`
	// The identifier for the route this Transaction came through. Routes are things
	// like cards and ACH details.
	RouteID *string `pjson:"route_id"`
	// The type of the route this Transaction came through.
	RouteType *string `pjson:"route_type"`
	// This is an object giving more details on the network-level event that caused the
	// Transaction. Note that for backwards compatibility reasons, additional
	// undocumented keys may appear in this object. These should be treated as
	// deprecated and will be removed in the future.
	Source *TransactionSource `pjson:"source"`
	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type       *TransactionType       `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into Transaction using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *Transaction) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes Transaction into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *Transaction) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier for the Account the Transaction belongs to.
func (r *Transaction) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The Transaction amount in the minor unit of its currency. For dollars, for
// example, this is cents.
func (r *Transaction) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// Transaction's currency. This will match the currency on the Transcation's
// Account.
func (r *Transaction) GetCurrency() (Currency TransactionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
// Transaction occured.
func (r *Transaction) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// For a Transaction related to a transfer, this is the description you provide.
// For a Transaction related to a payment, this is the description the vendor
// provides.
func (r *Transaction) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Transaction identifier.
func (r *Transaction) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The identifier for the route this Transaction came through. Routes are things
// like cards and ACH details.
func (r *Transaction) GetRouteID() (RouteID string) {
	if r != nil && r.RouteID != nil {
		RouteID = *r.RouteID
	}
	return
}

// The type of the route this Transaction came through.
func (r *Transaction) GetRouteType() (RouteType string) {
	if r != nil && r.RouteType != nil {
		RouteType = *r.RouteType
	}
	return
}

// This is an object giving more details on the network-level event that caused the
// Transaction. Note that for backwards compatibility reasons, additional
// undocumented keys may appear in this object. These should be treated as
// deprecated and will be removed in the future.
func (r *Transaction) GetSource() (Source TransactionSource) {
	if r != nil && r.Source != nil {
		Source = *r.Source
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `transaction`.
func (r *Transaction) GetType() (Type TransactionType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r Transaction) String() (result string) {
	return fmt.Sprintf("&Transaction{AccountID:%s Amount:%s Currency:%s CreatedAt:%s Description:%s ID:%s RouteID:%s RouteType:%s Source:%s Type:%s}", core.FmtP(r.AccountID), core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.CreatedAt), core.FmtP(r.Description), core.FmtP(r.ID), core.FmtP(r.RouteID), core.FmtP(r.RouteType), r.Source, core.FmtP(r.Type))
}

type TransactionCurrency string

const (
	TransactionCurrencyCad TransactionCurrency = "CAD"
	TransactionCurrencyChf TransactionCurrency = "CHF"
	TransactionCurrencyEur TransactionCurrency = "EUR"
	TransactionCurrencyGbp TransactionCurrency = "GBP"
	TransactionCurrencyJpy TransactionCurrency = "JPY"
	TransactionCurrencyUsd TransactionCurrency = "USD"
)

type TransactionSource struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category *TransactionSourceCategory `pjson:"category"`
	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention *TransactionSourceAccountTransferIntention `pjson:"account_transfer_intention"`
	// A ACH Check Conversion Return object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_check_conversion_return`.
	ACHCheckConversionReturn *TransactionSourceACHCheckConversionReturn `pjson:"ach_check_conversion_return"`
	// A ACH Check Conversion object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_check_conversion`.
	ACHCheckConversion *TransactionSourceACHCheckConversion `pjson:"ach_check_conversion"`
	// A ACH Transfer Intention object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention *TransactionSourceACHTransferIntention `pjson:"ach_transfer_intention"`
	// A ACH Transfer Rejection object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection *TransactionSourceACHTransferRejection `pjson:"ach_transfer_rejection"`
	// A ACH Transfer Return object. This field will be present in the JSON response if
	// and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn *TransactionSourceACHTransferReturn `pjson:"ach_transfer_return"`
	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance *TransactionSourceCardDisputeAcceptance `pjson:"card_dispute_acceptance"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund *TransactionSourceCardRefund `pjson:"card_refund"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement *TransactionSourceCardSettlement `pjson:"card_settlement"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance *TransactionSourceCheckDepositAcceptance `pjson:"check_deposit_acceptance"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn *TransactionSourceCheckDepositReturn `pjson:"check_deposit_return"`
	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention *TransactionSourceCheckTransferIntention `pjson:"check_transfer_intention"`
	// A Check Transfer Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_transfer_return`.
	CheckTransferReturn *TransactionSourceCheckTransferReturn `pjson:"check_transfer_return"`
	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection *TransactionSourceCheckTransferRejection `pjson:"check_transfer_rejection"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest *TransactionSourceCheckTransferStopPaymentRequest `pjson:"check_transfer_stop_payment_request"`
	// A Dispute Resolution object. This field will be present in the JSON response if
	// and only if `category` is equal to `dispute_resolution`.
	DisputeResolution *TransactionSourceDisputeResolution `pjson:"dispute_resolution"`
	// A Empyreal Cash Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `empyreal_cash_deposit`.
	EmpyrealCashDeposit *TransactionSourceEmpyrealCashDeposit `pjson:"empyreal_cash_deposit"`
	// A Inbound ACH Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer *TransactionSourceInboundACHTransfer `pjson:"inbound_ach_transfer"`
	// A Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck *TransactionSourceInboundCheck `pjson:"inbound_check"`
	// A Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer *TransactionSourceInboundInternationalACHTransfer `pjson:"inbound_international_ach_transfer"`
	// A Inbound Real Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation *TransactionSourceInboundRealTimePaymentsTransferConfirmation `pjson:"inbound_real_time_payments_transfer_confirmation"`
	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal *TransactionSourceInboundWireDrawdownPaymentReversal `pjson:"inbound_wire_drawdown_payment_reversal"`
	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment *TransactionSourceInboundWireDrawdownPayment `pjson:"inbound_wire_drawdown_payment"`
	// A Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal *TransactionSourceInboundWireReversal `pjson:"inbound_wire_reversal"`
	// A Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer *TransactionSourceInboundWireTransfer `pjson:"inbound_wire_transfer"`
	// A Internal Source object. This field will be present in the JSON response if and
	// only if `category` is equal to `internal_source`.
	InternalSource *TransactionSourceInternalSource `pjson:"internal_source"`
	// A Deprecated Card Refund object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_route_refund`.
	CardRouteRefund *TransactionSourceCardRouteRefund `pjson:"card_route_refund"`
	// A Deprecated Card Settlement object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_settlement`.
	CardRouteSettlement *TransactionSourceCardRouteSettlement `pjson:"card_route_settlement"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds *TransactionSourceSampleFunds `pjson:"sample_funds"`
	// A Wire Drawdown Payment Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_intention`.
	WireDrawdownPaymentIntention *TransactionSourceWireDrawdownPaymentIntention `pjson:"wire_drawdown_payment_intention"`
	// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_rejection`.
	WireDrawdownPaymentRejection *TransactionSourceWireDrawdownPaymentRejection `pjson:"wire_drawdown_payment_rejection"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention *TransactionSourceWireTransferIntention `pjson:"wire_transfer_intention"`
	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection *TransactionSourceWireTransferRejection `pjson:"wire_transfer_rejection"`
	jsonFields            map[string]interface{}                  `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into TransactionSource using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *TransactionSource) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSource into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *TransactionSource) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The type of transaction that took place. We may add additional possible values
// for this enum over time; your application should be able to handle such
// additions gracefully.
func (r *TransactionSource) GetCategory() (Category TransactionSourceCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// A Account Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_intention`.
func (r *TransactionSource) GetAccountTransferIntention() (AccountTransferIntention TransactionSourceAccountTransferIntention) {
	if r != nil && r.AccountTransferIntention != nil {
		AccountTransferIntention = *r.AccountTransferIntention
	}
	return
}

// A ACH Check Conversion Return object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_check_conversion_return`.
func (r *TransactionSource) GetACHCheckConversionReturn() (ACHCheckConversionReturn TransactionSourceACHCheckConversionReturn) {
	if r != nil && r.ACHCheckConversionReturn != nil {
		ACHCheckConversionReturn = *r.ACHCheckConversionReturn
	}
	return
}

// A ACH Check Conversion object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_check_conversion`.
func (r *TransactionSource) GetACHCheckConversion() (ACHCheckConversion TransactionSourceACHCheckConversion) {
	if r != nil && r.ACHCheckConversion != nil {
		ACHCheckConversion = *r.ACHCheckConversion
	}
	return
}

// A ACH Transfer Intention object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_intention`.
func (r *TransactionSource) GetACHTransferIntention() (ACHTransferIntention TransactionSourceACHTransferIntention) {
	if r != nil && r.ACHTransferIntention != nil {
		ACHTransferIntention = *r.ACHTransferIntention
	}
	return
}

// A ACH Transfer Rejection object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_rejection`.
func (r *TransactionSource) GetACHTransferRejection() (ACHTransferRejection TransactionSourceACHTransferRejection) {
	if r != nil && r.ACHTransferRejection != nil {
		ACHTransferRejection = *r.ACHTransferRejection
	}
	return
}

// A ACH Transfer Return object. This field will be present in the JSON response if
// and only if `category` is equal to `ach_transfer_return`.
func (r *TransactionSource) GetACHTransferReturn() (ACHTransferReturn TransactionSourceACHTransferReturn) {
	if r != nil && r.ACHTransferReturn != nil {
		ACHTransferReturn = *r.ACHTransferReturn
	}
	return
}

// A Card Dispute Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `card_dispute_acceptance`.
func (r *TransactionSource) GetCardDisputeAcceptance() (CardDisputeAcceptance TransactionSourceCardDisputeAcceptance) {
	if r != nil && r.CardDisputeAcceptance != nil {
		CardDisputeAcceptance = *r.CardDisputeAcceptance
	}
	return
}

// A Card Refund object. This field will be present in the JSON response if and
// only if `category` is equal to `card_refund`.
func (r *TransactionSource) GetCardRefund() (CardRefund TransactionSourceCardRefund) {
	if r != nil && r.CardRefund != nil {
		CardRefund = *r.CardRefund
	}
	return
}

// A Card Settlement object. This field will be present in the JSON response if and
// only if `category` is equal to `card_settlement`.
func (r *TransactionSource) GetCardSettlement() (CardSettlement TransactionSourceCardSettlement) {
	if r != nil && r.CardSettlement != nil {
		CardSettlement = *r.CardSettlement
	}
	return
}

// A Check Deposit Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_acceptance`.
func (r *TransactionSource) GetCheckDepositAcceptance() (CheckDepositAcceptance TransactionSourceCheckDepositAcceptance) {
	if r != nil && r.CheckDepositAcceptance != nil {
		CheckDepositAcceptance = *r.CheckDepositAcceptance
	}
	return
}

// A Check Deposit Return object. This field will be present in the JSON response
// if and only if `category` is equal to `check_deposit_return`.
func (r *TransactionSource) GetCheckDepositReturn() (CheckDepositReturn TransactionSourceCheckDepositReturn) {
	if r != nil && r.CheckDepositReturn != nil {
		CheckDepositReturn = *r.CheckDepositReturn
	}
	return
}

// A Check Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_intention`.
func (r *TransactionSource) GetCheckTransferIntention() (CheckTransferIntention TransactionSourceCheckTransferIntention) {
	if r != nil && r.CheckTransferIntention != nil {
		CheckTransferIntention = *r.CheckTransferIntention
	}
	return
}

// A Check Transfer Return object. This field will be present in the JSON response
// if and only if `category` is equal to `check_transfer_return`.
func (r *TransactionSource) GetCheckTransferReturn() (CheckTransferReturn TransactionSourceCheckTransferReturn) {
	if r != nil && r.CheckTransferReturn != nil {
		CheckTransferReturn = *r.CheckTransferReturn
	}
	return
}

// A Check Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_rejection`.
func (r *TransactionSource) GetCheckTransferRejection() (CheckTransferRejection TransactionSourceCheckTransferRejection) {
	if r != nil && r.CheckTransferRejection != nil {
		CheckTransferRejection = *r.CheckTransferRejection
	}
	return
}

// A Check Transfer Stop Payment Request object. This field will be present in the
// JSON response if and only if `category` is equal to
// `check_transfer_stop_payment_request`.
func (r *TransactionSource) GetCheckTransferStopPaymentRequest() (CheckTransferStopPaymentRequest TransactionSourceCheckTransferStopPaymentRequest) {
	if r != nil && r.CheckTransferStopPaymentRequest != nil {
		CheckTransferStopPaymentRequest = *r.CheckTransferStopPaymentRequest
	}
	return
}

// A Dispute Resolution object. This field will be present in the JSON response if
// and only if `category` is equal to `dispute_resolution`.
func (r *TransactionSource) GetDisputeResolution() (DisputeResolution TransactionSourceDisputeResolution) {
	if r != nil && r.DisputeResolution != nil {
		DisputeResolution = *r.DisputeResolution
	}
	return
}

// A Empyreal Cash Deposit object. This field will be present in the JSON response
// if and only if `category` is equal to `empyreal_cash_deposit`.
func (r *TransactionSource) GetEmpyrealCashDeposit() (EmpyrealCashDeposit TransactionSourceEmpyrealCashDeposit) {
	if r != nil && r.EmpyrealCashDeposit != nil {
		EmpyrealCashDeposit = *r.EmpyrealCashDeposit
	}
	return
}

// A Inbound ACH Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_ach_transfer`.
func (r *TransactionSource) GetInboundACHTransfer() (InboundACHTransfer TransactionSourceInboundACHTransfer) {
	if r != nil && r.InboundACHTransfer != nil {
		InboundACHTransfer = *r.InboundACHTransfer
	}
	return
}

// A Inbound Check object. This field will be present in the JSON response if and
// only if `category` is equal to `inbound_check`.
func (r *TransactionSource) GetInboundCheck() (InboundCheck TransactionSourceInboundCheck) {
	if r != nil && r.InboundCheck != nil {
		InboundCheck = *r.InboundCheck
	}
	return
}

// A Inbound International ACH Transfer object. This field will be present in the
// JSON response if and only if `category` is equal to
// `inbound_international_ach_transfer`.
func (r *TransactionSource) GetInboundInternationalACHTransfer() (InboundInternationalACHTransfer TransactionSourceInboundInternationalACHTransfer) {
	if r != nil && r.InboundInternationalACHTransfer != nil {
		InboundInternationalACHTransfer = *r.InboundInternationalACHTransfer
	}
	return
}

// A Inbound Real Time Payments Transfer Confirmation object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_confirmation`.
func (r *TransactionSource) GetInboundRealTimePaymentsTransferConfirmation() (InboundRealTimePaymentsTransferConfirmation TransactionSourceInboundRealTimePaymentsTransferConfirmation) {
	if r != nil && r.InboundRealTimePaymentsTransferConfirmation != nil {
		InboundRealTimePaymentsTransferConfirmation = *r.InboundRealTimePaymentsTransferConfirmation
	}
	return
}

// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
// the JSON response if and only if `category` is equal to
// `inbound_wire_drawdown_payment_reversal`.
func (r *TransactionSource) GetInboundWireDrawdownPaymentReversal() (InboundWireDrawdownPaymentReversal TransactionSourceInboundWireDrawdownPaymentReversal) {
	if r != nil && r.InboundWireDrawdownPaymentReversal != nil {
		InboundWireDrawdownPaymentReversal = *r.InboundWireDrawdownPaymentReversal
	}
	return
}

// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
func (r *TransactionSource) GetInboundWireDrawdownPayment() (InboundWireDrawdownPayment TransactionSourceInboundWireDrawdownPayment) {
	if r != nil && r.InboundWireDrawdownPayment != nil {
		InboundWireDrawdownPayment = *r.InboundWireDrawdownPayment
	}
	return
}

// A Inbound Wire Reversal object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_reversal`.
func (r *TransactionSource) GetInboundWireReversal() (InboundWireReversal TransactionSourceInboundWireReversal) {
	if r != nil && r.InboundWireReversal != nil {
		InboundWireReversal = *r.InboundWireReversal
	}
	return
}

// A Inbound Wire Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_transfer`.
func (r *TransactionSource) GetInboundWireTransfer() (InboundWireTransfer TransactionSourceInboundWireTransfer) {
	if r != nil && r.InboundWireTransfer != nil {
		InboundWireTransfer = *r.InboundWireTransfer
	}
	return
}

// A Internal Source object. This field will be present in the JSON response if and
// only if `category` is equal to `internal_source`.
func (r *TransactionSource) GetInternalSource() (InternalSource TransactionSourceInternalSource) {
	if r != nil && r.InternalSource != nil {
		InternalSource = *r.InternalSource
	}
	return
}

// A Deprecated Card Refund object. This field will be present in the JSON response
// if and only if `category` is equal to `card_route_refund`.
func (r *TransactionSource) GetCardRouteRefund() (CardRouteRefund TransactionSourceCardRouteRefund) {
	if r != nil && r.CardRouteRefund != nil {
		CardRouteRefund = *r.CardRouteRefund
	}
	return
}

// A Deprecated Card Settlement object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_settlement`.
func (r *TransactionSource) GetCardRouteSettlement() (CardRouteSettlement TransactionSourceCardRouteSettlement) {
	if r != nil && r.CardRouteSettlement != nil {
		CardRouteSettlement = *r.CardRouteSettlement
	}
	return
}

// A Sample Funds object. This field will be present in the JSON response if and
// only if `category` is equal to `sample_funds`.
func (r *TransactionSource) GetSampleFunds() (SampleFunds TransactionSourceSampleFunds) {
	if r != nil && r.SampleFunds != nil {
		SampleFunds = *r.SampleFunds
	}
	return
}

// A Wire Drawdown Payment Intention object. This field will be present in the JSON
// response if and only if `category` is equal to
// `wire_drawdown_payment_intention`.
func (r *TransactionSource) GetWireDrawdownPaymentIntention() (WireDrawdownPaymentIntention TransactionSourceWireDrawdownPaymentIntention) {
	if r != nil && r.WireDrawdownPaymentIntention != nil {
		WireDrawdownPaymentIntention = *r.WireDrawdownPaymentIntention
	}
	return
}

// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to
// `wire_drawdown_payment_rejection`.
func (r *TransactionSource) GetWireDrawdownPaymentRejection() (WireDrawdownPaymentRejection TransactionSourceWireDrawdownPaymentRejection) {
	if r != nil && r.WireDrawdownPaymentRejection != nil {
		WireDrawdownPaymentRejection = *r.WireDrawdownPaymentRejection
	}
	return
}

// A Wire Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_intention`.
func (r *TransactionSource) GetWireTransferIntention() (WireTransferIntention TransactionSourceWireTransferIntention) {
	if r != nil && r.WireTransferIntention != nil {
		WireTransferIntention = *r.WireTransferIntention
	}
	return
}

// A Wire Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_rejection`.
func (r *TransactionSource) GetWireTransferRejection() (WireTransferRejection TransactionSourceWireTransferRejection) {
	if r != nil && r.WireTransferRejection != nil {
		WireTransferRejection = *r.WireTransferRejection
	}
	return
}

func (r TransactionSource) String() (result string) {
	return fmt.Sprintf("&TransactionSource{Category:%s AccountTransferIntention:%s ACHCheckConversionReturn:%s ACHCheckConversion:%s ACHTransferIntention:%s ACHTransferRejection:%s ACHTransferReturn:%s CardDisputeAcceptance:%s CardRefund:%s CardSettlement:%s CheckDepositAcceptance:%s CheckDepositReturn:%s CheckTransferIntention:%s CheckTransferReturn:%s CheckTransferRejection:%s CheckTransferStopPaymentRequest:%s DisputeResolution:%s EmpyrealCashDeposit:%s InboundACHTransfer:%s InboundCheck:%s InboundInternationalACHTransfer:%s InboundRealTimePaymentsTransferConfirmation:%s InboundWireDrawdownPaymentReversal:%s InboundWireDrawdownPayment:%s InboundWireReversal:%s InboundWireTransfer:%s InternalSource:%s CardRouteRefund:%s CardRouteSettlement:%s SampleFunds:%s WireDrawdownPaymentIntention:%s WireDrawdownPaymentRejection:%s WireTransferIntention:%s WireTransferRejection:%s}", core.FmtP(r.Category), r.AccountTransferIntention, r.ACHCheckConversionReturn, r.ACHCheckConversion, r.ACHTransferIntention, r.ACHTransferRejection, r.ACHTransferReturn, r.CardDisputeAcceptance, r.CardRefund, r.CardSettlement, r.CheckDepositAcceptance, r.CheckDepositReturn, r.CheckTransferIntention, r.CheckTransferReturn, r.CheckTransferRejection, r.CheckTransferStopPaymentRequest, r.DisputeResolution, r.EmpyrealCashDeposit, r.InboundACHTransfer, r.InboundCheck, r.InboundInternationalACHTransfer, r.InboundRealTimePaymentsTransferConfirmation, r.InboundWireDrawdownPaymentReversal, r.InboundWireDrawdownPayment, r.InboundWireReversal, r.InboundWireTransfer, r.InternalSource, r.CardRouteRefund, r.CardRouteSettlement, r.SampleFunds, r.WireDrawdownPaymentIntention, r.WireDrawdownPaymentRejection, r.WireTransferIntention, r.WireTransferRejection)
}

type TransactionSourceCategory string

const (
	TransactionSourceCategoryAccountTransferIntention                    TransactionSourceCategory = "account_transfer_intention"
	TransactionSourceCategoryACHCheckConversionReturn                    TransactionSourceCategory = "ach_check_conversion_return"
	TransactionSourceCategoryACHCheckConversion                          TransactionSourceCategory = "ach_check_conversion"
	TransactionSourceCategoryACHTransferIntention                        TransactionSourceCategory = "ach_transfer_intention"
	TransactionSourceCategoryACHTransferRejection                        TransactionSourceCategory = "ach_transfer_rejection"
	TransactionSourceCategoryACHTransferReturn                           TransactionSourceCategory = "ach_transfer_return"
	TransactionSourceCategoryCardDisputeAcceptance                       TransactionSourceCategory = "card_dispute_acceptance"
	TransactionSourceCategoryCardRefund                                  TransactionSourceCategory = "card_refund"
	TransactionSourceCategoryCardSettlement                              TransactionSourceCategory = "card_settlement"
	TransactionSourceCategoryCheckDepositAcceptance                      TransactionSourceCategory = "check_deposit_acceptance"
	TransactionSourceCategoryCheckDepositReturn                          TransactionSourceCategory = "check_deposit_return"
	TransactionSourceCategoryCheckTransferIntention                      TransactionSourceCategory = "check_transfer_intention"
	TransactionSourceCategoryCheckTransferReturn                         TransactionSourceCategory = "check_transfer_return"
	TransactionSourceCategoryCheckTransferRejection                      TransactionSourceCategory = "check_transfer_rejection"
	TransactionSourceCategoryCheckTransferStopPaymentRequest             TransactionSourceCategory = "check_transfer_stop_payment_request"
	TransactionSourceCategoryDisputeResolution                           TransactionSourceCategory = "dispute_resolution"
	TransactionSourceCategoryEmpyrealCashDeposit                         TransactionSourceCategory = "empyreal_cash_deposit"
	TransactionSourceCategoryInboundACHTransfer                          TransactionSourceCategory = "inbound_ach_transfer"
	TransactionSourceCategoryInboundCheck                                TransactionSourceCategory = "inbound_check"
	TransactionSourceCategoryInboundInternationalACHTransfer             TransactionSourceCategory = "inbound_international_ach_transfer"
	TransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation TransactionSourceCategory = "inbound_real_time_payments_transfer_confirmation"
	TransactionSourceCategoryInboundWireDrawdownPaymentReversal          TransactionSourceCategory = "inbound_wire_drawdown_payment_reversal"
	TransactionSourceCategoryInboundWireDrawdownPayment                  TransactionSourceCategory = "inbound_wire_drawdown_payment"
	TransactionSourceCategoryInboundWireReversal                         TransactionSourceCategory = "inbound_wire_reversal"
	TransactionSourceCategoryInboundWireTransfer                         TransactionSourceCategory = "inbound_wire_transfer"
	TransactionSourceCategoryInternalGeneralLedgerTransaction            TransactionSourceCategory = "internal_general_ledger_transaction"
	TransactionSourceCategoryInternalSource                              TransactionSourceCategory = "internal_source"
	TransactionSourceCategoryCardRouteRefund                             TransactionSourceCategory = "card_route_refund"
	TransactionSourceCategoryCardRouteSettlement                         TransactionSourceCategory = "card_route_settlement"
	TransactionSourceCategoryRealTimePaymentsTransferAcknowledgement     TransactionSourceCategory = "real_time_payments_transfer_acknowledgement"
	TransactionSourceCategorySampleFunds                                 TransactionSourceCategory = "sample_funds"
	TransactionSourceCategoryWireDrawdownPaymentIntention                TransactionSourceCategory = "wire_drawdown_payment_intention"
	TransactionSourceCategoryWireDrawdownPaymentRejection                TransactionSourceCategory = "wire_drawdown_payment_rejection"
	TransactionSourceCategoryWireTransferIntention                       TransactionSourceCategory = "wire_transfer_intention"
	TransactionSourceCategoryWireTransferRejection                       TransactionSourceCategory = "wire_transfer_rejection"
	TransactionSourceCategoryOther                                       TransactionSourceCategory = "other"
)

type TransactionSourceAccountTransferIntention struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *TransactionSourceAccountTransferIntentionCurrency `pjson:"currency"`
	// The description you chose to give the transfer.
	Description *string `pjson:"description"`
	// The identifier of the Account to where the Account Transfer was sent.
	DestinationAccountID *string `pjson:"destination_account_id"`
	// The identifier of the Account from where the Account Transfer was sent.
	SourceAccountID *string `pjson:"source_account_id"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceAccountTransferIntention using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceAccountTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceAccountTransferIntention into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *TransactionSourceAccountTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *TransactionSourceAccountTransferIntention) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *TransactionSourceAccountTransferIntention) GetCurrency() (Currency TransactionSourceAccountTransferIntentionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The description you chose to give the transfer.
func (r *TransactionSourceAccountTransferIntention) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The identifier of the Account to where the Account Transfer was sent.
func (r *TransactionSourceAccountTransferIntention) GetDestinationAccountID() (DestinationAccountID string) {
	if r != nil && r.DestinationAccountID != nil {
		DestinationAccountID = *r.DestinationAccountID
	}
	return
}

// The identifier of the Account from where the Account Transfer was sent.
func (r *TransactionSourceAccountTransferIntention) GetSourceAccountID() (SourceAccountID string) {
	if r != nil && r.SourceAccountID != nil {
		SourceAccountID = *r.SourceAccountID
	}
	return
}

// The identifier of the Account Transfer that led to this Pending Transaction.
func (r *TransactionSourceAccountTransferIntention) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r TransactionSourceAccountTransferIntention) String() (result string) {
	return fmt.Sprintf("&TransactionSourceAccountTransferIntention{Amount:%s Currency:%s Description:%s DestinationAccountID:%s SourceAccountID:%s TransferID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.Description), core.FmtP(r.DestinationAccountID), core.FmtP(r.SourceAccountID), core.FmtP(r.TransferID))
}

type TransactionSourceAccountTransferIntentionCurrency string

const (
	TransactionSourceAccountTransferIntentionCurrencyCad TransactionSourceAccountTransferIntentionCurrency = "CAD"
	TransactionSourceAccountTransferIntentionCurrencyChf TransactionSourceAccountTransferIntentionCurrency = "CHF"
	TransactionSourceAccountTransferIntentionCurrencyEur TransactionSourceAccountTransferIntentionCurrency = "EUR"
	TransactionSourceAccountTransferIntentionCurrencyGbp TransactionSourceAccountTransferIntentionCurrency = "GBP"
	TransactionSourceAccountTransferIntentionCurrencyJpy TransactionSourceAccountTransferIntentionCurrency = "JPY"
	TransactionSourceAccountTransferIntentionCurrencyUsd TransactionSourceAccountTransferIntentionCurrency = "USD"
)

type TransactionSourceACHCheckConversionReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// Why the transfer was returned.
	ReturnReasonCode *string                `pjson:"return_reason_code"`
	jsonFields       map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceACHCheckConversionReturn using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceACHCheckConversionReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceACHCheckConversionReturn into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *TransactionSourceACHCheckConversionReturn) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceACHCheckConversionReturn) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// Why the transfer was returned.
func (r *TransactionSourceACHCheckConversionReturn) GetReturnReasonCode() (ReturnReasonCode string) {
	if r != nil && r.ReturnReasonCode != nil {
		ReturnReasonCode = *r.ReturnReasonCode
	}
	return
}

func (r TransactionSourceACHCheckConversionReturn) String() (result string) {
	return fmt.Sprintf("&TransactionSourceACHCheckConversionReturn{Amount:%s ReturnReasonCode:%s}", core.FmtP(r.Amount), core.FmtP(r.ReturnReasonCode))
}

type TransactionSourceACHCheckConversion struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The identifier of the File containing an image of the returned check.
	FileID     *string                `pjson:"file_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceACHCheckConversion using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceACHCheckConversion) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceACHCheckConversion into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceACHCheckConversion) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceACHCheckConversion) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The identifier of the File containing an image of the returned check.
func (r *TransactionSourceACHCheckConversion) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

func (r TransactionSourceACHCheckConversion) String() (result string) {
	return fmt.Sprintf("&TransactionSourceACHCheckConversion{Amount:%s FileID:%s}", core.FmtP(r.Amount), core.FmtP(r.FileID))
}

type TransactionSourceACHTransferIntention struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount              *int64  `pjson:"amount"`
	AccountNumber       *string `pjson:"account_number"`
	RoutingNumber       *string `pjson:"routing_number"`
	StatementDescriptor *string `pjson:"statement_descriptor"`
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceACHTransferIntention using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceACHTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceACHTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceACHTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceACHTransferIntention) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *TransactionSourceACHTransferIntention) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *TransactionSourceACHTransferIntention) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *TransactionSourceACHTransferIntention) GetStatementDescriptor() (StatementDescriptor string) {
	if r != nil && r.StatementDescriptor != nil {
		StatementDescriptor = *r.StatementDescriptor
	}
	return
}

// The identifier of the ACH Transfer that led to this Transaction.
func (r *TransactionSourceACHTransferIntention) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r TransactionSourceACHTransferIntention) String() (result string) {
	return fmt.Sprintf("&TransactionSourceACHTransferIntention{Amount:%s AccountNumber:%s RoutingNumber:%s StatementDescriptor:%s TransferID:%s}", core.FmtP(r.Amount), core.FmtP(r.AccountNumber), core.FmtP(r.RoutingNumber), core.FmtP(r.StatementDescriptor), core.FmtP(r.TransferID))
}

type TransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceACHTransferRejection using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceACHTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceACHTransferRejection into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceACHTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the ACH Transfer that led to this Transaction.
func (r *TransactionSourceACHTransferRejection) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r TransactionSourceACHTransferRejection) String() (result string) {
	return fmt.Sprintf("&TransactionSourceACHTransferRejection{TransferID:%s}", core.FmtP(r.TransferID))
}

type TransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `pjson:"created_at"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode *TransactionSourceACHTransferReturnReturnReasonCode `pjson:"return_reason_code"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID *string `pjson:"transfer_id"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID *string                `pjson:"transaction_id"`
	jsonFields    map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceACHTransferReturn using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceACHTransferReturn into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceACHTransferReturn) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r *TransactionSourceACHTransferReturn) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// Why the ACH Transfer was returned.
func (r *TransactionSourceACHTransferReturn) GetReturnReasonCode() (ReturnReasonCode TransactionSourceACHTransferReturnReturnReasonCode) {
	if r != nil && r.ReturnReasonCode != nil {
		ReturnReasonCode = *r.ReturnReasonCode
	}
	return
}

// The identifier of the ACH Transfer associated with this return.
func (r *TransactionSourceACHTransferReturn) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

// The identifier of the Tranasaction associated with this return.
func (r *TransactionSourceACHTransferReturn) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

func (r TransactionSourceACHTransferReturn) String() (result string) {
	return fmt.Sprintf("&TransactionSourceACHTransferReturn{CreatedAt:%s ReturnReasonCode:%s TransferID:%s TransactionID:%s}", core.FmtP(r.CreatedAt), core.FmtP(r.ReturnReasonCode), core.FmtP(r.TransferID), core.FmtP(r.TransactionID))
}

type TransactionSourceACHTransferReturnReturnReasonCode string

const (
	TransactionSourceACHTransferReturnReturnReasonCodeInsufficientFund                                          TransactionSourceACHTransferReturnReturnReasonCode = "insufficient_fund"
	TransactionSourceACHTransferReturnReturnReasonCodeNoAccount                                                 TransactionSourceACHTransferReturnReturnReasonCode = "no_account"
	TransactionSourceACHTransferReturnReturnReasonCodeAccountClosed                                             TransactionSourceACHTransferReturnReturnReasonCode = "account_closed"
	TransactionSourceACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure                             TransactionSourceACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	TransactionSourceACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction              TransactionSourceACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	TransactionSourceACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver                              TransactionSourceACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	TransactionSourceACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode   TransactionSourceACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	TransactionSourceACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized                     TransactionSourceACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	TransactionSourceACHTransferReturnReturnReasonCodePaymentStopped                                            TransactionSourceACHTransferReturnReturnReasonCode = "payment_stopped"
	TransactionSourceACHTransferReturnReturnReasonCodeNonTransactionAccount                                     TransactionSourceACHTransferReturnReturnReasonCode = "non_transaction_account"
	TransactionSourceACHTransferReturnReturnReasonCodeUncollectedFunds                                          TransactionSourceACHTransferReturnReturnReasonCode = "uncollected_funds"
	TransactionSourceACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError                              TransactionSourceACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	TransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete TransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	TransactionSourceACHTransferReturnReturnReasonCodeAmountFieldError                                          TransactionSourceACHTransferReturnReturnReasonCode = "amount_field_error"
	TransactionSourceACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer                            TransactionSourceACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	TransactionSourceACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber                                   TransactionSourceACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	TransactionSourceACHTransferReturnReturnReasonCodeFileRecordEditCriteria                                    TransactionSourceACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	TransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualName                                  TransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	TransactionSourceACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest                                    TransactionSourceACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	TransactionSourceACHTransferReturnReturnReasonCodeAddendaError                                              TransactionSourceACHTransferReturnReturnReasonCode = "addenda_error"
	TransactionSourceACHTransferReturnReturnReasonCodeLimitedParticipationDfi                                   TransactionSourceACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	TransactionSourceACHTransferReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment              TransactionSourceACHTransferReturnReturnReasonCode = "incorrectly_coded_outbound_international_payment"
	TransactionSourceACHTransferReturnReturnReasonCodeOther                                                     TransactionSourceACHTransferReturnReturnReasonCode = "other"
)

type TransactionSourceCardDisputeAcceptance struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was accepted.
	AcceptedAt *string `pjson:"accepted_at"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID *string `pjson:"card_dispute_id"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID *string                `pjson:"transaction_id"`
	jsonFields    map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceCardDisputeAcceptance using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceCardDisputeAcceptance) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceCardDisputeAcceptance into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceCardDisputeAcceptance) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card Dispute was accepted.
func (r *TransactionSourceCardDisputeAcceptance) GetAcceptedAt() (AcceptedAt string) {
	if r != nil && r.AcceptedAt != nil {
		AcceptedAt = *r.AcceptedAt
	}
	return
}

// The identifier of the Card Dispute that was accepted.
func (r *TransactionSourceCardDisputeAcceptance) GetCardDisputeID() (CardDisputeID string) {
	if r != nil && r.CardDisputeID != nil {
		CardDisputeID = *r.CardDisputeID
	}
	return
}

// The identifier of the Transaction that was created to return the disputed funds
// to your account.
func (r *TransactionSourceCardDisputeAcceptance) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

func (r TransactionSourceCardDisputeAcceptance) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCardDisputeAcceptance{AcceptedAt:%s CardDisputeID:%s TransactionID:%s}", core.FmtP(r.AcceptedAt), core.FmtP(r.CardDisputeID), core.FmtP(r.TransactionID))
}

type TransactionSourceCardRefund struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *TransactionSourceCardRefundCurrency `pjson:"currency"`
	// The identifier for the Transaction this refunds, if any.
	CardSettlementTransactionID *string `pjson:"card_settlement_transaction_id"`
	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type       *TransactionSourceCardRefundType `pjson:"type"`
	jsonFields map[string]interface{}           `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into TransactionSourceCardRefund
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *TransactionSourceCardRefund) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceCardRefund into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *TransactionSourceCardRefund) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *TransactionSourceCardRefund) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *TransactionSourceCardRefund) GetCurrency() (Currency TransactionSourceCardRefundCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier for the Transaction this refunds, if any.
func (r *TransactionSourceCardRefund) GetCardSettlementTransactionID() (CardSettlementTransactionID string) {
	if r != nil && r.CardSettlementTransactionID != nil {
		CardSettlementTransactionID = *r.CardSettlementTransactionID
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_refund`.
func (r *TransactionSourceCardRefund) GetType() (Type TransactionSourceCardRefundType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r TransactionSourceCardRefund) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCardRefund{Amount:%s Currency:%s CardSettlementTransactionID:%s Type:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.CardSettlementTransactionID), core.FmtP(r.Type))
}

type TransactionSourceCardRefundCurrency string

const (
	TransactionSourceCardRefundCurrencyCad TransactionSourceCardRefundCurrency = "CAD"
	TransactionSourceCardRefundCurrencyChf TransactionSourceCardRefundCurrency = "CHF"
	TransactionSourceCardRefundCurrencyEur TransactionSourceCardRefundCurrency = "EUR"
	TransactionSourceCardRefundCurrencyGbp TransactionSourceCardRefundCurrency = "GBP"
	TransactionSourceCardRefundCurrencyJpy TransactionSourceCardRefundCurrency = "JPY"
	TransactionSourceCardRefundCurrencyUsd TransactionSourceCardRefundCurrency = "USD"
)

type TransactionSourceCardRefundType string

const (
	TransactionSourceCardRefundTypeCardRefund TransactionSourceCardRefundType = "card_refund"
)

type TransactionSourceCardSettlement struct {
	// The amount in the minor unit of the transaction's settlement currency. For
	// dollars, for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's settlement currency.
	Currency *TransactionSourceCardSettlementCurrency `pjson:"currency"`
	// The amount in the minor unit of the transaction's presentment currency.
	PresentmentAmount *int64 `pjson:"presentment_amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's presentment currency.
	PresentmentCurrency  *string `pjson:"presentment_currency"`
	MerchantCity         *string `pjson:"merchant_city"`
	MerchantCountry      *string `pjson:"merchant_country"`
	MerchantName         *string `pjson:"merchant_name"`
	MerchantCategoryCode *string `pjson:"merchant_category_code"`
	MerchantState        *string `pjson:"merchant_state"`
	// The identifier of the Pending Transaction associated with this Transaction.
	PendingTransactionID *string `pjson:"pending_transaction_id"`
	// A constant representing the object's type. For this resource it will always be
	// `card_settlement`.
	Type       *TransactionSourceCardSettlementType `pjson:"type"`
	jsonFields map[string]interface{}               `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceCardSettlement using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *TransactionSourceCardSettlement) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceCardSettlement into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceCardSettlement) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's settlement currency. For
// dollars, for example, this is cents.
func (r *TransactionSourceCardSettlement) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's settlement currency.
func (r *TransactionSourceCardSettlement) GetCurrency() (Currency TransactionSourceCardSettlementCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The amount in the minor unit of the transaction's presentment currency.
func (r *TransactionSourceCardSettlement) GetPresentmentAmount() (PresentmentAmount int64) {
	if r != nil && r.PresentmentAmount != nil {
		PresentmentAmount = *r.PresentmentAmount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's presentment currency.
func (r *TransactionSourceCardSettlement) GetPresentmentCurrency() (PresentmentCurrency string) {
	if r != nil && r.PresentmentCurrency != nil {
		PresentmentCurrency = *r.PresentmentCurrency
	}
	return
}

func (r *TransactionSourceCardSettlement) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *TransactionSourceCardSettlement) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *TransactionSourceCardSettlement) GetMerchantName() (MerchantName string) {
	if r != nil && r.MerchantName != nil {
		MerchantName = *r.MerchantName
	}
	return
}

func (r *TransactionSourceCardSettlement) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

func (r *TransactionSourceCardSettlement) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

// The identifier of the Pending Transaction associated with this Transaction.
func (r *TransactionSourceCardSettlement) GetPendingTransactionID() (PendingTransactionID string) {
	if r != nil && r.PendingTransactionID != nil {
		PendingTransactionID = *r.PendingTransactionID
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_settlement`.
func (r *TransactionSourceCardSettlement) GetType() (Type TransactionSourceCardSettlementType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r TransactionSourceCardSettlement) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCardSettlement{Amount:%s Currency:%s PresentmentAmount:%s PresentmentCurrency:%s MerchantCity:%s MerchantCountry:%s MerchantName:%s MerchantCategoryCode:%s MerchantState:%s PendingTransactionID:%s Type:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.PresentmentAmount), core.FmtP(r.PresentmentCurrency), core.FmtP(r.MerchantCity), core.FmtP(r.MerchantCountry), core.FmtP(r.MerchantName), core.FmtP(r.MerchantCategoryCode), core.FmtP(r.MerchantState), core.FmtP(r.PendingTransactionID), core.FmtP(r.Type))
}

type TransactionSourceCardSettlementCurrency string

const (
	TransactionSourceCardSettlementCurrencyCad TransactionSourceCardSettlementCurrency = "CAD"
	TransactionSourceCardSettlementCurrencyChf TransactionSourceCardSettlementCurrency = "CHF"
	TransactionSourceCardSettlementCurrencyEur TransactionSourceCardSettlementCurrency = "EUR"
	TransactionSourceCardSettlementCurrencyGbp TransactionSourceCardSettlementCurrency = "GBP"
	TransactionSourceCardSettlementCurrencyJpy TransactionSourceCardSettlementCurrency = "JPY"
	TransactionSourceCardSettlementCurrencyUsd TransactionSourceCardSettlementCurrency = "USD"
)

type TransactionSourceCardSettlementType string

const (
	TransactionSourceCardSettlementTypeCardSettlement TransactionSourceCardSettlementType = "card_settlement"
)

type TransactionSourceCheckDepositAcceptance struct {
	// The amount to be deposited in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *TransactionSourceCheckDepositAcceptanceCurrency `pjson:"currency"`
	// The account number printed on the check.
	AccountNumber *string `pjson:"account_number"`
	// The routing number printed on the check.
	RoutingNumber *string `pjson:"routing_number"`
	// An additional line of metadata printed on the check. This typically includes the
	// check number for business checks.
	AuxiliaryOnUs *string `pjson:"auxiliary_on_us"`
	// The check serial number, if present, for consumer checks. For business checks,
	// the serial number is usually in the `auxiliary_on_us` field.
	SerialNumber *string `pjson:"serial_number"`
	// The ID of the Check Deposit that was accepted.
	CheckDepositID *string                `pjson:"check_deposit_id"`
	jsonFields     map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceCheckDepositAcceptance using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceCheckDepositAcceptance) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceCheckDepositAcceptance into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceCheckDepositAcceptance) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount to be deposited in the minor unit of the transaction's currency. For
// dollars, for example, this is cents.
func (r *TransactionSourceCheckDepositAcceptance) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *TransactionSourceCheckDepositAcceptance) GetCurrency() (Currency TransactionSourceCheckDepositAcceptanceCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The account number printed on the check.
func (r *TransactionSourceCheckDepositAcceptance) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The routing number printed on the check.
func (r *TransactionSourceCheckDepositAcceptance) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// An additional line of metadata printed on the check. This typically includes the
// check number for business checks.
func (r *TransactionSourceCheckDepositAcceptance) GetAuxiliaryOnUs() (AuxiliaryOnUs string) {
	if r != nil && r.AuxiliaryOnUs != nil {
		AuxiliaryOnUs = *r.AuxiliaryOnUs
	}
	return
}

// The check serial number, if present, for consumer checks. For business checks,
// the serial number is usually in the `auxiliary_on_us` field.
func (r *TransactionSourceCheckDepositAcceptance) GetSerialNumber() (SerialNumber string) {
	if r != nil && r.SerialNumber != nil {
		SerialNumber = *r.SerialNumber
	}
	return
}

// The ID of the Check Deposit that was accepted.
func (r *TransactionSourceCheckDepositAcceptance) GetCheckDepositID() (CheckDepositID string) {
	if r != nil && r.CheckDepositID != nil {
		CheckDepositID = *r.CheckDepositID
	}
	return
}

func (r TransactionSourceCheckDepositAcceptance) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCheckDepositAcceptance{Amount:%s Currency:%s AccountNumber:%s RoutingNumber:%s AuxiliaryOnUs:%s SerialNumber:%s CheckDepositID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.AccountNumber), core.FmtP(r.RoutingNumber), core.FmtP(r.AuxiliaryOnUs), core.FmtP(r.SerialNumber), core.FmtP(r.CheckDepositID))
}

type TransactionSourceCheckDepositAcceptanceCurrency string

const (
	TransactionSourceCheckDepositAcceptanceCurrencyCad TransactionSourceCheckDepositAcceptanceCurrency = "CAD"
	TransactionSourceCheckDepositAcceptanceCurrencyChf TransactionSourceCheckDepositAcceptanceCurrency = "CHF"
	TransactionSourceCheckDepositAcceptanceCurrencyEur TransactionSourceCheckDepositAcceptanceCurrency = "EUR"
	TransactionSourceCheckDepositAcceptanceCurrencyGbp TransactionSourceCheckDepositAcceptanceCurrency = "GBP"
	TransactionSourceCheckDepositAcceptanceCurrencyJpy TransactionSourceCheckDepositAcceptanceCurrency = "JPY"
	TransactionSourceCheckDepositAcceptanceCurrencyUsd TransactionSourceCheckDepositAcceptanceCurrency = "USD"
)

type TransactionSourceCheckDepositReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt *string `pjson:"returned_at"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *TransactionSourceCheckDepositReturnCurrency `pjson:"currency"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID *string `pjson:"check_deposit_id"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID *string                                          `pjson:"transaction_id"`
	ReturnReason  *TransactionSourceCheckDepositReturnReturnReason `pjson:"return_reason"`
	jsonFields    map[string]interface{}                           `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceCheckDepositReturn using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceCheckDepositReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceCheckDepositReturn into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceCheckDepositReturn) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceCheckDepositReturn) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the check deposit was returned.
func (r *TransactionSourceCheckDepositReturn) GetReturnedAt() (ReturnedAt string) {
	if r != nil && r.ReturnedAt != nil {
		ReturnedAt = *r.ReturnedAt
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *TransactionSourceCheckDepositReturn) GetCurrency() (Currency TransactionSourceCheckDepositReturnCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Check Deposit that was returned.
func (r *TransactionSourceCheckDepositReturn) GetCheckDepositID() (CheckDepositID string) {
	if r != nil && r.CheckDepositID != nil {
		CheckDepositID = *r.CheckDepositID
	}
	return
}

// The identifier of the transaction that reversed the original check deposit
// transaction.
func (r *TransactionSourceCheckDepositReturn) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

func (r *TransactionSourceCheckDepositReturn) GetReturnReason() (ReturnReason TransactionSourceCheckDepositReturnReturnReason) {
	if r != nil && r.ReturnReason != nil {
		ReturnReason = *r.ReturnReason
	}
	return
}

func (r TransactionSourceCheckDepositReturn) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCheckDepositReturn{Amount:%s ReturnedAt:%s Currency:%s CheckDepositID:%s TransactionID:%s ReturnReason:%s}", core.FmtP(r.Amount), core.FmtP(r.ReturnedAt), core.FmtP(r.Currency), core.FmtP(r.CheckDepositID), core.FmtP(r.TransactionID), core.FmtP(r.ReturnReason))
}

type TransactionSourceCheckDepositReturnCurrency string

const (
	TransactionSourceCheckDepositReturnCurrencyCad TransactionSourceCheckDepositReturnCurrency = "CAD"
	TransactionSourceCheckDepositReturnCurrencyChf TransactionSourceCheckDepositReturnCurrency = "CHF"
	TransactionSourceCheckDepositReturnCurrencyEur TransactionSourceCheckDepositReturnCurrency = "EUR"
	TransactionSourceCheckDepositReturnCurrencyGbp TransactionSourceCheckDepositReturnCurrency = "GBP"
	TransactionSourceCheckDepositReturnCurrencyJpy TransactionSourceCheckDepositReturnCurrency = "JPY"
	TransactionSourceCheckDepositReturnCurrencyUsd TransactionSourceCheckDepositReturnCurrency = "USD"
)

type TransactionSourceCheckDepositReturnReturnReason string

const (
	TransactionSourceCheckDepositReturnReturnReasonACHConversionNotSupported TransactionSourceCheckDepositReturnReturnReason = "ach_conversion_not_supported"
	TransactionSourceCheckDepositReturnReturnReasonClosedAccount             TransactionSourceCheckDepositReturnReturnReason = "closed_account"
	TransactionSourceCheckDepositReturnReturnReasonDuplicateSubmission       TransactionSourceCheckDepositReturnReturnReason = "duplicate_submission"
	TransactionSourceCheckDepositReturnReturnReasonInsufficientFunds         TransactionSourceCheckDepositReturnReturnReason = "insufficient_funds"
	TransactionSourceCheckDepositReturnReturnReasonNoAccount                 TransactionSourceCheckDepositReturnReturnReason = "no_account"
	TransactionSourceCheckDepositReturnReturnReasonNotAuthorized             TransactionSourceCheckDepositReturnReturnReason = "not_authorized"
	TransactionSourceCheckDepositReturnReturnReasonStaleDated                TransactionSourceCheckDepositReturnReturnReason = "stale_dated"
	TransactionSourceCheckDepositReturnReturnReasonStopPayment               TransactionSourceCheckDepositReturnReturnReason = "stop_payment"
	TransactionSourceCheckDepositReturnReturnReasonUnknownReason             TransactionSourceCheckDepositReturnReturnReason = "unknown_reason"
	TransactionSourceCheckDepositReturnReturnReasonUnmatchedDetails          TransactionSourceCheckDepositReturnReturnReason = "unmatched_details"
	TransactionSourceCheckDepositReturnReturnReasonUnreadableImage           TransactionSourceCheckDepositReturnReturnReason = "unreadable_image"
)

type TransactionSourceCheckTransferIntention struct {
	// The street address of the check's destination.
	AddressLine1 *string `pjson:"address_line1"`
	// The second line of the address of the check's destination.
	AddressLine2 *string `pjson:"address_line2"`
	// The city of the check's destination.
	AddressCity *string `pjson:"address_city"`
	// The state of the check's destination.
	AddressState *string `pjson:"address_state"`
	// The postal code of the check's destination.
	AddressZip *string `pjson:"address_zip"`
	// The transfer amount in USD cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency *TransactionSourceCheckTransferIntentionCurrency `pjson:"currency"`
	// The name that will be printed on the check.
	RecipientName *string `pjson:"recipient_name"`
	// The identifier of the Check Transfer with which this is associated.
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceCheckTransferIntention using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceCheckTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceCheckTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceCheckTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The street address of the check's destination.
func (r *TransactionSourceCheckTransferIntention) GetAddressLine1() (AddressLine1 string) {
	if r != nil && r.AddressLine1 != nil {
		AddressLine1 = *r.AddressLine1
	}
	return
}

// The second line of the address of the check's destination.
func (r *TransactionSourceCheckTransferIntention) GetAddressLine2() (AddressLine2 string) {
	if r != nil && r.AddressLine2 != nil {
		AddressLine2 = *r.AddressLine2
	}
	return
}

// The city of the check's destination.
func (r *TransactionSourceCheckTransferIntention) GetAddressCity() (AddressCity string) {
	if r != nil && r.AddressCity != nil {
		AddressCity = *r.AddressCity
	}
	return
}

// The state of the check's destination.
func (r *TransactionSourceCheckTransferIntention) GetAddressState() (AddressState string) {
	if r != nil && r.AddressState != nil {
		AddressState = *r.AddressState
	}
	return
}

// The postal code of the check's destination.
func (r *TransactionSourceCheckTransferIntention) GetAddressZip() (AddressZip string) {
	if r != nil && r.AddressZip != nil {
		AddressZip = *r.AddressZip
	}
	return
}

// The transfer amount in USD cents.
func (r *TransactionSourceCheckTransferIntention) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
func (r *TransactionSourceCheckTransferIntention) GetCurrency() (Currency TransactionSourceCheckTransferIntentionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The name that will be printed on the check.
func (r *TransactionSourceCheckTransferIntention) GetRecipientName() (RecipientName string) {
	if r != nil && r.RecipientName != nil {
		RecipientName = *r.RecipientName
	}
	return
}

// The identifier of the Check Transfer with which this is associated.
func (r *TransactionSourceCheckTransferIntention) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r TransactionSourceCheckTransferIntention) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCheckTransferIntention{AddressLine1:%s AddressLine2:%s AddressCity:%s AddressState:%s AddressZip:%s Amount:%s Currency:%s RecipientName:%s TransferID:%s}", core.FmtP(r.AddressLine1), core.FmtP(r.AddressLine2), core.FmtP(r.AddressCity), core.FmtP(r.AddressState), core.FmtP(r.AddressZip), core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.RecipientName), core.FmtP(r.TransferID))
}

type TransactionSourceCheckTransferIntentionCurrency string

const (
	TransactionSourceCheckTransferIntentionCurrencyCad TransactionSourceCheckTransferIntentionCurrency = "CAD"
	TransactionSourceCheckTransferIntentionCurrencyChf TransactionSourceCheckTransferIntentionCurrency = "CHF"
	TransactionSourceCheckTransferIntentionCurrencyEur TransactionSourceCheckTransferIntentionCurrency = "EUR"
	TransactionSourceCheckTransferIntentionCurrencyGbp TransactionSourceCheckTransferIntentionCurrency = "GBP"
	TransactionSourceCheckTransferIntentionCurrencyJpy TransactionSourceCheckTransferIntentionCurrency = "JPY"
	TransactionSourceCheckTransferIntentionCurrencyUsd TransactionSourceCheckTransferIntentionCurrency = "USD"
)

type TransactionSourceCheckTransferReturn struct {
	// The identifier of the returned Check Transfer.
	TransferID *string `pjson:"transfer_id"`
	// If available, a document with additional information about the return.
	FileID     *string                `pjson:"file_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceCheckTransferReturn using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceCheckTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceCheckTransferReturn into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceCheckTransferReturn) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the returned Check Transfer.
func (r *TransactionSourceCheckTransferReturn) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

// If available, a document with additional information about the return.
func (r *TransactionSourceCheckTransferReturn) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

func (r TransactionSourceCheckTransferReturn) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCheckTransferReturn{TransferID:%s FileID:%s}", core.FmtP(r.TransferID), core.FmtP(r.FileID))
}

type TransactionSourceCheckTransferRejection struct {
	// The identifier of the Check Transfer that led to this Transaction.
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceCheckTransferRejection using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceCheckTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceCheckTransferRejection into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceCheckTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the Check Transfer that led to this Transaction.
func (r *TransactionSourceCheckTransferRejection) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r TransactionSourceCheckTransferRejection) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCheckTransferRejection{TransferID:%s}", core.FmtP(r.TransferID))
}

type TransactionSourceCheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferID *string `pjson:"transfer_id"`
	// The transaction ID of the corresponding credit transaction.
	TransactionID *string `pjson:"transaction_id"`
	// The time the stop-payment was requested.
	RequestedAt *string `pjson:"requested_at"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type       *TransactionSourceCheckTransferStopPaymentRequestType `pjson:"type"`
	jsonFields map[string]interface{}                                `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceCheckTransferStopPaymentRequest using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceCheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceCheckTransferStopPaymentRequest into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *TransactionSourceCheckTransferStopPaymentRequest) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The ID of the check transfer that was stopped.
func (r *TransactionSourceCheckTransferStopPaymentRequest) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

// The transaction ID of the corresponding credit transaction.
func (r *TransactionSourceCheckTransferStopPaymentRequest) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

// The time the stop-payment was requested.
func (r *TransactionSourceCheckTransferStopPaymentRequest) GetRequestedAt() (RequestedAt string) {
	if r != nil && r.RequestedAt != nil {
		RequestedAt = *r.RequestedAt
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_stop_payment_request`.
func (r *TransactionSourceCheckTransferStopPaymentRequest) GetType() (Type TransactionSourceCheckTransferStopPaymentRequestType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r TransactionSourceCheckTransferStopPaymentRequest) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCheckTransferStopPaymentRequest{TransferID:%s TransactionID:%s RequestedAt:%s Type:%s}", core.FmtP(r.TransferID), core.FmtP(r.TransactionID), core.FmtP(r.RequestedAt), core.FmtP(r.Type))
}

type TransactionSourceCheckTransferStopPaymentRequestType string

const (
	TransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest TransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

type TransactionSourceDisputeResolution struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *TransactionSourceDisputeResolutionCurrency `pjson:"currency"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionID *string                `pjson:"disputed_transaction_id"`
	jsonFields            map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceDisputeResolution using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceDisputeResolution) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceDisputeResolution into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceDisputeResolution) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceDisputeResolution) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *TransactionSourceDisputeResolution) GetCurrency() (Currency TransactionSourceDisputeResolutionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Transaction that was disputed.
func (r *TransactionSourceDisputeResolution) GetDisputedTransactionID() (DisputedTransactionID string) {
	if r != nil && r.DisputedTransactionID != nil {
		DisputedTransactionID = *r.DisputedTransactionID
	}
	return
}

func (r TransactionSourceDisputeResolution) String() (result string) {
	return fmt.Sprintf("&TransactionSourceDisputeResolution{Amount:%s Currency:%s DisputedTransactionID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.DisputedTransactionID))
}

type TransactionSourceDisputeResolutionCurrency string

const (
	TransactionSourceDisputeResolutionCurrencyCad TransactionSourceDisputeResolutionCurrency = "CAD"
	TransactionSourceDisputeResolutionCurrencyChf TransactionSourceDisputeResolutionCurrency = "CHF"
	TransactionSourceDisputeResolutionCurrencyEur TransactionSourceDisputeResolutionCurrency = "EUR"
	TransactionSourceDisputeResolutionCurrencyGbp TransactionSourceDisputeResolutionCurrency = "GBP"
	TransactionSourceDisputeResolutionCurrencyJpy TransactionSourceDisputeResolutionCurrency = "JPY"
	TransactionSourceDisputeResolutionCurrencyUsd TransactionSourceDisputeResolutionCurrency = "USD"
)

type TransactionSourceEmpyrealCashDeposit struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount      *int64                 `pjson:"amount"`
	BagID       *string                `pjson:"bag_id"`
	DepositDate *string                `pjson:"deposit_date"`
	jsonFields  map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceEmpyrealCashDeposit using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceEmpyrealCashDeposit) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceEmpyrealCashDeposit into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceEmpyrealCashDeposit) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceEmpyrealCashDeposit) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *TransactionSourceEmpyrealCashDeposit) GetBagID() (BagID string) {
	if r != nil && r.BagID != nil {
		BagID = *r.BagID
	}
	return
}

func (r *TransactionSourceEmpyrealCashDeposit) GetDepositDate() (DepositDate string) {
	if r != nil && r.DepositDate != nil {
		DepositDate = *r.DepositDate
	}
	return
}

func (r TransactionSourceEmpyrealCashDeposit) String() (result string) {
	return fmt.Sprintf("&TransactionSourceEmpyrealCashDeposit{Amount:%s BagID:%s DepositDate:%s}", core.FmtP(r.Amount), core.FmtP(r.BagID), core.FmtP(r.DepositDate))
}

type TransactionSourceInboundACHTransfer struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount                             *int64                 `pjson:"amount"`
	OriginatorCompanyName              *string                `pjson:"originator_company_name"`
	OriginatorCompanyDescriptiveDate   *string                `pjson:"originator_company_descriptive_date"`
	OriginatorCompanyDiscretionaryData *string                `pjson:"originator_company_discretionary_data"`
	OriginatorCompanyEntryDescription  *string                `pjson:"originator_company_entry_description"`
	OriginatorCompanyID                *string                `pjson:"originator_company_id"`
	ReceiverIDNumber                   *string                `pjson:"receiver_id_number"`
	ReceiverName                       *string                `pjson:"receiver_name"`
	TraceNumber                        *string                `pjson:"trace_number"`
	jsonFields                         map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceInboundACHTransfer using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceInboundACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceInboundACHTransfer into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceInboundACHTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the destination account currency. For dollars,
// for example, this is cents.
func (r *TransactionSourceInboundACHTransfer) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetOriginatorCompanyName() (OriginatorCompanyName string) {
	if r != nil && r.OriginatorCompanyName != nil {
		OriginatorCompanyName = *r.OriginatorCompanyName
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetOriginatorCompanyDescriptiveDate() (OriginatorCompanyDescriptiveDate string) {
	if r != nil && r.OriginatorCompanyDescriptiveDate != nil {
		OriginatorCompanyDescriptiveDate = *r.OriginatorCompanyDescriptiveDate
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetOriginatorCompanyDiscretionaryData() (OriginatorCompanyDiscretionaryData string) {
	if r != nil && r.OriginatorCompanyDiscretionaryData != nil {
		OriginatorCompanyDiscretionaryData = *r.OriginatorCompanyDiscretionaryData
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r != nil && r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetOriginatorCompanyID() (OriginatorCompanyID string) {
	if r != nil && r.OriginatorCompanyID != nil {
		OriginatorCompanyID = *r.OriginatorCompanyID
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetReceiverIDNumber() (ReceiverIDNumber string) {
	if r != nil && r.ReceiverIDNumber != nil {
		ReceiverIDNumber = *r.ReceiverIDNumber
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetReceiverName() (ReceiverName string) {
	if r != nil && r.ReceiverName != nil {
		ReceiverName = *r.ReceiverName
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

func (r TransactionSourceInboundACHTransfer) String() (result string) {
	return fmt.Sprintf("&TransactionSourceInboundACHTransfer{Amount:%s OriginatorCompanyName:%s OriginatorCompanyDescriptiveDate:%s OriginatorCompanyDiscretionaryData:%s OriginatorCompanyEntryDescription:%s OriginatorCompanyID:%s ReceiverIDNumber:%s ReceiverName:%s TraceNumber:%s}", core.FmtP(r.Amount), core.FmtP(r.OriginatorCompanyName), core.FmtP(r.OriginatorCompanyDescriptiveDate), core.FmtP(r.OriginatorCompanyDiscretionaryData), core.FmtP(r.OriginatorCompanyEntryDescription), core.FmtP(r.OriginatorCompanyID), core.FmtP(r.ReceiverIDNumber), core.FmtP(r.ReceiverName), core.FmtP(r.TraceNumber))
}

type TransactionSourceInboundCheck struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency              *TransactionSourceInboundCheckCurrency `pjson:"currency"`
	CheckNumber           *string                                `pjson:"check_number"`
	CheckFrontImageFileID *string                                `pjson:"check_front_image_file_id"`
	CheckRearImageFileID  *string                                `pjson:"check_rear_image_file_id"`
	jsonFields            map[string]interface{}                 `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into TransactionSourceInboundCheck
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *TransactionSourceInboundCheck) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceInboundCheck into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceInboundCheck) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the destination account currency. For dollars,
// for example, this is cents.
func (r *TransactionSourceInboundCheck) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *TransactionSourceInboundCheck) GetCurrency() (Currency TransactionSourceInboundCheckCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *TransactionSourceInboundCheck) GetCheckNumber() (CheckNumber string) {
	if r != nil && r.CheckNumber != nil {
		CheckNumber = *r.CheckNumber
	}
	return
}

func (r *TransactionSourceInboundCheck) GetCheckFrontImageFileID() (CheckFrontImageFileID string) {
	if r != nil && r.CheckFrontImageFileID != nil {
		CheckFrontImageFileID = *r.CheckFrontImageFileID
	}
	return
}

func (r *TransactionSourceInboundCheck) GetCheckRearImageFileID() (CheckRearImageFileID string) {
	if r != nil && r.CheckRearImageFileID != nil {
		CheckRearImageFileID = *r.CheckRearImageFileID
	}
	return
}

func (r TransactionSourceInboundCheck) String() (result string) {
	return fmt.Sprintf("&TransactionSourceInboundCheck{Amount:%s Currency:%s CheckNumber:%s CheckFrontImageFileID:%s CheckRearImageFileID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.CheckNumber), core.FmtP(r.CheckFrontImageFileID), core.FmtP(r.CheckRearImageFileID))
}

type TransactionSourceInboundCheckCurrency string

const (
	TransactionSourceInboundCheckCurrencyCad TransactionSourceInboundCheckCurrency = "CAD"
	TransactionSourceInboundCheckCurrencyChf TransactionSourceInboundCheckCurrency = "CHF"
	TransactionSourceInboundCheckCurrencyEur TransactionSourceInboundCheckCurrency = "EUR"
	TransactionSourceInboundCheckCurrencyGbp TransactionSourceInboundCheckCurrency = "GBP"
	TransactionSourceInboundCheckCurrencyJpy TransactionSourceInboundCheckCurrency = "JPY"
	TransactionSourceInboundCheckCurrencyUsd TransactionSourceInboundCheckCurrency = "USD"
)

type TransactionSourceInboundInternationalACHTransfer struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount                                                 *int64                 `pjson:"amount"`
	ForeignExchangeIndicator                               *string                `pjson:"foreign_exchange_indicator"`
	ForeignExchangeReferenceIndicator                      *string                `pjson:"foreign_exchange_reference_indicator"`
	ForeignExchangeReference                               *string                `pjson:"foreign_exchange_reference"`
	DestinationCountryCode                                 *string                `pjson:"destination_country_code"`
	DestinationCurrencyCode                                *string                `pjson:"destination_currency_code"`
	ForeignPaymentAmount                                   *int64                 `pjson:"foreign_payment_amount"`
	ForeignTraceNumber                                     *string                `pjson:"foreign_trace_number"`
	InternationalTransactionTypeCode                       *string                `pjson:"international_transaction_type_code"`
	OriginatingCurrencyCode                                *string                `pjson:"originating_currency_code"`
	OriginatingDepositoryFinancialInstitutionName          *string                `pjson:"originating_depository_financial_institution_name"`
	OriginatingDepositoryFinancialInstitutionIDQualifier   *string                `pjson:"originating_depository_financial_institution_id_qualifier"`
	OriginatingDepositoryFinancialInstitutionID            *string                `pjson:"originating_depository_financial_institution_id"`
	OriginatingDepositoryFinancialInstitutionBranchCountry *string                `pjson:"originating_depository_financial_institution_branch_country"`
	OriginatorCity                                         *string                `pjson:"originator_city"`
	OriginatorCompanyEntryDescription                      *string                `pjson:"originator_company_entry_description"`
	OriginatorCountry                                      *string                `pjson:"originator_country"`
	OriginatorIdentification                               *string                `pjson:"originator_identification"`
	OriginatorName                                         *string                `pjson:"originator_name"`
	OriginatorPostalCode                                   *string                `pjson:"originator_postal_code"`
	OriginatorStreetAddress                                *string                `pjson:"originator_street_address"`
	OriginatorStateOrProvince                              *string                `pjson:"originator_state_or_province"`
	PaymentRelatedInformation                              *string                `pjson:"payment_related_information"`
	PaymentRelatedInformation2                             *string                `pjson:"payment_related_information2"`
	ReceiverIdentificationNumber                           *string                `pjson:"receiver_identification_number"`
	ReceiverStreetAddress                                  *string                `pjson:"receiver_street_address"`
	ReceiverCity                                           *string                `pjson:"receiver_city"`
	ReceiverStateOrProvince                                *string                `pjson:"receiver_state_or_province"`
	ReceiverCountry                                        *string                `pjson:"receiver_country"`
	ReceiverPostalCode                                     *string                `pjson:"receiver_postal_code"`
	ReceivingCompanyOrIndividualName                       *string                `pjson:"receiving_company_or_individual_name"`
	ReceivingDepositoryFinancialInstitutionName            *string                `pjson:"receiving_depository_financial_institution_name"`
	ReceivingDepositoryFinancialInstitutionIDQualifier     *string                `pjson:"receiving_depository_financial_institution_id_qualifier"`
	ReceivingDepositoryFinancialInstitutionID              *string                `pjson:"receiving_depository_financial_institution_id"`
	ReceivingDepositoryFinancialInstitutionCountry         *string                `pjson:"receiving_depository_financial_institution_country"`
	TraceNumber                                            *string                `pjson:"trace_number"`
	jsonFields                                             map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceInboundInternationalACHTransfer using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceInboundInternationalACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceInboundInternationalACHTransfer into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *TransactionSourceInboundInternationalACHTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the destination account currency. For dollars,
// for example, this is cents.
func (r *TransactionSourceInboundInternationalACHTransfer) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetForeignExchangeIndicator() (ForeignExchangeIndicator string) {
	if r != nil && r.ForeignExchangeIndicator != nil {
		ForeignExchangeIndicator = *r.ForeignExchangeIndicator
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetForeignExchangeReferenceIndicator() (ForeignExchangeReferenceIndicator string) {
	if r != nil && r.ForeignExchangeReferenceIndicator != nil {
		ForeignExchangeReferenceIndicator = *r.ForeignExchangeReferenceIndicator
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetForeignExchangeReference() (ForeignExchangeReference string) {
	if r != nil && r.ForeignExchangeReference != nil {
		ForeignExchangeReference = *r.ForeignExchangeReference
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetDestinationCountryCode() (DestinationCountryCode string) {
	if r != nil && r.DestinationCountryCode != nil {
		DestinationCountryCode = *r.DestinationCountryCode
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetDestinationCurrencyCode() (DestinationCurrencyCode string) {
	if r != nil && r.DestinationCurrencyCode != nil {
		DestinationCurrencyCode = *r.DestinationCurrencyCode
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetForeignPaymentAmount() (ForeignPaymentAmount int64) {
	if r != nil && r.ForeignPaymentAmount != nil {
		ForeignPaymentAmount = *r.ForeignPaymentAmount
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetForeignTraceNumber() (ForeignTraceNumber string) {
	if r != nil && r.ForeignTraceNumber != nil {
		ForeignTraceNumber = *r.ForeignTraceNumber
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetInternationalTransactionTypeCode() (InternationalTransactionTypeCode string) {
	if r != nil && r.InternationalTransactionTypeCode != nil {
		InternationalTransactionTypeCode = *r.InternationalTransactionTypeCode
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatingCurrencyCode() (OriginatingCurrencyCode string) {
	if r != nil && r.OriginatingCurrencyCode != nil {
		OriginatingCurrencyCode = *r.OriginatingCurrencyCode
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionName() (OriginatingDepositoryFinancialInstitutionName string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionName != nil {
		OriginatingDepositoryFinancialInstitutionName = *r.OriginatingDepositoryFinancialInstitutionName
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionIDQualifier() (OriginatingDepositoryFinancialInstitutionIDQualifier string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionIDQualifier != nil {
		OriginatingDepositoryFinancialInstitutionIDQualifier = *r.OriginatingDepositoryFinancialInstitutionIDQualifier
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionID() (OriginatingDepositoryFinancialInstitutionID string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionID != nil {
		OriginatingDepositoryFinancialInstitutionID = *r.OriginatingDepositoryFinancialInstitutionID
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionBranchCountry() (OriginatingDepositoryFinancialInstitutionBranchCountry string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionBranchCountry != nil {
		OriginatingDepositoryFinancialInstitutionBranchCountry = *r.OriginatingDepositoryFinancialInstitutionBranchCountry
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatorCity() (OriginatorCity string) {
	if r != nil && r.OriginatorCity != nil {
		OriginatorCity = *r.OriginatorCity
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r != nil && r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatorCountry() (OriginatorCountry string) {
	if r != nil && r.OriginatorCountry != nil {
		OriginatorCountry = *r.OriginatorCountry
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatorIdentification() (OriginatorIdentification string) {
	if r != nil && r.OriginatorIdentification != nil {
		OriginatorIdentification = *r.OriginatorIdentification
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatorPostalCode() (OriginatorPostalCode string) {
	if r != nil && r.OriginatorPostalCode != nil {
		OriginatorPostalCode = *r.OriginatorPostalCode
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatorStreetAddress() (OriginatorStreetAddress string) {
	if r != nil && r.OriginatorStreetAddress != nil {
		OriginatorStreetAddress = *r.OriginatorStreetAddress
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatorStateOrProvince() (OriginatorStateOrProvince string) {
	if r != nil && r.OriginatorStateOrProvince != nil {
		OriginatorStateOrProvince = *r.OriginatorStateOrProvince
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetPaymentRelatedInformation() (PaymentRelatedInformation string) {
	if r != nil && r.PaymentRelatedInformation != nil {
		PaymentRelatedInformation = *r.PaymentRelatedInformation
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetPaymentRelatedInformation2() (PaymentRelatedInformation2 string) {
	if r != nil && r.PaymentRelatedInformation2 != nil {
		PaymentRelatedInformation2 = *r.PaymentRelatedInformation2
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceiverIdentificationNumber() (ReceiverIdentificationNumber string) {
	if r != nil && r.ReceiverIdentificationNumber != nil {
		ReceiverIdentificationNumber = *r.ReceiverIdentificationNumber
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceiverStreetAddress() (ReceiverStreetAddress string) {
	if r != nil && r.ReceiverStreetAddress != nil {
		ReceiverStreetAddress = *r.ReceiverStreetAddress
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceiverCity() (ReceiverCity string) {
	if r != nil && r.ReceiverCity != nil {
		ReceiverCity = *r.ReceiverCity
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceiverStateOrProvince() (ReceiverStateOrProvince string) {
	if r != nil && r.ReceiverStateOrProvince != nil {
		ReceiverStateOrProvince = *r.ReceiverStateOrProvince
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceiverCountry() (ReceiverCountry string) {
	if r != nil && r.ReceiverCountry != nil {
		ReceiverCountry = *r.ReceiverCountry
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceiverPostalCode() (ReceiverPostalCode string) {
	if r != nil && r.ReceiverPostalCode != nil {
		ReceiverPostalCode = *r.ReceiverPostalCode
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceivingCompanyOrIndividualName() (ReceivingCompanyOrIndividualName string) {
	if r != nil && r.ReceivingCompanyOrIndividualName != nil {
		ReceivingCompanyOrIndividualName = *r.ReceivingCompanyOrIndividualName
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionName() (ReceivingDepositoryFinancialInstitutionName string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionName != nil {
		ReceivingDepositoryFinancialInstitutionName = *r.ReceivingDepositoryFinancialInstitutionName
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionIDQualifier() (ReceivingDepositoryFinancialInstitutionIDQualifier string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionIDQualifier != nil {
		ReceivingDepositoryFinancialInstitutionIDQualifier = *r.ReceivingDepositoryFinancialInstitutionIDQualifier
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionID() (ReceivingDepositoryFinancialInstitutionID string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionID != nil {
		ReceivingDepositoryFinancialInstitutionID = *r.ReceivingDepositoryFinancialInstitutionID
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionCountry() (ReceivingDepositoryFinancialInstitutionCountry string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionCountry != nil {
		ReceivingDepositoryFinancialInstitutionCountry = *r.ReceivingDepositoryFinancialInstitutionCountry
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

func (r TransactionSourceInboundInternationalACHTransfer) String() (result string) {
	return fmt.Sprintf("&TransactionSourceInboundInternationalACHTransfer{Amount:%s ForeignExchangeIndicator:%s ForeignExchangeReferenceIndicator:%s ForeignExchangeReference:%s DestinationCountryCode:%s DestinationCurrencyCode:%s ForeignPaymentAmount:%s ForeignTraceNumber:%s InternationalTransactionTypeCode:%s OriginatingCurrencyCode:%s OriginatingDepositoryFinancialInstitutionName:%s OriginatingDepositoryFinancialInstitutionIDQualifier:%s OriginatingDepositoryFinancialInstitutionID:%s OriginatingDepositoryFinancialInstitutionBranchCountry:%s OriginatorCity:%s OriginatorCompanyEntryDescription:%s OriginatorCountry:%s OriginatorIdentification:%s OriginatorName:%s OriginatorPostalCode:%s OriginatorStreetAddress:%s OriginatorStateOrProvince:%s PaymentRelatedInformation:%s PaymentRelatedInformation2:%s ReceiverIdentificationNumber:%s ReceiverStreetAddress:%s ReceiverCity:%s ReceiverStateOrProvince:%s ReceiverCountry:%s ReceiverPostalCode:%s ReceivingCompanyOrIndividualName:%s ReceivingDepositoryFinancialInstitutionName:%s ReceivingDepositoryFinancialInstitutionIDQualifier:%s ReceivingDepositoryFinancialInstitutionID:%s ReceivingDepositoryFinancialInstitutionCountry:%s TraceNumber:%s}", core.FmtP(r.Amount), core.FmtP(r.ForeignExchangeIndicator), core.FmtP(r.ForeignExchangeReferenceIndicator), core.FmtP(r.ForeignExchangeReference), core.FmtP(r.DestinationCountryCode), core.FmtP(r.DestinationCurrencyCode), core.FmtP(r.ForeignPaymentAmount), core.FmtP(r.ForeignTraceNumber), core.FmtP(r.InternationalTransactionTypeCode), core.FmtP(r.OriginatingCurrencyCode), core.FmtP(r.OriginatingDepositoryFinancialInstitutionName), core.FmtP(r.OriginatingDepositoryFinancialInstitutionIDQualifier), core.FmtP(r.OriginatingDepositoryFinancialInstitutionID), core.FmtP(r.OriginatingDepositoryFinancialInstitutionBranchCountry), core.FmtP(r.OriginatorCity), core.FmtP(r.OriginatorCompanyEntryDescription), core.FmtP(r.OriginatorCountry), core.FmtP(r.OriginatorIdentification), core.FmtP(r.OriginatorName), core.FmtP(r.OriginatorPostalCode), core.FmtP(r.OriginatorStreetAddress), core.FmtP(r.OriginatorStateOrProvince), core.FmtP(r.PaymentRelatedInformation), core.FmtP(r.PaymentRelatedInformation2), core.FmtP(r.ReceiverIdentificationNumber), core.FmtP(r.ReceiverStreetAddress), core.FmtP(r.ReceiverCity), core.FmtP(r.ReceiverStateOrProvince), core.FmtP(r.ReceiverCountry), core.FmtP(r.ReceiverPostalCode), core.FmtP(r.ReceivingCompanyOrIndividualName), core.FmtP(r.ReceivingDepositoryFinancialInstitutionName), core.FmtP(r.ReceivingDepositoryFinancialInstitutionIDQualifier), core.FmtP(r.ReceivingDepositoryFinancialInstitutionID), core.FmtP(r.ReceivingDepositoryFinancialInstitutionCountry), core.FmtP(r.TraceNumber))
}

type TransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real Time Payments transfer.
	Currency *TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency `pjson:"currency"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName *string `pjson:"creditor_name"`
	// The name provided by the sender of the transfer.
	DebtorName *string `pjson:"debtor_name"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber *string `pjson:"debtor_account_number"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber *string `pjson:"debtor_routing_number"`
	// The Real Time Payments network identification of the transfer
	TransactionIdentification *string `pjson:"transaction_identification"`
	// Additional information included with the transfer.
	RemittanceInformation *string                `pjson:"remittance_information"`
	jsonFields            map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceInboundRealTimePaymentsTransferConfirmation using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// TransactionSourceInboundRealTimePaymentsTransferConfirmation into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transfer's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
// currency. This will always be "USD" for a Real Time Payments transfer.
func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) GetCurrency() (Currency TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The name the sender of the transfer specified as the recipient of the transfer.
func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) GetCreditorName() (CreditorName string) {
	if r != nil && r.CreditorName != nil {
		CreditorName = *r.CreditorName
	}
	return
}

// The name provided by the sender of the transfer.
func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorName() (DebtorName string) {
	if r != nil && r.DebtorName != nil {
		DebtorName = *r.DebtorName
	}
	return
}

// The account number of the account that sent the transfer.
func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorAccountNumber() (DebtorAccountNumber string) {
	if r != nil && r.DebtorAccountNumber != nil {
		DebtorAccountNumber = *r.DebtorAccountNumber
	}
	return
}

// The routing number of the account that sent the transfer.
func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorRoutingNumber() (DebtorRoutingNumber string) {
	if r != nil && r.DebtorRoutingNumber != nil {
		DebtorRoutingNumber = *r.DebtorRoutingNumber
	}
	return
}

// The Real Time Payments network identification of the transfer
func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) GetTransactionIdentification() (TransactionIdentification string) {
	if r != nil && r.TransactionIdentification != nil {
		TransactionIdentification = *r.TransactionIdentification
	}
	return
}

// Additional information included with the transfer.
func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) GetRemittanceInformation() (RemittanceInformation string) {
	if r != nil && r.RemittanceInformation != nil {
		RemittanceInformation = *r.RemittanceInformation
	}
	return
}

func (r TransactionSourceInboundRealTimePaymentsTransferConfirmation) String() (result string) {
	return fmt.Sprintf("&TransactionSourceInboundRealTimePaymentsTransferConfirmation{Amount:%s Currency:%s CreditorName:%s DebtorName:%s DebtorAccountNumber:%s DebtorRoutingNumber:%s TransactionIdentification:%s RemittanceInformation:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.CreditorName), core.FmtP(r.DebtorName), core.FmtP(r.DebtorAccountNumber), core.FmtP(r.DebtorRoutingNumber), core.FmtP(r.TransactionIdentification), core.FmtP(r.RemittanceInformation))
}

type TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency string

const (
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyCad TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CAD"
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyChf TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CHF"
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyEur TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "EUR"
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyGbp TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "GBP"
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyJpy TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "JPY"
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyUsd TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "USD"
)

type TransactionSourceInboundWireDrawdownPaymentReversal struct {
	// The amount that was reversed.
	Amount *int64 `pjson:"amount"`
	// The description on the reversal message from Fedwire.
	Description *string `pjson:"description"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate *string `pjson:"input_cycle_date"`
	// The Fedwire sequence number.
	InputSequenceNumber *string `pjson:"input_sequence_number"`
	// The Fedwire input source identifier.
	InputSource *string `pjson:"input_source"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData *string `pjson:"input_message_accountability_data"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData *string `pjson:"previous_message_input_message_accountability_data"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate *string `pjson:"previous_message_input_cycle_date"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber *string `pjson:"previous_message_input_sequence_number"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource *string                `pjson:"previous_message_input_source"`
	jsonFields                 map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceInboundWireDrawdownPaymentReversal using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceInboundWireDrawdownPaymentReversal into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount that was reversed.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description on the reversal message from Fedwire.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Fedwire cycle date for the wire reversal.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) GetInputCycleDate() (InputCycleDate string) {
	if r != nil && r.InputCycleDate != nil {
		InputCycleDate = *r.InputCycleDate
	}
	return
}

// The Fedwire sequence number.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) GetInputSequenceNumber() (InputSequenceNumber string) {
	if r != nil && r.InputSequenceNumber != nil {
		InputSequenceNumber = *r.InputSequenceNumber
	}
	return
}

// The Fedwire input source identifier.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) GetInputSource() (InputSource string) {
	if r != nil && r.InputSource != nil {
		InputSource = *r.InputSource
	}
	return
}

// The Fedwire transaction identifier.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

// The Fedwire transaction identifier for the wire transfer that was reversed.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputMessageAccountabilityData() (PreviousMessageInputMessageAccountabilityData string) {
	if r != nil && r.PreviousMessageInputMessageAccountabilityData != nil {
		PreviousMessageInputMessageAccountabilityData = *r.PreviousMessageInputMessageAccountabilityData
	}
	return
}

// The Fedwire cycle date for the wire transfer that was reversed.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputCycleDate() (PreviousMessageInputCycleDate string) {
	if r != nil && r.PreviousMessageInputCycleDate != nil {
		PreviousMessageInputCycleDate = *r.PreviousMessageInputCycleDate
	}
	return
}

// The Fedwire sequence number for the wire transfer that was reversed.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputSequenceNumber() (PreviousMessageInputSequenceNumber string) {
	if r != nil && r.PreviousMessageInputSequenceNumber != nil {
		PreviousMessageInputSequenceNumber = *r.PreviousMessageInputSequenceNumber
	}
	return
}

// The Fedwire input source identifier for the wire transfer that was reversed.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputSource() (PreviousMessageInputSource string) {
	if r != nil && r.PreviousMessageInputSource != nil {
		PreviousMessageInputSource = *r.PreviousMessageInputSource
	}
	return
}

func (r TransactionSourceInboundWireDrawdownPaymentReversal) String() (result string) {
	return fmt.Sprintf("&TransactionSourceInboundWireDrawdownPaymentReversal{Amount:%s Description:%s InputCycleDate:%s InputSequenceNumber:%s InputSource:%s InputMessageAccountabilityData:%s PreviousMessageInputMessageAccountabilityData:%s PreviousMessageInputCycleDate:%s PreviousMessageInputSequenceNumber:%s PreviousMessageInputSource:%s}", core.FmtP(r.Amount), core.FmtP(r.Description), core.FmtP(r.InputCycleDate), core.FmtP(r.InputSequenceNumber), core.FmtP(r.InputSource), core.FmtP(r.InputMessageAccountabilityData), core.FmtP(r.PreviousMessageInputMessageAccountabilityData), core.FmtP(r.PreviousMessageInputCycleDate), core.FmtP(r.PreviousMessageInputSequenceNumber), core.FmtP(r.PreviousMessageInputSource))
}

type TransactionSourceInboundWireDrawdownPayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount                             *int64                 `pjson:"amount"`
	BeneficiaryAddressLine1            *string                `pjson:"beneficiary_address_line1"`
	BeneficiaryAddressLine2            *string                `pjson:"beneficiary_address_line2"`
	BeneficiaryAddressLine3            *string                `pjson:"beneficiary_address_line3"`
	BeneficiaryName                    *string                `pjson:"beneficiary_name"`
	BeneficiaryReference               *string                `pjson:"beneficiary_reference"`
	Description                        *string                `pjson:"description"`
	InputMessageAccountabilityData     *string                `pjson:"input_message_accountability_data"`
	OriginatorAddressLine1             *string                `pjson:"originator_address_line1"`
	OriginatorAddressLine2             *string                `pjson:"originator_address_line2"`
	OriginatorAddressLine3             *string                `pjson:"originator_address_line3"`
	OriginatorName                     *string                `pjson:"originator_name"`
	OriginatorToBeneficiaryInformation *string                `pjson:"originator_to_beneficiary_information"`
	jsonFields                         map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceInboundWireDrawdownPayment using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceInboundWireDrawdownPayment) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceInboundWireDrawdownPayment into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *TransactionSourceInboundWireDrawdownPayment) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceInboundWireDrawdownPayment) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r != nil && r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r != nil && r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r != nil && r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetBeneficiaryName() (BeneficiaryName string) {
	if r != nil && r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetBeneficiaryReference() (BeneficiaryReference string) {
	if r != nil && r.BeneficiaryReference != nil {
		BeneficiaryReference = *r.BeneficiaryReference
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine1() (OriginatorAddressLine1 string) {
	if r != nil && r.OriginatorAddressLine1 != nil {
		OriginatorAddressLine1 = *r.OriginatorAddressLine1
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine2() (OriginatorAddressLine2 string) {
	if r != nil && r.OriginatorAddressLine2 != nil {
		OriginatorAddressLine2 = *r.OriginatorAddressLine2
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine3() (OriginatorAddressLine3 string) {
	if r != nil && r.OriginatorAddressLine3 != nil {
		OriginatorAddressLine3 = *r.OriginatorAddressLine3
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetOriginatorToBeneficiaryInformation() (OriginatorToBeneficiaryInformation string) {
	if r != nil && r.OriginatorToBeneficiaryInformation != nil {
		OriginatorToBeneficiaryInformation = *r.OriginatorToBeneficiaryInformation
	}
	return
}

func (r TransactionSourceInboundWireDrawdownPayment) String() (result string) {
	return fmt.Sprintf("&TransactionSourceInboundWireDrawdownPayment{Amount:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s BeneficiaryName:%s BeneficiaryReference:%s Description:%s InputMessageAccountabilityData:%s OriginatorAddressLine1:%s OriginatorAddressLine2:%s OriginatorAddressLine3:%s OriginatorName:%s OriginatorToBeneficiaryInformation:%s}", core.FmtP(r.Amount), core.FmtP(r.BeneficiaryAddressLine1), core.FmtP(r.BeneficiaryAddressLine2), core.FmtP(r.BeneficiaryAddressLine3), core.FmtP(r.BeneficiaryName), core.FmtP(r.BeneficiaryReference), core.FmtP(r.Description), core.FmtP(r.InputMessageAccountabilityData), core.FmtP(r.OriginatorAddressLine1), core.FmtP(r.OriginatorAddressLine2), core.FmtP(r.OriginatorAddressLine3), core.FmtP(r.OriginatorName), core.FmtP(r.OriginatorToBeneficiaryInformation))
}

type TransactionSourceInboundWireReversal struct {
	// The amount that was reversed.
	Amount *int64 `pjson:"amount"`
	// The description on the reversal message from Fedwire.
	Description *string `pjson:"description"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate *string `pjson:"input_cycle_date"`
	// The Fedwire sequence number.
	InputSequenceNumber *string `pjson:"input_sequence_number"`
	// The Fedwire input source identifier.
	InputSource *string `pjson:"input_source"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData *string `pjson:"input_message_accountability_data"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData *string `pjson:"previous_message_input_message_accountability_data"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate *string `pjson:"previous_message_input_cycle_date"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber *string `pjson:"previous_message_input_sequence_number"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource *string `pjson:"previous_message_input_source"`
	// Information included in the wire reversal for the receiving financial
	// institution.
	ReceiverFinancialInstitutionInformation *string `pjson:"receiver_financial_institution_information"`
	// Additional financial institution information included in the wire reversal.
	FinancialInstitutionToFinancialInstitutionInformation *string                `pjson:"financial_institution_to_financial_institution_information"`
	jsonFields                                            map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceInboundWireReversal using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceInboundWireReversal) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceInboundWireReversal into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceInboundWireReversal) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount that was reversed.
func (r *TransactionSourceInboundWireReversal) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description on the reversal message from Fedwire.
func (r *TransactionSourceInboundWireReversal) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Fedwire cycle date for the wire reversal.
func (r *TransactionSourceInboundWireReversal) GetInputCycleDate() (InputCycleDate string) {
	if r != nil && r.InputCycleDate != nil {
		InputCycleDate = *r.InputCycleDate
	}
	return
}

// The Fedwire sequence number.
func (r *TransactionSourceInboundWireReversal) GetInputSequenceNumber() (InputSequenceNumber string) {
	if r != nil && r.InputSequenceNumber != nil {
		InputSequenceNumber = *r.InputSequenceNumber
	}
	return
}

// The Fedwire input source identifier.
func (r *TransactionSourceInboundWireReversal) GetInputSource() (InputSource string) {
	if r != nil && r.InputSource != nil {
		InputSource = *r.InputSource
	}
	return
}

// The Fedwire transaction identifier.
func (r *TransactionSourceInboundWireReversal) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

// The Fedwire transaction identifier for the wire transfer that was reversed.
func (r *TransactionSourceInboundWireReversal) GetPreviousMessageInputMessageAccountabilityData() (PreviousMessageInputMessageAccountabilityData string) {
	if r != nil && r.PreviousMessageInputMessageAccountabilityData != nil {
		PreviousMessageInputMessageAccountabilityData = *r.PreviousMessageInputMessageAccountabilityData
	}
	return
}

// The Fedwire cycle date for the wire transfer that was reversed.
func (r *TransactionSourceInboundWireReversal) GetPreviousMessageInputCycleDate() (PreviousMessageInputCycleDate string) {
	if r != nil && r.PreviousMessageInputCycleDate != nil {
		PreviousMessageInputCycleDate = *r.PreviousMessageInputCycleDate
	}
	return
}

// The Fedwire sequence number for the wire transfer that was reversed.
func (r *TransactionSourceInboundWireReversal) GetPreviousMessageInputSequenceNumber() (PreviousMessageInputSequenceNumber string) {
	if r != nil && r.PreviousMessageInputSequenceNumber != nil {
		PreviousMessageInputSequenceNumber = *r.PreviousMessageInputSequenceNumber
	}
	return
}

// The Fedwire input source identifier for the wire transfer that was reversed.
func (r *TransactionSourceInboundWireReversal) GetPreviousMessageInputSource() (PreviousMessageInputSource string) {
	if r != nil && r.PreviousMessageInputSource != nil {
		PreviousMessageInputSource = *r.PreviousMessageInputSource
	}
	return
}

// Information included in the wire reversal for the receiving financial
// institution.
func (r *TransactionSourceInboundWireReversal) GetReceiverFinancialInstitutionInformation() (ReceiverFinancialInstitutionInformation string) {
	if r != nil && r.ReceiverFinancialInstitutionInformation != nil {
		ReceiverFinancialInstitutionInformation = *r.ReceiverFinancialInstitutionInformation
	}
	return
}

// Additional financial institution information included in the wire reversal.
func (r *TransactionSourceInboundWireReversal) GetFinancialInstitutionToFinancialInstitutionInformation() (FinancialInstitutionToFinancialInstitutionInformation string) {
	if r != nil && r.FinancialInstitutionToFinancialInstitutionInformation != nil {
		FinancialInstitutionToFinancialInstitutionInformation = *r.FinancialInstitutionToFinancialInstitutionInformation
	}
	return
}

func (r TransactionSourceInboundWireReversal) String() (result string) {
	return fmt.Sprintf("&TransactionSourceInboundWireReversal{Amount:%s Description:%s InputCycleDate:%s InputSequenceNumber:%s InputSource:%s InputMessageAccountabilityData:%s PreviousMessageInputMessageAccountabilityData:%s PreviousMessageInputCycleDate:%s PreviousMessageInputSequenceNumber:%s PreviousMessageInputSource:%s ReceiverFinancialInstitutionInformation:%s FinancialInstitutionToFinancialInstitutionInformation:%s}", core.FmtP(r.Amount), core.FmtP(r.Description), core.FmtP(r.InputCycleDate), core.FmtP(r.InputSequenceNumber), core.FmtP(r.InputSource), core.FmtP(r.InputMessageAccountabilityData), core.FmtP(r.PreviousMessageInputMessageAccountabilityData), core.FmtP(r.PreviousMessageInputCycleDate), core.FmtP(r.PreviousMessageInputSequenceNumber), core.FmtP(r.PreviousMessageInputSource), core.FmtP(r.ReceiverFinancialInstitutionInformation), core.FmtP(r.FinancialInstitutionToFinancialInstitutionInformation))
}

type TransactionSourceInboundWireTransfer struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount                                  *int64                 `pjson:"amount"`
	BeneficiaryAddressLine1                 *string                `pjson:"beneficiary_address_line1"`
	BeneficiaryAddressLine2                 *string                `pjson:"beneficiary_address_line2"`
	BeneficiaryAddressLine3                 *string                `pjson:"beneficiary_address_line3"`
	BeneficiaryName                         *string                `pjson:"beneficiary_name"`
	BeneficiaryReference                    *string                `pjson:"beneficiary_reference"`
	Description                             *string                `pjson:"description"`
	InputMessageAccountabilityData          *string                `pjson:"input_message_accountability_data"`
	OriginatorAddressLine1                  *string                `pjson:"originator_address_line1"`
	OriginatorAddressLine2                  *string                `pjson:"originator_address_line2"`
	OriginatorAddressLine3                  *string                `pjson:"originator_address_line3"`
	OriginatorName                          *string                `pjson:"originator_name"`
	OriginatorToBeneficiaryInformationLine1 *string                `pjson:"originator_to_beneficiary_information_line1"`
	OriginatorToBeneficiaryInformationLine2 *string                `pjson:"originator_to_beneficiary_information_line2"`
	OriginatorToBeneficiaryInformationLine3 *string                `pjson:"originator_to_beneficiary_information_line3"`
	OriginatorToBeneficiaryInformationLine4 *string                `pjson:"originator_to_beneficiary_information_line4"`
	OriginatorToBeneficiaryInformation      *string                `pjson:"originator_to_beneficiary_information"`
	jsonFields                              map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceInboundWireTransfer using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceInboundWireTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceInboundWireTransfer into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceInboundWireTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceInboundWireTransfer) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r != nil && r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r != nil && r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r != nil && r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetBeneficiaryName() (BeneficiaryName string) {
	if r != nil && r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetBeneficiaryReference() (BeneficiaryReference string) {
	if r != nil && r.BeneficiaryReference != nil {
		BeneficiaryReference = *r.BeneficiaryReference
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetOriginatorAddressLine1() (OriginatorAddressLine1 string) {
	if r != nil && r.OriginatorAddressLine1 != nil {
		OriginatorAddressLine1 = *r.OriginatorAddressLine1
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetOriginatorAddressLine2() (OriginatorAddressLine2 string) {
	if r != nil && r.OriginatorAddressLine2 != nil {
		OriginatorAddressLine2 = *r.OriginatorAddressLine2
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetOriginatorAddressLine3() (OriginatorAddressLine3 string) {
	if r != nil && r.OriginatorAddressLine3 != nil {
		OriginatorAddressLine3 = *r.OriginatorAddressLine3
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetOriginatorToBeneficiaryInformationLine1() (OriginatorToBeneficiaryInformationLine1 string) {
	if r != nil && r.OriginatorToBeneficiaryInformationLine1 != nil {
		OriginatorToBeneficiaryInformationLine1 = *r.OriginatorToBeneficiaryInformationLine1
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetOriginatorToBeneficiaryInformationLine2() (OriginatorToBeneficiaryInformationLine2 string) {
	if r != nil && r.OriginatorToBeneficiaryInformationLine2 != nil {
		OriginatorToBeneficiaryInformationLine2 = *r.OriginatorToBeneficiaryInformationLine2
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetOriginatorToBeneficiaryInformationLine3() (OriginatorToBeneficiaryInformationLine3 string) {
	if r != nil && r.OriginatorToBeneficiaryInformationLine3 != nil {
		OriginatorToBeneficiaryInformationLine3 = *r.OriginatorToBeneficiaryInformationLine3
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetOriginatorToBeneficiaryInformationLine4() (OriginatorToBeneficiaryInformationLine4 string) {
	if r != nil && r.OriginatorToBeneficiaryInformationLine4 != nil {
		OriginatorToBeneficiaryInformationLine4 = *r.OriginatorToBeneficiaryInformationLine4
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetOriginatorToBeneficiaryInformation() (OriginatorToBeneficiaryInformation string) {
	if r != nil && r.OriginatorToBeneficiaryInformation != nil {
		OriginatorToBeneficiaryInformation = *r.OriginatorToBeneficiaryInformation
	}
	return
}

func (r TransactionSourceInboundWireTransfer) String() (result string) {
	return fmt.Sprintf("&TransactionSourceInboundWireTransfer{Amount:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s BeneficiaryName:%s BeneficiaryReference:%s Description:%s InputMessageAccountabilityData:%s OriginatorAddressLine1:%s OriginatorAddressLine2:%s OriginatorAddressLine3:%s OriginatorName:%s OriginatorToBeneficiaryInformationLine1:%s OriginatorToBeneficiaryInformationLine2:%s OriginatorToBeneficiaryInformationLine3:%s OriginatorToBeneficiaryInformationLine4:%s OriginatorToBeneficiaryInformation:%s}", core.FmtP(r.Amount), core.FmtP(r.BeneficiaryAddressLine1), core.FmtP(r.BeneficiaryAddressLine2), core.FmtP(r.BeneficiaryAddressLine3), core.FmtP(r.BeneficiaryName), core.FmtP(r.BeneficiaryReference), core.FmtP(r.Description), core.FmtP(r.InputMessageAccountabilityData), core.FmtP(r.OriginatorAddressLine1), core.FmtP(r.OriginatorAddressLine2), core.FmtP(r.OriginatorAddressLine3), core.FmtP(r.OriginatorName), core.FmtP(r.OriginatorToBeneficiaryInformationLine1), core.FmtP(r.OriginatorToBeneficiaryInformationLine2), core.FmtP(r.OriginatorToBeneficiaryInformationLine3), core.FmtP(r.OriginatorToBeneficiaryInformationLine4), core.FmtP(r.OriginatorToBeneficiaryInformation))
}

type TransactionSourceInternalSource struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency   *TransactionSourceInternalSourceCurrency `pjson:"currency"`
	Reason     *TransactionSourceInternalSourceReason   `pjson:"reason"`
	jsonFields map[string]interface{}                   `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceInternalSource using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *TransactionSourceInternalSource) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceInternalSource into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceInternalSource) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceInternalSource) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
func (r *TransactionSourceInternalSource) GetCurrency() (Currency TransactionSourceInternalSourceCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *TransactionSourceInternalSource) GetReason() (Reason TransactionSourceInternalSourceReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

func (r TransactionSourceInternalSource) String() (result string) {
	return fmt.Sprintf("&TransactionSourceInternalSource{Amount:%s Currency:%s Reason:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.Reason))
}

type TransactionSourceInternalSourceCurrency string

const (
	TransactionSourceInternalSourceCurrencyCad TransactionSourceInternalSourceCurrency = "CAD"
	TransactionSourceInternalSourceCurrencyChf TransactionSourceInternalSourceCurrency = "CHF"
	TransactionSourceInternalSourceCurrencyEur TransactionSourceInternalSourceCurrency = "EUR"
	TransactionSourceInternalSourceCurrencyGbp TransactionSourceInternalSourceCurrency = "GBP"
	TransactionSourceInternalSourceCurrencyJpy TransactionSourceInternalSourceCurrency = "JPY"
	TransactionSourceInternalSourceCurrencyUsd TransactionSourceInternalSourceCurrency = "USD"
)

type TransactionSourceInternalSourceReason string

const (
	TransactionSourceInternalSourceReasonBankMigration      TransactionSourceInternalSourceReason = "bank_migration"
	TransactionSourceInternalSourceReasonCashback           TransactionSourceInternalSourceReason = "cashback"
	TransactionSourceInternalSourceReasonEmpyrealAdjustment TransactionSourceInternalSourceReason = "empyreal_adjustment"
	TransactionSourceInternalSourceReasonError              TransactionSourceInternalSourceReason = "error"
	TransactionSourceInternalSourceReasonErrorCorrection    TransactionSourceInternalSourceReason = "error_correction"
	TransactionSourceInternalSourceReasonFees               TransactionSourceInternalSourceReason = "fees"
	TransactionSourceInternalSourceReasonInterest           TransactionSourceInternalSourceReason = "interest"
	TransactionSourceInternalSourceReasonSampleFunds        TransactionSourceInternalSourceReason = "sample_funds"
	TransactionSourceInternalSourceReasonSampleFundsReturn  TransactionSourceInternalSourceReason = "sample_funds_return"
)

type TransactionSourceCardRouteRefund struct {
	// The refunded amount in the minor unit of the refunded currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
	// currency.
	Currency             *TransactionSourceCardRouteRefundCurrency `pjson:"currency"`
	MerchantAcceptorID   *string                                   `pjson:"merchant_acceptor_id"`
	MerchantCity         *string                                   `pjson:"merchant_city"`
	MerchantCountry      *string                                   `pjson:"merchant_country"`
	MerchantDescriptor   *string                                   `pjson:"merchant_descriptor"`
	MerchantState        *string                                   `pjson:"merchant_state"`
	MerchantCategoryCode *string                                   `pjson:"merchant_category_code"`
	jsonFields           map[string]interface{}                    `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceCardRouteRefund using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *TransactionSourceCardRouteRefund) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceCardRouteRefund into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceCardRouteRefund) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The refunded amount in the minor unit of the refunded currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceCardRouteRefund) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
// currency.
func (r *TransactionSourceCardRouteRefund) GetCurrency() (Currency TransactionSourceCardRouteRefundCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *TransactionSourceCardRouteRefund) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r != nil && r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
}

func (r *TransactionSourceCardRouteRefund) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *TransactionSourceCardRouteRefund) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *TransactionSourceCardRouteRefund) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *TransactionSourceCardRouteRefund) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *TransactionSourceCardRouteRefund) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

func (r TransactionSourceCardRouteRefund) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCardRouteRefund{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.MerchantAcceptorID), core.FmtP(r.MerchantCity), core.FmtP(r.MerchantCountry), core.FmtP(r.MerchantDescriptor), core.FmtP(r.MerchantState), core.FmtP(r.MerchantCategoryCode))
}

type TransactionSourceCardRouteRefundCurrency string

const (
	TransactionSourceCardRouteRefundCurrencyCad TransactionSourceCardRouteRefundCurrency = "CAD"
	TransactionSourceCardRouteRefundCurrencyChf TransactionSourceCardRouteRefundCurrency = "CHF"
	TransactionSourceCardRouteRefundCurrencyEur TransactionSourceCardRouteRefundCurrency = "EUR"
	TransactionSourceCardRouteRefundCurrencyGbp TransactionSourceCardRouteRefundCurrency = "GBP"
	TransactionSourceCardRouteRefundCurrencyJpy TransactionSourceCardRouteRefundCurrency = "JPY"
	TransactionSourceCardRouteRefundCurrencyUsd TransactionSourceCardRouteRefundCurrency = "USD"
)

type TransactionSourceCardRouteSettlement struct {
	// The settled amount in the minor unit of the settlement currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
	// currency.
	Currency             *TransactionSourceCardRouteSettlementCurrency `pjson:"currency"`
	MerchantAcceptorID   *string                                       `pjson:"merchant_acceptor_id"`
	MerchantCity         *string                                       `pjson:"merchant_city"`
	MerchantCountry      *string                                       `pjson:"merchant_country"`
	MerchantDescriptor   *string                                       `pjson:"merchant_descriptor"`
	MerchantState        *string                                       `pjson:"merchant_state"`
	MerchantCategoryCode *string                                       `pjson:"merchant_category_code"`
	jsonFields           map[string]interface{}                        `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceCardRouteSettlement using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceCardRouteSettlement) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceCardRouteSettlement into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceCardRouteSettlement) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The settled amount in the minor unit of the settlement currency. For dollars,
// for example, this is cents.
func (r *TransactionSourceCardRouteSettlement) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
// currency.
func (r *TransactionSourceCardRouteSettlement) GetCurrency() (Currency TransactionSourceCardRouteSettlementCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *TransactionSourceCardRouteSettlement) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r != nil && r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
}

func (r *TransactionSourceCardRouteSettlement) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *TransactionSourceCardRouteSettlement) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *TransactionSourceCardRouteSettlement) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *TransactionSourceCardRouteSettlement) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *TransactionSourceCardRouteSettlement) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

func (r TransactionSourceCardRouteSettlement) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCardRouteSettlement{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.MerchantAcceptorID), core.FmtP(r.MerchantCity), core.FmtP(r.MerchantCountry), core.FmtP(r.MerchantDescriptor), core.FmtP(r.MerchantState), core.FmtP(r.MerchantCategoryCode))
}

type TransactionSourceCardRouteSettlementCurrency string

const (
	TransactionSourceCardRouteSettlementCurrencyCad TransactionSourceCardRouteSettlementCurrency = "CAD"
	TransactionSourceCardRouteSettlementCurrencyChf TransactionSourceCardRouteSettlementCurrency = "CHF"
	TransactionSourceCardRouteSettlementCurrencyEur TransactionSourceCardRouteSettlementCurrency = "EUR"
	TransactionSourceCardRouteSettlementCurrencyGbp TransactionSourceCardRouteSettlementCurrency = "GBP"
	TransactionSourceCardRouteSettlementCurrencyJpy TransactionSourceCardRouteSettlementCurrency = "JPY"
	TransactionSourceCardRouteSettlementCurrencyUsd TransactionSourceCardRouteSettlementCurrency = "USD"
)

type TransactionSourceSampleFunds struct {
	// Where the sample funds came from.
	Originator *string                `pjson:"originator"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into TransactionSourceSampleFunds
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *TransactionSourceSampleFunds) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceSampleFunds into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *TransactionSourceSampleFunds) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Where the sample funds came from.
func (r *TransactionSourceSampleFunds) GetOriginator() (Originator string) {
	if r != nil && r.Originator != nil {
		Originator = *r.Originator
	}
	return
}

func (r TransactionSourceSampleFunds) String() (result string) {
	return fmt.Sprintf("&TransactionSourceSampleFunds{Originator:%s}", core.FmtP(r.Originator))
}

type TransactionSourceWireDrawdownPaymentIntention struct {
	// The transfer amount in USD cents.
	Amount             *int64                 `pjson:"amount"`
	AccountNumber      *string                `pjson:"account_number"`
	RoutingNumber      *string                `pjson:"routing_number"`
	MessageToRecipient *string                `pjson:"message_to_recipient"`
	TransferID         *string                `pjson:"transfer_id"`
	jsonFields         map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceWireDrawdownPaymentIntention using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceWireDrawdownPaymentIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceWireDrawdownPaymentIntention into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *TransactionSourceWireDrawdownPaymentIntention) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The transfer amount in USD cents.
func (r *TransactionSourceWireDrawdownPaymentIntention) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *TransactionSourceWireDrawdownPaymentIntention) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *TransactionSourceWireDrawdownPaymentIntention) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *TransactionSourceWireDrawdownPaymentIntention) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r *TransactionSourceWireDrawdownPaymentIntention) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r TransactionSourceWireDrawdownPaymentIntention) String() (result string) {
	return fmt.Sprintf("&TransactionSourceWireDrawdownPaymentIntention{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s TransferID:%s}", core.FmtP(r.Amount), core.FmtP(r.AccountNumber), core.FmtP(r.RoutingNumber), core.FmtP(r.MessageToRecipient), core.FmtP(r.TransferID))
}

type TransactionSourceWireDrawdownPaymentRejection struct {
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceWireDrawdownPaymentRejection using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceWireDrawdownPaymentRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceWireDrawdownPaymentRejection into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *TransactionSourceWireDrawdownPaymentRejection) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

func (r *TransactionSourceWireDrawdownPaymentRejection) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r TransactionSourceWireDrawdownPaymentRejection) String() (result string) {
	return fmt.Sprintf("&TransactionSourceWireDrawdownPaymentRejection{TransferID:%s}", core.FmtP(r.TransferID))
}

type TransactionSourceWireTransferIntention struct {
	// The transfer amount in USD cents.
	Amount *int64 `pjson:"amount"`
	// The destination account number.
	AccountNumber *string `pjson:"account_number"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber *string `pjson:"routing_number"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient *string                `pjson:"message_to_recipient"`
	TransferID         *string                `pjson:"transfer_id"`
	jsonFields         map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceWireTransferIntention using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceWireTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceWireTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceWireTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The transfer amount in USD cents.
func (r *TransactionSourceWireTransferIntention) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The destination account number.
func (r *TransactionSourceWireTransferIntention) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r *TransactionSourceWireTransferIntention) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The message that will show on the recipient's bank statement.
func (r *TransactionSourceWireTransferIntention) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r *TransactionSourceWireTransferIntention) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r TransactionSourceWireTransferIntention) String() (result string) {
	return fmt.Sprintf("&TransactionSourceWireTransferIntention{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s TransferID:%s}", core.FmtP(r.Amount), core.FmtP(r.AccountNumber), core.FmtP(r.RoutingNumber), core.FmtP(r.MessageToRecipient), core.FmtP(r.TransferID))
}

type TransactionSourceWireTransferRejection struct {
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceWireTransferRejection using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceWireTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionSourceWireTransferRejection into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceWireTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

func (r *TransactionSourceWireTransferRejection) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r TransactionSourceWireTransferRejection) String() (result string) {
	return fmt.Sprintf("&TransactionSourceWireTransferRejection{TransferID:%s}", core.FmtP(r.TransferID))
}

type TransactionType string

const (
	TransactionTypeTransaction TransactionType = "transaction"
)

type TransactionListParams struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `query:"limit"`
	// Filter Transactions for those belonging to the specified Account.
	AccountID *string `query:"account_id"`
	// Filter Transactions for those belonging to the specified route.
	RouteID    *string                          `query:"route_id"`
	CreatedAt  *TransactionsListParamsCreatedAt `query:"created_at"`
	Category   *TransactionsListParamsCategory  `query:"category"`
	jsonFields map[string]interface{}           `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into TransactionListParams using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *TransactionListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionListParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *TransactionListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Return the page of entries after this one.
func (r *TransactionListParams) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *TransactionListParams) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Transactions for those belonging to the specified Account.
func (r *TransactionListParams) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// Filter Transactions for those belonging to the specified route.
func (r *TransactionListParams) GetRouteID() (RouteID string) {
	if r != nil && r.RouteID != nil {
		RouteID = *r.RouteID
	}
	return
}

func (r *TransactionListParams) GetCreatedAt() (CreatedAt TransactionsListParamsCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r *TransactionListParams) GetCategory() (Category TransactionsListParamsCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

func (r TransactionListParams) String() (result string) {
	return fmt.Sprintf("&TransactionListParams{Cursor:%s Limit:%s AccountID:%s RouteID:%s CreatedAt:%s Category:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), core.FmtP(r.AccountID), core.FmtP(r.RouteID), r.CreatedAt, r.Category)
}

type TransactionsListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *string `pjson:"after"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *string `pjson:"before"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *string `pjson:"on_or_after"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *string                `pjson:"on_or_before"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionsListParamsCreatedAt using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *TransactionsListParamsCreatedAt) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionsListParamsCreatedAt into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TransactionsListParamsCreatedAt) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *TransactionsListParamsCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *TransactionsListParamsCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *TransactionsListParamsCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *TransactionsListParamsCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

func (r TransactionsListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&TransactionsListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", core.FmtP(r.After), core.FmtP(r.Before), core.FmtP(r.OnOrAfter), core.FmtP(r.OnOrBefore))
}

type TransactionsListParamsCategory struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In         *[]TransactionsListParamsCategoryIn `pjson:"in"`
	jsonFields map[string]interface{}              `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionsListParamsCategory using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *TransactionsListParamsCategory) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionsListParamsCategory into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TransactionsListParamsCategory) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Return results whose value is in the provided list. For GET requests, this
// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
func (r *TransactionsListParamsCategory) GetIn() (In []TransactionsListParamsCategoryIn) {
	if r != nil && r.In != nil {
		In = *r.In
	}
	return
}

func (r TransactionsListParamsCategory) String() (result string) {
	return fmt.Sprintf("&TransactionsListParamsCategory{In:%s}", core.Fmt(r.In))
}

type TransactionsListParamsCategoryIn string

const (
	TransactionsListParamsCategoryInAccountTransferIntention                    TransactionsListParamsCategoryIn = "account_transfer_intention"
	TransactionsListParamsCategoryInACHCheckConversionReturn                    TransactionsListParamsCategoryIn = "ach_check_conversion_return"
	TransactionsListParamsCategoryInACHCheckConversion                          TransactionsListParamsCategoryIn = "ach_check_conversion"
	TransactionsListParamsCategoryInACHTransferIntention                        TransactionsListParamsCategoryIn = "ach_transfer_intention"
	TransactionsListParamsCategoryInACHTransferRejection                        TransactionsListParamsCategoryIn = "ach_transfer_rejection"
	TransactionsListParamsCategoryInACHTransferReturn                           TransactionsListParamsCategoryIn = "ach_transfer_return"
	TransactionsListParamsCategoryInCardDisputeAcceptance                       TransactionsListParamsCategoryIn = "card_dispute_acceptance"
	TransactionsListParamsCategoryInCardRefund                                  TransactionsListParamsCategoryIn = "card_refund"
	TransactionsListParamsCategoryInCardSettlement                              TransactionsListParamsCategoryIn = "card_settlement"
	TransactionsListParamsCategoryInCheckDepositAcceptance                      TransactionsListParamsCategoryIn = "check_deposit_acceptance"
	TransactionsListParamsCategoryInCheckDepositReturn                          TransactionsListParamsCategoryIn = "check_deposit_return"
	TransactionsListParamsCategoryInCheckTransferIntention                      TransactionsListParamsCategoryIn = "check_transfer_intention"
	TransactionsListParamsCategoryInCheckTransferReturn                         TransactionsListParamsCategoryIn = "check_transfer_return"
	TransactionsListParamsCategoryInCheckTransferRejection                      TransactionsListParamsCategoryIn = "check_transfer_rejection"
	TransactionsListParamsCategoryInCheckTransferStopPaymentRequest             TransactionsListParamsCategoryIn = "check_transfer_stop_payment_request"
	TransactionsListParamsCategoryInDisputeResolution                           TransactionsListParamsCategoryIn = "dispute_resolution"
	TransactionsListParamsCategoryInEmpyrealCashDeposit                         TransactionsListParamsCategoryIn = "empyreal_cash_deposit"
	TransactionsListParamsCategoryInInboundACHTransfer                          TransactionsListParamsCategoryIn = "inbound_ach_transfer"
	TransactionsListParamsCategoryInInboundCheck                                TransactionsListParamsCategoryIn = "inbound_check"
	TransactionsListParamsCategoryInInboundInternationalACHTransfer             TransactionsListParamsCategoryIn = "inbound_international_ach_transfer"
	TransactionsListParamsCategoryInInboundRealTimePaymentsTransferConfirmation TransactionsListParamsCategoryIn = "inbound_real_time_payments_transfer_confirmation"
	TransactionsListParamsCategoryInInboundWireDrawdownPaymentReversal          TransactionsListParamsCategoryIn = "inbound_wire_drawdown_payment_reversal"
	TransactionsListParamsCategoryInInboundWireDrawdownPayment                  TransactionsListParamsCategoryIn = "inbound_wire_drawdown_payment"
	TransactionsListParamsCategoryInInboundWireReversal                         TransactionsListParamsCategoryIn = "inbound_wire_reversal"
	TransactionsListParamsCategoryInInboundWireTransfer                         TransactionsListParamsCategoryIn = "inbound_wire_transfer"
	TransactionsListParamsCategoryInInternalGeneralLedgerTransaction            TransactionsListParamsCategoryIn = "internal_general_ledger_transaction"
	TransactionsListParamsCategoryInInternalSource                              TransactionsListParamsCategoryIn = "internal_source"
	TransactionsListParamsCategoryInCardRouteRefund                             TransactionsListParamsCategoryIn = "card_route_refund"
	TransactionsListParamsCategoryInCardRouteSettlement                         TransactionsListParamsCategoryIn = "card_route_settlement"
	TransactionsListParamsCategoryInRealTimePaymentsTransferAcknowledgement     TransactionsListParamsCategoryIn = "real_time_payments_transfer_acknowledgement"
	TransactionsListParamsCategoryInSampleFunds                                 TransactionsListParamsCategoryIn = "sample_funds"
	TransactionsListParamsCategoryInWireDrawdownPaymentIntention                TransactionsListParamsCategoryIn = "wire_drawdown_payment_intention"
	TransactionsListParamsCategoryInWireDrawdownPaymentRejection                TransactionsListParamsCategoryIn = "wire_drawdown_payment_rejection"
	TransactionsListParamsCategoryInWireTransferIntention                       TransactionsListParamsCategoryIn = "wire_transfer_intention"
	TransactionsListParamsCategoryInWireTransferRejection                       TransactionsListParamsCategoryIn = "wire_transfer_rejection"
	TransactionsListParamsCategoryInOther                                       TransactionsListParamsCategoryIn = "other"
)

type TransactionList struct {
	// The contents of the list.
	Data *[]Transaction `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into TransactionList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *TransactionList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes TransactionList into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *TransactionList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The contents of the list.
func (r *TransactionList) GetData() (Data []Transaction) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *TransactionList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r TransactionList) String() (result string) {
	return fmt.Sprintf("&TransactionList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type TransactionsPage struct {
	*pagination.Page[Transaction]
}

func (r *TransactionsPage) Transaction() *Transaction {
	return r.Current()
}

func (r *TransactionsPage) NextPage() (*TransactionsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &TransactionsPage{page}, nil
	}
}
