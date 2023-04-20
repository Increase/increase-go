package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
)

type Transaction struct {
	// The identifier for the Account the Transaction belongs to.
	AccountID string `json:"account_id,required"`
	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transcation's
	// Account.
	Currency TransactionCurrency `json:"currency,required"`
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
	RouteType TransactionRouteType `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Transaction. Note that for backwards compatibility reasons, additional
	// undocumented keys may appear in this object. These should be treated as
	// deprecated and will be removed in the future.
	Source TransactionSource `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type TransactionType `json:"type,required"`
	JSON TransactionJSON
}

type TransactionJSON struct {
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

// UnmarshalJSON deserializes the provided bytes into Transaction using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *Transaction) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionCurrency string

const (
	TransactionCurrencyCad TransactionCurrency = "CAD"
	TransactionCurrencyChf TransactionCurrency = "CHF"
	TransactionCurrencyEur TransactionCurrency = "EUR"
	TransactionCurrencyGbp TransactionCurrency = "GBP"
	TransactionCurrencyJpy TransactionCurrency = "JPY"
	TransactionCurrencyUsd TransactionCurrency = "USD"
)

type TransactionRouteType string

const (
	TransactionRouteTypeAccountNumber TransactionRouteType = "account_number"
	TransactionRouteTypeCard          TransactionRouteType = "card"
)

type TransactionSource struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category TransactionSourceCategory `json:"category,required"`
	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention TransactionSourceAccountTransferIntention `json:"account_transfer_intention,required,nullable"`
	// A ACH Check Conversion Return object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_check_conversion_return`.
	ACHCheckConversionReturn TransactionSourceACHCheckConversionReturn `json:"ach_check_conversion_return,required,nullable"`
	// A ACH Check Conversion object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_check_conversion`.
	ACHCheckConversion TransactionSourceACHCheckConversion `json:"ach_check_conversion,required,nullable"`
	// A ACH Transfer Intention object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention TransactionSourceACHTransferIntention `json:"ach_transfer_intention,required,nullable"`
	// A ACH Transfer Rejection object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection TransactionSourceACHTransferRejection `json:"ach_transfer_rejection,required,nullable"`
	// A ACH Transfer Return object. This field will be present in the JSON response if
	// and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn TransactionSourceACHTransferReturn `json:"ach_transfer_return,required,nullable"`
	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance TransactionSourceCardDisputeAcceptance `json:"card_dispute_acceptance,required,nullable"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund TransactionSourceCardRefund `json:"card_refund,required,nullable"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement TransactionSourceCardSettlement `json:"card_settlement,required,nullable"`
	// A Card Revenue Payment object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_revenue_payment`.
	CardRevenuePayment TransactionSourceCardRevenuePayment `json:"card_revenue_payment,required,nullable"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance TransactionSourceCheckDepositAcceptance `json:"check_deposit_acceptance,required,nullable"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn TransactionSourceCheckDepositReturn `json:"check_deposit_return,required,nullable"`
	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention TransactionSourceCheckTransferIntention `json:"check_transfer_intention,required,nullable"`
	// A Check Transfer Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_transfer_return`.
	CheckTransferReturn TransactionSourceCheckTransferReturn `json:"check_transfer_return,required,nullable"`
	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection TransactionSourceCheckTransferRejection `json:"check_transfer_rejection,required,nullable"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest TransactionSourceCheckTransferStopPaymentRequest `json:"check_transfer_stop_payment_request,required,nullable"`
	// A Dispute Resolution object. This field will be present in the JSON response if
	// and only if `category` is equal to `dispute_resolution`.
	DisputeResolution TransactionSourceDisputeResolution `json:"dispute_resolution,required,nullable"`
	// A Empyreal Cash Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `empyreal_cash_deposit`.
	EmpyrealCashDeposit TransactionSourceEmpyrealCashDeposit `json:"empyreal_cash_deposit,required,nullable"`
	// A Fee Payment object. This field will be present in the JSON response if and
	// only if `category` is equal to `fee_payment`.
	FeePayment TransactionSourceFeePayment `json:"fee_payment,required,nullable"`
	// A Inbound ACH Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer TransactionSourceInboundACHTransfer `json:"inbound_ach_transfer,required,nullable"`
	// A Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck TransactionSourceInboundCheck `json:"inbound_check,required,nullable"`
	// A Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer TransactionSourceInboundInternationalACHTransfer `json:"inbound_international_ach_transfer,required,nullable"`
	// A Inbound Real Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation TransactionSourceInboundRealTimePaymentsTransferConfirmation `json:"inbound_real_time_payments_transfer_confirmation,required,nullable"`
	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal TransactionSourceInboundWireDrawdownPaymentReversal `json:"inbound_wire_drawdown_payment_reversal,required,nullable"`
	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment TransactionSourceInboundWireDrawdownPayment `json:"inbound_wire_drawdown_payment,required,nullable"`
	// A Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal TransactionSourceInboundWireReversal `json:"inbound_wire_reversal,required,nullable"`
	// A Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer TransactionSourceInboundWireTransfer `json:"inbound_wire_transfer,required,nullable"`
	// A Interest Payment object. This field will be present in the JSON response if
	// and only if `category` is equal to `interest_payment`.
	InterestPayment TransactionSourceInterestPayment `json:"interest_payment,required,nullable"`
	// A Internal Source object. This field will be present in the JSON response if and
	// only if `category` is equal to `internal_source`.
	InternalSource TransactionSourceInternalSource `json:"internal_source,required,nullable"`
	// A Deprecated Card Refund object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_route_refund`.
	CardRouteRefund TransactionSourceCardRouteRefund `json:"card_route_refund,required,nullable"`
	// A Deprecated Card Settlement object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_settlement`.
	CardRouteSettlement TransactionSourceCardRouteSettlement `json:"card_route_settlement,required,nullable"`
	// A Real Time Payments Transfer Acknowledgement object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `real_time_payments_transfer_acknowledgement`.
	RealTimePaymentsTransferAcknowledgement TransactionSourceRealTimePaymentsTransferAcknowledgement `json:"real_time_payments_transfer_acknowledgement,required,nullable"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds TransactionSourceSampleFunds `json:"sample_funds,required,nullable"`
	// A Wire Drawdown Payment Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_intention`.
	WireDrawdownPaymentIntention TransactionSourceWireDrawdownPaymentIntention `json:"wire_drawdown_payment_intention,required,nullable"`
	// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_rejection`.
	WireDrawdownPaymentRejection TransactionSourceWireDrawdownPaymentRejection `json:"wire_drawdown_payment_rejection,required,nullable"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention TransactionSourceWireTransferIntention `json:"wire_transfer_intention,required,nullable"`
	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection TransactionSourceWireTransferRejection `json:"wire_transfer_rejection,required,nullable"`
	JSON                  TransactionSourceJSON
}

type TransactionSourceJSON struct {
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

// UnmarshalJSON deserializes the provided bytes into TransactionSource using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *TransactionSource) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceCategory string

const (
	TransactionSourceCategoryAccountTransferIntention                    TransactionSourceCategory = "account_transfer_intention"
	TransactionSourceCategoryACHCheckConversionReturn                    TransactionSourceCategory = "ach_check_conversion_return"
	TransactionSourceCategoryACHCheckConversion                          TransactionSourceCategory = "ach_check_conversion"
	TransactionSourceCategoryACHTransferIntention                        TransactionSourceCategory = "ach_transfer_intention"
	TransactionSourceCategoryACHTransferRejection                        TransactionSourceCategory = "ach_transfer_rejection"
	TransactionSourceCategoryACHTransferReturn                           TransactionSourceCategory = "ach_transfer_return"
	TransactionSourceCategoryCardDisputeAcceptance                       TransactionSourceCategory = "card_dispute_acceptance"
	TransactionSourceCategoryCardRefund                                  TransactionSourceCategory = "card_refund"
	TransactionSourceCategoryCardSettlement                              TransactionSourceCategory = "card_settlement"
	TransactionSourceCategoryCardRevenuePayment                          TransactionSourceCategory = "card_revenue_payment"
	TransactionSourceCategoryCheckDepositAcceptance                      TransactionSourceCategory = "check_deposit_acceptance"
	TransactionSourceCategoryCheckDepositReturn                          TransactionSourceCategory = "check_deposit_return"
	TransactionSourceCategoryCheckTransferIntention                      TransactionSourceCategory = "check_transfer_intention"
	TransactionSourceCategoryCheckTransferReturn                         TransactionSourceCategory = "check_transfer_return"
	TransactionSourceCategoryCheckTransferRejection                      TransactionSourceCategory = "check_transfer_rejection"
	TransactionSourceCategoryCheckTransferStopPaymentRequest             TransactionSourceCategory = "check_transfer_stop_payment_request"
	TransactionSourceCategoryDisputeResolution                           TransactionSourceCategory = "dispute_resolution"
	TransactionSourceCategoryEmpyrealCashDeposit                         TransactionSourceCategory = "empyreal_cash_deposit"
	TransactionSourceCategoryFeePayment                                  TransactionSourceCategory = "fee_payment"
	TransactionSourceCategoryInboundACHTransfer                          TransactionSourceCategory = "inbound_ach_transfer"
	TransactionSourceCategoryInboundACHTransferReturnIntention           TransactionSourceCategory = "inbound_ach_transfer_return_intention"
	TransactionSourceCategoryInboundCheck                                TransactionSourceCategory = "inbound_check"
	TransactionSourceCategoryInboundInternationalACHTransfer             TransactionSourceCategory = "inbound_international_ach_transfer"
	TransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation TransactionSourceCategory = "inbound_real_time_payments_transfer_confirmation"
	TransactionSourceCategoryInboundWireDrawdownPaymentReversal          TransactionSourceCategory = "inbound_wire_drawdown_payment_reversal"
	TransactionSourceCategoryInboundWireDrawdownPayment                  TransactionSourceCategory = "inbound_wire_drawdown_payment"
	TransactionSourceCategoryInboundWireReversal                         TransactionSourceCategory = "inbound_wire_reversal"
	TransactionSourceCategoryInboundWireTransfer                         TransactionSourceCategory = "inbound_wire_transfer"
	TransactionSourceCategoryInterestPayment                             TransactionSourceCategory = "interest_payment"
	TransactionSourceCategoryInternalGeneralLedgerTransaction            TransactionSourceCategory = "internal_general_ledger_transaction"
	TransactionSourceCategoryInternalSource                              TransactionSourceCategory = "internal_source"
	TransactionSourceCategoryCardRouteRefund                             TransactionSourceCategory = "card_route_refund"
	TransactionSourceCategoryCardRouteSettlement                         TransactionSourceCategory = "card_route_settlement"
	TransactionSourceCategoryRealTimePaymentsTransferAcknowledgement     TransactionSourceCategory = "real_time_payments_transfer_acknowledgement"
	TransactionSourceCategorySampleFunds                                 TransactionSourceCategory = "sample_funds"
	TransactionSourceCategoryWireDrawdownPaymentIntention                TransactionSourceCategory = "wire_drawdown_payment_intention"
	TransactionSourceCategoryWireDrawdownPaymentRejection                TransactionSourceCategory = "wire_drawdown_payment_rejection"
	TransactionSourceCategoryWireTransferIntention                       TransactionSourceCategory = "wire_transfer_intention"
	TransactionSourceCategoryWireTransferRejection                       TransactionSourceCategory = "wire_transfer_rejection"
	TransactionSourceCategoryOther                                       TransactionSourceCategory = "other"
)

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
	TransferID string `json:"transfer_id,required"`
	JSON       TransactionSourceAccountTransferIntentionJSON
}

type TransactionSourceAccountTransferIntentionJSON struct {
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
// TransactionSourceAccountTransferIntention using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceAccountTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceAccountTransferIntentionCurrency string

const (
	TransactionSourceAccountTransferIntentionCurrencyCad TransactionSourceAccountTransferIntentionCurrency = "CAD"
	TransactionSourceAccountTransferIntentionCurrencyChf TransactionSourceAccountTransferIntentionCurrency = "CHF"
	TransactionSourceAccountTransferIntentionCurrencyEur TransactionSourceAccountTransferIntentionCurrency = "EUR"
	TransactionSourceAccountTransferIntentionCurrencyGbp TransactionSourceAccountTransferIntentionCurrency = "GBP"
	TransactionSourceAccountTransferIntentionCurrencyJpy TransactionSourceAccountTransferIntentionCurrency = "JPY"
	TransactionSourceAccountTransferIntentionCurrencyUsd TransactionSourceAccountTransferIntentionCurrency = "USD"
)

type TransactionSourceACHCheckConversionReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// Why the transfer was returned.
	ReturnReasonCode string `json:"return_reason_code,required"`
	JSON             TransactionSourceACHCheckConversionReturnJSON
}

type TransactionSourceACHCheckConversionReturnJSON struct {
	Amount           pjson.Metadata
	ReturnReasonCode pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceACHCheckConversionReturn using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceACHCheckConversionReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceACHCheckConversion struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the File containing an image of the returned check.
	FileID string `json:"file_id,required"`
	JSON   TransactionSourceACHCheckConversionJSON
}

type TransactionSourceACHCheckConversionJSON struct {
	Amount pjson.Metadata
	FileID pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceACHCheckConversion using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceACHCheckConversion) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceACHTransferIntention struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount              int64  `json:"amount,required"`
	AccountNumber       string `json:"account_number,required"`
	RoutingNumber       string `json:"routing_number,required"`
	StatementDescriptor string `json:"statement_descriptor,required"`
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       TransactionSourceACHTransferIntentionJSON
}

type TransactionSourceACHTransferIntentionJSON struct {
	Amount              pjson.Metadata
	AccountNumber       pjson.Metadata
	RoutingNumber       pjson.Metadata
	StatementDescriptor pjson.Metadata
	TransferID          pjson.Metadata
	Raw                 []byte
	Extras              map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceACHTransferIntention using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceACHTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       TransactionSourceACHTransferRejectionJSON
}

type TransactionSourceACHTransferRejectionJSON struct {
	TransferID pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceACHTransferRejection using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceACHTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode TransactionSourceACHTransferReturnReturnReasonCode `json:"return_reason_code,required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID string `json:"transfer_id,required"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID string `json:"transaction_id,required"`
	JSON          TransactionSourceACHTransferReturnJSON
}

