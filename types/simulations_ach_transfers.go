package types

import (
	"fmt"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
)

type ACHTransferSimulation struct {
	// If the ACH Transfer attempt succeeds, this will contain the resulting
	// [Transaction](#transactions) object. The Transaction's `source` will be of
	// `category: inbound_ach_transfer`.
	Transaction *ACHTransferSimulationTransaction `pjson:"transaction"`
	// If the ACH Transfer attempt fails, this will contain the resulting
	// [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of `category: inbound_ach_transfer`.
	DeclinedTransaction *ACHTransferSimulationDeclinedTransaction `pjson:"declined_transaction"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_ach_transfer_simulation_result`.
	Type       *ACHTransferSimulationType `pjson:"type"`
	jsonFields map[string]interface{}     `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into ACHTransferSimulation using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulation) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulation into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulation) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// If the ACH Transfer attempt succeeds, this will contain the resulting
// [Transaction](#transactions) object. The Transaction's `source` will be of
// `category: inbound_ach_transfer`.
func (r *ACHTransferSimulation) GetTransaction() (Transaction ACHTransferSimulationTransaction) {
	if r != nil && r.Transaction != nil {
		Transaction = *r.Transaction
	}
	return
}

// If the ACH Transfer attempt fails, this will contain the resulting
// [Declined Transaction](#declined-transactions) object. The Declined
// Transaction's `source` will be of `category: inbound_ach_transfer`.
func (r *ACHTransferSimulation) GetDeclinedTransaction() (DeclinedTransaction ACHTransferSimulationDeclinedTransaction) {
	if r != nil && r.DeclinedTransaction != nil {
		DeclinedTransaction = *r.DeclinedTransaction
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `inbound_ach_transfer_simulation_result`.
func (r *ACHTransferSimulation) GetType() (Type ACHTransferSimulationType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r ACHTransferSimulation) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulation{Transaction:%s DeclinedTransaction:%s Type:%s}", r.Transaction, r.DeclinedTransaction, core.FmtP(r.Type))
}

type ACHTransferSimulationTransaction struct {
	// The identifier for the Account the Transaction belongs to.
	AccountID *string `pjson:"account_id"`
	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transcation's
	// Account.
	Currency *ACHTransferSimulationTransactionCurrency `pjson:"currency"`
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
	Source *ACHTransferSimulationTransactionSource `pjson:"source"`
	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type       *ACHTransferSimulationTransactionType `pjson:"type"`
	jsonFields map[string]interface{}                `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransaction using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransaction) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationTransaction into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransaction) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier for the Account the Transaction belongs to.
func (r *ACHTransferSimulationTransaction) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
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
func (r *ACHTransferSimulationTransaction) GetCurrency() (Currency ACHTransferSimulationTransactionCurrency) {
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
func (r *ACHTransferSimulationTransaction) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The identifier for the route this Transaction came through. Routes are things
// like cards and ACH details.
func (r *ACHTransferSimulationTransaction) GetRouteID() (RouteID string) {
	if r != nil && r.RouteID != nil {
		RouteID = *r.RouteID
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
func (r *ACHTransferSimulationTransaction) GetSource() (Source ACHTransferSimulationTransactionSource) {
	if r != nil && r.Source != nil {
		Source = *r.Source
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `transaction`.
func (r *ACHTransferSimulationTransaction) GetType() (Type ACHTransferSimulationTransactionType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r ACHTransferSimulationTransaction) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransaction{AccountID:%s Amount:%s Currency:%s CreatedAt:%s Description:%s ID:%s RouteID:%s RouteType:%s Source:%s Type:%s}", core.FmtP(r.AccountID), core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.CreatedAt), core.FmtP(r.Description), core.FmtP(r.ID), core.FmtP(r.RouteID), core.FmtP(r.RouteType), r.Source, core.FmtP(r.Type))
}

type ACHTransferSimulationTransactionCurrency string

const (
	ACHTransferSimulationTransactionCurrencyCad ACHTransferSimulationTransactionCurrency = "CAD"
	ACHTransferSimulationTransactionCurrencyChf ACHTransferSimulationTransactionCurrency = "CHF"
	ACHTransferSimulationTransactionCurrencyEur ACHTransferSimulationTransactionCurrency = "EUR"
	ACHTransferSimulationTransactionCurrencyGbp ACHTransferSimulationTransactionCurrency = "GBP"
	ACHTransferSimulationTransactionCurrencyJpy ACHTransferSimulationTransactionCurrency = "JPY"
	ACHTransferSimulationTransactionCurrencyUsd ACHTransferSimulationTransactionCurrency = "USD"
)

type ACHTransferSimulationTransactionSource struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category *ACHTransferSimulationTransactionSourceCategory `pjson:"category"`
	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention *ACHTransferSimulationTransactionSourceAccountTransferIntention `pjson:"account_transfer_intention"`
	// A ACH Check Conversion Return object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_check_conversion_return`.
	ACHCheckConversionReturn *ACHTransferSimulationTransactionSourceACHCheckConversionReturn `pjson:"ach_check_conversion_return"`
	// A ACH Check Conversion object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_check_conversion`.
	ACHCheckConversion *ACHTransferSimulationTransactionSourceACHCheckConversion `pjson:"ach_check_conversion"`
	// A ACH Transfer Intention object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention *ACHTransferSimulationTransactionSourceACHTransferIntention `pjson:"ach_transfer_intention"`
	// A ACH Transfer Rejection object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection *ACHTransferSimulationTransactionSourceACHTransferRejection `pjson:"ach_transfer_rejection"`
	// A ACH Transfer Return object. This field will be present in the JSON response if
	// and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn *ACHTransferSimulationTransactionSourceACHTransferReturn `pjson:"ach_transfer_return"`
	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance *ACHTransferSimulationTransactionSourceCardDisputeAcceptance `pjson:"card_dispute_acceptance"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund *ACHTransferSimulationTransactionSourceCardRefund `pjson:"card_refund"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement *ACHTransferSimulationTransactionSourceCardSettlement `pjson:"card_settlement"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance *ACHTransferSimulationTransactionSourceCheckDepositAcceptance `pjson:"check_deposit_acceptance"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn *ACHTransferSimulationTransactionSourceCheckDepositReturn `pjson:"check_deposit_return"`
	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention *ACHTransferSimulationTransactionSourceCheckTransferIntention `pjson:"check_transfer_intention"`
	// A Check Transfer Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_transfer_return`.
	CheckTransferReturn *ACHTransferSimulationTransactionSourceCheckTransferReturn `pjson:"check_transfer_return"`
	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection *ACHTransferSimulationTransactionSourceCheckTransferRejection `pjson:"check_transfer_rejection"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest *ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest `pjson:"check_transfer_stop_payment_request"`
	// A Dispute Resolution object. This field will be present in the JSON response if
	// and only if `category` is equal to `dispute_resolution`.
	DisputeResolution *ACHTransferSimulationTransactionSourceDisputeResolution `pjson:"dispute_resolution"`
	// A Empyreal Cash Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `empyreal_cash_deposit`.
	EmpyrealCashDeposit *ACHTransferSimulationTransactionSourceEmpyrealCashDeposit `pjson:"empyreal_cash_deposit"`
	// A Inbound ACH Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer *ACHTransferSimulationTransactionSourceInboundACHTransfer `pjson:"inbound_ach_transfer"`
	// A Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck *ACHTransferSimulationTransactionSourceInboundCheck `pjson:"inbound_check"`
	// A Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer `pjson:"inbound_international_ach_transfer"`
	// A Inbound Real Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation `pjson:"inbound_real_time_payments_transfer_confirmation"`
	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal `pjson:"inbound_wire_drawdown_payment_reversal"`
	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment `pjson:"inbound_wire_drawdown_payment"`
	// A Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal *ACHTransferSimulationTransactionSourceInboundWireReversal `pjson:"inbound_wire_reversal"`
	// A Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer *ACHTransferSimulationTransactionSourceInboundWireTransfer `pjson:"inbound_wire_transfer"`
	// A Interest Payment object. This field will be present in the JSON response if
	// and only if `category` is equal to `interest_payment`.
	InterestPayment *ACHTransferSimulationTransactionSourceInterestPayment `pjson:"interest_payment"`
	// A Internal Source object. This field will be present in the JSON response if and
	// only if `category` is equal to `internal_source`.
	InternalSource *ACHTransferSimulationTransactionSourceInternalSource `pjson:"internal_source"`
	// A Deprecated Card Refund object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_route_refund`.
	CardRouteRefund *ACHTransferSimulationTransactionSourceCardRouteRefund `pjson:"card_route_refund"`
	// A Deprecated Card Settlement object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_settlement`.
	CardRouteSettlement *ACHTransferSimulationTransactionSourceCardRouteSettlement `pjson:"card_route_settlement"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds *ACHTransferSimulationTransactionSourceSampleFunds `pjson:"sample_funds"`
	// A Wire Drawdown Payment Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_intention`.
	WireDrawdownPaymentIntention *ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention `pjson:"wire_drawdown_payment_intention"`
	// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_rejection`.
	WireDrawdownPaymentRejection *ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection `pjson:"wire_drawdown_payment_rejection"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention *ACHTransferSimulationTransactionSourceWireTransferIntention `pjson:"wire_transfer_intention"`
	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection *ACHTransferSimulationTransactionSourceWireTransferRejection `pjson:"wire_transfer_rejection"`
	jsonFields            map[string]interface{}                                       `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSource using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationTransactionSource into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSource) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The type of transaction that took place. We may add additional possible values
// for this enum over time; your application should be able to handle such
// additions gracefully.
func (r *ACHTransferSimulationTransactionSource) GetCategory() (Category ACHTransferSimulationTransactionSourceCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// A Account Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_intention`.
func (r *ACHTransferSimulationTransactionSource) GetAccountTransferIntention() (AccountTransferIntention ACHTransferSimulationTransactionSourceAccountTransferIntention) {
	if r != nil && r.AccountTransferIntention != nil {
		AccountTransferIntention = *r.AccountTransferIntention
	}
	return
}

// A ACH Check Conversion Return object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_check_conversion_return`.
func (r *ACHTransferSimulationTransactionSource) GetACHCheckConversionReturn() (ACHCheckConversionReturn ACHTransferSimulationTransactionSourceACHCheckConversionReturn) {
	if r != nil && r.ACHCheckConversionReturn != nil {
		ACHCheckConversionReturn = *r.ACHCheckConversionReturn
	}
	return
}

// A ACH Check Conversion object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_check_conversion`.
func (r *ACHTransferSimulationTransactionSource) GetACHCheckConversion() (ACHCheckConversion ACHTransferSimulationTransactionSourceACHCheckConversion) {
	if r != nil && r.ACHCheckConversion != nil {
		ACHCheckConversion = *r.ACHCheckConversion
	}
	return
}

// A ACH Transfer Intention object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_intention`.
func (r *ACHTransferSimulationTransactionSource) GetACHTransferIntention() (ACHTransferIntention ACHTransferSimulationTransactionSourceACHTransferIntention) {
	if r != nil && r.ACHTransferIntention != nil {
		ACHTransferIntention = *r.ACHTransferIntention
	}
	return
}

// A ACH Transfer Rejection object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_rejection`.
func (r *ACHTransferSimulationTransactionSource) GetACHTransferRejection() (ACHTransferRejection ACHTransferSimulationTransactionSourceACHTransferRejection) {
	if r != nil && r.ACHTransferRejection != nil {
		ACHTransferRejection = *r.ACHTransferRejection
	}
	return
}

// A ACH Transfer Return object. This field will be present in the JSON response if
// and only if `category` is equal to `ach_transfer_return`.
func (r *ACHTransferSimulationTransactionSource) GetACHTransferReturn() (ACHTransferReturn ACHTransferSimulationTransactionSourceACHTransferReturn) {
	if r != nil && r.ACHTransferReturn != nil {
		ACHTransferReturn = *r.ACHTransferReturn
	}
	return
}

// A Card Dispute Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `card_dispute_acceptance`.
func (r *ACHTransferSimulationTransactionSource) GetCardDisputeAcceptance() (CardDisputeAcceptance ACHTransferSimulationTransactionSourceCardDisputeAcceptance) {
	if r != nil && r.CardDisputeAcceptance != nil {
		CardDisputeAcceptance = *r.CardDisputeAcceptance
	}
	return
}

// A Card Refund object. This field will be present in the JSON response if and
// only if `category` is equal to `card_refund`.
func (r *ACHTransferSimulationTransactionSource) GetCardRefund() (CardRefund ACHTransferSimulationTransactionSourceCardRefund) {
	if r != nil && r.CardRefund != nil {
		CardRefund = *r.CardRefund
	}
	return
}

// A Card Settlement object. This field will be present in the JSON response if and
// only if `category` is equal to `card_settlement`.
func (r *ACHTransferSimulationTransactionSource) GetCardSettlement() (CardSettlement ACHTransferSimulationTransactionSourceCardSettlement) {
	if r != nil && r.CardSettlement != nil {
		CardSettlement = *r.CardSettlement
	}
	return
}

// A Check Deposit Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_acceptance`.
func (r *ACHTransferSimulationTransactionSource) GetCheckDepositAcceptance() (CheckDepositAcceptance ACHTransferSimulationTransactionSourceCheckDepositAcceptance) {
	if r != nil && r.CheckDepositAcceptance != nil {
		CheckDepositAcceptance = *r.CheckDepositAcceptance
	}
	return
}

// A Check Deposit Return object. This field will be present in the JSON response
// if and only if `category` is equal to `check_deposit_return`.
func (r *ACHTransferSimulationTransactionSource) GetCheckDepositReturn() (CheckDepositReturn ACHTransferSimulationTransactionSourceCheckDepositReturn) {
	if r != nil && r.CheckDepositReturn != nil {
		CheckDepositReturn = *r.CheckDepositReturn
	}
	return
}

// A Check Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_intention`.
func (r *ACHTransferSimulationTransactionSource) GetCheckTransferIntention() (CheckTransferIntention ACHTransferSimulationTransactionSourceCheckTransferIntention) {
	if r != nil && r.CheckTransferIntention != nil {
		CheckTransferIntention = *r.CheckTransferIntention
	}
	return
}

// A Check Transfer Return object. This field will be present in the JSON response
// if and only if `category` is equal to `check_transfer_return`.
func (r *ACHTransferSimulationTransactionSource) GetCheckTransferReturn() (CheckTransferReturn ACHTransferSimulationTransactionSourceCheckTransferReturn) {
	if r != nil && r.CheckTransferReturn != nil {
		CheckTransferReturn = *r.CheckTransferReturn
	}
	return
}

// A Check Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_rejection`.
func (r *ACHTransferSimulationTransactionSource) GetCheckTransferRejection() (CheckTransferRejection ACHTransferSimulationTransactionSourceCheckTransferRejection) {
	if r != nil && r.CheckTransferRejection != nil {
		CheckTransferRejection = *r.CheckTransferRejection
	}
	return
}

// A Check Transfer Stop Payment Request object. This field will be present in the
// JSON response if and only if `category` is equal to
// `check_transfer_stop_payment_request`.
func (r *ACHTransferSimulationTransactionSource) GetCheckTransferStopPaymentRequest() (CheckTransferStopPaymentRequest ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) {
	if r != nil && r.CheckTransferStopPaymentRequest != nil {
		CheckTransferStopPaymentRequest = *r.CheckTransferStopPaymentRequest
	}
	return
}

// A Dispute Resolution object. This field will be present in the JSON response if
// and only if `category` is equal to `dispute_resolution`.
func (r *ACHTransferSimulationTransactionSource) GetDisputeResolution() (DisputeResolution ACHTransferSimulationTransactionSourceDisputeResolution) {
	if r != nil && r.DisputeResolution != nil {
		DisputeResolution = *r.DisputeResolution
	}
	return
}

// A Empyreal Cash Deposit object. This field will be present in the JSON response
// if and only if `category` is equal to `empyreal_cash_deposit`.
func (r *ACHTransferSimulationTransactionSource) GetEmpyrealCashDeposit() (EmpyrealCashDeposit ACHTransferSimulationTransactionSourceEmpyrealCashDeposit) {
	if r != nil && r.EmpyrealCashDeposit != nil {
		EmpyrealCashDeposit = *r.EmpyrealCashDeposit
	}
	return
}

// A Inbound ACH Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_ach_transfer`.
func (r *ACHTransferSimulationTransactionSource) GetInboundACHTransfer() (InboundACHTransfer ACHTransferSimulationTransactionSourceInboundACHTransfer) {
	if r != nil && r.InboundACHTransfer != nil {
		InboundACHTransfer = *r.InboundACHTransfer
	}
	return
}

// A Inbound Check object. This field will be present in the JSON response if and
// only if `category` is equal to `inbound_check`.
func (r *ACHTransferSimulationTransactionSource) GetInboundCheck() (InboundCheck ACHTransferSimulationTransactionSourceInboundCheck) {
	if r != nil && r.InboundCheck != nil {
		InboundCheck = *r.InboundCheck
	}
	return
}

// A Inbound International ACH Transfer object. This field will be present in the
// JSON response if and only if `category` is equal to
// `inbound_international_ach_transfer`.
func (r *ACHTransferSimulationTransactionSource) GetInboundInternationalACHTransfer() (InboundInternationalACHTransfer ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) {
	if r != nil && r.InboundInternationalACHTransfer != nil {
		InboundInternationalACHTransfer = *r.InboundInternationalACHTransfer
	}
	return
}

// A Inbound Real Time Payments Transfer Confirmation object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_confirmation`.
func (r *ACHTransferSimulationTransactionSource) GetInboundRealTimePaymentsTransferConfirmation() (InboundRealTimePaymentsTransferConfirmation ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) {
	if r != nil && r.InboundRealTimePaymentsTransferConfirmation != nil {
		InboundRealTimePaymentsTransferConfirmation = *r.InboundRealTimePaymentsTransferConfirmation
	}
	return
}

// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
// the JSON response if and only if `category` is equal to
// `inbound_wire_drawdown_payment_reversal`.
func (r *ACHTransferSimulationTransactionSource) GetInboundWireDrawdownPaymentReversal() (InboundWireDrawdownPaymentReversal ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) {
	if r != nil && r.InboundWireDrawdownPaymentReversal != nil {
		InboundWireDrawdownPaymentReversal = *r.InboundWireDrawdownPaymentReversal
	}
	return
}

// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
func (r *ACHTransferSimulationTransactionSource) GetInboundWireDrawdownPayment() (InboundWireDrawdownPayment ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) {
	if r != nil && r.InboundWireDrawdownPayment != nil {
		InboundWireDrawdownPayment = *r.InboundWireDrawdownPayment
	}
	return
}

// A Inbound Wire Reversal object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_reversal`.
func (r *ACHTransferSimulationTransactionSource) GetInboundWireReversal() (InboundWireReversal ACHTransferSimulationTransactionSourceInboundWireReversal) {
	if r != nil && r.InboundWireReversal != nil {
		InboundWireReversal = *r.InboundWireReversal
	}
	return
}

// A Inbound Wire Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_transfer`.
func (r *ACHTransferSimulationTransactionSource) GetInboundWireTransfer() (InboundWireTransfer ACHTransferSimulationTransactionSourceInboundWireTransfer) {
	if r != nil && r.InboundWireTransfer != nil {
		InboundWireTransfer = *r.InboundWireTransfer
	}
	return
}

// A Interest Payment object. This field will be present in the JSON response if
// and only if `category` is equal to `interest_payment`.
func (r *ACHTransferSimulationTransactionSource) GetInterestPayment() (InterestPayment ACHTransferSimulationTransactionSourceInterestPayment) {
	if r != nil && r.InterestPayment != nil {
		InterestPayment = *r.InterestPayment
	}
	return
}

// A Internal Source object. This field will be present in the JSON response if and
// only if `category` is equal to `internal_source`.
func (r *ACHTransferSimulationTransactionSource) GetInternalSource() (InternalSource ACHTransferSimulationTransactionSourceInternalSource) {
	if r != nil && r.InternalSource != nil {
		InternalSource = *r.InternalSource
	}
	return
}

// A Deprecated Card Refund object. This field will be present in the JSON response
// if and only if `category` is equal to `card_route_refund`.
func (r *ACHTransferSimulationTransactionSource) GetCardRouteRefund() (CardRouteRefund ACHTransferSimulationTransactionSourceCardRouteRefund) {
	if r != nil && r.CardRouteRefund != nil {
		CardRouteRefund = *r.CardRouteRefund
	}
	return
}

// A Deprecated Card Settlement object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_settlement`.
func (r *ACHTransferSimulationTransactionSource) GetCardRouteSettlement() (CardRouteSettlement ACHTransferSimulationTransactionSourceCardRouteSettlement) {
	if r != nil && r.CardRouteSettlement != nil {
		CardRouteSettlement = *r.CardRouteSettlement
	}
	return
}

// A Sample Funds object. This field will be present in the JSON response if and
// only if `category` is equal to `sample_funds`.
func (r *ACHTransferSimulationTransactionSource) GetSampleFunds() (SampleFunds ACHTransferSimulationTransactionSourceSampleFunds) {
	if r != nil && r.SampleFunds != nil {
		SampleFunds = *r.SampleFunds
	}
	return
}

// A Wire Drawdown Payment Intention object. This field will be present in the JSON
// response if and only if `category` is equal to
// `wire_drawdown_payment_intention`.
func (r *ACHTransferSimulationTransactionSource) GetWireDrawdownPaymentIntention() (WireDrawdownPaymentIntention ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention) {
	if r != nil && r.WireDrawdownPaymentIntention != nil {
		WireDrawdownPaymentIntention = *r.WireDrawdownPaymentIntention
	}
	return
}

// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to
// `wire_drawdown_payment_rejection`.
func (r *ACHTransferSimulationTransactionSource) GetWireDrawdownPaymentRejection() (WireDrawdownPaymentRejection ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection) {
	if r != nil && r.WireDrawdownPaymentRejection != nil {
		WireDrawdownPaymentRejection = *r.WireDrawdownPaymentRejection
	}
	return
}

// A Wire Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_intention`.
func (r *ACHTransferSimulationTransactionSource) GetWireTransferIntention() (WireTransferIntention ACHTransferSimulationTransactionSourceWireTransferIntention) {
	if r != nil && r.WireTransferIntention != nil {
		WireTransferIntention = *r.WireTransferIntention
	}
	return
}

// A Wire Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_rejection`.
func (r *ACHTransferSimulationTransactionSource) GetWireTransferRejection() (WireTransferRejection ACHTransferSimulationTransactionSourceWireTransferRejection) {
	if r != nil && r.WireTransferRejection != nil {
		WireTransferRejection = *r.WireTransferRejection
	}
	return
}

func (r ACHTransferSimulationTransactionSource) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSource{Category:%s AccountTransferIntention:%s ACHCheckConversionReturn:%s ACHCheckConversion:%s ACHTransferIntention:%s ACHTransferRejection:%s ACHTransferReturn:%s CardDisputeAcceptance:%s CardRefund:%s CardSettlement:%s CheckDepositAcceptance:%s CheckDepositReturn:%s CheckTransferIntention:%s CheckTransferReturn:%s CheckTransferRejection:%s CheckTransferStopPaymentRequest:%s DisputeResolution:%s EmpyrealCashDeposit:%s InboundACHTransfer:%s InboundCheck:%s InboundInternationalACHTransfer:%s InboundRealTimePaymentsTransferConfirmation:%s InboundWireDrawdownPaymentReversal:%s InboundWireDrawdownPayment:%s InboundWireReversal:%s InboundWireTransfer:%s InterestPayment:%s InternalSource:%s CardRouteRefund:%s CardRouteSettlement:%s SampleFunds:%s WireDrawdownPaymentIntention:%s WireDrawdownPaymentRejection:%s WireTransferIntention:%s WireTransferRejection:%s}", core.FmtP(r.Category), r.AccountTransferIntention, r.ACHCheckConversionReturn, r.ACHCheckConversion, r.ACHTransferIntention, r.ACHTransferRejection, r.ACHTransferReturn, r.CardDisputeAcceptance, r.CardRefund, r.CardSettlement, r.CheckDepositAcceptance, r.CheckDepositReturn, r.CheckTransferIntention, r.CheckTransferReturn, r.CheckTransferRejection, r.CheckTransferStopPaymentRequest, r.DisputeResolution, r.EmpyrealCashDeposit, r.InboundACHTransfer, r.InboundCheck, r.InboundInternationalACHTransfer, r.InboundRealTimePaymentsTransferConfirmation, r.InboundWireDrawdownPaymentReversal, r.InboundWireDrawdownPayment, r.InboundWireReversal, r.InboundWireTransfer, r.InterestPayment, r.InternalSource, r.CardRouteRefund, r.CardRouteSettlement, r.SampleFunds, r.WireDrawdownPaymentIntention, r.WireDrawdownPaymentRejection, r.WireTransferIntention, r.WireTransferRejection)
}

type ACHTransferSimulationTransactionSourceCategory string

const (
	ACHTransferSimulationTransactionSourceCategoryAccountTransferIntention                    ACHTransferSimulationTransactionSourceCategory = "account_transfer_intention"
	ACHTransferSimulationTransactionSourceCategoryACHCheckConversionReturn                    ACHTransferSimulationTransactionSourceCategory = "ach_check_conversion_return"
	ACHTransferSimulationTransactionSourceCategoryACHCheckConversion                          ACHTransferSimulationTransactionSourceCategory = "ach_check_conversion"
	ACHTransferSimulationTransactionSourceCategoryACHTransferIntention                        ACHTransferSimulationTransactionSourceCategory = "ach_transfer_intention"
	ACHTransferSimulationTransactionSourceCategoryACHTransferRejection                        ACHTransferSimulationTransactionSourceCategory = "ach_transfer_rejection"
	ACHTransferSimulationTransactionSourceCategoryACHTransferReturn                           ACHTransferSimulationTransactionSourceCategory = "ach_transfer_return"
	ACHTransferSimulationTransactionSourceCategoryCardDisputeAcceptance                       ACHTransferSimulationTransactionSourceCategory = "card_dispute_acceptance"
	ACHTransferSimulationTransactionSourceCategoryCardRefund                                  ACHTransferSimulationTransactionSourceCategory = "card_refund"
	ACHTransferSimulationTransactionSourceCategoryCardSettlement                              ACHTransferSimulationTransactionSourceCategory = "card_settlement"
	ACHTransferSimulationTransactionSourceCategoryCheckDepositAcceptance                      ACHTransferSimulationTransactionSourceCategory = "check_deposit_acceptance"
	ACHTransferSimulationTransactionSourceCategoryCheckDepositReturn                          ACHTransferSimulationTransactionSourceCategory = "check_deposit_return"
	ACHTransferSimulationTransactionSourceCategoryCheckTransferIntention                      ACHTransferSimulationTransactionSourceCategory = "check_transfer_intention"
	ACHTransferSimulationTransactionSourceCategoryCheckTransferReturn                         ACHTransferSimulationTransactionSourceCategory = "check_transfer_return"
	ACHTransferSimulationTransactionSourceCategoryCheckTransferRejection                      ACHTransferSimulationTransactionSourceCategory = "check_transfer_rejection"
	ACHTransferSimulationTransactionSourceCategoryCheckTransferStopPaymentRequest             ACHTransferSimulationTransactionSourceCategory = "check_transfer_stop_payment_request"
	ACHTransferSimulationTransactionSourceCategoryDisputeResolution                           ACHTransferSimulationTransactionSourceCategory = "dispute_resolution"
	ACHTransferSimulationTransactionSourceCategoryEmpyrealCashDeposit                         ACHTransferSimulationTransactionSourceCategory = "empyreal_cash_deposit"
	ACHTransferSimulationTransactionSourceCategoryInboundACHTransfer                          ACHTransferSimulationTransactionSourceCategory = "inbound_ach_transfer"
	ACHTransferSimulationTransactionSourceCategoryInboundACHTransferReturnIntention           ACHTransferSimulationTransactionSourceCategory = "inbound_ach_transfer_return_intention"
	ACHTransferSimulationTransactionSourceCategoryInboundCheck                                ACHTransferSimulationTransactionSourceCategory = "inbound_check"
	ACHTransferSimulationTransactionSourceCategoryInboundInternationalACHTransfer             ACHTransferSimulationTransactionSourceCategory = "inbound_international_ach_transfer"
	ACHTransferSimulationTransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation ACHTransferSimulationTransactionSourceCategory = "inbound_real_time_payments_transfer_confirmation"
	ACHTransferSimulationTransactionSourceCategoryInboundWireDrawdownPaymentReversal          ACHTransferSimulationTransactionSourceCategory = "inbound_wire_drawdown_payment_reversal"
	ACHTransferSimulationTransactionSourceCategoryInboundWireDrawdownPayment                  ACHTransferSimulationTransactionSourceCategory = "inbound_wire_drawdown_payment"
	ACHTransferSimulationTransactionSourceCategoryInboundWireReversal                         ACHTransferSimulationTransactionSourceCategory = "inbound_wire_reversal"
	ACHTransferSimulationTransactionSourceCategoryInboundWireTransfer                         ACHTransferSimulationTransactionSourceCategory = "inbound_wire_transfer"
	ACHTransferSimulationTransactionSourceCategoryInterestPayment                             ACHTransferSimulationTransactionSourceCategory = "interest_payment"
	ACHTransferSimulationTransactionSourceCategoryInternalGeneralLedgerTransaction            ACHTransferSimulationTransactionSourceCategory = "internal_general_ledger_transaction"
	ACHTransferSimulationTransactionSourceCategoryInternalSource                              ACHTransferSimulationTransactionSourceCategory = "internal_source"
	ACHTransferSimulationTransactionSourceCategoryCardRouteRefund                             ACHTransferSimulationTransactionSourceCategory = "card_route_refund"
	ACHTransferSimulationTransactionSourceCategoryCardRouteSettlement                         ACHTransferSimulationTransactionSourceCategory = "card_route_settlement"
	ACHTransferSimulationTransactionSourceCategoryRealTimePaymentsTransferAcknowledgement     ACHTransferSimulationTransactionSourceCategory = "real_time_payments_transfer_acknowledgement"
	ACHTransferSimulationTransactionSourceCategorySampleFunds                                 ACHTransferSimulationTransactionSourceCategory = "sample_funds"
	ACHTransferSimulationTransactionSourceCategoryWireDrawdownPaymentIntention                ACHTransferSimulationTransactionSourceCategory = "wire_drawdown_payment_intention"
	ACHTransferSimulationTransactionSourceCategoryWireDrawdownPaymentRejection                ACHTransferSimulationTransactionSourceCategory = "wire_drawdown_payment_rejection"
	ACHTransferSimulationTransactionSourceCategoryWireTransferIntention                       ACHTransferSimulationTransactionSourceCategory = "wire_transfer_intention"
	ACHTransferSimulationTransactionSourceCategoryWireTransferRejection                       ACHTransferSimulationTransactionSourceCategory = "wire_transfer_rejection"
	ACHTransferSimulationTransactionSourceCategoryOther                                       ACHTransferSimulationTransactionSourceCategory = "other"
)

type ACHTransferSimulationTransactionSourceAccountTransferIntention struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency `pjson:"currency"`
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
// ACHTransferSimulationTransactionSourceAccountTransferIntention using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationTransactionSourceAccountTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceAccountTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSourceAccountTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *ACHTransferSimulationTransactionSourceAccountTransferIntention) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *ACHTransferSimulationTransactionSourceAccountTransferIntention) GetCurrency() (Currency ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The description you chose to give the transfer.
func (r *ACHTransferSimulationTransactionSourceAccountTransferIntention) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The identifier of the Account to where the Account Transfer was sent.
func (r *ACHTransferSimulationTransactionSourceAccountTransferIntention) GetDestinationAccountID() (DestinationAccountID string) {
	if r != nil && r.DestinationAccountID != nil {
		DestinationAccountID = *r.DestinationAccountID
	}
	return
}

// The identifier of the Account from where the Account Transfer was sent.
func (r *ACHTransferSimulationTransactionSourceAccountTransferIntention) GetSourceAccountID() (SourceAccountID string) {
	if r != nil && r.SourceAccountID != nil {
		SourceAccountID = *r.SourceAccountID
	}
	return
}

// The identifier of the Account Transfer that led to this Pending Transaction.
func (r *ACHTransferSimulationTransactionSourceAccountTransferIntention) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r ACHTransferSimulationTransactionSourceAccountTransferIntention) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceAccountTransferIntention{Amount:%s Currency:%s Description:%s DestinationAccountID:%s SourceAccountID:%s TransferID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.Description), core.FmtP(r.DestinationAccountID), core.FmtP(r.SourceAccountID), core.FmtP(r.TransferID))
}

type ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency string

const (
	ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrencyCad ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "CAD"
	ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrencyChf ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "CHF"
	ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrencyEur ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "EUR"
	ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrencyGbp ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "GBP"
	ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrencyJpy ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "JPY"
	ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrencyUsd ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceACHCheckConversionReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// Why the transfer was returned.
	ReturnReasonCode *string                `pjson:"return_reason_code"`
	jsonFields       map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceACHCheckConversionReturn using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationTransactionSourceACHCheckConversionReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceACHCheckConversionReturn into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSourceACHCheckConversionReturn) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransactionSourceACHCheckConversionReturn) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// Why the transfer was returned.
func (r *ACHTransferSimulationTransactionSourceACHCheckConversionReturn) GetReturnReasonCode() (ReturnReasonCode string) {
	if r != nil && r.ReturnReasonCode != nil {
		ReturnReasonCode = *r.ReturnReasonCode
	}
	return
}

func (r ACHTransferSimulationTransactionSourceACHCheckConversionReturn) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceACHCheckConversionReturn{Amount:%s ReturnReasonCode:%s}", core.FmtP(r.Amount), core.FmtP(r.ReturnReasonCode))
}

type ACHTransferSimulationTransactionSourceACHCheckConversion struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The identifier of the File containing an image of the returned check.
	FileID     *string                `pjson:"file_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceACHCheckConversion using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceACHCheckConversion) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceACHCheckConversion
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceACHCheckConversion) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransactionSourceACHCheckConversion) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The identifier of the File containing an image of the returned check.
func (r *ACHTransferSimulationTransactionSourceACHCheckConversion) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

func (r ACHTransferSimulationTransactionSourceACHCheckConversion) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceACHCheckConversion{Amount:%s FileID:%s}", core.FmtP(r.Amount), core.FmtP(r.FileID))
}

type ACHTransferSimulationTransactionSourceACHTransferIntention struct {
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
// ACHTransferSimulationTransactionSourceACHTransferIntention using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceACHTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceACHTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSourceACHTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransactionSourceACHTransferIntention) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceACHTransferIntention) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceACHTransferIntention) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceACHTransferIntention) GetStatementDescriptor() (StatementDescriptor string) {
	if r != nil && r.StatementDescriptor != nil {
		StatementDescriptor = *r.StatementDescriptor
	}
	return
}

// The identifier of the ACH Transfer that led to this Transaction.
func (r *ACHTransferSimulationTransactionSourceACHTransferIntention) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r ACHTransferSimulationTransactionSourceACHTransferIntention) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceACHTransferIntention{Amount:%s AccountNumber:%s RoutingNumber:%s StatementDescriptor:%s TransferID:%s}", core.FmtP(r.Amount), core.FmtP(r.AccountNumber), core.FmtP(r.RoutingNumber), core.FmtP(r.StatementDescriptor), core.FmtP(r.TransferID))
}

type ACHTransferSimulationTransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceACHTransferRejection using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceACHTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceACHTransferRejection into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSourceACHTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the ACH Transfer that led to this Transaction.
func (r *ACHTransferSimulationTransactionSourceACHTransferRejection) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r ACHTransferSimulationTransactionSourceACHTransferRejection) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceACHTransferRejection{TransferID:%s}", core.FmtP(r.TransferID))
}

type ACHTransferSimulationTransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `pjson:"created_at"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode *ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode `pjson:"return_reason_code"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID *string `pjson:"transfer_id"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID *string                `pjson:"transaction_id"`
	jsonFields    map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceACHTransferReturn using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceACHTransferReturn
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceACHTransferReturn) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r *ACHTransferSimulationTransactionSourceACHTransferReturn) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// Why the ACH Transfer was returned.
func (r *ACHTransferSimulationTransactionSourceACHTransferReturn) GetReturnReasonCode() (ReturnReasonCode ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode) {
	if r != nil && r.ReturnReasonCode != nil {
		ReturnReasonCode = *r.ReturnReasonCode
	}
	return
}

// The identifier of the ACH Transfer associated with this return.
func (r *ACHTransferSimulationTransactionSourceACHTransferReturn) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

// The identifier of the Tranasaction associated with this return.
func (r *ACHTransferSimulationTransactionSourceACHTransferReturn) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

