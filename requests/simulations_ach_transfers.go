package requests

import (
	"fmt"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/fields"
)

type ACHTransferSimulation struct {
	// If the ACH Transfer attempt succeeds, this will contain the resulting
	// [Transaction](#transactions) object. The Transaction's `source` will be of
	// `category: inbound_ach_transfer`.
	Transaction fields.Field[ACHTransferSimulationTransaction] `json:"transaction,required,nullable"`
	// If the ACH Transfer attempt fails, this will contain the resulting
	// [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of `category: inbound_ach_transfer`.
	DeclinedTransaction fields.Field[ACHTransferSimulationDeclinedTransaction] `json:"declined_transaction,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_ach_transfer_simulation_result`.
	Type fields.Field[ACHTransferSimulationType] `json:"type,required"`
}

// MarshalJSON serializes ACHTransferSimulation into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulation) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulation) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulation{Transaction:%s DeclinedTransaction:%s Type:%s}", r.Transaction, r.DeclinedTransaction, r.Type)
}

type ACHTransferSimulationTransaction struct {
	// The identifier for the Account the Transaction belongs to.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transcation's
	// Account.
	Currency fields.Field[ACHTransferSimulationTransactionCurrency] `json:"currency,required"`
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
	RouteType fields.Field[ACHTransferSimulationTransactionRouteType] `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Transaction. Note that for backwards compatibility reasons, additional
	// undocumented keys may appear in this object. These should be treated as
	// deprecated and will be removed in the future.
	Source fields.Field[ACHTransferSimulationTransactionSource] `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type fields.Field[ACHTransferSimulationTransactionType] `json:"type,required"`
}

// MarshalJSON serializes ACHTransferSimulationTransaction into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransaction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransaction) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransaction{AccountID:%s Amount:%s Currency:%s CreatedAt:%s Description:%s ID:%s RouteID:%s RouteType:%s Source:%s Type:%s}", r.AccountID, r.Amount, r.Currency, r.CreatedAt, r.Description, r.ID, r.RouteID, r.RouteType, r.Source, r.Type)
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

type ACHTransferSimulationTransactionRouteType string

const (
	ACHTransferSimulationTransactionRouteTypeAccountNumber ACHTransferSimulationTransactionRouteType = "account_number"
	ACHTransferSimulationTransactionRouteTypeCard          ACHTransferSimulationTransactionRouteType = "card"
)