type TransactionSourceACHTransferReturnJSON struct {
	CreatedAt        pjson.Metadata
	ReturnReasonCode pjson.Metadata
	TransferID       pjson.Metadata
	TransactionID    pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceACHTransferReturn using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceACHTransferReturnReturnReasonCode string

const (
	TransactionSourceACHTransferReturnReturnReasonCodeInsufficientFund                                          TransactionSourceACHTransferReturnReturnReasonCode = "insufficient_fund"
	TransactionSourceACHTransferReturnReturnReasonCodeNoAccount                                                 TransactionSourceACHTransferReturnReturnReasonCode = "no_account"
	TransactionSourceACHTransferReturnReturnReasonCodeAccountClosed                                             TransactionSourceACHTransferReturnReturnReasonCode = "account_closed"
	TransactionSourceACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure                             TransactionSourceACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	TransactionSourceACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction              TransactionSourceACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	TransactionSourceACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver                              TransactionSourceACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	TransactionSourceACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode   TransactionSourceACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	TransactionSourceACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized                     TransactionSourceACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	TransactionSourceACHTransferReturnReturnReasonCodePaymentStopped                                            TransactionSourceACHTransferReturnReturnReasonCode = "payment_stopped"
	TransactionSourceACHTransferReturnReturnReasonCodeNonTransactionAccount                                     TransactionSourceACHTransferReturnReturnReasonCode = "non_transaction_account"
	TransactionSourceACHTransferReturnReturnReasonCodeUncollectedFunds                                          TransactionSourceACHTransferReturnReturnReasonCode = "uncollected_funds"
	TransactionSourceACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError                              TransactionSourceACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	TransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete TransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	TransactionSourceACHTransferReturnReturnReasonCodeAmountFieldError                                          TransactionSourceACHTransferReturnReturnReasonCode = "amount_field_error"
	TransactionSourceACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer                            TransactionSourceACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	TransactionSourceACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber                                   TransactionSourceACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	TransactionSourceACHTransferReturnReturnReasonCodeFileRecordEditCriteria                                    TransactionSourceACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	TransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualName                                  TransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	TransactionSourceACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest                                    TransactionSourceACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	TransactionSourceACHTransferReturnReturnReasonCodeAddendaError                                              TransactionSourceACHTransferReturnReturnReasonCode = "addenda_error"
	TransactionSourceACHTransferReturnReturnReasonCodeLimitedParticipationDfi                                   TransactionSourceACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	TransactionSourceACHTransferReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment              TransactionSourceACHTransferReturnReturnReasonCode = "incorrectly_coded_outbound_international_payment"
	TransactionSourceACHTransferReturnReturnReasonCodeOther                                                     TransactionSourceACHTransferReturnReturnReasonCode = "other"
)

type TransactionSourceCardDisputeAcceptance struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was accepted.
	AcceptedAt time.Time `json:"accepted_at,required" format:"date-time"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID string `json:"card_dispute_id,required"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID string `json:"transaction_id,required"`
	JSON          TransactionSourceCardDisputeAcceptanceJSON
}

