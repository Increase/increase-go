package requests

import (
	"fmt"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/fields"
)

type WireTransferSimulation struct {
	// If the Wire Transfer attempt succeeds, this will contain the resulting
	// [Transaction](#transactions) object. The Transaction's `source` will be of
	// `category: inbound_wire_transfer`.
	Transaction fields.Field[WireTransferSimulationTransaction] `json:"transaction,required"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_wire_transfer_simulation_result`.
	Type fields.Field[WireTransferSimulationType] `json:"type,required"`
}

// MarshalJSON serializes WireTransferSimulation into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulation) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulation) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulation{Transaction:%s Type:%s}", r.Transaction, r.Type)
}

type WireTransferSimulationTransaction struct {
	// The identifier for the Account the Transaction belongs to.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transcation's
	// Account.
	Currency fields.Field[WireTransferSimulationTransactionCurrency] `json:"currency,required"`
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
	RouteType fields.Field[string] `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Transaction. Note that for backwards compatibility reasons, additional
	// undocumented keys may appear in this object. These should be treated as
	// deprecated and will be removed in the future.
	Source fields.Field[WireTransferSimulationTransactionSource] `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type fields.Field[WireTransferSimulationTransactionType] `json:"type,required"`
}

// MarshalJSON serializes WireTransferSimulationTransaction into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransaction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransaction) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransaction{AccountID:%s Amount:%s Currency:%s CreatedAt:%s Description:%s ID:%s RouteID:%s RouteType:%s Source:%s Type:%s}", r.AccountID, r.Amount, r.Currency, r.CreatedAt, r.Description, r.ID, r.RouteID, r.RouteType, r.Source, r.Type)
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
	Category fields.Field[WireTransferSimulationTransactionSourceCategory] `json:"category,required"`
	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention fields.Field[WireTransferSimulationTransactionSourceAccountTransferIntention] `json:"account_transfer_intention,required,nullable"`
	// A ACH Check Conversion Return object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_check_conversion_return`.
	ACHCheckConversionReturn fields.Field[WireTransferSimulationTransactionSourceACHCheckConversionReturn] `json:"ach_check_conversion_return,required,nullable"`
	// A ACH Check Conversion object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_check_conversion`.
	ACHCheckConversion fields.Field[WireTransferSimulationTransactionSourceACHCheckConversion] `json:"ach_check_conversion,required,nullable"`
	// A ACH Transfer Intention object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention fields.Field[WireTransferSimulationTransactionSourceACHTransferIntention] `json:"ach_transfer_intention,required,nullable"`
	// A ACH Transfer Rejection object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection fields.Field[WireTransferSimulationTransactionSourceACHTransferRejection] `json:"ach_transfer_rejection,required,nullable"`
	// A ACH Transfer Return object. This field will be present in the JSON response if
	// and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn fields.Field[WireTransferSimulationTransactionSourceACHTransferReturn] `json:"ach_transfer_return,required,nullable"`
	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance fields.Field[WireTransferSimulationTransactionSourceCardDisputeAcceptance] `json:"card_dispute_acceptance,required,nullable"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund fields.Field[WireTransferSimulationTransactionSourceCardRefund] `json:"card_refund,required,nullable"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement fields.Field[WireTransferSimulationTransactionSourceCardSettlement] `json:"card_settlement,required,nullable"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance fields.Field[WireTransferSimulationTransactionSourceCheckDepositAcceptance] `json:"check_deposit_acceptance,required,nullable"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn fields.Field[WireTransferSimulationTransactionSourceCheckDepositReturn] `json:"check_deposit_return,required,nullable"`
	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention fields.Field[WireTransferSimulationTransactionSourceCheckTransferIntention] `json:"check_transfer_intention,required,nullable"`
	// A Check Transfer Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_transfer_return`.
	CheckTransferReturn fields.Field[WireTransferSimulationTransactionSourceCheckTransferReturn] `json:"check_transfer_return,required,nullable"`
	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection fields.Field[WireTransferSimulationTransactionSourceCheckTransferRejection] `json:"check_transfer_rejection,required,nullable"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest fields.Field[WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest] `json:"check_transfer_stop_payment_request,required,nullable"`
	// A Dispute Resolution object. This field will be present in the JSON response if
	// and only if `category` is equal to `dispute_resolution`.
	DisputeResolution fields.Field[WireTransferSimulationTransactionSourceDisputeResolution] `json:"dispute_resolution,required,nullable"`
	// A Empyreal Cash Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `empyreal_cash_deposit`.
	EmpyrealCashDeposit fields.Field[WireTransferSimulationTransactionSourceEmpyrealCashDeposit] `json:"empyreal_cash_deposit,required,nullable"`
	// A Inbound ACH Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer fields.Field[WireTransferSimulationTransactionSourceInboundACHTransfer] `json:"inbound_ach_transfer,required,nullable"`
	// A Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck fields.Field[WireTransferSimulationTransactionSourceInboundCheck] `json:"inbound_check,required,nullable"`
	// A Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer fields.Field[WireTransferSimulationTransactionSourceInboundInternationalACHTransfer] `json:"inbound_international_ach_transfer,required,nullable"`
	// A Inbound Real Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation fields.Field[WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation] `json:"inbound_real_time_payments_transfer_confirmation,required,nullable"`
	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal fields.Field[WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal] `json:"inbound_wire_drawdown_payment_reversal,required,nullable"`
	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment fields.Field[WireTransferSimulationTransactionSourceInboundWireDrawdownPayment] `json:"inbound_wire_drawdown_payment,required,nullable"`
	// A Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal fields.Field[WireTransferSimulationTransactionSourceInboundWireReversal] `json:"inbound_wire_reversal,required,nullable"`
	// A Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer fields.Field[WireTransferSimulationTransactionSourceInboundWireTransfer] `json:"inbound_wire_transfer,required,nullable"`
	// A Interest Payment object. This field will be present in the JSON response if
	// and only if `category` is equal to `interest_payment`.
	InterestPayment fields.Field[WireTransferSimulationTransactionSourceInterestPayment] `json:"interest_payment,required,nullable"`
	// A Internal Source object. This field will be present in the JSON response if and
	// only if `category` is equal to `internal_source`.
	InternalSource fields.Field[WireTransferSimulationTransactionSourceInternalSource] `json:"internal_source,required,nullable"`
	// A Deprecated Card Refund object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_route_refund`.
	CardRouteRefund fields.Field[WireTransferSimulationTransactionSourceCardRouteRefund] `json:"card_route_refund,required,nullable"`
	// A Deprecated Card Settlement object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_settlement`.
	CardRouteSettlement fields.Field[WireTransferSimulationTransactionSourceCardRouteSettlement] `json:"card_route_settlement,required,nullable"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds fields.Field[WireTransferSimulationTransactionSourceSampleFunds] `json:"sample_funds,required,nullable"`
	// A Wire Drawdown Payment Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_intention`.
	WireDrawdownPaymentIntention fields.Field[WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention] `json:"wire_drawdown_payment_intention,required,nullable"`
	// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_rejection`.
	WireDrawdownPaymentRejection fields.Field[WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection] `json:"wire_drawdown_payment_rejection,required,nullable"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention fields.Field[WireTransferSimulationTransactionSourceWireTransferIntention] `json:"wire_transfer_intention,required,nullable"`
	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection fields.Field[WireTransferSimulationTransactionSourceWireTransferRejection] `json:"wire_transfer_rejection,required,nullable"`
}

// MarshalJSON serializes WireTransferSimulationTransactionSource into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSource) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSource) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSource{Category:%s AccountTransferIntention:%s ACHCheckConversionReturn:%s ACHCheckConversion:%s ACHTransferIntention:%s ACHTransferRejection:%s ACHTransferReturn:%s CardDisputeAcceptance:%s CardRefund:%s CardSettlement:%s CheckDepositAcceptance:%s CheckDepositReturn:%s CheckTransferIntention:%s CheckTransferReturn:%s CheckTransferRejection:%s CheckTransferStopPaymentRequest:%s DisputeResolution:%s EmpyrealCashDeposit:%s InboundACHTransfer:%s InboundCheck:%s InboundInternationalACHTransfer:%s InboundRealTimePaymentsTransferConfirmation:%s InboundWireDrawdownPaymentReversal:%s InboundWireDrawdownPayment:%s InboundWireReversal:%s InboundWireTransfer:%s InterestPayment:%s InternalSource:%s CardRouteRefund:%s CardRouteSettlement:%s SampleFunds:%s WireDrawdownPaymentIntention:%s WireDrawdownPaymentRejection:%s WireTransferIntention:%s WireTransferRejection:%s}", r.Category, r.AccountTransferIntention, r.ACHCheckConversionReturn, r.ACHCheckConversion, r.ACHTransferIntention, r.ACHTransferRejection, r.ACHTransferReturn, r.CardDisputeAcceptance, r.CardRefund, r.CardSettlement, r.CheckDepositAcceptance, r.CheckDepositReturn, r.CheckTransferIntention, r.CheckTransferReturn, r.CheckTransferRejection, r.CheckTransferStopPaymentRequest, r.DisputeResolution, r.EmpyrealCashDeposit, r.InboundACHTransfer, r.InboundCheck, r.InboundInternationalACHTransfer, r.InboundRealTimePaymentsTransferConfirmation, r.InboundWireDrawdownPaymentReversal, r.InboundWireDrawdownPayment, r.InboundWireReversal, r.InboundWireTransfer, r.InterestPayment, r.InternalSource, r.CardRouteRefund, r.CardRouteSettlement, r.SampleFunds, r.WireDrawdownPaymentIntention, r.WireDrawdownPaymentRejection, r.WireTransferIntention, r.WireTransferRejection)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency fields.Field[WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency] `json:"currency,required"`
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
// WireTransferSimulationTransactionSourceAccountTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceAccountTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceAccountTransferIntention) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceAccountTransferIntention{Amount:%s Currency:%s Description:%s DestinationAccountID:%s SourceAccountID:%s TransferID:%s}", r.Amount, r.Currency, r.Description, r.DestinationAccountID, r.SourceAccountID, r.TransferID)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// Why the transfer was returned.
	ReturnReasonCode fields.Field[string] `json:"return_reason_code,required"`
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceACHCheckConversionReturn into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceACHCheckConversionReturn) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceACHCheckConversionReturn) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceACHCheckConversionReturn{Amount:%s ReturnReasonCode:%s}", r.Amount, r.ReturnReasonCode)
}

type WireTransferSimulationTransactionSourceACHCheckConversion struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The identifier of the File containing an image of the returned check.
	FileID fields.Field[string] `json:"file_id,required"`
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceACHCheckConversion
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *WireTransferSimulationTransactionSourceACHCheckConversion) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceACHCheckConversion) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceACHCheckConversion{Amount:%s FileID:%s}", r.Amount, r.FileID)
}

type WireTransferSimulationTransactionSourceACHTransferIntention struct {
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
// WireTransferSimulationTransactionSourceACHTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceACHTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceACHTransferIntention) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceACHTransferIntention{Amount:%s AccountNumber:%s RoutingNumber:%s StatementDescriptor:%s TransferID:%s}", r.Amount, r.AccountNumber, r.RoutingNumber, r.StatementDescriptor, r.TransferID)
}

type WireTransferSimulationTransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceACHTransferRejection into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceACHTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceACHTransferRejection) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceACHTransferRejection{TransferID:%s}", r.TransferID)
}

type WireTransferSimulationTransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode fields.Field[WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode] `json:"return_reason_code,required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID fields.Field[string] `json:"transfer_id,required"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID fields.Field[string] `json:"transaction_id,required"`
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceACHTransferReturn
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *WireTransferSimulationTransactionSourceACHTransferReturn) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceACHTransferReturn) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceACHTransferReturn{CreatedAt:%s ReturnReasonCode:%s TransferID:%s TransactionID:%s}", r.CreatedAt, r.ReturnReasonCode, r.TransferID, r.TransactionID)
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
	AcceptedAt fields.Field[time.Time] `json:"accepted_at,required" format:"date-time"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID fields.Field[string] `json:"card_dispute_id,required"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID fields.Field[string] `json:"transaction_id,required,nullable"`
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceCardDisputeAcceptance into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceCardDisputeAcceptance) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceCardDisputeAcceptance) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCardDisputeAcceptance{AcceptedAt:%s CardDisputeID:%s TransactionID:%s}", r.AcceptedAt, r.CardDisputeID, r.TransactionID)
}