func (r ACHTransferSimulationTransactionSourceACHTransferReturn) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceACHTransferReturn{CreatedAt:%s ReturnReasonCode:%s TransferID:%s TransactionID:%s}", core.FmtP(r.CreatedAt), core.FmtP(r.ReturnReasonCode), core.FmtP(r.TransferID), core.FmtP(r.TransactionID))
}

type ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode string

const (
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInsufficientFund                                          ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "insufficient_fund"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeNoAccount                                                 ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "no_account"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAccountClosed                                             ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "account_closed"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure                             ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction              ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver                              ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode   ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized                     ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodePaymentStopped                                            ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "payment_stopped"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeNonTransactionAccount                                     ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "non_transaction_account"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeUncollectedFunds                                          ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "uncollected_funds"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError                              ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAmountFieldError                                          ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "amount_field_error"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer                            ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber                                   ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeFileRecordEditCriteria                                    ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualName                                  ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest                                    ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAddendaError                                              ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "addenda_error"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeLimitedParticipationDfi                                   ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment              ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "incorrectly_coded_outbound_international_payment"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeOther                                                     ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "other"
)

type ACHTransferSimulationTransactionSourceCardDisputeAcceptance struct {
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
// ACHTransferSimulationTransactionSourceCardDisputeAcceptance using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCardDisputeAcceptance) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceCardDisputeAcceptance into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSourceCardDisputeAcceptance) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card Dispute was accepted.
func (r *ACHTransferSimulationTransactionSourceCardDisputeAcceptance) GetAcceptedAt() (AcceptedAt string) {
	if r != nil && r.AcceptedAt != nil {
		AcceptedAt = *r.AcceptedAt
	}
	return
}

// The identifier of the Card Dispute that was accepted.
func (r *ACHTransferSimulationTransactionSourceCardDisputeAcceptance) GetCardDisputeID() (CardDisputeID string) {
	if r != nil && r.CardDisputeID != nil {
		CardDisputeID = *r.CardDisputeID
	}
	return
}

// The identifier of the Transaction that was created to return the disputed funds
// to your account.
func (r *ACHTransferSimulationTransactionSourceCardDisputeAcceptance) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

func (r ACHTransferSimulationTransactionSourceCardDisputeAcceptance) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCardDisputeAcceptance{AcceptedAt:%s CardDisputeID:%s TransactionID:%s}", core.FmtP(r.AcceptedAt), core.FmtP(r.CardDisputeID), core.FmtP(r.TransactionID))
}

type ACHTransferSimulationTransactionSourceCardRefund struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *ACHTransferSimulationTransactionSourceCardRefundCurrency `pjson:"currency"`
	// The identifier for the Transaction this refunds, if any.
	CardSettlementTransactionID *string `pjson:"card_settlement_transaction_id"`
	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type       *ACHTransferSimulationTransactionSourceCardRefundType `pjson:"type"`
	jsonFields map[string]interface{}                                `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceCardRefund using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCardRefund) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceCardRefund into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceCardRefund) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *ACHTransferSimulationTransactionSourceCardRefund) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *ACHTransferSimulationTransactionSourceCardRefund) GetCurrency() (Currency ACHTransferSimulationTransactionSourceCardRefundCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier for the Transaction this refunds, if any.
func (r *ACHTransferSimulationTransactionSourceCardRefund) GetCardSettlementTransactionID() (CardSettlementTransactionID string) {
	if r != nil && r.CardSettlementTransactionID != nil {
		CardSettlementTransactionID = *r.CardSettlementTransactionID
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_refund`.
func (r *ACHTransferSimulationTransactionSourceCardRefund) GetType() (Type ACHTransferSimulationTransactionSourceCardRefundType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r ACHTransferSimulationTransactionSourceCardRefund) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCardRefund{Amount:%s Currency:%s CardSettlementTransactionID:%s Type:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.CardSettlementTransactionID), core.FmtP(r.Type))
}

type ACHTransferSimulationTransactionSourceCardRefundCurrency string

const (
	ACHTransferSimulationTransactionSourceCardRefundCurrencyCad ACHTransferSimulationTransactionSourceCardRefundCurrency = "CAD"
	ACHTransferSimulationTransactionSourceCardRefundCurrencyChf ACHTransferSimulationTransactionSourceCardRefundCurrency = "CHF"
	ACHTransferSimulationTransactionSourceCardRefundCurrencyEur ACHTransferSimulationTransactionSourceCardRefundCurrency = "EUR"
	ACHTransferSimulationTransactionSourceCardRefundCurrencyGbp ACHTransferSimulationTransactionSourceCardRefundCurrency = "GBP"
	ACHTransferSimulationTransactionSourceCardRefundCurrencyJpy ACHTransferSimulationTransactionSourceCardRefundCurrency = "JPY"
	ACHTransferSimulationTransactionSourceCardRefundCurrencyUsd ACHTransferSimulationTransactionSourceCardRefundCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceCardRefundType string

const (
	ACHTransferSimulationTransactionSourceCardRefundTypeCardRefund ACHTransferSimulationTransactionSourceCardRefundType = "card_refund"
)

type ACHTransferSimulationTransactionSourceCardSettlement struct {
	// The amount in the minor unit of the transaction's settlement currency. For
	// dollars, for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's settlement currency.
	Currency *ACHTransferSimulationTransactionSourceCardSettlementCurrency `pjson:"currency"`
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
	Type       *ACHTransferSimulationTransactionSourceCardSettlementType `pjson:"type"`
	jsonFields map[string]interface{}                                    `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceCardSettlement using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCardSettlement) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceCardSettlement into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceCardSettlement) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's settlement currency. For
// dollars, for example, this is cents.
func (r *ACHTransferSimulationTransactionSourceCardSettlement) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's settlement currency.
func (r *ACHTransferSimulationTransactionSourceCardSettlement) GetCurrency() (Currency ACHTransferSimulationTransactionSourceCardSettlementCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The amount in the minor unit of the transaction's presentment currency.
func (r *ACHTransferSimulationTransactionSourceCardSettlement) GetPresentmentAmount() (PresentmentAmount int64) {
	if r != nil && r.PresentmentAmount != nil {
		PresentmentAmount = *r.PresentmentAmount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's presentment currency.
func (r *ACHTransferSimulationTransactionSourceCardSettlement) GetPresentmentCurrency() (PresentmentCurrency string) {
	if r != nil && r.PresentmentCurrency != nil {
		PresentmentCurrency = *r.PresentmentCurrency
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardSettlement) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardSettlement) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardSettlement) GetMerchantName() (MerchantName string) {
	if r != nil && r.MerchantName != nil {
		MerchantName = *r.MerchantName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardSettlement) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardSettlement) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

// The identifier of the Pending Transaction associated with this Transaction.
func (r *ACHTransferSimulationTransactionSourceCardSettlement) GetPendingTransactionID() (PendingTransactionID string) {
	if r != nil && r.PendingTransactionID != nil {
		PendingTransactionID = *r.PendingTransactionID
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_settlement`.
func (r *ACHTransferSimulationTransactionSourceCardSettlement) GetType() (Type ACHTransferSimulationTransactionSourceCardSettlementType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r ACHTransferSimulationTransactionSourceCardSettlement) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCardSettlement{Amount:%s Currency:%s PresentmentAmount:%s PresentmentCurrency:%s MerchantCity:%s MerchantCountry:%s MerchantName:%s MerchantCategoryCode:%s MerchantState:%s PendingTransactionID:%s Type:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.PresentmentAmount), core.FmtP(r.PresentmentCurrency), core.FmtP(r.MerchantCity), core.FmtP(r.MerchantCountry), core.FmtP(r.MerchantName), core.FmtP(r.MerchantCategoryCode), core.FmtP(r.MerchantState), core.FmtP(r.PendingTransactionID), core.FmtP(r.Type))
}

type ACHTransferSimulationTransactionSourceCardSettlementCurrency string

const (
	ACHTransferSimulationTransactionSourceCardSettlementCurrencyCad ACHTransferSimulationTransactionSourceCardSettlementCurrency = "CAD"
	ACHTransferSimulationTransactionSourceCardSettlementCurrencyChf ACHTransferSimulationTransactionSourceCardSettlementCurrency = "CHF"
	ACHTransferSimulationTransactionSourceCardSettlementCurrencyEur ACHTransferSimulationTransactionSourceCardSettlementCurrency = "EUR"
	ACHTransferSimulationTransactionSourceCardSettlementCurrencyGbp ACHTransferSimulationTransactionSourceCardSettlementCurrency = "GBP"
	ACHTransferSimulationTransactionSourceCardSettlementCurrencyJpy ACHTransferSimulationTransactionSourceCardSettlementCurrency = "JPY"
	ACHTransferSimulationTransactionSourceCardSettlementCurrencyUsd ACHTransferSimulationTransactionSourceCardSettlementCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceCardSettlementType string

const (
	ACHTransferSimulationTransactionSourceCardSettlementTypeCardSettlement ACHTransferSimulationTransactionSourceCardSettlementType = "card_settlement"
)

type ACHTransferSimulationTransactionSourceCheckDepositAcceptance struct {
	// The amount to be deposited in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency `pjson:"currency"`
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
// ACHTransferSimulationTransactionSourceCheckDepositAcceptance using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCheckDepositAcceptance) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceCheckDepositAcceptance into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSourceCheckDepositAcceptance) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount to be deposited in the minor unit of the transaction's currency. For
// dollars, for example, this is cents.
func (r *ACHTransferSimulationTransactionSourceCheckDepositAcceptance) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *ACHTransferSimulationTransactionSourceCheckDepositAcceptance) GetCurrency() (Currency ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The account number printed on the check.
func (r *ACHTransferSimulationTransactionSourceCheckDepositAcceptance) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The routing number printed on the check.
func (r *ACHTransferSimulationTransactionSourceCheckDepositAcceptance) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// An additional line of metadata printed on the check. This typically includes the
// check number for business checks.
func (r *ACHTransferSimulationTransactionSourceCheckDepositAcceptance) GetAuxiliaryOnUs() (AuxiliaryOnUs string) {
	if r != nil && r.AuxiliaryOnUs != nil {
		AuxiliaryOnUs = *r.AuxiliaryOnUs
	}
	return
}

// The check serial number, if present, for consumer checks. For business checks,
// the serial number is usually in the `auxiliary_on_us` field.
func (r *ACHTransferSimulationTransactionSourceCheckDepositAcceptance) GetSerialNumber() (SerialNumber string) {
	if r != nil && r.SerialNumber != nil {
		SerialNumber = *r.SerialNumber
	}
	return
}

// The ID of the Check Deposit that was accepted.
func (r *ACHTransferSimulationTransactionSourceCheckDepositAcceptance) GetCheckDepositID() (CheckDepositID string) {
	if r != nil && r.CheckDepositID != nil {
		CheckDepositID = *r.CheckDepositID
	}
	return
}

func (r ACHTransferSimulationTransactionSourceCheckDepositAcceptance) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCheckDepositAcceptance{Amount:%s Currency:%s AccountNumber:%s RoutingNumber:%s AuxiliaryOnUs:%s SerialNumber:%s CheckDepositID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.AccountNumber), core.FmtP(r.RoutingNumber), core.FmtP(r.AuxiliaryOnUs), core.FmtP(r.SerialNumber), core.FmtP(r.CheckDepositID))
}

type ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency string

const (
	ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyCad ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "CAD"
	ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyChf ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "CHF"
	ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyEur ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "EUR"
	ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyGbp ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "GBP"
	ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyJpy ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "JPY"
	ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyUsd ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceCheckDepositReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt *string `pjson:"returned_at"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency `pjson:"currency"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID *string `pjson:"check_deposit_id"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID *string                                                               `pjson:"transaction_id"`
	ReturnReason  *ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason `pjson:"return_reason"`
	jsonFields    map[string]interface{}                                                `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceCheckDepositReturn using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCheckDepositReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceCheckDepositReturn
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceCheckDepositReturn) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransactionSourceCheckDepositReturn) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the check deposit was returned.
func (r *ACHTransferSimulationTransactionSourceCheckDepositReturn) GetReturnedAt() (ReturnedAt string) {
	if r != nil && r.ReturnedAt != nil {
		ReturnedAt = *r.ReturnedAt
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *ACHTransferSimulationTransactionSourceCheckDepositReturn) GetCurrency() (Currency ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Check Deposit that was returned.
func (r *ACHTransferSimulationTransactionSourceCheckDepositReturn) GetCheckDepositID() (CheckDepositID string) {
	if r != nil && r.CheckDepositID != nil {
		CheckDepositID = *r.CheckDepositID
	}
	return
}

// The identifier of the transaction that reversed the original check deposit
// transaction.
func (r *ACHTransferSimulationTransactionSourceCheckDepositReturn) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCheckDepositReturn) GetReturnReason() (ReturnReason ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason) {
	if r != nil && r.ReturnReason != nil {
		ReturnReason = *r.ReturnReason
	}
	return
}

func (r ACHTransferSimulationTransactionSourceCheckDepositReturn) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCheckDepositReturn{Amount:%s ReturnedAt:%s Currency:%s CheckDepositID:%s TransactionID:%s ReturnReason:%s}", core.FmtP(r.Amount), core.FmtP(r.ReturnedAt), core.FmtP(r.Currency), core.FmtP(r.CheckDepositID), core.FmtP(r.TransactionID), core.FmtP(r.ReturnReason))
}

type ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency string

const (
	ACHTransferSimulationTransactionSourceCheckDepositReturnCurrencyCad ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency = "CAD"
	ACHTransferSimulationTransactionSourceCheckDepositReturnCurrencyChf ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency = "CHF"
	ACHTransferSimulationTransactionSourceCheckDepositReturnCurrencyEur ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency = "EUR"
	ACHTransferSimulationTransactionSourceCheckDepositReturnCurrencyGbp ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency = "GBP"
	ACHTransferSimulationTransactionSourceCheckDepositReturnCurrencyJpy ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency = "JPY"
	ACHTransferSimulationTransactionSourceCheckDepositReturnCurrencyUsd ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason string

const (
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonACHConversionNotSupported ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "ach_conversion_not_supported"
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonClosedAccount             ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "closed_account"
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonDuplicateSubmission       ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "duplicate_submission"
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonInsufficientFunds         ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "insufficient_funds"
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonNoAccount                 ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "no_account"
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonNotAuthorized             ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "not_authorized"
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonStaleDated                ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "stale_dated"
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonStopPayment               ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "stop_payment"
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonUnknownReason             ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "unknown_reason"
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonUnmatchedDetails          ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "unmatched_details"
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonUnreadableImage           ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "unreadable_image"
)

type ACHTransferSimulationTransactionSourceCheckTransferIntention struct {
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
	Currency *ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency `pjson:"currency"`
	// The name that will be printed on the check.
	RecipientName *string `pjson:"recipient_name"`
	// The identifier of the Check Transfer with which this is associated.
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceCheckTransferIntention using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceCheckTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The street address of the check's destination.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) GetAddressLine1() (AddressLine1 string) {
	if r != nil && r.AddressLine1 != nil {
		AddressLine1 = *r.AddressLine1
	}
	return
}

// The second line of the address of the check's destination.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) GetAddressLine2() (AddressLine2 string) {
	if r != nil && r.AddressLine2 != nil {
		AddressLine2 = *r.AddressLine2
	}
	return
}

// The city of the check's destination.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) GetAddressCity() (AddressCity string) {
	if r != nil && r.AddressCity != nil {
		AddressCity = *r.AddressCity
	}
	return
}

// The state of the check's destination.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) GetAddressState() (AddressState string) {
	if r != nil && r.AddressState != nil {
		AddressState = *r.AddressState
	}
	return
}

// The postal code of the check's destination.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) GetAddressZip() (AddressZip string) {
	if r != nil && r.AddressZip != nil {
		AddressZip = *r.AddressZip
	}
	return
}

// The transfer amount in USD cents.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) GetCurrency() (Currency ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The name that will be printed on the check.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) GetRecipientName() (RecipientName string) {
	if r != nil && r.RecipientName != nil {
		RecipientName = *r.RecipientName
	}
	return
}

// The identifier of the Check Transfer with which this is associated.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r ACHTransferSimulationTransactionSourceCheckTransferIntention) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCheckTransferIntention{AddressLine1:%s AddressLine2:%s AddressCity:%s AddressState:%s AddressZip:%s Amount:%s Currency:%s RecipientName:%s TransferID:%s}", core.FmtP(r.AddressLine1), core.FmtP(r.AddressLine2), core.FmtP(r.AddressCity), core.FmtP(r.AddressState), core.FmtP(r.AddressZip), core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.RecipientName), core.FmtP(r.TransferID))
}

type ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency string