type ACHTransferSimulationTransactionSource struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category fields.Field[ACHTransferSimulationTransactionSourceCategory] `json:"category,required"`
	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention fields.Field[ACHTransferSimulationTransactionSourceAccountTransferIntention] `json:"account_transfer_intention,required,nullable"`
	// A ACH Check Conversion Return object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_check_conversion_return`.
	ACHCheckConversionReturn fields.Field[ACHTransferSimulationTransactionSourceACHCheckConversionReturn] `json:"ach_check_conversion_return,required,nullable"`
	// A ACH Check Conversion object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_check_conversion`.
	ACHCheckConversion fields.Field[ACHTransferSimulationTransactionSourceACHCheckConversion] `json:"ach_check_conversion,required,nullable"`
	// A ACH Transfer Intention object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention fields.Field[ACHTransferSimulationTransactionSourceACHTransferIntention] `json:"ach_transfer_intention,required,nullable"`
	// A ACH Transfer Rejection object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection fields.Field[ACHTransferSimulationTransactionSourceACHTransferRejection] `json:"ach_transfer_rejection,required,nullable"`
	// A ACH Transfer Return object. This field will be present in the JSON response if
	// and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn fields.Field[ACHTransferSimulationTransactionSourceACHTransferReturn] `json:"ach_transfer_return,required,nullable"`
	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance fields.Field[ACHTransferSimulationTransactionSourceCardDisputeAcceptance] `json:"card_dispute_acceptance,required,nullable"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund fields.Field[ACHTransferSimulationTransactionSourceCardRefund] `json:"card_refund,required,nullable"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement fields.Field[ACHTransferSimulationTransactionSourceCardSettlement] `json:"card_settlement,required,nullable"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance fields.Field[ACHTransferSimulationTransactionSourceCheckDepositAcceptance] `json:"check_deposit_acceptance,required,nullable"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn fields.Field[ACHTransferSimulationTransactionSourceCheckDepositReturn] `json:"check_deposit_return,required,nullable"`
	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention fields.Field[ACHTransferSimulationTransactionSourceCheckTransferIntention] `json:"check_transfer_intention,required,nullable"`
	// A Check Transfer Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_transfer_return`.
	CheckTransferReturn fields.Field[ACHTransferSimulationTransactionSourceCheckTransferReturn] `json:"check_transfer_return,required,nullable"`
	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection fields.Field[ACHTransferSimulationTransactionSourceCheckTransferRejection] `json:"check_transfer_rejection,required,nullable"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest fields.Field[ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest] `json:"check_transfer_stop_payment_request,required,nullable"`
	// A Dispute Resolution object. This field will be present in the JSON response if
	// and only if `category` is equal to `dispute_resolution`.
	DisputeResolution fields.Field[ACHTransferSimulationTransactionSourceDisputeResolution] `json:"dispute_resolution,required,nullable"`
	// A Empyreal Cash Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `empyreal_cash_deposit`.
	EmpyrealCashDeposit fields.Field[ACHTransferSimulationTransactionSourceEmpyrealCashDeposit] `json:"empyreal_cash_deposit,required,nullable"`
	// A Inbound ACH Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer fields.Field[ACHTransferSimulationTransactionSourceInboundACHTransfer] `json:"inbound_ach_transfer,required,nullable"`
	// A Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck fields.Field[ACHTransferSimulationTransactionSourceInboundCheck] `json:"inbound_check,required,nullable"`
	// A Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer fields.Field[ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer] `json:"inbound_international_ach_transfer,required,nullable"`
	// A Inbound Real Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation fields.Field[ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation] `json:"inbound_real_time_payments_transfer_confirmation,required,nullable"`
	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal fields.Field[ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal] `json:"inbound_wire_drawdown_payment_reversal,required,nullable"`
	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment fields.Field[ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment] `json:"inbound_wire_drawdown_payment,required,nullable"`
	// A Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal fields.Field[ACHTransferSimulationTransactionSourceInboundWireReversal] `json:"inbound_wire_reversal,required,nullable"`
	// A Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer fields.Field[ACHTransferSimulationTransactionSourceInboundWireTransfer] `json:"inbound_wire_transfer,required,nullable"`
	// A Interest Payment object. This field will be present in the JSON response if
	// and only if `category` is equal to `interest_payment`.
	InterestPayment fields.Field[ACHTransferSimulationTransactionSourceInterestPayment] `json:"interest_payment,required,nullable"`
	// A Internal Source object. This field will be present in the JSON response if and
	// only if `category` is equal to `internal_source`.
	InternalSource fields.Field[ACHTransferSimulationTransactionSourceInternalSource] `json:"internal_source,required,nullable"`
	// A Deprecated Card Refund object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_route_refund`.
	CardRouteRefund fields.Field[ACHTransferSimulationTransactionSourceCardRouteRefund] `json:"card_route_refund,required,nullable"`
	// A Deprecated Card Settlement object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_settlement`.
	CardRouteSettlement fields.Field[ACHTransferSimulationTransactionSourceCardRouteSettlement] `json:"card_route_settlement,required,nullable"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds fields.Field[ACHTransferSimulationTransactionSourceSampleFunds] `json:"sample_funds,required,nullable"`
	// A Wire Drawdown Payment Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_intention`.
	WireDrawdownPaymentIntention fields.Field[ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention] `json:"wire_drawdown_payment_intention,required,nullable"`
	// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_rejection`.
	WireDrawdownPaymentRejection fields.Field[ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection] `json:"wire_drawdown_payment_rejection,required,nullable"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention fields.Field[ACHTransferSimulationTransactionSourceWireTransferIntention] `json:"wire_transfer_intention,required,nullable"`
	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection fields.Field[ACHTransferSimulationTransactionSourceWireTransferRejection] `json:"wire_transfer_rejection,required,nullable"`
}

// MarshalJSON serializes ACHTransferSimulationTransactionSource into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSource) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSource) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSource{Category:%s AccountTransferIntention:%s ACHCheckConversionReturn:%s ACHCheckConversion:%s ACHTransferIntention:%s ACHTransferRejection:%s ACHTransferReturn:%s CardDisputeAcceptance:%s CardRefund:%s CardSettlement:%s CheckDepositAcceptance:%s CheckDepositReturn:%s CheckTransferIntention:%s CheckTransferReturn:%s CheckTransferRejection:%s CheckTransferStopPaymentRequest:%s DisputeResolution:%s EmpyrealCashDeposit:%s InboundACHTransfer:%s InboundCheck:%s InboundInternationalACHTransfer:%s InboundRealTimePaymentsTransferConfirmation:%s InboundWireDrawdownPaymentReversal:%s InboundWireDrawdownPayment:%s InboundWireReversal:%s InboundWireTransfer:%s InterestPayment:%s InternalSource:%s CardRouteRefund:%s CardRouteSettlement:%s SampleFunds:%s WireDrawdownPaymentIntention:%s WireDrawdownPaymentRejection:%s WireTransferIntention:%s WireTransferRejection:%s}", r.Category, r.AccountTransferIntention, r.ACHCheckConversionReturn, r.ACHCheckConversion, r.ACHTransferIntention, r.ACHTransferRejection, r.ACHTransferReturn, r.CardDisputeAcceptance, r.CardRefund, r.CardSettlement, r.CheckDepositAcceptance, r.CheckDepositReturn, r.CheckTransferIntention, r.CheckTransferReturn, r.CheckTransferRejection, r.CheckTransferStopPaymentRequest, r.DisputeResolution, r.EmpyrealCashDeposit, r.InboundACHTransfer, r.InboundCheck, r.InboundInternationalACHTransfer, r.InboundRealTimePaymentsTransferConfirmation, r.InboundWireDrawdownPaymentReversal, r.InboundWireDrawdownPayment, r.InboundWireReversal, r.InboundWireTransfer, r.InterestPayment, r.InternalSource, r.CardRouteRefund, r.CardRouteSettlement, r.SampleFunds, r.WireDrawdownPaymentIntention, r.WireDrawdownPaymentRejection, r.WireTransferIntention, r.WireTransferRejection)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency fields.Field[ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency] `json:"currency,required"`
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
// ACHTransferSimulationTransactionSourceAccountTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSourceAccountTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceAccountTransferIntention) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceAccountTransferIntention{Amount:%s Currency:%s Description:%s DestinationAccountID:%s SourceAccountID:%s TransferID:%s}", r.Amount, r.Currency, r.Description, r.DestinationAccountID, r.SourceAccountID, r.TransferID)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// Why the transfer was returned.
	ReturnReasonCode fields.Field[string] `json:"return_reason_code,required"`
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceACHCheckConversionReturn into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSourceACHCheckConversionReturn) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceACHCheckConversionReturn) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceACHCheckConversionReturn{Amount:%s ReturnReasonCode:%s}", r.Amount, r.ReturnReasonCode)
}

type ACHTransferSimulationTransactionSourceACHCheckConversion struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The identifier of the File containing an image of the returned check.
	FileID fields.Field[string] `json:"file_id,required"`
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceACHCheckConversion
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceACHCheckConversion) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceACHCheckConversion) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceACHCheckConversion{Amount:%s FileID:%s}", r.Amount, r.FileID)
}

type ACHTransferSimulationTransactionSourceACHTransferIntention struct {
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
// ACHTransferSimulationTransactionSourceACHTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSourceACHTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceACHTransferIntention) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceACHTransferIntention{Amount:%s AccountNumber:%s RoutingNumber:%s StatementDescriptor:%s TransferID:%s}", r.Amount, r.AccountNumber, r.RoutingNumber, r.StatementDescriptor, r.TransferID)
}

type ACHTransferSimulationTransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceACHTransferRejection into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSourceACHTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceACHTransferRejection) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceACHTransferRejection{TransferID:%s}", r.TransferID)
}

type ACHTransferSimulationTransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode fields.Field[ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode] `json:"return_reason_code,required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID fields.Field[string] `json:"transfer_id,required"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID fields.Field[string] `json:"transaction_id,required"`
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceACHTransferReturn
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceACHTransferReturn) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceACHTransferReturn) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceACHTransferReturn{CreatedAt:%s ReturnReasonCode:%s TransferID:%s TransactionID:%s}", r.CreatedAt, r.ReturnReasonCode, r.TransferID, r.TransactionID)
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
	AcceptedAt fields.Field[time.Time] `json:"accepted_at,required" format:"date-time"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID fields.Field[string] `json:"card_dispute_id,required"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID fields.Field[string] `json:"transaction_id,required,nullable"`
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceCardDisputeAcceptance into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSourceCardDisputeAcceptance) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceCardDisputeAcceptance) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCardDisputeAcceptance{AcceptedAt:%s CardDisputeID:%s TransactionID:%s}", r.AcceptedAt, r.CardDisputeID, r.TransactionID)
}

