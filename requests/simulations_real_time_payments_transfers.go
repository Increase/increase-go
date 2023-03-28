package requests

import (
	"fmt"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/fields"
)

type InboundRealTimePaymentsTransferSimulationResult struct {
	// If the Real Time Payments Transfer attempt succeeds, this will contain the
	// resulting [Transaction](#transactions) object. The Transaction's `source` will
	// be of `category: inbound_real_time_payments_transfer_confirmation`.
	Transaction fields.Field[InboundRealTimePaymentsTransferSimulationResultTransaction] `json:"transaction,required,nullable"`
	// If the Real Time Payments Transfer attempt fails, this will contain the
	// resulting [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of
	// `category: inbound_real_time_payments_transfer_decline`.
	DeclinedTransaction fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction] `json:"declined_transaction,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_real_time_payments_transfer_simulation_result`.
	Type fields.Field[InboundRealTimePaymentsTransferSimulationResultType] `json:"type,required"`
}

// MarshalJSON serializes InboundRealTimePaymentsTransferSimulationResult into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *InboundRealTimePaymentsTransferSimulationResult) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResult) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResult{Transaction:%s DeclinedTransaction:%s Type:%s}", r.Transaction, r.DeclinedTransaction, r.Type)
}

type InboundRealTimePaymentsTransferSimulationResultTransaction struct {
	// The identifier for the Account the Transaction belongs to.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transcation's
	// Account.
	Currency fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionCurrency] `json:"currency,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// For a Transaction related to a transfer, this is the description you provide.
	// For a Transaction related to a payment, this is the description the vendor
	// provides.
	Description fields.Field[string] `json:"description,required"`
	// The Transaction identifier.
	ID fields.Field[string] `json:"id,required"`
	// The identifier for the route this Transaction came through. Routes are things
	// like cards and ACH details.
	RouteID fields.Field[string] `json:"route_id,required,nullable"`
	// The type of the route this Transaction came through.
	RouteType fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionRouteType] `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Transaction. Note that for backwards compatibility reasons, additional
	// undocumented keys may appear in this object. These should be treated as
	// deprecated and will be removed in the future.
	Source fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSource] `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionType] `json:"type,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransaction into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransaction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransaction) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransaction{AccountID:%s Amount:%s Currency:%s CreatedAt:%s Description:%s ID:%s RouteID:%s RouteType:%s Source:%s Type:%s}", r.AccountID, r.Amount, r.Currency, r.CreatedAt, r.Description, r.ID, r.RouteID, r.RouteType, r.Source, r.Type)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionRouteType string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionRouteTypeAccountNumber InboundRealTimePaymentsTransferSimulationResultTransactionRouteType = "account_number"
	InboundRealTimePaymentsTransferSimulationResultTransactionRouteTypeCard          InboundRealTimePaymentsTransferSimulationResultTransactionRouteType = "card"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSource struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory] `json:"category,required"`
	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention] `json:"account_transfer_intention,required,nullable"`
	// A ACH Check Conversion Return object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_check_conversion_return`.
	ACHCheckConversionReturn fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversionReturn] `json:"ach_check_conversion_return,required,nullable"`
	// A ACH Check Conversion object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_check_conversion`.
	ACHCheckConversion fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversion] `json:"ach_check_conversion,required,nullable"`
	// A ACH Transfer Intention object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention] `json:"ach_transfer_intention,required,nullable"`
	// A ACH Transfer Rejection object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection] `json:"ach_transfer_rejection,required,nullable"`
	// A ACH Transfer Return object. This field will be present in the JSON response if
	// and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn] `json:"ach_transfer_return,required,nullable"`
	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptance] `json:"card_dispute_acceptance,required,nullable"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund] `json:"card_refund,required,nullable"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement] `json:"card_settlement,required,nullable"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptance] `json:"check_deposit_acceptance,required,nullable"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn] `json:"check_deposit_return,required,nullable"`
	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention] `json:"check_transfer_intention,required,nullable"`
	// A Check Transfer Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_transfer_return`.
	CheckTransferReturn fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturn] `json:"check_transfer_return,required,nullable"`
	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejection] `json:"check_transfer_rejection,required,nullable"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest] `json:"check_transfer_stop_payment_request,required,nullable"`
	// A Dispute Resolution object. This field will be present in the JSON response if
	// and only if `category` is equal to `dispute_resolution`.
	DisputeResolution fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolution] `json:"dispute_resolution,required,nullable"`
	// A Empyreal Cash Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `empyreal_cash_deposit`.
	EmpyrealCashDeposit fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceEmpyrealCashDeposit] `json:"empyreal_cash_deposit,required,nullable"`
	// A Inbound ACH Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer] `json:"inbound_ach_transfer,required,nullable"`
	// A Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck] `json:"inbound_check,required,nullable"`
	// A Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer] `json:"inbound_international_ach_transfer,required,nullable"`
	// A Inbound Real Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation] `json:"inbound_real_time_payments_transfer_confirmation,required,nullable"`
	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal] `json:"inbound_wire_drawdown_payment_reversal,required,nullable"`
	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment] `json:"inbound_wire_drawdown_payment,required,nullable"`
	// A Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal] `json:"inbound_wire_reversal,required,nullable"`
	// A Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer] `json:"inbound_wire_transfer,required,nullable"`
	// A Interest Payment object. This field will be present in the JSON response if
	// and only if `category` is equal to `interest_payment`.
	InterestPayment fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPayment] `json:"interest_payment,required,nullable"`
	// A Internal Source object. This field will be present in the JSON response if and
	// only if `category` is equal to `internal_source`.
	InternalSource fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource] `json:"internal_source,required,nullable"`
	// A Deprecated Card Refund object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_route_refund`.
	CardRouteRefund fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefund] `json:"card_route_refund,required,nullable"`
	// A Deprecated Card Settlement object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_settlement`.
	CardRouteSettlement fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlement] `json:"card_route_settlement,required,nullable"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFunds] `json:"sample_funds,required,nullable"`
	// A Wire Drawdown Payment Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_intention`.
	WireDrawdownPaymentIntention fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentIntention] `json:"wire_drawdown_payment_intention,required,nullable"`
	// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_rejection`.
	WireDrawdownPaymentRejection fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentRejection] `json:"wire_drawdown_payment_rejection,required,nullable"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention] `json:"wire_transfer_intention,required,nullable"`
	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejection] `json:"wire_transfer_rejection,required,nullable"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSource into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSource) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSource{Category:%s AccountTransferIntention:%s ACHCheckConversionReturn:%s ACHCheckConversion:%s ACHTransferIntention:%s ACHTransferRejection:%s ACHTransferReturn:%s CardDisputeAcceptance:%s CardRefund:%s CardSettlement:%s CheckDepositAcceptance:%s CheckDepositReturn:%s CheckTransferIntention:%s CheckTransferReturn:%s CheckTransferRejection:%s CheckTransferStopPaymentRequest:%s DisputeResolution:%s EmpyrealCashDeposit:%s InboundACHTransfer:%s InboundCheck:%s InboundInternationalACHTransfer:%s InboundRealTimePaymentsTransferConfirmation:%s InboundWireDrawdownPaymentReversal:%s InboundWireDrawdownPayment:%s InboundWireReversal:%s InboundWireTransfer:%s InterestPayment:%s InternalSource:%s CardRouteRefund:%s CardRouteSettlement:%s SampleFunds:%s WireDrawdownPaymentIntention:%s WireDrawdownPaymentRejection:%s WireTransferIntention:%s WireTransferRejection:%s}", r.Category, r.AccountTransferIntention, r.ACHCheckConversionReturn, r.ACHCheckConversion, r.ACHTransferIntention, r.ACHTransferRejection, r.ACHTransferReturn, r.CardDisputeAcceptance, r.CardRefund, r.CardSettlement, r.CheckDepositAcceptance, r.CheckDepositReturn, r.CheckTransferIntention, r.CheckTransferReturn, r.CheckTransferRejection, r.CheckTransferStopPaymentRequest, r.DisputeResolution, r.EmpyrealCashDeposit, r.InboundACHTransfer, r.InboundCheck, r.InboundInternationalACHTransfer, r.InboundRealTimePaymentsTransferConfirmation, r.InboundWireDrawdownPaymentReversal, r.InboundWireDrawdownPayment, r.InboundWireReversal, r.InboundWireTransfer, r.InterestPayment, r.InternalSource, r.CardRouteRefund, r.CardRouteSettlement, r.SampleFunds, r.WireDrawdownPaymentIntention, r.WireDrawdownPaymentRejection, r.WireTransferIntention, r.WireTransferRejection)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryAccountTransferIntention                    InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "account_transfer_intention"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryACHCheckConversionReturn                    InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "ach_check_conversion_return"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryACHCheckConversion                          InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "ach_check_conversion"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryACHTransferIntention                        InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "ach_transfer_intention"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryACHTransferRejection                        InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "ach_transfer_rejection"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryACHTransferReturn                           InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "ach_transfer_return"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCardDisputeAcceptance                       InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "card_dispute_acceptance"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCardRefund                                  InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "card_refund"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCardSettlement                              InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "card_settlement"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckDepositAcceptance                      InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_deposit_acceptance"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckDepositReturn                          InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_deposit_return"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckTransferIntention                      InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_transfer_intention"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckTransferReturn                         InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_transfer_return"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckTransferRejection                      InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_transfer_rejection"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckTransferStopPaymentRequest             InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_transfer_stop_payment_request"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryDisputeResolution                           InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "dispute_resolution"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryEmpyrealCashDeposit                         InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "empyreal_cash_deposit"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundACHTransfer                          InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_ach_transfer"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundACHTransferReturnIntention           InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_ach_transfer_return_intention"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundCheck                                InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_check"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundInternationalACHTransfer             InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_international_ach_transfer"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_real_time_payments_transfer_confirmation"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundWireDrawdownPaymentReversal          InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_wire_drawdown_payment_reversal"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundWireDrawdownPayment                  InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_wire_drawdown_payment"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundWireReversal                         InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_wire_reversal"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundWireTransfer                         InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_wire_transfer"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInterestPayment                             InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "interest_payment"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInternalGeneralLedgerTransaction            InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "internal_general_ledger_transaction"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInternalSource                              InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "internal_source"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCardRouteRefund                             InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "card_route_refund"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCardRouteSettlement                         InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "card_route_settlement"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryRealTimePaymentsTransferAcknowledgement     InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "real_time_payments_transfer_acknowledgement"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategorySampleFunds                                 InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "sample_funds"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryWireDrawdownPaymentIntention                InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "wire_drawdown_payment_intention"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryWireDrawdownPaymentRejection                InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "wire_drawdown_payment_rejection"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryWireTransferIntention                       InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "wire_transfer_intention"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryWireTransferRejection                       InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "wire_transfer_rejection"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryOther                                       InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "other"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency] `json:"currency,required"`
	// The description you chose to give the transfer.
	Description fields.Field[string] `json:"description,required"`
	// The identifier of the Account to where the Account Transfer was sent.
	DestinationAccountID fields.Field[string] `json:"destination_account_id,required"`
	// The identifier of the Account from where the Account Transfer was sent.
	SourceAccountID fields.Field[string] `json:"source_account_id,required"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention{Amount:%s Currency:%s Description:%s DestinationAccountID:%s SourceAccountID:%s TransferID:%s}", r.Amount, r.Currency, r.Description, r.DestinationAccountID, r.SourceAccountID, r.TransferID)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversionReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// Why the transfer was returned.
	ReturnReasonCode fields.Field[string] `json:"return_reason_code,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversionReturn
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversionReturn) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversionReturn) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversionReturn{Amount:%s ReturnReasonCode:%s}", r.Amount, r.ReturnReasonCode)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversion struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The identifier of the File containing an image of the returned check.
	FileID fields.Field[string] `json:"file_id,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversion
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversion) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversion) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversion{Amount:%s FileID:%s}", r.Amount, r.FileID)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount              fields.Field[int64]  `json:"amount,required"`
	AccountNumber       fields.Field[string] `json:"account_number,required"`
	RoutingNumber       fields.Field[string] `json:"routing_number,required"`
	StatementDescriptor fields.Field[string] `json:"statement_descriptor,required"`
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention{Amount:%s AccountNumber:%s RoutingNumber:%s StatementDescriptor:%s TransferID:%s}", r.Amount, r.AccountNumber, r.RoutingNumber, r.StatementDescriptor, r.TransferID)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection{TransferID:%s}", r.TransferID)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode] `json:"return_reason_code,required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID fields.Field[string] `json:"transfer_id,required"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID fields.Field[string] `json:"transaction_id,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn{CreatedAt:%s ReturnReasonCode:%s TransferID:%s TransactionID:%s}", r.CreatedAt, r.ReturnReasonCode, r.TransferID, r.TransactionID)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInsufficientFund                                          InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "insufficient_fund"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeNoAccount                                                 InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "no_account"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAccountClosed                                             InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "account_closed"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure                             InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction              InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver                              InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode   InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized                     InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodePaymentStopped                                            InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "payment_stopped"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeNonTransactionAccount                                     InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "non_transaction_account"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeUncollectedFunds                                          InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "uncollected_funds"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError                              InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAmountFieldError                                          InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "amount_field_error"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer                            InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber                                   InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeFileRecordEditCriteria                                    InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualName                                  InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest                                    InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAddendaError                                              InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "addenda_error"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeLimitedParticipationDfi                                   InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment              InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "incorrectly_coded_outbound_international_payment"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeOther                                                     InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "other"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptance struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was accepted.
	AcceptedAt fields.Field[time.Time] `json:"accepted_at,required" format:"date-time"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID fields.Field[string] `json:"card_dispute_id,required"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID fields.Field[string] `json:"transaction_id,required,nullable"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptance
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptance) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptance) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptance{AcceptedAt:%s CardDisputeID:%s TransactionID:%s}", r.AcceptedAt, r.CardDisputeID, r.TransactionID)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency] `json:"currency,required"`
	// The identifier for the Transaction this refunds, if any.
	CardSettlementTransactionID fields.Field[string] `json:"card_settlement_transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundType] `json:"type,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund{Amount:%s Currency:%s CardSettlementTransactionID:%s Type:%s}", r.Amount, r.Currency, r.CardSettlementTransactionID, r.Type)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundType string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundTypeCardRefund InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundType = "card_refund"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement struct {
	// The amount in the minor unit of the transaction's settlement currency. For
	// dollars, for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's settlement currency.
	Currency fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency] `json:"currency,required"`
	// The amount in the minor unit of the transaction's presentment currency.
	PresentmentAmount fields.Field[int64] `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's presentment currency.
	PresentmentCurrency  fields.Field[string] `json:"presentment_currency,required"`
	MerchantCity         fields.Field[string] `json:"merchant_city,required,nullable"`
	MerchantCountry      fields.Field[string] `json:"merchant_country,required"`
	MerchantName         fields.Field[string] `json:"merchant_name,required,nullable"`
	MerchantCategoryCode fields.Field[string] `json:"merchant_category_code,required"`
	MerchantState        fields.Field[string] `json:"merchant_state,required,nullable"`
	// The identifier of the Pending Transaction associated with this Transaction.
	PendingTransactionID fields.Field[string] `json:"pending_transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_settlement`.
	Type fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementType] `json:"type,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement{Amount:%s Currency:%s PresentmentAmount:%s PresentmentCurrency:%s MerchantCity:%s MerchantCountry:%s MerchantName:%s MerchantCategoryCode:%s MerchantState:%s PendingTransactionID:%s Type:%s}", r.Amount, r.Currency, r.PresentmentAmount, r.PresentmentCurrency, r.MerchantCity, r.MerchantCountry, r.MerchantName, r.MerchantCategoryCode, r.MerchantState, r.PendingTransactionID, r.Type)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementType string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementTypeCardSettlement InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementType = "card_settlement"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptance struct {
	// The amount to be deposited in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency] `json:"currency,required"`
	// The account number printed on the check.
	AccountNumber fields.Field[string] `json:"account_number,required"`
	// The routing number printed on the check.
	RoutingNumber fields.Field[string] `json:"routing_number,required"`
	// An additional line of metadata printed on the check. This typically includes the
	// check number for business checks.
	AuxiliaryOnUs fields.Field[string] `json:"auxiliary_on_us,required,nullable"`
	// The check serial number, if present, for consumer checks. For business checks,
	// the serial number is usually in the `auxiliary_on_us` field.
	SerialNumber fields.Field[string] `json:"serial_number,required,nullable"`
	// The ID of the Check Deposit that was accepted.
	CheckDepositID fields.Field[string] `json:"check_deposit_id,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptance
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptance) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptance) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptance{Amount:%s Currency:%s AccountNumber:%s RoutingNumber:%s AuxiliaryOnUs:%s SerialNumber:%s CheckDepositID:%s}", r.Amount, r.Currency, r.AccountNumber, r.RoutingNumber, r.AuxiliaryOnUs, r.SerialNumber, r.CheckDepositID)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt fields.Field[time.Time] `json:"returned_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency] `json:"currency,required"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID fields.Field[string] `json:"check_deposit_id,required"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID fields.Field[string]                                                                                         `json:"transaction_id,required"`
	ReturnReason  fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason] `json:"return_reason,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn{Amount:%s ReturnedAt:%s Currency:%s CheckDepositID:%s TransactionID:%s ReturnReason:%s}", r.Amount, r.ReturnedAt, r.Currency, r.CheckDepositID, r.TransactionID, r.ReturnReason)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonACHConversionNotSupported InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "ach_conversion_not_supported"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonClosedAccount             InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "closed_account"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonDuplicateSubmission       InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "duplicate_submission"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonInsufficientFunds         InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "insufficient_funds"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonNoAccount                 InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "no_account"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonNotAuthorized             InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "not_authorized"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonStaleDated                InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "stale_dated"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonStopPayment               InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "stop_payment"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonUnknownReason             InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "unknown_reason"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonUnmatchedDetails          InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "unmatched_details"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonUnreadableImage           InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "unreadable_image"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention struct {
	// The street address of the check's destination.
	AddressLine1 fields.Field[string] `json:"address_line1,required"`
	// The second line of the address of the check's destination.
	AddressLine2 fields.Field[string] `json:"address_line2,required,nullable"`
	// The city of the check's destination.
	AddressCity fields.Field[string] `json:"address_city,required"`
	// The state of the check's destination.
	AddressState fields.Field[string] `json:"address_state,required"`
	// The postal code of the check's destination.
	AddressZip fields.Field[string] `json:"address_zip,required"`
	// The transfer amount in USD cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency] `json:"currency,required"`
	// The name that will be printed on the check.
	RecipientName fields.Field[string] `json:"recipient_name,required"`
	// The identifier of the Check Transfer with which this is associated.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention{AddressLine1:%s AddressLine2:%s AddressCity:%s AddressState:%s AddressZip:%s Amount:%s Currency:%s RecipientName:%s TransferID:%s}", r.AddressLine1, r.AddressLine2, r.AddressCity, r.AddressState, r.AddressZip, r.Amount, r.Currency, r.RecipientName, r.TransferID)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturn struct {
	// The identifier of the returned Check Transfer.
	TransferID fields.Field[string] `json:"transfer_id,required"`
	// If available, a document with additional information about the return.
	FileID fields.Field[string] `json:"file_id,required,nullable"`
	// The reason why the check was returned.
	Reason fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnReason] `json:"reason,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturn
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturn) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturn) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturn{TransferID:%s FileID:%s Reason:%s}", r.TransferID, r.FileID, r.Reason)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnReason string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnReasonMailDeliveryFailure InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnReason = "mail_delivery_failure"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnReasonRefusedByRecipient  InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnReason = "refused_by_recipient"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejection struct {
	// The identifier of the Check Transfer that led to this Transaction.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejection
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejection) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejection{TransferID:%s}", r.TransferID)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferID fields.Field[string] `json:"transfer_id,required"`
	// The transaction ID of the corresponding credit transaction.
	TransactionID fields.Field[string] `json:"transaction_id,required"`
	// The time the stop-payment was requested.
	RequestedAt fields.Field[time.Time] `json:"requested_at,required" format:"date-time"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestType] `json:"type,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest{TransferID:%s TransactionID:%s RequestedAt:%s Type:%s}", r.TransferID, r.TransactionID, r.RequestedAt, r.Type)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestType string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolution struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrency] `json:"currency,required"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionID fields.Field[string] `json:"disputed_transaction_id,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolution
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolution) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolution) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolution{Amount:%s Currency:%s DisputedTransactionID:%s}", r.Amount, r.Currency, r.DisputedTransactionID)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceEmpyrealCashDeposit struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount      fields.Field[int64]     `json:"amount,required"`
	BagID       fields.Field[string]    `json:"bag_id,required"`
	DepositDate fields.Field[time.Time] `json:"deposit_date,required" format:"date-time"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceEmpyrealCashDeposit
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceEmpyrealCashDeposit) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceEmpyrealCashDeposit) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceEmpyrealCashDeposit{Amount:%s BagID:%s DepositDate:%s}", r.Amount, r.BagID, r.DepositDate)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount                             fields.Field[int64]  `json:"amount,required"`
	OriginatorCompanyName              fields.Field[string] `json:"originator_company_name,required"`
	OriginatorCompanyDescriptiveDate   fields.Field[string] `json:"originator_company_descriptive_date,required,nullable"`
	OriginatorCompanyDiscretionaryData fields.Field[string] `json:"originator_company_discretionary_data,required,nullable"`
	OriginatorCompanyEntryDescription  fields.Field[string] `json:"originator_company_entry_description,required"`
	OriginatorCompanyID                fields.Field[string] `json:"originator_company_id,required"`
	ReceiverIDNumber                   fields.Field[string] `json:"receiver_id_number,required,nullable"`
	ReceiverName                       fields.Field[string] `json:"receiver_name,required,nullable"`
	TraceNumber                        fields.Field[string] `json:"trace_number,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer{Amount:%s OriginatorCompanyName:%s OriginatorCompanyDescriptiveDate:%s OriginatorCompanyDiscretionaryData:%s OriginatorCompanyEntryDescription:%s OriginatorCompanyID:%s ReceiverIDNumber:%s ReceiverName:%s TraceNumber:%s}", r.Amount, r.OriginatorCompanyName, r.OriginatorCompanyDescriptiveDate, r.OriginatorCompanyDiscretionaryData, r.OriginatorCompanyEntryDescription, r.OriginatorCompanyID, r.ReceiverIDNumber, r.ReceiverName, r.TraceNumber)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency              fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency] `json:"currency,required"`
	CheckNumber           fields.Field[string]                                                                               `json:"check_number,required,nullable"`
	CheckFrontImageFileID fields.Field[string]                                                                               `json:"check_front_image_file_id,required,nullable"`
	CheckRearImageFileID  fields.Field[string]                                                                               `json:"check_rear_image_file_id,required,nullable"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck{Amount:%s Currency:%s CheckNumber:%s CheckFrontImageFileID:%s CheckRearImageFileID:%s}", r.Amount, r.Currency, r.CheckNumber, r.CheckFrontImageFileID, r.CheckRearImageFileID)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount                                                 fields.Field[int64]  `json:"amount,required"`
	ForeignExchangeIndicator                               fields.Field[string] `json:"foreign_exchange_indicator,required"`
	ForeignExchangeReferenceIndicator                      fields.Field[string] `json:"foreign_exchange_reference_indicator,required"`
	ForeignExchangeReference                               fields.Field[string] `json:"foreign_exchange_reference,required,nullable"`
	DestinationCountryCode                                 fields.Field[string] `json:"destination_country_code,required"`
	DestinationCurrencyCode                                fields.Field[string] `json:"destination_currency_code,required"`
	ForeignPaymentAmount                                   fields.Field[int64]  `json:"foreign_payment_amount,required"`
	ForeignTraceNumber                                     fields.Field[string] `json:"foreign_trace_number,required,nullable"`
	InternationalTransactionTypeCode                       fields.Field[string] `json:"international_transaction_type_code,required"`
	OriginatingCurrencyCode                                fields.Field[string] `json:"originating_currency_code,required"`
	OriginatingDepositoryFinancialInstitutionName          fields.Field[string] `json:"originating_depository_financial_institution_name,required"`
	OriginatingDepositoryFinancialInstitutionIDQualifier   fields.Field[string] `json:"originating_depository_financial_institution_id_qualifier,required"`
	OriginatingDepositoryFinancialInstitutionID            fields.Field[string] `json:"originating_depository_financial_institution_id,required"`
	OriginatingDepositoryFinancialInstitutionBranchCountry fields.Field[string] `json:"originating_depository_financial_institution_branch_country,required"`
	OriginatorCity                                         fields.Field[string] `json:"originator_city,required"`
	OriginatorCompanyEntryDescription                      fields.Field[string] `json:"originator_company_entry_description,required"`
	OriginatorCountry                                      fields.Field[string] `json:"originator_country,required"`
	OriginatorIdentification                               fields.Field[string] `json:"originator_identification,required"`
	OriginatorName                                         fields.Field[string] `json:"originator_name,required"`
	OriginatorPostalCode                                   fields.Field[string] `json:"originator_postal_code,required,nullable"`
	OriginatorStreetAddress                                fields.Field[string] `json:"originator_street_address,required"`
	OriginatorStateOrProvince                              fields.Field[string] `json:"originator_state_or_province,required,nullable"`
	PaymentRelatedInformation                              fields.Field[string] `json:"payment_related_information,required,nullable"`
	PaymentRelatedInformation2                             fields.Field[string] `json:"payment_related_information2,required,nullable"`
	ReceiverIdentificationNumber                           fields.Field[string] `json:"receiver_identification_number,required,nullable"`
	ReceiverStreetAddress                                  fields.Field[string] `json:"receiver_street_address,required"`
	ReceiverCity                                           fields.Field[string] `json:"receiver_city,required"`
	ReceiverStateOrProvince                                fields.Field[string] `json:"receiver_state_or_province,required,nullable"`
	ReceiverCountry                                        fields.Field[string] `json:"receiver_country,required"`
	ReceiverPostalCode                                     fields.Field[string] `json:"receiver_postal_code,required,nullable"`
	ReceivingCompanyOrIndividualName                       fields.Field[string] `json:"receiving_company_or_individual_name,required"`
	ReceivingDepositoryFinancialInstitutionName            fields.Field[string] `json:"receiving_depository_financial_institution_name,required"`
	ReceivingDepositoryFinancialInstitutionIDQualifier     fields.Field[string] `json:"receiving_depository_financial_institution_id_qualifier,required"`
	ReceivingDepositoryFinancialInstitutionID              fields.Field[string] `json:"receiving_depository_financial_institution_id,required"`
	ReceivingDepositoryFinancialInstitutionCountry         fields.Field[string] `json:"receiving_depository_financial_institution_country,required"`
	TraceNumber                                            fields.Field[string] `json:"trace_number,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer{Amount:%s ForeignExchangeIndicator:%s ForeignExchangeReferenceIndicator:%s ForeignExchangeReference:%s DestinationCountryCode:%s DestinationCurrencyCode:%s ForeignPaymentAmount:%s ForeignTraceNumber:%s InternationalTransactionTypeCode:%s OriginatingCurrencyCode:%s OriginatingDepositoryFinancialInstitutionName:%s OriginatingDepositoryFinancialInstitutionIDQualifier:%s OriginatingDepositoryFinancialInstitutionID:%s OriginatingDepositoryFinancialInstitutionBranchCountry:%s OriginatorCity:%s OriginatorCompanyEntryDescription:%s OriginatorCountry:%s OriginatorIdentification:%s OriginatorName:%s OriginatorPostalCode:%s OriginatorStreetAddress:%s OriginatorStateOrProvince:%s PaymentRelatedInformation:%s PaymentRelatedInformation2:%s ReceiverIdentificationNumber:%s ReceiverStreetAddress:%s ReceiverCity:%s ReceiverStateOrProvince:%s ReceiverCountry:%s ReceiverPostalCode:%s ReceivingCompanyOrIndividualName:%s ReceivingDepositoryFinancialInstitutionName:%s ReceivingDepositoryFinancialInstitutionIDQualifier:%s ReceivingDepositoryFinancialInstitutionID:%s ReceivingDepositoryFinancialInstitutionCountry:%s TraceNumber:%s}", r.Amount, r.ForeignExchangeIndicator, r.ForeignExchangeReferenceIndicator, r.ForeignExchangeReference, r.DestinationCountryCode, r.DestinationCurrencyCode, r.ForeignPaymentAmount, r.ForeignTraceNumber, r.InternationalTransactionTypeCode, r.OriginatingCurrencyCode, r.OriginatingDepositoryFinancialInstitutionName, r.OriginatingDepositoryFinancialInstitutionIDQualifier, r.OriginatingDepositoryFinancialInstitutionID, r.OriginatingDepositoryFinancialInstitutionBranchCountry, r.OriginatorCity, r.OriginatorCompanyEntryDescription, r.OriginatorCountry, r.OriginatorIdentification, r.OriginatorName, r.OriginatorPostalCode, r.OriginatorStreetAddress, r.OriginatorStateOrProvince, r.PaymentRelatedInformation, r.PaymentRelatedInformation2, r.ReceiverIdentificationNumber, r.ReceiverStreetAddress, r.ReceiverCity, r.ReceiverStateOrProvince, r.ReceiverCountry, r.ReceiverPostalCode, r.ReceivingCompanyOrIndividualName, r.ReceivingDepositoryFinancialInstitutionName, r.ReceivingDepositoryFinancialInstitutionIDQualifier, r.ReceivingDepositoryFinancialInstitutionID, r.ReceivingDepositoryFinancialInstitutionCountry, r.TraceNumber)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real Time Payments transfer.
	Currency fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency] `json:"currency,required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName fields.Field[string] `json:"creditor_name,required"`
	// The name provided by the sender of the transfer.
	DebtorName fields.Field[string] `json:"debtor_name,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber fields.Field[string] `json:"debtor_account_number,required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber fields.Field[string] `json:"debtor_routing_number,required"`
	// The Real Time Payments network identification of the transfer
	TransactionIdentification fields.Field[string] `json:"transaction_identification,required"`
	// Additional information included with the transfer.
	RemittanceInformation fields.Field[string] `json:"remittance_information,required,nullable"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation{Amount:%s Currency:%s CreditorName:%s DebtorName:%s DebtorAccountNumber:%s DebtorRoutingNumber:%s TransactionIdentification:%s RemittanceInformation:%s}", r.Amount, r.Currency, r.CreditorName, r.DebtorName, r.DebtorAccountNumber, r.DebtorRoutingNumber, r.TransactionIdentification, r.RemittanceInformation)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal struct {
	// The amount that was reversed.
	Amount fields.Field[int64] `json:"amount,required"`
	// The description on the reversal message from Fedwire.
	Description fields.Field[string] `json:"description,required"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate fields.Field[time.Time] `json:"input_cycle_date,required" format:"date"`
	// The Fedwire sequence number.
	InputSequenceNumber fields.Field[string] `json:"input_sequence_number,required"`
	// The Fedwire input source identifier.
	InputSource fields.Field[string] `json:"input_source,required"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData fields.Field[string] `json:"input_message_accountability_data,required"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData fields.Field[string] `json:"previous_message_input_message_accountability_data,required"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate fields.Field[time.Time] `json:"previous_message_input_cycle_date,required" format:"date"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber fields.Field[string] `json:"previous_message_input_sequence_number,required"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource fields.Field[string] `json:"previous_message_input_source,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal{Amount:%s Description:%s InputCycleDate:%s InputSequenceNumber:%s InputSource:%s InputMessageAccountabilityData:%s PreviousMessageInputMessageAccountabilityData:%s PreviousMessageInputCycleDate:%s PreviousMessageInputSequenceNumber:%s PreviousMessageInputSource:%s}", r.Amount, r.Description, r.InputCycleDate, r.InputSequenceNumber, r.InputSource, r.InputMessageAccountabilityData, r.PreviousMessageInputMessageAccountabilityData, r.PreviousMessageInputCycleDate, r.PreviousMessageInputSequenceNumber, r.PreviousMessageInputSource)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount                             fields.Field[int64]  `json:"amount,required"`
	BeneficiaryAddressLine1            fields.Field[string] `json:"beneficiary_address_line1,required,nullable"`
	BeneficiaryAddressLine2            fields.Field[string] `json:"beneficiary_address_line2,required,nullable"`
	BeneficiaryAddressLine3            fields.Field[string] `json:"beneficiary_address_line3,required,nullable"`
	BeneficiaryName                    fields.Field[string] `json:"beneficiary_name,required,nullable"`
	BeneficiaryReference               fields.Field[string] `json:"beneficiary_reference,required,nullable"`
	Description                        fields.Field[string] `json:"description,required"`
	InputMessageAccountabilityData     fields.Field[string] `json:"input_message_accountability_data,required,nullable"`
	OriginatorAddressLine1             fields.Field[string] `json:"originator_address_line1,required,nullable"`
	OriginatorAddressLine2             fields.Field[string] `json:"originator_address_line2,required,nullable"`
	OriginatorAddressLine3             fields.Field[string] `json:"originator_address_line3,required,nullable"`
	OriginatorName                     fields.Field[string] `json:"originator_name,required,nullable"`
	OriginatorToBeneficiaryInformation fields.Field[string] `json:"originator_to_beneficiary_information,required,nullable"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment{Amount:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s BeneficiaryName:%s BeneficiaryReference:%s Description:%s InputMessageAccountabilityData:%s OriginatorAddressLine1:%s OriginatorAddressLine2:%s OriginatorAddressLine3:%s OriginatorName:%s OriginatorToBeneficiaryInformation:%s}", r.Amount, r.BeneficiaryAddressLine1, r.BeneficiaryAddressLine2, r.BeneficiaryAddressLine3, r.BeneficiaryName, r.BeneficiaryReference, r.Description, r.InputMessageAccountabilityData, r.OriginatorAddressLine1, r.OriginatorAddressLine2, r.OriginatorAddressLine3, r.OriginatorName, r.OriginatorToBeneficiaryInformation)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal struct {
	// The amount that was reversed.
	Amount fields.Field[int64] `json:"amount,required"`
	// The description on the reversal message from Fedwire.
	Description fields.Field[string] `json:"description,required"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate fields.Field[time.Time] `json:"input_cycle_date,required" format:"date"`
	// The Fedwire sequence number.
	InputSequenceNumber fields.Field[string] `json:"input_sequence_number,required"`
	// The Fedwire input source identifier.
	InputSource fields.Field[string] `json:"input_source,required"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData fields.Field[string] `json:"input_message_accountability_data,required"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData fields.Field[string] `json:"previous_message_input_message_accountability_data,required"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate fields.Field[time.Time] `json:"previous_message_input_cycle_date,required" format:"date"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber fields.Field[string] `json:"previous_message_input_sequence_number,required"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource fields.Field[string] `json:"previous_message_input_source,required"`
	// Information included in the wire reversal for the receiving financial
	// institution.
	ReceiverFinancialInstitutionInformation fields.Field[string] `json:"receiver_financial_institution_information,required,nullable"`
	// Additional financial institution information included in the wire reversal.
	FinancialInstitutionToFinancialInstitutionInformation fields.Field[string] `json:"financial_institution_to_financial_institution_information,required,nullable"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal{Amount:%s Description:%s InputCycleDate:%s InputSequenceNumber:%s InputSource:%s InputMessageAccountabilityData:%s PreviousMessageInputMessageAccountabilityData:%s PreviousMessageInputCycleDate:%s PreviousMessageInputSequenceNumber:%s PreviousMessageInputSource:%s ReceiverFinancialInstitutionInformation:%s FinancialInstitutionToFinancialInstitutionInformation:%s}", r.Amount, r.Description, r.InputCycleDate, r.InputSequenceNumber, r.InputSource, r.InputMessageAccountabilityData, r.PreviousMessageInputMessageAccountabilityData, r.PreviousMessageInputCycleDate, r.PreviousMessageInputSequenceNumber, r.PreviousMessageInputSource, r.ReceiverFinancialInstitutionInformation, r.FinancialInstitutionToFinancialInstitutionInformation)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount                                  fields.Field[int64]  `json:"amount,required"`
	BeneficiaryAddressLine1                 fields.Field[string] `json:"beneficiary_address_line1,required,nullable"`
	BeneficiaryAddressLine2                 fields.Field[string] `json:"beneficiary_address_line2,required,nullable"`
	BeneficiaryAddressLine3                 fields.Field[string] `json:"beneficiary_address_line3,required,nullable"`
	BeneficiaryName                         fields.Field[string] `json:"beneficiary_name,required,nullable"`
	BeneficiaryReference                    fields.Field[string] `json:"beneficiary_reference,required,nullable"`
	Description                             fields.Field[string] `json:"description,required"`
	InputMessageAccountabilityData          fields.Field[string] `json:"input_message_accountability_data,required,nullable"`
	OriginatorAddressLine1                  fields.Field[string] `json:"originator_address_line1,required,nullable"`
	OriginatorAddressLine2                  fields.Field[string] `json:"originator_address_line2,required,nullable"`
	OriginatorAddressLine3                  fields.Field[string] `json:"originator_address_line3,required,nullable"`
	OriginatorName                          fields.Field[string] `json:"originator_name,required,nullable"`
	OriginatorToBeneficiaryInformationLine1 fields.Field[string] `json:"originator_to_beneficiary_information_line1,required,nullable"`
	OriginatorToBeneficiaryInformationLine2 fields.Field[string] `json:"originator_to_beneficiary_information_line2,required,nullable"`
	OriginatorToBeneficiaryInformationLine3 fields.Field[string] `json:"originator_to_beneficiary_information_line3,required,nullable"`
	OriginatorToBeneficiaryInformationLine4 fields.Field[string] `json:"originator_to_beneficiary_information_line4,required,nullable"`
	OriginatorToBeneficiaryInformation      fields.Field[string] `json:"originator_to_beneficiary_information,required,nullable"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer{Amount:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s BeneficiaryName:%s BeneficiaryReference:%s Description:%s InputMessageAccountabilityData:%s OriginatorAddressLine1:%s OriginatorAddressLine2:%s OriginatorAddressLine3:%s OriginatorName:%s OriginatorToBeneficiaryInformationLine1:%s OriginatorToBeneficiaryInformationLine2:%s OriginatorToBeneficiaryInformationLine3:%s OriginatorToBeneficiaryInformationLine4:%s OriginatorToBeneficiaryInformation:%s}", r.Amount, r.BeneficiaryAddressLine1, r.BeneficiaryAddressLine2, r.BeneficiaryAddressLine3, r.BeneficiaryName, r.BeneficiaryReference, r.Description, r.InputMessageAccountabilityData, r.OriginatorAddressLine1, r.OriginatorAddressLine2, r.OriginatorAddressLine3, r.OriginatorName, r.OriginatorToBeneficiaryInformationLine1, r.OriginatorToBeneficiaryInformationLine2, r.OriginatorToBeneficiaryInformationLine3, r.OriginatorToBeneficiaryInformationLine4, r.OriginatorToBeneficiaryInformation)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrency] `json:"currency,required"`
	// The start of the period for which this transaction paid interest.
	PeriodStart fields.Field[time.Time] `json:"period_start,required" format:"date-time"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd fields.Field[time.Time] `json:"period_end,required" format:"date-time"`
	// The account on which the interest was accrued.
	AccruedOnAccountID fields.Field[string] `json:"accrued_on_account_id,required,nullable"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPayment
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPayment) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPayment) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPayment{Amount:%s Currency:%s PeriodStart:%s PeriodEnd:%s AccruedOnAccountID:%s}", r.Amount, r.Currency, r.PeriodStart, r.PeriodEnd, r.AccruedOnAccountID)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency] `json:"currency,required"`
	Reason   fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason]   `json:"reason,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource{Amount:%s Currency:%s Reason:%s}", r.Amount, r.Currency, r.Reason)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonBankMigration      InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "bank_migration"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonCashback           InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "cashback"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonEmpyrealAdjustment InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "empyreal_adjustment"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonError              InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "error"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonErrorCorrection    InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "error_correction"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonFees               InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "fees"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonInterest           InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "interest"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonSampleFunds        InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "sample_funds"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonSampleFundsReturn  InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "sample_funds_return"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefund struct {
	// The refunded amount in the minor unit of the refunded currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
	// currency.
	Currency             fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrency] `json:"currency,required"`
	MerchantAcceptorID   fields.Field[string]                                                                                  `json:"merchant_acceptor_id,required"`
	MerchantCity         fields.Field[string]                                                                                  `json:"merchant_city,required,nullable"`
	MerchantCountry      fields.Field[string]                                                                                  `json:"merchant_country,required"`
	MerchantDescriptor   fields.Field[string]                                                                                  `json:"merchant_descriptor,required"`
	MerchantState        fields.Field[string]                                                                                  `json:"merchant_state,required,nullable"`
	MerchantCategoryCode fields.Field[string]                                                                                  `json:"merchant_category_code,required,nullable"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefund
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefund) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefund) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefund{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", r.Amount, r.Currency, r.MerchantAcceptorID, r.MerchantCity, r.MerchantCountry, r.MerchantDescriptor, r.MerchantState, r.MerchantCategoryCode)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlement struct {
	// The settled amount in the minor unit of the settlement currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
	// currency.
	Currency             fields.Field[InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrency] `json:"currency,required"`
	MerchantAcceptorID   fields.Field[string]                                                                                      `json:"merchant_acceptor_id,required"`
	MerchantCity         fields.Field[string]                                                                                      `json:"merchant_city,required,nullable"`
	MerchantCountry      fields.Field[string]                                                                                      `json:"merchant_country,required,nullable"`
	MerchantDescriptor   fields.Field[string]                                                                                      `json:"merchant_descriptor,required"`
	MerchantState        fields.Field[string]                                                                                      `json:"merchant_state,required,nullable"`
	MerchantCategoryCode fields.Field[string]                                                                                      `json:"merchant_category_code,required,nullable"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlement
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlement) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlement) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlement{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", r.Amount, r.Currency, r.MerchantAcceptorID, r.MerchantCity, r.MerchantCountry, r.MerchantDescriptor, r.MerchantState, r.MerchantCategoryCode)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFunds struct {
	// Where the sample funds came from.
	Originator fields.Field[string] `json:"originator,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFunds into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFunds) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFunds) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFunds{Originator:%s}", r.Originator)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentIntention struct {
	// The transfer amount in USD cents.
	Amount             fields.Field[int64]  `json:"amount,required"`
	AccountNumber      fields.Field[string] `json:"account_number,required"`
	RoutingNumber      fields.Field[string] `json:"routing_number,required"`
	MessageToRecipient fields.Field[string] `json:"message_to_recipient,required"`
	TransferID         fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentIntention
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentIntention) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentIntention) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentIntention{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s TransferID:%s}", r.Amount, r.AccountNumber, r.RoutingNumber, r.MessageToRecipient, r.TransferID)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentRejection struct {
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentRejection
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentRejection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentRejection) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentRejection{TransferID:%s}", r.TransferID)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention struct {
	// The transfer amount in USD cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The destination account number.
	AccountNumber fields.Field[string] `json:"account_number,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber fields.Field[string] `json:"routing_number,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient fields.Field[string] `json:"message_to_recipient,required"`
	TransferID         fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s TransferID:%s}", r.Amount, r.AccountNumber, r.RoutingNumber, r.MessageToRecipient, r.TransferID)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejection struct {
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejection
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejection) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejection{TransferID:%s}", r.TransferID)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionType string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionTypeTransaction InboundRealTimePaymentsTransferSimulationResultTransactionType = "transaction"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction struct {
	// The identifier for the Account the Declined Transaction belongs to.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The Declined Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transcation's Account.
	Currency fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency] `json:"currency,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// This is the description the vendor provides.
	Description fields.Field[string] `json:"description,required"`
	// The Declined Transaction identifier.
	ID fields.Field[string] `json:"id,required"`
	// The identifier for the route this Declined Transaction came through. Routes are
	// things like cards and ACH details.
	RouteID fields.Field[string] `json:"route_id,required,nullable"`
	// The type of the route this Declined Transaction came through.
	RouteType fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionRouteType] `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Declined Transaction. For example, for a card transaction this lists the
	// merchant's industry and location. Note that for backwards compatibility reasons,
	// additional undocumented keys may appear in this object. These should be treated
	// as deprecated and will be removed in the future.
	Source fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource] `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionType] `json:"type,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction{AccountID:%s Amount:%s Currency:%s CreatedAt:%s Description:%s ID:%s RouteID:%s RouteType:%s Source:%s Type:%s}", r.AccountID, r.Amount, r.Currency, r.CreatedAt, r.Description, r.ID, r.RouteID, r.RouteType, r.Source, r.Type)
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrencyCad InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrencyChf InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrencyEur InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrencyGbp InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrencyJpy InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrencyUsd InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionRouteType string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionRouteTypeAccountNumber InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionRouteType = "account_number"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionRouteTypeCard          InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionRouteType = "card"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource struct {
	// The type of decline that took place. We may add additional possible values for
	// this enum over time; your application should be able to handle such additions
	// gracefully.
	Category fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory] `json:"category,required"`
	// A ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline] `json:"ach_decline,required,nullable"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline] `json:"card_decline,required,nullable"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline] `json:"check_decline,required,nullable"`
	// A Inbound Real Time Payments Transfer Decline object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline] `json:"inbound_real_time_payments_transfer_decline,required,nullable"`
	// A International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline] `json:"international_ach_decline,required,nullable"`
	// A Deprecated Card Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_decline`.
	CardRouteDecline fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline] `json:"card_route_decline,required,nullable"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource{Category:%s ACHDecline:%s CardDecline:%s CheckDecline:%s InboundRealTimePaymentsTransferDecline:%s InternationalACHDecline:%s CardRouteDecline:%s}", r.Category, r.ACHDecline, r.CardDecline, r.CheckDecline, r.InboundRealTimePaymentsTransferDecline, r.InternationalACHDecline, r.CardRouteDecline)
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryACHDecline                             InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "ach_decline"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryCardDecline                            InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "card_decline"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryCheckDecline                           InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "check_decline"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryInboundRealTimePaymentsTransferDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "inbound_real_time_payments_transfer_decline"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryInternationalACHDecline                InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "international_ach_decline"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryCardRouteDecline                       InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "card_route_decline"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryOther                                  InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "other"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount                             fields.Field[int64]  `json:"amount,required"`
	OriginatorCompanyName              fields.Field[string] `json:"originator_company_name,required"`
	OriginatorCompanyDescriptiveDate   fields.Field[string] `json:"originator_company_descriptive_date,required,nullable"`
	OriginatorCompanyDiscretionaryData fields.Field[string] `json:"originator_company_discretionary_data,required,nullable"`
	OriginatorCompanyID                fields.Field[string] `json:"originator_company_id,required"`
	// Why the ACH transfer was declined.
	Reason           fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason] `json:"reason,required"`
	ReceiverIDNumber fields.Field[string]                                                                                   `json:"receiver_id_number,required,nullable"`
	ReceiverName     fields.Field[string]                                                                                   `json:"receiver_name,required,nullable"`
	TraceNumber      fields.Field[string]                                                                                   `json:"trace_number,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline{Amount:%s OriginatorCompanyName:%s OriginatorCompanyDescriptiveDate:%s OriginatorCompanyDiscretionaryData:%s OriginatorCompanyID:%s Reason:%s ReceiverIDNumber:%s ReceiverName:%s TraceNumber:%s}", r.Amount, r.OriginatorCompanyName, r.OriginatorCompanyDescriptiveDate, r.OriginatorCompanyDiscretionaryData, r.OriginatorCompanyID, r.Reason, r.ReceiverIDNumber, r.ReceiverName, r.TraceNumber)
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonACHRouteCanceled             InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "ach_route_canceled"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonACHRouteDisabled             InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "ach_route_disabled"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonBreachesLimit                InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonDuplicateReturn              InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "duplicate_return"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonEntityNotActive              InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "entity_not_active"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonTransactionNotAllowed        InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "transaction_not_allowed"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonGroupLocked                  InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "group_locked"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonInsufficientFunds            InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "insufficient_funds"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonNoACHRoute                   InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "no_ach_route"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonOriginatorRequest            InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "originator_request"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline struct {
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID fields.Field[string] `json:"merchant_acceptor_id,required"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor fields.Field[string] `json:"merchant_descriptor,required"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode fields.Field[string] `json:"merchant_category_code,required,nullable"`
	// The city the merchant resides in.
	MerchantCity fields.Field[string] `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry fields.Field[string] `json:"merchant_country,required,nullable"`
	// The payment network used to process this card authorization
	Network fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetwork] `json:"network,required"`
	// Fields specific to the `network`
	NetworkDetails fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetails] `json:"network_details,required"`
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency] `json:"currency,required"`
	// Why the transaction was declined.
	Reason fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason] `json:"reason,required"`
	// The state the merchant resides in.
	MerchantState fields.Field[string] `json:"merchant_state,required,nullable"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID fields.Field[string] `json:"real_time_decision_id,required,nullable"`
	// If the authorization was attempted using a Digital Wallet Token (such as an
	// Apple Pay purchase), the identifier of the token that was used.
	DigitalWalletTokenID fields.Field[string] `json:"digital_wallet_token_id,required,nullable"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline{MerchantAcceptorID:%s MerchantDescriptor:%s MerchantCategoryCode:%s MerchantCity:%s MerchantCountry:%s Network:%s NetworkDetails:%s Amount:%s Currency:%s Reason:%s MerchantState:%s RealTimeDecisionID:%s DigitalWalletTokenID:%s}", r.MerchantAcceptorID, r.MerchantDescriptor, r.MerchantCategoryCode, r.MerchantCity, r.MerchantCountry, r.Network, r.NetworkDetails, r.Amount, r.Currency, r.Reason, r.MerchantState, r.RealTimeDecisionID, r.DigitalWalletTokenID)
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetwork string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkVisa InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetwork = "visa"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisa] `json:"visa,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetails
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetails) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetails) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetails{Visa:%s}", r.Visa)
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator] `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode fields.Field[PointOfServiceEntryMode] `json:"point_of_service_entry_mode,required,nullable"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisa
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisa{ElectronicCommerceIndicator:%s PointOfServiceEntryMode:%s}", r.ElectronicCommerceIndicator, r.PointOfServiceEntryMode)
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                           InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                                InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                              InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                    InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                 InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt_3dsCapableMerchant InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                      InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                     InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrencyCad InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrencyChf InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrencyEur InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrencyGbp InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrencyJpy InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrencyUsd InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonCardNotActive               InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "card_not_active"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonEntityNotActive             InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "entity_not_active"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonGroupLocked                 InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "group_locked"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonInsufficientFunds           InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "insufficient_funds"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonCvv2Mismatch                InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "cvv2_mismatch"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonTransactionNotAllowed       InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "transaction_not_allowed"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonBreachesLimit               InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "breaches_limit"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonWebhookDeclined             InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "webhook_declined"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonWebhookTimedOut             InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "webhook_timed_out"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonDeclinedByStandInProcessing InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "declined_by_stand_in_processing"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonInvalidPhysicalCard         InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "invalid_physical_card"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount        fields.Field[int64]  `json:"amount,required"`
	AuxiliaryOnUs fields.Field[string] `json:"auxiliary_on_us,required,nullable"`
	// Why the check was declined.
	Reason fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason] `json:"reason,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline{Amount:%s AuxiliaryOnUs:%s Reason:%s}", r.Amount, r.AuxiliaryOnUs, r.Reason)
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonACHRouteCanceled      InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "ach_route_canceled"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonACHRouteDisabled      InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "ach_route_disabled"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonBreachesLimit         InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "breaches_limit"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonEntityNotActive       InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "entity_not_active"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonGroupLocked           InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "group_locked"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonInsufficientFunds     InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "insufficient_funds"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonUnableToLocateAccount InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "unable_to_locate_account"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonUnableToProcess       InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "unable_to_process"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonReferToImage          InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "refer_to_image"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonStopPaymentRequested  InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "stop_payment_requested"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonReturned              InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "returned"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonDuplicatePresentment  InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "duplicate_presentment"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real Time Payments
	// transfer.
	Currency fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency] `json:"currency,required"`
	// Why the transfer was declined.
	Reason fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason] `json:"reason,required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName fields.Field[string] `json:"creditor_name,required"`
	// The name provided by the sender of the transfer.
	DebtorName fields.Field[string] `json:"debtor_name,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber fields.Field[string] `json:"debtor_account_number,required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber fields.Field[string] `json:"debtor_routing_number,required"`
	// The Real Time Payments network identification of the declined transfer.
	TransactionIdentification fields.Field[string] `json:"transaction_identification,required"`
	// Additional information included with the transfer.
	RemittanceInformation fields.Field[string] `json:"remittance_information,required,nullable"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline{Amount:%s Currency:%s Reason:%s CreditorName:%s DebtorName:%s DebtorAccountNumber:%s DebtorRoutingNumber:%s TransactionIdentification:%s RemittanceInformation:%s}", r.Amount, r.Currency, r.Reason, r.CreditorName, r.DebtorName, r.DebtorAccountNumber, r.DebtorRoutingNumber, r.TransactionIdentification, r.RemittanceInformation)
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyCad InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyChf InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyEur InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyGbp InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyJpy InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyUsd InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberCanceled      InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_canceled"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberDisabled      InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_disabled"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonGroupLocked                InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "group_locked"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonEntityNotActive            InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "entity_not_active"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "real_time_payments_not_enabled"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount                                                 fields.Field[int64]  `json:"amount,required"`
	ForeignExchangeIndicator                               fields.Field[string] `json:"foreign_exchange_indicator,required"`
	ForeignExchangeReferenceIndicator                      fields.Field[string] `json:"foreign_exchange_reference_indicator,required"`
	ForeignExchangeReference                               fields.Field[string] `json:"foreign_exchange_reference,required,nullable"`
	DestinationCountryCode                                 fields.Field[string] `json:"destination_country_code,required"`
	DestinationCurrencyCode                                fields.Field[string] `json:"destination_currency_code,required"`
	ForeignPaymentAmount                                   fields.Field[int64]  `json:"foreign_payment_amount,required"`
	ForeignTraceNumber                                     fields.Field[string] `json:"foreign_trace_number,required,nullable"`
	InternationalTransactionTypeCode                       fields.Field[string] `json:"international_transaction_type_code,required"`
	OriginatingCurrencyCode                                fields.Field[string] `json:"originating_currency_code,required"`
	OriginatingDepositoryFinancialInstitutionName          fields.Field[string] `json:"originating_depository_financial_institution_name,required"`
	OriginatingDepositoryFinancialInstitutionIDQualifier   fields.Field[string] `json:"originating_depository_financial_institution_id_qualifier,required"`
	OriginatingDepositoryFinancialInstitutionID            fields.Field[string] `json:"originating_depository_financial_institution_id,required"`
	OriginatingDepositoryFinancialInstitutionBranchCountry fields.Field[string] `json:"originating_depository_financial_institution_branch_country,required"`
	OriginatorCity                                         fields.Field[string] `json:"originator_city,required"`
	OriginatorCompanyEntryDescription                      fields.Field[string] `json:"originator_company_entry_description,required"`
	OriginatorCountry                                      fields.Field[string] `json:"originator_country,required"`
	OriginatorIdentification                               fields.Field[string] `json:"originator_identification,required"`
	OriginatorName                                         fields.Field[string] `json:"originator_name,required"`
	OriginatorPostalCode                                   fields.Field[string] `json:"originator_postal_code,required,nullable"`
	OriginatorStreetAddress                                fields.Field[string] `json:"originator_street_address,required"`
	OriginatorStateOrProvince                              fields.Field[string] `json:"originator_state_or_province,required,nullable"`
	PaymentRelatedInformation                              fields.Field[string] `json:"payment_related_information,required,nullable"`
	PaymentRelatedInformation2                             fields.Field[string] `json:"payment_related_information2,required,nullable"`
	ReceiverIdentificationNumber                           fields.Field[string] `json:"receiver_identification_number,required,nullable"`
	ReceiverStreetAddress                                  fields.Field[string] `json:"receiver_street_address,required"`
	ReceiverCity                                           fields.Field[string] `json:"receiver_city,required"`
	ReceiverStateOrProvince                                fields.Field[string] `json:"receiver_state_or_province,required,nullable"`
	ReceiverCountry                                        fields.Field[string] `json:"receiver_country,required"`
	ReceiverPostalCode                                     fields.Field[string] `json:"receiver_postal_code,required,nullable"`
	ReceivingCompanyOrIndividualName                       fields.Field[string] `json:"receiving_company_or_individual_name,required"`
	ReceivingDepositoryFinancialInstitutionName            fields.Field[string] `json:"receiving_depository_financial_institution_name,required"`
	ReceivingDepositoryFinancialInstitutionIDQualifier     fields.Field[string] `json:"receiving_depository_financial_institution_id_qualifier,required"`
	ReceivingDepositoryFinancialInstitutionID              fields.Field[string] `json:"receiving_depository_financial_institution_id,required"`
	ReceivingDepositoryFinancialInstitutionCountry         fields.Field[string] `json:"receiving_depository_financial_institution_country,required"`
	TraceNumber                                            fields.Field[string] `json:"trace_number,required"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline{Amount:%s ForeignExchangeIndicator:%s ForeignExchangeReferenceIndicator:%s ForeignExchangeReference:%s DestinationCountryCode:%s DestinationCurrencyCode:%s ForeignPaymentAmount:%s ForeignTraceNumber:%s InternationalTransactionTypeCode:%s OriginatingCurrencyCode:%s OriginatingDepositoryFinancialInstitutionName:%s OriginatingDepositoryFinancialInstitutionIDQualifier:%s OriginatingDepositoryFinancialInstitutionID:%s OriginatingDepositoryFinancialInstitutionBranchCountry:%s OriginatorCity:%s OriginatorCompanyEntryDescription:%s OriginatorCountry:%s OriginatorIdentification:%s OriginatorName:%s OriginatorPostalCode:%s OriginatorStreetAddress:%s OriginatorStateOrProvince:%s PaymentRelatedInformation:%s PaymentRelatedInformation2:%s ReceiverIdentificationNumber:%s ReceiverStreetAddress:%s ReceiverCity:%s ReceiverStateOrProvince:%s ReceiverCountry:%s ReceiverPostalCode:%s ReceivingCompanyOrIndividualName:%s ReceivingDepositoryFinancialInstitutionName:%s ReceivingDepositoryFinancialInstitutionIDQualifier:%s ReceivingDepositoryFinancialInstitutionID:%s ReceivingDepositoryFinancialInstitutionCountry:%s TraceNumber:%s}", r.Amount, r.ForeignExchangeIndicator, r.ForeignExchangeReferenceIndicator, r.ForeignExchangeReference, r.DestinationCountryCode, r.DestinationCurrencyCode, r.ForeignPaymentAmount, r.ForeignTraceNumber, r.InternationalTransactionTypeCode, r.OriginatingCurrencyCode, r.OriginatingDepositoryFinancialInstitutionName, r.OriginatingDepositoryFinancialInstitutionIDQualifier, r.OriginatingDepositoryFinancialInstitutionID, r.OriginatingDepositoryFinancialInstitutionBranchCountry, r.OriginatorCity, r.OriginatorCompanyEntryDescription, r.OriginatorCountry, r.OriginatorIdentification, r.OriginatorName, r.OriginatorPostalCode, r.OriginatorStreetAddress, r.OriginatorStateOrProvince, r.PaymentRelatedInformation, r.PaymentRelatedInformation2, r.ReceiverIdentificationNumber, r.ReceiverStreetAddress, r.ReceiverCity, r.ReceiverStateOrProvince, r.ReceiverCountry, r.ReceiverPostalCode, r.ReceivingCompanyOrIndividualName, r.ReceivingDepositoryFinancialInstitutionName, r.ReceivingDepositoryFinancialInstitutionIDQualifier, r.ReceivingDepositoryFinancialInstitutionID, r.ReceivingDepositoryFinancialInstitutionCountry, r.TraceNumber)
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency             fields.Field[InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrency] `json:"currency,required"`
	MerchantAcceptorID   fields.Field[string]                                                                                           `json:"merchant_acceptor_id,required"`
	MerchantCity         fields.Field[string]                                                                                           `json:"merchant_city,required,nullable"`
	MerchantCountry      fields.Field[string]                                                                                           `json:"merchant_country,required"`
	MerchantDescriptor   fields.Field[string]                                                                                           `json:"merchant_descriptor,required"`
	MerchantState        fields.Field[string]                                                                                           `json:"merchant_state,required,nullable"`
	MerchantCategoryCode fields.Field[string]                                                                                           `json:"merchant_category_code,required,nullable"`
}

