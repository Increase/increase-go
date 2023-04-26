package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
)

type InterestPaymentSimulationResult struct {
	// This will contain the resulting [Transaction](#transactions) object. The
	// Transaction's `source` will be of `category: interest_payment`.
	Transaction InterestPaymentSimulationResultTransaction `json:"transaction,required"`
	// A constant representing the object's type. For this resource it will always be
	// `interest_payment_simulation_result`.
	Type InterestPaymentSimulationResultType `json:"type,required"`
	JSON InterestPaymentSimulationResultJSON
}

type InterestPaymentSimulationResultJSON struct {
	Transaction pjson.Metadata
	Type        pjson.Metadata
	Raw         []byte
	Extras      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResult using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *InterestPaymentSimulationResult) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransaction struct {
	// The identifier for the Account the Transaction belongs to.
	AccountID string `json:"account_id,required"`
	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transcation's
	// Account.
	Currency InterestPaymentSimulationResultTransactionCurrency `json:"currency,required"`
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
	RouteType InterestPaymentSimulationResultTransactionRouteType `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Transaction. Note that for backwards compatibility reasons, additional
	// undocumented keys may appear in this object. These should be treated as
	// deprecated and will be removed in the future.
	Source InterestPaymentSimulationResultTransactionSource `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type InterestPaymentSimulationResultTransactionType `json:"type,required"`
	JSON InterestPaymentSimulationResultTransactionJSON
}

type InterestPaymentSimulationResultTransactionJSON struct {
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
	Raw         []byte
	Extras      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransaction using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *InterestPaymentSimulationResultTransaction) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionCurrency string

const (
	InterestPaymentSimulationResultTransactionCurrencyCad InterestPaymentSimulationResultTransactionCurrency = "CAD"
	InterestPaymentSimulationResultTransactionCurrencyChf InterestPaymentSimulationResultTransactionCurrency = "CHF"
	InterestPaymentSimulationResultTransactionCurrencyEur InterestPaymentSimulationResultTransactionCurrency = "EUR"
	InterestPaymentSimulationResultTransactionCurrencyGbp InterestPaymentSimulationResultTransactionCurrency = "GBP"
	InterestPaymentSimulationResultTransactionCurrencyJpy InterestPaymentSimulationResultTransactionCurrency = "JPY"
	InterestPaymentSimulationResultTransactionCurrencyUsd InterestPaymentSimulationResultTransactionCurrency = "USD"
)

type InterestPaymentSimulationResultTransactionRouteType string

const (
	InterestPaymentSimulationResultTransactionRouteTypeAccountNumber InterestPaymentSimulationResultTransactionRouteType = "account_number"
	InterestPaymentSimulationResultTransactionRouteTypeCard          InterestPaymentSimulationResultTransactionRouteType = "card"
)

