// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	ID string `json:"id" api:"required"`
	// The identifier for the Account the Transaction belongs to.
	AccountID string `json:"account_id" api:"required"`
	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occurred.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transaction's
	// Account.
	Currency TransactionCurrency `json:"currency" api:"required"`
	// An informational message describing this transaction. Use the fields in `source`
	// to get more detailed information. This field appears as the line-item on the
	// statement.
	Description string `json:"description" api:"required"`
	// The identifier for the route this Transaction came through. Routes are things
	// like cards and ACH details.
	RouteID string `json:"route_id" api:"required,nullable"`
	// The type of the route this Transaction came through.
	RouteType TransactionRouteType `json:"route_type" api:"required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Transaction. Note that for backwards compatibility reasons, additional
	// undocumented keys may appear in this object. These should be treated as
	// deprecated and will be removed in the future.
	Source TransactionSource `json:"source" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type        TransactionType        `json:"type" api:"required"`
	ExtraFields map[string]interface{} `json:"-" api:"extrafields"`
	JSON        transactionJSON        `json:"-"`
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
	TransactionCurrencyUsd TransactionCurrency = "USD"
)

func (r TransactionCurrency) IsKnown() bool {
	switch r {
	case TransactionCurrencyUsd:
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
	// The type of the resource. We may add additional possible values for this enum
	// over time; your application should be able to handle such additions gracefully.
	Category TransactionSourceCategory `json:"category" api:"required"`
	// An Account Revenue Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_revenue_payment`. An
	// Account Revenue Payment represents a payment made to an account from the bank.
	// Account revenue is a type of non-interest income.
	AccountRevenuePayment TransactionSourceAccountRevenuePayment `json:"account_revenue_payment" api:"nullable"`
	// An Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`. Two
	// Account Transfer Intentions are created from each Account Transfer. One
	// decrements the source account, and the other increments the destination account.
	AccountTransferIntention TransactionSourceAccountTransferIntention `json:"account_transfer_intention" api:"nullable"`
	// An ACH Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_intention`. An ACH
	// Transfer Intention is created from an ACH Transfer. It reflects the intention to
	// move money into or out of an Increase account via the ACH network.
	ACHTransferIntention TransactionSourceACHTransferIntention `json:"ach_transfer_intention" api:"nullable"`
	// An ACH Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_rejection`. An ACH
	// Transfer Rejection is created when an ACH Transfer is rejected by Increase. It
	// offsets the ACH Transfer Intention. These rejections are rare.
	ACHTransferRejection TransactionSourceACHTransferRejection `json:"ach_transfer_rejection" api:"nullable"`
	// An ACH Transfer Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_return`. An ACH Transfer
	// Return is created when an ACH Transfer is returned by the receiving bank. It
	// offsets the ACH Transfer Intention. ACH Transfer Returns usually occur within
	// the first two business days after the transfer is initiated, but can occur much
	// later.
	ACHTransferReturn TransactionSourceACHTransferReturn `json:"ach_transfer_return" api:"nullable"`
	// A Blockchain Off-Ramp Transfer Settlement object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `blockchain_offramp_transfer_settlement`.
	BlockchainOfframpTransferSettlement TransactionSourceBlockchainOfframpTransferSettlement `json:"blockchain_offramp_transfer_settlement" api:"nullable"`
	// A Blockchain On-Ramp Transfer Intention object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `blockchain_onramp_transfer_intention`.
	BlockchainOnrampTransferIntention TransactionSourceBlockchainOnrampTransferIntention `json:"blockchain_onramp_transfer_intention" api:"nullable"`
	// A Legacy Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	// Contains the details of a successful Card Dispute.
	CardDisputeAcceptance TransactionSourceCardDisputeAcceptance `json:"card_dispute_acceptance" api:"nullable"`
	// A Card Dispute Financial object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_dispute_financial`. Financial event
	// related to a Card Dispute.
	CardDisputeFinancial TransactionSourceCardDisputeFinancial `json:"card_dispute_financial" api:"nullable"`
	// A Legacy Card Dispute Loss object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_loss`. Contains the
	// details of a lost Card Dispute.
	CardDisputeLoss TransactionSourceCardDisputeLoss `json:"card_dispute_loss" api:"nullable"`
	// A Card Financial object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_financial`. Card Financials are temporary
	// holds placed on a customer's funds with the intent to later clear a transaction.
	CardFinancial TransactionSourceCardFinancial `json:"card_financial" api:"nullable"`
	// A Card Push Transfer Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_push_transfer_acceptance`.
	// A Card Push Transfer Acceptance is created when an Outbound Card Push Transfer
	// sent from Increase is accepted by the receiving bank.
	CardPushTransferAcceptance TransactionSourceCardPushTransferAcceptance `json:"card_push_transfer_acceptance" api:"nullable"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`. Card Refunds move money back to
	// the cardholder. While they are usually connected to a Card Settlement, an
	// acquirer can also refund money directly to a card without relation to a
	// transaction.
	CardRefund TransactionSourceCardRefund `json:"card_refund" api:"nullable"`
	// A Card Revenue Payment object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_revenue_payment`. Card Revenue
	// Payments reflect earnings from fees on card transactions.
	CardRevenuePayment TransactionSourceCardRevenuePayment `json:"card_revenue_payment" api:"nullable"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`. Card Settlements are card
	// transactions that have cleared and settled. While a settlement is usually
	// preceded by an authorization, an acquirer can also directly clear a transaction
	// without first authorizing it.
	CardSettlement TransactionSourceCardSettlement `json:"card_settlement" api:"nullable"`
	// A Cashback Payment object. This field will be present in the JSON response if
	// and only if `category` is equal to `cashback_payment`. A Cashback Payment
	// represents the cashback paid to a cardholder for a given period. Cashback is
	// usually paid monthly for the prior month's transactions.
	CashbackPayment TransactionSourceCashbackPayment `json:"cashback_payment" api:"nullable"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`. A
	// Check Deposit Acceptance is created when a Check Deposit is processed and its
	// details confirmed. Check Deposits may be returned by the receiving bank, which
	// will appear as a Check Deposit Return.
	CheckDepositAcceptance TransactionSourceCheckDepositAcceptance `json:"check_deposit_acceptance" api:"nullable"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`. A Check Deposit
	// Return is created when a Check Deposit is returned by the bank holding the
	// account it was drawn against. Check Deposits may be returned for a variety of
	// reasons, including insufficient funds or a mismatched account number. Usually,
	// checks are returned within the first 7 days after the deposit is made.
	CheckDepositReturn TransactionSourceCheckDepositReturn `json:"check_deposit_return" api:"nullable"`
	// A Check Transfer Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_transfer_deposit`. An Inbound Check
	// is a check drawn on an Increase account that has been deposited by an external
	// bank account. These types of checks are not pre-registered.
	CheckTransferDeposit TransactionSourceCheckTransferDeposit `json:"check_transfer_deposit" api:"nullable"`
	// A FedNow Transfer Acknowledgement object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `fednow_transfer_acknowledgement`. A FedNow Transfer Acknowledgement is created
	// when a FedNow Transfer sent from Increase is acknowledged by the receiving bank.
	FednowTransferAcknowledgement TransactionSourceFednowTransferAcknowledgement `json:"fednow_transfer_acknowledgement" api:"nullable"`
	// A Fee Payment object. This field will be present in the JSON response if and
	// only if `category` is equal to `fee_payment`. A Fee Payment represents a payment
	// made to Increase.
	FeePayment TransactionSourceFeePayment `json:"fee_payment" api:"nullable"`
	// An Inbound ACH Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_ach_transfer`. An
	// Inbound ACH Transfer Intention is created when an ACH transfer is initiated at
	// another bank and received by Increase.
	InboundACHTransfer TransactionSourceInboundACHTransfer `json:"inbound_ach_transfer" api:"nullable"`
	// An Inbound ACH Transfer Return Intention object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_ach_transfer_return_intention`. An Inbound ACH Transfer Return
	// Intention is created when an ACH transfer is initiated at another bank and
	// returned by Increase.
	InboundACHTransferReturnIntention TransactionSourceInboundACHTransferReturnIntention `json:"inbound_ach_transfer_return_intention" api:"nullable"`
	// An Inbound Check Adjustment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_check_adjustment`. An
	// Inbound Check Adjustment is created when Increase receives an adjustment for a
	// check or return deposited through Check21.
	InboundCheckAdjustment TransactionSourceInboundCheckAdjustment `json:"inbound_check_adjustment" api:"nullable"`
	// An Inbound Check Deposit Return Intention object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_check_deposit_return_intention`. An Inbound Check Deposit Return
	// Intention is created when Increase receives an Inbound Check and the User
	// requests that it be returned.
	InboundCheckDepositReturnIntention TransactionSourceInboundCheckDepositReturnIntention `json:"inbound_check_deposit_return_intention" api:"nullable"`
	// An Inbound FedNow Transfer Confirmation object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_fednow_transfer_confirmation`. An Inbound FedNow Transfer Confirmation
	// is created when a FedNow transfer is initiated at another bank and received by
	// Increase.
	InboundFednowTransferConfirmation TransactionSourceInboundFednowTransferConfirmation `json:"inbound_fednow_transfer_confirmation" api:"nullable"`
	// An Inbound Real-Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`. An Inbound Real-Time
	// Payments Transfer Confirmation is created when a Real-Time Payments transfer is
	// initiated at another bank and received by Increase.
	InboundRealTimePaymentsTransferConfirmation TransactionSourceInboundRealTimePaymentsTransferConfirmation `json:"inbound_real_time_payments_transfer_confirmation" api:"nullable"`
	// An Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`. An Inbound Wire
	// Reversal represents a reversal of a wire transfer that was initiated via
	// Increase. The other bank is sending the money back. This most often happens when
	// the original destination account details were incorrect.
	InboundWireReversal TransactionSourceInboundWireReversal `json:"inbound_wire_reversal" api:"nullable"`
	// An Inbound Wire Transfer Intention object. This field will be present in the
	// JSON response if and only if `category` is equal to `inbound_wire_transfer`. An
	// Inbound Wire Transfer Intention is created when a wire transfer is initiated at
	// another bank and received by Increase.
	InboundWireTransfer TransactionSourceInboundWireTransfer `json:"inbound_wire_transfer" api:"nullable"`
	// An Inbound Wire Transfer Reversal Intention object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_wire_transfer_reversal`. An Inbound Wire Transfer Reversal Intention is
	// created when Increase has received a wire and the User requests that it be
	// reversed.
	InboundWireTransferReversal TransactionSourceInboundWireTransferReversal `json:"inbound_wire_transfer_reversal" api:"nullable"`
	// An Interest Payment object. This field will be present in the JSON response if
	// and only if `category` is equal to `interest_payment`. An Interest Payment
	// represents a payment of interest on an account. Interest is usually paid
	// monthly.
	InterestPayment TransactionSourceInterestPayment `json:"interest_payment" api:"nullable"`
	// An Internal Source object. This field will be present in the JSON response if
	// and only if `category` is equal to `internal_source`. A transaction between the
	// user and Increase. See the `reason` attribute for more information.
	InternalSource TransactionSourceInternalSource `json:"internal_source" api:"nullable"`
	// If the category of this Transaction source is equal to `other`, this field will
	// contain an empty object, otherwise it will contain null.
	Other TransactionSourceOther `json:"other" api:"nullable"`
	// A Real-Time Payments Transfer Acknowledgement object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `real_time_payments_transfer_acknowledgement`. A Real-Time Payments Transfer
	// Acknowledgement is created when a Real-Time Payments Transfer sent from Increase
	// is acknowledged by the receiving bank.
	RealTimePaymentsTransferAcknowledgement TransactionSourceRealTimePaymentsTransferAcknowledgement `json:"real_time_payments_transfer_acknowledgement" api:"nullable"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`. Sample funds for testing
	// purposes.
	SampleFunds TransactionSourceSampleFunds `json:"sample_funds" api:"nullable"`
	// A Swift Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `swift_transfer_intention`. A
	// Swift Transfer initiated via Increase.
	SwiftTransferIntention TransactionSourceSwiftTransferIntention `json:"swift_transfer_intention" api:"nullable"`
	// A Swift Transfer Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `swift_transfer_return`. A Swift Transfer
	// Return is created when a Swift Transfer is returned by the receiving bank.
	SwiftTransferReturn TransactionSourceSwiftTransferReturn `json:"swift_transfer_return" api:"nullable"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`. A Wire
	// Transfer initiated via Increase and sent to a different bank.
	WireTransferIntention TransactionSourceWireTransferIntention `json:"wire_transfer_intention" api:"nullable"`
	ExtraFields           map[string]interface{}                 `json:"-" api:"extrafields"`
	JSON                  transactionSourceJSON                  `json:"-"`
}

// transactionSourceJSON contains the JSON metadata for the struct
// [TransactionSource]
type transactionSourceJSON struct {
	Category                                    apijson.Field
	AccountRevenuePayment                       apijson.Field
	AccountTransferIntention                    apijson.Field
	ACHTransferIntention                        apijson.Field
	ACHTransferRejection                        apijson.Field
	ACHTransferReturn                           apijson.Field
	BlockchainOfframpTransferSettlement         apijson.Field
	BlockchainOnrampTransferIntention           apijson.Field
	CardDisputeAcceptance                       apijson.Field
	CardDisputeFinancial                        apijson.Field
	CardDisputeLoss                             apijson.Field
	CardFinancial                               apijson.Field
	CardPushTransferAcceptance                  apijson.Field
	CardRefund                                  apijson.Field
	CardRevenuePayment                          apijson.Field
	CardSettlement                              apijson.Field
	CashbackPayment                             apijson.Field
	CheckDepositAcceptance                      apijson.Field
	CheckDepositReturn                          apijson.Field
	CheckTransferDeposit                        apijson.Field
	FednowTransferAcknowledgement               apijson.Field
	FeePayment                                  apijson.Field
	InboundACHTransfer                          apijson.Field
	InboundACHTransferReturnIntention           apijson.Field
	InboundCheckAdjustment                      apijson.Field
	InboundCheckDepositReturnIntention          apijson.Field
	InboundFednowTransferConfirmation           apijson.Field
	InboundRealTimePaymentsTransferConfirmation apijson.Field
	InboundWireReversal                         apijson.Field
	InboundWireTransfer                         apijson.Field
	InboundWireTransferReversal                 apijson.Field
	InterestPayment                             apijson.Field
	InternalSource                              apijson.Field
	Other                                       apijson.Field
	RealTimePaymentsTransferAcknowledgement     apijson.Field
	SampleFunds                                 apijson.Field
	SwiftTransferIntention                      apijson.Field
	SwiftTransferReturn                         apijson.Field
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
	TransactionSourceCategoryCardFinancial                               TransactionSourceCategory = "card_financial"
	TransactionSourceCategoryCardRevenuePayment                          TransactionSourceCategory = "card_revenue_payment"
	TransactionSourceCategoryCheckDepositAcceptance                      TransactionSourceCategory = "check_deposit_acceptance"
	TransactionSourceCategoryCheckDepositReturn                          TransactionSourceCategory = "check_deposit_return"
	TransactionSourceCategoryFednowTransferAcknowledgement               TransactionSourceCategory = "fednow_transfer_acknowledgement"
	TransactionSourceCategoryCheckTransferDeposit                        TransactionSourceCategory = "check_transfer_deposit"
	TransactionSourceCategoryFeePayment                                  TransactionSourceCategory = "fee_payment"
	TransactionSourceCategoryInboundACHTransfer                          TransactionSourceCategory = "inbound_ach_transfer"
	TransactionSourceCategoryInboundACHTransferReturnIntention           TransactionSourceCategory = "inbound_ach_transfer_return_intention"
	TransactionSourceCategoryInboundCheckDepositReturnIntention          TransactionSourceCategory = "inbound_check_deposit_return_intention"
	TransactionSourceCategoryInboundCheckAdjustment                      TransactionSourceCategory = "inbound_check_adjustment"
	TransactionSourceCategoryInboundFednowTransferConfirmation           TransactionSourceCategory = "inbound_fednow_transfer_confirmation"
	TransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation TransactionSourceCategory = "inbound_real_time_payments_transfer_confirmation"
	TransactionSourceCategoryInboundWireReversal                         TransactionSourceCategory = "inbound_wire_reversal"
	TransactionSourceCategoryInboundWireTransfer                         TransactionSourceCategory = "inbound_wire_transfer"
	TransactionSourceCategoryInboundWireTransferReversal                 TransactionSourceCategory = "inbound_wire_transfer_reversal"
	TransactionSourceCategoryInterestPayment                             TransactionSourceCategory = "interest_payment"
	TransactionSourceCategoryInternalSource                              TransactionSourceCategory = "internal_source"
	TransactionSourceCategoryRealTimePaymentsTransferAcknowledgement     TransactionSourceCategory = "real_time_payments_transfer_acknowledgement"
	TransactionSourceCategorySampleFunds                                 TransactionSourceCategory = "sample_funds"
	TransactionSourceCategoryWireTransferIntention                       TransactionSourceCategory = "wire_transfer_intention"
	TransactionSourceCategorySwiftTransferIntention                      TransactionSourceCategory = "swift_transfer_intention"
	TransactionSourceCategorySwiftTransferReturn                         TransactionSourceCategory = "swift_transfer_return"
	TransactionSourceCategoryCardPushTransferAcceptance                  TransactionSourceCategory = "card_push_transfer_acceptance"
	TransactionSourceCategoryAccountRevenuePayment                       TransactionSourceCategory = "account_revenue_payment"
	TransactionSourceCategoryBlockchainOnrampTransferIntention           TransactionSourceCategory = "blockchain_onramp_transfer_intention"
	TransactionSourceCategoryBlockchainOfframpTransferSettlement         TransactionSourceCategory = "blockchain_offramp_transfer_settlement"
	TransactionSourceCategoryOther                                       TransactionSourceCategory = "other"
)

func (r TransactionSourceCategory) IsKnown() bool {
	switch r {
	case TransactionSourceCategoryAccountTransferIntention, TransactionSourceCategoryACHTransferIntention, TransactionSourceCategoryACHTransferRejection, TransactionSourceCategoryACHTransferReturn, TransactionSourceCategoryCashbackPayment, TransactionSourceCategoryCardDisputeAcceptance, TransactionSourceCategoryCardDisputeFinancial, TransactionSourceCategoryCardDisputeLoss, TransactionSourceCategoryCardRefund, TransactionSourceCategoryCardSettlement, TransactionSourceCategoryCardFinancial, TransactionSourceCategoryCardRevenuePayment, TransactionSourceCategoryCheckDepositAcceptance, TransactionSourceCategoryCheckDepositReturn, TransactionSourceCategoryFednowTransferAcknowledgement, TransactionSourceCategoryCheckTransferDeposit, TransactionSourceCategoryFeePayment, TransactionSourceCategoryInboundACHTransfer, TransactionSourceCategoryInboundACHTransferReturnIntention, TransactionSourceCategoryInboundCheckDepositReturnIntention, TransactionSourceCategoryInboundCheckAdjustment, TransactionSourceCategoryInboundFednowTransferConfirmation, TransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation, TransactionSourceCategoryInboundWireReversal, TransactionSourceCategoryInboundWireTransfer, TransactionSourceCategoryInboundWireTransferReversal, TransactionSourceCategoryInterestPayment, TransactionSourceCategoryInternalSource, TransactionSourceCategoryRealTimePaymentsTransferAcknowledgement, TransactionSourceCategorySampleFunds, TransactionSourceCategoryWireTransferIntention, TransactionSourceCategorySwiftTransferIntention, TransactionSourceCategorySwiftTransferReturn, TransactionSourceCategoryCardPushTransferAcceptance, TransactionSourceCategoryAccountRevenuePayment, TransactionSourceCategoryBlockchainOnrampTransferIntention, TransactionSourceCategoryBlockchainOfframpTransferSettlement, TransactionSourceCategoryOther:
		return true
	}
	return false
}

// An Account Revenue Payment object. This field will be present in the JSON
// response if and only if `category` is equal to `account_revenue_payment`. An
// Account Revenue Payment represents a payment made to an account from the bank.
// Account revenue is a type of non-interest income.
type TransactionSourceAccountRevenuePayment struct {
	// The account on which the account revenue was accrued.
	AccruedOnAccountID string `json:"accrued_on_account_id" api:"required"`
	// The end of the period for which this transaction paid account revenue.
	PeriodEnd time.Time `json:"period_end" api:"required" format:"date-time"`
	// The start of the period for which this transaction paid account revenue.
	PeriodStart time.Time                                  `json:"period_start" api:"required" format:"date-time"`
	ExtraFields map[string]interface{}                     `json:"-" api:"extrafields"`
	JSON        transactionSourceAccountRevenuePaymentJSON `json:"-"`
}

// transactionSourceAccountRevenuePaymentJSON contains the JSON metadata for the
// struct [TransactionSourceAccountRevenuePayment]
type transactionSourceAccountRevenuePaymentJSON struct {
	AccruedOnAccountID apijson.Field
	PeriodEnd          apijson.Field
	PeriodStart        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TransactionSourceAccountRevenuePayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceAccountRevenuePaymentJSON) RawJSON() string {
	return r.raw
}

// An Account Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_intention`. Two
// Account Transfer Intentions are created from each Account Transfer. One
// decrements the source account, and the other increments the destination account.
type TransactionSourceAccountTransferIntention struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency TransactionSourceAccountTransferIntentionCurrency `json:"currency" api:"required"`
	// The description you chose to give the transfer.
	Description string `json:"description" api:"required"`
	// The identifier of the Account to where the Account Transfer was sent.
	DestinationAccountID string `json:"destination_account_id" api:"required"`
	// The identifier of the Account from where the Account Transfer was sent.
	SourceAccountID string `json:"source_account_id" api:"required"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID  string                                        `json:"transfer_id" api:"required"`
	ExtraFields map[string]interface{}                        `json:"-" api:"extrafields"`
	JSON        transactionSourceAccountTransferIntentionJSON `json:"-"`
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
	TransactionSourceAccountTransferIntentionCurrencyUsd TransactionSourceAccountTransferIntentionCurrency = "USD"
)

func (r TransactionSourceAccountTransferIntentionCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceAccountTransferIntentionCurrencyUsd:
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
	AccountNumber string `json:"account_number" api:"required"`
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount" api:"required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber string `json:"routing_number" api:"required"`
	// A description set when the ACH Transfer was created.
	StatementDescriptor string `json:"statement_descriptor" api:"required"`
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID  string                                    `json:"transfer_id" api:"required"`
	ExtraFields map[string]interface{}                    `json:"-" api:"extrafields"`
	JSON        transactionSourceACHTransferIntentionJSON `json:"-"`
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
	TransferID  string                                    `json:"transfer_id" api:"required"`
	ExtraFields map[string]interface{}                    `json:"-" api:"extrafields"`
	JSON        transactionSourceACHTransferRejectionJSON `json:"-"`
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
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The three character ACH return code, in the range R01 to R85.
	RawReturnReasonCode string `json:"raw_return_reason_code" api:"required"`
	// Why the ACH Transfer was returned. This reason code is sent by the receiving
	// bank back to Increase.
	ReturnReasonCode TransactionSourceACHTransferReturnReturnReasonCode `json:"return_reason_code" api:"required"`
	// A 15 digit number that was generated by the bank that initiated the return. The
	// trace number of the return is different than that of the original transfer. ACH
	// trace numbers are not unique, but along with the amount and date this number can
	// be used to identify the ACH return at the bank that initiated it.
	TraceNumber string `json:"trace_number" api:"required"`
	// The identifier of the Transaction associated with this return.
	TransactionID string `json:"transaction_id" api:"required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID  string                                 `json:"transfer_id" api:"required"`
	ExtraFields map[string]interface{}                 `json:"-" api:"extrafields"`
	JSON        transactionSourceACHTransferReturnJSON `json:"-"`
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

// A Blockchain Off-Ramp Transfer Settlement object. This field will be present in
// the JSON response if and only if `category` is equal to
// `blockchain_offramp_transfer_settlement`.
type TransactionSourceBlockchainOfframpTransferSettlement struct {
	// The identifier of the Blockchain Address the funds were received at.
	SourceBlockchainAddressID string `json:"source_blockchain_address_id" api:"required"`
	// The identifier of the Blockchain Off-Ramp Transfer that led to this Transaction.
	TransferID  string                                                   `json:"transfer_id" api:"required"`
	ExtraFields map[string]interface{}                                   `json:"-" api:"extrafields"`
	JSON        transactionSourceBlockchainOfframpTransferSettlementJSON `json:"-"`
}

// transactionSourceBlockchainOfframpTransferSettlementJSON contains the JSON
// metadata for the struct [TransactionSourceBlockchainOfframpTransferSettlement]
type transactionSourceBlockchainOfframpTransferSettlementJSON struct {
	SourceBlockchainAddressID apijson.Field
	TransferID                apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *TransactionSourceBlockchainOfframpTransferSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceBlockchainOfframpTransferSettlementJSON) RawJSON() string {
	return r.raw
}

// A Blockchain On-Ramp Transfer Intention object. This field will be present in
// the JSON response if and only if `category` is equal to
// `blockchain_onramp_transfer_intention`.
type TransactionSourceBlockchainOnrampTransferIntention struct {
	// The blockchain address the funds were sent to.
	DestinationBlockchainAddress string `json:"destination_blockchain_address" api:"required"`
	// The identifier of the Blockchain On-Ramp Transfer that led to this Transaction.
	TransferID  string                                                 `json:"transfer_id" api:"required"`
	ExtraFields map[string]interface{}                                 `json:"-" api:"extrafields"`
	JSON        transactionSourceBlockchainOnrampTransferIntentionJSON `json:"-"`
}

// transactionSourceBlockchainOnrampTransferIntentionJSON contains the JSON
// metadata for the struct [TransactionSourceBlockchainOnrampTransferIntention]
type transactionSourceBlockchainOnrampTransferIntentionJSON struct {
	DestinationBlockchainAddress apijson.Field
	TransferID                   apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *TransactionSourceBlockchainOnrampTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceBlockchainOnrampTransferIntentionJSON) RawJSON() string {
	return r.raw
}

// A Legacy Card Dispute Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `card_dispute_acceptance`.
// Contains the details of a successful Card Dispute.
type TransactionSourceCardDisputeAcceptance struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was accepted.
	AcceptedAt time.Time `json:"accepted_at" api:"required" format:"date-time"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID string                                     `json:"transaction_id" api:"required"`
	ExtraFields   map[string]interface{}                     `json:"-" api:"extrafields"`
	JSON          transactionSourceCardDisputeAcceptanceJSON `json:"-"`
}

// transactionSourceCardDisputeAcceptanceJSON contains the JSON metadata for the
// struct [TransactionSourceCardDisputeAcceptance]
type transactionSourceCardDisputeAcceptanceJSON struct {
	AcceptedAt    apijson.Field
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
	Amount int64 `json:"amount" api:"required"`
	// The network that the Card Dispute is associated with.
	Network TransactionSourceCardDisputeFinancialNetwork `json:"network" api:"required"`
	// The identifier of the Transaction that was created to credit or debit the
	// disputed funds to or from your account.
	TransactionID string `json:"transaction_id" api:"required"`
	// Information for events related to card dispute for card payments processed over
	// Visa's network. This field will be present in the JSON response if and only if
	// `network` is equal to `visa`.
	Visa        TransactionSourceCardDisputeFinancialVisa `json:"visa" api:"required,nullable"`
	ExtraFields map[string]interface{}                    `json:"-" api:"extrafields"`
	JSON        transactionSourceCardDisputeFinancialJSON `json:"-"`
}

// transactionSourceCardDisputeFinancialJSON contains the JSON metadata for the
// struct [TransactionSourceCardDisputeFinancial]
type transactionSourceCardDisputeFinancialJSON struct {
	Amount        apijson.Field
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
	TransactionSourceCardDisputeFinancialNetworkVisa  TransactionSourceCardDisputeFinancialNetwork = "visa"
	TransactionSourceCardDisputeFinancialNetworkPulse TransactionSourceCardDisputeFinancialNetwork = "pulse"
)

func (r TransactionSourceCardDisputeFinancialNetwork) IsKnown() bool {
	switch r {
	case TransactionSourceCardDisputeFinancialNetworkVisa, TransactionSourceCardDisputeFinancialNetworkPulse:
		return true
	}
	return false
}

// Information for events related to card dispute for card payments processed over
// Visa's network. This field will be present in the JSON response if and only if
// `network` is equal to `visa`.
type TransactionSourceCardDisputeFinancialVisa struct {
	// The type of card dispute financial event.
	EventType TransactionSourceCardDisputeFinancialVisaEventType `json:"event_type" api:"required"`
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

// A Legacy Card Dispute Loss object. This field will be present in the JSON
// response if and only if `category` is equal to `card_dispute_loss`. Contains the
// details of a lost Card Dispute.
type TransactionSourceCardDisputeLoss struct {
	// Why the Card Dispute was lost.
	Explanation string `json:"explanation" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was lost.
	LostAt time.Time `json:"lost_at" api:"required" format:"date-time"`
	// The identifier of the Transaction that was created to debit the disputed funds
	// from your account.
	TransactionID string                               `json:"transaction_id" api:"required"`
	ExtraFields   map[string]interface{}               `json:"-" api:"extrafields"`
	JSON          transactionSourceCardDisputeLossJSON `json:"-"`
}

