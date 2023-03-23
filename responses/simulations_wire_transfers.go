package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
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
	Transaction pjson.Metadata
	Type        pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into WireTransferSimulation using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulation) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	RouteType string `json:"route_type,required,nullable"`
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
	AccountID   pjson.Metadata
	Amount      pjson.Metadata
	Currency    pjson.Metadata
	CreatedAt   pjson.Metadata
	Description pjson.Metadata
	ID          pjson.Metadata
	RouteID     pjson.Metadata
	RouteType   pjson.Metadata
	Source      pjson.Metadata
	Type        pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransaction using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransaction) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Category                                    pjson.Metadata
	AccountTransferIntention                    pjson.Metadata
	ACHCheckConversionReturn                    pjson.Metadata
	ACHCheckConversion                          pjson.Metadata
	ACHTransferIntention                        pjson.Metadata
	ACHTransferRejection                        pjson.Metadata
	ACHTransferReturn                           pjson.Metadata
	CardDisputeAcceptance                       pjson.Metadata
	CardRefund                                  pjson.Metadata
	CardSettlement                              pjson.Metadata
	CheckDepositAcceptance                      pjson.Metadata
	CheckDepositReturn                          pjson.Metadata
	CheckTransferIntention                      pjson.Metadata
	CheckTransferReturn                         pjson.Metadata
	CheckTransferRejection                      pjson.Metadata
	CheckTransferStopPaymentRequest             pjson.Metadata
	DisputeResolution                           pjson.Metadata
	EmpyrealCashDeposit                         pjson.Metadata
	InboundACHTransfer                          pjson.Metadata
	InboundCheck                                pjson.Metadata
	InboundInternationalACHTransfer             pjson.Metadata
	InboundRealTimePaymentsTransferConfirmation pjson.Metadata
	InboundWireDrawdownPaymentReversal          pjson.Metadata
	InboundWireDrawdownPayment                  pjson.Metadata
	InboundWireReversal                         pjson.Metadata
	InboundWireTransfer                         pjson.Metadata
	InterestPayment                             pjson.Metadata
	InternalSource                              pjson.Metadata
	CardRouteRefund                             pjson.Metadata
	CardRouteSettlement                         pjson.Metadata
	SampleFunds                                 pjson.Metadata
	WireDrawdownPaymentIntention                pjson.Metadata
	WireDrawdownPaymentRejection                pjson.Metadata
	WireTransferIntention                       pjson.Metadata
	WireTransferRejection                       pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSource using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount               pjson.Metadata
	Currency             pjson.Metadata
	Description          pjson.Metadata
	DestinationAccountID pjson.Metadata
	SourceAccountID      pjson.Metadata
	TransferID           pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceAccountTransferIntention using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceAccountTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount           pjson.Metadata
	ReturnReasonCode pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceACHCheckConversionReturn using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceACHCheckConversionReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount pjson.Metadata
	FileID pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceACHCheckConversion using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceACHCheckConversion) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount              pjson.Metadata
	AccountNumber       pjson.Metadata
	RoutingNumber       pjson.Metadata
	StatementDescriptor pjson.Metadata
	TransferID          pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceACHTransferIntention using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceACHTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       WireTransferSimulationTransactionSourceACHTransferRejectionJSON
}

type WireTransferSimulationTransactionSourceACHTransferRejectionJSON struct {
	TransferID pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceACHTransferRejection using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceACHTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	CreatedAt        pjson.Metadata
	ReturnReasonCode pjson.Metadata
	TransferID       pjson.Metadata
	TransactionID    pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceACHTransferReturn using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	TransactionID string `json:"transaction_id,required,nullable"`
	JSON          WireTransferSimulationTransactionSourceCardDisputeAcceptanceJSON
}

