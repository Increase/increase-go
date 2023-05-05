package increase

import (
	"context"
	"net/http"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// SimulationInterestPaymentService contains methods and other services that help
// with interacting with the increase API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSimulationInterestPaymentService] method instead.
type SimulationInterestPaymentService struct {
	Options []option.RequestOption
}

// NewSimulationInterestPaymentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationInterestPaymentService(opts ...option.RequestOption) (r *SimulationInterestPaymentService) {
	r = &SimulationInterestPaymentService{}
	r.Options = opts
	return
}

// Simulates an interest payment to your account. In production, this happens
// automatically on the first of each month.
func (r *SimulationInterestPaymentService) New(ctx context.Context, body SimulationInterestPaymentNewParams, opts ...option.RequestOption) (res *InterestPaymentSimulationResult, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/interest_payment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The results of an Interest Payment simulation.
type InterestPaymentSimulationResult struct {
	// This will contain the resulting [Transaction](#transactions) object. The
	// Transaction's `source` will be of `category: interest_payment`.
	Transaction InterestPaymentSimulationResultTransaction `json:"transaction,required"`
	// A constant representing the object's type. For this resource it will always be
	// `interest_payment_simulation_result`.
	Type InterestPaymentSimulationResultType `json:"type,required"`
	JSON interestPaymentSimulationResultJSON
}

// interestPaymentSimulationResultJSON contains the JSON metadata for the struct
// [InterestPaymentSimulationResult]
type interestPaymentSimulationResultJSON struct {
	Transaction apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterestPaymentSimulationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// This will contain the resulting [Transaction](#transactions) object. The
// Transaction's `source` will be of `category: interest_payment`.
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
	JSON interestPaymentSimulationResultTransactionJSON
}

// interestPaymentSimulationResultTransactionJSON contains the JSON metadata for
// the struct [InterestPaymentSimulationResultTransaction]
type interestPaymentSimulationResultTransactionJSON struct {
	AccountID   apijson.Field
	Amount      apijson.Field
	Currency    apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	ID          apijson.Field
	RouteID     apijson.Field
	RouteType   apijson.Field
	Source      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// This is an object giving more details on the network-level event that caused the
// Transaction. Note that for backwards compatibility reasons, additional
// undocumented keys may appear in this object. These should be treated as
// deprecated and will be removed in the future.
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
	JSON                  interestPaymentSimulationResultTransactionSourceJSON
}

// interestPaymentSimulationResultTransactionSourceJSON contains the JSON metadata
// for the struct [InterestPaymentSimulationResultTransactionSource]
type interestPaymentSimulationResultTransactionSourceJSON struct {
	Category                                    apijson.Field
	AccountTransferIntention                    apijson.Field
	ACHCheckConversionReturn                    apijson.Field
	ACHCheckConversion                          apijson.Field
	ACHTransferIntention                        apijson.Field
	ACHTransferRejection                        apijson.Field
	ACHTransferReturn                           apijson.Field
	CardDisputeAcceptance                       apijson.Field
	CardRefund                                  apijson.Field
	CardSettlement                              apijson.Field
	CardRevenuePayment                          apijson.Field
	CheckDepositAcceptance                      apijson.Field
	CheckDepositReturn                          apijson.Field
	CheckTransferIntention                      apijson.Field
	CheckTransferReturn                         apijson.Field
	CheckTransferRejection                      apijson.Field
	CheckTransferStopPaymentRequest             apijson.Field
	DisputeResolution                           apijson.Field
	EmpyrealCashDeposit                         apijson.Field
	FeePayment                                  apijson.Field
	InboundACHTransfer                          apijson.Field
	InboundCheck                                apijson.Field
	InboundInternationalACHTransfer             apijson.Field
	InboundRealTimePaymentsTransferConfirmation apijson.Field
	InboundWireDrawdownPaymentReversal          apijson.Field
	InboundWireDrawdownPayment                  apijson.Field
	InboundWireReversal                         apijson.Field
	InboundWireTransfer                         apijson.Field
	InterestPayment                             apijson.Field
	InternalSource                              apijson.Field
	CardRouteRefund                             apijson.Field
	CardRouteSettlement                         apijson.Field
	RealTimePaymentsTransferAcknowledgement     apijson.Field
	SampleFunds                                 apijson.Field
	WireDrawdownPaymentIntention                apijson.Field
	WireDrawdownPaymentRejection                apijson.Field
	WireTransferIntention                       apijson.Field
	WireTransferRejection                       apijson.Field
	raw                                         string
	ExtraFields                                 map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	InterestPaymentSimulationResultTransactionSourceCategoryCollectionReceivable                        InterestPaymentSimulationResultTransactionSourceCategory = "collection_receivable"
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

// A Account Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_intention`.
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
	JSON       interestPaymentSimulationResultTransactionSourceAccountTransferIntentionJSON
}

// interestPaymentSimulationResultTransactionSourceAccountTransferIntentionJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceAccountTransferIntention]
type interestPaymentSimulationResultTransactionSourceAccountTransferIntentionJSON struct {
	Amount               apijson.Field
	Currency             apijson.Field
	Description          apijson.Field
	DestinationAccountID apijson.Field
	SourceAccountID      apijson.Field
	TransferID           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceAccountTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// A ACH Check Conversion Return object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_check_conversion_return`.
type InterestPaymentSimulationResultTransactionSourceACHCheckConversionReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// Why the transfer was returned.
	ReturnReasonCode string `json:"return_reason_code,required"`
	JSON             interestPaymentSimulationResultTransactionSourceACHCheckConversionReturnJSON
}

// interestPaymentSimulationResultTransactionSourceACHCheckConversionReturnJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceACHCheckConversionReturn]
type interestPaymentSimulationResultTransactionSourceACHCheckConversionReturnJSON struct {
	Amount           apijson.Field
	ReturnReasonCode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceACHCheckConversionReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ACH Check Conversion object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_check_conversion`.
type InterestPaymentSimulationResultTransactionSourceACHCheckConversion struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the File containing an image of the returned check.
	FileID string `json:"file_id,required"`
	JSON   interestPaymentSimulationResultTransactionSourceACHCheckConversionJSON
}

// interestPaymentSimulationResultTransactionSourceACHCheckConversionJSON contains
// the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceACHCheckConversion]
type interestPaymentSimulationResultTransactionSourceACHCheckConversionJSON struct {
	Amount      apijson.Field
	FileID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceACHCheckConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ACH Transfer Intention object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_intention`.
type InterestPaymentSimulationResultTransactionSourceACHTransferIntention struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount              int64  `json:"amount,required"`
	AccountNumber       string `json:"account_number,required"`
	RoutingNumber       string `json:"routing_number,required"`
	StatementDescriptor string `json:"statement_descriptor,required"`
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       interestPaymentSimulationResultTransactionSourceACHTransferIntentionJSON
}

// interestPaymentSimulationResultTransactionSourceACHTransferIntentionJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceACHTransferIntention]
type interestPaymentSimulationResultTransactionSourceACHTransferIntentionJSON struct {
	Amount              apijson.Field
	AccountNumber       apijson.Field
	RoutingNumber       apijson.Field
	StatementDescriptor apijson.Field
	TransferID          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceACHTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ACH Transfer Rejection object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_rejection`.
type InterestPaymentSimulationResultTransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       interestPaymentSimulationResultTransactionSourceACHTransferRejectionJSON
}

// interestPaymentSimulationResultTransactionSourceACHTransferRejectionJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceACHTransferRejection]
type interestPaymentSimulationResultTransactionSourceACHTransferRejectionJSON struct {
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceACHTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ACH Transfer Return object. This field will be present in the JSON response if
// and only if `category` is equal to `ach_transfer_return`.
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
	JSON          interestPaymentSimulationResultTransactionSourceACHTransferReturnJSON
}

// interestPaymentSimulationResultTransactionSourceACHTransferReturnJSON contains
// the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceACHTransferReturn]
type interestPaymentSimulationResultTransactionSourceACHTransferReturnJSON struct {
	CreatedAt        apijson.Field
	ReturnReasonCode apijson.Field
	TransferID       apijson.Field
	TransactionID    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// A Card Dispute Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `card_dispute_acceptance`.
type InterestPaymentSimulationResultTransactionSourceCardDisputeAcceptance struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was accepted.
	AcceptedAt time.Time `json:"accepted_at,required" format:"date-time"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID string `json:"card_dispute_id,required"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID string `json:"transaction_id,required"`
	JSON          interestPaymentSimulationResultTransactionSourceCardDisputeAcceptanceJSON
}

// interestPaymentSimulationResultTransactionSourceCardDisputeAcceptanceJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceCardDisputeAcceptance]
type interestPaymentSimulationResultTransactionSourceCardDisputeAcceptanceJSON struct {
	AcceptedAt    apijson.Field
	CardDisputeID apijson.Field
	TransactionID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceCardDisputeAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Card Refund object. This field will be present in the JSON response if and
// only if `category` is equal to `card_refund`.
type InterestPaymentSimulationResultTransactionSourceCardRefund struct {
	// The Card Refund identifier.
	ID string `json:"id,required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency InterestPaymentSimulationResultTransactionSourceCardRefundCurrency `json:"currency,required"`
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
	JSON interestPaymentSimulationResultTransactionSourceCardRefundJSON
}

// interestPaymentSimulationResultTransactionSourceCardRefundJSON contains the JSON
// metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceCardRefund]
type interestPaymentSimulationResultTransactionSourceCardRefundJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	Currency             apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCity         apijson.Field
	MerchantState        apijson.Field
	MerchantCountry      apijson.Field
	MerchantName         apijson.Field
	MerchantCategoryCode apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceCardRefund) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// A Card Settlement object. This field will be present in the JSON response if and
