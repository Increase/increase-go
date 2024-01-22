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
	JSON wireTransferSimulationJSON `json:"-"`
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
	// Transaction occurred.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transaction's
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
	JSON wireTransferSimulationTransactionJSON `json:"-"`
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
// Transaction's currency. This will match the currency on the Transaction's
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
	// An Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention WireTransferSimulationTransactionSourceAccountTransferIntention `json:"account_transfer_intention,required,nullable"`
	// An ACH Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention WireTransferSimulationTransactionSourceACHTransferIntention `json:"ach_transfer_intention,required,nullable"`
	// An ACH Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection WireTransferSimulationTransactionSourceACHTransferRejection `json:"ach_transfer_rejection,required,nullable"`
	// An ACH Transfer Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_return`.
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
	// The type of the resource. We may add additional possible values for this enum
	// over time; your application should be able to handle such additions gracefully.
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
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest `json:"check_transfer_stop_payment_request,required,nullable"`
	// A Fee Payment object. This field will be present in the JSON response if and
	// only if `category` is equal to `fee_payment`.
	FeePayment WireTransferSimulationTransactionSourceFeePayment `json:"fee_payment,required,nullable"`
	// An Inbound ACH Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer WireTransferSimulationTransactionSourceInboundACHTransfer `json:"inbound_ach_transfer,required,nullable"`
	// An Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck WireTransferSimulationTransactionSourceInboundCheck `json:"inbound_check,required,nullable"`
	// An Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer WireTransferSimulationTransactionSourceInboundInternationalACHTransfer `json:"inbound_international_ach_transfer,required,nullable"`
	// An Inbound Real-Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation `json:"inbound_real_time_payments_transfer_confirmation,required,nullable"`
	// An Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment WireTransferSimulationTransactionSourceInboundWireDrawdownPayment `json:"inbound_wire_drawdown_payment,required,nullable"`
	// An Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal `json:"inbound_wire_drawdown_payment_reversal,required,nullable"`
	// An Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal WireTransferSimulationTransactionSourceInboundWireReversal `json:"inbound_wire_reversal,required,nullable"`
	// An Inbound Wire Transfer Intention object. This field will be present in the
	// JSON response if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer WireTransferSimulationTransactionSourceInboundWireTransfer `json:"inbound_wire_transfer,required,nullable"`
	// An Interest Payment object. This field will be present in the JSON response if
	// and only if `category` is equal to `interest_payment`.
	InterestPayment WireTransferSimulationTransactionSourceInterestPayment `json:"interest_payment,required,nullable"`
	// An Internal Source object. This field will be present in the JSON response if
	// and only if `category` is equal to `internal_source`.
	InternalSource WireTransferSimulationTransactionSourceInternalSource `json:"internal_source,required,nullable"`
	// A Real-Time Payments Transfer Acknowledgement object. This field will be present
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
	JSON                  wireTransferSimulationTransactionSourceJSON                  `json:"-"`
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

// An Account Transfer Intention object. This field will be present in the JSON
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
	TransferID string                                                              `json:"transfer_id,required"`
	JSON       wireTransferSimulationTransactionSourceAccountTransferIntentionJSON `json:"-"`
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

// An ACH Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_transfer_intention`.
type WireTransferSimulationTransactionSourceACHTransferIntention struct {
	// The account number for the destination account.
	AccountNumber string `json:"account_number,required"`
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber string `json:"routing_number,required"`
	// A description set when the ACH Transfer was created.
	StatementDescriptor string `json:"statement_descriptor,required"`
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string                                                          `json:"transfer_id,required"`
	JSON       wireTransferSimulationTransactionSourceACHTransferIntentionJSON `json:"-"`
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

// An ACH Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_transfer_rejection`.
type WireTransferSimulationTransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string                                                          `json:"transfer_id,required"`
	JSON       wireTransferSimulationTransactionSourceACHTransferRejectionJSON `json:"-"`
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

// An ACH Transfer Return object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_return`.
type WireTransferSimulationTransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The three character ACH return code, in the range R01 to R85.
	RawReturnReasonCode string `json:"raw_return_reason_code,required"`
	// Why the ACH Transfer was returned. This reason code is sent by the receiving
	// bank back to Increase.
	ReturnReasonCode WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode `json:"return_reason_code,required"`
	// The identifier of the Transaction associated with this return.
	TransactionID string `json:"transaction_id,required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID string                                                       `json:"transfer_id,required"`
	JSON       wireTransferSimulationTransactionSourceACHTransferReturnJSON `json:"-"`
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

// Why the ACH Transfer was returned. This reason code is sent by the receiving
// bank back to Increase.
type WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode string

const (
	// Code R01. Insufficient funds in the receiving account. Sometimes abbreviated to
	// NSF.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInsufficientFund WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "insufficient_fund"
	// Code R03. The account does not exist or the receiving bank was unable to locate
	// it.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeNoAccount WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "no_account"
	// Code R02. The account is closed at the receiving bank.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAccountClosed WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "account_closed"
	// Code R04. The account number is invalid at the receiving bank.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	// Code R16. The account at the receiving bank was frozen per the Office of Foreign
	// Assets Control.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	// Code R23. The receiving bank account refused a credit transfer.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	// Code R05. The receiving bank rejected because of an incorrect Standard Entry
	// Class code.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	// Code R29. The corporate customer at the receiving bank reversed the transfer.
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
	// Code R10. The customer at the receiving bank reversed the transfer.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	// Code R19. The amount field is incorrect or too large.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAmountFieldError WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "amount_field_error"
	// Code R07. The customer at the receiving institution informed their bank that
	// they have revoked authorization for a previously authorized transfer.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	// Code R13. The routing number is invalid.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	// Code R17. The receiving bank is unable to process a field in the transfer.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeFileRecordEditCriteria WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	// Code R45. The individual name field was invalid.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualName WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	// Code R06. The originating financial institution asked for this transfer to be
	// returned. The receiving bank is complying with the request.
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	// Code R34. The receiving bank's regulatory supervisor has limited their
	// participation in the ACH network.
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
	// Code R62. A rare return reason. The originating financial institution made a
	// mistake and this return corrects it.
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
	// Code R27. A rare return reason. An ACH return's trace number does not match an
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
	TransactionID string                                                           `json:"transaction_id,required"`
	JSON          wireTransferSimulationTransactionSourceCardDisputeAcceptanceJSON `json:"-"`
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
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id,required,nullable"`
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
	// Network-specific identifiers for this refund.
	NetworkIdentifiers WireTransferSimulationTransactionSourceCardRefundNetworkIdentifiers `json:"network_identifiers,required"`
	// Additional details about the card purchase, such as tax and industry-specific
	// fields.
	PurchaseDetails WireTransferSimulationTransactionSourceCardRefundPurchaseDetails `json:"purchase_details,required,nullable"`
	// The identifier of the Transaction associated with this Transaction.
	TransactionID string `json:"transaction_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type WireTransferSimulationTransactionSourceCardRefundType `json:"type,required"`
	JSON wireTransferSimulationTransactionSourceCardRefundJSON `json:"-"`
}

// wireTransferSimulationTransactionSourceCardRefundJSON contains the JSON metadata
// for the struct [WireTransferSimulationTransactionSourceCardRefund]
type wireTransferSimulationTransactionSourceCardRefundJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CardPaymentID        apijson.Field
	Currency             apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantName         apijson.Field
	MerchantState        apijson.Field
	NetworkIdentifiers   apijson.Field
	PurchaseDetails      apijson.Field
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

// Network-specific identifiers for this refund.
type WireTransferSimulationTransactionSourceCardRefundNetworkIdentifiers struct {
	// A network assigned business ID that identifies the acquirer that processed this
	// transaction.
	AcquirerBusinessID string `json:"acquirer_business_id,required"`
	// A globally unique identifier for this settlement.
	AcquirerReferenceNumber string `json:"acquirer_reference_number,required"`
	// A globally unique transaction identifier provided by the card network, used
	// across multiple life-cycle requests.
	TransactionID string                                                                  `json:"transaction_id,required,nullable"`
	JSON          wireTransferSimulationTransactionSourceCardRefundNetworkIdentifiersJSON `json:"-"`
}

// wireTransferSimulationTransactionSourceCardRefundNetworkIdentifiersJSON contains
// the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCardRefundNetworkIdentifiers]
type wireTransferSimulationTransactionSourceCardRefundNetworkIdentifiersJSON struct {
	AcquirerBusinessID      apijson.Field
	AcquirerReferenceNumber apijson.Field
	TransactionID           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCardRefundNetworkIdentifiers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional details about the card purchase, such as tax and industry-specific
// fields.
type WireTransferSimulationTransactionSourceCardRefundPurchaseDetails struct {
	// Fields specific to car rentals.
	CarRental WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRental `json:"car_rental,required,nullable"`
	// An identifier from the merchant for the customer or consumer.
	CustomerReferenceIdentifier string `json:"customer_reference_identifier,required,nullable"`
	// The state or provincial tax amount in minor units.
	LocalTaxAmount int64 `json:"local_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	LocalTaxCurrency string `json:"local_tax_currency,required,nullable"`
	// Fields specific to lodging.
	Lodging WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodging `json:"lodging,required,nullable"`
	// The national tax amount in minor units.
	NationalTaxAmount int64 `json:"national_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	NationalTaxCurrency string `json:"national_tax_currency,required,nullable"`
	// An identifier from the merchant for the purchase to the issuer and cardholder.
	PurchaseIdentifier string `json:"purchase_identifier,required,nullable"`
	// The format of the purchase identifier.
	PurchaseIdentifierFormat WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat `json:"purchase_identifier_format,required,nullable"`
	// Fields specific to travel.
	Travel WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravel `json:"travel,required,nullable"`
	JSON   wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsJSON   `json:"-"`
}

// wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsJSON contains
// the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCardRefundPurchaseDetails]
type wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsJSON struct {
	CarRental                   apijson.Field
	CustomerReferenceIdentifier apijson.Field
	LocalTaxAmount              apijson.Field
	LocalTaxCurrency            apijson.Field
	Lodging                     apijson.Field
	NationalTaxAmount           apijson.Field
	NationalTaxCurrency         apijson.Field
	PurchaseIdentifier          apijson.Field
	PurchaseIdentifierFormat    apijson.Field
	Travel                      apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCardRefundPurchaseDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Fields specific to car rentals.
type WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRental struct {
	// Code indicating the vehicle's class.
	CarClassCode string `json:"car_class_code,required,nullable"`
	// Date the customer picked up the car or, in the case of a no-show or pre-pay
	// transaction, the scheduled pick up date.
	CheckoutDate time.Time `json:"checkout_date,required,nullable" format:"date"`
	// Daily rate being charged for the vehicle.
	DailyRentalRateAmount int64 `json:"daily_rental_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the daily rental
	// rate.
	DailyRentalRateCurrency string `json:"daily_rental_rate_currency,required,nullable"`
	// Number of days the vehicle was rented.
	DaysRented int64 `json:"days_rented,required,nullable"`
	// Additional charges (gas, late fee, etc.) being billed.
	ExtraCharges WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges `json:"extra_charges,required,nullable"`
	// Fuel charges for the vehicle.
	FuelChargesAmount int64 `json:"fuel_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the fuel charges
	// assessed.
	FuelChargesCurrency string `json:"fuel_charges_currency,required,nullable"`
	// Any insurance being charged for the vehicle.
	InsuranceChargesAmount int64 `json:"insurance_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the insurance
	// charges assessed.
	InsuranceChargesCurrency string `json:"insurance_charges_currency,required,nullable"`
	// An indicator that the cardholder is being billed for a reserved vehicle that was
	// not actually rented (that is, a "no-show" charge).
	NoShowIndicator WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicator `json:"no_show_indicator,required,nullable"`
	// Charges for returning the vehicle at a different location than where it was
	// picked up.
	OneWayDropOffChargesAmount int64 `json:"one_way_drop_off_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the one-way
	// drop-off charges assessed.
	OneWayDropOffChargesCurrency string `json:"one_way_drop_off_charges_currency,required,nullable"`
	// Name of the person renting the vehicle.
	RenterName string `json:"renter_name,required,nullable"`
	// Weekly rate being charged for the vehicle.
	WeeklyRentalRateAmount int64 `json:"weekly_rental_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the weekly
	// rental rate.
	WeeklyRentalRateCurrency string                                                                        `json:"weekly_rental_rate_currency,required,nullable"`
	JSON                     wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalJSON `json:"-"`
}

// wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRental]
type wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalJSON struct {
	CarClassCode                 apijson.Field
	CheckoutDate                 apijson.Field
	DailyRentalRateAmount        apijson.Field
	DailyRentalRateCurrency      apijson.Field
	DaysRented                   apijson.Field
	ExtraCharges                 apijson.Field
	FuelChargesAmount            apijson.Field
	FuelChargesCurrency          apijson.Field
	InsuranceChargesAmount       apijson.Field
	InsuranceChargesCurrency     apijson.Field
	NoShowIndicator              apijson.Field
	OneWayDropOffChargesAmount   apijson.Field
	OneWayDropOffChargesCurrency apijson.Field
	RenterName                   apijson.Field
	WeeklyRentalRateAmount       apijson.Field
	WeeklyRentalRateCurrency     apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRental) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional charges (gas, late fee, etc.) being billed.
type WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges string

const (
	// No extra charge
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesNoExtraCharge WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "no_extra_charge"
	// Gas
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesGas WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "gas"
	// Extra mileage
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesExtraMileage WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "extra_mileage"
	// Late return
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesLateReturn WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "late_return"
	// One way service fee
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesOneWayServiceFee WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "one_way_service_fee"
	// Parking violation
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesParkingViolation WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "parking_violation"
)

// An indicator that the cardholder is being billed for a reserved vehicle that was
// not actually rented (that is, a "no-show" charge).
type WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicator string

const (
	// Not applicable
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicatorNotApplicable WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicator = "not_applicable"
	// No show for specialized vehicle
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicatorNoShowForSpecializedVehicle WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicator = "no_show_for_specialized_vehicle"
)

// Fields specific to lodging.
type WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodging struct {
	// Date the customer checked in.
	CheckInDate time.Time `json:"check_in_date,required,nullable" format:"date"`
	// Daily rate being charged for the room.
	DailyRoomRateAmount int64 `json:"daily_room_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the daily room
	// rate.
	DailyRoomRateCurrency string `json:"daily_room_rate_currency,required,nullable"`
	// Additional charges (phone, late check-out, etc.) being billed.
	ExtraCharges WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges `json:"extra_charges,required,nullable"`
	// Folio cash advances for the room.
	FolioCashAdvancesAmount int64 `json:"folio_cash_advances_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the folio cash
	// advances.
	FolioCashAdvancesCurrency string `json:"folio_cash_advances_currency,required,nullable"`
	// Food and beverage charges for the room.
	FoodBeverageChargesAmount int64 `json:"food_beverage_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the food and
	// beverage charges.
	FoodBeverageChargesCurrency string `json:"food_beverage_charges_currency,required,nullable"`
	// Indicator that the cardholder is being billed for a reserved room that was not
	// actually used.
	NoShowIndicator WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicator `json:"no_show_indicator,required,nullable"`
	// Prepaid expenses being charged for the room.
	PrepaidExpensesAmount int64 `json:"prepaid_expenses_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the prepaid
	// expenses.
	PrepaidExpensesCurrency string `json:"prepaid_expenses_currency,required,nullable"`
	// Number of nights the room was rented.
	RoomNights int64 `json:"room_nights,required,nullable"`
	// Total room tax being charged.
	TotalRoomTaxAmount int64 `json:"total_room_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the total room
	// tax.
	TotalRoomTaxCurrency string `json:"total_room_tax_currency,required,nullable"`
	// Total tax being charged for the room.
	TotalTaxAmount int64 `json:"total_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the total tax
	// assessed.
	TotalTaxCurrency string                                                                      `json:"total_tax_currency,required,nullable"`
	JSON             wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingJSON `json:"-"`
}

// wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodging]
type wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingJSON struct {
	CheckInDate                 apijson.Field
	DailyRoomRateAmount         apijson.Field
	DailyRoomRateCurrency       apijson.Field
	ExtraCharges                apijson.Field
	FolioCashAdvancesAmount     apijson.Field
	FolioCashAdvancesCurrency   apijson.Field
	FoodBeverageChargesAmount   apijson.Field
	FoodBeverageChargesCurrency apijson.Field
	NoShowIndicator             apijson.Field
	PrepaidExpensesAmount       apijson.Field
	PrepaidExpensesCurrency     apijson.Field
	RoomNights                  apijson.Field
	TotalRoomTaxAmount          apijson.Field
	TotalRoomTaxCurrency        apijson.Field
	TotalTaxAmount              apijson.Field
	TotalTaxCurrency            apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional charges (phone, late check-out, etc.) being billed.
type WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges string

const (
	// No extra charge
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesNoExtraCharge WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "no_extra_charge"
	// Restaurant
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesRestaurant WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "restaurant"
	// Gift shop
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesGiftShop WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "gift_shop"
	// Mini bar
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesMiniBar WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "mini_bar"
	// Telephone
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesTelephone WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "telephone"
	// Other
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesOther WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "other"
	// Laundry
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesLaundry WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "laundry"
)

// Indicator that the cardholder is being billed for a reserved room that was not
// actually used.
type WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicator string

const (
	// Not applicable
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicatorNotApplicable WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicator = "not_applicable"
	// No show
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicatorNoShow WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicator = "no_show"
)

// The format of the purchase identifier.
type WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat string

const (
	// Free text
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatFreeText WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "free_text"
	// Order number
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatOrderNumber WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "order_number"
	// Rental agreement number
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatRentalAgreementNumber WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "rental_agreement_number"
	// Hotel folio number
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatHotelFolioNumber WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "hotel_folio_number"
	// Invoice number
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatInvoiceNumber WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "invoice_number"
)

// Fields specific to travel.
type WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravel struct {
	// Ancillary purchases in addition to the airfare.
	Ancillary WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillary `json:"ancillary,required,nullable"`
	// Indicates the computerized reservation system used to book the ticket.
	ComputerizedReservationSystem string `json:"computerized_reservation_system,required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator `json:"credit_reason_indicator,required,nullable"`
	// Date of departure.
	DepartureDate time.Time `json:"departure_date,required,nullable" format:"date"`
	// Code for the originating city or airport.
	OriginationCityAirportCode string `json:"origination_city_airport_code,required,nullable"`
	// Name of the passenger.
	PassengerName string `json:"passenger_name,required,nullable"`
	// Indicates whether this ticket is non-refundable.
	RestrictedTicketIndicator WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicator `json:"restricted_ticket_indicator,required,nullable"`
	// Indicates why a ticket was changed.
	TicketChangeIndicator WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator `json:"ticket_change_indicator,required,nullable"`
	// Ticket number.
	TicketNumber string `json:"ticket_number,required,nullable"`
	// Code for the travel agency if the ticket was issued by a travel agency.
	TravelAgencyCode string `json:"travel_agency_code,required,nullable"`
	// Name of the travel agency if the ticket was issued by a travel agency.
	TravelAgencyName string `json:"travel_agency_name,required,nullable"`
	// Fields specific to each leg of the journey.
	TripLegs []WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLeg `json:"trip_legs,required,nullable"`
	JSON     wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelJSON      `json:"-"`
}

// wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravel]
type wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelJSON struct {
	Ancillary                     apijson.Field
	ComputerizedReservationSystem apijson.Field
	CreditReasonIndicator         apijson.Field
	DepartureDate                 apijson.Field
	OriginationCityAirportCode    apijson.Field
	PassengerName                 apijson.Field
	RestrictedTicketIndicator     apijson.Field
	TicketChangeIndicator         apijson.Field
	TicketNumber                  apijson.Field
	TravelAgencyCode              apijson.Field
	TravelAgencyName              apijson.Field
	TripLegs                      apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Ancillary purchases in addition to the airfare.
type WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillary struct {
	// If this purchase has a connection or relationship to another purchase, such as a
	// baggage fee for a passenger transport ticket, this field should contain the
	// ticket document number for the other purchase.
	ConnectedTicketDocumentNumber string `json:"connected_ticket_document_number,required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator `json:"credit_reason_indicator,required,nullable"`
	// Name of the passenger or description of the ancillary purchase.
	PassengerNameOrDescription string `json:"passenger_name_or_description,required,nullable"`
	// Additional travel charges, such as baggage fees.
	Services []WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryService `json:"services,required"`
	// Ticket document number.
	TicketDocumentNumber string                                                                              `json:"ticket_document_number,required,nullable"`
	JSON                 wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryJSON `json:"-"`
}

// wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillary]
type wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryJSON struct {
	ConnectedTicketDocumentNumber apijson.Field
	CreditReasonIndicator         apijson.Field
	PassengerNameOrDescription    apijson.Field
	Services                      apijson.Field
	TicketDocumentNumber          apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the reason for a credit to the cardholder.
type WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator string

const (
	// No credit
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorNoCredit WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "no_credit"
	// Passenger transport ancillary purchase cancellation
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket and passenger transport ancillary purchase cancellation
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	// Other
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorOther WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "other"
)

type WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryService struct {
	// Category of the ancillary service.
	Category WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory `json:"category,required,nullable"`
	// Sub-category of the ancillary service, free-form.
	SubCategory string                                                                                     `json:"sub_category,required,nullable"`
	JSON        wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServiceJSON `json:"-"`
}

// wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServiceJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryService]
type wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServiceJSON struct {
	Category    apijson.Field
	SubCategory apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Category of the ancillary service.
type WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory string

const (
	// None
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryNone WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "none"
	// Bundled service
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryBundledService WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "bundled_service"
	// Baggage fee
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryBaggageFee WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "baggage_fee"
	// Change fee
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryChangeFee WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "change_fee"
	// Cargo
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryCargo WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "cargo"
	// Carbon offset
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryCarbonOffset WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "carbon_offset"
	// Frequent flyer
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryFrequentFlyer WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "frequent_flyer"
	// Gift card
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryGiftCard WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "gift_card"
	// Ground transport
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryGroundTransport WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "ground_transport"
	// In-flight entertainment
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryInFlightEntertainment WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "in_flight_entertainment"
	// Lounge
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryLounge WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "lounge"
	// Medical
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryMedical WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "medical"
	// Meal beverage
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryMealBeverage WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "meal_beverage"
	// Other
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryOther WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "other"
	// Passenger assist fee
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryPassengerAssistFee WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "passenger_assist_fee"
	// Pets
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryPets WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "pets"
	// Seat fees
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategorySeatFees WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "seat_fees"
	// Standby
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryStandby WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "standby"
	// Service fee
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryServiceFee WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "service_fee"
	// Store
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryStore WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "store"
	// Travel service
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryTravelService WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "travel_service"
	// Unaccompanied travel
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryUnaccompaniedTravel WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "unaccompanied_travel"
	// Upgrades
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryUpgrades WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "upgrades"
	// Wi-fi
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryWifi WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "wifi"
)

// Indicates the reason for a credit to the cardholder.
type WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator string

const (
	// No credit
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorNoCredit WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "no_credit"
	// Passenger transport ancillary purchase cancellation
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket and passenger transport ancillary purchase cancellation
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket cancellation
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketCancellation WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_cancellation"
	// Other
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorOther WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "other"
	// Partial refund of airline ticket
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorPartialRefundOfAirlineTicket WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "partial_refund_of_airline_ticket"
)

// Indicates whether this ticket is non-refundable.
type WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicator string

const (
	// No restrictions
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicatorNoRestrictions WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicator = "no_restrictions"
	// Restricted non-refundable ticket
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicatorRestrictedNonRefundableTicket WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicator = "restricted_non_refundable_ticket"
)

// Indicates why a ticket was changed.
type WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator string

const (
	// None
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicatorNone WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator = "none"
	// Change to existing ticket
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicatorChangeToExistingTicket WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator = "change_to_existing_ticket"
	// New ticket
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicatorNewTicket WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator = "new_ticket"
)

type WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLeg struct {
	// Carrier code (e.g., United Airlines, Jet Blue, etc.).
	CarrierCode string `json:"carrier_code,required,nullable"`
	// Code for the destination city or airport.
	DestinationCityAirportCode string `json:"destination_city_airport_code,required,nullable"`
	// Fare basis code.
	FareBasisCode string `json:"fare_basis_code,required,nullable"`
	// Flight number.
	FlightNumber string `json:"flight_number,required,nullable"`
	// Service class (e.g., first class, business class, etc.).
	ServiceClass string `json:"service_class,required,nullable"`
	// Indicates whether a stopover is allowed on this ticket.
	StopOverCode WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode `json:"stop_over_code,required,nullable"`
	JSON         wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegJSON          `json:"-"`
}

// wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLeg]
type wireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegJSON struct {
	CarrierCode                apijson.Field
	DestinationCityAirportCode apijson.Field
	FareBasisCode              apijson.Field
	FlightNumber               apijson.Field
	ServiceClass               apijson.Field
	StopOverCode               apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLeg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates whether a stopover is allowed on this ticket.
type WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode string

const (
	// None
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCodeNone WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode = "none"
	// Stop over allowed
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCodeStopOverAllowed WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_allowed"
	// Stop over not allowed
	WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCodeStopOverNotAllowed WireTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_not_allowed"
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
	TransactedOnAccountID string                                                        `json:"transacted_on_account_id,required,nullable"`
	JSON                  wireTransferSimulationTransactionSourceCardRevenuePaymentJSON `json:"-"`
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
	// The Card Authorization that was created prior to this Card Settlement, if one
	// exists.
	CardAuthorization string `json:"card_authorization,required,nullable"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id,required,nullable"`
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
	// Network-specific identifiers for this refund.
	NetworkIdentifiers WireTransferSimulationTransactionSourceCardSettlementNetworkIdentifiers `json:"network_identifiers,required"`
	// The identifier of the Pending Transaction associated with this Transaction.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// The amount in the minor unit of the transaction's presentment currency.
	PresentmentAmount int64 `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's presentment currency.
	PresentmentCurrency string `json:"presentment_currency,required"`
	// Additional details about the card purchase, such as tax and industry-specific
	// fields.
	PurchaseDetails WireTransferSimulationTransactionSourceCardSettlementPurchaseDetails `json:"purchase_details,required,nullable"`
	// The identifier of the Transaction associated with this Transaction.
	TransactionID string `json:"transaction_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_settlement`.
	Type WireTransferSimulationTransactionSourceCardSettlementType `json:"type,required"`
	JSON wireTransferSimulationTransactionSourceCardSettlementJSON `json:"-"`
}

// wireTransferSimulationTransactionSourceCardSettlementJSON contains the JSON
// metadata for the struct [WireTransferSimulationTransactionSourceCardSettlement]
type wireTransferSimulationTransactionSourceCardSettlementJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CardAuthorization    apijson.Field
	CardPaymentID        apijson.Field
	Currency             apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantName         apijson.Field
	MerchantState        apijson.Field
	NetworkIdentifiers   apijson.Field
	PendingTransactionID apijson.Field
	PresentmentAmount    apijson.Field
	PresentmentCurrency  apijson.Field
	PurchaseDetails      apijson.Field
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

// Network-specific identifiers for this refund.
type WireTransferSimulationTransactionSourceCardSettlementNetworkIdentifiers struct {
	// A network assigned business ID that identifies the acquirer that processed this
	// transaction.
	AcquirerBusinessID string `json:"acquirer_business_id,required"`
	// A globally unique identifier for this settlement.
	AcquirerReferenceNumber string `json:"acquirer_reference_number,required"`
	// A globally unique transaction identifier provided by the card network, used
	// across multiple life-cycle requests.
	TransactionID string                                                                      `json:"transaction_id,required,nullable"`
	JSON          wireTransferSimulationTransactionSourceCardSettlementNetworkIdentifiersJSON `json:"-"`
}

// wireTransferSimulationTransactionSourceCardSettlementNetworkIdentifiersJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCardSettlementNetworkIdentifiers]
type wireTransferSimulationTransactionSourceCardSettlementNetworkIdentifiersJSON struct {
	AcquirerBusinessID      apijson.Field
	AcquirerReferenceNumber apijson.Field
	TransactionID           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCardSettlementNetworkIdentifiers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional details about the card purchase, such as tax and industry-specific
// fields.
type WireTransferSimulationTransactionSourceCardSettlementPurchaseDetails struct {
	// Fields specific to car rentals.
	CarRental WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRental `json:"car_rental,required,nullable"`
	// An identifier from the merchant for the customer or consumer.
	CustomerReferenceIdentifier string `json:"customer_reference_identifier,required,nullable"`
	// The state or provincial tax amount in minor units.
	LocalTaxAmount int64 `json:"local_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	LocalTaxCurrency string `json:"local_tax_currency,required,nullable"`
	// Fields specific to lodging.
	Lodging WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodging `json:"lodging,required,nullable"`
	// The national tax amount in minor units.
	NationalTaxAmount int64 `json:"national_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	NationalTaxCurrency string `json:"national_tax_currency,required,nullable"`
	// An identifier from the merchant for the purchase to the issuer and cardholder.
	PurchaseIdentifier string `json:"purchase_identifier,required,nullable"`
	// The format of the purchase identifier.
	PurchaseIdentifierFormat WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat `json:"purchase_identifier_format,required,nullable"`
	// Fields specific to travel.
	Travel WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravel `json:"travel,required,nullable"`
	JSON   wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsJSON   `json:"-"`
}

// wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCardSettlementPurchaseDetails]
type wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsJSON struct {
	CarRental                   apijson.Field
	CustomerReferenceIdentifier apijson.Field
	LocalTaxAmount              apijson.Field
	LocalTaxCurrency            apijson.Field
	Lodging                     apijson.Field
	NationalTaxAmount           apijson.Field
	NationalTaxCurrency         apijson.Field
	PurchaseIdentifier          apijson.Field
	PurchaseIdentifierFormat    apijson.Field
	Travel                      apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCardSettlementPurchaseDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Fields specific to car rentals.
type WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRental struct {
	// Code indicating the vehicle's class.
	CarClassCode string `json:"car_class_code,required,nullable"`
	// Date the customer picked up the car or, in the case of a no-show or pre-pay
	// transaction, the scheduled pick up date.
	CheckoutDate time.Time `json:"checkout_date,required,nullable" format:"date"`
	// Daily rate being charged for the vehicle.
	DailyRentalRateAmount int64 `json:"daily_rental_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the daily rental
	// rate.
	DailyRentalRateCurrency string `json:"daily_rental_rate_currency,required,nullable"`
	// Number of days the vehicle was rented.
	DaysRented int64 `json:"days_rented,required,nullable"`
	// Additional charges (gas, late fee, etc.) being billed.
	ExtraCharges WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges `json:"extra_charges,required,nullable"`
	// Fuel charges for the vehicle.
	FuelChargesAmount int64 `json:"fuel_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the fuel charges
	// assessed.
	FuelChargesCurrency string `json:"fuel_charges_currency,required,nullable"`
	// Any insurance being charged for the vehicle.
	InsuranceChargesAmount int64 `json:"insurance_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the insurance
	// charges assessed.
	InsuranceChargesCurrency string `json:"insurance_charges_currency,required,nullable"`
	// An indicator that the cardholder is being billed for a reserved vehicle that was
	// not actually rented (that is, a "no-show" charge).
	NoShowIndicator WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicator `json:"no_show_indicator,required,nullable"`
	// Charges for returning the vehicle at a different location than where it was
	// picked up.
	OneWayDropOffChargesAmount int64 `json:"one_way_drop_off_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the one-way
	// drop-off charges assessed.
	OneWayDropOffChargesCurrency string `json:"one_way_drop_off_charges_currency,required,nullable"`
	// Name of the person renting the vehicle.
	RenterName string `json:"renter_name,required,nullable"`
	// Weekly rate being charged for the vehicle.
	WeeklyRentalRateAmount int64 `json:"weekly_rental_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the weekly
	// rental rate.
	WeeklyRentalRateCurrency string                                                                            `json:"weekly_rental_rate_currency,required,nullable"`
	JSON                     wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalJSON `json:"-"`
}

// wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRental]
type wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalJSON struct {
	CarClassCode                 apijson.Field
	CheckoutDate                 apijson.Field
	DailyRentalRateAmount        apijson.Field
	DailyRentalRateCurrency      apijson.Field
	DaysRented                   apijson.Field
	ExtraCharges                 apijson.Field
	FuelChargesAmount            apijson.Field
	FuelChargesCurrency          apijson.Field
	InsuranceChargesAmount       apijson.Field
	InsuranceChargesCurrency     apijson.Field
	NoShowIndicator              apijson.Field
	OneWayDropOffChargesAmount   apijson.Field
	OneWayDropOffChargesCurrency apijson.Field
	RenterName                   apijson.Field
	WeeklyRentalRateAmount       apijson.Field
	WeeklyRentalRateCurrency     apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRental) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional charges (gas, late fee, etc.) being billed.
type WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges string

const (
	// No extra charge
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesNoExtraCharge WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "no_extra_charge"
	// Gas
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesGas WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "gas"
	// Extra mileage
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesExtraMileage WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "extra_mileage"
	// Late return
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesLateReturn WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "late_return"
	// One way service fee
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesOneWayServiceFee WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "one_way_service_fee"
	// Parking violation
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesParkingViolation WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "parking_violation"
)

// An indicator that the cardholder is being billed for a reserved vehicle that was
// not actually rented (that is, a "no-show" charge).
type WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicator string

const (
	// Not applicable
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicatorNotApplicable WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicator = "not_applicable"
	// No show for specialized vehicle
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicatorNoShowForSpecializedVehicle WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicator = "no_show_for_specialized_vehicle"
)

// Fields specific to lodging.
type WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodging struct {
	// Date the customer checked in.
	CheckInDate time.Time `json:"check_in_date,required,nullable" format:"date"`
	// Daily rate being charged for the room.
	DailyRoomRateAmount int64 `json:"daily_room_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the daily room
	// rate.
	DailyRoomRateCurrency string `json:"daily_room_rate_currency,required,nullable"`
	// Additional charges (phone, late check-out, etc.) being billed.
	ExtraCharges WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges `json:"extra_charges,required,nullable"`
	// Folio cash advances for the room.
	FolioCashAdvancesAmount int64 `json:"folio_cash_advances_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the folio cash
	// advances.
	FolioCashAdvancesCurrency string `json:"folio_cash_advances_currency,required,nullable"`
	// Food and beverage charges for the room.
	FoodBeverageChargesAmount int64 `json:"food_beverage_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the food and
	// beverage charges.
	FoodBeverageChargesCurrency string `json:"food_beverage_charges_currency,required,nullable"`
	// Indicator that the cardholder is being billed for a reserved room that was not
	// actually used.
	NoShowIndicator WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicator `json:"no_show_indicator,required,nullable"`
	// Prepaid expenses being charged for the room.
	PrepaidExpensesAmount int64 `json:"prepaid_expenses_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the prepaid
	// expenses.
	PrepaidExpensesCurrency string `json:"prepaid_expenses_currency,required,nullable"`
	// Number of nights the room was rented.
	RoomNights int64 `json:"room_nights,required,nullable"`
	// Total room tax being charged.
	TotalRoomTaxAmount int64 `json:"total_room_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the total room
	// tax.
	TotalRoomTaxCurrency string `json:"total_room_tax_currency,required,nullable"`
	// Total tax being charged for the room.
	TotalTaxAmount int64 `json:"total_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the total tax
	// assessed.
	TotalTaxCurrency string                                                                          `json:"total_tax_currency,required,nullable"`
	JSON             wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingJSON `json:"-"`
}

// wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodging]
type wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingJSON struct {
	CheckInDate                 apijson.Field
	DailyRoomRateAmount         apijson.Field
	DailyRoomRateCurrency       apijson.Field
	ExtraCharges                apijson.Field
	FolioCashAdvancesAmount     apijson.Field
	FolioCashAdvancesCurrency   apijson.Field
	FoodBeverageChargesAmount   apijson.Field
	FoodBeverageChargesCurrency apijson.Field
	NoShowIndicator             apijson.Field
	PrepaidExpensesAmount       apijson.Field
	PrepaidExpensesCurrency     apijson.Field
	RoomNights                  apijson.Field
	TotalRoomTaxAmount          apijson.Field
	TotalRoomTaxCurrency        apijson.Field
	TotalTaxAmount              apijson.Field
	TotalTaxCurrency            apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional charges (phone, late check-out, etc.) being billed.
type WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges string

const (
	// No extra charge
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesNoExtraCharge WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "no_extra_charge"
	// Restaurant
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesRestaurant WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "restaurant"
	// Gift shop
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesGiftShop WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "gift_shop"
	// Mini bar
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesMiniBar WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "mini_bar"
	// Telephone
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesTelephone WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "telephone"
	// Other
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesOther WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "other"
	// Laundry
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesLaundry WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "laundry"
)

// Indicator that the cardholder is being billed for a reserved room that was not
// actually used.
type WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicator string

const (
	// Not applicable
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicatorNotApplicable WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicator = "not_applicable"
	// No show
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicatorNoShow WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicator = "no_show"
)

// The format of the purchase identifier.
type WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat string

const (
	// Free text
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatFreeText WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "free_text"
	// Order number
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatOrderNumber WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "order_number"
	// Rental agreement number
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatRentalAgreementNumber WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "rental_agreement_number"
	// Hotel folio number
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatHotelFolioNumber WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "hotel_folio_number"
	// Invoice number
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatInvoiceNumber WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "invoice_number"
)

// Fields specific to travel.
type WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravel struct {
	// Ancillary purchases in addition to the airfare.
	Ancillary WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillary `json:"ancillary,required,nullable"`
	// Indicates the computerized reservation system used to book the ticket.
	ComputerizedReservationSystem string `json:"computerized_reservation_system,required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator `json:"credit_reason_indicator,required,nullable"`
	// Date of departure.
	DepartureDate time.Time `json:"departure_date,required,nullable" format:"date"`
	// Code for the originating city or airport.
	OriginationCityAirportCode string `json:"origination_city_airport_code,required,nullable"`
	// Name of the passenger.
	PassengerName string `json:"passenger_name,required,nullable"`
	// Indicates whether this ticket is non-refundable.
	RestrictedTicketIndicator WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator `json:"restricted_ticket_indicator,required,nullable"`
	// Indicates why a ticket was changed.
	TicketChangeIndicator WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator `json:"ticket_change_indicator,required,nullable"`
	// Ticket number.
	TicketNumber string `json:"ticket_number,required,nullable"`
	// Code for the travel agency if the ticket was issued by a travel agency.
	TravelAgencyCode string `json:"travel_agency_code,required,nullable"`
	// Name of the travel agency if the ticket was issued by a travel agency.
	TravelAgencyName string `json:"travel_agency_name,required,nullable"`
	// Fields specific to each leg of the journey.
	TripLegs []WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLeg `json:"trip_legs,required,nullable"`
	JSON     wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelJSON      `json:"-"`
}

// wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravel]
type wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelJSON struct {
	Ancillary                     apijson.Field
	ComputerizedReservationSystem apijson.Field
	CreditReasonIndicator         apijson.Field
	DepartureDate                 apijson.Field
	OriginationCityAirportCode    apijson.Field
	PassengerName                 apijson.Field
	RestrictedTicketIndicator     apijson.Field
	TicketChangeIndicator         apijson.Field
	TicketNumber                  apijson.Field
	TravelAgencyCode              apijson.Field
	TravelAgencyName              apijson.Field
	TripLegs                      apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Ancillary purchases in addition to the airfare.
type WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillary struct {
	// If this purchase has a connection or relationship to another purchase, such as a
	// baggage fee for a passenger transport ticket, this field should contain the
	// ticket document number for the other purchase.
	ConnectedTicketDocumentNumber string `json:"connected_ticket_document_number,required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator `json:"credit_reason_indicator,required,nullable"`
	// Name of the passenger or description of the ancillary purchase.
	PassengerNameOrDescription string `json:"passenger_name_or_description,required,nullable"`
	// Additional travel charges, such as baggage fees.
	Services []WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryService `json:"services,required"`
	// Ticket document number.
	TicketDocumentNumber string                                                                                  `json:"ticket_document_number,required,nullable"`
	JSON                 wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryJSON `json:"-"`
}

// wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillary]
type wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryJSON struct {
	ConnectedTicketDocumentNumber apijson.Field
	CreditReasonIndicator         apijson.Field
	PassengerNameOrDescription    apijson.Field
	Services                      apijson.Field
	TicketDocumentNumber          apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the reason for a credit to the cardholder.
type WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator string

const (
	// No credit
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorNoCredit WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "no_credit"
	// Passenger transport ancillary purchase cancellation
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket and passenger transport ancillary purchase cancellation
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	// Other
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorOther WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "other"
)

type WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryService struct {
	// Category of the ancillary service.
	Category WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory `json:"category,required,nullable"`
	// Sub-category of the ancillary service, free-form.
	SubCategory string                                                                                         `json:"sub_category,required,nullable"`
	JSON        wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServiceJSON `json:"-"`
}

// wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServiceJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryService]
type wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServiceJSON struct {
	Category    apijson.Field
	SubCategory apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Category of the ancillary service.
type WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory string

const (
	// None
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryNone WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "none"
	// Bundled service
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryBundledService WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "bundled_service"
	// Baggage fee
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryBaggageFee WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "baggage_fee"
	// Change fee
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryChangeFee WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "change_fee"
	// Cargo
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryCargo WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "cargo"
	// Carbon offset
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryCarbonOffset WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "carbon_offset"
	// Frequent flyer
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryFrequentFlyer WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "frequent_flyer"
	// Gift card
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryGiftCard WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "gift_card"
	// Ground transport
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryGroundTransport WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "ground_transport"
	// In-flight entertainment
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryInFlightEntertainment WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "in_flight_entertainment"
	// Lounge
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryLounge WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "lounge"
	// Medical
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryMedical WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "medical"
	// Meal beverage
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryMealBeverage WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "meal_beverage"
	// Other
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryOther WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "other"
	// Passenger assist fee
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryPassengerAssistFee WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "passenger_assist_fee"
	// Pets
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryPets WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "pets"
	// Seat fees
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategorySeatFees WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "seat_fees"
	// Standby
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryStandby WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "standby"
	// Service fee
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryServiceFee WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "service_fee"
	// Store
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryStore WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "store"
	// Travel service
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryTravelService WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "travel_service"
	// Unaccompanied travel
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryUnaccompaniedTravel WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "unaccompanied_travel"
	// Upgrades
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryUpgrades WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "upgrades"
	// Wi-fi
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryWifi WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "wifi"
)

// Indicates the reason for a credit to the cardholder.
type WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator string

const (
	// No credit
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorNoCredit WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "no_credit"
	// Passenger transport ancillary purchase cancellation
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket and passenger transport ancillary purchase cancellation
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket cancellation
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketCancellation WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_cancellation"
	// Other
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorOther WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "other"
	// Partial refund of airline ticket
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorPartialRefundOfAirlineTicket WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "partial_refund_of_airline_ticket"
)

// Indicates whether this ticket is non-refundable.
type WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator string

const (
	// No restrictions
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicatorNoRestrictions WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator = "no_restrictions"
	// Restricted non-refundable ticket
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicatorRestrictedNonRefundableTicket WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator = "restricted_non_refundable_ticket"
)

// Indicates why a ticket was changed.
type WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator string

const (
	// None
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicatorNone WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "none"
	// Change to existing ticket
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicatorChangeToExistingTicket WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "change_to_existing_ticket"
	// New ticket
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicatorNewTicket WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "new_ticket"
)

type WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLeg struct {
	// Carrier code (e.g., United Airlines, Jet Blue, etc.).
	CarrierCode string `json:"carrier_code,required,nullable"`
	// Code for the destination city or airport.
	DestinationCityAirportCode string `json:"destination_city_airport_code,required,nullable"`
	// Fare basis code.
	FareBasisCode string `json:"fare_basis_code,required,nullable"`
	// Flight number.
	FlightNumber string `json:"flight_number,required,nullable"`
	// Service class (e.g., first class, business class, etc.).
	ServiceClass string `json:"service_class,required,nullable"`
	// Indicates whether a stopover is allowed on this ticket.
	StopOverCode WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode `json:"stop_over_code,required,nullable"`
	JSON         wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegJSON          `json:"-"`
}

// wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLeg]
type wireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegJSON struct {
	CarrierCode                apijson.Field
	DestinationCityAirportCode apijson.Field
	FareBasisCode              apijson.Field
	FlightNumber               apijson.Field
	ServiceClass               apijson.Field
	StopOverCode               apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLeg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates whether a stopover is allowed on this ticket.
type WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode string

const (
	// None
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeNone WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode = "none"
	// Stop over allowed
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeStopOverAllowed WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_allowed"
	// Stop over not allowed
	WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeStopOverNotAllowed WireTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_not_allowed"
)

// A constant representing the object's type. For this resource it will always be
// `card_settlement`.
type WireTransferSimulationTransactionSourceCardSettlementType string

const (
	WireTransferSimulationTransactionSourceCardSettlementTypeCardSettlement WireTransferSimulationTransactionSourceCardSettlementType = "card_settlement"
)

// The type of the resource. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
type WireTransferSimulationTransactionSourceCategory string

const (
	// Account Transfer Intention: details will be under the
	// `account_transfer_intention` object.
	WireTransferSimulationTransactionSourceCategoryAccountTransferIntention WireTransferSimulationTransactionSourceCategory = "account_transfer_intention"
	// ACH Transfer Intention: details will be under the `ach_transfer_intention`
	// object.
	WireTransferSimulationTransactionSourceCategoryACHTransferIntention WireTransferSimulationTransactionSourceCategory = "ach_transfer_intention"
	// ACH Transfer Rejection: details will be under the `ach_transfer_rejection`
	// object.
	WireTransferSimulationTransactionSourceCategoryACHTransferRejection WireTransferSimulationTransactionSourceCategory = "ach_transfer_rejection"
	// ACH Transfer Return: details will be under the `ach_transfer_return` object.
	WireTransferSimulationTransactionSourceCategoryACHTransferReturn WireTransferSimulationTransactionSourceCategory = "ach_transfer_return"
	// Card Dispute Acceptance: details will be under the `card_dispute_acceptance`
	// object.
	WireTransferSimulationTransactionSourceCategoryCardDisputeAcceptance WireTransferSimulationTransactionSourceCategory = "card_dispute_acceptance"
	// Card Refund: details will be under the `card_refund` object.
	WireTransferSimulationTransactionSourceCategoryCardRefund WireTransferSimulationTransactionSourceCategory = "card_refund"
	// Card Settlement: details will be under the `card_settlement` object.
	WireTransferSimulationTransactionSourceCategoryCardSettlement WireTransferSimulationTransactionSourceCategory = "card_settlement"
	// Card Revenue Payment: details will be under the `card_revenue_payment` object.
	WireTransferSimulationTransactionSourceCategoryCardRevenuePayment WireTransferSimulationTransactionSourceCategory = "card_revenue_payment"
	// Check Deposit Acceptance: details will be under the `check_deposit_acceptance`
	// object.
	WireTransferSimulationTransactionSourceCategoryCheckDepositAcceptance WireTransferSimulationTransactionSourceCategory = "check_deposit_acceptance"
	// Check Deposit Return: details will be under the `check_deposit_return` object.
	WireTransferSimulationTransactionSourceCategoryCheckDepositReturn WireTransferSimulationTransactionSourceCategory = "check_deposit_return"
	// Check Transfer Deposit: details will be under the `check_transfer_deposit`
	// object.
	WireTransferSimulationTransactionSourceCategoryCheckTransferDeposit WireTransferSimulationTransactionSourceCategory = "check_transfer_deposit"
	// Check Transfer Intention: details will be under the `check_transfer_intention`
	// object.
	WireTransferSimulationTransactionSourceCategoryCheckTransferIntention WireTransferSimulationTransactionSourceCategory = "check_transfer_intention"
	// Check Transfer Stop Payment Request: details will be under the
	// `check_transfer_stop_payment_request` object.
	WireTransferSimulationTransactionSourceCategoryCheckTransferStopPaymentRequest WireTransferSimulationTransactionSourceCategory = "check_transfer_stop_payment_request"
	// Fee Payment: details will be under the `fee_payment` object.
	WireTransferSimulationTransactionSourceCategoryFeePayment WireTransferSimulationTransactionSourceCategory = "fee_payment"
	// Inbound ACH Transfer Intention: details will be under the `inbound_ach_transfer`
	// object.
	WireTransferSimulationTransactionSourceCategoryInboundACHTransfer WireTransferSimulationTransactionSourceCategory = "inbound_ach_transfer"
	// Inbound ACH Transfer Return Intention: details will be under the
	// `inbound_ach_transfer_return_intention` object.
	WireTransferSimulationTransactionSourceCategoryInboundACHTransferReturnIntention WireTransferSimulationTransactionSourceCategory = "inbound_ach_transfer_return_intention"
	// Inbound Check: details will be under the `inbound_check` object.
	WireTransferSimulationTransactionSourceCategoryInboundCheck WireTransferSimulationTransactionSourceCategory = "inbound_check"
	// Inbound International ACH Transfer: details will be under the
	// `inbound_international_ach_transfer` object.
	WireTransferSimulationTransactionSourceCategoryInboundInternationalACHTransfer WireTransferSimulationTransactionSourceCategory = "inbound_international_ach_transfer"
	// Inbound Real-Time Payments Transfer Confirmation: details will be under the
	// `inbound_real_time_payments_transfer_confirmation` object.
	WireTransferSimulationTransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation WireTransferSimulationTransactionSourceCategory = "inbound_real_time_payments_transfer_confirmation"
	// Inbound Wire Drawdown Payment Reversal: details will be under the
	// `inbound_wire_drawdown_payment_reversal` object.
	WireTransferSimulationTransactionSourceCategoryInboundWireDrawdownPaymentReversal WireTransferSimulationTransactionSourceCategory = "inbound_wire_drawdown_payment_reversal"
	// Inbound Wire Drawdown Payment: details will be under the
	// `inbound_wire_drawdown_payment` object.
	WireTransferSimulationTransactionSourceCategoryInboundWireDrawdownPayment WireTransferSimulationTransactionSourceCategory = "inbound_wire_drawdown_payment"
	// Inbound Wire Reversal: details will be under the `inbound_wire_reversal` object.
	WireTransferSimulationTransactionSourceCategoryInboundWireReversal WireTransferSimulationTransactionSourceCategory = "inbound_wire_reversal"
	// Inbound Wire Transfer Intention: details will be under the
	// `inbound_wire_transfer` object.
	WireTransferSimulationTransactionSourceCategoryInboundWireTransfer WireTransferSimulationTransactionSourceCategory = "inbound_wire_transfer"
	// Interest Payment: details will be under the `interest_payment` object.
	WireTransferSimulationTransactionSourceCategoryInterestPayment WireTransferSimulationTransactionSourceCategory = "interest_payment"
	// Internal Source: details will be under the `internal_source` object.
	WireTransferSimulationTransactionSourceCategoryInternalSource WireTransferSimulationTransactionSourceCategory = "internal_source"
	// Real-Time Payments Transfer Acknowledgement: details will be under the
	// `real_time_payments_transfer_acknowledgement` object.
	WireTransferSimulationTransactionSourceCategoryRealTimePaymentsTransferAcknowledgement WireTransferSimulationTransactionSourceCategory = "real_time_payments_transfer_acknowledgement"
	// Sample Funds: details will be under the `sample_funds` object.
	WireTransferSimulationTransactionSourceCategorySampleFunds WireTransferSimulationTransactionSourceCategory = "sample_funds"
	// Wire Transfer Intention: details will be under the `wire_transfer_intention`
	// object.
	WireTransferSimulationTransactionSourceCategoryWireTransferIntention WireTransferSimulationTransactionSourceCategory = "wire_transfer_intention"
	// Wire Transfer Rejection: details will be under the `wire_transfer_rejection`
	// object.
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
	SerialNumber string                                                            `json:"serial_number,required,nullable"`
	JSON         wireTransferSimulationTransactionSourceCheckDepositAcceptanceJSON `json:"-"`
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
	Currency WireTransferSimulationTransactionSourceCheckDepositReturnCurrency `json:"currency,required"`
	// Why this check was returned by the bank holding the account it was drawn
	// against.
	ReturnReason WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason `json:"return_reason,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID string                                                        `json:"transaction_id,required"`
	JSON          wireTransferSimulationTransactionSourceCheckDepositReturnJSON `json:"-"`
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

// Why this check was returned by the bank holding the account it was drawn
// against.
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
	// The check present was either altered or fake.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonAlteredOrFictitiousItem WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "altered_or_fictitious_item"
	// The account this check is drawn on is frozen.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonFrozenOrBlockedAccount WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "frozen_or_blocked_account"
	// The check is post dated.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonPostDated WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "post_dated"
	// The endorsement was missing.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonEndorsementMissing WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "endorsement_missing"
	// The check signature was missing.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonSignatureMissing WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "signature_missing"
	// The bank suspects a stop payment will be placed.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonStopPaymentSuspect WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "stop_payment_suspect"
	// The bank cannot read the image.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonUnusableImage WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "unusable_image"
	// The check image fails the bank's security check.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonImageFailsSecurityCheck WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "image_fails_security_check"
	// The bank cannot determine the amount.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonCannotDetermineAmount WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "cannot_determine_amount"
	// The signature is inconsistent with prior signatures.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonSignatureIrregular WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "signature_irregular"
	// The check is a non-cash item and cannot be drawn against the account.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonNonCashItem WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "non_cash_item"
	// The bank is unable to process this check.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonUnableToProcess WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "unable_to_process"
	// The check exceeds the bank or customer's limit.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonItemExceedsDollarLimit WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "item_exceeds_dollar_limit"
	// The bank sold this account and no longer services this customer.
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonBranchOrAccountSold WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "branch_or_account_sold"
)

// A Check Transfer Deposit object. This field will be present in the JSON response
// if and only if `category` is equal to `check_transfer_deposit`.
type WireTransferSimulationTransactionSourceCheckTransferDeposit struct {
	// The identifier of the API File object containing an image of the back of the
	// deposited check.
	BackImageFileID string `json:"back_image_file_id,required,nullable"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// bank depositing this check. In some rare cases, this is not transmitted via
	// Check21 and the value will be null.
	BankOfFirstDepositRoutingNumber string `json:"bank_of_first_deposit_routing_number,required,nullable"`
	// When the check was deposited.
	DepositedAt time.Time `json:"deposited_at,required" format:"date-time"`
	// The identifier of the API File object containing an image of the front of the
	// deposited check.
	FrontImageFileID string `json:"front_image_file_id,required,nullable"`
	// The identifier of the Transaction object created when the check was deposited.
	TransactionID string `json:"transaction_id,required,nullable"`
	// The identifier of the Check Transfer object that was deposited.
	TransferID string `json:"transfer_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_deposit`.
	Type WireTransferSimulationTransactionSourceCheckTransferDepositType `json:"type,required"`
	JSON wireTransferSimulationTransactionSourceCheckTransferDepositJSON `json:"-"`
}

// wireTransferSimulationTransactionSourceCheckTransferDepositJSON contains the
// JSON metadata for the struct
// [WireTransferSimulationTransactionSourceCheckTransferDeposit]
type wireTransferSimulationTransactionSourceCheckTransferDepositJSON struct {
	BackImageFileID                 apijson.Field
	BankOfFirstDepositRoutingNumber apijson.Field
	DepositedAt                     apijson.Field
	FrontImageFileID                apijson.Field
	TransactionID                   apijson.Field
	TransferID                      apijson.Field
	Type                            apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
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
	TransferID string                                                            `json:"transfer_id,required"`
	JSON       wireTransferSimulationTransactionSourceCheckTransferIntentionJSON `json:"-"`
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
	JSON wireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestJSON `json:"-"`
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
	// The check was canceled by an Increase operator who will provide details
	// out-of-band.
	WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestReasonRejectedByIncrease WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestReason = "rejected_by_increase"
	// The check was not authorized.
	WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestReasonNotAuthorized WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestReason = "not_authorized"
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
	// The start of this payment's fee period, usually the first day of a month.
	FeePeriodStart time.Time                                             `json:"fee_period_start,required" format:"date"`
	JSON           wireTransferSimulationTransactionSourceFeePaymentJSON `json:"-"`
}

// wireTransferSimulationTransactionSourceFeePaymentJSON contains the JSON metadata
// for the struct [WireTransferSimulationTransactionSourceFeePayment]
type wireTransferSimulationTransactionSourceFeePaymentJSON struct {
	Amount         apijson.Field
	Currency       apijson.Field
	FeePeriodStart apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
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

// An Inbound ACH Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_ach_transfer`.
type WireTransferSimulationTransactionSourceInboundACHTransfer struct {
	// Additional information sent from the originator.
	Addenda WireTransferSimulationTransactionSourceInboundACHTransferAddenda `json:"addenda,required,nullable"`
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The description of the date of the transfer, usually in the format `YYMMDD`.
	OriginatorCompanyDescriptiveDate string `json:"originator_company_descriptive_date,required,nullable"`
	// Data set by the originator.
	OriginatorCompanyDiscretionaryData string `json:"originator_company_discretionary_data,required,nullable"`
	// An informational description of the transfer.
	OriginatorCompanyEntryDescription string `json:"originator_company_entry_description,required"`
	// An identifier for the originating company. This is generally, but not always, a
	// stable identifier across multiple transfers.
	OriginatorCompanyID string `json:"originator_company_id,required"`
	// A name set by the originator to identify themselves.
	OriginatorCompanyName string `json:"originator_company_name,required"`
	// The originator's identifier for the transfer receipient.
	ReceiverIDNumber string `json:"receiver_id_number,required,nullable"`
	// The name of the transfer recipient. This value is informational and not verified
	// by Increase.
	ReceiverName string `json:"receiver_name,required,nullable"`
	// A 15 digit number recorded in the Nacha file and available to both the
	// originating and receiving bank. Along with the amount, date, and originating
	// routing number, this can be used to identify the ACH transfer at either bank.
	// ACH trace numbers are not unique, but are
	// [used to correlate returns](https://increase.com/documentation/ach#returns).
	TraceNumber string `json:"trace_number,required"`
	// The inbound ach transfer's identifier.
	TransferID string                                                        `json:"transfer_id,required"`
	JSON       wireTransferSimulationTransactionSourceInboundACHTransferJSON `json:"-"`
}

// wireTransferSimulationTransactionSourceInboundACHTransferJSON contains the JSON
// metadata for the struct
// [WireTransferSimulationTransactionSourceInboundACHTransfer]
type wireTransferSimulationTransactionSourceInboundACHTransferJSON struct {
	Addenda                            apijson.Field
	Amount                             apijson.Field
	OriginatorCompanyDescriptiveDate   apijson.Field
	OriginatorCompanyDiscretionaryData apijson.Field
	OriginatorCompanyEntryDescription  apijson.Field
	OriginatorCompanyID                apijson.Field
	OriginatorCompanyName              apijson.Field
	ReceiverIDNumber                   apijson.Field
	ReceiverName                       apijson.Field
	TraceNumber                        apijson.Field
	TransferID                         apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceInboundACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional information sent from the originator.
type WireTransferSimulationTransactionSourceInboundACHTransferAddenda struct {
	// The type of addendum.
	Category WireTransferSimulationTransactionSourceInboundACHTransferAddendaCategory `json:"category,required"`
	// Unstructured `payment_related_information` passed through by the originator.
	Freeform WireTransferSimulationTransactionSourceInboundACHTransferAddendaFreeform `json:"freeform,required,nullable"`
	JSON     wireTransferSimulationTransactionSourceInboundACHTransferAddendaJSON     `json:"-"`
}

// wireTransferSimulationTransactionSourceInboundACHTransferAddendaJSON contains
// the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceInboundACHTransferAddenda]
type wireTransferSimulationTransactionSourceInboundACHTransferAddendaJSON struct {
	Category    apijson.Field
	Freeform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceInboundACHTransferAddenda) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of addendum.
type WireTransferSimulationTransactionSourceInboundACHTransferAddendaCategory string

const (
	// Unstructured addendum.
	WireTransferSimulationTransactionSourceInboundACHTransferAddendaCategoryFreeform WireTransferSimulationTransactionSourceInboundACHTransferAddendaCategory = "freeform"
)

// Unstructured `payment_related_information` passed through by the originator.
type WireTransferSimulationTransactionSourceInboundACHTransferAddendaFreeform struct {
	// Each entry represents an addendum received from the originator.
	Entries []WireTransferSimulationTransactionSourceInboundACHTransferAddendaFreeformEntry `json:"entries,required"`
	JSON    wireTransferSimulationTransactionSourceInboundACHTransferAddendaFreeformJSON    `json:"-"`
}

// wireTransferSimulationTransactionSourceInboundACHTransferAddendaFreeformJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceInboundACHTransferAddendaFreeform]
type wireTransferSimulationTransactionSourceInboundACHTransferAddendaFreeformJSON struct {
	Entries     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceInboundACHTransferAddendaFreeform) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WireTransferSimulationTransactionSourceInboundACHTransferAddendaFreeformEntry struct {
	// The payment related information passed in the addendum.
	PaymentRelatedInformation string                                                                            `json:"payment_related_information,required"`
	JSON                      wireTransferSimulationTransactionSourceInboundACHTransferAddendaFreeformEntryJSON `json:"-"`
}

// wireTransferSimulationTransactionSourceInboundACHTransferAddendaFreeformEntryJSON
// contains the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceInboundACHTransferAddendaFreeformEntry]
type wireTransferSimulationTransactionSourceInboundACHTransferAddendaFreeformEntryJSON struct {
	PaymentRelatedInformation apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceInboundACHTransferAddendaFreeformEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An Inbound Check object. This field will be present in the JSON response if and
// only if `category` is equal to `inbound_check`.
type WireTransferSimulationTransactionSourceInboundCheck struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// bank depositing this check. In some rare cases, this is not transmitted via
	// Check21 and the value will be null.
	BankOfFirstDepositRoutingNumber string `json:"bank_of_first_deposit_routing_number,required,nullable"`
	// The front image of the check. This is a black and white TIFF image file.
	CheckFrontImageFileID string `json:"check_front_image_file_id,required,nullable"`
	// The number of the check. This field is set by the depositing bank and can be
	// unreliable.
	CheckNumber string `json:"check_number,required,nullable"`
	// The rear image of the check. This is a black and white TIFF image file.
	CheckRearImageFileID string `json:"check_rear_image_file_id,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency WireTransferSimulationTransactionSourceInboundCheckCurrency `json:"currency,required"`
	JSON     wireTransferSimulationTransactionSourceInboundCheckJSON     `json:"-"`
}

// wireTransferSimulationTransactionSourceInboundCheckJSON contains the JSON
// metadata for the struct [WireTransferSimulationTransactionSourceInboundCheck]
type wireTransferSimulationTransactionSourceInboundCheckJSON struct {
	Amount                          apijson.Field
	BankOfFirstDepositRoutingNumber apijson.Field
	CheckFrontImageFileID           apijson.Field
	CheckNumber                     apijson.Field
	CheckRearImageFileID            apijson.Field
	Currency                        apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
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

// An Inbound International ACH Transfer object. This field will be present in the
// JSON response if and only if `category` is equal to
// `inbound_international_ach_transfer`.
type WireTransferSimulationTransactionSourceInboundInternationalACHTransfer struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2
	// country code of the destination country.
	DestinationCountryCode string `json:"destination_country_code,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code for the
	// destination bank account.
	DestinationCurrencyCode string `json:"destination_currency_code,required"`
	// A description of how the foreign exchange rate was calculated.
	ForeignExchangeIndicator WireTransferSimulationTransactionSourceInboundInternationalACHTransferForeignExchangeIndicator `json:"foreign_exchange_indicator,required"`
	// Depending on the `foreign_exchange_reference_indicator`, an exchange rate or a
	// reference to a well-known rate.
	ForeignExchangeReference string `json:"foreign_exchange_reference,required,nullable"`
	// An instruction of how to interpret the `foreign_exchange_reference` field for
	// this Transaction.
	ForeignExchangeReferenceIndicator WireTransferSimulationTransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicator `json:"foreign_exchange_reference_indicator,required"`
	// The amount in the minor unit of the foreign payment currency. For dollars, for
	// example, this is cents.
	ForeignPaymentAmount int64 `json:"foreign_payment_amount,required"`
	// A reference number in the foreign banking infrastructure.
	ForeignTraceNumber string `json:"foreign_trace_number,required,nullable"`
	// The type of transfer. Set by the originator.
	InternationalTransactionTypeCode WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode `json:"international_transaction_type_code,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code for the
	// originating bank account.
	OriginatingCurrencyCode string `json:"originating_currency_code,required"`
	// The [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2
	// country code of the originating branch country.
	OriginatingDepositoryFinancialInstitutionBranchCountry string `json:"originating_depository_financial_institution_branch_country,required"`
	// An identifier for the originating bank. One of an International Bank Account
	// Number (IBAN) bank identifier, SWIFT Bank Identification Code (BIC), or a
	// domestic identifier like a US Routing Number.
	OriginatingDepositoryFinancialInstitutionID string `json:"originating_depository_financial_institution_id,required"`
	// An instruction of how to interpret the
	// `originating_depository_financial_institution_id` field for this Transaction.
	OriginatingDepositoryFinancialInstitutionIDQualifier WireTransferSimulationTransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifier `json:"originating_depository_financial_institution_id_qualifier,required"`
	// The name of the originating bank. Sometimes this will refer to an American bank
	// and obscure the correspondent foreign bank.
	OriginatingDepositoryFinancialInstitutionName string `json:"originating_depository_financial_institution_name,required"`
	// A portion of the originator address. This may be incomplete.
	OriginatorCity string `json:"originator_city,required"`
	// A description field set by the originator.
	OriginatorCompanyEntryDescription string `json:"originator_company_entry_description,required"`
	// A portion of the originator address. The
	// [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2 country
	// code of the originator country.
	OriginatorCountry string `json:"originator_country,required"`
	// An identifier for the originating company. This is generally stable across
	// multiple ACH transfers.
	OriginatorIdentification string `json:"originator_identification,required"`
	// Either the name of the originator or an intermediary money transmitter.
	OriginatorName string `json:"originator_name,required"`
	// A portion of the originator address. This may be incomplete.
	OriginatorPostalCode string `json:"originator_postal_code,required,nullable"`
	// A portion of the originator address. This may be incomplete.
	OriginatorStateOrProvince string `json:"originator_state_or_province,required,nullable"`
	// A portion of the originator address. This may be incomplete.
	OriginatorStreetAddress string `json:"originator_street_address,required"`
	// A description field set by the originator.
	PaymentRelatedInformation string `json:"payment_related_information,required,nullable"`
	// A description field set by the originator.
	PaymentRelatedInformation2 string `json:"payment_related_information2,required,nullable"`
	// A portion of the receiver address. This may be incomplete.
	ReceiverCity string `json:"receiver_city,required"`
	// A portion of the receiver address. The
	// [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2 country
	// code of the receiver country.
	ReceiverCountry string `json:"receiver_country,required"`
	// An identification number the originator uses for the receiver.
	ReceiverIdentificationNumber string `json:"receiver_identification_number,required,nullable"`
	// A portion of the receiver address. This may be incomplete.
	ReceiverPostalCode string `json:"receiver_postal_code,required,nullable"`
	// A portion of the receiver address. This may be incomplete.
	ReceiverStateOrProvince string `json:"receiver_state_or_province,required,nullable"`
	// A portion of the receiver address. This may be incomplete.
	ReceiverStreetAddress string `json:"receiver_street_address,required"`
	// The name of the receiver of the transfer. This is not verified by Increase.
	ReceivingCompanyOrIndividualName string `json:"receiving_company_or_individual_name,required"`
	// The [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2
	// country code of the receiving bank country.
	ReceivingDepositoryFinancialInstitutionCountry string `json:"receiving_depository_financial_institution_country,required"`
	// An identifier for the receiving bank. One of an International Bank Account
	// Number (IBAN) bank identifier, SWIFT Bank Identification Code (BIC), or a
	// domestic identifier like a US Routing Number.
	ReceivingDepositoryFinancialInstitutionID string `json:"receiving_depository_financial_institution_id,required"`
	// An instruction of how to interpret the
	// `receiving_depository_financial_institution_id` field for this Transaction.
	ReceivingDepositoryFinancialInstitutionIDQualifier WireTransferSimulationTransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifier `json:"receiving_depository_financial_institution_id_qualifier,required"`
	// The name of the receiving bank, as set by the sending financial institution.
	ReceivingDepositoryFinancialInstitutionName string `json:"receiving_depository_financial_institution_name,required"`
	// A 15 digit number recorded in the Nacha file and available to both the
	// originating and receiving bank. Along with the amount, date, and originating
	// routing number, this can be used to identify the ACH transfer at either bank.
	// ACH trace numbers are not unique, but are
	// [used to correlate returns](https://increase.com/documentation/ach#returns).
	TraceNumber string                                                                     `json:"trace_number,required"`
	JSON        wireTransferSimulationTransactionSourceInboundInternationalACHTransferJSON `json:"-"`
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

// A description of how the foreign exchange rate was calculated.
type WireTransferSimulationTransactionSourceInboundInternationalACHTransferForeignExchangeIndicator string

const (
	// The originator chose an amount in their own currency. The settled amount in USD
	// was converted using the exchange rate.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferForeignExchangeIndicatorFixedToVariable WireTransferSimulationTransactionSourceInboundInternationalACHTransferForeignExchangeIndicator = "fixed_to_variable"
	// The originator chose an amount to settle in USD. The originator's amount was
	// variable; known only after the foreign exchange conversion.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferForeignExchangeIndicatorVariableToFixed WireTransferSimulationTransactionSourceInboundInternationalACHTransferForeignExchangeIndicator = "variable_to_fixed"
	// The amount was originated and settled as a fixed amount in USD. There is no
	// foreign exchange conversion.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferForeignExchangeIndicatorFixedToFixed WireTransferSimulationTransactionSourceInboundInternationalACHTransferForeignExchangeIndicator = "fixed_to_fixed"
)

// An instruction of how to interpret the `foreign_exchange_reference` field for
// this Transaction.
type WireTransferSimulationTransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicator string

const (
	// The ACH file contains a foreign exchange rate.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicatorForeignExchangeRate WireTransferSimulationTransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicator = "foreign_exchange_rate"
	// The ACH file contains a reference to a well-known foreign exchange rate.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicatorForeignExchangeReferenceNumber WireTransferSimulationTransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicator = "foreign_exchange_reference_number"
	// There is no foreign exchange for this transfer, so the
	// `foreign_exchange_reference` field is blank.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicatorBlank WireTransferSimulationTransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicator = "blank"
)

// The type of transfer. Set by the originator.
type WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode string

const (
	// Sent as `ANN` in the Nacha file.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeAnnuity WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "annuity"
	// Sent as `BUS` in the Nacha file.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeBusinessOrCommercial WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "business_or_commercial"
	// Sent as `DEP` in the Nacha file.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeDeposit WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "deposit"
	// Sent as `LOA` in the Nacha file.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeLoan WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "loan"
	// Sent as `MIS` in the Nacha file.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeMiscellaneous WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "miscellaneous"
	// Sent as `MOR` in the Nacha file.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeMortgage WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "mortgage"
	// Sent as `PEN` in the Nacha file.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodePension WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "pension"
	// Sent as `REM` in the Nacha file.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeRemittance WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "remittance"
	// Sent as `RLS` in the Nacha file.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeRentOrLease WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "rent_or_lease"
	// Sent as `SAL` in the Nacha file.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeSalaryOrPayroll WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "salary_or_payroll"
	// Sent as `TAX` in the Nacha file.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeTax WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "tax"
	// Sent as `ARC` in the Nacha file.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeAccountsReceivable WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "accounts_receivable"
	// Sent as `BOC` in the Nacha file.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeBackOfficeConversion WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "back_office_conversion"
	// Sent as `MTE` in the Nacha file.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeMachineTransfer WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "machine_transfer"
	// Sent as `POP` in the Nacha file.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodePointOfPurchase WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "point_of_purchase"
	// Sent as `POS` in the Nacha file.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodePointOfSale WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "point_of_sale"
	// Sent as `RCK` in the Nacha file.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeRepresentedCheck WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "represented_check"
	// Sent as `SHR` in the Nacha file.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeSharedNetworkTransaction WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "shared_network_transaction"
	// Sent as `TEL` in the Nacha file.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeTelphoneInitiated WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "telphone_initiated"
	// Sent as `WEB` in the Nacha file.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeInternetInitiated WireTransferSimulationTransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "internet_initiated"
)

// An instruction of how to interpret the
// `originating_depository_financial_institution_id` field for this Transaction.
type WireTransferSimulationTransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifier string

const (
	// A domestic clearing system number. In the US, for example, this is the American
	// Banking Association (ABA) routing number.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber WireTransferSimulationTransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifier = "national_clearing_system_number"
	// The SWIFT Bank Identifier Code (BIC) of the bank.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifierBicCode WireTransferSimulationTransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifier = "bic_code"
	// An International Bank Account Number.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifierIban WireTransferSimulationTransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifier = "iban"
)

// An instruction of how to interpret the
// `receiving_depository_financial_institution_id` field for this Transaction.
type WireTransferSimulationTransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifier string

const (
	// A domestic clearing system number. In the US, for example, this is the American
	// Banking Association (ABA) routing number.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber WireTransferSimulationTransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifier = "national_clearing_system_number"
	// The SWIFT Bank Identifier Code (BIC) of the bank.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifierBicCode WireTransferSimulationTransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifier = "bic_code"
	// An International Bank Account Number.
	WireTransferSimulationTransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifierIban WireTransferSimulationTransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifier = "iban"
)

// An Inbound Real-Time Payments Transfer Confirmation object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_confirmation`.
type WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real-Time Payments transfer.
	Currency WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency `json:"currency,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber string `json:"debtor_account_number,required"`
	// The name provided by the sender of the transfer.
	DebtorName string `json:"debtor_name,required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber string `json:"debtor_routing_number,required"`
	// Additional information included with the transfer.
	RemittanceInformation string `json:"remittance_information,required,nullable"`
	// The Real-Time Payments network identification of the transfer.
	TransactionIdentification string                                                                                 `json:"transaction_identification,required"`
	JSON                      wireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON `json:"-"`
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
// currency. This will always be "USD" for a Real-Time Payments transfer.
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

// An Inbound Wire Drawdown Payment object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
type WireTransferSimulationTransactionSourceInboundWireDrawdownPayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// A free-form address field set by the sender.
	BeneficiaryAddressLine1 string `json:"beneficiary_address_line1,required,nullable"`
	// A free-form address field set by the sender.
	BeneficiaryAddressLine2 string `json:"beneficiary_address_line2,required,nullable"`
	// A free-form address field set by the sender.
	BeneficiaryAddressLine3 string `json:"beneficiary_address_line3,required,nullable"`
	// A name set by the sender.
	BeneficiaryName string `json:"beneficiary_name,required,nullable"`
	// A free-form reference string set by the sender, to help identify the transfer.
	BeneficiaryReference string `json:"beneficiary_reference,required,nullable"`
	// An Increase-constructed description of the transfer.
	Description string `json:"description,required"`
	// A unique identifier available to the originating and receiving banks, commonly
	// abbreviated as IMAD. It is created when the wire is submitted to the Fedwire
	// service and is helpful when debugging wires with the receiving bank.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required,nullable"`
	// The address of the wire originator, set by the sending bank.
	OriginatorAddressLine1 string `json:"originator_address_line1,required,nullable"`
	// The address of the wire originator, set by the sending bank.
	OriginatorAddressLine2 string `json:"originator_address_line2,required,nullable"`
	// The address of the wire originator, set by the sending bank.
	OriginatorAddressLine3 string `json:"originator_address_line3,required,nullable"`
	// The originator of the wire, set by the sending bank.
	OriginatorName string `json:"originator_name,required,nullable"`
	// The American Banking Association (ABA) routing number of the bank originating
	// the transfer.
	OriginatorRoutingNumber string `json:"originator_routing_number,required,nullable"`
	// An Increase-created concatenation of the Originator-to-Beneficiary lines.
	OriginatorToBeneficiaryInformation string `json:"originator_to_beneficiary_information,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine1 string `json:"originator_to_beneficiary_information_line1,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine2 string `json:"originator_to_beneficiary_information_line2,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine3 string `json:"originator_to_beneficiary_information_line3,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine4 string                                                                `json:"originator_to_beneficiary_information_line4,required,nullable"`
	JSON                                    wireTransferSimulationTransactionSourceInboundWireDrawdownPaymentJSON `json:"-"`
}

// wireTransferSimulationTransactionSourceInboundWireDrawdownPaymentJSON contains
// the JSON metadata for the struct
// [WireTransferSimulationTransactionSourceInboundWireDrawdownPayment]
type wireTransferSimulationTransactionSourceInboundWireDrawdownPaymentJSON struct {
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
	OriginatorRoutingNumber                 apijson.Field
	OriginatorToBeneficiaryInformation      apijson.Field
	OriginatorToBeneficiaryInformationLine1 apijson.Field
	OriginatorToBeneficiaryInformationLine2 apijson.Field
	OriginatorToBeneficiaryInformationLine3 apijson.Field
	OriginatorToBeneficiaryInformationLine4 apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An Inbound Wire Drawdown Payment Reversal object. This field will be present in
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
	// The American Banking Association (ABA) routing number of the bank originating
	// the transfer.
	OriginatorRoutingNumber string `json:"originator_routing_number,required,nullable"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate time.Time `json:"previous_message_input_cycle_date,required" format:"date"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData string `json:"previous_message_input_message_accountability_data,required"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber string `json:"previous_message_input_sequence_number,required"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource string                                                                        `json:"previous_message_input_source,required"`
	JSON                       wireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversalJSON `json:"-"`
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
	OriginatorRoutingNumber                       apijson.Field
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

// An Inbound Wire Reversal object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_reversal`.
type WireTransferSimulationTransactionSourceInboundWireReversal struct {
	// The amount that was reversed in USD cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the reversal was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description on the reversal message from Fedwire, set by the reversing bank.
	Description string `json:"description,required"`
	// Additional financial institution information included in the wire reversal.
	FinancialInstitutionToFinancialInstitutionInformation string `json:"financial_institution_to_financial_institution_information,required,nullable"`
	// The Fedwire cycle date for the wire reversal. The "Fedwire day" begins at 9:00
	// PM Eastern Time on the evening before the `cycle date`.
	InputCycleDate time.Time `json:"input_cycle_date,required" format:"date"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required"`
	// The Fedwire sequence number.
	InputSequenceNumber string `json:"input_sequence_number,required"`
	// The Fedwire input source identifier.
	InputSource string `json:"input_source,required"`
	// The American Banking Association (ABA) routing number of the bank originating
	// the transfer.
	OriginatorRoutingNumber string `json:"originator_routing_number,required,nullable"`
	// The Fedwire cycle date for the wire transfer that is being reversed by this
	// message.
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
	TransactionID string `json:"transaction_id,required"`
	// The ID for the Wire Transfer that is being reversed.
	WireTransferID string                                                         `json:"wire_transfer_id,required"`
	JSON           wireTransferSimulationTransactionSourceInboundWireReversalJSON `json:"-"`
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
	OriginatorRoutingNumber                               apijson.Field
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

// An Inbound Wire Transfer Intention object. This field will be present in the
// JSON response if and only if `category` is equal to `inbound_wire_transfer`.
type WireTransferSimulationTransactionSourceInboundWireTransfer struct {
	// The amount in USD cents.
	Amount int64 `json:"amount,required"`
	// A free-form address field set by the sender.
	BeneficiaryAddressLine1 string `json:"beneficiary_address_line1,required,nullable"`
	// A free-form address field set by the sender.
	BeneficiaryAddressLine2 string `json:"beneficiary_address_line2,required,nullable"`
	// A free-form address field set by the sender.
	BeneficiaryAddressLine3 string `json:"beneficiary_address_line3,required,nullable"`
	// A name set by the sender.
	BeneficiaryName string `json:"beneficiary_name,required,nullable"`
	// A free-form reference string set by the sender, to help identify the transfer.
	BeneficiaryReference string `json:"beneficiary_reference,required,nullable"`
	// An Increase-constructed description of the transfer.
	Description string `json:"description,required"`
	// A unique identifier available to the originating and receiving banks, commonly
	// abbreviated as IMAD. It is created when the wire is submitted to the Fedwire
	// service and is helpful when debugging wires with the originating bank.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required,nullable"`
	// The address of the wire originator, set by the sending bank.
	OriginatorAddressLine1 string `json:"originator_address_line1,required,nullable"`
	// The address of the wire originator, set by the sending bank.
	OriginatorAddressLine2 string `json:"originator_address_line2,required,nullable"`
	// The address of the wire originator, set by the sending bank.
	OriginatorAddressLine3 string `json:"originator_address_line3,required,nullable"`
	// The originator of the wire, set by the sending bank.
	OriginatorName string `json:"originator_name,required,nullable"`
	// The American Banking Association (ABA) routing number of the bank originating
	// the transfer.
	OriginatorRoutingNumber string `json:"originator_routing_number,required,nullable"`
	// An Increase-created concatenation of the Originator-to-Beneficiary lines.
	OriginatorToBeneficiaryInformation string `json:"originator_to_beneficiary_information,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine1 string `json:"originator_to_beneficiary_information_line1,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine2 string `json:"originator_to_beneficiary_information_line2,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine3 string `json:"originator_to_beneficiary_information_line3,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine4 string `json:"originator_to_beneficiary_information_line4,required,nullable"`
	// The ID of the Inbound Wire Transfer object that resulted in this Transaction.
	TransferID string                                                         `json:"transfer_id,required"`
	JSON       wireTransferSimulationTransactionSourceInboundWireTransferJSON `json:"-"`
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
	OriginatorRoutingNumber                 apijson.Field
	OriginatorToBeneficiaryInformation      apijson.Field
	OriginatorToBeneficiaryInformationLine1 apijson.Field
	OriginatorToBeneficiaryInformationLine2 apijson.Field
	OriginatorToBeneficiaryInformationLine3 apijson.Field
	OriginatorToBeneficiaryInformationLine4 apijson.Field
	TransferID                              apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *WireTransferSimulationTransactionSourceInboundWireTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An Interest Payment object. This field will be present in the JSON response if
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
	PeriodStart time.Time                                                  `json:"period_start,required" format:"date-time"`
	JSON        wireTransferSimulationTransactionSourceInterestPaymentJSON `json:"-"`
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

// An Internal Source object. This field will be present in the JSON response if
// and only if `category` is equal to `internal_source`.
type WireTransferSimulationTransactionSourceInternalSource struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency WireTransferSimulationTransactionSourceInternalSourceCurrency `json:"currency,required"`
	// An Internal Source is a transaction between you and Increase. This describes the
	// reason for the transaction.
	Reason WireTransferSimulationTransactionSourceInternalSourceReason `json:"reason,required"`
	JSON   wireTransferSimulationTransactionSourceInternalSourceJSON   `json:"-"`
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

// An Internal Source is a transaction between you and Increase. This describes the
// reason for the transaction.
type WireTransferSimulationTransactionSourceInternalSourceReason string

const (
	// Account closure
	WireTransferSimulationTransactionSourceInternalSourceReasonAccountClosure WireTransferSimulationTransactionSourceInternalSourceReason = "account_closure"
	// Bank migration
	WireTransferSimulationTransactionSourceInternalSourceReasonBankMigration WireTransferSimulationTransactionSourceInternalSourceReason = "bank_migration"
	// Cashback
	WireTransferSimulationTransactionSourceInternalSourceReasonCashback WireTransferSimulationTransactionSourceInternalSourceReason = "cashback"
	// Check adjustment
	WireTransferSimulationTransactionSourceInternalSourceReasonCheckAdjustment WireTransferSimulationTransactionSourceInternalSourceReason = "check_adjustment"
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

// A Real-Time Payments Transfer Acknowledgement object. This field will be present
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
	// The identifier of the Real-Time Payments Transfer that led to this Transaction.
	TransferID string                                                                             `json:"transfer_id,required"`
	JSON       wireTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgementJSON `json:"-"`
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
	Originator string                                                 `json:"originator,required"`
	JSON       wireTransferSimulationTransactionSourceSampleFundsJSON `json:"-"`
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
	// The identifier of the Wire Transfer that led to this Transaction.
	TransferID string                                                           `json:"transfer_id,required"`
	JSON       wireTransferSimulationTransactionSourceWireTransferIntentionJSON `json:"-"`
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
	// The identifier of the Wire Transfer that led to this Transaction.
	TransferID string                                                           `json:"transfer_id,required"`
	JSON       wireTransferSimulationTransactionSourceWireTransferRejectionJSON `json:"-"`
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
	// The sending bank will set originator_routing_number in production. You can
	// simulate any value here.
	OriginatorRoutingNumber param.Field[string] `json:"originator_routing_number"`
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
