// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/pagination"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// TransactionService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewTransactionService] method
// instead.
type TransactionService struct {
	Options []option.RequestOption
}

// NewTransactionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionService(opts ...option.RequestOption) (r *TransactionService) {
	r = &TransactionService{}
	r.Options = opts
	return
}

// Retrieve a Transaction
func (r *TransactionService) Get(ctx context.Context, transactionID string, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("transactions/%s", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Transactions
func (r *TransactionService) List(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) (res *pagination.Page[Transaction], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "transactions"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Transactions
func (r *TransactionService) ListAutoPaging(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Transaction] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Transactions are the immutable additions and removals of money from your bank
// account. They're the equivalent of line items on your bank statement.
type Transaction struct {
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
	Currency TransactionCurrency `json:"currency,required"`
	// An informational message describing this transaction. Use the fields in `source`
	// to get more detailed information. This field appears as the line-item on the
	// statement.
	Description string `json:"description,required"`
	// The identifier for the route this Transaction came through. Routes are things
	// like cards and ACH details.
	RouteID string `json:"route_id,required,nullable"`
	// The type of the route this Transaction came through.
	RouteType TransactionRouteType `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Transaction. Note that for backwards compatibility reasons, additional
	// undocumented keys may appear in this object. These should be treated as
	// deprecated and will be removed in the future.
	Source TransactionSource `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type TransactionType `json:"type,required"`
	JSON transactionJSON `json:"-"`
}

// transactionJSON contains the JSON metadata for the struct [Transaction]
type transactionJSON struct {
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

func (r *Transaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// Transaction's currency. This will match the currency on the Transaction's
// Account.
type TransactionCurrency string

const (
	// Canadian Dollar (CAD)
	TransactionCurrencyCad TransactionCurrency = "CAD"
	// Swiss Franc (CHF)
	TransactionCurrencyChf TransactionCurrency = "CHF"
	// Euro (EUR)
	TransactionCurrencyEur TransactionCurrency = "EUR"
	// British Pound (GBP)
	TransactionCurrencyGbp TransactionCurrency = "GBP"
	// Japanese Yen (JPY)
	TransactionCurrencyJpy TransactionCurrency = "JPY"
	// US Dollar (USD)
	TransactionCurrencyUsd TransactionCurrency = "USD"
)

func (r TransactionCurrency) IsKnown() bool {
	switch r {
	case TransactionCurrencyCad, TransactionCurrencyChf, TransactionCurrencyEur, TransactionCurrencyGbp, TransactionCurrencyJpy, TransactionCurrencyUsd:
		return true
	}
	return false
}

// The type of the route this Transaction came through.
type TransactionRouteType string

const (
	// An Account Number.
	TransactionRouteTypeAccountNumber TransactionRouteType = "account_number"
	// A Card.
	TransactionRouteTypeCard TransactionRouteType = "card"
)

func (r TransactionRouteType) IsKnown() bool {
	switch r {
	case TransactionRouteTypeAccountNumber, TransactionRouteTypeCard:
		return true
	}
	return false
}

// This is an object giving more details on the network-level event that caused the
// Transaction. Note that for backwards compatibility reasons, additional
// undocumented keys may appear in this object. These should be treated as
// deprecated and will be removed in the future.
type TransactionSource struct {
	// An Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention TransactionSourceAccountTransferIntention `json:"account_transfer_intention,required,nullable"`
	// An ACH Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention TransactionSourceACHTransferIntention `json:"ach_transfer_intention,required,nullable"`
	// An ACH Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection TransactionSourceACHTransferRejection `json:"ach_transfer_rejection,required,nullable"`
	// An ACH Transfer Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn TransactionSourceACHTransferReturn `json:"ach_transfer_return,required,nullable"`
	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance TransactionSourceCardDisputeAcceptance `json:"card_dispute_acceptance,required,nullable"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund TransactionSourceCardRefund `json:"card_refund,required,nullable"`
	// A Card Revenue Payment object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_revenue_payment`.
	CardRevenuePayment TransactionSourceCardRevenuePayment `json:"card_revenue_payment,required,nullable"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement TransactionSourceCardSettlement `json:"card_settlement,required,nullable"`
	// The type of the resource. We may add additional possible values for this enum
	// over time; your application should be able to handle such additions gracefully.
	Category TransactionSourceCategory `json:"category,required"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance TransactionSourceCheckDepositAcceptance `json:"check_deposit_acceptance,required,nullable"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn TransactionSourceCheckDepositReturn `json:"check_deposit_return,required,nullable"`
	// A Check Transfer Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_transfer_deposit`.
	CheckTransferDeposit TransactionSourceCheckTransferDeposit `json:"check_transfer_deposit,required,nullable"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest TransactionSourceCheckTransferStopPaymentRequest `json:"check_transfer_stop_payment_request,required,nullable"`
	// A Fee Payment object. This field will be present in the JSON response if and
	// only if `category` is equal to `fee_payment`.
	FeePayment TransactionSourceFeePayment `json:"fee_payment,required,nullable"`
	// An Inbound ACH Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer TransactionSourceInboundACHTransfer `json:"inbound_ach_transfer,required,nullable"`
	// An Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer TransactionSourceInboundInternationalACHTransfer `json:"inbound_international_ach_transfer,required,nullable"`
	// An Inbound Real-Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation TransactionSourceInboundRealTimePaymentsTransferConfirmation `json:"inbound_real_time_payments_transfer_confirmation,required,nullable"`
	// An Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment TransactionSourceInboundWireDrawdownPayment `json:"inbound_wire_drawdown_payment,required,nullable"`
	// An Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal TransactionSourceInboundWireDrawdownPaymentReversal `json:"inbound_wire_drawdown_payment_reversal,required,nullable"`
	// An Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal TransactionSourceInboundWireReversal `json:"inbound_wire_reversal,required,nullable"`
	// An Inbound Wire Transfer Intention object. This field will be present in the
	// JSON response if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer TransactionSourceInboundWireTransfer `json:"inbound_wire_transfer,required,nullable"`
	// An Interest Payment object. This field will be present in the JSON response if
	// and only if `category` is equal to `interest_payment`.
	InterestPayment TransactionSourceInterestPayment `json:"interest_payment,required,nullable"`
	// An Internal Source object. This field will be present in the JSON response if
	// and only if `category` is equal to `internal_source`.
	InternalSource TransactionSourceInternalSource `json:"internal_source,required,nullable"`
	// A Real-Time Payments Transfer Acknowledgement object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `real_time_payments_transfer_acknowledgement`.
	RealTimePaymentsTransferAcknowledgement TransactionSourceRealTimePaymentsTransferAcknowledgement `json:"real_time_payments_transfer_acknowledgement,required,nullable"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds TransactionSourceSampleFunds `json:"sample_funds,required,nullable"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention TransactionSourceWireTransferIntention `json:"wire_transfer_intention,required,nullable"`
	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection TransactionSourceWireTransferRejection `json:"wire_transfer_rejection,required,nullable"`
	JSON                  transactionSourceJSON                  `json:"-"`
}

// transactionSourceJSON contains the JSON metadata for the struct
// [TransactionSource]
type transactionSourceJSON struct {
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
	CheckTransferStopPaymentRequest             apijson.Field
	FeePayment                                  apijson.Field
	InboundACHTransfer                          apijson.Field
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

func (r *TransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceJSON) RawJSON() string {
	return r.raw
}

// An Account Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_intention`.
type TransactionSourceAccountTransferIntention struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency TransactionSourceAccountTransferIntentionCurrency `json:"currency,required"`
	// The description you chose to give the transfer.
	Description string `json:"description,required"`
	// The identifier of the Account to where the Account Transfer was sent.
	DestinationAccountID string `json:"destination_account_id,required"`
	// The identifier of the Account from where the Account Transfer was sent.
	SourceAccountID string `json:"source_account_id,required"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID string                                        `json:"transfer_id,required"`
	JSON       transactionSourceAccountTransferIntentionJSON `json:"-"`
}

// transactionSourceAccountTransferIntentionJSON contains the JSON metadata for the
// struct [TransactionSourceAccountTransferIntention]
type transactionSourceAccountTransferIntentionJSON struct {
	Amount               apijson.Field
	Currency             apijson.Field
	Description          apijson.Field
	DestinationAccountID apijson.Field
	SourceAccountID      apijson.Field
	TransferID           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TransactionSourceAccountTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceAccountTransferIntentionJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
type TransactionSourceAccountTransferIntentionCurrency string

const (
	// Canadian Dollar (CAD)
	TransactionSourceAccountTransferIntentionCurrencyCad TransactionSourceAccountTransferIntentionCurrency = "CAD"
	// Swiss Franc (CHF)
	TransactionSourceAccountTransferIntentionCurrencyChf TransactionSourceAccountTransferIntentionCurrency = "CHF"
	// Euro (EUR)
	TransactionSourceAccountTransferIntentionCurrencyEur TransactionSourceAccountTransferIntentionCurrency = "EUR"
	// British Pound (GBP)
	TransactionSourceAccountTransferIntentionCurrencyGbp TransactionSourceAccountTransferIntentionCurrency = "GBP"
	// Japanese Yen (JPY)
	TransactionSourceAccountTransferIntentionCurrencyJpy TransactionSourceAccountTransferIntentionCurrency = "JPY"
	// US Dollar (USD)
	TransactionSourceAccountTransferIntentionCurrencyUsd TransactionSourceAccountTransferIntentionCurrency = "USD"
)

func (r TransactionSourceAccountTransferIntentionCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceAccountTransferIntentionCurrencyCad, TransactionSourceAccountTransferIntentionCurrencyChf, TransactionSourceAccountTransferIntentionCurrencyEur, TransactionSourceAccountTransferIntentionCurrencyGbp, TransactionSourceAccountTransferIntentionCurrencyJpy, TransactionSourceAccountTransferIntentionCurrencyUsd:
		return true
	}
	return false
}

// An ACH Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_transfer_intention`.
type TransactionSourceACHTransferIntention struct {
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
	TransferID string                                    `json:"transfer_id,required"`
	JSON       transactionSourceACHTransferIntentionJSON `json:"-"`
}

// transactionSourceACHTransferIntentionJSON contains the JSON metadata for the
// struct [TransactionSourceACHTransferIntention]
type transactionSourceACHTransferIntentionJSON struct {
	AccountNumber       apijson.Field
	Amount              apijson.Field
	RoutingNumber       apijson.Field
	StatementDescriptor apijson.Field
	TransferID          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *TransactionSourceACHTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceACHTransferIntentionJSON) RawJSON() string {
	return r.raw
}

// An ACH Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_transfer_rejection`.
type TransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string                                    `json:"transfer_id,required"`
	JSON       transactionSourceACHTransferRejectionJSON `json:"-"`
}

// transactionSourceACHTransferRejectionJSON contains the JSON metadata for the
// struct [TransactionSourceACHTransferRejection]
type transactionSourceACHTransferRejectionJSON struct {
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceACHTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceACHTransferRejectionJSON) RawJSON() string {
	return r.raw
}

// An ACH Transfer Return object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_return`.
type TransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The three character ACH return code, in the range R01 to R85.
	RawReturnReasonCode string `json:"raw_return_reason_code,required"`
	// Why the ACH Transfer was returned. This reason code is sent by the receiving
	// bank back to Increase.
	ReturnReasonCode TransactionSourceACHTransferReturnReturnReasonCode `json:"return_reason_code,required"`
	// The identifier of the Transaction associated with this return.
	TransactionID string `json:"transaction_id,required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID string                                 `json:"transfer_id,required"`
	JSON       transactionSourceACHTransferReturnJSON `json:"-"`
}

// transactionSourceACHTransferReturnJSON contains the JSON metadata for the struct
// [TransactionSourceACHTransferReturn]
type transactionSourceACHTransferReturnJSON struct {
	CreatedAt           apijson.Field
	RawReturnReasonCode apijson.Field
	ReturnReasonCode    apijson.Field
	TransactionID       apijson.Field
	TransferID          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *TransactionSourceACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceACHTransferReturnJSON) RawJSON() string {
	return r.raw
}

// Why the ACH Transfer was returned. This reason code is sent by the receiving
// bank back to Increase.
type TransactionSourceACHTransferReturnReturnReasonCode string

const (
	// Code R01. Insufficient funds in the receiving account. Sometimes abbreviated to
	// NSF.
	TransactionSourceACHTransferReturnReturnReasonCodeInsufficientFund TransactionSourceACHTransferReturnReturnReasonCode = "insufficient_fund"
	// Code R03. The account does not exist or the receiving bank was unable to locate
	// it.
	TransactionSourceACHTransferReturnReturnReasonCodeNoAccount TransactionSourceACHTransferReturnReturnReasonCode = "no_account"
	// Code R02. The account is closed at the receiving bank.
	TransactionSourceACHTransferReturnReturnReasonCodeAccountClosed TransactionSourceACHTransferReturnReturnReasonCode = "account_closed"
	// Code R04. The account number is invalid at the receiving bank.
	TransactionSourceACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure TransactionSourceACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	// Code R16. The account at the receiving bank was frozen per the Office of Foreign
	// Assets Control.
	TransactionSourceACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction TransactionSourceACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	// Code R23. The receiving bank account refused a credit transfer.
	TransactionSourceACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver TransactionSourceACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	// Code R05. The receiving bank rejected because of an incorrect Standard Entry
	// Class code.
	TransactionSourceACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode TransactionSourceACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	// Code R29. The corporate customer at the receiving bank reversed the transfer.
	TransactionSourceACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized TransactionSourceACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	// Code R08. The receiving bank stopped payment on this transfer.
	TransactionSourceACHTransferReturnReturnReasonCodePaymentStopped TransactionSourceACHTransferReturnReturnReasonCode = "payment_stopped"
	// Code R20. The receiving bank account does not perform transfers.
	TransactionSourceACHTransferReturnReturnReasonCodeNonTransactionAccount TransactionSourceACHTransferReturnReturnReasonCode = "non_transaction_account"
	// Code R09. The receiving bank account does not have enough available balance for
	// the transfer.
	TransactionSourceACHTransferReturnReturnReasonCodeUncollectedFunds TransactionSourceACHTransferReturnReturnReasonCode = "uncollected_funds"
	// Code R28. The routing number is incorrect.
	TransactionSourceACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError TransactionSourceACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	// Code R10. The customer at the receiving bank reversed the transfer.
	TransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete TransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	// Code R19. The amount field is incorrect or too large.
	TransactionSourceACHTransferReturnReturnReasonCodeAmountFieldError TransactionSourceACHTransferReturnReturnReasonCode = "amount_field_error"
	// Code R07. The customer at the receiving institution informed their bank that
	// they have revoked authorization for a previously authorized transfer.
	TransactionSourceACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer TransactionSourceACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	// Code R13. The routing number is invalid.
	TransactionSourceACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber TransactionSourceACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	// Code R17. The receiving bank is unable to process a field in the transfer.
	TransactionSourceACHTransferReturnReturnReasonCodeFileRecordEditCriteria TransactionSourceACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	// Code R45. The individual name field was invalid.
	TransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualName TransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	// Code R06. The originating financial institution asked for this transfer to be
	// returned. The receiving bank is complying with the request.
	TransactionSourceACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest TransactionSourceACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	// Code R34. The receiving bank's regulatory supervisor has limited their
	// participation in the ACH network.
	TransactionSourceACHTransferReturnReturnReasonCodeLimitedParticipationDfi TransactionSourceACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	// Code R85. The outbound international ACH transfer was incorrect.
	TransactionSourceACHTransferReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment TransactionSourceACHTransferReturnReturnReasonCode = "incorrectly_coded_outbound_international_payment"
	// Code R12. A rare return reason. The account was sold to another bank.
	TransactionSourceACHTransferReturnReturnReasonCodeAccountSoldToAnotherDfi TransactionSourceACHTransferReturnReturnReasonCode = "account_sold_to_another_dfi"
	// Code R25. The addenda record is incorrect or missing.
	TransactionSourceACHTransferReturnReturnReasonCodeAddendaError TransactionSourceACHTransferReturnReturnReasonCode = "addenda_error"
	// Code R15. A rare return reason. The account holder is deceased.
	TransactionSourceACHTransferReturnReturnReasonCodeBeneficiaryOrAccountHolderDeceased TransactionSourceACHTransferReturnReturnReasonCode = "beneficiary_or_account_holder_deceased"
	// Code R11. A rare return reason. The customer authorized some payment to the
	// sender, but this payment was not in error.
	TransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedNotWithinAuthorizationTerms TransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_not_within_authorization_terms"
	// Code R74. A rare return reason. Sent in response to a return that was returned
	// with code `field_error`. The latest return should include the corrected
	// field(s).
	TransactionSourceACHTransferReturnReturnReasonCodeCorrectedReturn TransactionSourceACHTransferReturnReturnReasonCode = "corrected_return"
	// Code R24. A rare return reason. The receiving bank received an exact duplicate
	// entry with the same trace number and amount.
	TransactionSourceACHTransferReturnReturnReasonCodeDuplicateEntry TransactionSourceACHTransferReturnReturnReasonCode = "duplicate_entry"
	// Code R67. A rare return reason. The return this message refers to was a
	// duplicate.
	TransactionSourceACHTransferReturnReturnReasonCodeDuplicateReturn TransactionSourceACHTransferReturnReturnReasonCode = "duplicate_return"
	// Code R47. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	TransactionSourceACHTransferReturnReturnReasonCodeEnrDuplicateEnrollment TransactionSourceACHTransferReturnReturnReasonCode = "enr_duplicate_enrollment"
	// Code R43. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	TransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidDfiAccountNumber TransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_dfi_account_number"
	// Code R44. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	TransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualIDNumber TransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_id_number"
	// Code R46. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	TransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidRepresentativePayeeIndicator TransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_representative_payee_indicator"
	// Code R41. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	TransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidTransactionCode TransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_transaction_code"
	// Code R40. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	TransactionSourceACHTransferReturnReturnReasonCodeEnrReturnOfEnrEntry TransactionSourceACHTransferReturnReturnReasonCode = "enr_return_of_enr_entry"
	// Code R42. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	TransactionSourceACHTransferReturnReturnReasonCodeEnrRoutingNumberCheckDigitError TransactionSourceACHTransferReturnReturnReasonCode = "enr_routing_number_check_digit_error"
	// Code R84. A rare return reason. The International ACH Transfer cannot be
	// processed by the gateway.
	TransactionSourceACHTransferReturnReturnReasonCodeEntryNotProcessedByGateway TransactionSourceACHTransferReturnReturnReasonCode = "entry_not_processed_by_gateway"
	// Code R69. A rare return reason. One or more of the fields in the ACH were
	// malformed.
	TransactionSourceACHTransferReturnReturnReasonCodeFieldError TransactionSourceACHTransferReturnReturnReasonCode = "field_error"
	// Code R83. A rare return reason. The Foreign receiving bank was unable to settle
	// this ACH transfer.
	TransactionSourceACHTransferReturnReturnReasonCodeForeignReceivingDfiUnableToSettle TransactionSourceACHTransferReturnReturnReasonCode = "foreign_receiving_dfi_unable_to_settle"
	// Code R80. A rare return reason. The International ACH Transfer is malformed.
	TransactionSourceACHTransferReturnReturnReasonCodeIatEntryCodingError TransactionSourceACHTransferReturnReturnReasonCode = "iat_entry_coding_error"
	// Code R18. A rare return reason. The ACH has an improper effective entry date
	// field.
	TransactionSourceACHTransferReturnReturnReasonCodeImproperEffectiveEntryDate TransactionSourceACHTransferReturnReturnReasonCode = "improper_effective_entry_date"
	// Code R39. A rare return reason. The source document related to this ACH, usually
	// an ACH check conversion, was presented to the bank.
	TransactionSourceACHTransferReturnReturnReasonCodeImproperSourceDocumentSourceDocumentPresented TransactionSourceACHTransferReturnReturnReasonCode = "improper_source_document_source_document_presented"
	// Code R21. A rare return reason. The Company ID field of the ACH was invalid.
	TransactionSourceACHTransferReturnReturnReasonCodeInvalidCompanyID TransactionSourceACHTransferReturnReturnReasonCode = "invalid_company_id"
	// Code R82. A rare return reason. The foreign receiving bank identifier for an
	// International ACH Transfer was invalid.
	TransactionSourceACHTransferReturnReturnReasonCodeInvalidForeignReceivingDfiIdentification TransactionSourceACHTransferReturnReturnReasonCode = "invalid_foreign_receiving_dfi_identification"
	// Code R22. A rare return reason. The Individual ID number field of the ACH was
	// invalid.
	TransactionSourceACHTransferReturnReturnReasonCodeInvalidIndividualIDNumber TransactionSourceACHTransferReturnReturnReasonCode = "invalid_individual_id_number"
	// Code R53. A rare return reason. Both the Represented Check ("RCK") entry and the
	// original check were presented to the bank.
	TransactionSourceACHTransferReturnReturnReasonCodeItemAndRckEntryPresentedForPayment TransactionSourceACHTransferReturnReturnReasonCode = "item_and_rck_entry_presented_for_payment"
	// Code R51. A rare return reason. The Represented Check ("RCK") entry is
	// ineligible.
	TransactionSourceACHTransferReturnReturnReasonCodeItemRelatedToRckEntryIsIneligible TransactionSourceACHTransferReturnReturnReasonCode = "item_related_to_rck_entry_is_ineligible"
	// Code R26. A rare return reason. The ACH is missing a required field.
	TransactionSourceACHTransferReturnReturnReasonCodeMandatoryFieldError TransactionSourceACHTransferReturnReturnReasonCode = "mandatory_field_error"
	// Code R71. A rare return reason. The receiving bank does not recognize the
	// routing number in a dishonored return entry.
	TransactionSourceACHTransferReturnReturnReasonCodeMisroutedDishonoredReturn TransactionSourceACHTransferReturnReturnReasonCode = "misrouted_dishonored_return"
	// Code R61. A rare return reason. The receiving bank does not recognize the
	// routing number in a return entry.
	TransactionSourceACHTransferReturnReturnReasonCodeMisroutedReturn TransactionSourceACHTransferReturnReturnReasonCode = "misrouted_return"
	// Code R76. A rare return reason. Sent in response to a return, the bank does not
	// find the errors alleged by the returning bank.
	TransactionSourceACHTransferReturnReturnReasonCodeNoErrorsFound TransactionSourceACHTransferReturnReturnReasonCode = "no_errors_found"
	// Code R77. A rare return reason. The receiving bank does not accept the return of
	// the erroneous debit. The funds are not available at the receiving bank.
	TransactionSourceACHTransferReturnReturnReasonCodeNonAcceptanceOfR62DishonoredReturn TransactionSourceACHTransferReturnReturnReasonCode = "non_acceptance_of_r62_dishonored_return"
	// Code R81. A rare return reason. The receiving bank does not accept International
	// ACH Transfers.
	TransactionSourceACHTransferReturnReturnReasonCodeNonParticipantInIatProgram TransactionSourceACHTransferReturnReturnReasonCode = "non_participant_in_iat_program"
	// Code R31. A rare return reason. A return that has been agreed to be accepted by
	// the receiving bank, despite falling outside of the usual return timeframe.
	TransactionSourceACHTransferReturnReturnReasonCodePermissibleReturnEntry TransactionSourceACHTransferReturnReturnReasonCode = "permissible_return_entry"
	// Code R70. A rare return reason. The receiving bank had not approved this return.
	TransactionSourceACHTransferReturnReturnReasonCodePermissibleReturnEntryNotAccepted TransactionSourceACHTransferReturnReturnReasonCode = "permissible_return_entry_not_accepted"
	// Code R32. A rare return reason. The receiving bank could not settle this
	// transaction.
	TransactionSourceACHTransferReturnReturnReasonCodeRdfiNonSettlement TransactionSourceACHTransferReturnReturnReasonCode = "rdfi_non_settlement"
	// Code R30. A rare return reason. The receiving bank does not accept Check
	// Truncation ACH transfers.
	TransactionSourceACHTransferReturnReturnReasonCodeRdfiParticipantInCheckTruncationProgram TransactionSourceACHTransferReturnReturnReasonCode = "rdfi_participant_in_check_truncation_program"
	// Code R14. A rare return reason. The payee is deceased.
	TransactionSourceACHTransferReturnReturnReasonCodeRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity TransactionSourceACHTransferReturnReturnReasonCode = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	// Code R75. A rare return reason. The originating bank disputes that an earlier
	// `duplicate_entry` return was actually a duplicate.
	TransactionSourceACHTransferReturnReturnReasonCodeReturnNotADuplicate TransactionSourceACHTransferReturnReturnReasonCode = "return_not_a_duplicate"
	// Code R62. A rare return reason. The originating financial institution made a
	// mistake and this return corrects it.
	TransactionSourceACHTransferReturnReturnReasonCodeReturnOfErroneousOrReversingDebit TransactionSourceACHTransferReturnReturnReasonCode = "return_of_erroneous_or_reversing_debit"
	// Code R36. A rare return reason. Return of a malformed credit entry.
	TransactionSourceACHTransferReturnReturnReasonCodeReturnOfImproperCreditEntry TransactionSourceACHTransferReturnReturnReasonCode = "return_of_improper_credit_entry"
	// Code R35. A rare return reason. Return of a malformed debit entry.
	TransactionSourceACHTransferReturnReturnReasonCodeReturnOfImproperDebitEntry TransactionSourceACHTransferReturnReturnReasonCode = "return_of_improper_debit_entry"
	// Code R33. A rare return reason. Return of a Destroyed Check ("XKC") entry.
	TransactionSourceACHTransferReturnReturnReasonCodeReturnOfXckEntry TransactionSourceACHTransferReturnReturnReasonCode = "return_of_xck_entry"
	// Code R37. A rare return reason. The source document related to this ACH, usually
	// an ACH check conversion, was presented to the bank.
	TransactionSourceACHTransferReturnReturnReasonCodeSourceDocumentPresentedForPayment TransactionSourceACHTransferReturnReturnReasonCode = "source_document_presented_for_payment"
	// Code R50. A rare return reason. State law prevents the bank from accepting the
	// Represented Check ("RCK") entry.
	TransactionSourceACHTransferReturnReturnReasonCodeStateLawAffectingRckAcceptance TransactionSourceACHTransferReturnReturnReasonCode = "state_law_affecting_rck_acceptance"
	// Code R52. A rare return reason. A stop payment was issued on a Represented Check
	// ("RCK") entry.
	TransactionSourceACHTransferReturnReturnReasonCodeStopPaymentOnItemRelatedToRckEntry TransactionSourceACHTransferReturnReturnReasonCode = "stop_payment_on_item_related_to_rck_entry"
	// Code R38. A rare return reason. The source attached to the ACH, usually an ACH
	// check conversion, includes a stop payment.
	TransactionSourceACHTransferReturnReturnReasonCodeStopPaymentOnSourceDocument TransactionSourceACHTransferReturnReturnReasonCode = "stop_payment_on_source_document"
	// Code R73. A rare return reason. The bank receiving an `untimely_return` believes
	// it was on time.
	TransactionSourceACHTransferReturnReturnReasonCodeTimelyOriginalReturn TransactionSourceACHTransferReturnReturnReasonCode = "timely_original_return"
	// Code R27. A rare return reason. An ACH return's trace number does not match an
	// originated ACH.
	TransactionSourceACHTransferReturnReturnReasonCodeTraceNumberError TransactionSourceACHTransferReturnReturnReasonCode = "trace_number_error"
	// Code R72. A rare return reason. The dishonored return was sent too late.
	TransactionSourceACHTransferReturnReturnReasonCodeUntimelyDishonoredReturn TransactionSourceACHTransferReturnReturnReasonCode = "untimely_dishonored_return"
	// Code R68. A rare return reason. The return was sent too late.
	TransactionSourceACHTransferReturnReturnReasonCodeUntimelyReturn TransactionSourceACHTransferReturnReturnReasonCode = "untimely_return"
)

func (r TransactionSourceACHTransferReturnReturnReasonCode) IsKnown() bool {
	switch r {
	case TransactionSourceACHTransferReturnReturnReasonCodeInsufficientFund, TransactionSourceACHTransferReturnReturnReasonCodeNoAccount, TransactionSourceACHTransferReturnReturnReasonCodeAccountClosed, TransactionSourceACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure, TransactionSourceACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction, TransactionSourceACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver, TransactionSourceACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode, TransactionSourceACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized, TransactionSourceACHTransferReturnReturnReasonCodePaymentStopped, TransactionSourceACHTransferReturnReturnReasonCodeNonTransactionAccount, TransactionSourceACHTransferReturnReturnReasonCodeUncollectedFunds, TransactionSourceACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError, TransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete, TransactionSourceACHTransferReturnReturnReasonCodeAmountFieldError, TransactionSourceACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer, TransactionSourceACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber, TransactionSourceACHTransferReturnReturnReasonCodeFileRecordEditCriteria, TransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualName, TransactionSourceACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest, TransactionSourceACHTransferReturnReturnReasonCodeLimitedParticipationDfi, TransactionSourceACHTransferReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment, TransactionSourceACHTransferReturnReturnReasonCodeAccountSoldToAnotherDfi, TransactionSourceACHTransferReturnReturnReasonCodeAddendaError, TransactionSourceACHTransferReturnReturnReasonCodeBeneficiaryOrAccountHolderDeceased, TransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedNotWithinAuthorizationTerms, TransactionSourceACHTransferReturnReturnReasonCodeCorrectedReturn, TransactionSourceACHTransferReturnReturnReasonCodeDuplicateEntry, TransactionSourceACHTransferReturnReturnReasonCodeDuplicateReturn, TransactionSourceACHTransferReturnReturnReasonCodeEnrDuplicateEnrollment, TransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidDfiAccountNumber, TransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualIDNumber, TransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidRepresentativePayeeIndicator, TransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidTransactionCode, TransactionSourceACHTransferReturnReturnReasonCodeEnrReturnOfEnrEntry, TransactionSourceACHTransferReturnReturnReasonCodeEnrRoutingNumberCheckDigitError, TransactionSourceACHTransferReturnReturnReasonCodeEntryNotProcessedByGateway, TransactionSourceACHTransferReturnReturnReasonCodeFieldError, TransactionSourceACHTransferReturnReturnReasonCodeForeignReceivingDfiUnableToSettle, TransactionSourceACHTransferReturnReturnReasonCodeIatEntryCodingError, TransactionSourceACHTransferReturnReturnReasonCodeImproperEffectiveEntryDate, TransactionSourceACHTransferReturnReturnReasonCodeImproperSourceDocumentSourceDocumentPresented, TransactionSourceACHTransferReturnReturnReasonCodeInvalidCompanyID, TransactionSourceACHTransferReturnReturnReasonCodeInvalidForeignReceivingDfiIdentification, TransactionSourceACHTransferReturnReturnReasonCodeInvalidIndividualIDNumber, TransactionSourceACHTransferReturnReturnReasonCodeItemAndRckEntryPresentedForPayment, TransactionSourceACHTransferReturnReturnReasonCodeItemRelatedToRckEntryIsIneligible, TransactionSourceACHTransferReturnReturnReasonCodeMandatoryFieldError, TransactionSourceACHTransferReturnReturnReasonCodeMisroutedDishonoredReturn, TransactionSourceACHTransferReturnReturnReasonCodeMisroutedReturn, TransactionSourceACHTransferReturnReturnReasonCodeNoErrorsFound, TransactionSourceACHTransferReturnReturnReasonCodeNonAcceptanceOfR62DishonoredReturn, TransactionSourceACHTransferReturnReturnReasonCodeNonParticipantInIatProgram, TransactionSourceACHTransferReturnReturnReasonCodePermissibleReturnEntry, TransactionSourceACHTransferReturnReturnReasonCodePermissibleReturnEntryNotAccepted, TransactionSourceACHTransferReturnReturnReasonCodeRdfiNonSettlement, TransactionSourceACHTransferReturnReturnReasonCodeRdfiParticipantInCheckTruncationProgram, TransactionSourceACHTransferReturnReturnReasonCodeRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity, TransactionSourceACHTransferReturnReturnReasonCodeReturnNotADuplicate, TransactionSourceACHTransferReturnReturnReasonCodeReturnOfErroneousOrReversingDebit, TransactionSourceACHTransferReturnReturnReasonCodeReturnOfImproperCreditEntry, TransactionSourceACHTransferReturnReturnReasonCodeReturnOfImproperDebitEntry, TransactionSourceACHTransferReturnReturnReasonCodeReturnOfXckEntry, TransactionSourceACHTransferReturnReturnReasonCodeSourceDocumentPresentedForPayment, TransactionSourceACHTransferReturnReturnReasonCodeStateLawAffectingRckAcceptance, TransactionSourceACHTransferReturnReturnReasonCodeStopPaymentOnItemRelatedToRckEntry, TransactionSourceACHTransferReturnReturnReasonCodeStopPaymentOnSourceDocument, TransactionSourceACHTransferReturnReturnReasonCodeTimelyOriginalReturn, TransactionSourceACHTransferReturnReturnReasonCodeTraceNumberError, TransactionSourceACHTransferReturnReturnReasonCodeUntimelyDishonoredReturn, TransactionSourceACHTransferReturnReturnReasonCodeUntimelyReturn:
		return true
	}
	return false
}

// A Card Dispute Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `card_dispute_acceptance`.
type TransactionSourceCardDisputeAcceptance struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was accepted.
	AcceptedAt time.Time `json:"accepted_at,required" format:"date-time"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID string `json:"card_dispute_id,required"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID string                                     `json:"transaction_id,required"`
	JSON          transactionSourceCardDisputeAcceptanceJSON `json:"-"`
}

// transactionSourceCardDisputeAcceptanceJSON contains the JSON metadata for the
// struct [TransactionSourceCardDisputeAcceptance]
type transactionSourceCardDisputeAcceptanceJSON struct {
	AcceptedAt    apijson.Field
	CardDisputeID apijson.Field
	TransactionID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TransactionSourceCardDisputeAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardDisputeAcceptanceJSON) RawJSON() string {
	return r.raw
}

// A Card Refund object. This field will be present in the JSON response if and
// only if `category` is equal to `card_refund`.
type TransactionSourceCardRefund struct {
	// The Card Refund identifier.
	ID string `json:"id,required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency TransactionSourceCardRefundCurrency `json:"currency,required"`
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
	NetworkIdentifiers TransactionSourceCardRefundNetworkIdentifiers `json:"network_identifiers,required"`
	// Additional details about the card purchase, such as tax and industry-specific
	// fields.
	PurchaseDetails TransactionSourceCardRefundPurchaseDetails `json:"purchase_details,required,nullable"`
	// The identifier of the Transaction associated with this Transaction.
	TransactionID string `json:"transaction_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type TransactionSourceCardRefundType `json:"type,required"`
	JSON transactionSourceCardRefundJSON `json:"-"`
}

// transactionSourceCardRefundJSON contains the JSON metadata for the struct
// [TransactionSourceCardRefund]
type transactionSourceCardRefundJSON struct {
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

func (r *TransactionSourceCardRefund) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardRefundJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type TransactionSourceCardRefundCurrency string

const (
	// Canadian Dollar (CAD)
	TransactionSourceCardRefundCurrencyCad TransactionSourceCardRefundCurrency = "CAD"
	// Swiss Franc (CHF)
	TransactionSourceCardRefundCurrencyChf TransactionSourceCardRefundCurrency = "CHF"
	// Euro (EUR)
	TransactionSourceCardRefundCurrencyEur TransactionSourceCardRefundCurrency = "EUR"
	// British Pound (GBP)
	TransactionSourceCardRefundCurrencyGbp TransactionSourceCardRefundCurrency = "GBP"
	// Japanese Yen (JPY)
	TransactionSourceCardRefundCurrencyJpy TransactionSourceCardRefundCurrency = "JPY"
	// US Dollar (USD)
	TransactionSourceCardRefundCurrencyUsd TransactionSourceCardRefundCurrency = "USD"
)

func (r TransactionSourceCardRefundCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCardRefundCurrencyCad, TransactionSourceCardRefundCurrencyChf, TransactionSourceCardRefundCurrencyEur, TransactionSourceCardRefundCurrencyGbp, TransactionSourceCardRefundCurrencyJpy, TransactionSourceCardRefundCurrencyUsd:
		return true
	}
	return false
}

// Network-specific identifiers for this refund.
type TransactionSourceCardRefundNetworkIdentifiers struct {
	// A network assigned business ID that identifies the acquirer that processed this
	// transaction.
	AcquirerBusinessID string `json:"acquirer_business_id,required"`
	// A globally unique identifier for this settlement.
	AcquirerReferenceNumber string `json:"acquirer_reference_number,required"`
	// A globally unique transaction identifier provided by the card network, used
	// across multiple life-cycle requests.
	TransactionID string                                            `json:"transaction_id,required,nullable"`
	JSON          transactionSourceCardRefundNetworkIdentifiersJSON `json:"-"`
}

// transactionSourceCardRefundNetworkIdentifiersJSON contains the JSON metadata for
// the struct [TransactionSourceCardRefundNetworkIdentifiers]
type transactionSourceCardRefundNetworkIdentifiersJSON struct {
	AcquirerBusinessID      apijson.Field
	AcquirerReferenceNumber apijson.Field
	TransactionID           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *TransactionSourceCardRefundNetworkIdentifiers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardRefundNetworkIdentifiersJSON) RawJSON() string {
	return r.raw
}

// Additional details about the card purchase, such as tax and industry-specific
// fields.
type TransactionSourceCardRefundPurchaseDetails struct {
	// Fields specific to car rentals.
	CarRental TransactionSourceCardRefundPurchaseDetailsCarRental `json:"car_rental,required,nullable"`
	// An identifier from the merchant for the customer or consumer.
	CustomerReferenceIdentifier string `json:"customer_reference_identifier,required,nullable"`
	// The state or provincial tax amount in minor units.
	LocalTaxAmount int64 `json:"local_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	LocalTaxCurrency string `json:"local_tax_currency,required,nullable"`
	// Fields specific to lodging.
	Lodging TransactionSourceCardRefundPurchaseDetailsLodging `json:"lodging,required,nullable"`
	// The national tax amount in minor units.
	NationalTaxAmount int64 `json:"national_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	NationalTaxCurrency string `json:"national_tax_currency,required,nullable"`
	// An identifier from the merchant for the purchase to the issuer and cardholder.
	PurchaseIdentifier string `json:"purchase_identifier,required,nullable"`
	// The format of the purchase identifier.
	PurchaseIdentifierFormat TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat `json:"purchase_identifier_format,required,nullable"`
	// Fields specific to travel.
	Travel TransactionSourceCardRefundPurchaseDetailsTravel `json:"travel,required,nullable"`
	JSON   transactionSourceCardRefundPurchaseDetailsJSON   `json:"-"`
}

// transactionSourceCardRefundPurchaseDetailsJSON contains the JSON metadata for
// the struct [TransactionSourceCardRefundPurchaseDetails]
type transactionSourceCardRefundPurchaseDetailsJSON struct {
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

func (r *TransactionSourceCardRefundPurchaseDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardRefundPurchaseDetailsJSON) RawJSON() string {
	return r.raw
}

// Fields specific to car rentals.
type TransactionSourceCardRefundPurchaseDetailsCarRental struct {
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
	ExtraCharges TransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges `json:"extra_charges,required,nullable"`
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
	NoShowIndicator TransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicator `json:"no_show_indicator,required,nullable"`
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
	WeeklyRentalRateCurrency string                                                  `json:"weekly_rental_rate_currency,required,nullable"`
	JSON                     transactionSourceCardRefundPurchaseDetailsCarRentalJSON `json:"-"`
}

// transactionSourceCardRefundPurchaseDetailsCarRentalJSON contains the JSON
// metadata for the struct [TransactionSourceCardRefundPurchaseDetailsCarRental]
type transactionSourceCardRefundPurchaseDetailsCarRentalJSON struct {
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

func (r *TransactionSourceCardRefundPurchaseDetailsCarRental) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardRefundPurchaseDetailsCarRentalJSON) RawJSON() string {
	return r.raw
}

// Additional charges (gas, late fee, etc.) being billed.
type TransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges string

const (
	// No extra charge
	TransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesNoExtraCharge TransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "no_extra_charge"
	// Gas
	TransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesGas TransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "gas"
	// Extra mileage
	TransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesExtraMileage TransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "extra_mileage"
	// Late return
	TransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesLateReturn TransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "late_return"
	// One way service fee
	TransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesOneWayServiceFee TransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "one_way_service_fee"
	// Parking violation
	TransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesParkingViolation TransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "parking_violation"
)

func (r TransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges) IsKnown() bool {
	switch r {
	case TransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesNoExtraCharge, TransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesGas, TransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesExtraMileage, TransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesLateReturn, TransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesOneWayServiceFee, TransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesParkingViolation:
		return true
	}
	return false
}

// An indicator that the cardholder is being billed for a reserved vehicle that was
// not actually rented (that is, a "no-show" charge).
type TransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicator string

const (
	// Not applicable
	TransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicatorNotApplicable TransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicator = "not_applicable"
	// No show for specialized vehicle
	TransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicatorNoShowForSpecializedVehicle TransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicator = "no_show_for_specialized_vehicle"
)

func (r TransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicator) IsKnown() bool {
	switch r {
	case TransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicatorNotApplicable, TransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicatorNoShowForSpecializedVehicle:
		return true
	}
	return false
}

// Fields specific to lodging.
type TransactionSourceCardRefundPurchaseDetailsLodging struct {
	// Date the customer checked in.
	CheckInDate time.Time `json:"check_in_date,required,nullable" format:"date"`
	// Daily rate being charged for the room.
	DailyRoomRateAmount int64 `json:"daily_room_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the daily room
	// rate.
	DailyRoomRateCurrency string `json:"daily_room_rate_currency,required,nullable"`
	// Additional charges (phone, late check-out, etc.) being billed.
	ExtraCharges TransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges `json:"extra_charges,required,nullable"`
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
	NoShowIndicator TransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicator `json:"no_show_indicator,required,nullable"`
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
	TotalTaxCurrency string                                                `json:"total_tax_currency,required,nullable"`
	JSON             transactionSourceCardRefundPurchaseDetailsLodgingJSON `json:"-"`
}

// transactionSourceCardRefundPurchaseDetailsLodgingJSON contains the JSON metadata
// for the struct [TransactionSourceCardRefundPurchaseDetailsLodging]
type transactionSourceCardRefundPurchaseDetailsLodgingJSON struct {
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

func (r *TransactionSourceCardRefundPurchaseDetailsLodging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardRefundPurchaseDetailsLodgingJSON) RawJSON() string {
	return r.raw
}

// Additional charges (phone, late check-out, etc.) being billed.
type TransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges string

const (
	// No extra charge
	TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesNoExtraCharge TransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "no_extra_charge"
	// Restaurant
	TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesRestaurant TransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "restaurant"
	// Gift shop
	TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesGiftShop TransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "gift_shop"
	// Mini bar
	TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesMiniBar TransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "mini_bar"
	// Telephone
	TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesTelephone TransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "telephone"
	// Other
	TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesOther TransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "other"
	// Laundry
	TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesLaundry TransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "laundry"
)

func (r TransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges) IsKnown() bool {
	switch r {
	case TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesNoExtraCharge, TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesRestaurant, TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesGiftShop, TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesMiniBar, TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesTelephone, TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesOther, TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesLaundry:
		return true
	}
	return false
}

// Indicator that the cardholder is being billed for a reserved room that was not
// actually used.
type TransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicator string

const (
	// Not applicable
	TransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicatorNotApplicable TransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicator = "not_applicable"
	// No show
	TransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicatorNoShow TransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicator = "no_show"
)

func (r TransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicator) IsKnown() bool {
	switch r {
	case TransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicatorNotApplicable, TransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicatorNoShow:
		return true
	}
	return false
}

// The format of the purchase identifier.
type TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat string

const (
	// Free text
	TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatFreeText TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "free_text"
	// Order number
	TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatOrderNumber TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "order_number"
	// Rental agreement number
	TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatRentalAgreementNumber TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "rental_agreement_number"
	// Hotel folio number
	TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatHotelFolioNumber TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "hotel_folio_number"
	// Invoice number
	TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatInvoiceNumber TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "invoice_number"
)

func (r TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat) IsKnown() bool {
	switch r {
	case TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatFreeText, TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatOrderNumber, TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatRentalAgreementNumber, TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatHotelFolioNumber, TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatInvoiceNumber:
		return true
	}
	return false
}

// Fields specific to travel.
type TransactionSourceCardRefundPurchaseDetailsTravel struct {
	// Ancillary purchases in addition to the airfare.
	Ancillary TransactionSourceCardRefundPurchaseDetailsTravelAncillary `json:"ancillary,required,nullable"`
	// Indicates the computerized reservation system used to book the ticket.
	ComputerizedReservationSystem string `json:"computerized_reservation_system,required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator `json:"credit_reason_indicator,required,nullable"`
	// Date of departure.
	DepartureDate time.Time `json:"departure_date,required,nullable" format:"date"`
	// Code for the originating city or airport.
	OriginationCityAirportCode string `json:"origination_city_airport_code,required,nullable"`
	// Name of the passenger.
	PassengerName string `json:"passenger_name,required,nullable"`
	// Indicates whether this ticket is non-refundable.
	RestrictedTicketIndicator TransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicator `json:"restricted_ticket_indicator,required,nullable"`
	// Indicates why a ticket was changed.
	TicketChangeIndicator TransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator `json:"ticket_change_indicator,required,nullable"`
	// Ticket number.
	TicketNumber string `json:"ticket_number,required,nullable"`
	// Code for the travel agency if the ticket was issued by a travel agency.
	TravelAgencyCode string `json:"travel_agency_code,required,nullable"`
	// Name of the travel agency if the ticket was issued by a travel agency.
	TravelAgencyName string `json:"travel_agency_name,required,nullable"`
	// Fields specific to each leg of the journey.
	TripLegs []TransactionSourceCardRefundPurchaseDetailsTravelTripLeg `json:"trip_legs,required,nullable"`
	JSON     transactionSourceCardRefundPurchaseDetailsTravelJSON      `json:"-"`
}

// transactionSourceCardRefundPurchaseDetailsTravelJSON contains the JSON metadata
// for the struct [TransactionSourceCardRefundPurchaseDetailsTravel]
type transactionSourceCardRefundPurchaseDetailsTravelJSON struct {
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

func (r *TransactionSourceCardRefundPurchaseDetailsTravel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardRefundPurchaseDetailsTravelJSON) RawJSON() string {
	return r.raw
}

// Ancillary purchases in addition to the airfare.
type TransactionSourceCardRefundPurchaseDetailsTravelAncillary struct {
	// If this purchase has a connection or relationship to another purchase, such as a
	// baggage fee for a passenger transport ticket, this field should contain the
	// ticket document number for the other purchase.
	ConnectedTicketDocumentNumber string `json:"connected_ticket_document_number,required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator `json:"credit_reason_indicator,required,nullable"`
	// Name of the passenger or description of the ancillary purchase.
	PassengerNameOrDescription string `json:"passenger_name_or_description,required,nullable"`
	// Additional travel charges, such as baggage fees.
	Services []TransactionSourceCardRefundPurchaseDetailsTravelAncillaryService `json:"services,required"`
	// Ticket document number.
	TicketDocumentNumber string                                                        `json:"ticket_document_number,required,nullable"`
	JSON                 transactionSourceCardRefundPurchaseDetailsTravelAncillaryJSON `json:"-"`
}

// transactionSourceCardRefundPurchaseDetailsTravelAncillaryJSON contains the JSON
// metadata for the struct
// [TransactionSourceCardRefundPurchaseDetailsTravelAncillary]
type transactionSourceCardRefundPurchaseDetailsTravelAncillaryJSON struct {
	ConnectedTicketDocumentNumber apijson.Field
	CreditReasonIndicator         apijson.Field
	PassengerNameOrDescription    apijson.Field
	Services                      apijson.Field
	TicketDocumentNumber          apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *TransactionSourceCardRefundPurchaseDetailsTravelAncillary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardRefundPurchaseDetailsTravelAncillaryJSON) RawJSON() string {
	return r.raw
}

// Indicates the reason for a credit to the cardholder.
type TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator string

const (
	// No credit
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorNoCredit TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "no_credit"
	// Passenger transport ancillary purchase cancellation
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket and passenger transport ancillary purchase cancellation
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	// Other
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorOther TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "other"
)

func (r TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator) IsKnown() bool {
	switch r {
	case TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorNoCredit, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorOther:
		return true
	}
	return false
}

type TransactionSourceCardRefundPurchaseDetailsTravelAncillaryService struct {
	// Category of the ancillary service.
	Category TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory `json:"category,required,nullable"`
	// Sub-category of the ancillary service, free-form.
	SubCategory string                                                               `json:"sub_category,required,nullable"`
	JSON        transactionSourceCardRefundPurchaseDetailsTravelAncillaryServiceJSON `json:"-"`
}

// transactionSourceCardRefundPurchaseDetailsTravelAncillaryServiceJSON contains
// the JSON metadata for the struct
// [TransactionSourceCardRefundPurchaseDetailsTravelAncillaryService]
type transactionSourceCardRefundPurchaseDetailsTravelAncillaryServiceJSON struct {
	Category    apijson.Field
	SubCategory apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardRefundPurchaseDetailsTravelAncillaryService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardRefundPurchaseDetailsTravelAncillaryServiceJSON) RawJSON() string {
	return r.raw
}

// Category of the ancillary service.
type TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory string

const (
	// None
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryNone TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "none"
	// Bundled service
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryBundledService TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "bundled_service"
	// Baggage fee
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryBaggageFee TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "baggage_fee"
	// Change fee
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryChangeFee TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "change_fee"
	// Cargo
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryCargo TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "cargo"
	// Carbon offset
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryCarbonOffset TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "carbon_offset"
	// Frequent flyer
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryFrequentFlyer TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "frequent_flyer"
	// Gift card
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryGiftCard TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "gift_card"
	// Ground transport
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryGroundTransport TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "ground_transport"
	// In-flight entertainment
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryInFlightEntertainment TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "in_flight_entertainment"
	// Lounge
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryLounge TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "lounge"
	// Medical
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryMedical TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "medical"
	// Meal beverage
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryMealBeverage TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "meal_beverage"
	// Other
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryOther TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "other"
	// Passenger assist fee
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryPassengerAssistFee TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "passenger_assist_fee"
	// Pets
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryPets TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "pets"
	// Seat fees
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategorySeatFees TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "seat_fees"
	// Standby
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryStandby TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "standby"
	// Service fee
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryServiceFee TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "service_fee"
	// Store
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryStore TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "store"
	// Travel service
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryTravelService TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "travel_service"
	// Unaccompanied travel
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryUnaccompaniedTravel TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "unaccompanied_travel"
	// Upgrades
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryUpgrades TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "upgrades"
	// Wi-fi
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryWifi TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "wifi"
)

func (r TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory) IsKnown() bool {
	switch r {
	case TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryNone, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryBundledService, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryBaggageFee, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryChangeFee, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryCargo, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryCarbonOffset, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryFrequentFlyer, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryGiftCard, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryGroundTransport, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryInFlightEntertainment, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryLounge, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryMedical, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryMealBeverage, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryOther, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryPassengerAssistFee, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryPets, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategorySeatFees, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryStandby, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryServiceFee, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryStore, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryTravelService, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryUnaccompaniedTravel, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryUpgrades, TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryWifi:
		return true
	}
	return false
}

// Indicates the reason for a credit to the cardholder.
type TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator string

const (
	// No credit
	TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorNoCredit TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "no_credit"
	// Passenger transport ancillary purchase cancellation
	TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket and passenger transport ancillary purchase cancellation
	TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket cancellation
	TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketCancellation TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_cancellation"
	// Other
	TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorOther TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "other"
	// Partial refund of airline ticket
	TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorPartialRefundOfAirlineTicket TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "partial_refund_of_airline_ticket"
)

func (r TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator) IsKnown() bool {
	switch r {
	case TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorNoCredit, TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation, TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation, TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketCancellation, TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorOther, TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorPartialRefundOfAirlineTicket:
		return true
	}
	return false
}

// Indicates whether this ticket is non-refundable.
type TransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicator string

const (
	// No restrictions
	TransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicatorNoRestrictions TransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicator = "no_restrictions"
	// Restricted non-refundable ticket
	TransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicatorRestrictedNonRefundableTicket TransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicator = "restricted_non_refundable_ticket"
)

func (r TransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicator) IsKnown() bool {
	switch r {
	case TransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicatorNoRestrictions, TransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicatorRestrictedNonRefundableTicket:
		return true
	}
	return false
}

// Indicates why a ticket was changed.
type TransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator string

const (
	// None
	TransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicatorNone TransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator = "none"
	// Change to existing ticket
	TransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicatorChangeToExistingTicket TransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator = "change_to_existing_ticket"
	// New ticket
	TransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicatorNewTicket TransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator = "new_ticket"
)

func (r TransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator) IsKnown() bool {
	switch r {
	case TransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicatorNone, TransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicatorChangeToExistingTicket, TransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicatorNewTicket:
		return true
	}
	return false
}

type TransactionSourceCardRefundPurchaseDetailsTravelTripLeg struct {
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
	StopOverCode TransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode `json:"stop_over_code,required,nullable"`
	JSON         transactionSourceCardRefundPurchaseDetailsTravelTripLegJSON          `json:"-"`
}

// transactionSourceCardRefundPurchaseDetailsTravelTripLegJSON contains the JSON
// metadata for the struct
// [TransactionSourceCardRefundPurchaseDetailsTravelTripLeg]
type transactionSourceCardRefundPurchaseDetailsTravelTripLegJSON struct {
	CarrierCode                apijson.Field
	DestinationCityAirportCode apijson.Field
	FareBasisCode              apijson.Field
	FlightNumber               apijson.Field
	ServiceClass               apijson.Field
	StopOverCode               apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *TransactionSourceCardRefundPurchaseDetailsTravelTripLeg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardRefundPurchaseDetailsTravelTripLegJSON) RawJSON() string {
	return r.raw
}

// Indicates whether a stopover is allowed on this ticket.
type TransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode string

const (
	// None
	TransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCodeNone TransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode = "none"
	// Stop over allowed
	TransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCodeStopOverAllowed TransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_allowed"
	// Stop over not allowed
	TransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCodeStopOverNotAllowed TransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_not_allowed"
)

func (r TransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode) IsKnown() bool {
	switch r {
	case TransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCodeNone, TransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCodeStopOverAllowed, TransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCodeStopOverNotAllowed:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `card_refund`.
type TransactionSourceCardRefundType string

const (
	TransactionSourceCardRefundTypeCardRefund TransactionSourceCardRefundType = "card_refund"
)

func (r TransactionSourceCardRefundType) IsKnown() bool {
	switch r {
	case TransactionSourceCardRefundTypeCardRefund:
		return true
	}
	return false
}

// A Card Revenue Payment object. This field will be present in the JSON response
// if and only if `category` is equal to `card_revenue_payment`.
type TransactionSourceCardRevenuePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency TransactionSourceCardRevenuePaymentCurrency `json:"currency,required"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time `json:"period_start,required" format:"date-time"`
	// The account the card belonged to.
	TransactedOnAccountID string                                  `json:"transacted_on_account_id,required,nullable"`
	JSON                  transactionSourceCardRevenuePaymentJSON `json:"-"`
}

// transactionSourceCardRevenuePaymentJSON contains the JSON metadata for the
// struct [TransactionSourceCardRevenuePayment]
type transactionSourceCardRevenuePaymentJSON struct {
	Amount                apijson.Field
	Currency              apijson.Field
	PeriodEnd             apijson.Field
	PeriodStart           apijson.Field
	TransactedOnAccountID apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *TransactionSourceCardRevenuePayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardRevenuePaymentJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type TransactionSourceCardRevenuePaymentCurrency string

const (
	// Canadian Dollar (CAD)
	TransactionSourceCardRevenuePaymentCurrencyCad TransactionSourceCardRevenuePaymentCurrency = "CAD"
	// Swiss Franc (CHF)
	TransactionSourceCardRevenuePaymentCurrencyChf TransactionSourceCardRevenuePaymentCurrency = "CHF"
	// Euro (EUR)
	TransactionSourceCardRevenuePaymentCurrencyEur TransactionSourceCardRevenuePaymentCurrency = "EUR"
	// British Pound (GBP)
	TransactionSourceCardRevenuePaymentCurrencyGbp TransactionSourceCardRevenuePaymentCurrency = "GBP"
	// Japanese Yen (JPY)
	TransactionSourceCardRevenuePaymentCurrencyJpy TransactionSourceCardRevenuePaymentCurrency = "JPY"
	// US Dollar (USD)
	TransactionSourceCardRevenuePaymentCurrencyUsd TransactionSourceCardRevenuePaymentCurrency = "USD"
)

func (r TransactionSourceCardRevenuePaymentCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCardRevenuePaymentCurrencyCad, TransactionSourceCardRevenuePaymentCurrencyChf, TransactionSourceCardRevenuePaymentCurrencyEur, TransactionSourceCardRevenuePaymentCurrencyGbp, TransactionSourceCardRevenuePaymentCurrencyJpy, TransactionSourceCardRevenuePaymentCurrencyUsd:
		return true
	}
	return false
}

// A Card Settlement object. This field will be present in the JSON response if and
// only if `category` is equal to `card_settlement`.
type TransactionSourceCardSettlement struct {
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
	Currency TransactionSourceCardSettlementCurrency `json:"currency,required"`
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
	NetworkIdentifiers TransactionSourceCardSettlementNetworkIdentifiers `json:"network_identifiers,required"`
	// The identifier of the Pending Transaction associated with this Transaction.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// The amount in the minor unit of the transaction's presentment currency.
	PresentmentAmount int64 `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's presentment currency.
	PresentmentCurrency string `json:"presentment_currency,required"`
	// Additional details about the card purchase, such as tax and industry-specific
	// fields.
	PurchaseDetails TransactionSourceCardSettlementPurchaseDetails `json:"purchase_details,required,nullable"`
	// The identifier of the Transaction associated with this Transaction.
	TransactionID string `json:"transaction_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_settlement`.
	Type TransactionSourceCardSettlementType `json:"type,required"`
	JSON transactionSourceCardSettlementJSON `json:"-"`
}

// transactionSourceCardSettlementJSON contains the JSON metadata for the struct
// [TransactionSourceCardSettlement]
type transactionSourceCardSettlementJSON struct {
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

func (r *TransactionSourceCardSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardSettlementJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's settlement currency.
type TransactionSourceCardSettlementCurrency string

const (
	// Canadian Dollar (CAD)
	TransactionSourceCardSettlementCurrencyCad TransactionSourceCardSettlementCurrency = "CAD"
	// Swiss Franc (CHF)
	TransactionSourceCardSettlementCurrencyChf TransactionSourceCardSettlementCurrency = "CHF"
	// Euro (EUR)
	TransactionSourceCardSettlementCurrencyEur TransactionSourceCardSettlementCurrency = "EUR"
	// British Pound (GBP)
	TransactionSourceCardSettlementCurrencyGbp TransactionSourceCardSettlementCurrency = "GBP"
	// Japanese Yen (JPY)
	TransactionSourceCardSettlementCurrencyJpy TransactionSourceCardSettlementCurrency = "JPY"
	// US Dollar (USD)
	TransactionSourceCardSettlementCurrencyUsd TransactionSourceCardSettlementCurrency = "USD"
)

func (r TransactionSourceCardSettlementCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCardSettlementCurrencyCad, TransactionSourceCardSettlementCurrencyChf, TransactionSourceCardSettlementCurrencyEur, TransactionSourceCardSettlementCurrencyGbp, TransactionSourceCardSettlementCurrencyJpy, TransactionSourceCardSettlementCurrencyUsd:
		return true
	}
	return false
}

// Network-specific identifiers for this refund.
type TransactionSourceCardSettlementNetworkIdentifiers struct {
	// A network assigned business ID that identifies the acquirer that processed this
	// transaction.
	AcquirerBusinessID string `json:"acquirer_business_id,required"`
	// A globally unique identifier for this settlement.
	AcquirerReferenceNumber string `json:"acquirer_reference_number,required"`
	// A globally unique transaction identifier provided by the card network, used
	// across multiple life-cycle requests.
	TransactionID string                                                `json:"transaction_id,required,nullable"`
	JSON          transactionSourceCardSettlementNetworkIdentifiersJSON `json:"-"`
}

// transactionSourceCardSettlementNetworkIdentifiersJSON contains the JSON metadata
// for the struct [TransactionSourceCardSettlementNetworkIdentifiers]
type transactionSourceCardSettlementNetworkIdentifiersJSON struct {
	AcquirerBusinessID      apijson.Field
	AcquirerReferenceNumber apijson.Field
	TransactionID           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *TransactionSourceCardSettlementNetworkIdentifiers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardSettlementNetworkIdentifiersJSON) RawJSON() string {
	return r.raw
}

// Additional details about the card purchase, such as tax and industry-specific
// fields.
type TransactionSourceCardSettlementPurchaseDetails struct {
	// Fields specific to car rentals.
	CarRental TransactionSourceCardSettlementPurchaseDetailsCarRental `json:"car_rental,required,nullable"`
	// An identifier from the merchant for the customer or consumer.
	CustomerReferenceIdentifier string `json:"customer_reference_identifier,required,nullable"`
	// The state or provincial tax amount in minor units.
	LocalTaxAmount int64 `json:"local_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	LocalTaxCurrency string `json:"local_tax_currency,required,nullable"`
	// Fields specific to lodging.
	Lodging TransactionSourceCardSettlementPurchaseDetailsLodging `json:"lodging,required,nullable"`
	// The national tax amount in minor units.
	NationalTaxAmount int64 `json:"national_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	NationalTaxCurrency string `json:"national_tax_currency,required,nullable"`
	// An identifier from the merchant for the purchase to the issuer and cardholder.
	PurchaseIdentifier string `json:"purchase_identifier,required,nullable"`
	// The format of the purchase identifier.
	PurchaseIdentifierFormat TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat `json:"purchase_identifier_format,required,nullable"`
	// Fields specific to travel.
	Travel TransactionSourceCardSettlementPurchaseDetailsTravel `json:"travel,required,nullable"`
	JSON   transactionSourceCardSettlementPurchaseDetailsJSON   `json:"-"`
}

// transactionSourceCardSettlementPurchaseDetailsJSON contains the JSON metadata
// for the struct [TransactionSourceCardSettlementPurchaseDetails]
type transactionSourceCardSettlementPurchaseDetailsJSON struct {
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

func (r *TransactionSourceCardSettlementPurchaseDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardSettlementPurchaseDetailsJSON) RawJSON() string {
	return r.raw
}

// Fields specific to car rentals.
type TransactionSourceCardSettlementPurchaseDetailsCarRental struct {
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
	ExtraCharges TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges `json:"extra_charges,required,nullable"`
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
	NoShowIndicator TransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicator `json:"no_show_indicator,required,nullable"`
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
	WeeklyRentalRateCurrency string                                                      `json:"weekly_rental_rate_currency,required,nullable"`
	JSON                     transactionSourceCardSettlementPurchaseDetailsCarRentalJSON `json:"-"`
}

// transactionSourceCardSettlementPurchaseDetailsCarRentalJSON contains the JSON
// metadata for the struct
// [TransactionSourceCardSettlementPurchaseDetailsCarRental]
type transactionSourceCardSettlementPurchaseDetailsCarRentalJSON struct {
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

func (r *TransactionSourceCardSettlementPurchaseDetailsCarRental) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardSettlementPurchaseDetailsCarRentalJSON) RawJSON() string {
	return r.raw
}

// Additional charges (gas, late fee, etc.) being billed.
type TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges string

const (
	// No extra charge
	TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesNoExtraCharge TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "no_extra_charge"
	// Gas
	TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesGas TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "gas"
	// Extra mileage
	TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesExtraMileage TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "extra_mileage"
	// Late return
	TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesLateReturn TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "late_return"
	// One way service fee
	TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesOneWayServiceFee TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "one_way_service_fee"
	// Parking violation
	TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesParkingViolation TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "parking_violation"
)

func (r TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges) IsKnown() bool {
	switch r {
	case TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesNoExtraCharge, TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesGas, TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesExtraMileage, TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesLateReturn, TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesOneWayServiceFee, TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesParkingViolation:
		return true
	}
	return false
}

// An indicator that the cardholder is being billed for a reserved vehicle that was
// not actually rented (that is, a "no-show" charge).
type TransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicator string

const (
	// Not applicable
	TransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicatorNotApplicable TransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicator = "not_applicable"
	// No show for specialized vehicle
	TransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicatorNoShowForSpecializedVehicle TransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicator = "no_show_for_specialized_vehicle"
)

func (r TransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicator) IsKnown() bool {
	switch r {
	case TransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicatorNotApplicable, TransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicatorNoShowForSpecializedVehicle:
		return true
	}
	return false
}

// Fields specific to lodging.
type TransactionSourceCardSettlementPurchaseDetailsLodging struct {
	// Date the customer checked in.
	CheckInDate time.Time `json:"check_in_date,required,nullable" format:"date"`
	// Daily rate being charged for the room.
	DailyRoomRateAmount int64 `json:"daily_room_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the daily room
	// rate.
	DailyRoomRateCurrency string `json:"daily_room_rate_currency,required,nullable"`
	// Additional charges (phone, late check-out, etc.) being billed.
	ExtraCharges TransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges `json:"extra_charges,required,nullable"`
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
	NoShowIndicator TransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicator `json:"no_show_indicator,required,nullable"`
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
	TotalTaxCurrency string                                                    `json:"total_tax_currency,required,nullable"`
	JSON             transactionSourceCardSettlementPurchaseDetailsLodgingJSON `json:"-"`
}

// transactionSourceCardSettlementPurchaseDetailsLodgingJSON contains the JSON
// metadata for the struct [TransactionSourceCardSettlementPurchaseDetailsLodging]
type transactionSourceCardSettlementPurchaseDetailsLodgingJSON struct {
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

func (r *TransactionSourceCardSettlementPurchaseDetailsLodging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardSettlementPurchaseDetailsLodgingJSON) RawJSON() string {
	return r.raw
}

// Additional charges (phone, late check-out, etc.) being billed.
type TransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges string

const (
	// No extra charge
	TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesNoExtraCharge TransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "no_extra_charge"
	// Restaurant
	TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesRestaurant TransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "restaurant"
	// Gift shop
	TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesGiftShop TransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "gift_shop"
	// Mini bar
	TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesMiniBar TransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "mini_bar"
	// Telephone
	TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesTelephone TransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "telephone"
	// Other
	TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesOther TransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "other"
	// Laundry
	TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesLaundry TransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "laundry"
)

func (r TransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges) IsKnown() bool {
	switch r {
	case TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesNoExtraCharge, TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesRestaurant, TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesGiftShop, TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesMiniBar, TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesTelephone, TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesOther, TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesLaundry:
		return true
	}
	return false
}

// Indicator that the cardholder is being billed for a reserved room that was not
// actually used.
type TransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicator string

const (
	// Not applicable
	TransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicatorNotApplicable TransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicator = "not_applicable"
	// No show
	TransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicatorNoShow TransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicator = "no_show"
)

func (r TransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicator) IsKnown() bool {
	switch r {
	case TransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicatorNotApplicable, TransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicatorNoShow:
		return true
	}
	return false
}

// The format of the purchase identifier.
type TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat string

const (
	// Free text
	TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatFreeText TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "free_text"
	// Order number
	TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatOrderNumber TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "order_number"
	// Rental agreement number
	TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatRentalAgreementNumber TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "rental_agreement_number"
	// Hotel folio number
	TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatHotelFolioNumber TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "hotel_folio_number"
	// Invoice number
	TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatInvoiceNumber TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "invoice_number"
)

func (r TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat) IsKnown() bool {
	switch r {
	case TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatFreeText, TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatOrderNumber, TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatRentalAgreementNumber, TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatHotelFolioNumber, TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatInvoiceNumber:
		return true
	}
	return false
}

// Fields specific to travel.
type TransactionSourceCardSettlementPurchaseDetailsTravel struct {
	// Ancillary purchases in addition to the airfare.
	Ancillary TransactionSourceCardSettlementPurchaseDetailsTravelAncillary `json:"ancillary,required,nullable"`
	// Indicates the computerized reservation system used to book the ticket.
	ComputerizedReservationSystem string `json:"computerized_reservation_system,required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator `json:"credit_reason_indicator,required,nullable"`
	// Date of departure.
	DepartureDate time.Time `json:"departure_date,required,nullable" format:"date"`
	// Code for the originating city or airport.
	OriginationCityAirportCode string `json:"origination_city_airport_code,required,nullable"`
	// Name of the passenger.
	PassengerName string `json:"passenger_name,required,nullable"`
	// Indicates whether this ticket is non-refundable.
	RestrictedTicketIndicator TransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator `json:"restricted_ticket_indicator,required,nullable"`
	// Indicates why a ticket was changed.
	TicketChangeIndicator TransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator `json:"ticket_change_indicator,required,nullable"`
	// Ticket number.
	TicketNumber string `json:"ticket_number,required,nullable"`
	// Code for the travel agency if the ticket was issued by a travel agency.
	TravelAgencyCode string `json:"travel_agency_code,required,nullable"`
	// Name of the travel agency if the ticket was issued by a travel agency.
	TravelAgencyName string `json:"travel_agency_name,required,nullable"`
	// Fields specific to each leg of the journey.
	TripLegs []TransactionSourceCardSettlementPurchaseDetailsTravelTripLeg `json:"trip_legs,required,nullable"`
	JSON     transactionSourceCardSettlementPurchaseDetailsTravelJSON      `json:"-"`
}

// transactionSourceCardSettlementPurchaseDetailsTravelJSON contains the JSON
// metadata for the struct [TransactionSourceCardSettlementPurchaseDetailsTravel]
type transactionSourceCardSettlementPurchaseDetailsTravelJSON struct {
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

func (r *TransactionSourceCardSettlementPurchaseDetailsTravel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardSettlementPurchaseDetailsTravelJSON) RawJSON() string {
	return r.raw
}

// Ancillary purchases in addition to the airfare.
type TransactionSourceCardSettlementPurchaseDetailsTravelAncillary struct {
	// If this purchase has a connection or relationship to another purchase, such as a
	// baggage fee for a passenger transport ticket, this field should contain the
	// ticket document number for the other purchase.
	ConnectedTicketDocumentNumber string `json:"connected_ticket_document_number,required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator `json:"credit_reason_indicator,required,nullable"`
	// Name of the passenger or description of the ancillary purchase.
	PassengerNameOrDescription string `json:"passenger_name_or_description,required,nullable"`
	// Additional travel charges, such as baggage fees.
	Services []TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryService `json:"services,required"`
	// Ticket document number.
	TicketDocumentNumber string                                                            `json:"ticket_document_number,required,nullable"`
	JSON                 transactionSourceCardSettlementPurchaseDetailsTravelAncillaryJSON `json:"-"`
}

// transactionSourceCardSettlementPurchaseDetailsTravelAncillaryJSON contains the
// JSON metadata for the struct
// [TransactionSourceCardSettlementPurchaseDetailsTravelAncillary]
type transactionSourceCardSettlementPurchaseDetailsTravelAncillaryJSON struct {
	ConnectedTicketDocumentNumber apijson.Field
	CreditReasonIndicator         apijson.Field
	PassengerNameOrDescription    apijson.Field
	Services                      apijson.Field
	TicketDocumentNumber          apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *TransactionSourceCardSettlementPurchaseDetailsTravelAncillary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardSettlementPurchaseDetailsTravelAncillaryJSON) RawJSON() string {
	return r.raw
}

// Indicates the reason for a credit to the cardholder.
type TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator string

const (
	// No credit
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorNoCredit TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "no_credit"
	// Passenger transport ancillary purchase cancellation
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket and passenger transport ancillary purchase cancellation
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	// Other
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorOther TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "other"
)

func (r TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator) IsKnown() bool {
	switch r {
	case TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorNoCredit, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorOther:
		return true
	}
	return false
}

type TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryService struct {
	// Category of the ancillary service.
	Category TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory `json:"category,required,nullable"`
	// Sub-category of the ancillary service, free-form.
	SubCategory string                                                                   `json:"sub_category,required,nullable"`
	JSON        transactionSourceCardSettlementPurchaseDetailsTravelAncillaryServiceJSON `json:"-"`
}

// transactionSourceCardSettlementPurchaseDetailsTravelAncillaryServiceJSON
// contains the JSON metadata for the struct
// [TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryService]
type transactionSourceCardSettlementPurchaseDetailsTravelAncillaryServiceJSON struct {
	Category    apijson.Field
	SubCategory apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardSettlementPurchaseDetailsTravelAncillaryServiceJSON) RawJSON() string {
	return r.raw
}

// Category of the ancillary service.
type TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory string

const (
	// None
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryNone TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "none"
	// Bundled service
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryBundledService TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "bundled_service"
	// Baggage fee
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryBaggageFee TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "baggage_fee"
	// Change fee
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryChangeFee TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "change_fee"
	// Cargo
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryCargo TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "cargo"
	// Carbon offset
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryCarbonOffset TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "carbon_offset"
	// Frequent flyer
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryFrequentFlyer TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "frequent_flyer"
	// Gift card
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryGiftCard TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "gift_card"
	// Ground transport
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryGroundTransport TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "ground_transport"
	// In-flight entertainment
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryInFlightEntertainment TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "in_flight_entertainment"
	// Lounge
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryLounge TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "lounge"
	// Medical
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryMedical TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "medical"
	// Meal beverage
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryMealBeverage TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "meal_beverage"
	// Other
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryOther TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "other"
	// Passenger assist fee
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryPassengerAssistFee TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "passenger_assist_fee"
	// Pets
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryPets TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "pets"
	// Seat fees
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategorySeatFees TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "seat_fees"
	// Standby
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryStandby TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "standby"
	// Service fee
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryServiceFee TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "service_fee"
	// Store
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryStore TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "store"
	// Travel service
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryTravelService TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "travel_service"
	// Unaccompanied travel
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryUnaccompaniedTravel TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "unaccompanied_travel"
	// Upgrades
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryUpgrades TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "upgrades"
	// Wi-fi
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryWifi TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "wifi"
)

func (r TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory) IsKnown() bool {
	switch r {
	case TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryNone, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryBundledService, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryBaggageFee, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryChangeFee, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryCargo, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryCarbonOffset, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryFrequentFlyer, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryGiftCard, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryGroundTransport, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryInFlightEntertainment, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryLounge, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryMedical, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryMealBeverage, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryOther, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryPassengerAssistFee, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryPets, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategorySeatFees, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryStandby, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryServiceFee, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryStore, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryTravelService, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryUnaccompaniedTravel, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryUpgrades, TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryWifi:
		return true
	}
	return false
}

// Indicates the reason for a credit to the cardholder.
type TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator string

const (
	// No credit
	TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorNoCredit TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "no_credit"
	// Passenger transport ancillary purchase cancellation
	TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket and passenger transport ancillary purchase cancellation
	TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket cancellation
	TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketCancellation TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_cancellation"
	// Other
	TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorOther TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "other"
	// Partial refund of airline ticket
	TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorPartialRefundOfAirlineTicket TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "partial_refund_of_airline_ticket"
)

func (r TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator) IsKnown() bool {
	switch r {
	case TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorNoCredit, TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation, TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation, TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketCancellation, TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorOther, TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorPartialRefundOfAirlineTicket:
		return true
	}
	return false
}

// Indicates whether this ticket is non-refundable.
type TransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator string

const (
	// No restrictions
	TransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicatorNoRestrictions TransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator = "no_restrictions"
	// Restricted non-refundable ticket
	TransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicatorRestrictedNonRefundableTicket TransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator = "restricted_non_refundable_ticket"
)

func (r TransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator) IsKnown() bool {
	switch r {
	case TransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicatorNoRestrictions, TransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicatorRestrictedNonRefundableTicket:
		return true
	}
	return false
}

// Indicates why a ticket was changed.
type TransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator string

const (
	// None
	TransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicatorNone TransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "none"
	// Change to existing ticket
	TransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicatorChangeToExistingTicket TransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "change_to_existing_ticket"
	// New ticket
	TransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicatorNewTicket TransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "new_ticket"
)

func (r TransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator) IsKnown() bool {
	switch r {
	case TransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicatorNone, TransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicatorChangeToExistingTicket, TransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicatorNewTicket:
		return true
	}
	return false
}

type TransactionSourceCardSettlementPurchaseDetailsTravelTripLeg struct {
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
	StopOverCode TransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode `json:"stop_over_code,required,nullable"`
	JSON         transactionSourceCardSettlementPurchaseDetailsTravelTripLegJSON          `json:"-"`
}

// transactionSourceCardSettlementPurchaseDetailsTravelTripLegJSON contains the
// JSON metadata for the struct
// [TransactionSourceCardSettlementPurchaseDetailsTravelTripLeg]
type transactionSourceCardSettlementPurchaseDetailsTravelTripLegJSON struct {
	CarrierCode                apijson.Field
	DestinationCityAirportCode apijson.Field
	FareBasisCode              apijson.Field
	FlightNumber               apijson.Field
	ServiceClass               apijson.Field
	StopOverCode               apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *TransactionSourceCardSettlementPurchaseDetailsTravelTripLeg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardSettlementPurchaseDetailsTravelTripLegJSON) RawJSON() string {
	return r.raw
}

// Indicates whether a stopover is allowed on this ticket.
type TransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode string

const (
	// None
	TransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeNone TransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode = "none"
	// Stop over allowed
	TransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeStopOverAllowed TransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_allowed"
	// Stop over not allowed
	TransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeStopOverNotAllowed TransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_not_allowed"
)

func (r TransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode) IsKnown() bool {
	switch r {
	case TransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeNone, TransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeStopOverAllowed, TransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeStopOverNotAllowed:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `card_settlement`.
type TransactionSourceCardSettlementType string

const (
	TransactionSourceCardSettlementTypeCardSettlement TransactionSourceCardSettlementType = "card_settlement"
)

func (r TransactionSourceCardSettlementType) IsKnown() bool {
	switch r {
	case TransactionSourceCardSettlementTypeCardSettlement:
		return true
	}
	return false
}

// The type of the resource. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
type TransactionSourceCategory string

const (
	// Account Transfer Intention: details will be under the
	// `account_transfer_intention` object.
	TransactionSourceCategoryAccountTransferIntention TransactionSourceCategory = "account_transfer_intention"
	// ACH Transfer Intention: details will be under the `ach_transfer_intention`
	// object.
	TransactionSourceCategoryACHTransferIntention TransactionSourceCategory = "ach_transfer_intention"
	// ACH Transfer Rejection: details will be under the `ach_transfer_rejection`
	// object.
	TransactionSourceCategoryACHTransferRejection TransactionSourceCategory = "ach_transfer_rejection"
	// ACH Transfer Return: details will be under the `ach_transfer_return` object.
	TransactionSourceCategoryACHTransferReturn TransactionSourceCategory = "ach_transfer_return"
	// Card Dispute Acceptance: details will be under the `card_dispute_acceptance`
	// object.
	TransactionSourceCategoryCardDisputeAcceptance TransactionSourceCategory = "card_dispute_acceptance"
	// Card Refund: details will be under the `card_refund` object.
	TransactionSourceCategoryCardRefund TransactionSourceCategory = "card_refund"
	// Card Settlement: details will be under the `card_settlement` object.
	TransactionSourceCategoryCardSettlement TransactionSourceCategory = "card_settlement"
	// Card Revenue Payment: details will be under the `card_revenue_payment` object.
	TransactionSourceCategoryCardRevenuePayment TransactionSourceCategory = "card_revenue_payment"
	// Check Deposit Acceptance: details will be under the `check_deposit_acceptance`
	// object.
	TransactionSourceCategoryCheckDepositAcceptance TransactionSourceCategory = "check_deposit_acceptance"
	// Check Deposit Return: details will be under the `check_deposit_return` object.
	TransactionSourceCategoryCheckDepositReturn TransactionSourceCategory = "check_deposit_return"
	// Check Transfer Deposit: details will be under the `check_transfer_deposit`
	// object.
	TransactionSourceCategoryCheckTransferDeposit TransactionSourceCategory = "check_transfer_deposit"
	// Check Transfer Stop Payment Request: details will be under the
	// `check_transfer_stop_payment_request` object.
	TransactionSourceCategoryCheckTransferStopPaymentRequest TransactionSourceCategory = "check_transfer_stop_payment_request"
	// Fee Payment: details will be under the `fee_payment` object.
	TransactionSourceCategoryFeePayment TransactionSourceCategory = "fee_payment"
	// Inbound ACH Transfer Intention: details will be under the `inbound_ach_transfer`
	// object.
	TransactionSourceCategoryInboundACHTransfer TransactionSourceCategory = "inbound_ach_transfer"
	// Inbound ACH Transfer Return Intention: details will be under the
	// `inbound_ach_transfer_return_intention` object.
	TransactionSourceCategoryInboundACHTransferReturnIntention TransactionSourceCategory = "inbound_ach_transfer_return_intention"
	// Inbound Check Deposit Return Intention: details will be under the
	// `inbound_check_deposit_return_intention` object.
	TransactionSourceCategoryInboundCheckDepositReturnIntention TransactionSourceCategory = "inbound_check_deposit_return_intention"
	// Inbound International ACH Transfer: details will be under the
	// `inbound_international_ach_transfer` object.
	TransactionSourceCategoryInboundInternationalACHTransfer TransactionSourceCategory = "inbound_international_ach_transfer"
	// Inbound Real-Time Payments Transfer Confirmation: details will be under the
	// `inbound_real_time_payments_transfer_confirmation` object.
	TransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation TransactionSourceCategory = "inbound_real_time_payments_transfer_confirmation"
	// Inbound Wire Drawdown Payment Reversal: details will be under the
	// `inbound_wire_drawdown_payment_reversal` object.
	TransactionSourceCategoryInboundWireDrawdownPaymentReversal TransactionSourceCategory = "inbound_wire_drawdown_payment_reversal"
	// Inbound Wire Drawdown Payment: details will be under the
	// `inbound_wire_drawdown_payment` object.
	TransactionSourceCategoryInboundWireDrawdownPayment TransactionSourceCategory = "inbound_wire_drawdown_payment"
	// Inbound Wire Reversal: details will be under the `inbound_wire_reversal` object.
	TransactionSourceCategoryInboundWireReversal TransactionSourceCategory = "inbound_wire_reversal"
	// Inbound Wire Transfer Intention: details will be under the
	// `inbound_wire_transfer` object.
	TransactionSourceCategoryInboundWireTransfer TransactionSourceCategory = "inbound_wire_transfer"
	// Inbound Wire Transfer Reversal Intention: details will be under the
	// `inbound_wire_transfer_reversal` object.
	TransactionSourceCategoryInboundWireTransferReversal TransactionSourceCategory = "inbound_wire_transfer_reversal"
	// Interest Payment: details will be under the `interest_payment` object.
	TransactionSourceCategoryInterestPayment TransactionSourceCategory = "interest_payment"
	// Internal Source: details will be under the `internal_source` object.
	TransactionSourceCategoryInternalSource TransactionSourceCategory = "internal_source"
	// Real-Time Payments Transfer Acknowledgement: details will be under the
	// `real_time_payments_transfer_acknowledgement` object.
	TransactionSourceCategoryRealTimePaymentsTransferAcknowledgement TransactionSourceCategory = "real_time_payments_transfer_acknowledgement"
	// Sample Funds: details will be under the `sample_funds` object.
	TransactionSourceCategorySampleFunds TransactionSourceCategory = "sample_funds"
	// Wire Transfer Intention: details will be under the `wire_transfer_intention`
	// object.
	TransactionSourceCategoryWireTransferIntention TransactionSourceCategory = "wire_transfer_intention"
	// Wire Transfer Rejection: details will be under the `wire_transfer_rejection`
	// object.
	TransactionSourceCategoryWireTransferRejection TransactionSourceCategory = "wire_transfer_rejection"
	// The Transaction was made for an undocumented or deprecated reason.
	TransactionSourceCategoryOther TransactionSourceCategory = "other"
)

func (r TransactionSourceCategory) IsKnown() bool {
	switch r {
	case TransactionSourceCategoryAccountTransferIntention, TransactionSourceCategoryACHTransferIntention, TransactionSourceCategoryACHTransferRejection, TransactionSourceCategoryACHTransferReturn, TransactionSourceCategoryCardDisputeAcceptance, TransactionSourceCategoryCardRefund, TransactionSourceCategoryCardSettlement, TransactionSourceCategoryCardRevenuePayment, TransactionSourceCategoryCheckDepositAcceptance, TransactionSourceCategoryCheckDepositReturn, TransactionSourceCategoryCheckTransferDeposit, TransactionSourceCategoryCheckTransferStopPaymentRequest, TransactionSourceCategoryFeePayment, TransactionSourceCategoryInboundACHTransfer, TransactionSourceCategoryInboundACHTransferReturnIntention, TransactionSourceCategoryInboundCheckDepositReturnIntention, TransactionSourceCategoryInboundInternationalACHTransfer, TransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation, TransactionSourceCategoryInboundWireDrawdownPaymentReversal, TransactionSourceCategoryInboundWireDrawdownPayment, TransactionSourceCategoryInboundWireReversal, TransactionSourceCategoryInboundWireTransfer, TransactionSourceCategoryInboundWireTransferReversal, TransactionSourceCategoryInterestPayment, TransactionSourceCategoryInternalSource, TransactionSourceCategoryRealTimePaymentsTransferAcknowledgement, TransactionSourceCategorySampleFunds, TransactionSourceCategoryWireTransferIntention, TransactionSourceCategoryWireTransferRejection, TransactionSourceCategoryOther:
		return true
	}
	return false
}

// A Check Deposit Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_acceptance`.
type TransactionSourceCheckDepositAcceptance struct {
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
	Currency TransactionSourceCheckDepositAcceptanceCurrency `json:"currency,required"`
	// The routing number printed on the check.
	RoutingNumber string `json:"routing_number,required"`
	// The check serial number, if present, for consumer checks. For business checks,
	// the serial number is usually in the `auxiliary_on_us` field.
	SerialNumber string                                      `json:"serial_number,required,nullable"`
	JSON         transactionSourceCheckDepositAcceptanceJSON `json:"-"`
}

// transactionSourceCheckDepositAcceptanceJSON contains the JSON metadata for the
// struct [TransactionSourceCheckDepositAcceptance]
type transactionSourceCheckDepositAcceptanceJSON struct {
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

func (r *TransactionSourceCheckDepositAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCheckDepositAcceptanceJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type TransactionSourceCheckDepositAcceptanceCurrency string

const (
	// Canadian Dollar (CAD)
	TransactionSourceCheckDepositAcceptanceCurrencyCad TransactionSourceCheckDepositAcceptanceCurrency = "CAD"
	// Swiss Franc (CHF)
	TransactionSourceCheckDepositAcceptanceCurrencyChf TransactionSourceCheckDepositAcceptanceCurrency = "CHF"
	// Euro (EUR)
	TransactionSourceCheckDepositAcceptanceCurrencyEur TransactionSourceCheckDepositAcceptanceCurrency = "EUR"
	// British Pound (GBP)
	TransactionSourceCheckDepositAcceptanceCurrencyGbp TransactionSourceCheckDepositAcceptanceCurrency = "GBP"
	// Japanese Yen (JPY)
	TransactionSourceCheckDepositAcceptanceCurrencyJpy TransactionSourceCheckDepositAcceptanceCurrency = "JPY"
	// US Dollar (USD)
	TransactionSourceCheckDepositAcceptanceCurrencyUsd TransactionSourceCheckDepositAcceptanceCurrency = "USD"
)

func (r TransactionSourceCheckDepositAcceptanceCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCheckDepositAcceptanceCurrencyCad, TransactionSourceCheckDepositAcceptanceCurrencyChf, TransactionSourceCheckDepositAcceptanceCurrencyEur, TransactionSourceCheckDepositAcceptanceCurrencyGbp, TransactionSourceCheckDepositAcceptanceCurrencyJpy, TransactionSourceCheckDepositAcceptanceCurrencyUsd:
		return true
	}
	return false
}

// A Check Deposit Return object. This field will be present in the JSON response
// if and only if `category` is equal to `check_deposit_return`.
type TransactionSourceCheckDepositReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID string `json:"check_deposit_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency TransactionSourceCheckDepositReturnCurrency `json:"currency,required"`
	// Why this check was returned by the bank holding the account it was drawn
	// against.
	ReturnReason TransactionSourceCheckDepositReturnReturnReason `json:"return_reason,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID string                                  `json:"transaction_id,required"`
	JSON          transactionSourceCheckDepositReturnJSON `json:"-"`
}

// transactionSourceCheckDepositReturnJSON contains the JSON metadata for the
// struct [TransactionSourceCheckDepositReturn]
type transactionSourceCheckDepositReturnJSON struct {
	Amount         apijson.Field
	CheckDepositID apijson.Field
	Currency       apijson.Field
	ReturnReason   apijson.Field
	ReturnedAt     apijson.Field
	TransactionID  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionSourceCheckDepositReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCheckDepositReturnJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type TransactionSourceCheckDepositReturnCurrency string

const (
	// Canadian Dollar (CAD)
	TransactionSourceCheckDepositReturnCurrencyCad TransactionSourceCheckDepositReturnCurrency = "CAD"
	// Swiss Franc (CHF)
	TransactionSourceCheckDepositReturnCurrencyChf TransactionSourceCheckDepositReturnCurrency = "CHF"
	// Euro (EUR)
	TransactionSourceCheckDepositReturnCurrencyEur TransactionSourceCheckDepositReturnCurrency = "EUR"
	// British Pound (GBP)
	TransactionSourceCheckDepositReturnCurrencyGbp TransactionSourceCheckDepositReturnCurrency = "GBP"
	// Japanese Yen (JPY)
	TransactionSourceCheckDepositReturnCurrencyJpy TransactionSourceCheckDepositReturnCurrency = "JPY"
	// US Dollar (USD)
	TransactionSourceCheckDepositReturnCurrencyUsd TransactionSourceCheckDepositReturnCurrency = "USD"
)

func (r TransactionSourceCheckDepositReturnCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCheckDepositReturnCurrencyCad, TransactionSourceCheckDepositReturnCurrencyChf, TransactionSourceCheckDepositReturnCurrencyEur, TransactionSourceCheckDepositReturnCurrencyGbp, TransactionSourceCheckDepositReturnCurrencyJpy, TransactionSourceCheckDepositReturnCurrencyUsd:
		return true
	}
	return false
}

// Why this check was returned by the bank holding the account it was drawn
// against.
type TransactionSourceCheckDepositReturnReturnReason string

const (
	// The check doesn't allow ACH conversion.
	TransactionSourceCheckDepositReturnReturnReasonACHConversionNotSupported TransactionSourceCheckDepositReturnReturnReason = "ach_conversion_not_supported"
	// The account is closed.
	TransactionSourceCheckDepositReturnReturnReasonClosedAccount TransactionSourceCheckDepositReturnReturnReason = "closed_account"
	// The check has already been deposited.
	TransactionSourceCheckDepositReturnReturnReasonDuplicateSubmission TransactionSourceCheckDepositReturnReturnReason = "duplicate_submission"
	// Insufficient funds
	TransactionSourceCheckDepositReturnReturnReasonInsufficientFunds TransactionSourceCheckDepositReturnReturnReason = "insufficient_funds"
	// No account was found matching the check details.
	TransactionSourceCheckDepositReturnReturnReasonNoAccount TransactionSourceCheckDepositReturnReturnReason = "no_account"
	// The check was not authorized.
	TransactionSourceCheckDepositReturnReturnReasonNotAuthorized TransactionSourceCheckDepositReturnReturnReason = "not_authorized"
	// The check is too old.
	TransactionSourceCheckDepositReturnReturnReasonStaleDated TransactionSourceCheckDepositReturnReturnReason = "stale_dated"
	// The payment has been stopped by the account holder.
	TransactionSourceCheckDepositReturnReturnReasonStopPayment TransactionSourceCheckDepositReturnReturnReason = "stop_payment"
	// The reason for the return is unknown.
	TransactionSourceCheckDepositReturnReturnReasonUnknownReason TransactionSourceCheckDepositReturnReturnReason = "unknown_reason"
	// The image doesn't match the details submitted.
	TransactionSourceCheckDepositReturnReturnReasonUnmatchedDetails TransactionSourceCheckDepositReturnReturnReason = "unmatched_details"
	// The image could not be read.
	TransactionSourceCheckDepositReturnReturnReasonUnreadableImage TransactionSourceCheckDepositReturnReturnReason = "unreadable_image"
	// The check endorsement was irregular.
	TransactionSourceCheckDepositReturnReturnReasonEndorsementIrregular TransactionSourceCheckDepositReturnReturnReason = "endorsement_irregular"
	// The check present was either altered or fake.
	TransactionSourceCheckDepositReturnReturnReasonAlteredOrFictitiousItem TransactionSourceCheckDepositReturnReturnReason = "altered_or_fictitious_item"
	// The account this check is drawn on is frozen.
	TransactionSourceCheckDepositReturnReturnReasonFrozenOrBlockedAccount TransactionSourceCheckDepositReturnReturnReason = "frozen_or_blocked_account"
	// The check is post dated.
	TransactionSourceCheckDepositReturnReturnReasonPostDated TransactionSourceCheckDepositReturnReturnReason = "post_dated"
	// The endorsement was missing.
	TransactionSourceCheckDepositReturnReturnReasonEndorsementMissing TransactionSourceCheckDepositReturnReturnReason = "endorsement_missing"
	// The check signature was missing.
	TransactionSourceCheckDepositReturnReturnReasonSignatureMissing TransactionSourceCheckDepositReturnReturnReason = "signature_missing"
	// The bank suspects a stop payment will be placed.
	TransactionSourceCheckDepositReturnReturnReasonStopPaymentSuspect TransactionSourceCheckDepositReturnReturnReason = "stop_payment_suspect"
	// The bank cannot read the image.
	TransactionSourceCheckDepositReturnReturnReasonUnusableImage TransactionSourceCheckDepositReturnReturnReason = "unusable_image"
	// The check image fails the bank's security check.
	TransactionSourceCheckDepositReturnReturnReasonImageFailsSecurityCheck TransactionSourceCheckDepositReturnReturnReason = "image_fails_security_check"
	// The bank cannot determine the amount.
	TransactionSourceCheckDepositReturnReturnReasonCannotDetermineAmount TransactionSourceCheckDepositReturnReturnReason = "cannot_determine_amount"
	// The signature is inconsistent with prior signatures.
	TransactionSourceCheckDepositReturnReturnReasonSignatureIrregular TransactionSourceCheckDepositReturnReturnReason = "signature_irregular"
	// The check is a non-cash item and cannot be drawn against the account.
	TransactionSourceCheckDepositReturnReturnReasonNonCashItem TransactionSourceCheckDepositReturnReturnReason = "non_cash_item"
	// The bank is unable to process this check.
	TransactionSourceCheckDepositReturnReturnReasonUnableToProcess TransactionSourceCheckDepositReturnReturnReason = "unable_to_process"
	// The check exceeds the bank or customer's limit.
	TransactionSourceCheckDepositReturnReturnReasonItemExceedsDollarLimit TransactionSourceCheckDepositReturnReturnReason = "item_exceeds_dollar_limit"
	// The bank sold this account and no longer services this customer.
	TransactionSourceCheckDepositReturnReturnReasonBranchOrAccountSold TransactionSourceCheckDepositReturnReturnReason = "branch_or_account_sold"
)

func (r TransactionSourceCheckDepositReturnReturnReason) IsKnown() bool {
	switch r {
	case TransactionSourceCheckDepositReturnReturnReasonACHConversionNotSupported, TransactionSourceCheckDepositReturnReturnReasonClosedAccount, TransactionSourceCheckDepositReturnReturnReasonDuplicateSubmission, TransactionSourceCheckDepositReturnReturnReasonInsufficientFunds, TransactionSourceCheckDepositReturnReturnReasonNoAccount, TransactionSourceCheckDepositReturnReturnReasonNotAuthorized, TransactionSourceCheckDepositReturnReturnReasonStaleDated, TransactionSourceCheckDepositReturnReturnReasonStopPayment, TransactionSourceCheckDepositReturnReturnReasonUnknownReason, TransactionSourceCheckDepositReturnReturnReasonUnmatchedDetails, TransactionSourceCheckDepositReturnReturnReasonUnreadableImage, TransactionSourceCheckDepositReturnReturnReasonEndorsementIrregular, TransactionSourceCheckDepositReturnReturnReasonAlteredOrFictitiousItem, TransactionSourceCheckDepositReturnReturnReasonFrozenOrBlockedAccount, TransactionSourceCheckDepositReturnReturnReasonPostDated, TransactionSourceCheckDepositReturnReturnReasonEndorsementMissing, TransactionSourceCheckDepositReturnReturnReasonSignatureMissing, TransactionSourceCheckDepositReturnReturnReasonStopPaymentSuspect, TransactionSourceCheckDepositReturnReturnReasonUnusableImage, TransactionSourceCheckDepositReturnReturnReasonImageFailsSecurityCheck, TransactionSourceCheckDepositReturnReturnReasonCannotDetermineAmount, TransactionSourceCheckDepositReturnReturnReasonSignatureIrregular, TransactionSourceCheckDepositReturnReturnReasonNonCashItem, TransactionSourceCheckDepositReturnReturnReasonUnableToProcess, TransactionSourceCheckDepositReturnReturnReasonItemExceedsDollarLimit, TransactionSourceCheckDepositReturnReturnReasonBranchOrAccountSold:
		return true
	}
	return false
}

// A Check Transfer Deposit object. This field will be present in the JSON response
// if and only if `category` is equal to `check_transfer_deposit`.
type TransactionSourceCheckTransferDeposit struct {
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
	TransferID string `json:"transfer_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_deposit`.
	Type TransactionSourceCheckTransferDepositType `json:"type,required"`
	JSON transactionSourceCheckTransferDepositJSON `json:"-"`
}

// transactionSourceCheckTransferDepositJSON contains the JSON metadata for the
// struct [TransactionSourceCheckTransferDeposit]
type transactionSourceCheckTransferDepositJSON struct {
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

func (r *TransactionSourceCheckTransferDeposit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCheckTransferDepositJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_deposit`.
type TransactionSourceCheckTransferDepositType string

const (
	TransactionSourceCheckTransferDepositTypeCheckTransferDeposit TransactionSourceCheckTransferDepositType = "check_transfer_deposit"
)

func (r TransactionSourceCheckTransferDepositType) IsKnown() bool {
	switch r {
	case TransactionSourceCheckTransferDepositTypeCheckTransferDeposit:
		return true
	}
	return false
}

// A Check Transfer Stop Payment Request object. This field will be present in the
// JSON response if and only if `category` is equal to
// `check_transfer_stop_payment_request`.
type TransactionSourceCheckTransferStopPaymentRequest struct {
	// The reason why this transfer was stopped.
	Reason TransactionSourceCheckTransferStopPaymentRequestReason `json:"reason,required"`
	// The time the stop-payment was requested.
	RequestedAt time.Time `json:"requested_at,required" format:"date-time"`
	// The ID of the check transfer that was stopped.
	TransferID string `json:"transfer_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type TransactionSourceCheckTransferStopPaymentRequestType `json:"type,required"`
	JSON transactionSourceCheckTransferStopPaymentRequestJSON `json:"-"`
}

// transactionSourceCheckTransferStopPaymentRequestJSON contains the JSON metadata
// for the struct [TransactionSourceCheckTransferStopPaymentRequest]
type transactionSourceCheckTransferStopPaymentRequestJSON struct {
	Reason      apijson.Field
	RequestedAt apijson.Field
	TransferID  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCheckTransferStopPaymentRequestJSON) RawJSON() string {
	return r.raw
}

// The reason why this transfer was stopped.
type TransactionSourceCheckTransferStopPaymentRequestReason string

const (
	// The check could not be delivered.
	TransactionSourceCheckTransferStopPaymentRequestReasonMailDeliveryFailed TransactionSourceCheckTransferStopPaymentRequestReason = "mail_delivery_failed"
	// The check was canceled by an Increase operator who will provide details
	// out-of-band.
	TransactionSourceCheckTransferStopPaymentRequestReasonRejectedByIncrease TransactionSourceCheckTransferStopPaymentRequestReason = "rejected_by_increase"
	// The check was not authorized.
	TransactionSourceCheckTransferStopPaymentRequestReasonNotAuthorized TransactionSourceCheckTransferStopPaymentRequestReason = "not_authorized"
	// The check was stopped for another reason.
	TransactionSourceCheckTransferStopPaymentRequestReasonUnknown TransactionSourceCheckTransferStopPaymentRequestReason = "unknown"
)

func (r TransactionSourceCheckTransferStopPaymentRequestReason) IsKnown() bool {
	switch r {
	case TransactionSourceCheckTransferStopPaymentRequestReasonMailDeliveryFailed, TransactionSourceCheckTransferStopPaymentRequestReasonRejectedByIncrease, TransactionSourceCheckTransferStopPaymentRequestReasonNotAuthorized, TransactionSourceCheckTransferStopPaymentRequestReasonUnknown:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_stop_payment_request`.
type TransactionSourceCheckTransferStopPaymentRequestType string

const (
	TransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest TransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

func (r TransactionSourceCheckTransferStopPaymentRequestType) IsKnown() bool {
	switch r {
	case TransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest:
		return true
	}
	return false
}

// A Fee Payment object. This field will be present in the JSON response if and
// only if `category` is equal to `fee_payment`.
type TransactionSourceFeePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency TransactionSourceFeePaymentCurrency `json:"currency,required"`
	// The start of this payment's fee period, usually the first day of a month.
	FeePeriodStart time.Time                       `json:"fee_period_start,required" format:"date"`
	JSON           transactionSourceFeePaymentJSON `json:"-"`
}

// transactionSourceFeePaymentJSON contains the JSON metadata for the struct
// [TransactionSourceFeePayment]
type transactionSourceFeePaymentJSON struct {
	Amount         apijson.Field
	Currency       apijson.Field
	FeePeriodStart apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionSourceFeePayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceFeePaymentJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type TransactionSourceFeePaymentCurrency string

const (
	// Canadian Dollar (CAD)
	TransactionSourceFeePaymentCurrencyCad TransactionSourceFeePaymentCurrency = "CAD"
	// Swiss Franc (CHF)
	TransactionSourceFeePaymentCurrencyChf TransactionSourceFeePaymentCurrency = "CHF"
	// Euro (EUR)
	TransactionSourceFeePaymentCurrencyEur TransactionSourceFeePaymentCurrency = "EUR"
	// British Pound (GBP)
	TransactionSourceFeePaymentCurrencyGbp TransactionSourceFeePaymentCurrency = "GBP"
	// Japanese Yen (JPY)
	TransactionSourceFeePaymentCurrencyJpy TransactionSourceFeePaymentCurrency = "JPY"
	// US Dollar (USD)
	TransactionSourceFeePaymentCurrencyUsd TransactionSourceFeePaymentCurrency = "USD"
)

func (r TransactionSourceFeePaymentCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceFeePaymentCurrencyCad, TransactionSourceFeePaymentCurrencyChf, TransactionSourceFeePaymentCurrencyEur, TransactionSourceFeePaymentCurrencyGbp, TransactionSourceFeePaymentCurrencyJpy, TransactionSourceFeePaymentCurrencyUsd:
		return true
	}
	return false
}

// An Inbound ACH Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_ach_transfer`.
type TransactionSourceInboundACHTransfer struct {
	// Additional information sent from the originator.
	Addenda TransactionSourceInboundACHTransferAddenda `json:"addenda,required,nullable"`
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
	// [used to correlate returns](https://increase.com/documentation/ach-returns#ach-returns).
	TraceNumber string `json:"trace_number,required"`
	// The inbound ach transfer's identifier.
	TransferID string                                  `json:"transfer_id,required"`
	JSON       transactionSourceInboundACHTransferJSON `json:"-"`
}

// transactionSourceInboundACHTransferJSON contains the JSON metadata for the
// struct [TransactionSourceInboundACHTransfer]
type transactionSourceInboundACHTransferJSON struct {
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

func (r *TransactionSourceInboundACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceInboundACHTransferJSON) RawJSON() string {
	return r.raw
}

// Additional information sent from the originator.
type TransactionSourceInboundACHTransferAddenda struct {
	// The type of addendum.
	Category TransactionSourceInboundACHTransferAddendaCategory `json:"category,required"`
	// Unstructured `payment_related_information` passed through by the originator.
	Freeform TransactionSourceInboundACHTransferAddendaFreeform `json:"freeform,required,nullable"`
	JSON     transactionSourceInboundACHTransferAddendaJSON     `json:"-"`
}

// transactionSourceInboundACHTransferAddendaJSON contains the JSON metadata for
// the struct [TransactionSourceInboundACHTransferAddenda]
type transactionSourceInboundACHTransferAddendaJSON struct {
	Category    apijson.Field
	Freeform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceInboundACHTransferAddenda) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceInboundACHTransferAddendaJSON) RawJSON() string {
	return r.raw
}

// The type of addendum.
type TransactionSourceInboundACHTransferAddendaCategory string

const (
	// Unstructured addendum.
	TransactionSourceInboundACHTransferAddendaCategoryFreeform TransactionSourceInboundACHTransferAddendaCategory = "freeform"
)

func (r TransactionSourceInboundACHTransferAddendaCategory) IsKnown() bool {
	switch r {
	case TransactionSourceInboundACHTransferAddendaCategoryFreeform:
		return true
	}
	return false
}

// Unstructured `payment_related_information` passed through by the originator.
type TransactionSourceInboundACHTransferAddendaFreeform struct {
	// Each entry represents an addendum received from the originator.
	Entries []TransactionSourceInboundACHTransferAddendaFreeformEntry `json:"entries,required"`
	JSON    transactionSourceInboundACHTransferAddendaFreeformJSON    `json:"-"`
}

// transactionSourceInboundACHTransferAddendaFreeformJSON contains the JSON
// metadata for the struct [TransactionSourceInboundACHTransferAddendaFreeform]
type transactionSourceInboundACHTransferAddendaFreeformJSON struct {
	Entries     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceInboundACHTransferAddendaFreeform) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceInboundACHTransferAddendaFreeformJSON) RawJSON() string {
	return r.raw
}

type TransactionSourceInboundACHTransferAddendaFreeformEntry struct {
	// The payment related information passed in the addendum.
	PaymentRelatedInformation string                                                      `json:"payment_related_information,required"`
	JSON                      transactionSourceInboundACHTransferAddendaFreeformEntryJSON `json:"-"`
}

// transactionSourceInboundACHTransferAddendaFreeformEntryJSON contains the JSON
// metadata for the struct
// [TransactionSourceInboundACHTransferAddendaFreeformEntry]
type transactionSourceInboundACHTransferAddendaFreeformEntryJSON struct {
	PaymentRelatedInformation apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *TransactionSourceInboundACHTransferAddendaFreeformEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceInboundACHTransferAddendaFreeformEntryJSON) RawJSON() string {
	return r.raw
}

// An Inbound International ACH Transfer object. This field will be present in the
// JSON response if and only if `category` is equal to
// `inbound_international_ach_transfer`.
type TransactionSourceInboundInternationalACHTransfer struct {
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
	ForeignExchangeIndicator TransactionSourceInboundInternationalACHTransferForeignExchangeIndicator `json:"foreign_exchange_indicator,required"`
	// Depending on the `foreign_exchange_reference_indicator`, an exchange rate or a
	// reference to a well-known rate.
	ForeignExchangeReference string `json:"foreign_exchange_reference,required,nullable"`
	// An instruction of how to interpret the `foreign_exchange_reference` field for
	// this Transaction.
	ForeignExchangeReferenceIndicator TransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicator `json:"foreign_exchange_reference_indicator,required"`
	// The amount in the minor unit of the foreign payment currency. For dollars, for
	// example, this is cents.
	ForeignPaymentAmount int64 `json:"foreign_payment_amount,required"`
	// A reference number in the foreign banking infrastructure.
	ForeignTraceNumber string `json:"foreign_trace_number,required,nullable"`
	// The type of transfer. Set by the originator.
	InternationalTransactionTypeCode TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode `json:"international_transaction_type_code,required"`
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
	OriginatingDepositoryFinancialInstitutionIDQualifier TransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifier `json:"originating_depository_financial_institution_id_qualifier,required"`
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
	ReceivingDepositoryFinancialInstitutionIDQualifier TransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifier `json:"receiving_depository_financial_institution_id_qualifier,required"`
	// The name of the receiving bank, as set by the sending financial institution.
	ReceivingDepositoryFinancialInstitutionName string `json:"receiving_depository_financial_institution_name,required"`
	// A 15 digit number recorded in the Nacha file and available to both the
	// originating and receiving bank. Along with the amount, date, and originating
	// routing number, this can be used to identify the ACH transfer at either bank.
	// ACH trace numbers are not unique, but are
	// [used to correlate returns](https://increase.com/documentation/ach-returns#ach-returns).
	TraceNumber string                                               `json:"trace_number,required"`
	JSON        transactionSourceInboundInternationalACHTransferJSON `json:"-"`
}

// transactionSourceInboundInternationalACHTransferJSON contains the JSON metadata
// for the struct [TransactionSourceInboundInternationalACHTransfer]
type transactionSourceInboundInternationalACHTransferJSON struct {
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

func (r *TransactionSourceInboundInternationalACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceInboundInternationalACHTransferJSON) RawJSON() string {
	return r.raw
}

// A description of how the foreign exchange rate was calculated.
type TransactionSourceInboundInternationalACHTransferForeignExchangeIndicator string

const (
	// The originator chose an amount in their own currency. The settled amount in USD
	// was converted using the exchange rate.
	TransactionSourceInboundInternationalACHTransferForeignExchangeIndicatorFixedToVariable TransactionSourceInboundInternationalACHTransferForeignExchangeIndicator = "fixed_to_variable"
	// The originator chose an amount to settle in USD. The originator's amount was
	// variable; known only after the foreign exchange conversion.
	TransactionSourceInboundInternationalACHTransferForeignExchangeIndicatorVariableToFixed TransactionSourceInboundInternationalACHTransferForeignExchangeIndicator = "variable_to_fixed"
	// The amount was originated and settled as a fixed amount in USD. There is no
	// foreign exchange conversion.
	TransactionSourceInboundInternationalACHTransferForeignExchangeIndicatorFixedToFixed TransactionSourceInboundInternationalACHTransferForeignExchangeIndicator = "fixed_to_fixed"
)

func (r TransactionSourceInboundInternationalACHTransferForeignExchangeIndicator) IsKnown() bool {
	switch r {
	case TransactionSourceInboundInternationalACHTransferForeignExchangeIndicatorFixedToVariable, TransactionSourceInboundInternationalACHTransferForeignExchangeIndicatorVariableToFixed, TransactionSourceInboundInternationalACHTransferForeignExchangeIndicatorFixedToFixed:
		return true
	}
	return false
}

// An instruction of how to interpret the `foreign_exchange_reference` field for
// this Transaction.
type TransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicator string

const (
	// The ACH file contains a foreign exchange rate.
	TransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicatorForeignExchangeRate TransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicator = "foreign_exchange_rate"
	// The ACH file contains a reference to a well-known foreign exchange rate.
	TransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicatorForeignExchangeReferenceNumber TransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicator = "foreign_exchange_reference_number"
	// There is no foreign exchange for this transfer, so the
	// `foreign_exchange_reference` field is blank.
	TransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicatorBlank TransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicator = "blank"
)

func (r TransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicator) IsKnown() bool {
	switch r {
	case TransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicatorForeignExchangeRate, TransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicatorForeignExchangeReferenceNumber, TransactionSourceInboundInternationalACHTransferForeignExchangeReferenceIndicatorBlank:
		return true
	}
	return false
}

// The type of transfer. Set by the originator.
type TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode string

const (
	// Sent as `ANN` in the Nacha file.
	TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeAnnuity TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "annuity"
	// Sent as `BUS` in the Nacha file.
	TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeBusinessOrCommercial TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "business_or_commercial"
	// Sent as `DEP` in the Nacha file.
	TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeDeposit TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "deposit"
	// Sent as `LOA` in the Nacha file.
	TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeLoan TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "loan"
	// Sent as `MIS` in the Nacha file.
	TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeMiscellaneous TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "miscellaneous"
	// Sent as `MOR` in the Nacha file.
	TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeMortgage TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "mortgage"
	// Sent as `PEN` in the Nacha file.
	TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodePension TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "pension"
	// Sent as `REM` in the Nacha file.
	TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeRemittance TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "remittance"
	// Sent as `RLS` in the Nacha file.
	TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeRentOrLease TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "rent_or_lease"
	// Sent as `SAL` in the Nacha file.
	TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeSalaryOrPayroll TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "salary_or_payroll"
	// Sent as `TAX` in the Nacha file.
	TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeTax TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "tax"
	// Sent as `ARC` in the Nacha file.
	TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeAccountsReceivable TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "accounts_receivable"
	// Sent as `BOC` in the Nacha file.
	TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeBackOfficeConversion TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "back_office_conversion"
	// Sent as `MTE` in the Nacha file.
	TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeMachineTransfer TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "machine_transfer"
	// Sent as `POP` in the Nacha file.
	TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodePointOfPurchase TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "point_of_purchase"
	// Sent as `POS` in the Nacha file.
	TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodePointOfSale TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "point_of_sale"
	// Sent as `RCK` in the Nacha file.
	TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeRepresentedCheck TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "represented_check"
	// Sent as `SHR` in the Nacha file.
	TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeSharedNetworkTransaction TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "shared_network_transaction"
	// Sent as `TEL` in the Nacha file.
	TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeTelphoneInitiated TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "telphone_initiated"
	// Sent as `WEB` in the Nacha file.
	TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeInternetInitiated TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode = "internet_initiated"
)

func (r TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCode) IsKnown() bool {
	switch r {
	case TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeAnnuity, TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeBusinessOrCommercial, TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeDeposit, TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeLoan, TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeMiscellaneous, TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeMortgage, TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodePension, TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeRemittance, TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeRentOrLease, TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeSalaryOrPayroll, TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeTax, TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeAccountsReceivable, TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeBackOfficeConversion, TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeMachineTransfer, TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodePointOfPurchase, TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodePointOfSale, TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeRepresentedCheck, TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeSharedNetworkTransaction, TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeTelphoneInitiated, TransactionSourceInboundInternationalACHTransferInternationalTransactionTypeCodeInternetInitiated:
		return true
	}
	return false
}

// An instruction of how to interpret the
// `originating_depository_financial_institution_id` field for this Transaction.
type TransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifier string

const (
	// A domestic clearing system number. In the US, for example, this is the American
	// Banking Association (ABA) routing number.
	TransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber TransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifier = "national_clearing_system_number"
	// The SWIFT Bank Identifier Code (BIC) of the bank.
	TransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifierBicCode TransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifier = "bic_code"
	// An International Bank Account Number.
	TransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifierIban TransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifier = "iban"
)

func (r TransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifier) IsKnown() bool {
	switch r {
	case TransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber, TransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifierBicCode, TransactionSourceInboundInternationalACHTransferOriginatingDepositoryFinancialInstitutionIDQualifierIban:
		return true
	}
	return false
}

// An instruction of how to interpret the
// `receiving_depository_financial_institution_id` field for this Transaction.
type TransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifier string

const (
	// A domestic clearing system number. In the US, for example, this is the American
	// Banking Association (ABA) routing number.
	TransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber TransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifier = "national_clearing_system_number"
	// The SWIFT Bank Identifier Code (BIC) of the bank.
	TransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifierBicCode TransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifier = "bic_code"
	// An International Bank Account Number.
	TransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifierIban TransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifier = "iban"
)

func (r TransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifier) IsKnown() bool {
	switch r {
	case TransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber, TransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifierBicCode, TransactionSourceInboundInternationalACHTransferReceivingDepositoryFinancialInstitutionIDQualifierIban:
		return true
	}
	return false
}

// An Inbound Real-Time Payments Transfer Confirmation object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_confirmation`.
type TransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real-Time Payments transfer.
	Currency TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency `json:"currency,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber string `json:"debtor_account_number,required"`
	// The name provided by the sender of the transfer.
	DebtorName string `json:"debtor_name,required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber string `json:"debtor_routing_number,required"`
	// Additional information included with the transfer.
	RemittanceInformation string `json:"remittance_information,required,nullable"`
	// The Real-Time Payments network identification of the transfer.
	TransactionIdentification string                                                           `json:"transaction_identification,required"`
	JSON                      transactionSourceInboundRealTimePaymentsTransferConfirmationJSON `json:"-"`
}

// transactionSourceInboundRealTimePaymentsTransferConfirmationJSON contains the
// JSON metadata for the struct
// [TransactionSourceInboundRealTimePaymentsTransferConfirmation]
type transactionSourceInboundRealTimePaymentsTransferConfirmationJSON struct {
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

func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceInboundRealTimePaymentsTransferConfirmationJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
// currency. This will always be "USD" for a Real-Time Payments transfer.
type TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency string

const (
	// Canadian Dollar (CAD)
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyCad TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CAD"
	// Swiss Franc (CHF)
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyChf TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CHF"
	// Euro (EUR)
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyEur TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "EUR"
	// British Pound (GBP)
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyGbp TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "GBP"
	// Japanese Yen (JPY)
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyJpy TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "JPY"
	// US Dollar (USD)
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyUsd TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "USD"
)

func (r TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyCad, TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyChf, TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyEur, TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyGbp, TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyJpy, TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyUsd:
		return true
	}
	return false
}

// An Inbound Wire Drawdown Payment object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
type TransactionSourceInboundWireDrawdownPayment struct {
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
	OriginatorToBeneficiaryInformationLine4 string                                          `json:"originator_to_beneficiary_information_line4,required,nullable"`
	JSON                                    transactionSourceInboundWireDrawdownPaymentJSON `json:"-"`
}

// transactionSourceInboundWireDrawdownPaymentJSON contains the JSON metadata for
// the struct [TransactionSourceInboundWireDrawdownPayment]
type transactionSourceInboundWireDrawdownPaymentJSON struct {
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

func (r *TransactionSourceInboundWireDrawdownPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceInboundWireDrawdownPaymentJSON) RawJSON() string {
	return r.raw
}

// An Inbound Wire Drawdown Payment Reversal object. This field will be present in
// the JSON response if and only if `category` is equal to
// `inbound_wire_drawdown_payment_reversal`.
type TransactionSourceInboundWireDrawdownPaymentReversal struct {
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
	PreviousMessageInputSource string                                                  `json:"previous_message_input_source,required"`
	JSON                       transactionSourceInboundWireDrawdownPaymentReversalJSON `json:"-"`
}

// transactionSourceInboundWireDrawdownPaymentReversalJSON contains the JSON
// metadata for the struct [TransactionSourceInboundWireDrawdownPaymentReversal]
type transactionSourceInboundWireDrawdownPaymentReversalJSON struct {
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

func (r *TransactionSourceInboundWireDrawdownPaymentReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceInboundWireDrawdownPaymentReversalJSON) RawJSON() string {
	return r.raw
}

// An Inbound Wire Reversal object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_reversal`.
type TransactionSourceInboundWireReversal struct {
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
	WireTransferID string                                   `json:"wire_transfer_id,required"`
	JSON           transactionSourceInboundWireReversalJSON `json:"-"`
}

// transactionSourceInboundWireReversalJSON contains the JSON metadata for the
// struct [TransactionSourceInboundWireReversal]
type transactionSourceInboundWireReversalJSON struct {
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

func (r *TransactionSourceInboundWireReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceInboundWireReversalJSON) RawJSON() string {
	return r.raw
}

// An Inbound Wire Transfer Intention object. This field will be present in the
// JSON response if and only if `category` is equal to `inbound_wire_transfer`.
type TransactionSourceInboundWireTransfer struct {
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
	TransferID string                                   `json:"transfer_id,required"`
	JSON       transactionSourceInboundWireTransferJSON `json:"-"`
}

// transactionSourceInboundWireTransferJSON contains the JSON metadata for the
// struct [TransactionSourceInboundWireTransfer]
type transactionSourceInboundWireTransferJSON struct {
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

func (r *TransactionSourceInboundWireTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceInboundWireTransferJSON) RawJSON() string {
	return r.raw
}

// An Interest Payment object. This field will be present in the JSON response if
// and only if `category` is equal to `interest_payment`.
type TransactionSourceInterestPayment struct {
	// The account on which the interest was accrued.
	AccruedOnAccountID string `json:"accrued_on_account_id,required,nullable"`
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency TransactionSourceInterestPaymentCurrency `json:"currency,required"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time                            `json:"period_start,required" format:"date-time"`
	JSON        transactionSourceInterestPaymentJSON `json:"-"`
}

// transactionSourceInterestPaymentJSON contains the JSON metadata for the struct
// [TransactionSourceInterestPayment]
type transactionSourceInterestPaymentJSON struct {
	AccruedOnAccountID apijson.Field
	Amount             apijson.Field
	Currency           apijson.Field
	PeriodEnd          apijson.Field
	PeriodStart        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TransactionSourceInterestPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceInterestPaymentJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type TransactionSourceInterestPaymentCurrency string

const (
	// Canadian Dollar (CAD)
	TransactionSourceInterestPaymentCurrencyCad TransactionSourceInterestPaymentCurrency = "CAD"
	// Swiss Franc (CHF)
	TransactionSourceInterestPaymentCurrencyChf TransactionSourceInterestPaymentCurrency = "CHF"
	// Euro (EUR)
	TransactionSourceInterestPaymentCurrencyEur TransactionSourceInterestPaymentCurrency = "EUR"
	// British Pound (GBP)
	TransactionSourceInterestPaymentCurrencyGbp TransactionSourceInterestPaymentCurrency = "GBP"
	// Japanese Yen (JPY)
	TransactionSourceInterestPaymentCurrencyJpy TransactionSourceInterestPaymentCurrency = "JPY"
	// US Dollar (USD)
	TransactionSourceInterestPaymentCurrencyUsd TransactionSourceInterestPaymentCurrency = "USD"
)

func (r TransactionSourceInterestPaymentCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceInterestPaymentCurrencyCad, TransactionSourceInterestPaymentCurrencyChf, TransactionSourceInterestPaymentCurrencyEur, TransactionSourceInterestPaymentCurrencyGbp, TransactionSourceInterestPaymentCurrencyJpy, TransactionSourceInterestPaymentCurrencyUsd:
		return true
	}
	return false
}

// An Internal Source object. This field will be present in the JSON response if
// and only if `category` is equal to `internal_source`.
type TransactionSourceInternalSource struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency TransactionSourceInternalSourceCurrency `json:"currency,required"`
	// An Internal Source is a transaction between you and Increase. This describes the
	// reason for the transaction.
	Reason TransactionSourceInternalSourceReason `json:"reason,required"`
	JSON   transactionSourceInternalSourceJSON   `json:"-"`
}

// transactionSourceInternalSourceJSON contains the JSON metadata for the struct
// [TransactionSourceInternalSource]
type transactionSourceInternalSourceJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceInternalSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceInternalSourceJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type TransactionSourceInternalSourceCurrency string

const (
	// Canadian Dollar (CAD)
	TransactionSourceInternalSourceCurrencyCad TransactionSourceInternalSourceCurrency = "CAD"
	// Swiss Franc (CHF)
	TransactionSourceInternalSourceCurrencyChf TransactionSourceInternalSourceCurrency = "CHF"
	// Euro (EUR)
	TransactionSourceInternalSourceCurrencyEur TransactionSourceInternalSourceCurrency = "EUR"
	// British Pound (GBP)
	TransactionSourceInternalSourceCurrencyGbp TransactionSourceInternalSourceCurrency = "GBP"
	// Japanese Yen (JPY)
	TransactionSourceInternalSourceCurrencyJpy TransactionSourceInternalSourceCurrency = "JPY"
	// US Dollar (USD)
	TransactionSourceInternalSourceCurrencyUsd TransactionSourceInternalSourceCurrency = "USD"
)

func (r TransactionSourceInternalSourceCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceInternalSourceCurrencyCad, TransactionSourceInternalSourceCurrencyChf, TransactionSourceInternalSourceCurrencyEur, TransactionSourceInternalSourceCurrencyGbp, TransactionSourceInternalSourceCurrencyJpy, TransactionSourceInternalSourceCurrencyUsd:
		return true
	}
	return false
}

// An Internal Source is a transaction between you and Increase. This describes the
// reason for the transaction.
type TransactionSourceInternalSourceReason string

const (
	// Account closure
	TransactionSourceInternalSourceReasonAccountClosure TransactionSourceInternalSourceReason = "account_closure"
	// Bank migration
	TransactionSourceInternalSourceReasonBankMigration TransactionSourceInternalSourceReason = "bank_migration"
	// Cashback
	TransactionSourceInternalSourceReasonCashback TransactionSourceInternalSourceReason = "cashback"
	// Check adjustment
	TransactionSourceInternalSourceReasonCheckAdjustment TransactionSourceInternalSourceReason = "check_adjustment"
	// Collection payment
	TransactionSourceInternalSourceReasonCollectionPayment TransactionSourceInternalSourceReason = "collection_payment"
	// Collection receivable
	TransactionSourceInternalSourceReasonCollectionReceivable TransactionSourceInternalSourceReason = "collection_receivable"
	// Empyreal adjustment
	TransactionSourceInternalSourceReasonEmpyrealAdjustment TransactionSourceInternalSourceReason = "empyreal_adjustment"
	// Error
	TransactionSourceInternalSourceReasonError TransactionSourceInternalSourceReason = "error"
	// Error correction
	TransactionSourceInternalSourceReasonErrorCorrection TransactionSourceInternalSourceReason = "error_correction"
	// Fees
	TransactionSourceInternalSourceReasonFees TransactionSourceInternalSourceReason = "fees"
	// Interest
	TransactionSourceInternalSourceReasonInterest TransactionSourceInternalSourceReason = "interest"
	// Negative balance forgiveness
	TransactionSourceInternalSourceReasonNegativeBalanceForgiveness TransactionSourceInternalSourceReason = "negative_balance_forgiveness"
	// Sample funds
	TransactionSourceInternalSourceReasonSampleFunds TransactionSourceInternalSourceReason = "sample_funds"
	// Sample funds return
	TransactionSourceInternalSourceReasonSampleFundsReturn TransactionSourceInternalSourceReason = "sample_funds_return"
)

func (r TransactionSourceInternalSourceReason) IsKnown() bool {
	switch r {
	case TransactionSourceInternalSourceReasonAccountClosure, TransactionSourceInternalSourceReasonBankMigration, TransactionSourceInternalSourceReasonCashback, TransactionSourceInternalSourceReasonCheckAdjustment, TransactionSourceInternalSourceReasonCollectionPayment, TransactionSourceInternalSourceReasonCollectionReceivable, TransactionSourceInternalSourceReasonEmpyrealAdjustment, TransactionSourceInternalSourceReasonError, TransactionSourceInternalSourceReasonErrorCorrection, TransactionSourceInternalSourceReasonFees, TransactionSourceInternalSourceReasonInterest, TransactionSourceInternalSourceReasonNegativeBalanceForgiveness, TransactionSourceInternalSourceReasonSampleFunds, TransactionSourceInternalSourceReasonSampleFundsReturn:
		return true
	}
	return false
}

// A Real-Time Payments Transfer Acknowledgement object. This field will be present
// in the JSON response if and only if `category` is equal to
// `real_time_payments_transfer_acknowledgement`.
type TransactionSourceRealTimePaymentsTransferAcknowledgement struct {
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The destination account number.
	DestinationAccountNumber string `json:"destination_account_number,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	DestinationRoutingNumber string `json:"destination_routing_number,required"`
	// Unstructured information that will show on the recipient's bank statement.
	RemittanceInformation string `json:"remittance_information,required"`
	// The identifier of the Real-Time Payments Transfer that led to this Transaction.
	TransferID string                                                       `json:"transfer_id,required"`
	JSON       transactionSourceRealTimePaymentsTransferAcknowledgementJSON `json:"-"`
}

// transactionSourceRealTimePaymentsTransferAcknowledgementJSON contains the JSON
// metadata for the struct
// [TransactionSourceRealTimePaymentsTransferAcknowledgement]
type transactionSourceRealTimePaymentsTransferAcknowledgementJSON struct {
	Amount                   apijson.Field
	DestinationAccountNumber apijson.Field
	DestinationRoutingNumber apijson.Field
	RemittanceInformation    apijson.Field
	TransferID               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *TransactionSourceRealTimePaymentsTransferAcknowledgement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceRealTimePaymentsTransferAcknowledgementJSON) RawJSON() string {
	return r.raw
}

// A Sample Funds object. This field will be present in the JSON response if and
// only if `category` is equal to `sample_funds`.
type TransactionSourceSampleFunds struct {
	// Where the sample funds came from.
	Originator string                           `json:"originator,required"`
	JSON       transactionSourceSampleFundsJSON `json:"-"`
}

// transactionSourceSampleFundsJSON contains the JSON metadata for the struct
// [TransactionSourceSampleFunds]
type transactionSourceSampleFundsJSON struct {
	Originator  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceSampleFunds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceSampleFundsJSON) RawJSON() string {
	return r.raw
}

// A Wire Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_intention`.
type TransactionSourceWireTransferIntention struct {
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient string `json:"message_to_recipient,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// The identifier of the Wire Transfer that led to this Transaction.
	TransferID string                                     `json:"transfer_id,required"`
	JSON       transactionSourceWireTransferIntentionJSON `json:"-"`
}

// transactionSourceWireTransferIntentionJSON contains the JSON metadata for the
// struct [TransactionSourceWireTransferIntention]
type transactionSourceWireTransferIntentionJSON struct {
	AccountNumber      apijson.Field
	Amount             apijson.Field
	MessageToRecipient apijson.Field
	RoutingNumber      apijson.Field
	TransferID         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TransactionSourceWireTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceWireTransferIntentionJSON) RawJSON() string {
	return r.raw
}

// A Wire Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_rejection`.
type TransactionSourceWireTransferRejection struct {
	// The identifier of the Wire Transfer that led to this Transaction.
	TransferID string                                     `json:"transfer_id,required"`
	JSON       transactionSourceWireTransferRejectionJSON `json:"-"`
}

// transactionSourceWireTransferRejectionJSON contains the JSON metadata for the
// struct [TransactionSourceWireTransferRejection]
type transactionSourceWireTransferRejectionJSON struct {
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceWireTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceWireTransferRejectionJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `transaction`.
type TransactionType string

const (
	TransactionTypeTransaction TransactionType = "transaction"
)

func (r TransactionType) IsKnown() bool {
	switch r {
	case TransactionTypeTransaction:
		return true
	}
	return false
}

type TransactionListParams struct {
	// Filter Transactions for those belonging to the specified Account.
	AccountID param.Field[string]                         `query:"account_id"`
	Category  param.Field[TransactionListParamsCategory]  `query:"category"`
	CreatedAt param.Field[TransactionListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
	// Filter Transactions for those belonging to the specified route. This could be a
	// Card ID or an Account Number ID.
	RouteID param.Field[string] `query:"route_id"`
}

// URLQuery serializes [TransactionListParams]'s query parameters as `url.Values`.
func (r TransactionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type TransactionListParamsCategory struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In param.Field[[]TransactionListParamsCategoryIn] `query:"in"`
}

// URLQuery serializes [TransactionListParamsCategory]'s query parameters as
// `url.Values`.
func (r TransactionListParamsCategory) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type TransactionListParamsCategoryIn string

const (
	// Account Transfer Intention: details will be under the
	// `account_transfer_intention` object.
	TransactionListParamsCategoryInAccountTransferIntention TransactionListParamsCategoryIn = "account_transfer_intention"
	// ACH Transfer Intention: details will be under the `ach_transfer_intention`
	// object.
	TransactionListParamsCategoryInACHTransferIntention TransactionListParamsCategoryIn = "ach_transfer_intention"
	// ACH Transfer Rejection: details will be under the `ach_transfer_rejection`
	// object.
	TransactionListParamsCategoryInACHTransferRejection TransactionListParamsCategoryIn = "ach_transfer_rejection"
	// ACH Transfer Return: details will be under the `ach_transfer_return` object.
	TransactionListParamsCategoryInACHTransferReturn TransactionListParamsCategoryIn = "ach_transfer_return"
	// Card Dispute Acceptance: details will be under the `card_dispute_acceptance`
	// object.
	TransactionListParamsCategoryInCardDisputeAcceptance TransactionListParamsCategoryIn = "card_dispute_acceptance"
	// Card Refund: details will be under the `card_refund` object.
	TransactionListParamsCategoryInCardRefund TransactionListParamsCategoryIn = "card_refund"
	// Card Settlement: details will be under the `card_settlement` object.
	TransactionListParamsCategoryInCardSettlement TransactionListParamsCategoryIn = "card_settlement"
	// Card Revenue Payment: details will be under the `card_revenue_payment` object.
	TransactionListParamsCategoryInCardRevenuePayment TransactionListParamsCategoryIn = "card_revenue_payment"
	// Check Deposit Acceptance: details will be under the `check_deposit_acceptance`
	// object.
	TransactionListParamsCategoryInCheckDepositAcceptance TransactionListParamsCategoryIn = "check_deposit_acceptance"
	// Check Deposit Return: details will be under the `check_deposit_return` object.
	TransactionListParamsCategoryInCheckDepositReturn TransactionListParamsCategoryIn = "check_deposit_return"
	// Check Transfer Deposit: details will be under the `check_transfer_deposit`
	// object.
	TransactionListParamsCategoryInCheckTransferDeposit TransactionListParamsCategoryIn = "check_transfer_deposit"
	// Check Transfer Stop Payment Request: details will be under the
	// `check_transfer_stop_payment_request` object.
	TransactionListParamsCategoryInCheckTransferStopPaymentRequest TransactionListParamsCategoryIn = "check_transfer_stop_payment_request"
	// Fee Payment: details will be under the `fee_payment` object.
	TransactionListParamsCategoryInFeePayment TransactionListParamsCategoryIn = "fee_payment"
	// Inbound ACH Transfer Intention: details will be under the `inbound_ach_transfer`
	// object.
	TransactionListParamsCategoryInInboundACHTransfer TransactionListParamsCategoryIn = "inbound_ach_transfer"
	// Inbound ACH Transfer Return Intention: details will be under the
	// `inbound_ach_transfer_return_intention` object.
	TransactionListParamsCategoryInInboundACHTransferReturnIntention TransactionListParamsCategoryIn = "inbound_ach_transfer_return_intention"
	// Inbound Check Deposit Return Intention: details will be under the
	// `inbound_check_deposit_return_intention` object.
	TransactionListParamsCategoryInInboundCheckDepositReturnIntention TransactionListParamsCategoryIn = "inbound_check_deposit_return_intention"
	// Inbound International ACH Transfer: details will be under the
	// `inbound_international_ach_transfer` object.
	TransactionListParamsCategoryInInboundInternationalACHTransfer TransactionListParamsCategoryIn = "inbound_international_ach_transfer"
	// Inbound Real-Time Payments Transfer Confirmation: details will be under the
	// `inbound_real_time_payments_transfer_confirmation` object.
	TransactionListParamsCategoryInInboundRealTimePaymentsTransferConfirmation TransactionListParamsCategoryIn = "inbound_real_time_payments_transfer_confirmation"
	// Inbound Wire Drawdown Payment Reversal: details will be under the
	// `inbound_wire_drawdown_payment_reversal` object.
	TransactionListParamsCategoryInInboundWireDrawdownPaymentReversal TransactionListParamsCategoryIn = "inbound_wire_drawdown_payment_reversal"
	// Inbound Wire Drawdown Payment: details will be under the
	// `inbound_wire_drawdown_payment` object.
	TransactionListParamsCategoryInInboundWireDrawdownPayment TransactionListParamsCategoryIn = "inbound_wire_drawdown_payment"
	// Inbound Wire Reversal: details will be under the `inbound_wire_reversal` object.
	TransactionListParamsCategoryInInboundWireReversal TransactionListParamsCategoryIn = "inbound_wire_reversal"
	// Inbound Wire Transfer Intention: details will be under the
	// `inbound_wire_transfer` object.
	TransactionListParamsCategoryInInboundWireTransfer TransactionListParamsCategoryIn = "inbound_wire_transfer"
	// Inbound Wire Transfer Reversal Intention: details will be under the
	// `inbound_wire_transfer_reversal` object.
	TransactionListParamsCategoryInInboundWireTransferReversal TransactionListParamsCategoryIn = "inbound_wire_transfer_reversal"
	// Interest Payment: details will be under the `interest_payment` object.
	TransactionListParamsCategoryInInterestPayment TransactionListParamsCategoryIn = "interest_payment"
	// Internal Source: details will be under the `internal_source` object.
	TransactionListParamsCategoryInInternalSource TransactionListParamsCategoryIn = "internal_source"
	// Real-Time Payments Transfer Acknowledgement: details will be under the
	// `real_time_payments_transfer_acknowledgement` object.
	TransactionListParamsCategoryInRealTimePaymentsTransferAcknowledgement TransactionListParamsCategoryIn = "real_time_payments_transfer_acknowledgement"
	// Sample Funds: details will be under the `sample_funds` object.
	TransactionListParamsCategoryInSampleFunds TransactionListParamsCategoryIn = "sample_funds"
	// Wire Transfer Intention: details will be under the `wire_transfer_intention`
	// object.
	TransactionListParamsCategoryInWireTransferIntention TransactionListParamsCategoryIn = "wire_transfer_intention"
	// Wire Transfer Rejection: details will be under the `wire_transfer_rejection`
	// object.
	TransactionListParamsCategoryInWireTransferRejection TransactionListParamsCategoryIn = "wire_transfer_rejection"
	// The Transaction was made for an undocumented or deprecated reason.
	TransactionListParamsCategoryInOther TransactionListParamsCategoryIn = "other"
)

func (r TransactionListParamsCategoryIn) IsKnown() bool {
	switch r {
	case TransactionListParamsCategoryInAccountTransferIntention, TransactionListParamsCategoryInACHTransferIntention, TransactionListParamsCategoryInACHTransferRejection, TransactionListParamsCategoryInACHTransferReturn, TransactionListParamsCategoryInCardDisputeAcceptance, TransactionListParamsCategoryInCardRefund, TransactionListParamsCategoryInCardSettlement, TransactionListParamsCategoryInCardRevenuePayment, TransactionListParamsCategoryInCheckDepositAcceptance, TransactionListParamsCategoryInCheckDepositReturn, TransactionListParamsCategoryInCheckTransferDeposit, TransactionListParamsCategoryInCheckTransferStopPaymentRequest, TransactionListParamsCategoryInFeePayment, TransactionListParamsCategoryInInboundACHTransfer, TransactionListParamsCategoryInInboundACHTransferReturnIntention, TransactionListParamsCategoryInInboundCheckDepositReturnIntention, TransactionListParamsCategoryInInboundInternationalACHTransfer, TransactionListParamsCategoryInInboundRealTimePaymentsTransferConfirmation, TransactionListParamsCategoryInInboundWireDrawdownPaymentReversal, TransactionListParamsCategoryInInboundWireDrawdownPayment, TransactionListParamsCategoryInInboundWireReversal, TransactionListParamsCategoryInInboundWireTransfer, TransactionListParamsCategoryInInboundWireTransferReversal, TransactionListParamsCategoryInInterestPayment, TransactionListParamsCategoryInInternalSource, TransactionListParamsCategoryInRealTimePaymentsTransferAcknowledgement, TransactionListParamsCategoryInSampleFunds, TransactionListParamsCategoryInWireTransferIntention, TransactionListParamsCategoryInWireTransferRejection, TransactionListParamsCategoryInOther:
		return true
	}
	return false
}

type TransactionListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After param.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter param.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore param.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes [TransactionListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r TransactionListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