type ACHTransferSimulationTransactionSourceCardRefund struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[ACHTransferSimulationTransactionSourceCardRefundCurrency] `json:"currency,required"`
	// The identifier for the Transaction this refunds, if any.
	CardSettlementTransactionID fields.Field[string] `json:"card_settlement_transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type fields.Field[ACHTransferSimulationTransactionSourceCardRefundType] `json:"type,required"`
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceCardRefund into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceCardRefund) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceCardRefund) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCardRefund{Amount:%s Currency:%s CardSettlementTransactionID:%s Type:%s}", r.Amount, r.Currency, r.CardSettlementTransactionID, r.Type)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's settlement currency.
	Currency fields.Field[ACHTransferSimulationTransactionSourceCardSettlementCurrency] `json:"currency,required"`
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
	Type fields.Field[ACHTransferSimulationTransactionSourceCardSettlementType] `json:"type,required"`
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceCardSettlement into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceCardSettlement) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceCardSettlement) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCardSettlement{Amount:%s Currency:%s PresentmentAmount:%s PresentmentCurrency:%s MerchantCity:%s MerchantCountry:%s MerchantName:%s MerchantCategoryCode:%s MerchantState:%s PendingTransactionID:%s Type:%s}", r.Amount, r.Currency, r.PresentmentAmount, r.PresentmentCurrency, r.MerchantCity, r.MerchantCountry, r.MerchantName, r.MerchantCategoryCode, r.MerchantState, r.PendingTransactionID, r.Type)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency] `json:"currency,required"`
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
// ACHTransferSimulationTransactionSourceCheckDepositAcceptance into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSourceCheckDepositAcceptance) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceCheckDepositAcceptance) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCheckDepositAcceptance{Amount:%s Currency:%s AccountNumber:%s RoutingNumber:%s AuxiliaryOnUs:%s SerialNumber:%s CheckDepositID:%s}", r.Amount, r.Currency, r.AccountNumber, r.RoutingNumber, r.AuxiliaryOnUs, r.SerialNumber, r.CheckDepositID)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt fields.Field[time.Time] `json:"returned_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency] `json:"currency,required"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID fields.Field[string] `json:"check_deposit_id,required"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID fields.Field[string]                                                               `json:"transaction_id,required"`
	ReturnReason  fields.Field[ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason] `json:"return_reason,required"`
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceCheckDepositReturn
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceCheckDepositReturn) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceCheckDepositReturn) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCheckDepositReturn{Amount:%s ReturnedAt:%s Currency:%s CheckDepositID:%s TransactionID:%s ReturnReason:%s}", r.Amount, r.ReturnedAt, r.Currency, r.CheckDepositID, r.TransactionID, r.ReturnReason)
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
	Currency fields.Field[ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency] `json:"currency,required"`
	// The name that will be printed on the check.
	RecipientName fields.Field[string] `json:"recipient_name,required"`
	// The identifier of the Check Transfer with which this is associated.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceCheckTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceCheckTransferIntention) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCheckTransferIntention{AddressLine1:%s AddressLine2:%s AddressCity:%s AddressState:%s AddressZip:%s Amount:%s Currency:%s RecipientName:%s TransferID:%s}", r.AddressLine1, r.AddressLine2, r.AddressCity, r.AddressState, r.AddressZip, r.Amount, r.Currency, r.RecipientName, r.TransferID)
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
	TransferID fields.Field[string] `json:"transfer_id,required"`
	// If available, a document with additional information about the return.
	FileID fields.Field[string] `json:"file_id,required,nullable"`
	// The reason why the check was returned.
	Reason fields.Field[ACHTransferSimulationTransactionSourceCheckTransferReturnReason] `json:"reason,required"`
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceCheckTransferReturn
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceCheckTransferReturn) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceCheckTransferReturn) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCheckTransferReturn{TransferID:%s FileID:%s Reason:%s}", r.TransferID, r.FileID, r.Reason)
}

type ACHTransferSimulationTransactionSourceCheckTransferReturnReason string

const (
	ACHTransferSimulationTransactionSourceCheckTransferReturnReasonMailDeliveryFailure ACHTransferSimulationTransactionSourceCheckTransferReturnReason = "mail_delivery_failure"
	ACHTransferSimulationTransactionSourceCheckTransferReturnReasonRefusedByRecipient  ACHTransferSimulationTransactionSourceCheckTransferReturnReason = "refused_by_recipient"
)

type ACHTransferSimulationTransactionSourceCheckTransferRejection struct {
	// The identifier of the Check Transfer that led to this Transaction.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceCheckTransferRejection into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSourceCheckTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceCheckTransferRejection) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCheckTransferRejection{TransferID:%s}", r.TransferID)
}

type ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferID fields.Field[string] `json:"transfer_id,required"`
	// The transaction ID of the corresponding credit transaction.
	TransactionID fields.Field[string] `json:"transaction_id,required"`
	// The time the stop-payment was requested.
	RequestedAt fields.Field[time.Time] `json:"requested_at,required" format:"date-time"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type fields.Field[ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType] `json:"type,required"`
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest{TransferID:%s TransactionID:%s RequestedAt:%s Type:%s}", r.TransferID, r.TransactionID, r.RequestedAt, r.Type)
}

type ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType string

const (
	ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

type ACHTransferSimulationTransactionSourceDisputeResolution struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[ACHTransferSimulationTransactionSourceDisputeResolutionCurrency] `json:"currency,required"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionID fields.Field[string] `json:"disputed_transaction_id,required"`
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceDisputeResolution
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceDisputeResolution) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceDisputeResolution) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceDisputeResolution{Amount:%s Currency:%s DisputedTransactionID:%s}", r.Amount, r.Currency, r.DisputedTransactionID)
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
	Amount      fields.Field[int64]     `json:"amount,required"`
	BagID       fields.Field[string]    `json:"bag_id,required"`
	DepositDate fields.Field[time.Time] `json:"deposit_date,required" format:"date-time"`
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceEmpyrealCashDeposit
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceEmpyrealCashDeposit) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceEmpyrealCashDeposit) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceEmpyrealCashDeposit{Amount:%s BagID:%s DepositDate:%s}", r.Amount, r.BagID, r.DepositDate)
}

type ACHTransferSimulationTransactionSourceInboundACHTransfer struct {
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

// MarshalJSON serializes ACHTransferSimulationTransactionSourceInboundACHTransfer
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceInboundACHTransfer) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceInboundACHTransfer{Amount:%s OriginatorCompanyName:%s OriginatorCompanyDescriptiveDate:%s OriginatorCompanyDiscretionaryData:%s OriginatorCompanyEntryDescription:%s OriginatorCompanyID:%s ReceiverIDNumber:%s ReceiverName:%s TraceNumber:%s}", r.Amount, r.OriginatorCompanyName, r.OriginatorCompanyDescriptiveDate, r.OriginatorCompanyDiscretionaryData, r.OriginatorCompanyEntryDescription, r.OriginatorCompanyID, r.ReceiverIDNumber, r.ReceiverName, r.TraceNumber)
}

type ACHTransferSimulationTransactionSourceInboundCheck struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency              fields.Field[ACHTransferSimulationTransactionSourceInboundCheckCurrency] `json:"currency,required"`
	CheckNumber           fields.Field[string]                                                     `json:"check_number,required,nullable"`
	CheckFrontImageFileID fields.Field[string]                                                     `json:"check_front_image_file_id,required,nullable"`
	CheckRearImageFileID  fields.Field[string]                                                     `json:"check_rear_image_file_id,required,nullable"`
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceInboundCheck into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceInboundCheck) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceInboundCheck) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceInboundCheck{Amount:%s Currency:%s CheckNumber:%s CheckFrontImageFileID:%s CheckRearImageFileID:%s}", r.Amount, r.Currency, r.CheckNumber, r.CheckFrontImageFileID, r.CheckRearImageFileID)
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
// ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer{Amount:%s ForeignExchangeIndicator:%s ForeignExchangeReferenceIndicator:%s ForeignExchangeReference:%s DestinationCountryCode:%s DestinationCurrencyCode:%s ForeignPaymentAmount:%s ForeignTraceNumber:%s InternationalTransactionTypeCode:%s OriginatingCurrencyCode:%s OriginatingDepositoryFinancialInstitutionName:%s OriginatingDepositoryFinancialInstitutionIDQualifier:%s OriginatingDepositoryFinancialInstitutionID:%s OriginatingDepositoryFinancialInstitutionBranchCountry:%s OriginatorCity:%s OriginatorCompanyEntryDescription:%s OriginatorCountry:%s OriginatorIdentification:%s OriginatorName:%s OriginatorPostalCode:%s OriginatorStreetAddress:%s OriginatorStateOrProvince:%s PaymentRelatedInformation:%s PaymentRelatedInformation2:%s ReceiverIdentificationNumber:%s ReceiverStreetAddress:%s ReceiverCity:%s ReceiverStateOrProvince:%s ReceiverCountry:%s ReceiverPostalCode:%s ReceivingCompanyOrIndividualName:%s ReceivingDepositoryFinancialInstitutionName:%s ReceivingDepositoryFinancialInstitutionIDQualifier:%s ReceivingDepositoryFinancialInstitutionID:%s ReceivingDepositoryFinancialInstitutionCountry:%s TraceNumber:%s}", r.Amount, r.ForeignExchangeIndicator, r.ForeignExchangeReferenceIndicator, r.ForeignExchangeReference, r.DestinationCountryCode, r.DestinationCurrencyCode, r.ForeignPaymentAmount, r.ForeignTraceNumber, r.InternationalTransactionTypeCode, r.OriginatingCurrencyCode, r.OriginatingDepositoryFinancialInstitutionName, r.OriginatingDepositoryFinancialInstitutionIDQualifier, r.OriginatingDepositoryFinancialInstitutionID, r.OriginatingDepositoryFinancialInstitutionBranchCountry, r.OriginatorCity, r.OriginatorCompanyEntryDescription, r.OriginatorCountry, r.OriginatorIdentification, r.OriginatorName, r.OriginatorPostalCode, r.OriginatorStreetAddress, r.OriginatorStateOrProvince, r.PaymentRelatedInformation, r.PaymentRelatedInformation2, r.ReceiverIdentificationNumber, r.ReceiverStreetAddress, r.ReceiverCity, r.ReceiverStateOrProvince, r.ReceiverCountry, r.ReceiverPostalCode, r.ReceivingCompanyOrIndividualName, r.ReceivingDepositoryFinancialInstitutionName, r.ReceivingDepositoryFinancialInstitutionIDQualifier, r.ReceivingDepositoryFinancialInstitutionID, r.ReceivingDepositoryFinancialInstitutionCountry, r.TraceNumber)
}

type ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real Time Payments transfer.
	Currency fields.Field[ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency] `json:"currency,required"`
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
// ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation{Amount:%s Currency:%s CreditorName:%s DebtorName:%s DebtorAccountNumber:%s DebtorRoutingNumber:%s TransactionIdentification:%s RemittanceInformation:%s}", r.Amount, r.Currency, r.CreditorName, r.DebtorName, r.DebtorAccountNumber, r.DebtorRoutingNumber, r.TransactionIdentification, r.RemittanceInformation)
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
// ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal{Amount:%s Description:%s InputCycleDate:%s InputSequenceNumber:%s InputSource:%s InputMessageAccountabilityData:%s PreviousMessageInputMessageAccountabilityData:%s PreviousMessageInputCycleDate:%s PreviousMessageInputSequenceNumber:%s PreviousMessageInputSource:%s}", r.Amount, r.Description, r.InputCycleDate, r.InputSequenceNumber, r.InputSource, r.InputMessageAccountabilityData, r.PreviousMessageInputMessageAccountabilityData, r.PreviousMessageInputCycleDate, r.PreviousMessageInputSequenceNumber, r.PreviousMessageInputSource)
}

type ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment struct {
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
// ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment{Amount:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s BeneficiaryName:%s BeneficiaryReference:%s Description:%s InputMessageAccountabilityData:%s OriginatorAddressLine1:%s OriginatorAddressLine2:%s OriginatorAddressLine3:%s OriginatorName:%s OriginatorToBeneficiaryInformation:%s}", r.Amount, r.BeneficiaryAddressLine1, r.BeneficiaryAddressLine2, r.BeneficiaryAddressLine3, r.BeneficiaryName, r.BeneficiaryReference, r.Description, r.InputMessageAccountabilityData, r.OriginatorAddressLine1, r.OriginatorAddressLine2, r.OriginatorAddressLine3, r.OriginatorName, r.OriginatorToBeneficiaryInformation)
}

type ACHTransferSimulationTransactionSourceInboundWireReversal struct {
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

// MarshalJSON serializes ACHTransferSimulationTransactionSourceInboundWireReversal
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceInboundWireReversal) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceInboundWireReversal{Amount:%s Description:%s InputCycleDate:%s InputSequenceNumber:%s InputSource:%s InputMessageAccountabilityData:%s PreviousMessageInputMessageAccountabilityData:%s PreviousMessageInputCycleDate:%s PreviousMessageInputSequenceNumber:%s PreviousMessageInputSource:%s ReceiverFinancialInstitutionInformation:%s FinancialInstitutionToFinancialInstitutionInformation:%s}", r.Amount, r.Description, r.InputCycleDate, r.InputSequenceNumber, r.InputSource, r.InputMessageAccountabilityData, r.PreviousMessageInputMessageAccountabilityData, r.PreviousMessageInputCycleDate, r.PreviousMessageInputSequenceNumber, r.PreviousMessageInputSource, r.ReceiverFinancialInstitutionInformation, r.FinancialInstitutionToFinancialInstitutionInformation)
}

type ACHTransferSimulationTransactionSourceInboundWireTransfer struct {
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

// MarshalJSON serializes ACHTransferSimulationTransactionSourceInboundWireTransfer
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceInboundWireTransfer) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceInboundWireTransfer{Amount:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s BeneficiaryName:%s BeneficiaryReference:%s Description:%s InputMessageAccountabilityData:%s OriginatorAddressLine1:%s OriginatorAddressLine2:%s OriginatorAddressLine3:%s OriginatorName:%s OriginatorToBeneficiaryInformationLine1:%s OriginatorToBeneficiaryInformationLine2:%s OriginatorToBeneficiaryInformationLine3:%s OriginatorToBeneficiaryInformationLine4:%s OriginatorToBeneficiaryInformation:%s}", r.Amount, r.BeneficiaryAddressLine1, r.BeneficiaryAddressLine2, r.BeneficiaryAddressLine3, r.BeneficiaryName, r.BeneficiaryReference, r.Description, r.InputMessageAccountabilityData, r.OriginatorAddressLine1, r.OriginatorAddressLine2, r.OriginatorAddressLine3, r.OriginatorName, r.OriginatorToBeneficiaryInformationLine1, r.OriginatorToBeneficiaryInformationLine2, r.OriginatorToBeneficiaryInformationLine3, r.OriginatorToBeneficiaryInformationLine4, r.OriginatorToBeneficiaryInformation)
}