type TransactionSourceCardDisputeAcceptanceJSON struct {
	AcceptedAt    pjson.Metadata
	CardDisputeID pjson.Metadata
	TransactionID pjson.Metadata
	Raw           []byte
	Extras        map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceCardDisputeAcceptance using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceCardDisputeAcceptance) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceCardRefund struct {
	// The Card Refund identifier.
	ID string `json:"id,required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency TransactionSourceCardRefundCurrency `json:"currency,required"`
	// The identifier for the Transaction this refunds, if any.
	CardSettlementTransactionID string `json:"card_settlement_transaction_id,required,nullable"`
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
	Type TransactionSourceCardRefundType `json:"type,required"`
	JSON TransactionSourceCardRefundJSON
}

type TransactionSourceCardRefundJSON struct {
	ID                          pjson.Metadata
	Amount                      pjson.Metadata
	Currency                    pjson.Metadata
	CardSettlementTransactionID pjson.Metadata
	MerchantCity                pjson.Metadata
	MerchantState               pjson.Metadata
	MerchantCountry             pjson.Metadata
	MerchantName                pjson.Metadata
	MerchantCategoryCode        pjson.Metadata
	Type                        pjson.Metadata
	Raw                         []byte
	Extras                      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into TransactionSourceCardRefund
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *TransactionSourceCardRefund) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceCardRefundCurrency string

const (
	TransactionSourceCardRefundCurrencyCad TransactionSourceCardRefundCurrency = "CAD"
	TransactionSourceCardRefundCurrencyChf TransactionSourceCardRefundCurrency = "CHF"
	TransactionSourceCardRefundCurrencyEur TransactionSourceCardRefundCurrency = "EUR"
	TransactionSourceCardRefundCurrencyGbp TransactionSourceCardRefundCurrency = "GBP"
	TransactionSourceCardRefundCurrencyJpy TransactionSourceCardRefundCurrency = "JPY"
	TransactionSourceCardRefundCurrencyUsd TransactionSourceCardRefundCurrency = "USD"
)

type TransactionSourceCardRefundType string

const (
	TransactionSourceCardRefundTypeCardRefund TransactionSourceCardRefundType = "card_refund"
)

type TransactionSourceCardSettlement struct {
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
	Currency TransactionSourceCardSettlementCurrency `json:"currency,required"`
	// The amount in the minor unit of the transaction's presentment currency.
	PresentmentAmount int64 `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's presentment currency.
	PresentmentCurrency string `json:"presentment_currency,required"`
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
	Type TransactionSourceCardSettlementType `json:"type,required"`
	JSON TransactionSourceCardSettlementJSON
}