type WireTransferSimulationTransactionSourceCardDisputeAcceptanceJSON struct {
	AcceptedAt    pjson.Metadata
	CardDisputeID pjson.Metadata
	TransactionID pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCardDisputeAcceptance using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCardDisputeAcceptance) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceCardRefund struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency WireTransferSimulationTransactionSourceCardRefundCurrency `json:"currency,required"`
	// The identifier for the Transaction this refunds, if any.
	CardSettlementTransactionID string `json:"card_settlement_transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type WireTransferSimulationTransactionSourceCardRefundType `json:"type,required"`
	JSON WireTransferSimulationTransactionSourceCardRefundJSON
}

type WireTransferSimulationTransactionSourceCardRefundJSON struct {
	Amount                      pjson.Metadata
	Currency                    pjson.Metadata
	CardSettlementTransactionID pjson.Metadata
	Type                        pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCardRefund using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCardRefund) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's settlement currency.
	Currency WireTransferSimulationTransactionSourceCardSettlementCurrency `json:"currency,required"`
	// The amount in the minor unit of the transaction's presentment currency.
	PresentmentAmount int64 `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's presentment currency.
	PresentmentCurrency  string `json:"presentment_currency,required"`
	MerchantCity         string `json:"merchant_city,required,nullable"`
	MerchantCountry      string `json:"merchant_country,required"`
	MerchantName         string `json:"merchant_name,required,nullable"`
	MerchantCategoryCode string `json:"merchant_category_code,required"`
	MerchantState        string `json:"merchant_state,required,nullable"`
	// The identifier of the Pending Transaction associated with this Transaction.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_settlement`.
	Type WireTransferSimulationTransactionSourceCardSettlementType `json:"type,required"`
	JSON WireTransferSimulationTransactionSourceCardSettlementJSON
}

type WireTransferSimulationTransactionSourceCardSettlementJSON struct {
	Amount               pjson.Metadata
	Currency             pjson.Metadata
	PresentmentAmount    pjson.Metadata
	PresentmentCurrency  pjson.Metadata
	MerchantCity         pjson.Metadata
	MerchantCountry      pjson.Metadata
	MerchantName         pjson.Metadata
	MerchantCategoryCode pjson.Metadata
	MerchantState        pjson.Metadata
	PendingTransactionID pjson.Metadata
	Type                 pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCardSettlement using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCardSettlement) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount         pjson.Metadata
	Currency       pjson.Metadata
	AccountNumber  pjson.Metadata
	RoutingNumber  pjson.Metadata
	AuxiliaryOnUs  pjson.Metadata
	SerialNumber   pjson.Metadata
	CheckDepositID pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCheckDepositAcceptance using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCheckDepositAcceptance) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount         pjson.Metadata
	ReturnedAt     pjson.Metadata
	Currency       pjson.Metadata
	CheckDepositID pjson.Metadata
	TransactionID  pjson.Metadata
	ReturnReason   pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCheckDepositReturn using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCheckDepositReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	AddressLine1  pjson.Metadata
	AddressLine2  pjson.Metadata
	AddressCity   pjson.Metadata
	AddressState  pjson.Metadata
	AddressZip    pjson.Metadata
	Amount        pjson.Metadata
	Currency      pjson.Metadata
	RecipientName pjson.Metadata
	TransferID    pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCheckTransferIntention using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCheckTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	// If available, a document with additional information about the return.
	FileID string `json:"file_id,required,nullable"`
	// The reason why the check was returned.
	Reason WireTransferSimulationTransactionSourceCheckTransferReturnReason `json:"reason,required"`
	JSON   WireTransferSimulationTransactionSourceCheckTransferReturnJSON
}

