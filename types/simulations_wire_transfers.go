package types

import (
	"fmt"
	"time"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
)

type WireTransferSimulation struct {
	// If the Wire Transfer attempt succeeds, this will contain the resulting
	// [Transaction](#transactions) object. The Transaction's `source` will be of
	// `category: inbound_wire_transfer`.
	Transaction *WireTransferSimulationTransaction `pjson:"transaction"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_wire_transfer_simulation_result`.
	Type       *WireTransferSimulationType `pjson:"type"`
	jsonFields map[string]interface{}      `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into WireTransferSimulation using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulation) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferSimulation into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulation) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// If the Wire Transfer attempt succeeds, this will contain the resulting
// [Transaction](#transactions) object. The Transaction's `source` will be of
// `category: inbound_wire_transfer`.
func (r WireTransferSimulation) GetTransaction() (Transaction WireTransferSimulationTransaction) {
	if r.Transaction != nil {
		Transaction = *r.Transaction
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `inbound_wire_transfer_simulation_result`.
func (r WireTransferSimulation) GetType() (Type WireTransferSimulationType) {
	if r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r WireTransferSimulation) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulation{Transaction:%s Type:%s}", r.Transaction, core.FmtP(r.Type))
}

type WireTransferSimulationTransaction struct {
	// The identifier for the Account the Transaction belongs to.
	AccountID *string `pjson:"account_id"`
	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transcation's
	// Account.
	Currency *WireTransferSimulationTransactionCurrency `pjson:"currency"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt *time.Time `pjson:"created_at" format:"2006-01-02T15:04:05Z07:00"`
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
	Source *WireTransferSimulationTransactionSource `pjson:"source"`
	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type       *WireTransferSimulationTransactionType `pjson:"type"`
	jsonFields map[string]interface{}                 `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransaction using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransaction) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferSimulationTransaction into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransaction) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier for the Account the Transaction belongs to.
func (r WireTransferSimulationTransaction) GetAccountID() (AccountID string) {
	if r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The Transaction amount in the minor unit of its currency. For dollars, for
// example, this is cents.
func (r WireTransferSimulationTransaction) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// Transaction's currency. This will match the currency on the Transcation's
// Account.
func (r WireTransferSimulationTransaction) GetCurrency() (Currency WireTransferSimulationTransactionCurrency) {
	if r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
// Transaction occured.
func (r WireTransferSimulationTransaction) GetCreatedAt() (CreatedAt time.Time) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// For a Transaction related to a transfer, this is the description you provide.
// For a Transaction related to a payment, this is the description the vendor
// provides.
func (r WireTransferSimulationTransaction) GetDescription() (Description string) {
	if r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Transaction identifier.
func (r WireTransferSimulationTransaction) GetID() (ID string) {
	if r.ID != nil {
		ID = *r.ID
	}
	return
}

// The identifier for the route this Transaction came through. Routes are things
// like cards and ACH details.
func (r WireTransferSimulationTransaction) GetRouteID() (RouteID string) {
	if r.RouteID != nil {
		RouteID = *r.RouteID
	}
	return
}

// The type of the route this Transaction came through.
func (r WireTransferSimulationTransaction) GetRouteType() (RouteType string) {
	if r.RouteType != nil {
		RouteType = *r.RouteType
	}
	return
}

// This is an object giving more details on the network-level event that caused the
// Transaction. Note that for backwards compatibility reasons, additional
// undocumented keys may appear in this object. These should be treated as
// deprecated and will be removed in the future.
func (r WireTransferSimulationTransaction) GetSource() (Source WireTransferSimulationTransactionSource) {
	if r.Source != nil {
		Source = *r.Source
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `transaction`.
func (r WireTransferSimulationTransaction) GetType() (Type WireTransferSimulationTransactionType) {
	if r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r WireTransferSimulationTransaction) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransaction{AccountID:%s Amount:%s Currency:%s CreatedAt:%s Description:%s ID:%s RouteID:%s RouteType:%s Source:%s Type:%s}", core.FmtP(r.AccountID), core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.CreatedAt), core.FmtP(r.Description), core.FmtP(r.ID), core.FmtP(r.RouteID), core.FmtP(r.RouteType), r.Source, core.FmtP(r.Type))
}

type WireTransferSimulationTransactionCurrency string

const (
	WireTransferSimulationTransactionCurrencyCad WireTransferSimulationTransactionCurrency = "CAD"
	WireTransferSimulationTransactionCurrencyChf WireTransferSimulationTransactionCurrency = "CHF"
	WireTransferSimulationTransactionCurrencyEur WireTransferSimulationTransactionCurrency = "EUR"
	WireTransferSimulationTransactionCurrencyGbp WireTransferSimulationTransactionCurrency = "GBP"
	WireTransferSimulationTransactionCurrencyJpy WireTransferSimulationTransactionCurrency = "JPY"
	WireTransferSimulationTransactionCurrencyUsd WireTransferSimulationTransactionCurrency = "USD"
)

type WireTransferSimulationTransactionSource struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category *WireTransferSimulationTransactionSourceCategory `pjson:"category"`
	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention *WireTransferSimulationTransactionSourceAccountTransferIntention `pjson:"account_transfer_intention"`
	// A ACH Check Conversion Return object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_check_conversion_return`.
	ACHCheckConversionReturn *WireTransferSimulationTransactionSourceACHCheckConversionReturn `pjson:"ach_check_conversion_return"`
	// A ACH Check Conversion object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_check_conversion`.
	ACHCheckConversion *WireTransferSimulationTransactionSourceACHCheckConversion `pjson:"ach_check_conversion"`
	// A ACH Transfer Intention object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention *WireTransferSimulationTransactionSourceACHTransferIntention `pjson:"ach_transfer_intention"`
	// A ACH Transfer Rejection object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection *WireTransferSimulationTransactionSourceACHTransferRejection `pjson:"ach_transfer_rejection"`
	// A ACH Transfer Return object. This field will be present in the JSON response if
	// and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn *WireTransferSimulationTransactionSourceACHTransferReturn `pjson:"ach_transfer_return"`
	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance *WireTransferSimulationTransactionSourceCardDisputeAcceptance `pjson:"card_dispute_acceptance"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund *WireTransferSimulationTransactionSourceCardRefund `pjson:"card_refund"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement *WireTransferSimulationTransactionSourceCardSettlement `pjson:"card_settlement"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance *WireTransferSimulationTransactionSourceCheckDepositAcceptance `pjson:"check_deposit_acceptance"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn *WireTransferSimulationTransactionSourceCheckDepositReturn `pjson:"check_deposit_return"`
	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention *WireTransferSimulationTransactionSourceCheckTransferIntention `pjson:"check_transfer_intention"`
	// A Check Transfer Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_transfer_return`.
	CheckTransferReturn *WireTransferSimulationTransactionSourceCheckTransferReturn `pjson:"check_transfer_return"`
	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection *WireTransferSimulationTransactionSourceCheckTransferRejection `pjson:"check_transfer_rejection"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest *WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest `pjson:"check_transfer_stop_payment_request"`
	// A Dispute Resolution object. This field will be present in the JSON response if
	// and only if `category` is equal to `dispute_resolution`.
	DisputeResolution *WireTransferSimulationTransactionSourceDisputeResolution `pjson:"dispute_resolution"`
	// A Empyreal Cash Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `empyreal_cash_deposit`.
	EmpyrealCashDeposit *WireTransferSimulationTransactionSourceEmpyrealCashDeposit `pjson:"empyreal_cash_deposit"`
	// A Inbound ACH Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer *WireTransferSimulationTransactionSourceInboundACHTransfer `pjson:"inbound_ach_transfer"`
	// A Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck *WireTransferSimulationTransactionSourceInboundCheck `pjson:"inbound_check"`
	// A Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer `pjson:"inbound_international_ach_transfer"`
	// A Inbound Real Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation *WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation `pjson:"inbound_real_time_payments_transfer_confirmation"`
	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal *WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal `pjson:"inbound_wire_drawdown_payment_reversal"`
	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment `pjson:"inbound_wire_drawdown_payment"`
	// A Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal *WireTransferSimulationTransactionSourceInboundWireReversal `pjson:"inbound_wire_reversal"`
	// A Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer *WireTransferSimulationTransactionSourceInboundWireTransfer `pjson:"inbound_wire_transfer"`
	// A Interest Payment object. This field will be present in the JSON response if
	// and only if `category` is equal to `interest_payment`.
	InterestPayment *WireTransferSimulationTransactionSourceInterestPayment `pjson:"interest_payment"`
	// A Internal Source object. This field will be present in the JSON response if and
	// only if `category` is equal to `internal_source`.
	InternalSource *WireTransferSimulationTransactionSourceInternalSource `pjson:"internal_source"`
	// A Deprecated Card Refund object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_route_refund`.
	CardRouteRefund *WireTransferSimulationTransactionSourceCardRouteRefund `pjson:"card_route_refund"`
	// A Deprecated Card Settlement object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_settlement`.
	CardRouteSettlement *WireTransferSimulationTransactionSourceCardRouteSettlement `pjson:"card_route_settlement"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds *WireTransferSimulationTransactionSourceSampleFunds `pjson:"sample_funds"`
	// A Wire Drawdown Payment Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_intention`.
	WireDrawdownPaymentIntention *WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention `pjson:"wire_drawdown_payment_intention"`
	// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_rejection`.
	WireDrawdownPaymentRejection *WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection `pjson:"wire_drawdown_payment_rejection"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention *WireTransferSimulationTransactionSourceWireTransferIntention `pjson:"wire_transfer_intention"`
	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection *WireTransferSimulationTransactionSourceWireTransferRejection `pjson:"wire_transfer_rejection"`
	jsonFields            map[string]interface{}                                        `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSource using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferSimulationTransactionSource into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSource) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The type of transaction that took place. We may add additional possible values
// for this enum over time; your application should be able to handle such
// additions gracefully.
func (r WireTransferSimulationTransactionSource) GetCategory() (Category WireTransferSimulationTransactionSourceCategory) {
	if r.Category != nil {
		Category = *r.Category
	}
	return
}

// A Account Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_intention`.
func (r WireTransferSimulationTransactionSource) GetAccountTransferIntention() (AccountTransferIntention WireTransferSimulationTransactionSourceAccountTransferIntention) {
	if r.AccountTransferIntention != nil {
		AccountTransferIntention = *r.AccountTransferIntention
	}
	return
}

// A ACH Check Conversion Return object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_check_conversion_return`.
func (r WireTransferSimulationTransactionSource) GetACHCheckConversionReturn() (ACHCheckConversionReturn WireTransferSimulationTransactionSourceACHCheckConversionReturn) {
	if r.ACHCheckConversionReturn != nil {
		ACHCheckConversionReturn = *r.ACHCheckConversionReturn
	}
	return
}

// A ACH Check Conversion object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_check_conversion`.
func (r WireTransferSimulationTransactionSource) GetACHCheckConversion() (ACHCheckConversion WireTransferSimulationTransactionSourceACHCheckConversion) {
	if r.ACHCheckConversion != nil {
		ACHCheckConversion = *r.ACHCheckConversion
	}
	return
}

// A ACH Transfer Intention object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_intention`.
func (r WireTransferSimulationTransactionSource) GetACHTransferIntention() (ACHTransferIntention WireTransferSimulationTransactionSourceACHTransferIntention) {
	if r.ACHTransferIntention != nil {
		ACHTransferIntention = *r.ACHTransferIntention
	}
	return
}

// A ACH Transfer Rejection object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_rejection`.
func (r WireTransferSimulationTransactionSource) GetACHTransferRejection() (ACHTransferRejection WireTransferSimulationTransactionSourceACHTransferRejection) {
	if r.ACHTransferRejection != nil {
		ACHTransferRejection = *r.ACHTransferRejection
	}
	return
}

// A ACH Transfer Return object. This field will be present in the JSON response if
// and only if `category` is equal to `ach_transfer_return`.
func (r WireTransferSimulationTransactionSource) GetACHTransferReturn() (ACHTransferReturn WireTransferSimulationTransactionSourceACHTransferReturn) {
	if r.ACHTransferReturn != nil {
		ACHTransferReturn = *r.ACHTransferReturn
	}
	return
}

// A Card Dispute Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `card_dispute_acceptance`.
func (r WireTransferSimulationTransactionSource) GetCardDisputeAcceptance() (CardDisputeAcceptance WireTransferSimulationTransactionSourceCardDisputeAcceptance) {
	if r.CardDisputeAcceptance != nil {
		CardDisputeAcceptance = *r.CardDisputeAcceptance
	}
	return
}

// A Card Refund object. This field will be present in the JSON response if and
// only if `category` is equal to `card_refund`.
func (r WireTransferSimulationTransactionSource) GetCardRefund() (CardRefund WireTransferSimulationTransactionSourceCardRefund) {
	if r.CardRefund != nil {
		CardRefund = *r.CardRefund
	}
	return
}

// A Card Settlement object. This field will be present in the JSON response if and
// only if `category` is equal to `card_settlement`.
func (r WireTransferSimulationTransactionSource) GetCardSettlement() (CardSettlement WireTransferSimulationTransactionSourceCardSettlement) {
	if r.CardSettlement != nil {
		CardSettlement = *r.CardSettlement
	}
	return
}

// A Check Deposit Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_acceptance`.
func (r WireTransferSimulationTransactionSource) GetCheckDepositAcceptance() (CheckDepositAcceptance WireTransferSimulationTransactionSourceCheckDepositAcceptance) {
	if r.CheckDepositAcceptance != nil {
		CheckDepositAcceptance = *r.CheckDepositAcceptance
	}
	return
}

// A Check Deposit Return object. This field will be present in the JSON response
// if and only if `category` is equal to `check_deposit_return`.
func (r WireTransferSimulationTransactionSource) GetCheckDepositReturn() (CheckDepositReturn WireTransferSimulationTransactionSourceCheckDepositReturn) {
	if r.CheckDepositReturn != nil {
		CheckDepositReturn = *r.CheckDepositReturn
	}
	return
}

// A Check Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_intention`.
func (r WireTransferSimulationTransactionSource) GetCheckTransferIntention() (CheckTransferIntention WireTransferSimulationTransactionSourceCheckTransferIntention) {
	if r.CheckTransferIntention != nil {
		CheckTransferIntention = *r.CheckTransferIntention
	}
	return
}

// A Check Transfer Return object. This field will be present in the JSON response
// if and only if `category` is equal to `check_transfer_return`.
func (r WireTransferSimulationTransactionSource) GetCheckTransferReturn() (CheckTransferReturn WireTransferSimulationTransactionSourceCheckTransferReturn) {
	if r.CheckTransferReturn != nil {
		CheckTransferReturn = *r.CheckTransferReturn
	}
	return
}

// A Check Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_rejection`.
func (r WireTransferSimulationTransactionSource) GetCheckTransferRejection() (CheckTransferRejection WireTransferSimulationTransactionSourceCheckTransferRejection) {
	if r.CheckTransferRejection != nil {
		CheckTransferRejection = *r.CheckTransferRejection
	}
	return
}

// A Check Transfer Stop Payment Request object. This field will be present in the
// JSON response if and only if `category` is equal to
// `check_transfer_stop_payment_request`.
func (r WireTransferSimulationTransactionSource) GetCheckTransferStopPaymentRequest() (CheckTransferStopPaymentRequest WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) {
	if r.CheckTransferStopPaymentRequest != nil {
		CheckTransferStopPaymentRequest = *r.CheckTransferStopPaymentRequest
	}
	return
}

