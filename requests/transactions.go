package requests

import (
	"fmt"
	"net/url"
	"time"

	"github.com/increase/increase-go/core"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type Transaction struct {
	// The identifier for the Account the Transaction belongs to.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transcation's
	// Account.
	Currency fields.Field[TransactionCurrency] `json:"currency,required"`
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
	Source fields.Field[TransactionSource] `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type fields.Field[TransactionType] `json:"type,required"`
}

// MarshalJSON serializes Transaction into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *Transaction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Transaction) String() (result string) {
	return fmt.Sprintf("&Transaction{AccountID:%s Amount:%s Currency:%s CreatedAt:%s Description:%s ID:%s RouteID:%s RouteType:%s Source:%s Type:%s}", r.AccountID, r.Amount, r.Currency, r.CreatedAt, r.Description, r.ID, r.RouteID, r.RouteType, r.Source, r.Type)
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
	Category fields.Field[TransactionSourceCategory] `json:"category,required"`
	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention fields.Field[TransactionSourceAccountTransferIntention] `json:"account_transfer_intention,required,nullable"`
	// A ACH Check Conversion Return object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_check_conversion_return`.
	ACHCheckConversionReturn fields.Field[TransactionSourceACHCheckConversionReturn] `json:"ach_check_conversion_return,required,nullable"`
	// A ACH Check Conversion object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_check_conversion`.
	ACHCheckConversion fields.Field[TransactionSourceACHCheckConversion] `json:"ach_check_conversion,required,nullable"`
	// A ACH Transfer Intention object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention fields.Field[TransactionSourceACHTransferIntention] `json:"ach_transfer_intention,required,nullable"`
	// A ACH Transfer Rejection object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection fields.Field[TransactionSourceACHTransferRejection] `json:"ach_transfer_rejection,required,nullable"`
	// A ACH Transfer Return object. This field will be present in the JSON response if
	// and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn fields.Field[TransactionSourceACHTransferReturn] `json:"ach_transfer_return,required,nullable"`
	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance fields.Field[TransactionSourceCardDisputeAcceptance] `json:"card_dispute_acceptance,required,nullable"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund fields.Field[TransactionSourceCardRefund] `json:"card_refund,required,nullable"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement fields.Field[TransactionSourceCardSettlement] `json:"card_settlement,required,nullable"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance fields.Field[TransactionSourceCheckDepositAcceptance] `json:"check_deposit_acceptance,required,nullable"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn fields.Field[TransactionSourceCheckDepositReturn] `json:"check_deposit_return,required,nullable"`
	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention fields.Field[TransactionSourceCheckTransferIntention] `json:"check_transfer_intention,required,nullable"`
	// A Check Transfer Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_transfer_return`.
	CheckTransferReturn fields.Field[TransactionSourceCheckTransferReturn] `json:"check_transfer_return,required,nullable"`
	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection fields.Field[TransactionSourceCheckTransferRejection] `json:"check_transfer_rejection,required,nullable"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest fields.Field[TransactionSourceCheckTransferStopPaymentRequest] `json:"check_transfer_stop_payment_request,required,nullable"`
	// A Dispute Resolution object. This field will be present in the JSON response if
	// and only if `category` is equal to `dispute_resolution`.
	DisputeResolution fields.Field[TransactionSourceDisputeResolution] `json:"dispute_resolution,required,nullable"`
	// A Empyreal Cash Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `empyreal_cash_deposit`.
	EmpyrealCashDeposit fields.Field[TransactionSourceEmpyrealCashDeposit] `json:"empyreal_cash_deposit,required,nullable"`
	// A Inbound ACH Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer fields.Field[TransactionSourceInboundACHTransfer] `json:"inbound_ach_transfer,required,nullable"`
	// A Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck fields.Field[TransactionSourceInboundCheck] `json:"inbound_check,required,nullable"`
	// A Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer fields.Field[TransactionSourceInboundInternationalACHTransfer] `json:"inbound_international_ach_transfer,required,nullable"`
	// A Inbound Real Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation fields.Field[TransactionSourceInboundRealTimePaymentsTransferConfirmation] `json:"inbound_real_time_payments_transfer_confirmation,required,nullable"`
	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal fields.Field[TransactionSourceInboundWireDrawdownPaymentReversal] `json:"inbound_wire_drawdown_payment_reversal,required,nullable"`
	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment fields.Field[TransactionSourceInboundWireDrawdownPayment] `json:"inbound_wire_drawdown_payment,required,nullable"`
	// A Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal fields.Field[TransactionSourceInboundWireReversal] `json:"inbound_wire_reversal,required,nullable"`
	// A Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer fields.Field[TransactionSourceInboundWireTransfer] `json:"inbound_wire_transfer,required,nullable"`
	// A Interest Payment object. This field will be present in the JSON response if
	// and only if `category` is equal to `interest_payment`.
	InterestPayment fields.Field[TransactionSourceInterestPayment] `json:"interest_payment,required,nullable"`
	// A Internal Source object. This field will be present in the JSON response if and
	// only if `category` is equal to `internal_source`.
	InternalSource fields.Field[TransactionSourceInternalSource] `json:"internal_source,required,nullable"`
	// A Deprecated Card Refund object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_route_refund`.
	CardRouteRefund fields.Field[TransactionSourceCardRouteRefund] `json:"card_route_refund,required,nullable"`
	// A Deprecated Card Settlement object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_settlement`.
	CardRouteSettlement fields.Field[TransactionSourceCardRouteSettlement] `json:"card_route_settlement,required,nullable"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds fields.Field[TransactionSourceSampleFunds] `json:"sample_funds,required,nullable"`
	// A Wire Drawdown Payment Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_intention`.
	WireDrawdownPaymentIntention fields.Field[TransactionSourceWireDrawdownPaymentIntention] `json:"wire_drawdown_payment_intention,required,nullable"`
	// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_rejection`.
	WireDrawdownPaymentRejection fields.Field[TransactionSourceWireDrawdownPaymentRejection] `json:"wire_drawdown_payment_rejection,required,nullable"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention fields.Field[TransactionSourceWireTransferIntention] `json:"wire_transfer_intention,required,nullable"`
	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection fields.Field[TransactionSourceWireTransferRejection] `json:"wire_transfer_rejection,required,nullable"`
}

// MarshalJSON serializes TransactionSource into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *TransactionSource) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSource) String() (result string) {
	return fmt.Sprintf("&TransactionSource{Category:%s AccountTransferIntention:%s ACHCheckConversionReturn:%s ACHCheckConversion:%s ACHTransferIntention:%s ACHTransferRejection:%s ACHTransferReturn:%s CardDisputeAcceptance:%s CardRefund:%s CardSettlement:%s CheckDepositAcceptance:%s CheckDepositReturn:%s CheckTransferIntention:%s CheckTransferReturn:%s CheckTransferRejection:%s CheckTransferStopPaymentRequest:%s DisputeResolution:%s EmpyrealCashDeposit:%s InboundACHTransfer:%s InboundCheck:%s InboundInternationalACHTransfer:%s InboundRealTimePaymentsTransferConfirmation:%s InboundWireDrawdownPaymentReversal:%s InboundWireDrawdownPayment:%s InboundWireReversal:%s InboundWireTransfer:%s InterestPayment:%s InternalSource:%s CardRouteRefund:%s CardRouteSettlement:%s SampleFunds:%s WireDrawdownPaymentIntention:%s WireDrawdownPaymentRejection:%s WireTransferIntention:%s WireTransferRejection:%s}", r.Category, r.AccountTransferIntention, r.ACHCheckConversionReturn, r.ACHCheckConversion, r.ACHTransferIntention, r.ACHTransferRejection, r.ACHTransferReturn, r.CardDisputeAcceptance, r.CardRefund, r.CardSettlement, r.CheckDepositAcceptance, r.CheckDepositReturn, r.CheckTransferIntention, r.CheckTransferReturn, r.CheckTransferRejection, r.CheckTransferStopPaymentRequest, r.DisputeResolution, r.EmpyrealCashDeposit, r.InboundACHTransfer, r.InboundCheck, r.InboundInternationalACHTransfer, r.InboundRealTimePaymentsTransferConfirmation, r.InboundWireDrawdownPaymentReversal, r.InboundWireDrawdownPayment, r.InboundWireReversal, r.InboundWireTransfer, r.InterestPayment, r.InternalSource, r.CardRouteRefund, r.CardRouteSettlement, r.SampleFunds, r.WireDrawdownPaymentIntention, r.WireDrawdownPaymentRejection, r.WireTransferIntention, r.WireTransferRejection)
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
	TransactionSourceCategoryInboundACHTransferReturnIntention           TransactionSourceCategory = "inbound_ach_transfer_return_intention"
	TransactionSourceCategoryInboundCheck                                TransactionSourceCategory = "inbound_check"
	TransactionSourceCategoryInboundInternationalACHTransfer             TransactionSourceCategory = "inbound_international_ach_transfer"
	TransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation TransactionSourceCategory = "inbound_real_time_payments_transfer_confirmation"
	TransactionSourceCategoryInboundWireDrawdownPaymentReversal          TransactionSourceCategory = "inbound_wire_drawdown_payment_reversal"
	TransactionSourceCategoryInboundWireDrawdownPayment                  TransactionSourceCategory = "inbound_wire_drawdown_payment"
	TransactionSourceCategoryInboundWireReversal                         TransactionSourceCategory = "inbound_wire_reversal"
	TransactionSourceCategoryInboundWireTransfer                         TransactionSourceCategory = "inbound_wire_transfer"
	TransactionSourceCategoryInterestPayment                             TransactionSourceCategory = "interest_payment"
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency fields.Field[TransactionSourceAccountTransferIntentionCurrency] `json:"currency,required"`
	// The description you chose to give the transfer.
	Description fields.Field[string] `json:"description,required"`
	// The identifier of the Account to where the Account Transfer was sent.
	DestinationAccountID fields.Field[string] `json:"destination_account_id,required"`
	// The identifier of the Account from where the Account Transfer was sent.
	SourceAccountID fields.Field[string] `json:"source_account_id,required"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes TransactionSourceAccountTransferIntention into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *TransactionSourceAccountTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceAccountTransferIntention) String() (result string) {
	return fmt.Sprintf("&TransactionSourceAccountTransferIntention{Amount:%s Currency:%s Description:%s DestinationAccountID:%s SourceAccountID:%s TransferID:%s}", r.Amount, r.Currency, r.Description, r.DestinationAccountID, r.SourceAccountID, r.TransferID)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// Why the transfer was returned.
	ReturnReasonCode fields.Field[string] `json:"return_reason_code,required"`
}

// MarshalJSON serializes TransactionSourceACHCheckConversionReturn into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *TransactionSourceACHCheckConversionReturn) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceACHCheckConversionReturn) String() (result string) {
	return fmt.Sprintf("&TransactionSourceACHCheckConversionReturn{Amount:%s ReturnReasonCode:%s}", r.Amount, r.ReturnReasonCode)
}