const (
	ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrencyCad ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "CAD"
	ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrencyChf ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "CHF"
	ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrencyEur ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "EUR"
	ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrencyGbp ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "GBP"
	ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrencyJpy ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "JPY"
	ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrencyUsd ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceCheckTransferReturn struct {
	// The identifier of the returned Check Transfer.
	TransferID *string `pjson:"transfer_id"`
	// If available, a document with additional information about the return.
	FileID     *string                `pjson:"file_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceCheckTransferReturn using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCheckTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceCheckTransferReturn
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceCheckTransferReturn) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the returned Check Transfer.
func (r *ACHTransferSimulationTransactionSourceCheckTransferReturn) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

// If available, a document with additional information about the return.
func (r *ACHTransferSimulationTransactionSourceCheckTransferReturn) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

func (r ACHTransferSimulationTransactionSourceCheckTransferReturn) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCheckTransferReturn{TransferID:%s FileID:%s}", core.FmtP(r.TransferID), core.FmtP(r.FileID))
}

type ACHTransferSimulationTransactionSourceCheckTransferRejection struct {
	// The identifier of the Check Transfer that led to this Transaction.
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceCheckTransferRejection using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCheckTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceCheckTransferRejection into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSourceCheckTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the Check Transfer that led to this Transaction.
func (r *ACHTransferSimulationTransactionSourceCheckTransferRejection) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r ACHTransferSimulationTransactionSourceCheckTransferRejection) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCheckTransferRejection{TransferID:%s}", core.FmtP(r.TransferID))
}

type ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferID *string `pjson:"transfer_id"`
	// The transaction ID of the corresponding credit transaction.
	TransactionID *string `pjson:"transaction_id"`
	// The time the stop-payment was requested.
	RequestedAt *string `pjson:"requested_at"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type       *ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType `pjson:"type"`
	jsonFields map[string]interface{}                                                     `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The ID of the check transfer that was stopped.
func (r *ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

// The transaction ID of the corresponding credit transaction.
func (r *ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

// The time the stop-payment was requested.
func (r *ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) GetRequestedAt() (RequestedAt string) {
	if r != nil && r.RequestedAt != nil {
		RequestedAt = *r.RequestedAt
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_stop_payment_request`.
func (r *ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) GetType() (Type ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest{TransferID:%s TransactionID:%s RequestedAt:%s Type:%s}", core.FmtP(r.TransferID), core.FmtP(r.TransactionID), core.FmtP(r.RequestedAt), core.FmtP(r.Type))
}

type ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType string

const (
	ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

type ACHTransferSimulationTransactionSourceDisputeResolution struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *ACHTransferSimulationTransactionSourceDisputeResolutionCurrency `pjson:"currency"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionID *string                `pjson:"disputed_transaction_id"`
	jsonFields            map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceDisputeResolution using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceDisputeResolution) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceDisputeResolution
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceDisputeResolution) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransactionSourceDisputeResolution) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *ACHTransferSimulationTransactionSourceDisputeResolution) GetCurrency() (Currency ACHTransferSimulationTransactionSourceDisputeResolutionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Transaction that was disputed.
func (r *ACHTransferSimulationTransactionSourceDisputeResolution) GetDisputedTransactionID() (DisputedTransactionID string) {
	if r != nil && r.DisputedTransactionID != nil {
		DisputedTransactionID = *r.DisputedTransactionID
	}
	return
}

func (r ACHTransferSimulationTransactionSourceDisputeResolution) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceDisputeResolution{Amount:%s Currency:%s DisputedTransactionID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.DisputedTransactionID))
}

type ACHTransferSimulationTransactionSourceDisputeResolutionCurrency string

const (
	ACHTransferSimulationTransactionSourceDisputeResolutionCurrencyCad ACHTransferSimulationTransactionSourceDisputeResolutionCurrency = "CAD"
	ACHTransferSimulationTransactionSourceDisputeResolutionCurrencyChf ACHTransferSimulationTransactionSourceDisputeResolutionCurrency = "CHF"
	ACHTransferSimulationTransactionSourceDisputeResolutionCurrencyEur ACHTransferSimulationTransactionSourceDisputeResolutionCurrency = "EUR"
	ACHTransferSimulationTransactionSourceDisputeResolutionCurrencyGbp ACHTransferSimulationTransactionSourceDisputeResolutionCurrency = "GBP"
	ACHTransferSimulationTransactionSourceDisputeResolutionCurrencyJpy ACHTransferSimulationTransactionSourceDisputeResolutionCurrency = "JPY"
	ACHTransferSimulationTransactionSourceDisputeResolutionCurrencyUsd ACHTransferSimulationTransactionSourceDisputeResolutionCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceEmpyrealCashDeposit struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount      *int64                 `pjson:"amount"`
	BagID       *string                `pjson:"bag_id"`
	DepositDate *string                `pjson:"deposit_date"`
	jsonFields  map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceEmpyrealCashDeposit using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceEmpyrealCashDeposit) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceEmpyrealCashDeposit
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceEmpyrealCashDeposit) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransactionSourceEmpyrealCashDeposit) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceEmpyrealCashDeposit) GetBagID() (BagID string) {
	if r != nil && r.BagID != nil {
		BagID = *r.BagID
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceEmpyrealCashDeposit) GetDepositDate() (DepositDate string) {
	if r != nil && r.DepositDate != nil {
		DepositDate = *r.DepositDate
	}
	return
}

func (r ACHTransferSimulationTransactionSourceEmpyrealCashDeposit) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceEmpyrealCashDeposit{Amount:%s BagID:%s DepositDate:%s}", core.FmtP(r.Amount), core.FmtP(r.BagID), core.FmtP(r.DepositDate))
}

type ACHTransferSimulationTransactionSourceInboundACHTransfer struct {
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
// ACHTransferSimulationTransactionSourceInboundACHTransfer using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceInboundACHTransfer
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the destination account currency. For dollars,
// for example, this is cents.
func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) GetOriginatorCompanyName() (OriginatorCompanyName string) {
	if r != nil && r.OriginatorCompanyName != nil {
		OriginatorCompanyName = *r.OriginatorCompanyName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) GetOriginatorCompanyDescriptiveDate() (OriginatorCompanyDescriptiveDate string) {
	if r != nil && r.OriginatorCompanyDescriptiveDate != nil {
		OriginatorCompanyDescriptiveDate = *r.OriginatorCompanyDescriptiveDate
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) GetOriginatorCompanyDiscretionaryData() (OriginatorCompanyDiscretionaryData string) {
	if r != nil && r.OriginatorCompanyDiscretionaryData != nil {
		OriginatorCompanyDiscretionaryData = *r.OriginatorCompanyDiscretionaryData
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r != nil && r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) GetOriginatorCompanyID() (OriginatorCompanyID string) {
	if r != nil && r.OriginatorCompanyID != nil {
		OriginatorCompanyID = *r.OriginatorCompanyID
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) GetReceiverIDNumber() (ReceiverIDNumber string) {
	if r != nil && r.ReceiverIDNumber != nil {
		ReceiverIDNumber = *r.ReceiverIDNumber
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) GetReceiverName() (ReceiverName string) {
	if r != nil && r.ReceiverName != nil {
		ReceiverName = *r.ReceiverName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

func (r ACHTransferSimulationTransactionSourceInboundACHTransfer) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceInboundACHTransfer{Amount:%s OriginatorCompanyName:%s OriginatorCompanyDescriptiveDate:%s OriginatorCompanyDiscretionaryData:%s OriginatorCompanyEntryDescription:%s OriginatorCompanyID:%s ReceiverIDNumber:%s ReceiverName:%s TraceNumber:%s}", core.FmtP(r.Amount), core.FmtP(r.OriginatorCompanyName), core.FmtP(r.OriginatorCompanyDescriptiveDate), core.FmtP(r.OriginatorCompanyDiscretionaryData), core.FmtP(r.OriginatorCompanyEntryDescription), core.FmtP(r.OriginatorCompanyID), core.FmtP(r.ReceiverIDNumber), core.FmtP(r.ReceiverName), core.FmtP(r.TraceNumber))
}

type ACHTransferSimulationTransactionSourceInboundCheck struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency              *ACHTransferSimulationTransactionSourceInboundCheckCurrency `pjson:"currency"`
	CheckNumber           *string                                                     `pjson:"check_number"`
	CheckFrontImageFileID *string                                                     `pjson:"check_front_image_file_id"`
	CheckRearImageFileID  *string                                                     `pjson:"check_rear_image_file_id"`
	jsonFields            map[string]interface{}                                      `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceInboundCheck using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceInboundCheck) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceInboundCheck into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceInboundCheck) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the destination account currency. For dollars,
// for example, this is cents.
func (r *ACHTransferSimulationTransactionSourceInboundCheck) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *ACHTransferSimulationTransactionSourceInboundCheck) GetCurrency() (Currency ACHTransferSimulationTransactionSourceInboundCheckCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundCheck) GetCheckNumber() (CheckNumber string) {
	if r != nil && r.CheckNumber != nil {
		CheckNumber = *r.CheckNumber
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundCheck) GetCheckFrontImageFileID() (CheckFrontImageFileID string) {
	if r != nil && r.CheckFrontImageFileID != nil {
		CheckFrontImageFileID = *r.CheckFrontImageFileID
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundCheck) GetCheckRearImageFileID() (CheckRearImageFileID string) {
	if r != nil && r.CheckRearImageFileID != nil {
		CheckRearImageFileID = *r.CheckRearImageFileID
	}
	return
}

func (r ACHTransferSimulationTransactionSourceInboundCheck) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceInboundCheck{Amount:%s Currency:%s CheckNumber:%s CheckFrontImageFileID:%s CheckRearImageFileID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.CheckNumber), core.FmtP(r.CheckFrontImageFileID), core.FmtP(r.CheckRearImageFileID))
}

type ACHTransferSimulationTransactionSourceInboundCheckCurrency string