type InterestPaymentSimulationResultTransactionSource struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category InterestPaymentSimulationResultTransactionSourceCategory `json:"category,required"`
	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention InterestPaymentSimulationResultTransactionSourceAccountTransferIntention `json:"account_transfer_intention,required,nullable"`
	// A ACH Check Conversion Return object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_check_conversion_return`.
	ACHCheckConversionReturn InterestPaymentSimulationResultTransactionSourceACHCheckConversionReturn `json:"ach_check_conversion_return,required,nullable"`
	// A ACH Check Conversion object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_check_conversion`.
	ACHCheckConversion InterestPaymentSimulationResultTransactionSourceACHCheckConversion `json:"ach_check_conversion,required,nullable"`
	// A ACH Transfer Intention object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention InterestPaymentSimulationResultTransactionSourceACHTransferIntention `json:"ach_transfer_intention,required,nullable"`
	// A ACH Transfer Rejection object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection InterestPaymentSimulationResultTransactionSourceACHTransferRejection `json:"ach_transfer_rejection,required,nullable"`
	// A ACH Transfer Return object. This field will be present in the JSON response if
	// and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn InterestPaymentSimulationResultTransactionSourceACHTransferReturn `json:"ach_transfer_return,required,nullable"`
	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance InterestPaymentSimulationResultTransactionSourceCardDisputeAcceptance `json:"card_dispute_acceptance,required,nullable"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund InterestPaymentSimulationResultTransactionSourceCardRefund `json:"card_refund,required,nullable"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement InterestPaymentSimulationResultTransactionSourceCardSettlement `json:"card_settlement,required,nullable"`
	// A Card Revenue Payment object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_revenue_payment`.
	CardRevenuePayment InterestPaymentSimulationResultTransactionSourceCardRevenuePayment `json:"card_revenue_payment,required,nullable"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptance `json:"check_deposit_acceptance,required,nullable"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn InterestPaymentSimulationResultTransactionSourceCheckDepositReturn `json:"check_deposit_return,required,nullable"`
	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention InterestPaymentSimulationResultTransactionSourceCheckTransferIntention `json:"check_transfer_intention,required,nullable"`
	// A Check Transfer Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_transfer_return`.
	CheckTransferReturn InterestPaymentSimulationResultTransactionSourceCheckTransferReturn `json:"check_transfer_return,required,nullable"`
	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection InterestPaymentSimulationResultTransactionSourceCheckTransferRejection `json:"check_transfer_rejection,required,nullable"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequest `json:"check_transfer_stop_payment_request,required,nullable"`
	// A Dispute Resolution object. This field will be present in the JSON response if
	// and only if `category` is equal to `dispute_resolution`.
	DisputeResolution InterestPaymentSimulationResultTransactionSourceDisputeResolution `json:"dispute_resolution,required,nullable"`
	// A Empyreal Cash Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `empyreal_cash_deposit`.
	EmpyrealCashDeposit InterestPaymentSimulationResultTransactionSourceEmpyrealCashDeposit `json:"empyreal_cash_deposit,required,nullable"`
	// A Fee Payment object. This field will be present in the JSON response if and
	// only if `category` is equal to `fee_payment`.
	FeePayment InterestPaymentSimulationResultTransactionSourceFeePayment `json:"fee_payment,required,nullable"`
	// A Inbound ACH Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer InterestPaymentSimulationResultTransactionSourceInboundACHTransfer `json:"inbound_ach_transfer,required,nullable"`
	// A Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck InterestPaymentSimulationResultTransactionSourceInboundCheck `json:"inbound_check,required,nullable"`
	// A Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer InterestPaymentSimulationResultTransactionSourceInboundInternationalACHTransfer `json:"inbound_international_ach_transfer,required,nullable"`
	// A Inbound Real Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation `json:"inbound_real_time_payments_transfer_confirmation,required,nullable"`
	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal InterestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal `json:"inbound_wire_drawdown_payment_reversal,required,nullable"`
	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment InterestPaymentSimulationResultTransactionSourceInboundWireDrawdownPayment `json:"inbound_wire_drawdown_payment,required,nullable"`
	// A Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal InterestPaymentSimulationResultTransactionSourceInboundWireReversal `json:"inbound_wire_reversal,required,nullable"`
	// A Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer InterestPaymentSimulationResultTransactionSourceInboundWireTransfer `json:"inbound_wire_transfer,required,nullable"`
	// A Interest Payment object. This field will be present in the JSON response if
	// and only if `category` is equal to `interest_payment`.
	InterestPayment InterestPaymentSimulationResultTransactionSourceInterestPayment `json:"interest_payment,required,nullable"`
	// A Internal Source object. This field will be present in the JSON response if and
	// only if `category` is equal to `internal_source`.
	InternalSource InterestPaymentSimulationResultTransactionSourceInternalSource `json:"internal_source,required,nullable"`
	// A Deprecated Card Refund object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_route_refund`.
	CardRouteRefund InterestPaymentSimulationResultTransactionSourceCardRouteRefund `json:"card_route_refund,required,nullable"`
	// A Deprecated Card Settlement object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_settlement`.
	CardRouteSettlement InterestPaymentSimulationResultTransactionSourceCardRouteSettlement `json:"card_route_settlement,required,nullable"`
	// A Real Time Payments Transfer Acknowledgement object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `real_time_payments_transfer_acknowledgement`.
	RealTimePaymentsTransferAcknowledgement InterestPaymentSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgement `json:"real_time_payments_transfer_acknowledgement,required,nullable"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds InterestPaymentSimulationResultTransactionSourceSampleFunds `json:"sample_funds,required,nullable"`
	// A Wire Drawdown Payment Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_intention`.
	WireDrawdownPaymentIntention InterestPaymentSimulationResultTransactionSourceWireDrawdownPaymentIntention `json:"wire_drawdown_payment_intention,required,nullable"`
	// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_rejection`.
	WireDrawdownPaymentRejection InterestPaymentSimulationResultTransactionSourceWireDrawdownPaymentRejection `json:"wire_drawdown_payment_rejection,required,nullable"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention InterestPaymentSimulationResultTransactionSourceWireTransferIntention `json:"wire_transfer_intention,required,nullable"`
	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection InterestPaymentSimulationResultTransactionSourceWireTransferRejection `json:"wire_transfer_rejection,required,nullable"`
	JSON                  InterestPaymentSimulationResultTransactionSourceJSON
}

type InterestPaymentSimulationResultTransactionSourceJSON struct {
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
	CardRevenuePayment                          pjson.Metadata
	CheckDepositAcceptance                      pjson.Metadata
	CheckDepositReturn                          pjson.Metadata
	CheckTransferIntention                      pjson.Metadata
	CheckTransferReturn                         pjson.Metadata
	CheckTransferRejection                      pjson.Metadata
	CheckTransferStopPaymentRequest             pjson.Metadata
	DisputeResolution                           pjson.Metadata
	EmpyrealCashDeposit                         pjson.Metadata
	FeePayment                                  pjson.Metadata
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
	RealTimePaymentsTransferAcknowledgement     pjson.Metadata
	SampleFunds                                 pjson.Metadata
	WireDrawdownPaymentIntention                pjson.Metadata
	WireDrawdownPaymentRejection                pjson.Metadata
	WireTransferIntention                       pjson.Metadata
	WireTransferRejection                       pjson.Metadata
	Raw                                         []byte
	Extras                                      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSource using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *InterestPaymentSimulationResultTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceCategory string

const (
	InterestPaymentSimulationResultTransactionSourceCategoryAccountTransferIntention                    InterestPaymentSimulationResultTransactionSourceCategory = "account_transfer_intention"
	InterestPaymentSimulationResultTransactionSourceCategoryACHCheckConversionReturn                    InterestPaymentSimulationResultTransactionSourceCategory = "ach_check_conversion_return"
	InterestPaymentSimulationResultTransactionSourceCategoryACHCheckConversion                          InterestPaymentSimulationResultTransactionSourceCategory = "ach_check_conversion"
	InterestPaymentSimulationResultTransactionSourceCategoryACHTransferIntention                        InterestPaymentSimulationResultTransactionSourceCategory = "ach_transfer_intention"
	InterestPaymentSimulationResultTransactionSourceCategoryACHTransferRejection                        InterestPaymentSimulationResultTransactionSourceCategory = "ach_transfer_rejection"
	InterestPaymentSimulationResultTransactionSourceCategoryACHTransferReturn                           InterestPaymentSimulationResultTransactionSourceCategory = "ach_transfer_return"
	InterestPaymentSimulationResultTransactionSourceCategoryCardDisputeAcceptance                       InterestPaymentSimulationResultTransactionSourceCategory = "card_dispute_acceptance"
	InterestPaymentSimulationResultTransactionSourceCategoryCardRefund                                  InterestPaymentSimulationResultTransactionSourceCategory = "card_refund"
	InterestPaymentSimulationResultTransactionSourceCategoryCardSettlement                              InterestPaymentSimulationResultTransactionSourceCategory = "card_settlement"
	InterestPaymentSimulationResultTransactionSourceCategoryCardRevenuePayment                          InterestPaymentSimulationResultTransactionSourceCategory = "card_revenue_payment"
	InterestPaymentSimulationResultTransactionSourceCategoryCheckDepositAcceptance                      InterestPaymentSimulationResultTransactionSourceCategory = "check_deposit_acceptance"
	InterestPaymentSimulationResultTransactionSourceCategoryCheckDepositReturn                          InterestPaymentSimulationResultTransactionSourceCategory = "check_deposit_return"
	InterestPaymentSimulationResultTransactionSourceCategoryCheckTransferIntention                      InterestPaymentSimulationResultTransactionSourceCategory = "check_transfer_intention"
	InterestPaymentSimulationResultTransactionSourceCategoryCheckTransferReturn                         InterestPaymentSimulationResultTransactionSourceCategory = "check_transfer_return"
	InterestPaymentSimulationResultTransactionSourceCategoryCheckTransferRejection                      InterestPaymentSimulationResultTransactionSourceCategory = "check_transfer_rejection"
	InterestPaymentSimulationResultTransactionSourceCategoryCheckTransferStopPaymentRequest             InterestPaymentSimulationResultTransactionSourceCategory = "check_transfer_stop_payment_request"
	InterestPaymentSimulationResultTransactionSourceCategoryDisputeResolution                           InterestPaymentSimulationResultTransactionSourceCategory = "dispute_resolution"
	InterestPaymentSimulationResultTransactionSourceCategoryEmpyrealCashDeposit                         InterestPaymentSimulationResultTransactionSourceCategory = "empyreal_cash_deposit"
	InterestPaymentSimulationResultTransactionSourceCategoryFeePayment                                  InterestPaymentSimulationResultTransactionSourceCategory = "fee_payment"
	InterestPaymentSimulationResultTransactionSourceCategoryInboundACHTransfer                          InterestPaymentSimulationResultTransactionSourceCategory = "inbound_ach_transfer"
	InterestPaymentSimulationResultTransactionSourceCategoryInboundACHTransferReturnIntention           InterestPaymentSimulationResultTransactionSourceCategory = "inbound_ach_transfer_return_intention"
	InterestPaymentSimulationResultTransactionSourceCategoryInboundCheck                                InterestPaymentSimulationResultTransactionSourceCategory = "inbound_check"
	InterestPaymentSimulationResultTransactionSourceCategoryInboundInternationalACHTransfer             InterestPaymentSimulationResultTransactionSourceCategory = "inbound_international_ach_transfer"
	InterestPaymentSimulationResultTransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation InterestPaymentSimulationResultTransactionSourceCategory = "inbound_real_time_payments_transfer_confirmation"
	InterestPaymentSimulationResultTransactionSourceCategoryInboundWireDrawdownPaymentReversal          InterestPaymentSimulationResultTransactionSourceCategory = "inbound_wire_drawdown_payment_reversal"
	InterestPaymentSimulationResultTransactionSourceCategoryInboundWireDrawdownPayment                  InterestPaymentSimulationResultTransactionSourceCategory = "inbound_wire_drawdown_payment"
	InterestPaymentSimulationResultTransactionSourceCategoryInboundWireReversal                         InterestPaymentSimulationResultTransactionSourceCategory = "inbound_wire_reversal"
	InterestPaymentSimulationResultTransactionSourceCategoryInboundWireTransfer                         InterestPaymentSimulationResultTransactionSourceCategory = "inbound_wire_transfer"
	InterestPaymentSimulationResultTransactionSourceCategoryInterestPayment                             InterestPaymentSimulationResultTransactionSourceCategory = "interest_payment"
	InterestPaymentSimulationResultTransactionSourceCategoryInternalGeneralLedgerTransaction            InterestPaymentSimulationResultTransactionSourceCategory = "internal_general_ledger_transaction"
	InterestPaymentSimulationResultTransactionSourceCategoryInternalSource                              InterestPaymentSimulationResultTransactionSourceCategory = "internal_source"
	InterestPaymentSimulationResultTransactionSourceCategoryCardRouteRefund                             InterestPaymentSimulationResultTransactionSourceCategory = "card_route_refund"
	InterestPaymentSimulationResultTransactionSourceCategoryCardRouteSettlement                         InterestPaymentSimulationResultTransactionSourceCategory = "card_route_settlement"
	InterestPaymentSimulationResultTransactionSourceCategoryRealTimePaymentsTransferAcknowledgement     InterestPaymentSimulationResultTransactionSourceCategory = "real_time_payments_transfer_acknowledgement"
	InterestPaymentSimulationResultTransactionSourceCategorySampleFunds                                 InterestPaymentSimulationResultTransactionSourceCategory = "sample_funds"
	InterestPaymentSimulationResultTransactionSourceCategoryWireDrawdownPaymentIntention                InterestPaymentSimulationResultTransactionSourceCategory = "wire_drawdown_payment_intention"
	InterestPaymentSimulationResultTransactionSourceCategoryWireDrawdownPaymentRejection                InterestPaymentSimulationResultTransactionSourceCategory = "wire_drawdown_payment_rejection"
	InterestPaymentSimulationResultTransactionSourceCategoryWireTransferIntention                       InterestPaymentSimulationResultTransactionSourceCategory = "wire_transfer_intention"
	InterestPaymentSimulationResultTransactionSourceCategoryWireTransferRejection                       InterestPaymentSimulationResultTransactionSourceCategory = "wire_transfer_rejection"
	InterestPaymentSimulationResultTransactionSourceCategoryOther                                       InterestPaymentSimulationResultTransactionSourceCategory = "other"
)

type InterestPaymentSimulationResultTransactionSourceAccountTransferIntention struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrency `json:"currency,required"`
	// The description you chose to give the transfer.
	Description string `json:"description,required"`
	// The identifier of the Account to where the Account Transfer was sent.
	DestinationAccountID string `json:"destination_account_id,required"`
	// The identifier of the Account from where the Account Transfer was sent.
	SourceAccountID string `json:"source_account_id,required"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionJSON
}

type InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionJSON struct {
	Amount               pjson.Metadata
	Currency             pjson.Metadata
	Description          pjson.Metadata
	DestinationAccountID pjson.Metadata
	SourceAccountID      pjson.Metadata
	TransferID           pjson.Metadata
	Raw                  []byte
	Extras               map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceAccountTransferIntention using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceAccountTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrency string

const (
	InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrencyCad InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrency = "CAD"
	InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrencyChf InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrency = "CHF"
	InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrencyEur InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrency = "EUR"
	InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrencyGbp InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrency = "GBP"
	InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrencyJpy InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrency = "JPY"
	InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrencyUsd InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrency = "USD"
)

type InterestPaymentSimulationResultTransactionSourceACHCheckConversionReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// Why the transfer was returned.
	ReturnReasonCode string `json:"return_reason_code,required"`
	JSON             InterestPaymentSimulationResultTransactionSourceACHCheckConversionReturnJSON
}

type InterestPaymentSimulationResultTransactionSourceACHCheckConversionReturnJSON struct {
	Amount           pjson.Metadata
	ReturnReasonCode pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceACHCheckConversionReturn using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceACHCheckConversionReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceACHCheckConversion struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the File containing an image of the returned check.
	FileID string `json:"file_id,required"`
	JSON   InterestPaymentSimulationResultTransactionSourceACHCheckConversionJSON
}

type InterestPaymentSimulationResultTransactionSourceACHCheckConversionJSON struct {
	Amount pjson.Metadata
	FileID pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceACHCheckConversion using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceACHCheckConversion) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceACHTransferIntention struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount              int64  `json:"amount,required"`
	AccountNumber       string `json:"account_number,required"`
	RoutingNumber       string `json:"routing_number,required"`
	StatementDescriptor string `json:"statement_descriptor,required"`
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       InterestPaymentSimulationResultTransactionSourceACHTransferIntentionJSON
}

type InterestPaymentSimulationResultTransactionSourceACHTransferIntentionJSON struct {
	Amount              pjson.Metadata
	AccountNumber       pjson.Metadata
	RoutingNumber       pjson.Metadata
	StatementDescriptor pjson.Metadata
	TransferID          pjson.Metadata
	Raw                 []byte
	Extras              map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceACHTransferIntention using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceACHTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       InterestPaymentSimulationResultTransactionSourceACHTransferRejectionJSON
}

type InterestPaymentSimulationResultTransactionSourceACHTransferRejectionJSON struct {
	TransferID pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceACHTransferRejection using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceACHTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode `json:"return_reason_code,required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID string `json:"transfer_id,required"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID string `json:"transaction_id,required"`
	JSON          InterestPaymentSimulationResultTransactionSourceACHTransferReturnJSON
}