type TransactionSourceACHCheckConversion struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The identifier of the File containing an image of the returned check.
	FileID fields.Field[string] `json:"file_id,required"`
}

// MarshalJSON serializes TransactionSourceACHCheckConversion into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceACHCheckConversion) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceACHCheckConversion) String() (result string) {
	return fmt.Sprintf("&TransactionSourceACHCheckConversion{Amount:%s FileID:%s}", r.Amount, r.FileID)
}

type TransactionSourceACHTransferIntention struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount              fields.Field[int64]  `json:"amount,required"`
	AccountNumber       fields.Field[string] `json:"account_number,required"`
	RoutingNumber       fields.Field[string] `json:"routing_number,required"`
	StatementDescriptor fields.Field[string] `json:"statement_descriptor,required"`
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes TransactionSourceACHTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceACHTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceACHTransferIntention) String() (result string) {
	return fmt.Sprintf("&TransactionSourceACHTransferIntention{Amount:%s AccountNumber:%s RoutingNumber:%s StatementDescriptor:%s TransferID:%s}", r.Amount, r.AccountNumber, r.RoutingNumber, r.StatementDescriptor, r.TransferID)
}

type TransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes TransactionSourceACHTransferRejection into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceACHTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceACHTransferRejection) String() (result string) {
	return fmt.Sprintf("&TransactionSourceACHTransferRejection{TransferID:%s}", r.TransferID)
}

type TransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode fields.Field[TransactionSourceACHTransferReturnReturnReasonCode] `json:"return_reason_code,required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID fields.Field[string] `json:"transfer_id,required"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID fields.Field[string] `json:"transaction_id,required"`
}

// MarshalJSON serializes TransactionSourceACHTransferReturn into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceACHTransferReturn) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceACHTransferReturn) String() (result string) {
	return fmt.Sprintf("&TransactionSourceACHTransferReturn{CreatedAt:%s ReturnReasonCode:%s TransferID:%s TransactionID:%s}", r.CreatedAt, r.ReturnReasonCode, r.TransferID, r.TransactionID)
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
	AcceptedAt fields.Field[time.Time] `json:"accepted_at,required" format:"date-time"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID fields.Field[string] `json:"card_dispute_id,required"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID fields.Field[string] `json:"transaction_id,required,nullable"`
}

// MarshalJSON serializes TransactionSourceCardDisputeAcceptance into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceCardDisputeAcceptance) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceCardDisputeAcceptance) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCardDisputeAcceptance{AcceptedAt:%s CardDisputeID:%s TransactionID:%s}", r.AcceptedAt, r.CardDisputeID, r.TransactionID)
}

type TransactionSourceCardRefund struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[TransactionSourceCardRefundCurrency] `json:"currency,required"`
	// The identifier for the Transaction this refunds, if any.
	CardSettlementTransactionID fields.Field[string] `json:"card_settlement_transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type fields.Field[TransactionSourceCardRefundType] `json:"type,required"`
}

// MarshalJSON serializes TransactionSourceCardRefund into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *TransactionSourceCardRefund) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceCardRefund) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCardRefund{Amount:%s Currency:%s CardSettlementTransactionID:%s Type:%s}", r.Amount, r.Currency, r.CardSettlementTransactionID, r.Type)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's settlement currency.
	Currency fields.Field[TransactionSourceCardSettlementCurrency] `json:"currency,required"`
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
	Type fields.Field[TransactionSourceCardSettlementType] `json:"type,required"`
}

// MarshalJSON serializes TransactionSourceCardSettlement into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceCardSettlement) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceCardSettlement) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCardSettlement{Amount:%s Currency:%s PresentmentAmount:%s PresentmentCurrency:%s MerchantCity:%s MerchantCountry:%s MerchantName:%s MerchantCategoryCode:%s MerchantState:%s PendingTransactionID:%s Type:%s}", r.Amount, r.Currency, r.PresentmentAmount, r.PresentmentCurrency, r.MerchantCity, r.MerchantCountry, r.MerchantName, r.MerchantCategoryCode, r.MerchantState, r.PendingTransactionID, r.Type)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[TransactionSourceCheckDepositAcceptanceCurrency] `json:"currency,required"`
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

// MarshalJSON serializes TransactionSourceCheckDepositAcceptance into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceCheckDepositAcceptance) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceCheckDepositAcceptance) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCheckDepositAcceptance{Amount:%s Currency:%s AccountNumber:%s RoutingNumber:%s AuxiliaryOnUs:%s SerialNumber:%s CheckDepositID:%s}", r.Amount, r.Currency, r.AccountNumber, r.RoutingNumber, r.AuxiliaryOnUs, r.SerialNumber, r.CheckDepositID)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt fields.Field[time.Time] `json:"returned_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[TransactionSourceCheckDepositReturnCurrency] `json:"currency,required"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID fields.Field[string] `json:"check_deposit_id,required"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID fields.Field[string]                                          `json:"transaction_id,required"`
	ReturnReason  fields.Field[TransactionSourceCheckDepositReturnReturnReason] `json:"return_reason,required"`
}

