package responses

import (
	"time"

	apijson "github.com/increase/increase-go/internal/json"
)

type InboundRealTimePaymentsTransferSimulationResult struct {
	// If the Real Time Payments Transfer attempt succeeds, this will contain the
	// resulting [Transaction](#transactions) object. The Transaction's `source` will
	// be of `category: inbound_real_time_payments_transfer_confirmation`.
	Transaction InboundRealTimePaymentsTransferSimulationResultTransaction `json:"transaction,required,nullable"`
	// If the Real Time Payments Transfer attempt fails, this will contain the
	// resulting [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of
	// `category: inbound_real_time_payments_transfer_decline`.
	DeclinedTransaction InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction `json:"declined_transaction,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_real_time_payments_transfer_simulation_result`.
	Type InboundRealTimePaymentsTransferSimulationResultType `json:"type,required"`
	JSON InboundRealTimePaymentsTransferSimulationResultJSON
}

type InboundRealTimePaymentsTransferSimulationResultJSON struct {
	Transaction         apijson.Metadata
	DeclinedTransaction apijson.Metadata
	Type                apijson.Metadata
	raw                 string
	Extras              map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResult using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransaction struct {
	// The identifier for the Account the Transaction belongs to.
	AccountID string `json:"account_id,required"`
	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transcation's
	// Account.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionCurrency `json:"currency,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// For a Transaction related to a transfer, this is the description you provide.
	// For a Transaction related to a payment, this is the description the vendor
	// provides.
	Description string `json:"description,required"`
	// The Transaction identifier.
	ID string `json:"id,required"`
	// The identifier for the route this Transaction came through. Routes are things
	// like cards and ACH details.
	RouteID string `json:"route_id,required,nullable"`
	// The type of the route this Transaction came through.
	RouteType InboundRealTimePaymentsTransferSimulationResultTransactionRouteType `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Transaction. Note that for backwards compatibility reasons, additional
	// undocumented keys may appear in this object. These should be treated as
	// deprecated and will be removed in the future.
	Source InboundRealTimePaymentsTransferSimulationResultTransactionSource `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type InboundRealTimePaymentsTransferSimulationResultTransactionType `json:"type,required"`
	JSON InboundRealTimePaymentsTransferSimulationResultTransactionJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionJSON struct {
	AccountID   apijson.Metadata
	Amount      apijson.Metadata
	Currency    apijson.Metadata
	CreatedAt   apijson.Metadata
	Description apijson.Metadata
	ID          apijson.Metadata
	RouteID     apijson.Metadata
	RouteType   apijson.Metadata
	Source      apijson.Metadata
	Type        apijson.Metadata
	raw         string
	Extras      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransaction using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Category InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory `json:"category,required"`
	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention `json:"account_transfer_intention,required,nullable"`
	// A ACH Check Conversion Return object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_check_conversion_return`.
	ACHCheckConversionReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversionReturn `json:"ach_check_conversion_return,required,nullable"`
	// A ACH Check Conversion object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_check_conversion`.
	ACHCheckConversion InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversion `json:"ach_check_conversion,required,nullable"`
	// A ACH Transfer Intention object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention `json:"ach_transfer_intention,required,nullable"`
	// A ACH Transfer Rejection object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection `json:"ach_transfer_rejection,required,nullable"`
	// A ACH Transfer Return object. This field will be present in the JSON response if
	// and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn `json:"ach_transfer_return,required,nullable"`
	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptance `json:"card_dispute_acceptance,required,nullable"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund `json:"card_refund,required,nullable"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement `json:"card_settlement,required,nullable"`
	// A Card Revenue Payment object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_revenue_payment`.
	CardRevenuePayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePayment `json:"card_revenue_payment,required,nullable"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptance `json:"check_deposit_acceptance,required,nullable"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn `json:"check_deposit_return,required,nullable"`
	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention `json:"check_transfer_intention,required,nullable"`
	// A Check Transfer Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_transfer_return`.
	CheckTransferReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturn `json:"check_transfer_return,required,nullable"`
	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejection `json:"check_transfer_rejection,required,nullable"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest `json:"check_transfer_stop_payment_request,required,nullable"`
	// A Dispute Resolution object. This field will be present in the JSON response if
	// and only if `category` is equal to `dispute_resolution`.
	DisputeResolution InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolution `json:"dispute_resolution,required,nullable"`
	// A Empyreal Cash Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `empyreal_cash_deposit`.
	EmpyrealCashDeposit InboundRealTimePaymentsTransferSimulationResultTransactionSourceEmpyrealCashDeposit `json:"empyreal_cash_deposit,required,nullable"`
	// A Fee Payment object. This field will be present in the JSON response if and
	// only if `category` is equal to `fee_payment`.
	FeePayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePayment `json:"fee_payment,required,nullable"`
	// A Inbound ACH Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer `json:"inbound_ach_transfer,required,nullable"`
	// A Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck `json:"inbound_check,required,nullable"`
	// A Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer `json:"inbound_international_ach_transfer,required,nullable"`
	// A Inbound Real Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation `json:"inbound_real_time_payments_transfer_confirmation,required,nullable"`
	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal `json:"inbound_wire_drawdown_payment_reversal,required,nullable"`
	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment `json:"inbound_wire_drawdown_payment,required,nullable"`
	// A Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal `json:"inbound_wire_reversal,required,nullable"`
	// A Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer `json:"inbound_wire_transfer,required,nullable"`
	// A Interest Payment object. This field will be present in the JSON response if
	// and only if `category` is equal to `interest_payment`.
	InterestPayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPayment `json:"interest_payment,required,nullable"`
	// A Internal Source object. This field will be present in the JSON response if and
	// only if `category` is equal to `internal_source`.
	InternalSource InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource `json:"internal_source,required,nullable"`
	// A Deprecated Card Refund object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_route_refund`.
	CardRouteRefund InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefund `json:"card_route_refund,required,nullable"`
	// A Deprecated Card Settlement object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_settlement`.
	CardRouteSettlement InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlement `json:"card_route_settlement,required,nullable"`
	// A Real Time Payments Transfer Acknowledgement object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `real_time_payments_transfer_acknowledgement`.
	RealTimePaymentsTransferAcknowledgement InboundRealTimePaymentsTransferSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgement `json:"real_time_payments_transfer_acknowledgement,required,nullable"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds InboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFunds `json:"sample_funds,required,nullable"`
	// A Wire Drawdown Payment Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_intention`.
	WireDrawdownPaymentIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentIntention `json:"wire_drawdown_payment_intention,required,nullable"`
	// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_rejection`.
	WireDrawdownPaymentRejection InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentRejection `json:"wire_drawdown_payment_rejection,required,nullable"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention `json:"wire_transfer_intention,required,nullable"`
	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejection `json:"wire_transfer_rejection,required,nullable"`
	JSON                  InboundRealTimePaymentsTransferSimulationResultTransactionSourceJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceJSON struct {
	Category                                    apijson.Metadata
	AccountTransferIntention                    apijson.Metadata
	ACHCheckConversionReturn                    apijson.Metadata
	ACHCheckConversion                          apijson.Metadata
	ACHTransferIntention                        apijson.Metadata
	ACHTransferRejection                        apijson.Metadata
	ACHTransferReturn                           apijson.Metadata
	CardDisputeAcceptance                       apijson.Metadata
	CardRefund                                  apijson.Metadata
	CardSettlement                              apijson.Metadata
	CardRevenuePayment                          apijson.Metadata
	CheckDepositAcceptance                      apijson.Metadata
	CheckDepositReturn                          apijson.Metadata
	CheckTransferIntention                      apijson.Metadata
	CheckTransferReturn                         apijson.Metadata
	CheckTransferRejection                      apijson.Metadata
	CheckTransferStopPaymentRequest             apijson.Metadata
	DisputeResolution                           apijson.Metadata
	EmpyrealCashDeposit                         apijson.Metadata
	FeePayment                                  apijson.Metadata
	InboundACHTransfer                          apijson.Metadata
	InboundCheck                                apijson.Metadata
	InboundInternationalACHTransfer             apijson.Metadata
	InboundRealTimePaymentsTransferConfirmation apijson.Metadata
	InboundWireDrawdownPaymentReversal          apijson.Metadata
	InboundWireDrawdownPayment                  apijson.Metadata
	InboundWireReversal                         apijson.Metadata
	InboundWireTransfer                         apijson.Metadata
	InterestPayment                             apijson.Metadata
	InternalSource                              apijson.Metadata
	CardRouteRefund                             apijson.Metadata
	CardRouteSettlement                         apijson.Metadata
	RealTimePaymentsTransferAcknowledgement     apijson.Metadata
	SampleFunds                                 apijson.Metadata
	WireDrawdownPaymentIntention                apijson.Metadata
	WireDrawdownPaymentRejection                apijson.Metadata
	WireTransferIntention                       apijson.Metadata
	WireTransferRejection                       apijson.Metadata
	raw                                         string
	Extras                                      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSource using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCardRevenuePayment                          InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "card_revenue_payment"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckDepositAcceptance                      InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_deposit_acceptance"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckDepositReturn                          InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_deposit_return"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckTransferIntention                      InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_transfer_intention"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckTransferReturn                         InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_transfer_return"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckTransferRejection                      InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_transfer_rejection"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckTransferStopPaymentRequest             InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_transfer_stop_payment_request"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryDisputeResolution                           InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "dispute_resolution"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryEmpyrealCashDeposit                         InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "empyreal_cash_deposit"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryFeePayment                                  InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "fee_payment"
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
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency `json:"currency,required"`
	// The description you chose to give the transfer.
	Description string `json:"description,required"`
	// The identifier of the Account to where the Account Transfer was sent.
	DestinationAccountID string `json:"destination_account_id,required"`
	// The identifier of the Account from where the Account Transfer was sent.
	SourceAccountID string `json:"source_account_id,required"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionJSON struct {
	Amount               apijson.Metadata
	Currency             apijson.Metadata
	Description          apijson.Metadata
	DestinationAccountID apijson.Metadata
	SourceAccountID      apijson.Metadata
	TransferID           apijson.Metadata
	raw                  string
	Extras               map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// Why the transfer was returned.
	ReturnReasonCode string `json:"return_reason_code,required"`
	JSON             InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversionReturnJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversionReturnJSON struct {
	Amount           apijson.Metadata
	ReturnReasonCode apijson.Metadata
	raw              string
	Extras           map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversionReturn
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversionReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversion struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the File containing an image of the returned check.
	FileID string `json:"file_id,required"`
	JSON   InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversionJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversionJSON struct {
	Amount apijson.Metadata
	FileID apijson.Metadata
	raw    string
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversion
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount              int64  `json:"amount,required"`
	AccountNumber       string `json:"account_number,required"`
	RoutingNumber       string `json:"routing_number,required"`
	StatementDescriptor string `json:"statement_descriptor,required"`
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntentionJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntentionJSON struct {
	Amount              apijson.Metadata
	AccountNumber       apijson.Metadata
	RoutingNumber       apijson.Metadata
	StatementDescriptor apijson.Metadata
	TransferID          apijson.Metadata
	raw                 string
	Extras              map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejectionJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejectionJSON struct {
	TransferID apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode `json:"return_reason_code,required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID string `json:"transfer_id,required"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID string `json:"transaction_id,required"`
	JSON          InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnJSON struct {
	CreatedAt        apijson.Metadata
	ReturnReasonCode apijson.Metadata
	TransferID       apijson.Metadata
	TransactionID    apijson.Metadata
	raw              string
	Extras           map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	AcceptedAt time.Time `json:"accepted_at,required" format:"date-time"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID string `json:"card_dispute_id,required"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID string `json:"transaction_id,required"`
	JSON          InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptanceJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptanceJSON struct {
	AcceptedAt    apijson.Metadata
	CardDisputeID apijson.Metadata
	TransactionID apijson.Metadata
	raw           string
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptance
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund struct {
	// The Card Refund identifier.
	ID string `json:"id,required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency `json:"currency,required"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required,nullable"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required,nullable"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required"`
	// The name of the merchant.
	MerchantName string `json:"merchant_name,required,nullable"`
	// The 4-digit MCC describing the merchant's business.
	MerchantCategoryCode string `json:"merchant_category_code,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundType `json:"type,required"`
	JSON InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundJSON struct {
	ID                   apijson.Metadata
	Amount               apijson.Metadata
	Currency             apijson.Metadata
	MerchantAcceptorID   apijson.Metadata
	MerchantCity         apijson.Metadata
	MerchantState        apijson.Metadata
	MerchantCountry      apijson.Metadata
	MerchantName         apijson.Metadata
	MerchantCategoryCode apijson.Metadata
	Type                 apijson.Metadata
	raw                  string
	Extras               map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	// The Card Settlement identifier.
	ID string `json:"id,required"`
	// The Card Authorization that was created prior to this Card Settlement, if on
	// exists.
	CardAuthorization string `json:"card_authorization,required,nullable"`
	// The amount in the minor unit of the transaction's settlement currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's settlement currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency `json:"currency,required"`
	// The amount in the minor unit of the transaction's presentment currency.
	PresentmentAmount int64 `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's presentment currency.
	PresentmentCurrency string `json:"presentment_currency,required"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required,nullable"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required,nullable"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required"`
	// The name of the merchant.
	MerchantName string `json:"merchant_name,required,nullable"`
	// The 4-digit MCC describing the merchant's business.
	MerchantCategoryCode string `json:"merchant_category_code,required"`
	// The identifier of the Pending Transaction associated with this Transaction.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_settlement`.
	Type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementType `json:"type,required"`
	JSON InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementJSON struct {
	ID                   apijson.Metadata
	CardAuthorization    apijson.Metadata
	Amount               apijson.Metadata
	Currency             apijson.Metadata
	PresentmentAmount    apijson.Metadata
	PresentmentCurrency  apijson.Metadata
	MerchantAcceptorID   apijson.Metadata
	MerchantCity         apijson.Metadata
	MerchantState        apijson.Metadata
	MerchantCountry      apijson.Metadata
	MerchantName         apijson.Metadata
	MerchantCategoryCode apijson.Metadata
	PendingTransactionID apijson.Metadata
	Type                 apijson.Metadata
	raw                  string
	Extras               map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrency `json:"currency,required"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time `json:"period_start,required" format:"date-time"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The account the card belonged to.
	TransactedOnAccountID string `json:"transacted_on_account_id,required,nullable"`
	JSON                  InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentJSON struct {
	Amount                apijson.Metadata
	Currency              apijson.Metadata
	PeriodStart           apijson.Metadata
	PeriodEnd             apijson.Metadata
	TransactedOnAccountID apijson.Metadata
	raw                   string
	Extras                map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePayment
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRevenuePaymentCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptance struct {
	// The amount to be deposited in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency `json:"currency,required"`
	// The account number printed on the check.
	AccountNumber string `json:"account_number,required"`
	// The routing number printed on the check.
	RoutingNumber string `json:"routing_number,required"`
	// An additional line of metadata printed on the check. This typically includes the
	// check number for business checks.
	AuxiliaryOnUs string `json:"auxiliary_on_us,required,nullable"`
	// The check serial number, if present, for consumer checks. For business checks,
	// the serial number is usually in the `auxiliary_on_us` field.
	SerialNumber string `json:"serial_number,required,nullable"`
	// The ID of the Check Deposit that was accepted.
	CheckDepositID string `json:"check_deposit_id,required"`
	JSON           InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceJSON struct {
	Amount         apijson.Metadata
	Currency       apijson.Metadata
	AccountNumber  apijson.Metadata
	RoutingNumber  apijson.Metadata
	AuxiliaryOnUs  apijson.Metadata
	SerialNumber   apijson.Metadata
	CheckDepositID apijson.Metadata
	raw            string
	Extras         map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptance
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency `json:"currency,required"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID string `json:"check_deposit_id,required"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID string                                                                                         `json:"transaction_id,required"`
	ReturnReason  InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason `json:"return_reason,required"`
	JSON          InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnJSON struct {
	Amount         apijson.Metadata
	ReturnedAt     apijson.Metadata
	Currency       apijson.Metadata
	CheckDepositID apijson.Metadata
	TransactionID  apijson.Metadata
	ReturnReason   apijson.Metadata
	raw            string
	Extras         map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	AddressLine1 string `json:"address_line1,required"`
	// The second line of the address of the check's destination.
	AddressLine2 string `json:"address_line2,required,nullable"`
	// The city of the check's destination.
	AddressCity string `json:"address_city,required"`
	// The state of the check's destination.
	AddressState string `json:"address_state,required"`
	// The postal code of the check's destination.
	AddressZip string `json:"address_zip,required"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency `json:"currency,required"`
	// The name that will be printed on the check.
	RecipientName string `json:"recipient_name,required"`
	// The identifier of the Check Transfer with which this is associated.
	TransferID string `json:"transfer_id,required"`
	JSON       InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionJSON struct {
	AddressLine1  apijson.Metadata
	AddressLine2  apijson.Metadata
	AddressCity   apijson.Metadata
	AddressState  apijson.Metadata
	AddressZip    apijson.Metadata
	Amount        apijson.Metadata
	Currency      apijson.Metadata
	RecipientName apijson.Metadata
	TransferID    apijson.Metadata
	raw           string
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	TransferID string `json:"transfer_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// If available, a document with additional information about the return.
	FileID string `json:"file_id,required,nullable"`
	// The reason why the check was returned.
	Reason InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnReason `json:"reason,required"`
	// The identifier of the Transaction that was created to credit you for the
	// returned check.
	TransactionID string `json:"transaction_id,required,nullable"`
	JSON          InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnJSON struct {
	TransferID    apijson.Metadata
	ReturnedAt    apijson.Metadata
	FileID        apijson.Metadata
	Reason        apijson.Metadata
	TransactionID apijson.Metadata
	raw           string
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturn
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnReason string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnReasonMailDeliveryFailure InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnReason = "mail_delivery_failure"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnReasonRefusedByRecipient  InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferReturnReason = "refused_by_recipient"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejection struct {
	// The identifier of the Check Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejectionJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejectionJSON struct {
	TransferID apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejection
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferID string `json:"transfer_id,required"`
	// The transaction ID of the corresponding credit transaction.
	TransactionID string `json:"transaction_id,required"`
	// The time the stop-payment was requested.
	RequestedAt time.Time `json:"requested_at,required" format:"date-time"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestType `json:"type,required"`
	JSON InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestJSON struct {
	TransferID    apijson.Metadata
	TransactionID apijson.Metadata
	RequestedAt   apijson.Metadata
	Type          apijson.Metadata
	raw           string
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestType string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolution struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrency `json:"currency,required"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionID string `json:"disputed_transaction_id,required"`
	JSON                  InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionJSON struct {
	Amount                apijson.Metadata
	Currency              apijson.Metadata
	DisputedTransactionID apijson.Metadata
	raw                   string
	Extras                map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolution
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount      int64     `json:"amount,required"`
	BagID       string    `json:"bag_id,required"`
	DepositDate time.Time `json:"deposit_date,required" format:"date-time"`
	JSON        InboundRealTimePaymentsTransferSimulationResultTransactionSourceEmpyrealCashDepositJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceEmpyrealCashDepositJSON struct {
	Amount      apijson.Metadata
	BagID       apijson.Metadata
	DepositDate apijson.Metadata
	raw         string
	Extras      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceEmpyrealCashDeposit
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceEmpyrealCashDeposit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrency `json:"currency,required"`
	JSON     InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentJSON struct {
	Amount   apijson.Metadata
	Currency apijson.Metadata
	raw      string
	Extras   map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePayment using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceFeePaymentCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount                             int64  `json:"amount,required"`
	OriginatorCompanyName              string `json:"originator_company_name,required"`
	OriginatorCompanyDescriptiveDate   string `json:"originator_company_descriptive_date,required,nullable"`
	OriginatorCompanyDiscretionaryData string `json:"originator_company_discretionary_data,required,nullable"`
	OriginatorCompanyEntryDescription  string `json:"originator_company_entry_description,required"`
	OriginatorCompanyID                string `json:"originator_company_id,required"`
	ReceiverIDNumber                   string `json:"receiver_id_number,required,nullable"`
	ReceiverName                       string `json:"receiver_name,required,nullable"`
	TraceNumber                        string `json:"trace_number,required"`
	JSON                               InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransferJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransferJSON struct {
	Amount                             apijson.Metadata
	OriginatorCompanyName              apijson.Metadata
	OriginatorCompanyDescriptiveDate   apijson.Metadata
	OriginatorCompanyDiscretionaryData apijson.Metadata
	OriginatorCompanyEntryDescription  apijson.Metadata
	OriginatorCompanyID                apijson.Metadata
	ReceiverIDNumber                   apijson.Metadata
	ReceiverName                       apijson.Metadata
	TraceNumber                        apijson.Metadata
	raw                                string
	Extras                             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency              InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency `json:"currency,required"`
	CheckNumber           string                                                                               `json:"check_number,required,nullable"`
	CheckFrontImageFileID string                                                                               `json:"check_front_image_file_id,required,nullable"`
	CheckRearImageFileID  string                                                                               `json:"check_rear_image_file_id,required,nullable"`
	JSON                  InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckJSON struct {
	Amount                apijson.Metadata
	Currency              apijson.Metadata
	CheckNumber           apijson.Metadata
	CheckFrontImageFileID apijson.Metadata
	CheckRearImageFileID  apijson.Metadata
	raw                   string
	Extras                map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount                                                 int64  `json:"amount,required"`
	ForeignExchangeIndicator                               string `json:"foreign_exchange_indicator,required"`
	ForeignExchangeReferenceIndicator                      string `json:"foreign_exchange_reference_indicator,required"`
	ForeignExchangeReference                               string `json:"foreign_exchange_reference,required,nullable"`
	DestinationCountryCode                                 string `json:"destination_country_code,required"`
	DestinationCurrencyCode                                string `json:"destination_currency_code,required"`
	ForeignPaymentAmount                                   int64  `json:"foreign_payment_amount,required"`
	ForeignTraceNumber                                     string `json:"foreign_trace_number,required,nullable"`
	InternationalTransactionTypeCode                       string `json:"international_transaction_type_code,required"`
	OriginatingCurrencyCode                                string `json:"originating_currency_code,required"`
	OriginatingDepositoryFinancialInstitutionName          string `json:"originating_depository_financial_institution_name,required"`
	OriginatingDepositoryFinancialInstitutionIDQualifier   string `json:"originating_depository_financial_institution_id_qualifier,required"`
	OriginatingDepositoryFinancialInstitutionID            string `json:"originating_depository_financial_institution_id,required"`
	OriginatingDepositoryFinancialInstitutionBranchCountry string `json:"originating_depository_financial_institution_branch_country,required"`
	OriginatorCity                                         string `json:"originator_city,required"`
	OriginatorCompanyEntryDescription                      string `json:"originator_company_entry_description,required"`
	OriginatorCountry                                      string `json:"originator_country,required"`
	OriginatorIdentification                               string `json:"originator_identification,required"`
	OriginatorName                                         string `json:"originator_name,required"`
	OriginatorPostalCode                                   string `json:"originator_postal_code,required,nullable"`
	OriginatorStreetAddress                                string `json:"originator_street_address,required"`
	OriginatorStateOrProvince                              string `json:"originator_state_or_province,required,nullable"`
	PaymentRelatedInformation                              string `json:"payment_related_information,required,nullable"`
	PaymentRelatedInformation2                             string `json:"payment_related_information2,required,nullable"`
	ReceiverIdentificationNumber                           string `json:"receiver_identification_number,required,nullable"`
	ReceiverStreetAddress                                  string `json:"receiver_street_address,required"`
	ReceiverCity                                           string `json:"receiver_city,required"`
	ReceiverStateOrProvince                                string `json:"receiver_state_or_province,required,nullable"`
	ReceiverCountry                                        string `json:"receiver_country,required"`
	ReceiverPostalCode                                     string `json:"receiver_postal_code,required,nullable"`
	ReceivingCompanyOrIndividualName                       string `json:"receiving_company_or_individual_name,required"`
	ReceivingDepositoryFinancialInstitutionName            string `json:"receiving_depository_financial_institution_name,required"`
	ReceivingDepositoryFinancialInstitutionIDQualifier     string `json:"receiving_depository_financial_institution_id_qualifier,required"`
	ReceivingDepositoryFinancialInstitutionID              string `json:"receiving_depository_financial_institution_id,required"`
	ReceivingDepositoryFinancialInstitutionCountry         string `json:"receiving_depository_financial_institution_country,required"`
	TraceNumber                                            string `json:"trace_number,required"`
	JSON                                                   InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransferJSON struct {
	Amount                                                 apijson.Metadata
	ForeignExchangeIndicator                               apijson.Metadata
	ForeignExchangeReferenceIndicator                      apijson.Metadata
	ForeignExchangeReference                               apijson.Metadata
	DestinationCountryCode                                 apijson.Metadata
	DestinationCurrencyCode                                apijson.Metadata
	ForeignPaymentAmount                                   apijson.Metadata
	ForeignTraceNumber                                     apijson.Metadata
	InternationalTransactionTypeCode                       apijson.Metadata
	OriginatingCurrencyCode                                apijson.Metadata
	OriginatingDepositoryFinancialInstitutionName          apijson.Metadata
	OriginatingDepositoryFinancialInstitutionIDQualifier   apijson.Metadata
	OriginatingDepositoryFinancialInstitutionID            apijson.Metadata
	OriginatingDepositoryFinancialInstitutionBranchCountry apijson.Metadata
	OriginatorCity                                         apijson.Metadata
	OriginatorCompanyEntryDescription                      apijson.Metadata
	OriginatorCountry                                      apijson.Metadata
	OriginatorIdentification                               apijson.Metadata
	OriginatorName                                         apijson.Metadata
	OriginatorPostalCode                                   apijson.Metadata
	OriginatorStreetAddress                                apijson.Metadata
	OriginatorStateOrProvince                              apijson.Metadata
	PaymentRelatedInformation                              apijson.Metadata
	PaymentRelatedInformation2                             apijson.Metadata
	ReceiverIdentificationNumber                           apijson.Metadata
	ReceiverStreetAddress                                  apijson.Metadata
	ReceiverCity                                           apijson.Metadata
	ReceiverStateOrProvince                                apijson.Metadata
	ReceiverCountry                                        apijson.Metadata
	ReceiverPostalCode                                     apijson.Metadata
	ReceivingCompanyOrIndividualName                       apijson.Metadata
	ReceivingDepositoryFinancialInstitutionName            apijson.Metadata
	ReceivingDepositoryFinancialInstitutionIDQualifier     apijson.Metadata
	ReceivingDepositoryFinancialInstitutionID              apijson.Metadata
	ReceivingDepositoryFinancialInstitutionCountry         apijson.Metadata
	TraceNumber                                            apijson.Metadata
	raw                                                    string
	Extras                                                 map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real Time Payments transfer.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency `json:"currency,required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The name provided by the sender of the transfer.
	DebtorName string `json:"debtor_name,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber string `json:"debtor_account_number,required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber string `json:"debtor_routing_number,required"`
	// The Real Time Payments network identification of the transfer
	TransactionIdentification string `json:"transaction_identification,required"`
	// Additional information included with the transfer.
	RemittanceInformation string `json:"remittance_information,required,nullable"`
	JSON                  InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON struct {
	Amount                    apijson.Metadata
	Currency                  apijson.Metadata
	CreditorName              apijson.Metadata
	DebtorName                apijson.Metadata
	DebtorAccountNumber       apijson.Metadata
	DebtorRoutingNumber       apijson.Metadata
	TransactionIdentification apijson.Metadata
	RemittanceInformation     apijson.Metadata
	raw                       string
	Extras                    map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The description on the reversal message from Fedwire.
	Description string `json:"description,required"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate time.Time `json:"input_cycle_date,required" format:"date"`
	// The Fedwire sequence number.
	InputSequenceNumber string `json:"input_sequence_number,required"`
	// The Fedwire input source identifier.
	InputSource string `json:"input_source,required"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData string `json:"previous_message_input_message_accountability_data,required"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate time.Time `json:"previous_message_input_cycle_date,required" format:"date"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber string `json:"previous_message_input_sequence_number,required"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource string `json:"previous_message_input_source,required"`
	JSON                       InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversalJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversalJSON struct {
	Amount                                        apijson.Metadata
	Description                                   apijson.Metadata
	InputCycleDate                                apijson.Metadata
	InputSequenceNumber                           apijson.Metadata
	InputSource                                   apijson.Metadata
	InputMessageAccountabilityData                apijson.Metadata
	PreviousMessageInputMessageAccountabilityData apijson.Metadata
	PreviousMessageInputCycleDate                 apijson.Metadata
	PreviousMessageInputSequenceNumber            apijson.Metadata
	PreviousMessageInputSource                    apijson.Metadata
	raw                                           string
	Extras                                        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount                             int64  `json:"amount,required"`
	BeneficiaryAddressLine1            string `json:"beneficiary_address_line1,required,nullable"`
	BeneficiaryAddressLine2            string `json:"beneficiary_address_line2,required,nullable"`
	BeneficiaryAddressLine3            string `json:"beneficiary_address_line3,required,nullable"`
	BeneficiaryName                    string `json:"beneficiary_name,required,nullable"`
	BeneficiaryReference               string `json:"beneficiary_reference,required,nullable"`
	Description                        string `json:"description,required"`
	InputMessageAccountabilityData     string `json:"input_message_accountability_data,required,nullable"`
	OriginatorAddressLine1             string `json:"originator_address_line1,required,nullable"`
	OriginatorAddressLine2             string `json:"originator_address_line2,required,nullable"`
	OriginatorAddressLine3             string `json:"originator_address_line3,required,nullable"`
	OriginatorName                     string `json:"originator_name,required,nullable"`
	OriginatorToBeneficiaryInformation string `json:"originator_to_beneficiary_information,required,nullable"`
	JSON                               InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentJSON struct {
	Amount                             apijson.Metadata
	BeneficiaryAddressLine1            apijson.Metadata
	BeneficiaryAddressLine2            apijson.Metadata
	BeneficiaryAddressLine3            apijson.Metadata
	BeneficiaryName                    apijson.Metadata
	BeneficiaryReference               apijson.Metadata
	Description                        apijson.Metadata
	InputMessageAccountabilityData     apijson.Metadata
	OriginatorAddressLine1             apijson.Metadata
	OriginatorAddressLine2             apijson.Metadata
	OriginatorAddressLine3             apijson.Metadata
	OriginatorName                     apijson.Metadata
	OriginatorToBeneficiaryInformation apijson.Metadata
	raw                                string
	Extras                             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal struct {
	// The amount that was reversed.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the reversal was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description on the reversal message from Fedwire.
	Description string `json:"description,required"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate time.Time `json:"input_cycle_date,required" format:"date"`
	// The Fedwire sequence number.
	InputSequenceNumber string `json:"input_sequence_number,required"`
	// The Fedwire input source identifier.
	InputSource string `json:"input_source,required"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData string `json:"previous_message_input_message_accountability_data,required"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate time.Time `json:"previous_message_input_cycle_date,required" format:"date"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber string `json:"previous_message_input_sequence_number,required"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource string `json:"previous_message_input_source,required"`
	// Information included in the wire reversal for the receiving financial
	// institution.
	ReceiverFinancialInstitutionInformation string `json:"receiver_financial_institution_information,required,nullable"`
	// Additional financial institution information included in the wire reversal.
	FinancialInstitutionToFinancialInstitutionInformation string `json:"financial_institution_to_financial_institution_information,required,nullable"`
	// The ID for the Transaction associated with the transfer reversal.
	TransactionID string `json:"transaction_id,required,nullable"`
	JSON          InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversalJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversalJSON struct {
	Amount                                                apijson.Metadata
	CreatedAt                                             apijson.Metadata
	Description                                           apijson.Metadata
	InputCycleDate                                        apijson.Metadata
	InputSequenceNumber                                   apijson.Metadata
	InputSource                                           apijson.Metadata
	InputMessageAccountabilityData                        apijson.Metadata
	PreviousMessageInputMessageAccountabilityData         apijson.Metadata
	PreviousMessageInputCycleDate                         apijson.Metadata
	PreviousMessageInputSequenceNumber                    apijson.Metadata
	PreviousMessageInputSource                            apijson.Metadata
	ReceiverFinancialInstitutionInformation               apijson.Metadata
	FinancialInstitutionToFinancialInstitutionInformation apijson.Metadata
	TransactionID                                         apijson.Metadata
	raw                                                   string
	Extras                                                map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount                                  int64  `json:"amount,required"`
	BeneficiaryAddressLine1                 string `json:"beneficiary_address_line1,required,nullable"`
	BeneficiaryAddressLine2                 string `json:"beneficiary_address_line2,required,nullable"`
	BeneficiaryAddressLine3                 string `json:"beneficiary_address_line3,required,nullable"`
	BeneficiaryName                         string `json:"beneficiary_name,required,nullable"`
	BeneficiaryReference                    string `json:"beneficiary_reference,required,nullable"`
	Description                             string `json:"description,required"`
	InputMessageAccountabilityData          string `json:"input_message_accountability_data,required,nullable"`
	OriginatorAddressLine1                  string `json:"originator_address_line1,required,nullable"`
	OriginatorAddressLine2                  string `json:"originator_address_line2,required,nullable"`
	OriginatorAddressLine3                  string `json:"originator_address_line3,required,nullable"`
	OriginatorName                          string `json:"originator_name,required,nullable"`
	OriginatorToBeneficiaryInformationLine1 string `json:"originator_to_beneficiary_information_line1,required,nullable"`
	OriginatorToBeneficiaryInformationLine2 string `json:"originator_to_beneficiary_information_line2,required,nullable"`
	OriginatorToBeneficiaryInformationLine3 string `json:"originator_to_beneficiary_information_line3,required,nullable"`
	OriginatorToBeneficiaryInformationLine4 string `json:"originator_to_beneficiary_information_line4,required,nullable"`
	OriginatorToBeneficiaryInformation      string `json:"originator_to_beneficiary_information,required,nullable"`
	JSON                                    InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransferJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransferJSON struct {
	Amount                                  apijson.Metadata
	BeneficiaryAddressLine1                 apijson.Metadata
	BeneficiaryAddressLine2                 apijson.Metadata
	BeneficiaryAddressLine3                 apijson.Metadata
	BeneficiaryName                         apijson.Metadata
	BeneficiaryReference                    apijson.Metadata
	Description                             apijson.Metadata
	InputMessageAccountabilityData          apijson.Metadata
	OriginatorAddressLine1                  apijson.Metadata
	OriginatorAddressLine2                  apijson.Metadata
	OriginatorAddressLine3                  apijson.Metadata
	OriginatorName                          apijson.Metadata
	OriginatorToBeneficiaryInformationLine1 apijson.Metadata
	OriginatorToBeneficiaryInformationLine2 apijson.Metadata
	OriginatorToBeneficiaryInformationLine3 apijson.Metadata
	OriginatorToBeneficiaryInformationLine4 apijson.Metadata
	OriginatorToBeneficiaryInformation      apijson.Metadata
	raw                                     string
	Extras                                  map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentCurrency `json:"currency,required"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time `json:"period_start,required" format:"date-time"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The account on which the interest was accrued.
	AccruedOnAccountID string `json:"accrued_on_account_id,required,nullable"`
	JSON               InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPaymentJSON struct {
	Amount             apijson.Metadata
	Currency           apijson.Metadata
	PeriodStart        apijson.Metadata
	PeriodEnd          apijson.Metadata
	AccruedOnAccountID apijson.Metadata
	raw                string
	Extras             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPayment
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInterestPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency `json:"currency,required"`
	Reason   InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason   `json:"reason,required"`
	JSON     InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceJSON struct {
	Amount   apijson.Metadata
	Currency apijson.Metadata
	Reason   apijson.Metadata
	raw      string
	Extras   map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
	// currency.
	Currency             InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrency `json:"currency,required"`
	MerchantAcceptorID   string                                                                                  `json:"merchant_acceptor_id,required"`
	MerchantCity         string                                                                                  `json:"merchant_city,required,nullable"`
	MerchantCountry      string                                                                                  `json:"merchant_country,required"`
	MerchantDescriptor   string                                                                                  `json:"merchant_descriptor,required"`
	MerchantState        string                                                                                  `json:"merchant_state,required,nullable"`
	MerchantCategoryCode string                                                                                  `json:"merchant_category_code,required,nullable"`
	JSON                 InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundJSON struct {
	Amount               apijson.Metadata
	Currency             apijson.Metadata
	MerchantAcceptorID   apijson.Metadata
	MerchantCity         apijson.Metadata
	MerchantCountry      apijson.Metadata
	MerchantDescriptor   apijson.Metadata
	MerchantState        apijson.Metadata
	MerchantCategoryCode apijson.Metadata
	raw                  string
	Extras               map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefund
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefund) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
	// currency.
	Currency             InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrency `json:"currency,required"`
	MerchantAcceptorID   string                                                                                      `json:"merchant_acceptor_id,required"`
	MerchantCity         string                                                                                      `json:"merchant_city,required,nullable"`
	MerchantCountry      string                                                                                      `json:"merchant_country,required,nullable"`
	MerchantDescriptor   string                                                                                      `json:"merchant_descriptor,required"`
	MerchantState        string                                                                                      `json:"merchant_state,required,nullable"`
	MerchantCategoryCode string                                                                                      `json:"merchant_category_code,required,nullable"`
	JSON                 InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementJSON struct {
	Amount               apijson.Metadata
	Currency             apijson.Metadata
	MerchantAcceptorID   apijson.Metadata
	MerchantCity         apijson.Metadata
	MerchantCountry      apijson.Metadata
	MerchantDescriptor   apijson.Metadata
	MerchantState        apijson.Metadata
	MerchantCategoryCode apijson.Metadata
	raw                  string
	Extras               map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlement
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgement struct {
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The destination account number.
	DestinationAccountNumber string `json:"destination_account_number,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	DestinationRoutingNumber string `json:"destination_routing_number,required"`
	// Unstructured information that will show on the recipient's bank statement.
	RemittanceInformation string `json:"remittance_information,required"`
	// The identifier of the Real Time Payments Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       InboundRealTimePaymentsTransferSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgementJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgementJSON struct {
	Amount                   apijson.Metadata
	DestinationAccountNumber apijson.Metadata
	DestinationRoutingNumber apijson.Metadata
	RemittanceInformation    apijson.Metadata
	TransferID               apijson.Metadata
	raw                      string
	Extras                   map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgement
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFunds struct {
	// Where the sample funds came from.
	Originator string `json:"originator,required"`
	JSON       InboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFundsJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFundsJSON struct {
	Originator apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFunds
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFunds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentIntention struct {
	// The transfer amount in USD cents.
	Amount             int64  `json:"amount,required"`
	AccountNumber      string `json:"account_number,required"`
	RoutingNumber      string `json:"routing_number,required"`
	MessageToRecipient string `json:"message_to_recipient,required"`
	TransferID         string `json:"transfer_id,required"`
	JSON               InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentIntentionJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentIntentionJSON struct {
	Amount             apijson.Metadata
	AccountNumber      apijson.Metadata
	RoutingNumber      apijson.Metadata
	MessageToRecipient apijson.Metadata
	TransferID         apijson.Metadata
	raw                string
	Extras             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentIntention
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentRejection struct {
	TransferID string `json:"transfer_id,required"`
	JSON       InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentRejectionJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentRejectionJSON struct {
	TransferID apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentRejection
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention struct {
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient string `json:"message_to_recipient,required"`
	TransferID         string `json:"transfer_id,required"`
	JSON               InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntentionJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntentionJSON struct {
	Amount             apijson.Metadata
	AccountNumber      apijson.Metadata
	RoutingNumber      apijson.Metadata
	MessageToRecipient apijson.Metadata
	TransferID         apijson.Metadata
	raw                string
	Extras             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejection struct {
	TransferID string `json:"transfer_id,required"`
	JSON       InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejectionJSON
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejectionJSON struct {
	TransferID apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejection
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultTransactionType string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionTypeTransaction InboundRealTimePaymentsTransferSimulationResultTransactionType = "transaction"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction struct {
	// The identifier for the Account the Declined Transaction belongs to.
	AccountID string `json:"account_id,required"`
	// The Declined Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transcation's Account.
	Currency InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency `json:"currency,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// This is the description the vendor provides.
	Description string `json:"description,required"`
	// The Declined Transaction identifier.
	ID string `json:"id,required"`
	// The identifier for the route this Declined Transaction came through. Routes are
	// things like cards and ACH details.
	RouteID string `json:"route_id,required,nullable"`
	// The type of the route this Declined Transaction came through.
	RouteType InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionRouteType `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Declined Transaction. For example, for a card transaction this lists the
	// merchant's industry and location. Note that for backwards compatibility reasons,
	// additional undocumented keys may appear in this object. These should be treated
	// as deprecated and will be removed in the future.
	Source InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionType `json:"type,required"`
	JSON InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionJSON
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionJSON struct {
	AccountID   apijson.Metadata
	Amount      apijson.Metadata
	Currency    apijson.Metadata
	CreatedAt   apijson.Metadata
	Description apijson.Metadata
	ID          apijson.Metadata
	RouteID     apijson.Metadata
	RouteType   apijson.Metadata
	Source      apijson.Metadata
	Type        apijson.Metadata
	raw         string
	Extras      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Category InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory `json:"category,required"`
	// A ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline `json:"ach_decline,required,nullable"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline `json:"card_decline,required,nullable"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline `json:"check_decline,required,nullable"`
	// A Inbound Real Time Payments Transfer Decline object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline `json:"inbound_real_time_payments_transfer_decline,required,nullable"`
	// A International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline `json:"international_ach_decline,required,nullable"`
	// A Deprecated Card Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_decline`.
	CardRouteDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline `json:"card_route_decline,required,nullable"`
	JSON             InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceJSON
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceJSON struct {
	Category                               apijson.Metadata
	ACHDecline                             apijson.Metadata
	CardDecline                            apijson.Metadata
	CheckDecline                           apijson.Metadata
	InboundRealTimePaymentsTransferDecline apijson.Metadata
	InternationalACHDecline                apijson.Metadata
	CardRouteDecline                       apijson.Metadata
	raw                                    string
	Extras                                 map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount                             int64  `json:"amount,required"`
	OriginatorCompanyName              string `json:"originator_company_name,required"`
	OriginatorCompanyDescriptiveDate   string `json:"originator_company_descriptive_date,required,nullable"`
	OriginatorCompanyDiscretionaryData string `json:"originator_company_discretionary_data,required,nullable"`
	OriginatorCompanyID                string `json:"originator_company_id,required"`
	// Why the ACH transfer was declined.
	Reason           InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason `json:"reason,required"`
	ReceiverIDNumber string                                                                                   `json:"receiver_id_number,required,nullable"`
	ReceiverName     string                                                                                   `json:"receiver_name,required,nullable"`
	TraceNumber      string                                                                                   `json:"trace_number,required"`
	JSON             InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineJSON
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineJSON struct {
	Amount                             apijson.Metadata
	OriginatorCompanyName              apijson.Metadata
	OriginatorCompanyDescriptiveDate   apijson.Metadata
	OriginatorCompanyDiscretionaryData apijson.Metadata
	OriginatorCompanyID                apijson.Metadata
	Reason                             apijson.Metadata
	ReceiverIDNumber                   apijson.Metadata
	ReceiverName                       apijson.Metadata
	TraceNumber                        apijson.Metadata
	raw                                string
	Extras                             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonACHRouteCanceled             InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "ach_route_canceled"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonACHRouteDisabled             InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "ach_route_disabled"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonBreachesLimit                InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonDuplicateReturn              InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "duplicate_return"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonEntityNotActive              InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "entity_not_active"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonGroupLocked                  InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "group_locked"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonInsufficientFunds            InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "insufficient_funds"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonMisroutedReturn              InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "misrouted_return"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonNoACHRoute                   InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "no_ach_route"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonOriginatorRequest            InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "originator_request"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonTransactionNotAllowed        InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "transaction_not_allowed"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline struct {
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor string `json:"merchant_descriptor,required"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode string `json:"merchant_category_code,required,nullable"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required,nullable"`
	// The payment network used to process this card authorization
	Network InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetwork `json:"network,required"`
	// Fields specific to the `network`
	NetworkDetails InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetails `json:"network_details,required"`
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency `json:"currency,required"`
	// Why the transaction was declined.
	Reason InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason `json:"reason,required"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state,required,nullable"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// If the authorization was attempted using a Digital Wallet Token (such as an
	// Apple Pay purchase), the identifier of the token that was used.
	DigitalWalletTokenID string `json:"digital_wallet_token_id,required,nullable"`
	JSON                 InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineJSON
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineJSON struct {
	MerchantAcceptorID   apijson.Metadata
	MerchantDescriptor   apijson.Metadata
	MerchantCategoryCode apijson.Metadata
	MerchantCity         apijson.Metadata
	MerchantCountry      apijson.Metadata
	Network              apijson.Metadata
	NetworkDetails       apijson.Metadata
	Amount               apijson.Metadata
	Currency             apijson.Metadata
	Reason               apijson.Metadata
	MerchantState        apijson.Metadata
	RealTimeDecisionID   apijson.Metadata
	DigitalWalletTokenID apijson.Metadata
	raw                  string
	Extras               map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetwork string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkVisa InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetwork = "visa"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisa `json:"visa,required"`
	JSON InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsJSON
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsJSON struct {
	Visa   apijson.Metadata
	raw    string
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetails
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode PointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Metadata
	PointOfServiceEntryMode     apijson.Metadata
	raw                         string
	Extras                      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisa
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                          InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                               InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                             InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                   InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                     InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                    InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
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
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonCardNotActive                InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "card_not_active"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonEntityNotActive              InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "entity_not_active"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonGroupLocked                  InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "group_locked"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonInsufficientFunds            InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "insufficient_funds"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonCvv2Mismatch                 InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "cvv2_mismatch"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonTransactionNotAllowed        InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "transaction_not_allowed"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonBreachesInternalLimit        InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "breaches_internal_limit"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonBreachesLimit                InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "breaches_limit"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonWebhookDeclined              InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "webhook_declined"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonWebhookTimedOut              InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "webhook_timed_out"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonDeclinedByStandInProcessing  InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "declined_by_stand_in_processing"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonInvalidPhysicalCard          InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "invalid_physical_card"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonMissingOriginalAuthorization InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "missing_original_authorization"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount        int64  `json:"amount,required"`
	AuxiliaryOnUs string `json:"auxiliary_on_us,required,nullable"`
	// Why the check was declined.
	Reason InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason `json:"reason,required"`
	JSON   InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineJSON
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineJSON struct {
	Amount        apijson.Metadata
	AuxiliaryOnUs apijson.Metadata
	Reason        apijson.Metadata
	raw           string
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonNotAuthorized         InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "not_authorized"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real Time Payments
	// transfer.
	Currency InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency `json:"currency,required"`
	// Why the transfer was declined.
	Reason InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason `json:"reason,required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The name provided by the sender of the transfer.
	DebtorName string `json:"debtor_name,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber string `json:"debtor_account_number,required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber string `json:"debtor_routing_number,required"`
	// The Real Time Payments network identification of the declined transfer.
	TransactionIdentification string `json:"transaction_identification,required"`
	// Additional information included with the transfer.
	RemittanceInformation string `json:"remittance_information,required,nullable"`
	JSON                  InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON struct {
	Amount                    apijson.Metadata
	Currency                  apijson.Metadata
	Reason                    apijson.Metadata
	CreditorName              apijson.Metadata
	DebtorName                apijson.Metadata
	DebtorAccountNumber       apijson.Metadata
	DebtorRoutingNumber       apijson.Metadata
	TransactionIdentification apijson.Metadata
	RemittanceInformation     apijson.Metadata
	raw                       string
	Extras                    map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount                                                 int64  `json:"amount,required"`
	ForeignExchangeIndicator                               string `json:"foreign_exchange_indicator,required"`
	ForeignExchangeReferenceIndicator                      string `json:"foreign_exchange_reference_indicator,required"`
	ForeignExchangeReference                               string `json:"foreign_exchange_reference,required,nullable"`
	DestinationCountryCode                                 string `json:"destination_country_code,required"`
	DestinationCurrencyCode                                string `json:"destination_currency_code,required"`
	ForeignPaymentAmount                                   int64  `json:"foreign_payment_amount,required"`
	ForeignTraceNumber                                     string `json:"foreign_trace_number,required,nullable"`
	InternationalTransactionTypeCode                       string `json:"international_transaction_type_code,required"`
	OriginatingCurrencyCode                                string `json:"originating_currency_code,required"`
	OriginatingDepositoryFinancialInstitutionName          string `json:"originating_depository_financial_institution_name,required"`
	OriginatingDepositoryFinancialInstitutionIDQualifier   string `json:"originating_depository_financial_institution_id_qualifier,required"`
	OriginatingDepositoryFinancialInstitutionID            string `json:"originating_depository_financial_institution_id,required"`
	OriginatingDepositoryFinancialInstitutionBranchCountry string `json:"originating_depository_financial_institution_branch_country,required"`
	OriginatorCity                                         string `json:"originator_city,required"`
	OriginatorCompanyEntryDescription                      string `json:"originator_company_entry_description,required"`
	OriginatorCountry                                      string `json:"originator_country,required"`
	OriginatorIdentification                               string `json:"originator_identification,required"`
	OriginatorName                                         string `json:"originator_name,required"`
	OriginatorPostalCode                                   string `json:"originator_postal_code,required,nullable"`
	OriginatorStreetAddress                                string `json:"originator_street_address,required"`
	OriginatorStateOrProvince                              string `json:"originator_state_or_province,required,nullable"`
	PaymentRelatedInformation                              string `json:"payment_related_information,required,nullable"`
	PaymentRelatedInformation2                             string `json:"payment_related_information2,required,nullable"`
	ReceiverIdentificationNumber                           string `json:"receiver_identification_number,required,nullable"`
	ReceiverStreetAddress                                  string `json:"receiver_street_address,required"`
	ReceiverCity                                           string `json:"receiver_city,required"`
	ReceiverStateOrProvince                                string `json:"receiver_state_or_province,required,nullable"`
	ReceiverCountry                                        string `json:"receiver_country,required"`
	ReceiverPostalCode                                     string `json:"receiver_postal_code,required,nullable"`
	ReceivingCompanyOrIndividualName                       string `json:"receiving_company_or_individual_name,required"`
	ReceivingDepositoryFinancialInstitutionName            string `json:"receiving_depository_financial_institution_name,required"`
	ReceivingDepositoryFinancialInstitutionIDQualifier     string `json:"receiving_depository_financial_institution_id_qualifier,required"`
	ReceivingDepositoryFinancialInstitutionID              string `json:"receiving_depository_financial_institution_id,required"`
	ReceivingDepositoryFinancialInstitutionCountry         string `json:"receiving_depository_financial_institution_country,required"`
	TraceNumber                                            string `json:"trace_number,required"`
	JSON                                                   InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineJSON
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDeclineJSON struct {
	Amount                                                 apijson.Metadata
	ForeignExchangeIndicator                               apijson.Metadata
	ForeignExchangeReferenceIndicator                      apijson.Metadata
	ForeignExchangeReference                               apijson.Metadata
	DestinationCountryCode                                 apijson.Metadata
	DestinationCurrencyCode                                apijson.Metadata
	ForeignPaymentAmount                                   apijson.Metadata
	ForeignTraceNumber                                     apijson.Metadata
	InternationalTransactionTypeCode                       apijson.Metadata
	OriginatingCurrencyCode                                apijson.Metadata
	OriginatingDepositoryFinancialInstitutionName          apijson.Metadata
	OriginatingDepositoryFinancialInstitutionIDQualifier   apijson.Metadata
	OriginatingDepositoryFinancialInstitutionID            apijson.Metadata
	OriginatingDepositoryFinancialInstitutionBranchCountry apijson.Metadata
	OriginatorCity                                         apijson.Metadata
	OriginatorCompanyEntryDescription                      apijson.Metadata
	OriginatorCountry                                      apijson.Metadata
	OriginatorIdentification                               apijson.Metadata
	OriginatorName                                         apijson.Metadata
	OriginatorPostalCode                                   apijson.Metadata
	OriginatorStreetAddress                                apijson.Metadata
	OriginatorStateOrProvince                              apijson.Metadata
	PaymentRelatedInformation                              apijson.Metadata
	PaymentRelatedInformation2                             apijson.Metadata
	ReceiverIdentificationNumber                           apijson.Metadata
	ReceiverStreetAddress                                  apijson.Metadata
	ReceiverCity                                           apijson.Metadata
	ReceiverStateOrProvince                                apijson.Metadata
	ReceiverCountry                                        apijson.Metadata
	ReceiverPostalCode                                     apijson.Metadata
	ReceivingCompanyOrIndividualName                       apijson.Metadata
	ReceivingDepositoryFinancialInstitutionName            apijson.Metadata
	ReceivingDepositoryFinancialInstitutionIDQualifier     apijson.Metadata
	ReceivingDepositoryFinancialInstitutionID              apijson.Metadata
	ReceivingDepositoryFinancialInstitutionCountry         apijson.Metadata
	TraceNumber                                            apijson.Metadata
	raw                                                    string
	Extras                                                 map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency             InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrency `json:"currency,required"`
	MerchantAcceptorID   string                                                                                           `json:"merchant_acceptor_id,required"`
	MerchantCity         string                                                                                           `json:"merchant_city,required,nullable"`
	MerchantCountry      string                                                                                           `json:"merchant_country,required"`
	MerchantDescriptor   string                                                                                           `json:"merchant_descriptor,required"`
	MerchantState        string                                                                                           `json:"merchant_state,required,nullable"`
	MerchantCategoryCode string                                                                                           `json:"merchant_category_code,required,nullable"`
	JSON                 InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineJSON
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineJSON struct {
	Amount               apijson.Metadata
	Currency             apijson.Metadata
	MerchantAcceptorID   apijson.Metadata
	MerchantCity         apijson.Metadata
	MerchantCountry      apijson.Metadata
	MerchantDescriptor   apijson.Metadata
	MerchantState        apijson.Metadata
	MerchantCategoryCode apijson.Metadata
	raw                  string
	Extras               map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