type WireTransferSimulationTransactionSourceCardRefund struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[WireTransferSimulationTransactionSourceCardRefundCurrency] `json:"currency,required"`
	// The identifier for the Transaction this refunds, if any.
	CardSettlementTransactionID fields.Field[string] `json:"card_settlement_transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type fields.Field[WireTransferSimulationTransactionSourceCardRefundType] `json:"type,required"`
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceCardRefund into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *WireTransferSimulationTransactionSourceCardRefund) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceCardRefund) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCardRefund{Amount:%s Currency:%s CardSettlementTransactionID:%s Type:%s}", r.Amount, r.Currency, r.CardSettlementTransactionID, r.Type)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's settlement currency.
	Currency fields.Field[WireTransferSimulationTransactionSourceCardSettlementCurrency] `json:"currency,required"`
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
	Type fields.Field[WireTransferSimulationTransactionSourceCardSettlementType] `json:"type,required"`
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceCardSettlement
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *WireTransferSimulationTransactionSourceCardSettlement) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceCardSettlement) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCardSettlement{Amount:%s Currency:%s PresentmentAmount:%s PresentmentCurrency:%s MerchantCity:%s MerchantCountry:%s MerchantName:%s MerchantCategoryCode:%s MerchantState:%s PendingTransactionID:%s Type:%s}", r.Amount, r.Currency, r.PresentmentAmount, r.PresentmentCurrency, r.MerchantCity, r.MerchantCountry, r.MerchantName, r.MerchantCategoryCode, r.MerchantState, r.PendingTransactionID, r.Type)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency] `json:"currency,required"`
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
// WireTransferSimulationTransactionSourceCheckDepositAcceptance into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceCheckDepositAcceptance) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceCheckDepositAcceptance) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCheckDepositAcceptance{Amount:%s Currency:%s AccountNumber:%s RoutingNumber:%s AuxiliaryOnUs:%s SerialNumber:%s CheckDepositID:%s}", r.Amount, r.Currency, r.AccountNumber, r.RoutingNumber, r.AuxiliaryOnUs, r.SerialNumber, r.CheckDepositID)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt fields.Field[time.Time] `json:"returned_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[WireTransferSimulationTransactionSourceCheckDepositReturnCurrency] `json:"currency,required"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID fields.Field[string] `json:"check_deposit_id,required"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID fields.Field[string]                                                                `json:"transaction_id,required"`
	ReturnReason  fields.Field[WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason] `json:"return_reason,required"`
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceCheckDepositReturn
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *WireTransferSimulationTransactionSourceCheckDepositReturn) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceCheckDepositReturn) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCheckDepositReturn{Amount:%s ReturnedAt:%s Currency:%s CheckDepositID:%s TransactionID:%s ReturnReason:%s}", r.Amount, r.ReturnedAt, r.Currency, r.CheckDepositID, r.TransactionID, r.ReturnReason)
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
	Currency fields.Field[WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency] `json:"currency,required"`
	// The name that will be printed on the check.
	RecipientName fields.Field[string] `json:"recipient_name,required"`
	// The identifier of the Check Transfer with which this is associated.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceCheckTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceCheckTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceCheckTransferIntention) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCheckTransferIntention{AddressLine1:%s AddressLine2:%s AddressCity:%s AddressState:%s AddressZip:%s Amount:%s Currency:%s RecipientName:%s TransferID:%s}", r.AddressLine1, r.AddressLine2, r.AddressCity, r.AddressState, r.AddressZip, r.Amount, r.Currency, r.RecipientName, r.TransferID)
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
	TransferID fields.Field[string] `json:"transfer_id,required"`
	// If available, a document with additional information about the return.
	FileID fields.Field[string] `json:"file_id,required,nullable"`
	// The reason why the check was returned.
	Reason fields.Field[WireTransferSimulationTransactionSourceCheckTransferReturnReason] `json:"reason,required"`
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceCheckTransferReturn into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceCheckTransferReturn) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceCheckTransferReturn) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCheckTransferReturn{TransferID:%s FileID:%s Reason:%s}", r.TransferID, r.FileID, r.Reason)
}