// A Dispute Resolution object. This field will be present in the JSON response if
// and only if `category` is equal to `dispute_resolution`.
func (r WireTransferSimulationTransactionSource) GetDisputeResolution() (DisputeResolution WireTransferSimulationTransactionSourceDisputeResolution) {
	if r.DisputeResolution != nil {
		DisputeResolution = *r.DisputeResolution
	}
	return
}

// A Empyreal Cash Deposit object. This field will be present in the JSON response
// if and only if `category` is equal to `empyreal_cash_deposit`.
func (r WireTransferSimulationTransactionSource) GetEmpyrealCashDeposit() (EmpyrealCashDeposit WireTransferSimulationTransactionSourceEmpyrealCashDeposit) {
	if r.EmpyrealCashDeposit != nil {
		EmpyrealCashDeposit = *r.EmpyrealCashDeposit
	}
	return
}

// A Inbound ACH Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_ach_transfer`.
func (r WireTransferSimulationTransactionSource) GetInboundACHTransfer() (InboundACHTransfer WireTransferSimulationTransactionSourceInboundACHTransfer) {
	if r.InboundACHTransfer != nil {
		InboundACHTransfer = *r.InboundACHTransfer
	}
	return
}

// A Inbound Check object. This field will be present in the JSON response if and
// only if `category` is equal to `inbound_check`.
func (r WireTransferSimulationTransactionSource) GetInboundCheck() (InboundCheck WireTransferSimulationTransactionSourceInboundCheck) {
	if r.InboundCheck != nil {
		InboundCheck = *r.InboundCheck
	}
	return
}

// A Inbound International ACH Transfer object. This field will be present in the
// JSON response if and only if `category` is equal to
// `inbound_international_ach_transfer`.
func (r WireTransferSimulationTransactionSource) GetInboundInternationalACHTransfer() (InboundInternationalACHTransfer WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) {
	if r.InboundInternationalACHTransfer != nil {
		InboundInternationalACHTransfer = *r.InboundInternationalACHTransfer
	}
	return
}

// A Inbound Real Time Payments Transfer Confirmation object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_confirmation`.
func (r WireTransferSimulationTransactionSource) GetInboundRealTimePaymentsTransferConfirmation() (InboundRealTimePaymentsTransferConfirmation WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) {
	if r.InboundRealTimePaymentsTransferConfirmation != nil {
		InboundRealTimePaymentsTransferConfirmation = *r.InboundRealTimePaymentsTransferConfirmation
	}
	return
}

// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
// the JSON response if and only if `category` is equal to
// `inbound_wire_drawdown_payment_reversal`.
func (r WireTransferSimulationTransactionSource) GetInboundWireDrawdownPaymentReversal() (InboundWireDrawdownPaymentReversal WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) {
	if r.InboundWireDrawdownPaymentReversal != nil {
		InboundWireDrawdownPaymentReversal = *r.InboundWireDrawdownPaymentReversal
	}
	return
}

// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
func (r WireTransferSimulationTransactionSource) GetInboundWireDrawdownPayment() (InboundWireDrawdownPayment WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) {
	if r.InboundWireDrawdownPayment != nil {
		InboundWireDrawdownPayment = *r.InboundWireDrawdownPayment
	}
	return
}

// A Inbound Wire Reversal object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_reversal`.
func (r WireTransferSimulationTransactionSource) GetInboundWireReversal() (InboundWireReversal WireTransferSimulationTransactionSourceInboundWireReversal) {
	if r.InboundWireReversal != nil {
		InboundWireReversal = *r.InboundWireReversal
	}
	return
}

// A Inbound Wire Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_transfer`.
func (r WireTransferSimulationTransactionSource) GetInboundWireTransfer() (InboundWireTransfer WireTransferSimulationTransactionSourceInboundWireTransfer) {
	if r.InboundWireTransfer != nil {
		InboundWireTransfer = *r.InboundWireTransfer
	}
	return
}

// A Interest Payment object. This field will be present in the JSON response if
// and only if `category` is equal to `interest_payment`.
func (r WireTransferSimulationTransactionSource) GetInterestPayment() (InterestPayment WireTransferSimulationTransactionSourceInterestPayment) {
	if r.InterestPayment != nil {
		InterestPayment = *r.InterestPayment
	}
	return
}

// A Internal Source object. This field will be present in the JSON response if and
// only if `category` is equal to `internal_source`.
func (r WireTransferSimulationTransactionSource) GetInternalSource() (InternalSource WireTransferSimulationTransactionSourceInternalSource) {
	if r.InternalSource != nil {
		InternalSource = *r.InternalSource
	}
	return
}

// A Deprecated Card Refund object. This field will be present in the JSON response
// if and only if `category` is equal to `card_route_refund`.
func (r WireTransferSimulationTransactionSource) GetCardRouteRefund() (CardRouteRefund WireTransferSimulationTransactionSourceCardRouteRefund) {
	if r.CardRouteRefund != nil {
		CardRouteRefund = *r.CardRouteRefund
	}
	return
}

// A Deprecated Card Settlement object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_settlement`.
func (r WireTransferSimulationTransactionSource) GetCardRouteSettlement() (CardRouteSettlement WireTransferSimulationTransactionSourceCardRouteSettlement) {
	if r.CardRouteSettlement != nil {
		CardRouteSettlement = *r.CardRouteSettlement
	}
	return
}

// A Sample Funds object. This field will be present in the JSON response if and
// only if `category` is equal to `sample_funds`.
func (r WireTransferSimulationTransactionSource) GetSampleFunds() (SampleFunds WireTransferSimulationTransactionSourceSampleFunds) {
	if r.SampleFunds != nil {
		SampleFunds = *r.SampleFunds
	}
	return
}

// A Wire Drawdown Payment Intention object. This field will be present in the JSON
// response if and only if `category` is equal to
// `wire_drawdown_payment_intention`.
func (r WireTransferSimulationTransactionSource) GetWireDrawdownPaymentIntention() (WireDrawdownPaymentIntention WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention) {
	if r.WireDrawdownPaymentIntention != nil {
		WireDrawdownPaymentIntention = *r.WireDrawdownPaymentIntention
	}
	return
}

// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to
// `wire_drawdown_payment_rejection`.
func (r WireTransferSimulationTransactionSource) GetWireDrawdownPaymentRejection() (WireDrawdownPaymentRejection WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection) {
	if r.WireDrawdownPaymentRejection != nil {
		WireDrawdownPaymentRejection = *r.WireDrawdownPaymentRejection
	}
	return
}

// A Wire Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_intention`.
func (r WireTransferSimulationTransactionSource) GetWireTransferIntention() (WireTransferIntention WireTransferSimulationTransactionSourceWireTransferIntention) {
	if r.WireTransferIntention != nil {
		WireTransferIntention = *r.WireTransferIntention
	}
	return
}

// A Wire Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_rejection`.
func (r WireTransferSimulationTransactionSource) GetWireTransferRejection() (WireTransferRejection WireTransferSimulationTransactionSourceWireTransferRejection) {
	if r.WireTransferRejection != nil {
		WireTransferRejection = *r.WireTransferRejection
	}
	return
}

func (r WireTransferSimulationTransactionSource) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSource{Category:%s AccountTransferIntention:%s ACHCheckConversionReturn:%s ACHCheckConversion:%s ACHTransferIntention:%s ACHTransferRejection:%s ACHTransferReturn:%s CardDisputeAcceptance:%s CardRefund:%s CardSettlement:%s CheckDepositAcceptance:%s CheckDepositReturn:%s CheckTransferIntention:%s CheckTransferReturn:%s CheckTransferRejection:%s CheckTransferStopPaymentRequest:%s DisputeResolution:%s EmpyrealCashDeposit:%s InboundACHTransfer:%s InboundCheck:%s InboundInternationalACHTransfer:%s InboundRealTimePaymentsTransferConfirmation:%s InboundWireDrawdownPaymentReversal:%s InboundWireDrawdownPayment:%s InboundWireReversal:%s InboundWireTransfer:%s InterestPayment:%s InternalSource:%s CardRouteRefund:%s CardRouteSettlement:%s SampleFunds:%s WireDrawdownPaymentIntention:%s WireDrawdownPaymentRejection:%s WireTransferIntention:%s WireTransferRejection:%s}", core.FmtP(r.Category), r.AccountTransferIntention, r.ACHCheckConversionReturn, r.ACHCheckConversion, r.ACHTransferIntention, r.ACHTransferRejection, r.ACHTransferReturn, r.CardDisputeAcceptance, r.CardRefund, r.CardSettlement, r.CheckDepositAcceptance, r.CheckDepositReturn, r.CheckTransferIntention, r.CheckTransferReturn, r.CheckTransferRejection, r.CheckTransferStopPaymentRequest, r.DisputeResolution, r.EmpyrealCashDeposit, r.InboundACHTransfer, r.InboundCheck, r.InboundInternationalACHTransfer, r.InboundRealTimePaymentsTransferConfirmation, r.InboundWireDrawdownPaymentReversal, r.InboundWireDrawdownPayment, r.InboundWireReversal, r.InboundWireTransfer, r.InterestPayment, r.InternalSource, r.CardRouteRefund, r.CardRouteSettlement, r.SampleFunds, r.WireDrawdownPaymentIntention, r.WireDrawdownPaymentRejection, r.WireTransferIntention, r.WireTransferRejection)
}

type WireTransferSimulationTransactionSourceCategory string

const (
	WireTransferSimulationTransactionSourceCategoryAccountTransferIntention                    WireTransferSimulationTransactionSourceCategory = "account_transfer_intention"
	WireTransferSimulationTransactionSourceCategoryACHCheckConversionReturn                    WireTransferSimulationTransactionSourceCategory = "ach_check_conversion_return"
	WireTransferSimulationTransactionSourceCategoryACHCheckConversion                          WireTransferSimulationTransactionSourceCategory = "ach_check_conversion"
	WireTransferSimulationTransactionSourceCategoryACHTransferIntention                        WireTransferSimulationTransactionSourceCategory = "ach_transfer_intention"
	WireTransferSimulationTransactionSourceCategoryACHTransferRejection                        WireTransferSimulationTransactionSourceCategory = "ach_transfer_rejection"
	WireTransferSimulationTransactionSourceCategoryACHTransferReturn                           WireTransferSimulationTransactionSourceCategory = "ach_transfer_return"
	WireTransferSimulationTransactionSourceCategoryCardDisputeAcceptance                       WireTransferSimulationTransactionSourceCategory = "card_dispute_acceptance"
	WireTransferSimulationTransactionSourceCategoryCardRefund                                  WireTransferSimulationTransactionSourceCategory = "card_refund"
	WireTransferSimulationTransactionSourceCategoryCardSettlement                              WireTransferSimulationTransactionSourceCategory = "card_settlement"
	WireTransferSimulationTransactionSourceCategoryCheckDepositAcceptance                      WireTransferSimulationTransactionSourceCategory = "check_deposit_acceptance"
	WireTransferSimulationTransactionSourceCategoryCheckDepositReturn                          WireTransferSimulationTransactionSourceCategory = "check_deposit_return"
	WireTransferSimulationTransactionSourceCategoryCheckTransferIntention                      WireTransferSimulationTransactionSourceCategory = "check_transfer_intention"
	WireTransferSimulationTransactionSourceCategoryCheckTransferReturn                         WireTransferSimulationTransactionSourceCategory = "check_transfer_return"
	WireTransferSimulationTransactionSourceCategoryCheckTransferRejection                      WireTransferSimulationTransactionSourceCategory = "check_transfer_rejection"
	WireTransferSimulationTransactionSourceCategoryCheckTransferStopPaymentRequest             WireTransferSimulationTransactionSourceCategory = "check_transfer_stop_payment_request"
	WireTransferSimulationTransactionSourceCategoryDisputeResolution                           WireTransferSimulationTransactionSourceCategory = "dispute_resolution"
	WireTransferSimulationTransactionSourceCategoryEmpyrealCashDeposit                         WireTransferSimulationTransactionSourceCategory = "empyreal_cash_deposit"
	WireTransferSimulationTransactionSourceCategoryInboundACHTransfer                          WireTransferSimulationTransactionSourceCategory = "inbound_ach_transfer"
	WireTransferSimulationTransactionSourceCategoryInboundACHTransferReturnIntention           WireTransferSimulationTransactionSourceCategory = "inbound_ach_transfer_return_intention"
	WireTransferSimulationTransactionSourceCategoryInboundCheck                                WireTransferSimulationTransactionSourceCategory = "inbound_check"
	WireTransferSimulationTransactionSourceCategoryInboundInternationalACHTransfer             WireTransferSimulationTransactionSourceCategory = "inbound_international_ach_transfer"
	WireTransferSimulationTransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation WireTransferSimulationTransactionSourceCategory = "inbound_real_time_payments_transfer_confirmation"
	WireTransferSimulationTransactionSourceCategoryInboundWireDrawdownPaymentReversal          WireTransferSimulationTransactionSourceCategory = "inbound_wire_drawdown_payment_reversal"
	WireTransferSimulationTransactionSourceCategoryInboundWireDrawdownPayment                  WireTransferSimulationTransactionSourceCategory = "inbound_wire_drawdown_payment"
	WireTransferSimulationTransactionSourceCategoryInboundWireReversal                         WireTransferSimulationTransactionSourceCategory = "inbound_wire_reversal"
	WireTransferSimulationTransactionSourceCategoryInboundWireTransfer                         WireTransferSimulationTransactionSourceCategory = "inbound_wire_transfer"
	WireTransferSimulationTransactionSourceCategoryInterestPayment                             WireTransferSimulationTransactionSourceCategory = "interest_payment"
	WireTransferSimulationTransactionSourceCategoryInternalGeneralLedgerTransaction            WireTransferSimulationTransactionSourceCategory = "internal_general_ledger_transaction"
	WireTransferSimulationTransactionSourceCategoryInternalSource                              WireTransferSimulationTransactionSourceCategory = "internal_source"
	WireTransferSimulationTransactionSourceCategoryCardRouteRefund                             WireTransferSimulationTransactionSourceCategory = "card_route_refund"
	WireTransferSimulationTransactionSourceCategoryCardRouteSettlement                         WireTransferSimulationTransactionSourceCategory = "card_route_settlement"
	WireTransferSimulationTransactionSourceCategoryRealTimePaymentsTransferAcknowledgement     WireTransferSimulationTransactionSourceCategory = "real_time_payments_transfer_acknowledgement"
	WireTransferSimulationTransactionSourceCategorySampleFunds                                 WireTransferSimulationTransactionSourceCategory = "sample_funds"
	WireTransferSimulationTransactionSourceCategoryWireDrawdownPaymentIntention                WireTransferSimulationTransactionSourceCategory = "wire_drawdown_payment_intention"
	WireTransferSimulationTransactionSourceCategoryWireDrawdownPaymentRejection                WireTransferSimulationTransactionSourceCategory = "wire_drawdown_payment_rejection"
	WireTransferSimulationTransactionSourceCategoryWireTransferIntention                       WireTransferSimulationTransactionSourceCategory = "wire_transfer_intention"
	WireTransferSimulationTransactionSourceCategoryWireTransferRejection                       WireTransferSimulationTransactionSourceCategory = "wire_transfer_rejection"
	WireTransferSimulationTransactionSourceCategoryOther                                       WireTransferSimulationTransactionSourceCategory = "other"
)