// transactionSourceCardDisputeLossJSON contains the JSON metadata for the struct
// [TransactionSourceCardDisputeLoss]
type transactionSourceCardDisputeLossJSON struct {
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

// A Card Financial object. This field will be present in the JSON response if and
// only if `category` is equal to `card_financial`. Card Financials are temporary
// holds placed on a customer's funds with the intent to later clear a transaction.
type TransactionSourceCardFinancial struct {
	// The Card Financial identifier.
	ID string `json:"id" api:"required"`
	// Whether this financial was approved by Increase, the card network through
	// stand-in processing, or the user through a real-time decision.
	Actioner TransactionSourceCardFinancialActioner `json:"actioner" api:"required"`
	// Additional amounts associated with the card authorization, such as ATM
	// surcharges fees. These are usually a subset of the `amount` field and are used
	// to provide more detailed information about the transaction.
	AdditionalAmounts TransactionSourceCardFinancialAdditionalAmounts `json:"additional_amounts" api:"required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount" api:"required"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency TransactionSourceCardFinancialCurrency `json:"currency" api:"required"`
	// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
	// purchase), the identifier of the token that was used.
	DigitalWalletTokenID string `json:"digital_wallet_token_id" api:"required,nullable"`
	// The direction describes the direction the funds will move, either from the
	// cardholder to the merchant or from the merchant to the cardholder.
	Direction TransactionSourceCardFinancialDirection `json:"direction" api:"required"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id" api:"required"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode string `json:"merchant_category_code" api:"required"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city" api:"required,nullable"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country" api:"required"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor string `json:"merchant_descriptor" api:"required"`
	// The merchant's postal code. For US merchants this is either a 5-digit or 9-digit
	// ZIP code, where the first 5 and last 4 are separated by a dash.
	MerchantPostalCode string `json:"merchant_postal_code" api:"required,nullable"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state" api:"required,nullable"`
	// Fields specific to the `network`.
	NetworkDetails TransactionSourceCardFinancialNetworkDetails `json:"network_details" api:"required"`
	// Network-specific identifiers for a specific request or transaction.
	NetworkIdentifiers TransactionSourceCardFinancialNetworkIdentifiers `json:"network_identifiers" api:"required"`
	// The risk score generated by the card network. For Visa this is the Visa Advanced
	// Authorization risk score, from 0 to 99, where 99 is the riskiest. For Pulse the
	// score is from 0 to 999, where 999 is the riskiest.
	NetworkRiskScore int64 `json:"network_risk_score" api:"required,nullable"`
	// If the authorization was made in-person with a physical card, the Physical Card
	// that was used.
	PhysicalCardID string `json:"physical_card_id" api:"required,nullable"`
	// The pending amount in the minor unit of the transaction's presentment currency.
	PresentmentAmount int64 `json:"presentment_amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's presentment currency.
	PresentmentCurrency string `json:"presentment_currency" api:"required"`
	// The processing category describes the intent behind the financial, such as
	// whether it was used for bill payments or an automatic fuel dispenser.
	ProcessingCategory TransactionSourceCardFinancialProcessingCategory `json:"processing_category" api:"required"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id" api:"required,nullable"`
	// The terminal identifier (commonly abbreviated as TID) of the terminal the card
	// is transacting with.
	TerminalID string `json:"terminal_id" api:"required,nullable"`
	// The identifier of the Transaction associated with this Transaction.
	TransactionID string `json:"transaction_id" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_financial`.
	Type TransactionSourceCardFinancialType `json:"type" api:"required"`
	// Fields related to verification of cardholder-provided values.
	Verification TransactionSourceCardFinancialVerification `json:"verification" api:"required"`
	ExtraFields  map[string]interface{}                     `json:"-" api:"extrafields"`
	JSON         transactionSourceCardFinancialJSON         `json:"-"`
}

// transactionSourceCardFinancialJSON contains the JSON metadata for the struct
// [TransactionSourceCardFinancial]
type transactionSourceCardFinancialJSON struct {
	ID                   apijson.Field
	Actioner             apijson.Field
	AdditionalAmounts    apijson.Field
	Amount               apijson.Field
	CardPaymentID        apijson.Field
	Currency             apijson.Field
	DigitalWalletTokenID apijson.Field
	Direction            apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantDescriptor   apijson.Field
	MerchantPostalCode   apijson.Field
	MerchantState        apijson.Field
	NetworkDetails       apijson.Field
	NetworkIdentifiers   apijson.Field
	NetworkRiskScore     apijson.Field
	PhysicalCardID       apijson.Field
	PresentmentAmount    apijson.Field
	PresentmentCurrency  apijson.Field
	ProcessingCategory   apijson.Field
	RealTimeDecisionID   apijson.Field
	TerminalID           apijson.Field
	TransactionID        apijson.Field
	Type                 apijson.Field
	Verification         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TransactionSourceCardFinancial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardFinancialJSON) RawJSON() string {
	return r.raw
}

// Whether this financial was approved by Increase, the card network through
// stand-in processing, or the user through a real-time decision.
type TransactionSourceCardFinancialActioner string

const (
	TransactionSourceCardFinancialActionerUser     TransactionSourceCardFinancialActioner = "user"
	TransactionSourceCardFinancialActionerIncrease TransactionSourceCardFinancialActioner = "increase"
	TransactionSourceCardFinancialActionerNetwork  TransactionSourceCardFinancialActioner = "network"
)

func (r TransactionSourceCardFinancialActioner) IsKnown() bool {
	switch r {
	case TransactionSourceCardFinancialActionerUser, TransactionSourceCardFinancialActionerIncrease, TransactionSourceCardFinancialActionerNetwork:
		return true
	}
	return false
}

// Additional amounts associated with the card authorization, such as ATM
// surcharges fees. These are usually a subset of the `amount` field and are used
// to provide more detailed information about the transaction.
type TransactionSourceCardFinancialAdditionalAmounts struct {
	// The part of this transaction amount that was for clinic-related services.
	Clinic TransactionSourceCardFinancialAdditionalAmountsClinic `json:"clinic" api:"required,nullable"`
	// The part of this transaction amount that was for dental-related services.
	Dental TransactionSourceCardFinancialAdditionalAmountsDental `json:"dental" api:"required,nullable"`
	// The original pre-authorized amount.
	Original TransactionSourceCardFinancialAdditionalAmountsOriginal `json:"original" api:"required,nullable"`
	// The part of this transaction amount that was for healthcare prescriptions.
	Prescription TransactionSourceCardFinancialAdditionalAmountsPrescription `json:"prescription" api:"required,nullable"`
	// The surcharge amount charged for this transaction by the merchant.
	Surcharge TransactionSourceCardFinancialAdditionalAmountsSurcharge `json:"surcharge" api:"required,nullable"`
	// The total amount of a series of incremental authorizations, optionally provided.
	TotalCumulative TransactionSourceCardFinancialAdditionalAmountsTotalCumulative `json:"total_cumulative" api:"required,nullable"`
	// The total amount of healthcare-related additional amounts.
	TotalHealthcare TransactionSourceCardFinancialAdditionalAmountsTotalHealthcare `json:"total_healthcare" api:"required,nullable"`
	// The part of this transaction amount that was for transit-related services.
	Transit TransactionSourceCardFinancialAdditionalAmountsTransit `json:"transit" api:"required,nullable"`
	// An unknown additional amount.
	Unknown TransactionSourceCardFinancialAdditionalAmountsUnknown `json:"unknown" api:"required,nullable"`
	// The part of this transaction amount that was for vision-related services.
	Vision TransactionSourceCardFinancialAdditionalAmountsVision `json:"vision" api:"required,nullable"`
	JSON   transactionSourceCardFinancialAdditionalAmountsJSON   `json:"-"`
}

// transactionSourceCardFinancialAdditionalAmountsJSON contains the JSON metadata
// for the struct [TransactionSourceCardFinancialAdditionalAmounts]
type transactionSourceCardFinancialAdditionalAmountsJSON struct {
	Clinic          apijson.Field
	Dental          apijson.Field
	Original        apijson.Field
	Prescription    apijson.Field
	Surcharge       apijson.Field
	TotalCumulative apijson.Field
	TotalHealthcare apijson.Field
	Transit         apijson.Field
	Unknown         apijson.Field
	Vision          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TransactionSourceCardFinancialAdditionalAmounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardFinancialAdditionalAmountsJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for clinic-related services.
type TransactionSourceCardFinancialAdditionalAmountsClinic struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                    `json:"currency" api:"required"`
	JSON     transactionSourceCardFinancialAdditionalAmountsClinicJSON `json:"-"`
}

// transactionSourceCardFinancialAdditionalAmountsClinicJSON contains the JSON
// metadata for the struct [TransactionSourceCardFinancialAdditionalAmountsClinic]
type transactionSourceCardFinancialAdditionalAmountsClinicJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardFinancialAdditionalAmountsClinic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardFinancialAdditionalAmountsClinicJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for dental-related services.
type TransactionSourceCardFinancialAdditionalAmountsDental struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                    `json:"currency" api:"required"`
	JSON     transactionSourceCardFinancialAdditionalAmountsDentalJSON `json:"-"`
}