type ACHTransferSimulationTransactionSourceInterestPayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency fields.Field[ACHTransferSimulationTransactionSourceInterestPaymentCurrency] `json:"currency,required"`
	// The start of the period for which this transaction paid interest.
	PeriodStart fields.Field[time.Time] `json:"period_start,required" format:"date-time"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd fields.Field[time.Time] `json:"period_end,required" format:"date-time"`
	// The account on which the interest was accrued.
	AccruedOnAccountID fields.Field[string] `json:"accrued_on_account_id,required,nullable"`
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceInterestPayment
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceInterestPayment) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceInterestPayment) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceInterestPayment{Amount:%s Currency:%s PeriodStart:%s PeriodEnd:%s AccruedOnAccountID:%s}", r.Amount, r.Currency, r.PeriodStart, r.PeriodEnd, r.AccruedOnAccountID)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency fields.Field[ACHTransferSimulationTransactionSourceInternalSourceCurrency] `json:"currency,required"`
	Reason   fields.Field[ACHTransferSimulationTransactionSourceInternalSourceReason]   `json:"reason,required"`
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceInternalSource into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceInternalSource) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceInternalSource) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceInternalSource{Amount:%s Currency:%s Reason:%s}", r.Amount, r.Currency, r.Reason)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
	// currency.
	Currency             fields.Field[ACHTransferSimulationTransactionSourceCardRouteRefundCurrency] `json:"currency,required"`
	MerchantAcceptorID   fields.Field[string]                                                        `json:"merchant_acceptor_id,required"`
	MerchantCity         fields.Field[string]                                                        `json:"merchant_city,required,nullable"`
	MerchantCountry      fields.Field[string]                                                        `json:"merchant_country,required"`
	MerchantDescriptor   fields.Field[string]                                                        `json:"merchant_descriptor,required"`
	MerchantState        fields.Field[string]                                                        `json:"merchant_state,required,nullable"`
	MerchantCategoryCode fields.Field[string]                                                        `json:"merchant_category_code,required,nullable"`
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceCardRouteRefund
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceCardRouteRefund) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceCardRouteRefund) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCardRouteRefund{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", r.Amount, r.Currency, r.MerchantAcceptorID, r.MerchantCity, r.MerchantCountry, r.MerchantDescriptor, r.MerchantState, r.MerchantCategoryCode)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
	// currency.
	Currency             fields.Field[ACHTransferSimulationTransactionSourceCardRouteSettlementCurrency] `json:"currency,required"`
	MerchantAcceptorID   fields.Field[string]                                                            `json:"merchant_acceptor_id,required"`
	MerchantCity         fields.Field[string]                                                            `json:"merchant_city,required,nullable"`
	MerchantCountry      fields.Field[string]                                                            `json:"merchant_country,required,nullable"`
	MerchantDescriptor   fields.Field[string]                                                            `json:"merchant_descriptor,required"`
	MerchantState        fields.Field[string]                                                            `json:"merchant_state,required,nullable"`
	MerchantCategoryCode fields.Field[string]                                                            `json:"merchant_category_code,required,nullable"`
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceCardRouteSettlement
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationTransactionSourceCardRouteSettlement) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceCardRouteSettlement) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceCardRouteSettlement{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", r.Amount, r.Currency, r.MerchantAcceptorID, r.MerchantCity, r.MerchantCountry, r.MerchantDescriptor, r.MerchantState, r.MerchantCategoryCode)
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
	Originator fields.Field[string] `json:"originator,required"`
}

// MarshalJSON serializes ACHTransferSimulationTransactionSourceSampleFunds into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceSampleFunds) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceSampleFunds) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceSampleFunds{Originator:%s}", r.Originator)
}

type ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention struct {
	// The transfer amount in USD cents.
	Amount             fields.Field[int64]  `json:"amount,required"`
	AccountNumber      fields.Field[string] `json:"account_number,required"`
	RoutingNumber      fields.Field[string] `json:"routing_number,required"`
	MessageToRecipient fields.Field[string] `json:"message_to_recipient,required"`
	TransferID         fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s TransferID:%s}", r.Amount, r.AccountNumber, r.RoutingNumber, r.MessageToRecipient, r.TransferID)
}

type ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection struct {
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection{TransferID:%s}", r.TransferID)
}

type ACHTransferSimulationTransactionSourceWireTransferIntention struct {
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
// ACHTransferSimulationTransactionSourceWireTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSourceWireTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceWireTransferIntention) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceWireTransferIntention{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s TransferID:%s}", r.Amount, r.AccountNumber, r.RoutingNumber, r.MessageToRecipient, r.TransferID)
}

type ACHTransferSimulationTransactionSourceWireTransferRejection struct {
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// ACHTransferSimulationTransactionSourceWireTransferRejection into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationTransactionSourceWireTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationTransactionSourceWireTransferRejection) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationTransactionSourceWireTransferRejection{TransferID:%s}", r.TransferID)
}

type ACHTransferSimulationTransactionType string

const (
	ACHTransferSimulationTransactionTypeTransaction ACHTransferSimulationTransactionType = "transaction"
)

type ACHTransferSimulationDeclinedTransaction struct {
	// The identifier for the Account the Declined Transaction belongs to.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The Declined Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transcation's Account.
	Currency fields.Field[ACHTransferSimulationDeclinedTransactionCurrency] `json:"currency,required"`
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
	RouteType fields.Field[ACHTransferSimulationDeclinedTransactionRouteType] `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Declined Transaction. For example, for a card transaction this lists the
	// merchant's industry and location. Note that for backwards compatibility reasons,
	// additional undocumented keys may appear in this object. These should be treated
	// as deprecated and will be removed in the future.
	Source fields.Field[ACHTransferSimulationDeclinedTransactionSource] `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type fields.Field[ACHTransferSimulationDeclinedTransactionType] `json:"type,required"`
}

// MarshalJSON serializes ACHTransferSimulationDeclinedTransaction into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationDeclinedTransaction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationDeclinedTransaction) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationDeclinedTransaction{AccountID:%s Amount:%s Currency:%s CreatedAt:%s Description:%s ID:%s RouteID:%s RouteType:%s Source:%s Type:%s}", r.AccountID, r.Amount, r.Currency, r.CreatedAt, r.Description, r.ID, r.RouteID, r.RouteType, r.Source, r.Type)
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

type ACHTransferSimulationDeclinedTransactionRouteType string

const (
	ACHTransferSimulationDeclinedTransactionRouteTypeAccountNumber ACHTransferSimulationDeclinedTransactionRouteType = "account_number"
	ACHTransferSimulationDeclinedTransactionRouteTypeCard          ACHTransferSimulationDeclinedTransactionRouteType = "card"
)

type ACHTransferSimulationDeclinedTransactionSource struct {
	// The type of decline that took place. We may add additional possible values for
	// this enum over time; your application should be able to handle such additions
	// gracefully.
	Category fields.Field[ACHTransferSimulationDeclinedTransactionSourceCategory] `json:"category,required"`
	// A ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline fields.Field[ACHTransferSimulationDeclinedTransactionSourceACHDecline] `json:"ach_decline,required,nullable"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline fields.Field[ACHTransferSimulationDeclinedTransactionSourceCardDecline] `json:"card_decline,required,nullable"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline fields.Field[ACHTransferSimulationDeclinedTransactionSourceCheckDecline] `json:"check_decline,required,nullable"`
	// A Inbound Real Time Payments Transfer Decline object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline fields.Field[ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline] `json:"inbound_real_time_payments_transfer_decline,required,nullable"`
	// A International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline fields.Field[ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline] `json:"international_ach_decline,required,nullable"`
	// A Deprecated Card Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_decline`.
	CardRouteDecline fields.Field[ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline] `json:"card_route_decline,required,nullable"`
}