type WireTransferSimulationTransactionSourceAccountTransferIntention struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency `pjson:"currency"`
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
// WireTransferSimulationTransactionSourceAccountTransferIntention using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceAccountTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceAccountTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceAccountTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r WireTransferSimulationTransactionSourceAccountTransferIntention) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r WireTransferSimulationTransactionSourceAccountTransferIntention) GetCurrency() (Currency WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency) {
	if r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The description you chose to give the transfer.
func (r WireTransferSimulationTransactionSourceAccountTransferIntention) GetDescription() (Description string) {
	if r.Description != nil {
		Description = *r.Description
	}
	return
}

// The identifier of the Account to where the Account Transfer was sent.
func (r WireTransferSimulationTransactionSourceAccountTransferIntention) GetDestinationAccountID() (DestinationAccountID string) {
	if r.DestinationAccountID != nil {
		DestinationAccountID = *r.DestinationAccountID
	}
	return
}

// The identifier of the Account from where the Account Transfer was sent.
func (r WireTransferSimulationTransactionSourceAccountTransferIntention) GetSourceAccountID() (SourceAccountID string) {
	if r.SourceAccountID != nil {
		SourceAccountID = *r.SourceAccountID
	}
	return
}

// The identifier of the Account Transfer that led to this Pending Transaction.
func (r WireTransferSimulationTransactionSourceAccountTransferIntention) GetTransferID() (TransferID string) {
	if r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r WireTransferSimulationTransactionSourceAccountTransferIntention) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceAccountTransferIntention{Amount:%s Currency:%s Description:%s DestinationAccountID:%s SourceAccountID:%s TransferID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.Description), core.FmtP(r.DestinationAccountID), core.FmtP(r.SourceAccountID), core.FmtP(r.TransferID))
}

type WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency string

const (
	WireTransferSimulationTransactionSourceAccountTransferIntentionCurrencyCad WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "CAD"
	WireTransferSimulationTransactionSourceAccountTransferIntentionCurrencyChf WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "CHF"
	WireTransferSimulationTransactionSourceAccountTransferIntentionCurrencyEur WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "EUR"
	WireTransferSimulationTransactionSourceAccountTransferIntentionCurrencyGbp WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "GBP"
	WireTransferSimulationTransactionSourceAccountTransferIntentionCurrencyJpy WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "JPY"
	WireTransferSimulationTransactionSourceAccountTransferIntentionCurrencyUsd WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "USD"
)

type WireTransferSimulationTransactionSourceACHCheckConversionReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// Why the transfer was returned.
	ReturnReasonCode *string                `pjson:"return_reason_code"`
	jsonFields       map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceACHCheckConversionReturn using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceACHCheckConversionReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceACHCheckConversionReturn into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceACHCheckConversionReturn) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r WireTransferSimulationTransactionSourceACHCheckConversionReturn) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// Why the transfer was returned.
func (r WireTransferSimulationTransactionSourceACHCheckConversionReturn) GetReturnReasonCode() (ReturnReasonCode string) {
	if r.ReturnReasonCode != nil {
		ReturnReasonCode = *r.ReturnReasonCode
	}
	return
}

func (r WireTransferSimulationTransactionSourceACHCheckConversionReturn) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceACHCheckConversionReturn{Amount:%s ReturnReasonCode:%s}", core.FmtP(r.Amount), core.FmtP(r.ReturnReasonCode))
}

type WireTransferSimulationTransactionSourceACHCheckConversion struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The identifier of the File containing an image of the returned check.
	FileID     *string                `pjson:"file_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceACHCheckConversion using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceACHCheckConversion) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceACHCheckConversion
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *WireTransferSimulationTransactionSourceACHCheckConversion) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r WireTransferSimulationTransactionSourceACHCheckConversion) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The identifier of the File containing an image of the returned check.
func (r WireTransferSimulationTransactionSourceACHCheckConversion) GetFileID() (FileID string) {
	if r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

func (r WireTransferSimulationTransactionSourceACHCheckConversion) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceACHCheckConversion{Amount:%s FileID:%s}", core.FmtP(r.Amount), core.FmtP(r.FileID))
}

type WireTransferSimulationTransactionSourceACHTransferIntention struct {
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
// WireTransferSimulationTransactionSourceACHTransferIntention using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceACHTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceACHTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceACHTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r WireTransferSimulationTransactionSourceACHTransferIntention) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r WireTransferSimulationTransactionSourceACHTransferIntention) GetAccountNumber() (AccountNumber string) {
	if r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r WireTransferSimulationTransactionSourceACHTransferIntention) GetRoutingNumber() (RoutingNumber string) {
	if r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r WireTransferSimulationTransactionSourceACHTransferIntention) GetStatementDescriptor() (StatementDescriptor string) {
	if r.StatementDescriptor != nil {
		StatementDescriptor = *r.StatementDescriptor
	}
	return
}

// The identifier of the ACH Transfer that led to this Transaction.
func (r WireTransferSimulationTransactionSourceACHTransferIntention) GetTransferID() (TransferID string) {
	if r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r WireTransferSimulationTransactionSourceACHTransferIntention) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceACHTransferIntention{Amount:%s AccountNumber:%s RoutingNumber:%s StatementDescriptor:%s TransferID:%s}", core.FmtP(r.Amount), core.FmtP(r.AccountNumber), core.FmtP(r.RoutingNumber), core.FmtP(r.StatementDescriptor), core.FmtP(r.TransferID))
}

type WireTransferSimulationTransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceACHTransferRejection using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceACHTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceACHTransferRejection into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceACHTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the ACH Transfer that led to this Transaction.
func (r WireTransferSimulationTransactionSourceACHTransferRejection) GetTransferID() (TransferID string) {
	if r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r WireTransferSimulationTransactionSourceACHTransferRejection) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceACHTransferRejection{TransferID:%s}", core.FmtP(r.TransferID))
}

type WireTransferSimulationTransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *time.Time `pjson:"created_at" format:"2006-01-02T15:04:05Z07:00"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode *WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode `pjson:"return_reason_code"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID *string `pjson:"transfer_id"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID *string                `pjson:"transaction_id"`
	jsonFields    map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceACHTransferReturn using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceACHTransferReturn
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *WireTransferSimulationTransactionSourceACHTransferReturn) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r WireTransferSimulationTransactionSourceACHTransferReturn) GetCreatedAt() (CreatedAt time.Time) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// Why the ACH Transfer was returned.
func (r WireTransferSimulationTransactionSourceACHTransferReturn) GetReturnReasonCode() (ReturnReasonCode WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode) {
	if r.ReturnReasonCode != nil {
		ReturnReasonCode = *r.ReturnReasonCode
	}
	return
}

// The identifier of the ACH Transfer associated with this return.
func (r WireTransferSimulationTransactionSourceACHTransferReturn) GetTransferID() (TransferID string) {
	if r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

// The identifier of the Tranasaction associated with this return.
func (r WireTransferSimulationTransactionSourceACHTransferReturn) GetTransactionID() (TransactionID string) {
	if r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

func (r WireTransferSimulationTransactionSourceACHTransferReturn) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceACHTransferReturn{CreatedAt:%s ReturnReasonCode:%s TransferID:%s TransactionID:%s}", core.FmtP(r.CreatedAt), core.FmtP(r.ReturnReasonCode), core.FmtP(r.TransferID), core.FmtP(r.TransactionID))
}

type WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode string

const (
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInsufficientFund                                          WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "insufficient_fund"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeNoAccount                                                 WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "no_account"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAccountClosed                                             WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "account_closed"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure                             WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction              WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver                              WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode   WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized                     WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodePaymentStopped                                            WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "payment_stopped"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeNonTransactionAccount                                     WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "non_transaction_account"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeUncollectedFunds                                          WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "uncollected_funds"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError                              WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAmountFieldError                                          WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "amount_field_error"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer                            WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber                                   WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeFileRecordEditCriteria                                    WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualName                                  WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest                                    WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAddendaError                                              WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "addenda_error"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeLimitedParticipationDfi                                   WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment              WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "incorrectly_coded_outbound_international_payment"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeOther                                                     WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "other"
)

type WireTransferSimulationTransactionSourceCardDisputeAcceptance struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was accepted.
	AcceptedAt *time.Time `pjson:"accepted_at" format:"2006-01-02T15:04:05Z07:00"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID *string `pjson:"card_dispute_id"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID *string                `pjson:"transaction_id"`
	jsonFields    map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCardDisputeAcceptance using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCardDisputeAcceptance) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceCardDisputeAcceptance into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceCardDisputeAcceptance) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card Dispute was accepted.
func (r WireTransferSimulationTransactionSourceCardDisputeAcceptance) GetAcceptedAt() (AcceptedAt time.Time) {
	if r.AcceptedAt != nil {
		AcceptedAt = *r.AcceptedAt
	}
	return
}

// The identifier of the Card Dispute that was accepted.
func (r WireTransferSimulationTransactionSourceCardDisputeAcceptance) GetCardDisputeID() (CardDisputeID string) {
	if r.CardDisputeID != nil {
		CardDisputeID = *r.CardDisputeID
	}
	return
}

// The identifier of the Transaction that was created to return the disputed funds
// to your account.
func (r WireTransferSimulationTransactionSourceCardDisputeAcceptance) GetTransactionID() (TransactionID string) {
	if r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardDisputeAcceptance) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCardDisputeAcceptance{AcceptedAt:%s CardDisputeID:%s TransactionID:%s}", core.FmtP(r.AcceptedAt), core.FmtP(r.CardDisputeID), core.FmtP(r.TransactionID))
}

type WireTransferSimulationTransactionSourceCardRefund struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *WireTransferSimulationTransactionSourceCardRefundCurrency `pjson:"currency"`
	// The identifier for the Transaction this refunds, if any.
	CardSettlementTransactionID *string `pjson:"card_settlement_transaction_id"`
	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type       *WireTransferSimulationTransactionSourceCardRefundType `pjson:"type"`
	jsonFields map[string]interface{}                                 `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCardRefund using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCardRefund) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceCardRefund into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *WireTransferSimulationTransactionSourceCardRefund) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r WireTransferSimulationTransactionSourceCardRefund) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r WireTransferSimulationTransactionSourceCardRefund) GetCurrency() (Currency WireTransferSimulationTransactionSourceCardRefundCurrency) {
	if r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier for the Transaction this refunds, if any.
func (r WireTransferSimulationTransactionSourceCardRefund) GetCardSettlementTransactionID() (CardSettlementTransactionID string) {
	if r.CardSettlementTransactionID != nil {
		CardSettlementTransactionID = *r.CardSettlementTransactionID
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_refund`.
func (r WireTransferSimulationTransactionSourceCardRefund) GetType() (Type WireTransferSimulationTransactionSourceCardRefundType) {
	if r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardRefund) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCardRefund{Amount:%s Currency:%s CardSettlementTransactionID:%s Type:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.CardSettlementTransactionID), core.FmtP(r.Type))
}

type WireTransferSimulationTransactionSourceCardRefundCurrency string

const (
	WireTransferSimulationTransactionSourceCardRefundCurrencyCad WireTransferSimulationTransactionSourceCardRefundCurrency = "CAD"
	WireTransferSimulationTransactionSourceCardRefundCurrencyChf WireTransferSimulationTransactionSourceCardRefundCurrency = "CHF"
	WireTransferSimulationTransactionSourceCardRefundCurrencyEur WireTransferSimulationTransactionSourceCardRefundCurrency = "EUR"
	WireTransferSimulationTransactionSourceCardRefundCurrencyGbp WireTransferSimulationTransactionSourceCardRefundCurrency = "GBP"
	WireTransferSimulationTransactionSourceCardRefundCurrencyJpy WireTransferSimulationTransactionSourceCardRefundCurrency = "JPY"
	WireTransferSimulationTransactionSourceCardRefundCurrencyUsd WireTransferSimulationTransactionSourceCardRefundCurrency = "USD"
)

type WireTransferSimulationTransactionSourceCardRefundType string

const (
	WireTransferSimulationTransactionSourceCardRefundTypeCardRefund WireTransferSimulationTransactionSourceCardRefundType = "card_refund"
)

type WireTransferSimulationTransactionSourceCardSettlement struct {
	// The amount in the minor unit of the transaction's settlement currency. For
	// dollars, for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's settlement currency.
	Currency *WireTransferSimulationTransactionSourceCardSettlementCurrency `pjson:"currency"`
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
	Type       *WireTransferSimulationTransactionSourceCardSettlementType `pjson:"type"`
	jsonFields map[string]interface{}                                     `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCardSettlement using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCardSettlement) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceCardSettlement
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *WireTransferSimulationTransactionSourceCardSettlement) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's settlement currency. For
// dollars, for example, this is cents.
func (r WireTransferSimulationTransactionSourceCardSettlement) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's settlement currency.
func (r WireTransferSimulationTransactionSourceCardSettlement) GetCurrency() (Currency WireTransferSimulationTransactionSourceCardSettlementCurrency) {
	if r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The amount in the minor unit of the transaction's presentment currency.
func (r WireTransferSimulationTransactionSourceCardSettlement) GetPresentmentAmount() (PresentmentAmount int64) {
	if r.PresentmentAmount != nil {
		PresentmentAmount = *r.PresentmentAmount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's presentment currency.
func (r WireTransferSimulationTransactionSourceCardSettlement) GetPresentmentCurrency() (PresentmentCurrency string) {
	if r.PresentmentCurrency != nil {
		PresentmentCurrency = *r.PresentmentCurrency
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardSettlement) GetMerchantCity() (MerchantCity string) {
	if r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardSettlement) GetMerchantCountry() (MerchantCountry string) {
	if r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardSettlement) GetMerchantName() (MerchantName string) {
	if r.MerchantName != nil {
		MerchantName = *r.MerchantName
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardSettlement) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardSettlement) GetMerchantState() (MerchantState string) {
	if r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

// The identifier of the Pending Transaction associated with this Transaction.
func (r WireTransferSimulationTransactionSourceCardSettlement) GetPendingTransactionID() (PendingTransactionID string) {
	if r.PendingTransactionID != nil {
		PendingTransactionID = *r.PendingTransactionID
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_settlement`.
func (r WireTransferSimulationTransactionSourceCardSettlement) GetType() (Type WireTransferSimulationTransactionSourceCardSettlementType) {
	if r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardSettlement) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCardSettlement{Amount:%s Currency:%s PresentmentAmount:%s PresentmentCurrency:%s MerchantCity:%s MerchantCountry:%s MerchantName:%s MerchantCategoryCode:%s MerchantState:%s PendingTransactionID:%s Type:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.PresentmentAmount), core.FmtP(r.PresentmentCurrency), core.FmtP(r.MerchantCity), core.FmtP(r.MerchantCountry), core.FmtP(r.MerchantName), core.FmtP(r.MerchantCategoryCode), core.FmtP(r.MerchantState), core.FmtP(r.PendingTransactionID), core.FmtP(r.Type))
}