type TransactionSourceCardSettlementJSON struct {
	ID                   pjson.Metadata
	CardAuthorization    pjson.Metadata
	Amount               pjson.Metadata
	Currency             pjson.Metadata
	PresentmentAmount    pjson.Metadata
	PresentmentCurrency  pjson.Metadata
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
// TransactionSourceCardSettlement using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *TransactionSourceCardSettlement) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceCardSettlementCurrency string

const (
	TransactionSourceCardSettlementCurrencyCad TransactionSourceCardSettlementCurrency = "CAD"
	TransactionSourceCardSettlementCurrencyChf TransactionSourceCardSettlementCurrency = "CHF"
	TransactionSourceCardSettlementCurrencyEur TransactionSourceCardSettlementCurrency = "EUR"
	TransactionSourceCardSettlementCurrencyGbp TransactionSourceCardSettlementCurrency = "GBP"
	TransactionSourceCardSettlementCurrencyJpy TransactionSourceCardSettlementCurrency = "JPY"
	TransactionSourceCardSettlementCurrencyUsd TransactionSourceCardSettlementCurrency = "USD"
)

type TransactionSourceCardSettlementType string

const (
	TransactionSourceCardSettlementTypeCardSettlement TransactionSourceCardSettlementType = "card_settlement"
)

type TransactionSourceCardRevenuePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency TransactionSourceCardRevenuePaymentCurrency `json:"currency,required"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time `json:"period_start,required" format:"date-time"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The account the card belonged to.
	TransactedOnAccountID string `json:"transacted_on_account_id,required,nullable"`
	JSON                  TransactionSourceCardRevenuePaymentJSON
}

type TransactionSourceCardRevenuePaymentJSON struct {
	Amount                pjson.Metadata
	Currency              pjson.Metadata
	PeriodStart           pjson.Metadata
	PeriodEnd             pjson.Metadata
	TransactedOnAccountID pjson.Metadata
	Raw                   []byte
	Extras                map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceCardRevenuePayment using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceCardRevenuePayment) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceCardRevenuePaymentCurrency string

const (
	TransactionSourceCardRevenuePaymentCurrencyCad TransactionSourceCardRevenuePaymentCurrency = "CAD"
	TransactionSourceCardRevenuePaymentCurrencyChf TransactionSourceCardRevenuePaymentCurrency = "CHF"
	TransactionSourceCardRevenuePaymentCurrencyEur TransactionSourceCardRevenuePaymentCurrency = "EUR"
	TransactionSourceCardRevenuePaymentCurrencyGbp TransactionSourceCardRevenuePaymentCurrency = "GBP"
	TransactionSourceCardRevenuePaymentCurrencyJpy TransactionSourceCardRevenuePaymentCurrency = "JPY"
	TransactionSourceCardRevenuePaymentCurrencyUsd TransactionSourceCardRevenuePaymentCurrency = "USD"
)

type TransactionSourceCheckDepositAcceptance struct {
	// The amount to be deposited in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency TransactionSourceCheckDepositAcceptanceCurrency `json:"currency,required"`
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
	JSON           TransactionSourceCheckDepositAcceptanceJSON
}