// transactionSourceCardFinancialAdditionalAmountsDentalJSON contains the JSON
// metadata for the struct [TransactionSourceCardFinancialAdditionalAmountsDental]
type transactionSourceCardFinancialAdditionalAmountsDentalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardFinancialAdditionalAmountsDental) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardFinancialAdditionalAmountsDentalJSON) RawJSON() string {
	return r.raw
}

// The original pre-authorized amount.
type TransactionSourceCardFinancialAdditionalAmountsOriginal struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                      `json:"currency" api:"required"`
	JSON     transactionSourceCardFinancialAdditionalAmountsOriginalJSON `json:"-"`
}

// transactionSourceCardFinancialAdditionalAmountsOriginalJSON contains the JSON
// metadata for the struct
// [TransactionSourceCardFinancialAdditionalAmountsOriginal]
type transactionSourceCardFinancialAdditionalAmountsOriginalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardFinancialAdditionalAmountsOriginal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardFinancialAdditionalAmountsOriginalJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for healthcare prescriptions.
type TransactionSourceCardFinancialAdditionalAmountsPrescription struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                          `json:"currency" api:"required"`
	JSON     transactionSourceCardFinancialAdditionalAmountsPrescriptionJSON `json:"-"`
}

// transactionSourceCardFinancialAdditionalAmountsPrescriptionJSON contains the
// JSON metadata for the struct
// [TransactionSourceCardFinancialAdditionalAmountsPrescription]
type transactionSourceCardFinancialAdditionalAmountsPrescriptionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardFinancialAdditionalAmountsPrescription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardFinancialAdditionalAmountsPrescriptionJSON) RawJSON() string {
	return r.raw
}

// The surcharge amount charged for this transaction by the merchant.
type TransactionSourceCardFinancialAdditionalAmountsSurcharge struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                       `json:"currency" api:"required"`
	JSON     transactionSourceCardFinancialAdditionalAmountsSurchargeJSON `json:"-"`
}

// transactionSourceCardFinancialAdditionalAmountsSurchargeJSON contains the JSON
// metadata for the struct
// [TransactionSourceCardFinancialAdditionalAmountsSurcharge]
type transactionSourceCardFinancialAdditionalAmountsSurchargeJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardFinancialAdditionalAmountsSurcharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardFinancialAdditionalAmountsSurchargeJSON) RawJSON() string {
	return r.raw
}

// The total amount of a series of incremental authorizations, optionally provided.
type TransactionSourceCardFinancialAdditionalAmountsTotalCumulative struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                             `json:"currency" api:"required"`
	JSON     transactionSourceCardFinancialAdditionalAmountsTotalCumulativeJSON `json:"-"`
}

// transactionSourceCardFinancialAdditionalAmountsTotalCumulativeJSON contains the
// JSON metadata for the struct
// [TransactionSourceCardFinancialAdditionalAmountsTotalCumulative]
type transactionSourceCardFinancialAdditionalAmountsTotalCumulativeJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardFinancialAdditionalAmountsTotalCumulative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardFinancialAdditionalAmountsTotalCumulativeJSON) RawJSON() string {
	return r.raw
}

// The total amount of healthcare-related additional amounts.
type TransactionSourceCardFinancialAdditionalAmountsTotalHealthcare struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                             `json:"currency" api:"required"`
	JSON     transactionSourceCardFinancialAdditionalAmountsTotalHealthcareJSON `json:"-"`
}

// transactionSourceCardFinancialAdditionalAmountsTotalHealthcareJSON contains the
// JSON metadata for the struct
// [TransactionSourceCardFinancialAdditionalAmountsTotalHealthcare]
type transactionSourceCardFinancialAdditionalAmountsTotalHealthcareJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardFinancialAdditionalAmountsTotalHealthcare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardFinancialAdditionalAmountsTotalHealthcareJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for transit-related services.
type TransactionSourceCardFinancialAdditionalAmountsTransit struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                     `json:"currency" api:"required"`
	JSON     transactionSourceCardFinancialAdditionalAmountsTransitJSON `json:"-"`
}

// transactionSourceCardFinancialAdditionalAmountsTransitJSON contains the JSON
// metadata for the struct [TransactionSourceCardFinancialAdditionalAmountsTransit]
type transactionSourceCardFinancialAdditionalAmountsTransitJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardFinancialAdditionalAmountsTransit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardFinancialAdditionalAmountsTransitJSON) RawJSON() string {
	return r.raw
}

// An unknown additional amount.
type TransactionSourceCardFinancialAdditionalAmountsUnknown struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                     `json:"currency" api:"required"`
	JSON     transactionSourceCardFinancialAdditionalAmountsUnknownJSON `json:"-"`
}

// transactionSourceCardFinancialAdditionalAmountsUnknownJSON contains the JSON
// metadata for the struct [TransactionSourceCardFinancialAdditionalAmountsUnknown]
type transactionSourceCardFinancialAdditionalAmountsUnknownJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardFinancialAdditionalAmountsUnknown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardFinancialAdditionalAmountsUnknownJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for vision-related services.
type TransactionSourceCardFinancialAdditionalAmountsVision struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                    `json:"currency" api:"required"`
	JSON     transactionSourceCardFinancialAdditionalAmountsVisionJSON `json:"-"`
}

// transactionSourceCardFinancialAdditionalAmountsVisionJSON contains the JSON
// metadata for the struct [TransactionSourceCardFinancialAdditionalAmountsVision]
type transactionSourceCardFinancialAdditionalAmountsVisionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardFinancialAdditionalAmountsVision) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardFinancialAdditionalAmountsVisionJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type TransactionSourceCardFinancialCurrency string

const (
	TransactionSourceCardFinancialCurrencyUsd TransactionSourceCardFinancialCurrency = "USD"
)

func (r TransactionSourceCardFinancialCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCardFinancialCurrencyUsd:
		return true
	}
	return false
}

// The direction describes the direction the funds will move, either from the
// cardholder to the merchant or from the merchant to the cardholder.
type TransactionSourceCardFinancialDirection string

const (
	TransactionSourceCardFinancialDirectionSettlement TransactionSourceCardFinancialDirection = "settlement"
	TransactionSourceCardFinancialDirectionRefund     TransactionSourceCardFinancialDirection = "refund"
)

func (r TransactionSourceCardFinancialDirection) IsKnown() bool {
	switch r {
	case TransactionSourceCardFinancialDirectionSettlement, TransactionSourceCardFinancialDirectionRefund:
		return true
	}
	return false
}

// Fields specific to the `network`.
type TransactionSourceCardFinancialNetworkDetails struct {
	// The payment network used to process this card authorization.
	Category TransactionSourceCardFinancialNetworkDetailsCategory `json:"category" api:"required"`
	// Fields specific to the `pulse` network.
	Pulse TransactionSourceCardFinancialNetworkDetailsPulse `json:"pulse" api:"required,nullable"`
	// Fields specific to the `visa` network.
	Visa TransactionSourceCardFinancialNetworkDetailsVisa `json:"visa" api:"required,nullable"`
	JSON transactionSourceCardFinancialNetworkDetailsJSON `json:"-"`
}

// transactionSourceCardFinancialNetworkDetailsJSON contains the JSON metadata for
// the struct [TransactionSourceCardFinancialNetworkDetails]
type transactionSourceCardFinancialNetworkDetailsJSON struct {
	Category    apijson.Field
	Pulse       apijson.Field
	Visa        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardFinancialNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardFinancialNetworkDetailsJSON) RawJSON() string {
	return r.raw
}

// The payment network used to process this card authorization.
type TransactionSourceCardFinancialNetworkDetailsCategory string

const (
	TransactionSourceCardFinancialNetworkDetailsCategoryVisa  TransactionSourceCardFinancialNetworkDetailsCategory = "visa"
	TransactionSourceCardFinancialNetworkDetailsCategoryPulse TransactionSourceCardFinancialNetworkDetailsCategory = "pulse"
)

func (r TransactionSourceCardFinancialNetworkDetailsCategory) IsKnown() bool {
	switch r {
	case TransactionSourceCardFinancialNetworkDetailsCategoryVisa, TransactionSourceCardFinancialNetworkDetailsCategoryPulse:
		return true
	}
	return false
}

// Fields specific to the `pulse` network.
type TransactionSourceCardFinancialNetworkDetailsPulse struct {
	JSON transactionSourceCardFinancialNetworkDetailsPulseJSON `json:"-"`
}

// transactionSourceCardFinancialNetworkDetailsPulseJSON contains the JSON metadata
// for the struct [TransactionSourceCardFinancialNetworkDetailsPulse]
type transactionSourceCardFinancialNetworkDetailsPulseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardFinancialNetworkDetailsPulse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardFinancialNetworkDetailsPulseJSON) RawJSON() string {
	return r.raw
}