// MarshalJSON serializes TransactionSourceCheckDepositReturn into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceCheckDepositReturn) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceCheckDepositReturn) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCheckDepositReturn{Amount:%s ReturnedAt:%s Currency:%s CheckDepositID:%s TransactionID:%s ReturnReason:%s}", r.Amount, r.ReturnedAt, r.Currency, r.CheckDepositID, r.TransactionID, r.ReturnReason)
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
	Currency fields.Field[TransactionSourceCheckTransferIntentionCurrency] `json:"currency,required"`
	// The name that will be printed on the check.
	RecipientName fields.Field[string] `json:"recipient_name,required"`
	// The identifier of the Check Transfer with which this is associated.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes TransactionSourceCheckTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceCheckTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceCheckTransferIntention) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCheckTransferIntention{AddressLine1:%s AddressLine2:%s AddressCity:%s AddressState:%s AddressZip:%s Amount:%s Currency:%s RecipientName:%s TransferID:%s}", r.AddressLine1, r.AddressLine2, r.AddressCity, r.AddressState, r.AddressZip, r.Amount, r.Currency, r.RecipientName, r.TransferID)
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
	TransferID fields.Field[string] `json:"transfer_id,required"`
	// If available, a document with additional information about the return.
	FileID fields.Field[string] `json:"file_id,required,nullable"`
	// The reason why the check was returned.
	Reason fields.Field[TransactionSourceCheckTransferReturnReason] `json:"reason,required"`
}

// MarshalJSON serializes TransactionSourceCheckTransferReturn into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceCheckTransferReturn) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceCheckTransferReturn) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCheckTransferReturn{TransferID:%s FileID:%s Reason:%s}", r.TransferID, r.FileID, r.Reason)
}

type TransactionSourceCheckTransferReturnReason string

const (
	TransactionSourceCheckTransferReturnReasonMailDeliveryFailure TransactionSourceCheckTransferReturnReason = "mail_delivery_failure"
	TransactionSourceCheckTransferReturnReasonRefusedByRecipient  TransactionSourceCheckTransferReturnReason = "refused_by_recipient"
)

type TransactionSourceCheckTransferRejection struct {
	// The identifier of the Check Transfer that led to this Transaction.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes TransactionSourceCheckTransferRejection into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceCheckTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceCheckTransferRejection) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCheckTransferRejection{TransferID:%s}", r.TransferID)
}

type TransactionSourceCheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferID fields.Field[string] `json:"transfer_id,required"`
	// The transaction ID of the corresponding credit transaction.
	TransactionID fields.Field[string] `json:"transaction_id,required"`
	// The time the stop-payment was requested.
	RequestedAt fields.Field[time.Time] `json:"requested_at,required" format:"date-time"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type fields.Field[TransactionSourceCheckTransferStopPaymentRequestType] `json:"type,required"`
}

// MarshalJSON serializes TransactionSourceCheckTransferStopPaymentRequest into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *TransactionSourceCheckTransferStopPaymentRequest) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceCheckTransferStopPaymentRequest) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCheckTransferStopPaymentRequest{TransferID:%s TransactionID:%s RequestedAt:%s Type:%s}", r.TransferID, r.TransactionID, r.RequestedAt, r.Type)
}

type TransactionSourceCheckTransferStopPaymentRequestType string

const (
	TransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest TransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

type TransactionSourceDisputeResolution struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[TransactionSourceDisputeResolutionCurrency] `json:"currency,required"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionID fields.Field[string] `json:"disputed_transaction_id,required"`
}

// MarshalJSON serializes TransactionSourceDisputeResolution into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceDisputeResolution) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceDisputeResolution) String() (result string) {
	return fmt.Sprintf("&TransactionSourceDisputeResolution{Amount:%s Currency:%s DisputedTransactionID:%s}", r.Amount, r.Currency, r.DisputedTransactionID)
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
	Amount      fields.Field[int64]     `json:"amount,required"`
	BagID       fields.Field[string]    `json:"bag_id,required"`
	DepositDate fields.Field[time.Time] `json:"deposit_date,required" format:"date-time"`
}

// MarshalJSON serializes TransactionSourceEmpyrealCashDeposit into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceEmpyrealCashDeposit) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceEmpyrealCashDeposit) String() (result string) {
	return fmt.Sprintf("&TransactionSourceEmpyrealCashDeposit{Amount:%s BagID:%s DepositDate:%s}", r.Amount, r.BagID, r.DepositDate)
}

type TransactionSourceInboundACHTransfer struct {
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

// MarshalJSON serializes TransactionSourceInboundACHTransfer into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceInboundACHTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceInboundACHTransfer) String() (result string) {
	return fmt.Sprintf("&TransactionSourceInboundACHTransfer{Amount:%s OriginatorCompanyName:%s OriginatorCompanyDescriptiveDate:%s OriginatorCompanyDiscretionaryData:%s OriginatorCompanyEntryDescription:%s OriginatorCompanyID:%s ReceiverIDNumber:%s ReceiverName:%s TraceNumber:%s}", r.Amount, r.OriginatorCompanyName, r.OriginatorCompanyDescriptiveDate, r.OriginatorCompanyDiscretionaryData, r.OriginatorCompanyEntryDescription, r.OriginatorCompanyID, r.ReceiverIDNumber, r.ReceiverName, r.TraceNumber)
}

type TransactionSourceInboundCheck struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency              fields.Field[TransactionSourceInboundCheckCurrency] `json:"currency,required"`
	CheckNumber           fields.Field[string]                                `json:"check_number,required,nullable"`
	CheckFrontImageFileID fields.Field[string]                                `json:"check_front_image_file_id,required,nullable"`
	CheckRearImageFileID  fields.Field[string]                                `json:"check_rear_image_file_id,required,nullable"`
}

