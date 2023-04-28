package responses

import (
	"time"

	apijson "github.com/increase/increase-go/internal/json"
)

type WireTransferSimulation struct {
	// If the Wire Transfer attempt succeeds, this will contain the resulting
	// [Transaction](#transactions) object. The Transaction's `source` will be of
	// `category: inbound_wire_transfer`.
	Transaction WireTransferSimulationTransaction `json:"transaction,required"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_wire_transfer_simulation_result`.
	Type WireTransferSimulationType `json:"type,required"`
	JSON WireTransferSimulationJSON
}

type WireTransferSimulationJSON struct {
	Transaction apijson.Metadata
	Type        apijson.Metadata
	raw         string
	Extras      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into WireTransferSimulation using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransaction struct {
	// The identifier for the Account the Transaction belongs to.
	AccountID string `json:"account_id,required"`
	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transcation's
	// Account.
	Currency WireTransferSimulationTransactionCurrency `json:"currency,required"`
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
	RouteType WireTransferSimulationTransactionRouteType `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Transaction. Note that for backwards compatibility reasons, additional
	// undocumented keys may appear in this object. These should be treated as
	// deprecated and will be removed in the future.
	Source WireTransferSimulationTransactionSource `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type WireTransferSimulationTransactionType `json:"type,required"`
	JSON WireTransferSimulationTransactionJSON
}

type WireTransferSimulationTransactionJSON struct {
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
// WireTransferSimulationTransaction using the internal json library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type WireTransferSimulationTransactionRouteType string

const (
	WireTransferSimulationTransactionRouteTypeAccountNumber WireTransferSimulationTransactionRouteType = "account_number"
	WireTransferSimulationTransactionRouteTypeCard          WireTransferSimulationTransactionRouteType = "card"
)

type WireTransferSimulationTransactionSource struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category WireTransferSimulationTransactionSourceCategory `json:"category,required"`
	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention WireTransferSimulationTransactionSourceAccountTransferIntention `json:"account_transfer_intention,required,nullable"`
	// A ACH Check Conversion Return object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_check_conversion_return`.
	ACHCheckConversionReturn WireTransferSimulationTransactionSourceACHCheckConversionReturn `json:"ach_check_conversion_return,required,nullable"`
	// A ACH Check Conversion object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_check_conversion`.
	ACHCheckConversion WireTransferSimulationTransactionSourceACHCheckConversion `json:"ach_check_conversion,required,nullable"`
	// A ACH Transfer Intention object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention WireTransferSimulationTransactionSourceACHTransferIntention `json:"ach_transfer_intention,required,nullable"`
	// A ACH Transfer Rejection object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection WireTransferSimulationTransactionSourceACHTransferRejection `json:"ach_transfer_rejection,required,nullable"`
	// A ACH Transfer Return object. This field will be present in the JSON response if
	// and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn WireTransferSimulationTransactionSourceACHTransferReturn `json:"ach_transfer_return,required,nullable"`
	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance WireTransferSimulationTransactionSourceCardDisputeAcceptance `json:"card_dispute_acceptance,required,nullable"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund WireTransferSimulationTransactionSourceCardRefund `json:"card_refund,required,nullable"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement WireTransferSimulationTransactionSourceCardSettlement `json:"card_settlement,required,nullable"`
	// A Card Revenue Payment object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_revenue_payment`.
	CardRevenuePayment WireTransferSimulationTransactionSourceCardRevenuePayment `json:"card_revenue_payment,required,nullable"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance WireTransferSimulationTransactionSourceCheckDepositAcceptance `json:"check_deposit_acceptance,required,nullable"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn WireTransferSimulationTransactionSourceCheckDepositReturn `json:"check_deposit_return,required,nullable"`
	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention WireTransferSimulationTransactionSourceCheckTransferIntention `json:"check_transfer_intention,required,nullable"`
	// A Check Transfer Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_transfer_return`.
	CheckTransferReturn WireTransferSimulationTransactionSourceCheckTransferReturn `json:"check_transfer_return,required,nullable"`
	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection WireTransferSimulationTransactionSourceCheckTransferRejection `json:"check_transfer_rejection,required,nullable"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest `json:"check_transfer_stop_payment_request,required,nullable"`
	// A Dispute Resolution object. This field will be present in the JSON response if
	// and only if `category` is equal to `dispute_resolution`.
	DisputeResolution WireTransferSimulationTransactionSourceDisputeResolution `json:"dispute_resolution,required,nullable"`
	// A Empyreal Cash Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `empyreal_cash_deposit`.
	EmpyrealCashDeposit WireTransferSimulationTransactionSourceEmpyrealCashDeposit `json:"empyreal_cash_deposit,required,nullable"`
	// A Fee Payment object. This field will be present in the JSON response if and
	// only if `category` is equal to `fee_payment`.
	FeePayment WireTransferSimulationTransactionSourceFeePayment `json:"fee_payment,required,nullable"`
	// A Inbound ACH Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer WireTransferSimulationTransactionSourceInboundACHTransfer `json:"inbound_ach_transfer,required,nullable"`
	// A Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck WireTransferSimulationTransactionSourceInboundCheck `json:"inbound_check,required,nullable"`
	// A Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer WireTransferSimulationTransactionSourceInboundInternationalACHTransfer `json:"inbound_international_ach_transfer,required,nullable"`
	// A Inbound Real Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation `json:"inbound_real_time_payments_transfer_confirmation,required,nullable"`
	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal `json:"inbound_wire_drawdown_payment_reversal,required,nullable"`
	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment WireTransferSimulationTransactionSourceInboundWireDrawdownPayment `json:"inbound_wire_drawdown_payment,required,nullable"`
	// A Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal WireTransferSimulationTransactionSourceInboundWireReversal `json:"inbound_wire_reversal,required,nullable"`
	// A Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer WireTransferSimulationTransactionSourceInboundWireTransfer `json:"inbound_wire_transfer,required,nullable"`
	// A Interest Payment object. This field will be present in the JSON response if
	// and only if `category` is equal to `interest_payment`.
	InterestPayment WireTransferSimulationTransactionSourceInterestPayment `json:"interest_payment,required,nullable"`
	// A Internal Source object. This field will be present in the JSON response if and
	// only if `category` is equal to `internal_source`.
	InternalSource WireTransferSimulationTransactionSourceInternalSource `json:"internal_source,required,nullable"`
	// A Deprecated Card Refund object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_route_refund`.
	CardRouteRefund WireTransferSimulationTransactionSourceCardRouteRefund `json:"card_route_refund,required,nullable"`
	// A Deprecated Card Settlement object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_settlement`.
	CardRouteSettlement WireTransferSimulationTransactionSourceCardRouteSettlement `json:"card_route_settlement,required,nullable"`
	// A Real Time Payments Transfer Acknowledgement object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `real_time_payments_transfer_acknowledgement`.
	RealTimePaymentsTransferAcknowledgement WireTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgement `json:"real_time_payments_transfer_acknowledgement,required,nullable"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds WireTransferSimulationTransactionSourceSampleFunds `json:"sample_funds,required,nullable"`
	// A Wire Drawdown Payment Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_intention`.
	WireDrawdownPaymentIntention WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention `json:"wire_drawdown_payment_intention,required,nullable"`
	// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_rejection`.
	WireDrawdownPaymentRejection WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection `json:"wire_drawdown_payment_rejection,required,nullable"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention WireTransferSimulationTransactionSourceWireTransferIntention `json:"wire_transfer_intention,required,nullable"`
	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection WireTransferSimulationTransactionSourceWireTransferRejection `json:"wire_transfer_rejection,required,nullable"`
	JSON                  WireTransferSimulationTransactionSourceJSON
}

type WireTransferSimulationTransactionSourceJSON struct {
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
// WireTransferSimulationTransactionSource using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	WireTransferSimulationTransactionSourceCategoryCardRevenuePayment                          WireTransferSimulationTransactionSourceCategory = "card_revenue_payment"
	WireTransferSimulationTransactionSourceCategoryCheckDepositAcceptance                      WireTransferSimulationTransactionSourceCategory = "check_deposit_acceptance"
	WireTransferSimulationTransactionSourceCategoryCheckDepositReturn                          WireTransferSimulationTransactionSourceCategory = "check_deposit_return"
	WireTransferSimulationTransactionSourceCategoryCheckTransferIntention                      WireTransferSimulationTransactionSourceCategory = "check_transfer_intention"
	WireTransferSimulationTransactionSourceCategoryCheckTransferReturn                         WireTransferSimulationTransactionSourceCategory = "check_transfer_return"
	WireTransferSimulationTransactionSourceCategoryCheckTransferRejection                      WireTransferSimulationTransactionSourceCategory = "check_transfer_rejection"
	WireTransferSimulationTransactionSourceCategoryCheckTransferStopPaymentRequest             WireTransferSimulationTransactionSourceCategory = "check_transfer_stop_payment_request"
	WireTransferSimulationTransactionSourceCategoryDisputeResolution                           WireTransferSimulationTransactionSourceCategory = "dispute_resolution"
	WireTransferSimulationTransactionSourceCategoryEmpyrealCashDeposit                         WireTransferSimulationTransactionSourceCategory = "empyreal_cash_deposit"
	WireTransferSimulationTransactionSourceCategoryFeePayment                                  WireTransferSimulationTransactionSourceCategory = "fee_payment"
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
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency `json:"currency,required"`
	// The description you chose to give the transfer.
	Description string `json:"description,required"`
	// The identifier of the Account to where the Account Transfer was sent.
	DestinationAccountID string `json:"destination_account_id,required"`
	// The identifier of the Account from where the Account Transfer was sent.
	SourceAccountID string `json:"source_account_id,required"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       WireTransferSimulationTransactionSourceAccountTransferIntentionJSON
}

type WireTransferSimulationTransactionSourceAccountTransferIntentionJSON struct {
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
// WireTransferSimulationTransactionSourceAccountTransferIntention using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceAccountTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// Why the transfer was returned.
	ReturnReasonCode string `json:"return_reason_code,required"`
	JSON             WireTransferSimulationTransactionSourceACHCheckConversionReturnJSON
}

type WireTransferSimulationTransactionSourceACHCheckConversionReturnJSON struct {
	Amount           apijson.Metadata
	ReturnReasonCode apijson.Metadata
	raw              string
	Extras           map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceACHCheckConversionReturn using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceACHCheckConversionReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceACHCheckConversion struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the File containing an image of the returned check.
	FileID string `json:"file_id,required"`
	JSON   WireTransferSimulationTransactionSourceACHCheckConversionJSON
}

type WireTransferSimulationTransactionSourceACHCheckConversionJSON struct {
	Amount apijson.Metadata
	FileID apijson.Metadata
	raw    string
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceACHCheckConversion using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceACHCheckConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceACHTransferIntention struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount              int64  `json:"amount,required"`
	AccountNumber       string `json:"account_number,required"`
	RoutingNumber       string `json:"routing_number,required"`
	StatementDescriptor string `json:"statement_descriptor,required"`
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       WireTransferSimulationTransactionSourceACHTransferIntentionJSON
}

type WireTransferSimulationTransactionSourceACHTransferIntentionJSON struct {
	Amount              apijson.Metadata
	AccountNumber       apijson.Metadata
	RoutingNumber       apijson.Metadata
	StatementDescriptor apijson.Metadata
	TransferID          apijson.Metadata
	raw                 string
	Extras              map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceACHTransferIntention using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceACHTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       WireTransferSimulationTransactionSourceACHTransferRejectionJSON
}

type WireTransferSimulationTransactionSourceACHTransferRejectionJSON struct {
	TransferID apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceACHTransferRejection using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceACHTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode `json:"return_reason_code,required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID string `json:"transfer_id,required"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID string `json:"transaction_id,required"`
	JSON          WireTransferSimulationTransactionSourceACHTransferReturnJSON
}

type WireTransferSimulationTransactionSourceACHTransferReturnJSON struct {
	CreatedAt        apijson.Metadata
	ReturnReasonCode apijson.Metadata
	TransferID       apijson.Metadata
	TransactionID    apijson.Metadata
	raw              string
	Extras           map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceACHTransferReturn using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	AcceptedAt time.Time `json:"accepted_at,required" format:"date-time"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID string `json:"card_dispute_id,required"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID string `json:"transaction_id,required"`
	JSON          WireTransferSimulationTransactionSourceCardDisputeAcceptanceJSON
}

type WireTransferSimulationTransactionSourceCardDisputeAcceptanceJSON struct {
	AcceptedAt    apijson.Metadata
	CardDisputeID apijson.Metadata
	TransactionID apijson.Metadata
	raw           string
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCardDisputeAcceptance using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCardDisputeAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceCardRefund struct {
	// The Card Refund identifier.
	ID string `json:"id,required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency WireTransferSimulationTransactionSourceCardRefundCurrency `json:"currency,required"`
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
	Type WireTransferSimulationTransactionSourceCardRefundType `json:"type,required"`
	JSON WireTransferSimulationTransactionSourceCardRefundJSON
}

type WireTransferSimulationTransactionSourceCardRefundJSON struct {
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
// WireTransferSimulationTransactionSourceCardRefund using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCardRefund) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Currency WireTransferSimulationTransactionSourceCardSettlementCurrency `json:"currency,required"`
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
	Type WireTransferSimulationTransactionSourceCardSettlementType `json:"type,required"`
	JSON WireTransferSimulationTransactionSourceCardSettlementJSON
}

type WireTransferSimulationTransactionSourceCardSettlementJSON struct {
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
// WireTransferSimulationTransactionSourceCardSettlement using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCardSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type WireTransferSimulationTransactionSourceCardRevenuePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency WireTransferSimulationTransactionSourceCardRevenuePaymentCurrency `json:"currency,required"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time `json:"period_start,required" format:"date-time"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The account the card belonged to.
	TransactedOnAccountID string `json:"transacted_on_account_id,required,nullable"`
	JSON                  WireTransferSimulationTransactionSourceCardRevenuePaymentJSON
}

type WireTransferSimulationTransactionSourceCardRevenuePaymentJSON struct {
	Amount                apijson.Metadata
	Currency              apijson.Metadata
	PeriodStart           apijson.Metadata
	PeriodEnd             apijson.Metadata
	TransactedOnAccountID apijson.Metadata
	raw                   string
	Extras                map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCardRevenuePayment using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCardRevenuePayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceCardRevenuePaymentCurrency string

const (
	WireTransferSimulationTransactionSourceCardRevenuePaymentCurrencyCad WireTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "CAD"
	WireTransferSimulationTransactionSourceCardRevenuePaymentCurrencyChf WireTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "CHF"
	WireTransferSimulationTransactionSourceCardRevenuePaymentCurrencyEur WireTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "EUR"
	WireTransferSimulationTransactionSourceCardRevenuePaymentCurrencyGbp WireTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "GBP"
	WireTransferSimulationTransactionSourceCardRevenuePaymentCurrencyJpy WireTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "JPY"
	WireTransferSimulationTransactionSourceCardRevenuePaymentCurrencyUsd WireTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "USD"
)

type WireTransferSimulationTransactionSourceCheckDepositAcceptance struct {
	// The amount to be deposited in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency `json:"currency,required"`
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
	JSON           WireTransferSimulationTransactionSourceCheckDepositAcceptanceJSON
}

type WireTransferSimulationTransactionSourceCheckDepositAcceptanceJSON struct {
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
// WireTransferSimulationTransactionSourceCheckDepositAcceptance using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCheckDepositAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency WireTransferSimulationTransactionSourceCheckDepositReturnCurrency `json:"currency,required"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID string `json:"check_deposit_id,required"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID string                                                                `json:"transaction_id,required"`
	ReturnReason  WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason `json:"return_reason,required"`
	JSON          WireTransferSimulationTransactionSourceCheckDepositReturnJSON
}

type WireTransferSimulationTransactionSourceCheckDepositReturnJSON struct {
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
// WireTransferSimulationTransactionSourceCheckDepositReturn using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCheckDepositReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Currency WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency `json:"currency,required"`
	// The name that will be printed on the check.
	RecipientName string `json:"recipient_name,required"`
	// The identifier of the Check Transfer with which this is associated.
	TransferID string `json:"transfer_id,required"`
	JSON       WireTransferSimulationTransactionSourceCheckTransferIntentionJSON
}

type WireTransferSimulationTransactionSourceCheckTransferIntentionJSON struct {
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
// WireTransferSimulationTransactionSourceCheckTransferIntention using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCheckTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	TransferID string `json:"transfer_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// If available, a document with additional information about the return.
	FileID string `json:"file_id,required,nullable"`
	// The reason why the check was returned.
	Reason WireTransferSimulationTransactionSourceCheckTransferReturnReason `json:"reason,required"`
	// The identifier of the Transaction that was created to credit you for the
	// returned check.
	TransactionID string `json:"transaction_id,required,nullable"`
	JSON          WireTransferSimulationTransactionSourceCheckTransferReturnJSON
}

type WireTransferSimulationTransactionSourceCheckTransferReturnJSON struct {
	TransferID    apijson.Metadata
	ReturnedAt    apijson.Metadata
	FileID        apijson.Metadata
	Reason        apijson.Metadata
	TransactionID apijson.Metadata
	raw           string
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCheckTransferReturn using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCheckTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceCheckTransferReturnReason string

const (
	WireTransferSimulationTransactionSourceCheckTransferReturnReasonMailDeliveryFailure WireTransferSimulationTransactionSourceCheckTransferReturnReason = "mail_delivery_failure"
	WireTransferSimulationTransactionSourceCheckTransferReturnReasonRefusedByRecipient  WireTransferSimulationTransactionSourceCheckTransferReturnReason = "refused_by_recipient"
)

type WireTransferSimulationTransactionSourceCheckTransferRejection struct {
	// The identifier of the Check Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       WireTransferSimulationTransactionSourceCheckTransferRejectionJSON
}

type WireTransferSimulationTransactionSourceCheckTransferRejectionJSON struct {
	TransferID apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCheckTransferRejection using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCheckTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferID string `json:"transfer_id,required"`
	// The transaction ID of the corresponding credit transaction.
	TransactionID string `json:"transaction_id,required"`
	// The time the stop-payment was requested.
	RequestedAt time.Time `json:"requested_at,required" format:"date-time"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType `json:"type,required"`
	JSON WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestJSON
}

type WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestJSON struct {
	TransferID    apijson.Metadata
	TransactionID apijson.Metadata
	RequestedAt   apijson.Metadata
	Type          apijson.Metadata
	raw           string
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType string

const (
	WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

type WireTransferSimulationTransactionSourceDisputeResolution struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency WireTransferSimulationTransactionSourceDisputeResolutionCurrency `json:"currency,required"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionID string `json:"disputed_transaction_id,required"`
	JSON                  WireTransferSimulationTransactionSourceDisputeResolutionJSON
}

type WireTransferSimulationTransactionSourceDisputeResolutionJSON struct {
	Amount                apijson.Metadata
	Currency              apijson.Metadata
	DisputedTransactionID apijson.Metadata
	raw                   string
	Extras                map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceDisputeResolution using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceDisputeResolution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount      int64     `json:"amount,required"`
	BagID       string    `json:"bag_id,required"`
	DepositDate time.Time `json:"deposit_date,required" format:"date-time"`
	JSON        WireTransferSimulationTransactionSourceEmpyrealCashDepositJSON
}

type WireTransferSimulationTransactionSourceEmpyrealCashDepositJSON struct {
	Amount      apijson.Metadata
	BagID       apijson.Metadata
	DepositDate apijson.Metadata
	raw         string
	Extras      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceEmpyrealCashDeposit using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceEmpyrealCashDeposit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceFeePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency WireTransferSimulationTransactionSourceFeePaymentCurrency `json:"currency,required"`
	JSON     WireTransferSimulationTransactionSourceFeePaymentJSON
}

type WireTransferSimulationTransactionSourceFeePaymentJSON struct {
	Amount   apijson.Metadata
	Currency apijson.Metadata
	raw      string
	Extras   map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceFeePayment using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceFeePayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceFeePaymentCurrency string

const (
	WireTransferSimulationTransactionSourceFeePaymentCurrencyCad WireTransferSimulationTransactionSourceFeePaymentCurrency = "CAD"
	WireTransferSimulationTransactionSourceFeePaymentCurrencyChf WireTransferSimulationTransactionSourceFeePaymentCurrency = "CHF"
	WireTransferSimulationTransactionSourceFeePaymentCurrencyEur WireTransferSimulationTransactionSourceFeePaymentCurrency = "EUR"
	WireTransferSimulationTransactionSourceFeePaymentCurrencyGbp WireTransferSimulationTransactionSourceFeePaymentCurrency = "GBP"
	WireTransferSimulationTransactionSourceFeePaymentCurrencyJpy WireTransferSimulationTransactionSourceFeePaymentCurrency = "JPY"
	WireTransferSimulationTransactionSourceFeePaymentCurrencyUsd WireTransferSimulationTransactionSourceFeePaymentCurrency = "USD"
)

type WireTransferSimulationTransactionSourceInboundACHTransfer struct {
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
	JSON                               WireTransferSimulationTransactionSourceInboundACHTransferJSON
}

type WireTransferSimulationTransactionSourceInboundACHTransferJSON struct {
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
// WireTransferSimulationTransactionSourceInboundACHTransfer using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInboundACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceInboundCheck struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency              WireTransferSimulationTransactionSourceInboundCheckCurrency `json:"currency,required"`
	CheckNumber           string                                                      `json:"check_number,required,nullable"`
	CheckFrontImageFileID string                                                      `json:"check_front_image_file_id,required,nullable"`
	CheckRearImageFileID  string                                                      `json:"check_rear_image_file_id,required,nullable"`
	JSON                  WireTransferSimulationTransactionSourceInboundCheckJSON
}

type WireTransferSimulationTransactionSourceInboundCheckJSON struct {
	Amount                apijson.Metadata
	Currency              apijson.Metadata
	CheckNumber           apijson.Metadata
	CheckFrontImageFileID apijson.Metadata
	CheckRearImageFileID  apijson.Metadata
	raw                   string
	Extras                map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceInboundCheck using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInboundCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	JSON                                                   WireTransferSimulationTransactionSourceInboundInternationalACHTransferJSON
}

type WireTransferSimulationTransactionSourceInboundInternationalACHTransferJSON struct {
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
// WireTransferSimulationTransactionSourceInboundInternationalACHTransfer using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real Time Payments transfer.
	Currency WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency `json:"currency,required"`
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
	JSON                  WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON
}

type WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON struct {
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
// WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	JSON                       WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversalJSON
}

type WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversalJSON struct {
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
// WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceInboundWireDrawdownPayment struct {
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
	JSON                               WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentJSON
}

type WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentJSON struct {
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
// WireTransferSimulationTransactionSourceInboundWireDrawdownPayment using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceInboundWireReversal struct {
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
	JSON          WireTransferSimulationTransactionSourceInboundWireReversalJSON
}

type WireTransferSimulationTransactionSourceInboundWireReversalJSON struct {
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
// WireTransferSimulationTransactionSourceInboundWireReversal using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInboundWireReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceInboundWireTransfer struct {
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
	JSON                                    WireTransferSimulationTransactionSourceInboundWireTransferJSON
}

type WireTransferSimulationTransactionSourceInboundWireTransferJSON struct {
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
// WireTransferSimulationTransactionSourceInboundWireTransfer using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInboundWireTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceInterestPayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency WireTransferSimulationTransactionSourceInterestPaymentCurrency `json:"currency,required"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time `json:"period_start,required" format:"date-time"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The account on which the interest was accrued.
	AccruedOnAccountID string `json:"accrued_on_account_id,required,nullable"`
	JSON               WireTransferSimulationTransactionSourceInterestPaymentJSON
}

type WireTransferSimulationTransactionSourceInterestPaymentJSON struct {
	Amount             apijson.Metadata
	Currency           apijson.Metadata
	PeriodStart        apijson.Metadata
	PeriodEnd          apijson.Metadata
	AccruedOnAccountID apijson.Metadata
	raw                string
	Extras             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceInterestPayment using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInterestPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency WireTransferSimulationTransactionSourceInternalSourceCurrency `json:"currency,required"`
	Reason   WireTransferSimulationTransactionSourceInternalSourceReason   `json:"reason,required"`
	JSON     WireTransferSimulationTransactionSourceInternalSourceJSON
}

type WireTransferSimulationTransactionSourceInternalSourceJSON struct {
	Amount   apijson.Metadata
	Currency apijson.Metadata
	Reason   apijson.Metadata
	raw      string
	Extras   map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceInternalSource using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInternalSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
	// currency.
	Currency             WireTransferSimulationTransactionSourceCardRouteRefundCurrency `json:"currency,required"`
	MerchantAcceptorID   string                                                         `json:"merchant_acceptor_id,required"`
	MerchantCity         string                                                         `json:"merchant_city,required,nullable"`
	MerchantCountry      string                                                         `json:"merchant_country,required"`
	MerchantDescriptor   string                                                         `json:"merchant_descriptor,required"`
	MerchantState        string                                                         `json:"merchant_state,required,nullable"`
	MerchantCategoryCode string                                                         `json:"merchant_category_code,required,nullable"`
	JSON                 WireTransferSimulationTransactionSourceCardRouteRefundJSON
}

type WireTransferSimulationTransactionSourceCardRouteRefundJSON struct {
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
// WireTransferSimulationTransactionSourceCardRouteRefund using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCardRouteRefund) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
	// currency.
	Currency             WireTransferSimulationTransactionSourceCardRouteSettlementCurrency `json:"currency,required"`
	MerchantAcceptorID   string                                                             `json:"merchant_acceptor_id,required"`
	MerchantCity         string                                                             `json:"merchant_city,required,nullable"`
	MerchantCountry      string                                                             `json:"merchant_country,required,nullable"`
	MerchantDescriptor   string                                                             `json:"merchant_descriptor,required"`
	MerchantState        string                                                             `json:"merchant_state,required,nullable"`
	MerchantCategoryCode string                                                             `json:"merchant_category_code,required,nullable"`
	JSON                 WireTransferSimulationTransactionSourceCardRouteSettlementJSON
}

type WireTransferSimulationTransactionSourceCardRouteSettlementJSON struct {
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
// WireTransferSimulationTransactionSourceCardRouteSettlement using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCardRouteSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type WireTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgement struct {
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
	JSON       WireTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgementJSON
}

type WireTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgementJSON struct {
	Amount                   apijson.Metadata
	DestinationAccountNumber apijson.Metadata
	DestinationRoutingNumber apijson.Metadata
	RemittanceInformation    apijson.Metadata
	TransferID               apijson.Metadata
	raw                      string
	Extras                   map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgement
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceSampleFunds struct {
	// Where the sample funds came from.
	Originator string `json:"originator,required"`
	JSON       WireTransferSimulationTransactionSourceSampleFundsJSON
}

type WireTransferSimulationTransactionSourceSampleFundsJSON struct {
	Originator apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceSampleFunds using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceSampleFunds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention struct {
	// The transfer amount in USD cents.
	Amount             int64  `json:"amount,required"`
	AccountNumber      string `json:"account_number,required"`
	RoutingNumber      string `json:"routing_number,required"`
	MessageToRecipient string `json:"message_to_recipient,required"`
	TransferID         string `json:"transfer_id,required"`
	JSON               WireTransferSimulationTransactionSourceWireDrawdownPaymentIntentionJSON
}

type WireTransferSimulationTransactionSourceWireDrawdownPaymentIntentionJSON struct {
	Amount             apijson.Metadata
	AccountNumber      apijson.Metadata
	RoutingNumber      apijson.Metadata
	MessageToRecipient apijson.Metadata
	TransferID         apijson.Metadata
	raw                string
	Extras             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection struct {
	TransferID string `json:"transfer_id,required"`
	JSON       WireTransferSimulationTransactionSourceWireDrawdownPaymentRejectionJSON
}

type WireTransferSimulationTransactionSourceWireDrawdownPaymentRejectionJSON struct {
	TransferID apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceWireTransferIntention struct {
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient string `json:"message_to_recipient,required"`
	TransferID         string `json:"transfer_id,required"`
	JSON               WireTransferSimulationTransactionSourceWireTransferIntentionJSON
}

type WireTransferSimulationTransactionSourceWireTransferIntentionJSON struct {
	Amount             apijson.Metadata
	AccountNumber      apijson.Metadata
	RoutingNumber      apijson.Metadata
	MessageToRecipient apijson.Metadata
	TransferID         apijson.Metadata
	raw                string
	Extras             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceWireTransferIntention using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceWireTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceWireTransferRejection struct {
	TransferID string `json:"transfer_id,required"`
	JSON       WireTransferSimulationTransactionSourceWireTransferRejectionJSON
}

type WireTransferSimulationTransactionSourceWireTransferRejectionJSON struct {
	TransferID apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceWireTransferRejection using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceWireTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionType string

const (
	WireTransferSimulationTransactionTypeTransaction WireTransferSimulationTransactionType = "transaction"
)

type WireTransferSimulationType string

const (
	WireTransferSimulationTypeInboundWireTransferSimulationResult WireTransferSimulationType = "inbound_wire_transfer_simulation_result"
)