type InterestPaymentSimulationResultTransactionSourceACHTransferReturnJSON struct {
	CreatedAt        pjson.Metadata
	ReturnReasonCode pjson.Metadata
	TransferID       pjson.Metadata
	TransactionID    pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceACHTransferReturn using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode string

const (
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInsufficientFund                                          InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "insufficient_fund"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeNoAccount                                                 InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "no_account"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAccountClosed                                             InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "account_closed"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure                             InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction              InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver                              InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode   InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized                     InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodePaymentStopped                                            InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "payment_stopped"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeNonTransactionAccount                                     InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "non_transaction_account"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeUncollectedFunds                                          InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "uncollected_funds"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError                              InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAmountFieldError                                          InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "amount_field_error"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer                            InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber                                   InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeFileRecordEditCriteria                                    InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualName                                  InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest                                    InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAddendaError                                              InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "addenda_error"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeLimitedParticipationDfi                                   InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment              InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "incorrectly_coded_outbound_international_payment"
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeOther                                                     InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "other"
)

type InterestPaymentSimulationResultTransactionSourceCardDisputeAcceptance struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was accepted.
	AcceptedAt time.Time `json:"accepted_at,required" format:"date-time"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID string `json:"card_dispute_id,required"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID string `json:"transaction_id,required"`
	JSON          InterestPaymentSimulationResultTransactionSourceCardDisputeAcceptanceJSON
}

type InterestPaymentSimulationResultTransactionSourceCardDisputeAcceptanceJSON struct {
	AcceptedAt    pjson.Metadata
	CardDisputeID pjson.Metadata
	TransactionID pjson.Metadata
	Raw           []byte
	Extras        map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceCardDisputeAcceptance using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceCardDisputeAcceptance) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceCardRefund struct {
	// The Card Refund identifier.
	ID string `json:"id,required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency InterestPaymentSimulationResultTransactionSourceCardRefundCurrency `json:"currency,required"`
	// The identifier for the Transaction this refunds, if any.
	CardSettlementTransactionID string `json:"card_settlement_transaction_id,required,nullable"`
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
	Type InterestPaymentSimulationResultTransactionSourceCardRefundType `json:"type,required"`
	JSON InterestPaymentSimulationResultTransactionSourceCardRefundJSON
}

type InterestPaymentSimulationResultTransactionSourceCardRefundJSON struct {
	ID                          pjson.Metadata
	Amount                      pjson.Metadata
	Currency                    pjson.Metadata
	CardSettlementTransactionID pjson.Metadata
	MerchantAcceptorID          pjson.Metadata
	MerchantCity                pjson.Metadata
	MerchantState               pjson.Metadata
	MerchantCountry             pjson.Metadata
	MerchantName                pjson.Metadata
	MerchantCategoryCode        pjson.Metadata
	Type                        pjson.Metadata
	Raw                         []byte
	Extras                      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceCardRefund using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *InterestPaymentSimulationResultTransactionSourceCardRefund) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceCardRefundCurrency string

const (
	InterestPaymentSimulationResultTransactionSourceCardRefundCurrencyCad InterestPaymentSimulationResultTransactionSourceCardRefundCurrency = "CAD"
	InterestPaymentSimulationResultTransactionSourceCardRefundCurrencyChf InterestPaymentSimulationResultTransactionSourceCardRefundCurrency = "CHF"
	InterestPaymentSimulationResultTransactionSourceCardRefundCurrencyEur InterestPaymentSimulationResultTransactionSourceCardRefundCurrency = "EUR"
	InterestPaymentSimulationResultTransactionSourceCardRefundCurrencyGbp InterestPaymentSimulationResultTransactionSourceCardRefundCurrency = "GBP"
	InterestPaymentSimulationResultTransactionSourceCardRefundCurrencyJpy InterestPaymentSimulationResultTransactionSourceCardRefundCurrency = "JPY"
	InterestPaymentSimulationResultTransactionSourceCardRefundCurrencyUsd InterestPaymentSimulationResultTransactionSourceCardRefundCurrency = "USD"
)

type InterestPaymentSimulationResultTransactionSourceCardRefundType string

const (
	InterestPaymentSimulationResultTransactionSourceCardRefundTypeCardRefund InterestPaymentSimulationResultTransactionSourceCardRefundType = "card_refund"
)

type InterestPaymentSimulationResultTransactionSourceCardSettlement struct {
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
	Currency InterestPaymentSimulationResultTransactionSourceCardSettlementCurrency `json:"currency,required"`
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
	Type InterestPaymentSimulationResultTransactionSourceCardSettlementType `json:"type,required"`
	JSON InterestPaymentSimulationResultTransactionSourceCardSettlementJSON
}

type InterestPaymentSimulationResultTransactionSourceCardSettlementJSON struct {
	ID                   pjson.Metadata
	CardAuthorization    pjson.Metadata
	Amount               pjson.Metadata
	Currency             pjson.Metadata
	PresentmentAmount    pjson.Metadata
	PresentmentCurrency  pjson.Metadata
	MerchantAcceptorID   pjson.Metadata
	MerchantCity         pjson.Metadata
	MerchantState        pjson.Metadata
	MerchantCountry      pjson.Metadata
	MerchantName         pjson.Metadata
	MerchantCategoryCode pjson.Metadata
	PendingTransactionID pjson.Metadata
	Type                 pjson.Metadata
	Raw                  []byte
	Extras               map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceCardSettlement using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceCardSettlement) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceCardSettlementCurrency string

const (
	InterestPaymentSimulationResultTransactionSourceCardSettlementCurrencyCad InterestPaymentSimulationResultTransactionSourceCardSettlementCurrency = "CAD"
	InterestPaymentSimulationResultTransactionSourceCardSettlementCurrencyChf InterestPaymentSimulationResultTransactionSourceCardSettlementCurrency = "CHF"
	InterestPaymentSimulationResultTransactionSourceCardSettlementCurrencyEur InterestPaymentSimulationResultTransactionSourceCardSettlementCurrency = "EUR"
	InterestPaymentSimulationResultTransactionSourceCardSettlementCurrencyGbp InterestPaymentSimulationResultTransactionSourceCardSettlementCurrency = "GBP"
	InterestPaymentSimulationResultTransactionSourceCardSettlementCurrencyJpy InterestPaymentSimulationResultTransactionSourceCardSettlementCurrency = "JPY"
	InterestPaymentSimulationResultTransactionSourceCardSettlementCurrencyUsd InterestPaymentSimulationResultTransactionSourceCardSettlementCurrency = "USD"
)

type InterestPaymentSimulationResultTransactionSourceCardSettlementType string

const (
	InterestPaymentSimulationResultTransactionSourceCardSettlementTypeCardSettlement InterestPaymentSimulationResultTransactionSourceCardSettlementType = "card_settlement"
)

type InterestPaymentSimulationResultTransactionSourceCardRevenuePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrency `json:"currency,required"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time `json:"period_start,required" format:"date-time"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The account the card belonged to.
	TransactedOnAccountID string `json:"transacted_on_account_id,required,nullable"`
	JSON                  InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentJSON
}

type InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentJSON struct {
	Amount                pjson.Metadata
	Currency              pjson.Metadata
	PeriodStart           pjson.Metadata
	PeriodEnd             pjson.Metadata
	TransactedOnAccountID pjson.Metadata
	Raw                   []byte
	Extras                map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceCardRevenuePayment using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceCardRevenuePayment) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrency string

const (
	InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrencyCad InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrency = "CAD"
	InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrencyChf InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrency = "CHF"
	InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrencyEur InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrency = "EUR"
	InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrencyGbp InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrency = "GBP"
	InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrencyJpy InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrency = "JPY"
	InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrencyUsd InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrency = "USD"
)

type InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptance struct {
	// The amount to be deposited in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrency `json:"currency,required"`
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
	JSON           InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceJSON
}

type InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceJSON struct {
	Amount         pjson.Metadata
	Currency       pjson.Metadata
	AccountNumber  pjson.Metadata
	RoutingNumber  pjson.Metadata
	AuxiliaryOnUs  pjson.Metadata
	SerialNumber   pjson.Metadata
	CheckDepositID pjson.Metadata
	Raw            []byte
	Extras         map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptance using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptance) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrency string

const (
	InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyCad InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "CAD"
	InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyChf InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "CHF"
	InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyEur InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "EUR"
	InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyGbp InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "GBP"
	InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyJpy InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "JPY"
	InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyUsd InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "USD"
)

type InterestPaymentSimulationResultTransactionSourceCheckDepositReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrency `json:"currency,required"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID string `json:"check_deposit_id,required"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID string                                                                         `json:"transaction_id,required"`
	ReturnReason  InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason `json:"return_reason,required"`
	JSON          InterestPaymentSimulationResultTransactionSourceCheckDepositReturnJSON
}

type InterestPaymentSimulationResultTransactionSourceCheckDepositReturnJSON struct {
	Amount         pjson.Metadata
	ReturnedAt     pjson.Metadata
	Currency       pjson.Metadata
	CheckDepositID pjson.Metadata
	TransactionID  pjson.Metadata
	ReturnReason   pjson.Metadata
	Raw            []byte
	Extras         map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceCheckDepositReturn using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceCheckDepositReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrency string

const (
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrencyCad InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrency = "CAD"
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrencyChf InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrency = "CHF"
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrencyEur InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrency = "EUR"
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrencyGbp InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrency = "GBP"
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrencyJpy InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrency = "JPY"
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrencyUsd InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrency = "USD"
)

type InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason string

const (
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonACHConversionNotSupported InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "ach_conversion_not_supported"
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonClosedAccount             InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "closed_account"
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonDuplicateSubmission       InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "duplicate_submission"
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonInsufficientFunds         InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "insufficient_funds"
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonNoAccount                 InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "no_account"
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonNotAuthorized             InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "not_authorized"
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonStaleDated                InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "stale_dated"
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonStopPayment               InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "stop_payment"
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonUnknownReason             InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "unknown_reason"
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonUnmatchedDetails          InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "unmatched_details"
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonUnreadableImage           InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "unreadable_image"
)

type InterestPaymentSimulationResultTransactionSourceCheckTransferIntention struct {
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
	Currency InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrency `json:"currency,required"`
	// The name that will be printed on the check.
	RecipientName string `json:"recipient_name,required"`
	// The identifier of the Check Transfer with which this is associated.
	TransferID string `json:"transfer_id,required"`
	JSON       InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionJSON
}

type InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionJSON struct {
	AddressLine1  pjson.Metadata
	AddressLine2  pjson.Metadata
	AddressCity   pjson.Metadata
	AddressState  pjson.Metadata
	AddressZip    pjson.Metadata
	Amount        pjson.Metadata
	Currency      pjson.Metadata
	RecipientName pjson.Metadata
	TransferID    pjson.Metadata
	Raw           []byte
	Extras        map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceCheckTransferIntention using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceCheckTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrency string

const (
	InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrencyCad InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrency = "CAD"
	InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrencyChf InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrency = "CHF"
	InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrencyEur InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrency = "EUR"
	InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrencyGbp InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrency = "GBP"
	InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrencyJpy InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrency = "JPY"
	InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrencyUsd InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrency = "USD"
)

type InterestPaymentSimulationResultTransactionSourceCheckTransferReturn struct {
	// The identifier of the returned Check Transfer.
	TransferID string `json:"transfer_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// If available, a document with additional information about the return.
	FileID string `json:"file_id,required,nullable"`
	// The reason why the check was returned.
	Reason InterestPaymentSimulationResultTransactionSourceCheckTransferReturnReason `json:"reason,required"`
	// The identifier of the Transaction that was created to credit you for the
	// returned check.
	TransactionID string `json:"transaction_id,required,nullable"`
	JSON          InterestPaymentSimulationResultTransactionSourceCheckTransferReturnJSON
}