// Fields specific to the `visa` network.
type TransactionSourceCardFinancialNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator" api:"required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date.
	PointOfServiceEntryMode TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryMode `json:"point_of_service_entry_mode" api:"required,nullable"`
	// Only present when `actioner: network`. Describes why a card authorization was
	// approved or declined by Visa through stand-in processing.
	StandInProcessingReason TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReason `json:"stand_in_processing_reason" api:"required,nullable"`
	// The capability of the terminal being used to read the card. Shows whether a
	// terminal can e.g., accept chip cards or if it only supports magnetic stripe
	// reads. This reflects the highest capability of the terminal — for example, a
	// terminal that supports both chip and magnetic stripe will be identified as
	// chip-capable.
	TerminalEntryCapability TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapability `json:"terminal_entry_capability" api:"required,nullable"`
	JSON                    transactionSourceCardFinancialNetworkDetailsVisaJSON                    `json:"-"`
}

// transactionSourceCardFinancialNetworkDetailsVisaJSON contains the JSON metadata
// for the struct [TransactionSourceCardFinancialNetworkDetailsVisa]
type transactionSourceCardFinancialNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Field
	PointOfServiceEntryMode     apijson.Field
	StandInProcessingReason     apijson.Field
	TerminalEntryCapability     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *TransactionSourceCardFinancialNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardFinancialNetworkDetailsVisaJSON) RawJSON() string {
	return r.raw
}

// For electronic commerce transactions, this identifies the level of security used
// in obtaining the customer's payment credential. For mail or telephone order
// transactions, identifies the type of mail or telephone order.
type TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicator string

const (
	TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                          TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                               TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                             TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                   TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                     TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                    TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

func (r TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicator) IsKnown() bool {
	switch r {
	case TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder, TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicatorRecurring, TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicatorInstallment, TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder, TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce, TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant, TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction, TransactionSourceCardFinancialNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction:
		return true
	}
	return false
}

// The method used to enter the cardholder's primary account number and card
// expiration date.
type TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryMode string

const (
	TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryModeUnknown                    TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryMode = "unknown"
	TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryModeManual                     TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryMode = "manual"
	TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv        TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryModeOpticalCode                TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryMode = "optical_code"
	TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard      TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card"
	TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryModeContactless                TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryMode = "contactless"
	TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile           TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryMode = "credential_on_file"
	TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe             TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe"
	TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe  TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryMode = "contactless_magnetic_stripe"
	TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)

func (r TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryMode) IsKnown() bool {
	switch r {
	case TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryModeUnknown, TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryModeManual, TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv, TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryModeOpticalCode, TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard, TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryModeContactless, TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile, TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe, TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe, TransactionSourceCardFinancialNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv:
		return true
	}
	return false
}

// Only present when `actioner: network`. Describes why a card authorization was
// approved or declined by Visa through stand-in processing.
type TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReason string

const (
	TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReasonIssuerError                                              TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReason = "issuer_error"
	TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReasonInvalidPhysicalCard                                      TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReason = "invalid_physical_card"
	TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReasonInvalidCardholderAuthenticationVerificationValue         TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReason = "invalid_cardholder_authentication_verification_value"
	TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReasonInternalVisaError                                        TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReason = "internal_visa_error"
	TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReasonMerchantTransactionAdvisoryServiceAuthenticationRequired TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReason = "merchant_transaction_advisory_service_authentication_required"
	TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReasonPaymentFraudDisruptionAcquirerBlock                      TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReason = "payment_fraud_disruption_acquirer_block"
	TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReasonOther                                                    TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReason = "other"
)

func (r TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReason) IsKnown() bool {
	switch r {
	case TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReasonIssuerError, TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReasonInvalidPhysicalCard, TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReasonInvalidCardholderAuthenticationVerificationValue, TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReasonInternalVisaError, TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReasonMerchantTransactionAdvisoryServiceAuthenticationRequired, TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReasonPaymentFraudDisruptionAcquirerBlock, TransactionSourceCardFinancialNetworkDetailsVisaStandInProcessingReasonOther:
		return true
	}
	return false
}

// The capability of the terminal being used to read the card. Shows whether a
// terminal can e.g., accept chip cards or if it only supports magnetic stripe
// reads. This reflects the highest capability of the terminal — for example, a
// terminal that supports both chip and magnetic stripe will be identified as
// chip-capable.
type TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapability string

const (
	TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapabilityUnknown                     TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapability = "unknown"
	TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapabilityTerminalNotUsed             TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapability = "terminal_not_used"
	TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapabilityMagneticStripe              TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapability = "magnetic_stripe"
	TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapabilityBarcode                     TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapability = "barcode"
	TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapabilityOpticalCharacterRecognition TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapability = "optical_character_recognition"
	TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapabilityChipOrContactless           TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapability = "chip_or_contactless"
	TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapabilityContactlessOnly             TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapability = "contactless_only"
	TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapabilityNoCapability                TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapability = "no_capability"
)

func (r TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapability) IsKnown() bool {
	switch r {
	case TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapabilityUnknown, TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapabilityTerminalNotUsed, TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapabilityMagneticStripe, TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapabilityBarcode, TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapabilityOpticalCharacterRecognition, TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapabilityChipOrContactless, TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapabilityContactlessOnly, TransactionSourceCardFinancialNetworkDetailsVisaTerminalEntryCapabilityNoCapability:
		return true
	}
	return false
}

// Network-specific identifiers for a specific request or transaction.
type TransactionSourceCardFinancialNetworkIdentifiers struct {
	// The randomly generated 6-character Authorization Identification Response code
	// sent back to the acquirer in an approved response.
	AuthorizationIdentificationResponse string `json:"authorization_identification_response" api:"required,nullable"`
	// A life-cycle identifier used across e.g., an authorization and a reversal.
	// Expected to be unique per acquirer within a window of time. For some card
	// networks the retrieval reference number includes the trace counter.
	RetrievalReferenceNumber string `json:"retrieval_reference_number" api:"required,nullable"`
	// A counter used to verify an individual authorization. Expected to be unique per
	// acquirer within a window of time.
	TraceNumber string `json:"trace_number" api:"required,nullable"`
	// A globally unique transaction identifier provided by the card network, used
	// across multiple life-cycle requests.
	TransactionID string                                               `json:"transaction_id" api:"required,nullable"`
	JSON          transactionSourceCardFinancialNetworkIdentifiersJSON `json:"-"`
}

// transactionSourceCardFinancialNetworkIdentifiersJSON contains the JSON metadata
// for the struct [TransactionSourceCardFinancialNetworkIdentifiers]
type transactionSourceCardFinancialNetworkIdentifiersJSON struct {
	AuthorizationIdentificationResponse apijson.Field
	RetrievalReferenceNumber            apijson.Field
	TraceNumber                         apijson.Field
	TransactionID                       apijson.Field
	raw                                 string
	ExtraFields                         map[string]apijson.Field
}

func (r *TransactionSourceCardFinancialNetworkIdentifiers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardFinancialNetworkIdentifiersJSON) RawJSON() string {
	return r.raw
}

// The processing category describes the intent behind the financial, such as
// whether it was used for bill payments or an automatic fuel dispenser.
type TransactionSourceCardFinancialProcessingCategory string

const (
	TransactionSourceCardFinancialProcessingCategoryAccountFunding         TransactionSourceCardFinancialProcessingCategory = "account_funding"
	TransactionSourceCardFinancialProcessingCategoryAutomaticFuelDispenser TransactionSourceCardFinancialProcessingCategory = "automatic_fuel_dispenser"
	TransactionSourceCardFinancialProcessingCategoryBillPayment            TransactionSourceCardFinancialProcessingCategory = "bill_payment"
	TransactionSourceCardFinancialProcessingCategoryOriginalCredit         TransactionSourceCardFinancialProcessingCategory = "original_credit"
	TransactionSourceCardFinancialProcessingCategoryPurchase               TransactionSourceCardFinancialProcessingCategory = "purchase"
	TransactionSourceCardFinancialProcessingCategoryQuasiCash              TransactionSourceCardFinancialProcessingCategory = "quasi_cash"
	TransactionSourceCardFinancialProcessingCategoryRefund                 TransactionSourceCardFinancialProcessingCategory = "refund"
	TransactionSourceCardFinancialProcessingCategoryCashDisbursement       TransactionSourceCardFinancialProcessingCategory = "cash_disbursement"
	TransactionSourceCardFinancialProcessingCategoryBalanceInquiry         TransactionSourceCardFinancialProcessingCategory = "balance_inquiry"
	TransactionSourceCardFinancialProcessingCategoryUnknown                TransactionSourceCardFinancialProcessingCategory = "unknown"
)

func (r TransactionSourceCardFinancialProcessingCategory) IsKnown() bool {
	switch r {
	case TransactionSourceCardFinancialProcessingCategoryAccountFunding, TransactionSourceCardFinancialProcessingCategoryAutomaticFuelDispenser, TransactionSourceCardFinancialProcessingCategoryBillPayment, TransactionSourceCardFinancialProcessingCategoryOriginalCredit, TransactionSourceCardFinancialProcessingCategoryPurchase, TransactionSourceCardFinancialProcessingCategoryQuasiCash, TransactionSourceCardFinancialProcessingCategoryRefund, TransactionSourceCardFinancialProcessingCategoryCashDisbursement, TransactionSourceCardFinancialProcessingCategoryBalanceInquiry, TransactionSourceCardFinancialProcessingCategoryUnknown:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `card_financial`.
type TransactionSourceCardFinancialType string

const (
	TransactionSourceCardFinancialTypeCardFinancial TransactionSourceCardFinancialType = "card_financial"
)

func (r TransactionSourceCardFinancialType) IsKnown() bool {
	switch r {
	case TransactionSourceCardFinancialTypeCardFinancial:
		return true
	}
	return false
}

// Fields related to verification of cardholder-provided values.
type TransactionSourceCardFinancialVerification struct {
	// Fields related to verification of the Card Verification Code, a 3-digit code on
	// the back of the card.
	CardVerificationCode TransactionSourceCardFinancialVerificationCardVerificationCode `json:"card_verification_code" api:"required"`
	// Cardholder address provided in the authorization request and the address on file
	// we verified it against.
	CardholderAddress TransactionSourceCardFinancialVerificationCardholderAddress `json:"cardholder_address" api:"required"`
	// Cardholder name provided in the authorization request.
	CardholderName TransactionSourceCardFinancialVerificationCardholderName `json:"cardholder_name" api:"required,nullable"`
	JSON           transactionSourceCardFinancialVerificationJSON           `json:"-"`
}

// transactionSourceCardFinancialVerificationJSON contains the JSON metadata for
// the struct [TransactionSourceCardFinancialVerification]
type transactionSourceCardFinancialVerificationJSON struct {
	CardVerificationCode apijson.Field
	CardholderAddress    apijson.Field
	CardholderName       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TransactionSourceCardFinancialVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardFinancialVerificationJSON) RawJSON() string {
	return r.raw
}

// Fields related to verification of the Card Verification Code, a 3-digit code on
// the back of the card.
type TransactionSourceCardFinancialVerificationCardVerificationCode struct {
	// The result of verifying the Card Verification Code.
	Result TransactionSourceCardFinancialVerificationCardVerificationCodeResult `json:"result" api:"required"`
	JSON   transactionSourceCardFinancialVerificationCardVerificationCodeJSON   `json:"-"`
}

// transactionSourceCardFinancialVerificationCardVerificationCodeJSON contains the
// JSON metadata for the struct
// [TransactionSourceCardFinancialVerificationCardVerificationCode]
type transactionSourceCardFinancialVerificationCardVerificationCodeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceCardFinancialVerificationCardVerificationCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardFinancialVerificationCardVerificationCodeJSON) RawJSON() string {
	return r.raw
}

// The result of verifying the Card Verification Code.
type TransactionSourceCardFinancialVerificationCardVerificationCodeResult string

const (
	TransactionSourceCardFinancialVerificationCardVerificationCodeResultNotChecked TransactionSourceCardFinancialVerificationCardVerificationCodeResult = "not_checked"
	TransactionSourceCardFinancialVerificationCardVerificationCodeResultMatch      TransactionSourceCardFinancialVerificationCardVerificationCodeResult = "match"
	TransactionSourceCardFinancialVerificationCardVerificationCodeResultNoMatch    TransactionSourceCardFinancialVerificationCardVerificationCodeResult = "no_match"
)

func (r TransactionSourceCardFinancialVerificationCardVerificationCodeResult) IsKnown() bool {
	switch r {
	case TransactionSourceCardFinancialVerificationCardVerificationCodeResultNotChecked, TransactionSourceCardFinancialVerificationCardVerificationCodeResultMatch, TransactionSourceCardFinancialVerificationCardVerificationCodeResultNoMatch:
		return true
	}
	return false
}

// Cardholder address provided in the authorization request and the address on file
// we verified it against.
type TransactionSourceCardFinancialVerificationCardholderAddress struct {
	// Line 1 of the address on file for the cardholder.
	ActualLine1 string `json:"actual_line1" api:"required,nullable"`
	// The postal code of the address on file for the cardholder.
	ActualPostalCode string `json:"actual_postal_code" api:"required,nullable"`
	// The cardholder address line 1 provided for verification in the authorization
	// request.
	ProvidedLine1 string `json:"provided_line1" api:"required,nullable"`
	// The postal code provided for verification in the authorization request.
	ProvidedPostalCode string `json:"provided_postal_code" api:"required,nullable"`
	// The address verification result returned to the card network.
	Result TransactionSourceCardFinancialVerificationCardholderAddressResult `json:"result" api:"required"`
	JSON   transactionSourceCardFinancialVerificationCardholderAddressJSON   `json:"-"`
}

// transactionSourceCardFinancialVerificationCardholderAddressJSON contains the
// JSON metadata for the struct
// [TransactionSourceCardFinancialVerificationCardholderAddress]
type transactionSourceCardFinancialVerificationCardholderAddressJSON struct {
	ActualLine1        apijson.Field
	ActualPostalCode   apijson.Field
	ProvidedLine1      apijson.Field
	ProvidedPostalCode apijson.Field
	Result             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TransactionSourceCardFinancialVerificationCardholderAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardFinancialVerificationCardholderAddressJSON) RawJSON() string {
	return r.raw
}

// The address verification result returned to the card network.
type TransactionSourceCardFinancialVerificationCardholderAddressResult string

const (
	TransactionSourceCardFinancialVerificationCardholderAddressResultNotChecked                       TransactionSourceCardFinancialVerificationCardholderAddressResult = "not_checked"
	TransactionSourceCardFinancialVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch    TransactionSourceCardFinancialVerificationCardholderAddressResult = "postal_code_match_address_no_match"
	TransactionSourceCardFinancialVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch    TransactionSourceCardFinancialVerificationCardholderAddressResult = "postal_code_no_match_address_match"
	TransactionSourceCardFinancialVerificationCardholderAddressResultMatch                            TransactionSourceCardFinancialVerificationCardholderAddressResult = "match"
	TransactionSourceCardFinancialVerificationCardholderAddressResultNoMatch                          TransactionSourceCardFinancialVerificationCardholderAddressResult = "no_match"
	TransactionSourceCardFinancialVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked TransactionSourceCardFinancialVerificationCardholderAddressResult = "postal_code_match_address_not_checked"
)

func (r TransactionSourceCardFinancialVerificationCardholderAddressResult) IsKnown() bool {
	switch r {
	case TransactionSourceCardFinancialVerificationCardholderAddressResultNotChecked, TransactionSourceCardFinancialVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch, TransactionSourceCardFinancialVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch, TransactionSourceCardFinancialVerificationCardholderAddressResultMatch, TransactionSourceCardFinancialVerificationCardholderAddressResultNoMatch, TransactionSourceCardFinancialVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked:
		return true
	}
	return false
}

// Cardholder name provided in the authorization request.
type TransactionSourceCardFinancialVerificationCardholderName struct {
	// The first name provided for verification in the authorization request.
	ProvidedFirstName string `json:"provided_first_name" api:"required,nullable"`
	// The last name provided for verification in the authorization request.
	ProvidedLastName string `json:"provided_last_name" api:"required,nullable"`
	// The middle name provided for verification in the authorization request.
	ProvidedMiddleName string                                                       `json:"provided_middle_name" api:"required,nullable"`
	JSON               transactionSourceCardFinancialVerificationCardholderNameJSON `json:"-"`
}

// transactionSourceCardFinancialVerificationCardholderNameJSON contains the JSON
// metadata for the struct
// [TransactionSourceCardFinancialVerificationCardholderName]
type transactionSourceCardFinancialVerificationCardholderNameJSON struct {
	ProvidedFirstName  apijson.Field
	ProvidedLastName   apijson.Field
	ProvidedMiddleName apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TransactionSourceCardFinancialVerificationCardholderName) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardFinancialVerificationCardholderNameJSON) RawJSON() string {
	return r.raw
}

// A Card Push Transfer Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `card_push_transfer_acceptance`.
// A Card Push Transfer Acceptance is created when an Outbound Card Push Transfer
// sent from Increase is accepted by the receiving bank.
type TransactionSourceCardPushTransferAcceptance struct {
	// The transfer amount in USD cents.
	SettlementAmount int64 `json:"settlement_amount" api:"required"`
	// The identifier of the Card Push Transfer that led to this Transaction.
	TransferID  string                                          `json:"transfer_id" api:"required"`
	ExtraFields map[string]interface{}                          `json:"-" api:"extrafields"`
	JSON        transactionSourceCardPushTransferAcceptanceJSON `json:"-"`
}

// transactionSourceCardPushTransferAcceptanceJSON contains the JSON metadata for
// the struct [TransactionSourceCardPushTransferAcceptance]
type transactionSourceCardPushTransferAcceptanceJSON struct {
	SettlementAmount apijson.Field
	TransferID       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransactionSourceCardPushTransferAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardPushTransferAcceptanceJSON) RawJSON() string {
	return r.raw
}

// A Card Refund object. This field will be present in the JSON response if and
// only if `category` is equal to `card_refund`. Card Refunds move money back to
// the cardholder. While they are usually connected to a Card Settlement, an
// acquirer can also refund money directly to a card without relation to a
// transaction.
type TransactionSourceCardRefund struct {
	// The Card Refund identifier.
	ID string `json:"id" api:"required"`
	// The amount in the minor unit of the transaction's settlement currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount" api:"required"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id" api:"required"`
	// Cashback debited for this transaction, if eligible. Cashback is paid out in
	// aggregate, monthly.
	Cashback TransactionSourceCardRefundCashback `json:"cashback" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's settlement currency.
	Currency TransactionSourceCardRefundCurrency `json:"currency" api:"required"`
	// Interchange assessed as a part of this transaction.
	Interchange TransactionSourceCardRefundInterchange `json:"interchange" api:"required,nullable"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id" api:"required"`
	// The 4-digit MCC describing the merchant's business.
	MerchantCategoryCode string `json:"merchant_category_code" api:"required"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city" api:"required"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country" api:"required"`
	// The name of the merchant.
	MerchantName string `json:"merchant_name" api:"required"`
	// The merchant's postal code. For US merchants this is always a 5-digit ZIP code.
	MerchantPostalCode string `json:"merchant_postal_code" api:"required,nullable"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state" api:"required,nullable"`
	// Network-specific identifiers for this refund.
	NetworkIdentifiers TransactionSourceCardRefundNetworkIdentifiers `json:"network_identifiers" api:"required"`
	// The amount in the minor unit of the transaction's presentment currency.
	PresentmentAmount int64 `json:"presentment_amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's presentment currency.
	PresentmentCurrency string `json:"presentment_currency" api:"required"`
	// Additional details about the card purchase, such as tax and industry-specific
	// fields.
	PurchaseDetails TransactionSourceCardRefundPurchaseDetails `json:"purchase_details" api:"required,nullable"`
	// The identifier of the Transaction associated with this Transaction.
	TransactionID string `json:"transaction_id" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type        TransactionSourceCardRefundType `json:"type" api:"required"`
	ExtraFields map[string]interface{}          `json:"-" api:"extrafields"`
	JSON        transactionSourceCardRefundJSON `json:"-"`
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
	Amount string `json:"amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the cashback.
	Currency TransactionSourceCardRefundCashbackCurrency `json:"currency" api:"required"`
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
	TransactionSourceCardRefundCashbackCurrencyUsd TransactionSourceCardRefundCashbackCurrency = "USD"
)

func (r TransactionSourceCardRefundCashbackCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCardRefundCashbackCurrencyUsd:
		return true
	}
	return false
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's settlement currency.
type TransactionSourceCardRefundCurrency string

const (
	TransactionSourceCardRefundCurrencyUsd TransactionSourceCardRefundCurrency = "USD"
)

func (r TransactionSourceCardRefundCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCardRefundCurrencyUsd:
		return true
	}
	return false
}

// Interchange assessed as a part of this transaction.
type TransactionSourceCardRefundInterchange struct {
	// The interchange amount given as a string containing a decimal number in major
	// units (so e.g., "3.14" for $3.14). The amount is a positive number if it is
	// credited to Increase (e.g., settlements) and a negative number if it is debited
	// (e.g., refunds).
	Amount string `json:"amount" api:"required"`
	// The card network specific interchange code.
	Code string `json:"code" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the interchange
	// reimbursement.
	Currency TransactionSourceCardRefundInterchangeCurrency `json:"currency" api:"required"`
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
	TransactionSourceCardRefundInterchangeCurrencyUsd TransactionSourceCardRefundInterchangeCurrency = "USD"
)

func (r TransactionSourceCardRefundInterchangeCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCardRefundInterchangeCurrencyUsd:
		return true
	}
	return false
}