type WireTransferSimulationTransactionSourceCheckTransferReturnReason string

const (
	WireTransferSimulationTransactionSourceCheckTransferReturnReasonMailDeliveryFailure WireTransferSimulationTransactionSourceCheckTransferReturnReason = "mail_delivery_failure"
	WireTransferSimulationTransactionSourceCheckTransferReturnReasonRefusedByRecipient  WireTransferSimulationTransactionSourceCheckTransferReturnReason = "refused_by_recipient"
)

type WireTransferSimulationTransactionSourceCheckTransferRejection struct {
	// The identifier of the Check Transfer that led to this Transaction.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceCheckTransferRejection into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceCheckTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceCheckTransferRejection) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCheckTransferRejection{TransferID:%s}", r.TransferID)
}

type WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferID fields.Field[string] `json:"transfer_id,required"`
	// The transaction ID of the corresponding credit transaction.
	TransactionID fields.Field[string] `json:"transaction_id,required"`
	// The time the stop-payment was requested.
	RequestedAt fields.Field[time.Time] `json:"requested_at,required" format:"date-time"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type fields.Field[WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType] `json:"type,required"`
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest{TransferID:%s TransactionID:%s RequestedAt:%s Type:%s}", r.TransferID, r.TransactionID, r.RequestedAt, r.Type)
}

type WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType string

const (
	WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

type WireTransferSimulationTransactionSourceDisputeResolution struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[WireTransferSimulationTransactionSourceDisputeResolutionCurrency] `json:"currency,required"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionID fields.Field[string] `json:"disputed_transaction_id,required"`
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceDisputeResolution
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *WireTransferSimulationTransactionSourceDisputeResolution) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceDisputeResolution) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceDisputeResolution{Amount:%s Currency:%s DisputedTransactionID:%s}", r.Amount, r.Currency, r.DisputedTransactionID)
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
	Amount      fields.Field[int64]     `json:"amount,required"`
	BagID       fields.Field[string]    `json:"bag_id,required"`
	DepositDate fields.Field[time.Time] `json:"deposit_date,required" format:"date-time"`
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceEmpyrealCashDeposit into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceEmpyrealCashDeposit) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceEmpyrealCashDeposit) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceEmpyrealCashDeposit{Amount:%s BagID:%s DepositDate:%s}", r.Amount, r.BagID, r.DepositDate)
}