type InterestPaymentSimulationResultTransactionSourceCheckTransferReturnJSON struct {
	TransferID    pjson.Metadata
	ReturnedAt    pjson.Metadata
	FileID        pjson.Metadata
	Reason        pjson.Metadata
	TransactionID pjson.Metadata
	Raw           []byte
	Extras        map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceCheckTransferReturn using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceCheckTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceCheckTransferReturnReason string

const (
	InterestPaymentSimulationResultTransactionSourceCheckTransferReturnReasonMailDeliveryFailure InterestPaymentSimulationResultTransactionSourceCheckTransferReturnReason = "mail_delivery_failure"
	InterestPaymentSimulationResultTransactionSourceCheckTransferReturnReasonRefusedByRecipient  InterestPaymentSimulationResultTransactionSourceCheckTransferReturnReason = "refused_by_recipient"
)

type InterestPaymentSimulationResultTransactionSourceCheckTransferRejection struct {
	// The identifier of the Check Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       InterestPaymentSimulationResultTransactionSourceCheckTransferRejectionJSON
}

type InterestPaymentSimulationResultTransactionSourceCheckTransferRejectionJSON struct {
	TransferID pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceCheckTransferRejection using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceCheckTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferID string `json:"transfer_id,required"`
	// The transaction ID of the corresponding credit transaction.
	TransactionID string `json:"transaction_id,required"`
	// The time the stop-payment was requested.
	RequestedAt time.Time `json:"requested_at,required" format:"date-time"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestType `json:"type,required"`
	JSON InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestJSON
}

type InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestJSON struct {
	TransferID    pjson.Metadata
	TransactionID pjson.Metadata
	RequestedAt   pjson.Metadata
	Type          pjson.Metadata
	Raw           []byte
	Extras        map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequest
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestType string

const (
	InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

type InterestPaymentSimulationResultTransactionSourceDisputeResolution struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency InterestPaymentSimulationResultTransactionSourceDisputeResolutionCurrency `json:"currency,required"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionID string `json:"disputed_transaction_id,required"`
	JSON                  InterestPaymentSimulationResultTransactionSourceDisputeResolutionJSON
}

type InterestPaymentSimulationResultTransactionSourceDisputeResolutionJSON struct {
	Amount                pjson.Metadata
	Currency              pjson.Metadata
	DisputedTransactionID pjson.Metadata
	Raw                   []byte
	Extras                map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceDisputeResolution using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceDisputeResolution) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceDisputeResolutionCurrency string

const (
	InterestPaymentSimulationResultTransactionSourceDisputeResolutionCurrencyCad InterestPaymentSimulationResultTransactionSourceDisputeResolutionCurrency = "CAD"
	InterestPaymentSimulationResultTransactionSourceDisputeResolutionCurrencyChf InterestPaymentSimulationResultTransactionSourceDisputeResolutionCurrency = "CHF"
	InterestPaymentSimulationResultTransactionSourceDisputeResolutionCurrencyEur InterestPaymentSimulationResultTransactionSourceDisputeResolutionCurrency = "EUR"
	InterestPaymentSimulationResultTransactionSourceDisputeResolutionCurrencyGbp InterestPaymentSimulationResultTransactionSourceDisputeResolutionCurrency = "GBP"
	InterestPaymentSimulationResultTransactionSourceDisputeResolutionCurrencyJpy InterestPaymentSimulationResultTransactionSourceDisputeResolutionCurrency = "JPY"
	InterestPaymentSimulationResultTransactionSourceDisputeResolutionCurrencyUsd InterestPaymentSimulationResultTransactionSourceDisputeResolutionCurrency = "USD"
)

type InterestPaymentSimulationResultTransactionSourceEmpyrealCashDeposit struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount      int64     `json:"amount,required"`
	BagID       string    `json:"bag_id,required"`
	DepositDate time.Time `json:"deposit_date,required" format:"date-time"`
	JSON        InterestPaymentSimulationResultTransactionSourceEmpyrealCashDepositJSON
}