// MarshalJSON serializes TransactionSourceInboundCheck into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceInboundCheck) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceInboundCheck) String() (result string) {
	return fmt.Sprintf("&TransactionSourceInboundCheck{Amount:%s Currency:%s CheckNumber:%s CheckFrontImageFileID:%s CheckRearImageFileID:%s}", r.Amount, r.Currency, r.CheckNumber, r.CheckFrontImageFileID, r.CheckRearImageFileID)
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

// MarshalJSON serializes TransactionSourceInboundInternationalACHTransfer into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *TransactionSourceInboundInternationalACHTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceInboundInternationalACHTransfer) String() (result string) {
	return fmt.Sprintf("&TransactionSourceInboundInternationalACHTransfer{Amount:%s ForeignExchangeIndicator:%s ForeignExchangeReferenceIndicator:%s ForeignExchangeReference:%s DestinationCountryCode:%s DestinationCurrencyCode:%s ForeignPaymentAmount:%s ForeignTraceNumber:%s InternationalTransactionTypeCode:%s OriginatingCurrencyCode:%s OriginatingDepositoryFinancialInstitutionName:%s OriginatingDepositoryFinancialInstitutionIDQualifier:%s OriginatingDepositoryFinancialInstitutionID:%s OriginatingDepositoryFinancialInstitutionBranchCountry:%s OriginatorCity:%s OriginatorCompanyEntryDescription:%s OriginatorCountry:%s OriginatorIdentification:%s OriginatorName:%s OriginatorPostalCode:%s OriginatorStreetAddress:%s OriginatorStateOrProvince:%s PaymentRelatedInformation:%s PaymentRelatedInformation2:%s ReceiverIdentificationNumber:%s ReceiverStreetAddress:%s ReceiverCity:%s ReceiverStateOrProvince:%s ReceiverCountry:%s ReceiverPostalCode:%s ReceivingCompanyOrIndividualName:%s ReceivingDepositoryFinancialInstitutionName:%s ReceivingDepositoryFinancialInstitutionIDQualifier:%s ReceivingDepositoryFinancialInstitutionID:%s ReceivingDepositoryFinancialInstitutionCountry:%s TraceNumber:%s}", r.Amount, r.ForeignExchangeIndicator, r.ForeignExchangeReferenceIndicator, r.ForeignExchangeReference, r.DestinationCountryCode, r.DestinationCurrencyCode, r.ForeignPaymentAmount, r.ForeignTraceNumber, r.InternationalTransactionTypeCode, r.OriginatingCurrencyCode, r.OriginatingDepositoryFinancialInstitutionName, r.OriginatingDepositoryFinancialInstitutionIDQualifier, r.OriginatingDepositoryFinancialInstitutionID, r.OriginatingDepositoryFinancialInstitutionBranchCountry, r.OriginatorCity, r.OriginatorCompanyEntryDescription, r.OriginatorCountry, r.OriginatorIdentification, r.OriginatorName, r.OriginatorPostalCode, r.OriginatorStreetAddress, r.OriginatorStateOrProvince, r.PaymentRelatedInformation, r.PaymentRelatedInformation2, r.ReceiverIdentificationNumber, r.ReceiverStreetAddress, r.ReceiverCity, r.ReceiverStateOrProvince, r.ReceiverCountry, r.ReceiverPostalCode, r.ReceivingCompanyOrIndividualName, r.ReceivingDepositoryFinancialInstitutionName, r.ReceivingDepositoryFinancialInstitutionIDQualifier, r.ReceivingDepositoryFinancialInstitutionID, r.ReceivingDepositoryFinancialInstitutionCountry, r.TraceNumber)
}

type TransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real Time Payments transfer.
	Currency fields.Field[TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency] `json:"currency,required"`
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
// TransactionSourceInboundRealTimePaymentsTransferConfirmation into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceInboundRealTimePaymentsTransferConfirmation) String() (result string) {
	return fmt.Sprintf("&TransactionSourceInboundRealTimePaymentsTransferConfirmation{Amount:%s Currency:%s CreditorName:%s DebtorName:%s DebtorAccountNumber:%s DebtorRoutingNumber:%s TransactionIdentification:%s RemittanceInformation:%s}", r.Amount, r.Currency, r.CreditorName, r.DebtorName, r.DebtorAccountNumber, r.DebtorRoutingNumber, r.TransactionIdentification, r.RemittanceInformation)
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

// MarshalJSON serializes TransactionSourceInboundWireDrawdownPaymentReversal into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceInboundWireDrawdownPaymentReversal) String() (result string) {
	return fmt.Sprintf("&TransactionSourceInboundWireDrawdownPaymentReversal{Amount:%s Description:%s InputCycleDate:%s InputSequenceNumber:%s InputSource:%s InputMessageAccountabilityData:%s PreviousMessageInputMessageAccountabilityData:%s PreviousMessageInputCycleDate:%s PreviousMessageInputSequenceNumber:%s PreviousMessageInputSource:%s}", r.Amount, r.Description, r.InputCycleDate, r.InputSequenceNumber, r.InputSource, r.InputMessageAccountabilityData, r.PreviousMessageInputMessageAccountabilityData, r.PreviousMessageInputCycleDate, r.PreviousMessageInputSequenceNumber, r.PreviousMessageInputSource)
}