// only if `category` is equal to `card_settlement`.
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
	JSON interestPaymentSimulationResultTransactionSourceCardSettlementJSON
}

// interestPaymentSimulationResultTransactionSourceCardSettlementJSON contains the
// JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceCardSettlement]
type interestPaymentSimulationResultTransactionSourceCardSettlementJSON struct {
	ID                   apijson.Field
	CardAuthorization    apijson.Field
	Amount               apijson.Field
	Currency             apijson.Field
	PresentmentAmount    apijson.Field
	PresentmentCurrency  apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCity         apijson.Field
	MerchantState        apijson.Field
	MerchantCountry      apijson.Field
	MerchantName         apijson.Field
	MerchantCategoryCode apijson.Field
	PendingTransactionID apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceCardSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// A Card Revenue Payment object. This field will be present in the JSON response
// if and only if `category` is equal to `card_revenue_payment`.
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
	JSON                  interestPaymentSimulationResultTransactionSourceCardRevenuePaymentJSON
}

// interestPaymentSimulationResultTransactionSourceCardRevenuePaymentJSON contains
// the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceCardRevenuePayment]
type interestPaymentSimulationResultTransactionSourceCardRevenuePaymentJSON struct {
	Amount                apijson.Field
	Currency              apijson.Field
	PeriodStart           apijson.Field
	PeriodEnd             apijson.Field
	TransactedOnAccountID apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceCardRevenuePayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// A Check Deposit Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_acceptance`.
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
	JSON           interestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceJSON
}

// interestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptance]
type interestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceJSON struct {
	Amount         apijson.Field
	Currency       apijson.Field
	AccountNumber  apijson.Field
	RoutingNumber  apijson.Field
	AuxiliaryOnUs  apijson.Field
	SerialNumber   apijson.Field
	CheckDepositID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// A Check Deposit Return object. This field will be present in the JSON response
// if and only if `category` is equal to `check_deposit_return`.
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
	JSON          interestPaymentSimulationResultTransactionSourceCheckDepositReturnJSON
}

// interestPaymentSimulationResultTransactionSourceCheckDepositReturnJSON contains
// the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceCheckDepositReturn]
type interestPaymentSimulationResultTransactionSourceCheckDepositReturnJSON struct {
	Amount         apijson.Field
	ReturnedAt     apijson.Field
	Currency       apijson.Field
	CheckDepositID apijson.Field
	TransactionID  apijson.Field
	ReturnReason   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceCheckDepositReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// A Check Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_intention`.
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
	JSON       interestPaymentSimulationResultTransactionSourceCheckTransferIntentionJSON
}

// interestPaymentSimulationResultTransactionSourceCheckTransferIntentionJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceCheckTransferIntention]
type interestPaymentSimulationResultTransactionSourceCheckTransferIntentionJSON struct {
	AddressLine1  apijson.Field
	AddressLine2  apijson.Field
	AddressCity   apijson.Field
	AddressState  apijson.Field
	AddressZip    apijson.Field
	Amount        apijson.Field
	Currency      apijson.Field
	RecipientName apijson.Field
	TransferID    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceCheckTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// A Check Transfer Return object. This field will be present in the JSON response
// if and only if `category` is equal to `check_transfer_return`.
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
	JSON          interestPaymentSimulationResultTransactionSourceCheckTransferReturnJSON
}