// MarshalJSON serializes ACHTransferSimulationDeclinedTransactionSource into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationDeclinedTransactionSource) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationDeclinedTransactionSource) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationDeclinedTransactionSource{Category:%s ACHDecline:%s CardDecline:%s CheckDecline:%s InboundRealTimePaymentsTransferDecline:%s InternationalACHDecline:%s CardRouteDecline:%s}", r.Category, r.ACHDecline, r.CardDecline, r.CheckDecline, r.InboundRealTimePaymentsTransferDecline, r.InternationalACHDecline, r.CardRouteDecline)
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
	Amount                             fields.Field[int64]  `json:"amount,required"`
	OriginatorCompanyName              fields.Field[string] `json:"originator_company_name,required"`
	OriginatorCompanyDescriptiveDate   fields.Field[string] `json:"originator_company_descriptive_date,required,nullable"`
	OriginatorCompanyDiscretionaryData fields.Field[string] `json:"originator_company_discretionary_data,required,nullable"`
	OriginatorCompanyID                fields.Field[string] `json:"originator_company_id,required"`
	// Why the ACH transfer was declined.
	Reason           fields.Field[ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason] `json:"reason,required"`
	ReceiverIDNumber fields.Field[string]                                                         `json:"receiver_id_number,required,nullable"`
	ReceiverName     fields.Field[string]                                                         `json:"receiver_name,required,nullable"`
	TraceNumber      fields.Field[string]                                                         `json:"trace_number,required"`
}

// MarshalJSON serializes ACHTransferSimulationDeclinedTransactionSourceACHDecline
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationDeclinedTransactionSourceACHDecline) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationDeclinedTransactionSourceACHDecline{Amount:%s OriginatorCompanyName:%s OriginatorCompanyDescriptiveDate:%s OriginatorCompanyDiscretionaryData:%s OriginatorCompanyID:%s Reason:%s ReceiverIDNumber:%s ReceiverName:%s TraceNumber:%s}", r.Amount, r.OriginatorCompanyName, r.OriginatorCompanyDescriptiveDate, r.OriginatorCompanyDiscretionaryData, r.OriginatorCompanyID, r.Reason, r.ReceiverIDNumber, r.ReceiverName, r.TraceNumber)
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
	Network fields.Field[ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetwork] `json:"network,required"`
	// Fields specific to the `network`
	NetworkDetails fields.Field[ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails] `json:"network_details,required"`
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency fields.Field[ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency] `json:"currency,required"`
	// Why the transaction was declined.
	Reason fields.Field[ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason] `json:"reason,required"`
	// The state the merchant resides in.
	MerchantState fields.Field[string] `json:"merchant_state,required,nullable"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID fields.Field[string] `json:"real_time_decision_id,required,nullable"`
	// If the authorization was attempted using a Digital Wallet Token (such as an
	// Apple Pay purchase), the identifier of the token that was used.
	DigitalWalletTokenID fields.Field[string] `json:"digital_wallet_token_id,required,nullable"`
}

// MarshalJSON serializes ACHTransferSimulationDeclinedTransactionSourceCardDecline
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationDeclinedTransactionSourceCardDecline) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationDeclinedTransactionSourceCardDecline{MerchantAcceptorID:%s MerchantDescriptor:%s MerchantCategoryCode:%s MerchantCity:%s MerchantCountry:%s Network:%s NetworkDetails:%s Amount:%s Currency:%s Reason:%s MerchantState:%s RealTimeDecisionID:%s DigitalWalletTokenID:%s}", r.MerchantAcceptorID, r.MerchantDescriptor, r.MerchantCategoryCode, r.MerchantCity, r.MerchantCountry, r.Network, r.NetworkDetails, r.Amount, r.Currency, r.Reason, r.MerchantState, r.RealTimeDecisionID, r.DigitalWalletTokenID)
}

type ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetwork string

const (
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkVisa ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetwork = "visa"
)

type ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa fields.Field[ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa] `json:"visa,required"`
}

// MarshalJSON serializes
// ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails{Visa:%s}", r.Visa)
}

type ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator fields.Field[ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator] `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode fields.Field[PointOfServiceEntryMode] `json:"point_of_service_entry_mode,required,nullable"`
}

// MarshalJSON serializes
// ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa{ElectronicCommerceIndicator:%s PointOfServiceEntryMode:%s}", r.ElectronicCommerceIndicator, r.PointOfServiceEntryMode)
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
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonInvalidPhysicalCard         ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "invalid_physical_card"
)

type ACHTransferSimulationDeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount        fields.Field[int64]  `json:"amount,required"`
	AuxiliaryOnUs fields.Field[string] `json:"auxiliary_on_us,required,nullable"`
	// Why the check was declined.
	Reason fields.Field[ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason] `json:"reason,required"`
}

// MarshalJSON serializes
// ACHTransferSimulationDeclinedTransactionSourceCheckDecline into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationDeclinedTransactionSourceCheckDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationDeclinedTransactionSourceCheckDecline) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationDeclinedTransactionSourceCheckDecline{Amount:%s AuxiliaryOnUs:%s Reason:%s}", r.Amount, r.AuxiliaryOnUs, r.Reason)
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
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonDuplicatePresentment  ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "duplicate_presentment"
)

type ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real Time Payments
	// transfer.
	Currency fields.Field[ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency] `json:"currency,required"`
	// Why the transfer was declined.
	Reason fields.Field[ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason] `json:"reason,required"`
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
// ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline{Amount:%s Currency:%s Reason:%s CreditorName:%s DebtorName:%s DebtorAccountNumber:%s DebtorRoutingNumber:%s TransactionIdentification:%s RemittanceInformation:%s}", r.Amount, r.Currency, r.Reason, r.CreditorName, r.DebtorName, r.DebtorAccountNumber, r.DebtorRoutingNumber, r.TransactionIdentification, r.RemittanceInformation)
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
// ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline{Amount:%s ForeignExchangeIndicator:%s ForeignExchangeReferenceIndicator:%s ForeignExchangeReference:%s DestinationCountryCode:%s DestinationCurrencyCode:%s ForeignPaymentAmount:%s ForeignTraceNumber:%s InternationalTransactionTypeCode:%s OriginatingCurrencyCode:%s OriginatingDepositoryFinancialInstitutionName:%s OriginatingDepositoryFinancialInstitutionIDQualifier:%s OriginatingDepositoryFinancialInstitutionID:%s OriginatingDepositoryFinancialInstitutionBranchCountry:%s OriginatorCity:%s OriginatorCompanyEntryDescription:%s OriginatorCountry:%s OriginatorIdentification:%s OriginatorName:%s OriginatorPostalCode:%s OriginatorStreetAddress:%s OriginatorStateOrProvince:%s PaymentRelatedInformation:%s PaymentRelatedInformation2:%s ReceiverIdentificationNumber:%s ReceiverStreetAddress:%s ReceiverCity:%s ReceiverStateOrProvince:%s ReceiverCountry:%s ReceiverPostalCode:%s ReceivingCompanyOrIndividualName:%s ReceivingDepositoryFinancialInstitutionName:%s ReceivingDepositoryFinancialInstitutionIDQualifier:%s ReceivingDepositoryFinancialInstitutionID:%s ReceivingDepositoryFinancialInstitutionCountry:%s TraceNumber:%s}", r.Amount, r.ForeignExchangeIndicator, r.ForeignExchangeReferenceIndicator, r.ForeignExchangeReference, r.DestinationCountryCode, r.DestinationCurrencyCode, r.ForeignPaymentAmount, r.ForeignTraceNumber, r.InternationalTransactionTypeCode, r.OriginatingCurrencyCode, r.OriginatingDepositoryFinancialInstitutionName, r.OriginatingDepositoryFinancialInstitutionIDQualifier, r.OriginatingDepositoryFinancialInstitutionID, r.OriginatingDepositoryFinancialInstitutionBranchCountry, r.OriginatorCity, r.OriginatorCompanyEntryDescription, r.OriginatorCountry, r.OriginatorIdentification, r.OriginatorName, r.OriginatorPostalCode, r.OriginatorStreetAddress, r.OriginatorStateOrProvince, r.PaymentRelatedInformation, r.PaymentRelatedInformation2, r.ReceiverIdentificationNumber, r.ReceiverStreetAddress, r.ReceiverCity, r.ReceiverStateOrProvince, r.ReceiverCountry, r.ReceiverPostalCode, r.ReceivingCompanyOrIndividualName, r.ReceivingDepositoryFinancialInstitutionName, r.ReceivingDepositoryFinancialInstitutionIDQualifier, r.ReceivingDepositoryFinancialInstitutionID, r.ReceivingDepositoryFinancialInstitutionCountry, r.TraceNumber)
}

type ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency             fields.Field[ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrency] `json:"currency,required"`
	MerchantAcceptorID   fields.Field[string]                                                                 `json:"merchant_acceptor_id,required"`
	MerchantCity         fields.Field[string]                                                                 `json:"merchant_city,required,nullable"`
	MerchantCountry      fields.Field[string]                                                                 `json:"merchant_country,required"`
	MerchantDescriptor   fields.Field[string]                                                                 `json:"merchant_descriptor,required"`
	MerchantState        fields.Field[string]                                                                 `json:"merchant_state,required,nullable"`
	MerchantCategoryCode fields.Field[string]                                                                 `json:"merchant_category_code,required,nullable"`
}

// MarshalJSON serializes
// ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) String() (result string) {
	return fmt.Sprintf("&ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", r.Amount, r.Currency, r.MerchantAcceptorID, r.MerchantCity, r.MerchantCountry, r.MerchantDescriptor, r.MerchantState, r.MerchantCategoryCode)
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
	AccountNumberID fields.Field[string] `json:"account_number_id,required"`
	// The transfer amount in cents. A positive amount originates a credit transfer
	// pushing funds to the receiving account. A negative amount originates a debit
	// transfer pulling funds from the receiving account.
	Amount fields.Field[int64] `json:"amount,required"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate fields.Field[string] `json:"company_descriptive_date"`
	// Data associated with the transfer set by the sender.
	CompanyDiscretionaryData fields.Field[string] `json:"company_discretionary_data"`
	// The description of the transfer set by the sender.
	CompanyEntryDescription fields.Field[string] `json:"company_entry_description"`
	// The name of the sender.
	CompanyName fields.Field[string] `json:"company_name"`
	// The sender's company id.
	CompanyID fields.Field[string] `json:"company_id"`
}

// MarshalJSON serializes SimulateAnACHTransferToYourAccountParameters into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *SimulateAnACHTransferToYourAccountParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SimulateAnACHTransferToYourAccountParameters) String() (result string) {
	return fmt.Sprintf("&SimulateAnACHTransferToYourAccountParameters{AccountNumberID:%s Amount:%s CompanyDescriptiveDate:%s CompanyDiscretionaryData:%s CompanyEntryDescription:%s CompanyName:%s CompanyID:%s}", r.AccountNumberID, r.Amount, r.CompanyDescriptiveDate, r.CompanyDiscretionaryData, r.CompanyEntryDescription, r.CompanyName, r.CompanyID)
}

type ReturnASandboxACHTransferParameters struct {
	// The reason why the Federal Reserve or destination bank returned this transfer.
	// Defaults to `no_account`.
	Reason fields.Field[ReturnASandboxACHTransferParametersReason] `json:"reason"`
}

// MarshalJSON serializes ReturnASandboxACHTransferParameters into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ReturnASandboxACHTransferParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ReturnASandboxACHTransferParameters) String() (result string) {
	return fmt.Sprintf("&ReturnASandboxACHTransferParameters{Reason:%s}", r.Reason)
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