type TransactionSourceInboundWireDrawdownPayment struct {
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

// MarshalJSON serializes TransactionSourceInboundWireDrawdownPayment into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *TransactionSourceInboundWireDrawdownPayment) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceInboundWireDrawdownPayment) String() (result string) {
	return fmt.Sprintf("&TransactionSourceInboundWireDrawdownPayment{Amount:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s BeneficiaryName:%s BeneficiaryReference:%s Description:%s InputMessageAccountabilityData:%s OriginatorAddressLine1:%s OriginatorAddressLine2:%s OriginatorAddressLine3:%s OriginatorName:%s OriginatorToBeneficiaryInformation:%s}", r.Amount, r.BeneficiaryAddressLine1, r.BeneficiaryAddressLine2, r.BeneficiaryAddressLine3, r.BeneficiaryName, r.BeneficiaryReference, r.Description, r.InputMessageAccountabilityData, r.OriginatorAddressLine1, r.OriginatorAddressLine2, r.OriginatorAddressLine3, r.OriginatorName, r.OriginatorToBeneficiaryInformation)
}

type TransactionSourceInboundWireReversal struct {
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

// MarshalJSON serializes TransactionSourceInboundWireReversal into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceInboundWireReversal) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceInboundWireReversal) String() (result string) {
	return fmt.Sprintf("&TransactionSourceInboundWireReversal{Amount:%s Description:%s InputCycleDate:%s InputSequenceNumber:%s InputSource:%s InputMessageAccountabilityData:%s PreviousMessageInputMessageAccountabilityData:%s PreviousMessageInputCycleDate:%s PreviousMessageInputSequenceNumber:%s PreviousMessageInputSource:%s ReceiverFinancialInstitutionInformation:%s FinancialInstitutionToFinancialInstitutionInformation:%s}", r.Amount, r.Description, r.InputCycleDate, r.InputSequenceNumber, r.InputSource, r.InputMessageAccountabilityData, r.PreviousMessageInputMessageAccountabilityData, r.PreviousMessageInputCycleDate, r.PreviousMessageInputSequenceNumber, r.PreviousMessageInputSource, r.ReceiverFinancialInstitutionInformation, r.FinancialInstitutionToFinancialInstitutionInformation)
}

type TransactionSourceInboundWireTransfer struct {
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

// MarshalJSON serializes TransactionSourceInboundWireTransfer into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceInboundWireTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceInboundWireTransfer) String() (result string) {
	return fmt.Sprintf("&TransactionSourceInboundWireTransfer{Amount:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s BeneficiaryName:%s BeneficiaryReference:%s Description:%s InputMessageAccountabilityData:%s OriginatorAddressLine1:%s OriginatorAddressLine2:%s OriginatorAddressLine3:%s OriginatorName:%s OriginatorToBeneficiaryInformationLine1:%s OriginatorToBeneficiaryInformationLine2:%s OriginatorToBeneficiaryInformationLine3:%s OriginatorToBeneficiaryInformationLine4:%s OriginatorToBeneficiaryInformation:%s}", r.Amount, r.BeneficiaryAddressLine1, r.BeneficiaryAddressLine2, r.BeneficiaryAddressLine3, r.BeneficiaryName, r.BeneficiaryReference, r.Description, r.InputMessageAccountabilityData, r.OriginatorAddressLine1, r.OriginatorAddressLine2, r.OriginatorAddressLine3, r.OriginatorName, r.OriginatorToBeneficiaryInformationLine1, r.OriginatorToBeneficiaryInformationLine2, r.OriginatorToBeneficiaryInformationLine3, r.OriginatorToBeneficiaryInformationLine4, r.OriginatorToBeneficiaryInformation)
}

type TransactionSourceInterestPayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency fields.Field[TransactionSourceInterestPaymentCurrency] `json:"currency,required"`
	// The start of the period for which this transaction paid interest.
	PeriodStart fields.Field[time.Time] `json:"period_start,required" format:"date-time"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd fields.Field[time.Time] `json:"period_end,required" format:"date-time"`
	// The account on which the interest was accrued.
	AccruedOnAccountID fields.Field[string] `json:"accrued_on_account_id,required,nullable"`
}

// MarshalJSON serializes TransactionSourceInterestPayment into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceInterestPayment) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceInterestPayment) String() (result string) {
	return fmt.Sprintf("&TransactionSourceInterestPayment{Amount:%s Currency:%s PeriodStart:%s PeriodEnd:%s AccruedOnAccountID:%s}", r.Amount, r.Currency, r.PeriodStart, r.PeriodEnd, r.AccruedOnAccountID)
}

type TransactionSourceInterestPaymentCurrency string

const (
	TransactionSourceInterestPaymentCurrencyCad TransactionSourceInterestPaymentCurrency = "CAD"
	TransactionSourceInterestPaymentCurrencyChf TransactionSourceInterestPaymentCurrency = "CHF"
	TransactionSourceInterestPaymentCurrencyEur TransactionSourceInterestPaymentCurrency = "EUR"
	TransactionSourceInterestPaymentCurrencyGbp TransactionSourceInterestPaymentCurrency = "GBP"
	TransactionSourceInterestPaymentCurrencyJpy TransactionSourceInterestPaymentCurrency = "JPY"
	TransactionSourceInterestPaymentCurrencyUsd TransactionSourceInterestPaymentCurrency = "USD"
)

type TransactionSourceInternalSource struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency fields.Field[TransactionSourceInternalSourceCurrency] `json:"currency,required"`
	Reason   fields.Field[TransactionSourceInternalSourceReason]   `json:"reason,required"`
}

// MarshalJSON serializes TransactionSourceInternalSource into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceInternalSource) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceInternalSource) String() (result string) {
	return fmt.Sprintf("&TransactionSourceInternalSource{Amount:%s Currency:%s Reason:%s}", r.Amount, r.Currency, r.Reason)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
	// currency.
	Currency             fields.Field[TransactionSourceCardRouteRefundCurrency] `json:"currency,required"`
	MerchantAcceptorID   fields.Field[string]                                   `json:"merchant_acceptor_id,required"`
	MerchantCity         fields.Field[string]                                   `json:"merchant_city,required,nullable"`
	MerchantCountry      fields.Field[string]                                   `json:"merchant_country,required"`
	MerchantDescriptor   fields.Field[string]                                   `json:"merchant_descriptor,required"`
	MerchantState        fields.Field[string]                                   `json:"merchant_state,required,nullable"`
	MerchantCategoryCode fields.Field[string]                                   `json:"merchant_category_code,required,nullable"`
}

// MarshalJSON serializes TransactionSourceCardRouteRefund into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceCardRouteRefund) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceCardRouteRefund) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCardRouteRefund{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", r.Amount, r.Currency, r.MerchantAcceptorID, r.MerchantCity, r.MerchantCountry, r.MerchantDescriptor, r.MerchantState, r.MerchantCategoryCode)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
	// currency.
	Currency             fields.Field[TransactionSourceCardRouteSettlementCurrency] `json:"currency,required"`
	MerchantAcceptorID   fields.Field[string]                                       `json:"merchant_acceptor_id,required"`
	MerchantCity         fields.Field[string]                                       `json:"merchant_city,required,nullable"`
	MerchantCountry      fields.Field[string]                                       `json:"merchant_country,required,nullable"`
	MerchantDescriptor   fields.Field[string]                                       `json:"merchant_descriptor,required"`
	MerchantState        fields.Field[string]                                       `json:"merchant_state,required,nullable"`
	MerchantCategoryCode fields.Field[string]                                       `json:"merchant_category_code,required,nullable"`
}

// MarshalJSON serializes TransactionSourceCardRouteSettlement into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceCardRouteSettlement) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceCardRouteSettlement) String() (result string) {
	return fmt.Sprintf("&TransactionSourceCardRouteSettlement{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", r.Amount, r.Currency, r.MerchantAcceptorID, r.MerchantCity, r.MerchantCountry, r.MerchantDescriptor, r.MerchantState, r.MerchantCategoryCode)
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
	Originator fields.Field[string] `json:"originator,required"`
}

// MarshalJSON serializes TransactionSourceSampleFunds into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *TransactionSourceSampleFunds) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceSampleFunds) String() (result string) {
	return fmt.Sprintf("&TransactionSourceSampleFunds{Originator:%s}", r.Originator)
}

type TransactionSourceWireDrawdownPaymentIntention struct {
	// The transfer amount in USD cents.
	Amount             fields.Field[int64]  `json:"amount,required"`
	AccountNumber      fields.Field[string] `json:"account_number,required"`
	RoutingNumber      fields.Field[string] `json:"routing_number,required"`
	MessageToRecipient fields.Field[string] `json:"message_to_recipient,required"`
	TransferID         fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes TransactionSourceWireDrawdownPaymentIntention into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *TransactionSourceWireDrawdownPaymentIntention) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceWireDrawdownPaymentIntention) String() (result string) {
	return fmt.Sprintf("&TransactionSourceWireDrawdownPaymentIntention{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s TransferID:%s}", r.Amount, r.AccountNumber, r.RoutingNumber, r.MessageToRecipient, r.TransferID)
}

type TransactionSourceWireDrawdownPaymentRejection struct {
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes TransactionSourceWireDrawdownPaymentRejection into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *TransactionSourceWireDrawdownPaymentRejection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceWireDrawdownPaymentRejection) String() (result string) {
	return fmt.Sprintf("&TransactionSourceWireDrawdownPaymentRejection{TransferID:%s}", r.TransferID)
}

type TransactionSourceWireTransferIntention struct {
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

// MarshalJSON serializes TransactionSourceWireTransferIntention into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceWireTransferIntention) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceWireTransferIntention) String() (result string) {
	return fmt.Sprintf("&TransactionSourceWireTransferIntention{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s TransferID:%s}", r.Amount, r.AccountNumber, r.RoutingNumber, r.MessageToRecipient, r.TransferID)
}

type TransactionSourceWireTransferRejection struct {
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes TransactionSourceWireTransferRejection into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *TransactionSourceWireTransferRejection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TransactionSourceWireTransferRejection) String() (result string) {
	return fmt.Sprintf("&TransactionSourceWireTransferRejection{TransferID:%s}", r.TransferID)
}

type TransactionType string

const (
	TransactionTypeTransaction TransactionType = "transaction"
)

type TransactionListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter Transactions for those belonging to the specified Account.
	AccountID fields.Field[string] `query:"account_id"`
	// Filter Transactions for those belonging to the specified route.
	RouteID   fields.Field[string]                         `query:"route_id"`
	CreatedAt fields.Field[TransactionListParamsCreatedAt] `query:"created_at"`
	Category  fields.Field[TransactionListParamsCategory]  `query:"category"`
}

// URLQuery serializes TransactionListParams into a url.Values of the query
// parameters associated with this value
func (r *TransactionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r TransactionListParams) String() (result string) {
	return fmt.Sprintf("&TransactionListParams{Cursor:%s Limit:%s AccountID:%s RouteID:%s CreatedAt:%s Category:%s}", r.Cursor, r.Limit, r.AccountID, r.RouteID, r.CreatedAt, r.Category)
}

type TransactionListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After fields.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before fields.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter fields.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore fields.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes TransactionListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r *TransactionListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r TransactionListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&TransactionListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}

type TransactionListParamsCategory struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In fields.Field[[]TransactionListParamsCategoryIn] `query:"in"`
}

// URLQuery serializes TransactionListParamsCategory into a url.Values of the query
// parameters associated with this value
func (r *TransactionListParamsCategory) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r TransactionListParamsCategory) String() (result string) {
	return fmt.Sprintf("&TransactionListParamsCategory{In:%s}", core.Fmt(r.In))
}

type TransactionListParamsCategoryIn string

const (
	TransactionListParamsCategoryInAccountTransferIntention                    TransactionListParamsCategoryIn = "account_transfer_intention"
	TransactionListParamsCategoryInACHCheckConversionReturn                    TransactionListParamsCategoryIn = "ach_check_conversion_return"
	TransactionListParamsCategoryInACHCheckConversion                          TransactionListParamsCategoryIn = "ach_check_conversion"
	TransactionListParamsCategoryInACHTransferIntention                        TransactionListParamsCategoryIn = "ach_transfer_intention"
	TransactionListParamsCategoryInACHTransferRejection                        TransactionListParamsCategoryIn = "ach_transfer_rejection"
	TransactionListParamsCategoryInACHTransferReturn                           TransactionListParamsCategoryIn = "ach_transfer_return"
	TransactionListParamsCategoryInCardDisputeAcceptance                       TransactionListParamsCategoryIn = "card_dispute_acceptance"
	TransactionListParamsCategoryInCardRefund                                  TransactionListParamsCategoryIn = "card_refund"
	TransactionListParamsCategoryInCardSettlement                              TransactionListParamsCategoryIn = "card_settlement"
	TransactionListParamsCategoryInCheckDepositAcceptance                      TransactionListParamsCategoryIn = "check_deposit_acceptance"
	TransactionListParamsCategoryInCheckDepositReturn                          TransactionListParamsCategoryIn = "check_deposit_return"
	TransactionListParamsCategoryInCheckTransferIntention                      TransactionListParamsCategoryIn = "check_transfer_intention"
	TransactionListParamsCategoryInCheckTransferReturn                         TransactionListParamsCategoryIn = "check_transfer_return"
	TransactionListParamsCategoryInCheckTransferRejection                      TransactionListParamsCategoryIn = "check_transfer_rejection"
	TransactionListParamsCategoryInCheckTransferStopPaymentRequest             TransactionListParamsCategoryIn = "check_transfer_stop_payment_request"
	TransactionListParamsCategoryInDisputeResolution                           TransactionListParamsCategoryIn = "dispute_resolution"
	TransactionListParamsCategoryInEmpyrealCashDeposit                         TransactionListParamsCategoryIn = "empyreal_cash_deposit"
	TransactionListParamsCategoryInInboundACHTransfer                          TransactionListParamsCategoryIn = "inbound_ach_transfer"
	TransactionListParamsCategoryInInboundACHTransferReturnIntention           TransactionListParamsCategoryIn = "inbound_ach_transfer_return_intention"
	TransactionListParamsCategoryInInboundCheck                                TransactionListParamsCategoryIn = "inbound_check"
	TransactionListParamsCategoryInInboundInternationalACHTransfer             TransactionListParamsCategoryIn = "inbound_international_ach_transfer"
	TransactionListParamsCategoryInInboundRealTimePaymentsTransferConfirmation TransactionListParamsCategoryIn = "inbound_real_time_payments_transfer_confirmation"
	TransactionListParamsCategoryInInboundWireDrawdownPaymentReversal          TransactionListParamsCategoryIn = "inbound_wire_drawdown_payment_reversal"
	TransactionListParamsCategoryInInboundWireDrawdownPayment                  TransactionListParamsCategoryIn = "inbound_wire_drawdown_payment"
	TransactionListParamsCategoryInInboundWireReversal                         TransactionListParamsCategoryIn = "inbound_wire_reversal"
	TransactionListParamsCategoryInInboundWireTransfer                         TransactionListParamsCategoryIn = "inbound_wire_transfer"
	TransactionListParamsCategoryInInterestPayment                             TransactionListParamsCategoryIn = "interest_payment"
	TransactionListParamsCategoryInInternalGeneralLedgerTransaction            TransactionListParamsCategoryIn = "internal_general_ledger_transaction"
	TransactionListParamsCategoryInInternalSource                              TransactionListParamsCategoryIn = "internal_source"
	TransactionListParamsCategoryInCardRouteRefund                             TransactionListParamsCategoryIn = "card_route_refund"
	TransactionListParamsCategoryInCardRouteSettlement                         TransactionListParamsCategoryIn = "card_route_settlement"
	TransactionListParamsCategoryInRealTimePaymentsTransferAcknowledgement     TransactionListParamsCategoryIn = "real_time_payments_transfer_acknowledgement"
	TransactionListParamsCategoryInSampleFunds                                 TransactionListParamsCategoryIn = "sample_funds"
	TransactionListParamsCategoryInWireDrawdownPaymentIntention                TransactionListParamsCategoryIn = "wire_drawdown_payment_intention"
	TransactionListParamsCategoryInWireDrawdownPaymentRejection                TransactionListParamsCategoryIn = "wire_drawdown_payment_rejection"
	TransactionListParamsCategoryInWireTransferIntention                       TransactionListParamsCategoryIn = "wire_transfer_intention"
	TransactionListParamsCategoryInWireTransferRejection                       TransactionListParamsCategoryIn = "wire_transfer_rejection"
	TransactionListParamsCategoryInOther                                       TransactionListParamsCategoryIn = "other"
)