// MarshalJSON serializes
// InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline) String() (result string) {
	return fmt.Sprintf("&InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", r.Amount, r.Currency, r.MerchantAcceptorID, r.MerchantCity, r.MerchantCountry, r.MerchantDescriptor, r.MerchantState, r.MerchantCategoryCode)
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrencyCad InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrencyChf InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrencyEur InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrencyGbp InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrencyJpy InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrencyUsd InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionType string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionTypeDeclinedTransaction InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionType = "declined_transaction"
)

type InboundRealTimePaymentsTransferSimulationResultType string

const (
	InboundRealTimePaymentsTransferSimulationResultTypeInboundRealTimePaymentsTransferSimulationResult InboundRealTimePaymentsTransferSimulationResultType = "inbound_real_time_payments_transfer_simulation_result"
)

type SimulateARealTimePaymentsTransferToYourAccountParameters struct {
	// The identifier of the Account Number the inbound Real Time Payments Transfer is
	// for.
	AccountNumberID fields.Field[string] `json:"account_number_id,required"`
	// The transfer amount in USD cents. Must be positive.
	Amount fields.Field[int64] `json:"amount,required"`
	// The identifier of a pending Request for Payment that this transfer will fulfill.
	RequestForPaymentID fields.Field[string] `json:"request_for_payment_id"`
	// The name provided by the sender of the transfer.
	DebtorName fields.Field[string] `json:"debtor_name"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber fields.Field[string] `json:"debtor_account_number"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber fields.Field[string] `json:"debtor_routing_number"`
	// Additional information included with the transfer.
	RemittanceInformation fields.Field[string] `json:"remittance_information"`
}

// MarshalJSON serializes SimulateARealTimePaymentsTransferToYourAccountParameters
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *SimulateARealTimePaymentsTransferToYourAccountParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SimulateARealTimePaymentsTransferToYourAccountParameters) String() (result string) {
	return fmt.Sprintf("&SimulateARealTimePaymentsTransferToYourAccountParameters{AccountNumberID:%s Amount:%s RequestForPaymentID:%s DebtorName:%s DebtorAccountNumber:%s DebtorRoutingNumber:%s RemittanceInformation:%s}", r.AccountNumberID, r.Amount, r.RequestForPaymentID, r.DebtorName, r.DebtorAccountNumber, r.DebtorRoutingNumber, r.RemittanceInformation)
}