type WireTransferSimulationTransactionSourceCheckTransferReturnJSON struct {
	TransferID pjson.Metadata
	FileID     pjson.Metadata
	Reason     pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCheckTransferReturn using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCheckTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	TransferID pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCheckTransferRejection using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCheckTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	TransferID    pjson.Metadata
	TransactionID pjson.Metadata
	RequestedAt   pjson.Metadata
	Type          pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount                pjson.Metadata
	Currency              pjson.Metadata
	DisputedTransactionID pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceDisputeResolution using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceDisputeResolution) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount      pjson.Metadata
	BagID       pjson.Metadata
	DepositDate pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceEmpyrealCashDeposit using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceEmpyrealCashDeposit) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

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
	Amount                             pjson.Metadata
	OriginatorCompanyName              pjson.Metadata
	OriginatorCompanyDescriptiveDate   pjson.Metadata
	OriginatorCompanyDiscretionaryData pjson.Metadata
	OriginatorCompanyEntryDescription  pjson.Metadata
	OriginatorCompanyID                pjson.Metadata
	ReceiverIDNumber                   pjson.Metadata
	ReceiverName                       pjson.Metadata
	TraceNumber                        pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceInboundACHTransfer using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInboundACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount                pjson.Metadata
	Currency              pjson.Metadata
	CheckNumber           pjson.Metadata
	CheckFrontImageFileID pjson.Metadata
	CheckRearImageFileID  pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceInboundCheck using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInboundCheck) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount                                                 pjson.Metadata
	ForeignExchangeIndicator                               pjson.Metadata
	ForeignExchangeReferenceIndicator                      pjson.Metadata
	ForeignExchangeReference                               pjson.Metadata
	DestinationCountryCode                                 pjson.Metadata
	DestinationCurrencyCode                                pjson.Metadata
	ForeignPaymentAmount                                   pjson.Metadata
	ForeignTraceNumber                                     pjson.Metadata
	InternationalTransactionTypeCode                       pjson.Metadata
	OriginatingCurrencyCode                                pjson.Metadata
	OriginatingDepositoryFinancialInstitutionName          pjson.Metadata
	OriginatingDepositoryFinancialInstitutionIDQualifier   pjson.Metadata
	OriginatingDepositoryFinancialInstitutionID            pjson.Metadata
	OriginatingDepositoryFinancialInstitutionBranchCountry pjson.Metadata
	OriginatorCity                                         pjson.Metadata
	OriginatorCompanyEntryDescription                      pjson.Metadata
	OriginatorCountry                                      pjson.Metadata
	OriginatorIdentification                               pjson.Metadata
	OriginatorName                                         pjson.Metadata
	OriginatorPostalCode                                   pjson.Metadata
	OriginatorStreetAddress                                pjson.Metadata
	OriginatorStateOrProvince                              pjson.Metadata
	PaymentRelatedInformation                              pjson.Metadata
	PaymentRelatedInformation2                             pjson.Metadata
	ReceiverIdentificationNumber                           pjson.Metadata
	ReceiverStreetAddress                                  pjson.Metadata
	ReceiverCity                                           pjson.Metadata
	ReceiverStateOrProvince                                pjson.Metadata
	ReceiverCountry                                        pjson.Metadata
	ReceiverPostalCode                                     pjson.Metadata
	ReceivingCompanyOrIndividualName                       pjson.Metadata
	ReceivingDepositoryFinancialInstitutionName            pjson.Metadata
	ReceivingDepositoryFinancialInstitutionIDQualifier     pjson.Metadata
	ReceivingDepositoryFinancialInstitutionID              pjson.Metadata
	ReceivingDepositoryFinancialInstitutionCountry         pjson.Metadata
	TraceNumber                                            pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceInboundInternationalACHTransfer using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount                    pjson.Metadata
	Currency                  pjson.Metadata
	CreditorName              pjson.Metadata
	DebtorName                pjson.Metadata
	DebtorAccountNumber       pjson.Metadata
	DebtorRoutingNumber       pjson.Metadata
	TransactionIdentification pjson.Metadata
	RemittanceInformation     pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount                                        pjson.Metadata
	Description                                   pjson.Metadata
	InputCycleDate                                pjson.Metadata
	InputSequenceNumber                           pjson.Metadata
	InputSource                                   pjson.Metadata
	InputMessageAccountabilityData                pjson.Metadata
	PreviousMessageInputMessageAccountabilityData pjson.Metadata
	PreviousMessageInputCycleDate                 pjson.Metadata
	PreviousMessageInputSequenceNumber            pjson.Metadata
	PreviousMessageInputSource                    pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount                             pjson.Metadata
	BeneficiaryAddressLine1            pjson.Metadata
	BeneficiaryAddressLine2            pjson.Metadata
	BeneficiaryAddressLine3            pjson.Metadata
	BeneficiaryName                    pjson.Metadata
	BeneficiaryReference               pjson.Metadata
	Description                        pjson.Metadata
	InputMessageAccountabilityData     pjson.Metadata
	OriginatorAddressLine1             pjson.Metadata
	OriginatorAddressLine2             pjson.Metadata
	OriginatorAddressLine3             pjson.Metadata
	OriginatorName                     pjson.Metadata
	OriginatorToBeneficiaryInformation pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceInboundWireDrawdownPayment using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceInboundWireReversal struct {
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
	// Information included in the wire reversal for the receiving financial
	// institution.
	ReceiverFinancialInstitutionInformation string `json:"receiver_financial_institution_information,required,nullable"`
	// Additional financial institution information included in the wire reversal.
	FinancialInstitutionToFinancialInstitutionInformation string `json:"financial_institution_to_financial_institution_information,required,nullable"`
	JSON                                                  WireTransferSimulationTransactionSourceInboundWireReversalJSON
}