type TransactionSourceCheckDepositAcceptanceJSON struct {
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
// TransactionSourceCheckDepositAcceptance using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceCheckDepositAcceptance) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceCheckDepositAcceptanceCurrency string

const (
	TransactionSourceCheckDepositAcceptanceCurrencyCad TransactionSourceCheckDepositAcceptanceCurrency = "CAD"
	TransactionSourceCheckDepositAcceptanceCurrencyChf TransactionSourceCheckDepositAcceptanceCurrency = "CHF"
	TransactionSourceCheckDepositAcceptanceCurrencyEur TransactionSourceCheckDepositAcceptanceCurrency = "EUR"
	TransactionSourceCheckDepositAcceptanceCurrencyGbp TransactionSourceCheckDepositAcceptanceCurrency = "GBP"
	TransactionSourceCheckDepositAcceptanceCurrencyJpy TransactionSourceCheckDepositAcceptanceCurrency = "JPY"
	TransactionSourceCheckDepositAcceptanceCurrencyUsd TransactionSourceCheckDepositAcceptanceCurrency = "USD"
)

type TransactionSourceCheckDepositReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency TransactionSourceCheckDepositReturnCurrency `json:"currency,required"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID string `json:"check_deposit_id,required"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID string                                          `json:"transaction_id,required"`
	ReturnReason  TransactionSourceCheckDepositReturnReturnReason `json:"return_reason,required"`
	JSON          TransactionSourceCheckDepositReturnJSON
}

type TransactionSourceCheckDepositReturnJSON struct {
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
// TransactionSourceCheckDepositReturn using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceCheckDepositReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceCheckDepositReturnCurrency string

const (
	TransactionSourceCheckDepositReturnCurrencyCad TransactionSourceCheckDepositReturnCurrency = "CAD"
	TransactionSourceCheckDepositReturnCurrencyChf TransactionSourceCheckDepositReturnCurrency = "CHF"
	TransactionSourceCheckDepositReturnCurrencyEur TransactionSourceCheckDepositReturnCurrency = "EUR"
	TransactionSourceCheckDepositReturnCurrencyGbp TransactionSourceCheckDepositReturnCurrency = "GBP"
	TransactionSourceCheckDepositReturnCurrencyJpy TransactionSourceCheckDepositReturnCurrency = "JPY"
	TransactionSourceCheckDepositReturnCurrencyUsd TransactionSourceCheckDepositReturnCurrency = "USD"
)

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
)

type TransactionSourceCheckTransferIntention struct {
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
	Currency TransactionSourceCheckTransferIntentionCurrency `json:"currency,required"`
	// The name that will be printed on the check.
	RecipientName string `json:"recipient_name,required"`
	// The identifier of the Check Transfer with which this is associated.
	TransferID string `json:"transfer_id,required"`
	JSON       TransactionSourceCheckTransferIntentionJSON
}