// interestPaymentSimulationResultTransactionSourceCheckTransferReturnJSON contains
// the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceCheckTransferReturn]
type interestPaymentSimulationResultTransactionSourceCheckTransferReturnJSON struct {
	TransferID    apijson.Field
	ReturnedAt    apijson.Field
	FileID        apijson.Field
	Reason        apijson.Field
	TransactionID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceCheckTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceCheckTransferReturnReason string

const (
	InterestPaymentSimulationResultTransactionSourceCheckTransferReturnReasonMailDeliveryFailure   InterestPaymentSimulationResultTransactionSourceCheckTransferReturnReason = "mail_delivery_failure"
	InterestPaymentSimulationResultTransactionSourceCheckTransferReturnReasonRefusedByRecipient    InterestPaymentSimulationResultTransactionSourceCheckTransferReturnReason = "refused_by_recipient"
	InterestPaymentSimulationResultTransactionSourceCheckTransferReturnReasonReturnedNotAuthorized InterestPaymentSimulationResultTransactionSourceCheckTransferReturnReason = "returned_not_authorized"
)

// A Check Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_rejection`.
type InterestPaymentSimulationResultTransactionSourceCheckTransferRejection struct {
	// The identifier of the Check Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       interestPaymentSimulationResultTransactionSourceCheckTransferRejectionJSON
}

// interestPaymentSimulationResultTransactionSourceCheckTransferRejectionJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceCheckTransferRejection]
type interestPaymentSimulationResultTransactionSourceCheckTransferRejectionJSON struct {
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceCheckTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Check Transfer Stop Payment Request object. This field will be present in the
// JSON response if and only if `category` is equal to
// `check_transfer_stop_payment_request`.
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
	JSON interestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestJSON
}

// interestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequest]
type interestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestJSON struct {
	TransferID    apijson.Field
	TransactionID apijson.Field
	RequestedAt   apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestType string

const (
	InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

// A Dispute Resolution object. This field will be present in the JSON response if
// and only if `category` is equal to `dispute_resolution`.
type InterestPaymentSimulationResultTransactionSourceDisputeResolution struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency InterestPaymentSimulationResultTransactionSourceDisputeResolutionCurrency `json:"currency,required"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionID string `json:"disputed_transaction_id,required"`
	JSON                  interestPaymentSimulationResultTransactionSourceDisputeResolutionJSON
}

// interestPaymentSimulationResultTransactionSourceDisputeResolutionJSON contains
// the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceDisputeResolution]
type interestPaymentSimulationResultTransactionSourceDisputeResolutionJSON struct {
	Amount                apijson.Field
	Currency              apijson.Field
	DisputedTransactionID apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceDisputeResolution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// A Empyreal Cash Deposit object. This field will be present in the JSON response
// if and only if `category` is equal to `empyreal_cash_deposit`.
type InterestPaymentSimulationResultTransactionSourceEmpyrealCashDeposit struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount      int64     `json:"amount,required"`
	BagID       string    `json:"bag_id,required"`
	DepositDate time.Time `json:"deposit_date,required" format:"date-time"`
	JSON        interestPaymentSimulationResultTransactionSourceEmpyrealCashDepositJSON
}

// interestPaymentSimulationResultTransactionSourceEmpyrealCashDepositJSON contains
// the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceEmpyrealCashDeposit]
type interestPaymentSimulationResultTransactionSourceEmpyrealCashDepositJSON struct {
	Amount      apijson.Field
	BagID       apijson.Field
	DepositDate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceEmpyrealCashDeposit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Fee Payment object. This field will be present in the JSON response if and
// only if `category` is equal to `fee_payment`.
type InterestPaymentSimulationResultTransactionSourceFeePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency InterestPaymentSimulationResultTransactionSourceFeePaymentCurrency `json:"currency,required"`
	JSON     interestPaymentSimulationResultTransactionSourceFeePaymentJSON
}

// interestPaymentSimulationResultTransactionSourceFeePaymentJSON contains the JSON
// metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceFeePayment]
type interestPaymentSimulationResultTransactionSourceFeePaymentJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceFeePayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// A Inbound ACH Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_ach_transfer`.
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
	JSON                               interestPaymentSimulationResultTransactionSourceInboundACHTransferJSON
}

// interestPaymentSimulationResultTransactionSourceInboundACHTransferJSON contains
// the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceInboundACHTransfer]
type interestPaymentSimulationResultTransactionSourceInboundACHTransferJSON struct {
	Amount                             apijson.Field
	OriginatorCompanyName              apijson.Field
	OriginatorCompanyDescriptiveDate   apijson.Field
	OriginatorCompanyDiscretionaryData apijson.Field
	OriginatorCompanyEntryDescription  apijson.Field
	OriginatorCompanyID                apijson.Field
	ReceiverIDNumber                   apijson.Field
	ReceiverName                       apijson.Field
	TraceNumber                        apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceInboundACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Inbound Check object. This field will be present in the JSON response if and
// only if `category` is equal to `inbound_check`.
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
	JSON                  interestPaymentSimulationResultTransactionSourceInboundCheckJSON
}

// interestPaymentSimulationResultTransactionSourceInboundCheckJSON contains the
// JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceInboundCheck]
type interestPaymentSimulationResultTransactionSourceInboundCheckJSON struct {
	Amount                apijson.Field
	Currency              apijson.Field
	CheckNumber           apijson.Field
	CheckFrontImageFileID apijson.Field
	CheckRearImageFileID  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceInboundCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// A Inbound International ACH Transfer object. This field will be present in the
// JSON response if and only if `category` is equal to
// `inbound_international_ach_transfer`.
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
	JSON                                                   interestPaymentSimulationResultTransactionSourceInboundInternationalACHTransferJSON
}

// interestPaymentSimulationResultTransactionSourceInboundInternationalACHTransferJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceInboundInternationalACHTransfer]
type interestPaymentSimulationResultTransactionSourceInboundInternationalACHTransferJSON struct {
	Amount                                                 apijson.Field
	ForeignExchangeIndicator                               apijson.Field
	ForeignExchangeReferenceIndicator                      apijson.Field
	ForeignExchangeReference                               apijson.Field
	DestinationCountryCode                                 apijson.Field
	DestinationCurrencyCode                                apijson.Field
	ForeignPaymentAmount                                   apijson.Field
	ForeignTraceNumber                                     apijson.Field
	InternationalTransactionTypeCode                       apijson.Field
	OriginatingCurrencyCode                                apijson.Field
	OriginatingDepositoryFinancialInstitutionName          apijson.Field
	OriginatingDepositoryFinancialInstitutionIDQualifier   apijson.Field
	OriginatingDepositoryFinancialInstitutionID            apijson.Field
	OriginatingDepositoryFinancialInstitutionBranchCountry apijson.Field
	OriginatorCity                                         apijson.Field
	OriginatorCompanyEntryDescription                      apijson.Field
	OriginatorCountry                                      apijson.Field
	OriginatorIdentification                               apijson.Field
	OriginatorName                                         apijson.Field
	OriginatorPostalCode                                   apijson.Field
	OriginatorStreetAddress                                apijson.Field
	OriginatorStateOrProvince                              apijson.Field
	PaymentRelatedInformation                              apijson.Field
	PaymentRelatedInformation2                             apijson.Field
	ReceiverIdentificationNumber                           apijson.Field
	ReceiverStreetAddress                                  apijson.Field
	ReceiverCity                                           apijson.Field
	ReceiverStateOrProvince                                apijson.Field
	ReceiverCountry                                        apijson.Field
	ReceiverPostalCode                                     apijson.Field
	ReceivingCompanyOrIndividualName                       apijson.Field
	ReceivingDepositoryFinancialInstitutionName            apijson.Field
	ReceivingDepositoryFinancialInstitutionIDQualifier     apijson.Field
	ReceivingDepositoryFinancialInstitutionID              apijson.Field
	ReceivingDepositoryFinancialInstitutionCountry         apijson.Field
	TraceNumber                                            apijson.Field
	raw                                                    string
	ExtraFields                                            map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceInboundInternationalACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Inbound Real Time Payments Transfer Confirmation object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_confirmation`.
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
	JSON                  interestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON
}

// interestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation]
type interestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON struct {
	Amount                    apijson.Field
	Currency                  apijson.Field
	CreditorName              apijson.Field
	DebtorName                apijson.Field
	DebtorAccountNumber       apijson.Field
	DebtorRoutingNumber       apijson.Field
	TransactionIdentification apijson.Field
	RemittanceInformation     apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
// the JSON response if and only if `category` is equal to
// `inbound_wire_drawdown_payment_reversal`.
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
	JSON                       interestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentReversalJSON
}

// interestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentReversalJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal]
type interestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentReversalJSON struct {
	Amount                                        apijson.Field
	Description                                   apijson.Field
	InputCycleDate                                apijson.Field
	InputSequenceNumber                           apijson.Field
	InputSource                                   apijson.Field
	InputMessageAccountabilityData                apijson.Field
	PreviousMessageInputMessageAccountabilityData apijson.Field
	PreviousMessageInputCycleDate                 apijson.Field
	PreviousMessageInputSequenceNumber            apijson.Field
	PreviousMessageInputSource                    apijson.Field
	raw                                           string
	ExtraFields                                   map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
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
	JSON                               interestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentJSON
}

// interestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceInboundWireDrawdownPayment]
type interestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentJSON struct {
	Amount                             apijson.Field
	BeneficiaryAddressLine1            apijson.Field
	BeneficiaryAddressLine2            apijson.Field
	BeneficiaryAddressLine3            apijson.Field
	BeneficiaryName                    apijson.Field
	BeneficiaryReference               apijson.Field
	Description                        apijson.Field
	InputMessageAccountabilityData     apijson.Field
	OriginatorAddressLine1             apijson.Field
	OriginatorAddressLine2             apijson.Field
	OriginatorAddressLine3             apijson.Field
	OriginatorName                     apijson.Field
	OriginatorToBeneficiaryInformation apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceInboundWireDrawdownPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Inbound Wire Reversal object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_reversal`.
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
	// The ID for the Wire Transfer that is being reversed.
	WireTransferID string `json:"wire_transfer_id,required"`
	JSON           interestPaymentSimulationResultTransactionSourceInboundWireReversalJSON
}

// interestPaymentSimulationResultTransactionSourceInboundWireReversalJSON contains
// the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceInboundWireReversal]
type interestPaymentSimulationResultTransactionSourceInboundWireReversalJSON struct {
	Amount                                                apijson.Field
	CreatedAt                                             apijson.Field
	Description                                           apijson.Field
	InputCycleDate                                        apijson.Field
	InputSequenceNumber                                   apijson.Field
	InputSource                                           apijson.Field
	InputMessageAccountabilityData                        apijson.Field
	PreviousMessageInputMessageAccountabilityData         apijson.Field
	PreviousMessageInputCycleDate                         apijson.Field
	PreviousMessageInputSequenceNumber                    apijson.Field
	PreviousMessageInputSource                            apijson.Field
	ReceiverFinancialInstitutionInformation               apijson.Field
	FinancialInstitutionToFinancialInstitutionInformation apijson.Field
	TransactionID                                         apijson.Field
	WireTransferID                                        apijson.Field
	raw                                                   string
	ExtraFields                                           map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceInboundWireReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Inbound Wire Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_transfer`.
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
	JSON                                    interestPaymentSimulationResultTransactionSourceInboundWireTransferJSON
}

// interestPaymentSimulationResultTransactionSourceInboundWireTransferJSON contains
// the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceInboundWireTransfer]
type interestPaymentSimulationResultTransactionSourceInboundWireTransferJSON struct {
	Amount                                  apijson.Field
	BeneficiaryAddressLine1                 apijson.Field
	BeneficiaryAddressLine2                 apijson.Field
	BeneficiaryAddressLine3                 apijson.Field
	BeneficiaryName                         apijson.Field
	BeneficiaryReference                    apijson.Field
	Description                             apijson.Field
	InputMessageAccountabilityData          apijson.Field
	OriginatorAddressLine1                  apijson.Field
	OriginatorAddressLine2                  apijson.Field
	OriginatorAddressLine3                  apijson.Field
	OriginatorName                          apijson.Field
	OriginatorToBeneficiaryInformationLine1 apijson.Field
	OriginatorToBeneficiaryInformationLine2 apijson.Field
	OriginatorToBeneficiaryInformationLine3 apijson.Field
	OriginatorToBeneficiaryInformationLine4 apijson.Field
	OriginatorToBeneficiaryInformation      apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceInboundWireTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Interest Payment object. This field will be present in the JSON response if
// and only if `category` is equal to `interest_payment`.
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
	JSON               interestPaymentSimulationResultTransactionSourceInterestPaymentJSON
}

// interestPaymentSimulationResultTransactionSourceInterestPaymentJSON contains the
// JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceInterestPayment]
type interestPaymentSimulationResultTransactionSourceInterestPaymentJSON struct {
	Amount             apijson.Field
	Currency           apijson.Field
	PeriodStart        apijson.Field
	PeriodEnd          apijson.Field
	AccruedOnAccountID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceInterestPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// A Internal Source object. This field will be present in the JSON response if and
// only if `category` is equal to `internal_source`.
type InterestPaymentSimulationResultTransactionSourceInternalSource struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency InterestPaymentSimulationResultTransactionSourceInternalSourceCurrency `json:"currency,required"`
	Reason   InterestPaymentSimulationResultTransactionSourceInternalSourceReason   `json:"reason,required"`
	JSON     interestPaymentSimulationResultTransactionSourceInternalSourceJSON
}

// interestPaymentSimulationResultTransactionSourceInternalSourceJSON contains the
// JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceInternalSource]
type interestPaymentSimulationResultTransactionSourceInternalSourceJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceInternalSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// A Deprecated Card Refund object. This field will be present in the JSON response
// if and only if `category` is equal to `card_route_refund`.
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
	JSON                 interestPaymentSimulationResultTransactionSourceCardRouteRefundJSON
}

// interestPaymentSimulationResultTransactionSourceCardRouteRefundJSON contains the
// JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceCardRouteRefund]
type interestPaymentSimulationResultTransactionSourceCardRouteRefundJSON struct {
	Amount               apijson.Field
	Currency             apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantDescriptor   apijson.Field
	MerchantState        apijson.Field
	MerchantCategoryCode apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceCardRouteRefund) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// A Deprecated Card Settlement object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_settlement`.
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
	JSON                 interestPaymentSimulationResultTransactionSourceCardRouteSettlementJSON
}

// interestPaymentSimulationResultTransactionSourceCardRouteSettlementJSON contains
// the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceCardRouteSettlement]
type interestPaymentSimulationResultTransactionSourceCardRouteSettlementJSON struct {
	Amount               apijson.Field
	Currency             apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantDescriptor   apijson.Field
	MerchantState        apijson.Field
	MerchantCategoryCode apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceCardRouteSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// A Real Time Payments Transfer Acknowledgement object. This field will be present
// in the JSON response if and only if `category` is equal to
// `real_time_payments_transfer_acknowledgement`.
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
	JSON       interestPaymentSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgementJSON
}

// interestPaymentSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgementJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgement]
type interestPaymentSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgementJSON struct {
	Amount                   apijson.Field
	DestinationAccountNumber apijson.Field
	DestinationRoutingNumber apijson.Field
	RemittanceInformation    apijson.Field
	TransferID               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Sample Funds object. This field will be present in the JSON response if and
// only if `category` is equal to `sample_funds`.
type InterestPaymentSimulationResultTransactionSourceSampleFunds struct {
	// Where the sample funds came from.
	Originator string `json:"originator,required"`
	JSON       interestPaymentSimulationResultTransactionSourceSampleFundsJSON
}

// interestPaymentSimulationResultTransactionSourceSampleFundsJSON contains the
// JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceSampleFunds]
type interestPaymentSimulationResultTransactionSourceSampleFundsJSON struct {
	Originator  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceSampleFunds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Wire Drawdown Payment Intention object. This field will be present in the JSON
// response if and only if `category` is equal to
// `wire_drawdown_payment_intention`.
type InterestPaymentSimulationResultTransactionSourceWireDrawdownPaymentIntention struct {
	// The transfer amount in USD cents.
	Amount             int64  `json:"amount,required"`
	AccountNumber      string `json:"account_number,required"`
	RoutingNumber      string `json:"routing_number,required"`
	MessageToRecipient string `json:"message_to_recipient,required"`
	TransferID         string `json:"transfer_id,required"`
	JSON               interestPaymentSimulationResultTransactionSourceWireDrawdownPaymentIntentionJSON
}

// interestPaymentSimulationResultTransactionSourceWireDrawdownPaymentIntentionJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceWireDrawdownPaymentIntention]
type interestPaymentSimulationResultTransactionSourceWireDrawdownPaymentIntentionJSON struct {
	Amount             apijson.Field
	AccountNumber      apijson.Field
	RoutingNumber      apijson.Field
	MessageToRecipient apijson.Field
	TransferID         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceWireDrawdownPaymentIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to
// `wire_drawdown_payment_rejection`.
type InterestPaymentSimulationResultTransactionSourceWireDrawdownPaymentRejection struct {
	TransferID string `json:"transfer_id,required"`
	JSON       interestPaymentSimulationResultTransactionSourceWireDrawdownPaymentRejectionJSON
}

// interestPaymentSimulationResultTransactionSourceWireDrawdownPaymentRejectionJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceWireDrawdownPaymentRejection]
type interestPaymentSimulationResultTransactionSourceWireDrawdownPaymentRejectionJSON struct {
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceWireDrawdownPaymentRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Wire Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_intention`.
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
	JSON               interestPaymentSimulationResultTransactionSourceWireTransferIntentionJSON
}

// interestPaymentSimulationResultTransactionSourceWireTransferIntentionJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceWireTransferIntention]
type interestPaymentSimulationResultTransactionSourceWireTransferIntentionJSON struct {
	Amount             apijson.Field
	AccountNumber      apijson.Field
	RoutingNumber      apijson.Field
	MessageToRecipient apijson.Field
	TransferID         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceWireTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Wire Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_rejection`.
type InterestPaymentSimulationResultTransactionSourceWireTransferRejection struct {
	TransferID string `json:"transfer_id,required"`
	JSON       interestPaymentSimulationResultTransactionSourceWireTransferRejectionJSON
}

// interestPaymentSimulationResultTransactionSourceWireTransferRejectionJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceWireTransferRejection]
type interestPaymentSimulationResultTransactionSourceWireTransferRejectionJSON struct {
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceWireTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InterestPaymentSimulationResultTransactionType string

const (
	InterestPaymentSimulationResultTransactionTypeTransaction InterestPaymentSimulationResultTransactionType = "transaction"
)

type InterestPaymentSimulationResultType string

const (
	InterestPaymentSimulationResultTypeInterestPaymentSimulationResult InterestPaymentSimulationResultType = "interest_payment_simulation_result"
)

type SimulationInterestPaymentNewParams struct {
	// The identifier of the Account Number the Interest Payment is for.
	AccountID param.Field[string] `json:"account_id,required"`
	// The interest amount in cents. Must be positive.
	Amount param.Field[int64] `json:"amount,required"`
}

func (r SimulationInterestPaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