type WireTransferSimulationTransactionSourceInboundACHTransfer struct {
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

// MarshalJSON serializes WireTransferSimulationTransactionSourceInboundACHTransfer
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *WireTransferSimulationTransactionSourceInboundACHTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceInboundACHTransfer) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceInboundACHTransfer{Amount:%s OriginatorCompanyName:%s OriginatorCompanyDescriptiveDate:%s OriginatorCompanyDiscretionaryData:%s OriginatorCompanyEntryDescription:%s OriginatorCompanyID:%s ReceiverIDNumber:%s ReceiverName:%s TraceNumber:%s}", r.Amount, r.OriginatorCompanyName, r.OriginatorCompanyDescriptiveDate, r.OriginatorCompanyDiscretionaryData, r.OriginatorCompanyEntryDescription, r.OriginatorCompanyID, r.ReceiverIDNumber, r.ReceiverName, r.TraceNumber)
}

type WireTransferSimulationTransactionSourceInboundCheck struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency              fields.Field[WireTransferSimulationTransactionSourceInboundCheckCurrency] `json:"currency,required"`
	CheckNumber           fields.Field[string]                                                      `json:"check_number,required,nullable"`
	CheckFrontImageFileID fields.Field[string]                                                      `json:"check_front_image_file_id,required,nullable"`
	CheckRearImageFileID  fields.Field[string]                                                      `json:"check_rear_image_file_id,required,nullable"`
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceInboundCheck into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *WireTransferSimulationTransactionSourceInboundCheck) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceInboundCheck) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceInboundCheck{Amount:%s Currency:%s CheckNumber:%s CheckFrontImageFileID:%s CheckRearImageFileID:%s}", r.Amount, r.Currency, r.CheckNumber, r.CheckFrontImageFileID, r.CheckRearImageFileID)
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
// WireTransferSimulationTransactionSourceInboundInternationalACHTransfer into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceInboundInternationalACHTransfer{Amount:%s ForeignExchangeIndicator:%s ForeignExchangeReferenceIndicator:%s ForeignExchangeReference:%s DestinationCountryCode:%s DestinationCurrencyCode:%s ForeignPaymentAmount:%s ForeignTraceNumber:%s InternationalTransactionTypeCode:%s OriginatingCurrencyCode:%s OriginatingDepositoryFinancialInstitutionName:%s OriginatingDepositoryFinancialInstitutionIDQualifier:%s OriginatingDepositoryFinancialInstitutionID:%s OriginatingDepositoryFinancialInstitutionBranchCountry:%s OriginatorCity:%s OriginatorCompanyEntryDescription:%s OriginatorCountry:%s OriginatorIdentification:%s OriginatorName:%s OriginatorPostalCode:%s OriginatorStreetAddress:%s OriginatorStateOrProvince:%s PaymentRelatedInformation:%s PaymentRelatedInformation2:%s ReceiverIdentificationNumber:%s ReceiverStreetAddress:%s ReceiverCity:%s ReceiverStateOrProvince:%s ReceiverCountry:%s ReceiverPostalCode:%s ReceivingCompanyOrIndividualName:%s ReceivingDepositoryFinancialInstitutionName:%s ReceivingDepositoryFinancialInstitutionIDQualifier:%s ReceivingDepositoryFinancialInstitutionID:%s ReceivingDepositoryFinancialInstitutionCountry:%s TraceNumber:%s}", r.Amount, r.ForeignExchangeIndicator, r.ForeignExchangeReferenceIndicator, r.ForeignExchangeReference, r.DestinationCountryCode, r.DestinationCurrencyCode, r.ForeignPaymentAmount, r.ForeignTraceNumber, r.InternationalTransactionTypeCode, r.OriginatingCurrencyCode, r.OriginatingDepositoryFinancialInstitutionName, r.OriginatingDepositoryFinancialInstitutionIDQualifier, r.OriginatingDepositoryFinancialInstitutionID, r.OriginatingDepositoryFinancialInstitutionBranchCountry, r.OriginatorCity, r.OriginatorCompanyEntryDescription, r.OriginatorCountry, r.OriginatorIdentification, r.OriginatorName, r.OriginatorPostalCode, r.OriginatorStreetAddress, r.OriginatorStateOrProvince, r.PaymentRelatedInformation, r.PaymentRelatedInformation2, r.ReceiverIdentificationNumber, r.ReceiverStreetAddress, r.ReceiverCity, r.ReceiverStateOrProvince, r.ReceiverCountry, r.ReceiverPostalCode, r.ReceivingCompanyOrIndividualName, r.ReceivingDepositoryFinancialInstitutionName, r.ReceivingDepositoryFinancialInstitutionIDQualifier, r.ReceivingDepositoryFinancialInstitutionID, r.ReceivingDepositoryFinancialInstitutionCountry, r.TraceNumber)
}

type WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real Time Payments transfer.
	Currency fields.Field[WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency] `json:"currency,required"`
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
// WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation{Amount:%s Currency:%s CreditorName:%s DebtorName:%s DebtorAccountNumber:%s DebtorRoutingNumber:%s TransactionIdentification:%s RemittanceInformation:%s}", r.Amount, r.Currency, r.CreditorName, r.DebtorName, r.DebtorAccountNumber, r.DebtorRoutingNumber, r.TransactionIdentification, r.RemittanceInformation)
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
// WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal{Amount:%s Description:%s InputCycleDate:%s InputSequenceNumber:%s InputSource:%s InputMessageAccountabilityData:%s PreviousMessageInputMessageAccountabilityData:%s PreviousMessageInputCycleDate:%s PreviousMessageInputSequenceNumber:%s PreviousMessageInputSource:%s}", r.Amount, r.Description, r.InputCycleDate, r.InputSequenceNumber, r.InputSource, r.InputMessageAccountabilityData, r.PreviousMessageInputMessageAccountabilityData, r.PreviousMessageInputCycleDate, r.PreviousMessageInputSequenceNumber, r.PreviousMessageInputSource)
}

type WireTransferSimulationTransactionSourceInboundWireDrawdownPayment struct {
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
// WireTransferSimulationTransactionSourceInboundWireDrawdownPayment into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceInboundWireDrawdownPayment{Amount:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s BeneficiaryName:%s BeneficiaryReference:%s Description:%s InputMessageAccountabilityData:%s OriginatorAddressLine1:%s OriginatorAddressLine2:%s OriginatorAddressLine3:%s OriginatorName:%s OriginatorToBeneficiaryInformation:%s}", r.Amount, r.BeneficiaryAddressLine1, r.BeneficiaryAddressLine2, r.BeneficiaryAddressLine3, r.BeneficiaryName, r.BeneficiaryReference, r.Description, r.InputMessageAccountabilityData, r.OriginatorAddressLine1, r.OriginatorAddressLine2, r.OriginatorAddressLine3, r.OriginatorName, r.OriginatorToBeneficiaryInformation)
}