type WireTransferSimulationTransactionSourceInboundWireReversalJSON struct {
	Amount                                                pjson.Metadata
	Description                                           pjson.Metadata
	InputCycleDate                                        pjson.Metadata
	InputSequenceNumber                                   pjson.Metadata
	InputSource                                           pjson.Metadata
	InputMessageAccountabilityData                        pjson.Metadata
	PreviousMessageInputMessageAccountabilityData         pjson.Metadata
	PreviousMessageInputCycleDate                         pjson.Metadata
	PreviousMessageInputSequenceNumber                    pjson.Metadata
	PreviousMessageInputSource                            pjson.Metadata
	ReceiverFinancialInstitutionInformation               pjson.Metadata
	FinancialInstitutionToFinancialInstitutionInformation pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceInboundWireReversal using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInboundWireReversal) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount                                  pjson.Metadata
	BeneficiaryAddressLine1                 pjson.Metadata
	BeneficiaryAddressLine2                 pjson.Metadata
	BeneficiaryAddressLine3                 pjson.Metadata
	BeneficiaryName                         pjson.Metadata
	BeneficiaryReference                    pjson.Metadata
	Description                             pjson.Metadata
	InputMessageAccountabilityData          pjson.Metadata
	OriginatorAddressLine1                  pjson.Metadata
	OriginatorAddressLine2                  pjson.Metadata
	OriginatorAddressLine3                  pjson.Metadata
	OriginatorName                          pjson.Metadata
	OriginatorToBeneficiaryInformationLine1 pjson.Metadata
	OriginatorToBeneficiaryInformationLine2 pjson.Metadata
	OriginatorToBeneficiaryInformationLine3 pjson.Metadata
	OriginatorToBeneficiaryInformationLine4 pjson.Metadata
	OriginatorToBeneficiaryInformation      pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceInboundWireTransfer using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInboundWireTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount             pjson.Metadata
	Currency           pjson.Metadata
	PeriodStart        pjson.Metadata
	PeriodEnd          pjson.Metadata
	AccruedOnAccountID pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceInterestPayment using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInterestPayment) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount   pjson.Metadata
	Currency pjson.Metadata
	Reason   pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceInternalSource using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceInternalSource) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount               pjson.Metadata
	Currency             pjson.Metadata
	MerchantAcceptorID   pjson.Metadata
	MerchantCity         pjson.Metadata
	MerchantCountry      pjson.Metadata
	MerchantDescriptor   pjson.Metadata
	MerchantState        pjson.Metadata
	MerchantCategoryCode pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCardRouteRefund using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCardRouteRefund) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount               pjson.Metadata
	Currency             pjson.Metadata
	MerchantAcceptorID   pjson.Metadata
	MerchantCity         pjson.Metadata
	MerchantCountry      pjson.Metadata
	MerchantDescriptor   pjson.Metadata
	MerchantState        pjson.Metadata
	MerchantCategoryCode pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceCardRouteSettlement using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceCardRouteSettlement) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Originator string `json:"originator,required"`
	JSON       WireTransferSimulationTransactionSourceSampleFundsJSON
}

type WireTransferSimulationTransactionSourceSampleFundsJSON struct {
	Originator pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceSampleFunds using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceSampleFunds) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount             pjson.Metadata
	AccountNumber      pjson.Metadata
	RoutingNumber      pjson.Metadata
	MessageToRecipient pjson.Metadata
	TransferID         pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection struct {
	TransferID string `json:"transfer_id,required"`
	JSON       WireTransferSimulationTransactionSourceWireDrawdownPaymentRejectionJSON
}

type WireTransferSimulationTransactionSourceWireDrawdownPaymentRejectionJSON struct {
	TransferID pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount             pjson.Metadata
	AccountNumber      pjson.Metadata
	RoutingNumber      pjson.Metadata
	MessageToRecipient pjson.Metadata
	TransferID         pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceWireTransferIntention using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceWireTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceWireTransferRejection struct {
	TransferID string `json:"transfer_id,required"`
	JSON       WireTransferSimulationTransactionSourceWireTransferRejectionJSON
}

type WireTransferSimulationTransactionSourceWireTransferRejectionJSON struct {
	TransferID pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransferSimulationTransactionSourceWireTransferRejection using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *WireTransferSimulationTransactionSourceWireTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionType string

const (
	WireTransferSimulationTransactionTypeTransaction WireTransferSimulationTransactionType = "transaction"
)

type WireTransferSimulationType string

const (
	WireTransferSimulationTypeInboundWireTransferSimulationResult WireTransferSimulationType = "inbound_wire_transfer_simulation_result"
)
