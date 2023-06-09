// File generated from our OpenAPI spec by Stainless.

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

// SimulationWireTransferService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSimulationWireTransferService]
// method instead.
type SimulationWireTransferService struct {
	Options []option.RequestOption
}

// NewSimulationWireTransferService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationWireTransferService(opts ...option.RequestOption) (r *SimulationWireTransferService) {
	r = &SimulationWireTransferService{}
	r.Options = opts
	return
}

// Simulates an inbound Wire Transfer to your account.
func (r *SimulationWireTransferService) NewInbound(ctx context.Context, body SimulationWireTransferNewInboundParams, opts ...option.RequestOption) (res *WireTransferSimulation, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_wire_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The results of an inbound Wire Transfer simulation.
type WireTransferSimulation struct {
	// If the Wire Transfer attempt succeeds, this will contain the resulting
	// [Transaction](#transactions) object. The Transaction's `source` will be of
	// `category: inbound_wire_transfer`.
	Transaction WireTransferSimulationTransaction `json:"transaction,required"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_wire_transfer_simulation_result`.
	Type WireTransferSimulationType `json:"type,required"`
	JSON wireTransferSimulationJSON
}

// wireTransferSimulationJSON contains the JSON metadata for the struct
// [WireTransferSimulation]
type wireTransferSimulationJSON struct {
	Transaction apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireTransferSimulation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If the Wire Transfer attempt succeeds, this will contain the resulting
// [Transaction](#transactions) object. The Transaction's `source` will be of
// `category: inbound_wire_transfer`.
type WireTransferSimulationTransaction struct {
	// The Transaction identifier.
	ID string `json:"id,required"`
	// The identifier for the Account the Transaction belongs to.
	AccountID string `json:"account_id,required"`
	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transcation's
	// Account.
	Currency WireTransferSimulationTransactionCurrency `json:"currency,required"`
	// An informational message describing this transaction. Use the fields in `source`
	// to get more detailed information. This field appears as the line-item on the
	// statement.
	Description string `json:"description,required"`
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
	JSON wireTransferSimulationTransactionJSON
}

// wireTransferSimulationTransactionJSON contains the JSON metadata for the struct
// [WireTransferSimulationTransaction]
type wireTransferSimulationTransactionJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Amount      apijson.Field
	CreatedAt   apijson.Field
	Currency    apijson.Field
	Description apijson.Field
	RouteID     apijson.Field
	RouteType   apijson.Field
	Source      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireTransferSimulationTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// Transaction's currency. This will match the currency on the Transcation's
// Account.
type WireTransferSimulationTransactionCurrency string

const (
	// Canadian Dollar (CAD)
	WireTransferSimulationTransactionCurrencyCad WireTransferSimulationTransactionCurrency = "CAD"
	// Swiss Franc (CHF)
	WireTransferSimulationTransactionCurrencyChf WireTransferSimulationTransactionCurrency = "CHF"
	// Euro (EUR)
	WireTransferSimulationTransactionCurrencyEur WireTransferSimulationTransactionCurrency = "EUR"
	// British Pound (GBP)
	WireTransferSimulationTransactionCurrencyGbp WireTransferSimulationTransactionCurrency = "GBP"
	// Japanese Yen (JPY)
	WireTransferSimulationTransactionCurrencyJpy WireTransferSimulationTransactionCurrency = "JPY"
	// US Dollar (USD)
	WireTransferSimulationTransactionCurrencyUsd WireTransferSimulationTransactionCurrency = "USD"
)

// The type of the route this Transaction came through.
type WireTransferSimulationTransactionRouteType string

const (
	// An Account Number.
	WireTransferSimulationTransactionRouteTypeAccountNumber WireTransferSimulationTransactionRouteType = "account_number"
	// A Card.
	WireTransferSimulationTransactionRouteTypeCard WireTransferSimulationTransactionRouteType = "card"
)

// This is an object giving more details on the network-level event that caused the
// Transaction. Note that for backwards compatibility reasons, additional
// undocumented keys may appear in this object. These should be treated as
// deprecated and will be removed in the future.
type WireTransferSimulationTransactionSource struct {
	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention WireTransferSimulationTransactionSourceAccountTransferIntention `json:"account_transfer_intention,required,nullable"`
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
	// A Card Revenue Payment object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_revenue_payment`.
	CardRevenuePayment WireTransferSimulationTransactionSourceCardRevenuePayment `json:"card_revenue_payment,required,nullable"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement WireTransferSimulationTransactionSourceCardSettlement `json:"card_settlement,required,nullable"`
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category WireTransferSimulationTransactionSourceCategory `json:"category,required"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance WireTransferSimulationTransactionSourceCheckDepositAcceptance `json:"check_deposit_acceptance,required,nullable"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn WireTransferSimulationTransactionSourceCheckDepositReturn `json:"check_deposit_return,required,nullable"`
	// A Check Transfer Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_transfer_deposit`.
	CheckTransferDeposit WireTransferSimulationTransactionSourceCheckTransferDeposit `json:"check_transfer_deposit,required,nullable"`
	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention WireTransferSimulationTransactionSourceCheckTransferIntention `json:"check_transfer_intention,required,nullable"`
	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection WireTransferSimulationTransactionSourceCheckTransferRejection `json:"check_transfer_rejection,required,nullable"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest `json:"check_transfer_stop_payment_request,required,nullable"`
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
	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment WireTransferSimulationTransactionSourceInboundWireDrawdownPayment `json:"inbound_wire_drawdown_payment,required,nullable"`
	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal `json:"inbound_wire_drawdown_payment_reversal,required,nullable"`
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
	// A Real Time Payments Transfer Acknowledgement object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `real_time_payments_transfer_acknowledgement`.
	RealTimePaymentsTransferAcknowledgement WireTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgement `json:"real_time_payments_transfer_acknowledgement,required,nullable"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds WireTransferSimulationTransactionSourceSampleFunds `json:"sample_funds,required,nullable"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention WireTransferSimulationTransactionSourceWireTransferIntention `json:"wire_transfer_intention,required,nullable"`
	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection WireTransferSimulationTransactionSourceWireTransferRejection `json:"wire_transfer_rejection,required,nullable"`
	JSON                  wireTransferSimulationTransactionSourceJSON
}

// wireTransferSimulationTransactionSourceJSON contains the JSON metadata for the
// struct [WireTransferSimulationTransactionSource]
type wireTransferSimulationTransactionSourceJSON struct {
	AccountTransferIntention                    apijson.Field
	ACHTransferIntention                        apijson.Field
	ACHTransferRejection                        apijson.Field
	ACHTransferReturn                           apijson.Field
	CardDisputeAcceptance                       apijson.Field
	CardRefund                                  apijson.Field
	CardRevenuePayment                          apijson.Field
	CardSettlement                              apijson.Field
	Category                                    apijson.Field
	CheckDepositAcceptance                      apijson.Field
	CheckDepositReturn                          apijson.Field
	CheckTransferDeposit                        apijson.Field
	CheckTransferIntention                      apijson.Field
	CheckTransferRejection                      apijson.Field
	CheckTransferStopPaymentRequest             apijson.Field
	FeePayment                                  apijson.Field
	InboundACHTransfer                          apijson.Field
	InboundCheck                                apijson.Field
	InboundInternationalACHTransfer             apijson.Field
	InboundRealTimePaymentsTransferConfirmation apijson.Field
	InboundWireDrawdownPayment                  apijson.Field
	InboundWireDrawdownPaymentReversal          apijson.Field
	InboundWireReversal                         apijson.Field
	InboundWireTransfer                         apijson.Field
	InterestPayment                             apijson.Field
	InternalSource                              apijson.Field
	RealTimePaymentsTransferAcknowledgement     apijson.Field
	SampleFunds                                 apijson.Field
	WireTransferIntention                       apijson.Field
	WireTransferRejection                       apijson.Field
	raw                                         string
	ExtraFields                                 map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Account Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_intention`.
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
	JSON       wireTransferSimulationTransactionSourceAccountTransferIntentionJSON
}

// wireTransferSimulationTransactionSourceAccountTransferIntentionJSON contains the
// JSON metadata for the struct
// [WireTransferSimulationTransactionSourceAccountTransferIntention]
type wireTransferSimulationTransactionSourceAccountTransferIntentionJSON struct {
	Amount               apijson.Field
	Currency             apijson.Field
	Description          apijson.Field
	DestinationAccountID apijson.Field
	SourceAccountID      apijson.Field
	TransferID           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceAccountTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
type WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency string

const (
	// Canadian Dollar (CAD)
	WireTransferSimulationTransactionSourceAccountTransferIntentionCurrencyCad WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "CAD"
	// Swiss Franc (CHF)
	WireTransferSimulationTransactionSourceAccountTransferIntentionCurrencyChf WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "CHF"
	// Euro (EUR)
	WireTransferSimulationTransactionSourceAccountTransferIntentionCurrencyEur WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "EUR"
	// British Pound (GBP)
	WireTransferSimulationTransactionSourceAccountTransferIntentionCurrencyGbp WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "GBP"
	// Japanese Yen (JPY)
	WireTransferSimulationTransactionSourceAccountTransferIntentionCurrencyJpy WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "JPY"
	// US Dollar (USD)
	WireTransferSimulationTransactionSourceAccountTransferIntentionCurrencyUsd WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "USD"
)

// A ACH Transfer Intention object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_intention`.
type WireTransferSimulationTransactionSourceACHTransferIntention struct {
	AccountNumber string `json:"account_number,required"`
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount              int64  `json:"amount,required"`
	RoutingNumber       string `json:"routing_number,required"`
	StatementDescriptor string `json:"statement_descriptor,required"`
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       wireTransferSimulationTransactionSourceACHTransferIntentionJSON
}

// wireTransferSimulationTransactionSourceACHTransferIntentionJSON contains the
// JSON metadata for the struct
// [WireTransferSimulationTransactionSourceACHTransferIntention]
type wireTransferSimulationTransactionSourceACHTransferIntentionJSON struct {
	AccountNumber       apijson.Field
	Amount              apijson.Field
	RoutingNumber       apijson.Field
	StatementDescriptor apijson.Field
	TransferID          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceACHTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ACH Transfer Rejection object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_rejection`.
type WireTransferSimulationTransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       wireTransferSimulationTransactionSourceACHTransferRejectionJSON
}

// wireTransferSimulationTransactionSourceACHTransferRejectionJSON contains the
// JSON metadata for the struct
// [WireTransferSimulationTransactionSourceACHTransferRejection]
type wireTransferSimulationTransactionSourceACHTransferRejectionJSON struct {
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceACHTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ACH Transfer Return object. This field will be present in the JSON response if
// and only if `category` is equal to `ach_transfer_return`.
type WireTransferSimulationTransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The three character ACH return code, in the range R01 to R85.
	RawReturnReasonCode string `json:"raw_return_reason_code,required"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode `json:"return_reason_code,required"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID string `json:"transaction_id,required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID string `json:"transfer_id,required"`
	JSON       wireTransferSimulationTransactionSourceACHTransferReturnJSON
}

// wireTransferSimulationTransactionSourceACHTransferReturnJSON contains the JSON
// metadata for the struct
// [WireTransferSimulationTransactionSourceACHTransferReturn]
type wireTransferSimulationTransactionSourceACHTransferReturnJSON struct {
	CreatedAt           apijson.Field
	RawReturnReasonCode apijson.Field
	ReturnReasonCode    apijson.Field
	TransactionID       apijson.Field
	TransferID          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Why the ACH Transfer was returned.
type WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode string

const (
	// Code R01. Insufficient funds in the source account.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInsufficientFund WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "insufficient_fund"
	// Code R03. The account does not exist or the receiving bank was unable to locate
	// it.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeNoAccount WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "no_account"
	// Code R02. The account is closed.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAccountClosed WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "account_closed"
	// Code R04. The account number is invalid at the receiving bank.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	// Code R16. The account was frozen per the Office of Foreign Assets Control.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	// Code R23. The receiving bank account refused a credit transfer.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	// Code R05. The receiving bank rejected because of an incorrect Standard Entry
	// Class code.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	// Code R29. The corporate customer reversed the transfer.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	// Code R08. The receiving bank stopped payment on this transfer.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodePaymentStopped WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "payment_stopped"
	// Code R20. The receiving bank account does not perform transfers.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeNonTransactionAccount WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "non_transaction_account"
	// Code R09. The receiving bank account does not have enough available balance for
	// the transfer.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeUncollectedFunds WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "uncollected_funds"
	// Code R28. The routing number is incorrect.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	// Code R10. The customer reversed the transfer.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	// Code R19. The amount field is incorrect or too large.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAmountFieldError WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "amount_field_error"
	// Code R07. The customer who initiated the transfer revoked authorization.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	// Code R13. The routing number is invalid.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	// Code R17. The receiving bank is unable to process a field in the transfer.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeFileRecordEditCriteria WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	// Code R45. The individual name field was invalid.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualName WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	// Code R06. The originating financial institution reversed the transfer.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	// Code R34. The receiving bank's regulatory supervisor has limited their
	// participation.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeLimitedParticipationDfi WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	// Code R85. The outbound international ACH transfer was incorrect.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "incorrectly_coded_outbound_international_payment"
	// Code R12. A rare return reason. The account was sold to another bank.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAccountSoldToAnotherDfi WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "account_sold_to_another_dfi"
	// Code R25. The addenda record is incorrect or missing.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAddendaError WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "addenda_error"
	// Code R15. A rare return reason. The account holder is deceased.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeBeneficiaryOrAccountHolderDeceased WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "beneficiary_or_account_holder_deceased"
	// Code R11. A rare return reason. The customer authorized some payment to the
	// sender, but this payment was not in error.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedNotWithinAuthorizationTerms WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_not_within_authorization_terms"
	// Code R74. A rare return reason. Sent in response to a return that was returned
	// with code `field_error`. The latest return should include the corrected
	// field(s).
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCorrectedReturn WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "corrected_return"
	// Code R24. A rare return reason. The receiving bank received an exact duplicate
	// entry with the same trace number and amount.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeDuplicateEntry WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "duplicate_entry"
	// Code R67. A rare return reason. The return this message refers to was a
	// duplicate.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeDuplicateReturn WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "duplicate_return"
	// Code R47. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrDuplicateEnrollment WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_duplicate_enrollment"
	// Code R43. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidDfiAccountNumber WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_dfi_account_number"
	// Code R44. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualIDNumber WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_id_number"
	// Code R46. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidRepresentativePayeeIndicator WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_representative_payee_indicator"
	// Code R41. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidTransactionCode WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_transaction_code"
	// Code R40. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrReturnOfEnrEntry WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_return_of_enr_entry"
	// Code R42. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrRoutingNumberCheckDigitError WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_routing_number_check_digit_error"
	// Code R84. A rare return reason. The International ACH Transfer cannot be
	// processed by the gateway.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEntryNotProcessedByGateway WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "entry_not_processed_by_gateway"
	// Code R69. A rare return reason. One or more of the fields in the ACH were
	// malformed.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeFieldError WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "field_error"
	// Code R83. A rare return reason. The Foreign receiving bank was unable to settle
	// this ACH transfer.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeForeignReceivingDfiUnableToSettle WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "foreign_receiving_dfi_unable_to_settle"
	// Code R80. A rare return reason. The International ACH Transfer is malformed.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeIatEntryCodingError WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "iat_entry_coding_error"
	// Code R18. A rare return reason. The ACH has an improper effective entry date
	// field.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeImproperEffectiveEntryDate WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "improper_effective_entry_date"
	// Code R39. A rare return reason. The source document related to this ACH, usually
	// an ACH check conversion, was presented to the bank.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeImproperSourceDocumentSourceDocumentPresented WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "improper_source_document_source_document_presented"
	// Code R21. A rare return reason. The Company ID field of the ACH was invalid.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInvalidCompanyID WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "invalid_company_id"
	// Code R82. A rare return reason. The foreign receiving bank identifier for an
	// International ACH Transfer was invalid.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInvalidForeignReceivingDfiIdentification WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "invalid_foreign_receiving_dfi_identification"
	// Code R22. A rare return reason. The Individual ID number field of the ACH was
	// invalid.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInvalidIndividualIDNumber WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "invalid_individual_id_number"
	// Code R53. A rare return reason. Both the Represented Check ("RCK") entry and the
	// original check were presented to the bank.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeItemAndRckEntryPresentedForPayment WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "item_and_rck_entry_presented_for_payment"
	// Code R51. A rare return reason. The Represented Check ("RCK") entry is
	// ineligible.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeItemRelatedToRckEntryIsIneligible WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "item_related_to_rck_entry_is_ineligible"
	// Code R26. A rare return reason. The ACH is missing a required field.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeMandatoryFieldError WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "mandatory_field_error"
	// Code R71. A rare return reason. The receiving bank does not recognize the
	// routing number in a dishonored return entry.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeMisroutedDishonoredReturn WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "misrouted_dishonored_return"
	// Code R61. A rare return reason. The receiving bank does not recognize the
	// routing number in a return entry.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeMisroutedReturn WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "misrouted_return"
	// Code R76. A rare return reason. Sent in response to a return, the bank does not
	// find the errors alleged by the returning bank.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeNoErrorsFound WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "no_errors_found"
	// Code R77. A rare return reason. The receiving bank does not accept the return of
	// the erroneous debit. The funds are not available at the receiving bank.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeNonAcceptanceOfR62DishonoredReturn WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "non_acceptance_of_r62_dishonored_return"
	// Code R81. A rare return reason. The receiving bank does not accept International
	// ACH Transfers.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeNonParticipantInIatProgram WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "non_participant_in_iat_program"
	// Code R31. A rare return reason. A return that has been agreed to be accepted by
	// the receiving bank, despite falling outside of the usual return timeframe.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodePermissibleReturnEntry WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "permissible_return_entry"
	// Code R70. A rare return reason. The receiving bank had not approved this return.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodePermissibleReturnEntryNotAccepted WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "permissible_return_entry_not_accepted"
	// Code R32. A rare return reason. The receiving bank could not settle this
	// transaction.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeRdfiNonSettlement WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "rdfi_non_settlement"
	// Code R30. A rare return reason. The receiving bank does not accept Check
	// Truncation ACH transfers.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeRdfiParticipantInCheckTruncationProgram WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "rdfi_participant_in_check_truncation_program"
	// Code R14. A rare return reason. The payee is deceased.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	// Code R75. A rare return reason. The originating bank disputes that an earlier
	// `duplicate_entry` return was actually a duplicate.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeReturnNotADuplicate WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "return_not_a_duplicate"
	// Code R62. A rare return reason. The originating bank made a mistake earlier and
	// this return corrects it.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeReturnOfErroneousOrReversingDebit WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "return_of_erroneous_or_reversing_debit"
	// Code R36. A rare return reason. Return of a malformed credit entry.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeReturnOfImproperCreditEntry WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "return_of_improper_credit_entry"
	// Code R35. A rare return reason. Return of a malformed debit entry.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeReturnOfImproperDebitEntry WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "return_of_improper_debit_entry"
	// Code R33. A rare return reason. Return of a Destroyed Check ("XKC") entry.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeReturnOfXckEntry WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "return_of_xck_entry"
	// Code R37. A rare return reason. The source document related to this ACH, usually
	// an ACH check conversion, was presented to the bank.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeSourceDocumentPresentedForPayment WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "source_document_presented_for_payment"
	// Code R50. A rare return reason. State law prevents the bank from accepting the
	// Represented Check ("RCK") entry.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeStateLawAffectingRckAcceptance WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "state_law_affecting_rck_acceptance"
	// Code R52. A rare return reason. A stop payment was issued on a Represented Check
	// ("RCK") entry.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeStopPaymentOnItemRelatedToRckEntry WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "stop_payment_on_item_related_to_rck_entry"
	// Code R38. A rare return reason. The source attached to the ACH, usually an ACH
	// check conversion, includes a stop payment.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeStopPaymentOnSourceDocument WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "stop_payment_on_source_document"
	// Code R73. A rare return reason. The bank receiving an `untimely_return` believes
	// it was on time.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeTimelyOriginalReturn WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "timely_original_return"
	// Code R27. A rare return reason. An ACH Return's trace number does not match an
	// originated ACH.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeTraceNumberError WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "trace_number_error"
	// Code R72. A rare return reason. The dishonored return was sent too late.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeUntimelyDishonoredReturn WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "untimely_dishonored_return"
	// Code R68. A rare return reason. The return was sent too late.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeUntimelyReturn WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "untimely_return"
)

// A Card Dispute Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `card_dispute_acceptance`.
type WireTransferSimulationTransactionSourceCardDisputeAcceptance struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was accepted.
	AcceptedAt time.Time `json:"accepted_at,required" format:"date-time"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID string `json:"card_dispute_id,required"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID string `json:"transaction_id,required"`
	JSON          wireTransferSimulationTransactionSourceCardDisputeAcceptanceJSON
}

// wireTransferSimulationTransactionSourceCardDisputeAcceptanceJSON contains the
// JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCardDisputeAcceptance]
type wireTransferSimulationTransactionSourceCardDisputeAcceptanceJSON struct {
	AcceptedAt    apijson.Field
	CardDisputeID apijson.Field
	TransactionID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCardDisputeAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Card Refund object. This field will be present in the JSON response if and
// only if `category` is equal to `card_refund`.
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
	// The 4-digit MCC describing the merchant's business.
	MerchantCategoryCode string `json:"merchant_category_code,required"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required"`
	// The name of the merchant.
	MerchantName string `json:"merchant_name,required,nullable"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state,required,nullable"`
	// The identifier of the Transaction associated with this Transaction.
	TransactionID string `json:"transaction_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type WireTransferSimulationTransactionSourceCardRefundType `json:"type,required"`
	JSON wireTransferSimulationTransactionSourceCardRefundJSON
}

// wireTransferSimulationTransactionSourceCardRefundJSON contains the JSON metadata
// for the struct [WireTransferSimulationTransactionSourceCardRefund]
type wireTransferSimulationTransactionSourceCardRefundJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	Currency             apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantName         apijson.Field
	MerchantState        apijson.Field
	TransactionID        apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCardRefund) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type WireTransferSimulationTransactionSourceCardRefundCurrency string

const (
	// Canadian Dollar (CAD)
	WireTransferSimulationTransactionSourceCardRefundCurrencyCad WireTransferSimulationTransactionSourceCardRefundCurrency = "CAD"
	// Swiss Franc (CHF)
	WireTransferSimulationTransactionSourceCardRefundCurrencyChf WireTransferSimulationTransactionSourceCardRefundCurrency = "CHF"
	// Euro (EUR)
	WireTransferSimulationTransactionSourceCardRefundCurrencyEur WireTransferSimulationTransactionSourceCardRefundCurrency = "EUR"
	// British Pound (GBP)
	WireTransferSimulationTransactionSourceCardRefundCurrencyGbp WireTransferSimulationTransactionSourceCardRefundCurrency = "GBP"
	// Japanese Yen (JPY)
	WireTransferSimulationTransactionSourceCardRefundCurrencyJpy WireTransferSimulationTransactionSourceCardRefundCurrency = "JPY"
	// US Dollar (USD)
	WireTransferSimulationTransactionSourceCardRefundCurrencyUsd WireTransferSimulationTransactionSourceCardRefundCurrency = "USD"
)

// A constant representing the object's type. For this resource it will always be
// `card_refund`.
type WireTransferSimulationTransactionSourceCardRefundType string

const (
	WireTransferSimulationTransactionSourceCardRefundTypeCardRefund WireTransferSimulationTransactionSourceCardRefundType = "card_refund"
)

// A Card Revenue Payment object. This field will be present in the JSON response
// if and only if `category` is equal to `card_revenue_payment`.
type WireTransferSimulationTransactionSourceCardRevenuePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency WireTransferSimulationTransactionSourceCardRevenuePaymentCurrency `json:"currency,required"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time `json:"period_start,required" format:"date-time"`
	// The account the card belonged to.
	TransactedOnAccountID string `json:"transacted_on_account_id,required,nullable"`
	JSON                  wireTransferSimulationTransactionSourceCardRevenuePaymentJSON
}

// wireTransferSimulationTransactionSourceCardRevenuePaymentJSON contains the JSON
// metadata for the struct
// [WireTransferSimulationTransactionSourceCardRevenuePayment]
type wireTransferSimulationTransactionSourceCardRevenuePaymentJSON struct {
	Amount                apijson.Field
	Currency              apijson.Field
	PeriodEnd             apijson.Field
	PeriodStart           apijson.Field
	TransactedOnAccountID apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCardRevenuePayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type WireTransferSimulationTransactionSourceCardRevenuePaymentCurrency string

const (
	// Canadian Dollar (CAD)
	WireTransferSimulationTransactionSourceCardRevenuePaymentCurrencyCad WireTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "CAD"
	// Swiss Franc (CHF)
	WireTransferSimulationTransactionSourceCardRevenuePaymentCurrencyChf WireTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "CHF"
	// Euro (EUR)
	WireTransferSimulationTransactionSourceCardRevenuePaymentCurrencyEur WireTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "EUR"
	// British Pound (GBP)
	WireTransferSimulationTransactionSourceCardRevenuePaymentCurrencyGbp WireTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "GBP"
	// Japanese Yen (JPY)
	WireTransferSimulationTransactionSourceCardRevenuePaymentCurrencyJpy WireTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "JPY"
	// US Dollar (USD)
	WireTransferSimulationTransactionSourceCardRevenuePaymentCurrencyUsd WireTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "USD"
)

// A Card Settlement object. This field will be present in the JSON response if and
// only if `category` is equal to `card_settlement`.
type WireTransferSimulationTransactionSourceCardSettlement struct {
	// The Card Settlement identifier.
	ID string `json:"id,required"`
	// The amount in the minor unit of the transaction's settlement currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The Card Authorization that was created prior to this Card Settlement, if on
	// exists.
	CardAuthorization string `json:"card_authorization,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's settlement currency.
	Currency WireTransferSimulationTransactionSourceCardSettlementCurrency `json:"currency,required"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required,nullable"`
	// The 4-digit MCC describing the merchant's business.
	MerchantCategoryCode string `json:"merchant_category_code,required"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required"`
	// The name of the merchant.
	MerchantName string `json:"merchant_name,required,nullable"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state,required,nullable"`
	// The identifier of the Pending Transaction associated with this Transaction.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// The amount in the minor unit of the transaction's presentment currency.
	PresentmentAmount int64 `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's presentment currency.
	PresentmentCurrency string `json:"presentment_currency,required"`
	// The identifier of the Transaction associated with this Transaction.
	TransactionID string `json:"transaction_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_settlement`.
	Type WireTransferSimulationTransactionSourceCardSettlementType `json:"type,required"`
	JSON wireTransferSimulationTransactionSourceCardSettlementJSON
}

// wireTransferSimulationTransactionSourceCardSettlementJSON contains the JSON
// metadata for the struct [WireTransferSimulationTransactionSourceCardSettlement]
type wireTransferSimulationTransactionSourceCardSettlementJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CardAuthorization    apijson.Field
	Currency             apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantName         apijson.Field
	MerchantState        apijson.Field
	PendingTransactionID apijson.Field
	PresentmentAmount    apijson.Field
	PresentmentCurrency  apijson.Field
	TransactionID        apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCardSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's settlement currency.
type WireTransferSimulationTransactionSourceCardSettlementCurrency string

const (
	// Canadian Dollar (CAD)
	WireTransferSimulationTransactionSourceCardSettlementCurrencyCad WireTransferSimulationTransactionSourceCardSettlementCurrency = "CAD"
	// Swiss Franc (CHF)
	WireTransferSimulationTransactionSourceCardSettlementCurrencyChf WireTransferSimulationTransactionSourceCardSettlementCurrency = "CHF"
	// Euro (EUR)
	WireTransferSimulationTransactionSourceCardSettlementCurrencyEur WireTransferSimulationTransactionSourceCardSettlementCurrency = "EUR"
	// British Pound (GBP)
	WireTransferSimulationTransactionSourceCardSettlementCurrencyGbp WireTransferSimulationTransactionSourceCardSettlementCurrency = "GBP"
	// Japanese Yen (JPY)
	WireTransferSimulationTransactionSourceCardSettlementCurrencyJpy WireTransferSimulationTransactionSourceCardSettlementCurrency = "JPY"
	// US Dollar (USD)
	WireTransferSimulationTransactionSourceCardSettlementCurrencyUsd WireTransferSimulationTransactionSourceCardSettlementCurrency = "USD"
)

// A constant representing the object's type. For this resource it will always be
// `card_settlement`.
type WireTransferSimulationTransactionSourceCardSettlementType string

const (
	WireTransferSimulationTransactionSourceCardSettlementTypeCardSettlement WireTransferSimulationTransactionSourceCardSettlementType = "card_settlement"
)

// The type of transaction that took place. We may add additional possible values
// for this enum over time; your application should be able to handle such
// additions gracefully.
type WireTransferSimulationTransactionSourceCategory string

const (
	// The Transaction was created by a Account Transfer Intention object. Details will
	// be under the `account_transfer_intention` object.
	WireTransferSimulationTransactionSourceCategoryAccountTransferIntention WireTransferSimulationTransactionSourceCategory = "account_transfer_intention"
	// The Transaction was created by a ACH Transfer Intention object. Details will be
	// under the `ach_transfer_intention` object.
	WireTransferSimulationTransactionSourceCategoryACHTransferIntention WireTransferSimulationTransactionSourceCategory = "ach_transfer_intention"
	// The Transaction was created by a ACH Transfer Rejection object. Details will be
	// under the `ach_transfer_rejection` object.
	WireTransferSimulationTransactionSourceCategoryACHTransferRejection WireTransferSimulationTransactionSourceCategory = "ach_transfer_rejection"
	// The Transaction was created by a ACH Transfer Return object. Details will be
	// under the `ach_transfer_return` object.
	WireTransferSimulationTransactionSourceCategoryACHTransferReturn WireTransferSimulationTransactionSourceCategory = "ach_transfer_return"
	// The Transaction was created by a Card Dispute Acceptance object. Details will be
	// under the `card_dispute_acceptance` object.
	WireTransferSimulationTransactionSourceCategoryCardDisputeAcceptance WireTransferSimulationTransactionSourceCategory = "card_dispute_acceptance"
	// The Transaction was created by a Card Refund object. Details will be under the
	// `card_refund` object.
	WireTransferSimulationTransactionSourceCategoryCardRefund WireTransferSimulationTransactionSourceCategory = "card_refund"
	// The Transaction was created by a Card Revenue Payment object. Details will be
	// under the `card_revenue_payment` object.
	WireTransferSimulationTransactionSourceCategoryCardRevenuePayment WireTransferSimulationTransactionSourceCategory = "card_revenue_payment"
	// The Transaction was created by a Card Settlement object. Details will be under
	// the `card_settlement` object.
	WireTransferSimulationTransactionSourceCategoryCardSettlement WireTransferSimulationTransactionSourceCategory = "card_settlement"
	// The Transaction was created by a Check Deposit Acceptance object. Details will
	// be under the `check_deposit_acceptance` object.
	WireTransferSimulationTransactionSourceCategoryCheckDepositAcceptance WireTransferSimulationTransactionSourceCategory = "check_deposit_acceptance"
	// The Transaction was created by a Check Deposit Return object. Details will be
	// under the `check_deposit_return` object.
	WireTransferSimulationTransactionSourceCategoryCheckDepositReturn WireTransferSimulationTransactionSourceCategory = "check_deposit_return"
	// The Transaction was created by a Check Transfer Deposit object. Details will be
	// under the `check_transfer_deposit` object.
	WireTransferSimulationTransactionSourceCategoryCheckTransferDeposit WireTransferSimulationTransactionSourceCategory = "check_transfer_deposit"
	// The Transaction was created by a Check Transfer Intention object. Details will
	// be under the `check_transfer_intention` object.
	WireTransferSimulationTransactionSourceCategoryCheckTransferIntention WireTransferSimulationTransactionSourceCategory = "check_transfer_intention"
	// The Transaction was created by a Check Transfer Rejection object. Details will
	// be under the `check_transfer_rejection` object.
	WireTransferSimulationTransactionSourceCategoryCheckTransferRejection WireTransferSimulationTransactionSourceCategory = "check_transfer_rejection"
	// The Transaction was created by a Check Transfer Stop Payment Request object.
	// Details will be under the `check_transfer_stop_payment_request` object.
	WireTransferSimulationTransactionSourceCategoryCheckTransferStopPaymentRequest WireTransferSimulationTransactionSourceCategory = "check_transfer_stop_payment_request"
	// The Transaction was created by a Fee Payment object. Details will be under the
	// `fee_payment` object.
	WireTransferSimulationTransactionSourceCategoryFeePayment WireTransferSimulationTransactionSourceCategory = "fee_payment"
	// The Transaction was created by a Inbound ACH Transfer object. Details will be
	// under the `inbound_ach_transfer` object.
	WireTransferSimulationTransactionSourceCategoryInboundACHTransfer WireTransferSimulationTransactionSourceCategory = "inbound_ach_transfer"
	// The Transaction was created by a Inbound ACH Transfer Return Intention object.
	// Details will be under the `inbound_ach_transfer_return_intention` object.
	WireTransferSimulationTransactionSourceCategoryInboundACHTransferReturnIntention WireTransferSimulationTransactionSourceCategory = "inbound_ach_transfer_return_intention"
	// The Transaction was created by a Inbound Check object. Details will be under the
	// `inbound_check` object.
	WireTransferSimulationTransactionSourceCategoryInboundCheck WireTransferSimulationTransactionSourceCategory = "inbound_check"
	// The Transaction was created by a Inbound International ACH Transfer object.
	// Details will be under the `inbound_international_ach_transfer` object.
	WireTransferSimulationTransactionSourceCategoryInboundInternationalACHTransfer WireTransferSimulationTransactionSourceCategory = "inbound_international_ach_transfer"
	// The Transaction was created by a Inbound Real Time Payments Transfer
	// Confirmation object. Details will be under the
	// `inbound_real_time_payments_transfer_confirmation` object.
	WireTransferSimulationTransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation WireTransferSimulationTransactionSourceCategory = "inbound_real_time_payments_transfer_confirmation"
	// The Transaction was created by a Inbound Wire Drawdown Payment object. Details
	// will be under the `inbound_wire_drawdown_payment` object.
	WireTransferSimulationTransactionSourceCategoryInboundWireDrawdownPayment WireTransferSimulationTransactionSourceCategory = "inbound_wire_drawdown_payment"
	// The Transaction was created by a Inbound Wire Drawdown Payment Reversal object.
	// Details will be under the `inbound_wire_drawdown_payment_reversal` object.
	WireTransferSimulationTransactionSourceCategoryInboundWireDrawdownPaymentReversal WireTransferSimulationTransactionSourceCategory = "inbound_wire_drawdown_payment_reversal"
	// The Transaction was created by a Inbound Wire Reversal object. Details will be
	// under the `inbound_wire_reversal` object.
	WireTransferSimulationTransactionSourceCategoryInboundWireReversal WireTransferSimulationTransactionSourceCategory = "inbound_wire_reversal"
	// The Transaction was created by a Inbound Wire Transfer object. Details will be
	// under the `inbound_wire_transfer` object.
	WireTransferSimulationTransactionSourceCategoryInboundWireTransfer WireTransferSimulationTransactionSourceCategory = "inbound_wire_transfer"
	// The Transaction was created by a Interest Payment object. Details will be under
	// the `interest_payment` object.
	WireTransferSimulationTransactionSourceCategoryInterestPayment WireTransferSimulationTransactionSourceCategory = "interest_payment"
	// The Transaction was created by a Internal Source object. Details will be under
	// the `internal_source` object.
	WireTransferSimulationTransactionSourceCategoryInternalSource WireTransferSimulationTransactionSourceCategory = "internal_source"
	// The Transaction was created by a Real Time Payments Transfer Acknowledgement
	// object. Details will be under the `real_time_payments_transfer_acknowledgement`
	// object.
	WireTransferSimulationTransactionSourceCategoryRealTimePaymentsTransferAcknowledgement WireTransferSimulationTransactionSourceCategory = "real_time_payments_transfer_acknowledgement"
	// The Transaction was created by a Sample Funds object. Details will be under the
	// `sample_funds` object.
	WireTransferSimulationTransactionSourceCategorySampleFunds WireTransferSimulationTransactionSourceCategory = "sample_funds"
	// The Transaction was created by a Wire Transfer Intention object. Details will be
	// under the `wire_transfer_intention` object.
	WireTransferSimulationTransactionSourceCategoryWireTransferIntention WireTransferSimulationTransactionSourceCategory = "wire_transfer_intention"
	// The Transaction was created by a Wire Transfer Rejection object. Details will be
	// under the `wire_transfer_rejection` object.
	WireTransferSimulationTransactionSourceCategoryWireTransferRejection WireTransferSimulationTransactionSourceCategory = "wire_transfer_rejection"
	// The Transaction was made for an undocumented or deprecated reason.
	WireTransferSimulationTransactionSourceCategoryOther WireTransferSimulationTransactionSourceCategory = "other"
)

// A Check Deposit Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_acceptance`.
type WireTransferSimulationTransactionSourceCheckDepositAcceptance struct {
	// The account number printed on the check.
	AccountNumber string `json:"account_number,required"`
	// The amount to be deposited in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// An additional line of metadata printed on the check. This typically includes the
	// check number for business checks.
	AuxiliaryOnUs string `json:"auxiliary_on_us,required,nullable"`
	// The ID of the Check Deposit that was accepted.
	CheckDepositID string `json:"check_deposit_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency `json:"currency,required"`
	// The routing number printed on the check.
	RoutingNumber string `json:"routing_number,required"`
	// The check serial number, if present, for consumer checks. For business checks,
	// the serial number is usually in the `auxiliary_on_us` field.
	SerialNumber string `json:"serial_number,required,nullable"`
	JSON         wireTransferSimulationTransactionSourceCheckDepositAcceptanceJSON
}

// wireTransferSimulationTransactionSourceCheckDepositAcceptanceJSON contains the
// JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCheckDepositAcceptance]
type wireTransferSimulationTransactionSourceCheckDepositAcceptanceJSON struct {
	AccountNumber  apijson.Field
	Amount         apijson.Field
	AuxiliaryOnUs  apijson.Field
	CheckDepositID apijson.Field
	Currency       apijson.Field
	RoutingNumber  apijson.Field
	SerialNumber   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCheckDepositAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency string

const (
	// Canadian Dollar (CAD)
	WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyCad WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "CAD"
	// Swiss Franc (CHF)
	WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyChf WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "CHF"
	// Euro (EUR)
	WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyEur WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "EUR"
	// British Pound (GBP)
	WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyGbp WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "GBP"
	// Japanese Yen (JPY)
	WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyJpy WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "JPY"
	// US Dollar (USD)
	WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyUsd WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "USD"
)

// A Check Deposit Return object. This field will be present in the JSON response
// if and only if `category` is equal to `check_deposit_return`.
type WireTransferSimulationTransactionSourceCheckDepositReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID string `json:"check_deposit_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency     WireTransferSimulationTransactionSourceCheckDepositReturnCurrency     `json:"currency,required"`
	ReturnReason WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason `json:"return_reason,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID string `json:"transaction_id,required"`
	JSON          wireTransferSimulationTransactionSourceCheckDepositReturnJSON
}

// wireTransferSimulationTransactionSourceCheckDepositReturnJSON contains the JSON
// metadata for the struct
// [WireTransferSimulationTransactionSourceCheckDepositReturn]
type wireTransferSimulationTransactionSourceCheckDepositReturnJSON struct {
	Amount         apijson.Field
	CheckDepositID apijson.Field
	Currency       apijson.Field
	ReturnReason   apijson.Field
	ReturnedAt     apijson.Field
	TransactionID  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCheckDepositReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type WireTransferSimulationTransactionSourceCheckDepositReturnCurrency string

const (
	// Canadian Dollar (CAD)
	WireTransferSimulationTransactionSourceCheckDepositReturnCurrencyCad WireTransferSimulationTransactionSourceCheckDepositReturnCurrency = "CAD"
	// Swiss Franc (CHF)
	WireTransferSimulationTransactionSourceCheckDepositReturnCurrencyChf WireTransferSimulationTransactionSourceCheckDepositReturnCurrency = "CHF"
	// Euro (EUR)
	WireTransferSimulationTransactionSourceCheckDepositReturnCurrencyEur WireTransferSimulationTransactionSourceCheckDepositReturnCurrency = "EUR"
	// British Pound (GBP)
	WireTransferSimulationTransactionSourceCheckDepositReturnCurrencyGbp WireTransferSimulationTransactionSourceCheckDepositReturnCurrency = "GBP"
	// Japanese Yen (JPY)
	WireTransferSimulationTransactionSourceCheckDepositReturnCurrencyJpy WireTransferSimulationTransactionSourceCheckDepositReturnCurrency = "JPY"
	// US Dollar (USD)
	WireTransferSimulationTransactionSourceCheckDepositReturnCurrencyUsd WireTransferSimulationTransactionSourceCheckDepositReturnCurrency = "USD"
)

type WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason string

const (
	// The check doesn't allow ACH conversion.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonACHConversionNotSupported WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "ach_conversion_not_supported"
	// The account is closed.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonClosedAccount WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "closed_account"
	// The check has already been deposited.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonDuplicateSubmission WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "duplicate_submission"
	// Insufficient funds
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonInsufficientFunds WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "insufficient_funds"
	// No account was found matching the check details.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonNoAccount WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "no_account"
	// The check was not authorized.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonNotAuthorized WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "not_authorized"
	// The check is too old.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonStaleDated WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "stale_dated"
	// The payment has been stopped by the account holder.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonStopPayment WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "stop_payment"
	// The reason for the return is unknown.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonUnknownReason WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "unknown_reason"
	// The image doesn't match the details submitted.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonUnmatchedDetails WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "unmatched_details"
	// The image could not be read.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonUnreadableImage WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "unreadable_image"
	// The check endorsement was irregular.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonEndorsementIrregular WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "endorsement_irregular"
)

// A Check Transfer Deposit object. This field will be present in the JSON response
// if and only if `category` is equal to `check_transfer_deposit`.
type WireTransferSimulationTransactionSourceCheckTransferDeposit struct {
	// The identifier of the API File object containing an image of the back of the
	// deposited check.
	BackImageFileID string `json:"back_image_file_id,required,nullable"`
	// When the check was deposited.
	DepositedAt time.Time `json:"deposited_at,required" format:"date-time"`
	// The identifier of the API File object containing an image of the front of the
	// deposited check.
	FrontImageFileID string `json:"front_image_file_id,required,nullable"`
	// The identifier of the Transaction object created when the check was deposited.
	TransactionID string `json:"transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_deposit`.
	Type WireTransferSimulationTransactionSourceCheckTransferDepositType `json:"type,required"`
	JSON wireTransferSimulationTransactionSourceCheckTransferDepositJSON
}

// wireTransferSimulationTransactionSourceCheckTransferDepositJSON contains the
// JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCheckTransferDeposit]
type wireTransferSimulationTransactionSourceCheckTransferDepositJSON struct {
	BackImageFileID  apijson.Field
	DepositedAt      apijson.Field
	FrontImageFileID apijson.Field
	TransactionID    apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCheckTransferDeposit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_deposit`.
type WireTransferSimulationTransactionSourceCheckTransferDepositType string

const (
	WireTransferSimulationTransactionSourceCheckTransferDepositTypeCheckTransferDeposit WireTransferSimulationTransactionSourceCheckTransferDepositType = "check_transfer_deposit"
)

// A Check Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_intention`.
type WireTransferSimulationTransactionSourceCheckTransferIntention struct {
	// The city of the check's destination.
	AddressCity string `json:"address_city,required,nullable"`
	// The street address of the check's destination.
	AddressLine1 string `json:"address_line1,required,nullable"`
	// The second line of the address of the check's destination.
	AddressLine2 string `json:"address_line2,required,nullable"`
	// The state of the check's destination.
	AddressState string `json:"address_state,required,nullable"`
	// The postal code of the check's destination.
	AddressZip string `json:"address_zip,required,nullable"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency `json:"currency,required"`
	// The name that will be printed on the check.
	RecipientName string `json:"recipient_name,required,nullable"`
	// The identifier of the Check Transfer with which this is associated.
	TransferID string `json:"transfer_id,required"`
	JSON       wireTransferSimulationTransactionSourceCheckTransferIntentionJSON
}

// wireTransferSimulationTransactionSourceCheckTransferIntentionJSON contains the
// JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCheckTransferIntention]
type wireTransferSimulationTransactionSourceCheckTransferIntentionJSON struct {
	AddressCity   apijson.Field
	AddressLine1  apijson.Field
	AddressLine2  apijson.Field
	AddressState  apijson.Field
	AddressZip    apijson.Field
	Amount        apijson.Field
	Currency      apijson.Field
	RecipientName apijson.Field
	TransferID    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCheckTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
type WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency string

const (
	// Canadian Dollar (CAD)
	WireTransferSimulationTransactionSourceCheckTransferIntentionCurrencyCad WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "CAD"
	// Swiss Franc (CHF)
	WireTransferSimulationTransactionSourceCheckTransferIntentionCurrencyChf WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "CHF"
	// Euro (EUR)
	WireTransferSimulationTransactionSourceCheckTransferIntentionCurrencyEur WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "EUR"
	// British Pound (GBP)
	WireTransferSimulationTransactionSourceCheckTransferIntentionCurrencyGbp WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "GBP"
	// Japanese Yen (JPY)
	WireTransferSimulationTransactionSourceCheckTransferIntentionCurrencyJpy WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "JPY"
	// US Dollar (USD)
	WireTransferSimulationTransactionSourceCheckTransferIntentionCurrencyUsd WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "USD"
)

// A Check Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_rejection`.
type WireTransferSimulationTransactionSourceCheckTransferRejection struct {
	// The identifier of the Check Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       wireTransferSimulationTransactionSourceCheckTransferRejectionJSON
}

// wireTransferSimulationTransactionSourceCheckTransferRejectionJSON contains the
// JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCheckTransferRejection]
type wireTransferSimulationTransactionSourceCheckTransferRejectionJSON struct {
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCheckTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Check Transfer Stop Payment Request object. This field will be present in the
// JSON response if and only if `category` is equal to
// `check_transfer_stop_payment_request`.
type WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest struct {
	// The reason why this transfer was stopped.
	Reason WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestReason `json:"reason,required"`
	// The time the stop-payment was requested.
	RequestedAt time.Time `json:"requested_at,required" format:"date-time"`
	// The ID of the check transfer that was stopped.
	TransferID string `json:"transfer_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType `json:"type,required"`
	JSON wireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestJSON
}

// wireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest]
type wireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestJSON struct {
	Reason      apijson.Field
	RequestedAt apijson.Field
	TransferID  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The reason why this transfer was stopped.
type WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestReason string

const (
	// The check could not be delivered.
	WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestReasonMailDeliveryFailed WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestReason = "mail_delivery_failed"
	// The check was stopped for another reason.
	WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestReasonUnknown WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestReason = "unknown"
)

// A constant representing the object's type. For this resource it will always be
// `check_transfer_stop_payment_request`.
type WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType string

const (
	WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

// A Fee Payment object. This field will be present in the JSON response if and
// only if `category` is equal to `fee_payment`.
type WireTransferSimulationTransactionSourceFeePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency WireTransferSimulationTransactionSourceFeePaymentCurrency `json:"currency,required"`
	JSON     wireTransferSimulationTransactionSourceFeePaymentJSON
}

// wireTransferSimulationTransactionSourceFeePaymentJSON contains the JSON metadata
// for the struct [WireTransferSimulationTransactionSourceFeePayment]
type wireTransferSimulationTransactionSourceFeePaymentJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceFeePayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type WireTransferSimulationTransactionSourceFeePaymentCurrency string

const (
	// Canadian Dollar (CAD)
	WireTransferSimulationTransactionSourceFeePaymentCurrencyCad WireTransferSimulationTransactionSourceFeePaymentCurrency = "CAD"
	// Swiss Franc (CHF)
	WireTransferSimulationTransactionSourceFeePaymentCurrencyChf WireTransferSimulationTransactionSourceFeePaymentCurrency = "CHF"
	// Euro (EUR)
	WireTransferSimulationTransactionSourceFeePaymentCurrencyEur WireTransferSimulationTransactionSourceFeePaymentCurrency = "EUR"
	// British Pound (GBP)
	WireTransferSimulationTransactionSourceFeePaymentCurrencyGbp WireTransferSimulationTransactionSourceFeePaymentCurrency = "GBP"
	// Japanese Yen (JPY)
	WireTransferSimulationTransactionSourceFeePaymentCurrencyJpy WireTransferSimulationTransactionSourceFeePaymentCurrency = "JPY"
	// US Dollar (USD)
	WireTransferSimulationTransactionSourceFeePaymentCurrencyUsd WireTransferSimulationTransactionSourceFeePaymentCurrency = "USD"
)

// A Inbound ACH Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_ach_transfer`.
type WireTransferSimulationTransactionSourceInboundACHTransfer struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount                             int64  `json:"amount,required"`
	OriginatorCompanyDescriptiveDate   string `json:"originator_company_descriptive_date,required,nullable"`
	OriginatorCompanyDiscretionaryData string `json:"originator_company_discretionary_data,required,nullable"`
	OriginatorCompanyEntryDescription  string `json:"originator_company_entry_description,required"`
	OriginatorCompanyID                string `json:"originator_company_id,required"`
	OriginatorCompanyName              string `json:"originator_company_name,required"`
	ReceiverIDNumber                   string `json:"receiver_id_number,required,nullable"`
	ReceiverName                       string `json:"receiver_name,required,nullable"`
	TraceNumber                        string `json:"trace_number,required"`
	JSON                               wireTransferSimulationTransactionSourceInboundACHTransferJSON
}

// wireTransferSimulationTransactionSourceInboundACHTransferJSON contains the JSON
// metadata for the struct
// [WireTransferSimulationTransactionSourceInboundACHTransfer]
type wireTransferSimulationTransactionSourceInboundACHTransferJSON struct {
	Amount                             apijson.Field
	OriginatorCompanyDescriptiveDate   apijson.Field
	OriginatorCompanyDiscretionaryData apijson.Field
	OriginatorCompanyEntryDescription  apijson.Field
	OriginatorCompanyID                apijson.Field
	OriginatorCompanyName              apijson.Field
	ReceiverIDNumber                   apijson.Field
	ReceiverName                       apijson.Field
	TraceNumber                        apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceInboundACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Inbound Check object. This field will be present in the JSON response if and
// only if `category` is equal to `inbound_check`.
type WireTransferSimulationTransactionSourceInboundCheck struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount                int64  `json:"amount,required"`
	CheckFrontImageFileID string `json:"check_front_image_file_id,required,nullable"`
	CheckNumber           string `json:"check_number,required,nullable"`
	CheckRearImageFileID  string `json:"check_rear_image_file_id,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency WireTransferSimulationTransactionSourceInboundCheckCurrency `json:"currency,required"`
	JSON     wireTransferSimulationTransactionSourceInboundCheckJSON
}

// wireTransferSimulationTransactionSourceInboundCheckJSON contains the JSON
// metadata for the struct [WireTransferSimulationTransactionSourceInboundCheck]
type wireTransferSimulationTransactionSourceInboundCheckJSON struct {
	Amount                apijson.Field
	CheckFrontImageFileID apijson.Field
	CheckNumber           apijson.Field
	CheckRearImageFileID  apijson.Field
	Currency              apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceInboundCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type WireTransferSimulationTransactionSourceInboundCheckCurrency string

const (
	// Canadian Dollar (CAD)
	WireTransferSimulationTransactionSourceInboundCheckCurrencyCad WireTransferSimulationTransactionSourceInboundCheckCurrency = "CAD"
	// Swiss Franc (CHF)
	WireTransferSimulationTransactionSourceInboundCheckCurrencyChf WireTransferSimulationTransactionSourceInboundCheckCurrency = "CHF"
	// Euro (EUR)
	WireTransferSimulationTransactionSourceInboundCheckCurrencyEur WireTransferSimulationTransactionSourceInboundCheckCurrency = "EUR"
	// British Pound (GBP)
	WireTransferSimulationTransactionSourceInboundCheckCurrencyGbp WireTransferSimulationTransactionSourceInboundCheckCurrency = "GBP"
	// Japanese Yen (JPY)
	WireTransferSimulationTransactionSourceInboundCheckCurrencyJpy WireTransferSimulationTransactionSourceInboundCheckCurrency = "JPY"
	// US Dollar (USD)
	WireTransferSimulationTransactionSourceInboundCheckCurrencyUsd WireTransferSimulationTransactionSourceInboundCheckCurrency = "USD"
)

// A Inbound International ACH Transfer object. This field will be present in the
// JSON response if and only if `category` is equal to
// `inbound_international_ach_transfer`.
type WireTransferSimulationTransactionSourceInboundInternationalACHTransfer struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount                                                 int64  `json:"amount,required"`
	DestinationCountryCode                                 string `json:"destination_country_code,required"`
	DestinationCurrencyCode                                string `json:"destination_currency_code,required"`
	ForeignExchangeIndicator                               string `json:"foreign_exchange_indicator,required"`
	ForeignExchangeReference                               string `json:"foreign_exchange_reference,required,nullable"`
	ForeignExchangeReferenceIndicator                      string `json:"foreign_exchange_reference_indicator,required"`
	ForeignPaymentAmount                                   int64  `json:"foreign_payment_amount,required"`
	ForeignTraceNumber                                     string `json:"foreign_trace_number,required,nullable"`
	InternationalTransactionTypeCode                       string `json:"international_transaction_type_code,required"`
	OriginatingCurrencyCode                                string `json:"originating_currency_code,required"`
	OriginatingDepositoryFinancialInstitutionBranchCountry string `json:"originating_depository_financial_institution_branch_country,required"`
	OriginatingDepositoryFinancialInstitutionID            string `json:"originating_depository_financial_institution_id,required"`
	OriginatingDepositoryFinancialInstitutionIDQualifier   string `json:"originating_depository_financial_institution_id_qualifier,required"`
	OriginatingDepositoryFinancialInstitutionName          string `json:"originating_depository_financial_institution_name,required"`
	OriginatorCity                                         string `json:"originator_city,required"`
	OriginatorCompanyEntryDescription                      string `json:"originator_company_entry_description,required"`
	OriginatorCountry                                      string `json:"originator_country,required"`
	OriginatorIdentification                               string `json:"originator_identification,required"`
	OriginatorName                                         string `json:"originator_name,required"`
	OriginatorPostalCode                                   string `json:"originator_postal_code,required,nullable"`
	OriginatorStateOrProvince                              string `json:"originator_state_or_province,required,nullable"`
	OriginatorStreetAddress                                string `json:"originator_street_address,required"`
	PaymentRelatedInformation                              string `json:"payment_related_information,required,nullable"`
	PaymentRelatedInformation2                             string `json:"payment_related_information2,required,nullable"`
	ReceiverCity                                           string `json:"receiver_city,required"`
	ReceiverCountry                                        string `json:"receiver_country,required"`
	ReceiverIdentificationNumber                           string `json:"receiver_identification_number,required,nullable"`
	ReceiverPostalCode                                     string `json:"receiver_postal_code,required,nullable"`
	ReceiverStateOrProvince                                string `json:"receiver_state_or_province,required,nullable"`
	ReceiverStreetAddress                                  string `json:"receiver_street_address,required"`
	ReceivingCompanyOrIndividualName                       string `json:"receiving_company_or_individual_name,required"`
	ReceivingDepositoryFinancialInstitutionCountry         string `json:"receiving_depository_financial_institution_country,required"`
	ReceivingDepositoryFinancialInstitutionID              string `json:"receiving_depository_financial_institution_id,required"`
	ReceivingDepositoryFinancialInstitutionIDQualifier     string `json:"receiving_depository_financial_institution_id_qualifier,required"`
	ReceivingDepositoryFinancialInstitutionName            string `json:"receiving_depository_financial_institution_name,required"`
	TraceNumber                                            string `json:"trace_number,required"`
	JSON                                                   wireTransferSimulationTransactionSourceInboundInternationalACHTransferJSON
}

// wireTransferSimulationTransactionSourceInboundInternationalACHTransferJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceInboundInternationalACHTransfer]
type wireTransferSimulationTransactionSourceInboundInternationalACHTransferJSON struct {
	Amount                                                 apijson.Field
	DestinationCountryCode                                 apijson.Field
	DestinationCurrencyCode                                apijson.Field
	ForeignExchangeIndicator                               apijson.Field
	ForeignExchangeReference                               apijson.Field
	ForeignExchangeReferenceIndicator                      apijson.Field
	ForeignPaymentAmount                                   apijson.Field
	ForeignTraceNumber                                     apijson.Field
	InternationalTransactionTypeCode                       apijson.Field
	OriginatingCurrencyCode                                apijson.Field
	OriginatingDepositoryFinancialInstitutionBranchCountry apijson.Field
	OriginatingDepositoryFinancialInstitutionID            apijson.Field
	OriginatingDepositoryFinancialInstitutionIDQualifier   apijson.Field
	OriginatingDepositoryFinancialInstitutionName          apijson.Field
	OriginatorCity                                         apijson.Field
	OriginatorCompanyEntryDescription                      apijson.Field
	OriginatorCountry                                      apijson.Field
	OriginatorIdentification                               apijson.Field
	OriginatorName                                         apijson.Field
	OriginatorPostalCode                                   apijson.Field
	OriginatorStateOrProvince                              apijson.Field
	OriginatorStreetAddress                                apijson.Field
	PaymentRelatedInformation                              apijson.Field
	PaymentRelatedInformation2                             apijson.Field
	ReceiverCity                                           apijson.Field
	ReceiverCountry                                        apijson.Field
	ReceiverIdentificationNumber                           apijson.Field
	ReceiverPostalCode                                     apijson.Field
	ReceiverStateOrProvince                                apijson.Field
	ReceiverStreetAddress                                  apijson.Field
	ReceivingCompanyOrIndividualName                       apijson.Field
	ReceivingDepositoryFinancialInstitutionCountry         apijson.Field
	ReceivingDepositoryFinancialInstitutionID              apijson.Field
	ReceivingDepositoryFinancialInstitutionIDQualifier     apijson.Field
	ReceivingDepositoryFinancialInstitutionName            apijson.Field
	TraceNumber                                            apijson.Field
	raw                                                    string
	ExtraFields                                            map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Inbound Real Time Payments Transfer Confirmation object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_confirmation`.
type WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real Time Payments transfer.
	Currency WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency `json:"currency,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber string `json:"debtor_account_number,required"`
	// The name provided by the sender of the transfer.
	DebtorName string `json:"debtor_name,required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber string `json:"debtor_routing_number,required"`
	// Additional information included with the transfer.
	RemittanceInformation string `json:"remittance_information,required,nullable"`
	// The Real Time Payments network identification of the transfer
	TransactionIdentification string `json:"transaction_identification,required"`
	JSON                      wireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON
}

// wireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation]
type wireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON struct {
	Amount                    apijson.Field
	CreditorName              apijson.Field
	Currency                  apijson.Field
	DebtorAccountNumber       apijson.Field
	DebtorName                apijson.Field
	DebtorRoutingNumber       apijson.Field
	RemittanceInformation     apijson.Field
	TransactionIdentification apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
// currency. This will always be "USD" for a Real Time Payments transfer.
type WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency string

const (
	// Canadian Dollar (CAD)
	WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyCad WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CAD"
	// Swiss Franc (CHF)
	WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyChf WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CHF"
	// Euro (EUR)
	WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyEur WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "EUR"
	// British Pound (GBP)
	WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyGbp WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "GBP"
	// Japanese Yen (JPY)
	WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyJpy WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "JPY"
	// US Dollar (USD)
	WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyUsd WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "USD"
)

// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
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
	JSON                               wireTransferSimulationTransactionSourceInboundWireDrawdownPaymentJSON
}

// wireTransferSimulationTransactionSourceInboundWireDrawdownPaymentJSON contains
// the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceInboundWireDrawdownPayment]
type wireTransferSimulationTransactionSourceInboundWireDrawdownPaymentJSON struct {
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

func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
// the JSON response if and only if `category` is equal to
// `inbound_wire_drawdown_payment_reversal`.
type WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal struct {
	// The amount that was reversed.
	Amount int64 `json:"amount,required"`
	// The description on the reversal message from Fedwire.
	Description string `json:"description,required"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate time.Time `json:"input_cycle_date,required" format:"date"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required"`
	// The Fedwire sequence number.
	InputSequenceNumber string `json:"input_sequence_number,required"`
	// The Fedwire input source identifier.
	InputSource string `json:"input_source,required"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate time.Time `json:"previous_message_input_cycle_date,required" format:"date"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData string `json:"previous_message_input_message_accountability_data,required"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber string `json:"previous_message_input_sequence_number,required"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource string `json:"previous_message_input_source,required"`
	JSON                       wireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversalJSON
}

// wireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversalJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal]
type wireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversalJSON struct {
	Amount                                        apijson.Field
	Description                                   apijson.Field
	InputCycleDate                                apijson.Field
	InputMessageAccountabilityData                apijson.Field
	InputSequenceNumber                           apijson.Field
	InputSource                                   apijson.Field
	PreviousMessageInputCycleDate                 apijson.Field
	PreviousMessageInputMessageAccountabilityData apijson.Field
	PreviousMessageInputSequenceNumber            apijson.Field
	PreviousMessageInputSource                    apijson.Field
	raw                                           string
	ExtraFields                                   map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Inbound Wire Reversal object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_reversal`.
type WireTransferSimulationTransactionSourceInboundWireReversal struct {
	// The amount that was reversed.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the reversal was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description on the reversal message from Fedwire.
	Description string `json:"description,required"`
	// Additional financial institution information included in the wire reversal.
	FinancialInstitutionToFinancialInstitutionInformation string `json:"financial_institution_to_financial_institution_information,required,nullable"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate time.Time `json:"input_cycle_date,required" format:"date"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required"`
	// The Fedwire sequence number.
	InputSequenceNumber string `json:"input_sequence_number,required"`
	// The Fedwire input source identifier.
	InputSource string `json:"input_source,required"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate time.Time `json:"previous_message_input_cycle_date,required" format:"date"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData string `json:"previous_message_input_message_accountability_data,required"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber string `json:"previous_message_input_sequence_number,required"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource string `json:"previous_message_input_source,required"`
	// Information included in the wire reversal for the receiving financial
	// institution.
	ReceiverFinancialInstitutionInformation string `json:"receiver_financial_institution_information,required,nullable"`
	// The ID for the Transaction associated with the transfer reversal.
	TransactionID string `json:"transaction_id,required,nullable"`
	// The ID for the Wire Transfer that is being reversed.
	WireTransferID string `json:"wire_transfer_id,required"`
	JSON           wireTransferSimulationTransactionSourceInboundWireReversalJSON
}

// wireTransferSimulationTransactionSourceInboundWireReversalJSON contains the JSON
// metadata for the struct
// [WireTransferSimulationTransactionSourceInboundWireReversal]
type wireTransferSimulationTransactionSourceInboundWireReversalJSON struct {
	Amount                                                apijson.Field
	CreatedAt                                             apijson.Field
	Description                                           apijson.Field
	FinancialInstitutionToFinancialInstitutionInformation apijson.Field
	InputCycleDate                                        apijson.Field
	InputMessageAccountabilityData                        apijson.Field
	InputSequenceNumber                                   apijson.Field
	InputSource                                           apijson.Field
	PreviousMessageInputCycleDate                         apijson.Field
	PreviousMessageInputMessageAccountabilityData         apijson.Field
	PreviousMessageInputSequenceNumber                    apijson.Field
	PreviousMessageInputSource                            apijson.Field
	ReceiverFinancialInstitutionInformation               apijson.Field
	TransactionID                                         apijson.Field
	WireTransferID                                        apijson.Field
	raw                                                   string
	ExtraFields                                           map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceInboundWireReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Inbound Wire Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_transfer`.
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
	OriginatorToBeneficiaryInformation      string `json:"originator_to_beneficiary_information,required,nullable"`
	OriginatorToBeneficiaryInformationLine1 string `json:"originator_to_beneficiary_information_line1,required,nullable"`
	OriginatorToBeneficiaryInformationLine2 string `json:"originator_to_beneficiary_information_line2,required,nullable"`
	OriginatorToBeneficiaryInformationLine3 string `json:"originator_to_beneficiary_information_line3,required,nullable"`
	OriginatorToBeneficiaryInformationLine4 string `json:"originator_to_beneficiary_information_line4,required,nullable"`
	JSON                                    wireTransferSimulationTransactionSourceInboundWireTransferJSON
}

// wireTransferSimulationTransactionSourceInboundWireTransferJSON contains the JSON
// metadata for the struct
// [WireTransferSimulationTransactionSourceInboundWireTransfer]
type wireTransferSimulationTransactionSourceInboundWireTransferJSON struct {
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
	OriginatorToBeneficiaryInformation      apijson.Field
	OriginatorToBeneficiaryInformationLine1 apijson.Field
	OriginatorToBeneficiaryInformationLine2 apijson.Field
	OriginatorToBeneficiaryInformationLine3 apijson.Field
	OriginatorToBeneficiaryInformationLine4 apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceInboundWireTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Interest Payment object. This field will be present in the JSON response if
// and only if `category` is equal to `interest_payment`.
type WireTransferSimulationTransactionSourceInterestPayment struct {
	// The account on which the interest was accrued.
	AccruedOnAccountID string `json:"accrued_on_account_id,required,nullable"`
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency WireTransferSimulationTransactionSourceInterestPaymentCurrency `json:"currency,required"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time `json:"period_start,required" format:"date-time"`
	JSON        wireTransferSimulationTransactionSourceInterestPaymentJSON
}

// wireTransferSimulationTransactionSourceInterestPaymentJSON contains the JSON
// metadata for the struct [WireTransferSimulationTransactionSourceInterestPayment]
type wireTransferSimulationTransactionSourceInterestPaymentJSON struct {
	AccruedOnAccountID apijson.Field
	Amount             apijson.Field
	Currency           apijson.Field
	PeriodEnd          apijson.Field
	PeriodStart        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceInterestPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type WireTransferSimulationTransactionSourceInterestPaymentCurrency string

const (
	// Canadian Dollar (CAD)
	WireTransferSimulationTransactionSourceInterestPaymentCurrencyCad WireTransferSimulationTransactionSourceInterestPaymentCurrency = "CAD"
	// Swiss Franc (CHF)
	WireTransferSimulationTransactionSourceInterestPaymentCurrencyChf WireTransferSimulationTransactionSourceInterestPaymentCurrency = "CHF"
	// Euro (EUR)
	WireTransferSimulationTransactionSourceInterestPaymentCurrencyEur WireTransferSimulationTransactionSourceInterestPaymentCurrency = "EUR"
	// British Pound (GBP)
	WireTransferSimulationTransactionSourceInterestPaymentCurrencyGbp WireTransferSimulationTransactionSourceInterestPaymentCurrency = "GBP"
	// Japanese Yen (JPY)
	WireTransferSimulationTransactionSourceInterestPaymentCurrencyJpy WireTransferSimulationTransactionSourceInterestPaymentCurrency = "JPY"
	// US Dollar (USD)
	WireTransferSimulationTransactionSourceInterestPaymentCurrencyUsd WireTransferSimulationTransactionSourceInterestPaymentCurrency = "USD"
)

// A Internal Source object. This field will be present in the JSON response if and
// only if `category` is equal to `internal_source`.
type WireTransferSimulationTransactionSourceInternalSource struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency WireTransferSimulationTransactionSourceInternalSourceCurrency `json:"currency,required"`
	Reason   WireTransferSimulationTransactionSourceInternalSourceReason   `json:"reason,required"`
	JSON     wireTransferSimulationTransactionSourceInternalSourceJSON
}

// wireTransferSimulationTransactionSourceInternalSourceJSON contains the JSON
// metadata for the struct [WireTransferSimulationTransactionSourceInternalSource]
type wireTransferSimulationTransactionSourceInternalSourceJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceInternalSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type WireTransferSimulationTransactionSourceInternalSourceCurrency string

const (
	// Canadian Dollar (CAD)
	WireTransferSimulationTransactionSourceInternalSourceCurrencyCad WireTransferSimulationTransactionSourceInternalSourceCurrency = "CAD"
	// Swiss Franc (CHF)
	WireTransferSimulationTransactionSourceInternalSourceCurrencyChf WireTransferSimulationTransactionSourceInternalSourceCurrency = "CHF"
	// Euro (EUR)
	WireTransferSimulationTransactionSourceInternalSourceCurrencyEur WireTransferSimulationTransactionSourceInternalSourceCurrency = "EUR"
	// British Pound (GBP)
	WireTransferSimulationTransactionSourceInternalSourceCurrencyGbp WireTransferSimulationTransactionSourceInternalSourceCurrency = "GBP"
	// Japanese Yen (JPY)
	WireTransferSimulationTransactionSourceInternalSourceCurrencyJpy WireTransferSimulationTransactionSourceInternalSourceCurrency = "JPY"
	// US Dollar (USD)
	WireTransferSimulationTransactionSourceInternalSourceCurrencyUsd WireTransferSimulationTransactionSourceInternalSourceCurrency = "USD"
)

type WireTransferSimulationTransactionSourceInternalSourceReason string

const (
	// Account closure
	WireTransferSimulationTransactionSourceInternalSourceReasonAccountClosure WireTransferSimulationTransactionSourceInternalSourceReason = "account_closure"
	// Bank migration
	WireTransferSimulationTransactionSourceInternalSourceReasonBankMigration WireTransferSimulationTransactionSourceInternalSourceReason = "bank_migration"
	// Cashback
	WireTransferSimulationTransactionSourceInternalSourceReasonCashback WireTransferSimulationTransactionSourceInternalSourceReason = "cashback"
	// Collection receivable
	WireTransferSimulationTransactionSourceInternalSourceReasonCollectionReceivable WireTransferSimulationTransactionSourceInternalSourceReason = "collection_receivable"
	// Empyreal adjustment
	WireTransferSimulationTransactionSourceInternalSourceReasonEmpyrealAdjustment WireTransferSimulationTransactionSourceInternalSourceReason = "empyreal_adjustment"
	// Error
	WireTransferSimulationTransactionSourceInternalSourceReasonError WireTransferSimulationTransactionSourceInternalSourceReason = "error"
	// Error correction
	WireTransferSimulationTransactionSourceInternalSourceReasonErrorCorrection WireTransferSimulationTransactionSourceInternalSourceReason = "error_correction"
	// Fees
	WireTransferSimulationTransactionSourceInternalSourceReasonFees WireTransferSimulationTransactionSourceInternalSourceReason = "fees"
	// Interest
	WireTransferSimulationTransactionSourceInternalSourceReasonInterest WireTransferSimulationTransactionSourceInternalSourceReason = "interest"
	// Negative balance forgiveness
	WireTransferSimulationTransactionSourceInternalSourceReasonNegativeBalanceForgiveness WireTransferSimulationTransactionSourceInternalSourceReason = "negative_balance_forgiveness"
	// Sample funds
	WireTransferSimulationTransactionSourceInternalSourceReasonSampleFunds WireTransferSimulationTransactionSourceInternalSourceReason = "sample_funds"
	// Sample funds return
	WireTransferSimulationTransactionSourceInternalSourceReasonSampleFundsReturn WireTransferSimulationTransactionSourceInternalSourceReason = "sample_funds_return"
)

// A Real Time Payments Transfer Acknowledgement object. This field will be present
// in the JSON response if and only if `category` is equal to
// `real_time_payments_transfer_acknowledgement`.
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
	JSON       wireTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgementJSON
}

// wireTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgementJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgement]
type wireTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgementJSON struct {
	Amount                   apijson.Field
	DestinationAccountNumber apijson.Field
	DestinationRoutingNumber apijson.Field
	RemittanceInformation    apijson.Field
	TransferID               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Sample Funds object. This field will be present in the JSON response if and
// only if `category` is equal to `sample_funds`.
type WireTransferSimulationTransactionSourceSampleFunds struct {
	// Where the sample funds came from.
	Originator string `json:"originator,required"`
	JSON       wireTransferSimulationTransactionSourceSampleFundsJSON
}

// wireTransferSimulationTransactionSourceSampleFundsJSON contains the JSON
// metadata for the struct [WireTransferSimulationTransactionSourceSampleFunds]
type wireTransferSimulationTransactionSourceSampleFundsJSON struct {
	Originator  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceSampleFunds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Wire Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_intention`.
type WireTransferSimulationTransactionSourceWireTransferIntention struct {
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient string `json:"message_to_recipient,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	TransferID    string `json:"transfer_id,required"`
	JSON          wireTransferSimulationTransactionSourceWireTransferIntentionJSON
}

// wireTransferSimulationTransactionSourceWireTransferIntentionJSON contains the
// JSON metadata for the struct
// [WireTransferSimulationTransactionSourceWireTransferIntention]
type wireTransferSimulationTransactionSourceWireTransferIntentionJSON struct {
	AccountNumber      apijson.Field
	Amount             apijson.Field
	MessageToRecipient apijson.Field
	RoutingNumber      apijson.Field
	TransferID         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceWireTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Wire Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_rejection`.
type WireTransferSimulationTransactionSourceWireTransferRejection struct {
	TransferID string `json:"transfer_id,required"`
	JSON       wireTransferSimulationTransactionSourceWireTransferRejectionJSON
}

// wireTransferSimulationTransactionSourceWireTransferRejectionJSON contains the
// JSON metadata for the struct
// [WireTransferSimulationTransactionSourceWireTransferRejection]
type wireTransferSimulationTransactionSourceWireTransferRejectionJSON struct {
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceWireTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A constant representing the object's type. For this resource it will always be
// `transaction`.
type WireTransferSimulationTransactionType string

const (
	WireTransferSimulationTransactionTypeTransaction WireTransferSimulationTransactionType = "transaction"
)

// A constant representing the object's type. For this resource it will always be
// `inbound_wire_transfer_simulation_result`.
type WireTransferSimulationType string

const (
	WireTransferSimulationTypeInboundWireTransferSimulationResult WireTransferSimulationType = "inbound_wire_transfer_simulation_result"
)

type SimulationWireTransferNewInboundParams struct {
	// The identifier of the Account Number the inbound Wire Transfer is for.
	AccountNumberID param.Field[string] `json:"account_number_id,required"`
	// The transfer amount in cents. Must be positive.
	Amount param.Field[int64] `json:"amount,required"`
	// The sending bank will set beneficiary_address_line1 in production. You can
	// simulate any value here.
	BeneficiaryAddressLine1 param.Field[string] `json:"beneficiary_address_line1"`
	// The sending bank will set beneficiary_address_line2 in production. You can
	// simulate any value here.
	BeneficiaryAddressLine2 param.Field[string] `json:"beneficiary_address_line2"`
	// The sending bank will set beneficiary_address_line3 in production. You can
	// simulate any value here.
	BeneficiaryAddressLine3 param.Field[string] `json:"beneficiary_address_line3"`
	// The sending bank will set beneficiary_name in production. You can simulate any
	// value here.
	BeneficiaryName param.Field[string] `json:"beneficiary_name"`
	// The sending bank will set beneficiary_reference in production. You can simulate
	// any value here.
	BeneficiaryReference param.Field[string] `json:"beneficiary_reference"`
	// The sending bank will set originator_address_line1 in production. You can
	// simulate any value here.
	OriginatorAddressLine1 param.Field[string] `json:"originator_address_line1"`
	// The sending bank will set originator_address_line2 in production. You can
	// simulate any value here.
	OriginatorAddressLine2 param.Field[string] `json:"originator_address_line2"`
	// The sending bank will set originator_address_line3 in production. You can
	// simulate any value here.
	OriginatorAddressLine3 param.Field[string] `json:"originator_address_line3"`
	// The sending bank will set originator_name in production. You can simulate any
	// value here.
	OriginatorName param.Field[string] `json:"originator_name"`
	// The sending bank will set originator_to_beneficiary_information_line1 in
	// production. You can simulate any value here.
	OriginatorToBeneficiaryInformationLine1 param.Field[string] `json:"originator_to_beneficiary_information_line1"`
	// The sending bank will set originator_to_beneficiary_information_line2 in
	// production. You can simulate any value here.
	OriginatorToBeneficiaryInformationLine2 param.Field[string] `json:"originator_to_beneficiary_information_line2"`
	// The sending bank will set originator_to_beneficiary_information_line3 in
	// production. You can simulate any value here.
	OriginatorToBeneficiaryInformationLine3 param.Field[string] `json:"originator_to_beneficiary_information_line3"`
	// The sending bank will set originator_to_beneficiary_information_line4 in
	// production. You can simulate any value here.
	OriginatorToBeneficiaryInformationLine4 param.Field[string] `json:"originator_to_beneficiary_information_line4"`
}

func (r SimulationWireTransferNewInboundParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