// Network-specific identifiers for this refund.
type TransactionSourceCardRefundNetworkIdentifiers struct {
	// A network assigned business ID that identifies the acquirer that processed this
	// transaction.
	AcquirerBusinessID string `json:"acquirer_business_id" api:"required"`
	// A globally unique identifier for this settlement.
	AcquirerReferenceNumber string `json:"acquirer_reference_number" api:"required"`
	// The randomly generated 6-character Authorization Identification Response code
	// sent back to the acquirer in an approved response.
	AuthorizationIdentificationResponse string `json:"authorization_identification_response" api:"required,nullable"`
	// A globally unique transaction identifier provided by the card network, used
	// across multiple life-cycle requests.
	TransactionID string                                            `json:"transaction_id" api:"required,nullable"`
	JSON          transactionSourceCardRefundNetworkIdentifiersJSON `json:"-"`
}

// transactionSourceCardRefundNetworkIdentifiersJSON contains the JSON metadata for
// the struct [TransactionSourceCardRefundNetworkIdentifiers]
type transactionSourceCardRefundNetworkIdentifiersJSON struct {
	AcquirerBusinessID                  apijson.Field
	AcquirerReferenceNumber             apijson.Field
	AuthorizationIdentificationResponse apijson.Field
	TransactionID                       apijson.Field
	raw                                 string
	ExtraFields                         map[string]apijson.Field
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
	CarRental TransactionSourceCardRefundPurchaseDetailsCarRental `json:"car_rental" api:"required,nullable"`
	// An identifier from the merchant for the customer or consumer.
	CustomerReferenceIdentifier string `json:"customer_reference_identifier" api:"required,nullable"`
	// The state or provincial tax amount in minor units.
	LocalTaxAmount int64 `json:"local_tax_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	LocalTaxCurrency string `json:"local_tax_currency" api:"required,nullable"`
	// Fields specific to lodging.
	Lodging TransactionSourceCardRefundPurchaseDetailsLodging `json:"lodging" api:"required,nullable"`
	// The national tax amount in minor units.
	NationalTaxAmount int64 `json:"national_tax_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	NationalTaxCurrency string `json:"national_tax_currency" api:"required,nullable"`
	// An identifier from the merchant for the purchase to the issuer and cardholder.
	PurchaseIdentifier string `json:"purchase_identifier" api:"required,nullable"`
	// The format of the purchase identifier.
	PurchaseIdentifierFormat TransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat `json:"purchase_identifier_format" api:"required,nullable"`
	// Fields specific to travel.
	Travel TransactionSourceCardRefundPurchaseDetailsTravel `json:"travel" api:"required,nullable"`
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
	CarClassCode string `json:"car_class_code" api:"required,nullable"`
	// Date the customer picked up the car or, in the case of a no-show or pre-pay
	// transaction, the scheduled pick up date.
	CheckoutDate time.Time `json:"checkout_date" api:"required,nullable" format:"date"`
	// Daily rate being charged for the vehicle.
	DailyRentalRateAmount int64 `json:"daily_rental_rate_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the daily rental
	// rate.
	DailyRentalRateCurrency string `json:"daily_rental_rate_currency" api:"required,nullable"`
	// Number of days the vehicle was rented.
	DaysRented int64 `json:"days_rented" api:"required,nullable"`
	// Additional charges (gas, late fee, etc.) being billed.
	ExtraCharges TransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges `json:"extra_charges" api:"required,nullable"`
	// Fuel charges for the vehicle.
	FuelChargesAmount int64 `json:"fuel_charges_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the fuel charges
	// assessed.
	FuelChargesCurrency string `json:"fuel_charges_currency" api:"required,nullable"`
	// Any insurance being charged for the vehicle.
	InsuranceChargesAmount int64 `json:"insurance_charges_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the insurance
	// charges assessed.
	InsuranceChargesCurrency string `json:"insurance_charges_currency" api:"required,nullable"`
	// An indicator that the cardholder is being billed for a reserved vehicle that was
	// not actually rented (that is, a "no-show" charge).
	NoShowIndicator TransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicator `json:"no_show_indicator" api:"required,nullable"`
	// Charges for returning the vehicle at a different location than where it was
	// picked up.
	OneWayDropOffChargesAmount int64 `json:"one_way_drop_off_charges_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the one-way
	// drop-off charges assessed.
	OneWayDropOffChargesCurrency string `json:"one_way_drop_off_charges_currency" api:"required,nullable"`
	// Name of the person renting the vehicle.
	RenterName string `json:"renter_name" api:"required,nullable"`
	// Weekly rate being charged for the vehicle.
	WeeklyRentalRateAmount int64 `json:"weekly_rental_rate_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the weekly
	// rental rate.
	WeeklyRentalRateCurrency string                                                  `json:"weekly_rental_rate_currency" api:"required,nullable"`
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
	CheckInDate time.Time `json:"check_in_date" api:"required,nullable" format:"date"`
	// Daily rate being charged for the room.
	DailyRoomRateAmount int64 `json:"daily_room_rate_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the daily room
	// rate.
	DailyRoomRateCurrency string `json:"daily_room_rate_currency" api:"required,nullable"`
	// Additional charges (phone, late check-out, etc.) being billed.
	ExtraCharges TransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges `json:"extra_charges" api:"required,nullable"`
	// Folio cash advances for the room.
	FolioCashAdvancesAmount int64 `json:"folio_cash_advances_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the folio cash
	// advances.
	FolioCashAdvancesCurrency string `json:"folio_cash_advances_currency" api:"required,nullable"`
	// Food and beverage charges for the room.
	FoodBeverageChargesAmount int64 `json:"food_beverage_charges_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the food and
	// beverage charges.
	FoodBeverageChargesCurrency string `json:"food_beverage_charges_currency" api:"required,nullable"`
	// Indicator that the cardholder is being billed for a reserved room that was not
	// actually used.
	NoShowIndicator TransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicator `json:"no_show_indicator" api:"required,nullable"`
	// Prepaid expenses being charged for the room.
	PrepaidExpensesAmount int64 `json:"prepaid_expenses_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the prepaid
	// expenses.
	PrepaidExpensesCurrency string `json:"prepaid_expenses_currency" api:"required,nullable"`
	// Number of nights the room was rented.
	RoomNights int64 `json:"room_nights" api:"required,nullable"`
	// Total room tax being charged.
	TotalRoomTaxAmount int64 `json:"total_room_tax_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the total room
	// tax.
	TotalRoomTaxCurrency string `json:"total_room_tax_currency" api:"required,nullable"`
	// Total tax being charged for the room.
	TotalTaxAmount int64 `json:"total_tax_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the total tax
	// assessed.
	TotalTaxCurrency string                                                `json:"total_tax_currency" api:"required,nullable"`
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
	Ancillary TransactionSourceCardRefundPurchaseDetailsTravelAncillary `json:"ancillary" api:"required,nullable"`
	// Indicates the computerized reservation system used to book the ticket.
	ComputerizedReservationSystem string `json:"computerized_reservation_system" api:"required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator TransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator `json:"credit_reason_indicator" api:"required,nullable"`
	// Date of departure.
	DepartureDate time.Time `json:"departure_date" api:"required,nullable" format:"date"`
	// Code for the originating city or airport.
	OriginationCityAirportCode string `json:"origination_city_airport_code" api:"required,nullable"`
	// Name of the passenger.
	PassengerName string `json:"passenger_name" api:"required,nullable"`
	// Indicates whether this ticket is non-refundable.
	RestrictedTicketIndicator TransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicator `json:"restricted_ticket_indicator" api:"required,nullable"`
	// Indicates why a ticket was changed.
	TicketChangeIndicator TransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator `json:"ticket_change_indicator" api:"required,nullable"`
	// Ticket number.
	TicketNumber string `json:"ticket_number" api:"required,nullable"`
	// Code for the travel agency if the ticket was issued by a travel agency.
	TravelAgencyCode string `json:"travel_agency_code" api:"required,nullable"`
	// Name of the travel agency if the ticket was issued by a travel agency.
	TravelAgencyName string `json:"travel_agency_name" api:"required,nullable"`
	// Fields specific to each leg of the journey.
	TripLegs []TransactionSourceCardRefundPurchaseDetailsTravelTripLeg `json:"trip_legs" api:"required,nullable"`
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
	ConnectedTicketDocumentNumber string `json:"connected_ticket_document_number" api:"required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator TransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator `json:"credit_reason_indicator" api:"required,nullable"`
	// Name of the passenger or description of the ancillary purchase.
	PassengerNameOrDescription string `json:"passenger_name_or_description" api:"required,nullable"`
	// Additional travel charges, such as baggage fees.
	Services []TransactionSourceCardRefundPurchaseDetailsTravelAncillaryService `json:"services" api:"required"`
	// Ticket document number.
	TicketDocumentNumber string                                                        `json:"ticket_document_number" api:"required,nullable"`
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
	Category TransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory `json:"category" api:"required,nullable"`
	// Sub-category of the ancillary service, free-form.
	SubCategory string                                                               `json:"sub_category" api:"required,nullable"`
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
	CarrierCode string `json:"carrier_code" api:"required,nullable"`
	// Code for the destination city or airport.
	DestinationCityAirportCode string `json:"destination_city_airport_code" api:"required,nullable"`
	// Fare basis code.
	FareBasisCode string `json:"fare_basis_code" api:"required,nullable"`
	// Flight number.
	FlightNumber string `json:"flight_number" api:"required,nullable"`
	// Service class (e.g., first class, business class, etc.).
	ServiceClass string `json:"service_class" api:"required,nullable"`
	// Indicates whether a stopover is allowed on this ticket.
	StopOverCode TransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode `json:"stop_over_code" api:"required,nullable"`
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
	Amount int64 `json:"amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency TransactionSourceCardRevenuePaymentCurrency `json:"currency" api:"required"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end" api:"required" format:"date-time"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time `json:"period_start" api:"required" format:"date-time"`
	// The account the card belonged to.
	TransactedOnAccountID string                                  `json:"transacted_on_account_id" api:"required,nullable"`
	ExtraFields           map[string]interface{}                  `json:"-" api:"extrafields"`
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
	TransactionSourceCardRevenuePaymentCurrencyUsd TransactionSourceCardRevenuePaymentCurrency = "USD"
)

func (r TransactionSourceCardRevenuePaymentCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCardRevenuePaymentCurrencyUsd:
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
	ID string `json:"id" api:"required"`
	// The amount in the minor unit of the transaction's settlement currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount" api:"required"`
	// The Card Authorization that was created prior to this Card Settlement, if one
	// exists.
	CardAuthorization string `json:"card_authorization" api:"required,nullable"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id" api:"required"`
	// Cashback earned on this transaction, if eligible. Cashback is paid out in
	// aggregate, monthly.
	Cashback TransactionSourceCardSettlementCashback `json:"cashback" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's settlement currency.
	Currency TransactionSourceCardSettlementCurrency `json:"currency" api:"required"`
	// Interchange assessed as a part of this transaction.
	Interchange TransactionSourceCardSettlementInterchange `json:"interchange" api:"required,nullable"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id" api:"required"`
	// The 4-digit MCC describing the merchant's business.
	MerchantCategoryCode string `json:"merchant_category_code" api:"required"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city" api:"required"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country" api:"required"`
	// The name of the merchant.
	MerchantName string `json:"merchant_name" api:"required"`
	// The merchant's postal code. For US merchants this is always a 5-digit ZIP code.
	MerchantPostalCode string `json:"merchant_postal_code" api:"required,nullable"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state" api:"required,nullable"`
	// The card network on which this transaction was processed.
	Network TransactionSourceCardSettlementNetwork `json:"network" api:"required"`
	// Network-specific identifiers for this refund.
	NetworkIdentifiers TransactionSourceCardSettlementNetworkIdentifiers `json:"network_identifiers" api:"required"`
	// The identifier of the Pending Transaction associated with this Transaction.
	PendingTransactionID string `json:"pending_transaction_id" api:"required,nullable"`
	// The amount in the minor unit of the transaction's presentment currency.
	PresentmentAmount int64 `json:"presentment_amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's presentment currency.
	PresentmentCurrency string `json:"presentment_currency" api:"required"`
	// Additional details about the card purchase, such as tax and industry-specific
	// fields.
	PurchaseDetails TransactionSourceCardSettlementPurchaseDetails `json:"purchase_details" api:"required,nullable"`
	// Surcharge amount details, if applicable. The amount is positive if the surcharge
	// is added to the overall transaction amount (surcharge), and negative if the
	// surcharge is deducted from the overall transaction amount (discount).
	Surcharge TransactionSourceCardSettlementSurcharge `json:"surcharge" api:"required,nullable"`
	// The identifier of the Transaction associated with this Transaction.
	TransactionID string `json:"transaction_id" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_settlement`.
	Type        TransactionSourceCardSettlementType `json:"type" api:"required"`
	ExtraFields map[string]interface{}              `json:"-" api:"extrafields"`
	JSON        transactionSourceCardSettlementJSON `json:"-"`
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
	Network              apijson.Field
	NetworkIdentifiers   apijson.Field
	PendingTransactionID apijson.Field
	PresentmentAmount    apijson.Field
	PresentmentCurrency  apijson.Field
	PurchaseDetails      apijson.Field
	Surcharge            apijson.Field
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
	Amount string `json:"amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the cashback.
	Currency TransactionSourceCardSettlementCashbackCurrency `json:"currency" api:"required"`
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
	TransactionSourceCardSettlementCashbackCurrencyUsd TransactionSourceCardSettlementCashbackCurrency = "USD"
)

func (r TransactionSourceCardSettlementCashbackCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCardSettlementCashbackCurrencyUsd:
		return true
	}
	return false
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's settlement currency.
type TransactionSourceCardSettlementCurrency string

const (
	TransactionSourceCardSettlementCurrencyUsd TransactionSourceCardSettlementCurrency = "USD"
)

func (r TransactionSourceCardSettlementCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCardSettlementCurrencyUsd:
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
	Amount string `json:"amount" api:"required"`
	// The card network specific interchange code.
	Code string `json:"code" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the interchange
	// reimbursement.
	Currency TransactionSourceCardSettlementInterchangeCurrency `json:"currency" api:"required"`
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
	TransactionSourceCardSettlementInterchangeCurrencyUsd TransactionSourceCardSettlementInterchangeCurrency = "USD"
)

func (r TransactionSourceCardSettlementInterchangeCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCardSettlementInterchangeCurrencyUsd:
		return true
	}
	return false
}

// The card network on which this transaction was processed.
type TransactionSourceCardSettlementNetwork string

const (
	TransactionSourceCardSettlementNetworkVisa  TransactionSourceCardSettlementNetwork = "visa"
	TransactionSourceCardSettlementNetworkPulse TransactionSourceCardSettlementNetwork = "pulse"
)

func (r TransactionSourceCardSettlementNetwork) IsKnown() bool {
	switch r {
	case TransactionSourceCardSettlementNetworkVisa, TransactionSourceCardSettlementNetworkPulse:
		return true
	}
	return false
}

// Network-specific identifiers for this refund.
type TransactionSourceCardSettlementNetworkIdentifiers struct {
	// A network assigned business ID that identifies the acquirer that processed this
	// transaction.
	AcquirerBusinessID string `json:"acquirer_business_id" api:"required"`
	// A globally unique identifier for this settlement.
	AcquirerReferenceNumber string `json:"acquirer_reference_number" api:"required"`
	// The randomly generated 6-character Authorization Identification Response code
	// sent back to the acquirer in an approved response.
	AuthorizationIdentificationResponse string `json:"authorization_identification_response" api:"required,nullable"`
	// A globally unique transaction identifier provided by the card network, used
	// across multiple life-cycle requests.
	TransactionID string                                                `json:"transaction_id" api:"required,nullable"`
	JSON          transactionSourceCardSettlementNetworkIdentifiersJSON `json:"-"`
}

// transactionSourceCardSettlementNetworkIdentifiersJSON contains the JSON metadata
// for the struct [TransactionSourceCardSettlementNetworkIdentifiers]
type transactionSourceCardSettlementNetworkIdentifiersJSON struct {
	AcquirerBusinessID                  apijson.Field
	AcquirerReferenceNumber             apijson.Field
	AuthorizationIdentificationResponse apijson.Field
	TransactionID                       apijson.Field
	raw                                 string
	ExtraFields                         map[string]apijson.Field
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
	CarRental TransactionSourceCardSettlementPurchaseDetailsCarRental `json:"car_rental" api:"required,nullable"`
	// An identifier from the merchant for the customer or consumer.
	CustomerReferenceIdentifier string `json:"customer_reference_identifier" api:"required,nullable"`
	// The state or provincial tax amount in minor units.
	LocalTaxAmount int64 `json:"local_tax_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	LocalTaxCurrency string `json:"local_tax_currency" api:"required,nullable"`
	// Fields specific to lodging.
	Lodging TransactionSourceCardSettlementPurchaseDetailsLodging `json:"lodging" api:"required,nullable"`
	// The national tax amount in minor units.
	NationalTaxAmount int64 `json:"national_tax_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	NationalTaxCurrency string `json:"national_tax_currency" api:"required,nullable"`
	// An identifier from the merchant for the purchase to the issuer and cardholder.
	PurchaseIdentifier string `json:"purchase_identifier" api:"required,nullable"`
	// The format of the purchase identifier.
	PurchaseIdentifierFormat TransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat `json:"purchase_identifier_format" api:"required,nullable"`
	// Fields specific to travel.
	Travel TransactionSourceCardSettlementPurchaseDetailsTravel `json:"travel" api:"required,nullable"`
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
	CarClassCode string `json:"car_class_code" api:"required,nullable"`
	// Date the customer picked up the car or, in the case of a no-show or pre-pay
	// transaction, the scheduled pick up date.
	CheckoutDate time.Time `json:"checkout_date" api:"required,nullable" format:"date"`
	// Daily rate being charged for the vehicle.
	DailyRentalRateAmount int64 `json:"daily_rental_rate_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the daily rental
	// rate.
	DailyRentalRateCurrency string `json:"daily_rental_rate_currency" api:"required,nullable"`
	// Number of days the vehicle was rented.
	DaysRented int64 `json:"days_rented" api:"required,nullable"`
	// Additional charges (gas, late fee, etc.) being billed.
	ExtraCharges TransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges `json:"extra_charges" api:"required,nullable"`
	// Fuel charges for the vehicle.
	FuelChargesAmount int64 `json:"fuel_charges_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the fuel charges
	// assessed.
	FuelChargesCurrency string `json:"fuel_charges_currency" api:"required,nullable"`
	// Any insurance being charged for the vehicle.
	InsuranceChargesAmount int64 `json:"insurance_charges_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the insurance
	// charges assessed.
	InsuranceChargesCurrency string `json:"insurance_charges_currency" api:"required,nullable"`
	// An indicator that the cardholder is being billed for a reserved vehicle that was
	// not actually rented (that is, a "no-show" charge).
	NoShowIndicator TransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicator `json:"no_show_indicator" api:"required,nullable"`
	// Charges for returning the vehicle at a different location than where it was
	// picked up.
	OneWayDropOffChargesAmount int64 `json:"one_way_drop_off_charges_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the one-way
	// drop-off charges assessed.
	OneWayDropOffChargesCurrency string `json:"one_way_drop_off_charges_currency" api:"required,nullable"`
	// Name of the person renting the vehicle.
	RenterName string `json:"renter_name" api:"required,nullable"`
	// Weekly rate being charged for the vehicle.
	WeeklyRentalRateAmount int64 `json:"weekly_rental_rate_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the weekly
	// rental rate.
	WeeklyRentalRateCurrency string                                                      `json:"weekly_rental_rate_currency" api:"required,nullable"`
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
	CheckInDate time.Time `json:"check_in_date" api:"required,nullable" format:"date"`
	// Daily rate being charged for the room.
	DailyRoomRateAmount int64 `json:"daily_room_rate_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the daily room
	// rate.
	DailyRoomRateCurrency string `json:"daily_room_rate_currency" api:"required,nullable"`
	// Additional charges (phone, late check-out, etc.) being billed.
	ExtraCharges TransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges `json:"extra_charges" api:"required,nullable"`
	// Folio cash advances for the room.
	FolioCashAdvancesAmount int64 `json:"folio_cash_advances_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the folio cash
	// advances.
	FolioCashAdvancesCurrency string `json:"folio_cash_advances_currency" api:"required,nullable"`
	// Food and beverage charges for the room.
	FoodBeverageChargesAmount int64 `json:"food_beverage_charges_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the food and
	// beverage charges.
	FoodBeverageChargesCurrency string `json:"food_beverage_charges_currency" api:"required,nullable"`
	// Indicator that the cardholder is being billed for a reserved room that was not
	// actually used.
	NoShowIndicator TransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicator `json:"no_show_indicator" api:"required,nullable"`
	// Prepaid expenses being charged for the room.
	PrepaidExpensesAmount int64 `json:"prepaid_expenses_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the prepaid
	// expenses.
	PrepaidExpensesCurrency string `json:"prepaid_expenses_currency" api:"required,nullable"`
	// Number of nights the room was rented.
	RoomNights int64 `json:"room_nights" api:"required,nullable"`
	// Total room tax being charged.
	TotalRoomTaxAmount int64 `json:"total_room_tax_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the total room
	// tax.
	TotalRoomTaxCurrency string `json:"total_room_tax_currency" api:"required,nullable"`
	// Total tax being charged for the room.
	TotalTaxAmount int64 `json:"total_tax_amount" api:"required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the total tax
	// assessed.
	TotalTaxCurrency string                                                    `json:"total_tax_currency" api:"required,nullable"`
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
	Ancillary TransactionSourceCardSettlementPurchaseDetailsTravelAncillary `json:"ancillary" api:"required,nullable"`
	// Indicates the computerized reservation system used to book the ticket.
	ComputerizedReservationSystem string `json:"computerized_reservation_system" api:"required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator TransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator `json:"credit_reason_indicator" api:"required,nullable"`
	// Date of departure.
	DepartureDate time.Time `json:"departure_date" api:"required,nullable" format:"date"`
	// Code for the originating city or airport.
	OriginationCityAirportCode string `json:"origination_city_airport_code" api:"required,nullable"`
	// Name of the passenger.
	PassengerName string `json:"passenger_name" api:"required,nullable"`
	// Indicates whether this ticket is non-refundable.
	RestrictedTicketIndicator TransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator `json:"restricted_ticket_indicator" api:"required,nullable"`
	// Indicates why a ticket was changed.
	TicketChangeIndicator TransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator `json:"ticket_change_indicator" api:"required,nullable"`
	// Ticket number.
	TicketNumber string `json:"ticket_number" api:"required,nullable"`
	// Code for the travel agency if the ticket was issued by a travel agency.
	TravelAgencyCode string `json:"travel_agency_code" api:"required,nullable"`
	// Name of the travel agency if the ticket was issued by a travel agency.
	TravelAgencyName string `json:"travel_agency_name" api:"required,nullable"`
	// Fields specific to each leg of the journey.
	TripLegs []TransactionSourceCardSettlementPurchaseDetailsTravelTripLeg `json:"trip_legs" api:"required,nullable"`
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
	ConnectedTicketDocumentNumber string `json:"connected_ticket_document_number" api:"required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator `json:"credit_reason_indicator" api:"required,nullable"`
	// Name of the passenger or description of the ancillary purchase.
	PassengerNameOrDescription string `json:"passenger_name_or_description" api:"required,nullable"`
	// Additional travel charges, such as baggage fees.
	Services []TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryService `json:"services" api:"required"`
	// Ticket document number.
	TicketDocumentNumber string                                                            `json:"ticket_document_number" api:"required,nullable"`
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
	Category TransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory `json:"category" api:"required,nullable"`
	// Sub-category of the ancillary service, free-form.
	SubCategory string                                                                   `json:"sub_category" api:"required,nullable"`
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
	CarrierCode string `json:"carrier_code" api:"required,nullable"`
	// Code for the destination city or airport.
	DestinationCityAirportCode string `json:"destination_city_airport_code" api:"required,nullable"`
	// Fare basis code.
	FareBasisCode string `json:"fare_basis_code" api:"required,nullable"`
	// Flight number.
	FlightNumber string `json:"flight_number" api:"required,nullable"`
	// Service class (e.g., first class, business class, etc.).
	ServiceClass string `json:"service_class" api:"required,nullable"`
	// Indicates whether a stopover is allowed on this ticket.
	StopOverCode TransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode `json:"stop_over_code" api:"required,nullable"`
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

// Surcharge amount details, if applicable. The amount is positive if the surcharge
// is added to the overall transaction amount (surcharge), and negative if the
// surcharge is deducted from the overall transaction amount (discount).
type TransactionSourceCardSettlementSurcharge struct {
	// The surcharge amount in the minor unit of the transaction's settlement currency.
	Amount int64 `json:"amount" api:"required"`
	// The surcharge amount in the minor unit of the transaction's presentment
	// currency.
	PresentmentAmount int64                                        `json:"presentment_amount" api:"required"`
	JSON              transactionSourceCardSettlementSurchargeJSON `json:"-"`
}

// transactionSourceCardSettlementSurchargeJSON contains the JSON metadata for the
// struct [TransactionSourceCardSettlementSurcharge]
type transactionSourceCardSettlementSurchargeJSON struct {
	Amount            apijson.Field
	PresentmentAmount apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TransactionSourceCardSettlementSurcharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceCardSettlementSurchargeJSON) RawJSON() string {
	return r.raw
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
	AccruedOnCardID string `json:"accrued_on_card_id" api:"required,nullable"`
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency TransactionSourceCashbackPaymentCurrency `json:"currency" api:"required"`
	// The end of the period for which this transaction paid cashback.
	PeriodEnd time.Time `json:"period_end" api:"required" format:"date-time"`
	// The start of the period for which this transaction paid cashback.
	PeriodStart time.Time                            `json:"period_start" api:"required" format:"date-time"`
	ExtraFields map[string]interface{}               `json:"-" api:"extrafields"`
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
	TransactionSourceCashbackPaymentCurrencyUsd TransactionSourceCashbackPaymentCurrency = "USD"
)

func (r TransactionSourceCashbackPaymentCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCashbackPaymentCurrencyUsd:
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
	// The account number printed on the check. This is an account at the bank that
	// issued the check.
	AccountNumber string `json:"account_number" api:"required"`
	// The amount to be deposited in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount" api:"required"`
	// An additional line of metadata printed on the check. This typically includes the
	// check number for business checks.
	AuxiliaryOnUs string `json:"auxiliary_on_us" api:"required,nullable"`
	// The ID of the Check Deposit that was accepted.
	CheckDepositID string `json:"check_deposit_id" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency TransactionSourceCheckDepositAcceptanceCurrency `json:"currency" api:"required"`
	// The routing number printed on the check. This is a routing number for the bank
	// that issued the check.
	RoutingNumber string `json:"routing_number" api:"required"`
	// The check serial number, if present, for consumer checks. For business checks,
	// the serial number is usually in the `auxiliary_on_us` field.
	SerialNumber string                                      `json:"serial_number" api:"required,nullable"`
	ExtraFields  map[string]interface{}                      `json:"-" api:"extrafields"`
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
	TransactionSourceCheckDepositAcceptanceCurrencyUsd TransactionSourceCheckDepositAcceptanceCurrency = "USD"
)

func (r TransactionSourceCheckDepositAcceptanceCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCheckDepositAcceptanceCurrencyUsd:
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
	Amount int64 `json:"amount" api:"required"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID string `json:"check_deposit_id" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency TransactionSourceCheckDepositReturnCurrency `json:"currency" api:"required"`
	// Why this check was returned by the bank holding the account it was drawn
	// against.
	ReturnReason TransactionSourceCheckDepositReturnReturnReason `json:"return_reason" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt time.Time `json:"returned_at" api:"required" format:"date-time"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID string                                  `json:"transaction_id" api:"required"`
	ExtraFields   map[string]interface{}                  `json:"-" api:"extrafields"`
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
	TransactionSourceCheckDepositReturnCurrencyUsd TransactionSourceCheckDepositReturnCurrency = "USD"
)

func (r TransactionSourceCheckDepositReturnCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceCheckDepositReturnCurrencyUsd:
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
	BackImageFileID string `json:"back_image_file_id" api:"required,nullable"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// bank depositing this check. In some rare cases, this is not transmitted via
	// Check21 and the value will be null.
	BankOfFirstDepositRoutingNumber string `json:"bank_of_first_deposit_routing_number" api:"required,nullable"`
	// When the check was deposited.
	DepositedAt time.Time `json:"deposited_at" api:"required" format:"date-time"`
	// The identifier of the API File object containing an image of the front of the
	// deposited check.
	FrontImageFileID string `json:"front_image_file_id" api:"required,nullable"`
	// The identifier of the Inbound Check Deposit object associated with this
	// transaction.
	InboundCheckDepositID string `json:"inbound_check_deposit_id" api:"required,nullable"`
	// The identifier of the Transaction object created when the check was deposited.
	TransactionID string `json:"transaction_id" api:"required,nullable"`
	// The identifier of the Check Transfer object that was deposited.
	TransferID string `json:"transfer_id" api:"required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_deposit`.
	Type        TransactionSourceCheckTransferDepositType `json:"type" api:"required"`
	ExtraFields map[string]interface{}                    `json:"-" api:"extrafields"`
	JSON        transactionSourceCheckTransferDepositJSON `json:"-"`
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

// A FedNow Transfer Acknowledgement object. This field will be present in the JSON
// response if and only if `category` is equal to
// `fednow_transfer_acknowledgement`. A FedNow Transfer Acknowledgement is created
// when a FedNow Transfer sent from Increase is acknowledged by the receiving bank.
type TransactionSourceFednowTransferAcknowledgement struct {
	// The identifier of the FedNow Transfer that led to this Transaction.
	TransferID  string                                             `json:"transfer_id" api:"required"`
	ExtraFields map[string]interface{}                             `json:"-" api:"extrafields"`
	JSON        transactionSourceFednowTransferAcknowledgementJSON `json:"-"`
}

// transactionSourceFednowTransferAcknowledgementJSON contains the JSON metadata
// for the struct [TransactionSourceFednowTransferAcknowledgement]
type transactionSourceFednowTransferAcknowledgementJSON struct {
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceFednowTransferAcknowledgement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceFednowTransferAcknowledgementJSON) RawJSON() string {
	return r.raw
}

// A Fee Payment object. This field will be present in the JSON response if and
// only if `category` is equal to `fee_payment`. A Fee Payment represents a payment
// made to Increase.
type TransactionSourceFeePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency TransactionSourceFeePaymentCurrency `json:"currency" api:"required"`
	// The start of this payment's fee period, usually the first day of a month.
	FeePeriodStart time.Time `json:"fee_period_start" api:"required" format:"date"`
	// The Program for which this fee was incurred.
	ProgramID   string                          `json:"program_id" api:"required,nullable"`
	ExtraFields map[string]interface{}          `json:"-" api:"extrafields"`
	JSON        transactionSourceFeePaymentJSON `json:"-"`
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
	TransactionSourceFeePaymentCurrencyUsd TransactionSourceFeePaymentCurrency = "USD"
)

func (r TransactionSourceFeePaymentCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceFeePaymentCurrencyUsd:
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
	Addenda TransactionSourceInboundACHTransferAddenda `json:"addenda" api:"required,nullable"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount" api:"required"`
	// The description of the date of the transfer, usually in the format `YYMMDD`.
	OriginatorCompanyDescriptiveDate string `json:"originator_company_descriptive_date" api:"required,nullable"`
	// Data set by the originator.
	OriginatorCompanyDiscretionaryData string `json:"originator_company_discretionary_data" api:"required,nullable"`
	// An informational description of the transfer.
	OriginatorCompanyEntryDescription string `json:"originator_company_entry_description" api:"required"`
	// An identifier for the originating company. This is generally, but not always, a
	// stable identifier across multiple transfers.
	OriginatorCompanyID string `json:"originator_company_id" api:"required"`
	// A name set by the originator to identify themselves.
	OriginatorCompanyName string `json:"originator_company_name" api:"required"`
	// The originator's identifier for the transfer recipient.
	ReceiverIDNumber string `json:"receiver_id_number" api:"required,nullable"`
	// The name of the transfer recipient. This value is informational and not verified
	// by Increase.
	ReceiverName string `json:"receiver_name" api:"required,nullable"`
	// A 15 digit number recorded in the Nacha file and available to both the
	// originating and receiving bank. Along with the amount, date, and originating
	// routing number, this can be used to identify the ACH transfer at either bank.
	// ACH trace numbers are not unique, but are
	// [used to correlate returns](https://increase.com/documentation/ach-returns#ach-returns).
	TraceNumber string `json:"trace_number" api:"required"`
	// The Inbound ACH Transfer's identifier.
	TransferID  string                                  `json:"transfer_id" api:"required"`
	ExtraFields map[string]interface{}                  `json:"-" api:"extrafields"`
	JSON        transactionSourceInboundACHTransferJSON `json:"-"`
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
	Category TransactionSourceInboundACHTransferAddendaCategory `json:"category" api:"required"`
	// Unstructured `payment_related_information` passed through by the originator.
	Freeform TransactionSourceInboundACHTransferAddendaFreeform `json:"freeform" api:"required,nullable"`
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
	Entries []TransactionSourceInboundACHTransferAddendaFreeformEntry `json:"entries" api:"required"`
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
	PaymentRelatedInformation string                                                      `json:"payment_related_information" api:"required"`
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
	InboundACHTransferID string                                                 `json:"inbound_ach_transfer_id" api:"required"`
	ExtraFields          map[string]interface{}                                 `json:"-" api:"extrafields"`
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
	AdjustedTransactionID string `json:"adjusted_transaction_id" api:"required"`
	// The amount of the check adjustment.
	Amount int64 `json:"amount" api:"required"`
	// The reason for the adjustment.
	Reason      TransactionSourceInboundCheckAdjustmentReason `json:"reason" api:"required"`
	ExtraFields map[string]interface{}                        `json:"-" api:"extrafields"`
	JSON        transactionSourceInboundCheckAdjustmentJSON   `json:"-"`
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
	TransactionSourceInboundCheckAdjustmentReasonPaid              TransactionSourceInboundCheckAdjustmentReason = "paid"
)

func (r TransactionSourceInboundCheckAdjustmentReason) IsKnown() bool {
	switch r {
	case TransactionSourceInboundCheckAdjustmentReasonLateReturn, TransactionSourceInboundCheckAdjustmentReasonWrongPayeeCredit, TransactionSourceInboundCheckAdjustmentReasonAdjustedAmount, TransactionSourceInboundCheckAdjustmentReasonNonConformingItem, TransactionSourceInboundCheckAdjustmentReasonPaid:
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
	InboundCheckDepositID string `json:"inbound_check_deposit_id" api:"required"`
	// The identifier of the Check Transfer object that was deposited.
	TransferID  string                                                  `json:"transfer_id" api:"required,nullable"`
	ExtraFields map[string]interface{}                                  `json:"-" api:"extrafields"`
	JSON        transactionSourceInboundCheckDepositReturnIntentionJSON `json:"-"`
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

// An Inbound FedNow Transfer Confirmation object. This field will be present in
// the JSON response if and only if `category` is equal to
// `inbound_fednow_transfer_confirmation`. An Inbound FedNow Transfer Confirmation
// is created when a FedNow transfer is initiated at another bank and received by
// Increase.
type TransactionSourceInboundFednowTransferConfirmation struct {
	// The identifier of the FedNow Transfer that led to this Transaction.
	TransferID  string                                                 `json:"transfer_id" api:"required"`
	ExtraFields map[string]interface{}                                 `json:"-" api:"extrafields"`
	JSON        transactionSourceInboundFednowTransferConfirmationJSON `json:"-"`
}

// transactionSourceInboundFednowTransferConfirmationJSON contains the JSON
// metadata for the struct [TransactionSourceInboundFednowTransferConfirmation]
type transactionSourceInboundFednowTransferConfirmationJSON struct {
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceInboundFednowTransferConfirmation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceInboundFednowTransferConfirmationJSON) RawJSON() string {
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
	Amount int64 `json:"amount" api:"required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real-Time Payments transfer.
	Currency TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency `json:"currency" api:"required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber string `json:"debtor_account_number" api:"required"`
	// The name provided by the sender of the transfer.
	DebtorName string `json:"debtor_name" api:"required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber string `json:"debtor_routing_number" api:"required"`
	// Additional information included with the transfer.
	RemittanceInformation string `json:"remittance_information" api:"required,nullable"`
	// The Real-Time Payments network identification of the transfer.
	TransactionIdentification string `json:"transaction_identification" api:"required"`
	// The identifier of the Real-Time Payments Transfer that led to this Transaction.
	TransferID  string                                                           `json:"transfer_id" api:"required"`
	ExtraFields map[string]interface{}                                           `json:"-" api:"extrafields"`
	JSON        transactionSourceInboundRealTimePaymentsTransferConfirmationJSON `json:"-"`
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
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyUsd TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "USD"
)

func (r TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyUsd:
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
	Amount int64 `json:"amount" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the reversal was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The debtor's routing number.
	DebtorRoutingNumber string `json:"debtor_routing_number" api:"required,nullable"`
	// The description on the reversal message from Fedwire, set by the reversing bank.
	Description string `json:"description" api:"required"`
	// The Fedwire cycle date for the wire reversal. The "Fedwire day" begins at 9:00
	// PM Eastern Time on the evening before the `cycle date`.
	InputCycleDate time.Time `json:"input_cycle_date" api:"required" format:"date"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData string `json:"input_message_accountability_data" api:"required"`
	// The Fedwire sequence number.
	InputSequenceNumber string `json:"input_sequence_number" api:"required"`
	// The Fedwire input source identifier.
	InputSource string `json:"input_source" api:"required"`
	// The sending bank's identifier for the reversal.
	InstructionIdentification string `json:"instruction_identification" api:"required,nullable"`
	// Additional information about the reason for the reversal.
	ReturnReasonAdditionalInformation string `json:"return_reason_additional_information" api:"required,nullable"`
	// A code provided by the sending bank giving a reason for the reversal. The common
	// return reason codes are
	// [documented here](/documentation/wire-reversals#reversal-reason-codes).
	ReturnReasonCode string `json:"return_reason_code" api:"required,nullable"`
	// An Increase-generated description of the `return_reason_code`.
	ReturnReasonCodeDescription string `json:"return_reason_code_description" api:"required,nullable"`
	// The ID for the Transaction associated with the transfer reversal.
	TransactionID string `json:"transaction_id" api:"required"`
	// The ID for the Wire Transfer that is being reversed.
	WireTransferID string                                   `json:"wire_transfer_id" api:"required"`
	ExtraFields    map[string]interface{}                   `json:"-" api:"extrafields"`
	JSON           transactionSourceInboundWireReversalJSON `json:"-"`
}

// transactionSourceInboundWireReversalJSON contains the JSON metadata for the
// struct [TransactionSourceInboundWireReversal]
type transactionSourceInboundWireReversalJSON struct {
	Amount                            apijson.Field
	CreatedAt                         apijson.Field
	DebtorRoutingNumber               apijson.Field
	Description                       apijson.Field
	InputCycleDate                    apijson.Field
	InputMessageAccountabilityData    apijson.Field
	InputSequenceNumber               apijson.Field
	InputSource                       apijson.Field
	InstructionIdentification         apijson.Field
	ReturnReasonAdditionalInformation apijson.Field
	ReturnReasonCode                  apijson.Field
	ReturnReasonCodeDescription       apijson.Field
	TransactionID                     apijson.Field
	WireTransferID                    apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
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
	Amount int64 `json:"amount" api:"required"`
	// A free-form address field set by the sender.
	CreditorAddressLine1 string `json:"creditor_address_line1" api:"required,nullable"`
	// A free-form address field set by the sender.
	CreditorAddressLine2 string `json:"creditor_address_line2" api:"required,nullable"`
	// A free-form address field set by the sender.
	CreditorAddressLine3 string `json:"creditor_address_line3" api:"required,nullable"`
	// A name set by the sender.
	CreditorName string `json:"creditor_name" api:"required,nullable"`
	// A free-form address field set by the sender.
	DebtorAddressLine1 string `json:"debtor_address_line1" api:"required,nullable"`
	// A free-form address field set by the sender.
	DebtorAddressLine2 string `json:"debtor_address_line2" api:"required,nullable"`
	// A free-form address field set by the sender.
	DebtorAddressLine3 string `json:"debtor_address_line3" api:"required,nullable"`
	// A name set by the sender.
	DebtorName string `json:"debtor_name" api:"required,nullable"`
	// An Increase-constructed description of the transfer.
	Description string `json:"description" api:"required"`
	// A free-form reference string set by the sender, to help identify the transfer.
	EndToEndIdentification string `json:"end_to_end_identification" api:"required,nullable"`
	// A unique identifier available to the originating and receiving banks, commonly
	// abbreviated as IMAD. It is created when the wire is submitted to the Fedwire
	// service and is helpful when debugging wires with the originating bank.
	InputMessageAccountabilityData string `json:"input_message_accountability_data" api:"required,nullable"`
	// The American Banking Association (ABA) routing number of the bank that sent the
	// wire.
	InstructingAgentRoutingNumber string `json:"instructing_agent_routing_number" api:"required,nullable"`
	// The sending bank's identifier for the wire transfer.
	InstructionIdentification string `json:"instruction_identification" api:"required,nullable"`
	// The ID of the Inbound Wire Transfer object that resulted in this Transaction.
	TransferID string `json:"transfer_id" api:"required"`
	// The Unique End-to-end Transaction Reference
	// ([UETR](https://www.swift.com/payments/what-unique-end-end-transaction-reference-uetr))
	// of the transfer.
	UniqueEndToEndTransactionReference string `json:"unique_end_to_end_transaction_reference" api:"required,nullable"`
	// A free-form message set by the sender.
	UnstructuredRemittanceInformation string                                   `json:"unstructured_remittance_information" api:"required,nullable"`
	ExtraFields                       map[string]interface{}                   `json:"-" api:"extrafields"`
	JSON                              transactionSourceInboundWireTransferJSON `json:"-"`
}

// transactionSourceInboundWireTransferJSON contains the JSON metadata for the
// struct [TransactionSourceInboundWireTransfer]
type transactionSourceInboundWireTransferJSON struct {
	Amount                             apijson.Field
	CreditorAddressLine1               apijson.Field
	CreditorAddressLine2               apijson.Field
	CreditorAddressLine3               apijson.Field
	CreditorName                       apijson.Field
	DebtorAddressLine1                 apijson.Field
	DebtorAddressLine2                 apijson.Field
	DebtorAddressLine3                 apijson.Field
	DebtorName                         apijson.Field
	Description                        apijson.Field
	EndToEndIdentification             apijson.Field
	InputMessageAccountabilityData     apijson.Field
	InstructingAgentRoutingNumber      apijson.Field
	InstructionIdentification          apijson.Field
	TransferID                         apijson.Field
	UniqueEndToEndTransactionReference apijson.Field
	UnstructuredRemittanceInformation  apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
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
	InboundWireTransferID string                                           `json:"inbound_wire_transfer_id" api:"required"`
	ExtraFields           map[string]interface{}                           `json:"-" api:"extrafields"`
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
	AccruedOnAccountID string `json:"accrued_on_account_id" api:"required"`
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency TransactionSourceInterestPaymentCurrency `json:"currency" api:"required"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end" api:"required" format:"date-time"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time                            `json:"period_start" api:"required" format:"date-time"`
	ExtraFields map[string]interface{}               `json:"-" api:"extrafields"`
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
	TransactionSourceInterestPaymentCurrencyUsd TransactionSourceInterestPaymentCurrency = "USD"
)

func (r TransactionSourceInterestPaymentCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceInterestPaymentCurrencyUsd:
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
	Amount int64 `json:"amount" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency TransactionSourceInternalSourceCurrency `json:"currency" api:"required"`
	// An Internal Source is a transaction between you and Increase. This describes the
	// reason for the transaction.
	Reason      TransactionSourceInternalSourceReason `json:"reason" api:"required"`
	ExtraFields map[string]interface{}                `json:"-" api:"extrafields"`
	JSON        transactionSourceInternalSourceJSON   `json:"-"`
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
	TransactionSourceInternalSourceCurrencyUsd TransactionSourceInternalSourceCurrency = "USD"
)

func (r TransactionSourceInternalSourceCurrency) IsKnown() bool {
	switch r {
	case TransactionSourceInternalSourceCurrencyUsd:
		return true
	}
	return false
}

// An Internal Source is a transaction between you and Increase. This describes the
// reason for the transaction.
type TransactionSourceInternalSourceReason string

const (
	TransactionSourceInternalSourceReasonAccountClosure                    TransactionSourceInternalSourceReason = "account_closure"
	TransactionSourceInternalSourceReasonAccountRevenuePaymentDistribution TransactionSourceInternalSourceReason = "account_revenue_payment_distribution"
	TransactionSourceInternalSourceReasonBankDrawnCheck                    TransactionSourceInternalSourceReason = "bank_drawn_check"
	TransactionSourceInternalSourceReasonBankDrawnCheckCredit              TransactionSourceInternalSourceReason = "bank_drawn_check_credit"
	TransactionSourceInternalSourceReasonBankMigration                     TransactionSourceInternalSourceReason = "bank_migration"
	TransactionSourceInternalSourceReasonCheckAdjustment                   TransactionSourceInternalSourceReason = "check_adjustment"
	TransactionSourceInternalSourceReasonCollectionPayment                 TransactionSourceInternalSourceReason = "collection_payment"
	TransactionSourceInternalSourceReasonCollectionReceivable              TransactionSourceInternalSourceReason = "collection_receivable"
	TransactionSourceInternalSourceReasonDishonoredACHReturn               TransactionSourceInternalSourceReason = "dishonored_ach_return"
	TransactionSourceInternalSourceReasonEmpyrealAdjustment                TransactionSourceInternalSourceReason = "empyreal_adjustment"
	TransactionSourceInternalSourceReasonError                             TransactionSourceInternalSourceReason = "error"
	TransactionSourceInternalSourceReasonErrorCorrection                   TransactionSourceInternalSourceReason = "error_correction"
	TransactionSourceInternalSourceReasonFees                              TransactionSourceInternalSourceReason = "fees"
	TransactionSourceInternalSourceReasonGeneralLedgerTransfer             TransactionSourceInternalSourceReason = "general_ledger_transfer"
	TransactionSourceInternalSourceReasonInterest                          TransactionSourceInternalSourceReason = "interest"
	TransactionSourceInternalSourceReasonNegativeBalanceForgiveness        TransactionSourceInternalSourceReason = "negative_balance_forgiveness"
	TransactionSourceInternalSourceReasonSampleFunds                       TransactionSourceInternalSourceReason = "sample_funds"
	TransactionSourceInternalSourceReasonSampleFundsReturn                 TransactionSourceInternalSourceReason = "sample_funds_return"
)

func (r TransactionSourceInternalSourceReason) IsKnown() bool {
	switch r {
	case TransactionSourceInternalSourceReasonAccountClosure, TransactionSourceInternalSourceReasonAccountRevenuePaymentDistribution, TransactionSourceInternalSourceReasonBankDrawnCheck, TransactionSourceInternalSourceReasonBankDrawnCheckCredit, TransactionSourceInternalSourceReasonBankMigration, TransactionSourceInternalSourceReasonCheckAdjustment, TransactionSourceInternalSourceReasonCollectionPayment, TransactionSourceInternalSourceReasonCollectionReceivable, TransactionSourceInternalSourceReasonDishonoredACHReturn, TransactionSourceInternalSourceReasonEmpyrealAdjustment, TransactionSourceInternalSourceReasonError, TransactionSourceInternalSourceReasonErrorCorrection, TransactionSourceInternalSourceReasonFees, TransactionSourceInternalSourceReasonGeneralLedgerTransfer, TransactionSourceInternalSourceReasonInterest, TransactionSourceInternalSourceReasonNegativeBalanceForgiveness, TransactionSourceInternalSourceReasonSampleFunds, TransactionSourceInternalSourceReasonSampleFundsReturn:
		return true
	}
	return false
}

// If the category of this Transaction source is equal to `other`, this field will
// contain an empty object, otherwise it will contain null.
type TransactionSourceOther struct {
	JSON transactionSourceOtherJSON `json:"-"`
}

// transactionSourceOtherJSON contains the JSON metadata for the struct
// [TransactionSourceOther]
type transactionSourceOtherJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceOther) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceOtherJSON) RawJSON() string {
	return r.raw
}

// A Real-Time Payments Transfer Acknowledgement object. This field will be present
// in the JSON response if and only if `category` is equal to
// `real_time_payments_transfer_acknowledgement`. A Real-Time Payments Transfer
// Acknowledgement is created when a Real-Time Payments Transfer sent from Increase
// is acknowledged by the receiving bank.
type TransactionSourceRealTimePaymentsTransferAcknowledgement struct {
	// The transfer amount in USD cents.
	Amount int64 `json:"amount" api:"required"`
	// The destination account number.
	DestinationAccountNumber string `json:"destination_account_number" api:"required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	DestinationRoutingNumber string `json:"destination_routing_number" api:"required"`
	// Unstructured information that will show on the recipient's bank statement.
	RemittanceInformation string `json:"remittance_information" api:"required"`
	// The identifier of the Real-Time Payments Transfer that led to this Transaction.
	TransferID  string                                                       `json:"transfer_id" api:"required"`
	ExtraFields map[string]interface{}                                       `json:"-" api:"extrafields"`
	JSON        transactionSourceRealTimePaymentsTransferAcknowledgementJSON `json:"-"`
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
	Originator  string                           `json:"originator" api:"required"`
	ExtraFields map[string]interface{}           `json:"-" api:"extrafields"`
	JSON        transactionSourceSampleFundsJSON `json:"-"`
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
	TransferID  string                                      `json:"transfer_id" api:"required"`
	ExtraFields map[string]interface{}                      `json:"-" api:"extrafields"`
	JSON        transactionSourceSwiftTransferIntentionJSON `json:"-"`
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

// A Swift Transfer Return object. This field will be present in the JSON response
// if and only if `category` is equal to `swift_transfer_return`. A Swift Transfer
// Return is created when a Swift Transfer is returned by the receiving bank.
type TransactionSourceSwiftTransferReturn struct {
	// The identifier of the Swift Transfer that led to this Transaction.
	TransferID  string                                   `json:"transfer_id" api:"required"`
	ExtraFields map[string]interface{}                   `json:"-" api:"extrafields"`
	JSON        transactionSourceSwiftTransferReturnJSON `json:"-"`
}

// transactionSourceSwiftTransferReturnJSON contains the JSON metadata for the
// struct [TransactionSourceSwiftTransferReturn]
type transactionSourceSwiftTransferReturnJSON struct {
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSourceSwiftTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSourceSwiftTransferReturnJSON) RawJSON() string {
	return r.raw
}

// A Wire Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_intention`. A Wire
// Transfer initiated via Increase and sent to a different bank.
type TransactionSourceWireTransferIntention struct {
	// The destination account number.
	AccountNumber string `json:"account_number" api:"required"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount" api:"required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient string `json:"message_to_recipient" api:"required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number" api:"required"`
	// The identifier of the Wire Transfer that led to this Transaction.
	TransferID  string                                     `json:"transfer_id" api:"required"`
	ExtraFields map[string]interface{}                     `json:"-" api:"extrafields"`
	JSON        transactionSourceWireTransferIntentionJSON `json:"-"`
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
	TransactionListParamsCategoryInCardFinancial                               TransactionListParamsCategoryIn = "card_financial"
	TransactionListParamsCategoryInCardRevenuePayment                          TransactionListParamsCategoryIn = "card_revenue_payment"
	TransactionListParamsCategoryInCheckDepositAcceptance                      TransactionListParamsCategoryIn = "check_deposit_acceptance"
	TransactionListParamsCategoryInCheckDepositReturn                          TransactionListParamsCategoryIn = "check_deposit_return"
	TransactionListParamsCategoryInFednowTransferAcknowledgement               TransactionListParamsCategoryIn = "fednow_transfer_acknowledgement"
	TransactionListParamsCategoryInCheckTransferDeposit                        TransactionListParamsCategoryIn = "check_transfer_deposit"
	TransactionListParamsCategoryInFeePayment                                  TransactionListParamsCategoryIn = "fee_payment"
	TransactionListParamsCategoryInInboundACHTransfer                          TransactionListParamsCategoryIn = "inbound_ach_transfer"
	TransactionListParamsCategoryInInboundACHTransferReturnIntention           TransactionListParamsCategoryIn = "inbound_ach_transfer_return_intention"
	TransactionListParamsCategoryInInboundCheckDepositReturnIntention          TransactionListParamsCategoryIn = "inbound_check_deposit_return_intention"
	TransactionListParamsCategoryInInboundCheckAdjustment                      TransactionListParamsCategoryIn = "inbound_check_adjustment"
	TransactionListParamsCategoryInInboundFednowTransferConfirmation           TransactionListParamsCategoryIn = "inbound_fednow_transfer_confirmation"
	TransactionListParamsCategoryInInboundRealTimePaymentsTransferConfirmation TransactionListParamsCategoryIn = "inbound_real_time_payments_transfer_confirmation"
	TransactionListParamsCategoryInInboundWireReversal                         TransactionListParamsCategoryIn = "inbound_wire_reversal"
	TransactionListParamsCategoryInInboundWireTransfer                         TransactionListParamsCategoryIn = "inbound_wire_transfer"
	TransactionListParamsCategoryInInboundWireTransferReversal                 TransactionListParamsCategoryIn = "inbound_wire_transfer_reversal"
	TransactionListParamsCategoryInInterestPayment                             TransactionListParamsCategoryIn = "interest_payment"
	TransactionListParamsCategoryInInternalSource                              TransactionListParamsCategoryIn = "internal_source"
	TransactionListParamsCategoryInRealTimePaymentsTransferAcknowledgement     TransactionListParamsCategoryIn = "real_time_payments_transfer_acknowledgement"
	TransactionListParamsCategoryInSampleFunds                                 TransactionListParamsCategoryIn = "sample_funds"
	TransactionListParamsCategoryInWireTransferIntention                       TransactionListParamsCategoryIn = "wire_transfer_intention"
	TransactionListParamsCategoryInSwiftTransferIntention                      TransactionListParamsCategoryIn = "swift_transfer_intention"
	TransactionListParamsCategoryInSwiftTransferReturn                         TransactionListParamsCategoryIn = "swift_transfer_return"
	TransactionListParamsCategoryInCardPushTransferAcceptance                  TransactionListParamsCategoryIn = "card_push_transfer_acceptance"
	TransactionListParamsCategoryInAccountRevenuePayment                       TransactionListParamsCategoryIn = "account_revenue_payment"
	TransactionListParamsCategoryInBlockchainOnrampTransferIntention           TransactionListParamsCategoryIn = "blockchain_onramp_transfer_intention"
	TransactionListParamsCategoryInBlockchainOfframpTransferSettlement         TransactionListParamsCategoryIn = "blockchain_offramp_transfer_settlement"
	TransactionListParamsCategoryInOther                                       TransactionListParamsCategoryIn = "other"
)

func (r TransactionListParamsCategoryIn) IsKnown() bool {
	switch r {
	case TransactionListParamsCategoryInAccountTransferIntention, TransactionListParamsCategoryInACHTransferIntention, TransactionListParamsCategoryInACHTransferRejection, TransactionListParamsCategoryInACHTransferReturn, TransactionListParamsCategoryInCashbackPayment, TransactionListParamsCategoryInCardDisputeAcceptance, TransactionListParamsCategoryInCardDisputeFinancial, TransactionListParamsCategoryInCardDisputeLoss, TransactionListParamsCategoryInCardRefund, TransactionListParamsCategoryInCardSettlement, TransactionListParamsCategoryInCardFinancial, TransactionListParamsCategoryInCardRevenuePayment, TransactionListParamsCategoryInCheckDepositAcceptance, TransactionListParamsCategoryInCheckDepositReturn, TransactionListParamsCategoryInFednowTransferAcknowledgement, TransactionListParamsCategoryInCheckTransferDeposit, TransactionListParamsCategoryInFeePayment, TransactionListParamsCategoryInInboundACHTransfer, TransactionListParamsCategoryInInboundACHTransferReturnIntention, TransactionListParamsCategoryInInboundCheckDepositReturnIntention, TransactionListParamsCategoryInInboundCheckAdjustment, TransactionListParamsCategoryInInboundFednowTransferConfirmation, TransactionListParamsCategoryInInboundRealTimePaymentsTransferConfirmation, TransactionListParamsCategoryInInboundWireReversal, TransactionListParamsCategoryInInboundWireTransfer, TransactionListParamsCategoryInInboundWireTransferReversal, TransactionListParamsCategoryInInterestPayment, TransactionListParamsCategoryInInternalSource, TransactionListParamsCategoryInRealTimePaymentsTransferAcknowledgement, TransactionListParamsCategoryInSampleFunds, TransactionListParamsCategoryInWireTransferIntention, TransactionListParamsCategoryInSwiftTransferIntention, TransactionListParamsCategoryInSwiftTransferReturn, TransactionListParamsCategoryInCardPushTransferAcceptance, TransactionListParamsCategoryInAccountRevenuePayment, TransactionListParamsCategoryInBlockchainOnrampTransferIntention, TransactionListParamsCategoryInBlockchainOfframpTransferSettlement, TransactionListParamsCategoryInOther:
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