type InterestPaymentSimulationResultTransactionSourceEmpyrealCashDepositJSON struct {
	Amount      pjson.Metadata
	BagID       pjson.Metadata
	DepositDate pjson.Metadata
	Raw         []byte
	Extras      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceEmpyrealCashDeposit using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceEmpyrealCashDeposit) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceFeePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency InterestPaymentSimulationResultTransactionSourceFeePaymentCurrency `json:"currency,required"`
	JSON     InterestPaymentSimulationResultTransactionSourceFeePaymentJSON
}

type InterestPaymentSimulationResultTransactionSourceFeePaymentJSON struct {
	Amount   pjson.Metadata
	Currency pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceFeePayment using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *InterestPaymentSimulationResultTransactionSourceFeePayment) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceFeePaymentCurrency string

const (
	InterestPaymentSimulationResultTransactionSourceFeePaymentCurrencyCad InterestPaymentSimulationResultTransactionSourceFeePaymentCurrency = "CAD"
	InterestPaymentSimulationResultTransactionSourceFeePaymentCurrencyChf InterestPaymentSimulationResultTransactionSourceFeePaymentCurrency = "CHF"
	InterestPaymentSimulationResultTransactionSourceFeePaymentCurrencyEur InterestPaymentSimulationResultTransactionSourceFeePaymentCurrency = "EUR"
	InterestPaymentSimulationResultTransactionSourceFeePaymentCurrencyGbp InterestPaymentSimulationResultTransactionSourceFeePaymentCurrency = "GBP"
	InterestPaymentSimulationResultTransactionSourceFeePaymentCurrencyJpy InterestPaymentSimulationResultTransactionSourceFeePaymentCurrency = "JPY"
	InterestPaymentSimulationResultTransactionSourceFeePaymentCurrencyUsd InterestPaymentSimulationResultTransactionSourceFeePaymentCurrency = "USD"
)

type InterestPaymentSimulationResultTransactionSourceInboundACHTransfer struct {
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
	JSON                               InterestPaymentSimulationResultTransactionSourceInboundACHTransferJSON
}

type InterestPaymentSimulationResultTransactionSourceInboundACHTransferJSON struct {
	Amount                             pjson.Metadata
	OriginatorCompanyName              pjson.Metadata
	OriginatorCompanyDescriptiveDate   pjson.Metadata
	OriginatorCompanyDiscretionaryData pjson.Metadata
	OriginatorCompanyEntryDescription  pjson.Metadata
	OriginatorCompanyID                pjson.Metadata
	ReceiverIDNumber                   pjson.Metadata
	ReceiverName                       pjson.Metadata
	TraceNumber                        pjson.Metadata
	Raw                                []byte
	Extras                             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceInboundACHTransfer using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceInboundACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceInboundCheck struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency              InterestPaymentSimulationResultTransactionSourceInboundCheckCurrency `json:"currency,required"`
	CheckNumber           string                                                               `json:"check_number,required,nullable"`
	CheckFrontImageFileID string                                                               `json:"check_front_image_file_id,required,nullable"`
	CheckRearImageFileID  string                                                               `json:"check_rear_image_file_id,required,nullable"`
	JSON                  InterestPaymentSimulationResultTransactionSourceInboundCheckJSON
}

type InterestPaymentSimulationResultTransactionSourceInboundCheckJSON struct {
	Amount                pjson.Metadata
	Currency              pjson.Metadata
	CheckNumber           pjson.Metadata
	CheckFrontImageFileID pjson.Metadata
	CheckRearImageFileID  pjson.Metadata
	Raw                   []byte
	Extras                map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceInboundCheck using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *InterestPaymentSimulationResultTransactionSourceInboundCheck) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceInboundCheckCurrency string

const (
	InterestPaymentSimulationResultTransactionSourceInboundCheckCurrencyCad InterestPaymentSimulationResultTransactionSourceInboundCheckCurrency = "CAD"
	InterestPaymentSimulationResultTransactionSourceInboundCheckCurrencyChf InterestPaymentSimulationResultTransactionSourceInboundCheckCurrency = "CHF"
	InterestPaymentSimulationResultTransactionSourceInboundCheckCurrencyEur InterestPaymentSimulationResultTransactionSourceInboundCheckCurrency = "EUR"
	InterestPaymentSimulationResultTransactionSourceInboundCheckCurrencyGbp InterestPaymentSimulationResultTransactionSourceInboundCheckCurrency = "GBP"
	InterestPaymentSimulationResultTransactionSourceInboundCheckCurrencyJpy InterestPaymentSimulationResultTransactionSourceInboundCheckCurrency = "JPY"
	InterestPaymentSimulationResultTransactionSourceInboundCheckCurrencyUsd InterestPaymentSimulationResultTransactionSourceInboundCheckCurrency = "USD"
)

type InterestPaymentSimulationResultTransactionSourceInboundInternationalACHTransfer struct {
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
	JSON                                                   InterestPaymentSimulationResultTransactionSourceInboundInternationalACHTransferJSON
}

type InterestPaymentSimulationResultTransactionSourceInboundInternationalACHTransferJSON struct {
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
	Raw                                                    []byte
	Extras                                                 map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceInboundInternationalACHTransfer
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InterestPaymentSimulationResultTransactionSourceInboundInternationalACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real Time Payments transfer.
	Currency InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency `json:"currency,required"`
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
	JSON                  InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON
}

type InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON struct {
	Amount                    pjson.Metadata
	Currency                  pjson.Metadata
	CreditorName              pjson.Metadata
	DebtorName                pjson.Metadata
	DebtorAccountNumber       pjson.Metadata
	DebtorRoutingNumber       pjson.Metadata
	TransactionIdentification pjson.Metadata
	RemittanceInformation     pjson.Metadata
	Raw                       []byte
	Extras                    map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency string

const (
	InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyCad InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CAD"
	InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyChf InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CHF"
	InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyEur InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "EUR"
	InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyGbp InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "GBP"
	InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyJpy InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "JPY"
	InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyUsd InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "USD"
)

type InterestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal struct {
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
	JSON                       InterestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentReversalJSON
}

type InterestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentReversalJSON struct {
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
	Raw                                           []byte
	Extras                                        map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InterestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceInboundWireDrawdownPayment struct {
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
	JSON                               InterestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentJSON
}

type InterestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentJSON struct {
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
	Raw                                []byte
	Extras                             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceInboundWireDrawdownPayment using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceInboundWireDrawdownPayment) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceInboundWireReversal struct {
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
	JSON          InterestPaymentSimulationResultTransactionSourceInboundWireReversalJSON
}

type InterestPaymentSimulationResultTransactionSourceInboundWireReversalJSON struct {
	Amount                                                pjson.Metadata
	CreatedAt                                             pjson.Metadata
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
	TransactionID                                         pjson.Metadata
	Raw                                                   []byte
	Extras                                                map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceInboundWireReversal using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceInboundWireReversal) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceInboundWireTransfer struct {
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
	JSON                                    InterestPaymentSimulationResultTransactionSourceInboundWireTransferJSON
}

type InterestPaymentSimulationResultTransactionSourceInboundWireTransferJSON struct {
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
	Raw                                     []byte
	Extras                                  map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceInboundWireTransfer using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceInboundWireTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceInterestPayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrency `json:"currency,required"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time `json:"period_start,required" format:"date-time"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The account on which the interest was accrued.
	AccruedOnAccountID string `json:"accrued_on_account_id,required,nullable"`
	JSON               InterestPaymentSimulationResultTransactionSourceInterestPaymentJSON
}

type InterestPaymentSimulationResultTransactionSourceInterestPaymentJSON struct {
	Amount             pjson.Metadata
	Currency           pjson.Metadata
	PeriodStart        pjson.Metadata
	PeriodEnd          pjson.Metadata
	AccruedOnAccountID pjson.Metadata
	Raw                []byte
	Extras             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceInterestPayment using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceInterestPayment) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrency string

const (
	InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrencyCad InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrency = "CAD"
	InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrencyChf InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrency = "CHF"
	InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrencyEur InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrency = "EUR"
	InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrencyGbp InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrency = "GBP"
	InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrencyJpy InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrency = "JPY"
	InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrencyUsd InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrency = "USD"
)

type InterestPaymentSimulationResultTransactionSourceInternalSource struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency InterestPaymentSimulationResultTransactionSourceInternalSourceCurrency `json:"currency,required"`
	Reason   InterestPaymentSimulationResultTransactionSourceInternalSourceReason   `json:"reason,required"`
	JSON     InterestPaymentSimulationResultTransactionSourceInternalSourceJSON
}

type InterestPaymentSimulationResultTransactionSourceInternalSourceJSON struct {
	Amount   pjson.Metadata
	Currency pjson.Metadata
	Reason   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceInternalSource using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceInternalSource) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceInternalSourceCurrency string

const (
	InterestPaymentSimulationResultTransactionSourceInternalSourceCurrencyCad InterestPaymentSimulationResultTransactionSourceInternalSourceCurrency = "CAD"
	InterestPaymentSimulationResultTransactionSourceInternalSourceCurrencyChf InterestPaymentSimulationResultTransactionSourceInternalSourceCurrency = "CHF"
	InterestPaymentSimulationResultTransactionSourceInternalSourceCurrencyEur InterestPaymentSimulationResultTransactionSourceInternalSourceCurrency = "EUR"
	InterestPaymentSimulationResultTransactionSourceInternalSourceCurrencyGbp InterestPaymentSimulationResultTransactionSourceInternalSourceCurrency = "GBP"
	InterestPaymentSimulationResultTransactionSourceInternalSourceCurrencyJpy InterestPaymentSimulationResultTransactionSourceInternalSourceCurrency = "JPY"
	InterestPaymentSimulationResultTransactionSourceInternalSourceCurrencyUsd InterestPaymentSimulationResultTransactionSourceInternalSourceCurrency = "USD"
)

type InterestPaymentSimulationResultTransactionSourceInternalSourceReason string

const (
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonBankMigration      InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "bank_migration"
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonCashback           InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "cashback"
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonEmpyrealAdjustment InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "empyreal_adjustment"
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonError              InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "error"
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonErrorCorrection    InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "error_correction"
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonFees               InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "fees"
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonInterest           InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "interest"
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonSampleFunds        InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "sample_funds"
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonSampleFundsReturn  InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "sample_funds_return"
)

type InterestPaymentSimulationResultTransactionSourceCardRouteRefund struct {
	// The refunded amount in the minor unit of the refunded currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
	// currency.
	Currency             InterestPaymentSimulationResultTransactionSourceCardRouteRefundCurrency `json:"currency,required"`
	MerchantAcceptorID   string                                                                  `json:"merchant_acceptor_id,required"`
	MerchantCity         string                                                                  `json:"merchant_city,required,nullable"`
	MerchantCountry      string                                                                  `json:"merchant_country,required"`
	MerchantDescriptor   string                                                                  `json:"merchant_descriptor,required"`
	MerchantState        string                                                                  `json:"merchant_state,required,nullable"`
	MerchantCategoryCode string                                                                  `json:"merchant_category_code,required,nullable"`
	JSON                 InterestPaymentSimulationResultTransactionSourceCardRouteRefundJSON
}

type InterestPaymentSimulationResultTransactionSourceCardRouteRefundJSON struct {
	Amount               pjson.Metadata
	Currency             pjson.Metadata
	MerchantAcceptorID   pjson.Metadata
	MerchantCity         pjson.Metadata
	MerchantCountry      pjson.Metadata
	MerchantDescriptor   pjson.Metadata
	MerchantState        pjson.Metadata
	MerchantCategoryCode pjson.Metadata
	Raw                  []byte
	Extras               map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceCardRouteRefund using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceCardRouteRefund) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceCardRouteRefundCurrency string

const (
	InterestPaymentSimulationResultTransactionSourceCardRouteRefundCurrencyCad InterestPaymentSimulationResultTransactionSourceCardRouteRefundCurrency = "CAD"
	InterestPaymentSimulationResultTransactionSourceCardRouteRefundCurrencyChf InterestPaymentSimulationResultTransactionSourceCardRouteRefundCurrency = "CHF"
	InterestPaymentSimulationResultTransactionSourceCardRouteRefundCurrencyEur InterestPaymentSimulationResultTransactionSourceCardRouteRefundCurrency = "EUR"
	InterestPaymentSimulationResultTransactionSourceCardRouteRefundCurrencyGbp InterestPaymentSimulationResultTransactionSourceCardRouteRefundCurrency = "GBP"
	InterestPaymentSimulationResultTransactionSourceCardRouteRefundCurrencyJpy InterestPaymentSimulationResultTransactionSourceCardRouteRefundCurrency = "JPY"
	InterestPaymentSimulationResultTransactionSourceCardRouteRefundCurrencyUsd InterestPaymentSimulationResultTransactionSourceCardRouteRefundCurrency = "USD"
)

type InterestPaymentSimulationResultTransactionSourceCardRouteSettlement struct {
	// The settled amount in the minor unit of the settlement currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
	// currency.
	Currency             InterestPaymentSimulationResultTransactionSourceCardRouteSettlementCurrency `json:"currency,required"`
	MerchantAcceptorID   string                                                                      `json:"merchant_acceptor_id,required"`
	MerchantCity         string                                                                      `json:"merchant_city,required,nullable"`
	MerchantCountry      string                                                                      `json:"merchant_country,required,nullable"`
	MerchantDescriptor   string                                                                      `json:"merchant_descriptor,required"`
	MerchantState        string                                                                      `json:"merchant_state,required,nullable"`
	MerchantCategoryCode string                                                                      `json:"merchant_category_code,required,nullable"`
	JSON                 InterestPaymentSimulationResultTransactionSourceCardRouteSettlementJSON
}

type InterestPaymentSimulationResultTransactionSourceCardRouteSettlementJSON struct {
	Amount               pjson.Metadata
	Currency             pjson.Metadata
	MerchantAcceptorID   pjson.Metadata
	MerchantCity         pjson.Metadata
	MerchantCountry      pjson.Metadata
	MerchantDescriptor   pjson.Metadata
	MerchantState        pjson.Metadata
	MerchantCategoryCode pjson.Metadata
	Raw                  []byte
	Extras               map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceCardRouteSettlement using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceCardRouteSettlement) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceCardRouteSettlementCurrency string

const (
	InterestPaymentSimulationResultTransactionSourceCardRouteSettlementCurrencyCad InterestPaymentSimulationResultTransactionSourceCardRouteSettlementCurrency = "CAD"
	InterestPaymentSimulationResultTransactionSourceCardRouteSettlementCurrencyChf InterestPaymentSimulationResultTransactionSourceCardRouteSettlementCurrency = "CHF"
	InterestPaymentSimulationResultTransactionSourceCardRouteSettlementCurrencyEur InterestPaymentSimulationResultTransactionSourceCardRouteSettlementCurrency = "EUR"
	InterestPaymentSimulationResultTransactionSourceCardRouteSettlementCurrencyGbp InterestPaymentSimulationResultTransactionSourceCardRouteSettlementCurrency = "GBP"
	InterestPaymentSimulationResultTransactionSourceCardRouteSettlementCurrencyJpy InterestPaymentSimulationResultTransactionSourceCardRouteSettlementCurrency = "JPY"
	InterestPaymentSimulationResultTransactionSourceCardRouteSettlementCurrencyUsd InterestPaymentSimulationResultTransactionSourceCardRouteSettlementCurrency = "USD"
)

type InterestPaymentSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgement struct {
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
	JSON       InterestPaymentSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgementJSON
}

type InterestPaymentSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgementJSON struct {
	Amount                   pjson.Metadata
	DestinationAccountNumber pjson.Metadata
	DestinationRoutingNumber pjson.Metadata
	RemittanceInformation    pjson.Metadata
	TransferID               pjson.Metadata
	Raw                      []byte
	Extras                   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgement
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InterestPaymentSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgement) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceSampleFunds struct {
	// Where the sample funds came from.
	Originator string `json:"originator,required"`
	JSON       InterestPaymentSimulationResultTransactionSourceSampleFundsJSON
}

type InterestPaymentSimulationResultTransactionSourceSampleFundsJSON struct {
	Originator pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceSampleFunds using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *InterestPaymentSimulationResultTransactionSourceSampleFunds) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceWireDrawdownPaymentIntention struct {
	// The transfer amount in USD cents.
	Amount             int64  `json:"amount,required"`
	AccountNumber      string `json:"account_number,required"`
	RoutingNumber      string `json:"routing_number,required"`
	MessageToRecipient string `json:"message_to_recipient,required"`
	TransferID         string `json:"transfer_id,required"`
	JSON               InterestPaymentSimulationResultTransactionSourceWireDrawdownPaymentIntentionJSON
}

type InterestPaymentSimulationResultTransactionSourceWireDrawdownPaymentIntentionJSON struct {
	Amount             pjson.Metadata
	AccountNumber      pjson.Metadata
	RoutingNumber      pjson.Metadata
	MessageToRecipient pjson.Metadata
	TransferID         pjson.Metadata
	Raw                []byte
	Extras             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceWireDrawdownPaymentIntention
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InterestPaymentSimulationResultTransactionSourceWireDrawdownPaymentIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceWireDrawdownPaymentRejection struct {
	TransferID string `json:"transfer_id,required"`
	JSON       InterestPaymentSimulationResultTransactionSourceWireDrawdownPaymentRejectionJSON
}

type InterestPaymentSimulationResultTransactionSourceWireDrawdownPaymentRejectionJSON struct {
	TransferID pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceWireDrawdownPaymentRejection
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InterestPaymentSimulationResultTransactionSourceWireDrawdownPaymentRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceWireTransferIntention struct {
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient string `json:"message_to_recipient,required"`
	TransferID         string `json:"transfer_id,required"`
	JSON               InterestPaymentSimulationResultTransactionSourceWireTransferIntentionJSON
}

type InterestPaymentSimulationResultTransactionSourceWireTransferIntentionJSON struct {
	Amount             pjson.Metadata
	AccountNumber      pjson.Metadata
	RoutingNumber      pjson.Metadata
	MessageToRecipient pjson.Metadata
	TransferID         pjson.Metadata
	Raw                []byte
	Extras             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceWireTransferIntention using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceWireTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceWireTransferRejection struct {
	TransferID string `json:"transfer_id,required"`
	JSON       InterestPaymentSimulationResultTransactionSourceWireTransferRejectionJSON
}

type InterestPaymentSimulationResultTransactionSourceWireTransferRejectionJSON struct {
	TransferID pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InterestPaymentSimulationResultTransactionSourceWireTransferRejection using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InterestPaymentSimulationResultTransactionSourceWireTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionType string

const (
	InterestPaymentSimulationResultTransactionTypeTransaction InterestPaymentSimulationResultTransactionType = "transaction"
)

type InterestPaymentSimulationResultType string

const (
	InterestPaymentSimulationResultTypeInterestPaymentSimulationResult InterestPaymentSimulationResultType = "interest_payment_simulation_result"
)