const (
	ACHTransferSimulationTransactionSourceInboundCheckCurrencyCad ACHTransferSimulationTransactionSourceInboundCheckCurrency = "CAD"
	ACHTransferSimulationTransactionSourceInboundCheckCurrencyChf ACHTransferSimulationTransactionSourceInboundCheckCurrency = "CHF"
	ACHTransferSimulationTransactionSourceInboundCheckCurrencyEur ACHTransferSimulationTransactionSourceInboundCheckCurrency = "EUR"
	ACHTransferSimulationTransactionSourceInboundCheckCurrencyGbp ACHTransferSimulationTransactionSourceInboundCheckCurrency = "GBP"
	ACHTransferSimulationTransactionSourceInboundCheckCurrencyJpy ACHTransferSimulationTransactionSourceInboundCheckCurrency = "JPY"
	ACHTransferSimulationTransactionSourceInboundCheckCurrencyUsd ACHTransferSimulationTransactionSourceInboundCheckCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer struct {
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
// ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the destination account currency. For dollars,
// for example, this is cents.
func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetForeignExchangeIndicator() (ForeignExchangeIndicator string) {
	if r != nil && r.ForeignExchangeIndicator != nil {
		ForeignExchangeIndicator = *r.ForeignExchangeIndicator
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetForeignExchangeReferenceIndicator() (ForeignExchangeReferenceIndicator string) {
	if r != nil && r.ForeignExchangeReferenceIndicator != nil {
		ForeignExchangeReferenceIndicator = *r.ForeignExchangeReferenceIndicator
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetForeignExchangeReference() (ForeignExchangeReference string) {
	if r != nil && r.ForeignExchangeReference != nil {
		ForeignExchangeReference = *r.ForeignExchangeReference
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetDestinationCountryCode() (DestinationCountryCode string) {
	if r != nil && r.DestinationCountryCode != nil {
		DestinationCountryCode = *r.DestinationCountryCode
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetDestinationCurrencyCode() (DestinationCurrencyCode string) {
	if r != nil && r.DestinationCurrencyCode != nil {
		DestinationCurrencyCode = *r.DestinationCurrencyCode
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetForeignPaymentAmount() (ForeignPaymentAmount int64) {
	if r != nil && r.ForeignPaymentAmount != nil {
		ForeignPaymentAmount = *r.ForeignPaymentAmount
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetForeignTraceNumber() (ForeignTraceNumber string) {
	if r != nil && r.ForeignTraceNumber != nil {
		ForeignTraceNumber = *r.ForeignTraceNumber
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetInternationalTransactionTypeCode() (InternationalTransactionTypeCode string) {
	if r != nil && r.InternationalTransactionTypeCode != nil {
		InternationalTransactionTypeCode = *r.InternationalTransactionTypeCode
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatingCurrencyCode() (OriginatingCurrencyCode string) {
	if r != nil && r.OriginatingCurrencyCode != nil {
		OriginatingCurrencyCode = *r.OriginatingCurrencyCode
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionName() (OriginatingDepositoryFinancialInstitutionName string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionName != nil {
		OriginatingDepositoryFinancialInstitutionName = *r.OriginatingDepositoryFinancialInstitutionName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionIDQualifier() (OriginatingDepositoryFinancialInstitutionIDQualifier string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionIDQualifier != nil {
		OriginatingDepositoryFinancialInstitutionIDQualifier = *r.OriginatingDepositoryFinancialInstitutionIDQualifier
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionID() (OriginatingDepositoryFinancialInstitutionID string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionID != nil {
		OriginatingDepositoryFinancialInstitutionID = *r.OriginatingDepositoryFinancialInstitutionID
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionBranchCountry() (OriginatingDepositoryFinancialInstitutionBranchCountry string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionBranchCountry != nil {
		OriginatingDepositoryFinancialInstitutionBranchCountry = *r.OriginatingDepositoryFinancialInstitutionBranchCountry
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorCity() (OriginatorCity string) {
	if r != nil && r.OriginatorCity != nil {
		OriginatorCity = *r.OriginatorCity
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r != nil && r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorCountry() (OriginatorCountry string) {
	if r != nil && r.OriginatorCountry != nil {
		OriginatorCountry = *r.OriginatorCountry
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorIdentification() (OriginatorIdentification string) {
	if r != nil && r.OriginatorIdentification != nil {
		OriginatorIdentification = *r.OriginatorIdentification
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorPostalCode() (OriginatorPostalCode string) {
	if r != nil && r.OriginatorPostalCode != nil {
		OriginatorPostalCode = *r.OriginatorPostalCode
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorStreetAddress() (OriginatorStreetAddress string) {
	if r != nil && r.OriginatorStreetAddress != nil {
		OriginatorStreetAddress = *r.OriginatorStreetAddress
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorStateOrProvince() (OriginatorStateOrProvince string) {
	if r != nil && r.OriginatorStateOrProvince != nil {
		OriginatorStateOrProvince = *r.OriginatorStateOrProvince
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetPaymentRelatedInformation() (PaymentRelatedInformation string) {
	if r != nil && r.PaymentRelatedInformation != nil {
		PaymentRelatedInformation = *r.PaymentRelatedInformation
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetPaymentRelatedInformation2() (PaymentRelatedInformation2 string) {
	if r != nil && r.PaymentRelatedInformation2 != nil {
		PaymentRelatedInformation2 = *r.PaymentRelatedInformation2
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverIdentificationNumber() (ReceiverIdentificationNumber string) {
	if r != nil && r.ReceiverIdentificationNumber != nil {
		ReceiverIdentificationNumber = *r.ReceiverIdentificationNumber
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverStreetAddress() (ReceiverStreetAddress string) {
	if r != nil && r.ReceiverStreetAddress != nil {
		ReceiverStreetAddress = *r.ReceiverStreetAddress
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverCity() (ReceiverCity string) {
	if r != nil && r.ReceiverCity != nil {
		ReceiverCity = *r.ReceiverCity
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverStateOrProvince() (ReceiverStateOrProvince string) {
	if r != nil && r.ReceiverStateOrProvince != nil {
		ReceiverStateOrProvince = *r.ReceiverStateOrProvince
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverCountry() (ReceiverCountry string) {
	if r != nil && r.ReceiverCountry != nil {
		ReceiverCountry = *r.ReceiverCountry
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverPostalCode() (ReceiverPostalCode string) {
	if r != nil && r.ReceiverPostalCode != nil {
		ReceiverPostalCode = *r.ReceiverPostalCode
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceivingCompanyOrIndividualName() (ReceivingCompanyOrIndividualName string) {
	if r != nil && r.ReceivingCompanyOrIndividualName != nil {
		ReceivingCompanyOrIndividualName = *r.ReceivingCompanyOrIndividualName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionName() (ReceivingDepositoryFinancialInstitutionName string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionName != nil {
		ReceivingDepositoryFinancialInstitutionName = *r.ReceivingDepositoryFinancialInstitutionName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionIDQualifier() (ReceivingDepositoryFinancialInstitutionIDQualifier string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionIDQualifier != nil {
		ReceivingDepositoryFinancialInstitutionIDQualifier = *r.ReceivingDepositoryFinancialInstitutionIDQualifier
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionID() (ReceivingDepositoryFinancialInstitutionID string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionID != nil {
		ReceivingDepositoryFinancialInstitutionID = *r.ReceivingDepositoryFinancialInstitutionID
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionCountry() (ReceivingDepositoryFinancialInstitutionCountry string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionCountry != nil {
		ReceivingDepositoryFinancialInstitutionCountry = *r.ReceivingDepositoryFinancialInstitutionCountry
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

func (r ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer{Amount:%s ForeignExchangeIndicator:%s ForeignExchangeReferenceIndicator:%s ForeignExchangeReference:%s DestinationCountryCode:%s DestinationCurrencyCode:%s ForeignPaymentAmount:%s ForeignTraceNumber:%s InternationalTransactionTypeCode:%s OriginatingCurrencyCode:%s OriginatingDepositoryFinancialInstitutionName:%s OriginatingDepositoryFinancialInstitutionIDQualifier:%s OriginatingDepositoryFinancialInstitutionID:%s OriginatingDepositoryFinancialInstitutionBranchCountry:%s OriginatorCity:%s OriginatorCompanyEntryDescription:%s OriginatorCountry:%s OriginatorIdentification:%s OriginatorName:%s OriginatorPostalCode:%s OriginatorStreetAddress:%s OriginatorStateOrProvince:%s PaymentRelatedInformation:%s PaymentRelatedInformation2:%s ReceiverIdentificationNumber:%s ReceiverStreetAddress:%s ReceiverCity:%s ReceiverStateOrProvince:%s ReceiverCountry:%s ReceiverPostalCode:%s ReceivingCompanyOrIndividualName:%s ReceivingDepositoryFinancialInstitutionName:%s ReceivingDepositoryFinancialInstitutionIDQualifier:%s ReceivingDepositoryFinancialInstitutionID:%s ReceivingDepositoryFinancialInstitutionCountry:%s TraceNumber:%s}", core.FmtP(r.Amount), core.FmtP(r.ForeignExchangeIndicator), core.FmtP(r.ForeignExchangeReferenceIndicator), core.FmtP(r.ForeignExchangeReference), core.FmtP(r.DestinationCountryCode), core.FmtP(r.DestinationCurrencyCode), core.FmtP(r.ForeignPaymentAmount), core.FmtP(r.ForeignTraceNumber), core.FmtP(r.InternationalTransactionTypeCode), core.FmtP(r.OriginatingCurrencyCode), core.FmtP(r.OriginatingDepositoryFinancialInstitutionName), core.FmtP(r.OriginatingDepositoryFinancialInstitutionIDQualifier), core.FmtP(r.OriginatingDepositoryFinancialInstitutionID), core.FmtP(r.OriginatingDepositoryFinancialInstitutionBranchCountry), core.FmtP(r.OriginatorCity), core.FmtP(r.OriginatorCompanyEntryDescription), core.FmtP(r.OriginatorCountry), core.FmtP(r.OriginatorIdentification), core.FmtP(r.OriginatorName), core.FmtP(r.OriginatorPostalCode), core.FmtP(r.OriginatorStreetAddress), core.FmtP(r.OriginatorStateOrProvince), core.FmtP(r.PaymentRelatedInformation), core.FmtP(r.PaymentRelatedInformation2), core.FmtP(r.ReceiverIdentificationNumber), core.FmtP(r.ReceiverStreetAddress), core.FmtP(r.ReceiverCity), core.FmtP(r.ReceiverStateOrProvince), core.FmtP(r.ReceiverCountry), core.FmtP(r.ReceiverPostalCode), core.FmtP(r.ReceivingCompanyOrIndividualName), core.FmtP(r.ReceivingDepositoryFinancialInstitutionName), core.FmtP(r.ReceivingDepositoryFinancialInstitutionIDQualifier), core.FmtP(r.ReceivingDepositoryFinancialInstitutionID), core.FmtP(r.ReceivingDepositoryFinancialInstitutionCountry), core.FmtP(r.TraceNumber))
}

type ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real Time Payments transfer.
	Currency *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency `pjson:"currency"`
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
// ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transfer's currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
// currency. This will always be "USD" for a Real Time Payments transfer.
func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetCurrency() (Currency ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The name the sender of the transfer specified as the recipient of the transfer.
func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetCreditorName() (CreditorName string) {
	if r != nil && r.CreditorName != nil {
		CreditorName = *r.CreditorName
	}
	return
}

// The name provided by the sender of the transfer.
func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorName() (DebtorName string) {
	if r != nil && r.DebtorName != nil {
		DebtorName = *r.DebtorName
	}
	return
}

// The account number of the account that sent the transfer.
func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorAccountNumber() (DebtorAccountNumber string) {
	if r != nil && r.DebtorAccountNumber != nil {
		DebtorAccountNumber = *r.DebtorAccountNumber
	}
	return
}

// The routing number of the account that sent the transfer.
func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorRoutingNumber() (DebtorRoutingNumber string) {
	if r != nil && r.DebtorRoutingNumber != nil {
		DebtorRoutingNumber = *r.DebtorRoutingNumber
	}
	return
}

// The Real Time Payments network identification of the transfer
func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetTransactionIdentification() (TransactionIdentification string) {
	if r != nil && r.TransactionIdentification != nil {
		TransactionIdentification = *r.TransactionIdentification
	}
	return
}

// Additional information included with the transfer.
func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetRemittanceInformation() (RemittanceInformation string) {
	if r != nil && r.RemittanceInformation != nil {
		RemittanceInformation = *r.RemittanceInformation
	}
	return
}

func (r ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation{Amount:%s Currency:%s CreditorName:%s DebtorName:%s DebtorAccountNumber:%s DebtorRoutingNumber:%s TransactionIdentification:%s RemittanceInformation:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.CreditorName), core.FmtP(r.DebtorName), core.FmtP(r.DebtorAccountNumber), core.FmtP(r.DebtorRoutingNumber), core.FmtP(r.TransactionIdentification), core.FmtP(r.RemittanceInformation))
}

type ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency string

const (
	ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyCad ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CAD"
	ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyChf ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CHF"
	ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyEur ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "EUR"
	ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyGbp ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "GBP"
	ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyJpy ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "JPY"
	ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyUsd ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal struct {
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
// ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount that was reversed.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description on the reversal message from Fedwire.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Fedwire cycle date for the wire reversal.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetInputCycleDate() (InputCycleDate string) {
	if r != nil && r.InputCycleDate != nil {
		InputCycleDate = *r.InputCycleDate
	}
	return
}

// The Fedwire sequence number.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetInputSequenceNumber() (InputSequenceNumber string) {
	if r != nil && r.InputSequenceNumber != nil {
		InputSequenceNumber = *r.InputSequenceNumber
	}
	return
}

// The Fedwire input source identifier.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetInputSource() (InputSource string) {
	if r != nil && r.InputSource != nil {
		InputSource = *r.InputSource
	}
	return
}

// The Fedwire transaction identifier.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

// The Fedwire transaction identifier for the wire transfer that was reversed.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputMessageAccountabilityData() (PreviousMessageInputMessageAccountabilityData string) {
	if r != nil && r.PreviousMessageInputMessageAccountabilityData != nil {
		PreviousMessageInputMessageAccountabilityData = *r.PreviousMessageInputMessageAccountabilityData
	}
	return
}

// The Fedwire cycle date for the wire transfer that was reversed.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputCycleDate() (PreviousMessageInputCycleDate string) {
	if r != nil && r.PreviousMessageInputCycleDate != nil {
		PreviousMessageInputCycleDate = *r.PreviousMessageInputCycleDate
	}
	return
}

// The Fedwire sequence number for the wire transfer that was reversed.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputSequenceNumber() (PreviousMessageInputSequenceNumber string) {
	if r != nil && r.PreviousMessageInputSequenceNumber != nil {
		PreviousMessageInputSequenceNumber = *r.PreviousMessageInputSequenceNumber
	}
	return
}

// The Fedwire input source identifier for the wire transfer that was reversed.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputSource() (PreviousMessageInputSource string) {
	if r != nil && r.PreviousMessageInputSource != nil {
		PreviousMessageInputSource = *r.PreviousMessageInputSource
	}
	return
}

func (r ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal{Amount:%s Description:%s InputCycleDate:%s InputSequenceNumber:%s InputSource:%s InputMessageAccountabilityData:%s PreviousMessageInputMessageAccountabilityData:%s PreviousMessageInputCycleDate:%s PreviousMessageInputSequenceNumber:%s PreviousMessageInputSource:%s}", core.FmtP(r.Amount), core.FmtP(r.Description), core.FmtP(r.InputCycleDate), core.FmtP(r.InputSequenceNumber), core.FmtP(r.InputSource), core.FmtP(r.InputMessageAccountabilityData), core.FmtP(r.PreviousMessageInputMessageAccountabilityData), core.FmtP(r.PreviousMessageInputCycleDate), core.FmtP(r.PreviousMessageInputSequenceNumber), core.FmtP(r.PreviousMessageInputSource))
}

type ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment struct {
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
// ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r != nil && r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r != nil && r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r != nil && r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryName() (BeneficiaryName string) {
	if r != nil && r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryReference() (BeneficiaryReference string) {
	if r != nil && r.BeneficiaryReference != nil {
		BeneficiaryReference = *r.BeneficiaryReference
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine1() (OriginatorAddressLine1 string) {
	if r != nil && r.OriginatorAddressLine1 != nil {
		OriginatorAddressLine1 = *r.OriginatorAddressLine1
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine2() (OriginatorAddressLine2 string) {
	if r != nil && r.OriginatorAddressLine2 != nil {
		OriginatorAddressLine2 = *r.OriginatorAddressLine2
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine3() (OriginatorAddressLine3 string) {
	if r != nil && r.OriginatorAddressLine3 != nil {
		OriginatorAddressLine3 = *r.OriginatorAddressLine3
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetOriginatorToBeneficiaryInformation() (OriginatorToBeneficiaryInformation string) {
	if r != nil && r.OriginatorToBeneficiaryInformation != nil {
		OriginatorToBeneficiaryInformation = *r.OriginatorToBeneficiaryInformation
	}
	return
}

func (r ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment{Amount:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s BeneficiaryName:%s BeneficiaryReference:%s Description:%s InputMessageAccountabilityData:%s OriginatorAddressLine1:%s OriginatorAddressLine2:%s OriginatorAddressLine3:%s OriginatorName:%s OriginatorToBeneficiaryInformation:%s}", core.FmtP(r.Amount), core.FmtP(r.BeneficiaryAddressLine1), core.FmtP(r.BeneficiaryAddressLine2), core.FmtP(r.BeneficiaryAddressLine3), core.FmtP(r.BeneficiaryName), core.FmtP(r.BeneficiaryReference), core.FmtP(r.Description), core.FmtP(r.InputMessageAccountabilityData), core.FmtP(r.OriginatorAddressLine1), core.FmtP(r.OriginatorAddressLine2), core.FmtP(r.OriginatorAddressLine3), core.FmtP(r.OriginatorName), core.FmtP(r.OriginatorToBeneficiaryInformation))
}

type ACHTransferSimulationTransactionSourceInboundWireReversal struct {
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
// ACHTransferSimulationTransactionSourceInboundWireReversal using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceInboundWireReversal
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount that was reversed.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description on the reversal message from Fedwire.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Fedwire cycle date for the wire reversal.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetInputCycleDate() (InputCycleDate string) {
	if r != nil && r.InputCycleDate != nil {
		InputCycleDate = *r.InputCycleDate
	}
	return
}

// The Fedwire sequence number.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetInputSequenceNumber() (InputSequenceNumber string) {
	if r != nil && r.InputSequenceNumber != nil {
		InputSequenceNumber = *r.InputSequenceNumber
	}
	return
}

// The Fedwire input source identifier.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetInputSource() (InputSource string) {
	if r != nil && r.InputSource != nil {
		InputSource = *r.InputSource
	}
	return
}

// The Fedwire transaction identifier.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

// The Fedwire transaction identifier for the wire transfer that was reversed.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetPreviousMessageInputMessageAccountabilityData() (PreviousMessageInputMessageAccountabilityData string) {
	if r != nil && r.PreviousMessageInputMessageAccountabilityData != nil {
		PreviousMessageInputMessageAccountabilityData = *r.PreviousMessageInputMessageAccountabilityData
	}
	return
}

// The Fedwire cycle date for the wire transfer that was reversed.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetPreviousMessageInputCycleDate() (PreviousMessageInputCycleDate string) {
	if r != nil && r.PreviousMessageInputCycleDate != nil {
		PreviousMessageInputCycleDate = *r.PreviousMessageInputCycleDate
	}
	return
}

// The Fedwire sequence number for the wire transfer that was reversed.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetPreviousMessageInputSequenceNumber() (PreviousMessageInputSequenceNumber string) {
	if r != nil && r.PreviousMessageInputSequenceNumber != nil {
		PreviousMessageInputSequenceNumber = *r.PreviousMessageInputSequenceNumber
	}
	return
}

// The Fedwire input source identifier for the wire transfer that was reversed.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetPreviousMessageInputSource() (PreviousMessageInputSource string) {
	if r != nil && r.PreviousMessageInputSource != nil {
		PreviousMessageInputSource = *r.PreviousMessageInputSource
	}
	return
}

// Information included in the wire reversal for the receiving financial
// institution.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetReceiverFinancialInstitutionInformation() (ReceiverFinancialInstitutionInformation string) {
	if r != nil && r.ReceiverFinancialInstitutionInformation != nil {
		ReceiverFinancialInstitutionInformation = *r.ReceiverFinancialInstitutionInformation
	}
	return
}

// Additional financial institution information included in the wire reversal.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetFinancialInstitutionToFinancialInstitutionInformation() (FinancialInstitutionToFinancialInstitutionInformation string) {
	if r != nil && r.FinancialInstitutionToFinancialInstitutionInformation != nil {
		FinancialInstitutionToFinancialInstitutionInformation = *r.FinancialInstitutionToFinancialInstitutionInformation
	}
	return
}

func (r ACHTransferSimulationTransactionSourceInboundWireReversal) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceInboundWireReversal{Amount:%s Description:%s InputCycleDate:%s InputSequenceNumber:%s InputSource:%s InputMessageAccountabilityData:%s PreviousMessageInputMessageAccountabilityData:%s PreviousMessageInputCycleDate:%s PreviousMessageInputSequenceNumber:%s PreviousMessageInputSource:%s ReceiverFinancialInstitutionInformation:%s FinancialInstitutionToFinancialInstitutionInformation:%s}", core.FmtP(r.Amount), core.FmtP(r.Description), core.FmtP(r.InputCycleDate), core.FmtP(r.InputSequenceNumber), core.FmtP(r.InputSource), core.FmtP(r.InputMessageAccountabilityData), core.FmtP(r.PreviousMessageInputMessageAccountabilityData), core.FmtP(r.PreviousMessageInputCycleDate), core.FmtP(r.PreviousMessageInputSequenceNumber), core.FmtP(r.PreviousMessageInputSource), core.FmtP(r.ReceiverFinancialInstitutionInformation), core.FmtP(r.FinancialInstitutionToFinancialInstitutionInformation))
}

type ACHTransferSimulationTransactionSourceInboundWireTransfer struct {
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
// ACHTransferSimulationTransactionSourceInboundWireTransfer using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceInboundWireTransfer
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r != nil && r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r != nil && r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r != nil && r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetBeneficiaryName() (BeneficiaryName string) {
	if r != nil && r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetBeneficiaryReference() (BeneficiaryReference string) {
	if r != nil && r.BeneficiaryReference != nil {
		BeneficiaryReference = *r.BeneficiaryReference
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorAddressLine1() (OriginatorAddressLine1 string) {
	if r != nil && r.OriginatorAddressLine1 != nil {
		OriginatorAddressLine1 = *r.OriginatorAddressLine1
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorAddressLine2() (OriginatorAddressLine2 string) {
	if r != nil && r.OriginatorAddressLine2 != nil {
		OriginatorAddressLine2 = *r.OriginatorAddressLine2
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorAddressLine3() (OriginatorAddressLine3 string) {
	if r != nil && r.OriginatorAddressLine3 != nil {
		OriginatorAddressLine3 = *r.OriginatorAddressLine3
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorToBeneficiaryInformationLine1() (OriginatorToBeneficiaryInformationLine1 string) {
	if r != nil && r.OriginatorToBeneficiaryInformationLine1 != nil {
		OriginatorToBeneficiaryInformationLine1 = *r.OriginatorToBeneficiaryInformationLine1
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorToBeneficiaryInformationLine2() (OriginatorToBeneficiaryInformationLine2 string) {
	if r != nil && r.OriginatorToBeneficiaryInformationLine2 != nil {
		OriginatorToBeneficiaryInformationLine2 = *r.OriginatorToBeneficiaryInformationLine2
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorToBeneficiaryInformationLine3() (OriginatorToBeneficiaryInformationLine3 string) {
	if r != nil && r.OriginatorToBeneficiaryInformationLine3 != nil {
		OriginatorToBeneficiaryInformationLine3 = *r.OriginatorToBeneficiaryInformationLine3
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorToBeneficiaryInformationLine4() (OriginatorToBeneficiaryInformationLine4 string) {
	if r != nil && r.OriginatorToBeneficiaryInformationLine4 != nil {
		OriginatorToBeneficiaryInformationLine4 = *r.OriginatorToBeneficiaryInformationLine4
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorToBeneficiaryInformation() (OriginatorToBeneficiaryInformation string) {
	if r != nil && r.OriginatorToBeneficiaryInformation != nil {
		OriginatorToBeneficiaryInformation = *r.OriginatorToBeneficiaryInformation
	}
	return
}

func (r ACHTransferSimulationTransactionSourceInboundWireTransfer) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceInboundWireTransfer{Amount:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s BeneficiaryName:%s BeneficiaryReference:%s Description:%s InputMessageAccountabilityData:%s OriginatorAddressLine1:%s OriginatorAddressLine2:%s OriginatorAddressLine3:%s OriginatorName:%s OriginatorToBeneficiaryInformationLine1:%s OriginatorToBeneficiaryInformationLine2:%s OriginatorToBeneficiaryInformationLine3:%s OriginatorToBeneficiaryInformationLine4:%s OriginatorToBeneficiaryInformation:%s}", core.FmtP(r.Amount), core.FmtP(r.BeneficiaryAddressLine1), core.FmtP(r.BeneficiaryAddressLine2), core.FmtP(r.BeneficiaryAddressLine3), core.FmtP(r.BeneficiaryName), core.FmtP(r.BeneficiaryReference), core.FmtP(r.Description), core.FmtP(r.InputMessageAccountabilityData), core.FmtP(r.OriginatorAddressLine1), core.FmtP(r.OriginatorAddressLine2), core.FmtP(r.OriginatorAddressLine3), core.FmtP(r.OriginatorName), core.FmtP(r.OriginatorToBeneficiaryInformationLine1), core.FmtP(r.OriginatorToBeneficiaryInformationLine2), core.FmtP(r.OriginatorToBeneficiaryInformationLine3), core.FmtP(r.OriginatorToBeneficiaryInformationLine4), core.FmtP(r.OriginatorToBeneficiaryInformation))
}

type ACHTransferSimulationTransactionSourceInterestPayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency *ACHTransferSimulationTransactionSourceInterestPaymentCurrency `pjson:"currency"`
	// The start of the period for which this transaction paid interest.
	PeriodStart *string `pjson:"period_start"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd *string `pjson:"period_end"`
	// The account on which the interest was accrued.
	AccruedOnAccountID *string                `pjson:"accrued_on_account_id"`
	jsonFields         map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceInterestPayment using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceInterestPayment) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceInterestPayment
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceInterestPayment) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransactionSourceInterestPayment) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
func (r *ACHTransferSimulationTransactionSourceInterestPayment) GetCurrency() (Currency ACHTransferSimulationTransactionSourceInterestPaymentCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The start of the period for which this transaction paid interest.
func (r *ACHTransferSimulationTransactionSourceInterestPayment) GetPeriodStart() (PeriodStart string) {
	if r != nil && r.PeriodStart != nil {
		PeriodStart = *r.PeriodStart
	}
	return
}

// The end of the period for which this transaction paid interest.
func (r *ACHTransferSimulationTransactionSourceInterestPayment) GetPeriodEnd() (PeriodEnd string) {
	if r != nil && r.PeriodEnd != nil {
		PeriodEnd = *r.PeriodEnd
	}
	return
}

// The account on which the interest was accrued.
func (r *ACHTransferSimulationTransactionSourceInterestPayment) GetAccruedOnAccountID() (AccruedOnAccountID string) {
	if r != nil && r.AccruedOnAccountID != nil {
		AccruedOnAccountID = *r.AccruedOnAccountID
	}
	return
}

func (r ACHTransferSimulationTransactionSourceInterestPayment) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceInterestPayment{Amount:%s Currency:%s PeriodStart:%s PeriodEnd:%s AccruedOnAccountID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.PeriodStart), core.FmtP(r.PeriodEnd), core.FmtP(r.AccruedOnAccountID))
}

type ACHTransferSimulationTransactionSourceInterestPaymentCurrency string

const (
	ACHTransferSimulationTransactionSourceInterestPaymentCurrencyCad ACHTransferSimulationTransactionSourceInterestPaymentCurrency = "CAD"
	ACHTransferSimulationTransactionSourceInterestPaymentCurrencyChf ACHTransferSimulationTransactionSourceInterestPaymentCurrency = "CHF"
	ACHTransferSimulationTransactionSourceInterestPaymentCurrencyEur ACHTransferSimulationTransactionSourceInterestPaymentCurrency = "EUR"
	ACHTransferSimulationTransactionSourceInterestPaymentCurrencyGbp ACHTransferSimulationTransactionSourceInterestPaymentCurrency = "GBP"
	ACHTransferSimulationTransactionSourceInterestPaymentCurrencyJpy ACHTransferSimulationTransactionSourceInterestPaymentCurrency = "JPY"
	ACHTransferSimulationTransactionSourceInterestPaymentCurrencyUsd ACHTransferSimulationTransactionSourceInterestPaymentCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceInternalSource struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency   *ACHTransferSimulationTransactionSourceInternalSourceCurrency `pjson:"currency"`
	Reason     *ACHTransferSimulationTransactionSourceInternalSourceReason   `pjson:"reason"`
	jsonFields map[string]interface{}                                        `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceInternalSource using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceInternalSource) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceInternalSource into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceInternalSource) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransactionSourceInternalSource) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
func (r *ACHTransferSimulationTransactionSourceInternalSource) GetCurrency() (Currency ACHTransferSimulationTransactionSourceInternalSourceCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInternalSource) GetReason() (Reason ACHTransferSimulationTransactionSourceInternalSourceReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

func (r ACHTransferSimulationTransactionSourceInternalSource) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceInternalSource{Amount:%s Currency:%s Reason:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.Reason))
}

type ACHTransferSimulationTransactionSourceInternalSourceCurrency string

const (
	ACHTransferSimulationTransactionSourceInternalSourceCurrencyCad ACHTransferSimulationTransactionSourceInternalSourceCurrency = "CAD"
	ACHTransferSimulationTransactionSourceInternalSourceCurrencyChf ACHTransferSimulationTransactionSourceInternalSourceCurrency = "CHF"
	ACHTransferSimulationTransactionSourceInternalSourceCurrencyEur ACHTransferSimulationTransactionSourceInternalSourceCurrency = "EUR"
	ACHTransferSimulationTransactionSourceInternalSourceCurrencyGbp ACHTransferSimulationTransactionSourceInternalSourceCurrency = "GBP"
	ACHTransferSimulationTransactionSourceInternalSourceCurrencyJpy ACHTransferSimulationTransactionSourceInternalSourceCurrency = "JPY"
	ACHTransferSimulationTransactionSourceInternalSourceCurrencyUsd ACHTransferSimulationTransactionSourceInternalSourceCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceInternalSourceReason string

const (
	ACHTransferSimulationTransactionSourceInternalSourceReasonBankMigration      ACHTransferSimulationTransactionSourceInternalSourceReason = "bank_migration"
	ACHTransferSimulationTransactionSourceInternalSourceReasonCashback           ACHTransferSimulationTransactionSourceInternalSourceReason = "cashback"
	ACHTransferSimulationTransactionSourceInternalSourceReasonEmpyrealAdjustment ACHTransferSimulationTransactionSourceInternalSourceReason = "empyreal_adjustment"
	ACHTransferSimulationTransactionSourceInternalSourceReasonError              ACHTransferSimulationTransactionSourceInternalSourceReason = "error"
	ACHTransferSimulationTransactionSourceInternalSourceReasonErrorCorrection    ACHTransferSimulationTransactionSourceInternalSourceReason = "error_correction"
	ACHTransferSimulationTransactionSourceInternalSourceReasonFees               ACHTransferSimulationTransactionSourceInternalSourceReason = "fees"
	ACHTransferSimulationTransactionSourceInternalSourceReasonInterest           ACHTransferSimulationTransactionSourceInternalSourceReason = "interest"
	ACHTransferSimulationTransactionSourceInternalSourceReasonSampleFunds        ACHTransferSimulationTransactionSourceInternalSourceReason = "sample_funds"
	ACHTransferSimulationTransactionSourceInternalSourceReasonSampleFundsReturn  ACHTransferSimulationTransactionSourceInternalSourceReason = "sample_funds_return"
)

type ACHTransferSimulationTransactionSourceCardRouteRefund struct {
	// The refunded amount in the minor unit of the refunded currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
	// currency.
	Currency             *ACHTransferSimulationTransactionSourceCardRouteRefundCurrency `pjson:"currency"`
	MerchantAcceptorID   *string                                                        `pjson:"merchant_acceptor_id"`
	MerchantCity         *string                                                        `pjson:"merchant_city"`
	MerchantCountry      *string                                                        `pjson:"merchant_country"`
	MerchantDescriptor   *string                                                        `pjson:"merchant_descriptor"`
	MerchantState        *string                                                        `pjson:"merchant_state"`
	MerchantCategoryCode *string                                                        `pjson:"merchant_category_code"`
	jsonFields           map[string]interface{}                                         `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceCardRouteRefund using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCardRouteRefund) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceCardRouteRefund
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceCardRouteRefund) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The refunded amount in the minor unit of the refunded currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransactionSourceCardRouteRefund) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
// currency.
func (r *ACHTransferSimulationTransactionSourceCardRouteRefund) GetCurrency() (Currency ACHTransferSimulationTransactionSourceCardRouteRefundCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteRefund) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r != nil && r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteRefund) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteRefund) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteRefund) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteRefund) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteRefund) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

func (r ACHTransferSimulationTransactionSourceCardRouteRefund) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCardRouteRefund{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.MerchantAcceptorID), core.FmtP(r.MerchantCity), core.FmtP(r.MerchantCountry), core.FmtP(r.MerchantDescriptor), core.FmtP(r.MerchantState), core.FmtP(r.MerchantCategoryCode))
}

type ACHTransferSimulationTransactionSourceCardRouteRefundCurrency string

const (
	ACHTransferSimulationTransactionSourceCardRouteRefundCurrencyCad ACHTransferSimulationTransactionSourceCardRouteRefundCurrency = "CAD"
	ACHTransferSimulationTransactionSourceCardRouteRefundCurrencyChf ACHTransferSimulationTransactionSourceCardRouteRefundCurrency = "CHF"
	ACHTransferSimulationTransactionSourceCardRouteRefundCurrencyEur ACHTransferSimulationTransactionSourceCardRouteRefundCurrency = "EUR"
	ACHTransferSimulationTransactionSourceCardRouteRefundCurrencyGbp ACHTransferSimulationTransactionSourceCardRouteRefundCurrency = "GBP"
	ACHTransferSimulationTransactionSourceCardRouteRefundCurrencyJpy ACHTransferSimulationTransactionSourceCardRouteRefundCurrency = "JPY"
	ACHTransferSimulationTransactionSourceCardRouteRefundCurrencyUsd ACHTransferSimulationTransactionSourceCardRouteRefundCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceCardRouteSettlement struct {
	// The settled amount in the minor unit of the settlement currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
	// currency.
	Currency             *ACHTransferSimulationTransactionSourceCardRouteSettlementCurrency `pjson:"currency"`
	MerchantAcceptorID   *string                                                            `pjson:"merchant_acceptor_id"`
	MerchantCity         *string                                                            `pjson:"merchant_city"`
	MerchantCountry      *string                                                            `pjson:"merchant_country"`
	MerchantDescriptor   *string                                                            `pjson:"merchant_descriptor"`
	MerchantState        *string                                                            `pjson:"merchant_state"`
	MerchantCategoryCode *string                                                            `pjson:"merchant_category_code"`
	jsonFields           map[string]interface{}                                             `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceCardRouteSettlement using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCardRouteSettlement) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceCardRouteSettlement
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceCardRouteSettlement) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The settled amount in the minor unit of the settlement currency. For dollars,
// for example, this is cents.
func (r *ACHTransferSimulationTransactionSourceCardRouteSettlement) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
// currency.
func (r *ACHTransferSimulationTransactionSourceCardRouteSettlement) GetCurrency() (Currency ACHTransferSimulationTransactionSourceCardRouteSettlementCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r != nil && r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

func (r ACHTransferSimulationTransactionSourceCardRouteSettlement) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCardRouteSettlement{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.MerchantAcceptorID), core.FmtP(r.MerchantCity), core.FmtP(r.MerchantCountry), core.FmtP(r.MerchantDescriptor), core.FmtP(r.MerchantState), core.FmtP(r.MerchantCategoryCode))
}

type ACHTransferSimulationTransactionSourceCardRouteSettlementCurrency string

const (
	ACHTransferSimulationTransactionSourceCardRouteSettlementCurrencyCad ACHTransferSimulationTransactionSourceCardRouteSettlementCurrency = "CAD"
	ACHTransferSimulationTransactionSourceCardRouteSettlementCurrencyChf ACHTransferSimulationTransactionSourceCardRouteSettlementCurrency = "CHF"
	ACHTransferSimulationTransactionSourceCardRouteSettlementCurrencyEur ACHTransferSimulationTransactionSourceCardRouteSettlementCurrency = "EUR"
	ACHTransferSimulationTransactionSourceCardRouteSettlementCurrencyGbp ACHTransferSimulationTransactionSourceCardRouteSettlementCurrency = "GBP"
	ACHTransferSimulationTransactionSourceCardRouteSettlementCurrencyJpy ACHTransferSimulationTransactionSourceCardRouteSettlementCurrency = "JPY"
	ACHTransferSimulationTransactionSourceCardRouteSettlementCurrencyUsd ACHTransferSimulationTransactionSourceCardRouteSettlementCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceSampleFunds struct {
	// Where the sample funds came from.
	Originator *string                `pjson:"originator"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceSampleFunds using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceSampleFunds) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceSampleFunds into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceSampleFunds) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Where the sample funds came from.
func (r *ACHTransferSimulationTransactionSourceSampleFunds) GetOriginator() (Originator string) {
	if r != nil && r.Originator != nil {
		Originator = *r.Originator
	}
	return
}

func (r ACHTransferSimulationTransactionSourceSampleFunds) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceSampleFunds{Originator:%s}", core.FmtP(r.Originator))
}

type ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention struct {
	// The transfer amount in USD cents.
	Amount             *int64                 `pjson:"amount"`
	AccountNumber      *string                `pjson:"account_number"`
	RoutingNumber      *string                `pjson:"routing_number"`
	MessageToRecipient *string                `pjson:"message_to_recipient"`
	TransferID         *string                `pjson:"transfer_id"`
	jsonFields         map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The transfer amount in USD cents.
func (r *ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s TransferID:%s}", core.FmtP(r.Amount), core.FmtP(r.AccountNumber), core.FmtP(r.RoutingNumber), core.FmtP(r.MessageToRecipient), core.FmtP(r.TransferID))
}

type ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection struct {
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

func (r *ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection{TransferID:%s}", core.FmtP(r.TransferID))
}

type ACHTransferSimulationTransactionSourceWireTransferIntention struct {
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
// ACHTransferSimulationTransactionSourceWireTransferIntention using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceWireTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceWireTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSourceWireTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The transfer amount in USD cents.
func (r *ACHTransferSimulationTransactionSourceWireTransferIntention) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The destination account number.
func (r *ACHTransferSimulationTransactionSourceWireTransferIntention) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r *ACHTransferSimulationTransactionSourceWireTransferIntention) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The message that will show on the recipient's bank statement.
func (r *ACHTransferSimulationTransactionSourceWireTransferIntention) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceWireTransferIntention) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r ACHTransferSimulationTransactionSourceWireTransferIntention) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceWireTransferIntention{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s TransferID:%s}", core.FmtP(r.Amount), core.FmtP(r.AccountNumber), core.FmtP(r.RoutingNumber), core.FmtP(r.MessageToRecipient), core.FmtP(r.TransferID))
}

type ACHTransferSimulationTransactionSourceWireTransferRejection struct {
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceWireTransferRejection using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceWireTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceWireTransferRejection into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSourceWireTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

func (r *ACHTransferSimulationTransactionSourceWireTransferRejection) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r ACHTransferSimulationTransactionSourceWireTransferRejection) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceWireTransferRejection{TransferID:%s}", core.FmtP(r.TransferID))
}

type ACHTransferSimulationTransactionType string

const (
	ACHTransferSimulationTransactionTypeTransaction ACHTransferSimulationTransactionType = "transaction"
)

type ACHTransferSimulationDeclinedTransaction struct {
	// The identifier for the Account the Declined Transaction belongs to.
	AccountID *string `pjson:"account_id"`
	// The Declined Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transcation's Account.
	Currency *ACHTransferSimulationDeclinedTransactionCurrency `pjson:"currency"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt *string `pjson:"created_at"`
	// This is the description the vendor provides.
	Description *string `pjson:"description"`
	// The Declined Transaction identifier.
	ID *string `pjson:"id"`
	// The identifier for the route this Declined Transaction came through. Routes are
	// things like cards and ACH details.
	RouteID *string `pjson:"route_id"`
	// The type of the route this Declined Transaction came through.
	RouteType *string `pjson:"route_type"`
	// This is an object giving more details on the network-level event that caused the
	// Declined Transaction. For example, for a card transaction this lists the
	// merchant's industry and location. Note that for backwards compatibility reasons,
	// additional undocumented keys may appear in this object. These should be treated
	// as deprecated and will be removed in the future.
	Source *ACHTransferSimulationDeclinedTransactionSource `pjson:"source"`
	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type       *ACHTransferSimulationDeclinedTransactionType `pjson:"type"`
	jsonFields map[string]interface{}                        `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationDeclinedTransaction using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationDeclinedTransaction) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationDeclinedTransaction into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationDeclinedTransaction) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier for the Account the Declined Transaction belongs to.
func (r *ACHTransferSimulationDeclinedTransaction) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The Declined Transaction amount in the minor unit of its currency. For dollars,
// for example, this is cents.
func (r *ACHTransferSimulationDeclinedTransaction) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
// Transaction's currency. This will match the currency on the Declined
// Transcation's Account.
func (r *ACHTransferSimulationDeclinedTransaction) GetCurrency() (Currency ACHTransferSimulationDeclinedTransactionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
// Transaction occured.
func (r *ACHTransferSimulationDeclinedTransaction) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// This is the description the vendor provides.
func (r *ACHTransferSimulationDeclinedTransaction) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Declined Transaction identifier.
func (r *ACHTransferSimulationDeclinedTransaction) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The identifier for the route this Declined Transaction came through. Routes are
// things like cards and ACH details.
func (r *ACHTransferSimulationDeclinedTransaction) GetRouteID() (RouteID string) {
	if r != nil && r.RouteID != nil {
		RouteID = *r.RouteID
	}
	return
}

// The type of the route this Declined Transaction came through.
func (r *ACHTransferSimulationDeclinedTransaction) GetRouteType() (RouteType string) {
	if r != nil && r.RouteType != nil {
		RouteType = *r.RouteType
	}
	return
}

// This is an object giving more details on the network-level event that caused the
// Declined Transaction. For example, for a card transaction this lists the
// merchant's industry and location. Note that for backwards compatibility reasons,
// additional undocumented keys may appear in this object. These should be treated
// as deprecated and will be removed in the future.
func (r *ACHTransferSimulationDeclinedTransaction) GetSource() (Source ACHTransferSimulationDeclinedTransactionSource) {
	if r != nil && r.Source != nil {
		Source = *r.Source
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `declined_transaction`.
func (r *ACHTransferSimulationDeclinedTransaction) GetType() (Type ACHTransferSimulationDeclinedTransactionType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r ACHTransferSimulationDeclinedTransaction) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationDeclinedTransaction{AccountID:%s Amount:%s Currency:%s CreatedAt:%s Description:%s ID:%s RouteID:%s RouteType:%s Source:%s Type:%s}", core.FmtP(r.AccountID), core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.CreatedAt), core.FmtP(r.Description), core.FmtP(r.ID), core.FmtP(r.RouteID), core.FmtP(r.RouteType), r.Source, core.FmtP(r.Type))
}

type ACHTransferSimulationDeclinedTransactionCurrency string

const (
	ACHTransferSimulationDeclinedTransactionCurrencyCad ACHTransferSimulationDeclinedTransactionCurrency = "CAD"
	ACHTransferSimulationDeclinedTransactionCurrencyChf ACHTransferSimulationDeclinedTransactionCurrency = "CHF"
	ACHTransferSimulationDeclinedTransactionCurrencyEur ACHTransferSimulationDeclinedTransactionCurrency = "EUR"
	ACHTransferSimulationDeclinedTransactionCurrencyGbp ACHTransferSimulationDeclinedTransactionCurrency = "GBP"
	ACHTransferSimulationDeclinedTransactionCurrencyJpy ACHTransferSimulationDeclinedTransactionCurrency = "JPY"
	ACHTransferSimulationDeclinedTransactionCurrencyUsd ACHTransferSimulationDeclinedTransactionCurrency = "USD"
)

type ACHTransferSimulationDeclinedTransactionSource struct {
	// The type of decline that took place. We may add additional possible values for
	// this enum over time; your application should be able to handle such additions
	// gracefully.
	Category *ACHTransferSimulationDeclinedTransactionSourceCategory `pjson:"category"`
	// A ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline *ACHTransferSimulationDeclinedTransactionSourceACHDecline `pjson:"ach_decline"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline *ACHTransferSimulationDeclinedTransactionSourceCardDecline `pjson:"card_decline"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline *ACHTransferSimulationDeclinedTransactionSourceCheckDecline `pjson:"check_decline"`
	// A Inbound Real Time Payments Transfer Decline object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline `pjson:"inbound_real_time_payments_transfer_decline"`
	// A International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline `pjson:"international_ach_decline"`
	// A Deprecated Card Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_decline`.
	CardRouteDecline *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline `pjson:"card_route_decline"`
	jsonFields       map[string]interface{}                                          `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationDeclinedTransactionSource using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationDeclinedTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationDeclinedTransactionSource into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationDeclinedTransactionSource) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The type of decline that took place. We may add additional possible values for
// this enum over time; your application should be able to handle such additions
// gracefully.
func (r *ACHTransferSimulationDeclinedTransactionSource) GetCategory() (Category ACHTransferSimulationDeclinedTransactionSourceCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// A ACH Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `ach_decline`.
func (r *ACHTransferSimulationDeclinedTransactionSource) GetACHDecline() (ACHDecline ACHTransferSimulationDeclinedTransactionSourceACHDecline) {
	if r != nil && r.ACHDecline != nil {
		ACHDecline = *r.ACHDecline
	}
	return
}

// A Card Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `card_decline`.
func (r *ACHTransferSimulationDeclinedTransactionSource) GetCardDecline() (CardDecline ACHTransferSimulationDeclinedTransactionSourceCardDecline) {
	if r != nil && r.CardDecline != nil {
		CardDecline = *r.CardDecline
	}
	return
}

// A Check Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `check_decline`.
func (r *ACHTransferSimulationDeclinedTransactionSource) GetCheckDecline() (CheckDecline ACHTransferSimulationDeclinedTransactionSourceCheckDecline) {
	if r != nil && r.CheckDecline != nil {
		CheckDecline = *r.CheckDecline
	}
	return
}

// A Inbound Real Time Payments Transfer Decline object. This field will be present
// in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_decline`.
func (r *ACHTransferSimulationDeclinedTransactionSource) GetInboundRealTimePaymentsTransferDecline() (InboundRealTimePaymentsTransferDecline ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) {
	if r != nil && r.InboundRealTimePaymentsTransferDecline != nil {
		InboundRealTimePaymentsTransferDecline = *r.InboundRealTimePaymentsTransferDecline
	}
	return
}

// A International ACH Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `international_ach_decline`.
func (r *ACHTransferSimulationDeclinedTransactionSource) GetInternationalACHDecline() (InternationalACHDecline ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) {
	if r != nil && r.InternationalACHDecline != nil {
		InternationalACHDecline = *r.InternationalACHDecline
	}
	return
}

// A Deprecated Card Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_decline`.
func (r *ACHTransferSimulationDeclinedTransactionSource) GetCardRouteDecline() (CardRouteDecline ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) {
	if r != nil && r.CardRouteDecline != nil {
		CardRouteDecline = *r.CardRouteDecline
	}
	return
}

func (r ACHTransferSimulationDeclinedTransactionSource) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationDeclinedTransactionSource{Category:%s ACHDecline:%s CardDecline:%s CheckDecline:%s InboundRealTimePaymentsTransferDecline:%s InternationalACHDecline:%s CardRouteDecline:%s}", core.FmtP(r.Category), r.ACHDecline, r.CardDecline, r.CheckDecline, r.InboundRealTimePaymentsTransferDecline, r.InternationalACHDecline, r.CardRouteDecline)
}

type ACHTransferSimulationDeclinedTransactionSourceCategory string

const (
	ACHTransferSimulationDeclinedTransactionSourceCategoryACHDecline                             ACHTransferSimulationDeclinedTransactionSourceCategory = "ach_decline"
	ACHTransferSimulationDeclinedTransactionSourceCategoryCardDecline                            ACHTransferSimulationDeclinedTransactionSourceCategory = "card_decline"
	ACHTransferSimulationDeclinedTransactionSourceCategoryCheckDecline                           ACHTransferSimulationDeclinedTransactionSourceCategory = "check_decline"
	ACHTransferSimulationDeclinedTransactionSourceCategoryInboundRealTimePaymentsTransferDecline ACHTransferSimulationDeclinedTransactionSourceCategory = "inbound_real_time_payments_transfer_decline"
	ACHTransferSimulationDeclinedTransactionSourceCategoryInternationalACHDecline                ACHTransferSimulationDeclinedTransactionSourceCategory = "international_ach_decline"
	ACHTransferSimulationDeclinedTransactionSourceCategoryCardRouteDecline                       ACHTransferSimulationDeclinedTransactionSourceCategory = "card_route_decline"
	ACHTransferSimulationDeclinedTransactionSourceCategoryOther                                  ACHTransferSimulationDeclinedTransactionSourceCategory = "other"
)

type ACHTransferSimulationDeclinedTransactionSourceACHDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount                             *int64  `pjson:"amount"`
	OriginatorCompanyName              *string `pjson:"originator_company_name"`
	OriginatorCompanyDescriptiveDate   *string `pjson:"originator_company_descriptive_date"`
	OriginatorCompanyDiscretionaryData *string `pjson:"originator_company_discretionary_data"`
	OriginatorCompanyID                *string `pjson:"originator_company_id"`
	// Why the ACH transfer was declined.
	Reason           *ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason `pjson:"reason"`
	ReceiverIDNumber *string                                                         `pjson:"receiver_id_number"`
	ReceiverName     *string                                                         `pjson:"receiver_name"`
	TraceNumber      *string                                                         `pjson:"trace_number"`
	jsonFields       map[string]interface{}                                          `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationDeclinedTransactionSourceACHDecline using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationDeclinedTransactionSourceACHDecline
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) GetOriginatorCompanyName() (OriginatorCompanyName string) {
	if r != nil && r.OriginatorCompanyName != nil {
		OriginatorCompanyName = *r.OriginatorCompanyName
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) GetOriginatorCompanyDescriptiveDate() (OriginatorCompanyDescriptiveDate string) {
	if r != nil && r.OriginatorCompanyDescriptiveDate != nil {
		OriginatorCompanyDescriptiveDate = *r.OriginatorCompanyDescriptiveDate
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) GetOriginatorCompanyDiscretionaryData() (OriginatorCompanyDiscretionaryData string) {
	if r != nil && r.OriginatorCompanyDiscretionaryData != nil {
		OriginatorCompanyDiscretionaryData = *r.OriginatorCompanyDiscretionaryData
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) GetOriginatorCompanyID() (OriginatorCompanyID string) {
	if r != nil && r.OriginatorCompanyID != nil {
		OriginatorCompanyID = *r.OriginatorCompanyID
	}
	return
}

// Why the ACH transfer was declined.
func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) GetReason() (Reason ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) GetReceiverIDNumber() (ReceiverIDNumber string) {
	if r != nil && r.ReceiverIDNumber != nil {
		ReceiverIDNumber = *r.ReceiverIDNumber
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) GetReceiverName() (ReceiverName string) {
	if r != nil && r.ReceiverName != nil {
		ReceiverName = *r.ReceiverName
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

func (r ACHTransferSimulationDeclinedTransactionSourceACHDecline) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationDeclinedTransactionSourceACHDecline{Amount:%s OriginatorCompanyName:%s OriginatorCompanyDescriptiveDate:%s OriginatorCompanyDiscretionaryData:%s OriginatorCompanyID:%s Reason:%s ReceiverIDNumber:%s ReceiverName:%s TraceNumber:%s}", core.FmtP(r.Amount), core.FmtP(r.OriginatorCompanyName), core.FmtP(r.OriginatorCompanyDescriptiveDate), core.FmtP(r.OriginatorCompanyDiscretionaryData), core.FmtP(r.OriginatorCompanyID), core.FmtP(r.Reason), core.FmtP(r.ReceiverIDNumber), core.FmtP(r.ReceiverName), core.FmtP(r.TraceNumber))
}

type ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason string

const (
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteCanceled             ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "ach_route_canceled"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteDisabled             ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "ach_route_disabled"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonBreachesLimit                ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonDuplicateReturn              ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "duplicate_return"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonEntityNotActive              ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "entity_not_active"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonTransactionNotAllowed        ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "transaction_not_allowed"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonGroupLocked                  ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "group_locked"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonInsufficientFunds            ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "insufficient_funds"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonNoACHRoute                   ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "no_ach_route"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonOriginatorRequest            ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "originator_request"
)

type ACHTransferSimulationDeclinedTransactionSourceCardDecline struct {
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID *string `pjson:"merchant_acceptor_id"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor *string `pjson:"merchant_descriptor"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode *string `pjson:"merchant_category_code"`
	// The city the merchant resides in.
	MerchantCity *string `pjson:"merchant_city"`
	// The country the merchant resides in.
	MerchantCountry *string `pjson:"merchant_country"`
	// The payment network used to process this card authorization
	Network *ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetwork `pjson:"network"`
	// Fields specific to the `network`
	NetworkDetails *ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails `pjson:"network_details"`
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency `pjson:"currency"`
	// Why the transaction was declined.
	Reason *ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason `pjson:"reason"`
	// The state the merchant resides in.
	MerchantState *string `pjson:"merchant_state"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID *string `pjson:"real_time_decision_id"`
	// If the authorization was attempted using a Digital Wallet Token (such as an
	// Apple Pay purchase), the identifier of the token that was used.
	DigitalWalletTokenID *string                `pjson:"digital_wallet_token_id"`
	jsonFields           map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationDeclinedTransactionSourceCardDecline using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSimulationDeclinedTransactionSourceCardDecline
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The merchant identifier (commonly abbreviated as MID) of the merchant the card
// is transacting with.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r != nil && r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
}

// The merchant descriptor of the merchant the card is transacting with.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
// card is transacting with.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

// The city the merchant resides in.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

// The country the merchant resides in.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

// The payment network used to process this card authorization
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetNetwork() (Network ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetwork) {
	if r != nil && r.Network != nil {
		Network = *r.Network
	}
	return
}

// Fields specific to the `network`
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetNetworkDetails() (NetworkDetails ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails) {
	if r != nil && r.NetworkDetails != nil {
		NetworkDetails = *r.NetworkDetails
	}
	return
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetCurrency() (Currency ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// Why the transaction was declined.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetReason() (Reason ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

// The state the merchant resides in.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

// The identifier of the Real-Time Decision sent to approve or decline this
// transaction.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetRealTimeDecisionID() (RealTimeDecisionID string) {
	if r != nil && r.RealTimeDecisionID != nil {
		RealTimeDecisionID = *r.RealTimeDecisionID
	}
	return
}

// If the authorization was attempted using a Digital Wallet Token (such as an
// Apple Pay purchase), the identifier of the token that was used.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetDigitalWalletTokenID() (DigitalWalletTokenID string) {
	if r != nil && r.DigitalWalletTokenID != nil {
		DigitalWalletTokenID = *r.DigitalWalletTokenID
	}
	return
}

func (r ACHTransferSimulationDeclinedTransactionSourceCardDecline) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationDeclinedTransactionSourceCardDecline{MerchantAcceptorID:%s MerchantDescriptor:%s MerchantCategoryCode:%s MerchantCity:%s MerchantCountry:%s Network:%s NetworkDetails:%s Amount:%s Currency:%s Reason:%s MerchantState:%s RealTimeDecisionID:%s DigitalWalletTokenID:%s}", core.FmtP(r.MerchantAcceptorID), core.FmtP(r.MerchantDescriptor), core.FmtP(r.MerchantCategoryCode), core.FmtP(r.MerchantCity), core.FmtP(r.MerchantCountry), core.FmtP(r.Network), r.NetworkDetails, core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.Reason), core.FmtP(r.MerchantState), core.FmtP(r.RealTimeDecisionID), core.FmtP(r.DigitalWalletTokenID))
}

type ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetwork string

const (
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkVisa ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetwork = "visa"
)

type ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa       *ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa `pjson:"visa"`
	jsonFields map[string]interface{}                                                       `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Fields specific to the `visa` network
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails) GetVisa() (Visa ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) {
	if r != nil && r.Visa != nil {
		Visa = *r.Visa
	}
	return
}

func (r ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails{Visa:%s}", r.Visa)
}

type ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator *ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator `pjson:"electronic_commerce_indicator"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode *PointOfServiceEntryMode `pjson:"point_of_service_entry_mode"`
	jsonFields              map[string]interface{}   `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// For electronic commerce transactions, this identifies the level of security used
// in obtaining the customer's payment credential. For mail or telephone order
// transactions, identifies the type of mail or telephone order.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) GetElectronicCommerceIndicator() (ElectronicCommerceIndicator ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator) {
	if r != nil && r.ElectronicCommerceIndicator != nil {
		ElectronicCommerceIndicator = *r.ElectronicCommerceIndicator
	}
	return
}

// The method used to enter the cardholder's primary account number and card
// expiration date
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) GetPointOfServiceEntryMode() (PointOfServiceEntryMode PointOfServiceEntryMode) {
	if r != nil && r.PointOfServiceEntryMode != nil {
		PointOfServiceEntryMode = *r.PointOfServiceEntryMode
	}
	return
}

func (r ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa{ElectronicCommerceIndicator:%s PointOfServiceEntryMode:%s}", core.FmtP(r.ElectronicCommerceIndicator), core.FmtP(r.PointOfServiceEntryMode))
}

type ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator string

const (
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                           ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                                ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                              ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                    ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                 ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt_3dsCapableMerchant ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                      ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                     ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

type ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency string

const (
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrencyCad ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency = "CAD"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrencyChf ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency = "CHF"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrencyEur ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency = "EUR"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrencyGbp ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency = "GBP"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrencyJpy ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency = "JPY"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrencyUsd ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency = "USD"
)

type ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason string

const (
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonCardNotActive               ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "card_not_active"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonEntityNotActive             ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "entity_not_active"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonGroupLocked                 ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "group_locked"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonInsufficientFunds           ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "insufficient_funds"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonCvv2Mismatch                ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "cvv2_mismatch"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonTransactionNotAllowed       ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "transaction_not_allowed"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonBreachesLimit               ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "breaches_limit"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonWebhookDeclined             ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "webhook_declined"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonWebhookTimedOut             ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "webhook_timed_out"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonDeclinedByStandInProcessing ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "declined_by_stand_in_processing"
)

type ACHTransferSimulationDeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount        *int64  `pjson:"amount"`
	AuxiliaryOnUs *string `pjson:"auxiliary_on_us"`
	// Why the check was declined.
	Reason     *ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason `pjson:"reason"`
	jsonFields map[string]interface{}                                            `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationDeclinedTransactionSourceCheckDecline using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationDeclinedTransactionSourceCheckDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationDeclinedTransactionSourceCheckDecline into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationDeclinedTransactionSourceCheckDecline) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *ACHTransferSimulationDeclinedTransactionSourceCheckDecline) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCheckDecline) GetAuxiliaryOnUs() (AuxiliaryOnUs string) {
	if r != nil && r.AuxiliaryOnUs != nil {
		AuxiliaryOnUs = *r.AuxiliaryOnUs
	}
	return
}

// Why the check was declined.
func (r *ACHTransferSimulationDeclinedTransactionSourceCheckDecline) GetReason() (Reason ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

func (r ACHTransferSimulationDeclinedTransactionSourceCheckDecline) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationDeclinedTransactionSourceCheckDecline{Amount:%s AuxiliaryOnUs:%s Reason:%s}", core.FmtP(r.Amount), core.FmtP(r.AuxiliaryOnUs), core.FmtP(r.Reason))
}

type ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason string

const (
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonACHRouteCanceled      ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "ach_route_canceled"
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonACHRouteDisabled      ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "ach_route_disabled"
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonBreachesLimit         ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "breaches_limit"
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonEntityNotActive       ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "entity_not_active"
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonGroupLocked           ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "group_locked"
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonInsufficientFunds     ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "insufficient_funds"
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonUnableToLocateAccount ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "unable_to_locate_account"
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonUnableToProcess       ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "unable_to_process"
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonReferToImage          ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "refer_to_image"
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonStopPaymentRequested  ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "stop_payment_requested"
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonReturned              ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "returned"
)

type ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real Time Payments
	// transfer.
	Currency *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency `pjson:"currency"`
	// Why the transfer was declined.
	Reason *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason `pjson:"reason"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName *string `pjson:"creditor_name"`
	// The name provided by the sender of the transfer.
	DebtorName *string `pjson:"debtor_name"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber *string `pjson:"debtor_account_number"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber *string `pjson:"debtor_routing_number"`
	// The Real Time Payments network identification of the declined transfer.
	TransactionIdentification *string `pjson:"transaction_identification"`
	// Additional information included with the transfer.
	RemittanceInformation *string                `pjson:"remittance_information"`
	jsonFields            map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
// transfer's currency. This will always be "USD" for a Real Time Payments
// transfer.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetCurrency() (Currency ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// Why the transfer was declined.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetReason() (Reason ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

// The name the sender of the transfer specified as the recipient of the transfer.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetCreditorName() (CreditorName string) {
	if r != nil && r.CreditorName != nil {
		CreditorName = *r.CreditorName
	}
	return
}

// The name provided by the sender of the transfer.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetDebtorName() (DebtorName string) {
	if r != nil && r.DebtorName != nil {
		DebtorName = *r.DebtorName
	}
	return
}

// The account number of the account that sent the transfer.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetDebtorAccountNumber() (DebtorAccountNumber string) {
	if r != nil && r.DebtorAccountNumber != nil {
		DebtorAccountNumber = *r.DebtorAccountNumber
	}
	return
}

// The routing number of the account that sent the transfer.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetDebtorRoutingNumber() (DebtorRoutingNumber string) {
	if r != nil && r.DebtorRoutingNumber != nil {
		DebtorRoutingNumber = *r.DebtorRoutingNumber
	}
	return
}

// The Real Time Payments network identification of the declined transfer.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetTransactionIdentification() (TransactionIdentification string) {
	if r != nil && r.TransactionIdentification != nil {
		TransactionIdentification = *r.TransactionIdentification
	}
	return
}

// Additional information included with the transfer.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetRemittanceInformation() (RemittanceInformation string) {
	if r != nil && r.RemittanceInformation != nil {
		RemittanceInformation = *r.RemittanceInformation
	}
	return
}

func (r ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline{Amount:%s Currency:%s Reason:%s CreditorName:%s DebtorName:%s DebtorAccountNumber:%s DebtorRoutingNumber:%s TransactionIdentification:%s RemittanceInformation:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.Reason), core.FmtP(r.CreditorName), core.FmtP(r.DebtorName), core.FmtP(r.DebtorAccountNumber), core.FmtP(r.DebtorRoutingNumber), core.FmtP(r.TransactionIdentification), core.FmtP(r.RemittanceInformation))
}

type ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency string

const (
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyCad ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CAD"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyChf ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CHF"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyEur ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "EUR"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyGbp ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "GBP"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyJpy ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "JPY"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyUsd ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "USD"
)

type ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason string

const (
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberCanceled      ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_canceled"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberDisabled      ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_disabled"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonGroupLocked                ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "group_locked"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonEntityNotActive            ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "entity_not_active"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "real_time_payments_not_enabled"
)

type ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
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
// ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetForeignExchangeIndicator() (ForeignExchangeIndicator string) {
	if r != nil && r.ForeignExchangeIndicator != nil {
		ForeignExchangeIndicator = *r.ForeignExchangeIndicator
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetForeignExchangeReferenceIndicator() (ForeignExchangeReferenceIndicator string) {
	if r != nil && r.ForeignExchangeReferenceIndicator != nil {
		ForeignExchangeReferenceIndicator = *r.ForeignExchangeReferenceIndicator
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetForeignExchangeReference() (ForeignExchangeReference string) {
	if r != nil && r.ForeignExchangeReference != nil {
		ForeignExchangeReference = *r.ForeignExchangeReference
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetDestinationCountryCode() (DestinationCountryCode string) {
	if r != nil && r.DestinationCountryCode != nil {
		DestinationCountryCode = *r.DestinationCountryCode
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetDestinationCurrencyCode() (DestinationCurrencyCode string) {
	if r != nil && r.DestinationCurrencyCode != nil {
		DestinationCurrencyCode = *r.DestinationCurrencyCode
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetForeignPaymentAmount() (ForeignPaymentAmount int64) {
	if r != nil && r.ForeignPaymentAmount != nil {
		ForeignPaymentAmount = *r.ForeignPaymentAmount
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetForeignTraceNumber() (ForeignTraceNumber string) {
	if r != nil && r.ForeignTraceNumber != nil {
		ForeignTraceNumber = *r.ForeignTraceNumber
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetInternationalTransactionTypeCode() (InternationalTransactionTypeCode string) {
	if r != nil && r.InternationalTransactionTypeCode != nil {
		InternationalTransactionTypeCode = *r.InternationalTransactionTypeCode
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatingCurrencyCode() (OriginatingCurrencyCode string) {
	if r != nil && r.OriginatingCurrencyCode != nil {
		OriginatingCurrencyCode = *r.OriginatingCurrencyCode
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionName() (OriginatingDepositoryFinancialInstitutionName string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionName != nil {
		OriginatingDepositoryFinancialInstitutionName = *r.OriginatingDepositoryFinancialInstitutionName
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionIDQualifier() (OriginatingDepositoryFinancialInstitutionIDQualifier string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionIDQualifier != nil {
		OriginatingDepositoryFinancialInstitutionIDQualifier = *r.OriginatingDepositoryFinancialInstitutionIDQualifier
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionID() (OriginatingDepositoryFinancialInstitutionID string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionID != nil {
		OriginatingDepositoryFinancialInstitutionID = *r.OriginatingDepositoryFinancialInstitutionID
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionBranchCountry() (OriginatingDepositoryFinancialInstitutionBranchCountry string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionBranchCountry != nil {
		OriginatingDepositoryFinancialInstitutionBranchCountry = *r.OriginatingDepositoryFinancialInstitutionBranchCountry
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorCity() (OriginatorCity string) {
	if r != nil && r.OriginatorCity != nil {
		OriginatorCity = *r.OriginatorCity
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r != nil && r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorCountry() (OriginatorCountry string) {
	if r != nil && r.OriginatorCountry != nil {
		OriginatorCountry = *r.OriginatorCountry
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorIdentification() (OriginatorIdentification string) {
	if r != nil && r.OriginatorIdentification != nil {
		OriginatorIdentification = *r.OriginatorIdentification
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorPostalCode() (OriginatorPostalCode string) {
	if r != nil && r.OriginatorPostalCode != nil {
		OriginatorPostalCode = *r.OriginatorPostalCode
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorStreetAddress() (OriginatorStreetAddress string) {
	if r != nil && r.OriginatorStreetAddress != nil {
		OriginatorStreetAddress = *r.OriginatorStreetAddress
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorStateOrProvince() (OriginatorStateOrProvince string) {
	if r != nil && r.OriginatorStateOrProvince != nil {
		OriginatorStateOrProvince = *r.OriginatorStateOrProvince
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetPaymentRelatedInformation() (PaymentRelatedInformation string) {
	if r != nil && r.PaymentRelatedInformation != nil {
		PaymentRelatedInformation = *r.PaymentRelatedInformation
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetPaymentRelatedInformation2() (PaymentRelatedInformation2 string) {
	if r != nil && r.PaymentRelatedInformation2 != nil {
		PaymentRelatedInformation2 = *r.PaymentRelatedInformation2
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceiverIdentificationNumber() (ReceiverIdentificationNumber string) {
	if r != nil && r.ReceiverIdentificationNumber != nil {
		ReceiverIdentificationNumber = *r.ReceiverIdentificationNumber
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceiverStreetAddress() (ReceiverStreetAddress string) {
	if r != nil && r.ReceiverStreetAddress != nil {
		ReceiverStreetAddress = *r.ReceiverStreetAddress
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceiverCity() (ReceiverCity string) {
	if r != nil && r.ReceiverCity != nil {
		ReceiverCity = *r.ReceiverCity
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceiverStateOrProvince() (ReceiverStateOrProvince string) {
	if r != nil && r.ReceiverStateOrProvince != nil {
		ReceiverStateOrProvince = *r.ReceiverStateOrProvince
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceiverCountry() (ReceiverCountry string) {
	if r != nil && r.ReceiverCountry != nil {
		ReceiverCountry = *r.ReceiverCountry
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceiverPostalCode() (ReceiverPostalCode string) {
	if r != nil && r.ReceiverPostalCode != nil {
		ReceiverPostalCode = *r.ReceiverPostalCode
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceivingCompanyOrIndividualName() (ReceivingCompanyOrIndividualName string) {
	if r != nil && r.ReceivingCompanyOrIndividualName != nil {
		ReceivingCompanyOrIndividualName = *r.ReceivingCompanyOrIndividualName
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionName() (ReceivingDepositoryFinancialInstitutionName string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionName != nil {
		ReceivingDepositoryFinancialInstitutionName = *r.ReceivingDepositoryFinancialInstitutionName
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionIDQualifier() (ReceivingDepositoryFinancialInstitutionIDQualifier string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionIDQualifier != nil {
		ReceivingDepositoryFinancialInstitutionIDQualifier = *r.ReceivingDepositoryFinancialInstitutionIDQualifier
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionID() (ReceivingDepositoryFinancialInstitutionID string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionID != nil {
		ReceivingDepositoryFinancialInstitutionID = *r.ReceivingDepositoryFinancialInstitutionID
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionCountry() (ReceivingDepositoryFinancialInstitutionCountry string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionCountry != nil {
		ReceivingDepositoryFinancialInstitutionCountry = *r.ReceivingDepositoryFinancialInstitutionCountry
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

func (r ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline{Amount:%s ForeignExchangeIndicator:%s ForeignExchangeReferenceIndicator:%s ForeignExchangeReference:%s DestinationCountryCode:%s DestinationCurrencyCode:%s ForeignPaymentAmount:%s ForeignTraceNumber:%s InternationalTransactionTypeCode:%s OriginatingCurrencyCode:%s OriginatingDepositoryFinancialInstitutionName:%s OriginatingDepositoryFinancialInstitutionIDQualifier:%s OriginatingDepositoryFinancialInstitutionID:%s OriginatingDepositoryFinancialInstitutionBranchCountry:%s OriginatorCity:%s OriginatorCompanyEntryDescription:%s OriginatorCountry:%s OriginatorIdentification:%s OriginatorName:%s OriginatorPostalCode:%s OriginatorStreetAddress:%s OriginatorStateOrProvince:%s PaymentRelatedInformation:%s PaymentRelatedInformation2:%s ReceiverIdentificationNumber:%s ReceiverStreetAddress:%s ReceiverCity:%s ReceiverStateOrProvince:%s ReceiverCountry:%s ReceiverPostalCode:%s ReceivingCompanyOrIndividualName:%s ReceivingDepositoryFinancialInstitutionName:%s ReceivingDepositoryFinancialInstitutionIDQualifier:%s ReceivingDepositoryFinancialInstitutionID:%s ReceivingDepositoryFinancialInstitutionCountry:%s TraceNumber:%s}", core.FmtP(r.Amount), core.FmtP(r.ForeignExchangeIndicator), core.FmtP(r.ForeignExchangeReferenceIndicator), core.FmtP(r.ForeignExchangeReference), core.FmtP(r.DestinationCountryCode), core.FmtP(r.DestinationCurrencyCode), core.FmtP(r.ForeignPaymentAmount), core.FmtP(r.ForeignTraceNumber), core.FmtP(r.InternationalTransactionTypeCode), core.FmtP(r.OriginatingCurrencyCode), core.FmtP(r.OriginatingDepositoryFinancialInstitutionName), core.FmtP(r.OriginatingDepositoryFinancialInstitutionIDQualifier), core.FmtP(r.OriginatingDepositoryFinancialInstitutionID), core.FmtP(r.OriginatingDepositoryFinancialInstitutionBranchCountry), core.FmtP(r.OriginatorCity), core.FmtP(r.OriginatorCompanyEntryDescription), core.FmtP(r.OriginatorCountry), core.FmtP(r.OriginatorIdentification), core.FmtP(r.OriginatorName), core.FmtP(r.OriginatorPostalCode), core.FmtP(r.OriginatorStreetAddress), core.FmtP(r.OriginatorStateOrProvince), core.FmtP(r.PaymentRelatedInformation), core.FmtP(r.PaymentRelatedInformation2), core.FmtP(r.ReceiverIdentificationNumber), core.FmtP(r.ReceiverStreetAddress), core.FmtP(r.ReceiverCity), core.FmtP(r.ReceiverStateOrProvince), core.FmtP(r.ReceiverCountry), core.FmtP(r.ReceiverPostalCode), core.FmtP(r.ReceivingCompanyOrIndividualName), core.FmtP(r.ReceivingDepositoryFinancialInstitutionName), core.FmtP(r.ReceivingDepositoryFinancialInstitutionIDQualifier), core.FmtP(r.ReceivingDepositoryFinancialInstitutionID), core.FmtP(r.ReceivingDepositoryFinancialInstitutionCountry), core.FmtP(r.TraceNumber))
}

type ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency             *ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrency `pjson:"currency"`
	MerchantAcceptorID   *string                                                                 `pjson:"merchant_acceptor_id"`
	MerchantCity         *string                                                                 `pjson:"merchant_city"`
	MerchantCountry      *string                                                                 `pjson:"merchant_country"`
	MerchantDescriptor   *string                                                                 `pjson:"merchant_descriptor"`
	MerchantState        *string                                                                 `pjson:"merchant_state"`
	MerchantCategoryCode *string                                                                 `pjson:"merchant_category_code"`
	jsonFields           map[string]interface{}                                                  `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) GetCurrency() (Currency ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r != nil && r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

func (r ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.MerchantAcceptorID), core.FmtP(r.MerchantCity), core.FmtP(r.MerchantCountry), core.FmtP(r.MerchantDescriptor), core.FmtP(r.MerchantState), core.FmtP(r.MerchantCategoryCode))
}

type ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrency string

const (
	ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyCad ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "CAD"
	ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyChf ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "CHF"
	ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyEur ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "EUR"
	ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyGbp ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "GBP"
	ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyJpy ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "JPY"
	ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyUsd ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "USD"
)

type ACHTransferSimulationDeclinedTransactionType string

const (
	ACHTransferSimulationDeclinedTransactionTypeDeclinedTransaction ACHTransferSimulationDeclinedTransactionType = "declined_transaction"
)

type ACHTransferSimulationType string

const (
	ACHTransferSimulationTypeInboundACHTransferSimulationResult ACHTransferSimulationType = "inbound_ach_transfer_simulation_result"
)

type SimulateAnACHTransferToYourAccountParameters struct {
	// The identifier of the Account Number the inbound ACH Transfer is for.
	AccountNumberID *string `pjson:"account_number_id"`
	// The transfer amount in cents. A positive amount originates a credit transfer
	// pushing funds to the receiving account. A negative amount originates a debit
	// transfer pulling funds from the receiving account.
	Amount *int64 `pjson:"amount"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate *string `pjson:"company_descriptive_date"`
	// Data associated with the transfer set by the sender.
	CompanyDiscretionaryData *string `pjson:"company_discretionary_data"`
	// The description of the transfer set by the sender.
	CompanyEntryDescription *string `pjson:"company_entry_description"`
	// The name of the sender.
	CompanyName *string `pjson:"company_name"`
	// The sender's company id.
	CompanyID  *string                `pjson:"company_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// SimulateAnACHTransferToYourAccountParameters using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *SimulateAnACHTransferToYourAccountParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes SimulateAnACHTransferToYourAccountParameters into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *SimulateAnACHTransferToYourAccountParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the Account Number the inbound ACH Transfer is for.
func (r *SimulateAnACHTransferToYourAccountParameters) GetAccountNumberID() (AccountNumberID string) {
	if r != nil && r.AccountNumberID != nil {
		AccountNumberID = *r.AccountNumberID
	}
	return
}

// The transfer amount in cents. A positive amount originates a credit transfer
// pushing funds to the receiving account. A negative amount originates a debit
// transfer pulling funds from the receiving account.
func (r *SimulateAnACHTransferToYourAccountParameters) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description of the date of the transfer.
func (r *SimulateAnACHTransferToYourAccountParameters) GetCompanyDescriptiveDate() (CompanyDescriptiveDate string) {
	if r != nil && r.CompanyDescriptiveDate != nil {
		CompanyDescriptiveDate = *r.CompanyDescriptiveDate
	}
	return
}

// Data associated with the transfer set by the sender.
func (r *SimulateAnACHTransferToYourAccountParameters) GetCompanyDiscretionaryData() (CompanyDiscretionaryData string) {
	if r != nil && r.CompanyDiscretionaryData != nil {
		CompanyDiscretionaryData = *r.CompanyDiscretionaryData
	}
	return
}

// The description of the transfer set by the sender.
func (r *SimulateAnACHTransferToYourAccountParameters) GetCompanyEntryDescription() (CompanyEntryDescription string) {
	if r != nil && r.CompanyEntryDescription != nil {
		CompanyEntryDescription = *r.CompanyEntryDescription
	}
	return
}

// The name of the sender.
func (r *SimulateAnACHTransferToYourAccountParameters) GetCompanyName() (CompanyName string) {
	if r != nil && r.CompanyName != nil {
		CompanyName = *r.CompanyName
	}
	return
}

// The sender's company id.
func (r *SimulateAnACHTransferToYourAccountParameters) GetCompanyID() (CompanyID string) {
	if r != nil && r.CompanyID != nil {
		CompanyID = *r.CompanyID
	}
	return
}

func (r SimulateAnACHTransferToYourAccountParameters) String() (result string) {
	return fmt.Sprintf("&SimulateAnACHTransferToYourAccountParameters{AccountNumberID:%s Amount:%s CompanyDescriptiveDate:%s CompanyDiscretionaryData:%s CompanyEntryDescription:%s CompanyName:%s CompanyID:%s}", core.FmtP(r.AccountNumberID), core.FmtP(r.Amount), core.FmtP(r.CompanyDescriptiveDate), core.FmtP(r.CompanyDiscretionaryData), core.FmtP(r.CompanyEntryDescription), core.FmtP(r.CompanyName), core.FmtP(r.CompanyID))
}

type ReturnASandboxACHTransferParameters struct {
	// The reason why the Federal Reserve or destination bank returned this transfer.
	// Defaults to `no_account`.
	Reason     *ReturnASandboxACHTransferParametersReason `pjson:"reason"`
	jsonFields map[string]interface{}                     `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ReturnASandboxACHTransferParameters using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *ReturnASandboxACHTransferParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ReturnASandboxACHTransferParameters into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ReturnASandboxACHTransferParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The reason why the Federal Reserve or destination bank returned this transfer.
// Defaults to `no_account`.
func (r *ReturnASandboxACHTransferParameters) GetReason() (Reason ReturnASandboxACHTransferParametersReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

func (r ReturnASandboxACHTransferParameters) String() (result string) {
	return fmt.Sprintf("&ReturnASandboxACHTransferParameters{Reason:%s}", core.FmtP(r.Reason))
}

type ReturnASandboxACHTransferParametersReason string

const (
	ReturnASandboxACHTransferParametersReasonInsufficientFund                                          ReturnASandboxACHTransferParametersReason = "insufficient_fund"
	ReturnASandboxACHTransferParametersReasonNoAccount                                                 ReturnASandboxACHTransferParametersReason = "no_account"
	ReturnASandboxACHTransferParametersReasonAccountClosed                                             ReturnASandboxACHTransferParametersReason = "account_closed"
	ReturnASandboxACHTransferParametersReasonInvalidAccountNumberStructure                             ReturnASandboxACHTransferParametersReason = "invalid_account_number_structure"
	ReturnASandboxACHTransferParametersReasonAccountFrozenEntryReturnedPerOfacInstruction              ReturnASandboxACHTransferParametersReason = "account_frozen_entry_returned_per_ofac_instruction"
	ReturnASandboxACHTransferParametersReasonCreditEntryRefusedByReceiver                              ReturnASandboxACHTransferParametersReason = "credit_entry_refused_by_receiver"
	ReturnASandboxACHTransferParametersReasonUnauthorizedDebitToConsumerAccountUsingCorporateSecCode   ReturnASandboxACHTransferParametersReason = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	ReturnASandboxACHTransferParametersReasonCorporateCustomerAdvisedNotAuthorized                     ReturnASandboxACHTransferParametersReason = "corporate_customer_advised_not_authorized"
	ReturnASandboxACHTransferParametersReasonPaymentStopped                                            ReturnASandboxACHTransferParametersReason = "payment_stopped"
	ReturnASandboxACHTransferParametersReasonNonTransactionAccount                                     ReturnASandboxACHTransferParametersReason = "non_transaction_account"
	ReturnASandboxACHTransferParametersReasonUncollectedFunds                                          ReturnASandboxACHTransferParametersReason = "uncollected_funds"
	ReturnASandboxACHTransferParametersReasonRoutingNumberCheckDigitError                              ReturnASandboxACHTransferParametersReason = "routing_number_check_digit_error"
	ReturnASandboxACHTransferParametersReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete ReturnASandboxACHTransferParametersReason = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	ReturnASandboxACHTransferParametersReasonAmountFieldError                                          ReturnASandboxACHTransferParametersReason = "amount_field_error"
	ReturnASandboxACHTransferParametersReasonAuthorizationRevokedByCustomer                            ReturnASandboxACHTransferParametersReason = "authorization_revoked_by_customer"
	ReturnASandboxACHTransferParametersReasonInvalidACHRoutingNumber                                   ReturnASandboxACHTransferParametersReason = "invalid_ach_routing_number"
	ReturnASandboxACHTransferParametersReasonFileRecordEditCriteria                                    ReturnASandboxACHTransferParametersReason = "file_record_edit_criteria"
	ReturnASandboxACHTransferParametersReasonEnrInvalidIndividualName                                  ReturnASandboxACHTransferParametersReason = "enr_invalid_individual_name"
	ReturnASandboxACHTransferParametersReasonReturnedPerOdfiRequest                                    ReturnASandboxACHTransferParametersReason = "returned_per_odfi_request"
	ReturnASandboxACHTransferParametersReasonAddendaError                                              ReturnASandboxACHTransferParametersReason = "addenda_error"
	ReturnASandboxACHTransferParametersReasonLimitedParticipationDfi                                   ReturnASandboxACHTransferParametersReason = "limited_participation_dfi"
	ReturnASandboxACHTransferParametersReasonIncorrectlyCodedOutboundInternationalPayment              ReturnASandboxACHTransferParametersReason = "incorrectly_coded_outbound_international_payment"
	ReturnASandboxACHTransferParametersReasonOther                                                     ReturnASandboxACHTransferParametersReason = "other"
)
