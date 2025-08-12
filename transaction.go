// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
	"github.com/Increase/increase-go/packages/pagination"
)

// TransactionService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionService] method instead.
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
	if transactionID == "" {
		err = errors.New("missing required transaction_id parameter")
		return
	}
	path := fmt.Sprintf("transactions/%s", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Transactions
func (r *TransactionService) List(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) (res *pagination.Page[Transaction], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
// account. They're the equivalent of line items on your bank statement. To learn
// more, see [Transactions and Transfers](/documentation/transactions-transfers).
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
	TransactionCurrencyCad TransactionCurrency = "CAD"
	TransactionCurrencyChf TransactionCurrency = "CHF"
	TransactionCurrencyEur TransactionCurrency = "EUR"
	TransactionCurrencyGbp TransactionCurrency = "GBP"
	TransactionCurrencyJpy TransactionCurrency = "JPY"
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
	TransactionRouteTypeAccountNumber TransactionRouteType = "account_number"
	TransactionRouteTypeCard          TransactionRouteType = "card"
	TransactionRouteTypeLockbox       TransactionRouteType = "lockbox"
)

func (r TransactionRouteType) IsKnown() bool {
	switch r {
	case TransactionRouteTypeAccountNumber, TransactionRouteTypeCard, TransactionRouteTypeLockbox:
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
	// response if and only if `category` is equal to `account_transfer_intention`. Two
	// Account Transfer Intentions are created from each Account Transfer. One
	// decrements the source account, and the other increments the destination account.
	AccountTransferIntention TransactionSourceAccountTransferIntention `json:"account_transfer_intention,required,nullable"`
	// An ACH Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_intention`. An ACH
	// Transfer Intention is created from an ACH Transfer. It reflects the intention to
	// move money into or out of an Increase account via the ACH network.
	ACHTransferIntention TransactionSourceACHTransferIntention `json:"ach_transfer_intention,required,nullable"`
	// An ACH Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_rejection`. An ACH
	// Transfer Rejection is created when an ACH Transfer is rejected by Increase. It
	// offsets the ACH Transfer Intention. These rejections are rare.
	ACHTransferRejection TransactionSourceACHTransferRejection `json:"ach_transfer_rejection,required,nullable"`
	// An ACH Transfer Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_return`. An ACH Transfer
	// Return is created when an ACH Transfer is returned by the receiving bank. It
	// offsets the ACH Transfer Intention. ACH Transfer Returns usually occur within
	// the first two business days after the transfer is initiated, but can occur much
	// later.
	ACHTransferReturn TransactionSourceACHTransferReturn `json:"ach_transfer_return,required,nullable"`
	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	// Contains the details of a successful Card Dispute.
	CardDisputeAcceptance TransactionSourceCardDisputeAcceptance `json:"card_dispute_acceptance,required,nullable"`
	// A Card Dispute Financial object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_dispute_financial`. Financial event
	// related to a Card Dispute.
	CardDisputeFinancial TransactionSourceCardDisputeFinancial `json:"card_dispute_financial,required,nullable"`
	// A Card Dispute Loss object. This field will be present in the JSON response if
	// and only if `category` is equal to `card_dispute_loss`. Contains the details of
	// a lost Card Dispute.
	CardDisputeLoss TransactionSourceCardDisputeLoss `json:"card_dispute_loss,required,nullable"`
	// A Card Push Transfer Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_push_transfer_acceptance`.
	// A Card Push Transfer Acceptance is created when an Outbound Card Push Transfer
	// sent from Increase is accepted by the receiving bank.
	CardPushTransferAcceptance TransactionSourceCardPushTransferAcceptance `json:"card_push_transfer_acceptance,required,nullable"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`. Card Refunds move money back to
	// the cardholder. While they are usually connected to a Card Settlement an
	// acquirer can also refund money directly to a card without relation to a
	// transaction.
	CardRefund TransactionSourceCardRefund `json:"card_refund,required,nullable"`
	// A Card Revenue Payment object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_revenue_payment`. Card Revenue
	// Payments reflect earnings from fees on card transactions.
	CardRevenuePayment TransactionSourceCardRevenuePayment `json:"card_revenue_payment,required,nullable"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`. Card Settlements are card
	// transactions that have cleared and settled. While a settlement is usually
	// preceded by an authorization, an acquirer can also directly clear a transaction
	// without first authorizing it.
	CardSettlement TransactionSourceCardSettlement `json:"card_settlement,required,nullable"`
	// A Cashback Payment object. This field will be present in the JSON response if
	// and only if `category` is equal to `cashback_payment`. A Cashback Payment
	// represents the cashback paid to a cardholder for a given period. Cashback is
	// usually paid monthly for the prior month's transactions.
	CashbackPayment TransactionSourceCashbackPayment `json:"cashback_payment,required,nullable"`
	// The type of the resource. We may add additional possible values for this enum
	// over time; your application should be able to handle such additions gracefully.
	Category TransactionSourceCategory `json:"category,required"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`. A
	// Check Deposit Acceptance is created when a Check Deposit is processed and its
	// details confirmed. Check Deposits may be returned by the receiving bank, which
	// will appear as a Check Deposit Return.
	CheckDepositAcceptance TransactionSourceCheckDepositAcceptance `json:"check_deposit_acceptance,required,nullable"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`. A Check Deposit
	// Return is created when a Check Deposit is returned by the bank holding the
	// account it was drawn against. Check Deposits may be returned for a variety of
	// reasons, including insufficient funds or a mismatched account number. Usually,
	// checks are returned within the first 7 days after the deposit is made.
	CheckDepositReturn TransactionSourceCheckDepositReturn `json:"check_deposit_return,required,nullable"`
	// A Check Transfer Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_transfer_deposit`. An Inbound Check
	// is a check drawn on an Increase account that has been deposited by an external
	// bank account. These types of checks are not pre-registered.
	CheckTransferDeposit TransactionSourceCheckTransferDeposit `json:"check_transfer_deposit,required,nullable"`
	// A Fee Payment object. This field will be present in the JSON response if and
	// only if `category` is equal to `fee_payment`. A Fee Payment represents a payment
	// made to Increase.
	FeePayment TransactionSourceFeePayment `json:"fee_payment,required,nullable"`
	// An Inbound ACH Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_ach_transfer`. An
	// Inbound ACH Transfer Intention is created when an ACH transfer is initiated at
	// another bank and received by Increase.
	InboundACHTransfer TransactionSourceInboundACHTransfer `json:"inbound_ach_transfer,required,nullable"`
	// An Inbound ACH Transfer Return Intention object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_ach_transfer_return_intention`. An Inbound ACH Transfer Return
	// Intention is created when an ACH transfer is initiated at another bank and
	// returned by Increase.
	InboundACHTransferReturnIntention TransactionSourceInboundACHTransferReturnIntention `json:"inbound_ach_transfer_return_intention,required,nullable"`
	// An Inbound Check Adjustment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_check_adjustment`. An
	// Inbound Check Adjustment is created when Increase receives an adjustment for a
	// check or return deposited through Check21.
	InboundCheckAdjustment TransactionSourceInboundCheckAdjustment `json:"inbound_check_adjustment,required,nullable"`
	// An Inbound Check Deposit Return Intention object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_check_deposit_return_intention`. An Inbound Check Deposit Return
	// Intention is created when Increase receives an Inbound Check and the User
	// requests that it be returned.
	InboundCheckDepositReturnIntention TransactionSourceInboundCheckDepositReturnIntention `json:"inbound_check_deposit_return_intention,required,nullable"`
	// An Inbound Real-Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`. An Inbound Real-Time
	// Payments Transfer Confirmation is created when a Real-Time Payments transfer is
	// initiated at another bank and received by Increase.
	InboundRealTimePaymentsTransferConfirmation TransactionSourceInboundRealTimePaymentsTransferConfirmation `json:"inbound_real_time_payments_transfer_confirmation,required,nullable"`
	// An Inbound Real-Time Payments Transfer Decline object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline TransactionSourceInboundRealTimePaymentsTransferDecline `json:"inbound_real_time_payments_transfer_decline,required,nullable"`
	// An Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`. An Inbound Wire
	// Reversal represents a reversal of a wire transfer that was initiated via
	// Increase. The other bank is sending the money back. This most often happens when
	// the original destination account details were incorrect.
	InboundWireReversal TransactionSourceInboundWireReversal `json:"inbound_wire_reversal,required,nullable"`
	// An Inbound Wire Transfer Intention object. This field will be present in the
	// JSON response if and only if `category` is equal to `inbound_wire_transfer`. An
	// Inbound Wire Transfer Intention is created when a wire transfer is initiated at
	// another bank and received by Increase.
	InboundWireTransfer TransactionSourceInboundWireTransfer `json:"inbound_wire_transfer,required,nullable"`
	// An Inbound Wire Transfer Reversal Intention object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_wire_transfer_reversal`. An Inbound Wire Transfer Reversal Intention is
	// created when Increase has received a wire and the User requests that it be
	// reversed.
	InboundWireTransferReversal TransactionSourceInboundWireTransferReversal `json:"inbound_wire_transfer_reversal,required,nullable"`
	// An Interest Payment object. This field will be present in the JSON response if
	// and only if `category` is equal to `interest_payment`. An Interest Payment
	// represents a payment of interest on an account. Interest is usually paid
	// monthly.
	InterestPayment TransactionSourceInterestPayment `json:"interest_payment,required,nullable"`
	// An Internal Source object. This field will be present in the JSON response if
	// and only if `category` is equal to `internal_source`. A transaction between the
	// user and Increase. See the `reason` attribute for more information.
	InternalSource TransactionSourceInternalSource `json:"internal_source,required,nullable"`
	// If the category of this Transaction source is equal to `other`, this field will
	// contain an empty object, otherwise it will contain null.
	Other interface{} `json:"other,required,nullable"`
	// A Real-Time Payments Transfer Acknowledgement object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `real_time_payments_transfer_acknowledgement`. A Real-Time Payments Transfer
	// Acknowledgement is created when a Real-Time Payments Transfer sent from Increase
	// is acknowledged by the receiving bank.
	RealTimePaymentsTransferAcknowledgement TransactionSourceRealTimePaymentsTransferAcknowledgement `json:"real_time_payments_transfer_acknowledgement,required,nullable"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`. Sample funds for testing
	// purposes.
	SampleFunds TransactionSourceSampleFunds `json:"sample_funds,required,nullable"`
	// A Swift Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `swift_transfer_intention`. A
	// Swift Transfer initiated via Increase.
	SwiftTransferIntention TransactionSourceSwiftTransferIntention `json:"swift_transfer_intention,required,nullable"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`. A Wire
	// Transfer initiated via Increase and sent to a different bank.
	WireTransferIntention TransactionSourceWireTransferIntention `json:"wire_transfer_intention,required,nullable"`
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
	CardDisputeFinancial                        apijson.Field
	CardDisputeLoss                             apijson.Field
	CardPushTransferAcceptance                  apijson.Field
	CardRefund                                  apijson.Field
	CardRevenuePayment                          apijson.Field
	CardSettlement                              apijson.Field
	CashbackPayment                             apijson.Field
	Category                                    apijson.Field
	CheckDepositAcceptance                      apijson.Field
	CheckDepositReturn                          apijson.Field
	CheckTransferDeposit                        apijson.Field
	FeePayment                                  apijson.Field
	InboundACHTransfer                          apijson.Field
	InboundACHTransferReturnIntention           apijson.Field
	InboundCheckAdjustment                      apijson.Field
	InboundCheckDepositReturnIntention          apijson.Field
	InboundRealTimePaymentsTransferConfirmation apijson.Field
	InboundRealTimePaymentsTransferDecline      apijson.Field
	InboundWireReversal                         apijson.Field
	InboundWireTransfer                         apijson.Field
	InboundWireTransferReversal                 apijson.Field
	InterestPayment                             apijson.Field
	InternalSource                              apijson.Field
	Other                                       apijson.Field
	RealTimePaymentsTransferAcknowledgement     apijson.Field
	SampleFunds                                 apijson.Field
	SwiftTransferIntention                      apijson.Field
	WireTransferIntention                       apijson.Field
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
// response if and only if `category` is equal to `account_transfer_intention`. Two
// Account Transfer Intentions are created from each Account Transfer. One
// decrements the source account, and the other increments the destination account.
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
	TransactionSourceAccountTransferIntentionCurrencyCad TransactionSourceAccountTransferIntentionCurrency = "CAD"
	TransactionSourceAccountTransferIntentionCurrencyChf TransactionSourceAccountTransferIntentionCurrency = "CHF"
	TransactionSourceAccountTransferIntentionCurrencyEur TransactionSourceAccountTransferIntentionCurrency = "EUR"
	TransactionSourceAccountTransferIntentionCurrencyGbp TransactionSourceAccountTransferIntentionCurrency = "GBP"
	TransactionSourceAccountTransferIntentionCurrencyJpy TransactionSourceAccountTransferIntentionCurrency = "JPY"
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
// response if and only if `category` is equal to `ach_transfer_intention`. An ACH
// Transfer Intention is created from an ACH Transfer. It reflects the intention to
// move money into or out of an Increase account via the ACH network.
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
// response if and only if `category` is equal to `ach_transfer_rejection`. An ACH
// Transfer Rejection is created when an ACH Transfer is rejected by Increase. It
// offsets the ACH Transfer Intention. These rejections are rare.
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
// if and only if `category` is equal to `ach_transfer_return`. An ACH Transfer
// Return is created when an ACH Transfer is returned by the receiving bank. It
// offsets the ACH Transfer Intention. ACH Transfer Returns usually occur within
// the first two business days after the transfer is initiated, but can occur much
// later.
type TransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The three character ACH return code, in the range R01 to R85.
	RawReturnReasonCode string `json:"raw_return_reason_code,required"`
	// Why the ACH Transfer was returned. This reason code is sent by the receiving
	// bank back to Increase.
	ReturnReasonCode TransactionSourceACHTransferReturnReturnReasonCode `json:"return_reason_code,required"`
	// A 15 digit number that was generated by the bank that initiated the return. The
	// trace number of the return is different than that of the original transfer. ACH
	// trace numbers are not unique, but along with the amount and date this number can
	// be used to identify the ACH return at the bank that initiated it.
	TraceNumber string `json:"trace_number,required"`
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
	TraceNumber         apijson.Field
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
	TransactionSourceACHTransferReturnReturnReasonCodeInsufficientFund                                            TransactionSourceACHTransferReturnReturnReasonCode = "insufficient_fund"
	TransactionSourceACHTransferReturnReturnReasonCodeNoAccount                                                   TransactionSourceACHTransferReturnReturnReasonCode = "no_account"
	TransactionSourceACHTransferReturnReturnReasonCodeAccountClosed                                               TransactionSourceACHTransferReturnReturnReasonCode = "account_closed"
	TransactionSourceACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure                               TransactionSourceACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	TransactionSourceACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction                TransactionSourceACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	TransactionSourceACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver                                TransactionSourceACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	TransactionSourceACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode     TransactionSourceACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	TransactionSourceACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized                       TransactionSourceACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	TransactionSourceACHTransferReturnReturnReasonCodePaymentStopped                                              TransactionSourceACHTransferReturnReturnReasonCode = "payment_stopped"
	TransactionSourceACHTransferReturnReturnReasonCodeNonTransactionAccount                                       TransactionSourceACHTransferReturnReturnReasonCode = "non_transaction_account"
	TransactionSourceACHTransferReturnReturnReasonCodeUncollectedFunds                                            TransactionSourceACHTransferReturnReturnReasonCode = "uncollected_funds"
	TransactionSourceACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError                                TransactionSourceACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	TransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete   TransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	TransactionSourceACHTransferReturnReturnReasonCodeAmountFieldError                                            TransactionSourceACHTransferReturnReturnReasonCode = "amount_field_error"
	TransactionSourceACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer                              TransactionSourceACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	TransactionSourceACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber                                     TransactionSourceACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	TransactionSourceACHTransferReturnReturnReasonCodeFileRecordEditCriteria                                      TransactionSourceACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	TransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualName                                    TransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	TransactionSourceACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest                                      TransactionSourceACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	TransactionSourceACHTransferReturnReturnReasonCodeLimitedParticipationDfi                                     TransactionSourceACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	TransactionSourceACHTransferReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment                TransactionSourceACHTransferReturnReturnReasonCode = "incorrectly_coded_outbound_international_payment"
	TransactionSourceACHTransferReturnReturnReasonCodeAccountSoldToAnotherDfi                                     TransactionSourceACHTransferReturnReturnReasonCode = "account_sold_to_another_dfi"
	TransactionSourceACHTransferReturnReturnReasonCodeAddendaError                                                TransactionSourceACHTransferReturnReturnReasonCode = "addenda_error"
	TransactionSourceACHTransferReturnReturnReasonCodeBeneficiaryOrAccountHolderDeceased                          TransactionSourceACHTransferReturnReturnReasonCode = "beneficiary_or_account_holder_deceased"
	TransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedNotWithinAuthorizationTerms                  TransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_not_within_authorization_terms"
	TransactionSourceACHTransferReturnReturnReasonCodeCorrectedReturn                                             TransactionSourceACHTransferReturnReturnReasonCode = "corrected_return"
	TransactionSourceACHTransferReturnReturnReasonCodeDuplicateEntry                                              TransactionSourceACHTransferReturnReturnReasonCode = "duplicate_entry"
	TransactionSourceACHTransferReturnReturnReasonCodeDuplicateReturn                                             TransactionSourceACHTransferReturnReturnReasonCode = "duplicate_return"
	TransactionSourceACHTransferReturnReturnReasonCodeEnrDuplicateEnrollment                                      TransactionSourceACHTransferReturnReturnReasonCode = "enr_duplicate_enrollment"
	TransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidDfiAccountNumber                                  TransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_dfi_account_number"
	TransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualIDNumber                                TransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_id_number"
	TransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidRepresentativePayeeIndicator                      TransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_representative_payee_indicator"
	TransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidTransactionCode                                   TransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_transaction_code"
	TransactionSourceACHTransferReturnReturnReasonCodeEnrReturnOfEnrEntry                                         TransactionSourceACHTransferReturnReturnReasonCode = "enr_return_of_enr_entry"
	TransactionSourceACHTransferReturnReturnReasonCodeEnrRoutingNumberCheckDigitError                             TransactionSourceACHTransferReturnReturnReasonCode = "enr_routing_number_check_digit_error"
	TransactionSourceACHTransferReturnReturnReasonCodeEntryNotProcessedByGateway                                  TransactionSourceACHTransferReturnReturnReasonCode = "entry_not_processed_by_gateway"
	TransactionSourceACHTransferReturnReturnReasonCodeFieldError                                                  TransactionSourceACHTransferReturnReturnReasonCode = "field_error"
	TransactionSourceACHTransferReturnReturnReasonCodeForeignReceivingDfiUnableToSettle                           TransactionSourceACHTransferReturnReturnReasonCode = "foreign_receiving_dfi_unable_to_settle"
	TransactionSourceACHTransferReturnReturnReasonCodeIatEntryCodingError                                         TransactionSourceACHTransferReturnReturnReasonCode = "iat_entry_coding_error"
	TransactionSourceACHTransferReturnReturnReasonCodeImproperEffectiveEntryDate                                  TransactionSourceACHTransferReturnReturnReasonCode = "improper_effective_entry_date"
	TransactionSourceACHTransferReturnReturnReasonCodeImproperSourceDocumentSourceDocumentPresented               TransactionSourceACHTransferReturnReturnReasonCode = "improper_source_document_source_document_presented"
	TransactionSourceACHTransferReturnReturnReasonCodeInvalidCompanyID                                            TransactionSourceACHTransferReturnReturnReasonCode = "invalid_company_id"
	TransactionSourceACHTransferReturnReturnReasonCodeInvalidForeignReceivingDfiIdentification                    TransactionSourceACHTransferReturnReturnReasonCode = "invalid_foreign_receiving_dfi_identification"
	TransactionSourceACHTransferReturnReturnReasonCodeInvalidIndividualIDNumber                                   TransactionSourceACHTransferReturnReturnReasonCode = "invalid_individual_id_number"
	TransactionSourceACHTransferReturnReturnReasonCodeItemAndRckEntryPresentedForPayment                          TransactionSourceACHTransferReturnReturnReasonCode = "item_and_rck_entry_presented_for_payment"
	TransactionSourceACHTransferReturnReturnReasonCodeItemRelatedToRckEntryIsIneligible                           TransactionSourceACHTransferReturnReturnReasonCode = "item_related_to_rck_entry_is_ineligible"
	TransactionSourceACHTransferReturnReturnReasonCodeMandatoryFieldError                                         TransactionSourceACHTransferReturnReturnReasonCode = "mandatory_field_error"
	TransactionSourceACHTransferReturnReturnReasonCodeMisroutedDishonoredReturn                                   TransactionSourceACHTransferReturnReturnReasonCode = "misrouted_dishonored_return"
	TransactionSourceACHTransferReturnReturnReasonCodeMisroutedReturn                                             TransactionSourceACHTransferReturnReturnReasonCode = "misrouted_return"
	TransactionSourceACHTransferReturnReturnReasonCodeNoErrorsFound                                               TransactionSourceACHTransferReturnReturnReasonCode = "no_errors_found"
	TransactionSourceACHTransferReturnReturnReasonCodeNonAcceptanceOfR62DishonoredReturn                          TransactionSourceACHTransferReturnReturnReasonCode = "non_acceptance_of_r62_dishonored_return"
	TransactionSourceACHTransferReturnReturnReasonCodeNonParticipantInIatProgram                                  TransactionSourceACHTransferReturnReturnReasonCode = "non_participant_in_iat_program"
	TransactionSourceACHTransferReturnReturnReasonCodePermissibleReturnEntry                                      TransactionSourceACHTransferReturnReturnReasonCode = "permissible_return_entry"
	TransactionSourceACHTransferReturnReturnReasonCodePermissibleReturnEntryNotAccepted                           TransactionSourceACHTransferReturnReturnReasonCode = "permissible_return_entry_not_accepted"
	TransactionSourceACHTransferReturnReturnReasonCodeRdfiNonSettlement                                           TransactionSourceACHTransferReturnReturnReasonCode = "rdfi_non_settlement"
	TransactionSourceACHTransferReturnReturnReasonCodeRdfiParticipantInCheckTruncationProgram                     TransactionSourceACHTransferReturnReturnReasonCode = "rdfi_participant_in_check_truncation_program"
	TransactionSourceACHTransferReturnReturnReasonCodeRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity TransactionSourceACHTransferReturnReturnReasonCode = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	TransactionSourceACHTransferReturnReturnReasonCodeReturnNotADuplicate                                         TransactionSourceACHTransferReturnReturnReasonCode = "return_not_a_duplicate"
	TransactionSourceACHTransferReturnReturnReasonCodeReturnOfErroneousOrReversingDebit                           TransactionSourceACHTransferReturnReturnReasonCode = "return_of_erroneous_or_reversing_debit"
	TransactionSourceACHTransferReturnReturnReasonCodeReturnOfImproperCreditEntry                                 TransactionSourceACHTransferReturnReturnReasonCode = "return_of_improper_credit_entry"
	TransactionSourceACHTransferReturnReturnReasonCodeReturnOfImproperDebitEntry                                  TransactionSourceACHTransferReturnReturnReasonCode = "return_of_improper_debit_entry"
	TransactionSourceACHTransferReturnReturnReasonCodeReturnOfXckEntry                                            TransactionSourceACHTransferReturnReturnReasonCode = "return_of_xck_entry"
	TransactionSourceACHTransferReturnReturnReasonCodeSourceDocumentPresentedForPayment                           TransactionSourceACHTransferReturnReturnReasonCode = "source_document_presented_for_payment"
	TransactionSourceACHTransferReturnReturnReasonCodeStateLawAffectingRckAcceptance                              TransactionSourceACHTransferReturnReturnReasonCode = "state_law_affecting_rck_acceptance"
	TransactionSourceACHTransferReturnReturnReasonCodeStopPaymentOnItemRelatedToRckEntry                          TransactionSourceACHTransferReturnReturnReasonCode = "stop_payment_on_item_related_to_rck_entry"
	TransactionSourceACHTransferReturnReturnReasonCodeStopPaymentOnSourceDocument                                 TransactionSourceACHTransferReturnReturnReasonCode = "stop_payment_on_source_document"
	TransactionSourceACHTransferReturnReturnReasonCodeTimelyOriginalReturn                                        TransactionSourceACHTransferReturnReturnReasonCode = "timely_original_return"
	TransactionSourceACHTransferReturnReturnReasonCodeTraceNumberError                                            TransactionSourceACHTransferReturnReturnReasonCode = "trace_number_error"
	TransactionSourceACHTransferReturnReturnReasonCodeUntimelyDishonoredReturn                                    TransactionSourceACHTransferReturnReturnReasonCode = "untimely_dishonored_return"
	TransactionSourceACHTransferReturnReturnReasonCodeUntimelyReturn                                              TransactionSourceACHTransferReturnReturnReasonCode = "untimely_return"
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
// Contains the details of a successful Card Dispute.
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

// A Card Dispute Financial object. This field will be present in the JSON response
// if and only if `category` is equal to `card_dispute_financial`. Financial event
// related to a Card Dispute.
type TransactionSourceCardDisputeFinancial struct {
	// The amount of the financial event.
	Amount int64 `json:"amount,required"`
	// The identifier of the Card Dispute the financial event is associated with.
	CardDisputeID string `json:"card_dispute_id,required"`
	// The network that the Card Dispute is associated with.
	Network TransactionSourceCardDisputeFinancialNetwork `json:"network,required"`
	// The identifier of the Transaction that was created to credit or debit the
	// disputed funds to or from your account.
	TransactionID string `json:"transaction_id,required"`
	// Information for events related to card dispute for card payments processed over
	// Visa's network. This field will be present in the JSON response if and only if
	// `network` is equal to `visa`.
	Visa TransactionSourceCardDisputeFinancialVisa `json:"visa,required,nullable"`
	JSON transactionSourceCardDisputeFinancialJSON `json:"-"`
}

// transactionSourceCardDisputeFinancialJSON contains the JSON metadata for the
// struct [TransactionSourceCardDisputeFinancial]
type transactionSourceCardDisputeFinancialJSON struct {
	Amount        apijson.Field
	CardDisputeID apijson.Field
	Network       apijson.Field
	TransactionID apijson.Field
	Visa          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TransactionSourceCardDisputeFinancial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardDisputeFinancialJSON) RawJSON() string {
	return r.raw
}

// The network that the Card Dispute is associated with.
type TransactionSourceCardDisputeFinancialNetwork string

const (
	TransactionSourceCardDisputeFinancialNetworkVisa TransactionSourceCardDisputeFinancialNetwork = "visa"
)

func (r TransactionSourceCardDisputeFinancialNetwork) IsKnown() bool {
	switch r {
	case TransactionSourceCardDisputeFinancialNetworkVisa:
		return true
	}
	return false
}

// Information for events related to card dispute for card payments processed over
// Visa's network. This field will be present in the JSON response if and only if
// `network` is equal to `visa`.
type TransactionSourceCardDisputeFinancialVisa struct {
	// The type of card dispute financial event.
	EventType TransactionSourceCardDisputeFinancialVisaEventType `json:"event_type,required"`
	JSON      transactionSourceCardDisputeFinancialVisaJSON      `json:"-"`
}

// transactionSourceCardDisputeFinancialVisaJSON contains the JSON metadata for the
// struct [TransactionSourceCardDisputeFinancialVisa]
type transactionSourceCardDisputeFinancialVisaJSON struct {
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardDisputeFinancialVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardDisputeFinancialVisaJSON) RawJSON() string {
	return r.raw
}

// The type of card dispute financial event.
type TransactionSourceCardDisputeFinancialVisaEventType string

const (
	TransactionSourceCardDisputeFinancialVisaEventTypeChargebackSubmitted                    TransactionSourceCardDisputeFinancialVisaEventType = "chargeback_submitted"
	TransactionSourceCardDisputeFinancialVisaEventTypeMerchantPrearbitrationDeclineSubmitted TransactionSourceCardDisputeFinancialVisaEventType = "merchant_prearbitration_decline_submitted"
	TransactionSourceCardDisputeFinancialVisaEventTypeMerchantPrearbitrationReceived         TransactionSourceCardDisputeFinancialVisaEventType = "merchant_prearbitration_received"
	TransactionSourceCardDisputeFinancialVisaEventTypeRepresented                            TransactionSourceCardDisputeFinancialVisaEventType = "represented"
	TransactionSourceCardDisputeFinancialVisaEventTypeUserPrearbitrationDeclineReceived      TransactionSourceCardDisputeFinancialVisaEventType = "user_prearbitration_decline_received"
	TransactionSourceCardDisputeFinancialVisaEventTypeUserPrearbitrationSubmitted            TransactionSourceCardDisputeFinancialVisaEventType = "user_prearbitration_submitted"
	TransactionSourceCardDisputeFinancialVisaEventTypeUserWithdrawalSubmitted                TransactionSourceCardDisputeFinancialVisaEventType = "user_withdrawal_submitted"
)

func (r TransactionSourceCardDisputeFinancialVisaEventType) IsKnown() bool {
	switch r {
	case TransactionSourceCardDisputeFinancialVisaEventTypeChargebackSubmitted, TransactionSourceCardDisputeFinancialVisaEventTypeMerchantPrearbitrationDeclineSubmitted, TransactionSourceCardDisputeFinancialVisaEventTypeMerchantPrearbitrationReceived, TransactionSourceCardDisputeFinancialVisaEventTypeRepresented, TransactionSourceCardDisputeFinancialVisaEventTypeUserPrearbitrationDeclineReceived, TransactionSourceCardDisputeFinancialVisaEventTypeUserPrearbitrationSubmitted, TransactionSourceCardDisputeFinancialVisaEventTypeUserWithdrawalSubmitted:
		return true
	}
	return false
}

// A Card Dispute Loss object. This field will be present in the JSON response if
// and only if `category` is equal to `card_dispute_loss`. Contains the details of
// a lost Card Dispute.
type TransactionSourceCardDisputeLoss struct {
	// The identifier of the Card Dispute that was lost.
	CardDisputeID string `json:"card_dispute_id,required"`
	// Why the Card Dispute was lost.
	Explanation string `json:"explanation,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was lost.
	LostAt time.Time `json:"lost_at,required" format:"date-time"`
	// The identifier of the Transaction that was created to debit the disputed funds
	// from your account.
	TransactionID string                               `json:"transaction_id,required"`
	JSON          transactionSourceCardDisputeLossJSON `json:"-"`
}

// transactionSourceCardDisputeLossJSON contains the JSON metadata for the struct
// [TransactionSourceCardDisputeLoss]
type transactionSourceCardDisputeLossJSON struct {
	CardDisputeID apijson.Field
	Explanation   apijson.Field
	LostAt        apijson.Field
	TransactionID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TransactionSourceCardDisputeLoss) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardDisputeLossJSON) RawJSON() string {
	return r.raw
}

// A Card Push Transfer Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `card_push_transfer_acceptance`.
// A Card Push Transfer Acceptance is created when an Outbound Card Push Transfer
// sent from Increase is accepted by the receiving bank.
type TransactionSourceCardPushTransferAcceptance struct {
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the Card Push Transfer that led to this Transaction.
	TransferID string                                          `json:"transfer_id,required"`
	JSON       transactionSourceCardPushTransferAcceptanceJSON `json:"-"`
}

// transactionSourceCardPushTransferAcceptanceJSON contains the JSON metadata for
// the struct [TransactionSourceCardPushTransferAcceptance]
type transactionSourceCardPushTransferAcceptanceJSON struct {
	Amount      apijson.Field
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardPushTransferAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardPushTransferAcceptanceJSON) RawJSON() string {
	return r.raw
}

// A Card Refund object. This field will be present in the JSON response if and
// only if `category` is equal to `card_refund`. Card Refunds move money back to
// the cardholder. While they are usually connected to a Card Settlement an
// acquirer can also refund money directly to a card without relation to a
// transaction.
type TransactionSourceCardRefund struct {
	// The Card Refund identifier.
	ID string `json:"id,required"`
	// The amount in the minor unit of the transaction's settlement currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id,required"`
	// Cashback debited for this transaction, if eligible. Cashback is paid out in
	// aggregate, monthly.
	Cashback TransactionSourceCardRefundCashback `json:"cashback,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's settlement currency.
	Currency TransactionSourceCardRefundCurrency `json:"currency,required"`
	// Interchange assessed as a part of this transaciton.
	Interchange TransactionSourceCardRefundInterchange `json:"interchange,required,nullable"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required"`
	// The 4-digit MCC describing the merchant's business.
	MerchantCategoryCode string `json:"merchant_category_code,required"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required"`
	// The name of the merchant.
	MerchantName string `json:"merchant_name,required"`
	// The merchant's postal code. For US merchants this is always a 5-digit ZIP code.
	MerchantPostalCode string `json:"merchant_postal_code,required,nullable"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state,required,nullable"`
	// Network-specific identifiers for this refund.
	NetworkIdentifiers TransactionSourceCardRefundNetworkIdentifiers `json:"network_identifiers,required"`
	// The amount in the minor unit of the transaction's presentment currency.
	PresentmentAmount int64 `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's presentment currency.
	PresentmentCurrency string `json:"presentment_currency,required"`
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
	Cashback             apijson.Field
	Currency             apijson.Field
	Interchange          apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantName         apijson.Field
	MerchantPostalCode   apijson.Field
	MerchantState        apijson.Field
	NetworkIdentifiers   apijson.Field
	PresentmentAmount    apijson.Field
	PresentmentCurrency  apijson.Field
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

// Cashback debited for this transaction, if eligible. Cashback is paid out in
// aggregate, monthly.
type TransactionSourceCardRefundCashback struct {
	// The cashback amount given as a string containing a decimal number. The amount is
	// a positive number if it will be credited to you (e.g., settlements) and a
	// negative number if it will be debited (e.g., refunds).
	Amount string `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the cashback.
	Currency TransactionSourceCardRefundCashbackCurrency `json:"currency,required"`
	JSON     transactionSourceCardRefundCashbackJSON     `json:"-"`
}

// transactionSourceCardRefundCashbackJSON contains the JSON metadata for the
// struct [TransactionSourceCardRefundCashback]
type transactionSourceCardRefundCashbackJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardRefundCashback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardRefundCashbackJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the cashback.
type TransactionSourceCardRefundCashbackCurrency string

const (
	TransactionSourceCardRefundCashbackCurrencyCad TransactionSourceCardRefundCashbackCurrency = "CAD"
	TransactionSourceCardRefundCashbackCurrencyChf TransactionSourceCardRefundCashbackCurrency = "CHF"
	TransactionSourceCardRefundCashbackCurrencyEur TransactionSourceCardRefundCashbackCurrency = "EUR"
	TransactionSourceCardRefundCashbackCurrencyGbp TransactionSourceCardRefundCashbackCurrency = "GBP"
	TransactionSourceCardRefundCashbackCurrencyJpy TransactionSourceCardRefundCashbackCurrency = "JPY"
	TransactionSourceCardRefundCashbackCurrencyUsd TransactionSourceCardRefundCashbackCurrency = "USD"
)

func (r TransactionSourceCardRefundCashbackCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCardRefundCashbackCurrencyCad, TransactionSourceCardRefundCashbackCurrencyChf, TransactionSourceCardRefundCashbackCurrencyEur, TransactionSourceCardRefundCashbackCurrencyGbp, TransactionSourceCardRefundCashbackCurrencyJpy, TransactionSourceCardRefundCashbackCurrencyUsd:
		return true
	}
	return false
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's settlement currency.
type TransactionSourceCardRefundCurrency string

const (
	TransactionSourceCardRefundCurrencyCad TransactionSourceCardRefundCurrency = "CAD"
	TransactionSourceCardRefundCurrencyChf TransactionSourceCardRefundCurrency = "CHF"
	TransactionSourceCardRefundCurrencyEur TransactionSourceCardRefundCurrency = "EUR"
	TransactionSourceCardRefundCurrencyGbp TransactionSourceCardRefundCurrency = "GBP"
	TransactionSourceCardRefundCurrencyJpy TransactionSourceCardRefundCurrency = "JPY"
	TransactionSourceCardRefundCurrencyUsd TransactionSourceCardRefundCurrency = "USD"
)

func (r TransactionSourceCardRefundCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCardRefundCurrencyCad, TransactionSourceCardRefundCurrencyChf, TransactionSourceCardRefundCurrencyEur, TransactionSourceCardRefundCurrencyGbp, TransactionSourceCardRefundCurrencyJpy, TransactionSourceCardRefundCurrencyUsd:
		return true
	}
	return false
}

// Interchange assessed as a part of this transaciton.
type TransactionSourceCardRefundInterchange struct {
	// The interchange amount given as a string containing a decimal number in major
	// units (so e.g., "3.14" for $3.14). The amount is a positive number if it is
	// credited to Increase (e.g., settlements) and a negative number if it is debited
	// (e.g., refunds).
	Amount string `json:"amount,required"`
	// The card network specific interchange code.
	Code string `json:"code,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the interchange
	// reimbursement.
	Currency TransactionSourceCardRefundInterchangeCurrency `json:"currency,required"`
	JSON     transactionSourceCardRefundInterchangeJSON     `json:"-"`
}

// transactionSourceCardRefundInterchangeJSON contains the JSON metadata for the
// struct [TransactionSourceCardRefundInterchange]
type transactionSourceCardRefundInterchangeJSON struct {
	Amount      apijson.Field
	Code        apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardRefundInterchange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardRefundInterchangeJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the interchange
// reimbursement.
type TransactionSourceCardRefundInterchangeCurrency string

const (
	TransactionSourceCardRefundInterchangeCurrencyCad TransactionSourceCardRefundInterchangeCurrency = "CAD"
	TransactionSourceCardRefundInterchangeCurrencyChf TransactionSourceCardRefundInterchangeCurrency = "CHF"
	TransactionSourceCardRefundInterchangeCurrencyEur TransactionSourceCardRefundInterchangeCurrency = "EUR"
	TransactionSourceCardRefundInterchangeCurrencyGbp TransactionSourceCardRefundInterchangeCurrency = "GBP"
	TransactionSourceCardRefundInterchangeCurrencyJpy TransactionSourceCardRefundInterchangeCurrency = "JPY"
	TransactionSourceCardRefundInterchangeCurrencyUsd TransactionSourceCardRefundInterchangeCurrency = "USD"
)

func (r TransactionSourceCardRefundInterchangeCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCardRefundInterchangeCurrencyCad, TransactionSourceCardRefundInterchangeCurrencyChf, TransactionSourceCardRefundInterchangeCurrencyEur, TransactionSourceCardRefundInterchangeCurrencyGbp, TransactionSourceCardRefundInterchangeCurrencyJpy, TransactionSourceCardRefundInterchangeCurrencyUsd:
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
	TransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesNoExtraCharge    TransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "no_extra_charge"
	TransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesGas              TransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "gas"
	TransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesExtraMileage     TransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "extra_mileage"
	TransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesLateReturn       TransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "late_return"
	TransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesOneWayServiceFee TransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "one_way_service_fee"
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
	TransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicatorNotApplicable               TransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicator = "not_applicable"
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
	TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesNoExtraCharge TransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "no_extra_charge"
	TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesRestaurant    TransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "restaurant"
	TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesGiftShop      TransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "gift_shop"
	TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesMiniBar       TransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "mini_bar"
	TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesTelephone     TransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "telephone"
	TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesOther         TransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "other"
	TransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesLaundry       TransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "laundry"
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
	TransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicatorNotApplicable TransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicator = "not_applicable"
	TransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicatorNoShow        TransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicator = "no_show"
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
	TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatFreeText              TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "free_text"
	TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatOrderNumber           TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "order_number"
	TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatRentalAgreementNumber TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "rental_agreement_number"
	TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatHotelFolioNumber      TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "hotel_folio_number"
	TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatInvoiceNumber         TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "invoice_number"
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
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorNoCredit                                                        TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "no_credit"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation                 TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorOther                                                           TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "other"
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
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryNone                  TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "none"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryBundledService        TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "bundled_service"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryBaggageFee            TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "baggage_fee"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryChangeFee             TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "change_fee"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryCargo                 TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "cargo"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryCarbonOffset          TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "carbon_offset"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryFrequentFlyer         TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "frequent_flyer"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryGiftCard              TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "gift_card"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryGroundTransport       TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "ground_transport"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryInFlightEntertainment TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "in_flight_entertainment"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryLounge                TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "lounge"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryMedical               TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "medical"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryMealBeverage          TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "meal_beverage"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryOther                 TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "other"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryPassengerAssistFee    TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "passenger_assist_fee"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryPets                  TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "pets"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategorySeatFees              TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "seat_fees"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryStandby               TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "standby"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryServiceFee            TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "service_fee"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryStore                 TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "store"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryTravelService         TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "travel_service"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryUnaccompaniedTravel   TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "unaccompanied_travel"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryUpgrades              TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "upgrades"
	TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryWifi                  TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "wifi"
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
	TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorNoCredit                                                        TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "no_credit"
	TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation                 TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketCancellation                                       TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_cancellation"
	TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorOther                                                           TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "other"
	TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorPartialRefundOfAirlineTicket                                    TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "partial_refund_of_airline_ticket"
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
	TransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicatorNoRestrictions                TransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicator = "no_restrictions"
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
	TransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicatorNone                   TransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator = "none"
	TransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicatorChangeToExistingTicket TransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator = "change_to_existing_ticket"
	TransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicatorNewTicket              TransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator = "new_ticket"
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
	TransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCodeNone               TransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode = "none"
	TransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCodeStopOverAllowed    TransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_allowed"
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
// if and only if `category` is equal to `card_revenue_payment`. Card Revenue
// Payments reflect earnings from fees on card transactions.
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
	TransactionSourceCardRevenuePaymentCurrencyCad TransactionSourceCardRevenuePaymentCurrency = "CAD"
	TransactionSourceCardRevenuePaymentCurrencyChf TransactionSourceCardRevenuePaymentCurrency = "CHF"
	TransactionSourceCardRevenuePaymentCurrencyEur TransactionSourceCardRevenuePaymentCurrency = "EUR"
	TransactionSourceCardRevenuePaymentCurrencyGbp TransactionSourceCardRevenuePaymentCurrency = "GBP"
	TransactionSourceCardRevenuePaymentCurrencyJpy TransactionSourceCardRevenuePaymentCurrency = "JPY"
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
// only if `category` is equal to `card_settlement`. Card Settlements are card
// transactions that have cleared and settled. While a settlement is usually
// preceded by an authorization, an acquirer can also directly clear a transaction
// without first authorizing it.
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
	CardPaymentID string `json:"card_payment_id,required"`
	// Cashback earned on this transaction, if eligible. Cashback is paid out in
	// aggregate, monthly.
	Cashback TransactionSourceCardSettlementCashback `json:"cashback,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's settlement currency.
	Currency TransactionSourceCardSettlementCurrency `json:"currency,required"`
	// Interchange assessed as a part of this transaction.
	Interchange TransactionSourceCardSettlementInterchange `json:"interchange,required,nullable"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required"`
	// The 4-digit MCC describing the merchant's business.
	MerchantCategoryCode string `json:"merchant_category_code,required"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required"`
	// The name of the merchant.
	MerchantName string `json:"merchant_name,required"`
	// The merchant's postal code. For US merchants this is always a 5-digit ZIP code.
	MerchantPostalCode string `json:"merchant_postal_code,required,nullable"`
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
	Cashback             apijson.Field
	Currency             apijson.Field
	Interchange          apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantName         apijson.Field
	MerchantPostalCode   apijson.Field
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

// Cashback earned on this transaction, if eligible. Cashback is paid out in
// aggregate, monthly.
type TransactionSourceCardSettlementCashback struct {
	// The cashback amount given as a string containing a decimal number. The amount is
	// a positive number if it will be credited to you (e.g., settlements) and a
	// negative number if it will be debited (e.g., refunds).
	Amount string `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the cashback.
	Currency TransactionSourceCardSettlementCashbackCurrency `json:"currency,required"`
	JSON     transactionSourceCardSettlementCashbackJSON     `json:"-"`
}

// transactionSourceCardSettlementCashbackJSON contains the JSON metadata for the
// struct [TransactionSourceCardSettlementCashback]
type transactionSourceCardSettlementCashbackJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardSettlementCashback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardSettlementCashbackJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the cashback.
type TransactionSourceCardSettlementCashbackCurrency string

const (
	TransactionSourceCardSettlementCashbackCurrencyCad TransactionSourceCardSettlementCashbackCurrency = "CAD"
	TransactionSourceCardSettlementCashbackCurrencyChf TransactionSourceCardSettlementCashbackCurrency = "CHF"
	TransactionSourceCardSettlementCashbackCurrencyEur TransactionSourceCardSettlementCashbackCurrency = "EUR"
	TransactionSourceCardSettlementCashbackCurrencyGbp TransactionSourceCardSettlementCashbackCurrency = "GBP"
	TransactionSourceCardSettlementCashbackCurrencyJpy TransactionSourceCardSettlementCashbackCurrency = "JPY"
	TransactionSourceCardSettlementCashbackCurrencyUsd TransactionSourceCardSettlementCashbackCurrency = "USD"
)

func (r TransactionSourceCardSettlementCashbackCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCardSettlementCashbackCurrencyCad, TransactionSourceCardSettlementCashbackCurrencyChf, TransactionSourceCardSettlementCashbackCurrencyEur, TransactionSourceCardSettlementCashbackCurrencyGbp, TransactionSourceCardSettlementCashbackCurrencyJpy, TransactionSourceCardSettlementCashbackCurrencyUsd:
		return true
	}
	return false
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's settlement currency.
type TransactionSourceCardSettlementCurrency string

const (
	TransactionSourceCardSettlementCurrencyCad TransactionSourceCardSettlementCurrency = "CAD"
	TransactionSourceCardSettlementCurrencyChf TransactionSourceCardSettlementCurrency = "CHF"
	TransactionSourceCardSettlementCurrencyEur TransactionSourceCardSettlementCurrency = "EUR"
	TransactionSourceCardSettlementCurrencyGbp TransactionSourceCardSettlementCurrency = "GBP"
	TransactionSourceCardSettlementCurrencyJpy TransactionSourceCardSettlementCurrency = "JPY"
	TransactionSourceCardSettlementCurrencyUsd TransactionSourceCardSettlementCurrency = "USD"
)

func (r TransactionSourceCardSettlementCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCardSettlementCurrencyCad, TransactionSourceCardSettlementCurrencyChf, TransactionSourceCardSettlementCurrencyEur, TransactionSourceCardSettlementCurrencyGbp, TransactionSourceCardSettlementCurrencyJpy, TransactionSourceCardSettlementCurrencyUsd:
		return true
	}
	return false
}

// Interchange assessed as a part of this transaction.
type TransactionSourceCardSettlementInterchange struct {
	// The interchange amount given as a string containing a decimal number in major
	// units (so e.g., "3.14" for $3.14). The amount is a positive number if it is
	// credited to Increase (e.g., settlements) and a negative number if it is debited
	// (e.g., refunds).
	Amount string `json:"amount,required"`
	// The card network specific interchange code.
	Code string `json:"code,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the interchange
	// reimbursement.
	Currency TransactionSourceCardSettlementInterchangeCurrency `json:"currency,required"`
	JSON     transactionSourceCardSettlementInterchangeJSON     `json:"-"`
}

// transactionSourceCardSettlementInterchangeJSON contains the JSON metadata for
// the struct [TransactionSourceCardSettlementInterchange]
type transactionSourceCardSettlementInterchangeJSON struct {
	Amount      apijson.Field
	Code        apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardSettlementInterchange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardSettlementInterchangeJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the interchange
// reimbursement.
type TransactionSourceCardSettlementInterchangeCurrency string

const (
	TransactionSourceCardSettlementInterchangeCurrencyCad TransactionSourceCardSettlementInterchangeCurrency = "CAD"
	TransactionSourceCardSettlementInterchangeCurrencyChf TransactionSourceCardSettlementInterchangeCurrency = "CHF"
	TransactionSourceCardSettlementInterchangeCurrencyEur TransactionSourceCardSettlementInterchangeCurrency = "EUR"
	TransactionSourceCardSettlementInterchangeCurrencyGbp TransactionSourceCardSettlementInterchangeCurrency = "GBP"
	TransactionSourceCardSettlementInterchangeCurrencyJpy TransactionSourceCardSettlementInterchangeCurrency = "JPY"
	TransactionSourceCardSettlementInterchangeCurrencyUsd TransactionSourceCardSettlementInterchangeCurrency = "USD"
)

func (r TransactionSourceCardSettlementInterchangeCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCardSettlementInterchangeCurrencyCad, TransactionSourceCardSettlementInterchangeCurrencyChf, TransactionSourceCardSettlementInterchangeCurrencyEur, TransactionSourceCardSettlementInterchangeCurrencyGbp, TransactionSourceCardSettlementInterchangeCurrencyJpy, TransactionSourceCardSettlementInterchangeCurrencyUsd:
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
	TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesNoExtraCharge    TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "no_extra_charge"
	TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesGas              TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "gas"
	TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesExtraMileage     TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "extra_mileage"
	TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesLateReturn       TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "late_return"
	TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesOneWayServiceFee TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "one_way_service_fee"
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
	TransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicatorNotApplicable               TransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicator = "not_applicable"
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
	TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesNoExtraCharge TransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "no_extra_charge"
	TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesRestaurant    TransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "restaurant"
	TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesGiftShop      TransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "gift_shop"
	TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesMiniBar       TransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "mini_bar"
	TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesTelephone     TransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "telephone"
	TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesOther         TransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "other"
	TransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesLaundry       TransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "laundry"
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
	TransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicatorNotApplicable TransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicator = "not_applicable"
	TransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicatorNoShow        TransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicator = "no_show"
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
	TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatFreeText              TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "free_text"
	TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatOrderNumber           TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "order_number"
	TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatRentalAgreementNumber TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "rental_agreement_number"
	TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatHotelFolioNumber      TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "hotel_folio_number"
	TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatInvoiceNumber         TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "invoice_number"
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
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorNoCredit                                                        TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "no_credit"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation                 TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorOther                                                           TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "other"
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
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryNone                  TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "none"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryBundledService        TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "bundled_service"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryBaggageFee            TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "baggage_fee"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryChangeFee             TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "change_fee"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryCargo                 TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "cargo"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryCarbonOffset          TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "carbon_offset"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryFrequentFlyer         TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "frequent_flyer"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryGiftCard              TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "gift_card"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryGroundTransport       TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "ground_transport"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryInFlightEntertainment TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "in_flight_entertainment"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryLounge                TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "lounge"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryMedical               TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "medical"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryMealBeverage          TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "meal_beverage"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryOther                 TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "other"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryPassengerAssistFee    TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "passenger_assist_fee"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryPets                  TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "pets"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategorySeatFees              TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "seat_fees"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryStandby               TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "standby"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryServiceFee            TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "service_fee"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryStore                 TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "store"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryTravelService         TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "travel_service"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryUnaccompaniedTravel   TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "unaccompanied_travel"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryUpgrades              TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "upgrades"
	TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryWifi                  TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "wifi"
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
	TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorNoCredit                                                        TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "no_credit"
	TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation                 TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketCancellation                                       TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_cancellation"
	TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorOther                                                           TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "other"
	TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorPartialRefundOfAirlineTicket                                    TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "partial_refund_of_airline_ticket"
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
	TransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicatorNoRestrictions                TransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator = "no_restrictions"
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
	TransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicatorNone                   TransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "none"
	TransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicatorChangeToExistingTicket TransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "change_to_existing_ticket"
	TransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicatorNewTicket              TransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "new_ticket"
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
	TransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeNone               TransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode = "none"
	TransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeStopOverAllowed    TransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_allowed"
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

// A Cashback Payment object. This field will be present in the JSON response if
// and only if `category` is equal to `cashback_payment`. A Cashback Payment
// represents the cashback paid to a cardholder for a given period. Cashback is
// usually paid monthly for the prior month's transactions.
type TransactionSourceCashbackPayment struct {
	// The card on which the cashback was accrued.
	AccruedOnCardID string `json:"accrued_on_card_id,required,nullable"`
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency TransactionSourceCashbackPaymentCurrency `json:"currency,required"`
	// The end of the period for which this transaction paid cashback.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The start of the period for which this transaction paid cashback.
	PeriodStart time.Time                            `json:"period_start,required" format:"date-time"`
	JSON        transactionSourceCashbackPaymentJSON `json:"-"`
}

// transactionSourceCashbackPaymentJSON contains the JSON metadata for the struct
// [TransactionSourceCashbackPayment]
type transactionSourceCashbackPaymentJSON struct {
	AccruedOnCardID apijson.Field
	Amount          apijson.Field
	Currency        apijson.Field
	PeriodEnd       apijson.Field
	PeriodStart     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TransactionSourceCashbackPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCashbackPaymentJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type TransactionSourceCashbackPaymentCurrency string

const (
	TransactionSourceCashbackPaymentCurrencyCad TransactionSourceCashbackPaymentCurrency = "CAD"
	TransactionSourceCashbackPaymentCurrencyChf TransactionSourceCashbackPaymentCurrency = "CHF"
	TransactionSourceCashbackPaymentCurrencyEur TransactionSourceCashbackPaymentCurrency = "EUR"
	TransactionSourceCashbackPaymentCurrencyGbp TransactionSourceCashbackPaymentCurrency = "GBP"
	TransactionSourceCashbackPaymentCurrencyJpy TransactionSourceCashbackPaymentCurrency = "JPY"
	TransactionSourceCashbackPaymentCurrencyUsd TransactionSourceCashbackPaymentCurrency = "USD"
)

func (r TransactionSourceCashbackPaymentCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCashbackPaymentCurrencyCad, TransactionSourceCashbackPaymentCurrencyChf, TransactionSourceCashbackPaymentCurrencyEur, TransactionSourceCashbackPaymentCurrencyGbp, TransactionSourceCashbackPaymentCurrencyJpy, TransactionSourceCashbackPaymentCurrencyUsd:
		return true
	}
	return false
}

// The type of the resource. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
type TransactionSourceCategory string

const (
	TransactionSourceCategoryAccountTransferIntention                    TransactionSourceCategory = "account_transfer_intention"
	TransactionSourceCategoryACHTransferIntention                        TransactionSourceCategory = "ach_transfer_intention"
	TransactionSourceCategoryACHTransferRejection                        TransactionSourceCategory = "ach_transfer_rejection"
	TransactionSourceCategoryACHTransferReturn                           TransactionSourceCategory = "ach_transfer_return"
	TransactionSourceCategoryCashbackPayment                             TransactionSourceCategory = "cashback_payment"
	TransactionSourceCategoryCardDisputeAcceptance                       TransactionSourceCategory = "card_dispute_acceptance"
	TransactionSourceCategoryCardDisputeFinancial                        TransactionSourceCategory = "card_dispute_financial"
	TransactionSourceCategoryCardDisputeLoss                             TransactionSourceCategory = "card_dispute_loss"
	TransactionSourceCategoryCardRefund                                  TransactionSourceCategory = "card_refund"
	TransactionSourceCategoryCardSettlement                              TransactionSourceCategory = "card_settlement"
	TransactionSourceCategoryCardRevenuePayment                          TransactionSourceCategory = "card_revenue_payment"
	TransactionSourceCategoryCheckDepositAcceptance                      TransactionSourceCategory = "check_deposit_acceptance"
	TransactionSourceCategoryCheckDepositReturn                          TransactionSourceCategory = "check_deposit_return"
	TransactionSourceCategoryCheckTransferDeposit                        TransactionSourceCategory = "check_transfer_deposit"
	TransactionSourceCategoryFeePayment                                  TransactionSourceCategory = "fee_payment"
	TransactionSourceCategoryInboundACHTransfer                          TransactionSourceCategory = "inbound_ach_transfer"
	TransactionSourceCategoryInboundACHTransferReturnIntention           TransactionSourceCategory = "inbound_ach_transfer_return_intention"
	TransactionSourceCategoryInboundCheckDepositReturnIntention          TransactionSourceCategory = "inbound_check_deposit_return_intention"
	TransactionSourceCategoryInboundCheckAdjustment                      TransactionSourceCategory = "inbound_check_adjustment"
	TransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation TransactionSourceCategory = "inbound_real_time_payments_transfer_confirmation"
	TransactionSourceCategoryInboundRealTimePaymentsTransferDecline      TransactionSourceCategory = "inbound_real_time_payments_transfer_decline"
	TransactionSourceCategoryInboundWireReversal                         TransactionSourceCategory = "inbound_wire_reversal"
	TransactionSourceCategoryInboundWireTransfer                         TransactionSourceCategory = "inbound_wire_transfer"
	TransactionSourceCategoryInboundWireTransferReversal                 TransactionSourceCategory = "inbound_wire_transfer_reversal"
	TransactionSourceCategoryInterestPayment                             TransactionSourceCategory = "interest_payment"
	TransactionSourceCategoryInternalSource                              TransactionSourceCategory = "internal_source"
	TransactionSourceCategoryRealTimePaymentsTransferAcknowledgement     TransactionSourceCategory = "real_time_payments_transfer_acknowledgement"
	TransactionSourceCategorySampleFunds                                 TransactionSourceCategory = "sample_funds"
	TransactionSourceCategoryWireTransferIntention                       TransactionSourceCategory = "wire_transfer_intention"
	TransactionSourceCategorySwiftTransferIntention                      TransactionSourceCategory = "swift_transfer_intention"
	TransactionSourceCategoryCardPushTransferAcceptance                  TransactionSourceCategory = "card_push_transfer_acceptance"
	TransactionSourceCategoryOther                                       TransactionSourceCategory = "other"
)

func (r TransactionSourceCategory) IsKnown() bool {
	switch r {
	case TransactionSourceCategoryAccountTransferIntention, TransactionSourceCategoryACHTransferIntention, TransactionSourceCategoryACHTransferRejection, TransactionSourceCategoryACHTransferReturn, TransactionSourceCategoryCashbackPayment, TransactionSourceCategoryCardDisputeAcceptance, TransactionSourceCategoryCardDisputeFinancial, TransactionSourceCategoryCardDisputeLoss, TransactionSourceCategoryCardRefund, TransactionSourceCategoryCardSettlement, TransactionSourceCategoryCardRevenuePayment, TransactionSourceCategoryCheckDepositAcceptance, TransactionSourceCategoryCheckDepositReturn, TransactionSourceCategoryCheckTransferDeposit, TransactionSourceCategoryFeePayment, TransactionSourceCategoryInboundACHTransfer, TransactionSourceCategoryInboundACHTransferReturnIntention, TransactionSourceCategoryInboundCheckDepositReturnIntention, TransactionSourceCategoryInboundCheckAdjustment, TransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation, TransactionSourceCategoryInboundRealTimePaymentsTransferDecline, TransactionSourceCategoryInboundWireReversal, TransactionSourceCategoryInboundWireTransfer, TransactionSourceCategoryInboundWireTransferReversal, TransactionSourceCategoryInterestPayment, TransactionSourceCategoryInternalSource, TransactionSourceCategoryRealTimePaymentsTransferAcknowledgement, TransactionSourceCategorySampleFunds, TransactionSourceCategoryWireTransferIntention, TransactionSourceCategorySwiftTransferIntention, TransactionSourceCategoryCardPushTransferAcceptance, TransactionSourceCategoryOther:
		return true
	}
	return false
}

// A Check Deposit Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_acceptance`. A
// Check Deposit Acceptance is created when a Check Deposit is processed and its
// details confirmed. Check Deposits may be returned by the receiving bank, which
// will appear as a Check Deposit Return.
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
	TransactionSourceCheckDepositAcceptanceCurrencyCad TransactionSourceCheckDepositAcceptanceCurrency = "CAD"
	TransactionSourceCheckDepositAcceptanceCurrencyChf TransactionSourceCheckDepositAcceptanceCurrency = "CHF"
	TransactionSourceCheckDepositAcceptanceCurrencyEur TransactionSourceCheckDepositAcceptanceCurrency = "EUR"
	TransactionSourceCheckDepositAcceptanceCurrencyGbp TransactionSourceCheckDepositAcceptanceCurrency = "GBP"
	TransactionSourceCheckDepositAcceptanceCurrencyJpy TransactionSourceCheckDepositAcceptanceCurrency = "JPY"
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
// if and only if `category` is equal to `check_deposit_return`. A Check Deposit
// Return is created when a Check Deposit is returned by the bank holding the
// account it was drawn against. Check Deposits may be returned for a variety of
// reasons, including insufficient funds or a mismatched account number. Usually,
// checks are returned within the first 7 days after the deposit is made.
type TransactionSourceCheckDepositReturn struct {
	// The returned amount in USD cents.
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
	TransactionSourceCheckDepositReturnCurrencyCad TransactionSourceCheckDepositReturnCurrency = "CAD"
	TransactionSourceCheckDepositReturnCurrencyChf TransactionSourceCheckDepositReturnCurrency = "CHF"
	TransactionSourceCheckDepositReturnCurrencyEur TransactionSourceCheckDepositReturnCurrency = "EUR"
	TransactionSourceCheckDepositReturnCurrencyGbp TransactionSourceCheckDepositReturnCurrency = "GBP"
	TransactionSourceCheckDepositReturnCurrencyJpy TransactionSourceCheckDepositReturnCurrency = "JPY"
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
	TransactionSourceCheckDepositReturnReturnReasonACHConversionNotSupported TransactionSourceCheckDepositReturnReturnReason = "ach_conversion_not_supported"
	TransactionSourceCheckDepositReturnReturnReasonClosedAccount             TransactionSourceCheckDepositReturnReturnReason = "closed_account"
	TransactionSourceCheckDepositReturnReturnReasonDuplicateSubmission       TransactionSourceCheckDepositReturnReturnReason = "duplicate_submission"
	TransactionSourceCheckDepositReturnReturnReasonInsufficientFunds         TransactionSourceCheckDepositReturnReturnReason = "insufficient_funds"
	TransactionSourceCheckDepositReturnReturnReasonNoAccount                 TransactionSourceCheckDepositReturnReturnReason = "no_account"
	TransactionSourceCheckDepositReturnReturnReasonNotAuthorized             TransactionSourceCheckDepositReturnReturnReason = "not_authorized"
	TransactionSourceCheckDepositReturnReturnReasonStaleDated                TransactionSourceCheckDepositReturnReturnReason = "stale_dated"
	TransactionSourceCheckDepositReturnReturnReasonStopPayment               TransactionSourceCheckDepositReturnReturnReason = "stop_payment"
	TransactionSourceCheckDepositReturnReturnReasonUnknownReason             TransactionSourceCheckDepositReturnReturnReason = "unknown_reason"
	TransactionSourceCheckDepositReturnReturnReasonUnmatchedDetails          TransactionSourceCheckDepositReturnReturnReason = "unmatched_details"
	TransactionSourceCheckDepositReturnReturnReasonUnreadableImage           TransactionSourceCheckDepositReturnReturnReason = "unreadable_image"
	TransactionSourceCheckDepositReturnReturnReasonEndorsementIrregular      TransactionSourceCheckDepositReturnReturnReason = "endorsement_irregular"
	TransactionSourceCheckDepositReturnReturnReasonAlteredOrFictitiousItem   TransactionSourceCheckDepositReturnReturnReason = "altered_or_fictitious_item"
	TransactionSourceCheckDepositReturnReturnReasonFrozenOrBlockedAccount    TransactionSourceCheckDepositReturnReturnReason = "frozen_or_blocked_account"
	TransactionSourceCheckDepositReturnReturnReasonPostDated                 TransactionSourceCheckDepositReturnReturnReason = "post_dated"
	TransactionSourceCheckDepositReturnReturnReasonEndorsementMissing        TransactionSourceCheckDepositReturnReturnReason = "endorsement_missing"
	TransactionSourceCheckDepositReturnReturnReasonSignatureMissing          TransactionSourceCheckDepositReturnReturnReason = "signature_missing"
	TransactionSourceCheckDepositReturnReturnReasonStopPaymentSuspect        TransactionSourceCheckDepositReturnReturnReason = "stop_payment_suspect"
	TransactionSourceCheckDepositReturnReturnReasonUnusableImage             TransactionSourceCheckDepositReturnReturnReason = "unusable_image"
	TransactionSourceCheckDepositReturnReturnReasonImageFailsSecurityCheck   TransactionSourceCheckDepositReturnReturnReason = "image_fails_security_check"
	TransactionSourceCheckDepositReturnReturnReasonCannotDetermineAmount     TransactionSourceCheckDepositReturnReturnReason = "cannot_determine_amount"
	TransactionSourceCheckDepositReturnReturnReasonSignatureIrregular        TransactionSourceCheckDepositReturnReturnReason = "signature_irregular"
	TransactionSourceCheckDepositReturnReturnReasonNonCashItem               TransactionSourceCheckDepositReturnReturnReason = "non_cash_item"
	TransactionSourceCheckDepositReturnReturnReasonUnableToProcess           TransactionSourceCheckDepositReturnReturnReason = "unable_to_process"
	TransactionSourceCheckDepositReturnReturnReasonItemExceedsDollarLimit    TransactionSourceCheckDepositReturnReturnReason = "item_exceeds_dollar_limit"
	TransactionSourceCheckDepositReturnReturnReasonBranchOrAccountSold       TransactionSourceCheckDepositReturnReturnReason = "branch_or_account_sold"
)

func (r TransactionSourceCheckDepositReturnReturnReason) IsKnown() bool {
	switch r {
	case TransactionSourceCheckDepositReturnReturnReasonACHConversionNotSupported, TransactionSourceCheckDepositReturnReturnReasonClosedAccount, TransactionSourceCheckDepositReturnReturnReasonDuplicateSubmission, TransactionSourceCheckDepositReturnReturnReasonInsufficientFunds, TransactionSourceCheckDepositReturnReturnReasonNoAccount, TransactionSourceCheckDepositReturnReturnReasonNotAuthorized, TransactionSourceCheckDepositReturnReturnReasonStaleDated, TransactionSourceCheckDepositReturnReturnReasonStopPayment, TransactionSourceCheckDepositReturnReturnReasonUnknownReason, TransactionSourceCheckDepositReturnReturnReasonUnmatchedDetails, TransactionSourceCheckDepositReturnReturnReasonUnreadableImage, TransactionSourceCheckDepositReturnReturnReasonEndorsementIrregular, TransactionSourceCheckDepositReturnReturnReasonAlteredOrFictitiousItem, TransactionSourceCheckDepositReturnReturnReasonFrozenOrBlockedAccount, TransactionSourceCheckDepositReturnReturnReasonPostDated, TransactionSourceCheckDepositReturnReturnReasonEndorsementMissing, TransactionSourceCheckDepositReturnReturnReasonSignatureMissing, TransactionSourceCheckDepositReturnReturnReasonStopPaymentSuspect, TransactionSourceCheckDepositReturnReturnReasonUnusableImage, TransactionSourceCheckDepositReturnReturnReasonImageFailsSecurityCheck, TransactionSourceCheckDepositReturnReturnReasonCannotDetermineAmount, TransactionSourceCheckDepositReturnReturnReasonSignatureIrregular, TransactionSourceCheckDepositReturnReturnReasonNonCashItem, TransactionSourceCheckDepositReturnReturnReasonUnableToProcess, TransactionSourceCheckDepositReturnReturnReasonItemExceedsDollarLimit, TransactionSourceCheckDepositReturnReturnReasonBranchOrAccountSold:
		return true
	}
	return false
}

// A Check Transfer Deposit object. This field will be present in the JSON response
// if and only if `category` is equal to `check_transfer_deposit`. An Inbound Check
// is a check drawn on an Increase account that has been deposited by an external
// bank account. These types of checks are not pre-registered.
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
	// The identifier of the Inbound Check Deposit object associated with this
	// transaction.
	InboundCheckDepositID string `json:"inbound_check_deposit_id,required,nullable"`
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
	InboundCheckDepositID           apijson.Field
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

// A Fee Payment object. This field will be present in the JSON response if and
// only if `category` is equal to `fee_payment`. A Fee Payment represents a payment
// made to Increase.
type TransactionSourceFeePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency TransactionSourceFeePaymentCurrency `json:"currency,required"`
	// The start of this payment's fee period, usually the first day of a month.
	FeePeriodStart time.Time `json:"fee_period_start,required" format:"date"`
	// The Program for which this fee was incurred.
	ProgramID string                          `json:"program_id,required,nullable"`
	JSON      transactionSourceFeePaymentJSON `json:"-"`
}

// transactionSourceFeePaymentJSON contains the JSON metadata for the struct
// [TransactionSourceFeePayment]
type transactionSourceFeePaymentJSON struct {
	Amount         apijson.Field
	Currency       apijson.Field
	FeePeriodStart apijson.Field
	ProgramID      apijson.Field
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
	TransactionSourceFeePaymentCurrencyCad TransactionSourceFeePaymentCurrency = "CAD"
	TransactionSourceFeePaymentCurrencyChf TransactionSourceFeePaymentCurrency = "CHF"
	TransactionSourceFeePaymentCurrencyEur TransactionSourceFeePaymentCurrency = "EUR"
	TransactionSourceFeePaymentCurrencyGbp TransactionSourceFeePaymentCurrency = "GBP"
	TransactionSourceFeePaymentCurrencyJpy TransactionSourceFeePaymentCurrency = "JPY"
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
// response if and only if `category` is equal to `inbound_ach_transfer`. An
// Inbound ACH Transfer Intention is created when an ACH transfer is initiated at
// another bank and received by Increase.
type TransactionSourceInboundACHTransfer struct {
	// Additional information sent from the originator.
	Addenda TransactionSourceInboundACHTransferAddenda `json:"addenda,required,nullable"`
	// The transfer amount in USD cents.
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
	// The originator's identifier for the transfer recipient.
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
	// The Inbound ACH Transfer's identifier.
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

// An Inbound ACH Transfer Return Intention object. This field will be present in
// the JSON response if and only if `category` is equal to
// `inbound_ach_transfer_return_intention`. An Inbound ACH Transfer Return
// Intention is created when an ACH transfer is initiated at another bank and
// returned by Increase.
type TransactionSourceInboundACHTransferReturnIntention struct {
	// The ID of the Inbound ACH Transfer that is being returned.
	InboundACHTransferID string                                                 `json:"inbound_ach_transfer_id,required"`
	JSON                 transactionSourceInboundACHTransferReturnIntentionJSON `json:"-"`
}

// transactionSourceInboundACHTransferReturnIntentionJSON contains the JSON
// metadata for the struct [TransactionSourceInboundACHTransferReturnIntention]
type transactionSourceInboundACHTransferReturnIntentionJSON struct {
	InboundACHTransferID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TransactionSourceInboundACHTransferReturnIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceInboundACHTransferReturnIntentionJSON) RawJSON() string {
	return r.raw
}

// An Inbound Check Adjustment object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_check_adjustment`. An
// Inbound Check Adjustment is created when Increase receives an adjustment for a
// check or return deposited through Check21.
type TransactionSourceInboundCheckAdjustment struct {
	// The ID of the transaction that was adjusted.
	AdjustedTransactionID string `json:"adjusted_transaction_id,required"`
	// The amount of the check adjustment.
	Amount int64 `json:"amount,required"`
	// The reason for the adjustment.
	Reason TransactionSourceInboundCheckAdjustmentReason `json:"reason,required"`
	JSON   transactionSourceInboundCheckAdjustmentJSON   `json:"-"`
}

// transactionSourceInboundCheckAdjustmentJSON contains the JSON metadata for the
// struct [TransactionSourceInboundCheckAdjustment]
type transactionSourceInboundCheckAdjustmentJSON struct {
	AdjustedTransactionID apijson.Field
	Amount                apijson.Field
	Reason                apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *TransactionSourceInboundCheckAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceInboundCheckAdjustmentJSON) RawJSON() string {
	return r.raw
}

// The reason for the adjustment.
type TransactionSourceInboundCheckAdjustmentReason string

const (
	TransactionSourceInboundCheckAdjustmentReasonLateReturn        TransactionSourceInboundCheckAdjustmentReason = "late_return"
	TransactionSourceInboundCheckAdjustmentReasonWrongPayeeCredit  TransactionSourceInboundCheckAdjustmentReason = "wrong_payee_credit"
	TransactionSourceInboundCheckAdjustmentReasonAdjustedAmount    TransactionSourceInboundCheckAdjustmentReason = "adjusted_amount"
	TransactionSourceInboundCheckAdjustmentReasonNonConformingItem TransactionSourceInboundCheckAdjustmentReason = "non_conforming_item"
)

func (r TransactionSourceInboundCheckAdjustmentReason) IsKnown() bool {
	switch r {
	case TransactionSourceInboundCheckAdjustmentReasonLateReturn, TransactionSourceInboundCheckAdjustmentReasonWrongPayeeCredit, TransactionSourceInboundCheckAdjustmentReasonAdjustedAmount, TransactionSourceInboundCheckAdjustmentReasonNonConformingItem:
		return true
	}
	return false
}

// An Inbound Check Deposit Return Intention object. This field will be present in
// the JSON response if and only if `category` is equal to
// `inbound_check_deposit_return_intention`. An Inbound Check Deposit Return
// Intention is created when Increase receives an Inbound Check and the User
// requests that it be returned.
type TransactionSourceInboundCheckDepositReturnIntention struct {
	// The ID of the Inbound Check Deposit that is being returned.
	InboundCheckDepositID string `json:"inbound_check_deposit_id,required"`
	// The identifier of the Check Transfer object that was deposited.
	TransferID string                                                  `json:"transfer_id,required,nullable"`
	JSON       transactionSourceInboundCheckDepositReturnIntentionJSON `json:"-"`
}

// transactionSourceInboundCheckDepositReturnIntentionJSON contains the JSON
// metadata for the struct [TransactionSourceInboundCheckDepositReturnIntention]
type transactionSourceInboundCheckDepositReturnIntentionJSON struct {
	InboundCheckDepositID apijson.Field
	TransferID            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *TransactionSourceInboundCheckDepositReturnIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceInboundCheckDepositReturnIntentionJSON) RawJSON() string {
	return r.raw
}

// An Inbound Real-Time Payments Transfer Confirmation object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_confirmation`. An Inbound Real-Time
// Payments Transfer Confirmation is created when a Real-Time Payments transfer is
// initiated at another bank and received by Increase.
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
	TransactionIdentification string `json:"transaction_identification,required"`
	// The identifier of the Real-Time Payments Transfer that led to this Transaction.
	TransferID string                                                           `json:"transfer_id,required"`
	JSON       transactionSourceInboundRealTimePaymentsTransferConfirmationJSON `json:"-"`
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
	TransferID                apijson.Field
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
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyCad TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CAD"
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyChf TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CHF"
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyEur TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "EUR"
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyGbp TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "GBP"
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyJpy TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "JPY"
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyUsd TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "USD"
)

func (r TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyCad, TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyChf, TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyEur, TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyGbp, TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyJpy, TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyUsd:
		return true
	}
	return false
}

// An Inbound Real-Time Payments Transfer Decline object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_decline`.
type TransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real-Time Payments
	// transfer.
	Currency TransactionSourceInboundRealTimePaymentsTransferDeclineCurrency `json:"currency,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber string `json:"debtor_account_number,required"`
	// The name provided by the sender of the transfer.
	DebtorName string `json:"debtor_name,required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber string `json:"debtor_routing_number,required"`
	// Why the transfer was declined.
	Reason TransactionSourceInboundRealTimePaymentsTransferDeclineReason `json:"reason,required"`
	// Additional information included with the transfer.
	RemittanceInformation string `json:"remittance_information,required,nullable"`
	// The Real-Time Payments network identification of the declined transfer.
	TransactionIdentification string `json:"transaction_identification,required"`
	// The identifier of the Real-Time Payments Transfer that led to this Transaction.
	TransferID string                                                      `json:"transfer_id,required"`
	JSON       transactionSourceInboundRealTimePaymentsTransferDeclineJSON `json:"-"`
}

// transactionSourceInboundRealTimePaymentsTransferDeclineJSON contains the JSON
// metadata for the struct
// [TransactionSourceInboundRealTimePaymentsTransferDecline]
type transactionSourceInboundRealTimePaymentsTransferDeclineJSON struct {
	Amount                    apijson.Field
	CreditorName              apijson.Field
	Currency                  apijson.Field
	DebtorAccountNumber       apijson.Field
	DebtorName                apijson.Field
	DebtorRoutingNumber       apijson.Field
	Reason                    apijson.Field
	RemittanceInformation     apijson.Field
	TransactionIdentification apijson.Field
	TransferID                apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *TransactionSourceInboundRealTimePaymentsTransferDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceInboundRealTimePaymentsTransferDeclineJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
// transfer's currency. This will always be "USD" for a Real-Time Payments
// transfer.
type TransactionSourceInboundRealTimePaymentsTransferDeclineCurrency string

const (
	TransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyCad TransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CAD"
	TransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyChf TransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CHF"
	TransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyEur TransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "EUR"
	TransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyGbp TransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "GBP"
	TransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyJpy TransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "JPY"
	TransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyUsd TransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "USD"
)

func (r TransactionSourceInboundRealTimePaymentsTransferDeclineCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyCad, TransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyChf, TransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyEur, TransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyGbp, TransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyJpy, TransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyUsd:
		return true
	}
	return false
}

// Why the transfer was declined.
type TransactionSourceInboundRealTimePaymentsTransferDeclineReason string

const (
	TransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberCanceled      TransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_canceled"
	TransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberDisabled      TransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_disabled"
	TransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountRestricted          TransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_restricted"
	TransactionSourceInboundRealTimePaymentsTransferDeclineReasonGroupLocked                TransactionSourceInboundRealTimePaymentsTransferDeclineReason = "group_locked"
	TransactionSourceInboundRealTimePaymentsTransferDeclineReasonEntityNotActive            TransactionSourceInboundRealTimePaymentsTransferDeclineReason = "entity_not_active"
	TransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled TransactionSourceInboundRealTimePaymentsTransferDeclineReason = "real_time_payments_not_enabled"
)

func (r TransactionSourceInboundRealTimePaymentsTransferDeclineReason) IsKnown() bool {
	switch r {
	case TransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberCanceled, TransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberDisabled, TransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountRestricted, TransactionSourceInboundRealTimePaymentsTransferDeclineReasonGroupLocked, TransactionSourceInboundRealTimePaymentsTransferDeclineReasonEntityNotActive, TransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled:
		return true
	}
	return false
}

// An Inbound Wire Reversal object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_reversal`. An Inbound Wire
// Reversal represents a reversal of a wire transfer that was initiated via
// Increase. The other bank is sending the money back. This most often happens when
// the original destination account details were incorrect.
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
	// Additional information included in the wire reversal by the reversal originator.
	OriginatorToBeneficiaryInformation string `json:"originator_to_beneficiary_information,required,nullable"`
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
	// The sending bank's reference number for the wire reversal.
	SenderReference string `json:"sender_reference,required,nullable"`
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
	OriginatorToBeneficiaryInformation                    apijson.Field
	PreviousMessageInputCycleDate                         apijson.Field
	PreviousMessageInputMessageAccountabilityData         apijson.Field
	PreviousMessageInputSequenceNumber                    apijson.Field
	PreviousMessageInputSource                            apijson.Field
	ReceiverFinancialInstitutionInformation               apijson.Field
	SenderReference                                       apijson.Field
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
// JSON response if and only if `category` is equal to `inbound_wire_transfer`. An
// Inbound Wire Transfer Intention is created when a wire transfer is initiated at
// another bank and received by Increase.
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

// An Inbound Wire Transfer Reversal Intention object. This field will be present
// in the JSON response if and only if `category` is equal to
// `inbound_wire_transfer_reversal`. An Inbound Wire Transfer Reversal Intention is
// created when Increase has received a wire and the User requests that it be
// reversed.
type TransactionSourceInboundWireTransferReversal struct {
	// The ID of the Inbound Wire Transfer that is being reversed.
	InboundWireTransferID string                                           `json:"inbound_wire_transfer_id,required"`
	JSON                  transactionSourceInboundWireTransferReversalJSON `json:"-"`
}

// transactionSourceInboundWireTransferReversalJSON contains the JSON metadata for
// the struct [TransactionSourceInboundWireTransferReversal]
type transactionSourceInboundWireTransferReversalJSON struct {
	InboundWireTransferID apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *TransactionSourceInboundWireTransferReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceInboundWireTransferReversalJSON) RawJSON() string {
	return r.raw
}

// An Interest Payment object. This field will be present in the JSON response if
// and only if `category` is equal to `interest_payment`. An Interest Payment
// represents a payment of interest on an account. Interest is usually paid
// monthly.
type TransactionSourceInterestPayment struct {
	// The account on which the interest was accrued.
	AccruedOnAccountID string `json:"accrued_on_account_id,required"`
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
	TransactionSourceInterestPaymentCurrencyCad TransactionSourceInterestPaymentCurrency = "CAD"
	TransactionSourceInterestPaymentCurrencyChf TransactionSourceInterestPaymentCurrency = "CHF"
	TransactionSourceInterestPaymentCurrencyEur TransactionSourceInterestPaymentCurrency = "EUR"
	TransactionSourceInterestPaymentCurrencyGbp TransactionSourceInterestPaymentCurrency = "GBP"
	TransactionSourceInterestPaymentCurrencyJpy TransactionSourceInterestPaymentCurrency = "JPY"
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
// and only if `category` is equal to `internal_source`. A transaction between the
// user and Increase. See the `reason` attribute for more information.
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
	TransactionSourceInternalSourceCurrencyCad TransactionSourceInternalSourceCurrency = "CAD"
	TransactionSourceInternalSourceCurrencyChf TransactionSourceInternalSourceCurrency = "CHF"
	TransactionSourceInternalSourceCurrencyEur TransactionSourceInternalSourceCurrency = "EUR"
	TransactionSourceInternalSourceCurrencyGbp TransactionSourceInternalSourceCurrency = "GBP"
	TransactionSourceInternalSourceCurrencyJpy TransactionSourceInternalSourceCurrency = "JPY"
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
	TransactionSourceInternalSourceReasonAccountClosure             TransactionSourceInternalSourceReason = "account_closure"
	TransactionSourceInternalSourceReasonBankDrawnCheck             TransactionSourceInternalSourceReason = "bank_drawn_check"
	TransactionSourceInternalSourceReasonBankDrawnCheckCredit       TransactionSourceInternalSourceReason = "bank_drawn_check_credit"
	TransactionSourceInternalSourceReasonBankMigration              TransactionSourceInternalSourceReason = "bank_migration"
	TransactionSourceInternalSourceReasonCheckAdjustment            TransactionSourceInternalSourceReason = "check_adjustment"
	TransactionSourceInternalSourceReasonCollectionPayment          TransactionSourceInternalSourceReason = "collection_payment"
	TransactionSourceInternalSourceReasonCollectionReceivable       TransactionSourceInternalSourceReason = "collection_receivable"
	TransactionSourceInternalSourceReasonEmpyrealAdjustment         TransactionSourceInternalSourceReason = "empyreal_adjustment"
	TransactionSourceInternalSourceReasonError                      TransactionSourceInternalSourceReason = "error"
	TransactionSourceInternalSourceReasonErrorCorrection            TransactionSourceInternalSourceReason = "error_correction"
	TransactionSourceInternalSourceReasonFees                       TransactionSourceInternalSourceReason = "fees"
	TransactionSourceInternalSourceReasonInterest                   TransactionSourceInternalSourceReason = "interest"
	TransactionSourceInternalSourceReasonNegativeBalanceForgiveness TransactionSourceInternalSourceReason = "negative_balance_forgiveness"
	TransactionSourceInternalSourceReasonSampleFunds                TransactionSourceInternalSourceReason = "sample_funds"
	TransactionSourceInternalSourceReasonSampleFundsReturn          TransactionSourceInternalSourceReason = "sample_funds_return"
)

func (r TransactionSourceInternalSourceReason) IsKnown() bool {
	switch r {
	case TransactionSourceInternalSourceReasonAccountClosure, TransactionSourceInternalSourceReasonBankDrawnCheck, TransactionSourceInternalSourceReasonBankDrawnCheckCredit, TransactionSourceInternalSourceReasonBankMigration, TransactionSourceInternalSourceReasonCheckAdjustment, TransactionSourceInternalSourceReasonCollectionPayment, TransactionSourceInternalSourceReasonCollectionReceivable, TransactionSourceInternalSourceReasonEmpyrealAdjustment, TransactionSourceInternalSourceReasonError, TransactionSourceInternalSourceReasonErrorCorrection, TransactionSourceInternalSourceReasonFees, TransactionSourceInternalSourceReasonInterest, TransactionSourceInternalSourceReasonNegativeBalanceForgiveness, TransactionSourceInternalSourceReasonSampleFunds, TransactionSourceInternalSourceReasonSampleFundsReturn:
		return true
	}
	return false
}

// A Real-Time Payments Transfer Acknowledgement object. This field will be present
// in the JSON response if and only if `category` is equal to
// `real_time_payments_transfer_acknowledgement`. A Real-Time Payments Transfer
// Acknowledgement is created when a Real-Time Payments Transfer sent from Increase
// is acknowledged by the receiving bank.
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
// only if `category` is equal to `sample_funds`. Sample funds for testing
// purposes.
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

// A Swift Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `swift_transfer_intention`. A
// Swift Transfer initiated via Increase.
type TransactionSourceSwiftTransferIntention struct {
	// The identifier of the Swift Transfer that led to this Transaction.
	TransferID string                                      `json:"transfer_id,required"`
	JSON       transactionSourceSwiftTransferIntentionJSON `json:"-"`
}

// transactionSourceSwiftTransferIntentionJSON contains the JSON metadata for the
// struct [TransactionSourceSwiftTransferIntention]
type transactionSourceSwiftTransferIntentionJSON struct {
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceSwiftTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceSwiftTransferIntentionJSON) RawJSON() string {
	return r.raw
}

// A Wire Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_intention`. A Wire
// Transfer initiated via Increase and sent to a different bank.
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
	TransactionListParamsCategoryInAccountTransferIntention                    TransactionListParamsCategoryIn = "account_transfer_intention"
	TransactionListParamsCategoryInACHTransferIntention                        TransactionListParamsCategoryIn = "ach_transfer_intention"
	TransactionListParamsCategoryInACHTransferRejection                        TransactionListParamsCategoryIn = "ach_transfer_rejection"
	TransactionListParamsCategoryInACHTransferReturn                           TransactionListParamsCategoryIn = "ach_transfer_return"
	TransactionListParamsCategoryInCashbackPayment                             TransactionListParamsCategoryIn = "cashback_payment"
	TransactionListParamsCategoryInCardDisputeAcceptance                       TransactionListParamsCategoryIn = "card_dispute_acceptance"
	TransactionListParamsCategoryInCardDisputeFinancial                        TransactionListParamsCategoryIn = "card_dispute_financial"
	TransactionListParamsCategoryInCardDisputeLoss                             TransactionListParamsCategoryIn = "card_dispute_loss"
	TransactionListParamsCategoryInCardRefund                                  TransactionListParamsCategoryIn = "card_refund"
	TransactionListParamsCategoryInCardSettlement                              TransactionListParamsCategoryIn = "card_settlement"
	TransactionListParamsCategoryInCardRevenuePayment                          TransactionListParamsCategoryIn = "card_revenue_payment"
	TransactionListParamsCategoryInCheckDepositAcceptance                      TransactionListParamsCategoryIn = "check_deposit_acceptance"
	TransactionListParamsCategoryInCheckDepositReturn                          TransactionListParamsCategoryIn = "check_deposit_return"
	TransactionListParamsCategoryInCheckTransferDeposit                        TransactionListParamsCategoryIn = "check_transfer_deposit"
	TransactionListParamsCategoryInFeePayment                                  TransactionListParamsCategoryIn = "fee_payment"
	TransactionListParamsCategoryInInboundACHTransfer                          TransactionListParamsCategoryIn = "inbound_ach_transfer"
	TransactionListParamsCategoryInInboundACHTransferReturnIntention           TransactionListParamsCategoryIn = "inbound_ach_transfer_return_intention"
	TransactionListParamsCategoryInInboundCheckDepositReturnIntention          TransactionListParamsCategoryIn = "inbound_check_deposit_return_intention"
	TransactionListParamsCategoryInInboundCheckAdjustment                      TransactionListParamsCategoryIn = "inbound_check_adjustment"
	TransactionListParamsCategoryInInboundRealTimePaymentsTransferConfirmation TransactionListParamsCategoryIn = "inbound_real_time_payments_transfer_confirmation"
	TransactionListParamsCategoryInInboundRealTimePaymentsTransferDecline      TransactionListParamsCategoryIn = "inbound_real_time_payments_transfer_decline"
	TransactionListParamsCategoryInInboundWireReversal                         TransactionListParamsCategoryIn = "inbound_wire_reversal"
	TransactionListParamsCategoryInInboundWireTransfer                         TransactionListParamsCategoryIn = "inbound_wire_transfer"
	TransactionListParamsCategoryInInboundWireTransferReversal                 TransactionListParamsCategoryIn = "inbound_wire_transfer_reversal"
	TransactionListParamsCategoryInInterestPayment                             TransactionListParamsCategoryIn = "interest_payment"
	TransactionListParamsCategoryInInternalSource                              TransactionListParamsCategoryIn = "internal_source"
	TransactionListParamsCategoryInRealTimePaymentsTransferAcknowledgement     TransactionListParamsCategoryIn = "real_time_payments_transfer_acknowledgement"
	TransactionListParamsCategoryInSampleFunds                                 TransactionListParamsCategoryIn = "sample_funds"
	TransactionListParamsCategoryInWireTransferIntention                       TransactionListParamsCategoryIn = "wire_transfer_intention"
	TransactionListParamsCategoryInSwiftTransferIntention                      TransactionListParamsCategoryIn = "swift_transfer_intention"
	TransactionListParamsCategoryInCardPushTransferAcceptance                  TransactionListParamsCategoryIn = "card_push_transfer_acceptance"
	TransactionListParamsCategoryInOther                                       TransactionListParamsCategoryIn = "other"
)

func (r TransactionListParamsCategoryIn) IsKnown() bool {
	switch r {
	case TransactionListParamsCategoryInAccountTransferIntention, TransactionListParamsCategoryInACHTransferIntention, TransactionListParamsCategoryInACHTransferRejection, TransactionListParamsCategoryInACHTransferReturn, TransactionListParamsCategoryInCashbackPayment, TransactionListParamsCategoryInCardDisputeAcceptance, TransactionListParamsCategoryInCardDisputeFinancial, TransactionListParamsCategoryInCardDisputeLoss, TransactionListParamsCategoryInCardRefund, TransactionListParamsCategoryInCardSettlement, TransactionListParamsCategoryInCardRevenuePayment, TransactionListParamsCategoryInCheckDepositAcceptance, TransactionListParamsCategoryInCheckDepositReturn, TransactionListParamsCategoryInCheckTransferDeposit, TransactionListParamsCategoryInFeePayment, TransactionListParamsCategoryInInboundACHTransfer, TransactionListParamsCategoryInInboundACHTransferReturnIntention, TransactionListParamsCategoryInInboundCheckDepositReturnIntention, TransactionListParamsCategoryInInboundCheckAdjustment, TransactionListParamsCategoryInInboundRealTimePaymentsTransferConfirmation, TransactionListParamsCategoryInInboundRealTimePaymentsTransferDecline, TransactionListParamsCategoryInInboundWireReversal, TransactionListParamsCategoryInInboundWireTransfer, TransactionListParamsCategoryInInboundWireTransferReversal, TransactionListParamsCategoryInInterestPayment, TransactionListParamsCategoryInInternalSource, TransactionListParamsCategoryInRealTimePaymentsTransferAcknowledgement, TransactionListParamsCategoryInSampleFunds, TransactionListParamsCategoryInWireTransferIntention, TransactionListParamsCategoryInSwiftTransferIntention, TransactionListParamsCategoryInCardPushTransferAcceptance, TransactionListParamsCategoryInOther:
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
