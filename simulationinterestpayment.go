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
	Currency InterestPaymentSimulationResultTransactionCurrency `json:"currency,required"`
	// An informational message describing this transaction. Use the fields in `source`
	// to get more detailed information. This field appears as the line-item on the
	// statement.
	Description string `json:"description,required"`
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

func (r *InterestPaymentSimulationResultTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// Transaction's currency. This will match the currency on the Transcation's
// Account.
type InterestPaymentSimulationResultTransactionCurrency string

const (
	// Canadian Dollar (CAD)
	InterestPaymentSimulationResultTransactionCurrencyCad InterestPaymentSimulationResultTransactionCurrency = "CAD"
	// Swiss Franc (CHF)
	InterestPaymentSimulationResultTransactionCurrencyChf InterestPaymentSimulationResultTransactionCurrency = "CHF"
	// Euro (EUR)
	InterestPaymentSimulationResultTransactionCurrencyEur InterestPaymentSimulationResultTransactionCurrency = "EUR"
	// British Pound (GBP)
	InterestPaymentSimulationResultTransactionCurrencyGbp InterestPaymentSimulationResultTransactionCurrency = "GBP"
	// Japanese Yen (JPY)
	InterestPaymentSimulationResultTransactionCurrencyJpy InterestPaymentSimulationResultTransactionCurrency = "JPY"
	// US Dollar (USD)
	InterestPaymentSimulationResultTransactionCurrencyUsd InterestPaymentSimulationResultTransactionCurrency = "USD"
)

// The type of the route this Transaction came through.
type InterestPaymentSimulationResultTransactionRouteType string

const (
	// An Account Number.
	InterestPaymentSimulationResultTransactionRouteTypeAccountNumber InterestPaymentSimulationResultTransactionRouteType = "account_number"
	// A Card.
	InterestPaymentSimulationResultTransactionRouteTypeCard InterestPaymentSimulationResultTransactionRouteType = "card"
)

// This is an object giving more details on the network-level event that caused the
// Transaction. Note that for backwards compatibility reasons, additional
// undocumented keys may appear in this object. These should be treated as
// deprecated and will be removed in the future.
type InterestPaymentSimulationResultTransactionSource struct {
	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention InterestPaymentSimulationResultTransactionSourceAccountTransferIntention `json:"account_transfer_intention,required,nullable"`
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
	// A Card Revenue Payment object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_revenue_payment`.
	CardRevenuePayment InterestPaymentSimulationResultTransactionSourceCardRevenuePayment `json:"card_revenue_payment,required,nullable"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement InterestPaymentSimulationResultTransactionSourceCardSettlement `json:"card_settlement,required,nullable"`
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category InterestPaymentSimulationResultTransactionSourceCategory `json:"category,required"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptance `json:"check_deposit_acceptance,required,nullable"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn InterestPaymentSimulationResultTransactionSourceCheckDepositReturn `json:"check_deposit_return,required,nullable"`
	// A Check Transfer Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_transfer_deposit`.
	CheckTransferDeposit InterestPaymentSimulationResultTransactionSourceCheckTransferDeposit `json:"check_transfer_deposit,required,nullable"`
	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention InterestPaymentSimulationResultTransactionSourceCheckTransferIntention `json:"check_transfer_intention,required,nullable"`
	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection InterestPaymentSimulationResultTransactionSourceCheckTransferRejection `json:"check_transfer_rejection,required,nullable"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequest `json:"check_transfer_stop_payment_request,required,nullable"`
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
	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment InterestPaymentSimulationResultTransactionSourceInboundWireDrawdownPayment `json:"inbound_wire_drawdown_payment,required,nullable"`
	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal InterestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal `json:"inbound_wire_drawdown_payment_reversal,required,nullable"`
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
	// A Real Time Payments Transfer Acknowledgement object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `real_time_payments_transfer_acknowledgement`.
	RealTimePaymentsTransferAcknowledgement InterestPaymentSimulationResultTransactionSourceRealTimePaymentsTransferAcknowledgement `json:"real_time_payments_transfer_acknowledgement,required,nullable"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds InterestPaymentSimulationResultTransactionSourceSampleFunds `json:"sample_funds,required,nullable"`
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

func (r *InterestPaymentSimulationResultTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
type InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrency string

const (
	// Canadian Dollar (CAD)
	InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrencyCad InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrency = "CAD"
	// Swiss Franc (CHF)
	InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrencyChf InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrency = "CHF"
	// Euro (EUR)
	InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrencyEur InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrency = "EUR"
	// British Pound (GBP)
	InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrencyGbp InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrency = "GBP"
	// Japanese Yen (JPY)
	InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrencyJpy InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrency = "JPY"
	// US Dollar (USD)
	InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrencyUsd InterestPaymentSimulationResultTransactionSourceAccountTransferIntentionCurrency = "USD"
)

// A ACH Transfer Intention object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_intention`.
type InterestPaymentSimulationResultTransactionSourceACHTransferIntention struct {
	AccountNumber string `json:"account_number,required"`
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount              int64  `json:"amount,required"`
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
	AccountNumber       apijson.Field
	Amount              apijson.Field
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
	// The three character ACH return code, in the range R01 to R85.
	RawReturnReasonCode string `json:"raw_return_reason_code,required"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode `json:"return_reason_code,required"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID string `json:"transaction_id,required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID string `json:"transfer_id,required"`
	JSON       interestPaymentSimulationResultTransactionSourceACHTransferReturnJSON
}

// interestPaymentSimulationResultTransactionSourceACHTransferReturnJSON contains
// the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceACHTransferReturn]
type interestPaymentSimulationResultTransactionSourceACHTransferReturnJSON struct {
	CreatedAt           apijson.Field
	RawReturnReasonCode apijson.Field
	ReturnReasonCode    apijson.Field
	TransactionID       apijson.Field
	TransferID          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Why the ACH Transfer was returned.
type InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode string

const (
	// Code R01. Insufficient funds in the source account.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInsufficientFund InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "insufficient_fund"
	// Code R03. The account does not exist or the receiving bank was unable to locate
	// it.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeNoAccount InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "no_account"
	// Code R02. The account is closed.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAccountClosed InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "account_closed"
	// Code R04. The account number is invalid at the receiving bank.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	// Code R16. The account was frozen per the Office of Foreign Assets Control.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	// Code R23. The receiving bank account refused a credit transfer.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	// Code R05. The receiving bank rejected because of an incorrect Standard Entry
	// Class code.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	// Code R29. The corporate customer reversed the transfer.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	// Code R08. The receiving bank stopped payment on this transfer.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodePaymentStopped InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "payment_stopped"
	// Code R20. The receiving bank account does not perform transfers.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeNonTransactionAccount InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "non_transaction_account"
	// Code R09. The receiving bank account does not have enough available balance for
	// the transfer.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeUncollectedFunds InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "uncollected_funds"
	// Code R28. The routing number is incorrect.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	// Code R10. The customer reversed the transfer.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	// Code R19. The amount field is incorrect or too large.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAmountFieldError InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "amount_field_error"
	// Code R07. The customer who initiated the transfer revoked authorization.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	// Code R13. The routing number is invalid.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	// Code R17. The receiving bank is unable to process a field in the transfer.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeFileRecordEditCriteria InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	// Code R45. The individual name field was invalid.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualName InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	// Code R06. The originating financial institution reversed the transfer.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	// Code R34. The receiving bank's regulatory supervisor has limited their
	// participation.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeLimitedParticipationDfi InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	// Code R85. The outbound international ACH transfer was incorrect.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "incorrectly_coded_outbound_international_payment"
	// Code R12. A rare return reason. The account was sold to another bank.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAccountSoldToAnotherDfi InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "account_sold_to_another_dfi"
	// Code R25. The addenda record is incorrect or missing.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAddendaError InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "addenda_error"
	// Code R15. A rare return reason. The account holder is deceased.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeBeneficiaryOrAccountHolderDeceased InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "beneficiary_or_account_holder_deceased"
	// Code R11. A rare return reason. The customer authorized some payment to the
	// sender, but this payment was not in error.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedNotWithinAuthorizationTerms InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_not_within_authorization_terms"
	// Code R74. A rare return reason. Sent in response to a return that was returned
	// with code `field_error`. The latest return should include the corrected
	// field(s).
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCorrectedReturn InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "corrected_return"
	// Code R24. A rare return reason. The receiving bank received an exact duplicate
	// entry with the same trace number and amount.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeDuplicateEntry InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "duplicate_entry"
	// Code R67. A rare return reason. The return this message refers to was a
	// duplicate.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeDuplicateReturn InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "duplicate_return"
	// Code R47. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEnrDuplicateEnrollment InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "enr_duplicate_enrollment"
	// Code R43. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidDfiAccountNumber InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_dfi_account_number"
	// Code R44. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualIDNumber InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_id_number"
	// Code R46. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidRepresentativePayeeIndicator InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_representative_payee_indicator"
	// Code R41. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidTransactionCode InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_transaction_code"
	// Code R40. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEnrReturnOfEnrEntry InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "enr_return_of_enr_entry"
	// Code R42. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEnrRoutingNumberCheckDigitError InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "enr_routing_number_check_digit_error"
	// Code R84. A rare return reason. The International ACH Transfer cannot be
	// processed by the gateway.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEntryNotProcessedByGateway InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "entry_not_processed_by_gateway"
	// Code R69. A rare return reason. One or more of the fields in the ACH were
	// malformed.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeFieldError InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "field_error"
	// Code R83. A rare return reason. The Foreign receiving bank was unable to settle
	// this ACH transfer.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeForeignReceivingDfiUnableToSettle InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "foreign_receiving_dfi_unable_to_settle"
	// Code R80. A rare return reason. The International ACH Transfer is malformed.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeIatEntryCodingError InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "iat_entry_coding_error"
	// Code R18. A rare return reason. The ACH has an improper effective entry date
	// field.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeImproperEffectiveEntryDate InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "improper_effective_entry_date"
	// Code R39. A rare return reason. The source document related to this ACH, usually
	// an ACH check conversion, was presented to the bank.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeImproperSourceDocumentSourceDocumentPresented InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "improper_source_document_source_document_presented"
	// Code R21. A rare return reason. The Company ID field of the ACH was invalid.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInvalidCompanyID InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "invalid_company_id"
	// Code R82. A rare return reason. The foreign receiving bank identifier for an
	// International ACH Transfer was invalid.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInvalidForeignReceivingDfiIdentification InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "invalid_foreign_receiving_dfi_identification"
	// Code R22. A rare return reason. The Individual ID number field of the ACH was
	// invalid.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInvalidIndividualIDNumber InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "invalid_individual_id_number"
	// Code R53. A rare return reason. Both the Represented Check ("RCK") entry and the
	// original check were presented to the bank.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeItemAndRckEntryPresentedForPayment InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "item_and_rck_entry_presented_for_payment"
	// Code R51. A rare return reason. The Represented Check ("RCK") entry is
	// ineligible.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeItemRelatedToRckEntryIsIneligible InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "item_related_to_rck_entry_is_ineligible"
	// Code R26. A rare return reason. The ACH is missing a required field.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeMandatoryFieldError InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "mandatory_field_error"
	// Code R71. A rare return reason. The receiving bank does not recognize the
	// routing number in a dishonored return entry.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeMisroutedDishonoredReturn InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "misrouted_dishonored_return"
	// Code R61. A rare return reason. The receiving bank does not recognize the
	// routing number in a return entry.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeMisroutedReturn InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "misrouted_return"
	// Code R76. A rare return reason. Sent in response to a return, the bank does not
	// find the errors alleged by the returning bank.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeNoErrorsFound InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "no_errors_found"
	// Code R77. A rare return reason. The receiving bank does not accept the return of
	// the erroneous debit. The funds are not available at the receiving bank.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeNonAcceptanceOfR62DishonoredReturn InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "non_acceptance_of_r62_dishonored_return"
	// Code R81. A rare return reason. The receiving bank does not accept International
	// ACH Transfers.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeNonParticipantInIatProgram InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "non_participant_in_iat_program"
	// Code R31. A rare return reason. A return that has been agreed to be accepted by
	// the receiving bank, despite falling outside of the usual return timeframe.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodePermissibleReturnEntry InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "permissible_return_entry"
	// Code R70. A rare return reason. The receiving bank had not approved this return.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodePermissibleReturnEntryNotAccepted InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "permissible_return_entry_not_accepted"
	// Code R32. A rare return reason. The receiving bank could not settle this
	// transaction.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeRdfiNonSettlement InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "rdfi_non_settlement"
	// Code R30. A rare return reason. The receiving bank does not accept Check
	// Truncation ACH transfers.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeRdfiParticipantInCheckTruncationProgram InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "rdfi_participant_in_check_truncation_program"
	// Code R14. A rare return reason. The payee is deceased.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	// Code R75. A rare return reason. The originating bank disputes that an earlier
	// `duplicate_entry` return was actually a duplicate.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeReturnNotADuplicate InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "return_not_a_duplicate"
	// Code R62. A rare return reason. The originating bank made a mistake earlier and
	// this return corrects it.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeReturnOfErroneousOrReversingDebit InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "return_of_erroneous_or_reversing_debit"
	// Code R36. A rare return reason. Return of a malformed credit entry.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeReturnOfImproperCreditEntry InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "return_of_improper_credit_entry"
	// Code R35. A rare return reason. Return of a malformed debit entry.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeReturnOfImproperDebitEntry InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "return_of_improper_debit_entry"
	// Code R33. A rare return reason. Return of a Destroyed Check ("XKC") entry.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeReturnOfXckEntry InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "return_of_xck_entry"
	// Code R37. A rare return reason. The source document related to this ACH, usually
	// an ACH check conversion, was presented to the bank.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeSourceDocumentPresentedForPayment InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "source_document_presented_for_payment"
	// Code R50. A rare return reason. State law prevents the bank from accepting the
	// Represented Check ("RCK") entry.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeStateLawAffectingRckAcceptance InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "state_law_affecting_rck_acceptance"
	// Code R52. A rare return reason. A stop payment was issued on a Represented Check
	// ("RCK") entry.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeStopPaymentOnItemRelatedToRckEntry InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "stop_payment_on_item_related_to_rck_entry"
	// Code R38. A rare return reason. The source attached to the ACH, usually an ACH
	// check conversion, includes a stop payment.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeStopPaymentOnSourceDocument InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "stop_payment_on_source_document"
	// Code R73. A rare return reason. The bank receiving an `untimely_return` believes
	// it was on time.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeTimelyOriginalReturn InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "timely_original_return"
	// Code R27. A rare return reason. An ACH Return's trace number does not match an
	// originated ACH.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeTraceNumberError InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "trace_number_error"
	// Code R72. A rare return reason. The dishonored return was sent too late.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeUntimelyDishonoredReturn InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "untimely_dishonored_return"
	// Code R68. A rare return reason. The return was sent too late.
	InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeUntimelyReturn InterestPaymentSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "untimely_return"
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

func (r *InterestPaymentSimulationResultTransactionSourceCardRefund) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type InterestPaymentSimulationResultTransactionSourceCardRefundCurrency string

const (
	// Canadian Dollar (CAD)
	InterestPaymentSimulationResultTransactionSourceCardRefundCurrencyCad InterestPaymentSimulationResultTransactionSourceCardRefundCurrency = "CAD"
	// Swiss Franc (CHF)
	InterestPaymentSimulationResultTransactionSourceCardRefundCurrencyChf InterestPaymentSimulationResultTransactionSourceCardRefundCurrency = "CHF"
	// Euro (EUR)
	InterestPaymentSimulationResultTransactionSourceCardRefundCurrencyEur InterestPaymentSimulationResultTransactionSourceCardRefundCurrency = "EUR"
	// British Pound (GBP)
	InterestPaymentSimulationResultTransactionSourceCardRefundCurrencyGbp InterestPaymentSimulationResultTransactionSourceCardRefundCurrency = "GBP"
	// Japanese Yen (JPY)
	InterestPaymentSimulationResultTransactionSourceCardRefundCurrencyJpy InterestPaymentSimulationResultTransactionSourceCardRefundCurrency = "JPY"
	// US Dollar (USD)
	InterestPaymentSimulationResultTransactionSourceCardRefundCurrencyUsd InterestPaymentSimulationResultTransactionSourceCardRefundCurrency = "USD"
)

// A constant representing the object's type. For this resource it will always be
// `card_refund`.
type InterestPaymentSimulationResultTransactionSourceCardRefundType string

const (
	InterestPaymentSimulationResultTransactionSourceCardRefundTypeCardRefund InterestPaymentSimulationResultTransactionSourceCardRefundType = "card_refund"
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
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time `json:"period_start,required" format:"date-time"`
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
	PeriodEnd             apijson.Field
	PeriodStart           apijson.Field
	TransactedOnAccountID apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceCardRevenuePayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrency string

const (
	// Canadian Dollar (CAD)
	InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrencyCad InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrency = "CAD"
	// Swiss Franc (CHF)
	InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrencyChf InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrency = "CHF"
	// Euro (EUR)
	InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrencyEur InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrency = "EUR"
	// British Pound (GBP)
	InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrencyGbp InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrency = "GBP"
	// Japanese Yen (JPY)
	InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrencyJpy InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrency = "JPY"
	// US Dollar (USD)
	InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrencyUsd InterestPaymentSimulationResultTransactionSourceCardRevenuePaymentCurrency = "USD"
)

// A Card Settlement object. This field will be present in the JSON response if and
// only if `category` is equal to `card_settlement`.
type InterestPaymentSimulationResultTransactionSourceCardSettlement struct {
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
	Currency InterestPaymentSimulationResultTransactionSourceCardSettlementCurrency `json:"currency,required"`
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
	Type InterestPaymentSimulationResultTransactionSourceCardSettlementType `json:"type,required"`
	JSON interestPaymentSimulationResultTransactionSourceCardSettlementJSON
}

// interestPaymentSimulationResultTransactionSourceCardSettlementJSON contains the
// JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceCardSettlement]
type interestPaymentSimulationResultTransactionSourceCardSettlementJSON struct {
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

func (r *InterestPaymentSimulationResultTransactionSourceCardSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's settlement currency.
type InterestPaymentSimulationResultTransactionSourceCardSettlementCurrency string

const (
	// Canadian Dollar (CAD)
	InterestPaymentSimulationResultTransactionSourceCardSettlementCurrencyCad InterestPaymentSimulationResultTransactionSourceCardSettlementCurrency = "CAD"
	// Swiss Franc (CHF)
	InterestPaymentSimulationResultTransactionSourceCardSettlementCurrencyChf InterestPaymentSimulationResultTransactionSourceCardSettlementCurrency = "CHF"
	// Euro (EUR)
	InterestPaymentSimulationResultTransactionSourceCardSettlementCurrencyEur InterestPaymentSimulationResultTransactionSourceCardSettlementCurrency = "EUR"
	// British Pound (GBP)
	InterestPaymentSimulationResultTransactionSourceCardSettlementCurrencyGbp InterestPaymentSimulationResultTransactionSourceCardSettlementCurrency = "GBP"
	// Japanese Yen (JPY)
	InterestPaymentSimulationResultTransactionSourceCardSettlementCurrencyJpy InterestPaymentSimulationResultTransactionSourceCardSettlementCurrency = "JPY"
	// US Dollar (USD)
	InterestPaymentSimulationResultTransactionSourceCardSettlementCurrencyUsd InterestPaymentSimulationResultTransactionSourceCardSettlementCurrency = "USD"
)

// A constant representing the object's type. For this resource it will always be
// `card_settlement`.
type InterestPaymentSimulationResultTransactionSourceCardSettlementType string

const (
	InterestPaymentSimulationResultTransactionSourceCardSettlementTypeCardSettlement InterestPaymentSimulationResultTransactionSourceCardSettlementType = "card_settlement"
)

// The type of transaction that took place. We may add additional possible values
// for this enum over time; your application should be able to handle such
// additions gracefully.
type InterestPaymentSimulationResultTransactionSourceCategory string

const (
	// The Transaction was created by a Account Transfer Intention object. Details will
	// be under the `account_transfer_intention` object.
	InterestPaymentSimulationResultTransactionSourceCategoryAccountTransferIntention InterestPaymentSimulationResultTransactionSourceCategory = "account_transfer_intention"
	// The Transaction was created by a ACH Transfer Intention object. Details will be
	// under the `ach_transfer_intention` object.
	InterestPaymentSimulationResultTransactionSourceCategoryACHTransferIntention InterestPaymentSimulationResultTransactionSourceCategory = "ach_transfer_intention"
	// The Transaction was created by a ACH Transfer Rejection object. Details will be
	// under the `ach_transfer_rejection` object.
	InterestPaymentSimulationResultTransactionSourceCategoryACHTransferRejection InterestPaymentSimulationResultTransactionSourceCategory = "ach_transfer_rejection"
	// The Transaction was created by a ACH Transfer Return object. Details will be
	// under the `ach_transfer_return` object.
	InterestPaymentSimulationResultTransactionSourceCategoryACHTransferReturn InterestPaymentSimulationResultTransactionSourceCategory = "ach_transfer_return"
	// The Transaction was created by a Card Dispute Acceptance object. Details will be
	// under the `card_dispute_acceptance` object.
	InterestPaymentSimulationResultTransactionSourceCategoryCardDisputeAcceptance InterestPaymentSimulationResultTransactionSourceCategory = "card_dispute_acceptance"
	// The Transaction was created by a Card Refund object. Details will be under the
	// `card_refund` object.
	InterestPaymentSimulationResultTransactionSourceCategoryCardRefund InterestPaymentSimulationResultTransactionSourceCategory = "card_refund"
	// The Transaction was created by a Card Revenue Payment object. Details will be
	// under the `card_revenue_payment` object.
	InterestPaymentSimulationResultTransactionSourceCategoryCardRevenuePayment InterestPaymentSimulationResultTransactionSourceCategory = "card_revenue_payment"
	// The Transaction was created by a Card Settlement object. Details will be under
	// the `card_settlement` object.
	InterestPaymentSimulationResultTransactionSourceCategoryCardSettlement InterestPaymentSimulationResultTransactionSourceCategory = "card_settlement"
	// The Transaction was created by a Check Deposit Acceptance object. Details will
	// be under the `check_deposit_acceptance` object.
	InterestPaymentSimulationResultTransactionSourceCategoryCheckDepositAcceptance InterestPaymentSimulationResultTransactionSourceCategory = "check_deposit_acceptance"
	// The Transaction was created by a Check Deposit Return object. Details will be
	// under the `check_deposit_return` object.
	InterestPaymentSimulationResultTransactionSourceCategoryCheckDepositReturn InterestPaymentSimulationResultTransactionSourceCategory = "check_deposit_return"
	// The Transaction was created by a Check Transfer Deposit object. Details will be
	// under the `check_transfer_deposit` object.
	InterestPaymentSimulationResultTransactionSourceCategoryCheckTransferDeposit InterestPaymentSimulationResultTransactionSourceCategory = "check_transfer_deposit"
	// The Transaction was created by a Check Transfer Intention object. Details will
	// be under the `check_transfer_intention` object.
	InterestPaymentSimulationResultTransactionSourceCategoryCheckTransferIntention InterestPaymentSimulationResultTransactionSourceCategory = "check_transfer_intention"
	// The Transaction was created by a Check Transfer Rejection object. Details will
	// be under the `check_transfer_rejection` object.
	InterestPaymentSimulationResultTransactionSourceCategoryCheckTransferRejection InterestPaymentSimulationResultTransactionSourceCategory = "check_transfer_rejection"
	// The Transaction was created by a Check Transfer Stop Payment Request object.
	// Details will be under the `check_transfer_stop_payment_request` object.
	InterestPaymentSimulationResultTransactionSourceCategoryCheckTransferStopPaymentRequest InterestPaymentSimulationResultTransactionSourceCategory = "check_transfer_stop_payment_request"
	// The Transaction was created by a Fee Payment object. Details will be under the
	// `fee_payment` object.
	InterestPaymentSimulationResultTransactionSourceCategoryFeePayment InterestPaymentSimulationResultTransactionSourceCategory = "fee_payment"
	// The Transaction was created by a Inbound ACH Transfer object. Details will be
	// under the `inbound_ach_transfer` object.
	InterestPaymentSimulationResultTransactionSourceCategoryInboundACHTransfer InterestPaymentSimulationResultTransactionSourceCategory = "inbound_ach_transfer"
	// The Transaction was created by a Inbound ACH Transfer Return Intention object.
	// Details will be under the `inbound_ach_transfer_return_intention` object.
	InterestPaymentSimulationResultTransactionSourceCategoryInboundACHTransferReturnIntention InterestPaymentSimulationResultTransactionSourceCategory = "inbound_ach_transfer_return_intention"
	// The Transaction was created by a Inbound Check object. Details will be under the
	// `inbound_check` object.
	InterestPaymentSimulationResultTransactionSourceCategoryInboundCheck InterestPaymentSimulationResultTransactionSourceCategory = "inbound_check"
	// The Transaction was created by a Inbound International ACH Transfer object.
	// Details will be under the `inbound_international_ach_transfer` object.
	InterestPaymentSimulationResultTransactionSourceCategoryInboundInternationalACHTransfer InterestPaymentSimulationResultTransactionSourceCategory = "inbound_international_ach_transfer"
	// The Transaction was created by a Inbound Real Time Payments Transfer
	// Confirmation object. Details will be under the
	// `inbound_real_time_payments_transfer_confirmation` object.
	InterestPaymentSimulationResultTransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation InterestPaymentSimulationResultTransactionSourceCategory = "inbound_real_time_payments_transfer_confirmation"
	// The Transaction was created by a Inbound Wire Drawdown Payment object. Details
	// will be under the `inbound_wire_drawdown_payment` object.
	InterestPaymentSimulationResultTransactionSourceCategoryInboundWireDrawdownPayment InterestPaymentSimulationResultTransactionSourceCategory = "inbound_wire_drawdown_payment"
	// The Transaction was created by a Inbound Wire Drawdown Payment Reversal object.
	// Details will be under the `inbound_wire_drawdown_payment_reversal` object.
	InterestPaymentSimulationResultTransactionSourceCategoryInboundWireDrawdownPaymentReversal InterestPaymentSimulationResultTransactionSourceCategory = "inbound_wire_drawdown_payment_reversal"
	// The Transaction was created by a Inbound Wire Reversal object. Details will be
	// under the `inbound_wire_reversal` object.
	InterestPaymentSimulationResultTransactionSourceCategoryInboundWireReversal InterestPaymentSimulationResultTransactionSourceCategory = "inbound_wire_reversal"
	// The Transaction was created by a Inbound Wire Transfer object. Details will be
	// under the `inbound_wire_transfer` object.
	InterestPaymentSimulationResultTransactionSourceCategoryInboundWireTransfer InterestPaymentSimulationResultTransactionSourceCategory = "inbound_wire_transfer"
	// The Transaction was created by a Interest Payment object. Details will be under
	// the `interest_payment` object.
	InterestPaymentSimulationResultTransactionSourceCategoryInterestPayment InterestPaymentSimulationResultTransactionSourceCategory = "interest_payment"
	// The Transaction was created by a Internal Source object. Details will be under
	// the `internal_source` object.
	InterestPaymentSimulationResultTransactionSourceCategoryInternalSource InterestPaymentSimulationResultTransactionSourceCategory = "internal_source"
	// The Transaction was created by a Real Time Payments Transfer Acknowledgement
	// object. Details will be under the `real_time_payments_transfer_acknowledgement`
	// object.
	InterestPaymentSimulationResultTransactionSourceCategoryRealTimePaymentsTransferAcknowledgement InterestPaymentSimulationResultTransactionSourceCategory = "real_time_payments_transfer_acknowledgement"
	// The Transaction was created by a Sample Funds object. Details will be under the
	// `sample_funds` object.
	InterestPaymentSimulationResultTransactionSourceCategorySampleFunds InterestPaymentSimulationResultTransactionSourceCategory = "sample_funds"
	// The Transaction was created by a Wire Transfer Intention object. Details will be
	// under the `wire_transfer_intention` object.
	InterestPaymentSimulationResultTransactionSourceCategoryWireTransferIntention InterestPaymentSimulationResultTransactionSourceCategory = "wire_transfer_intention"
	// The Transaction was created by a Wire Transfer Rejection object. Details will be
	// under the `wire_transfer_rejection` object.
	InterestPaymentSimulationResultTransactionSourceCategoryWireTransferRejection InterestPaymentSimulationResultTransactionSourceCategory = "wire_transfer_rejection"
	// The Transaction was made for an undocumented or deprecated reason.
	InterestPaymentSimulationResultTransactionSourceCategoryOther InterestPaymentSimulationResultTransactionSourceCategory = "other"
)

// A Check Deposit Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_acceptance`.
type InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptance struct {
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
	Currency InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrency `json:"currency,required"`
	// The routing number printed on the check.
	RoutingNumber string `json:"routing_number,required"`
	// The check serial number, if present, for consumer checks. For business checks,
	// the serial number is usually in the `auxiliary_on_us` field.
	SerialNumber string `json:"serial_number,required,nullable"`
	JSON         interestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceJSON
}

// interestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptance]
type interestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceJSON struct {
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

func (r *InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrency string

const (
	// Canadian Dollar (CAD)
	InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyCad InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "CAD"
	// Swiss Franc (CHF)
	InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyChf InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "CHF"
	// Euro (EUR)
	InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyEur InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "EUR"
	// British Pound (GBP)
	InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyGbp InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "GBP"
	// Japanese Yen (JPY)
	InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyJpy InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "JPY"
	// US Dollar (USD)
	InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyUsd InterestPaymentSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "USD"
)

// A Check Deposit Return object. This field will be present in the JSON response
// if and only if `category` is equal to `check_deposit_return`.
type InterestPaymentSimulationResultTransactionSourceCheckDepositReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID string `json:"check_deposit_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency     InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrency     `json:"currency,required"`
	ReturnReason InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason `json:"return_reason,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID string `json:"transaction_id,required"`
	JSON          interestPaymentSimulationResultTransactionSourceCheckDepositReturnJSON
}

// interestPaymentSimulationResultTransactionSourceCheckDepositReturnJSON contains
// the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceCheckDepositReturn]
type interestPaymentSimulationResultTransactionSourceCheckDepositReturnJSON struct {
	Amount         apijson.Field
	CheckDepositID apijson.Field
	Currency       apijson.Field
	ReturnReason   apijson.Field
	ReturnedAt     apijson.Field
	TransactionID  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceCheckDepositReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrency string

const (
	// Canadian Dollar (CAD)
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrencyCad InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrency = "CAD"
	// Swiss Franc (CHF)
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrencyChf InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrency = "CHF"
	// Euro (EUR)
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrencyEur InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrency = "EUR"
	// British Pound (GBP)
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrencyGbp InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrency = "GBP"
	// Japanese Yen (JPY)
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrencyJpy InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrency = "JPY"
	// US Dollar (USD)
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrencyUsd InterestPaymentSimulationResultTransactionSourceCheckDepositReturnCurrency = "USD"
)

type InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason string

const (
	// The check doesn't allow ACH conversion.
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonACHConversionNotSupported InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "ach_conversion_not_supported"
	// The account is closed.
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonClosedAccount InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "closed_account"
	// The check has already been deposited.
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonDuplicateSubmission InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "duplicate_submission"
	// Insufficient funds
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonInsufficientFunds InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "insufficient_funds"
	// No account was found matching the check details.
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonNoAccount InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "no_account"
	// The check was not authorized.
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonNotAuthorized InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "not_authorized"
	// The check is too old.
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonStaleDated InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "stale_dated"
	// The payment has been stopped by the account holder.
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonStopPayment InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "stop_payment"
	// The reason for the return is unknown.
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonUnknownReason InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "unknown_reason"
	// The image doesn't match the details submitted.
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonUnmatchedDetails InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "unmatched_details"
	// The image could not be read.
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonUnreadableImage InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "unreadable_image"
	// The check endorsement was irregular.
	InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReasonEndorsementIrregular InterestPaymentSimulationResultTransactionSourceCheckDepositReturnReturnReason = "endorsement_irregular"
)

// A Check Transfer Deposit object. This field will be present in the JSON response
// if and only if `category` is equal to `check_transfer_deposit`.
type InterestPaymentSimulationResultTransactionSourceCheckTransferDeposit struct {
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
	Type InterestPaymentSimulationResultTransactionSourceCheckTransferDepositType `json:"type,required"`
	JSON interestPaymentSimulationResultTransactionSourceCheckTransferDepositJSON
}

// interestPaymentSimulationResultTransactionSourceCheckTransferDepositJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceCheckTransferDeposit]
type interestPaymentSimulationResultTransactionSourceCheckTransferDepositJSON struct {
	BackImageFileID  apijson.Field
	DepositedAt      apijson.Field
	FrontImageFileID apijson.Field
	TransactionID    apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceCheckTransferDeposit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_deposit`.
type InterestPaymentSimulationResultTransactionSourceCheckTransferDepositType string

const (
	InterestPaymentSimulationResultTransactionSourceCheckTransferDepositTypeCheckTransferDeposit InterestPaymentSimulationResultTransactionSourceCheckTransferDepositType = "check_transfer_deposit"
)

// A Check Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_intention`.
type InterestPaymentSimulationResultTransactionSourceCheckTransferIntention struct {
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
	Currency InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrency `json:"currency,required"`
	// The name that will be printed on the check.
	RecipientName string `json:"recipient_name,required,nullable"`
	// The identifier of the Check Transfer with which this is associated.
	TransferID string `json:"transfer_id,required"`
	JSON       interestPaymentSimulationResultTransactionSourceCheckTransferIntentionJSON
}

// interestPaymentSimulationResultTransactionSourceCheckTransferIntentionJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceCheckTransferIntention]
type interestPaymentSimulationResultTransactionSourceCheckTransferIntentionJSON struct {
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

func (r *InterestPaymentSimulationResultTransactionSourceCheckTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
type InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrency string

const (
	// Canadian Dollar (CAD)
	InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrencyCad InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrency = "CAD"
	// Swiss Franc (CHF)
	InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrencyChf InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrency = "CHF"
	// Euro (EUR)
	InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrencyEur InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrency = "EUR"
	// British Pound (GBP)
	InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrencyGbp InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrency = "GBP"
	// Japanese Yen (JPY)
	InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrencyJpy InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrency = "JPY"
	// US Dollar (USD)
	InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrencyUsd InterestPaymentSimulationResultTransactionSourceCheckTransferIntentionCurrency = "USD"
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
	// The reason why this transfer was stopped.
	Reason InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestReason `json:"reason,required"`
	// The time the stop-payment was requested.
	RequestedAt time.Time `json:"requested_at,required" format:"date-time"`
	// The ID of the check transfer that was stopped.
	TransferID string `json:"transfer_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestType `json:"type,required"`
	JSON interestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestJSON
}

// interestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequest]
type interestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestJSON struct {
	Reason      apijson.Field
	RequestedAt apijson.Field
	TransferID  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The reason why this transfer was stopped.
type InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestReason string

const (
	// The check could not be delivered.
	InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestReasonMailDeliveryFailed InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestReason = "mail_delivery_failed"
	// The check was stopped for another reason.
	InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestReasonUnknown InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestReason = "unknown"
)

// A constant representing the object's type. For this resource it will always be
// `check_transfer_stop_payment_request`.
type InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestType string

const (
	InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest InterestPaymentSimulationResultTransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

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

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type InterestPaymentSimulationResultTransactionSourceFeePaymentCurrency string

const (
	// Canadian Dollar (CAD)
	InterestPaymentSimulationResultTransactionSourceFeePaymentCurrencyCad InterestPaymentSimulationResultTransactionSourceFeePaymentCurrency = "CAD"
	// Swiss Franc (CHF)
	InterestPaymentSimulationResultTransactionSourceFeePaymentCurrencyChf InterestPaymentSimulationResultTransactionSourceFeePaymentCurrency = "CHF"
	// Euro (EUR)
	InterestPaymentSimulationResultTransactionSourceFeePaymentCurrencyEur InterestPaymentSimulationResultTransactionSourceFeePaymentCurrency = "EUR"
	// British Pound (GBP)
	InterestPaymentSimulationResultTransactionSourceFeePaymentCurrencyGbp InterestPaymentSimulationResultTransactionSourceFeePaymentCurrency = "GBP"
	// Japanese Yen (JPY)
	InterestPaymentSimulationResultTransactionSourceFeePaymentCurrencyJpy InterestPaymentSimulationResultTransactionSourceFeePaymentCurrency = "JPY"
	// US Dollar (USD)
	InterestPaymentSimulationResultTransactionSourceFeePaymentCurrencyUsd InterestPaymentSimulationResultTransactionSourceFeePaymentCurrency = "USD"
)

// A Inbound ACH Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_ach_transfer`.
type InterestPaymentSimulationResultTransactionSourceInboundACHTransfer struct {
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
	JSON                               interestPaymentSimulationResultTransactionSourceInboundACHTransferJSON
}

// interestPaymentSimulationResultTransactionSourceInboundACHTransferJSON contains
// the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceInboundACHTransfer]
type interestPaymentSimulationResultTransactionSourceInboundACHTransferJSON struct {
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

func (r *InterestPaymentSimulationResultTransactionSourceInboundACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Inbound Check object. This field will be present in the JSON response if and
// only if `category` is equal to `inbound_check`.
type InterestPaymentSimulationResultTransactionSourceInboundCheck struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount                int64  `json:"amount,required"`
	CheckFrontImageFileID string `json:"check_front_image_file_id,required,nullable"`
	CheckNumber           string `json:"check_number,required,nullable"`
	CheckRearImageFileID  string `json:"check_rear_image_file_id,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency InterestPaymentSimulationResultTransactionSourceInboundCheckCurrency `json:"currency,required"`
	JSON     interestPaymentSimulationResultTransactionSourceInboundCheckJSON
}

// interestPaymentSimulationResultTransactionSourceInboundCheckJSON contains the
// JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceInboundCheck]
type interestPaymentSimulationResultTransactionSourceInboundCheckJSON struct {
	Amount                apijson.Field
	CheckFrontImageFileID apijson.Field
	CheckNumber           apijson.Field
	CheckRearImageFileID  apijson.Field
	Currency              apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceInboundCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type InterestPaymentSimulationResultTransactionSourceInboundCheckCurrency string

const (
	// Canadian Dollar (CAD)
	InterestPaymentSimulationResultTransactionSourceInboundCheckCurrencyCad InterestPaymentSimulationResultTransactionSourceInboundCheckCurrency = "CAD"
	// Swiss Franc (CHF)
	InterestPaymentSimulationResultTransactionSourceInboundCheckCurrencyChf InterestPaymentSimulationResultTransactionSourceInboundCheckCurrency = "CHF"
	// Euro (EUR)
	InterestPaymentSimulationResultTransactionSourceInboundCheckCurrencyEur InterestPaymentSimulationResultTransactionSourceInboundCheckCurrency = "EUR"
	// British Pound (GBP)
	InterestPaymentSimulationResultTransactionSourceInboundCheckCurrencyGbp InterestPaymentSimulationResultTransactionSourceInboundCheckCurrency = "GBP"
	// Japanese Yen (JPY)
	InterestPaymentSimulationResultTransactionSourceInboundCheckCurrencyJpy InterestPaymentSimulationResultTransactionSourceInboundCheckCurrency = "JPY"
	// US Dollar (USD)
	InterestPaymentSimulationResultTransactionSourceInboundCheckCurrencyUsd InterestPaymentSimulationResultTransactionSourceInboundCheckCurrency = "USD"
)

// A Inbound International ACH Transfer object. This field will be present in the
// JSON response if and only if `category` is equal to
// `inbound_international_ach_transfer`.
type InterestPaymentSimulationResultTransactionSourceInboundInternationalACHTransfer struct {
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
	JSON                                                   interestPaymentSimulationResultTransactionSourceInboundInternationalACHTransferJSON
}

// interestPaymentSimulationResultTransactionSourceInboundInternationalACHTransferJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceInboundInternationalACHTransfer]
type interestPaymentSimulationResultTransactionSourceInboundInternationalACHTransferJSON struct {
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
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real Time Payments transfer.
	Currency InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency `json:"currency,required"`
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
	JSON                      interestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON
}

// interestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation]
type interestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON struct {
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

func (r *InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
// currency. This will always be "USD" for a Real Time Payments transfer.
type InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency string

const (
	// Canadian Dollar (CAD)
	InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyCad InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CAD"
	// Swiss Franc (CHF)
	InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyChf InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CHF"
	// Euro (EUR)
	InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyEur InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "EUR"
	// British Pound (GBP)
	InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyGbp InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "GBP"
	// Japanese Yen (JPY)
	InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyJpy InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "JPY"
	// US Dollar (USD)
	InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyUsd InterestPaymentSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "USD"
)

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
	JSON                       interestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentReversalJSON
}

// interestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentReversalJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal]
type interestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentReversalJSON struct {
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

func (r *InterestPaymentSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal) UnmarshalJSON(data []byte) (err error) {
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
	JSON           interestPaymentSimulationResultTransactionSourceInboundWireReversalJSON
}

// interestPaymentSimulationResultTransactionSourceInboundWireReversalJSON contains
// the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceInboundWireReversal]
type interestPaymentSimulationResultTransactionSourceInboundWireReversalJSON struct {
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
	OriginatorToBeneficiaryInformation      string `json:"originator_to_beneficiary_information,required,nullable"`
	OriginatorToBeneficiaryInformationLine1 string `json:"originator_to_beneficiary_information_line1,required,nullable"`
	OriginatorToBeneficiaryInformationLine2 string `json:"originator_to_beneficiary_information_line2,required,nullable"`
	OriginatorToBeneficiaryInformationLine3 string `json:"originator_to_beneficiary_information_line3,required,nullable"`
	OriginatorToBeneficiaryInformationLine4 string `json:"originator_to_beneficiary_information_line4,required,nullable"`
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
	OriginatorToBeneficiaryInformation      apijson.Field
	OriginatorToBeneficiaryInformationLine1 apijson.Field
	OriginatorToBeneficiaryInformationLine2 apijson.Field
	OriginatorToBeneficiaryInformationLine3 apijson.Field
	OriginatorToBeneficiaryInformationLine4 apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceInboundWireTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Interest Payment object. This field will be present in the JSON response if
// and only if `category` is equal to `interest_payment`.
type InterestPaymentSimulationResultTransactionSourceInterestPayment struct {
	// The account on which the interest was accrued.
	AccruedOnAccountID string `json:"accrued_on_account_id,required,nullable"`
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrency `json:"currency,required"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time `json:"period_start,required" format:"date-time"`
	JSON        interestPaymentSimulationResultTransactionSourceInterestPaymentJSON
}

// interestPaymentSimulationResultTransactionSourceInterestPaymentJSON contains the
// JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceInterestPayment]
type interestPaymentSimulationResultTransactionSourceInterestPaymentJSON struct {
	AccruedOnAccountID apijson.Field
	Amount             apijson.Field
	Currency           apijson.Field
	PeriodEnd          apijson.Field
	PeriodStart        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InterestPaymentSimulationResultTransactionSourceInterestPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrency string

const (
	// Canadian Dollar (CAD)
	InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrencyCad InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrency = "CAD"
	// Swiss Franc (CHF)
	InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrencyChf InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrency = "CHF"
	// Euro (EUR)
	InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrencyEur InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrency = "EUR"
	// British Pound (GBP)
	InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrencyGbp InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrency = "GBP"
	// Japanese Yen (JPY)
	InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrencyJpy InterestPaymentSimulationResultTransactionSourceInterestPaymentCurrency = "JPY"
	// US Dollar (USD)
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

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type InterestPaymentSimulationResultTransactionSourceInternalSourceCurrency string

const (
	// Canadian Dollar (CAD)
	InterestPaymentSimulationResultTransactionSourceInternalSourceCurrencyCad InterestPaymentSimulationResultTransactionSourceInternalSourceCurrency = "CAD"
	// Swiss Franc (CHF)
	InterestPaymentSimulationResultTransactionSourceInternalSourceCurrencyChf InterestPaymentSimulationResultTransactionSourceInternalSourceCurrency = "CHF"
	// Euro (EUR)
	InterestPaymentSimulationResultTransactionSourceInternalSourceCurrencyEur InterestPaymentSimulationResultTransactionSourceInternalSourceCurrency = "EUR"
	// British Pound (GBP)
	InterestPaymentSimulationResultTransactionSourceInternalSourceCurrencyGbp InterestPaymentSimulationResultTransactionSourceInternalSourceCurrency = "GBP"
	// Japanese Yen (JPY)
	InterestPaymentSimulationResultTransactionSourceInternalSourceCurrencyJpy InterestPaymentSimulationResultTransactionSourceInternalSourceCurrency = "JPY"
	// US Dollar (USD)
	InterestPaymentSimulationResultTransactionSourceInternalSourceCurrencyUsd InterestPaymentSimulationResultTransactionSourceInternalSourceCurrency = "USD"
)

type InterestPaymentSimulationResultTransactionSourceInternalSourceReason string

const (
	// Account closure
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonAccountClosure InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "account_closure"
	// Bank migration
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonBankMigration InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "bank_migration"
	// Cashback
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonCashback InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "cashback"
	// Collection receivable
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonCollectionReceivable InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "collection_receivable"
	// Empyreal adjustment
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonEmpyrealAdjustment InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "empyreal_adjustment"
	// Error
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonError InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "error"
	// Error correction
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonErrorCorrection InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "error_correction"
	// Fees
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonFees InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "fees"
	// Interest
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonInterest InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "interest"
	// Negative balance forgiveness
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonNegativeBalanceForgiveness InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "negative_balance_forgiveness"
	// Sample funds
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonSampleFunds InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "sample_funds"
	// Sample funds return
	InterestPaymentSimulationResultTransactionSourceInternalSourceReasonSampleFundsReturn InterestPaymentSimulationResultTransactionSourceInternalSourceReason = "sample_funds_return"
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

// A Wire Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_intention`.
type InterestPaymentSimulationResultTransactionSourceWireTransferIntention struct {
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient string `json:"message_to_recipient,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	TransferID    string `json:"transfer_id,required"`
	JSON          interestPaymentSimulationResultTransactionSourceWireTransferIntentionJSON
}

// interestPaymentSimulationResultTransactionSourceWireTransferIntentionJSON
// contains the JSON metadata for the struct
// [InterestPaymentSimulationResultTransactionSourceWireTransferIntention]
type interestPaymentSimulationResultTransactionSourceWireTransferIntentionJSON struct {
	AccountNumber      apijson.Field
	Amount             apijson.Field
	MessageToRecipient apijson.Field
	RoutingNumber      apijson.Field
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

// A constant representing the object's type. For this resource it will always be
// `transaction`.
type InterestPaymentSimulationResultTransactionType string

const (
	InterestPaymentSimulationResultTransactionTypeTransaction InterestPaymentSimulationResultTransactionType = "transaction"
)

// A constant representing the object's type. For this resource it will always be
// `interest_payment_simulation_result`.
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