type TransactionSourceCheckTransferIntentionJSON struct {
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
// TransactionSourceCheckTransferIntention using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceCheckTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceCheckTransferIntentionCurrency string

const (
	TransactionSourceCheckTransferIntentionCurrencyCad TransactionSourceCheckTransferIntentionCurrency = "CAD"
	TransactionSourceCheckTransferIntentionCurrencyChf TransactionSourceCheckTransferIntentionCurrency = "CHF"
	TransactionSourceCheckTransferIntentionCurrencyEur TransactionSourceCheckTransferIntentionCurrency = "EUR"
	TransactionSourceCheckTransferIntentionCurrencyGbp TransactionSourceCheckTransferIntentionCurrency = "GBP"
	TransactionSourceCheckTransferIntentionCurrencyJpy TransactionSourceCheckTransferIntentionCurrency = "JPY"
	TransactionSourceCheckTransferIntentionCurrencyUsd TransactionSourceCheckTransferIntentionCurrency = "USD"
)

type TransactionSourceCheckTransferReturn struct {
	// The identifier of the returned Check Transfer.
	TransferID string `json:"transfer_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// If available, a document with additional information about the return.
	FileID string `json:"file_id,required,nullable"`
	// The reason why the check was returned.
	Reason TransactionSourceCheckTransferReturnReason `json:"reason,required"`
	// The identifier of the Transaction that was created to credit you for the
	// returned check.
	TransactionID string `json:"transaction_id,required,nullable"`
	JSON          TransactionSourceCheckTransferReturnJSON
}

type TransactionSourceCheckTransferReturnJSON struct {
	TransferID    pjson.Metadata
	ReturnedAt    pjson.Metadata
	FileID        pjson.Metadata
	Reason        pjson.Metadata
	TransactionID pjson.Metadata
	Raw           []byte
	Extras        map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceCheckTransferReturn using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceCheckTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceCheckTransferReturnReason string

const (
	TransactionSourceCheckTransferReturnReasonMailDeliveryFailure TransactionSourceCheckTransferReturnReason = "mail_delivery_failure"
	TransactionSourceCheckTransferReturnReasonRefusedByRecipient  TransactionSourceCheckTransferReturnReason = "refused_by_recipient"
)

type TransactionSourceCheckTransferRejection struct {
	// The identifier of the Check Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       TransactionSourceCheckTransferRejectionJSON
}

type TransactionSourceCheckTransferRejectionJSON struct {
	TransferID pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceCheckTransferRejection using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceCheckTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceCheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferID string `json:"transfer_id,required"`
	// The transaction ID of the corresponding credit transaction.
	TransactionID string `json:"transaction_id,required"`
	// The time the stop-payment was requested.
	RequestedAt time.Time `json:"requested_at,required" format:"date-time"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type TransactionSourceCheckTransferStopPaymentRequestType `json:"type,required"`
	JSON TransactionSourceCheckTransferStopPaymentRequestJSON
}

type TransactionSourceCheckTransferStopPaymentRequestJSON struct {
	TransferID    pjson.Metadata
	TransactionID pjson.Metadata
	RequestedAt   pjson.Metadata
	Type          pjson.Metadata
	Raw           []byte
	Extras        map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceCheckTransferStopPaymentRequest using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceCheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceCheckTransferStopPaymentRequestType string

const (
	TransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest TransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

type TransactionSourceDisputeResolution struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency TransactionSourceDisputeResolutionCurrency `json:"currency,required"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionID string `json:"disputed_transaction_id,required"`
	JSON                  TransactionSourceDisputeResolutionJSON
}

type TransactionSourceDisputeResolutionJSON struct {
	Amount                pjson.Metadata
	Currency              pjson.Metadata
	DisputedTransactionID pjson.Metadata
	Raw                   []byte
	Extras                map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceDisputeResolution using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceDisputeResolution) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceDisputeResolutionCurrency string

const (
	TransactionSourceDisputeResolutionCurrencyCad TransactionSourceDisputeResolutionCurrency = "CAD"
	TransactionSourceDisputeResolutionCurrencyChf TransactionSourceDisputeResolutionCurrency = "CHF"
	TransactionSourceDisputeResolutionCurrencyEur TransactionSourceDisputeResolutionCurrency = "EUR"
	TransactionSourceDisputeResolutionCurrencyGbp TransactionSourceDisputeResolutionCurrency = "GBP"
	TransactionSourceDisputeResolutionCurrencyJpy TransactionSourceDisputeResolutionCurrency = "JPY"
	TransactionSourceDisputeResolutionCurrencyUsd TransactionSourceDisputeResolutionCurrency = "USD"
)

type TransactionSourceEmpyrealCashDeposit struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount      int64     `json:"amount,required"`
	BagID       string    `json:"bag_id,required"`
	DepositDate time.Time `json:"deposit_date,required" format:"date-time"`
	JSON        TransactionSourceEmpyrealCashDepositJSON
}

type TransactionSourceEmpyrealCashDepositJSON struct {
	Amount      pjson.Metadata
	BagID       pjson.Metadata
	DepositDate pjson.Metadata
	Raw         []byte
	Extras      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceEmpyrealCashDeposit using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceEmpyrealCashDeposit) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceFeePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency TransactionSourceFeePaymentCurrency `json:"currency,required"`
	JSON     TransactionSourceFeePaymentJSON
}

type TransactionSourceFeePaymentJSON struct {
	Amount   pjson.Metadata
	Currency pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into TransactionSourceFeePayment
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *TransactionSourceFeePayment) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceFeePaymentCurrency string

const (
	TransactionSourceFeePaymentCurrencyCad TransactionSourceFeePaymentCurrency = "CAD"
	TransactionSourceFeePaymentCurrencyChf TransactionSourceFeePaymentCurrency = "CHF"
	TransactionSourceFeePaymentCurrencyEur TransactionSourceFeePaymentCurrency = "EUR"
	TransactionSourceFeePaymentCurrencyGbp TransactionSourceFeePaymentCurrency = "GBP"
	TransactionSourceFeePaymentCurrencyJpy TransactionSourceFeePaymentCurrency = "JPY"
	TransactionSourceFeePaymentCurrencyUsd TransactionSourceFeePaymentCurrency = "USD"
)

type TransactionSourceInboundACHTransfer struct {
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
	JSON                               TransactionSourceInboundACHTransferJSON
}

type TransactionSourceInboundACHTransferJSON struct {
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
// TransactionSourceInboundACHTransfer using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceInboundACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceInboundCheck struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency              TransactionSourceInboundCheckCurrency `json:"currency,required"`
	CheckNumber           string                                `json:"check_number,required,nullable"`
	CheckFrontImageFileID string                                `json:"check_front_image_file_id,required,nullable"`
	CheckRearImageFileID  string                                `json:"check_rear_image_file_id,required,nullable"`
	JSON                  TransactionSourceInboundCheckJSON
}

type TransactionSourceInboundCheckJSON struct {
	Amount                pjson.Metadata
	Currency              pjson.Metadata
	CheckNumber           pjson.Metadata
	CheckFrontImageFileID pjson.Metadata
	CheckRearImageFileID  pjson.Metadata
	Raw                   []byte
	Extras                map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into TransactionSourceInboundCheck
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *TransactionSourceInboundCheck) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceInboundCheckCurrency string

const (
	TransactionSourceInboundCheckCurrencyCad TransactionSourceInboundCheckCurrency = "CAD"
	TransactionSourceInboundCheckCurrencyChf TransactionSourceInboundCheckCurrency = "CHF"
	TransactionSourceInboundCheckCurrencyEur TransactionSourceInboundCheckCurrency = "EUR"
	TransactionSourceInboundCheckCurrencyGbp TransactionSourceInboundCheckCurrency = "GBP"
	TransactionSourceInboundCheckCurrencyJpy TransactionSourceInboundCheckCurrency = "JPY"
	TransactionSourceInboundCheckCurrencyUsd TransactionSourceInboundCheckCurrency = "USD"
)

type TransactionSourceInboundInternationalACHTransfer struct {
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
	JSON                                                   TransactionSourceInboundInternationalACHTransferJSON
}

type TransactionSourceInboundInternationalACHTransferJSON struct {
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
// TransactionSourceInboundInternationalACHTransfer using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceInboundInternationalACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real Time Payments transfer.
	Currency TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency `json:"currency,required"`
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
	JSON                  TransactionSourceInboundRealTimePaymentsTransferConfirmationJSON
}

type TransactionSourceInboundRealTimePaymentsTransferConfirmationJSON struct {
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
// TransactionSourceInboundRealTimePaymentsTransferConfirmation using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency string

const (
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyCad TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CAD"
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyChf TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CHF"
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyEur TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "EUR"
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyGbp TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "GBP"
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyJpy TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "JPY"
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyUsd TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "USD"
)

type TransactionSourceInboundWireDrawdownPaymentReversal struct {
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
	JSON                       TransactionSourceInboundWireDrawdownPaymentReversalJSON
}

type TransactionSourceInboundWireDrawdownPaymentReversalJSON struct {
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
// TransactionSourceInboundWireDrawdownPaymentReversal using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceInboundWireDrawdownPayment struct {
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
	JSON                               TransactionSourceInboundWireDrawdownPaymentJSON
}

type TransactionSourceInboundWireDrawdownPaymentJSON struct {
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
// TransactionSourceInboundWireDrawdownPayment using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceInboundWireDrawdownPayment) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceInboundWireReversal struct {
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
	JSON          TransactionSourceInboundWireReversalJSON
}

type TransactionSourceInboundWireReversalJSON struct {
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
// TransactionSourceInboundWireReversal using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceInboundWireReversal) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceInboundWireTransfer struct {
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
	JSON                                    TransactionSourceInboundWireTransferJSON
}

type TransactionSourceInboundWireTransferJSON struct {
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
// TransactionSourceInboundWireTransfer using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceInboundWireTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceInterestPayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency TransactionSourceInterestPaymentCurrency `json:"currency,required"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time `json:"period_start,required" format:"date-time"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The account on which the interest was accrued.
	AccruedOnAccountID string `json:"accrued_on_account_id,required,nullable"`
	JSON               TransactionSourceInterestPaymentJSON
}

type TransactionSourceInterestPaymentJSON struct {
	Amount             pjson.Metadata
	Currency           pjson.Metadata
	PeriodStart        pjson.Metadata
	PeriodEnd          pjson.Metadata
	AccruedOnAccountID pjson.Metadata
	Raw                []byte
	Extras             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceInterestPayment using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *TransactionSourceInterestPayment) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceInterestPaymentCurrency string

const (
	TransactionSourceInterestPaymentCurrencyCad TransactionSourceInterestPaymentCurrency = "CAD"
	TransactionSourceInterestPaymentCurrencyChf TransactionSourceInterestPaymentCurrency = "CHF"
	TransactionSourceInterestPaymentCurrencyEur TransactionSourceInterestPaymentCurrency = "EUR"
	TransactionSourceInterestPaymentCurrencyGbp TransactionSourceInterestPaymentCurrency = "GBP"
	TransactionSourceInterestPaymentCurrencyJpy TransactionSourceInterestPaymentCurrency = "JPY"
	TransactionSourceInterestPaymentCurrencyUsd TransactionSourceInterestPaymentCurrency = "USD"
)

type TransactionSourceInternalSource struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency TransactionSourceInternalSourceCurrency `json:"currency,required"`
	Reason   TransactionSourceInternalSourceReason   `json:"reason,required"`
	JSON     TransactionSourceInternalSourceJSON
}

type TransactionSourceInternalSourceJSON struct {
	Amount   pjson.Metadata
	Currency pjson.Metadata
	Reason   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceInternalSource using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *TransactionSourceInternalSource) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceInternalSourceCurrency string

const (
	TransactionSourceInternalSourceCurrencyCad TransactionSourceInternalSourceCurrency = "CAD"
	TransactionSourceInternalSourceCurrencyChf TransactionSourceInternalSourceCurrency = "CHF"
	TransactionSourceInternalSourceCurrencyEur TransactionSourceInternalSourceCurrency = "EUR"
	TransactionSourceInternalSourceCurrencyGbp TransactionSourceInternalSourceCurrency = "GBP"
	TransactionSourceInternalSourceCurrencyJpy TransactionSourceInternalSourceCurrency = "JPY"
	TransactionSourceInternalSourceCurrencyUsd TransactionSourceInternalSourceCurrency = "USD"
)

type TransactionSourceInternalSourceReason string

const (
	TransactionSourceInternalSourceReasonBankMigration      TransactionSourceInternalSourceReason = "bank_migration"
	TransactionSourceInternalSourceReasonCashback           TransactionSourceInternalSourceReason = "cashback"
	TransactionSourceInternalSourceReasonEmpyrealAdjustment TransactionSourceInternalSourceReason = "empyreal_adjustment"
	TransactionSourceInternalSourceReasonError              TransactionSourceInternalSourceReason = "error"
	TransactionSourceInternalSourceReasonErrorCorrection    TransactionSourceInternalSourceReason = "error_correction"
	TransactionSourceInternalSourceReasonFees               TransactionSourceInternalSourceReason = "fees"
	TransactionSourceInternalSourceReasonInterest           TransactionSourceInternalSourceReason = "interest"
	TransactionSourceInternalSourceReasonSampleFunds        TransactionSourceInternalSourceReason = "sample_funds"
	TransactionSourceInternalSourceReasonSampleFundsReturn  TransactionSourceInternalSourceReason = "sample_funds_return"
)

type TransactionSourceCardRouteRefund struct {
	// The refunded amount in the minor unit of the refunded currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
	// currency.
	Currency             TransactionSourceCardRouteRefundCurrency `json:"currency,required"`
	MerchantAcceptorID   string                                   `json:"merchant_acceptor_id,required"`
	MerchantCity         string                                   `json:"merchant_city,required,nullable"`
	MerchantCountry      string                                   `json:"merchant_country,required"`
	MerchantDescriptor   string                                   `json:"merchant_descriptor,required"`
	MerchantState        string                                   `json:"merchant_state,required,nullable"`
	MerchantCategoryCode string                                   `json:"merchant_category_code,required,nullable"`
	JSON                 TransactionSourceCardRouteRefundJSON
}

type TransactionSourceCardRouteRefundJSON struct {
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
// TransactionSourceCardRouteRefund using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *TransactionSourceCardRouteRefund) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceCardRouteRefundCurrency string

const (
	TransactionSourceCardRouteRefundCurrencyCad TransactionSourceCardRouteRefundCurrency = "CAD"
	TransactionSourceCardRouteRefundCurrencyChf TransactionSourceCardRouteRefundCurrency = "CHF"
	TransactionSourceCardRouteRefundCurrencyEur TransactionSourceCardRouteRefundCurrency = "EUR"
	TransactionSourceCardRouteRefundCurrencyGbp TransactionSourceCardRouteRefundCurrency = "GBP"
	TransactionSourceCardRouteRefundCurrencyJpy TransactionSourceCardRouteRefundCurrency = "JPY"
	TransactionSourceCardRouteRefundCurrencyUsd TransactionSourceCardRouteRefundCurrency = "USD"
)

type TransactionSourceCardRouteSettlement struct {
	// The settled amount in the minor unit of the settlement currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
	// currency.
	Currency             TransactionSourceCardRouteSettlementCurrency `json:"currency,required"`
	MerchantAcceptorID   string                                       `json:"merchant_acceptor_id,required"`
	MerchantCity         string                                       `json:"merchant_city,required,nullable"`
	MerchantCountry      string                                       `json:"merchant_country,required,nullable"`
	MerchantDescriptor   string                                       `json:"merchant_descriptor,required"`
	MerchantState        string                                       `json:"merchant_state,required,nullable"`
	MerchantCategoryCode string                                       `json:"merchant_category_code,required,nullable"`
	JSON                 TransactionSourceCardRouteSettlementJSON
}

type TransactionSourceCardRouteSettlementJSON struct {
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
// TransactionSourceCardRouteSettlement using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceCardRouteSettlement) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceCardRouteSettlementCurrency string

const (
	TransactionSourceCardRouteSettlementCurrencyCad TransactionSourceCardRouteSettlementCurrency = "CAD"
	TransactionSourceCardRouteSettlementCurrencyChf TransactionSourceCardRouteSettlementCurrency = "CHF"
	TransactionSourceCardRouteSettlementCurrencyEur TransactionSourceCardRouteSettlementCurrency = "EUR"
	TransactionSourceCardRouteSettlementCurrencyGbp TransactionSourceCardRouteSettlementCurrency = "GBP"
	TransactionSourceCardRouteSettlementCurrencyJpy TransactionSourceCardRouteSettlementCurrency = "JPY"
	TransactionSourceCardRouteSettlementCurrencyUsd TransactionSourceCardRouteSettlementCurrency = "USD"
)

type TransactionSourceRealTimePaymentsTransferAcknowledgement struct {
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
	JSON       TransactionSourceRealTimePaymentsTransferAcknowledgementJSON
}

type TransactionSourceRealTimePaymentsTransferAcknowledgementJSON struct {
	Amount                   pjson.Metadata
	DestinationAccountNumber pjson.Metadata
	DestinationRoutingNumber pjson.Metadata
	RemittanceInformation    pjson.Metadata
	TransferID               pjson.Metadata
	Raw                      []byte
	Extras                   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceRealTimePaymentsTransferAcknowledgement using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceRealTimePaymentsTransferAcknowledgement) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceSampleFunds struct {
	// Where the sample funds came from.
	Originator string `json:"originator,required"`
	JSON       TransactionSourceSampleFundsJSON
}

type TransactionSourceSampleFundsJSON struct {
	Originator pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into TransactionSourceSampleFunds
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *TransactionSourceSampleFunds) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceWireDrawdownPaymentIntention struct {
	// The transfer amount in USD cents.
	Amount             int64  `json:"amount,required"`
	AccountNumber      string `json:"account_number,required"`
	RoutingNumber      string `json:"routing_number,required"`
	MessageToRecipient string `json:"message_to_recipient,required"`
	TransferID         string `json:"transfer_id,required"`
	JSON               TransactionSourceWireDrawdownPaymentIntentionJSON
}

type TransactionSourceWireDrawdownPaymentIntentionJSON struct {
	Amount             pjson.Metadata
	AccountNumber      pjson.Metadata
	RoutingNumber      pjson.Metadata
	MessageToRecipient pjson.Metadata
	TransferID         pjson.Metadata
	Raw                []byte
	Extras             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceWireDrawdownPaymentIntention using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceWireDrawdownPaymentIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceWireDrawdownPaymentRejection struct {
	TransferID string `json:"transfer_id,required"`
	JSON       TransactionSourceWireDrawdownPaymentRejectionJSON
}

type TransactionSourceWireDrawdownPaymentRejectionJSON struct {
	TransferID pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceWireDrawdownPaymentRejection using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceWireDrawdownPaymentRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceWireTransferIntention struct {
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient string `json:"message_to_recipient,required"`
	TransferID         string `json:"transfer_id,required"`
	JSON               TransactionSourceWireTransferIntentionJSON
}

type TransactionSourceWireTransferIntentionJSON struct {
	Amount             pjson.Metadata
	AccountNumber      pjson.Metadata
	RoutingNumber      pjson.Metadata
	MessageToRecipient pjson.Metadata
	TransferID         pjson.Metadata
	Raw                []byte
	Extras             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceWireTransferIntention using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceWireTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionSourceWireTransferRejection struct {
	TransferID string `json:"transfer_id,required"`
	JSON       TransactionSourceWireTransferRejectionJSON
}

type TransactionSourceWireTransferRejectionJSON struct {
	TransferID pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TransactionSourceWireTransferRejection using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *TransactionSourceWireTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionType string

const (
	TransactionTypeTransaction TransactionType = "transaction"
)

type TransactionListResponse struct {
	// The contents of the list.
	Data []Transaction `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       TransactionListResponseJSON
}

type TransactionListResponseJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into TransactionListResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *TransactionListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
