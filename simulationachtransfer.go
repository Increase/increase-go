package increase

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

type SimulationACHTransferService struct {
	Options []option.RequestOption
}

func NewSimulationACHTransferService(opts ...option.RequestOption) (r *SimulationACHTransferService) {
	r = &SimulationACHTransferService{}
	r.Options = opts
	return
}

// Simulates an inbound ACH transfer to your account. This imitates initiating a
// transaction to an Increase account from a different financial institution. The
// transfer may be either a credit or a debit depending on if the `amount` is
// positive or negative. The result of calling this API will be either a
// [Transaction](#transactions) or a [Declined Transaction](#declined-transactions)
// depending on whether or not the transfer is allowed.
func (r *SimulationACHTransferService) NewInbound(ctx context.Context, body SimulationACHTransferNewInboundParams, opts ...option.RequestOption) (res *ACHTransferSimulation, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_ach_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates the return of an [ACH Transfer](#ach-transfers) by the Federal Reserve
// due to an error condition. This will also create a Transaction to account for
// the returned funds. This transfer must first have a `status` of `submitted`.
func (r *SimulationACHTransferService) Return(ctx context.Context, ach_transfer_id string, body SimulationACHTransferReturnParams, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/ach_transfers/%s/return", ach_transfer_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates the submission of an [ACH Transfer](#ach-transfers) to the Federal
// Reserve. This transfer must first have a `status` of `pending_approval` or
// `pending_submission`. In production, Increase submits ACH Transfers to the
// Federal Reserve three times per day on weekdays. Since sandbox ACH Transfers are
// not submitted to the Federal Reserve, this endpoint allows you to skip that
// delay and transition the ACH Transfer to a status of `submitted`.
func (r *SimulationACHTransferService) Submit(ctx context.Context, ach_transfer_id string, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/ach_transfers/%s/submit", ach_transfer_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type ACHTransferSimulation struct {
	// If the ACH Transfer attempt succeeds, this will contain the resulting
	// [Transaction](#transactions) object. The Transaction's `source` will be of
	// `category: inbound_ach_transfer`.
	Transaction ACHTransferSimulationTransaction `json:"transaction,required,nullable"`
	// If the ACH Transfer attempt fails, this will contain the resulting
	// [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of `category: inbound_ach_transfer`.
	DeclinedTransaction ACHTransferSimulationDeclinedTransaction `json:"declined_transaction,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_ach_transfer_simulation_result`.
	Type ACHTransferSimulationType `json:"type,required"`
	JSON ACHTransferSimulationJSON
}

type ACHTransferSimulationJSON struct {
	Transaction         apijson.Metadata
	DeclinedTransaction apijson.Metadata
	Type                apijson.Metadata
	raw                 string
	Extras              map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ACHTransferSimulation using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransaction struct {
	// The identifier for the Account the Transaction belongs to.
	AccountID string `json:"account_id,required"`
	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transcation's
	// Account.
	Currency ACHTransferSimulationTransactionCurrency `json:"currency,required"`
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
	RouteType ACHTransferSimulationTransactionRouteType `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Transaction. Note that for backwards compatibility reasons, additional
	// undocumented keys may appear in this object. These should be treated as
	// deprecated and will be removed in the future.
	Source ACHTransferSimulationTransactionSource `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type ACHTransferSimulationTransactionType `json:"type,required"`
	JSON ACHTransferSimulationTransactionJSON
}

type ACHTransferSimulationTransactionJSON struct {
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
// ACHTransferSimulationTransaction using the internal json library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Category ACHTransferSimulationTransactionSourceCategory `json:"category,required"`
	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention ACHTransferSimulationTransactionSourceAccountTransferIntention `json:"account_transfer_intention,required,nullable"`
	// A ACH Check Conversion Return object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_check_conversion_return`.
	ACHCheckConversionReturn ACHTransferSimulationTransactionSourceACHCheckConversionReturn `json:"ach_check_conversion_return,required,nullable"`
	// A ACH Check Conversion object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_check_conversion`.
	ACHCheckConversion ACHTransferSimulationTransactionSourceACHCheckConversion `json:"ach_check_conversion,required,nullable"`
	// A ACH Transfer Intention object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention ACHTransferSimulationTransactionSourceACHTransferIntention `json:"ach_transfer_intention,required,nullable"`
	// A ACH Transfer Rejection object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection ACHTransferSimulationTransactionSourceACHTransferRejection `json:"ach_transfer_rejection,required,nullable"`
	// A ACH Transfer Return object. This field will be present in the JSON response if
	// and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn ACHTransferSimulationTransactionSourceACHTransferReturn `json:"ach_transfer_return,required,nullable"`
	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance ACHTransferSimulationTransactionSourceCardDisputeAcceptance `json:"card_dispute_acceptance,required,nullable"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund ACHTransferSimulationTransactionSourceCardRefund `json:"card_refund,required,nullable"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement ACHTransferSimulationTransactionSourceCardSettlement `json:"card_settlement,required,nullable"`
	// A Card Revenue Payment object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_revenue_payment`.
	CardRevenuePayment ACHTransferSimulationTransactionSourceCardRevenuePayment `json:"card_revenue_payment,required,nullable"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance ACHTransferSimulationTransactionSourceCheckDepositAcceptance `json:"check_deposit_acceptance,required,nullable"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn ACHTransferSimulationTransactionSourceCheckDepositReturn `json:"check_deposit_return,required,nullable"`
	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention ACHTransferSimulationTransactionSourceCheckTransferIntention `json:"check_transfer_intention,required,nullable"`
	// A Check Transfer Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_transfer_return`.
	CheckTransferReturn ACHTransferSimulationTransactionSourceCheckTransferReturn `json:"check_transfer_return,required,nullable"`
	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection ACHTransferSimulationTransactionSourceCheckTransferRejection `json:"check_transfer_rejection,required,nullable"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest `json:"check_transfer_stop_payment_request,required,nullable"`
	// A Dispute Resolution object. This field will be present in the JSON response if
	// and only if `category` is equal to `dispute_resolution`.
	DisputeResolution ACHTransferSimulationTransactionSourceDisputeResolution `json:"dispute_resolution,required,nullable"`
	// A Empyreal Cash Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `empyreal_cash_deposit`.
	EmpyrealCashDeposit ACHTransferSimulationTransactionSourceEmpyrealCashDeposit `json:"empyreal_cash_deposit,required,nullable"`
	// A Fee Payment object. This field will be present in the JSON response if and
	// only if `category` is equal to `fee_payment`.
	FeePayment ACHTransferSimulationTransactionSourceFeePayment `json:"fee_payment,required,nullable"`
	// A Inbound ACH Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer ACHTransferSimulationTransactionSourceInboundACHTransfer `json:"inbound_ach_transfer,required,nullable"`
	// A Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck ACHTransferSimulationTransactionSourceInboundCheck `json:"inbound_check,required,nullable"`
	// A Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer `json:"inbound_international_ach_transfer,required,nullable"`
	// A Inbound Real Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation `json:"inbound_real_time_payments_transfer_confirmation,required,nullable"`
	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal `json:"inbound_wire_drawdown_payment_reversal,required,nullable"`
	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment `json:"inbound_wire_drawdown_payment,required,nullable"`
	// A Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal ACHTransferSimulationTransactionSourceInboundWireReversal `json:"inbound_wire_reversal,required,nullable"`
	// A Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer ACHTransferSimulationTransactionSourceInboundWireTransfer `json:"inbound_wire_transfer,required,nullable"`
	// A Interest Payment object. This field will be present in the JSON response if
	// and only if `category` is equal to `interest_payment`.
	InterestPayment ACHTransferSimulationTransactionSourceInterestPayment `json:"interest_payment,required,nullable"`
	// A Internal Source object. This field will be present in the JSON response if and
	// only if `category` is equal to `internal_source`.
	InternalSource ACHTransferSimulationTransactionSourceInternalSource `json:"internal_source,required,nullable"`
	// A Deprecated Card Refund object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_route_refund`.
	CardRouteRefund ACHTransferSimulationTransactionSourceCardRouteRefund `json:"card_route_refund,required,nullable"`
	// A Deprecated Card Settlement object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_settlement`.
	CardRouteSettlement ACHTransferSimulationTransactionSourceCardRouteSettlement `json:"card_route_settlement,required,nullable"`
	// A Real Time Payments Transfer Acknowledgement object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `real_time_payments_transfer_acknowledgement`.
	RealTimePaymentsTransferAcknowledgement ACHTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgement `json:"real_time_payments_transfer_acknowledgement,required,nullable"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds ACHTransferSimulationTransactionSourceSampleFunds `json:"sample_funds,required,nullable"`
	// A Wire Drawdown Payment Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_intention`.
	WireDrawdownPaymentIntention ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention `json:"wire_drawdown_payment_intention,required,nullable"`
	// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_rejection`.
	WireDrawdownPaymentRejection ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection `json:"wire_drawdown_payment_rejection,required,nullable"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention ACHTransferSimulationTransactionSourceWireTransferIntention `json:"wire_transfer_intention,required,nullable"`
	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection ACHTransferSimulationTransactionSourceWireTransferRejection `json:"wire_transfer_rejection,required,nullable"`
	JSON                  ACHTransferSimulationTransactionSourceJSON
}

type ACHTransferSimulationTransactionSourceJSON struct {
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
// ACHTransferSimulationTransactionSource using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	ACHTransferSimulationTransactionSourceCategoryCardRevenuePayment                          ACHTransferSimulationTransactionSourceCategory = "card_revenue_payment"
	ACHTransferSimulationTransactionSourceCategoryCheckDepositAcceptance                      ACHTransferSimulationTransactionSourceCategory = "check_deposit_acceptance"
	ACHTransferSimulationTransactionSourceCategoryCheckDepositReturn                          ACHTransferSimulationTransactionSourceCategory = "check_deposit_return"
	ACHTransferSimulationTransactionSourceCategoryCheckTransferIntention                      ACHTransferSimulationTransactionSourceCategory = "check_transfer_intention"
	ACHTransferSimulationTransactionSourceCategoryCheckTransferReturn                         ACHTransferSimulationTransactionSourceCategory = "check_transfer_return"
	ACHTransferSimulationTransactionSourceCategoryCheckTransferRejection                      ACHTransferSimulationTransactionSourceCategory = "check_transfer_rejection"
	ACHTransferSimulationTransactionSourceCategoryCheckTransferStopPaymentRequest             ACHTransferSimulationTransactionSourceCategory = "check_transfer_stop_payment_request"
	ACHTransferSimulationTransactionSourceCategoryDisputeResolution                           ACHTransferSimulationTransactionSourceCategory = "dispute_resolution"
	ACHTransferSimulationTransactionSourceCategoryEmpyrealCashDeposit                         ACHTransferSimulationTransactionSourceCategory = "empyreal_cash_deposit"
	ACHTransferSimulationTransactionSourceCategoryFeePayment                                  ACHTransferSimulationTransactionSourceCategory = "fee_payment"
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
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency `json:"currency,required"`
	// The description you chose to give the transfer.
	Description string `json:"description,required"`
	// The identifier of the Account to where the Account Transfer was sent.
	DestinationAccountID string `json:"destination_account_id,required"`
	// The identifier of the Account from where the Account Transfer was sent.
	SourceAccountID string `json:"source_account_id,required"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       ACHTransferSimulationTransactionSourceAccountTransferIntentionJSON
}

type ACHTransferSimulationTransactionSourceAccountTransferIntentionJSON struct {
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
// ACHTransferSimulationTransactionSourceAccountTransferIntention using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationTransactionSourceAccountTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// Why the transfer was returned.
	ReturnReasonCode string `json:"return_reason_code,required"`
	JSON             ACHTransferSimulationTransactionSourceACHCheckConversionReturnJSON
}

type ACHTransferSimulationTransactionSourceACHCheckConversionReturnJSON struct {
	Amount           apijson.Metadata
	ReturnReasonCode apijson.Metadata
	raw              string
	Extras           map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceACHCheckConversionReturn using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationTransactionSourceACHCheckConversionReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceACHCheckConversion struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the File containing an image of the returned check.
	FileID string `json:"file_id,required"`
	JSON   ACHTransferSimulationTransactionSourceACHCheckConversionJSON
}

type ACHTransferSimulationTransactionSourceACHCheckConversionJSON struct {
	Amount apijson.Metadata
	FileID apijson.Metadata
	raw    string
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceACHCheckConversion using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceACHCheckConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceACHTransferIntention struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount              int64  `json:"amount,required"`
	AccountNumber       string `json:"account_number,required"`
	RoutingNumber       string `json:"routing_number,required"`
	StatementDescriptor string `json:"statement_descriptor,required"`
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       ACHTransferSimulationTransactionSourceACHTransferIntentionJSON
}

type ACHTransferSimulationTransactionSourceACHTransferIntentionJSON struct {
	Amount              apijson.Metadata
	AccountNumber       apijson.Metadata
	RoutingNumber       apijson.Metadata
	StatementDescriptor apijson.Metadata
	TransferID          apijson.Metadata
	raw                 string
	Extras              map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceACHTransferIntention using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceACHTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       ACHTransferSimulationTransactionSourceACHTransferRejectionJSON
}

type ACHTransferSimulationTransactionSourceACHTransferRejectionJSON struct {
	TransferID apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceACHTransferRejection using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceACHTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode `json:"return_reason_code,required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID string `json:"transfer_id,required"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID string `json:"transaction_id,required"`
	JSON          ACHTransferSimulationTransactionSourceACHTransferReturnJSON
}

type ACHTransferSimulationTransactionSourceACHTransferReturnJSON struct {
	CreatedAt        apijson.Metadata
	ReturnReasonCode apijson.Metadata
	TransferID       apijson.Metadata
	TransactionID    apijson.Metadata
	raw              string
	Extras           map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceACHTransferReturn using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	AcceptedAt time.Time `json:"accepted_at,required" format:"date-time"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID string `json:"card_dispute_id,required"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID string `json:"transaction_id,required"`
	JSON          ACHTransferSimulationTransactionSourceCardDisputeAcceptanceJSON
}

type ACHTransferSimulationTransactionSourceCardDisputeAcceptanceJSON struct {
	AcceptedAt    apijson.Metadata
	CardDisputeID apijson.Metadata
	TransactionID apijson.Metadata
	raw           string
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceCardDisputeAcceptance using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCardDisputeAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceCardRefund struct {
	// The Card Refund identifier.
	ID string `json:"id,required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency ACHTransferSimulationTransactionSourceCardRefundCurrency `json:"currency,required"`
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
	Type ACHTransferSimulationTransactionSourceCardRefundType `json:"type,required"`
	JSON ACHTransferSimulationTransactionSourceCardRefundJSON
}

type ACHTransferSimulationTransactionSourceCardRefundJSON struct {
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
// ACHTransferSimulationTransactionSourceCardRefund using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCardRefund) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Currency ACHTransferSimulationTransactionSourceCardSettlementCurrency `json:"currency,required"`
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
	Type ACHTransferSimulationTransactionSourceCardSettlementType `json:"type,required"`
	JSON ACHTransferSimulationTransactionSourceCardSettlementJSON
}

type ACHTransferSimulationTransactionSourceCardSettlementJSON struct {
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
// ACHTransferSimulationTransactionSourceCardSettlement using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCardSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type ACHTransferSimulationTransactionSourceCardRevenuePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrency `json:"currency,required"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time `json:"period_start,required" format:"date-time"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The account the card belonged to.
	TransactedOnAccountID string `json:"transacted_on_account_id,required,nullable"`
	JSON                  ACHTransferSimulationTransactionSourceCardRevenuePaymentJSON
}

type ACHTransferSimulationTransactionSourceCardRevenuePaymentJSON struct {
	Amount                apijson.Metadata
	Currency              apijson.Metadata
	PeriodStart           apijson.Metadata
	PeriodEnd             apijson.Metadata
	TransactedOnAccountID apijson.Metadata
	raw                   string
	Extras                map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceCardRevenuePayment using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCardRevenuePayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrency string

const (
	ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrencyCad ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "CAD"
	ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrencyChf ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "CHF"
	ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrencyEur ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "EUR"
	ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrencyGbp ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "GBP"
	ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrencyJpy ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "JPY"
	ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrencyUsd ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceCheckDepositAcceptance struct {
	// The amount to be deposited in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency `json:"currency,required"`
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
	JSON           ACHTransferSimulationTransactionSourceCheckDepositAcceptanceJSON
}

type ACHTransferSimulationTransactionSourceCheckDepositAcceptanceJSON struct {
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
// ACHTransferSimulationTransactionSourceCheckDepositAcceptance using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCheckDepositAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency `json:"currency,required"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID string `json:"check_deposit_id,required"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID string                                                               `json:"transaction_id,required"`
	ReturnReason  ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason `json:"return_reason,required"`
	JSON          ACHTransferSimulationTransactionSourceCheckDepositReturnJSON
}

type ACHTransferSimulationTransactionSourceCheckDepositReturnJSON struct {
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
// ACHTransferSimulationTransactionSourceCheckDepositReturn using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCheckDepositReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Currency ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency `json:"currency,required"`
	// The name that will be printed on the check.
	RecipientName string `json:"recipient_name,required"`
	// The identifier of the Check Transfer with which this is associated.
	TransferID string `json:"transfer_id,required"`
	JSON       ACHTransferSimulationTransactionSourceCheckTransferIntentionJSON
}

type ACHTransferSimulationTransactionSourceCheckTransferIntentionJSON struct {
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
// ACHTransferSimulationTransactionSourceCheckTransferIntention using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	TransferID string `json:"transfer_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// If available, a document with additional information about the return.
	FileID string `json:"file_id,required,nullable"`
	// The reason why the check was returned.
	Reason ACHTransferSimulationTransactionSourceCheckTransferReturnReason `json:"reason,required"`
	// The identifier of the Transaction that was created to credit you for the
	// returned check.
	TransactionID string `json:"transaction_id,required,nullable"`
	JSON          ACHTransferSimulationTransactionSourceCheckTransferReturnJSON
}

type ACHTransferSimulationTransactionSourceCheckTransferReturnJSON struct {
	TransferID    apijson.Metadata
	ReturnedAt    apijson.Metadata
	FileID        apijson.Metadata
	Reason        apijson.Metadata
	TransactionID apijson.Metadata
	raw           string
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceCheckTransferReturn using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCheckTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceCheckTransferReturnReason string

const (
	ACHTransferSimulationTransactionSourceCheckTransferReturnReasonMailDeliveryFailure ACHTransferSimulationTransactionSourceCheckTransferReturnReason = "mail_delivery_failure"
	ACHTransferSimulationTransactionSourceCheckTransferReturnReasonRefusedByRecipient  ACHTransferSimulationTransactionSourceCheckTransferReturnReason = "refused_by_recipient"
)

type ACHTransferSimulationTransactionSourceCheckTransferRejection struct {
	// The identifier of the Check Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       ACHTransferSimulationTransactionSourceCheckTransferRejectionJSON
}

type ACHTransferSimulationTransactionSourceCheckTransferRejectionJSON struct {
	TransferID apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceCheckTransferRejection using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCheckTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferID string `json:"transfer_id,required"`
	// The transaction ID of the corresponding credit transaction.
	TransactionID string `json:"transaction_id,required"`
	// The time the stop-payment was requested.
	RequestedAt time.Time `json:"requested_at,required" format:"date-time"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType `json:"type,required"`
	JSON ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestJSON
}

type ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestJSON struct {
	TransferID    apijson.Metadata
	TransactionID apijson.Metadata
	RequestedAt   apijson.Metadata
	Type          apijson.Metadata
	raw           string
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType string

const (
	ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

type ACHTransferSimulationTransactionSourceDisputeResolution struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency ACHTransferSimulationTransactionSourceDisputeResolutionCurrency `json:"currency,required"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionID string `json:"disputed_transaction_id,required"`
	JSON                  ACHTransferSimulationTransactionSourceDisputeResolutionJSON
}

type ACHTransferSimulationTransactionSourceDisputeResolutionJSON struct {
	Amount                apijson.Metadata
	Currency              apijson.Metadata
	DisputedTransactionID apijson.Metadata
	raw                   string
	Extras                map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceDisputeResolution using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceDisputeResolution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount      int64     `json:"amount,required"`
	BagID       string    `json:"bag_id,required"`
	DepositDate time.Time `json:"deposit_date,required" format:"date-time"`
	JSON        ACHTransferSimulationTransactionSourceEmpyrealCashDepositJSON
}

type ACHTransferSimulationTransactionSourceEmpyrealCashDepositJSON struct {
	Amount      apijson.Metadata
	BagID       apijson.Metadata
	DepositDate apijson.Metadata
	raw         string
	Extras      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceEmpyrealCashDeposit using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceEmpyrealCashDeposit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceFeePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency ACHTransferSimulationTransactionSourceFeePaymentCurrency `json:"currency,required"`
	JSON     ACHTransferSimulationTransactionSourceFeePaymentJSON
}

type ACHTransferSimulationTransactionSourceFeePaymentJSON struct {
	Amount   apijson.Metadata
	Currency apijson.Metadata
	raw      string
	Extras   map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceFeePayment using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceFeePayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceFeePaymentCurrency string

const (
	ACHTransferSimulationTransactionSourceFeePaymentCurrencyCad ACHTransferSimulationTransactionSourceFeePaymentCurrency = "CAD"
	ACHTransferSimulationTransactionSourceFeePaymentCurrencyChf ACHTransferSimulationTransactionSourceFeePaymentCurrency = "CHF"
	ACHTransferSimulationTransactionSourceFeePaymentCurrencyEur ACHTransferSimulationTransactionSourceFeePaymentCurrency = "EUR"
	ACHTransferSimulationTransactionSourceFeePaymentCurrencyGbp ACHTransferSimulationTransactionSourceFeePaymentCurrency = "GBP"
	ACHTransferSimulationTransactionSourceFeePaymentCurrencyJpy ACHTransferSimulationTransactionSourceFeePaymentCurrency = "JPY"
	ACHTransferSimulationTransactionSourceFeePaymentCurrencyUsd ACHTransferSimulationTransactionSourceFeePaymentCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceInboundACHTransfer struct {
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
	JSON                               ACHTransferSimulationTransactionSourceInboundACHTransferJSON
}

type ACHTransferSimulationTransactionSourceInboundACHTransferJSON struct {
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
// ACHTransferSimulationTransactionSourceInboundACHTransfer using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceInboundCheck struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency              ACHTransferSimulationTransactionSourceInboundCheckCurrency `json:"currency,required"`
	CheckNumber           string                                                     `json:"check_number,required,nullable"`
	CheckFrontImageFileID string                                                     `json:"check_front_image_file_id,required,nullable"`
	CheckRearImageFileID  string                                                     `json:"check_rear_image_file_id,required,nullable"`
	JSON                  ACHTransferSimulationTransactionSourceInboundCheckJSON
}

type ACHTransferSimulationTransactionSourceInboundCheckJSON struct {
	Amount                apijson.Metadata
	Currency              apijson.Metadata
	CheckNumber           apijson.Metadata
	CheckFrontImageFileID apijson.Metadata
	CheckRearImageFileID  apijson.Metadata
	raw                   string
	Extras                map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceInboundCheck using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceInboundCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	JSON                                                   ACHTransferSimulationTransactionSourceInboundInternationalACHTransferJSON
}

type ACHTransferSimulationTransactionSourceInboundInternationalACHTransferJSON struct {
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
// ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real Time Payments transfer.
	Currency ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency `json:"currency,required"`
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
	JSON                  ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON
}

type ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON struct {
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
// ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	JSON                       ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversalJSON
}

type ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversalJSON struct {
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
// ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment struct {
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
	JSON                               ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentJSON
}

type ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentJSON struct {
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
// ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceInboundWireReversal struct {
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
	JSON           ACHTransferSimulationTransactionSourceInboundWireReversalJSON
}

type ACHTransferSimulationTransactionSourceInboundWireReversalJSON struct {
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
	WireTransferID                                        apijson.Metadata
	raw                                                   string
	Extras                                                map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceInboundWireReversal using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceInboundWireTransfer struct {
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
	JSON                                    ACHTransferSimulationTransactionSourceInboundWireTransferJSON
}

type ACHTransferSimulationTransactionSourceInboundWireTransferJSON struct {
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
// ACHTransferSimulationTransactionSourceInboundWireTransfer using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceInterestPayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency ACHTransferSimulationTransactionSourceInterestPaymentCurrency `json:"currency,required"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time `json:"period_start,required" format:"date-time"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The account on which the interest was accrued.
	AccruedOnAccountID string `json:"accrued_on_account_id,required,nullable"`
	JSON               ACHTransferSimulationTransactionSourceInterestPaymentJSON
}

type ACHTransferSimulationTransactionSourceInterestPaymentJSON struct {
	Amount             apijson.Metadata
	Currency           apijson.Metadata
	PeriodStart        apijson.Metadata
	PeriodEnd          apijson.Metadata
	AccruedOnAccountID apijson.Metadata
	raw                string
	Extras             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceInterestPayment using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceInterestPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency ACHTransferSimulationTransactionSourceInternalSourceCurrency `json:"currency,required"`
	Reason   ACHTransferSimulationTransactionSourceInternalSourceReason   `json:"reason,required"`
	JSON     ACHTransferSimulationTransactionSourceInternalSourceJSON
}

type ACHTransferSimulationTransactionSourceInternalSourceJSON struct {
	Amount   apijson.Metadata
	Currency apijson.Metadata
	Reason   apijson.Metadata
	raw      string
	Extras   map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceInternalSource using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceInternalSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
	// currency.
	Currency             ACHTransferSimulationTransactionSourceCardRouteRefundCurrency `json:"currency,required"`
	MerchantAcceptorID   string                                                        `json:"merchant_acceptor_id,required"`
	MerchantCity         string                                                        `json:"merchant_city,required,nullable"`
	MerchantCountry      string                                                        `json:"merchant_country,required"`
	MerchantDescriptor   string                                                        `json:"merchant_descriptor,required"`
	MerchantState        string                                                        `json:"merchant_state,required,nullable"`
	MerchantCategoryCode string                                                        `json:"merchant_category_code,required,nullable"`
	JSON                 ACHTransferSimulationTransactionSourceCardRouteRefundJSON
}

type ACHTransferSimulationTransactionSourceCardRouteRefundJSON struct {
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
// ACHTransferSimulationTransactionSourceCardRouteRefund using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCardRouteRefund) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
	// currency.
	Currency             ACHTransferSimulationTransactionSourceCardRouteSettlementCurrency `json:"currency,required"`
	MerchantAcceptorID   string                                                            `json:"merchant_acceptor_id,required"`
	MerchantCity         string                                                            `json:"merchant_city,required,nullable"`
	MerchantCountry      string                                                            `json:"merchant_country,required,nullable"`
	MerchantDescriptor   string                                                            `json:"merchant_descriptor,required"`
	MerchantState        string                                                            `json:"merchant_state,required,nullable"`
	MerchantCategoryCode string                                                            `json:"merchant_category_code,required,nullable"`
	JSON                 ACHTransferSimulationTransactionSourceCardRouteSettlementJSON
}

type ACHTransferSimulationTransactionSourceCardRouteSettlementJSON struct {
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
// ACHTransferSimulationTransactionSourceCardRouteSettlement using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceCardRouteSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type ACHTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgement struct {
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
	JSON       ACHTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgementJSON
}

type ACHTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgementJSON struct {
	Amount                   apijson.Metadata
	DestinationAccountNumber apijson.Metadata
	DestinationRoutingNumber apijson.Metadata
	RemittanceInformation    apijson.Metadata
	TransferID               apijson.Metadata
	raw                      string
	Extras                   map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgement
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceSampleFunds struct {
	// Where the sample funds came from.
	Originator string `json:"originator,required"`
	JSON       ACHTransferSimulationTransactionSourceSampleFundsJSON
}

type ACHTransferSimulationTransactionSourceSampleFundsJSON struct {
	Originator apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceSampleFunds using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceSampleFunds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention struct {
	// The transfer amount in USD cents.
	Amount             int64  `json:"amount,required"`
	AccountNumber      string `json:"account_number,required"`
	RoutingNumber      string `json:"routing_number,required"`
	MessageToRecipient string `json:"message_to_recipient,required"`
	TransferID         string `json:"transfer_id,required"`
	JSON               ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntentionJSON
}

type ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntentionJSON struct {
	Amount             apijson.Metadata
	AccountNumber      apijson.Metadata
	RoutingNumber      apijson.Metadata
	MessageToRecipient apijson.Metadata
	TransferID         apijson.Metadata
	raw                string
	Extras             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection struct {
	TransferID string `json:"transfer_id,required"`
	JSON       ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejectionJSON
}

type ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejectionJSON struct {
	TransferID apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceWireTransferIntention struct {
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient string `json:"message_to_recipient,required"`
	TransferID         string `json:"transfer_id,required"`
	JSON               ACHTransferSimulationTransactionSourceWireTransferIntentionJSON
}

type ACHTransferSimulationTransactionSourceWireTransferIntentionJSON struct {
	Amount             apijson.Metadata
	AccountNumber      apijson.Metadata
	RoutingNumber      apijson.Metadata
	MessageToRecipient apijson.Metadata
	TransferID         apijson.Metadata
	raw                string
	Extras             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceWireTransferIntention using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceWireTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionSourceWireTransferRejection struct {
	TransferID string `json:"transfer_id,required"`
	JSON       ACHTransferSimulationTransactionSourceWireTransferRejectionJSON
}

type ACHTransferSimulationTransactionSourceWireTransferRejectionJSON struct {
	TransferID apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationTransactionSourceWireTransferRejection using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationTransactionSourceWireTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationTransactionType string

const (
	ACHTransferSimulationTransactionTypeTransaction ACHTransferSimulationTransactionType = "transaction"
)

type ACHTransferSimulationDeclinedTransaction struct {
	// The identifier for the Account the Declined Transaction belongs to.
	AccountID string `json:"account_id,required"`
	// The Declined Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transcation's Account.
	Currency ACHTransferSimulationDeclinedTransactionCurrency `json:"currency,required"`
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
	RouteType ACHTransferSimulationDeclinedTransactionRouteType `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Declined Transaction. For example, for a card transaction this lists the
	// merchant's industry and location. Note that for backwards compatibility reasons,
	// additional undocumented keys may appear in this object. These should be treated
	// as deprecated and will be removed in the future.
	Source ACHTransferSimulationDeclinedTransactionSource `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type ACHTransferSimulationDeclinedTransactionType `json:"type,required"`
	JSON ACHTransferSimulationDeclinedTransactionJSON
}

type ACHTransferSimulationDeclinedTransactionJSON struct {
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
// ACHTransferSimulationDeclinedTransaction using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationDeclinedTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Category ACHTransferSimulationDeclinedTransactionSourceCategory `json:"category,required"`
	// A ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline ACHTransferSimulationDeclinedTransactionSourceACHDecline `json:"ach_decline,required,nullable"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline ACHTransferSimulationDeclinedTransactionSourceCardDecline `json:"card_decline,required,nullable"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline ACHTransferSimulationDeclinedTransactionSourceCheckDecline `json:"check_decline,required,nullable"`
	// A Inbound Real Time Payments Transfer Decline object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline `json:"inbound_real_time_payments_transfer_decline,required,nullable"`
	// A International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline `json:"international_ach_decline,required,nullable"`
	// A Deprecated Card Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_decline`.
	CardRouteDecline ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline `json:"card_route_decline,required,nullable"`
	JSON             ACHTransferSimulationDeclinedTransactionSourceJSON
}

type ACHTransferSimulationDeclinedTransactionSourceJSON struct {
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
// ACHTransferSimulationDeclinedTransactionSource using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationDeclinedTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount                             int64  `json:"amount,required"`
	OriginatorCompanyName              string `json:"originator_company_name,required"`
	OriginatorCompanyDescriptiveDate   string `json:"originator_company_descriptive_date,required,nullable"`
	OriginatorCompanyDiscretionaryData string `json:"originator_company_discretionary_data,required,nullable"`
	OriginatorCompanyID                string `json:"originator_company_id,required"`
	// Why the ACH transfer was declined.
	Reason           ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason `json:"reason,required"`
	ReceiverIDNumber string                                                         `json:"receiver_id_number,required,nullable"`
	ReceiverName     string                                                         `json:"receiver_name,required,nullable"`
	TraceNumber      string                                                         `json:"trace_number,required"`
	JSON             ACHTransferSimulationDeclinedTransactionSourceACHDeclineJSON
}

type ACHTransferSimulationDeclinedTransactionSourceACHDeclineJSON struct {
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
// ACHTransferSimulationDeclinedTransactionSourceACHDecline using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason string

const (
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteCanceled             ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "ach_route_canceled"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteDisabled             ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "ach_route_disabled"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonBreachesLimit                ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonDuplicateReturn              ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "duplicate_return"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonEntityNotActive              ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "entity_not_active"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonGroupLocked                  ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "group_locked"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonInsufficientFunds            ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "insufficient_funds"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonMisroutedReturn              ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "misrouted_return"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonNoACHRoute                   ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "no_ach_route"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonOriginatorRequest            ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "originator_request"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonTransactionNotAllowed        ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "transaction_not_allowed"
)

type ACHTransferSimulationDeclinedTransactionSourceCardDecline struct {
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
	Network ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetwork `json:"network,required"`
	// Fields specific to the `network`
	NetworkDetails ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails `json:"network_details,required"`
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency `json:"currency,required"`
	// Why the transaction was declined.
	Reason ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason `json:"reason,required"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state,required,nullable"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// If the authorization was attempted using a Digital Wallet Token (such as an
	// Apple Pay purchase), the identifier of the token that was used.
	DigitalWalletTokenID string `json:"digital_wallet_token_id,required,nullable"`
	JSON                 ACHTransferSimulationDeclinedTransactionSourceCardDeclineJSON
}

type ACHTransferSimulationDeclinedTransactionSourceCardDeclineJSON struct {
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
// ACHTransferSimulationDeclinedTransactionSourceCardDecline using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetwork string

const (
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkVisa ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetwork = "visa"
)

type ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa `json:"visa,required"`
	JSON ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsJSON
}

type ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsJSON struct {
	Visa   apijson.Metadata
	raw    string
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode shared.PointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON
}

type ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Metadata
	PointOfServiceEntryMode     apijson.Metadata
	raw                         string
	Extras                      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator string

const (
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                          ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                               ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                             ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                   ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                     ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                    ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
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
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonCardNotActive                ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "card_not_active"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonEntityNotActive              ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "entity_not_active"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonGroupLocked                  ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "group_locked"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonInsufficientFunds            ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "insufficient_funds"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonCvv2Mismatch                 ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "cvv2_mismatch"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonTransactionNotAllowed        ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "transaction_not_allowed"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonBreachesInternalLimit        ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "breaches_internal_limit"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonBreachesLimit                ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "breaches_limit"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonWebhookDeclined              ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "webhook_declined"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonWebhookTimedOut              ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "webhook_timed_out"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonDeclinedByStandInProcessing  ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "declined_by_stand_in_processing"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonInvalidPhysicalCard          ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "invalid_physical_card"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonMissingOriginalAuthorization ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "missing_original_authorization"
)

type ACHTransferSimulationDeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount        int64  `json:"amount,required"`
	AuxiliaryOnUs string `json:"auxiliary_on_us,required,nullable"`
	// Why the check was declined.
	Reason ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason `json:"reason,required"`
	JSON   ACHTransferSimulationDeclinedTransactionSourceCheckDeclineJSON
}

type ACHTransferSimulationDeclinedTransactionSourceCheckDeclineJSON struct {
	Amount        apijson.Metadata
	AuxiliaryOnUs apijson.Metadata
	Reason        apijson.Metadata
	raw           string
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferSimulationDeclinedTransactionSourceCheckDecline using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHTransferSimulationDeclinedTransactionSourceCheckDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonNotAuthorized         ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "not_authorized"
)

type ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real Time Payments
	// transfer.
	Currency ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency `json:"currency,required"`
	// Why the transfer was declined.
	Reason ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason `json:"reason,required"`
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
	JSON                  ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON
}

type ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON struct {
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
// ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountRestricted          ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_restricted"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonGroupLocked                ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "group_locked"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonEntityNotActive            ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "entity_not_active"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "real_time_payments_not_enabled"
)

type ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline struct {
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
	JSON                                                   ACHTransferSimulationDeclinedTransactionSourceInternationalACHDeclineJSON
}

type ACHTransferSimulationDeclinedTransactionSourceInternationalACHDeclineJSON struct {
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
// ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency             ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrency `json:"currency,required"`
	MerchantAcceptorID   string                                                                 `json:"merchant_acceptor_id,required"`
	MerchantCity         string                                                                 `json:"merchant_city,required,nullable"`
	MerchantCountry      string                                                                 `json:"merchant_country,required"`
	MerchantDescriptor   string                                                                 `json:"merchant_descriptor,required"`
	MerchantState        string                                                                 `json:"merchant_state,required,nullable"`
	MerchantCategoryCode string                                                                 `json:"merchant_category_code,required,nullable"`
	JSON                 ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineJSON
}

type ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineJSON struct {
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
// ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type SimulationACHTransferNewInboundParams struct {
	// The identifier of the Account Number the inbound ACH Transfer is for.
	AccountNumberID field.Field[string] `json:"account_number_id,required"`
	// The transfer amount in cents. A positive amount originates a credit transfer
	// pushing funds to the receiving account. A negative amount originates a debit
	// transfer pulling funds from the receiving account.
	Amount field.Field[int64] `json:"amount,required"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate field.Field[string] `json:"company_descriptive_date"`
	// Data associated with the transfer set by the sender.
	CompanyDiscretionaryData field.Field[string] `json:"company_discretionary_data"`
	// The description of the transfer set by the sender.
	CompanyEntryDescription field.Field[string] `json:"company_entry_description"`
	// The name of the sender.
	CompanyName field.Field[string] `json:"company_name"`
	// The sender's company id.
	CompanyID field.Field[string] `json:"company_id"`
}

// MarshalJSON serializes SimulationACHTransferNewInboundParams into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r SimulationACHTransferNewInboundParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SimulationACHTransferReturnParams struct {
	// The reason why the Federal Reserve or destination bank returned this transfer.
	// Defaults to `no_account`.
	Reason field.Field[SimulationACHTransferReturnParamsReason] `json:"reason"`
}

// MarshalJSON serializes SimulationACHTransferReturnParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r SimulationACHTransferReturnParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SimulationACHTransferReturnParamsReason string

const (
	SimulationACHTransferReturnParamsReasonInsufficientFund                                          SimulationACHTransferReturnParamsReason = "insufficient_fund"
	SimulationACHTransferReturnParamsReasonNoAccount                                                 SimulationACHTransferReturnParamsReason = "no_account"
	SimulationACHTransferReturnParamsReasonAccountClosed                                             SimulationACHTransferReturnParamsReason = "account_closed"
	SimulationACHTransferReturnParamsReasonInvalidAccountNumberStructure                             SimulationACHTransferReturnParamsReason = "invalid_account_number_structure"
	SimulationACHTransferReturnParamsReasonAccountFrozenEntryReturnedPerOfacInstruction              SimulationACHTransferReturnParamsReason = "account_frozen_entry_returned_per_ofac_instruction"
	SimulationACHTransferReturnParamsReasonCreditEntryRefusedByReceiver                              SimulationACHTransferReturnParamsReason = "credit_entry_refused_by_receiver"
	SimulationACHTransferReturnParamsReasonUnauthorizedDebitToConsumerAccountUsingCorporateSecCode   SimulationACHTransferReturnParamsReason = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	SimulationACHTransferReturnParamsReasonCorporateCustomerAdvisedNotAuthorized                     SimulationACHTransferReturnParamsReason = "corporate_customer_advised_not_authorized"
	SimulationACHTransferReturnParamsReasonPaymentStopped                                            SimulationACHTransferReturnParamsReason = "payment_stopped"
	SimulationACHTransferReturnParamsReasonNonTransactionAccount                                     SimulationACHTransferReturnParamsReason = "non_transaction_account"
	SimulationACHTransferReturnParamsReasonUncollectedFunds                                          SimulationACHTransferReturnParamsReason = "uncollected_funds"
	SimulationACHTransferReturnParamsReasonRoutingNumberCheckDigitError                              SimulationACHTransferReturnParamsReason = "routing_number_check_digit_error"
	SimulationACHTransferReturnParamsReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete SimulationACHTransferReturnParamsReason = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	SimulationACHTransferReturnParamsReasonAmountFieldError                                          SimulationACHTransferReturnParamsReason = "amount_field_error"
	SimulationACHTransferReturnParamsReasonAuthorizationRevokedByCustomer                            SimulationACHTransferReturnParamsReason = "authorization_revoked_by_customer"
	SimulationACHTransferReturnParamsReasonInvalidACHRoutingNumber                                   SimulationACHTransferReturnParamsReason = "invalid_ach_routing_number"
	SimulationACHTransferReturnParamsReasonFileRecordEditCriteria                                    SimulationACHTransferReturnParamsReason = "file_record_edit_criteria"
	SimulationACHTransferReturnParamsReasonEnrInvalidIndividualName                                  SimulationACHTransferReturnParamsReason = "enr_invalid_individual_name"
	SimulationACHTransferReturnParamsReasonReturnedPerOdfiRequest                                    SimulationACHTransferReturnParamsReason = "returned_per_odfi_request"
	SimulationACHTransferReturnParamsReasonAddendaError                                              SimulationACHTransferReturnParamsReason = "addenda_error"
	SimulationACHTransferReturnParamsReasonLimitedParticipationDfi                                   SimulationACHTransferReturnParamsReason = "limited_participation_dfi"
	SimulationACHTransferReturnParamsReasonIncorrectlyCodedOutboundInternationalPayment              SimulationACHTransferReturnParamsReason = "incorrectly_coded_outbound_international_payment"
	SimulationACHTransferReturnParamsReasonOther                                                     SimulationACHTransferReturnParamsReason = "other"
)