type WireTransferSimulationTransactionSourceCardSettlementCurrency string

const (
	WireTransferSimulationTransactionSourceCardSettlementCurrencyCad WireTransferSimulationTransactionSourceCardSettlementCurrency = "CAD"
	WireTransferSimulationTransactionSourceCardSettlementCurrencyChf WireTransferSimulationTransactionSourceCardSettlementCurrency = "CHF"
	WireTransferSimulationTransactionSourceCardSettlementCurrencyEur WireTransferSimulationTransactionSourceCardSettlementCurrency = "EUR"
	WireTransferSimulationTransactionSourceCardSettlementCurrencyGbp WireTransferSimulationTransactionSourceCardSettlementCurrency = "GBP"
	WireTransferSimulationTransactionSourceCardSettlementCurrencyJpy WireTransferSimulationTransactionSourceCardSettlementCurrency = "JPY"
	WireTransferSimulationTransactionSourceCardSettlementCurrencyUsd WireTransferSimulationTransactionSourceCardSettlementCurrency = "USD"
)

type WireTransferSimulationTransactionSourceCardSettlementType string

const (
	WireTransferSimulationTransactionSourceCardSettlementTypeCardSettlement WireTransferSimulationTransactionSourceCardSettlementType = "card_settlement"
)

type WireTransferSimulationTransactionSourceCheckDepositAcceptance struct {
	// The amount to be deposited in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency `pjson:"currency"`
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
// WireTransferSimulationTransactionSourceCheckDepositAcceptance using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCheckDepositAcceptance) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceCheckDepositAcceptance into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceCheckDepositAcceptance) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount to be deposited in the minor unit of the transaction's currency. For
// dollars, for example, this is cents.
func (r WireTransferSimulationTransactionSourceCheckDepositAcceptance) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r WireTransferSimulationTransactionSourceCheckDepositAcceptance) GetCurrency() (Currency WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency) {
	if r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The account number printed on the check.
func (r WireTransferSimulationTransactionSourceCheckDepositAcceptance) GetAccountNumber() (AccountNumber string) {
	if r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The routing number printed on the check.
func (r WireTransferSimulationTransactionSourceCheckDepositAcceptance) GetRoutingNumber() (RoutingNumber string) {
	if r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// An additional line of metadata printed on the check. This typically includes the
// check number for business checks.
func (r WireTransferSimulationTransactionSourceCheckDepositAcceptance) GetAuxiliaryOnUs() (AuxiliaryOnUs string) {
	if r.AuxiliaryOnUs != nil {
		AuxiliaryOnUs = *r.AuxiliaryOnUs
	}
	return
}

// The check serial number, if present, for consumer checks. For business checks,
// the serial number is usually in the `auxiliary_on_us` field.
func (r WireTransferSimulationTransactionSourceCheckDepositAcceptance) GetSerialNumber() (SerialNumber string) {
	if r.SerialNumber != nil {
		SerialNumber = *r.SerialNumber
	}
	return
}

// The ID of the Check Deposit that was accepted.
func (r WireTransferSimulationTransactionSourceCheckDepositAcceptance) GetCheckDepositID() (CheckDepositID string) {
	if r.CheckDepositID != nil {
		CheckDepositID = *r.CheckDepositID
	}
	return
}

func (r WireTransferSimulationTransactionSourceCheckDepositAcceptance) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCheckDepositAcceptance{Amount:%s Currency:%s AccountNumber:%s RoutingNumber:%s AuxiliaryOnUs:%s SerialNumber:%s CheckDepositID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.AccountNumber), core.FmtP(r.RoutingNumber), core.FmtP(r.AuxiliaryOnUs), core.FmtP(r.SerialNumber), core.FmtP(r.CheckDepositID))
}

type WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency string

const (
	WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyCad WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "CAD"
	WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyChf WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "CHF"
	WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyEur WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "EUR"
	WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyGbp WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "GBP"
	WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyJpy WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "JPY"
	WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyUsd WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "USD"
)

type WireTransferSimulationTransactionSourceCheckDepositReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt *time.Time `pjson:"returned_at" format:"2006-01-02T15:04:05Z07:00"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *WireTransferSimulationTransactionSourceCheckDepositReturnCurrency `pjson:"currency"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID *string `pjson:"check_deposit_id"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID *string                                                                `pjson:"transaction_id"`
	ReturnReason  *WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason `pjson:"return_reason"`
	jsonFields    map[string]interface{}                                                 `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCheckDepositReturn using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCheckDepositReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceCheckDepositReturn
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *WireTransferSimulationTransactionSourceCheckDepositReturn) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r WireTransferSimulationTransactionSourceCheckDepositReturn) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the check deposit was returned.
func (r WireTransferSimulationTransactionSourceCheckDepositReturn) GetReturnedAt() (ReturnedAt time.Time) {
	if r.ReturnedAt != nil {
		ReturnedAt = *r.ReturnedAt
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r WireTransferSimulationTransactionSourceCheckDepositReturn) GetCurrency() (Currency WireTransferSimulationTransactionSourceCheckDepositReturnCurrency) {
	if r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Check Deposit that was returned.
func (r WireTransferSimulationTransactionSourceCheckDepositReturn) GetCheckDepositID() (CheckDepositID string) {
	if r.CheckDepositID != nil {
		CheckDepositID = *r.CheckDepositID
	}
	return
}

// The identifier of the transaction that reversed the original check deposit
// transaction.
func (r WireTransferSimulationTransactionSourceCheckDepositReturn) GetTransactionID() (TransactionID string) {
	if r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

func (r WireTransferSimulationTransactionSourceCheckDepositReturn) GetReturnReason() (ReturnReason WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason) {
	if r.ReturnReason != nil {
		ReturnReason = *r.ReturnReason
	}
	return
}

func (r WireTransferSimulationTransactionSourceCheckDepositReturn) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCheckDepositReturn{Amount:%s ReturnedAt:%s Currency:%s CheckDepositID:%s TransactionID:%s ReturnReason:%s}", core.FmtP(r.Amount), core.FmtP(r.ReturnedAt), core.FmtP(r.Currency), core.FmtP(r.CheckDepositID), core.FmtP(r.TransactionID), core.FmtP(r.ReturnReason))
}

type WireTransferSimulationTransactionSourceCheckDepositReturnCurrency string

const (
	WireTransferSimulationTransactionSourceCheckDepositReturnCurrencyCad WireTransferSimulationTransactionSourceCheckDepositReturnCurrency = "CAD"
	WireTransferSimulationTransactionSourceCheckDepositReturnCurrencyChf WireTransferSimulationTransactionSourceCheckDepositReturnCurrency = "CHF"
	WireTransferSimulationTransactionSourceCheckDepositReturnCurrencyEur WireTransferSimulationTransactionSourceCheckDepositReturnCurrency = "EUR"
	WireTransferSimulationTransactionSourceCheckDepositReturnCurrencyGbp WireTransferSimulationTransactionSourceCheckDepositReturnCurrency = "GBP"
	WireTransferSimulationTransactionSourceCheckDepositReturnCurrencyJpy WireTransferSimulationTransactionSourceCheckDepositReturnCurrency = "JPY"
	WireTransferSimulationTransactionSourceCheckDepositReturnCurrencyUsd WireTransferSimulationTransactionSourceCheckDepositReturnCurrency = "USD"
)

type WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason string

const (
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonACHConversionNotSupported WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "ach_conversion_not_supported"
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonClosedAccount             WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "closed_account"
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonDuplicateSubmission       WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "duplicate_submission"
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonInsufficientFunds         WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "insufficient_funds"
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonNoAccount                 WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "no_account"
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonNotAuthorized             WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "not_authorized"
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonStaleDated                WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "stale_dated"
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonStopPayment               WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "stop_payment"
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonUnknownReason             WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "unknown_reason"
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonUnmatchedDetails          WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "unmatched_details"
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonUnreadableImage           WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "unreadable_image"
)

type WireTransferSimulationTransactionSourceCheckTransferIntention struct {
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
	Currency *WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency `pjson:"currency"`
	// The name that will be printed on the check.
	RecipientName *string `pjson:"recipient_name"`
	// The identifier of the Check Transfer with which this is associated.
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCheckTransferIntention using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCheckTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceCheckTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceCheckTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The street address of the check's destination.
func (r WireTransferSimulationTransactionSourceCheckTransferIntention) GetAddressLine1() (AddressLine1 string) {
	if r.AddressLine1 != nil {
		AddressLine1 = *r.AddressLine1
	}
	return
}

// The second line of the address of the check's destination.
func (r WireTransferSimulationTransactionSourceCheckTransferIntention) GetAddressLine2() (AddressLine2 string) {
	if r.AddressLine2 != nil {
		AddressLine2 = *r.AddressLine2
	}
	return
}

// The city of the check's destination.
func (r WireTransferSimulationTransactionSourceCheckTransferIntention) GetAddressCity() (AddressCity string) {
	if r.AddressCity != nil {
		AddressCity = *r.AddressCity
	}
	return
}

// The state of the check's destination.
func (r WireTransferSimulationTransactionSourceCheckTransferIntention) GetAddressState() (AddressState string) {
	if r.AddressState != nil {
		AddressState = *r.AddressState
	}
	return
}

// The postal code of the check's destination.
func (r WireTransferSimulationTransactionSourceCheckTransferIntention) GetAddressZip() (AddressZip string) {
	if r.AddressZip != nil {
		AddressZip = *r.AddressZip
	}
	return
}

// The transfer amount in USD cents.
func (r WireTransferSimulationTransactionSourceCheckTransferIntention) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
func (r WireTransferSimulationTransactionSourceCheckTransferIntention) GetCurrency() (Currency WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency) {
	if r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The name that will be printed on the check.
func (r WireTransferSimulationTransactionSourceCheckTransferIntention) GetRecipientName() (RecipientName string) {
	if r.RecipientName != nil {
		RecipientName = *r.RecipientName
	}
	return
}

// The identifier of the Check Transfer with which this is associated.
func (r WireTransferSimulationTransactionSourceCheckTransferIntention) GetTransferID() (TransferID string) {
	if r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r WireTransferSimulationTransactionSourceCheckTransferIntention) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCheckTransferIntention{AddressLine1:%s AddressLine2:%s AddressCity:%s AddressState:%s AddressZip:%s Amount:%s Currency:%s RecipientName:%s TransferID:%s}", core.FmtP(r.AddressLine1), core.FmtP(r.AddressLine2), core.FmtP(r.AddressCity), core.FmtP(r.AddressState), core.FmtP(r.AddressZip), core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.RecipientName), core.FmtP(r.TransferID))
}

type WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency string

const (
	WireTransferSimulationTransactionSourceCheckTransferIntentionCurrencyCad WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "CAD"
	WireTransferSimulationTransactionSourceCheckTransferIntentionCurrencyChf WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "CHF"
	WireTransferSimulationTransactionSourceCheckTransferIntentionCurrencyEur WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "EUR"
	WireTransferSimulationTransactionSourceCheckTransferIntentionCurrencyGbp WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "GBP"
	WireTransferSimulationTransactionSourceCheckTransferIntentionCurrencyJpy WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "JPY"
	WireTransferSimulationTransactionSourceCheckTransferIntentionCurrencyUsd WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "USD"
)

type WireTransferSimulationTransactionSourceCheckTransferReturn struct {
	// The identifier of the returned Check Transfer.
	TransferID *string `pjson:"transfer_id"`
	// If available, a document with additional information about the return.
	FileID *string `pjson:"file_id"`
	// The reason why the check was returned.
	Reason     *WireTransferSimulationTransactionSourceCheckTransferReturnReason `pjson:"reason"`
	jsonFields map[string]interface{}                                            `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCheckTransferReturn using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCheckTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceCheckTransferReturn into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceCheckTransferReturn) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the returned Check Transfer.
func (r WireTransferSimulationTransactionSourceCheckTransferReturn) GetTransferID() (TransferID string) {
	if r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

// If available, a document with additional information about the return.
func (r WireTransferSimulationTransactionSourceCheckTransferReturn) GetFileID() (FileID string) {
	if r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

// The reason why the check was returned.
func (r WireTransferSimulationTransactionSourceCheckTransferReturn) GetReason() (Reason WireTransferSimulationTransactionSourceCheckTransferReturnReason) {
	if r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

func (r WireTransferSimulationTransactionSourceCheckTransferReturn) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCheckTransferReturn{TransferID:%s FileID:%s Reason:%s}", core.FmtP(r.TransferID), core.FmtP(r.FileID), core.FmtP(r.Reason))
}

type WireTransferSimulationTransactionSourceCheckTransferReturnReason string

const (
	WireTransferSimulationTransactionSourceCheckTransferReturnReasonMailDeliveryFailure WireTransferSimulationTransactionSourceCheckTransferReturnReason = "mail_delivery_failure"
	WireTransferSimulationTransactionSourceCheckTransferReturnReasonRefusedByRecipient  WireTransferSimulationTransactionSourceCheckTransferReturnReason = "refused_by_recipient"
)

type WireTransferSimulationTransactionSourceCheckTransferRejection struct {
	// The identifier of the Check Transfer that led to this Transaction.
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCheckTransferRejection using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCheckTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceCheckTransferRejection into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceCheckTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the Check Transfer that led to this Transaction.
func (r WireTransferSimulationTransactionSourceCheckTransferRejection) GetTransferID() (TransferID string) {
	if r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r WireTransferSimulationTransactionSourceCheckTransferRejection) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCheckTransferRejection{TransferID:%s}", core.FmtP(r.TransferID))
}

type WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferID *string `pjson:"transfer_id"`
	// The transaction ID of the corresponding credit transaction.
	TransactionID *string `pjson:"transaction_id"`
	// The time the stop-payment was requested.
	RequestedAt *time.Time `pjson:"requested_at" format:"2006-01-02T15:04:05Z07:00"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type       *WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType `pjson:"type"`
	jsonFields map[string]interface{}                                                      `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The ID of the check transfer that was stopped.
func (r WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) GetTransferID() (TransferID string) {
	if r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

// The transaction ID of the corresponding credit transaction.
func (r WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) GetTransactionID() (TransactionID string) {
	if r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

// The time the stop-payment was requested.
func (r WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) GetRequestedAt() (RequestedAt time.Time) {
	if r.RequestedAt != nil {
		RequestedAt = *r.RequestedAt
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_stop_payment_request`.
func (r WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) GetType() (Type WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType) {
	if r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest{TransferID:%s TransactionID:%s RequestedAt:%s Type:%s}", core.FmtP(r.TransferID), core.FmtP(r.TransactionID), core.FmtP(r.RequestedAt), core.FmtP(r.Type))
}

type WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType string

const (
	WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

type WireTransferSimulationTransactionSourceDisputeResolution struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *WireTransferSimulationTransactionSourceDisputeResolutionCurrency `pjson:"currency"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionID *string                `pjson:"disputed_transaction_id"`
	jsonFields            map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceDisputeResolution using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceDisputeResolution) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceDisputeResolution
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *WireTransferSimulationTransactionSourceDisputeResolution) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r WireTransferSimulationTransactionSourceDisputeResolution) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r WireTransferSimulationTransactionSourceDisputeResolution) GetCurrency() (Currency WireTransferSimulationTransactionSourceDisputeResolutionCurrency) {
	if r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Transaction that was disputed.
func (r WireTransferSimulationTransactionSourceDisputeResolution) GetDisputedTransactionID() (DisputedTransactionID string) {
	if r.DisputedTransactionID != nil {
		DisputedTransactionID = *r.DisputedTransactionID
	}
	return
}

func (r WireTransferSimulationTransactionSourceDisputeResolution) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceDisputeResolution{Amount:%s Currency:%s DisputedTransactionID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.DisputedTransactionID))
}

type WireTransferSimulationTransactionSourceDisputeResolutionCurrency string

const (
	WireTransferSimulationTransactionSourceDisputeResolutionCurrencyCad WireTransferSimulationTransactionSourceDisputeResolutionCurrency = "CAD"
	WireTransferSimulationTransactionSourceDisputeResolutionCurrencyChf WireTransferSimulationTransactionSourceDisputeResolutionCurrency = "CHF"
	WireTransferSimulationTransactionSourceDisputeResolutionCurrencyEur WireTransferSimulationTransactionSourceDisputeResolutionCurrency = "EUR"
	WireTransferSimulationTransactionSourceDisputeResolutionCurrencyGbp WireTransferSimulationTransactionSourceDisputeResolutionCurrency = "GBP"
	WireTransferSimulationTransactionSourceDisputeResolutionCurrencyJpy WireTransferSimulationTransactionSourceDisputeResolutionCurrency = "JPY"
	WireTransferSimulationTransactionSourceDisputeResolutionCurrencyUsd WireTransferSimulationTransactionSourceDisputeResolutionCurrency = "USD"
)

type WireTransferSimulationTransactionSourceEmpyrealCashDeposit struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount      *int64                 `pjson:"amount"`
	BagID       *string                `pjson:"bag_id"`
	DepositDate *time.Time             `pjson:"deposit_date" format:"2006-01-02T15:04:05Z07:00"`
	jsonFields  map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceEmpyrealCashDeposit using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceEmpyrealCashDeposit) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceEmpyrealCashDeposit into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceEmpyrealCashDeposit) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r WireTransferSimulationTransactionSourceEmpyrealCashDeposit) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r WireTransferSimulationTransactionSourceEmpyrealCashDeposit) GetBagID() (BagID string) {
	if r.BagID != nil {
		BagID = *r.BagID
	}
	return
}

func (r WireTransferSimulationTransactionSourceEmpyrealCashDeposit) GetDepositDate() (DepositDate time.Time) {
	if r.DepositDate != nil {
		DepositDate = *r.DepositDate
	}
	return
}

func (r WireTransferSimulationTransactionSourceEmpyrealCashDeposit) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceEmpyrealCashDeposit{Amount:%s BagID:%s DepositDate:%s}", core.FmtP(r.Amount), core.FmtP(r.BagID), core.FmtP(r.DepositDate))
}

type WireTransferSimulationTransactionSourceInboundACHTransfer struct {
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
// WireTransferSimulationTransactionSourceInboundACHTransfer using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInboundACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceInboundACHTransfer
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *WireTransferSimulationTransactionSourceInboundACHTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the destination account currency. For dollars,
// for example, this is cents.
func (r WireTransferSimulationTransactionSourceInboundACHTransfer) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundACHTransfer) GetOriginatorCompanyName() (OriginatorCompanyName string) {
	if r.OriginatorCompanyName != nil {
		OriginatorCompanyName = *r.OriginatorCompanyName
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundACHTransfer) GetOriginatorCompanyDescriptiveDate() (OriginatorCompanyDescriptiveDate string) {
	if r.OriginatorCompanyDescriptiveDate != nil {
		OriginatorCompanyDescriptiveDate = *r.OriginatorCompanyDescriptiveDate
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundACHTransfer) GetOriginatorCompanyDiscretionaryData() (OriginatorCompanyDiscretionaryData string) {
	if r.OriginatorCompanyDiscretionaryData != nil {
		OriginatorCompanyDiscretionaryData = *r.OriginatorCompanyDiscretionaryData
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundACHTransfer) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundACHTransfer) GetOriginatorCompanyID() (OriginatorCompanyID string) {
	if r.OriginatorCompanyID != nil {
		OriginatorCompanyID = *r.OriginatorCompanyID
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundACHTransfer) GetReceiverIDNumber() (ReceiverIDNumber string) {
	if r.ReceiverIDNumber != nil {
		ReceiverIDNumber = *r.ReceiverIDNumber
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundACHTransfer) GetReceiverName() (ReceiverName string) {
	if r.ReceiverName != nil {
		ReceiverName = *r.ReceiverName
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundACHTransfer) GetTraceNumber() (TraceNumber string) {
	if r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundACHTransfer) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceInboundACHTransfer{Amount:%s OriginatorCompanyName:%s OriginatorCompanyDescriptiveDate:%s OriginatorCompanyDiscretionaryData:%s OriginatorCompanyEntryDescription:%s OriginatorCompanyID:%s ReceiverIDNumber:%s ReceiverName:%s TraceNumber:%s}", core.FmtP(r.Amount), core.FmtP(r.OriginatorCompanyName), core.FmtP(r.OriginatorCompanyDescriptiveDate), core.FmtP(r.OriginatorCompanyDiscretionaryData), core.FmtP(r.OriginatorCompanyEntryDescription), core.FmtP(r.OriginatorCompanyID), core.FmtP(r.ReceiverIDNumber), core.FmtP(r.ReceiverName), core.FmtP(r.TraceNumber))
}

type WireTransferSimulationTransactionSourceInboundCheck struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency              *WireTransferSimulationTransactionSourceInboundCheckCurrency `pjson:"currency"`
	CheckNumber           *string                                                      `pjson:"check_number"`
	CheckFrontImageFileID *string                                                      `pjson:"check_front_image_file_id"`
	CheckRearImageFileID  *string                                                      `pjson:"check_rear_image_file_id"`
	jsonFields            map[string]interface{}                                       `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceInboundCheck using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInboundCheck) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceInboundCheck into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *WireTransferSimulationTransactionSourceInboundCheck) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the destination account currency. For dollars,
// for example, this is cents.
func (r WireTransferSimulationTransactionSourceInboundCheck) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r WireTransferSimulationTransactionSourceInboundCheck) GetCurrency() (Currency WireTransferSimulationTransactionSourceInboundCheckCurrency) {
	if r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundCheck) GetCheckNumber() (CheckNumber string) {
	if r.CheckNumber != nil {
		CheckNumber = *r.CheckNumber
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundCheck) GetCheckFrontImageFileID() (CheckFrontImageFileID string) {
	if r.CheckFrontImageFileID != nil {
		CheckFrontImageFileID = *r.CheckFrontImageFileID
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundCheck) GetCheckRearImageFileID() (CheckRearImageFileID string) {
	if r.CheckRearImageFileID != nil {
		CheckRearImageFileID = *r.CheckRearImageFileID
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundCheck) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceInboundCheck{Amount:%s Currency:%s CheckNumber:%s CheckFrontImageFileID:%s CheckRearImageFileID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.CheckNumber), core.FmtP(r.CheckFrontImageFileID), core.FmtP(r.CheckRearImageFileID))
}

type WireTransferSimulationTransactionSourceInboundCheckCurrency string

const (
	WireTransferSimulationTransactionSourceInboundCheckCurrencyCad WireTransferSimulationTransactionSourceInboundCheckCurrency = "CAD"
	WireTransferSimulationTransactionSourceInboundCheckCurrencyChf WireTransferSimulationTransactionSourceInboundCheckCurrency = "CHF"
	WireTransferSimulationTransactionSourceInboundCheckCurrencyEur WireTransferSimulationTransactionSourceInboundCheckCurrency = "EUR"
	WireTransferSimulationTransactionSourceInboundCheckCurrencyGbp WireTransferSimulationTransactionSourceInboundCheckCurrency = "GBP"
	WireTransferSimulationTransactionSourceInboundCheckCurrencyJpy WireTransferSimulationTransactionSourceInboundCheckCurrency = "JPY"
	WireTransferSimulationTransactionSourceInboundCheckCurrencyUsd WireTransferSimulationTransactionSourceInboundCheckCurrency = "USD"
)

type WireTransferSimulationTransactionSourceInboundInternationalACHTransfer struct {
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
// WireTransferSimulationTransactionSourceInboundInternationalACHTransfer using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceInboundInternationalACHTransfer into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the destination account currency. For dollars,
// for example, this is cents.
func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetForeignExchangeIndicator() (ForeignExchangeIndicator string) {
	if r.ForeignExchangeIndicator != nil {
		ForeignExchangeIndicator = *r.ForeignExchangeIndicator
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetForeignExchangeReferenceIndicator() (ForeignExchangeReferenceIndicator string) {
	if r.ForeignExchangeReferenceIndicator != nil {
		ForeignExchangeReferenceIndicator = *r.ForeignExchangeReferenceIndicator
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetForeignExchangeReference() (ForeignExchangeReference string) {
	if r.ForeignExchangeReference != nil {
		ForeignExchangeReference = *r.ForeignExchangeReference
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetDestinationCountryCode() (DestinationCountryCode string) {
	if r.DestinationCountryCode != nil {
		DestinationCountryCode = *r.DestinationCountryCode
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetDestinationCurrencyCode() (DestinationCurrencyCode string) {
	if r.DestinationCurrencyCode != nil {
		DestinationCurrencyCode = *r.DestinationCurrencyCode
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetForeignPaymentAmount() (ForeignPaymentAmount int64) {
	if r.ForeignPaymentAmount != nil {
		ForeignPaymentAmount = *r.ForeignPaymentAmount
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetForeignTraceNumber() (ForeignTraceNumber string) {
	if r.ForeignTraceNumber != nil {
		ForeignTraceNumber = *r.ForeignTraceNumber
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetInternationalTransactionTypeCode() (InternationalTransactionTypeCode string) {
	if r.InternationalTransactionTypeCode != nil {
		InternationalTransactionTypeCode = *r.InternationalTransactionTypeCode
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatingCurrencyCode() (OriginatingCurrencyCode string) {
	if r.OriginatingCurrencyCode != nil {
		OriginatingCurrencyCode = *r.OriginatingCurrencyCode
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionName() (OriginatingDepositoryFinancialInstitutionName string) {
	if r.OriginatingDepositoryFinancialInstitutionName != nil {
		OriginatingDepositoryFinancialInstitutionName = *r.OriginatingDepositoryFinancialInstitutionName
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionIDQualifier() (OriginatingDepositoryFinancialInstitutionIDQualifier string) {
	if r.OriginatingDepositoryFinancialInstitutionIDQualifier != nil {
		OriginatingDepositoryFinancialInstitutionIDQualifier = *r.OriginatingDepositoryFinancialInstitutionIDQualifier
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionID() (OriginatingDepositoryFinancialInstitutionID string) {
	if r.OriginatingDepositoryFinancialInstitutionID != nil {
		OriginatingDepositoryFinancialInstitutionID = *r.OriginatingDepositoryFinancialInstitutionID
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionBranchCountry() (OriginatingDepositoryFinancialInstitutionBranchCountry string) {
	if r.OriginatingDepositoryFinancialInstitutionBranchCountry != nil {
		OriginatingDepositoryFinancialInstitutionBranchCountry = *r.OriginatingDepositoryFinancialInstitutionBranchCountry
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorCity() (OriginatorCity string) {
	if r.OriginatorCity != nil {
		OriginatorCity = *r.OriginatorCity
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorCountry() (OriginatorCountry string) {
	if r.OriginatorCountry != nil {
		OriginatorCountry = *r.OriginatorCountry
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorIdentification() (OriginatorIdentification string) {
	if r.OriginatorIdentification != nil {
		OriginatorIdentification = *r.OriginatorIdentification
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorName() (OriginatorName string) {
	if r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorPostalCode() (OriginatorPostalCode string) {
	if r.OriginatorPostalCode != nil {
		OriginatorPostalCode = *r.OriginatorPostalCode
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorStreetAddress() (OriginatorStreetAddress string) {
	if r.OriginatorStreetAddress != nil {
		OriginatorStreetAddress = *r.OriginatorStreetAddress
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorStateOrProvince() (OriginatorStateOrProvince string) {
	if r.OriginatorStateOrProvince != nil {
		OriginatorStateOrProvince = *r.OriginatorStateOrProvince
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetPaymentRelatedInformation() (PaymentRelatedInformation string) {
	if r.PaymentRelatedInformation != nil {
		PaymentRelatedInformation = *r.PaymentRelatedInformation
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetPaymentRelatedInformation2() (PaymentRelatedInformation2 string) {
	if r.PaymentRelatedInformation2 != nil {
		PaymentRelatedInformation2 = *r.PaymentRelatedInformation2
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverIdentificationNumber() (ReceiverIdentificationNumber string) {
	if r.ReceiverIdentificationNumber != nil {
		ReceiverIdentificationNumber = *r.ReceiverIdentificationNumber
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverStreetAddress() (ReceiverStreetAddress string) {
	if r.ReceiverStreetAddress != nil {
		ReceiverStreetAddress = *r.ReceiverStreetAddress
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverCity() (ReceiverCity string) {
	if r.ReceiverCity != nil {
		ReceiverCity = *r.ReceiverCity
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverStateOrProvince() (ReceiverStateOrProvince string) {
	if r.ReceiverStateOrProvince != nil {
		ReceiverStateOrProvince = *r.ReceiverStateOrProvince
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverCountry() (ReceiverCountry string) {
	if r.ReceiverCountry != nil {
		ReceiverCountry = *r.ReceiverCountry
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverPostalCode() (ReceiverPostalCode string) {
	if r.ReceiverPostalCode != nil {
		ReceiverPostalCode = *r.ReceiverPostalCode
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceivingCompanyOrIndividualName() (ReceivingCompanyOrIndividualName string) {
	if r.ReceivingCompanyOrIndividualName != nil {
		ReceivingCompanyOrIndividualName = *r.ReceivingCompanyOrIndividualName
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionName() (ReceivingDepositoryFinancialInstitutionName string) {
	if r.ReceivingDepositoryFinancialInstitutionName != nil {
		ReceivingDepositoryFinancialInstitutionName = *r.ReceivingDepositoryFinancialInstitutionName
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionIDQualifier() (ReceivingDepositoryFinancialInstitutionIDQualifier string) {
	if r.ReceivingDepositoryFinancialInstitutionIDQualifier != nil {
		ReceivingDepositoryFinancialInstitutionIDQualifier = *r.ReceivingDepositoryFinancialInstitutionIDQualifier
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionID() (ReceivingDepositoryFinancialInstitutionID string) {
	if r.ReceivingDepositoryFinancialInstitutionID != nil {
		ReceivingDepositoryFinancialInstitutionID = *r.ReceivingDepositoryFinancialInstitutionID
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionCountry() (ReceivingDepositoryFinancialInstitutionCountry string) {
	if r.ReceivingDepositoryFinancialInstitutionCountry != nil {
		ReceivingDepositoryFinancialInstitutionCountry = *r.ReceivingDepositoryFinancialInstitutionCountry
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetTraceNumber() (TraceNumber string) {
	if r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceInboundInternationalACHTransfer{Amount:%s ForeignExchangeIndicator:%s ForeignExchangeReferenceIndicator:%s ForeignExchangeReference:%s DestinationCountryCode:%s DestinationCurrencyCode:%s ForeignPaymentAmount:%s ForeignTraceNumber:%s InternationalTransactionTypeCode:%s OriginatingCurrencyCode:%s OriginatingDepositoryFinancialInstitutionName:%s OriginatingDepositoryFinancialInstitutionIDQualifier:%s OriginatingDepositoryFinancialInstitutionID:%s OriginatingDepositoryFinancialInstitutionBranchCountry:%s OriginatorCity:%s OriginatorCompanyEntryDescription:%s OriginatorCountry:%s OriginatorIdentification:%s OriginatorName:%s OriginatorPostalCode:%s OriginatorStreetAddress:%s OriginatorStateOrProvince:%s PaymentRelatedInformation:%s PaymentRelatedInformation2:%s ReceiverIdentificationNumber:%s ReceiverStreetAddress:%s ReceiverCity:%s ReceiverStateOrProvince:%s ReceiverCountry:%s ReceiverPostalCode:%s ReceivingCompanyOrIndividualName:%s ReceivingDepositoryFinancialInstitutionName:%s ReceivingDepositoryFinancialInstitutionIDQualifier:%s ReceivingDepositoryFinancialInstitutionID:%s ReceivingDepositoryFinancialInstitutionCountry:%s TraceNumber:%s}", core.FmtP(r.Amount), core.FmtP(r.ForeignExchangeIndicator), core.FmtP(r.ForeignExchangeReferenceIndicator), core.FmtP(r.ForeignExchangeReference), core.FmtP(r.DestinationCountryCode), core.FmtP(r.DestinationCurrencyCode), core.FmtP(r.ForeignPaymentAmount), core.FmtP(r.ForeignTraceNumber), core.FmtP(r.InternationalTransactionTypeCode), core.FmtP(r.OriginatingCurrencyCode), core.FmtP(r.OriginatingDepositoryFinancialInstitutionName), core.FmtP(r.OriginatingDepositoryFinancialInstitutionIDQualifier), core.FmtP(r.OriginatingDepositoryFinancialInstitutionID), core.FmtP(r.OriginatingDepositoryFinancialInstitutionBranchCountry), core.FmtP(r.OriginatorCity), core.FmtP(r.OriginatorCompanyEntryDescription), core.FmtP(r.OriginatorCountry), core.FmtP(r.OriginatorIdentification), core.FmtP(r.OriginatorName), core.FmtP(r.OriginatorPostalCode), core.FmtP(r.OriginatorStreetAddress), core.FmtP(r.OriginatorStateOrProvince), core.FmtP(r.PaymentRelatedInformation), core.FmtP(r.PaymentRelatedInformation2), core.FmtP(r.ReceiverIdentificationNumber), core.FmtP(r.ReceiverStreetAddress), core.FmtP(r.ReceiverCity), core.FmtP(r.ReceiverStateOrProvince), core.FmtP(r.ReceiverCountry), core.FmtP(r.ReceiverPostalCode), core.FmtP(r.ReceivingCompanyOrIndividualName), core.FmtP(r.ReceivingDepositoryFinancialInstitutionName), core.FmtP(r.ReceivingDepositoryFinancialInstitutionIDQualifier), core.FmtP(r.ReceivingDepositoryFinancialInstitutionID), core.FmtP(r.ReceivingDepositoryFinancialInstitutionCountry), core.FmtP(r.TraceNumber))
}

type WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real Time Payments transfer.
	Currency *WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency `pjson:"currency"`
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
// WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transfer's currency. For dollars, for
// example, this is cents.
func (r WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
// currency. This will always be "USD" for a Real Time Payments transfer.
func (r WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetCurrency() (Currency WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency) {
	if r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The name the sender of the transfer specified as the recipient of the transfer.
func (r WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetCreditorName() (CreditorName string) {
	if r.CreditorName != nil {
		CreditorName = *r.CreditorName
	}
	return
}

// The name provided by the sender of the transfer.
func (r WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorName() (DebtorName string) {
	if r.DebtorName != nil {
		DebtorName = *r.DebtorName
	}
	return
}

// The account number of the account that sent the transfer.
func (r WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorAccountNumber() (DebtorAccountNumber string) {
	if r.DebtorAccountNumber != nil {
		DebtorAccountNumber = *r.DebtorAccountNumber
	}
	return
}

// The routing number of the account that sent the transfer.
func (r WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorRoutingNumber() (DebtorRoutingNumber string) {
	if r.DebtorRoutingNumber != nil {
		DebtorRoutingNumber = *r.DebtorRoutingNumber
	}
	return
}

// The Real Time Payments network identification of the transfer
func (r WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetTransactionIdentification() (TransactionIdentification string) {
	if r.TransactionIdentification != nil {
		TransactionIdentification = *r.TransactionIdentification
	}
	return
}

// Additional information included with the transfer.
func (r WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetRemittanceInformation() (RemittanceInformation string) {
	if r.RemittanceInformation != nil {
		RemittanceInformation = *r.RemittanceInformation
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation{Amount:%s Currency:%s CreditorName:%s DebtorName:%s DebtorAccountNumber:%s DebtorRoutingNumber:%s TransactionIdentification:%s RemittanceInformation:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.CreditorName), core.FmtP(r.DebtorName), core.FmtP(r.DebtorAccountNumber), core.FmtP(r.DebtorRoutingNumber), core.FmtP(r.TransactionIdentification), core.FmtP(r.RemittanceInformation))
}

type WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency string

const (
	WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyCad WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CAD"
	WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyChf WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CHF"
	WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyEur WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "EUR"
	WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyGbp WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "GBP"
	WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyJpy WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "JPY"
	WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyUsd WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "USD"
)

type WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal struct {
	// The amount that was reversed.
	Amount *int64 `pjson:"amount"`
	// The description on the reversal message from Fedwire.
	Description *string `pjson:"description"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate *time.Time `pjson:"input_cycle_date" format:"2006-01-02"`
	// The Fedwire sequence number.
	InputSequenceNumber *string `pjson:"input_sequence_number"`
	// The Fedwire input source identifier.
	InputSource *string `pjson:"input_source"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData *string `pjson:"input_message_accountability_data"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData *string `pjson:"previous_message_input_message_accountability_data"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate *time.Time `pjson:"previous_message_input_cycle_date" format:"2006-01-02"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber *string `pjson:"previous_message_input_sequence_number"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource *string                `pjson:"previous_message_input_source"`
	jsonFields                 map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount that was reversed.
func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description on the reversal message from Fedwire.
func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetDescription() (Description string) {
	if r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Fedwire cycle date for the wire reversal.
func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetInputCycleDate() (InputCycleDate time.Time) {
	if r.InputCycleDate != nil {
		InputCycleDate = *r.InputCycleDate
	}
	return
}

// The Fedwire sequence number.
func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetInputSequenceNumber() (InputSequenceNumber string) {
	if r.InputSequenceNumber != nil {
		InputSequenceNumber = *r.InputSequenceNumber
	}
	return
}

// The Fedwire input source identifier.
func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetInputSource() (InputSource string) {
	if r.InputSource != nil {
		InputSource = *r.InputSource
	}
	return
}

// The Fedwire transaction identifier.
func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

// The Fedwire transaction identifier for the wire transfer that was reversed.
func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputMessageAccountabilityData() (PreviousMessageInputMessageAccountabilityData string) {
	if r.PreviousMessageInputMessageAccountabilityData != nil {
		PreviousMessageInputMessageAccountabilityData = *r.PreviousMessageInputMessageAccountabilityData
	}
	return
}

// The Fedwire cycle date for the wire transfer that was reversed.
func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputCycleDate() (PreviousMessageInputCycleDate time.Time) {
	if r.PreviousMessageInputCycleDate != nil {
		PreviousMessageInputCycleDate = *r.PreviousMessageInputCycleDate
	}
	return
}

// The Fedwire sequence number for the wire transfer that was reversed.
func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputSequenceNumber() (PreviousMessageInputSequenceNumber string) {
	if r.PreviousMessageInputSequenceNumber != nil {
		PreviousMessageInputSequenceNumber = *r.PreviousMessageInputSequenceNumber
	}
	return
}

// The Fedwire input source identifier for the wire transfer that was reversed.
func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputSource() (PreviousMessageInputSource string) {
	if r.PreviousMessageInputSource != nil {
		PreviousMessageInputSource = *r.PreviousMessageInputSource
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal{Amount:%s Description:%s InputCycleDate:%s InputSequenceNumber:%s InputSource:%s InputMessageAccountabilityData:%s PreviousMessageInputMessageAccountabilityData:%s PreviousMessageInputCycleDate:%s PreviousMessageInputSequenceNumber:%s PreviousMessageInputSource:%s}", core.FmtP(r.Amount), core.FmtP(r.Description), core.FmtP(r.InputCycleDate), core.FmtP(r.InputSequenceNumber), core.FmtP(r.InputSource), core.FmtP(r.InputMessageAccountabilityData), core.FmtP(r.PreviousMessageInputMessageAccountabilityData), core.FmtP(r.PreviousMessageInputCycleDate), core.FmtP(r.PreviousMessageInputSequenceNumber), core.FmtP(r.PreviousMessageInputSource))
}

type WireTransferSimulationTransactionSourceInboundWireDrawdownPayment struct {
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
// WireTransferSimulationTransactionSourceInboundWireDrawdownPayment using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceInboundWireDrawdownPayment into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryName() (BeneficiaryName string) {
	if r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryReference() (BeneficiaryReference string) {
	if r.BeneficiaryReference != nil {
		BeneficiaryReference = *r.BeneficiaryReference
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetDescription() (Description string) {
	if r.Description != nil {
		Description = *r.Description
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine1() (OriginatorAddressLine1 string) {
	if r.OriginatorAddressLine1 != nil {
		OriginatorAddressLine1 = *r.OriginatorAddressLine1
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine2() (OriginatorAddressLine2 string) {
	if r.OriginatorAddressLine2 != nil {
		OriginatorAddressLine2 = *r.OriginatorAddressLine2
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine3() (OriginatorAddressLine3 string) {
	if r.OriginatorAddressLine3 != nil {
		OriginatorAddressLine3 = *r.OriginatorAddressLine3
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetOriginatorName() (OriginatorName string) {
	if r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetOriginatorToBeneficiaryInformation() (OriginatorToBeneficiaryInformation string) {
	if r.OriginatorToBeneficiaryInformation != nil {
		OriginatorToBeneficiaryInformation = *r.OriginatorToBeneficiaryInformation
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceInboundWireDrawdownPayment{Amount:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s BeneficiaryName:%s BeneficiaryReference:%s Description:%s InputMessageAccountabilityData:%s OriginatorAddressLine1:%s OriginatorAddressLine2:%s OriginatorAddressLine3:%s OriginatorName:%s OriginatorToBeneficiaryInformation:%s}", core.FmtP(r.Amount), core.FmtP(r.BeneficiaryAddressLine1), core.FmtP(r.BeneficiaryAddressLine2), core.FmtP(r.BeneficiaryAddressLine3), core.FmtP(r.BeneficiaryName), core.FmtP(r.BeneficiaryReference), core.FmtP(r.Description), core.FmtP(r.InputMessageAccountabilityData), core.FmtP(r.OriginatorAddressLine1), core.FmtP(r.OriginatorAddressLine2), core.FmtP(r.OriginatorAddressLine3), core.FmtP(r.OriginatorName), core.FmtP(r.OriginatorToBeneficiaryInformation))
}

type WireTransferSimulationTransactionSourceInboundWireReversal struct {
	// The amount that was reversed.
	Amount *int64 `pjson:"amount"`
	// The description on the reversal message from Fedwire.
	Description *string `pjson:"description"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate *time.Time `pjson:"input_cycle_date" format:"2006-01-02"`
	// The Fedwire sequence number.
	InputSequenceNumber *string `pjson:"input_sequence_number"`
	// The Fedwire input source identifier.
	InputSource *string `pjson:"input_source"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData *string `pjson:"input_message_accountability_data"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData *string `pjson:"previous_message_input_message_accountability_data"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate *time.Time `pjson:"previous_message_input_cycle_date" format:"2006-01-02"`
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
// WireTransferSimulationTransactionSourceInboundWireReversal using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInboundWireReversal) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceInboundWireReversal into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceInboundWireReversal) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount that was reversed.
func (r WireTransferSimulationTransactionSourceInboundWireReversal) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description on the reversal message from Fedwire.
func (r WireTransferSimulationTransactionSourceInboundWireReversal) GetDescription() (Description string) {
	if r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Fedwire cycle date for the wire reversal.
func (r WireTransferSimulationTransactionSourceInboundWireReversal) GetInputCycleDate() (InputCycleDate time.Time) {
	if r.InputCycleDate != nil {
		InputCycleDate = *r.InputCycleDate
	}
	return
}

// The Fedwire sequence number.
func (r WireTransferSimulationTransactionSourceInboundWireReversal) GetInputSequenceNumber() (InputSequenceNumber string) {
	if r.InputSequenceNumber != nil {
		InputSequenceNumber = *r.InputSequenceNumber
	}
	return
}

// The Fedwire input source identifier.
func (r WireTransferSimulationTransactionSourceInboundWireReversal) GetInputSource() (InputSource string) {
	if r.InputSource != nil {
		InputSource = *r.InputSource
	}
	return
}

// The Fedwire transaction identifier.
func (r WireTransferSimulationTransactionSourceInboundWireReversal) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

// The Fedwire transaction identifier for the wire transfer that was reversed.
func (r WireTransferSimulationTransactionSourceInboundWireReversal) GetPreviousMessageInputMessageAccountabilityData() (PreviousMessageInputMessageAccountabilityData string) {
	if r.PreviousMessageInputMessageAccountabilityData != nil {
		PreviousMessageInputMessageAccountabilityData = *r.PreviousMessageInputMessageAccountabilityData
	}
	return
}

// The Fedwire cycle date for the wire transfer that was reversed.
func (r WireTransferSimulationTransactionSourceInboundWireReversal) GetPreviousMessageInputCycleDate() (PreviousMessageInputCycleDate time.Time) {
	if r.PreviousMessageInputCycleDate != nil {
		PreviousMessageInputCycleDate = *r.PreviousMessageInputCycleDate
	}
	return
}

// The Fedwire sequence number for the wire transfer that was reversed.
func (r WireTransferSimulationTransactionSourceInboundWireReversal) GetPreviousMessageInputSequenceNumber() (PreviousMessageInputSequenceNumber string) {
	if r.PreviousMessageInputSequenceNumber != nil {
		PreviousMessageInputSequenceNumber = *r.PreviousMessageInputSequenceNumber
	}
	return
}

// The Fedwire input source identifier for the wire transfer that was reversed.
func (r WireTransferSimulationTransactionSourceInboundWireReversal) GetPreviousMessageInputSource() (PreviousMessageInputSource string) {
	if r.PreviousMessageInputSource != nil {
		PreviousMessageInputSource = *r.PreviousMessageInputSource
	}
	return
}

// Information included in the wire reversal for the receiving financial
// institution.
func (r WireTransferSimulationTransactionSourceInboundWireReversal) GetReceiverFinancialInstitutionInformation() (ReceiverFinancialInstitutionInformation string) {
	if r.ReceiverFinancialInstitutionInformation != nil {
		ReceiverFinancialInstitutionInformation = *r.ReceiverFinancialInstitutionInformation
	}
	return
}

// Additional financial institution information included in the wire reversal.
func (r WireTransferSimulationTransactionSourceInboundWireReversal) GetFinancialInstitutionToFinancialInstitutionInformation() (FinancialInstitutionToFinancialInstitutionInformation string) {
	if r.FinancialInstitutionToFinancialInstitutionInformation != nil {
		FinancialInstitutionToFinancialInstitutionInformation = *r.FinancialInstitutionToFinancialInstitutionInformation
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireReversal) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceInboundWireReversal{Amount:%s Description:%s InputCycleDate:%s InputSequenceNumber:%s InputSource:%s InputMessageAccountabilityData:%s PreviousMessageInputMessageAccountabilityData:%s PreviousMessageInputCycleDate:%s PreviousMessageInputSequenceNumber:%s PreviousMessageInputSource:%s ReceiverFinancialInstitutionInformation:%s FinancialInstitutionToFinancialInstitutionInformation:%s}", core.FmtP(r.Amount), core.FmtP(r.Description), core.FmtP(r.InputCycleDate), core.FmtP(r.InputSequenceNumber), core.FmtP(r.InputSource), core.FmtP(r.InputMessageAccountabilityData), core.FmtP(r.PreviousMessageInputMessageAccountabilityData), core.FmtP(r.PreviousMessageInputCycleDate), core.FmtP(r.PreviousMessageInputSequenceNumber), core.FmtP(r.PreviousMessageInputSource), core.FmtP(r.ReceiverFinancialInstitutionInformation), core.FmtP(r.FinancialInstitutionToFinancialInstitutionInformation))
}

type WireTransferSimulationTransactionSourceInboundWireTransfer struct {
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
// WireTransferSimulationTransactionSourceInboundWireTransfer using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInboundWireTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceInboundWireTransfer into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceInboundWireTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r WireTransferSimulationTransactionSourceInboundWireTransfer) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireTransfer) GetBeneficiaryName() (BeneficiaryName string) {
	if r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireTransfer) GetBeneficiaryReference() (BeneficiaryReference string) {
	if r.BeneficiaryReference != nil {
		BeneficiaryReference = *r.BeneficiaryReference
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireTransfer) GetDescription() (Description string) {
	if r.Description != nil {
		Description = *r.Description
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireTransfer) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorAddressLine1() (OriginatorAddressLine1 string) {
	if r.OriginatorAddressLine1 != nil {
		OriginatorAddressLine1 = *r.OriginatorAddressLine1
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorAddressLine2() (OriginatorAddressLine2 string) {
	if r.OriginatorAddressLine2 != nil {
		OriginatorAddressLine2 = *r.OriginatorAddressLine2
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorAddressLine3() (OriginatorAddressLine3 string) {
	if r.OriginatorAddressLine3 != nil {
		OriginatorAddressLine3 = *r.OriginatorAddressLine3
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorName() (OriginatorName string) {
	if r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorToBeneficiaryInformationLine1() (OriginatorToBeneficiaryInformationLine1 string) {
	if r.OriginatorToBeneficiaryInformationLine1 != nil {
		OriginatorToBeneficiaryInformationLine1 = *r.OriginatorToBeneficiaryInformationLine1
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorToBeneficiaryInformationLine2() (OriginatorToBeneficiaryInformationLine2 string) {
	if r.OriginatorToBeneficiaryInformationLine2 != nil {
		OriginatorToBeneficiaryInformationLine2 = *r.OriginatorToBeneficiaryInformationLine2
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorToBeneficiaryInformationLine3() (OriginatorToBeneficiaryInformationLine3 string) {
	if r.OriginatorToBeneficiaryInformationLine3 != nil {
		OriginatorToBeneficiaryInformationLine3 = *r.OriginatorToBeneficiaryInformationLine3
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorToBeneficiaryInformationLine4() (OriginatorToBeneficiaryInformationLine4 string) {
	if r.OriginatorToBeneficiaryInformationLine4 != nil {
		OriginatorToBeneficiaryInformationLine4 = *r.OriginatorToBeneficiaryInformationLine4
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorToBeneficiaryInformation() (OriginatorToBeneficiaryInformation string) {
	if r.OriginatorToBeneficiaryInformation != nil {
		OriginatorToBeneficiaryInformation = *r.OriginatorToBeneficiaryInformation
	}
	return
}

func (r WireTransferSimulationTransactionSourceInboundWireTransfer) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceInboundWireTransfer{Amount:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s BeneficiaryName:%s BeneficiaryReference:%s Description:%s InputMessageAccountabilityData:%s OriginatorAddressLine1:%s OriginatorAddressLine2:%s OriginatorAddressLine3:%s OriginatorName:%s OriginatorToBeneficiaryInformationLine1:%s OriginatorToBeneficiaryInformationLine2:%s OriginatorToBeneficiaryInformationLine3:%s OriginatorToBeneficiaryInformationLine4:%s OriginatorToBeneficiaryInformation:%s}", core.FmtP(r.Amount), core.FmtP(r.BeneficiaryAddressLine1), core.FmtP(r.BeneficiaryAddressLine2), core.FmtP(r.BeneficiaryAddressLine3), core.FmtP(r.BeneficiaryName), core.FmtP(r.BeneficiaryReference), core.FmtP(r.Description), core.FmtP(r.InputMessageAccountabilityData), core.FmtP(r.OriginatorAddressLine1), core.FmtP(r.OriginatorAddressLine2), core.FmtP(r.OriginatorAddressLine3), core.FmtP(r.OriginatorName), core.FmtP(r.OriginatorToBeneficiaryInformationLine1), core.FmtP(r.OriginatorToBeneficiaryInformationLine2), core.FmtP(r.OriginatorToBeneficiaryInformationLine3), core.FmtP(r.OriginatorToBeneficiaryInformationLine4), core.FmtP(r.OriginatorToBeneficiaryInformation))
}

type WireTransferSimulationTransactionSourceInterestPayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency *WireTransferSimulationTransactionSourceInterestPaymentCurrency `pjson:"currency"`
	// The start of the period for which this transaction paid interest.
	PeriodStart *time.Time `pjson:"period_start" format:"2006-01-02T15:04:05Z07:00"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd *time.Time `pjson:"period_end" format:"2006-01-02T15:04:05Z07:00"`
	// The account on which the interest was accrued.
	AccruedOnAccountID *string                `pjson:"accrued_on_account_id"`
	jsonFields         map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceInterestPayment using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInterestPayment) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceInterestPayment
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *WireTransferSimulationTransactionSourceInterestPayment) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r WireTransferSimulationTransactionSourceInterestPayment) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
func (r WireTransferSimulationTransactionSourceInterestPayment) GetCurrency() (Currency WireTransferSimulationTransactionSourceInterestPaymentCurrency) {
	if r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The start of the period for which this transaction paid interest.
func (r WireTransferSimulationTransactionSourceInterestPayment) GetPeriodStart() (PeriodStart time.Time) {
	if r.PeriodStart != nil {
		PeriodStart = *r.PeriodStart
	}
	return
}

// The end of the period for which this transaction paid interest.
func (r WireTransferSimulationTransactionSourceInterestPayment) GetPeriodEnd() (PeriodEnd time.Time) {
	if r.PeriodEnd != nil {
		PeriodEnd = *r.PeriodEnd
	}
	return
}

// The account on which the interest was accrued.
func (r WireTransferSimulationTransactionSourceInterestPayment) GetAccruedOnAccountID() (AccruedOnAccountID string) {
	if r.AccruedOnAccountID != nil {
		AccruedOnAccountID = *r.AccruedOnAccountID
	}
	return
}

func (r WireTransferSimulationTransactionSourceInterestPayment) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceInterestPayment{Amount:%s Currency:%s PeriodStart:%s PeriodEnd:%s AccruedOnAccountID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.PeriodStart), core.FmtP(r.PeriodEnd), core.FmtP(r.AccruedOnAccountID))
}

type WireTransferSimulationTransactionSourceInterestPaymentCurrency string

const (
	WireTransferSimulationTransactionSourceInterestPaymentCurrencyCad WireTransferSimulationTransactionSourceInterestPaymentCurrency = "CAD"
	WireTransferSimulationTransactionSourceInterestPaymentCurrencyChf WireTransferSimulationTransactionSourceInterestPaymentCurrency = "CHF"
	WireTransferSimulationTransactionSourceInterestPaymentCurrencyEur WireTransferSimulationTransactionSourceInterestPaymentCurrency = "EUR"
	WireTransferSimulationTransactionSourceInterestPaymentCurrencyGbp WireTransferSimulationTransactionSourceInterestPaymentCurrency = "GBP"
	WireTransferSimulationTransactionSourceInterestPaymentCurrencyJpy WireTransferSimulationTransactionSourceInterestPaymentCurrency = "JPY"
	WireTransferSimulationTransactionSourceInterestPaymentCurrencyUsd WireTransferSimulationTransactionSourceInterestPaymentCurrency = "USD"
)

type WireTransferSimulationTransactionSourceInternalSource struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency   *WireTransferSimulationTransactionSourceInternalSourceCurrency `pjson:"currency"`
	Reason     *WireTransferSimulationTransactionSourceInternalSourceReason   `pjson:"reason"`
	jsonFields map[string]interface{}                                         `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceInternalSource using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInternalSource) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceInternalSource
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *WireTransferSimulationTransactionSourceInternalSource) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r WireTransferSimulationTransactionSourceInternalSource) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
func (r WireTransferSimulationTransactionSourceInternalSource) GetCurrency() (Currency WireTransferSimulationTransactionSourceInternalSourceCurrency) {
	if r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r WireTransferSimulationTransactionSourceInternalSource) GetReason() (Reason WireTransferSimulationTransactionSourceInternalSourceReason) {
	if r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

func (r WireTransferSimulationTransactionSourceInternalSource) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceInternalSource{Amount:%s Currency:%s Reason:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.Reason))
}

type WireTransferSimulationTransactionSourceInternalSourceCurrency string

const (
	WireTransferSimulationTransactionSourceInternalSourceCurrencyCad WireTransferSimulationTransactionSourceInternalSourceCurrency = "CAD"
	WireTransferSimulationTransactionSourceInternalSourceCurrencyChf WireTransferSimulationTransactionSourceInternalSourceCurrency = "CHF"
	WireTransferSimulationTransactionSourceInternalSourceCurrencyEur WireTransferSimulationTransactionSourceInternalSourceCurrency = "EUR"
	WireTransferSimulationTransactionSourceInternalSourceCurrencyGbp WireTransferSimulationTransactionSourceInternalSourceCurrency = "GBP"
	WireTransferSimulationTransactionSourceInternalSourceCurrencyJpy WireTransferSimulationTransactionSourceInternalSourceCurrency = "JPY"
	WireTransferSimulationTransactionSourceInternalSourceCurrencyUsd WireTransferSimulationTransactionSourceInternalSourceCurrency = "USD"
)

type WireTransferSimulationTransactionSourceInternalSourceReason string

const (
	WireTransferSimulationTransactionSourceInternalSourceReasonBankMigration      WireTransferSimulationTransactionSourceInternalSourceReason = "bank_migration"
	WireTransferSimulationTransactionSourceInternalSourceReasonCashback           WireTransferSimulationTransactionSourceInternalSourceReason = "cashback"
	WireTransferSimulationTransactionSourceInternalSourceReasonEmpyrealAdjustment WireTransferSimulationTransactionSourceInternalSourceReason = "empyreal_adjustment"
	WireTransferSimulationTransactionSourceInternalSourceReasonError              WireTransferSimulationTransactionSourceInternalSourceReason = "error"
	WireTransferSimulationTransactionSourceInternalSourceReasonErrorCorrection    WireTransferSimulationTransactionSourceInternalSourceReason = "error_correction"
	WireTransferSimulationTransactionSourceInternalSourceReasonFees               WireTransferSimulationTransactionSourceInternalSourceReason = "fees"
	WireTransferSimulationTransactionSourceInternalSourceReasonInterest           WireTransferSimulationTransactionSourceInternalSourceReason = "interest"
	WireTransferSimulationTransactionSourceInternalSourceReasonSampleFunds        WireTransferSimulationTransactionSourceInternalSourceReason = "sample_funds"
	WireTransferSimulationTransactionSourceInternalSourceReasonSampleFundsReturn  WireTransferSimulationTransactionSourceInternalSourceReason = "sample_funds_return"
)

type WireTransferSimulationTransactionSourceCardRouteRefund struct {
	// The refunded amount in the minor unit of the refunded currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
	// currency.
	Currency             *WireTransferSimulationTransactionSourceCardRouteRefundCurrency `pjson:"currency"`
	MerchantAcceptorID   *string                                                         `pjson:"merchant_acceptor_id"`
	MerchantCity         *string                                                         `pjson:"merchant_city"`
	MerchantCountry      *string                                                         `pjson:"merchant_country"`
	MerchantDescriptor   *string                                                         `pjson:"merchant_descriptor"`
	MerchantState        *string                                                         `pjson:"merchant_state"`
	MerchantCategoryCode *string                                                         `pjson:"merchant_category_code"`
	jsonFields           map[string]interface{}                                          `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCardRouteRefund using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCardRouteRefund) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceCardRouteRefund
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *WireTransferSimulationTransactionSourceCardRouteRefund) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The refunded amount in the minor unit of the refunded currency. For dollars, for
// example, this is cents.
func (r WireTransferSimulationTransactionSourceCardRouteRefund) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
// currency.
func (r WireTransferSimulationTransactionSourceCardRouteRefund) GetCurrency() (Currency WireTransferSimulationTransactionSourceCardRouteRefundCurrency) {
	if r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardRouteRefund) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardRouteRefund) GetMerchantCity() (MerchantCity string) {
	if r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardRouteRefund) GetMerchantCountry() (MerchantCountry string) {
	if r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardRouteRefund) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardRouteRefund) GetMerchantState() (MerchantState string) {
	if r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardRouteRefund) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardRouteRefund) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCardRouteRefund{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.MerchantAcceptorID), core.FmtP(r.MerchantCity), core.FmtP(r.MerchantCountry), core.FmtP(r.MerchantDescriptor), core.FmtP(r.MerchantState), core.FmtP(r.MerchantCategoryCode))
}

type WireTransferSimulationTransactionSourceCardRouteRefundCurrency string

const (
	WireTransferSimulationTransactionSourceCardRouteRefundCurrencyCad WireTransferSimulationTransactionSourceCardRouteRefundCurrency = "CAD"
	WireTransferSimulationTransactionSourceCardRouteRefundCurrencyChf WireTransferSimulationTransactionSourceCardRouteRefundCurrency = "CHF"
	WireTransferSimulationTransactionSourceCardRouteRefundCurrencyEur WireTransferSimulationTransactionSourceCardRouteRefundCurrency = "EUR"
	WireTransferSimulationTransactionSourceCardRouteRefundCurrencyGbp WireTransferSimulationTransactionSourceCardRouteRefundCurrency = "GBP"
	WireTransferSimulationTransactionSourceCardRouteRefundCurrencyJpy WireTransferSimulationTransactionSourceCardRouteRefundCurrency = "JPY"
	WireTransferSimulationTransactionSourceCardRouteRefundCurrencyUsd WireTransferSimulationTransactionSourceCardRouteRefundCurrency = "USD"
)

type WireTransferSimulationTransactionSourceCardRouteSettlement struct {
	// The settled amount in the minor unit of the settlement currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
	// currency.
	Currency             *WireTransferSimulationTransactionSourceCardRouteSettlementCurrency `pjson:"currency"`
	MerchantAcceptorID   *string                                                             `pjson:"merchant_acceptor_id"`
	MerchantCity         *string                                                             `pjson:"merchant_city"`
	MerchantCountry      *string                                                             `pjson:"merchant_country"`
	MerchantDescriptor   *string                                                             `pjson:"merchant_descriptor"`
	MerchantState        *string                                                             `pjson:"merchant_state"`
	MerchantCategoryCode *string                                                             `pjson:"merchant_category_code"`
	jsonFields           map[string]interface{}                                              `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCardRouteSettlement using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCardRouteSettlement) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceCardRouteSettlement into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceCardRouteSettlement) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The settled amount in the minor unit of the settlement currency. For dollars,
// for example, this is cents.
func (r WireTransferSimulationTransactionSourceCardRouteSettlement) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
// currency.
func (r WireTransferSimulationTransactionSourceCardRouteSettlement) GetCurrency() (Currency WireTransferSimulationTransactionSourceCardRouteSettlementCurrency) {
	if r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantCity() (MerchantCity string) {
	if r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantCountry() (MerchantCountry string) {
	if r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantState() (MerchantState string) {
	if r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

func (r WireTransferSimulationTransactionSourceCardRouteSettlement) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCardRouteSettlement{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.MerchantAcceptorID), core.FmtP(r.MerchantCity), core.FmtP(r.MerchantCountry), core.FmtP(r.MerchantDescriptor), core.FmtP(r.MerchantState), core.FmtP(r.MerchantCategoryCode))
}

type WireTransferSimulationTransactionSourceCardRouteSettlementCurrency string

const (
	WireTransferSimulationTransactionSourceCardRouteSettlementCurrencyCad WireTransferSimulationTransactionSourceCardRouteSettlementCurrency = "CAD"
	WireTransferSimulationTransactionSourceCardRouteSettlementCurrencyChf WireTransferSimulationTransactionSourceCardRouteSettlementCurrency = "CHF"
	WireTransferSimulationTransactionSourceCardRouteSettlementCurrencyEur WireTransferSimulationTransactionSourceCardRouteSettlementCurrency = "EUR"
	WireTransferSimulationTransactionSourceCardRouteSettlementCurrencyGbp WireTransferSimulationTransactionSourceCardRouteSettlementCurrency = "GBP"
	WireTransferSimulationTransactionSourceCardRouteSettlementCurrencyJpy WireTransferSimulationTransactionSourceCardRouteSettlementCurrency = "JPY"
	WireTransferSimulationTransactionSourceCardRouteSettlementCurrencyUsd WireTransferSimulationTransactionSourceCardRouteSettlementCurrency = "USD"
)

type WireTransferSimulationTransactionSourceSampleFunds struct {
	// Where the sample funds came from.
	Originator *string                `pjson:"originator"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceSampleFunds using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceSampleFunds) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceSampleFunds into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *WireTransferSimulationTransactionSourceSampleFunds) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Where the sample funds came from.
func (r WireTransferSimulationTransactionSourceSampleFunds) GetOriginator() (Originator string) {
	if r.Originator != nil {
		Originator = *r.Originator
	}
	return
}

func (r WireTransferSimulationTransactionSourceSampleFunds) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceSampleFunds{Originator:%s}", core.FmtP(r.Originator))
}

type WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention struct {
	// The transfer amount in USD cents.
	Amount             *int64                 `pjson:"amount"`
	AccountNumber      *string                `pjson:"account_number"`
	RoutingNumber      *string                `pjson:"routing_number"`
	MessageToRecipient *string                `pjson:"message_to_recipient"`
	TransferID         *string                `pjson:"transfer_id"`
	jsonFields         map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The transfer amount in USD cents.
func (r WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention) GetAccountNumber() (AccountNumber string) {
	if r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention) GetRoutingNumber() (RoutingNumber string) {
	if r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention) GetMessageToRecipient() (MessageToRecipient string) {
	if r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention) GetTransferID() (TransferID string) {
	if r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s TransferID:%s}", core.FmtP(r.Amount), core.FmtP(r.AccountNumber), core.FmtP(r.RoutingNumber), core.FmtP(r.MessageToRecipient), core.FmtP(r.TransferID))
}

type WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection struct {
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

func (r WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection) GetTransferID() (TransferID string) {
	if r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection{TransferID:%s}", core.FmtP(r.TransferID))
}

type WireTransferSimulationTransactionSourceWireTransferIntention struct {
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
// WireTransferSimulationTransactionSourceWireTransferIntention using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceWireTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceWireTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceWireTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The transfer amount in USD cents.
func (r WireTransferSimulationTransactionSourceWireTransferIntention) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The destination account number.
func (r WireTransferSimulationTransactionSourceWireTransferIntention) GetAccountNumber() (AccountNumber string) {
	if r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r WireTransferSimulationTransactionSourceWireTransferIntention) GetRoutingNumber() (RoutingNumber string) {
	if r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The message that will show on the recipient's bank statement.
func (r WireTransferSimulationTransactionSourceWireTransferIntention) GetMessageToRecipient() (MessageToRecipient string) {
	if r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r WireTransferSimulationTransactionSourceWireTransferIntention) GetTransferID() (TransferID string) {
	if r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r WireTransferSimulationTransactionSourceWireTransferIntention) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceWireTransferIntention{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s TransferID:%s}", core.FmtP(r.Amount), core.FmtP(r.AccountNumber), core.FmtP(r.RoutingNumber), core.FmtP(r.MessageToRecipient), core.FmtP(r.TransferID))
}

type WireTransferSimulationTransactionSourceWireTransferRejection struct {
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceWireTransferRejection using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceWireTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceWireTransferRejection into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceWireTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

func (r WireTransferSimulationTransactionSourceWireTransferRejection) GetTransferID() (TransferID string) {
	if r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r WireTransferSimulationTransactionSourceWireTransferRejection) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceWireTransferRejection{TransferID:%s}", core.FmtP(r.TransferID))
}

type WireTransferSimulationTransactionType string

const (
	WireTransferSimulationTransactionTypeTransaction WireTransferSimulationTransactionType = "transaction"
)

type WireTransferSimulationType string

const (
	WireTransferSimulationTypeInboundWireTransferSimulationResult WireTransferSimulationType = "inbound_wire_transfer_simulation_result"
)

type SimulateAWireTransferToYourAccountParameters struct {
	// The identifier of the Account Number the inbound Wire Transfer is for.
	AccountNumberID *string `pjson:"account_number_id"`
	// The transfer amount in cents. Must be positive.
	Amount *int64 `pjson:"amount"`
	// The sending bank will set beneficiary_address_line1 in production. You can
	// simulate any value here.
	BeneficiaryAddressLine1 *string `pjson:"beneficiary_address_line1"`
	// The sending bank will set beneficiary_address_line2 in production. You can
	// simulate any value here.
	BeneficiaryAddressLine2 *string `pjson:"beneficiary_address_line2"`
	// The sending bank will set beneficiary_address_line3 in production. You can
	// simulate any value here.
	BeneficiaryAddressLine3 *string `pjson:"beneficiary_address_line3"`
	// The sending bank will set beneficiary_name in production. You can simulate any
	// value here.
	BeneficiaryName *string `pjson:"beneficiary_name"`
	// The sending bank will set beneficiary_reference in production. You can simulate
	// any value here.
	BeneficiaryReference *string `pjson:"beneficiary_reference"`
	// The sending bank will set originator_address_line1 in production. You can
	// simulate any value here.
	OriginatorAddressLine1 *string `pjson:"originator_address_line1"`
	// The sending bank will set originator_address_line2 in production. You can
	// simulate any value here.
	OriginatorAddressLine2 *string `pjson:"originator_address_line2"`
	// The sending bank will set originator_address_line3 in production. You can
	// simulate any value here.
	OriginatorAddressLine3 *string `pjson:"originator_address_line3"`
	// The sending bank will set originator_name in production. You can simulate any
	// value here.
	OriginatorName *string `pjson:"originator_name"`
	// The sending bank will set originator_to_beneficiary_information_line1 in
	// production. You can simulate any value here.
	OriginatorToBeneficiaryInformationLine1 *string `pjson:"originator_to_beneficiary_information_line1"`
	// The sending bank will set originator_to_beneficiary_information_line2 in
	// production. You can simulate any value here.
	OriginatorToBeneficiaryInformationLine2 *string `pjson:"originator_to_beneficiary_information_line2"`
	// The sending bank will set originator_to_beneficiary_information_line3 in
	// production. You can simulate any value here.
	OriginatorToBeneficiaryInformationLine3 *string `pjson:"originator_to_beneficiary_information_line3"`
	// The sending bank will set originator_to_beneficiary_information_line4 in
	// production. You can simulate any value here.
	OriginatorToBeneficiaryInformationLine4 *string                `pjson:"originator_to_beneficiary_information_line4"`
	jsonFields                              map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// SimulateAWireTransferToYourAccountParameters using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *SimulateAWireTransferToYourAccountParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes SimulateAWireTransferToYourAccountParameters into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *SimulateAWireTransferToYourAccountParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the Account Number the inbound Wire Transfer is for.
func (r SimulateAWireTransferToYourAccountParameters) GetAccountNumberID() (AccountNumberID string) {
	if r.AccountNumberID != nil {
		AccountNumberID = *r.AccountNumberID
	}
	return
}

// The transfer amount in cents. Must be positive.
func (r SimulateAWireTransferToYourAccountParameters) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The sending bank will set beneficiary_address_line1 in production. You can
// simulate any value here.
func (r SimulateAWireTransferToYourAccountParameters) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

// The sending bank will set beneficiary_address_line2 in production. You can
// simulate any value here.
func (r SimulateAWireTransferToYourAccountParameters) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

// The sending bank will set beneficiary_address_line3 in production. You can
// simulate any value here.
func (r SimulateAWireTransferToYourAccountParameters) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

// The sending bank will set beneficiary_name in production. You can simulate any
// value here.
func (r SimulateAWireTransferToYourAccountParameters) GetBeneficiaryName() (BeneficiaryName string) {
	if r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

// The sending bank will set beneficiary_reference in production. You can simulate
// any value here.
func (r SimulateAWireTransferToYourAccountParameters) GetBeneficiaryReference() (BeneficiaryReference string) {
	if r.BeneficiaryReference != nil {
		BeneficiaryReference = *r.BeneficiaryReference
	}
	return
}

// The sending bank will set originator_address_line1 in production. You can
// simulate any value here.
func (r SimulateAWireTransferToYourAccountParameters) GetOriginatorAddressLine1() (OriginatorAddressLine1 string) {
	if r.OriginatorAddressLine1 != nil {
		OriginatorAddressLine1 = *r.OriginatorAddressLine1
	}
	return
}

// The sending bank will set originator_address_line2 in production. You can
// simulate any value here.
func (r SimulateAWireTransferToYourAccountParameters) GetOriginatorAddressLine2() (OriginatorAddressLine2 string) {
	if r.OriginatorAddressLine2 != nil {
		OriginatorAddressLine2 = *r.OriginatorAddressLine2
	}
	return
}

// The sending bank will set originator_address_line3 in production. You can
// simulate any value here.
func (r SimulateAWireTransferToYourAccountParameters) GetOriginatorAddressLine3() (OriginatorAddressLine3 string) {
	if r.OriginatorAddressLine3 != nil {
		OriginatorAddressLine3 = *r.OriginatorAddressLine3
	}
	return
}

// The sending bank will set originator_name in production. You can simulate any
// value here.
func (r SimulateAWireTransferToYourAccountParameters) GetOriginatorName() (OriginatorName string) {
	if r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

// The sending bank will set originator_to_beneficiary_information_line1 in
// production. You can simulate any value here.
func (r SimulateAWireTransferToYourAccountParameters) GetOriginatorToBeneficiaryInformationLine1() (OriginatorToBeneficiaryInformationLine1 string) {
	if r.OriginatorToBeneficiaryInformationLine1 != nil {
		OriginatorToBeneficiaryInformationLine1 = *r.OriginatorToBeneficiaryInformationLine1
	}
	return
}

// The sending bank will set originator_to_beneficiary_information_line2 in
// production. You can simulate any value here.
func (r SimulateAWireTransferToYourAccountParameters) GetOriginatorToBeneficiaryInformationLine2() (OriginatorToBeneficiaryInformationLine2 string) {
	if r.OriginatorToBeneficiaryInformationLine2 != nil {
		OriginatorToBeneficiaryInformationLine2 = *r.OriginatorToBeneficiaryInformationLine2
	}
	return
}

// The sending bank will set originator_to_beneficiary_information_line3 in
// production. You can simulate any value here.
func (r SimulateAWireTransferToYourAccountParameters) GetOriginatorToBeneficiaryInformationLine3() (OriginatorToBeneficiaryInformationLine3 string) {
	if r.OriginatorToBeneficiaryInformationLine3 != nil {
		OriginatorToBeneficiaryInformationLine3 = *r.OriginatorToBeneficiaryInformationLine3
	}
	return
}

// The sending bank will set originator_to_beneficiary_information_line4 in
// production. You can simulate any value here.
func (r SimulateAWireTransferToYourAccountParameters) GetOriginatorToBeneficiaryInformationLine4() (OriginatorToBeneficiaryInformationLine4 string) {
	if r.OriginatorToBeneficiaryInformationLine4 != nil {
		OriginatorToBeneficiaryInformationLine4 = *r.OriginatorToBeneficiaryInformationLine4
	}
	return
}

func (r SimulateAWireTransferToYourAccountParameters) String() (result string) {
	return fmt.Sprintf("&SimulateAWireTransferToYourAccountParameters{AccountNumberID:%s Amount:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s BeneficiaryName:%s BeneficiaryReference:%s OriginatorAddressLine1:%s OriginatorAddressLine2:%s OriginatorAddressLine3:%s OriginatorName:%s OriginatorToBeneficiaryInformationLine1:%s OriginatorToBeneficiaryInformationLine2:%s OriginatorToBeneficiaryInformationLine3:%s OriginatorToBeneficiaryInformationLine4:%s}", core.FmtP(r.AccountNumberID), core.FmtP(r.Amount), core.FmtP(r.BeneficiaryAddressLine1), core.FmtP(r.BeneficiaryAddressLine2), core.FmtP(r.BeneficiaryAddressLine3), core.FmtP(r.BeneficiaryName), core.FmtP(r.BeneficiaryReference), core.FmtP(r.OriginatorAddressLine1), core.FmtP(r.OriginatorAddressLine2), core.FmtP(r.OriginatorAddressLine3), core.FmtP(r.OriginatorName), core.FmtP(r.OriginatorToBeneficiaryInformationLine1), core.FmtP(r.OriginatorToBeneficiaryInformationLine2), core.FmtP(r.OriginatorToBeneficiaryInformationLine3), core.FmtP(r.OriginatorToBeneficiaryInformationLine4))
}