type WireTransferSimulationTransactionSourceInboundWireReversal struct {
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
// WireTransferSimulationTransactionSourceInboundWireReversal into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceInboundWireReversal) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceInboundWireReversal) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceInboundWireReversal{Amount:%s Description:%s InputCycleDate:%s InputSequenceNumber:%s InputSource:%s InputMessageAccountabilityData:%s PreviousMessageInputMessageAccountabilityData:%s PreviousMessageInputCycleDate:%s PreviousMessageInputSequenceNumber:%s PreviousMessageInputSource:%s ReceiverFinancialInstitutionInformation:%s FinancialInstitutionToFinancialInstitutionInformation:%s}", r.Amount, r.Description, r.InputCycleDate, r.InputSequenceNumber, r.InputSource, r.InputMessageAccountabilityData, r.PreviousMessageInputMessageAccountabilityData, r.PreviousMessageInputCycleDate, r.PreviousMessageInputSequenceNumber, r.PreviousMessageInputSource, r.ReceiverFinancialInstitutionInformation, r.FinancialInstitutionToFinancialInstitutionInformation)
}

type WireTransferSimulationTransactionSourceInboundWireTransfer struct {
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
// WireTransferSimulationTransactionSourceInboundWireTransfer into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceInboundWireTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceInboundWireTransfer) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceInboundWireTransfer{Amount:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s BeneficiaryName:%s BeneficiaryReference:%s Description:%s InputMessageAccountabilityData:%s OriginatorAddressLine1:%s OriginatorAddressLine2:%s OriginatorAddressLine3:%s OriginatorName:%s OriginatorToBeneficiaryInformationLine1:%s OriginatorToBeneficiaryInformationLine2:%s OriginatorToBeneficiaryInformationLine3:%s OriginatorToBeneficiaryInformationLine4:%s OriginatorToBeneficiaryInformation:%s}", r.Amount, r.BeneficiaryAddressLine1, r.BeneficiaryAddressLine2, r.BeneficiaryAddressLine3, r.BeneficiaryName, r.BeneficiaryReference, r.Description, r.InputMessageAccountabilityData, r.OriginatorAddressLine1, r.OriginatorAddressLine2, r.OriginatorAddressLine3, r.OriginatorName, r.OriginatorToBeneficiaryInformationLine1, r.OriginatorToBeneficiaryInformationLine2, r.OriginatorToBeneficiaryInformationLine3, r.OriginatorToBeneficiaryInformationLine4, r.OriginatorToBeneficiaryInformation)
}

type WireTransferSimulationTransactionSourceInterestPayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency fields.Field[WireTransferSimulationTransactionSourceInterestPaymentCurrency] `json:"currency,required"`
	// The start of the period for which this transaction paid interest.
	PeriodStart fields.Field[time.Time] `json:"period_start,required" format:"date-time"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd fields.Field[time.Time] `json:"period_end,required" format:"date-time"`
	// The account on which the interest was accrued.
	AccruedOnAccountID fields.Field[string] `json:"accrued_on_account_id,required,nullable"`
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceInterestPayment
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *WireTransferSimulationTransactionSourceInterestPayment) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceInterestPayment) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceInterestPayment{Amount:%s Currency:%s PeriodStart:%s PeriodEnd:%s AccruedOnAccountID:%s}", r.Amount, r.Currency, r.PeriodStart, r.PeriodEnd, r.AccruedOnAccountID)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency fields.Field[WireTransferSimulationTransactionSourceInternalSourceCurrency] `json:"currency,required"`
	Reason   fields.Field[WireTransferSimulationTransactionSourceInternalSourceReason]   `json:"reason,required"`
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceInternalSource
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *WireTransferSimulationTransactionSourceInternalSource) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceInternalSource) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceInternalSource{Amount:%s Currency:%s Reason:%s}", r.Amount, r.Currency, r.Reason)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
	// currency.
	Currency             fields.Field[WireTransferSimulationTransactionSourceCardRouteRefundCurrency] `json:"currency,required"`
	MerchantAcceptorID   fields.Field[string]                                                         `json:"merchant_acceptor_id,required"`
	MerchantCity         fields.Field[string]                                                         `json:"merchant_city,required,nullable"`
	MerchantCountry      fields.Field[string]                                                         `json:"merchant_country,required"`
	MerchantDescriptor   fields.Field[string]                                                         `json:"merchant_descriptor,required"`
	MerchantState        fields.Field[string]                                                         `json:"merchant_state,required,nullable"`
	MerchantCategoryCode fields.Field[string]                                                         `json:"merchant_category_code,required,nullable"`
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceCardRouteRefund
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *WireTransferSimulationTransactionSourceCardRouteRefund) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceCardRouteRefund) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCardRouteRefund{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", r.Amount, r.Currency, r.MerchantAcceptorID, r.MerchantCity, r.MerchantCountry, r.MerchantDescriptor, r.MerchantState, r.MerchantCategoryCode)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
	// currency.
	Currency             fields.Field[WireTransferSimulationTransactionSourceCardRouteSettlementCurrency] `json:"currency,required"`
	MerchantAcceptorID   fields.Field[string]                                                             `json:"merchant_acceptor_id,required"`
	MerchantCity         fields.Field[string]                                                             `json:"merchant_city,required,nullable"`
	MerchantCountry      fields.Field[string]                                                             `json:"merchant_country,required,nullable"`
	MerchantDescriptor   fields.Field[string]                                                             `json:"merchant_descriptor,required"`
	MerchantState        fields.Field[string]                                                             `json:"merchant_state,required,nullable"`
	MerchantCategoryCode fields.Field[string]                                                             `json:"merchant_category_code,required,nullable"`
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceCardRouteSettlement into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceCardRouteSettlement) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceCardRouteSettlement) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceCardRouteSettlement{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", r.Amount, r.Currency, r.MerchantAcceptorID, r.MerchantCity, r.MerchantCountry, r.MerchantDescriptor, r.MerchantState, r.MerchantCategoryCode)
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
	Originator fields.Field[string] `json:"originator,required"`
}

// MarshalJSON serializes WireTransferSimulationTransactionSourceSampleFunds into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *WireTransferSimulationTransactionSourceSampleFunds) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceSampleFunds) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceSampleFunds{Originator:%s}", r.Originator)
}

type WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention struct {
	// The transfer amount in USD cents.
	Amount             fields.Field[int64]  `json:"amount,required"`
	AccountNumber      fields.Field[string] `json:"account_number,required"`
	RoutingNumber      fields.Field[string] `json:"routing_number,required"`
	MessageToRecipient fields.Field[string] `json:"message_to_recipient,required"`
	TransferID         fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s TransferID:%s}", r.Amount, r.AccountNumber, r.RoutingNumber, r.MessageToRecipient, r.TransferID)
}

type WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection struct {
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection{TransferID:%s}", r.TransferID)
}

type WireTransferSimulationTransactionSourceWireTransferIntention struct {
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
// WireTransferSimulationTransactionSourceWireTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceWireTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceWireTransferIntention) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceWireTransferIntention{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s TransferID:%s}", r.Amount, r.AccountNumber, r.RoutingNumber, r.MessageToRecipient, r.TransferID)
}

type WireTransferSimulationTransactionSourceWireTransferRejection struct {
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// WireTransferSimulationTransactionSourceWireTransferRejection into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *WireTransferSimulationTransactionSourceWireTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSimulationTransactionSourceWireTransferRejection) String() (result string) {
	return fmt.Sprintf("&WireTransferSimulationTransactionSourceWireTransferRejection{TransferID:%s}", r.TransferID)
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
	AccountNumberID fields.Field[string] `json:"account_number_id,required"`
	// The transfer amount in cents. Must be positive.
	Amount fields.Field[int64] `json:"amount,required"`
	// The sending bank will set beneficiary_address_line1 in production. You can
	// simulate any value here.
	BeneficiaryAddressLine1 fields.Field[string] `json:"beneficiary_address_line1"`
	// The sending bank will set beneficiary_address_line2 in production. You can
	// simulate any value here.
	BeneficiaryAddressLine2 fields.Field[string] `json:"beneficiary_address_line2"`
	// The sending bank will set beneficiary_address_line3 in production. You can
	// simulate any value here.
	BeneficiaryAddressLine3 fields.Field[string] `json:"beneficiary_address_line3"`
	// The sending bank will set beneficiary_name in production. You can simulate any
	// value here.
	BeneficiaryName fields.Field[string] `json:"beneficiary_name"`
	// The sending bank will set beneficiary_reference in production. You can simulate
	// any value here.
	BeneficiaryReference fields.Field[string] `json:"beneficiary_reference"`
	// The sending bank will set originator_address_line1 in production. You can
	// simulate any value here.
	OriginatorAddressLine1 fields.Field[string] `json:"originator_address_line1"`
	// The sending bank will set originator_address_line2 in production. You can
	// simulate any value here.
	OriginatorAddressLine2 fields.Field[string] `json:"originator_address_line2"`
	// The sending bank will set originator_address_line3 in production. You can
	// simulate any value here.
	OriginatorAddressLine3 fields.Field[string] `json:"originator_address_line3"`
	// The sending bank will set originator_name in production. You can simulate any
	// value here.
	OriginatorName fields.Field[string] `json:"originator_name"`
	// The sending bank will set originator_to_beneficiary_information_line1 in
	// production. You can simulate any value here.
	OriginatorToBeneficiaryInformationLine1 fields.Field[string] `json:"originator_to_beneficiary_information_line1"`
	// The sending bank will set originator_to_beneficiary_information_line2 in
	// production. You can simulate any value here.
	OriginatorToBeneficiaryInformationLine2 fields.Field[string] `json:"originator_to_beneficiary_information_line2"`
	// The sending bank will set originator_to_beneficiary_information_line3 in
	// production. You can simulate any value here.
	OriginatorToBeneficiaryInformationLine3 fields.Field[string] `json:"originator_to_beneficiary_information_line3"`
	// The sending bank will set originator_to_beneficiary_information_line4 in
	// production. You can simulate any value here.
	OriginatorToBeneficiaryInformationLine4 fields.Field[string] `json:"originator_to_beneficiary_information_line4"`
}

// MarshalJSON serializes SimulateAWireTransferToYourAccountParameters into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *SimulateAWireTransferToYourAccountParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SimulateAWireTransferToYourAccountParameters) String() (result string) {
	return fmt.Sprintf("&SimulateAWireTransferToYourAccountParameters{AccountNumberID:%s Amount:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s BeneficiaryName:%s BeneficiaryReference:%s OriginatorAddressLine1:%s OriginatorAddressLine2:%s OriginatorAddressLine3:%s OriginatorName:%s OriginatorToBeneficiaryInformationLine1:%s OriginatorToBeneficiaryInformationLine2:%s OriginatorToBeneficiaryInformationLine3:%s OriginatorToBeneficiaryInformationLine4:%s}", r.AccountNumberID, r.Amount, r.BeneficiaryAddressLine1, r.BeneficiaryAddressLine2, r.BeneficiaryAddressLine3, r.BeneficiaryName, r.BeneficiaryReference, r.OriginatorAddressLine1, r.OriginatorAddressLine2, r.OriginatorAddressLine3, r.OriginatorName, r.OriginatorToBeneficiaryInformationLine1, r.OriginatorToBeneficiaryInformationLine2, r.OriginatorToBeneficiaryInformationLine3, r.OriginatorToBeneficiaryInformationLine4)
}
